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

package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go/aws"
	securityv1alpha1 "github.com/nirmata/kyverno-aws-adapter/api/v1alpha1"
)

const (
	PollFailure securityv1alpha1.PollStatus = "failure"
	PollSuccess securityv1alpha1.PollStatus = "success"
)

// AWSAdapterConfigReconciler reconciles a AWSAdapterConfig object
type AWSAdapterConfigReconciler struct {
	client.Client
	Scheme          *runtime.Scheme
	RequeueInterval time.Duration
}

//+kubebuilder:rbac:groups=security.nirmata.io,resources=awsadapterconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=security.nirmata.io,resources=awsadapterconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=security.nirmata.io,resources=awsadapterconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *AWSAdapterConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	objOld := &securityv1alpha1.AWSAdapterConfig{}
	err := r.Get(ctx, req.NamespacedName, objOld)
	if err != nil {
		if client.IgnoreNotFound(err) == nil {
			l.Info("Warning: Resource deleted or does not exist")
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if !objOld.DeletionTimestamp.IsZero() {
		l.Info("Warning: Deleting resource. Hope this is done intentionally.")
		return ctrl.Result{}, nil
	}

	if objOld.Status != (securityv1alpha1.AWSAdapterConfigStatus{}) {
		if metav1.Now().Time.Before(objOld.Status.LastPollInfo.Timestamp.Add(r.RequeueInterval)) {
			return ctrl.Result{}, nil
		}
	}
	l.Info("Reconciling", "req", req)

	l.Info("Loading AWS SDK config")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*objOld.Spec.Region))
	if err != nil {
		l.Error(err, "error occurred while loading aws sdk config")
		return r.updateLastPollStatusFailure(ctx, objOld, "error occurred while loading aws sdk config", err, &l, time.Now())
	}
	l.Info("AWS SDK config loaded successfully")
	eksClient := eks.NewFromConfig(cfg)
	ec2Client := ec2.NewFromConfig(cfg)

	objNew := objOld.DeepCopy()
	objNew.Status.EKSCluster = &securityv1alpha1.EKSCluster{}

	clusterFound := false
	if x, err := eksClient.ListClusters(context.TODO(), &eks.ListClustersInput{}); err == nil {
		if x.NextToken != nil {
			l.Info("Warning: more than 100 clusters found in the AWS account, fetching only 100")
		}

		for _, v := range x.Clusters {
			if c, err := eksClient.DescribeCluster(context.TODO(), &eks.DescribeClusterInput{Name: &v}); err == nil {
				if v == *objOld.Spec.Name && strings.ToLower(string(c.Cluster.Status)) != "deleting" {
					clusterFound = true
					break
				}
			} else {
				objOld.Status.LastPollInfo.Status = PollFailure
				l.Error(err, "error occurred while fetching cluster details")
				return r.updateLastPollStatusFailure(ctx, objOld, "error occurred while fetching cluster details", err, &l, time.Now())
			}
		}
	} else {
		objOld.Status.LastPollInfo.Status = PollFailure
		l.Error(err, "error occurred while fetching cluster details")
		return r.updateLastPollStatusFailure(ctx, objOld, "error occurred while fetching cluster details", err, &l, time.Now())
	}

	if !clusterFound {
		l.Error(fmt.Errorf("cluster not found"), fmt.Sprintf("could not find cluster '%s' in the given region '%s'", *objOld.Spec.Name, *objOld.Spec.Region))
		return r.updateLastPollStatusFailure(ctx, objOld, fmt.Sprintf("could not find cluster '%s' in the given region '%s'", *objOld.Spec.Name, *objOld.Spec.Region), fmt.Errorf("cluster not found"), &l, time.Now())
	}

	if x, err := eksClient.DescribeCluster(context.TODO(), &eks.DescribeClusterInput{Name: objOld.Spec.Name}); err == nil {
		tmpEncConf := []*securityv1alpha1.EKSEncryptionConfig{}
		for _, encConf := range x.Cluster.EncryptionConfig {
			tmpEncConf = append(tmpEncConf, &securityv1alpha1.EKSEncryptionConfig{
				KeyARN:    encConf.Provider.KeyArn,
				Resources: encConf.Resources,
			})
		}

		if describeFlowLogsOutput, err := ec2Client.DescribeFlowLogs(context.TODO(), &ec2.DescribeFlowLogsInput{Filter: []types.Filter{
			{
				Name: aws.String("resource-id"),
				Values: []string{
					*x.Cluster.ResourcesVpcConfig.VpcId,
				},
			},
		}}); err == nil {
			objNew.Status.EKSCluster = &securityv1alpha1.EKSCluster{
				CreatedAt:         x.Cluster.CreatedAt.String(),
				Endpoint:          x.Cluster.Endpoint,
				ID:                x.Cluster.Id,
				Name:              x.Cluster.Name,
				PlatformVersion:   x.Cluster.PlatformVersion,
				Region:            objOld.Spec.Region,
				RoleArn:           x.Cluster.RoleArn,
				Status:            string(x.Cluster.Status),
				KubernetesVersion: x.Cluster.Version,
				Arn:               x.Cluster.Arn,
				Certificate:       x.Cluster.CertificateAuthority.Data,
				EncryptionConfig:  tmpEncConf,
				Networking: &securityv1alpha1.EKSNetworking{
					VPC: &securityv1alpha1.EKSVpcConfig{
						ClusterSecurityGroupID: x.Cluster.ResourcesVpcConfig.ClusterSecurityGroupId,
						EndpointPrivateAccess:  x.Cluster.ResourcesVpcConfig.EndpointPrivateAccess,
						EndpointPublicAccess:   x.Cluster.ResourcesVpcConfig.EndpointPublicAccess,
						PublicAccessCIDRs:      x.Cluster.ResourcesVpcConfig.PublicAccessCidrs,
						SecurityGroupIDs:       x.Cluster.ResourcesVpcConfig.SecurityGroupIds,
						SubnetIDs:              x.Cluster.ResourcesVpcConfig.SubnetIds,
						VpcID:                  x.Cluster.ResourcesVpcConfig.VpcId,
						FlowLogsEnabled:        len(describeFlowLogsOutput.FlowLogs) != 0,
					},
					ServiceIPv4CIDR: x.Cluster.KubernetesNetworkConfig.ServiceIpv4Cidr,
					ServiceIPv6CIDR: x.Cluster.KubernetesNetworkConfig.ServiceIpv6Cidr,
					IPFamily:        string(x.Cluster.KubernetesNetworkConfig.IpFamily),
				},
				Compute: &securityv1alpha1.EKSCompute{},
				Logging: &securityv1alpha1.EKSLogging{},
				Tags:    x.Cluster.Tags,
			}
		} else {
			msg := "error occurred while fetching VPC flow logs"
			l.Error(err, msg)
			return r.updateLastPollStatusFailure(ctx, objOld, msg, err, &l, time.Now())
		}

		for _, v := range x.Cluster.Logging.ClusterLogging {
			for _, t := range v.Types {
				switch t {
				case "api":
					objNew.Status.EKSCluster.Logging.APIServer = v.Enabled
				case "audit":
					objNew.Status.EKSCluster.Logging.Audit = v.Enabled
				case "authenticator":
					objNew.Status.EKSCluster.Logging.Authenticator = v.Enabled
				case "controllerManager":
					objNew.Status.EKSCluster.Logging.ControllerManager = v.Enabled
				case "scheduler":
					objNew.Status.EKSCluster.Logging.Scheduler = v.Enabled
				}
			}
		}
	} else {
		l.Error(err, "error fetching cluster details")
		return r.updateLastPollStatusFailure(ctx, objOld, "error fetching cluster details", err, &l, time.Now())
	}

	if x, err := eksClient.ListFargateProfiles(context.TODO(), &eks.ListFargateProfilesInput{ClusterName: objOld.Spec.Name}); err == nil {
		objNew.Status.EKSCluster.Compute.FargateProfiles = x.FargateProfileNames
	} else {
		l.Error(err, "error listing fargate profiles")
		return r.updateLastPollStatusFailure(ctx, objOld, "error listing fargate profiles", err, &l, time.Now())
	}

	if x, err := eksClient.ListNodegroups(context.TODO(), &eks.ListNodegroupsInput{ClusterName: objOld.Spec.Name}); err == nil {
		for _, v := range x.Nodegroups {
			if y, err := eksClient.DescribeNodegroup(context.TODO(), &eks.DescribeNodegroupInput{ClusterName: objOld.Spec.Name, NodegroupName: &v}); err == nil {
				objNew.Status.EKSCluster.Compute.NodeGroups = []*securityv1alpha1.EKSNodeGroup{}
				var launchTemplate *securityv1alpha1.EC2LaunchTemplate
				if y.Nodegroup.LaunchTemplate != nil {
					launchTemplate = &securityv1alpha1.EC2LaunchTemplate{
						ID:      y.Nodegroup.LaunchTemplate.Id,
						Name:    y.Nodegroup.LaunchTemplate.Name,
						Version: y.Nodegroup.LaunchTemplate.Version,
					}
				}

				healthIssues := []*securityv1alpha1.EKSNodeGroupHealthIssue{}
				for _, issue := range y.Nodegroup.Health.Issues {
					healthIssues = append(healthIssues, &securityv1alpha1.EKSNodeGroupHealthIssue{
						Code:        string(issue.Code),
						Message:     issue.Message,
						ResourceIDs: issue.ResourceIds,
					})

				}

				autoScalingGroups := []string{}
				for _, asg := range y.Nodegroup.Resources.AutoScalingGroups {
					autoScalingGroups = append(autoScalingGroups, *asg.Name)
				}

				taints := []*securityv1alpha1.EKSNodeGroupTaint{}
				for _, taint := range y.Nodegroup.Taints {
					taints = append(taints, &securityv1alpha1.EKSNodeGroupTaint{
						Effect: string(taint.Effect),
						Key:    taint.Key,
						Value:  taint.Value,
					})
				}

				var remoteAccessConfig *securityv1alpha1.EKSNodeGroupRemoteAccessConfig
				if y.Nodegroup.RemoteAccess != nil {
					remoteAccessConfig = &securityv1alpha1.EKSNodeGroupRemoteAccessConfig{
						Ec2SshKey:            y.Nodegroup.RemoteAccess.Ec2SshKey,
						SourceSecurityGroups: y.Nodegroup.RemoteAccess.SourceSecurityGroups,
					}
				}

				objNew.Status.EKSCluster.Compute.NodeGroups = append(objNew.Status.EKSCluster.Compute.NodeGroups, &securityv1alpha1.EKSNodeGroup{
					Name: v,
					ScalingConfig: &securityv1alpha1.EKSNodeGroupScalingConfig{
						DesiredSize: y.Nodegroup.ScalingConfig.DesiredSize,
						MinSize:     y.Nodegroup.ScalingConfig.MinSize,
						MaxSize:     y.Nodegroup.ScalingConfig.MaxSize,
					},
					LaunchTemplate:     launchTemplate,
					Status:             string(y.Nodegroup.Status),
					AMIReleaseVersion:  y.Nodegroup.ReleaseVersion,
					HealthIssues:       healthIssues,
					AMIType:            string(y.Nodegroup.AmiType),
					CapacityType:       string(y.Nodegroup.CapacityType),
					CreatedAt:          y.Nodegroup.CreatedAt.String(),
					DiskSize:           y.Nodegroup.DiskSize,
					InstanceTypes:      y.Nodegroup.InstanceTypes,
					NodegroupArn:       y.Nodegroup.NodegroupArn,
					NodeRole:           y.Nodegroup.NodeRole,
					RemoteAccessConfig: remoteAccessConfig,
					Resources: &securityv1alpha1.EKSNodeGroupResources{
						RemoteAccessSecurityGroup: y.Nodegroup.Resources.RemoteAccessSecurityGroup,
						AutoScalingGroups:         autoScalingGroups,
					},
					Subnets: y.Nodegroup.Subnets,
					Tags:    y.Nodegroup.Tags,
					Taints:  taints,
					Labels:  y.Nodegroup.Labels,
					UpdateConfig: &securityv1alpha1.EKSNodeGroupUpdateConfig{
						MaxUnavailable:           y.Nodegroup.UpdateConfig.MaxUnavailable,
						MaxUnavailablePercentage: y.Nodegroup.UpdateConfig.MaxUnavailablePercentage,
					},
				})
			} else {
				l.Error(err, fmt.Sprintf("error describing nodegroup '%s'", v))
				return r.updateLastPollStatusFailure(ctx, objOld, fmt.Sprintf("error describing nodegroup '%s'", v), err, &l, time.Now())
			}
		}
	} else {
		l.Error(err, "error listing nodegroups", *objOld.Spec.Name, *objOld.Spec.Region)
		return r.updateLastPollStatusFailure(ctx, objOld, "error listing nodegroups", err, &l, time.Now())
	}

	if x, err := eksClient.ListAddons(context.TODO(), &eks.ListAddonsInput{ClusterName: objOld.Spec.Name}); err == nil {
		objNew.Status.EKSCluster.Addons = x.Addons
	} else {
		l.Error(err, "error listing addons")
		return r.updateLastPollStatusFailure(ctx, objOld, "error listing addons", err, &l, time.Now())
	}

	if x, err := eksClient.ListIdentityProviderConfigs(context.TODO(), &eks.ListIdentityProviderConfigsInput{ClusterName: objOld.Spec.Name}); err == nil {
		objNew.Status.EKSCluster.IdentityProviderConfigs = []*string{}
		for _, v := range x.IdentityProviderConfigs {
			objNew.Status.EKSCluster.IdentityProviderConfigs = append(objNew.Status.EKSCluster.IdentityProviderConfigs, v.Name)
		}
	} else {
		l.Error(err, "error listing identity provider configs")
		return r.updateLastPollStatusFailure(ctx, objOld, "error listing identity provider configs", err, &l, time.Now())
	}

	if x, err := ec2Client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name: aws.String("tag:aws:eks:cluster-name"),
				Values: []string{
					*objOld.Spec.Name,
				},
			},
		},
	},
	); err == nil {
		for _, r := range x.Reservations {
			tmpRes := []*securityv1alpha1.Reservation{}
			for _, i := range r.Instances {
				tmpIn := []*securityv1alpha1.Instance{}
				tmpIn = append(tmpIn, &securityv1alpha1.Instance{
					PublicDnsName:           i.PublicDnsName,
					HttpPutResponseHopLimit: i.MetadataOptions.HttpPutResponseHopLimit,
				})
				tmpRes = append(tmpRes, &securityv1alpha1.Reservation{
					Instances: tmpIn,
				})
			}
			objNew.Status.EKSCluster.Compute.Reservations = tmpRes
		}
	} else {
		l.Error(err, "error occurred while fetching EC2 instances")
		return r.updateLastPollStatusFailure(ctx, objOld, "error occurred while fetching EC2 instances", err, &l, time.Now())
	}

	currentPollTimestamp := time.Now()
	objNew.Status.LastPollInfo = securityv1alpha1.LastPollInfo{
		Timestamp: &metav1.Time{Time: currentPollTimestamp},
		Status:    PollSuccess,
	}

	if !cmp.Equal(objNew.Status.EKSCluster, objOld.Status.EKSCluster) {
		objNew.Status.LastUpdatedTimestamp = &metav1.Time{Time: currentPollTimestamp}
		if err := r.Status().Update(ctx, objNew); err != nil {
			l.Error(err, "error updating status")
		}
	}

	return ctrl.Result{RequeueAfter: r.RequeueInterval}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AWSAdapterConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&securityv1alpha1.AWSAdapterConfig{}).
		Complete(r)
}

func (r *AWSAdapterConfigReconciler) updateLastPollStatusFailure(ctx context.Context, objOld *securityv1alpha1.AWSAdapterConfig, msg string, err error, l *logr.Logger, currentPollTimestamp time.Time) (ctrl.Result, error) {
	objOld.Status.LastPollInfo.Status = PollFailure
	objOld.Status.LastPollInfo.Timestamp = &metav1.Time{Time: currentPollTimestamp}
	objOld.Status.LastPollInfo.Failure = &securityv1alpha1.PollFailure{
		Message: msg,
		Error:   err.Error(),
	}

	if err := r.Status().Update(ctx, objOld); err != nil {
		l.Error(err, "error updating last poll's status failure")
	}

	return ctrl.Result{RequeueAfter: r.RequeueInterval}, nil
}
