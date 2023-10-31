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

type ConfigurationInitParameters struct {

	// Description of the configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions []*string `json:"kafkaVersions,omitempty" tf:"kafka_versions,omitempty"`

	// Name of the configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Contents of the server.properties file. Supported properties are documented in the MSK Developer Guide.
	ServerProperties *string `json:"serverProperties,omitempty" tf:"server_properties,omitempty"`
}

type ConfigurationObservation struct {

	// Amazon Resource Name (ARN) of the configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions []*string `json:"kafkaVersions,omitempty" tf:"kafka_versions,omitempty"`

	// Latest revision of the configuration.
	LatestRevision *float64 `json:"latestRevision,omitempty" tf:"latest_revision,omitempty"`

	// Name of the configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Contents of the server.properties file. Supported properties are documented in the MSK Developer Guide.
	ServerProperties *string `json:"serverProperties,omitempty" tf:"server_properties,omitempty"`
}

type ConfigurationParameters struct {

	// Description of the configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of Apache Kafka versions which can use this configuration.
	// +kubebuilder:validation:Optional
	KafkaVersions []*string `json:"kafkaVersions,omitempty" tf:"kafka_versions,omitempty"`

	// Name of the configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Contents of the server.properties file. Supported properties are documented in the MSK Developer Guide.
	// +kubebuilder:validation:Optional
	ServerProperties *string `json:"serverProperties,omitempty" tf:"server_properties,omitempty"`
}

// ConfigurationSpec defines the desired state of Configuration
type ConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationParameters `json:"forProvider"`
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
	InitProvider ConfigurationInitParameters `json:"initProvider,omitempty"`
}

// ConfigurationStatus defines the observed state of Configuration.
type ConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Configuration is the Schema for the Configurations API. provider resource for managing an amazon managed streaming for kafka configuration
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Configuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serverProperties) || (has(self.initProvider) && has(self.initProvider.serverProperties))",message="spec.forProvider.serverProperties is a required parameter"
	Spec   ConfigurationSpec   `json:"spec"`
	Status ConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationList contains a list of Configurations
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Configuration `json:"items"`
}

// Repository type metadata.
var (
	Configuration_Kind             = "Configuration"
	Configuration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Configuration_Kind}.String()
	Configuration_KindAPIVersion   = Configuration_Kind + "." + CRDGroupVersion.String()
	Configuration_GroupVersionKind = CRDGroupVersion.WithKind(Configuration_Kind)
)

func init() {
	SchemeBuilder.Register(&Configuration{}, &ConfigurationList{})
}
