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

type APIDestinationInitParameters struct {

	// The description of the new API Destination. Maximum of 512 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Select the HTTP method used for the invocation endpoint, such as GET, POST, PUT, etc.
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// URL endpoint to invoke as a target. This could be a valid endpoint generated by a partner service. You can include "*" as path parameters wildcards to be set from the Target HttpParameters.
	InvocationEndpoint *string `json:"invocationEndpoint,omitempty" tf:"invocation_endpoint,omitempty"`

	// Enter the maximum number of invocations per second to allow for this destination. Enter a value greater than 0 (default 300).
	InvocationRateLimitPerSecond *int64 `json:"invocationRateLimitPerSecond,omitempty" tf:"invocation_rate_limit_per_second,omitempty"`
}

type APIDestinationObservation struct {

	// The Amazon Resource Name (ARN) of the event API Destination.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of the EventBridge Connection to use for the API Destination.
	ConnectionArn *string `json:"connectionArn,omitempty" tf:"connection_arn,omitempty"`

	// The description of the new API Destination. Maximum of 512 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Select the HTTP method used for the invocation endpoint, such as GET, POST, PUT, etc.
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// URL endpoint to invoke as a target. This could be a valid endpoint generated by a partner service. You can include "*" as path parameters wildcards to be set from the Target HttpParameters.
	InvocationEndpoint *string `json:"invocationEndpoint,omitempty" tf:"invocation_endpoint,omitempty"`

	// Enter the maximum number of invocations per second to allow for this destination. Enter a value greater than 0 (default 300).
	InvocationRateLimitPerSecond *int64 `json:"invocationRateLimitPerSecond,omitempty" tf:"invocation_rate_limit_per_second,omitempty"`
}

type APIDestinationParameters struct {

	// ARN of the EventBridge Connection to use for the API Destination.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatchevents/v1beta1.Connection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ConnectionArn *string `json:"connectionArn,omitempty" tf:"connection_arn,omitempty"`

	// Reference to a Connection in cloudwatchevents to populate connectionArn.
	// +kubebuilder:validation:Optional
	ConnectionArnRef *v1.Reference `json:"connectionArnRef,omitempty" tf:"-"`

	// Selector for a Connection in cloudwatchevents to populate connectionArn.
	// +kubebuilder:validation:Optional
	ConnectionArnSelector *v1.Selector `json:"connectionArnSelector,omitempty" tf:"-"`

	// The description of the new API Destination. Maximum of 512 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Select the HTTP method used for the invocation endpoint, such as GET, POST, PUT, etc.
	// +kubebuilder:validation:Optional
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// URL endpoint to invoke as a target. This could be a valid endpoint generated by a partner service. You can include "*" as path parameters wildcards to be set from the Target HttpParameters.
	// +kubebuilder:validation:Optional
	InvocationEndpoint *string `json:"invocationEndpoint,omitempty" tf:"invocation_endpoint,omitempty"`

	// Enter the maximum number of invocations per second to allow for this destination. Enter a value greater than 0 (default 300).
	// +kubebuilder:validation:Optional
	InvocationRateLimitPerSecond *int64 `json:"invocationRateLimitPerSecond,omitempty" tf:"invocation_rate_limit_per_second,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// APIDestinationSpec defines the desired state of APIDestination
type APIDestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIDestinationParameters `json:"forProvider"`
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
	InitProvider APIDestinationInitParameters `json:"initProvider,omitempty"`
}

// APIDestinationStatus defines the observed state of APIDestination.
type APIDestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIDestination is the Schema for the APIDestinations API. Provides an EventBridge event API Destination resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type APIDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.httpMethod) || (has(self.initProvider) && has(self.initProvider.httpMethod))",message="spec.forProvider.httpMethod is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.invocationEndpoint) || (has(self.initProvider) && has(self.initProvider.invocationEndpoint))",message="spec.forProvider.invocationEndpoint is a required parameter"
	Spec   APIDestinationSpec   `json:"spec"`
	Status APIDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIDestinationList contains a list of APIDestinations
type APIDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIDestination `json:"items"`
}

// Repository type metadata.
var (
	APIDestination_Kind             = "APIDestination"
	APIDestination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIDestination_Kind}.String()
	APIDestination_KindAPIVersion   = APIDestination_Kind + "." + CRDGroupVersion.String()
	APIDestination_GroupVersionKind = CRDGroupVersion.WithKind(APIDestination_Kind)
)

func init() {
	SchemeBuilder.Register(&APIDestination{}, &APIDestinationList{})
}
