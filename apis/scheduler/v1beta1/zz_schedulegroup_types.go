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

type ScheduleGroupInitParameters struct {

	// Name of the schedule group. Conflicts with name_prefix.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ScheduleGroupObservation struct {

	// ARN of the schedule group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Time at which the schedule group was created.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// Name of the schedule group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Time at which the schedule group was last modified.
	LastModificationDate *string `json:"lastModificationDate,omitempty" tf:"last_modification_date,omitempty"`

	// Name of the schedule group. Conflicts with name_prefix.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// State of the schedule group. Can be ACTIVE or DELETING.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ScheduleGroupParameters struct {

	// Name of the schedule group. Conflicts with name_prefix.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ScheduleGroupSpec defines the desired state of ScheduleGroup
type ScheduleGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScheduleGroupParameters `json:"forProvider"`
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
	InitProvider ScheduleGroupInitParameters `json:"initProvider,omitempty"`
}

// ScheduleGroupStatus defines the observed state of ScheduleGroup.
type ScheduleGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScheduleGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ScheduleGroup is the Schema for the ScheduleGroups API. Provides an EventBridge Scheduler Schedule Group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ScheduleGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScheduleGroupSpec   `json:"spec"`
	Status            ScheduleGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduleGroupList contains a list of ScheduleGroups
type ScheduleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScheduleGroup `json:"items"`
}

// Repository type metadata.
var (
	ScheduleGroup_Kind             = "ScheduleGroup"
	ScheduleGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScheduleGroup_Kind}.String()
	ScheduleGroup_KindAPIVersion   = ScheduleGroup_Kind + "." + CRDGroupVersion.String()
	ScheduleGroup_GroupVersionKind = CRDGroupVersion.WithKind(ScheduleGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ScheduleGroup{}, &ScheduleGroupList{})
}
