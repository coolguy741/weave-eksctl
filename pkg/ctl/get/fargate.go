package get

import (
	"github.com/kris-nova/logger"
	"github.com/spf13/pflag"
	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
	"github.com/weaveworks/eksctl/pkg/fargate"
)

func getFargateProfile(cmd *cmdutils.Cmd) {
	cmd.ClusterConfig = api.NewClusterConfig()
	cmd.SetDescription(
		"fargateprofile",
		"Get Fargate profile(s)",
		"",
	)
	opts := configureGetFargateProfileCmd(cmd)
	cmd.SetRunFuncWithNameArg(func() error {
		return doGetFargateProfile(cmd, opts)
	})
}

func configureGetFargateProfileCmd(cmd *cmdutils.Cmd) *fargate.Options {
	var opts fargate.Options
	cmd.FlagSetGroup.InFlagSet("Fargate", func(fs *pflag.FlagSet) {
		cmdutils.AddCommonFlagsForFargate(fs, &opts)
	})
	cmd.FlagSetGroup.InFlagSet("General", func(fs *pflag.FlagSet) {
		cmdutils.AddClusterFlag(fs, cmd.ClusterConfig.Metadata)
		cmdutils.AddRegionFlag(fs, cmd.ProviderConfig)
		cmdutils.AddConfigFileFlag(fs, &cmd.ClusterConfigFile)
		cmdutils.AddTimeoutFlag(fs, &cmd.ProviderConfig.WaitTimeout)
	})
	cmdutils.AddCommonFlagsForAWS(cmd.FlagSetGroup, cmd.ProviderConfig, false)
	return &opts
}

func doGetFargateProfile(cmd *cmdutils.Cmd, opts *fargate.Options) error {
	if err := opts.Validate(); err != nil {
		return err
	}
	logger.Info("getting Fargate profile \"%s\"", opts.ProfileName)
	return nil
}
