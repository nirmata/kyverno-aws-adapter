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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PollStatus string

// AWSAdapterConfigSpec defines the desired state of AWSAdapterConfig
type AWSAdapterConfigSpec struct {
	// EKS cluster's name
	Name *string `json:"name"`
	// EKS cluster's region
	Region *string `json:"region"`
}

// EKSCluster contains the EKS cluster's details
type EKSCluster struct {
	ID                      *string                `json:"id,omitempty"`
	KubernetesVersion       *string                `json:"kubernetesVersion,omitempty"`
	Name                    *string                `json:"name"`
	Status                  string                 `json:"status"`
	Region                  *string                `json:"region"`
	Endpoint                *string                `json:"endpoint,omitempty"`
	OIDCProvider            *string                `json:"oidcProvider,omitempty"`
	Certificate             *string                `json:"certificate,omitempty"`
	Arn                     *string                `json:"arn,omitempty"`
	PlatformVersion         *string                `json:"platformVersion,omitempty"`
	RoleArn                 *string                `json:"roleArn,omitempty"`
	CreatedAt               string                 `json:"createdAt,omitempty"`
	EncryptionConfig        []*EKSEncryptionConfig `json:"encryptionConfig,omitempty"`
	Compute                 *EKSCompute            `json:"compute,omitempty"`
	Networking              *EKSNetworking         `json:"networking,omitempty"`
	Logging                 *EKSLogging            `json:"logging,omitempty"`
	Addons                  []string               `json:"addons,omitempty"`
	IdentityProviderConfigs []*string              `json:"identityProviderConfigs,omitempty"`
	Tags                    map[string]string      `json:"tags,omitempty"`
}

// EKSEncryptionConfig contains encryption configuration of the EKS cluster
type EKSEncryptionConfig struct {
	KeyARN    *string  `json:"keyARN,omitempty"`
	Resources []string `json:"resources,omitempty"`
}

// EKSCompute contains node groups and fargate profiles of the EKS cluster
type EKSCompute struct {
	NodeGroups      []*EKSNodeGroup `json:"nodeGroups,omitempty"`
	FargateProfiles []string        `json:"fargateProfiles,omitempty"`
	Reservations    []*Reservation  `json:"reservations,omitempty"`
}

// EKSNodeGroup contains info of the EKS cluster's node group
type EKSNodeGroup struct {
	Name               string                          `json:"name,omitempty"`
	NodegroupArn       *string                         `json:"nodeGroupArn,omitempty"`
	NodeRole           *string                         `json:"nodeRole,omitempty"`
	CreatedAt          string                          `json:"createdAt,omitempty"`
	Status             string                          `json:"status,omitempty"`
	DiskSize           *int32                          `json:"diskSize,omitempty"`
	AMIType            string                          `json:"amiType,omitempty"`
	CapacityType       string                          `json:"capacityType,omitempty"`
	AMIReleaseVersion  *string                         `json:"amiReleaseVersion,omitempty"`
	Subnets            []string                        `json:"subnets,omitempty"`
	InstanceTypes      []string                        `json:"instanceTypes,omitempty"`
	UpdateConfig       *EKSNodeGroupUpdateConfig       `json:"updateConfig,omitempty"`
	ScalingConfig      *EKSNodeGroupScalingConfig      `json:"scalingConfig,omitempty"`
	LaunchTemplate     *EC2LaunchTemplate              `json:"launchTemplate,omitempty"`
	RemoteAccessConfig *EKSNodeGroupRemoteAccessConfig `json:"remoteAccessConfig,omitempty"`
	Resources          *EKSNodeGroupResources          `json:"resources,omitempty"`
	HealthIssues       []*EKSNodeGroupHealthIssue      `json:"healthIssues,omitempty"`
	Taints             []*EKSNodeGroupTaint            `json:"taints,omitempty"`
	Labels             map[string]string               `json:"labels,omitempty"`
	Tags               map[string]string               `json:"tags,omitempty"`
}

// EKSNodeGroupUpdateConfig contains number/percentage of node groups that can be updated in parallel
type EKSNodeGroupUpdateConfig struct {
	MaxUnavailable           *int32 `json:"maxUnavailable,omitempty"`
	MaxUnavailablePercentage *int32 `json:"maxUnavailablePercentage,omitempty"`
}

type Reservation struct {
	Instances []*Instance `json:"instances,omitempty"`
}

type Instance struct {
	HttpPutResponseHopLimit *int32  `json:"httpPutResponseHopLimit,omitempty"`
	PublicDnsName           *string `json:"publicDnsName,omitempty"`
}

// EKSNodeGroupResources contains info of ASG and remote access SG for node group
type EKSNodeGroupResources struct {
	AutoScalingGroups         []string `json:"autoScalingGroups,omitempty"`
	RemoteAccessSecurityGroup *string  `json:"remoteAccessSecurityGroup,omitempty"`
}

