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

type ContainerPolicyInitParameters struct {

	// The name of the container.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/mediastore/v1beta1.Container
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// Reference to a Container in mediastore to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameRef *v1.Reference `json:"containerNameRef,omitempty" tf:"-"`

	// Selector for a Container in mediastore to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameSelector *v1.Selector `json:"containerNameSelector,omitempty" tf:"-"`

	// The contents of the policy.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type ContainerPolicyObservation struct {

	// The name of the container.
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The contents of the policy.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type ContainerPolicyParameters struct {

	// The name of the container.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/mediastore/v1beta1.Container
	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// Reference to a Container in mediastore to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameRef *v1.Reference `json:"containerNameRef,omitempty" tf:"-"`

	// Selector for a Container in mediastore to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameSelector *v1.Selector `json:"containerNameSelector,omitempty" tf:"-"`

	// The contents of the policy.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ContainerPolicySpec defines the desired state of ContainerPolicy
type ContainerPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerPolicyParameters `json:"forProvider"`
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
	InitProvider ContainerPolicyInitParameters `json:"initProvider,omitempty"`
}

// ContainerPolicyStatus defines the observed state of ContainerPolicy.
type ContainerPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ContainerPolicy is the Schema for the ContainerPolicys API. Provides a MediaStore Container Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ContainerPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   ContainerPolicySpec   `json:"spec"`
	Status ContainerPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerPolicyList contains a list of ContainerPolicys
type ContainerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerPolicy `json:"items"`
}

// Repository type metadata.
var (
	ContainerPolicy_Kind             = "ContainerPolicy"
	ContainerPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContainerPolicy_Kind}.String()
	ContainerPolicy_KindAPIVersion   = ContainerPolicy_Kind + "." + CRDGroupVersion.String()
	ContainerPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ContainerPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ContainerPolicy{}, &ContainerPolicyList{})
}
