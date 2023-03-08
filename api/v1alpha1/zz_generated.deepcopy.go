//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAdapterConfig) DeepCopyInto(out *AWSAdapterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAdapterConfig.
func (in *AWSAdapterConfig) DeepCopy() *AWSAdapterConfig {
	if in == nil {
		return nil
	}
	out := new(AWSAdapterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSAdapterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAdapterConfigList) DeepCopyInto(out *AWSAdapterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSAdapterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAdapterConfigList.
func (in *AWSAdapterConfigList) DeepCopy() *AWSAdapterConfigList {
	if in == nil {
		return nil
	}
	out := new(AWSAdapterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSAdapterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAdapterConfigSpec) DeepCopyInto(out *AWSAdapterConfigSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAdapterConfigSpec.
func (in *AWSAdapterConfigSpec) DeepCopy() *AWSAdapterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AWSAdapterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAdapterConfigStatus) DeepCopyInto(out *AWSAdapterConfigStatus) {
	*out = *in
	if in.LastUpdatedTimestamp != nil {
		in, out := &in.LastUpdatedTimestamp, &out.LastUpdatedTimestamp
		*out = (*in).DeepCopy()
	}
	in.LastPollInfo.DeepCopyInto(&out.LastPollInfo)
	if in.AccountData != nil {
		in, out := &in.AccountData, &out.AccountData
		*out = new(AccountData)
		(*in).DeepCopyInto(*out)
	}
	if in.EKSCluster != nil {
		in, out := &in.EKSCluster, &out.EKSCluster
		*out = new(EKSCluster)
		(*in).DeepCopyInto(*out)
	}
	if in.ECRRepositories != nil {
		in, out := &in.ECRRepositories, &out.ECRRepositories
		*out = make([]*ECRRepository, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ECRRepository)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAdapterConfigStatus.
func (in *AWSAdapterConfigStatus) DeepCopy() *AWSAdapterConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AWSAdapterConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountData) DeepCopyInto(out *AccountData) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InspectorEnabledEC2 != nil {
		in, out := &in.InspectorEnabledEC2, &out.InspectorEnabledEC2
		*out = new(bool)
		**out = **in
	}
	if in.InspectorEnabledECR != nil {
		in, out := &in.InspectorEnabledECR, &out.InspectorEnabledECR
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountData.
func (in *AccountData) DeepCopy() *AccountData {
	if in == nil {
		return nil
	}
	out := new(AccountData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EC2LaunchTemplate) DeepCopyInto(out *EC2LaunchTemplate) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EC2LaunchTemplate.
func (in *EC2LaunchTemplate) DeepCopy() *EC2LaunchTemplate {
	if in == nil {
		return nil
	}
	out := new(EC2LaunchTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECRRepository) DeepCopyInto(out *ECRRepository) {
	*out = *in
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
	if in.RepositoryUri != nil {
		in, out := &in.RepositoryUri, &out.RepositoryUri
		*out = new(string)
		**out = **in
	}
	if in.ImageTagMutable != nil {
		in, out := &in.ImageTagMutable, &out.ImageTagMutable
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECRRepository.
func (in *ECRRepository) DeepCopy() *ECRRepository {
	if in == nil {
		return nil
	}
	out := new(ECRRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCluster) DeepCopyInto(out *EKSCluster) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.OIDCProvider != nil {
		in, out := &in.OIDCProvider, &out.OIDCProvider
		*out = new(string)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.PlatformVersion != nil {
		in, out := &in.PlatformVersion, &out.PlatformVersion
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.EncryptionConfig != nil {
		in, out := &in.EncryptionConfig, &out.EncryptionConfig
		*out = make([]*EKSEncryptionConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EKSEncryptionConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Compute != nil {
		in, out := &in.Compute, &out.Compute
		*out = new(EKSCompute)
		(*in).DeepCopyInto(*out)
	}
	if in.Networking != nil {
		in, out := &in.Networking, &out.Networking
		*out = new(EKSNetworking)
		(*in).DeepCopyInto(*out)
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(EKSLogging)
		(*in).DeepCopyInto(*out)
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IdentityProviderConfigs != nil {
		in, out := &in.IdentityProviderConfigs, &out.IdentityProviderConfigs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCluster.
func (in *EKSCluster) DeepCopy() *EKSCluster {
	if in == nil {
		return nil
	}
	out := new(EKSCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCompute) DeepCopyInto(out *EKSCompute) {
	*out = *in
	if in.NodeGroups != nil {
		in, out := &in.NodeGroups, &out.NodeGroups
		*out = make([]*EKSNodeGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EKSNodeGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.FargateProfiles != nil {
		in, out := &in.FargateProfiles, &out.FargateProfiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Reservations != nil {
		in, out := &in.Reservations, &out.Reservations
		*out = make([]*Reservation, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Reservation)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCompute.
func (in *EKSCompute) DeepCopy() *EKSCompute {
	if in == nil {
		return nil
	}
	out := new(EKSCompute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSEncryptionConfig) DeepCopyInto(out *EKSEncryptionConfig) {
	*out = *in
	if in.KeyARN != nil {
		in, out := &in.KeyARN, &out.KeyARN
		*out = new(string)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSEncryptionConfig.
func (in *EKSEncryptionConfig) DeepCopy() *EKSEncryptionConfig {
	if in == nil {
		return nil
	}
	out := new(EKSEncryptionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSLogging) DeepCopyInto(out *EKSLogging) {
	*out = *in
	if in.APIServer != nil {
		in, out := &in.APIServer, &out.APIServer
		*out = new(bool)
		**out = **in
	}
	if in.Audit != nil {
		in, out := &in.Audit, &out.Audit
		*out = new(bool)
		**out = **in
	}
	if in.Authenticator != nil {
		in, out := &in.Authenticator, &out.Authenticator
		*out = new(bool)
		**out = **in
	}
	if in.ControllerManager != nil {
		in, out := &in.ControllerManager, &out.ControllerManager
		*out = new(bool)
		**out = **in
	}
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSLogging.
func (in *EKSLogging) DeepCopy() *EKSLogging {
	if in == nil {
		return nil
	}
	out := new(EKSLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNetworking) DeepCopyInto(out *EKSNetworking) {
	*out = *in
	if in.VPC != nil {
		in, out := &in.VPC, &out.VPC
		*out = new(EKSVpcConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceIPv4CIDR != nil {
		in, out := &in.ServiceIPv4CIDR, &out.ServiceIPv4CIDR
		*out = new(string)
		**out = **in
	}
	if in.ServiceIPv6CIDR != nil {
		in, out := &in.ServiceIPv6CIDR, &out.ServiceIPv6CIDR
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNetworking.
func (in *EKSNetworking) DeepCopy() *EKSNetworking {
	if in == nil {
		return nil
	}
	out := new(EKSNetworking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNodeGroup) DeepCopyInto(out *EKSNodeGroup) {
	*out = *in
	if in.NodegroupArn != nil {
		in, out := &in.NodegroupArn, &out.NodegroupArn
		*out = new(string)
		**out = **in
	}
	if in.NodeRole != nil {
		in, out := &in.NodeRole, &out.NodeRole
		*out = new(string)
		**out = **in
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(int32)
		**out = **in
	}
	if in.AMIReleaseVersion != nil {
		in, out := &in.AMIReleaseVersion, &out.AMIReleaseVersion
		*out = new(string)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InstanceTypes != nil {
		in, out := &in.InstanceTypes, &out.InstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UpdateConfig != nil {
		in, out := &in.UpdateConfig, &out.UpdateConfig
		*out = new(EKSNodeGroupUpdateConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ScalingConfig != nil {
		in, out := &in.ScalingConfig, &out.ScalingConfig
		*out = new(EKSNodeGroupScalingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.LaunchTemplate != nil {
		in, out := &in.LaunchTemplate, &out.LaunchTemplate
		*out = new(EC2LaunchTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.RemoteAccessConfig != nil {
		in, out := &in.RemoteAccessConfig, &out.RemoteAccessConfig
		*out = new(EKSNodeGroupRemoteAccessConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(EKSNodeGroupResources)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthIssues != nil {
		in, out := &in.HealthIssues, &out.HealthIssues
		*out = make([]*EKSNodeGroupHealthIssue, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EKSNodeGroupHealthIssue)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]*EKSNodeGroupTaint, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EKSNodeGroupTaint)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNodeGroup.
func (in *EKSNodeGroup) DeepCopy() *EKSNodeGroup {
	if in == nil {
		return nil
	}
	out := new(EKSNodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNodeGroupHealthIssue) DeepCopyInto(out *EKSNodeGroupHealthIssue) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.ResourceIDs != nil {
		in, out := &in.ResourceIDs, &out.ResourceIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNodeGroupHealthIssue.
func (in *EKSNodeGroupHealthIssue) DeepCopy() *EKSNodeGroupHealthIssue {
	if in == nil {
		return nil
	}
	out := new(EKSNodeGroupHealthIssue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNodeGroupRemoteAccessConfig) DeepCopyInto(out *EKSNodeGroupRemoteAccessConfig) {
	*out = *in
	if in.Ec2SshKey != nil {
		in, out := &in.Ec2SshKey, &out.Ec2SshKey
		*out = new(string)
		**out = **in
	}
	if in.SourceSecurityGroups != nil {
		in, out := &in.SourceSecurityGroups, &out.SourceSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNodeGroupRemoteAccessConfig.
func (in *EKSNodeGroupRemoteAccessConfig) DeepCopy() *EKSNodeGroupRemoteAccessConfig {
	if in == nil {
		return nil
	}
	out := new(EKSNodeGroupRemoteAccessConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNodeGroupResources) DeepCopyInto(out *EKSNodeGroupResources) {
	*out = *in
	if in.AutoScalingGroups != nil {
		in, out := &in.AutoScalingGroups, &out.AutoScalingGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RemoteAccessSecurityGroup != nil {
		in, out := &in.RemoteAccessSecurityGroup, &out.RemoteAccessSecurityGroup
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNodeGroupResources.
func (in *EKSNodeGroupResources) DeepCopy() *EKSNodeGroupResources {
	if in == nil {
		return nil
	}
	out := new(EKSNodeGroupResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNodeGroupScalingConfig) DeepCopyInto(out *EKSNodeGroupScalingConfig) {
	*out = *in
	if in.DesiredSize != nil {
		in, out := &in.DesiredSize, &out.DesiredSize
		*out = new(int32)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(int32)
		**out = **in
	}
	if in.MinSize != nil {
		in, out := &in.MinSize, &out.MinSize
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNodeGroupScalingConfig.
func (in *EKSNodeGroupScalingConfig) DeepCopy() *EKSNodeGroupScalingConfig {
	if in == nil {
		return nil
	}
	out := new(EKSNodeGroupScalingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNodeGroupTaint) DeepCopyInto(out *EKSNodeGroupTaint) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNodeGroupTaint.
func (in *EKSNodeGroupTaint) DeepCopy() *EKSNodeGroupTaint {
	if in == nil {
		return nil
	}
	out := new(EKSNodeGroupTaint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNodeGroupUpdateConfig) DeepCopyInto(out *EKSNodeGroupUpdateConfig) {
	*out = *in
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(int32)
		**out = **in
	}
	if in.MaxUnavailablePercentage != nil {
		in, out := &in.MaxUnavailablePercentage, &out.MaxUnavailablePercentage
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNodeGroupUpdateConfig.
func (in *EKSNodeGroupUpdateConfig) DeepCopy() *EKSNodeGroupUpdateConfig {
	if in == nil {
		return nil
	}
	out := new(EKSNodeGroupUpdateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSVpcConfig) DeepCopyInto(out *EKSVpcConfig) {
	*out = *in
	if in.ClusterSecurityGroupID != nil {
		in, out := &in.ClusterSecurityGroupID, &out.ClusterSecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.PublicAccessCIDRs != nil {
		in, out := &in.PublicAccessCIDRs, &out.PublicAccessCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VpcID != nil {
		in, out := &in.VpcID, &out.VpcID
		*out = new(string)
		**out = **in
	}
	if in.FlowLogsEnabled != nil {
		in, out := &in.FlowLogsEnabled, &out.FlowLogsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSVpcConfig.
func (in *EKSVpcConfig) DeepCopy() *EKSVpcConfig {
	if in == nil {
		return nil
	}
	out := new(EKSVpcConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	if in.HttpPutResponseHopLimit != nil {
		in, out := &in.HttpPutResponseHopLimit, &out.HttpPutResponseHopLimit
		*out = new(int32)
		**out = **in
	}
	if in.PublicDnsName != nil {
		in, out := &in.PublicDnsName, &out.PublicDnsName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastPollInfo) DeepCopyInto(out *LastPollInfo) {
	*out = *in
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = (*in).DeepCopy()
	}
	if in.Failure != nil {
		in, out := &in.Failure, &out.Failure
		*out = new(PollFailure)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastPollInfo.
func (in *LastPollInfo) DeepCopy() *LastPollInfo {
	if in == nil {
		return nil
	}
	out := new(LastPollInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PollFailure) DeepCopyInto(out *PollFailure) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PollFailure.
func (in *PollFailure) DeepCopy() *PollFailure {
	if in == nil {
		return nil
	}
	out := new(PollFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reservation) DeepCopyInto(out *Reservation) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]*Instance, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Instance)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reservation.
func (in *Reservation) DeepCopy() *Reservation {
	if in == nil {
		return nil
	}
	out := new(Reservation)
	in.DeepCopyInto(out)
	return out
}