// EKSNodeGroupTaint contains info of taints in the EKS cluster's node group
type EKSNodeGroupTaint struct {
	Effect string  `json:"effect,omitempty"`
	Key    *string `json:"key,omitempty"`
	Value  *string `json:"value,omitempty"`
}

// EKSNodeGroupRemoteAccessConfig contains remote access configuration of the EKS cluster's node group
type EKSNodeGroupRemoteAccessConfig struct {
	Ec2SshKey            *string  `json:"ec2SSHKey,omitempty"`
	SourceSecurityGroups []string `json:"sourceSecurityGroups,omitempty"`
}

// EKSNodeGroupHealthIssue contains info of any health issue in the EKS cluster's node group
type EKSNodeGroupHealthIssue struct {
	Code        string   `json:"code,omitempty"`
	Message     *string  `json:"message,omitempty"`
	ResourceIDs []string `json:"resourceIDs,omitempty"`
}

// EKSNodeGroupScalingConfig contains scaling configuration of  the EKS cluster's node group
type EKSNodeGroupScalingConfig struct {
	DesiredSize *int32 `json:"desiredSize,omitempty"`
	MaxSize     *int32 `json:"maxSize,omitempty"`
	MinSize     *int32 `json:"minSize,omitempty"`
}

// EC2LaunchTemplate contains launch template info the EKS cluster's node group
type EC2LaunchTemplate struct {
	ID      *string `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// EKSVpcConfig contains VPC configuration of the EKS cluster
type EKSVpcConfig struct {
	ClusterSecurityGroupID *string  `json:"clusterSecurityGroupID,omitempty"`
	EndpointPrivateAccess  bool     `json:"endpointPrivateAccess,omitempty"`
	EndpointPublicAccess   bool     `json:"endpointPublicAccess,omitempty"`
	PublicAccessCIDRs      []string `json:"publicAccessCIDRs,omitempty"`
	SecurityGroupIDs       []string `json:"securityGroupIDs,omitempty"`
	SubnetIDs              []string `json:"subnetIDs,omitempty"`
	VpcID                  *string  `json:"vpcID,omitempty"`
	FlowLogsEnabled        bool     `json:"flowLogsEnabled,omitempty"`
}

// EKSNetworking contains networking configuration of the EKS cluster
type EKSNetworking struct {
	VPC             *EKSVpcConfig `json:"vpc,omitempty"`
	IPFamily        string        `json:"ipFamily,omitempty"`
	ServiceIPv4CIDR *string       `json:"serviceIPv4CIDR,omitempty"`
	ServiceIPv6CIDR *string       `json:"serviceIPv6CIDR,omitempty"`
}

// EKSLogging contains info of which logs are enabled
type EKSLogging struct {
	APIServer         *bool `json:"apiServer,omitempty"`
	Audit             *bool `json:"audit,omitempty"`
	Authenticator     *bool `json:"authenticator,omitempty"`
	ControllerManager *bool `json:"controllerManager,omitempty"`
	Scheduler         *bool `json:"scheduler,omitempty"`
}

// PollFailure contains the Error and relevant Message if got Failure in last poll
type PollFailure struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// LastPollInfo contains Timestamp, Status and Failure info of last poll
type LastPollInfo struct {
	Timestamp *metav1.Time `json:"timestamp,omitempty"`
	Status    PollStatus   `json:"status,omitempty"`
	Failure   *PollFailure `json:"failure,omitempty"`
}

// AWSAdapterConfigStatus defines the observed state of AWSAdapterConfig
type AWSAdapterConfigStatus struct {
	// Timestamp when the Status was last updated
	LastUpdatedTimestamp *metav1.Time `json:"lastUpdatedTimestamp,omitempty"`
	// Information on when the adapter last tried to fetch the EKS cluster details
	LastPollInfo LastPollInfo `json:"lastPollInfo"`
	// EKS cluster details fetched from AWS
	// For details of individual fields, refer to AWS SDK docs:
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/eks@v1.22.1/types#Cluster
	EKSCluster *EKSCluster `json:"eksCluster,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:shortName="awsacfg"
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Cluster Name",type=string,JSONPath=`.spec.name`
//+kubebuilder:printcolumn:name="Region",type=string,JSONPath=`.spec.region`
//+kubebuilder:printcolumn:name="Cluster Status",type=string,JSONPath=`.status.eksCluster.status`
//+kubebuilder:printcolumn:name="Kubernetes Version",type=string,JSONPath=`.status.eksCluster.kubernetesVersion`
//+kubebuilder:printcolumn:name="Last Polled Status",type=string,JSONPath=`.status.lastPollInfo.status`

// AWSAdapterConfig is the Schema for the awsadapterconfigs API
type AWSAdapterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AWSAdapterConfigSpec   `json:"spec,omitempty"`
	Status AWSAdapterConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AWSAdapterConfigList contains a list of AWSAdapterConfig
type AWSAdapterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSAdapterConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AWSAdapterConfig{}, &AWSAdapterConfigList{})
}
