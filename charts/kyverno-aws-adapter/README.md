# kyverno-aws-adapter

Helm chart for the Kyverno AWS Adapter

![Version: v0.0.1](https://img.shields.io/badge/Version-v0.0.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: v0.0.1](https://img.shields.io/badge/AppVersion-v0.0.1-informational?style=flat-square)

## Description

Kyverno AWS Adapter is a Kubernetes controller for the `AWSAdapterConfig` CRD. As of now, it observes the realtime state of an EKS cluster and reconciles it with the currently stored state, but can be further expanded to other AWS services later on by extending the current API with the help of [AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2).

## Installation

You’ll need an [EKS](https://aws.amazon.com/eks/) cluster to run against.

### Running on the EKS cluster

1. Make sure that you have configured an [IAM role for the service account](#IAM-Role-for-Service-Account) `kyverno-aws-adapter-sa` in your desired namespace (configured in `values.yaml`) and specified the role's ARN in the `roleArn` field inside `values.yaml` file.
2. Install the Helm chart after making any necessary changes to `charts/kyverno-aws-adapter/values.yaml`
   ```sh
   helm install kyverno-aws-adapter charts/kyverno-aws-adapter
   ```
3. Check the `status` field of the `<cluster-name>-config` custom resource in the namespace specified in `values.yaml`. For instance, if the cluster name is `eks-test` and namespace is `kyverno-aws-adapter`, then:
   ```sh
   kubectl get awsacfg eks-test-config -n kyverno-aws-adapter -o yaml
   ```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| nameOverride | string | `nil` | Override the name of the chart |
| fullnameOverride | string | `nil` | Override the expanded name of the chart |
| roleArn | string | `nil` | Role for accessing AWS API (REQUIRED) |
| pollInterval | int | `30` | Interval at which the controller reconciles in minutes |
| eksCluster.name | string | `nil` | EKS cluster name |
| eksCluster.region | string | `nil` | EKS cluster region |
| registryConfig.username | string | `nil` | Username to pull the private image (ghcr.io) |
| registryConfig.password | string | `nil` | Password to pull the private image (ghcr.io) |
| rbac.create | bool | `true` | Enable RBAC resources creation |
| rbac.serviceAccount.name | string | `nil` | Service account name, you MUST provide one when `rbac.create` is set to `false` |
| image.repository | string | `"ghcr.io/nirmata/kyverno-aws-adapter"` | Image repository |
| image.pullPolicy | string | `"Always"` | Image pull policy |
| image.tag | string | `nil` | Image tag (defaults to chart app version) |

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

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| Nirmata |  | <https://nirmata.com/> |

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

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
