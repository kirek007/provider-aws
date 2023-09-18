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

type TransitVirtualInterfaceInitParameters struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BGPAsn *int64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
	// The MTU of a virtual transit interface can be either 1500 or 8500 (jumbo frames). Default is 1500.
	Mtu *int64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name for the virtual interface.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates whether to enable or disable SiteLink.
	SitelinkEnabled *bool `json:"sitelinkEnabled,omitempty" tf:"sitelink_enabled,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VLAN ID.
	Vlan *int64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type TransitVirtualInterfaceObservation struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn,omitempty"`

	// The ARN of the virtual interface.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BGPAsn *int64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// The ID of the virtual interface.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether jumbo frames (8500 MTU) are supported.
	JumboFrameCapable *bool `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable,omitempty"`

	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
	// The MTU of a virtual transit interface can be either 1500 or 8500 (jumbo frames). Default is 1500.
	Mtu *int64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name for the virtual interface.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates whether to enable or disable SiteLink.
	SitelinkEnabled *bool `json:"sitelinkEnabled,omitempty" tf:"sitelink_enabled,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The VLAN ID.
	Vlan *int64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type TransitVirtualInterfaceParameters struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	// +kubebuilder:validation:Optional
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	// +kubebuilder:validation:Optional
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	// +kubebuilder:validation:Optional
	BGPAsn *int64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	// +kubebuilder:validation:Optional
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/directconnect/v1beta1.Connection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Reference to a Connection in directconnect to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	// Selector for a Connection in directconnect to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	// +kubebuilder:validation:Optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/directconnect/v1beta1.Gateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// Reference to a Gateway in directconnect to populate dxGatewayId.
	// +kubebuilder:validation:Optional
	DxGatewayIDRef *v1.Reference `json:"dxGatewayIdRef,omitempty" tf:"-"`

	// Selector for a Gateway in directconnect to populate dxGatewayId.
	// +kubebuilder:validation:Optional
	DxGatewayIDSelector *v1.Selector `json:"dxGatewayIdSelector,omitempty" tf:"-"`

	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
	// The MTU of a virtual transit interface can be either 1500 or 8500 (jumbo frames). Default is 1500.
	// +kubebuilder:validation:Optional
	Mtu *int64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name for the virtual interface.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Indicates whether to enable or disable SiteLink.
	// +kubebuilder:validation:Optional
	SitelinkEnabled *bool `json:"sitelinkEnabled,omitempty" tf:"sitelink_enabled,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VLAN ID.
	// +kubebuilder:validation:Optional
	Vlan *int64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

// TransitVirtualInterfaceSpec defines the desired state of TransitVirtualInterface
type TransitVirtualInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitVirtualInterfaceParameters `json:"forProvider"`
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
	InitProvider TransitVirtualInterfaceInitParameters `json:"initProvider,omitempty"`
}

// TransitVirtualInterfaceStatus defines the observed state of TransitVirtualInterface.
type TransitVirtualInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitVirtualInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitVirtualInterface is the Schema for the TransitVirtualInterfaces API. Provides a Direct Connect transit virtual interface resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addressFamily) || (has(self.initProvider) && has(self.initProvider.addressFamily))",message="spec.forProvider.addressFamily is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bgpAsn) || (has(self.initProvider) && has(self.initProvider.bgpAsn))",message="spec.forProvider.bgpAsn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vlan) || (has(self.initProvider) && has(self.initProvider.vlan))",message="spec.forProvider.vlan is a required parameter"
	Spec   TransitVirtualInterfaceSpec   `json:"spec"`
	Status TransitVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitVirtualInterfaceList contains a list of TransitVirtualInterfaces
type TransitVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitVirtualInterface `json:"items"`
}

// Repository type metadata.
var (
	TransitVirtualInterface_Kind             = "TransitVirtualInterface"
	TransitVirtualInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitVirtualInterface_Kind}.String()
	TransitVirtualInterface_KindAPIVersion   = TransitVirtualInterface_Kind + "." + CRDGroupVersion.String()
	TransitVirtualInterface_GroupVersionKind = CRDGroupVersion.WithKind(TransitVirtualInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitVirtualInterface{}, &TransitVirtualInterfaceList{})
}
