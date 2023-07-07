# nirmata-aws-adapter

## Description
Nirmata AWS Adapter is a Kubernetes controller for the `AWSConfig` CRD. As of now, it observes the realtime state of an EKS cluster and reconciles it with the currently stored state, but can be further expanded to other AWS services later on by extending the current API with the help of [AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2).

## Installation
You’ll need an [EKS](https://aws.amazon.com/eks/) cluster to run against.

### Running on the EKS cluster
1. Make sure that you have configured an [IAM role for the service account](#IAM-Role-for-Service-Account) `nirmata-aws-adapter-sa` in your desired namespace (configured in `values.yaml`) and specified the role's ARN in the `roleArn` field inside `values.yaml` file.
2. Install the Helm chart after making any necessary changes to `charts/aws-adapter/values.yaml`
   ```sh
   helm install nirmata-aws-adapter charts/aws-adapter
   ```
3. Check the `status` field of the `<cluster-name>-config` custom resource in the namespace specified in `values.yaml`. For instance, if the cluster name is `eks-test` and namespace is `nirmata`, then:
   ```sh
   kubectl get awsconfig eks-test-config -n nirmata -o yaml 
   ```

## Helm Values
Currently supported values for the Helm chart are as follows:
| Value | Description |
-- | ---
| `namespace` | Namespace for installing the controller and CRD |
| `eksCluster` | Configuration for EKS cluster's `name` and `region` |
| `registryConfig` | ghcr.io `username` and `password` configuration for the image secret |
| `pollInterval` | Interval for controller reconciliation |
| `image` | Configuration for image `name`, `tag` and `pullPolicy` |
| `roleArn` | IAM Role ARN with required permissions for the EKS cluster |
| `nameOverride` | Override the chart name |
| `fullnameOverride` | Override the entire generated name |


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

You can specify the Role's ARN in the `roleArn` field inside the Helm chart's `values.yaml` file.

Please ensure that the trust relationship policy for your IAM role resembles the following format:
```json
{
  "Version": "YYYY-MM-DD",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::$account_id:oidc-provider/$oidc_provider"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "$oidc_provider:aud": "sts.amazonaws.com",
          "$oidc_provider:sub": "system:serviceaccount:$namespace:$service_account"
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

