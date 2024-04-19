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

type ConditionsInitParameters struct {
}

type ConditionsObservation struct {

	// (Output)
	// Human readable message indicating details about the current status.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (Output)
	// One-word CamelCase reason for the condition's current status.
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	// The current status of the DomainMapping.
	// Structure is documented below.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Resource record type. Example: AAAA.
	// Possible values are: A, AAAA, CNAME.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConditionsParameters struct {
}

type DomainMappingInitParameters struct {

	// The location of the cloud run instance. eg us-central1
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Metadata associated with this DomainMapping.
	// Structure is documented below.
	Metadata *MetadataInitParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Name should be a verified domain
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The spec for this DomainMapping.
	// Structure is documented below.
	Spec *SpecInitParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type DomainMappingObservation struct {

	// an identifier for the resource with format locations/{{location}}/namespaces/{{project}}/domainmappings/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the cloud run instance. eg us-central1
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Metadata associated with this DomainMapping.
	// Structure is documented below.
	Metadata *MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Name should be a verified domain
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The spec for this DomainMapping.
	// Structure is documented below.
	Spec *SpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`

	// The current status of the DomainMapping.
	// Structure is documented below.
	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`
}

type DomainMappingParameters struct {

	// The location of the cloud run instance. eg us-central1
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Metadata associated with this DomainMapping.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Metadata *MetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Name should be a verified domain
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The spec for this DomainMapping.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Spec *SpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type MetadataInitParameters struct {

	// Annotations is a key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata. More
	// info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
	// Note: The Cloud Run API may add additional annotations that were not provided in your config.ignore_changes rule to the metadata.0.annotations field.
	// Note: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field effective_annotations for all of the annotations present on the resource.
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and routes.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// In Cloud Run the namespace must be equal to either the
	// project ID or project number.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Reference to a Project in cloudplatform to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceRef *v1.Reference `json:"namespaceRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceSelector *v1.Selector `json:"namespaceSelector,omitempty" tf:"-"`
}

type MetadataObservation struct {

	// Annotations is a key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata. More
	// info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
	// Note: The Cloud Run API may add additional annotations that were not provided in your config.ignore_changes rule to the metadata.0.annotations field.
	// Note: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field effective_annotations for all of the annotations present on the resource.
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +mapType=granular
	EffectiveAnnotations map[string]*string `json:"effectiveAnnotations,omitempty" tf:"effective_annotations,omitempty"`

	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// (Output)
	// A sequence number representing a specific generation of the desired state.
	Generation *float64 `json:"generation,omitempty" tf:"generation,omitempty"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and routes.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// In Cloud Run the namespace must be equal to either the
	// project ID or project number.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (Output)
	// An opaque value that represents the internal version of this object that
	// can be used by clients to determine when objects have changed. May be used
	// for optimistic concurrency, change detection, and the watch operation on a
	// resource or set of resources. They may only be valid for a
	// particular resource or set of resources.
	// More info:
	// https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	ResourceVersion *string `json:"resourceVersion,omitempty" tf:"resource_version,omitempty"`

	// (Output)
	// SelfLink is a URL representing this object.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// (Output)
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// (Output)
	// UID is a unique id generated by the server on successful creation of a resource and is not
	// allowed to change on PUT operations.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type MetadataParameters struct {

	// Annotations is a key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata. More
	// info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
	// Note: The Cloud Run API may add additional annotations that were not provided in your config.ignore_changes rule to the metadata.0.annotations field.
	// Note: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field effective_annotations for all of the annotations present on the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and routes.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// In Cloud Run the namespace must be equal to either the
	// project ID or project number.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Reference to a Project in cloudplatform to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceRef *v1.Reference `json:"namespaceRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceSelector *v1.Selector `json:"namespaceSelector,omitempty" tf:"-"`
}

type ResourceRecordsInitParameters struct {
}

type ResourceRecordsObservation struct {

	// (Output)
	// Relative name of the object affected by this record. Only applicable for
	// CNAME records. Example: 'www'.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Output)
	// Data for this record. Values vary by record type, as defined in RFC 1035
	// (section 5) and RFC 1034 (section 3.6.1).
	Rrdata *string `json:"rrdata,omitempty" tf:"rrdata,omitempty"`

	// Resource record type. Example: AAAA.
	// Possible values are: A, AAAA, CNAME.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceRecordsParameters struct {
}

type SpecInitParameters struct {

	// The mode of the certificate.
	// Default value is AUTOMATIC.
	// Possible values are: NONE, AUTOMATIC.
	CertificateMode *string `json:"certificateMode,omitempty" tf:"certificate_mode,omitempty"`

	// If set, the mapping will override any mapping set before this spec was set.
	// It is recommended that the user leaves this empty to receive an error
	// warning about a potential conflict and only set it once the respective UI
	// has given such a warning.
	ForceOverride *bool `json:"forceOverride,omitempty" tf:"force_override,omitempty"`

	// The name of the Cloud Run Service that this DomainMapping applies to.
	// The route must exist.
	// +crossplane:generate:reference:type=Service
	RouteName *string `json:"routeName,omitempty" tf:"route_name,omitempty"`

	// Reference to a Service to populate routeName.
	// +kubebuilder:validation:Optional
	RouteNameRef *v1.Reference `json:"routeNameRef,omitempty" tf:"-"`

	// Selector for a Service to populate routeName.
	// +kubebuilder:validation:Optional
	RouteNameSelector *v1.Selector `json:"routeNameSelector,omitempty" tf:"-"`
}

type SpecObservation struct {

	// The mode of the certificate.
	// Default value is AUTOMATIC.
	// Possible values are: NONE, AUTOMATIC.
	CertificateMode *string `json:"certificateMode,omitempty" tf:"certificate_mode,omitempty"`

	// If set, the mapping will override any mapping set before this spec was set.
	// It is recommended that the user leaves this empty to receive an error
	// warning about a potential conflict and only set it once the respective UI
	// has given such a warning.
	ForceOverride *bool `json:"forceOverride,omitempty" tf:"force_override,omitempty"`

	// The name of the Cloud Run Service that this DomainMapping applies to.
	// The route must exist.
	RouteName *string `json:"routeName,omitempty" tf:"route_name,omitempty"`
}

type SpecParameters struct {

	// The mode of the certificate.
	// Default value is AUTOMATIC.
	// Possible values are: NONE, AUTOMATIC.
	// +kubebuilder:validation:Optional
	CertificateMode *string `json:"certificateMode,omitempty" tf:"certificate_mode,omitempty"`

	// If set, the mapping will override any mapping set before this spec was set.
	// It is recommended that the user leaves this empty to receive an error
	// warning about a potential conflict and only set it once the respective UI
	// has given such a warning.
	// +kubebuilder:validation:Optional
	ForceOverride *bool `json:"forceOverride,omitempty" tf:"force_override,omitempty"`

	// The name of the Cloud Run Service that this DomainMapping applies to.
	// The route must exist.
	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	RouteName *string `json:"routeName,omitempty" tf:"route_name,omitempty"`

	// Reference to a Service to populate routeName.
	// +kubebuilder:validation:Optional
	RouteNameRef *v1.Reference `json:"routeNameRef,omitempty" tf:"-"`

	// Selector for a Service to populate routeName.
	// +kubebuilder:validation:Optional
	RouteNameSelector *v1.Selector `json:"routeNameSelector,omitempty" tf:"-"`
}

type StatusInitParameters struct {
}

type StatusObservation struct {

	// (Output)
	// Array of observed DomainMappingConditions, indicating the current state
	// of the DomainMapping.
	// Structure is documented below.
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// (Output)
	// The name of the route that the mapping currently points to.
	MappedRouteName *string `json:"mappedRouteName,omitempty" tf:"mapped_route_name,omitempty"`

	// (Output)
	// ObservedGeneration is the 'Generation' of the DomainMapping that
	// was last processed by the controller.
	ObservedGeneration *float64 `json:"observedGeneration,omitempty" tf:"observed_generation,omitempty"`

	// The resource records required to configure this domain mapping. These
	// records must be added to the domain's DNS configuration in order to
	// serve the application via this domain mapping.
	// Structure is documented below.
	ResourceRecords []ResourceRecordsObservation `json:"resourceRecords,omitempty" tf:"resource_records,omitempty"`
}

type StatusParameters struct {
}

// DomainMappingSpec defines the desired state of DomainMapping
type DomainMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainMappingParameters `json:"forProvider"`
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
	InitProvider DomainMappingInitParameters `json:"initProvider,omitempty"`
}

// DomainMappingStatus defines the observed state of DomainMapping.
type DomainMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DomainMapping is the Schema for the DomainMappings API. Resource to hold the state and status of a user's domain mapping.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DomainMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.spec) || (has(self.initProvider) && has(self.initProvider.spec))",message="spec.forProvider.spec is a required parameter"
	Spec   DomainMappingSpec   `json:"spec"`
	Status DomainMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainMappingList contains a list of DomainMappings
type DomainMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainMapping `json:"items"`
}

// Repository type metadata.
var (
	DomainMapping_Kind             = "DomainMapping"
	DomainMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainMapping_Kind}.String()
	DomainMapping_KindAPIVersion   = DomainMapping_Kind + "." + CRDGroupVersion.String()
	DomainMapping_GroupVersionKind = CRDGroupVersion.WithKind(DomainMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainMapping{}, &DomainMappingList{})
}
