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

type TrafficPolicyInstanceObservation struct {

	// ID of the hosted zone that you want Amazon Route 53 to create resource record sets in by using the configuration in a traffic policy.
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// ID of traffic policy instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Domain name for which Amazon Route 53 responds to DNS queries by using the resource record sets that Route 53 creates for this traffic policy instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// TTL that you want Amazon Route 53 to assign to all the resource record sets that it creates in the specified hosted zone.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// ID of the traffic policy that you want to use to create resource record sets in the specified hosted zone.
	TrafficPolicyID *string `json:"trafficPolicyId,omitempty" tf:"traffic_policy_id,omitempty"`

	// Version of the traffic policy
	TrafficPolicyVersion *float64 `json:"trafficPolicyVersion,omitempty" tf:"traffic_policy_version,omitempty"`
}

type TrafficPolicyInstanceParameters struct {

	// ID of the hosted zone that you want Amazon Route 53 to create resource record sets in by using the configuration in a traffic policy.
	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// Reference to a Zone to populate hostedZoneId.
	// +kubebuilder:validation:Optional
	HostedZoneIDRef *v1.Reference `json:"hostedZoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone to populate hostedZoneId.
	// +kubebuilder:validation:Optional
	HostedZoneIDSelector *v1.Selector `json:"hostedZoneIdSelector,omitempty" tf:"-"`

	// Domain name for which Amazon Route 53 responds to DNS queries by using the resource record sets that Route 53 creates for this traffic policy instance.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// TTL that you want Amazon Route 53 to assign to all the resource record sets that it creates in the specified hosted zone.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// ID of the traffic policy that you want to use to create resource record sets in the specified hosted zone.
	// +crossplane:generate:reference:type=TrafficPolicy
	// +kubebuilder:validation:Optional
	TrafficPolicyID *string `json:"trafficPolicyId,omitempty" tf:"traffic_policy_id,omitempty"`

	// Reference to a TrafficPolicy to populate trafficPolicyId.
	// +kubebuilder:validation:Optional
	TrafficPolicyIDRef *v1.Reference `json:"trafficPolicyIdRef,omitempty" tf:"-"`

	// Selector for a TrafficPolicy to populate trafficPolicyId.
	// +kubebuilder:validation:Optional
	TrafficPolicyIDSelector *v1.Selector `json:"trafficPolicyIdSelector,omitempty" tf:"-"`

	// Version of the traffic policy
	// +kubebuilder:validation:Optional
	TrafficPolicyVersion *float64 `json:"trafficPolicyVersion,omitempty" tf:"traffic_policy_version,omitempty"`
}

// TrafficPolicyInstanceSpec defines the desired state of TrafficPolicyInstance
type TrafficPolicyInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficPolicyInstanceParameters `json:"forProvider"`
}

// TrafficPolicyInstanceStatus defines the observed state of TrafficPolicyInstance.
type TrafficPolicyInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficPolicyInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficPolicyInstance is the Schema for the TrafficPolicyInstances API. Provides a Route53 traffic policy instance resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TrafficPolicyInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.trafficPolicyVersion)",message="trafficPolicyVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.ttl)",message="ttl is a required parameter"
	Spec   TrafficPolicyInstanceSpec   `json:"spec"`
	Status TrafficPolicyInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficPolicyInstanceList contains a list of TrafficPolicyInstances
type TrafficPolicyInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficPolicyInstance `json:"items"`
}

// Repository type metadata.
var (
	TrafficPolicyInstance_Kind             = "TrafficPolicyInstance"
	TrafficPolicyInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficPolicyInstance_Kind}.String()
	TrafficPolicyInstance_KindAPIVersion   = TrafficPolicyInstance_Kind + "." + CRDGroupVersion.String()
	TrafficPolicyInstance_GroupVersionKind = CRDGroupVersion.WithKind(TrafficPolicyInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficPolicyInstance{}, &TrafficPolicyInstanceList{})
}
