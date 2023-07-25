//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2023 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentCondition) DeepCopyInto(out *AgentCondition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentCondition.
func (in *AgentCondition) DeepCopy() *AgentCondition {
	if in == nil {
		return nil
	}
	out := new(AgentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaAgentInfo) DeepCopyInto(out *AntreaAgentInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.PodRef = in.PodRef
	out.NodeRef = in.NodeRef
	if in.NodeSubnets != nil {
		in, out := &in.NodeSubnets, &out.NodeSubnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.OVSInfo.DeepCopyInto(&out.OVSInfo)
	out.NetworkPolicyControllerInfo = in.NetworkPolicyControllerInfo
	if in.AgentConditions != nil {
		in, out := &in.AgentConditions, &out.AgentConditions
		*out = make([]AgentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.APICABundle != nil {
		in, out := &in.APICABundle, &out.APICABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaAgentInfo.
func (in *AntreaAgentInfo) DeepCopy() *AntreaAgentInfo {
	if in == nil {
		return nil
	}
	out := new(AntreaAgentInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AntreaAgentInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaAgentInfoList) DeepCopyInto(out *AntreaAgentInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AntreaAgentInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaAgentInfoList.
func (in *AntreaAgentInfoList) DeepCopy() *AntreaAgentInfoList {
	if in == nil {
		return nil
	}
	out := new(AntreaAgentInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AntreaAgentInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaControllerInfo) DeepCopyInto(out *AntreaControllerInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.PodRef = in.PodRef
	out.NodeRef = in.NodeRef
	out.ServiceRef = in.ServiceRef
	out.NetworkPolicyControllerInfo = in.NetworkPolicyControllerInfo
	if in.ControllerConditions != nil {
		in, out := &in.ControllerConditions, &out.ControllerConditions
		*out = make([]ControllerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaControllerInfo.
func (in *AntreaControllerInfo) DeepCopy() *AntreaControllerInfo {
	if in == nil {
		return nil
	}
	out := new(AntreaControllerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AntreaControllerInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaControllerInfoList) DeepCopyInto(out *AntreaControllerInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AntreaControllerInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaControllerInfoList.
func (in *AntreaControllerInfoList) DeepCopy() *AntreaControllerInfoList {
	if in == nil {
		return nil
	}
	out := new(AntreaControllerInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AntreaControllerInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedTo) DeepCopyInto(out *AppliedTo) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalEntitySelector != nil {
		in, out := &in.ExternalEntitySelector, &out.ExternalEntitySelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(NamespacedName)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(NamespacedName)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedTo.
func (in *AppliedTo) DeepCopy() *AppliedTo {
	if in == nil {
		return nil
	}
	out := new(AppliedTo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGroup) DeepCopyInto(out *ClusterGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGroup.
func (in *ClusterGroup) DeepCopy() *ClusterGroup {
	if in == nil {
		return nil
	}
	out := new(ClusterGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGroupList) DeepCopyInto(out *ClusterGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGroupList.
func (in *ClusterGroupList) DeepCopy() *ClusterGroupList {
	if in == nil {
		return nil
	}
	out := new(ClusterGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkPolicy) DeepCopyInto(out *ClusterNetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkPolicy.
func (in *ClusterNetworkPolicy) DeepCopy() *ClusterNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterNetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkPolicyList) DeepCopyInto(out *ClusterNetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkPolicyList.
func (in *ClusterNetworkPolicyList) DeepCopy() *ClusterNetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterNetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkPolicySpec) DeepCopyInto(out *ClusterNetworkPolicySpec) {
	*out = *in
	if in.AppliedTo != nil {
		in, out := &in.AppliedTo, &out.AppliedTo
		*out = make([]AppliedTo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkPolicySpec.
func (in *ClusterNetworkPolicySpec) DeepCopy() *ClusterNetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerCondition) DeepCopyInto(out *ControllerCondition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerCondition.
func (in *ControllerCondition) DeepCopy() *ControllerCondition {
	if in == nil {
		return nil
	}
	out := new(ControllerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalIPPool) DeepCopyInto(out *ExternalIPPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalIPPool.
func (in *ExternalIPPool) DeepCopy() *ExternalIPPool {
	if in == nil {
		return nil
	}
	out := new(ExternalIPPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalIPPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalIPPoolList) DeepCopyInto(out *ExternalIPPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalIPPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalIPPoolList.
func (in *ExternalIPPoolList) DeepCopy() *ExternalIPPoolList {
	if in == nil {
		return nil
	}
	out := new(ExternalIPPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalIPPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalIPPoolSpec) DeepCopyInto(out *ExternalIPPoolSpec) {
	*out = *in
	if in.IPRanges != nil {
		in, out := &in.IPRanges, &out.IPRanges
		*out = make([]IPRange, len(*in))
		copy(*out, *in)
	}
	in.NodeSelector.DeepCopyInto(&out.NodeSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalIPPoolSpec.
func (in *ExternalIPPoolSpec) DeepCopy() *ExternalIPPoolSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalIPPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalIPPoolStatus) DeepCopyInto(out *ExternalIPPoolStatus) {
	*out = *in
	out.Usage = in.Usage
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalIPPoolStatus.
func (in *ExternalIPPoolStatus) DeepCopy() *ExternalIPPoolStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalIPPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Group) DeepCopyInto(out *Group) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Group.
func (in *Group) DeepCopy() *Group {
	if in == nil {
		return nil
	}
	out := new(Group)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Group) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupCondition) DeepCopyInto(out *GroupCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupCondition.
func (in *GroupCondition) DeepCopy() *GroupCondition {
	if in == nil {
		return nil
	}
	out := new(GroupCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupList) DeepCopyInto(out *GroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Group, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupList.
func (in *GroupList) DeepCopy() *GroupList {
	if in == nil {
		return nil
	}
	out := new(GroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSpec) DeepCopyInto(out *GroupSpec) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.IPBlocks != nil {
		in, out := &in.IPBlocks, &out.IPBlocks
		*out = make([]IPBlock, len(*in))
		copy(*out, *in)
	}
	if in.ServiceReference != nil {
		in, out := &in.ServiceReference, &out.ServiceReference
		*out = new(NamespacedName)
		**out = **in
	}
	if in.ExternalEntitySelector != nil {
		in, out := &in.ExternalEntitySelector, &out.ExternalEntitySelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ChildGroups != nil {
		in, out := &in.ChildGroups, &out.ChildGroups
		*out = make([]ClusterGroupReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSpec.
func (in *GroupSpec) DeepCopy() *GroupSpec {
	if in == nil {
		return nil
	}
	out := new(GroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupStatus) DeepCopyInto(out *GroupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]GroupCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupStatus.
func (in *GroupStatus) DeepCopy() *GroupStatus {
	if in == nil {
		return nil
	}
	out := new(GroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProtocol) DeepCopyInto(out *HTTPProtocol) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProtocol.
func (in *HTTPProtocol) DeepCopy() *HTTPProtocol {
	if in == nil {
		return nil
	}
	out := new(HTTPProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ICMPProtocol) DeepCopyInto(out *ICMPProtocol) {
	*out = *in
	if in.ICMPType != nil {
		in, out := &in.ICMPType, &out.ICMPType
		*out = new(int32)
		**out = **in
	}
	if in.ICMPCode != nil {
		in, out := &in.ICMPCode, &out.ICMPCode
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMPProtocol.
func (in *ICMPProtocol) DeepCopy() *ICMPProtocol {
	if in == nil {
		return nil
	}
	out := new(ICMPProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IGMPProtocol) DeepCopyInto(out *IGMPProtocol) {
	*out = *in
	if in.IGMPType != nil {
		in, out := &in.IGMPType, &out.IGMPType
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IGMPProtocol.
func (in *IGMPProtocol) DeepCopy() *IGMPProtocol {
	if in == nil {
		return nil
	}
	out := new(IGMPProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlock) DeepCopyInto(out *IPBlock) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlock.
func (in *IPBlock) DeepCopy() *IPBlock {
	if in == nil {
		return nil
	}
	out := new(IPBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolUsage) DeepCopyInto(out *IPPoolUsage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolUsage.
func (in *IPPoolUsage) DeepCopy() *IPPoolUsage {
	if in == nil {
		return nil
	}
	out := new(IPPoolUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPRange) DeepCopyInto(out *IPRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPRange.
func (in *IPRange) DeepCopy() *IPRange {
	if in == nil {
		return nil
	}
	out := new(IPRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L7Protocol) DeepCopyInto(out *L7Protocol) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(HTTPProtocol)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSProtocol)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L7Protocol.
func (in *L7Protocol) DeepCopy() *L7Protocol {
	if in == nil {
		return nil
	}
	out := new(L7Protocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicy) DeepCopyInto(out *NetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicy.
func (in *NetworkPolicy) DeepCopy() *NetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyCondition) DeepCopyInto(out *NetworkPolicyCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyCondition.
func (in *NetworkPolicyCondition) DeepCopy() *NetworkPolicyCondition {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyControllerInfo) DeepCopyInto(out *NetworkPolicyControllerInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyControllerInfo.
func (in *NetworkPolicyControllerInfo) DeepCopy() *NetworkPolicyControllerInfo {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyControllerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyList) DeepCopyInto(out *NetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyList.
func (in *NetworkPolicyList) DeepCopy() *NetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyPeer) DeepCopyInto(out *NetworkPolicyPeer) {
	*out = *in
	if in.IPBlock != nil {
		in, out := &in.IPBlock, &out.IPBlock
		*out = new(IPBlock)
		**out = **in
	}
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = new(PeerNamespaces)
		**out = **in
	}
	if in.ExternalEntitySelector != nil {
		in, out := &in.ExternalEntitySelector, &out.ExternalEntitySelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(NamespacedName)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyPeer.
func (in *NetworkPolicyPeer) DeepCopy() *NetworkPolicyPeer {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyPort) DeepCopyInto(out *NetworkPolicyPort) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(corev1.Protocol)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.EndPort != nil {
		in, out := &in.EndPort, &out.EndPort
		*out = new(int32)
		**out = **in
	}
	if in.SourcePort != nil {
		in, out := &in.SourcePort, &out.SourcePort
		*out = new(int32)
		**out = **in
	}
	if in.SourceEndPort != nil {
		in, out := &in.SourceEndPort, &out.SourceEndPort
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyPort.
func (in *NetworkPolicyPort) DeepCopy() *NetworkPolicyPort {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyProtocol) DeepCopyInto(out *NetworkPolicyProtocol) {
	*out = *in
	if in.ICMP != nil {
		in, out := &in.ICMP, &out.ICMP
		*out = new(ICMPProtocol)
		(*in).DeepCopyInto(*out)
	}
	if in.IGMP != nil {
		in, out := &in.IGMP, &out.IGMP
		*out = new(IGMPProtocol)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyProtocol.
func (in *NetworkPolicyProtocol) DeepCopy() *NetworkPolicyProtocol {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicySpec) DeepCopyInto(out *NetworkPolicySpec) {
	*out = *in
	if in.AppliedTo != nil {
		in, out := &in.AppliedTo, &out.AppliedTo
		*out = make([]AppliedTo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicySpec.
func (in *NetworkPolicySpec) DeepCopy() *NetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyStatus) DeepCopyInto(out *NetworkPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]NetworkPolicyCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyStatus.
func (in *NetworkPolicyStatus) DeepCopy() *NetworkPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OVSInfo) DeepCopyInto(out *OVSInfo) {
	*out = *in
	if in.FlowTable != nil {
		in, out := &in.FlowTable, &out.FlowTable
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OVSInfo.
func (in *OVSInfo) DeepCopy() *OVSInfo {
	if in == nil {
		return nil
	}
	out := new(OVSInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerNamespaces) DeepCopyInto(out *PeerNamespaces) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerNamespaces.
func (in *PeerNamespaces) DeepCopy() *PeerNamespaces {
	if in == nil {
		return nil
	}
	out := new(PeerNamespaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerService) DeepCopyInto(out *PeerService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerService.
func (in *PeerService) DeepCopy() *PeerService {
	if in == nil {
		return nil
	}
	out := new(PeerService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(RuleAction)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]NetworkPolicyPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]NetworkPolicyProtocol, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.L7Protocols != nil {
		in, out := &in.L7Protocols, &out.L7Protocols
		*out = make([]L7Protocol, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]NetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]NetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToServices != nil {
		in, out := &in.ToServices, &out.ToServices
		*out = make([]PeerService, len(*in))
		copy(*out, *in)
	}
	if in.AppliedTo != nil {
		in, out := &in.AppliedTo, &out.AppliedTo
		*out = make([]AppliedTo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSProtocol) DeepCopyInto(out *TLSProtocol) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSProtocol.
func (in *TLSProtocol) DeepCopy() *TLSProtocol {
	if in == nil {
		return nil
	}
	out := new(TLSProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tier) DeepCopyInto(out *Tier) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tier.
func (in *Tier) DeepCopy() *Tier {
	if in == nil {
		return nil
	}
	out := new(Tier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tier) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TierList) DeepCopyInto(out *TierList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TierList.
func (in *TierList) DeepCopy() *TierList {
	if in == nil {
		return nil
	}
	out := new(TierList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TierList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TierSpec) DeepCopyInto(out *TierSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TierSpec.
func (in *TierSpec) DeepCopy() *TierSpec {
	if in == nil {
		return nil
	}
	out := new(TierSpec)
	in.DeepCopyInto(out)
	return out
}
