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

type AmisObservation struct {

	// Account identifier of the AMI.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Description of the AMI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Identifier of the AMI.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Name of the AMI.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region of the AMI.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type AmisParameters struct {
}

type ImageObservation struct {

	// Amazon Resource Name (ARN) of the image.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// - Amazon Resource Name (ARN) of the container recipe.
	ContainerRecipeArn *string `json:"containerRecipeArn,omitempty" tf:"container_recipe_arn,omitempty"`

	// Date the image was created.
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn *string `json:"distributionConfigurationArn,omitempty" tf:"distribution_configuration_arn,omitempty"`

	// Whether additional information about the image being created is collected. Defaults to true.
	EnhancedImageMetadataEnabled *bool `json:"enhancedImageMetadataEnabled,omitempty" tf:"enhanced_image_metadata_enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn *string `json:"imageRecipeArn,omitempty" tf:"image_recipe_arn,omitempty"`

	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration []ImageTestsConfigurationObservation `json:"imageTestsConfiguration,omitempty" tf:"image_tests_configuration,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn *string `json:"infrastructureConfigurationArn,omitempty" tf:"infrastructure_configuration_arn,omitempty"`

	// Name of the AMI.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operating System version of the image.
	OsVersion *string `json:"osVersion,omitempty" tf:"os_version,omitempty"`

	// List of objects with resources created by the image.
	OutputResources []OutputResourcesObservation `json:"outputResources,omitempty" tf:"output_resources,omitempty"`

	// Platform of the image.
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Version of the image.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ImageParameters struct {

	// - Amazon Resource Name (ARN) of the container recipe.
	// +kubebuilder:validation:Optional
	ContainerRecipeArn *string `json:"containerRecipeArn,omitempty" tf:"container_recipe_arn,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/imagebuilder/v1beta1.DistributionConfiguration
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	DistributionConfigurationArn *string `json:"distributionConfigurationArn,omitempty" tf:"distribution_configuration_arn,omitempty"`

	// Reference to a DistributionConfiguration in imagebuilder to populate distributionConfigurationArn.
	// +kubebuilder:validation:Optional
	DistributionConfigurationArnRef *v1.Reference `json:"distributionConfigurationArnRef,omitempty" tf:"-"`

	// Selector for a DistributionConfiguration in imagebuilder to populate distributionConfigurationArn.
	// +kubebuilder:validation:Optional
	DistributionConfigurationArnSelector *v1.Selector `json:"distributionConfigurationArnSelector,omitempty" tf:"-"`

	// Whether additional information about the image being created is collected. Defaults to true.
	// +kubebuilder:validation:Optional
	EnhancedImageMetadataEnabled *bool `json:"enhancedImageMetadataEnabled,omitempty" tf:"enhanced_image_metadata_enabled,omitempty"`

	// Amazon Resource Name (ARN) of the image recipe.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/imagebuilder/v1beta1.ImageRecipe
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ImageRecipeArn *string `json:"imageRecipeArn,omitempty" tf:"image_recipe_arn,omitempty"`

	// Reference to a ImageRecipe in imagebuilder to populate imageRecipeArn.
	// +kubebuilder:validation:Optional
	ImageRecipeArnRef *v1.Reference `json:"imageRecipeArnRef,omitempty" tf:"-"`

	// Selector for a ImageRecipe in imagebuilder to populate imageRecipeArn.
	// +kubebuilder:validation:Optional
	ImageRecipeArnSelector *v1.Selector `json:"imageRecipeArnSelector,omitempty" tf:"-"`

	// Configuration block with image tests configuration. Detailed below.
	// +kubebuilder:validation:Optional
	ImageTestsConfiguration []ImageTestsConfigurationParameters `json:"imageTestsConfiguration,omitempty" tf:"image_tests_configuration,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/imagebuilder/v1beta1.InfrastructureConfiguration
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	InfrastructureConfigurationArn *string `json:"infrastructureConfigurationArn,omitempty" tf:"infrastructure_configuration_arn,omitempty"`

	// Reference to a InfrastructureConfiguration in imagebuilder to populate infrastructureConfigurationArn.
	// +kubebuilder:validation:Optional
	InfrastructureConfigurationArnRef *v1.Reference `json:"infrastructureConfigurationArnRef,omitempty" tf:"-"`

	// Selector for a InfrastructureConfiguration in imagebuilder to populate infrastructureConfigurationArn.
	// +kubebuilder:validation:Optional
	InfrastructureConfigurationArnSelector *v1.Selector `json:"infrastructureConfigurationArnSelector,omitempty" tf:"-"`

	// Region of the AMI.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ImageTestsConfigurationObservation struct {

	// Whether image tests are enabled. Defaults to true.
	ImageTestsEnabled *bool `json:"imageTestsEnabled,omitempty" tf:"image_tests_enabled,omitempty"`

	// Number of minutes before image tests time out. Valid values are between 60 and 1440. Defaults to 720.
	TimeoutMinutes *float64 `json:"timeoutMinutes,omitempty" tf:"timeout_minutes,omitempty"`
}

type ImageTestsConfigurationParameters struct {

	// Whether image tests are enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	ImageTestsEnabled *bool `json:"imageTestsEnabled,omitempty" tf:"image_tests_enabled,omitempty"`

	// Number of minutes before image tests time out. Valid values are between 60 and 1440. Defaults to 720.
	// +kubebuilder:validation:Optional
	TimeoutMinutes *float64 `json:"timeoutMinutes,omitempty" tf:"timeout_minutes,omitempty"`
}

type OutputResourcesObservation struct {

	// Set of objects with each Amazon Machine Image (AMI) created.
	Amis []AmisObservation `json:"amis,omitempty" tf:"amis,omitempty"`
}

type OutputResourcesParameters struct {
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageParameters `json:"forProvider"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Image is the Schema for the Images API. Manages an Image Builder Image
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpec   `json:"spec"`
	Status            ImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

// Repository type metadata.
var (
	Image_Kind             = "Image"
	Image_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Image_Kind}.String()
	Image_KindAPIVersion   = Image_Kind + "." + CRDGroupVersion.String()
	Image_GroupVersionKind = CRDGroupVersion.WithKind(Image_Kind)
)

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}
