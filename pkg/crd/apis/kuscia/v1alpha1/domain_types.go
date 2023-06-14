// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// Domain is the Schema for the domain API.
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              DomainSpec `json:"spec"`
	// +optional
	Status *DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DomainList contains a list of domains.
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

type DomainRole string

const (
	Partner DomainRole = "partner"
)

// DomainSpec defines the details of domain.
type DomainSpec struct {
	// Role is used to represent the role of domain. Default is omit empty.
	// If it is used to represent a P2P domain, please set the value to partner.
	// kuscia scheduler will check the role, if it is partner, pods will be ignored by scheduler.
	// +optional
	Role DomainRole `json:"role,omitempty"`
	// +optional
	Cert string `json:"cert,omitempty"`
	// +optional
	ResourceQuota *DomainResourceQuota `json:"resourceQuota,omitempty"`
}

// DomainResourceQuota defines domain resource quota.
type DomainResourceQuota struct {
	// +optional
	// +kubebuilder:validation:Minimum=0
	PodMaxCount *int `json:"podMaxCount,omitempty"`
}

// DomainStatus defines domain status.
type DomainStatus struct {
	// +optional
	NodeStatuses []NodeStatus `json:"nodeStatuses,omitempty"`
}

// NodeStatus defines node status under domain.
type NodeStatus struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"`
	// +optional
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}
