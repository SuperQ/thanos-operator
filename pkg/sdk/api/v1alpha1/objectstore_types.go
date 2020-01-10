// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/secret"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ObjectStoreSpec defines the desired state of ObjectStore
type ObjectStoreSpec struct {
	// Config
	Config    secret.Secret `json:"config"`
	Compactor *Compactor    `json:"compactor,omitempty"`
	Bucket    *Bucket       `json:"bucket,omitempty"`
}

type Compactor struct {
}

type Bucket struct {
}

// ObjectStoreStatus defines the observed state of ObjectStore
type ObjectStoreStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// ObjectStore is the Schema for the objectstores API
type ObjectStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObjectStoreSpec   `json:"spec,omitempty"`
	Status ObjectStoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectStoreList contains a list of ObjectStore
type ObjectStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectStore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ObjectStore{}, &ObjectStoreList{})
}