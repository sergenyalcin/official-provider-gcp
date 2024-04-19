// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RegistryInitParameters struct {

	// The location of the registry. One of ASIA, EU, US or not specified. See the official documentation for more information on registry locations.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type RegistryObservation struct {

	// The URI of the created resource.
	BucketSelfLink *string `json:"bucketSelfLink,omitempty" tf:"bucket_self_link,omitempty"`

	// The name of the bucket that supports the Container Registry. In the form of artifacts.{project}.appspot.com or {location}.artifacts.{project}.appspot.com if location is specified.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the registry. One of ASIA, EU, US or not specified. See the official documentation for more information on registry locations.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type RegistryParameters struct {

	// The location of the registry. One of ASIA, EU, US or not specified. See the official documentation for more information on registry locations.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// RegistrySpec defines the desired state of Registry
type RegistrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RegistryInitParameters `json:"initProvider,omitempty"`
}

// RegistryStatus defines the observed state of Registry.
type RegistryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Registry is the Schema for the Registrys API. Ensures the GCS bucket backing Google Container Registry exists.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Registry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegistrySpec   `json:"spec"`
	Status            RegistryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryList contains a list of Registrys
type RegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registry `json:"items"`
}

// Repository type metadata.
var (
	Registry_Kind             = "Registry"
	Registry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Registry_Kind}.String()
	Registry_KindAPIVersion   = Registry_Kind + "." + CRDGroupVersion.String()
	Registry_GroupVersionKind = CRDGroupVersion.WithKind(Registry_Kind)
)

func init() {
	SchemeBuilder.Register(&Registry{}, &RegistryList{})
}
