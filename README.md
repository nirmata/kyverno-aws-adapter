# kyverno-aws-adapter

## Description
Kyverno AWS Adapter is a Kubernetes controller for the `AWSAdapterConfig` CRD. As of now, it observes the realtime state of an EKS cluster and reconciles it with the currently stored state, but can be further expanded to other AWS services later on by extending the current API with the help of [AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2).

## Installation
You’ll need an [EKS](https://aws.amazon.com/eks/) cluster to run against.

### Running on the EKS cluster
1. Make sure that you have configured an [IAM role for the service account](#IAM-Role-for-Service-Account) `kyverno-aws-adapter-sa` in your desired namespace (configured in `values.yaml`) and specified the role's ARN in the `roleArn` field inside `values.yaml` file.
2. Install the Helm Chart. Follow instructions given [here](/charts/kyverno-aws-adapter#installation).
3. Check the `status` field of the `<cluster-name>-config` custom resource in the namespace specified in `values.yaml`. For instance, if the cluster name is `eks-test` and namespace is `kyverno-aws-adapter`, then:
   ```sh
   kubectl get awsacfg eks-test-config -n kyverno-aws-adapter -o yaml
   ```

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

You can specify the Role's ARN through the `roleArn` setting in the [Helm chart](https://github.com/nirmata/kyverno-aws-adapter/tree/main/charts/kyverno-aws-adapter#installation).

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

## License

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
