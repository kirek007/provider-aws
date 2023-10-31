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

type SpaceInitParameters struct {

	// The name of the space.
	SpaceName *string `json:"spaceName,omitempty" tf:"space_name,omitempty"`

	// A collection of space settings. See Space Settings below.
	SpaceSettings []SpaceSettingsInitParameters `json:"spaceSettings,omitempty" tf:"space_settings,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SpaceObservation struct {

	// The space's Amazon Resource Name (ARN).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the associated Domain.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The ID of the space's profile in the Amazon Elastic File System volume.
	HomeEFSFileSystemUID *string `json:"homeEfsFileSystemUid,omitempty" tf:"home_efs_file_system_uid,omitempty"`

	// The space's Amazon Resource Name (ARN).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the space.
	SpaceName *string `json:"spaceName,omitempty" tf:"space_name,omitempty"`

	// A collection of space settings. See Space Settings below.
	SpaceSettings []SpaceSettingsObservation `json:"spaceSettings,omitempty" tf:"space_settings,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SpaceParameters struct {

	// The ID of the associated Domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta1.Domain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// Reference to a Domain in sagemaker to populate domainId.
	// +kubebuilder:validation:Optional
	DomainIDRef *v1.Reference `json:"domainIdRef,omitempty" tf:"-"`

	// Selector for a Domain in sagemaker to populate domainId.
	// +kubebuilder:validation:Optional
	DomainIDSelector *v1.Selector `json:"domainIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the space.
	// +kubebuilder:validation:Optional
	SpaceName *string `json:"spaceName,omitempty" tf:"space_name,omitempty"`

	// A collection of space settings. See Space Settings below.
	// +kubebuilder:validation:Optional
	SpaceSettings []SpaceSettingsParameters `json:"spaceSettings,omitempty" tf:"space_settings,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SpaceSettingsInitParameters struct {

	// The Jupyter server's app settings. See Jupyter Server App Settings below.
	JupyterServerAppSettings []SpaceSettingsJupyterServerAppSettingsInitParameters `json:"jupyterServerAppSettings,omitempty" tf:"jupyter_server_app_settings,omitempty"`

	// The kernel gateway app settings. See Kernel Gateway App Settings below.
	KernelGatewayAppSettings []SpaceSettingsKernelGatewayAppSettingsInitParameters `json:"kernelGatewayAppSettings,omitempty" tf:"kernel_gateway_app_settings,omitempty"`
}

type SpaceSettingsJupyterServerAppSettingsCodeRepositoryInitParameters struct {

	// The URL of the Git repository.
	RepositoryURL *string `json:"repositoryUrl,omitempty" tf:"repository_url,omitempty"`
}

type SpaceSettingsJupyterServerAppSettingsCodeRepositoryObservation struct {

	// The URL of the Git repository.
	RepositoryURL *string `json:"repositoryUrl,omitempty" tf:"repository_url,omitempty"`
}

type SpaceSettingsJupyterServerAppSettingsCodeRepositoryParameters struct {

	// The URL of the Git repository.
	// +kubebuilder:validation:Optional
	RepositoryURL *string `json:"repositoryUrl" tf:"repository_url,omitempty"`
}

type SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecInitParameters struct {

	// The instance type.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecObservation struct {

	// The instance type.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecParameters struct {

	// The instance type.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	// +kubebuilder:validation:Optional
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	// +kubebuilder:validation:Optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	// +kubebuilder:validation:Optional
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type SpaceSettingsJupyterServerAppSettingsInitParameters struct {

	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterServer application. see Code Repository below.
	CodeRepository []SpaceSettingsJupyterServerAppSettingsCodeRepositoryInitParameters `json:"codeRepository,omitempty" tf:"code_repository,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	DefaultResourceSpec []SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecInitParameters `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	LifecycleConfigArns []*string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns,omitempty"`
}

type SpaceSettingsJupyterServerAppSettingsObservation struct {

	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterServer application. see Code Repository below.
	CodeRepository []SpaceSettingsJupyterServerAppSettingsCodeRepositoryObservation `json:"codeRepository,omitempty" tf:"code_repository,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	DefaultResourceSpec []SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecObservation `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	LifecycleConfigArns []*string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns,omitempty"`
}

type SpaceSettingsJupyterServerAppSettingsParameters struct {

	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterServer application. see Code Repository below.
	// +kubebuilder:validation:Optional
	CodeRepository []SpaceSettingsJupyterServerAppSettingsCodeRepositoryParameters `json:"codeRepository,omitempty" tf:"code_repository,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	// +kubebuilder:validation:Optional
	DefaultResourceSpec []SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec" tf:"default_resource_spec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	// +kubebuilder:validation:Optional
	LifecycleConfigArns []*string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns,omitempty"`
}

type SpaceSettingsKernelGatewayAppSettingsCustomImageInitParameters struct {

	// The name of the App Image Config.
	AppImageConfigName *string `json:"appImageConfigName,omitempty" tf:"app_image_config_name,omitempty"`

	// The name of the Custom Image.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// The version number of the Custom Image.
	ImageVersionNumber *float64 `json:"imageVersionNumber,omitempty" tf:"image_version_number,omitempty"`
}

type SpaceSettingsKernelGatewayAppSettingsCustomImageObservation struct {

	// The name of the App Image Config.
	AppImageConfigName *string `json:"appImageConfigName,omitempty" tf:"app_image_config_name,omitempty"`

	// The name of the Custom Image.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// The version number of the Custom Image.
	ImageVersionNumber *float64 `json:"imageVersionNumber,omitempty" tf:"image_version_number,omitempty"`
}

type SpaceSettingsKernelGatewayAppSettingsCustomImageParameters struct {

	// The name of the App Image Config.
	// +kubebuilder:validation:Optional
	AppImageConfigName *string `json:"appImageConfigName" tf:"app_image_config_name,omitempty"`

	// The name of the Custom Image.
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName" tf:"image_name,omitempty"`

	// The version number of the Custom Image.
	// +kubebuilder:validation:Optional
	ImageVersionNumber *float64 `json:"imageVersionNumber,omitempty" tf:"image_version_number,omitempty"`
}

type SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecInitParameters struct {

	// The instance type.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecObservation struct {

	// The instance type.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecParameters struct {

	// The instance type.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	// +kubebuilder:validation:Optional
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	// +kubebuilder:validation:Optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	// +kubebuilder:validation:Optional
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type SpaceSettingsKernelGatewayAppSettingsInitParameters struct {

	// A list of custom SageMaker images that are configured to run as a KernelGateway app. see Custom Image below.
	CustomImage []SpaceSettingsKernelGatewayAppSettingsCustomImageInitParameters `json:"customImage,omitempty" tf:"custom_image,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	DefaultResourceSpec []SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecInitParameters `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	LifecycleConfigArns []*string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns,omitempty"`
}

type SpaceSettingsKernelGatewayAppSettingsObservation struct {

	// A list of custom SageMaker images that are configured to run as a KernelGateway app. see Custom Image below.
	CustomImage []SpaceSettingsKernelGatewayAppSettingsCustomImageObservation `json:"customImage,omitempty" tf:"custom_image,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	DefaultResourceSpec []SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecObservation `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	LifecycleConfigArns []*string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns,omitempty"`
}

type SpaceSettingsKernelGatewayAppSettingsParameters struct {

	// A list of custom SageMaker images that are configured to run as a KernelGateway app. see Custom Image below.
	// +kubebuilder:validation:Optional
	CustomImage []SpaceSettingsKernelGatewayAppSettingsCustomImageParameters `json:"customImage,omitempty" tf:"custom_image,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	// +kubebuilder:validation:Optional
	DefaultResourceSpec []SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec" tf:"default_resource_spec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	// +kubebuilder:validation:Optional
	LifecycleConfigArns []*string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns,omitempty"`
}

type SpaceSettingsObservation struct {

	// The Jupyter server's app settings. See Jupyter Server App Settings below.
	JupyterServerAppSettings []SpaceSettingsJupyterServerAppSettingsObservation `json:"jupyterServerAppSettings,omitempty" tf:"jupyter_server_app_settings,omitempty"`

	// The kernel gateway app settings. See Kernel Gateway App Settings below.
	KernelGatewayAppSettings []SpaceSettingsKernelGatewayAppSettingsObservation `json:"kernelGatewayAppSettings,omitempty" tf:"kernel_gateway_app_settings,omitempty"`
}

type SpaceSettingsParameters struct {

	// The Jupyter server's app settings. See Jupyter Server App Settings below.
	// +kubebuilder:validation:Optional
	JupyterServerAppSettings []SpaceSettingsJupyterServerAppSettingsParameters `json:"jupyterServerAppSettings,omitempty" tf:"jupyter_server_app_settings,omitempty"`

	// The kernel gateway app settings. See Kernel Gateway App Settings below.
	// +kubebuilder:validation:Optional
	KernelGatewayAppSettings []SpaceSettingsKernelGatewayAppSettingsParameters `json:"kernelGatewayAppSettings,omitempty" tf:"kernel_gateway_app_settings,omitempty"`
}

// SpaceSpec defines the desired state of Space
type SpaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpaceParameters `json:"forProvider"`
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
	InitProvider SpaceInitParameters `json:"initProvider,omitempty"`
}

// SpaceStatus defines the observed state of Space.
type SpaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Space is the Schema for the Spaces API. Provides a SageMaker Space resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Space struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.spaceName) || (has(self.initProvider) && has(self.initProvider.spaceName))",message="spec.forProvider.spaceName is a required parameter"
	Spec   SpaceSpec   `json:"spec"`
	Status SpaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpaceList contains a list of Spaces
type SpaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Space `json:"items"`
}

// Repository type metadata.
var (
	Space_Kind             = "Space"
	Space_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Space_Kind}.String()
	Space_KindAPIVersion   = Space_Kind + "." + CRDGroupVersion.String()
	Space_GroupVersionKind = CRDGroupVersion.WithKind(Space_Kind)
)

func init() {
	SchemeBuilder.Register(&Space{}, &SpaceList{})
}
