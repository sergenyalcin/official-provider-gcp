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

type FieldsInitParameters struct {

	// Holds the value for a tag field with boolean type.
	BoolValue *bool `json:"boolValue,omitempty" tf:"bool_value,omitempty"`

	// Holds the value for a tag field with double type.
	DoubleValue *float64 `json:"doubleValue,omitempty" tf:"double_value,omitempty"`

	// Holds the value for a tag field with enum type. This value must be one of the allowed values in the definition of this enum.
	EnumValue *string `json:"enumValue,omitempty" tf:"enum_value,omitempty"`

	// The identifier for this object. Format specified above.
	FieldName *string `json:"fieldName,omitempty" tf:"field_name,omitempty"`

	// Holds the value for a tag field with string type.
	StringValue *string `json:"stringValue,omitempty" tf:"string_value,omitempty"`

	// Holds the value for a tag field with timestamp type.
	TimestampValue *string `json:"timestampValue,omitempty" tf:"timestamp_value,omitempty"`
}

type FieldsObservation struct {

	// Holds the value for a tag field with boolean type.
	BoolValue *bool `json:"boolValue,omitempty" tf:"bool_value,omitempty"`

	// (Output)
	// The display name of this field
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Holds the value for a tag field with double type.
	DoubleValue *float64 `json:"doubleValue,omitempty" tf:"double_value,omitempty"`

	// Holds the value for a tag field with enum type. This value must be one of the allowed values in the definition of this enum.
	EnumValue *string `json:"enumValue,omitempty" tf:"enum_value,omitempty"`

	// The identifier for this object. Format specified above.
	FieldName *string `json:"fieldName,omitempty" tf:"field_name,omitempty"`

	// (Output)
	// The order of this field with respect to other fields in this tag. For example, a higher value can indicate
	// a more important field. The value can be negative. Multiple fields can have the same order, and field orders
	// within a tag do not have to be sequential.
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// Holds the value for a tag field with string type.
	StringValue *string `json:"stringValue,omitempty" tf:"string_value,omitempty"`

	// Holds the value for a tag field with timestamp type.
	TimestampValue *string `json:"timestampValue,omitempty" tf:"timestamp_value,omitempty"`
}

type FieldsParameters struct {

	// Holds the value for a tag field with boolean type.
	// +kubebuilder:validation:Optional
	BoolValue *bool `json:"boolValue,omitempty" tf:"bool_value,omitempty"`

	// Holds the value for a tag field with double type.
	// +kubebuilder:validation:Optional
	DoubleValue *float64 `json:"doubleValue,omitempty" tf:"double_value,omitempty"`

	// Holds the value for a tag field with enum type. This value must be one of the allowed values in the definition of this enum.
	// +kubebuilder:validation:Optional
	EnumValue *string `json:"enumValue,omitempty" tf:"enum_value,omitempty"`

	// The identifier for this object. Format specified above.
	// +kubebuilder:validation:Optional
	FieldName *string `json:"fieldName" tf:"field_name,omitempty"`

	// Holds the value for a tag field with string type.
	// +kubebuilder:validation:Optional
	StringValue *string `json:"stringValue,omitempty" tf:"string_value,omitempty"`

	// Holds the value for a tag field with timestamp type.
	// +kubebuilder:validation:Optional
	TimestampValue *string `json:"timestampValue,omitempty" tf:"timestamp_value,omitempty"`
}

type TagInitParameters struct {

	// Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an
	// individual column based on that schema.
	// For attaching a tag to a nested column, use . to separate the column names. Example:
	// outer_column.inner_column
	Column *string `json:"column,omitempty" tf:"column,omitempty"`

	// This maps the ID of a tag field to the value of and additional information about that field.
	// Valid field IDs are defined by the tag's template. A tag must have at least 1 field and at most 500 fields.
	// Structure is documented below.
	Fields []FieldsInitParameters `json:"fields,omitempty" tf:"fields,omitempty"`

	// The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group, the tag will be attached to
	// all entries in that group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datacatalog/v1beta1.Entry
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Reference to a Entry in datacatalog to populate parent.
	// +kubebuilder:validation:Optional
	ParentRef *v1.Reference `json:"parentRef,omitempty" tf:"-"`

	// Selector for a Entry in datacatalog to populate parent.
	// +kubebuilder:validation:Optional
	ParentSelector *v1.Selector `json:"parentSelector,omitempty" tf:"-"`

	// The resource name of the tag template that this tag uses. Example:
	// projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
	// This field cannot be modified after creation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datacatalog/v1beta1.TagTemplate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// Reference to a TagTemplate in datacatalog to populate template.
	// +kubebuilder:validation:Optional
	TemplateRef *v1.Reference `json:"templateRef,omitempty" tf:"-"`

	// Selector for a TagTemplate in datacatalog to populate template.
	// +kubebuilder:validation:Optional
	TemplateSelector *v1.Selector `json:"templateSelector,omitempty" tf:"-"`
}

