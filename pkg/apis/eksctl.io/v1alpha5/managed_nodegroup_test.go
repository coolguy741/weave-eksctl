package v1alpha5

import (
	"github.com/aws/aws-sdk-go/aws"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

type nodeGroupCase struct {
	ng     *ManagedNodeGroup
	errMsg string
}

var _ = Describe("Managed Nodegroup Validation", func() {
	DescribeTable("Unsupported field combinations", func(n *nodeGroupCase) {
		SetManagedNodeGroupDefaults(n.ng, &ClusterMeta{Name: "managed-cluster"})
		err := ValidateManagedNodeGroup(n.ng, 0)
		if n.errMsg == "" {
			Expect(err).ToNot(HaveOccurred())
			return
		}
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring(n.errMsg))

	},
		Entry("Unsupported AMI family", &nodeGroupCase{
			ng: &ManagedNodeGroup{
				NodeGroupBase: &NodeGroupBase{
					AMI:                      "",
					OverrideBootstrapCommand: nil,
					PreBootstrapCommands:     nil,
					AMIFamily:                "Ubuntu1804",
				},
			},
			errMsg: "only AmazonLinux2 is supported",
		}),
		Entry("Custom AMI without overrideBootstrapCommand", &nodeGroupCase{
			ng: &ManagedNodeGroup{
				NodeGroupBase: &NodeGroupBase{
					AMI: "ami-custom",
				},
			},
			errMsg: "overrideBootstrapCommand is required when a custom AMI",
		}),

		Entry("launchTemplate with no ID", &nodeGroupCase{
			ng: &ManagedNodeGroup{
				NodeGroupBase:  &NodeGroupBase{},
				LaunchTemplate: &LaunchTemplate{},
			},
			errMsg: "launchTemplate.id is required",
		}),

		Entry("launchTemplate with invalid version", &nodeGroupCase{
			ng: &ManagedNodeGroup{
				NodeGroupBase: &NodeGroupBase{},
				LaunchTemplate: &LaunchTemplate{
					ID:      "lt-custom",
					Version: aws.Int(0),
				},
			},
			errMsg: "launchTemplate.version must be >= 1",
		}),
	)

	DescribeTable("User-supplied launch template with unsupported fields", func(ngBase *NodeGroupBase) {
		mng := &ManagedNodeGroup{
			NodeGroupBase: ngBase,
			LaunchTemplate: &LaunchTemplate{
				ID: "lt-custom",
			},
		}
		SetManagedNodeGroupDefaults(mng, &ClusterMeta{Name: "managed-cluster"})
		err := ValidateManagedNodeGroup(mng, 0)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("cannot set instanceType, ami, ssh.allow, ssh.sourceSecurityGroupIds, securityGroups, " +
			"volumeSize, preBootstrapCommands or overrideBootstrapCommand in managedNodeGroup when a launch template is supplied"))
	},
		Entry("instanceType", &NodeGroupBase{
			InstanceType: "m5.xlarge",
		}),
		Entry("AMI", &NodeGroupBase{
			AMI: "ami-custom",
		}),
		Entry("SSH", &NodeGroupBase{
			SSH: &NodeGroupSSH{
				Allow: Enabled(),
			},
		}),
		Entry("volumeSize", &NodeGroupBase{
			VolumeSize: aws.Int(100),
		}),
		Entry("preBootstrapCommands", &NodeGroupBase{
			PreBootstrapCommands: []string{"echo test"},
		}),
		Entry("overrideBootstrapCommand", &NodeGroupBase{
			OverrideBootstrapCommand: aws.String("bootstrap.sh"),
		}),
		Entry("overrideBootstrapCommand", &NodeGroupBase{
			SecurityGroups: &NodeGroupSGs{
				AttachIDs: []string{"sg-custom"},
			},
		}),
	)
})
