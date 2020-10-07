// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
    cpv1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RepositorySpecParams defines the desired state of Repository
type RepositorySpecParams struct {
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`
	ImageScanningConfiguration *ImageScanningConfiguration `json:"imageScanningConfiguration,omitempty"`
	ImageTagMutability *string `json:"imageTagMutability,omitempty"`
	RepositoryName *string `json:"repositoryName,omitempty"`
	Tags []*Tag `json:"tags,omitempty"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
    cpv1alpha1.ResourceSpec `json:",inline"`
    ForProvider RepositorySpecParams `json:"forProvider"`
}

// RepositoryExternalStatus defines the observed state of Repository
type RepositoryExternalStatus struct {
    // TODO(negz): place common Crossplane-y stuff.
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	RegistryID *string `json:"registryID,omitempty"`
	RepositoryURI *string `json:"repositoryURI,omitempty"`
}

// RepositoryStatus defines the observed state of Repository
type RepositoryStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider RepositoryExternalStatus `json:"atProvider"`
}

// Repository is the Schema for the Repositories API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   RepositorySpec   `json:"spec,omitempty"`
	Status RepositoryStatus `json:"status,omitempty"`
}

// RepositoryList contains a list of Repository
// +kubebuilder:object:root=true
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []Repository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}

