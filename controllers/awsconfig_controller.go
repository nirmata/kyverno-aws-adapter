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

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	securityv1alpha1 "github.com/nirmata/kyverno-aws-adapter/api/v1alpha1"
)

// AWSConfigReconciler reconciles a AWSConfig object
type AWSConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=security.nirmata.io,resources=awsconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=security.nirmata.io,resources=awsconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=security.nirmata.io,resources=awsconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AWSConfig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *AWSConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	l.Info("Reconciling", "req", req)

	objOld := &securityv1alpha1.AWSConfig{}
	err := r.Get(ctx, types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, objOld)
	if err != nil {
		l.Error(err, "error occurred while retrieving awsconfig")
		return ctrl.Result{}, nil
	}
	objNew := objOld.DeepCopy()
	objNew.Status.EKSCluster = &securityv1alpha1.EKSCluster{}

	l.Info("Loading AWS SDK config")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-1"))
	if err != nil {
		l.Error(err, "error occurred while loading aws sdk config")
		return ctrl.Result{RequeueAfter: time.Duration(10) * time.Second}, nil
	}
	l.Info("AWS SDK config loaded successfully")
	svc := eks.NewFromConfig(cfg)

	// TODO: list all clusters instead of going through just first 100
	clusterFound := false
	if x, err := svc.ListClusters(context.TODO(), &eks.ListClustersInput{}); err == nil {
		for _, v := range x.Clusters {
			if c, err := svc.DescribeCluster(context.TODO(), &eks.DescribeClusterInput{Name: &v}); err == nil {
				if v == *objOld.Spec.Name && strings.ToLower(string(c.Cluster.Status)) != "deleting" {
					clusterFound = true
					break
				}
			}
		}
	}

	if !clusterFound {
		l.Error(fmt.Errorf("cluster not found"), fmt.Sprintf("could not find cluster '%s' in the given region '%s'", *objOld.Spec.Name, *objOld.Spec.Region))
		return ctrl.Result{RequeueAfter: time.Duration(10) * time.Second}, nil
	}

	if x, err := svc.DescribeCluster(context.TODO(), &eks.DescribeClusterInput{Name: objOld.Spec.Name}); err == nil {
		tmpEncConf := []*securityv1alpha1.EKSEncryptionConfig{}
		for _, encConf := range x.Cluster.EncryptionConfig {
			tmpEncConf = append(tmpEncConf, &securityv1alpha1.EKSEncryptionConfig{
				KeyARN:    encConf.Provider.KeyArn,
				Resources: encConf.Resources,
			})
		}

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
				},
				ServiceIPv4CIDR: x.Cluster.KubernetesNetworkConfig.ServiceIpv4Cidr,
				ServiceIPv6CIDR: x.Cluster.KubernetesNetworkConfig.ServiceIpv6Cidr,
				IPFamily:        string(x.Cluster.KubernetesNetworkConfig.IpFamily),
			},
			Compute: &securityv1alpha1.EKSCompute{},
			Logging: &securityv1alpha1.EKSLogging{},
			Tags:    x.Cluster.Tags,
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
	}

	if x, err := svc.ListFargateProfiles(context.TODO(), &eks.ListFargateProfilesInput{ClusterName: objOld.Spec.Name}); err == nil {
		objNew.Status.EKSCluster.Compute.FargateProfiles = x.FargateProfileNames
	} else {
		l.Error(err, "error listing fargate profiles")
	}

	if x, err := svc.ListNodegroups(context.TODO(), &eks.ListNodegroupsInput{ClusterName: objOld.Spec.Name}); err == nil {
		for _, v := range x.Nodegroups {
			if y, err := svc.DescribeNodegroup(context.TODO(), &eks.DescribeNodegroupInput{ClusterName: objOld.Spec.Name, NodegroupName: &v}); err == nil {
				objNew.Status.EKSCluster.Compute.NodeGroups = []*securityv1alpha1.EKSNodeGroup{}
				launchTemplate := &securityv1alpha1.EC2LaunchTemplate{}
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

				remoteAccessConfig := &securityv1alpha1.EKSNodeGroupRemoteAccessConfig{}
				if y.Nodegroup.RemoteAccess != nil {
					remoteAccessConfig.Ec2SshKey = y.Nodegroup.RemoteAccess.Ec2SshKey
					remoteAccessConfig.SourceSecurityGroups = y.Nodegroup.RemoteAccess.SourceSecurityGroups
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
			}
		}
	} else {
		l.Error(err, "error listing nodegroups")
	}

	if x, err := svc.ListAddons(context.TODO(), &eks.ListAddonsInput{ClusterName: objOld.Spec.Name}); err == nil {
		objNew.Status.EKSCluster.Addons = x.Addons
	} else {
		l.Error(err, "error listing addons")
	}

	if x, err := svc.ListIdentityProviderConfigs(context.TODO(), &eks.ListIdentityProviderConfigsInput{ClusterName: objOld.Spec.Name}); err == nil {
		for _, v := range x.IdentityProviderConfigs {
			objNew.Status.EKSCluster.IdentityProviderConfigs = []*string{}
			objNew.Status.EKSCluster.IdentityProviderConfigs = append(objNew.Status.EKSCluster.IdentityProviderConfigs, v.Name)
		}
	} else {
		l.Error(err, "error listing identity provider configs")
	}

	if !cmp.Equal(objNew.Status, objOld.Status) {
		if err := r.Status().Update(ctx, objNew); err != nil {
			l.Error(err, "error updating status")
		}
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AWSConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&securityv1alpha1.AWSConfig{}).
		Complete(r)
}
