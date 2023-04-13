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

type CertificateObservation struct {

	// ARN of the certificate.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// PEM-encoded certificate value.
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// ARN of the certificate authority.
	CertificateAuthorityArn *string `json:"certificateAuthorityArn,omitempty" tf:"certificate_authority_arn,omitempty"`

	// PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA.
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Algorithm to use to sign certificate requests. Valid values: SHA256WITHRSA, SHA256WITHECDSA, SHA384WITHRSA, SHA384WITHECDSA, SHA512WITHRSA, SHA512WITHECDSA.
	SigningAlgorithm *string `json:"signingAlgorithm,omitempty" tf:"signing_algorithm,omitempty"`

	// Template to use when issuing a certificate.
	// See ACM PCA Documentation for more information.
	TemplateArn *string `json:"templateArn,omitempty" tf:"template_arn,omitempty"`

	// Configures end of the validity period for the certificate. See validity block below.
	Validity []ValidityObservation `json:"validity,omitempty" tf:"validity,omitempty"`
}

type CertificateParameters struct {

	// ARN of the certificate authority.
	// +crossplane:generate:reference:type=CertificateAuthority
	// +kubebuilder:validation:Optional
	CertificateAuthorityArn *string `json:"certificateAuthorityArn,omitempty" tf:"certificate_authority_arn,omitempty"`

	// Reference to a CertificateAuthority to populate certificateAuthorityArn.
	// +kubebuilder:validation:Optional
	CertificateAuthorityArnRef *v1.Reference `json:"certificateAuthorityArnRef,omitempty" tf:"-"`

	// Selector for a CertificateAuthority to populate certificateAuthorityArn.
	// +kubebuilder:validation:Optional
	CertificateAuthorityArnSelector *v1.Selector `json:"certificateAuthorityArnSelector,omitempty" tf:"-"`

	// Certificate Signing Request in PEM format.
	// +kubebuilder:validation:Optional
	CertificateSigningRequestSecretRef v1.SecretKeySelector `json:"certificateSigningRequestSecretRef" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Algorithm to use to sign certificate requests. Valid values: SHA256WITHRSA, SHA256WITHECDSA, SHA384WITHRSA, SHA384WITHECDSA, SHA512WITHRSA, SHA512WITHECDSA.
	// +kubebuilder:validation:Optional
	SigningAlgorithm *string `json:"signingAlgorithm,omitempty" tf:"signing_algorithm,omitempty"`

	// Template to use when issuing a certificate.
	// See ACM PCA Documentation for more information.
	// +kubebuilder:validation:Optional
	TemplateArn *string `json:"templateArn,omitempty" tf:"template_arn,omitempty"`

	// Configures end of the validity period for the certificate. See validity block below.
	// +kubebuilder:validation:Optional
	Validity []ValidityParameters `json:"validity,omitempty" tf:"validity,omitempty"`
}

type ValidityObservation struct {

	// Determines how value is interpreted. Valid values: DAYS, MONTHS, YEARS, ABSOLUTE, END_DATE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// If type is DAYS, MONTHS, or YEARS, the relative time until the certificate expires. If type is ABSOLUTE, the date in seconds since the Unix epoch. If type is END_DATE, the  date in RFC 3339 format.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ValidityParameters struct {

	// Determines how value is interpreted. Valid values: DAYS, MONTHS, YEARS, ABSOLUTE, END_DATE.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// If type is DAYS, MONTHS, or YEARS, the relative time until the certificate expires. If type is ABSOLUTE, the date in seconds since the Unix epoch. If type is END_DATE, the  date in RFC 3339 format.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API. Provides a resource to issue a certificate using AWS Certificate Manager Private Certificate Authority (ACM PCA)
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.certificateSigningRequestSecretRef)",message="certificateSigningRequestSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.signingAlgorithm)",message="signingAlgorithm is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.validity)",message="validity is a required parameter"
	Spec   CertificateSpec   `json:"spec"`
	Status CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
