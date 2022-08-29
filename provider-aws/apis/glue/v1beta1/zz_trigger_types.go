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

type ActionsNotificationPropertyObservation struct {
}

type ActionsNotificationPropertyParameters struct {

	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	// +kubebuilder:validation:Optional
	NotifyDelayAfter *float64 `json:"notifyDelayAfter,omitempty" tf:"notify_delay_after,omitempty"`
}

type ActionsObservation struct {
}

type ActionsParameters struct {

	// Arguments to be passed to the job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes.
	// +kubebuilder:validation:Optional
	Arguments map[string]*string `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// +kubebuilder:validation:Optional
	CrawlerName *string `json:"crawlerName,omitempty" tf:"crawler_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/glue/v1beta1.Job
	// +kubebuilder:validation:Optional
	JobName *string `json:"jobName,omitempty" tf:"job_name,omitempty"`

	// Reference to a Job in glue to populate jobName.
	// +kubebuilder:validation:Optional
	JobNameRef *v1.Reference `json:"jobNameRef,omitempty" tf:"-"`

	// Selector for a Job in glue to populate jobName.
	// +kubebuilder:validation:Optional
	JobNameSelector *v1.Selector `json:"jobNameSelector,omitempty" tf:"-"`

	// Specifies configuration properties of a job run notification. See Notification Property details below.
	// +kubebuilder:validation:Optional
	NotificationProperty []ActionsNotificationPropertyParameters `json:"notificationProperty,omitempty" tf:"notification_property,omitempty"`

	// The name of the Security Configuration structure to be used with this action.
	// +kubebuilder:validation:Optional
	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`

	// The job run timeout in minutes. It overrides the timeout value of the job.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ConditionsObservation struct {
}

type ConditionsParameters struct {

	// The condition crawl state. Currently, the values supported are RUNNING, SUCCEEDED, CANCELLED, and FAILED. If this is specified, crawler_name must also be specified. Conflicts with state.
	// +kubebuilder:validation:Optional
	CrawlState *string `json:"crawlState,omitempty" tf:"crawl_state,omitempty"`

	// +kubebuilder:validation:Optional
	CrawlerName *string `json:"crawlerName,omitempty" tf:"crawler_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/glue/v1beta1.Job
	// +kubebuilder:validation:Optional
	JobName *string `json:"jobName,omitempty" tf:"job_name,omitempty"`

	// Reference to a Job in glue to populate jobName.
	// +kubebuilder:validation:Optional
	JobNameRef *v1.Reference `json:"jobNameRef,omitempty" tf:"-"`

	// Selector for a Job in glue to populate jobName.
	// +kubebuilder:validation:Optional
	JobNameSelector *v1.Selector `json:"jobNameSelector,omitempty" tf:"-"`

	// A logical operator. Defaults to EQUALS.
	// +kubebuilder:validation:Optional
	LogicalOperator *string `json:"logicalOperator,omitempty" tf:"logical_operator,omitempty"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type EventBatchingConditionObservation struct {
}

type EventBatchingConditionParameters struct {

	// Number of events that must be received from Amazon EventBridge before EventBridge  event trigger fires.
	// +kubebuilder:validation:Required
	BatchSize *float64 `json:"batchSize" tf:"batch_size,omitempty"`

	// Window of time in seconds after which EventBridge event trigger fires. Window starts when first event is received. Default value is 900.
	// +kubebuilder:validation:Optional
	BatchWindow *float64 `json:"batchWindow,omitempty" tf:"batch_window,omitempty"`
}

type PredicateObservation struct {
}

type PredicateParameters struct {

	// A list of the conditions that determine when the trigger will fire. See Conditions.
	// +kubebuilder:validation:Required
	Conditions []ConditionsParameters `json:"conditions" tf:"conditions,omitempty"`

	// How to handle multiple conditions. Defaults to AND. Valid values are AND or ANY.
	// +kubebuilder:validation:Optional
	Logical *string `json:"logical,omitempty" tf:"logical,omitempty"`
}

type TriggerObservation struct {

	// Amazon Resource Name  of Glue Trigger
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Trigger name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TriggerParameters struct {

	// –  List of actions initiated by this trigger when it fires. See Actions Below.
	// +kubebuilder:validation:Required
	Actions []ActionsParameters `json:"actions" tf:"actions,omitempty"`

	// –  A description of the new trigger.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  Start the trigger. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Batch condition that must be met  before EventBridge event trigger fires. See Event Batching Condition.
	// +kubebuilder:validation:Optional
	EventBatchingCondition []EventBatchingConditionParameters `json:"eventBatchingCondition,omitempty" tf:"event_batching_condition,omitempty"`

	// –  A predicate to specify when the new trigger should fire. Required when trigger type is CONDITIONAL. See Predicate Below.
	// +kubebuilder:validation:Optional
	Predicate []PredicateParameters `json:"predicate,omitempty" tf:"predicate,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ased Schedules for Jobs and Crawlers
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// –  Set to true to start SCHEDULED and CONDITIONAL triggers when created. True is not supported for ON_DEMAND triggers.
	// +kubebuilder:validation:Optional
	StartOnCreation *bool `json:"startOnCreation,omitempty" tf:"start_on_creation,omitempty"`

	// Key-value map of resource tags. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// –  The type of trigger. Valid values are CONDITIONAL, ON_DEMAND, and SCHEDULED.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// A workflow to which the trigger should be associated to. Every workflow graph  needs a starting trigger  and can contain multiple additional CONDITIONAL triggers.
	// +kubebuilder:validation:Optional
	WorkflowName *string `json:"workflowName,omitempty" tf:"workflow_name,omitempty"`
}

// TriggerSpec defines the desired state of Trigger
type TriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerParameters `json:"forProvider"`
}

// TriggerStatus defines the observed state of Trigger.
type TriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Trigger is the Schema for the Triggers API. Manages a Glue Trigger resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Trigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerSpec   `json:"spec"`
	Status            TriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerList contains a list of Triggers
type TriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Trigger `json:"items"`
}

// Repository type metadata.
var (
	Trigger_Kind             = "Trigger"
	Trigger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Trigger_Kind}.String()
	Trigger_KindAPIVersion   = Trigger_Kind + "." + CRDGroupVersion.String()
	Trigger_GroupVersionKind = CRDGroupVersion.WithKind(Trigger_Kind)
)

func init() {
	SchemeBuilder.Register(&Trigger{}, &TriggerList{})
}
