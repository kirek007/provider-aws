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

type VPCEndpointConnectionNotificationInitParameters struct {

	// One or more endpoint events for which to receive notifications.
	// +listType=set
	ConnectionEvents []*string `json:"connectionEvents,omitempty" tf:"connection_events,omitempty"`

	// The ARN of the SNS topic for the notifications.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	ConnectionNotificationArn *string `json:"connectionNotificationArn,omitempty" tf:"connection_notification_arn,omitempty"`

	// Reference to a Topic in sns to populate connectionNotificationArn.
	// +kubebuilder:validation:Optional
	ConnectionNotificationArnRef *v1.Reference `json:"connectionNotificationArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate connectionNotificationArn.
	// +kubebuilder:validation:Optional
	ConnectionNotificationArnSelector *v1.Selector `json:"connectionNotificationArnSelector,omitempty" tf:"-"`

	// The ID of the VPC Endpoint to receive notifications for.
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The ID of the VPC Endpoint Service to receive notifications for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCEndpointService
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VPCEndpointServiceID *string `json:"vpcEndpointServiceId,omitempty" tf:"vpc_endpoint_service_id,omitempty"`

	// Reference to a VPCEndpointService in ec2 to populate vpcEndpointServiceId.
	// +kubebuilder:validation:Optional
	VPCEndpointServiceIDRef *v1.Reference `json:"vpcEndpointServiceIdRef,omitempty" tf:"-"`

	// Selector for a VPCEndpointService in ec2 to populate vpcEndpointServiceId.
	// +kubebuilder:validation:Optional
	VPCEndpointServiceIDSelector *v1.Selector `json:"vpcEndpointServiceIdSelector,omitempty" tf:"-"`
}

type VPCEndpointConnectionNotificationObservation struct {

	// One or more endpoint events for which to receive notifications.
	// +listType=set
	ConnectionEvents []*string `json:"connectionEvents,omitempty" tf:"connection_events,omitempty"`

	// The ARN of the SNS topic for the notifications.
	ConnectionNotificationArn *string `json:"connectionNotificationArn,omitempty" tf:"connection_notification_arn,omitempty"`

	// The ID of the VPC connection notification.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of notification.
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`

	// The state of the notification.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The ID of the VPC Endpoint to receive notifications for.
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The ID of the VPC Endpoint Service to receive notifications for.
	VPCEndpointServiceID *string `json:"vpcEndpointServiceId,omitempty" tf:"vpc_endpoint_service_id,omitempty"`
}

type VPCEndpointConnectionNotificationParameters struct {

	// One or more endpoint events for which to receive notifications.
	// +kubebuilder:validation:Optional
	// +listType=set
	ConnectionEvents []*string `json:"connectionEvents,omitempty" tf:"connection_events,omitempty"`

	// The ARN of the SNS topic for the notifications.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ConnectionNotificationArn *string `json:"connectionNotificationArn,omitempty" tf:"connection_notification_arn,omitempty"`

	// Reference to a Topic in sns to populate connectionNotificationArn.
	// +kubebuilder:validation:Optional
	ConnectionNotificationArnRef *v1.Reference `json:"connectionNotificationArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate connectionNotificationArn.
	// +kubebuilder:validation:Optional
	ConnectionNotificationArnSelector *v1.Selector `json:"connectionNotificationArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the VPC Endpoint to receive notifications for.
	// +kubebuilder:validation:Optional
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The ID of the VPC Endpoint Service to receive notifications for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCEndpointService
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCEndpointServiceID *string `json:"vpcEndpointServiceId,omitempty" tf:"vpc_endpoint_service_id,omitempty"`

	// Reference to a VPCEndpointService in ec2 to populate vpcEndpointServiceId.
	// +kubebuilder:validation:Optional
	VPCEndpointServiceIDRef *v1.Reference `json:"vpcEndpointServiceIdRef,omitempty" tf:"-"`

	// Selector for a VPCEndpointService in ec2 to populate vpcEndpointServiceId.
	// +kubebuilder:validation:Optional
	VPCEndpointServiceIDSelector *v1.Selector `json:"vpcEndpointServiceIdSelector,omitempty" tf:"-"`
}

// VPCEndpointConnectionNotificationSpec defines the desired state of VPCEndpointConnectionNotification
type VPCEndpointConnectionNotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointConnectionNotificationParameters `json:"forProvider"`
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
	InitProvider VPCEndpointConnectionNotificationInitParameters `json:"initProvider,omitempty"`
}

// VPCEndpointConnectionNotificationStatus defines the observed state of VPCEndpointConnectionNotification.
type VPCEndpointConnectionNotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointConnectionNotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPCEndpointConnectionNotification is the Schema for the VPCEndpointConnectionNotifications API. Provides a VPC Endpoint connection notification resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCEndpointConnectionNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionEvents) || (has(self.initProvider) && has(self.initProvider.connectionEvents))",message="spec.forProvider.connectionEvents is a required parameter"
	Spec   VPCEndpointConnectionNotificationSpec   `json:"spec"`
	Status VPCEndpointConnectionNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointConnectionNotificationList contains a list of VPCEndpointConnectionNotifications
type VPCEndpointConnectionNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpointConnectionNotification `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpointConnectionNotification_Kind             = "VPCEndpointConnectionNotification"
	VPCEndpointConnectionNotification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpointConnectionNotification_Kind}.String()
	VPCEndpointConnectionNotification_KindAPIVersion   = VPCEndpointConnectionNotification_Kind + "." + CRDGroupVersion.String()
	VPCEndpointConnectionNotification_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpointConnectionNotification_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpointConnectionNotification{}, &VPCEndpointConnectionNotificationList{})
}
