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

type ConformancePackInitParameters struct {

	// Amazon S3 bucket where AWS Config stores conformance pack templates. Maximum length of 63.
	DeliveryS3Bucket *string `json:"deliveryS3Bucket,omitempty" tf:"delivery_s3_bucket,omitempty"`

	// The prefix for the Amazon S3 bucket. Maximum length of 1024.
	DeliveryS3KeyPrefix *string `json:"deliveryS3KeyPrefix,omitempty" tf:"delivery_s3_key_prefix,omitempty"`

	// Set of configuration blocks describing input parameters passed to the conformance pack template. Documented below. When configured, the parameters must also be included in the template_body or in the template stored in Amazon S3 if using template_s3_uri.
	InputParameter []InputParameterInitParameters `json:"inputParameter,omitempty" tf:"input_parameter,omitempty"`

	// A string containing full conformance pack template body. Maximum length of 51200. Drift detection is not possible with this argument.
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// Location of file, e.g., s3://bucketname/prefix, containing the template body. The uri must point to the conformance pack template that is located in an Amazon S3 bucket in the same region as the conformance pack. Maximum length of 1024. Drift detection is not possible with this argument.
	TemplateS3URI *string `json:"templateS3Uri,omitempty" tf:"template_s3_uri,omitempty"`
}

type ConformancePackObservation struct {

	// Amazon Resource Name (ARN) of the conformance pack.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Amazon S3 bucket where AWS Config stores conformance pack templates. Maximum length of 63.
	DeliveryS3Bucket *string `json:"deliveryS3Bucket,omitempty" tf:"delivery_s3_bucket,omitempty"`

	// The prefix for the Amazon S3 bucket. Maximum length of 1024.
	DeliveryS3KeyPrefix *string `json:"deliveryS3KeyPrefix,omitempty" tf:"delivery_s3_key_prefix,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of configuration blocks describing input parameters passed to the conformance pack template. Documented below. When configured, the parameters must also be included in the template_body or in the template stored in Amazon S3 if using template_s3_uri.
	InputParameter []InputParameterObservation `json:"inputParameter,omitempty" tf:"input_parameter,omitempty"`

	// A string containing full conformance pack template body. Maximum length of 51200. Drift detection is not possible with this argument.
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// Location of file, e.g., s3://bucketname/prefix, containing the template body. The uri must point to the conformance pack template that is located in an Amazon S3 bucket in the same region as the conformance pack. Maximum length of 1024. Drift detection is not possible with this argument.
	TemplateS3URI *string `json:"templateS3Uri,omitempty" tf:"template_s3_uri,omitempty"`
}

type ConformancePackParameters struct {

	// Amazon S3 bucket where AWS Config stores conformance pack templates. Maximum length of 63.
	// +kubebuilder:validation:Optional
	DeliveryS3Bucket *string `json:"deliveryS3Bucket,omitempty" tf:"delivery_s3_bucket,omitempty"`

	// The prefix for the Amazon S3 bucket. Maximum length of 1024.
	// +kubebuilder:validation:Optional
	DeliveryS3KeyPrefix *string `json:"deliveryS3KeyPrefix,omitempty" tf:"delivery_s3_key_prefix,omitempty"`

	// Set of configuration blocks describing input parameters passed to the conformance pack template. Documented below. When configured, the parameters must also be included in the template_body or in the template stored in Amazon S3 if using template_s3_uri.
	// +kubebuilder:validation:Optional
	InputParameter []InputParameterParameters `json:"inputParameter,omitempty" tf:"input_parameter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A string containing full conformance pack template body. Maximum length of 51200. Drift detection is not possible with this argument.
	// +kubebuilder:validation:Optional
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// Location of file, e.g., s3://bucketname/prefix, containing the template body. The uri must point to the conformance pack template that is located in an Amazon S3 bucket in the same region as the conformance pack. Maximum length of 1024. Drift detection is not possible with this argument.
	// +kubebuilder:validation:Optional
	TemplateS3URI *string `json:"templateS3Uri,omitempty" tf:"template_s3_uri,omitempty"`
}

type InputParameterInitParameters struct {

	// The input key.
	ParameterName *string `json:"parameterName,omitempty" tf:"parameter_name,omitempty"`

	// The input value.
	ParameterValue *string `json:"parameterValue,omitempty" tf:"parameter_value,omitempty"`
}

type InputParameterObservation struct {

	// The input key.
	ParameterName *string `json:"parameterName,omitempty" tf:"parameter_name,omitempty"`

	// The input value.
	ParameterValue *string `json:"parameterValue,omitempty" tf:"parameter_value,omitempty"`
}

type InputParameterParameters struct {

	// The input key.
	// +kubebuilder:validation:Optional
	ParameterName *string `json:"parameterName" tf:"parameter_name,omitempty"`

	// The input value.
	// +kubebuilder:validation:Optional
	ParameterValue *string `json:"parameterValue" tf:"parameter_value,omitempty"`
}

// ConformancePackSpec defines the desired state of ConformancePack
type ConformancePackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConformancePackParameters `json:"forProvider"`
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
	InitProvider ConformancePackInitParameters `json:"initProvider,omitempty"`
}

// ConformancePackStatus defines the observed state of ConformancePack.
type ConformancePackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConformancePackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ConformancePack is the Schema for the ConformancePacks API. Manages a Config Conformance Pack
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConformancePack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConformancePackSpec   `json:"spec"`
	Status            ConformancePackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConformancePackList contains a list of ConformancePacks
type ConformancePackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConformancePack `json:"items"`
}

// Repository type metadata.
var (
	ConformancePack_Kind             = "ConformancePack"
	ConformancePack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConformancePack_Kind}.String()
	ConformancePack_KindAPIVersion   = ConformancePack_Kind + "." + CRDGroupVersion.String()
	ConformancePack_GroupVersionKind = CRDGroupVersion.WithKind(ConformancePack_Kind)
)

func init() {
	SchemeBuilder.Register(&ConformancePack{}, &ConformancePackList{})
}
