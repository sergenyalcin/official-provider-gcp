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

type InstanceInitParameters struct {

	// Optional. Customer accept list represents the list of projects (id/number) on customer
	// side that can privately connect to the service attachment. It is an optional field
	// which the customers can provide during the instance creation. By default, the customer
	// project associated with the Apigee organization will be included to the list.
	ConsumerAcceptList []*string `json:"consumerAcceptList,omitempty" tf:"consumer_accept_list,omitempty"`

	// Description of the instance.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
	// Use the following format: projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DiskEncryptionKeyName *string `json:"diskEncryptionKeyName,omitempty" tf:"disk_encryption_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate diskEncryptionKeyName.
	// +kubebuilder:validation:Optional
	DiskEncryptionKeyNameRef *v1.Reference `json:"diskEncryptionKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate diskEncryptionKeyName.
	// +kubebuilder:validation:Optional
	DiskEncryptionKeyNameSelector *v1.Selector `json:"diskEncryptionKeyNameSelector,omitempty" tf:"-"`

	// Display name of the instance.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// IP range represents the customer-provided CIDR block of length 22 that will be used for
	// the Apigee instance creation. This optional range, if provided, should be freely
	// available as part of larger named range the customer has allocated to the Service
	// Networking peering. If this is not provided, Apigee will automatically request for any
	// available /22 CIDR block from Service Networking. The customer should use this CIDR block
	// for configuring their firewall needs to allow traffic from Apigee.
	// Input format: "a.b.c.d/22"
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// Required. Compute Engine location where the instance resides.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The size of the CIDR block range that will be reserved by the instance. For valid values,
	// see CidrRange on the documentation.
	PeeringCidrRange *string `json:"peeringCidrRange,omitempty" tf:"peering_cidr_range,omitempty"`
}

type InstanceObservation struct {

	// Optional. Customer accept list represents the list of projects (id/number) on customer
	// side that can privately connect to the service attachment. It is an optional field
	// which the customers can provide during the instance creation. By default, the customer
	// project associated with the Apigee organization will be included to the list.
	ConsumerAcceptList []*string `json:"consumerAcceptList,omitempty" tf:"consumer_accept_list,omitempty"`

	// Description of the instance.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
	// Use the following format: projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)
	DiskEncryptionKeyName *string `json:"diskEncryptionKeyName,omitempty" tf:"disk_encryption_key_name,omitempty"`

	// Display name of the instance.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// an identifier for the resource with format {{org_id}}/instances/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP range represents the customer-provided CIDR block of length 22 that will be used for
	// the Apigee instance creation. This optional range, if provided, should be freely
	// available as part of larger named range the customer has allocated to the Service
	// Networking peering. If this is not provided, Apigee will automatically request for any
	// available /22 CIDR block from Service Networking. The customer should use this CIDR block
	// for configuring their firewall needs to allow traffic from Apigee.
	// Input format: "a.b.c.d/22"
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// Required. Compute Engine location where the instance resides.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Apigee Organization associated with the Apigee instance,
	// in the format organizations/{{org_name}}.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The size of the CIDR block range that will be reserved by the instance. For valid values,
	// see CidrRange on the documentation.
	PeeringCidrRange *string `json:"peeringCidrRange,omitempty" tf:"peering_cidr_range,omitempty"`

	// Output only. Port number of the exposed Apigee endpoint.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Output only. Resource name of the service attachment created for the instance in
	// the format: projects//regions//serviceAttachments/* Apigee customers can privately
	// forward traffic to this service attachment using the PSC endpoints.
	ServiceAttachment *string `json:"serviceAttachment,omitempty" tf:"service_attachment,omitempty"`
}

type InstanceParameters struct {

	// Optional. Customer accept list represents the list of projects (id/number) on customer
	// side that can privately connect to the service attachment. It is an optional field
	// which the customers can provide during the instance creation. By default, the customer
	// project associated with the Apigee organization will be included to the list.
	// +kubebuilder:validation:Optional
	ConsumerAcceptList []*string `json:"consumerAcceptList,omitempty" tf:"consumer_accept_list,omitempty"`

	// Description of the instance.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
	// Use the following format: projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DiskEncryptionKeyName *string `json:"diskEncryptionKeyName,omitempty" tf:"disk_encryption_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate diskEncryptionKeyName.
	// +kubebuilder:validation:Optional
	DiskEncryptionKeyNameRef *v1.Reference `json:"diskEncryptionKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate diskEncryptionKeyName.
	// +kubebuilder:validation:Optional
	DiskEncryptionKeyNameSelector *v1.Selector `json:"diskEncryptionKeyNameSelector,omitempty" tf:"-"`

	// Display name of the instance.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// IP range represents the customer-provided CIDR block of length 22 that will be used for
	// the Apigee instance creation. This optional range, if provided, should be freely
	// available as part of larger named range the customer has allocated to the Service
	// Networking peering. If this is not provided, Apigee will automatically request for any
	// available /22 CIDR block from Service Networking. The customer should use this CIDR block
	// for configuring their firewall needs to allow traffic from Apigee.
	// Input format: "a.b.c.d/22"
	// +kubebuilder:validation:Optional
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// Required. Compute Engine location where the instance resides.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Apigee Organization associated with the Apigee instance,
	// in the format organizations/{{org_name}}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/apigee/v1beta1.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Organization in apigee to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDRef *v1.Reference `json:"orgIdRef,omitempty" tf:"-"`

	// Selector for a Organization in apigee to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDSelector *v1.Selector `json:"orgIdSelector,omitempty" tf:"-"`

	// The size of the CIDR block range that will be reserved by the instance. For valid values,
	// see CidrRange on the documentation.
	// +kubebuilder:validation:Optional
	PeeringCidrRange *string `json:"peeringCidrRange,omitempty" tf:"peering_cidr_range,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
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
	InitProvider InstanceInitParameters `json:"initProvider,omitempty"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Instance is the Schema for the Instances API. An
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
