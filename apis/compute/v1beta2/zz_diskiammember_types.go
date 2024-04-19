// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type DiskIAMMemberInitParameters struct {
	Condition *ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +crossplane:generate:reference:type=Disk
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Disk to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Disk to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DiskIAMMemberObservation struct {
	Condition *ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DiskIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition *ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +crossplane:generate:reference:type=Disk
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Disk to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Disk to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// DiskIAMMemberSpec defines the desired state of DiskIAMMember
type DiskIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskIAMMemberParameters `json:"forProvider"`
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
	InitProvider DiskIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// DiskIAMMemberStatus defines the observed state of DiskIAMMember.
type DiskIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DiskIAMMember is the Schema for the DiskIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DiskIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.member) || (has(self.initProvider) && has(self.initProvider.member))",message="spec.forProvider.member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   DiskIAMMemberSpec   `json:"spec"`
	Status DiskIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskIAMMemberList contains a list of DiskIAMMembers
type DiskIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskIAMMember `json:"items"`
}

// Repository type metadata.
var (
	DiskIAMMember_Kind             = "DiskIAMMember"
	DiskIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskIAMMember_Kind}.String()
	DiskIAMMember_KindAPIVersion   = DiskIAMMember_Kind + "." + CRDGroupVersion.String()
	DiskIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(DiskIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskIAMMember{}, &DiskIAMMemberList{})
}
