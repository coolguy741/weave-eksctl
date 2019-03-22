// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha4

import (
	ipnet "github.com/weaveworks/eksctl/pkg/utils/ipnet"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(ClusterMeta)
		(*in).DeepCopyInto(*out)
	}
	out.IAM = in.IAM
	if in.VPC != nil {
		in, out := &in.VPC, &out.VPC
		*out = new(ClusterVPC)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeGroups != nil {
		in, out := &in.NodeGroups, &out.NodeGroups
		*out = make([]*NodeGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NodeGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ClusterStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigList) DeepCopyInto(out *ClusterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigList.
func (in *ClusterConfigList) DeepCopy() *ClusterConfigList {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIAM) DeepCopyInto(out *ClusterIAM) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIAM.
func (in *ClusterIAM) DeepCopy() *ClusterIAM {
	if in == nil {
		return nil
	}
	out := new(ClusterIAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMeta) DeepCopyInto(out *ClusterMeta) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMeta.
func (in *ClusterMeta) DeepCopy() *ClusterMeta {
	if in == nil {
		return nil
	}
	out := new(ClusterMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSubnets) DeepCopyInto(out *ClusterSubnets) {
	*out = *in
	if in.Private != nil {
		in, out := &in.Private, &out.Private
		*out = make(map[string]Network, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = make(map[string]Network, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSubnets.
func (in *ClusterSubnets) DeepCopy() *ClusterSubnets {
	if in == nil {
		return nil
	}
	out := new(ClusterSubnets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVPC) DeepCopyInto(out *ClusterVPC) {
	*out = *in
	in.Network.DeepCopyInto(&out.Network)
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = new(ClusterSubnets)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraCIDRs != nil {
		in, out := &in.ExtraCIDRs, &out.ExtraCIDRs
		*out = make([]*ipnet.IPNet, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVPC.
func (in *ClusterVPC) DeepCopy() *ClusterVPC {
	if in == nil {
		return nil
	}
	out := new(ClusterVPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.CIDR != nil {
		in, out := &in.CIDR, &out.CIDR
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroup) DeepCopyInto(out *NodeGroup) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = new(NodeGroupSGs)
		(*in).DeepCopyInto(*out)
	}
	if in.DesiredCapacity != nil {
		in, out := &in.DesiredCapacity, &out.DesiredCapacity
		*out = new(int)
		**out = **in
	}
	if in.MinSize != nil {
		in, out := &in.MinSize, &out.MinSize
		*out = new(int)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(int)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SSHPublicKey != nil {
		in, out := &in.SSHPublicKey, &out.SSHPublicKey
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.IAM != nil {
		in, out := &in.IAM, &out.IAM
		*out = new(NodeGroupIAM)
		(*in).DeepCopyInto(*out)
	}
	if in.PreBootstrapCommands != nil {
		in, out := &in.PreBootstrapCommands, &out.PreBootstrapCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OverrideBootstrapCommand != nil {
		in, out := &in.OverrideBootstrapCommand, &out.OverrideBootstrapCommand
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroup.
func (in *NodeGroup) DeepCopy() *NodeGroup {
	if in == nil {
		return nil
	}
	out := new(NodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupIAM) DeepCopyInto(out *NodeGroupIAM) {
	*out = *in
	if in.AttachPolicyARNs != nil {
		in, out := &in.AttachPolicyARNs, &out.AttachPolicyARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.WithAddonPolicies.DeepCopyInto(&out.WithAddonPolicies)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupIAM.
func (in *NodeGroupIAM) DeepCopy() *NodeGroupIAM {
	if in == nil {
		return nil
	}
	out := new(NodeGroupIAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupIAMAddonPolicies) DeepCopyInto(out *NodeGroupIAMAddonPolicies) {
	*out = *in
	if in.ImageBuilder != nil {
		in, out := &in.ImageBuilder, &out.ImageBuilder
		*out = new(bool)
		**out = **in
	}
	if in.AutoScaler != nil {
		in, out := &in.AutoScaler, &out.AutoScaler
		*out = new(bool)
		**out = **in
	}
	if in.ExternalDNS != nil {
		in, out := &in.ExternalDNS, &out.ExternalDNS
		*out = new(bool)
		**out = **in
	}
	if in.AppMesh != nil {
		in, out := &in.AppMesh, &out.AppMesh
		*out = new(bool)
		**out = **in
	}
	if in.EBSCSI != nil {
		in, out := &in.EBSCSI, &out.EBSCSI
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupIAMAddonPolicies.
func (in *NodeGroupIAMAddonPolicies) DeepCopy() *NodeGroupIAMAddonPolicies {
	if in == nil {
		return nil
	}
	out := new(NodeGroupIAMAddonPolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupSGs) DeepCopyInto(out *NodeGroupSGs) {
	*out = *in
	if in.AttachIDs != nil {
		in, out := &in.AttachIDs, &out.AttachIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WithShared != nil {
		in, out := &in.WithShared, &out.WithShared
		*out = new(bool)
		**out = **in
	}
	if in.WithLocal != nil {
		in, out := &in.WithLocal, &out.WithLocal
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupSGs.
func (in *NodeGroupSGs) DeepCopy() *NodeGroupSGs {
	if in == nil {
		return nil
	}
	out := new(NodeGroupSGs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfig) DeepCopyInto(out *ProviderConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfig.
func (in *ProviderConfig) DeepCopy() *ProviderConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderConfig)
	in.DeepCopyInto(out)
	return out
}
