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

type DatabaseInitParameters struct {

	// The dialect of the Cloud Spanner Database.
	// If it is not provided, "GOOGLE_STANDARD_SQL" will be used.
	// Possible values are: GOOGLE_STANDARD_SQL, POSTGRESQL.
	DatabaseDialect *string `json:"databaseDialect,omitempty" tf:"database_dialect,omitempty"`

	// An optional list of DDL statements to run inside the newly created
	// database. Statements can create tables, indexes, etc. These statements
	// execute atomically with the creation of the database: if there is an
	// error in any statement, the database is not created.
	Ddl []*string `json:"ddl,omitempty" tf:"ddl,omitempty"`

	// Defaults to true.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Whether drop protection is enabled for this database. Defaults to false.
	// whereas setting “enableDropProtection” to true protects the database from deletions in all interfaces.
	// (2) Setting "enableDropProtection" to true also prevents the deletion of the parent instance containing the database.
	// "deletion_protection" attribute does not provide protection against the deletion of the parent instance.
	EnableDropProtection *bool `json:"enableDropProtection,omitempty" tf:"enable_drop_protection,omitempty"`

	// Encryption configuration for the database
	// Structure is documented below.
	EncryptionConfig *EncryptionConfigInitParameters `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The retention period for the database. The retention period must be between 1 hour
	// and 7 days, and can be specified in days, hours, minutes, or seconds. For example,
	// the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is 1h.
	// If this property is used, you must avoid adding new DDL statements to ddl that
	// update the database's version_retention_period.
	VersionRetentionPeriod *string `json:"versionRetentionPeriod,omitempty" tf:"version_retention_period,omitempty"`
}

type DatabaseObservation struct {

	// The dialect of the Cloud Spanner Database.
	// If it is not provided, "GOOGLE_STANDARD_SQL" will be used.
	// Possible values are: GOOGLE_STANDARD_SQL, POSTGRESQL.
	DatabaseDialect *string `json:"databaseDialect,omitempty" tf:"database_dialect,omitempty"`

	// An optional list of DDL statements to run inside the newly created
	// database. Statements can create tables, indexes, etc. These statements
	// execute atomically with the creation of the database: if there is an
	// error in any statement, the database is not created.
	Ddl []*string `json:"ddl,omitempty" tf:"ddl,omitempty"`

	// Defaults to true.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Whether drop protection is enabled for this database. Defaults to false.
	// whereas setting “enableDropProtection” to true protects the database from deletions in all interfaces.
	// (2) Setting "enableDropProtection" to true also prevents the deletion of the parent instance containing the database.
	// "deletion_protection" attribute does not provide protection against the deletion of the parent instance.
	EnableDropProtection *bool `json:"enableDropProtection,omitempty" tf:"enable_drop_protection,omitempty"`

	// Encryption configuration for the database
	// Structure is documented below.
	EncryptionConfig *EncryptionConfigObservation `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// an identifier for the resource with format {{instance}}/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance to create the database on.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// An explanation of the status of the database.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The retention period for the database. The retention period must be between 1 hour
	// and 7 days, and can be specified in days, hours, minutes, or seconds. For example,
	// the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is 1h.
	// If this property is used, you must avoid adding new DDL statements to ddl that
	// update the database's version_retention_period.
	VersionRetentionPeriod *string `json:"versionRetentionPeriod,omitempty" tf:"version_retention_period,omitempty"`
}

type DatabaseParameters struct {

	// The dialect of the Cloud Spanner Database.
	// If it is not provided, "GOOGLE_STANDARD_SQL" will be used.
	// Possible values are: GOOGLE_STANDARD_SQL, POSTGRESQL.
	// +kubebuilder:validation:Optional
	DatabaseDialect *string `json:"databaseDialect,omitempty" tf:"database_dialect,omitempty"`

	// An optional list of DDL statements to run inside the newly created
	// database. Statements can create tables, indexes, etc. These statements
	// execute atomically with the creation of the database: if there is an
	// error in any statement, the database is not created.
	// +kubebuilder:validation:Optional
	Ddl []*string `json:"ddl,omitempty" tf:"ddl,omitempty"`

	// Defaults to true.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Whether drop protection is enabled for this database. Defaults to false.
	// whereas setting “enableDropProtection” to true protects the database from deletions in all interfaces.
	// (2) Setting "enableDropProtection" to true also prevents the deletion of the parent instance containing the database.
	// "deletion_protection" attribute does not provide protection against the deletion of the parent instance.
	// +kubebuilder:validation:Optional
	EnableDropProtection *bool `json:"enableDropProtection,omitempty" tf:"enable_drop_protection,omitempty"`

	// Encryption configuration for the database
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EncryptionConfig *EncryptionConfigParameters `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// The instance to create the database on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/spanner/v1beta2.Instance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a Instance in spanner to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a Instance in spanner to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The retention period for the database. The retention period must be between 1 hour
	// and 7 days, and can be specified in days, hours, minutes, or seconds. For example,
	// the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is 1h.
	// If this property is used, you must avoid adding new DDL statements to ddl that
	// update the database's version_retention_period.
	// +kubebuilder:validation:Optional
	VersionRetentionPeriod *string `json:"versionRetentionPeriod,omitempty" tf:"version_retention_period,omitempty"`
}

type EncryptionConfigInitParameters struct {

	// Fully qualified name of the KMS key to use to encrypt this database. This key must exist
	// in the same location as the Spanner Database.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type EncryptionConfigObservation struct {

	// Fully qualified name of the KMS key to use to encrypt this database. This key must exist
	// in the same location as the Spanner Database.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type EncryptionConfigParameters struct {

	// Fully qualified name of the KMS key to use to encrypt this database. This key must exist
	// in the same location as the Spanner Database.
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName" tf:"kms_key_name,omitempty"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseParameters `json:"forProvider"`
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
	InitProvider DatabaseInitParameters `json:"initProvider,omitempty"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Database is the Schema for the Databases API. A Cloud Spanner Database which is hosted on a Spanner instance.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	Database_Kind             = "Database"
	Database_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Database_Kind}.String()
	Database_KindAPIVersion   = Database_Kind + "." + CRDGroupVersion.String()
	Database_GroupVersionKind = CRDGroupVersion.WithKind(Database_Kind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
