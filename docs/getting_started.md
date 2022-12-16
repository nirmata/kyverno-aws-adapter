# Getting Started

This is a guide on how to get started with the Nirmata Kyverno Adapter for AWS. To learn more about Kyverno, check out the official documentation.

## Prerequisites
- a running EKS Cluster (refer [Creating an Amazon EKS Cluster](https://docs.aws.amazon.com/eks/latest/userguide/create-cluster.html))
- AWS CLI (refer [Installing the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html))
- eksctl (refer [Installing eksctl](https://docs.aws.amazon.com/eks/latest/userguide/eksctl.html))
- Helm CLI (refer [Installing Helm](https://helm.sh/docs/intro/install/))

## Installing the Nirmata Kyverno Adapter for AWS
There are a few steps we need to follow before installing the Helm chart.

### Creating the IAM Policy
To fetch the EKS Cluster configuration, the AWS Adapter needs the below permissions that has to be expressed via an IAM Policy

```bash
cat >my-policy.json <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "eks:AccessKubernetesApi",
                  "eks:DescribeAddon",
                  "eks:DescribeAddonVersions",
                  "eks:DescribeCluster",
                  "eks:DescribeFargateProfile",
                  "eks:DescribeIdentityProviderConfig",
                  "eks:DescribeNodegroup",
                  "eks:DescribeUpdate",
                  "eks:ListAddons",
                  "eks:ListClusters",
                  "eks:ListFargateProfiles",
                  "eks:ListIdentityProviderConfigs",
                  "eks:ListNodegroups",
                "eks:ListTagsForResource",
                "eks:ListUpdates"
            ],
            "Resource": [
                "arn:aws:eks:*:111122223333:identityproviderconfig/*/*/*/*",
                "arn:aws:eks:*:111122223333:fargateprofile/*/*/*",
                "arn:aws:eks:*:111122223333:nodegroup/*/*/*",
                "arn:aws:eks:*:111122223333:cluster/*",
                "arn:aws:eks:*:111122223333:addon/*/*/*"
            ]
        }
    ]
}
EOF
aws iam create-policy --policy-name kyverno-aws-adapter-policy --policy-document file://my-policy.json
```

**Note:** Make sure your AWS CLI is configured correctly. Follow the [official guide](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html) for setting it up


### Creating IAM OIDC provider
An IAM Open ID Connect provider for the cluster is required to provide the reference in the IAM Role.

```bash
eksctl utils associate-iam-oidc-provider --cluster <cluster-name> --approve --region <region-code>
```

### Creating the IAM Role
Create an IAM Role that references the policy we created earlier.

**Note:** We will use `eksctl` in this guide to create the IAM Role. The serviceaccount creation will be done by the Helm chart so that appropriate RBAC is assigned by the chart itself. If you wish to use your own serviceaccount, then make sure all the necessary rolebinding and clusterrolebinding are present.

```bash
eksctl create iamserviceaccount --name nirmata-aws-adapter  --namespace nirmata-aws-adapter --cluster <cluster-name> --role-name nirmata-adapter-role --attach-policy-arn arn:aws:iam::111122223333:policy/kyverno-aws-adapter-policy   --role-only  --approve
```

This will create a new IAM Role that references the policy we created above. This also creates the trust-relationship for us. If you wish to create the IAM Role via the management console or the AWS CLI, make sure to create the trust-relationship so that the AWS Adapter can assume this Role to fetch EKS Cluster details.


### Installing the AWS Adapter Helm chart
First we need to set the values that are needed to install the Helm chart. You can either pass them as arguments via the [command line](https://helm.sh/docs/helm/helm_install/#helm-install) or set them in a [values file](https://helm.sh/docs/chart_template_guide/values_files/)

As an example, here are the minimum values that you need to set
```bash
# cat myvalues.yaml
eksCluster:
  name: cluster-name
  region: cluster-region

roleArn: arn:aws:iam::111122223333:role/nirmata-adapter-role

rbac:
  create: true
  serviceAccount:
    name: nirmata-aws-adapter
```

**Note:** Update the myvalues.yaml file as per information specific to your account and cluster

We will install the adapter in nirmata-aws-adapter namespace. Create the namespace using ```kubectl create namespace nirmata-aws-adapter```

Now let's install the Helm chart

```bash
helm repo add nirmata-kyverno-aws-adapter https://nirmata.github.io/kyverno-aws-adapter/
helm repo update nirmata-kyverno-aws-adapter

helm install kyverno-aws-adapter nirmata-kyverno-aws-adapter/kyverno-aws-adapter -f myvalues.yaml --devel --namespace nirmata-aws-adapter
```

**Note:** We are using --devel argument in the above helm install command to include pre-releases. We should be using only stable releases in production. We will drop this command from the guide once we have a stable version released.

### Verifying the AWS Adapter installation
Perform the below steps to see if the adapter is installed correctly
```bash
# check if the controller pod is running
kubectl get pods -n nirmata-aws-adapter
NAME                                              READY   STATUS    RESTARTS   AGE
aws-adapter-kyverno-aws-adapter-997f45bb9-c2z5j   1/1     Running   0          109m

# check the awsadapterconfig status
kubectl get awsacfg -n nirmata-aws-adapter
NAME                             ...   CLUSTER NAME     ...  STATUS
aws-adapter-kyverno-aws-adapter  ...   cluster-name     ...  success
```