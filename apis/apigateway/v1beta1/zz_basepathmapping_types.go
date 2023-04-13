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

type BasePathMappingObservation struct {

	// ID of the API to connect.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
	BasePath *string `json:"basePath,omitempty" tf:"base_path,omitempty"`

	// Already-registered domain name to connect the API to.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
	StageName *string `json:"stageName,omitempty" tf:"stage_name,omitempty"`
}

type BasePathMappingParameters struct {

	// ID of the API to connect.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
	// +kubebuilder:validation:Optional
	BasePath *string `json:"basePath,omitempty" tf:"base_path,omitempty"`

	// Already-registered domain name to connect the API to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.DomainName
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("domain_name",false)
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Reference to a DomainName in apigateway to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameRef *v1.Reference `json:"domainNameRef,omitempty" tf:"-"`

	// Selector for a DomainName in apigateway to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameSelector *v1.Selector `json:"domainNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Stage
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("stage_name",false)
	// +kubebuilder:validation:Optional
	StageName *string `json:"stageName,omitempty" tf:"stage_name,omitempty"`

	// Reference to a Stage in apigateway to populate stageName.
	// +kubebuilder:validation:Optional
	StageNameRef *v1.Reference `json:"stageNameRef,omitempty" tf:"-"`

	// Selector for a Stage in apigateway to populate stageName.
	// +kubebuilder:validation:Optional
	StageNameSelector *v1.Selector `json:"stageNameSelector,omitempty" tf:"-"`
}

// BasePathMappingSpec defines the desired state of BasePathMapping
type BasePathMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BasePathMappingParameters `json:"forProvider"`
}

// BasePathMappingStatus defines the observed state of BasePathMapping.
type BasePathMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BasePathMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BasePathMapping is the Schema for the BasePathMappings API. Connects a custom domain with a deployed API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BasePathMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BasePathMappingSpec   `json:"spec"`
	Status            BasePathMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BasePathMappingList contains a list of BasePathMappings
type BasePathMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BasePathMapping `json:"items"`
}

// Repository type metadata.
var (
	BasePathMapping_Kind             = "BasePathMapping"
	BasePathMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BasePathMapping_Kind}.String()
	BasePathMapping_KindAPIVersion   = BasePathMapping_Kind + "." + CRDGroupVersion.String()
	BasePathMapping_GroupVersionKind = CRDGroupVersion.WithKind(BasePathMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&BasePathMapping{}, &BasePathMappingList{})
}
