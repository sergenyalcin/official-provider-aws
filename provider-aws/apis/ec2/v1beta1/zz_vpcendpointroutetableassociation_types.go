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

type VPCEndpointRouteTableAssociationObservation struct {

	// A hash of the EC2 Route Table and VPC Endpoint identifiers.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPCEndpointRouteTableAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.RouteTable
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a RouteTable in ec2 to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a RouteTable in ec2 to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.VPCEndpoint
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// Reference to a VPCEndpoint in ec2 to populate vpcEndpointId.
	// +kubebuilder:validation:Optional
	VPCEndpointIDRef *v1.Reference `json:"vpcEndpointIdRef,omitempty" tf:"-"`

	// Selector for a VPCEndpoint in ec2 to populate vpcEndpointId.
	// +kubebuilder:validation:Optional
	VPCEndpointIDSelector *v1.Selector `json:"vpcEndpointIdSelector,omitempty" tf:"-"`
}

// VPCEndpointRouteTableAssociationSpec defines the desired state of VPCEndpointRouteTableAssociation
type VPCEndpointRouteTableAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointRouteTableAssociationParameters `json:"forProvider"`
}

// VPCEndpointRouteTableAssociationStatus defines the observed state of VPCEndpointRouteTableAssociation.
type VPCEndpointRouteTableAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointRouteTableAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointRouteTableAssociation is the Schema for the VPCEndpointRouteTableAssociations API. Manages a VPC Endpoint Route Table Association
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCEndpointRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCEndpointRouteTableAssociationSpec   `json:"spec"`
	Status            VPCEndpointRouteTableAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointRouteTableAssociationList contains a list of VPCEndpointRouteTableAssociations
type VPCEndpointRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpointRouteTableAssociation `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpointRouteTableAssociation_Kind             = "VPCEndpointRouteTableAssociation"
	VPCEndpointRouteTableAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpointRouteTableAssociation_Kind}.String()
	VPCEndpointRouteTableAssociation_KindAPIVersion   = VPCEndpointRouteTableAssociation_Kind + "." + CRDGroupVersion.String()
	VPCEndpointRouteTableAssociation_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpointRouteTableAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpointRouteTableAssociation{}, &VPCEndpointRouteTableAssociationList{})
}
