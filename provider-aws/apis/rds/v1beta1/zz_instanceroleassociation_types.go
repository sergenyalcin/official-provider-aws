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

type InstanceRoleAssociationObservation struct {

	// DB Instance Identifier and IAM Role ARN separated by a comma
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type InstanceRoleAssociationParameters struct {

	// DB Instance Identifier to associate with the IAM Role.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/rds/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DBInstanceIdentifier *string `json:"dbInstanceIdentifier,omitempty" tf:"db_instance_identifier,omitempty"`

	// Reference to a Instance in rds to populate dbInstanceIdentifier.
	// +kubebuilder:validation:Optional
	DBInstanceIdentifierRef *v1.Reference `json:"dbInstanceIdentifierRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate dbInstanceIdentifier.
	// +kubebuilder:validation:Optional
	DBInstanceIdentifierSelector *v1.Selector `json:"dbInstanceIdentifierSelector,omitempty" tf:"-"`

	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the SupportedFeatureNames list returned by AWS CLI rds describe-db-engine-versions.
	// +kubebuilder:validation:Required
	FeatureName *string `json:"featureName" tf:"feature_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Resource Name  of the IAM Role to associate with the DB Instance.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

// InstanceRoleAssociationSpec defines the desired state of InstanceRoleAssociation
type InstanceRoleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceRoleAssociationParameters `json:"forProvider"`
}

// InstanceRoleAssociationStatus defines the observed state of InstanceRoleAssociation.
type InstanceRoleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceRoleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceRoleAssociation is the Schema for the InstanceRoleAssociations API. Manages an RDS DB Instance association with an IAM Role.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InstanceRoleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceRoleAssociationSpec   `json:"spec"`
	Status            InstanceRoleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceRoleAssociationList contains a list of InstanceRoleAssociations
type InstanceRoleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceRoleAssociation `json:"items"`
}

// Repository type metadata.
var (
	InstanceRoleAssociation_Kind             = "InstanceRoleAssociation"
	InstanceRoleAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceRoleAssociation_Kind}.String()
	InstanceRoleAssociation_KindAPIVersion   = InstanceRoleAssociation_Kind + "." + CRDGroupVersion.String()
	InstanceRoleAssociation_GroupVersionKind = CRDGroupVersion.WithKind(InstanceRoleAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceRoleAssociation{}, &InstanceRoleAssociationList{})
}
