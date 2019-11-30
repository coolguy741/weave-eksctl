package create

import (
	"fmt"
	"time"

	"github.com/kris-nova/logger"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
	"github.com/weaveworks/eksctl/pkg/eks"
	"github.com/weaveworks/eksctl/pkg/fargate"
	"github.com/weaveworks/eksctl/pkg/fargate/coredns"
	"github.com/weaveworks/eksctl/pkg/utils/retry"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func createFargateProfile(cmd *cmdutils.Cmd) {
	cmd.ClusterConfig = api.NewClusterConfig()
	cmd.SetDescription(
		"fargateprofile",
		"Create a Fargate profile",
		"",
	)
	options := configureCreateFargateProfileCmd(cmd)
	cmd.SetRunFuncWithNameArg(func() error {
		return doCreateFargateProfile(cmd, options)
	})
}

func configureCreateFargateProfileCmd(cmd *cmdutils.Cmd) *fargate.CreateOptions {
	var options fargate.CreateOptions
	cmd.FlagSetGroup.InFlagSet("Fargate", func(fs *pflag.FlagSet) {
		cmdutils.AddFlagsForFargateProfileCreation(fs, &options)
	})
	cmd.FlagSetGroup.InFlagSet("General", func(fs *pflag.FlagSet) {
		cmdutils.AddClusterFlag(fs, cmd.ClusterConfig.Metadata)
		cmdutils.AddRegionFlag(fs, cmd.ProviderConfig)
		cmdutils.AddConfigFileFlag(fs, &cmd.ClusterConfigFile)
		cmdutils.AddWaitFlagWithFullDescription(fs, &cmd.Wait, "wait for the creation of the Fargate profile before exiting. Profile creation may take a few seconds to a couple of minutes.")
		cmdutils.AddTimeoutFlag(fs, &cmd.ProviderConfig.WaitTimeout)
	})
	cmdutils.AddCommonFlagsForAWS(cmd.FlagSetGroup, cmd.ProviderConfig, false)
	return &options
}

func doCreateFargateProfile(cmd *cmdutils.Cmd, options *fargate.CreateOptions) error {
	if err := cmdutils.NewCreateFargateProfileLoader(cmd, options).Load(); err != nil {
		cmd.CobraCommand.Help()
		return err
	}
	ctl, err := cmd.NewCtl()
	if err != nil {
		return err
	}
	if err := ctl.CheckAuth(); err != nil {
		return err
	}
	cfg := cmd.ClusterConfig
	if ok, err := ctl.CanOperate(cfg); !ok {
		return err
	}

	roleARN, err := getClusterRoleARN(ctl, cfg.Metadata)
	if err != nil {
		return err
	}
	if err := doCreateFargateProfiles(cmd, ctl, roleARN, cmd.Wait); err != nil {
		return err
	}
	clientSet, err := clientSet(cfg, ctl)
	if err != nil {
		return err
	}
	return scheduleCoreDNSOnFargateIfRelevant(cmd, clientSet)
}

func clientSet(cfg *api.ClusterConfig, ctl *eks.ClusterProvider) (kubernetes.Interface, error) {
	kubernetesClientConfigs, err := ctl.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	k8sConfig := kubernetesClientConfigs.Config
	k8sRestConfig, err := clientcmd.NewDefaultClientConfig(*k8sConfig, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, err
	}
	k8sClientSet, err := kubernetes.NewForConfig(k8sRestConfig)
	if err != nil {
		return nil, err
	}
	return k8sClientSet, nil
}

func getClusterRoleARN(ctl *eks.ClusterProvider, meta *api.ClusterMeta) (string, error) {
	eksCluster, err := ctl.DescribeControlPlane(meta)
	if err != nil {
		return "", errors.Wrapf(err, "failed to retrieve EKS cluster role ARN for %q", meta.Name)
	}
	roleARN := *eksCluster.RoleArn
	logger.Debug("default Fargate profile pod execution role ARN: %v", roleARN)
	return roleARN, nil
}

func doCreateFargateProfiles(cmd *cmdutils.Cmd, ctl *eks.ClusterProvider, defaultPodExecRoleARN string, wait bool) error {
	clusterName := cmd.ClusterConfig.Metadata.Name
	awsClient := fargate.NewClientWithWaitTimeout(clusterName, ctl.Provider.EKS(), cmd.ProviderConfig.WaitTimeout)
	for _, profile := range cmd.ClusterConfig.FargateProfiles {
		if wait {
			logger.Info(creatingFargateProfileMsg(clusterName, profile.Name))
		} else {
			logger.Debug(creatingFargateProfileMsg(clusterName, profile.Name))
		}

		// Default the pod execution role ARN to be the same as the cluster
		// role defined in CloudFormation:
		if profile.PodExecutionRoleARN == "" {
			profile.PodExecutionRoleARN = defaultPodExecRoleARN
		}
		if err := awsClient.CreateProfile(profile, wait); err != nil {
			return errors.Wrapf(err, "failed to create Fargate profile %q on EKS cluster %q", profile.Name, clusterName)
		}
		logger.Info("created Fargate profile %q on EKS cluster %q", profile.Name, clusterName)
	}
	return nil
}

func creatingFargateProfileMsg(clusterName, profileName string) string {
	return fmt.Sprintf("creating Fargate profile %q on EKS cluster %q", profileName, clusterName)
}

func scheduleCoreDNSOnFargateIfRelevant(cmd *cmdutils.Cmd, clientSet kubernetes.Interface) error {
	if coredns.IsSchedulableOnFargate(cmd.ClusterConfig.FargateProfiles) {
		scheduled, err := coredns.IsScheduledOnFargate(clientSet)
		if err != nil {
			return err
		}
		if !scheduled {
			if err := coredns.ScheduleOnFargate(clientSet); err != nil {
				return err
			}
			retryPolicy := &retry.TimingOutExponentialBackoff{
				Timeout:  cmd.ProviderConfig.WaitTimeout,
				TimeUnit: time.Second,
			}
			if err := coredns.WaitForScheduleOnFargate(clientSet, retryPolicy); err != nil {
				return err
			}
		}
	}
	return nil
}
