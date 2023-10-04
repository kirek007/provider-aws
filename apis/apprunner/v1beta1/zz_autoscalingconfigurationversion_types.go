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

type AutoScalingConfigurationVersionInitParameters struct {

	// Name of the auto scaling configuration.
	AutoScalingConfigurationName *string `json:"autoScalingConfigurationName,omitempty" tf:"auto_scaling_configuration_name,omitempty"`

	// Maximal number of concurrent requests that you want an instance to process. When the number of concurrent requests goes over this limit, App Runner scales up your service.
	MaxConcurrency *int64 `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// Maximal number of instances that App Runner provisions for your service.
	MaxSize *int64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Minimal number of instances that App Runner provisions for your service.
	MinSize *int64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AutoScalingConfigurationVersionObservation struct {

	// ARN of this auto scaling configuration version.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Name of the auto scaling configuration.
	AutoScalingConfigurationName *string `json:"autoScalingConfigurationName,omitempty" tf:"auto_scaling_configuration_name,omitempty"`

	// The revision of this auto scaling configuration.
	AutoScalingConfigurationRevision *int64 `json:"autoScalingConfigurationRevision,omitempty" tf:"auto_scaling_configuration_revision,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the auto scaling configuration has the highest auto_scaling_configuration_revision among all configurations that share the same auto_scaling_configuration_name.
	Latest *bool `json:"latest,omitempty" tf:"latest,omitempty"`

	// Maximal number of concurrent requests that you want an instance to process. When the number of concurrent requests goes over this limit, App Runner scales up your service.
	MaxConcurrency *int64 `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// Maximal number of instances that App Runner provisions for your service.
	MaxSize *int64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Minimal number of instances that App Runner provisions for your service.
	MinSize *int64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// Current state of the auto scaling configuration. An INACTIVE configuration revision has been deleted and can't be used. It is permanently removed some time after deletion.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AutoScalingConfigurationVersionParameters struct {

	// Name of the auto scaling configuration.
	// +kubebuilder:validation:Optional
	AutoScalingConfigurationName *string `json:"autoScalingConfigurationName,omitempty" tf:"auto_scaling_configuration_name,omitempty"`

	// Maximal number of concurrent requests that you want an instance to process. When the number of concurrent requests goes over this limit, App Runner scales up your service.
	// +kubebuilder:validation:Optional
	MaxConcurrency *int64 `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// Maximal number of instances that App Runner provisions for your service.
	// +kubebuilder:validation:Optional
	MaxSize *int64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Minimal number of instances that App Runner provisions for your service.
	// +kubebuilder:validation:Optional
	MinSize *int64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AutoScalingConfigurationVersionSpec defines the desired state of AutoScalingConfigurationVersion
type AutoScalingConfigurationVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutoScalingConfigurationVersionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AutoScalingConfigurationVersionInitParameters `json:"initProvider,omitempty"`
}

// AutoScalingConfigurationVersionStatus defines the observed state of AutoScalingConfigurationVersion.
type AutoScalingConfigurationVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutoScalingConfigurationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoScalingConfigurationVersion is the Schema for the AutoScalingConfigurationVersions API. Manages an App Runner AutoScaling Configuration Version.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AutoScalingConfigurationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autoScalingConfigurationName) || has(self.initProvider.autoScalingConfigurationName)",message="autoScalingConfigurationName is a required parameter"
	Spec   AutoScalingConfigurationVersionSpec   `json:"spec"`
	Status AutoScalingConfigurationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoScalingConfigurationVersionList contains a list of AutoScalingConfigurationVersions
type AutoScalingConfigurationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoScalingConfigurationVersion `json:"items"`
}

// Repository type metadata.
var (
	AutoScalingConfigurationVersion_Kind             = "AutoScalingConfigurationVersion"
	AutoScalingConfigurationVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutoScalingConfigurationVersion_Kind}.String()
	AutoScalingConfigurationVersion_KindAPIVersion   = AutoScalingConfigurationVersion_Kind + "." + CRDGroupVersion.String()
	AutoScalingConfigurationVersion_GroupVersionKind = CRDGroupVersion.WithKind(AutoScalingConfigurationVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&AutoScalingConfigurationVersion{}, &AutoScalingConfigurationVersionList{})
}
