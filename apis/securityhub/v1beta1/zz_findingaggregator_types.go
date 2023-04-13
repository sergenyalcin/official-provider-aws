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

type FindingAggregatorObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether to aggregate findings from all of the available Regions or from a specified list. The options are ALL_REGIONS, ALL_REGIONS_EXCEPT_SPECIFIED or SPECIFIED_REGIONS. When ALL_REGIONS or ALL_REGIONS_EXCEPT_SPECIFIED are used, Security Hub will automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.
	LinkingMode *string `json:"linkingMode,omitempty" tf:"linking_mode,omitempty"`

	// List of regions to include or exclude
	SpecifiedRegions []*string `json:"specifiedRegions,omitempty" tf:"specified_regions,omitempty"`
}

type FindingAggregatorParameters struct {

	// Indicates whether to aggregate findings from all of the available Regions or from a specified list. The options are ALL_REGIONS, ALL_REGIONS_EXCEPT_SPECIFIED or SPECIFIED_REGIONS. When ALL_REGIONS or ALL_REGIONS_EXCEPT_SPECIFIED are used, Security Hub will automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.
	// +kubebuilder:validation:Optional
	LinkingMode *string `json:"linkingMode,omitempty" tf:"linking_mode,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of regions to include or exclude
	// +kubebuilder:validation:Optional
	SpecifiedRegions []*string `json:"specifiedRegions,omitempty" tf:"specified_regions,omitempty"`
}

// FindingAggregatorSpec defines the desired state of FindingAggregator
type FindingAggregatorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FindingAggregatorParameters `json:"forProvider"`
}

// FindingAggregatorStatus defines the observed state of FindingAggregator.
type FindingAggregatorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FindingAggregatorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FindingAggregator is the Schema for the FindingAggregators API. Manages a Security Hub finding aggregator
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FindingAggregator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.linkingMode)",message="linkingMode is a required parameter"
	Spec   FindingAggregatorSpec   `json:"spec"`
	Status FindingAggregatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FindingAggregatorList contains a list of FindingAggregators
type FindingAggregatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FindingAggregator `json:"items"`
}

// Repository type metadata.
var (
	FindingAggregator_Kind             = "FindingAggregator"
	FindingAggregator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FindingAggregator_Kind}.String()
	FindingAggregator_KindAPIVersion   = FindingAggregator_Kind + "." + CRDGroupVersion.String()
	FindingAggregator_GroupVersionKind = CRDGroupVersion.WithKind(FindingAggregator_Kind)
)

func init() {
	SchemeBuilder.Register(&FindingAggregator{}, &FindingAggregatorList{})
}
