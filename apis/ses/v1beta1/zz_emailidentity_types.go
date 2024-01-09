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

type EmailIdentityInitParameters struct {

	// The email address to assign to SES.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`
}

type EmailIdentityObservation struct {

	// The ARN of the email identity.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The email address to assign to SES.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EmailIdentityParameters struct {

	// The email address to assign to SES.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// EmailIdentitySpec defines the desired state of EmailIdentity
type EmailIdentitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmailIdentityParameters `json:"forProvider"`
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
	InitProvider EmailIdentityInitParameters `json:"initProvider,omitempty"`
}

// EmailIdentityStatus defines the observed state of EmailIdentity.
type EmailIdentityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmailIdentityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EmailIdentity is the Schema for the EmailIdentitys API. Provides an SES email identity resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EmailIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	Spec   EmailIdentitySpec   `json:"spec"`
	Status EmailIdentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailIdentityList contains a list of EmailIdentitys
type EmailIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailIdentity `json:"items"`
}

// Repository type metadata.
var (
	EmailIdentity_Kind             = "EmailIdentity"
	EmailIdentity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmailIdentity_Kind}.String()
	EmailIdentity_KindAPIVersion   = EmailIdentity_Kind + "." + CRDGroupVersion.String()
	EmailIdentity_GroupVersionKind = CRDGroupVersion.WithKind(EmailIdentity_Kind)
)

func init() {
	SchemeBuilder.Register(&EmailIdentity{}, &EmailIdentityList{})
}
