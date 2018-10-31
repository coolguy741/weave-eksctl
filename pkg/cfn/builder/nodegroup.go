package builder

import (
	"fmt"

	cfn "github.com/aws/aws-sdk-go/service/cloudformation"
	gfn "github.com/awslabs/goformation/cloudformation"

	"github.com/weaveworks/eksctl/pkg/eks/api"

	"github.com/weaveworks/eksctl/pkg/nodebootstrap"
)

// NodeGroupResourceSet stores the resource information of the node group
type NodeGroupResourceSet struct {
	rs               *resourceSet
	id               int
	clusterSpec      *api.ClusterConfig
	spec             *api.NodeGroup
	clusterStackName string
	nodeGroupName    string
	instanceProfile  *gfn.Value
	securityGroups   []*gfn.Value
	vpc              *gfn.Value
	userData         *gfn.Value
}

// NewNodeGroupResourceSet returns a resource set for the new node group
func NewNodeGroupResourceSet(spec *api.ClusterConfig, clusterStackName string, id int) *NodeGroupResourceSet {
	return &NodeGroupResourceSet{
		rs:               newResourceSet(),
		id:               id,
		clusterStackName: clusterStackName,
		nodeGroupName:    fmt.Sprintf("%s-%d", spec.ClusterName, id),
		clusterSpec:      spec,
		spec:             spec.NodeGroups[id],
	}
}

// AddAllResources adds all the information about the node group to the resource set
func (n *NodeGroupResourceSet) AddAllResources() error {
	n.rs.template.Description = nodeGroupTemplateDescription
	n.rs.template.Description += nodeGroupTemplateDescriptionDefaultFeatures
	n.rs.template.Description += templateDescriptionSuffix

	n.vpc = makeImportValue(n.clusterStackName, cfnOutputClusterVPC)

	userData, err := nodebootstrap.NewUserData(n.clusterSpec, n.id)
	if err != nil {
		return err
	}
	n.userData = gfn.NewString(userData)

	if n.spec.MinSize == 0 && n.spec.MaxSize == 0 {
		n.spec.MinSize = n.spec.DesiredCapacity
		n.spec.MaxSize = n.spec.DesiredCapacity
	}

	n.addResourcesForIAM()
	n.addResourcesForSecurityGroups()

	return n.addResourcesForNodeGroup()
}

// RenderJSON returns the rendered JSON
func (n *NodeGroupResourceSet) RenderJSON() ([]byte, error) {
	return n.rs.renderJSON()
}

// Template returns the CloudFormation template
func (n *NodeGroupResourceSet) Template() gfn.Template {
	return *n.rs.template
}

func (n *NodeGroupResourceSet) newResource(name string, resource interface{}) *gfn.Value {
	return n.rs.newResource(name, resource)
}

func (n *NodeGroupResourceSet) addResourcesForNodeGroup() error {
	lc := &gfn.AWSAutoScalingLaunchConfiguration{
		IamInstanceProfile: n.instanceProfile,
		SecurityGroups:     n.securityGroups,
		ImageId:            gfn.NewString(n.spec.AMI),
		InstanceType:       gfn.NewString(n.spec.InstanceType),
		UserData:           n.userData,
	}
	if n.spec.AllowSSH {
		lc.KeyName = gfn.NewString(n.spec.SSHPublicKeyName)
	}
	if n.spec.PrivateNetworking {
		lc.AssociatePublicIpAddress = gfn.False()
	} else {
		lc.AssociatePublicIpAddress = gfn.True()
	}
	if n.spec.VolumeSize > 0 {
		lc.BlockDeviceMappings = []gfn.AWSAutoScalingLaunchConfiguration_BlockDeviceMapping{
			{
				DeviceName: gfn.NewString("/dev/xvda"),
				Ebs: &gfn.AWSAutoScalingLaunchConfiguration_BlockDevice{
					VolumeSize: gfn.NewInteger(n.spec.VolumeSize),
				},
			},
		}
	}
	refLC := n.newResource("NodeLaunchConfig", lc)
	// currently goformation type system doesn't allow specifying `VPCZoneIdentifier: { "Fn::ImportValue": ... }`,
	// and tags don't have `PropagateAtLaunch` field, so we have a custom method here until this gets resolved
	var vpcZoneIdentifier interface{}
	if numNodeGroupsAZs := len(n.spec.AvailabilityZones); numNodeGroupsAZs > 0 {
		subnets := n.clusterSpec.VPC.Subnets[n.spec.SubnetTopology()]
		errorDesc := fmt.Sprintf("(subnets=%#v AZs=%#v)", subnets, n.spec.AvailabilityZones)
		if len(subnets) < numNodeGroupsAZs {
			return fmt.Errorf("VPC doesn't have enough subnets for nodegroup AZs %s", errorDesc)
		}
		vpcZoneIdentifier = make([]interface{}, numNodeGroupsAZs)
		for i, az := range n.spec.AvailabilityZones {
			subnet, ok := subnets[az]
			if !ok {
				return fmt.Errorf("VPC doesn't have subnets in %s %s", az, errorDesc)
			}
			vpcZoneIdentifier.([]interface{})[i] = subnet.ID
		}
	} else {
		vpcZoneIdentifier = map[string][]interface{}{
			gfn.FnSplit: []interface{}{
				",",
				makeImportValue(n.clusterStackName, cfnOutputClusterSubnets+string(n.spec.SubnetTopology())),
			},
		}
	}
	n.newResource("NodeGroup", &awsCloudFormationResource{
		Type: "AWS::AutoScaling::AutoScalingGroup",
		Properties: map[string]interface{}{
			"LaunchConfigurationName": refLC,
			"DesiredCapacity":         fmt.Sprintf("%d", n.spec.DesiredCapacity),
			"MinSize":                 fmt.Sprintf("%d", n.spec.MinSize),
			"MaxSize":                 fmt.Sprintf("%d", n.spec.MaxSize),
			"VPCZoneIdentifier":       vpcZoneIdentifier,
			"Tags": []map[string]interface{}{
				{"Key": "Name", "Value": fmt.Sprintf("%s-Node", n.nodeGroupName), "PropagateAtLaunch": "true"},
				{"Key": "kubernetes.io/cluster/" + n.clusterSpec.ClusterName, "Value": "owned", "PropagateAtLaunch": "true"},
			},
		},
		UpdatePolicy: map[string]map[string]string{
			"AutoScalingRollingUpdate": {
				"MinInstancesInService": "1",
				"MaxBatchSize":          "1",
			},
		},
	})

	return nil
}

// GetAllOutputs collects all outputs of the node group
func (n *NodeGroupResourceSet) GetAllOutputs(stack cfn.Stack) error {
	return n.rs.GetAllOutputs(stack, n.spec)
}
