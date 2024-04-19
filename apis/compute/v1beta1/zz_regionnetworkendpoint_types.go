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

type RegionNetworkEndpointInitParameters struct {

	// Fully qualified domain name of network endpoint.
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// IPv4 address external endpoint.
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_IP_PORT.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Port number of network endpoint.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the containing network endpoint group is located.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The network endpoint group this endpoint is part of.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta2.RegionNetworkEndpointGroup
	RegionNetworkEndpointGroup *string `json:"regionNetworkEndpointGroup,omitempty" tf:"region_network_endpoint_group,omitempty"`

	// Reference to a RegionNetworkEndpointGroup in compute to populate regionNetworkEndpointGroup.
	// +kubebuilder:validation:Optional
	RegionNetworkEndpointGroupRef *v1.Reference `json:"regionNetworkEndpointGroupRef,omitempty" tf:"-"`

	// Selector for a RegionNetworkEndpointGroup in compute to populate regionNetworkEndpointGroup.
	// +kubebuilder:validation:Optional
	RegionNetworkEndpointGroupSelector *v1.Selector `json:"regionNetworkEndpointGroupSelector,omitempty" tf:"-"`
}

type RegionNetworkEndpointObservation struct {

	// Fully qualified domain name of network endpoint.
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// an identifier for the resource with format {{project}}/{{region}}/{{region_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IPv4 address external endpoint.
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_IP_PORT.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Port number of network endpoint.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the containing network endpoint group is located.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The network endpoint group this endpoint is part of.
	RegionNetworkEndpointGroup *string `json:"regionNetworkEndpointGroup,omitempty" tf:"region_network_endpoint_group,omitempty"`
}

type RegionNetworkEndpointParameters struct {

	// Fully qualified domain name of network endpoint.
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// IPv4 address external endpoint.
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_IP_PORT.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Port number of network endpoint.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the containing network endpoint group is located.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The network endpoint group this endpoint is part of.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta2.RegionNetworkEndpointGroup
	// +kubebuilder:validation:Optional
	RegionNetworkEndpointGroup *string `json:"regionNetworkEndpointGroup,omitempty" tf:"region_network_endpoint_group,omitempty"`

	// Reference to a RegionNetworkEndpointGroup in compute to populate regionNetworkEndpointGroup.
	// +kubebuilder:validation:Optional
	RegionNetworkEndpointGroupRef *v1.Reference `json:"regionNetworkEndpointGroupRef,omitempty" tf:"-"`

	// Selector for a RegionNetworkEndpointGroup in compute to populate regionNetworkEndpointGroup.
	// +kubebuilder:validation:Optional
	RegionNetworkEndpointGroupSelector *v1.Selector `json:"regionNetworkEndpointGroupSelector,omitempty" tf:"-"`
}

// RegionNetworkEndpointSpec defines the desired state of RegionNetworkEndpoint
type RegionNetworkEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionNetworkEndpointParameters `json:"forProvider"`
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
	InitProvider RegionNetworkEndpointInitParameters `json:"initProvider,omitempty"`
}

// RegionNetworkEndpointStatus defines the observed state of RegionNetworkEndpoint.
type RegionNetworkEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionNetworkEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RegionNetworkEndpoint is the Schema for the RegionNetworkEndpoints API. A Region network endpoint represents a IP address/FQDN and port combination that is part of a specific network endpoint group (NEG).
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RegionNetworkEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.port) || (has(self.initProvider) && has(self.initProvider.port))",message="spec.forProvider.port is a required parameter"
	Spec   RegionNetworkEndpointSpec   `json:"spec"`
	Status RegionNetworkEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionNetworkEndpointList contains a list of RegionNetworkEndpoints
type RegionNetworkEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionNetworkEndpoint `json:"items"`
}

// Repository type metadata.
var (
	RegionNetworkEndpoint_Kind             = "RegionNetworkEndpoint"
	RegionNetworkEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionNetworkEndpoint_Kind}.String()
	RegionNetworkEndpoint_KindAPIVersion   = RegionNetworkEndpoint_Kind + "." + CRDGroupVersion.String()
	RegionNetworkEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(RegionNetworkEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionNetworkEndpoint{}, &RegionNetworkEndpointList{})
}
