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

type ProxyTargetObservation struct {

	// Hostname for the target RDS DB Instance. Only returned for RDS_INSTANCE type.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Identifier of  db_proxy_name, target_group_name, target type , and resource identifier separated by forward slashes .
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Port for the target RDS DB Instance or Aurora DB Cluster.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Identifier representing the DB Instance or DB Cluster target.
	RDSResourceID *string `json:"rdsResourceId,omitempty" tf:"rds_resource_id,omitempty"`

	// Amazon Resource Name  for the DB instance or DB cluster. Currently not returned by the RDS API.
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`

	// DB Cluster identifier for the DB Instance target. Not returned unless manually importing an RDS_INSTANCE target that is part of a DB Cluster.
	TrackedClusterID *string `json:"trackedClusterId,omitempty" tf:"tracked_cluster_id,omitempty"`

	// Type of targetE.g., RDS_INSTANCE or TRACKED_CLUSTER
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProxyTargetParameters struct {

	// DB cluster identifier.
	// +kubebuilder:validation:Optional
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty" tf:"db_cluster_identifier,omitempty"`

	// DB instance identifier.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/rds/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DBInstanceIdentifier *string `json:"dbInstanceIdentifier,omitempty" tf:"db_instance_identifier,omitempty"`

	// Reference to a Instance in rds to populate dbInstanceIdentifier.
	// +kubebuilder:validation:Optional
	DBInstanceIdentifierRef *v1.Reference `json:"dbInstanceIdentifierRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate dbInstanceIdentifier.
	// +kubebuilder:validation:Optional
	DBInstanceIdentifierSelector *v1.Selector `json:"dbInstanceIdentifierSelector,omitempty" tf:"-"`

	// The name of the DB proxy.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/rds/v1beta1.Proxy
	// +kubebuilder:validation:Optional
	DBProxyName *string `json:"dbProxyName,omitempty" tf:"db_proxy_name,omitempty"`

	// Reference to a Proxy in rds to populate dbProxyName.
	// +kubebuilder:validation:Optional
	DBProxyNameRef *v1.Reference `json:"dbProxyNameRef,omitempty" tf:"-"`

	// Selector for a Proxy in rds to populate dbProxyName.
	// +kubebuilder:validation:Optional
	DBProxyNameSelector *v1.Selector `json:"dbProxyNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the target group.
	// +kubebuilder:validation:Required
	TargetGroupName *string `json:"targetGroupName" tf:"target_group_name,omitempty"`
}

// ProxyTargetSpec defines the desired state of ProxyTarget
type ProxyTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProxyTargetParameters `json:"forProvider"`
}

// ProxyTargetStatus defines the observed state of ProxyTarget.
type ProxyTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProxyTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyTarget is the Schema for the ProxyTargets API. Provides an RDS DB proxy target resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ProxyTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProxyTargetSpec   `json:"spec"`
	Status            ProxyTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyTargetList contains a list of ProxyTargets
type ProxyTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProxyTarget `json:"items"`
}

// Repository type metadata.
var (
	ProxyTarget_Kind             = "ProxyTarget"
	ProxyTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProxyTarget_Kind}.String()
	ProxyTarget_KindAPIVersion   = ProxyTarget_Kind + "." + CRDGroupVersion.String()
	ProxyTarget_GroupVersionKind = CRDGroupVersion.WithKind(ProxyTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&ProxyTarget{}, &ProxyTargetList{})
}
