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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type TargetSecretStores struct {
	// Name of the SecretStore resource
	Name string `json:"name"`

	// Kind of the SecretStore resource (SecretStore or ClusterSecretStore)
	// Defaults to `SecretStore`
	// +optional
	Kind string `json:"kind,omitempty"`
}

type SecretRef struct {
	Name string `json:"name"`

	//Namespace string `json:"namespace"`

	Data map[string][]byte `json:"data,omitempty"`
}

// DistributedSecretsSpec defines the desired state of DistributedSecrets
type DistributedSecretsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	TargetSecretStores []TargetSecretStores `json:"targetSecretStores,omitempty"`

	SecretRef SecretRef `json:"secretRef"`
}

// DistributedSecretsStatus defines the observed state of DistributedSecrets
type DistributedSecretsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DistributedSecrets is the Schema for the distributedsecrets API
type DistributedSecrets struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DistributedSecretsSpec   `json:"spec,omitempty"`
	Status DistributedSecretsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DistributedSecretsList contains a list of DistributedSecrets
type DistributedSecretsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DistributedSecrets `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DistributedSecrets{}, &DistributedSecretsList{})
}
