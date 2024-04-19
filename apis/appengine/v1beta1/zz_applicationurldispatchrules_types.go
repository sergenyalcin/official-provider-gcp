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

type ApplicationURLDispatchRulesInitParameters struct {

	// Rules to match an HTTP request and dispatch that request to a service.
	// Structure is documented below.
	DispatchRules []DispatchRulesInitParameters `json:"dispatchRules,omitempty" tf:"dispatch_rules,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ApplicationURLDispatchRulesObservation struct {

	// Rules to match an HTTP request and dispatch that request to a service.
	// Structure is documented below.
	DispatchRules []DispatchRulesObservation `json:"dispatchRules,omitempty" tf:"dispatch_rules,omitempty"`

	// an identifier for the resource with format {{project}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ApplicationURLDispatchRulesParameters struct {

	// Rules to match an HTTP request and dispatch that request to a service.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DispatchRules []DispatchRulesParameters `json:"dispatchRules,omitempty" tf:"dispatch_rules,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type DispatchRulesInitParameters struct {

	// Domain name to match against. The wildcard "" is supported if specified before a period: ".".
	// Defaults to matching all domains: "*".
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/appengine/v1beta2.StandardAppVersion
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("service",false)
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// Reference to a StandardAppVersion in appengine to populate service.
	// +kubebuilder:validation:Optional
	ServiceRef *v1.Reference `json:"serviceRef,omitempty" tf:"-"`

	// Selector for a StandardAppVersion in appengine to populate service.
	// +kubebuilder:validation:Optional
	ServiceSelector *v1.Selector `json:"serviceSelector,omitempty" tf:"-"`
}

type DispatchRulesObservation struct {

	// Domain name to match against. The wildcard "" is supported if specified before a period: ".".
	// Defaults to matching all domains: "*".
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type DispatchRulesParameters struct {

	// Domain name to match against. The wildcard "" is supported if specified before a period: ".".
	// Defaults to matching all domains: "*".
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`

	// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/appengine/v1beta2.StandardAppVersion
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("service",false)
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// Reference to a StandardAppVersion in appengine to populate service.
	// +kubebuilder:validation:Optional
	ServiceRef *v1.Reference `json:"serviceRef,omitempty" tf:"-"`

	// Selector for a StandardAppVersion in appengine to populate service.
	// +kubebuilder:validation:Optional
	ServiceSelector *v1.Selector `json:"serviceSelector,omitempty" tf:"-"`
}

// ApplicationURLDispatchRulesSpec defines the desired state of ApplicationURLDispatchRules
type ApplicationURLDispatchRulesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationURLDispatchRulesParameters `json:"forProvider"`
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
	InitProvider ApplicationURLDispatchRulesInitParameters `json:"initProvider,omitempty"`
}

// ApplicationURLDispatchRulesStatus defines the observed state of ApplicationURLDispatchRules.
type ApplicationURLDispatchRulesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationURLDispatchRulesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ApplicationURLDispatchRules is the Schema for the ApplicationURLDispatchRuless API. Rules to match an HTTP request and dispatch that request to a service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ApplicationURLDispatchRules struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dispatchRules) || (has(self.initProvider) && has(self.initProvider.dispatchRules))",message="spec.forProvider.dispatchRules is a required parameter"
	Spec   ApplicationURLDispatchRulesSpec   `json:"spec"`
	Status ApplicationURLDispatchRulesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationURLDispatchRulesList contains a list of ApplicationURLDispatchRuless
type ApplicationURLDispatchRulesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationURLDispatchRules `json:"items"`
}

// Repository type metadata.
var (
	ApplicationURLDispatchRules_Kind             = "ApplicationURLDispatchRules"
	ApplicationURLDispatchRules_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationURLDispatchRules_Kind}.String()
	ApplicationURLDispatchRules_KindAPIVersion   = ApplicationURLDispatchRules_Kind + "." + CRDGroupVersion.String()
	ApplicationURLDispatchRules_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationURLDispatchRules_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationURLDispatchRules{}, &ApplicationURLDispatchRulesList{})
}
