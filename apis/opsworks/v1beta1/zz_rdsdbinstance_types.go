// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type RDSDBInstanceInitParameters struct {

	// A db username
	DBUser *string `json:"dbUser,omitempty" tf:"db_user,omitempty"`

	// The db instance to register for this stack. Changing this will force a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta2.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	RDSDBInstanceArn *string `json:"rdsDbInstanceArn,omitempty" tf:"rds_db_instance_arn,omitempty"`

	// Reference to a Instance in rds to populate rdsDbInstanceArn.
	// +kubebuilder:validation:Optional
	RDSDBInstanceArnRef *v1.Reference `json:"rdsDbInstanceArnRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate rdsDbInstanceArn.
	// +kubebuilder:validation:Optional
	RDSDBInstanceArnSelector *v1.Selector `json:"rdsDbInstanceArnSelector,omitempty" tf:"-"`

	// The stack to register a db instance for. Changing this will force a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/opsworks/v1beta1.Stack
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// Reference to a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDRef *v1.Reference `json:"stackIdRef,omitempty" tf:"-"`

	// Selector for a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDSelector *v1.Selector `json:"stackIdSelector,omitempty" tf:"-"`
}

type RDSDBInstanceObservation struct {

	// A db username
	DBUser *string `json:"dbUser,omitempty" tf:"db_user,omitempty"`

	// The computed id. Please note that this is only used internally to identify the stack <-> instance relation. This value is not used in aws.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The db instance to register for this stack. Changing this will force a new resource.
	RDSDBInstanceArn *string `json:"rdsDbInstanceArn,omitempty" tf:"rds_db_instance_arn,omitempty"`

	// The stack to register a db instance for. Changing this will force a new resource.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`
}

type RDSDBInstanceParameters struct {

	// A db password
	// +kubebuilder:validation:Optional
	DBPasswordSecretRef v1.SecretKeySelector `json:"dbPasswordSecretRef" tf:"-"`

	// A db username
	// +kubebuilder:validation:Optional
	DBUser *string `json:"dbUser,omitempty" tf:"db_user,omitempty"`

	// The db instance to register for this stack. Changing this will force a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta2.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RDSDBInstanceArn *string `json:"rdsDbInstanceArn,omitempty" tf:"rds_db_instance_arn,omitempty"`

	// Reference to a Instance in rds to populate rdsDbInstanceArn.
	// +kubebuilder:validation:Optional
	RDSDBInstanceArnRef *v1.Reference `json:"rdsDbInstanceArnRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate rdsDbInstanceArn.
	// +kubebuilder:validation:Optional
	RDSDBInstanceArnSelector *v1.Selector `json:"rdsDbInstanceArnSelector,omitempty" tf:"-"`

	// The stack to register a db instance for. Changing this will force a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/opsworks/v1beta1.Stack
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// Reference to a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDRef *v1.Reference `json:"stackIdRef,omitempty" tf:"-"`

	// Selector for a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDSelector *v1.Selector `json:"stackIdSelector,omitempty" tf:"-"`
}

// RDSDBInstanceSpec defines the desired state of RDSDBInstance
type RDSDBInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RDSDBInstanceParameters `json:"forProvider"`
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
	InitProvider RDSDBInstanceInitParameters `json:"initProvider,omitempty"`
}

// RDSDBInstanceStatus defines the observed state of RDSDBInstance.
type RDSDBInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RDSDBInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RDSDBInstance is the Schema for the RDSDBInstances API. Provides an OpsWorks RDS DB Instance resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RDSDBInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dbPasswordSecretRef)",message="spec.forProvider.dbPasswordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dbUser) || (has(self.initProvider) && has(self.initProvider.dbUser))",message="spec.forProvider.dbUser is a required parameter"
	Spec   RDSDBInstanceSpec   `json:"spec"`
	Status RDSDBInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RDSDBInstanceList contains a list of RDSDBInstances
type RDSDBInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RDSDBInstance `json:"items"`
}

// Repository type metadata.
var (
	RDSDBInstance_Kind             = "RDSDBInstance"
	RDSDBInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RDSDBInstance_Kind}.String()
	RDSDBInstance_KindAPIVersion   = RDSDBInstance_Kind + "." + CRDGroupVersion.String()
	RDSDBInstance_GroupVersionKind = CRDGroupVersion.WithKind(RDSDBInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&RDSDBInstance{}, &RDSDBInstanceList{})
}
