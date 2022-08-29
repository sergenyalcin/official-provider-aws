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

type PermissionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PermissionParameters struct {

	// The AWS Lambda action you want to allow in this statement.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// The Event Source Token to validate.  Used with Alexa Skills.
	// +kubebuilder:validation:Optional
	EventSourceToken *string `json:"eventSourceToken,omitempty" tf:"event_source_token,omitempty"`

	// Name of the Lambda function whose resource policy you are updating
	// +crossplane:generate:reference:type=Function
	// +kubebuilder:validation:Optional
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// Reference to a Function to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameRef *v1.Reference `json:"functionNameRef,omitempty" tf:"-"`

	// Selector for a Function to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameSelector *v1.Selector `json:"functionNameSelector,omitempty" tf:"-"`

	// Lambda Function URLs authentication type. Valid values are: AWS_IAM or NONE.
	// +kubebuilder:validation:Optional
	FunctionURLAuthType *string `json:"functionUrlAuthType,omitempty" tf:"function_url_auth_type,omitempty"`

	// The principal who is getting this permission e.g., s3.amazonaws.com, an AWS account ID, or any valid AWS service principal such as events.amazonaws.com or sns.amazonaws.com.
	// +kubebuilder:validation:Required
	Principal *string `json:"principal" tf:"principal,omitempty"`

	// The identifier for your organization in AWS Organizations. Use this to grant permissions to all the AWS accounts under this organization.
	// +kubebuilder:validation:Optional
	PrincipalOrgID *string `json:"principalOrgId,omitempty" tf:"principal_org_id,omitempty"`

	// Query parameter to specify function version or alias name. The permission will then apply to the specific qualified ARN e.g., arn:aws:lambda:aws-region:acct-id:function:function-name:2
	// +crossplane:generate:reference:type=Alias
	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`

	// Reference to a Alias to populate qualifier.
	// +kubebuilder:validation:Optional
	QualifierRef *v1.Reference `json:"qualifierRef,omitempty" tf:"-"`

	// Selector for a Alias to populate qualifier.
	// +kubebuilder:validation:Optional
	QualifierSelector *v1.Selector `json:"qualifierSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// This parameter is used for S3 and SES. The AWS account ID  of the source owner.
	// +kubebuilder:validation:Optional
	SourceAccount *string `json:"sourceAccount,omitempty" tf:"source_account,omitempty"`

	// When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
	// Without this, any resource from principal will be granted permission – even if that resource is from another account.
	// For S3, this should be the ARN of the S3 Bucket.
	// For EventBridge events, this should be the ARN of the EventBridge Rule.
	// For API Gateway, this should be the ARN of the API, as described here.
	// +kubebuilder:validation:Optional
	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`

	// A unique statement identifier. By default generated by Terraform.
	// +kubebuilder:validation:Optional
	StatementID *string `json:"statementId,omitempty" tf:"statement_id,omitempty"`

	// A statement identifier prefix. Terraform will generate a unique suffix. Conflicts with statement_id.
	// +kubebuilder:validation:Optional
	StatementIDPrefix *string `json:"statementIdPrefix,omitempty" tf:"statement_id_prefix,omitempty"`
}

// PermissionSpec defines the desired state of Permission
type PermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionParameters `json:"forProvider"`
}

// PermissionStatus defines the observed state of Permission.
type PermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Permission is the Schema for the Permissions API. Creates a Lambda function permission.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Permission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionSpec   `json:"spec"`
	Status            PermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionList contains a list of Permissions
type PermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Permission `json:"items"`
}

// Repository type metadata.
var (
	Permission_Kind             = "Permission"
	Permission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Permission_Kind}.String()
	Permission_KindAPIVersion   = Permission_Kind + "." + CRDGroupVersion.String()
	Permission_GroupVersionKind = CRDGroupVersion.WithKind(Permission_Kind)
)

func init() {
	SchemeBuilder.Register(&Permission{}, &PermissionList{})
}
