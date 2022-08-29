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

type InvocationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// String result of the lambda function invocation.
	Result *string `json:"result,omitempty" tf:"result,omitempty"`
}

type InvocationParameters struct {

	// Name of the lambda function.
	// +crossplane:generate:reference:type=Function
	// +kubebuilder:validation:Optional
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// Reference to a Function to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameRef *v1.Reference `json:"functionNameRef,omitempty" tf:"-"`

	// Selector for a Function to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameSelector *v1.Selector `json:"functionNameSelector,omitempty" tf:"-"`

	// JSON payload to the lambda function.
	// +kubebuilder:validation:Required
	Input *string `json:"input" tf:"input,omitempty"`

	// Qualifier  of the lambda function. Defaults to $LATEST.
	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Map of arbitrary keys and values that, when changed, will trigger a re-invocation. To force a re-invocation without changing these keys/values, use the terraform taint command.
	// +kubebuilder:validation:Optional
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

// InvocationSpec defines the desired state of Invocation
type InvocationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InvocationParameters `json:"forProvider"`
}

// InvocationStatus defines the observed state of Invocation.
type InvocationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InvocationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Invocation is the Schema for the Invocations API. Invoke AWS Lambda Function
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Invocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InvocationSpec   `json:"spec"`
	Status            InvocationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InvocationList contains a list of Invocations
type InvocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Invocation `json:"items"`
}

// Repository type metadata.
var (
	Invocation_Kind             = "Invocation"
	Invocation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Invocation_Kind}.String()
	Invocation_KindAPIVersion   = Invocation_Kind + "." + CRDGroupVersion.String()
	Invocation_GroupVersionKind = CRDGroupVersion.WithKind(Invocation_Kind)
)

func init() {
	SchemeBuilder.Register(&Invocation{}, &InvocationList{})
}
