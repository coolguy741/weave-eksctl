package delete

import (
	"context"
	"fmt"
	"time"

	"github.com/kris-nova/logger"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/weaveworks/eksctl/pkg/eks"
	"github.com/weaveworks/eksctl/pkg/fargate"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/cfn/manager"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
	"github.com/weaveworks/eksctl/pkg/elb"
	iamoidc "github.com/weaveworks/eksctl/pkg/iam/oidc"
	"github.com/weaveworks/eksctl/pkg/kubernetes"
	"github.com/weaveworks/eksctl/pkg/printers"
	ssh "github.com/weaveworks/eksctl/pkg/ssh/client"
	"github.com/weaveworks/eksctl/pkg/utils/kubeconfig"
	"github.com/weaveworks/eksctl/pkg/vpc"
)

func deleteClusterCmd(cmd *cmdutils.Cmd) {
	cfg := api.NewClusterConfig()
	cmd.ClusterConfig = cfg

	cmd.SetDescription("cluster", "Delete a cluster", "")

	cmd.SetRunFuncWithNameArg(func() error {
		return doDeleteCluster(cmd)
	})

	cmd.FlagSetGroup.InFlagSet("General", func(fs *pflag.FlagSet) {
		fs.StringVarP(&cfg.Metadata.Name, "name", "n", "", "EKS cluster name")
		cmdutils.AddRegionFlag(fs, cmd.ProviderConfig)

		cmd.Wait = false
		cmdutils.AddWaitFlag(fs, &cmd.Wait, "deletion of all resources")

		cmdutils.AddConfigFileFlag(fs, &cmd.ClusterConfigFile)
		cmdutils.AddTimeoutFlag(fs, &cmd.ProviderConfig.WaitTimeout)
	})

	cmdutils.AddCommonFlagsForAWS(cmd.FlagSetGroup, cmd.ProviderConfig, true)
}

func handleErrors(errs []error, subject string) error {
	logger.Info("%d error(s) occurred while deleting %s", len(errs), subject)
	for _, err := range errs {
		logger.Critical("%s\n", err.Error())
	}
	return fmt.Errorf("failed to delete %s", subject)
}

func deleteDeprecatedStacks(stackManager *manager.StackCollection) (bool, error) {
	tasks, err := stackManager.DeleteTasksForDeprecatedStacks()
	if err != nil {
		return true, err
	}
	if count := tasks.Len(); count > 0 {
		logger.Info(tasks.Describe())
		if errs := tasks.DoAllSync(); len(errs) > 0 {
			return true, handleErrors(errs, "deprecated stacks")
		}
		logger.Success("deleted all %s deperecated stacks", count)
		return true, nil
	}
	return false, nil
}

func doDeleteCluster(cmd *cmdutils.Cmd) error {
	if err := cmdutils.NewMetadataLoader(cmd).Load(); err != nil {
		return err
	}

	cfg := cmd.ClusterConfig
	meta := cmd.ClusterConfig.Metadata

	printer := printers.NewJSONPrinter()

	ctl, err := cmd.NewCtl()
	if err != nil {
		return err
	}
	cmdutils.LogRegionAndVersionInfo(meta)

	if err := ctl.CheckAuth(); err != nil {
		return err
	}

	logger.Info("deleting EKS cluster %q", meta.Name)
	if err := printer.LogObj(logger.Debug, "cfg.json = \\\n%s\n", cfg); err != nil {
		return err
	}

	if ok, err := ctl.CanDelete(cfg); !ok {
		return err
	}

	var (
		clientSet kubernetes.Interface
		oidc      *iamoidc.OpenIDConnectManager
	)

	clusterOperable, _ := ctl.CanOperate(cfg)
	oidcSupported := true
	if clusterOperable {
		clientSet, err = ctl.NewStdClientSet(cfg)
		if err != nil {
			return err
		}

		oidc, err = ctl.NewOpenIDConnectManager(cfg)
		if err != nil {
			if _, ok := err.(*eks.UnsupportedOIDCError); !ok {
				return err
			}
			oidcSupported = false
		}
	}

	stackManager := ctl.NewStackManager(cfg)

	if err := deleteFargateProfiles(cmd, ctl); err != nil {
		return err
	}

	ssh.DeleteKeys(meta.Name, ctl.Provider.EC2())

	kubeconfig.MaybeDeleteConfig(meta)

	if hasDeprecatedStacks, err := deleteDeprecatedStacks(stackManager); hasDeprecatedStacks {
		if err != nil {
			return err
		}
		return nil
	}

	{
		// only need to cleanup ELBs if the cluster has already been created.
		if clusterOperable {
			ctx, cleanup := context.WithTimeout(context.Background(), 10*time.Minute)
			defer cleanup()

			logger.Info("cleaning up LoadBalancer services")
			if err := elb.Cleanup(ctx, ctl.Provider.EC2(), ctl.Provider.ELB(), ctl.Provider.ELBV2(), clientSet, cfg); err != nil {
				return err
			}
		}

		deleteOIDCProvider := clusterOperable && oidcSupported
		tasks, err := stackManager.NewTasksToDeleteClusterWithNodeGroups(deleteOIDCProvider, oidc, kubernetes.NewCachedClientSet(clientSet), cmd.Wait, func(errs chan error, _ string) error {
			logger.Info("trying to cleanup dangling network interfaces")
			if err := ctl.LoadClusterVPC(cfg); err != nil {
				return errors.Wrapf(err, "getting VPC configuration for cluster %q", cfg.Metadata.Name)
			}

			go func() {
				errs <- vpc.CleanupNetworkInterfaces(ctl.Provider.EC2(), cfg)
				close(errs)
			}()
			return nil
		})

		if err != nil {
			return err
		}

		if tasks.Len() == 0 {
			logger.Warning("no cluster resources were found for %q", meta.Name)
			return nil
		}

		logger.Info(tasks.Describe())
		if errs := tasks.DoAllSync(); len(errs) > 0 {
			return handleErrors(errs, "cluster with nodegroup(s)")
		}

		logger.Success("all cluster resources were deleted")
	}

	return nil
}

func deleteFargateProfiles(cmd *cmdutils.Cmd, ctl *eks.ClusterProvider) error {
	awsClient := fargate.NewClientWithWaitTimeout(
		cmd.ClusterConfig.Metadata.Name,
		ctl.Provider.EKS(),
		cmd.ProviderConfig.WaitTimeout,
	)
	profileNames, err := awsClient.ListProfiles()
	if err != nil {
		return err
	}
	// Linearise the deleting of Fargate profiles by passing as the API
	// otherwise errors out with:
	//   ResourceInUseException: Cannot delete Fargate Profile ${name2} because
	//   cluster ${clusterName} currently has Fargate profile ${name1} in
	//   status DELETING
	waitForDeletion := cmd.Wait || len(profileNames) > 1
	if len(profileNames) > 0 {
		for _, profileName := range profileNames {
			logger.Info("deleting Fargate profile %q", *profileName)
			if err := awsClient.DeleteProfile(*profileName, waitForDeletion); err != nil {
				return err
			}
			logger.Info("deleted Fargate profile %q", *profileName)
		}
	}
	logger.Info("deleted %v Fargate profile(s)", len(profileNames))
	return nil
}
