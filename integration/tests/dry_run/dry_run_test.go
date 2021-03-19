// +build integration

package dry_run

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"github.com/weaveworks/eksctl/integration/tests"
	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/eks"
	"github.com/weaveworks/eksctl/pkg/testutils"
	"github.com/weaveworks/eksctl/pkg/utils/ipnet"

	. "github.com/onsi/ginkgo"
)

var params *tests.Params

func init() {
	// Call testing.Init() prior to tests.NewParams(), as otherwise -test.* will not be recognised. See also: https://golang.org/doc/go1.13#testing
	testing.Init()
	// No cleanup required for dry-run clusters
	params = tests.NewParams("dry-run")
	if err := api.Register(); err != nil {
		panic(errors.Wrap(err, "unexpected error registering API scheme"))
	}
}

func TestDryRun(t *testing.T) {
	testutils.RegisterAndRun(t)
}

const defaultClusterConfig = `
apiVersion: eksctl.io/v1alpha5
cloudWatch:
  clusterLogging: {}
iam:
  vpcResourceControllerPolicy: true
  withOIDC: false
kind: ClusterConfig
metadata:
  name: dry-run-cluster
  region: us-west-2
  version: "1.18"
nodeGroups:
- amiFamily: AmazonLinux2
  disableIMDSv1: false
  disablePodIMDS: false
  iam:
    withAddonPolicies:
      albIngress: false
      appMesh: false
      appMeshPreview: false
      autoScaler: false
      certManager: false
      cloudWatch: false
      ebs: false
      efs: false
      externalDNS: false
      fsx: false
      imageBuilder: false
      xRay: false
  instanceType: m5.large
  labels:
    alpha.eksctl.io/cluster-name: dry-run-cluster
    alpha.eksctl.io/nodegroup-name: ng-default
  name: ng-default
  privateNetworking: false
  securityGroups:
    withLocal: true
    withShared: true
  ssh:
    allow: false
    enableSsm: false
  volumeIOPS: 3000
  volumeSize: 80
  volumeThroughput: 125
  volumeType: gp3

managedNodeGroups:
- amiFamily: AmazonLinux2
  desiredCapacity: 2
  disableIMDSv1: false
  disablePodIMDS: false
  iam:
    withAddonPolicies:
      albIngress: false
      appMesh: false
      appMeshPreview: false
      autoScaler: false
      certManager: false
      cloudWatch: false
      ebs: false
      efs: false
      externalDNS: false
      fsx: false
      imageBuilder: false
      xRay: false
  instanceType: m5.large
  labels:
    alpha.eksctl.io/cluster-name: dry-run-cluster
    alpha.eksctl.io/nodegroup-name: ng-default
  maxSize: 2
  minSize: 2
  name: ng-default
  privateNetworking: false
  securityGroups:
    withLocal: null
    withShared: null
  ssh:
    allow: false
    enableSsm: false
    publicKeyPath: ""
  tags:
    alpha.eksctl.io/nodegroup-name: ng-default
    alpha.eksctl.io/nodegroup-type: managed
  volumeIOPS: 3000
  volumeSize: 80
  volumeThroughput: 125
  volumeType: gp3

privateCluster:
  enabled: false
vpc:
  autoAllocateIPv6: false
  cidr: 192.168.0.0/16
  clusterEndpoints:
    privateAccess: false
    publicAccess: true
  manageSharedNodeSecurityGroupRules: true
  nat:
    gateway: Single
`

var _ = Describe("(Integration) [Dry-Run test]", func() {
	DescribeTable("`create cluster` with --dry-run", func(setValues func(*api.ClusterConfig), createArgs ...string) {
		cmd := params.EksctlCreateCmd.
			WithArgs(
				"cluster",
				"--dry-run",
				"--name=dry-run-cluster",
			).
			WithArgs(createArgs...)

		session := cmd.Run()
		Expect(session.ExitCode()).To(Equal(0))

		output := session.Buffer().Contents()
		actual, err := eks.ParseConfig(output)
		Expect(err).ToNot(HaveOccurred())

		expected, err := eks.ParseConfig([]byte(defaultClusterConfig))
		Expect(err).ToNot(HaveOccurred())
		setValues(expected)
		Expect(actual).To(Equal(expected))
	},
		Entry("default values", func(c *api.ClusterConfig) {
			c.ManagedNodeGroups = nil
		}, "--nodegroup-name=ng-default"),

		Entry("managed nodegroup defaults", func(c *api.ClusterConfig) {
			c.NodeGroups = nil
			ng := c.ManagedNodeGroups[0]
			ng.VolumeSize = aws.Int(101)
		}, "--managed", "--nodegroup-name=ng-default", "--node-volume-size=101"),

		Entry("override nodegroup defaults", func(c *api.ClusterConfig) {
			c.NodeGroups = nil
			ng := c.ManagedNodeGroups[0]
			ng.IAM.WithAddonPolicies.ExternalDNS = aws.Bool(true)
			ng.VolumeSize = aws.Int(42)
			ng.PrivateNetworking = true
		}, "--managed", "--nodegroup-name=ng-default", "--node-volume-size=42",
			"--external-dns-access", "--node-private-networking"),

		Entry("override cluster-wide defaults", func(c *api.ClusterConfig) {
			c.ManagedNodeGroups = nil
			c.NodeGroups = nil
			cidr, err := ipnet.ParseCIDR("192.168.0.0/24")
			ExpectWithOffset(1, err).ToNot(HaveOccurred(), "unexpected error parsing CIDR")
			c.VPC.CIDR = cidr
			c.VPC.NAT.Gateway = aws.String("HighlyAvailable")
			c.IAM.WithOIDC = aws.Bool(true)
		}, "--vpc-cidr=192.168.0.0/24", "--without-nodegroup", "--vpc-nat-mode=HighlyAvailable", "--with-oidc"),
	)

})
