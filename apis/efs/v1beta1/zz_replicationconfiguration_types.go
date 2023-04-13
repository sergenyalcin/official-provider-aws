/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DestinationObservation struct {

	// The availability zone in which the replica should be created. If specified, the replica will be created with One Zone storage. If omitted, regional storage will be used.
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty" tf:"availability_zone_name,omitempty"`

	// The fs ID of the replica.
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// The Key ID, ARN, alias, or alias ARN of the KMS key that should be used to encrypt the replica file system. If omitted, the default KMS key for EFS /aws/elasticfilesystem will be used.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The region in which the replica should be created.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The status of the replication.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DestinationParameters struct {

	// The availability zone in which the replica should be created. If specified, the replica will be created with One Zone storage. If omitted, regional storage will be used.
	// +kubebuilder:validation:Optional
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty" tf:"availability_zone_name,omitempty"`

	// The Key ID, ARN, alias, or alias ARN of the KMS key that should be used to encrypt the replica file system. If omitted, the default KMS key for EFS /aws/elasticfilesystem will be used.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The region in which the replica should be created.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type ReplicationConfigurationObservation struct {

	// When the replication configuration was created.
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// A destination configuration block (documented below).
	Destination []DestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) of the original source Amazon EFS file system in the replication configuration.
	OriginalSourceFileSystemArn *string `json:"originalSourceFileSystemArn,omitempty" tf:"original_source_file_system_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the current source file system in the replication configuration.
	SourceFileSystemArn *string `json:"sourceFileSystemArn,omitempty" tf:"source_file_system_arn,omitempty"`

	// The ID of the file system that is to be replicated.
	SourceFileSystemID *string `json:"sourceFileSystemId,omitempty" tf:"source_file_system_id,omitempty"`

	// The AWS Region in which the source Amazon EFS file system is located.
	SourceFileSystemRegion *string `json:"sourceFileSystemRegion,omitempty" tf:"source_file_system_region,omitempty"`
}

type ReplicationConfigurationParameters struct {

	// A destination configuration block (documented below).
	// +kubebuilder:validation:Optional
	Destination []DestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// The region in which the replica should be created.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the file system that is to be replicated.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/efs/v1beta1.FileSystem
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceFileSystemID *string `json:"sourceFileSystemId,omitempty" tf:"source_file_system_id,omitempty"`

	// Reference to a FileSystem in efs to populate sourceFileSystemId.
	// +kubebuilder:validation:Optional
	SourceFileSystemIDRef *v1.Reference `json:"sourceFileSystemIdRef,omitempty" tf:"-"`

	// Selector for a FileSystem in efs to populate sourceFileSystemId.
	// +kubebuilder:validation:Optional
	SourceFileSystemIDSelector *v1.Selector `json:"sourceFileSystemIdSelector,omitempty" tf:"-"`
}

// ReplicationConfigurationSpec defines the desired state of ReplicationConfiguration
type ReplicationConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicationConfigurationParameters `json:"forProvider"`
}

// ReplicationConfigurationStatus defines the observed state of ReplicationConfiguration.
type ReplicationConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicationConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationConfiguration is the Schema for the ReplicationConfigurations API. Provides an Elastic File System (EFS) Replication Configuration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReplicationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.destination)",message="destination is a required parameter"
	Spec   ReplicationConfigurationSpec   `json:"spec"`
	Status ReplicationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationConfigurationList contains a list of ReplicationConfigurations
type ReplicationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationConfiguration `json:"items"`
}

// Repository type metadata.
var (
	ReplicationConfiguration_Kind             = "ReplicationConfiguration"
	ReplicationConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReplicationConfiguration_Kind}.String()
	ReplicationConfiguration_KindAPIVersion   = ReplicationConfiguration_Kind + "." + CRDGroupVersion.String()
	ReplicationConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(ReplicationConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&ReplicationConfiguration{}, &ReplicationConfigurationList{})
}
