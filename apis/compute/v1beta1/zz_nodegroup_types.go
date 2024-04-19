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

type MaintenanceWindowInitParameters struct {

	// instances.start time of the window. This must be in UTC format that resolves to one of 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example, both 13:00-5 and 08:00 are valid.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type MaintenanceWindowObservation struct {

	// instances.start time of the window. This must be in UTC format that resolves to one of 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example, both 13:00-5 and 08:00 are valid.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type MaintenanceWindowParameters struct {

	// instances.start time of the window. This must be in UTC format that resolves to one of 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example, both 13:00-5 and 08:00 are valid.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type NodeGroupAutoscalingPolicyInitParameters struct {

	// Maximum size of the node group. Set to a value less than or equal
	// to 100 and greater than or equal to min-nodes.
	MaxNodes *float64 `json:"maxNodes,omitempty" tf:"max_nodes,omitempty"`

	// Minimum size of the node group. Must be less
	// than or equal to max-nodes. The default value is 0.
	MinNodes *float64 `json:"minNodes,omitempty" tf:"min_nodes,omitempty"`

	// The autoscaling mode. Set to one of the following:
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type NodeGroupAutoscalingPolicyObservation struct {

	// Maximum size of the node group. Set to a value less than or equal
	// to 100 and greater than or equal to min-nodes.
	MaxNodes *float64 `json:"maxNodes,omitempty" tf:"max_nodes,omitempty"`

	// Minimum size of the node group. Must be less
	// than or equal to max-nodes. The default value is 0.
	MinNodes *float64 `json:"minNodes,omitempty" tf:"min_nodes,omitempty"`

	// The autoscaling mode. Set to one of the following:
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type NodeGroupAutoscalingPolicyParameters struct {

	// Maximum size of the node group. Set to a value less than or equal
	// to 100 and greater than or equal to min-nodes.
	// +kubebuilder:validation:Optional
	MaxNodes *float64 `json:"maxNodes,omitempty" tf:"max_nodes,omitempty"`

	// Minimum size of the node group. Must be less
	// than or equal to max-nodes. The default value is 0.
	// +kubebuilder:validation:Optional
	MinNodes *float64 `json:"minNodes,omitempty" tf:"min_nodes,omitempty"`

	// The autoscaling mode. Set to one of the following:
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type NodeGroupInitParameters struct {

	// If you use sole-tenant nodes for your workloads, you can use the node
	// group autoscaler to automatically manage the sizes of your node groups.
	// One of initial_size or autoscaling_policy must be configured on resource creation.
	// Structure is documented below.
	AutoscalingPolicy []NodeGroupAutoscalingPolicyInitParameters `json:"autoscalingPolicy,omitempty" tf:"autoscaling_policy,omitempty"`

	// An optional textual description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The initial number of nodes in the node group. One of initial_size or autoscaling_policy must be configured on resource creation.
	InitialSize *float64 `json:"initialSize,omitempty" tf:"initial_size,omitempty"`

	// Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
	MaintenancePolicy *string `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// contains properties for the timeframe of maintenance
	// Structure is documented below.
	MaintenanceWindow []MaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The URL of the node template to which this node group belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.NodeTemplate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	NodeTemplate *string `json:"nodeTemplate,omitempty" tf:"node_template,omitempty"`

	// Reference to a NodeTemplate in compute to populate nodeTemplate.
	// +kubebuilder:validation:Optional
	NodeTemplateRef *v1.Reference `json:"nodeTemplateRef,omitempty" tf:"-"`

	// Selector for a NodeTemplate in compute to populate nodeTemplate.
	// +kubebuilder:validation:Optional
	NodeTemplateSelector *v1.Selector `json:"nodeTemplateSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Share settings for the node group.
	// Structure is documented below.
	ShareSettings []ShareSettingsInitParameters `json:"shareSettings,omitempty" tf:"share_settings,omitempty"`
}

type NodeGroupObservation struct {

	// If you use sole-tenant nodes for your workloads, you can use the node
	// group autoscaler to automatically manage the sizes of your node groups.
	// One of initial_size or autoscaling_policy must be configured on resource creation.
	// Structure is documented below.
	AutoscalingPolicy []NodeGroupAutoscalingPolicyObservation `json:"autoscalingPolicy,omitempty" tf:"autoscaling_policy,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional textual description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The initial number of nodes in the node group. One of initial_size or autoscaling_policy must be configured on resource creation.
	InitialSize *float64 `json:"initialSize,omitempty" tf:"initial_size,omitempty"`

	// Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
	MaintenancePolicy *string `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// contains properties for the timeframe of maintenance
	// Structure is documented below.
	MaintenanceWindow []MaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The URL of the node template to which this node group belongs.
	NodeTemplate *string `json:"nodeTemplate,omitempty" tf:"node_template,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Share settings for the node group.
	// Structure is documented below.
	ShareSettings []ShareSettingsObservation `json:"shareSettings,omitempty" tf:"share_settings,omitempty"`

	// The total number of nodes in the node group.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Zone where this node group is located
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NodeGroupParameters struct {

	// If you use sole-tenant nodes for your workloads, you can use the node
	// group autoscaler to automatically manage the sizes of your node groups.
	// One of initial_size or autoscaling_policy must be configured on resource creation.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AutoscalingPolicy []NodeGroupAutoscalingPolicyParameters `json:"autoscalingPolicy,omitempty" tf:"autoscaling_policy,omitempty"`

	// An optional textual description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The initial number of nodes in the node group. One of initial_size or autoscaling_policy must be configured on resource creation.
	// +kubebuilder:validation:Optional
	InitialSize *float64 `json:"initialSize,omitempty" tf:"initial_size,omitempty"`

	// Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
	// +kubebuilder:validation:Optional
	MaintenancePolicy *string `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// contains properties for the timeframe of maintenance
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The URL of the node template to which this node group belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.NodeTemplate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NodeTemplate *string `json:"nodeTemplate,omitempty" tf:"node_template,omitempty"`

	// Reference to a NodeTemplate in compute to populate nodeTemplate.
	// +kubebuilder:validation:Optional
	NodeTemplateRef *v1.Reference `json:"nodeTemplateRef,omitempty" tf:"-"`

	// Selector for a NodeTemplate in compute to populate nodeTemplate.
	// +kubebuilder:validation:Optional
	NodeTemplateSelector *v1.Selector `json:"nodeTemplateSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Share settings for the node group.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ShareSettings []ShareSettingsParameters `json:"shareSettings,omitempty" tf:"share_settings,omitempty"`

	// Zone where this node group is located
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type ProjectMapInitParameters struct {

	// The identifier for this object. Format specified above.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project_id",false)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Project in cloudplatform to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// The project id/number should be the same as the key of this project config in the project map.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project_id",false)
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in cloudplatform to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

type ProjectMapObservation struct {

	// The identifier for this object. Format specified above.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The project id/number should be the same as the key of this project config in the project map.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type ProjectMapParameters struct {

	// The identifier for this object. Format specified above.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project_id",false)
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Project in cloudplatform to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// The project id/number should be the same as the key of this project config in the project map.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project_id",false)
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in cloudplatform to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

type ShareSettingsInitParameters struct {

	// A map of project id and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
	// Structure is documented below.
	ProjectMap []ProjectMapInitParameters `json:"projectMap,omitempty" tf:"project_map,omitempty"`

	// Node group sharing type.
	// Possible values are: ORGANIZATION, SPECIFIC_PROJECTS, LOCAL.
	ShareType *string `json:"shareType,omitempty" tf:"share_type,omitempty"`
}

type ShareSettingsObservation struct {

	// A map of project id and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
	// Structure is documented below.
	ProjectMap []ProjectMapObservation `json:"projectMap,omitempty" tf:"project_map,omitempty"`

	// Node group sharing type.
	// Possible values are: ORGANIZATION, SPECIFIC_PROJECTS, LOCAL.
	ShareType *string `json:"shareType,omitempty" tf:"share_type,omitempty"`
}

type ShareSettingsParameters struct {

	// A map of project id and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ProjectMap []ProjectMapParameters `json:"projectMap,omitempty" tf:"project_map,omitempty"`

	// Node group sharing type.
	// Possible values are: ORGANIZATION, SPECIFIC_PROJECTS, LOCAL.
	// +kubebuilder:validation:Optional
	ShareType *string `json:"shareType" tf:"share_type,omitempty"`
}

// NodeGroupSpec defines the desired state of NodeGroup
type NodeGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeGroupParameters `json:"forProvider"`
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
	InitProvider NodeGroupInitParameters `json:"initProvider,omitempty"`
}

// NodeGroupStatus defines the observed state of NodeGroup.
type NodeGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NodeGroup is the Schema for the NodeGroups API. Represents a NodeGroup resource to manage a group of sole-tenant nodes.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeGroupSpec   `json:"spec"`
	Status            NodeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroupList contains a list of NodeGroups
type NodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeGroup `json:"items"`
}

// Repository type metadata.
var (
	NodeGroup_Kind             = "NodeGroup"
	NodeGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeGroup_Kind}.String()
	NodeGroup_KindAPIVersion   = NodeGroup_Kind + "." + CRDGroupVersion.String()
	NodeGroup_GroupVersionKind = CRDGroupVersion.WithKind(NodeGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeGroup{}, &NodeGroupList{})
}
