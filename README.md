# kyverno-aws-adapter

## Description
Kyverno AWS Adapter is a Kubernetes controller for the `AWSAdapterConfig` CRD. As of now, it observes the realtime state of an EKS cluster and reconciles it with the currently stored state, but can be further expanded to other AWS services later on by extending the current API with the help of [AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2).

## Getting Started
Check out the [getting_started.md](docs/getting_started.md) guide for installing the Nirmata Kyverno Adapter for AWS.


## Local Dev Installation
### Prerequisites
You’ll need an [EKS](https://aws.amazon.com/eks/) cluster to run against.

### Running on the EKS cluster
1. Make sure that you have configured an [IAM role for the service account](#IAM-Role-for-Service-Account) to be used by the Kyverno AWS Adapter.

2. Install the Helm Chart and verify that the Adapter works as expected. Follow instructions given [here](/charts/kyverno-aws-adapter#installation)

## IAM Role for Service Account
This adapter utilizes the ARN of a user-defined IAM Role associated with any policy that has `Full: List, Read` permissions for the `EKS` service, including the following:

| Permission |
| --- |
| ListAddons |
| ListClusters |
| ListFargateProfiles |
| ListIdentityProviderConfigs |
| ListNodeGroups |
| ListUpdates |
| AccessKubernetesApi |
| DescribeAddon |
| DescribeAddonVersions |
| DescribeCluster |
| DescribeFargateProfile |
| DescribeIdentityProviderConfig |
| DescribeNodegroup |
| DescribeUpdate |
| ListTagsForResource |

You can specify the Role's ARN through the `roleArn` setting in the [Helm chart](/charts/kyverno-aws-adapter#installation).

Please ensure that the trust relationship policy for your IAM role resembles the following format:
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::<account_id>:oidc-provider/oidc.eks.<region>.amazonaws.com/id/<oidc_provider_id>"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "oidc.eks.<region>.amazonaws.com/id/<oidc_provider_id>:aud": "sts.amazonaws.com",
          "oidc.eks.<region>.amazonaws.com/id/<oidc_provider_id>:sub": "system:serviceaccount:$namespace:<service_account>"
        }
      }
    }
  ]
}
```

For detailed instructions on how to configure the IAM role for service account, check out the official AWS documentation on [IAM roles for service accounts](https://docs.aws.amazon.com/eks/latest/userguide/associate-service-account-role.html).
