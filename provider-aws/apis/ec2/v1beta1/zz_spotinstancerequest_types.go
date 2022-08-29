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

type SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTargetObservation struct {
}

type SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTargetParameters struct {

	// +kubebuilder:validation:Optional
	CapacityReservationID *string `json:"capacityReservationId,omitempty" tf:"capacity_reservation_id,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityReservationResourceGroupArn *string `json:"capacityReservationResourceGroupArn,omitempty" tf:"capacity_reservation_resource_group_arn,omitempty"`
}

type SpotInstanceRequestCapacityReservationSpecificationObservation struct {
}

type SpotInstanceRequestCapacityReservationSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	CapacityReservationPreference *string `json:"capacityReservationPreference,omitempty" tf:"capacity_reservation_preference,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityReservationTarget []SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTargetParameters `json:"capacityReservationTarget,omitempty" tf:"capacity_reservation_target,omitempty"`
}

type SpotInstanceRequestCreditSpecificationObservation struct {
}

type SpotInstanceRequestCreditSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	CPUCredits *string `json:"cpuCredits,omitempty" tf:"cpu_credits,omitempty"`
}

type SpotInstanceRequestEBSBlockDeviceObservation struct {
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type SpotInstanceRequestEBSBlockDeviceParameters struct {

	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// A map of tags to assign to the Spot Instance Request. These tags are not automatically applied to the launched Instance. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type SpotInstanceRequestEnclaveOptionsObservation struct {
}

type SpotInstanceRequestEnclaveOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SpotInstanceRequestEphemeralBlockDeviceObservation struct {
}

type SpotInstanceRequestEphemeralBlockDeviceParameters struct {

	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type SpotInstanceRequestLaunchTemplateObservation struct {
}

type SpotInstanceRequestLaunchTemplateParameters struct {

	// The Spot Instance Request ID.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SpotInstanceRequestMaintenanceOptionsObservation struct {
}

type SpotInstanceRequestMaintenanceOptionsParameters struct {

	// +kubebuilder:validation:Optional
	AutoRecovery *string `json:"autoRecovery,omitempty" tf:"auto_recovery,omitempty"`
}

type SpotInstanceRequestMetadataOptionsObservation struct {
}

type SpotInstanceRequestMetadataOptionsParameters struct {

	// +kubebuilder:validation:Optional
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPPutResponseHopLimit *float64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPTokens *string `json:"httpTokens,omitempty" tf:"http_tokens,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceMetadataTags *string `json:"instanceMetadataTags,omitempty" tf:"instance_metadata_tags,omitempty"`
}

type SpotInstanceRequestNetworkInterfaceObservation struct {
}

type SpotInstanceRequestNetworkInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// +kubebuilder:validation:Required
	DeviceIndex *float64 `json:"deviceIndex" tf:"device_index,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkCardIndex *float64 `json:"networkCardIndex,omitempty" tf:"network_card_index,omitempty"`

	// +kubebuilder:validation:Required
	NetworkInterfaceID *string `json:"networkInterfaceId" tf:"network_interface_id,omitempty"`
}

type SpotInstanceRequestObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// +kubebuilder:validation:Optional
	EBSBlockDevice []SpotInstanceRequestEBSBlockDeviceObservation `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// The Spot Instance Request ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceState *string `json:"instanceState,omitempty" tf:"instance_state,omitempty"`

	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	PasswordData *string `json:"passwordData,omitempty" tf:"password_data,omitempty"`

	PrimaryNetworkInterfaceID *string `json:"primaryNetworkInterfaceId,omitempty" tf:"primary_network_interface_id,omitempty"`

	// The private DNS name assigned to the instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC
	PrivateDNS *string `json:"privateDns,omitempty" tf:"private_dns,omitempty"`

	// The public DNS name assigned to the instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC
	PublicDNS *string `json:"publicDns,omitempty" tf:"public_dns,omitempty"`

	// The public IP address assigned to the instance, if applicable.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// +kubebuilder:validation:Optional
	RootBlockDevice []SpotInstanceRequestRootBlockDeviceObservation `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`

	// The current bid
	// status
	// of the Spot Instance Request.
	SpotBidStatus *string `json:"spotBidStatus,omitempty" tf:"spot_bid_status,omitempty"`

	// The Instance ID  that is currently fulfilling
	// the Spot Instance request.
	SpotInstanceID *string `json:"spotInstanceId,omitempty" tf:"spot_instance_id,omitempty"`

	// The current request
	// state
	// of the Spot Instance Request.
	SpotRequestState *string `json:"spotRequestState,omitempty" tf:"spot_request_state,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SpotInstanceRequestParameters struct {

	// +kubebuilder:validation:Optional
	AMI *string `json:"ami,omitempty" tf:"ami,omitempty"`

	// +kubebuilder:validation:Optional
	AssociatePublicIPAddress *bool `json:"associatePublicIpAddress,omitempty" tf:"associate_public_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The required duration for the Spot instances, in minutes. This value must be a multiple of 60 .
	// The duration period starts as soon as your Spot instance receives its instance ID. At the end of the duration period, Amazon EC2 marks the Spot instance for termination and provides a Spot instance termination notice, which gives the instance a two-minute warning before it terminates.
	// Note that you can't specify an Availability Zone group or a launch group if you specify a duration.
	// +kubebuilder:validation:Optional
	BlockDurationMinutes *float64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	CPUCoreCount *float64 `json:"cpuCoreCount,omitempty" tf:"cpu_core_count,omitempty"`

	// +kubebuilder:validation:Optional
	CPUThreadsPerCore *float64 `json:"cpuThreadsPerCore,omitempty" tf:"cpu_threads_per_core,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityReservationSpecification []SpotInstanceRequestCapacityReservationSpecificationParameters `json:"capacityReservationSpecification,omitempty" tf:"capacity_reservation_specification,omitempty"`

	// +kubebuilder:validation:Optional
	CreditSpecification []SpotInstanceRequestCreditSpecificationParameters `json:"creditSpecification,omitempty" tf:"credit_specification,omitempty"`

	// +kubebuilder:validation:Optional
	DisableAPITermination *bool `json:"disableApiTermination,omitempty" tf:"disable_api_termination,omitempty"`

	// +kubebuilder:validation:Optional
	EBSBlockDevice []SpotInstanceRequestEBSBlockDeviceParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	EBSOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`

	// +kubebuilder:validation:Optional
	EnclaveOptions []SpotInstanceRequestEnclaveOptionsParameters `json:"enclaveOptions,omitempty" tf:"enclave_options,omitempty"`

	// +kubebuilder:validation:Optional
	EphemeralBlockDevice []SpotInstanceRequestEphemeralBlockDeviceParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	GetPasswordData *bool `json:"getPasswordData,omitempty" tf:"get_password_data,omitempty"`

	// +kubebuilder:validation:Optional
	Hibernation *bool `json:"hibernation,omitempty" tf:"hibernation,omitempty"`

	// +kubebuilder:validation:Optional
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// +kubebuilder:validation:Optional
	IAMInstanceProfile *string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6AddressCount *float64 `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Addresses []*string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceInitiatedShutdownBehavior *string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior,omitempty"`

	// Indicates Spot instance behavior when it is interrupted. Valid values are terminate, stop, or hibernate. Default value is terminate.
	// +kubebuilder:validation:Optional
	InstanceInterruptionBehavior *string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// A launch group is a group of spot instances that launch together and terminate together.
	// If left empty instances are launched and terminated individually.
	// +kubebuilder:validation:Optional
	LaunchGroup *string `json:"launchGroup,omitempty" tf:"launch_group,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplate []SpotInstanceRequestLaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceOptions []SpotInstanceRequestMaintenanceOptionsParameters `json:"maintenanceOptions,omitempty" tf:"maintenance_options,omitempty"`

	// +kubebuilder:validation:Optional
	MetadataOptions []SpotInstanceRequestMetadataOptionsParameters `json:"metadataOptions,omitempty" tf:"metadata_options,omitempty"`

	// +kubebuilder:validation:Optional
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterface []SpotInstanceRequestNetworkInterfaceParameters `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Optional
	PlacementGroup *string `json:"placementGroup,omitempty" tf:"placement_group,omitempty"`

	// +kubebuilder:validation:Optional
	PlacementPartitionNumber *float64 `json:"placementPartitionNumber,omitempty" tf:"placement_partition_number,omitempty"`

	// The private IP address assigned to the instance
	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RootBlockDevice []SpotInstanceRequestRootBlockDeviceParameters `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	SecondaryPrivateIps []*string `json:"secondaryPrivateIps,omitempty" tf:"secondary_private_ips,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// +kubebuilder:validation:Optional
	SourceDestCheck *bool `json:"sourceDestCheck,omitempty" tf:"source_dest_check,omitempty"`

	// The maximum price to request on the spot market.
	// +kubebuilder:validation:Optional
	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// If set to one-time, after
	// the instance is terminated, the spot request will be closed.
	// +kubebuilder:validation:Optional
	SpotType *string `json:"spotType,omitempty" tf:"spot_type,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A map of tags to assign to the Spot Instance Request. These tags are not automatically applied to the launched Instance. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`

	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// +kubebuilder:validation:Optional
	UserDataBase64 *string `json:"userDataBase64,omitempty" tf:"user_data_base64,omitempty"`

	// +kubebuilder:validation:Optional
	UserDataReplaceOnChange *bool `json:"userDataReplaceOnChange,omitempty" tf:"user_data_replace_on_change,omitempty"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// The start date and time of the request, in UTC RFC3339 format. The default is to start fulfilling the request immediately.
	// +kubebuilder:validation:Optional
	ValidFrom *string `json:"validFrom,omitempty" tf:"valid_from,omitempty"`

	// The end date and time of the request, in UTC RFC3339 format. At this point, no new Spot instance requests are placed or enabled to fulfill the request. The default end date is 7 days from the current date.
	// +kubebuilder:validation:Optional
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeTags map[string]*string `json:"volumeTags,omitempty" tf:"volume_tags,omitempty"`

	// If set, Terraform will
	// wait for the Spot Request to be fulfilled, and will throw an error if the
	// timeout of 10m is reached.
	// +kubebuilder:validation:Optional
	WaitForFulfillment *bool `json:"waitForFulfillment,omitempty" tf:"wait_for_fulfillment,omitempty"`
}

type SpotInstanceRequestRootBlockDeviceObservation struct {
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type SpotInstanceRequestRootBlockDeviceParameters struct {

	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// A map of tags to assign to the Spot Instance Request. These tags are not automatically applied to the launched Instance. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

// SpotInstanceRequestSpec defines the desired state of SpotInstanceRequest
type SpotInstanceRequestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpotInstanceRequestParameters `json:"forProvider"`
}

// SpotInstanceRequestStatus defines the observed state of SpotInstanceRequest.
type SpotInstanceRequestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpotInstanceRequestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpotInstanceRequest is the Schema for the SpotInstanceRequests API. Provides a Spot Instance Request resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SpotInstanceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpotInstanceRequestSpec   `json:"spec"`
	Status            SpotInstanceRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotInstanceRequestList contains a list of SpotInstanceRequests
type SpotInstanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotInstanceRequest `json:"items"`
}

// Repository type metadata.
var (
	SpotInstanceRequest_Kind             = "SpotInstanceRequest"
	SpotInstanceRequest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpotInstanceRequest_Kind}.String()
	SpotInstanceRequest_KindAPIVersion   = SpotInstanceRequest_Kind + "." + CRDGroupVersion.String()
	SpotInstanceRequest_GroupVersionKind = CRDGroupVersion.WithKind(SpotInstanceRequest_Kind)
)

func init() {
	SchemeBuilder.Register(&SpotInstanceRequest{}, &SpotInstanceRequestList{})
}
