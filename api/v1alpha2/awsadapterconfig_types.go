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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AWSAdapterConfigSpec defines the desired state of AWSAdapterConfig
type AWSAdapterConfigSpec struct {
	// EKS cluster's name
	Name *string `json:"name"`
	// EKS cluster's region
	Region *string `json:"region"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AWSAdapterConfig is the Schema for the awsadapterconfigs API
type AWSAdapterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AWSAdapterConfigSpec `json:"spec,omitempty"`
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