type TagObservation struct {

	// Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an
	// individual column based on that schema.
	// For attaching a tag to a nested column, use . to separate the column names. Example:
	// outer_column.inner_column
	Column *string `json:"column,omitempty" tf:"column,omitempty"`

	// This maps the ID of a tag field to the value of and additional information about that field.
	// Valid field IDs are defined by the tag's template. A tag must have at least 1 field and at most 500 fields.
	// Structure is documented below.
	Fields []FieldsObservation `json:"fields,omitempty" tf:"fields,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource name of the tag in URL format. Example:
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}/tags/{tag_id} or
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id}
	// where tag_id is a system-generated identifier. Note that this Tag may not actually be stored in the location in this name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group, the tag will be attached to
	// all entries in that group.
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// The resource name of the tag template that this tag uses. Example:
	// projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
	// This field cannot be modified after creation.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// The display name of the tag template.
	TemplateDisplayname *string `json:"templateDisplayname,omitempty" tf:"template_displayname,omitempty"`
}

type TagParameters struct {

	// Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an
	// individual column based on that schema.
	// For attaching a tag to a nested column, use . to separate the column names. Example:
	// outer_column.inner_column
	// +kubebuilder:validation:Optional
	Column *string `json:"column,omitempty" tf:"column,omitempty"`

	// This maps the ID of a tag field to the value of and additional information about that field.
	// Valid field IDs are defined by the tag's template. A tag must have at least 1 field and at most 500 fields.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Fields []FieldsParameters `json:"fields,omitempty" tf:"fields,omitempty"`

	// The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group, the tag will be attached to
	// all entries in that group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datacatalog/v1beta1.Entry
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Reference to a Entry in datacatalog to populate parent.
	// +kubebuilder:validation:Optional
	ParentRef *v1.Reference `json:"parentRef,omitempty" tf:"-"`

	// Selector for a Entry in datacatalog to populate parent.
	// +kubebuilder:validation:Optional
	ParentSelector *v1.Selector `json:"parentSelector,omitempty" tf:"-"`

	// The resource name of the tag template that this tag uses. Example:
	// projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
	// This field cannot be modified after creation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datacatalog/v1beta1.TagTemplate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// Reference to a TagTemplate in datacatalog to populate template.
	// +kubebuilder:validation:Optional
	TemplateRef *v1.Reference `json:"templateRef,omitempty" tf:"-"`

	// Selector for a TagTemplate in datacatalog to populate template.
	// +kubebuilder:validation:Optional
	TemplateSelector *v1.Selector `json:"templateSelector,omitempty" tf:"-"`
}

// TagSpec defines the desired state of Tag
type TagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagParameters `json:"forProvider"`
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
	InitProvider TagInitParameters `json:"initProvider,omitempty"`
}

// TagStatus defines the observed state of Tag.
type TagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Tag is the Schema for the Tags API. Tags are used to attach custom metadata to Data Catalog resources.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Tag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fields) || (has(self.initProvider) && has(self.initProvider.fields))",message="spec.forProvider.fields is a required parameter"
	Spec   TagSpec   `json:"spec"`
	Status TagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagList contains a list of Tags
type TagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tag `json:"items"`
}

// Repository type metadata.
var (
	Tag_Kind             = "Tag"
	Tag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tag_Kind}.String()
	Tag_KindAPIVersion   = Tag_Kind + "." + CRDGroupVersion.String()
	Tag_GroupVersionKind = CRDGroupVersion.WithKind(Tag_Kind)
)

func init() {
	SchemeBuilder.Register(&Tag{}, &TagList{})
}
