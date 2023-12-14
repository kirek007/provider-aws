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

type CustomerGatewayInitParameters struct {

	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BGPAsn *string `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// A name for the customer gateway device.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// The IPv4 address for the customer gateway device's outside interface.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of customer gateway. The only type AWS
	// supports at this time is "ipsec.1".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CustomerGatewayObservation struct {

	// The ARN of the customer gateway.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BGPAsn *string `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// A name for the customer gateway device.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// The amazon-assigned ID of the gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IPv4 address for the customer gateway device's outside interface.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The type of customer gateway. The only type AWS
	// supports at this time is "ipsec.1".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CustomerGatewayParameters struct {

	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	// +kubebuilder:validation:Optional
	BGPAsn *string `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// A name for the customer gateway device.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// The IPv4 address for the customer gateway device's outside interface.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of customer gateway. The only type AWS
	// supports at this time is "ipsec.1".
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// CustomerGatewaySpec defines the desired state of CustomerGateway
type CustomerGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomerGatewayParameters `json:"forProvider"`
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
	InitProvider CustomerGatewayInitParameters `json:"initProvider,omitempty"`
}

// CustomerGatewayStatus defines the observed state of CustomerGateway.
type CustomerGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomerGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomerGateway is the Schema for the CustomerGateways API. Provides a customer gateway inside a VPC. These objects can be connected to VPN gateways via VPN connections, and allow you to establish tunnels between your network and the VPC.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CustomerGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bgpAsn) || (has(self.initProvider) && has(self.initProvider.bgpAsn))",message="spec.forProvider.bgpAsn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   CustomerGatewaySpec   `json:"spec"`
	Status CustomerGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomerGatewayList contains a list of CustomerGateways
type CustomerGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomerGateway `json:"items"`
}

// Repository type metadata.
var (
	CustomerGateway_Kind             = "CustomerGateway"
	CustomerGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomerGateway_Kind}.String()
	CustomerGateway_KindAPIVersion   = CustomerGateway_Kind + "." + CRDGroupVersion.String()
	CustomerGateway_GroupVersionKind = CRDGroupVersion.WithKind(CustomerGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomerGateway{}, &CustomerGatewayList{})
}
