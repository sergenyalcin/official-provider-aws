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

type GlobalNetworkObservation struct {

	// Global Network Amazon Resource Name (ARN)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the Global Network.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type GlobalNetworkParameters struct {

	// Description of the Global Network.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// GlobalNetworkSpec defines the desired state of GlobalNetwork
type GlobalNetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalNetworkParameters `json:"forProvider"`
}

// GlobalNetworkStatus defines the observed state of GlobalNetwork.
type GlobalNetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalNetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalNetwork is the Schema for the GlobalNetworks API. Provides a global network resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GlobalNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalNetworkSpec   `json:"spec"`
	Status            GlobalNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalNetworkList contains a list of GlobalNetworks
type GlobalNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalNetwork `json:"items"`
}

// Repository type metadata.
var (
	GlobalNetwork_Kind             = "GlobalNetwork"
	GlobalNetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalNetwork_Kind}.String()
	GlobalNetwork_KindAPIVersion   = GlobalNetwork_Kind + "." + CRDGroupVersion.String()
	GlobalNetwork_GroupVersionKind = CRDGroupVersion.WithKind(GlobalNetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalNetwork{}, &GlobalNetworkList{})
}
