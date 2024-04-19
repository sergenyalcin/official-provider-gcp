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

type DashboardInitParameters struct {

	// The JSON representation of a dashboard, following the format at https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
	// The representation of an existing dashboard can be found by using the API Explorer
	DashboardJSON *string `json:"dashboardJson,omitempty" tf:"dashboard_json,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type DashboardObservation struct {

	// The JSON representation of a dashboard, following the format at https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
	// The representation of an existing dashboard can be found by using the API Explorer
	DashboardJSON *string `json:"dashboardJson,omitempty" tf:"dashboard_json,omitempty"`

	// an identifier for the resource with format projects/{project_id_or_number}/dashboards/{dashboard_id}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type DashboardParameters struct {

	// The JSON representation of a dashboard, following the format at https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
	// The representation of an existing dashboard can be found by using the API Explorer
	// +kubebuilder:validation:Optional
	DashboardJSON *string `json:"dashboardJson,omitempty" tf:"dashboard_json,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// DashboardSpec defines the desired state of Dashboard
type DashboardSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DashboardParameters `json:"forProvider"`
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
	InitProvider DashboardInitParameters `json:"initProvider,omitempty"`
}

// DashboardStatus defines the observed state of Dashboard.
type DashboardStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DashboardObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Dashboard is the Schema for the Dashboards API. A Google Stackdriver dashboard.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Dashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dashboardJson) || (has(self.initProvider) && has(self.initProvider.dashboardJson))",message="spec.forProvider.dashboardJson is a required parameter"
	Spec   DashboardSpec   `json:"spec"`
	Status DashboardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DashboardList contains a list of Dashboards
type DashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dashboard `json:"items"`
}

// Repository type metadata.
var (
	Dashboard_Kind             = "Dashboard"
	Dashboard_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dashboard_Kind}.String()
	Dashboard_KindAPIVersion   = Dashboard_Kind + "." + CRDGroupVersion.String()
	Dashboard_GroupVersionKind = CRDGroupVersion.WithKind(Dashboard_Kind)
)

func init() {
	SchemeBuilder.Register(&Dashboard{}, &DashboardList{})
}
