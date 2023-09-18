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

type CidrAuthorizationContextInitParameters struct {

	// The plain-text authorization message for the prefix and account.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The signed authorization message for the prefix and account.
	Signature *string `json:"signature,omitempty" tf:"signature,omitempty"`
}

type CidrAuthorizationContextObservation struct {

	// The plain-text authorization message for the prefix and account.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The signed authorization message for the prefix and account.
	Signature *string `json:"signature,omitempty" tf:"signature,omitempty"`
}

type CidrAuthorizationContextParameters struct {

	// The plain-text authorization message for the prefix and account.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The signed authorization message for the prefix and account.
	// +kubebuilder:validation:Optional
	Signature *string `json:"signature,omitempty" tf:"signature,omitempty"`
}

type VPCIpamPoolCidrInitParameters struct {

	// The CIDR you want to assign to the pool. Conflicts with netmask_length.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A signed document that proves that you are authorized to bring the specified IP address range to Amazon using BYOIP. This is not stored in the state file. See cidr_authorization_context for more information.
	CidrAuthorizationContext []CidrAuthorizationContextInitParameters `json:"cidrAuthorizationContext,omitempty" tf:"cidr_authorization_context,omitempty"`

	// If provided, the cidr provisioned into the specified pool will be the next available cidr given this declared netmask length. Conflicts with cidr.
	NetmaskLength *int64 `json:"netmaskLength,omitempty" tf:"netmask_length,omitempty"`
}

type VPCIpamPoolCidrObservation struct {

	// The CIDR you want to assign to the pool. Conflicts with netmask_length.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A signed document that proves that you are authorized to bring the specified IP address range to Amazon using BYOIP. This is not stored in the state file. See cidr_authorization_context for more information.
	CidrAuthorizationContext []CidrAuthorizationContextObservation `json:"cidrAuthorizationContext,omitempty" tf:"cidr_authorization_context,omitempty"`

	// The ID of the IPAM Pool Cidr concatenated with the IPAM Pool ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique ID generated by AWS for the pool cidr.
	IpamPoolCidrID *string `json:"ipamPoolCidrId,omitempty" tf:"ipam_pool_cidr_id,omitempty"`

	// The ID of the pool to which you want to assign a CIDR.
	IpamPoolID *string `json:"ipamPoolId,omitempty" tf:"ipam_pool_id,omitempty"`

	// If provided, the cidr provisioned into the specified pool will be the next available cidr given this declared netmask length. Conflicts with cidr.
	NetmaskLength *int64 `json:"netmaskLength,omitempty" tf:"netmask_length,omitempty"`
}

type VPCIpamPoolCidrParameters struct {

	// The CIDR you want to assign to the pool. Conflicts with netmask_length.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// A signed document that proves that you are authorized to bring the specified IP address range to Amazon using BYOIP. This is not stored in the state file. See cidr_authorization_context for more information.
	// +kubebuilder:validation:Optional
	CidrAuthorizationContext []CidrAuthorizationContextParameters `json:"cidrAuthorizationContext,omitempty" tf:"cidr_authorization_context,omitempty"`

	// The ID of the pool to which you want to assign a CIDR.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCIpamPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IpamPoolID *string `json:"ipamPoolId,omitempty" tf:"ipam_pool_id,omitempty"`

	// Reference to a VPCIpamPool in ec2 to populate ipamPoolId.
	// +kubebuilder:validation:Optional
	IpamPoolIDRef *v1.Reference `json:"ipamPoolIdRef,omitempty" tf:"-"`

	// Selector for a VPCIpamPool in ec2 to populate ipamPoolId.
	// +kubebuilder:validation:Optional
	IpamPoolIDSelector *v1.Selector `json:"ipamPoolIdSelector,omitempty" tf:"-"`

	// If provided, the cidr provisioned into the specified pool will be the next available cidr given this declared netmask length. Conflicts with cidr.
	// +kubebuilder:validation:Optional
	NetmaskLength *int64 `json:"netmaskLength,omitempty" tf:"netmask_length,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// VPCIpamPoolCidrSpec defines the desired state of VPCIpamPoolCidr
type VPCIpamPoolCidrSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCIpamPoolCidrParameters `json:"forProvider"`
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
	InitProvider VPCIpamPoolCidrInitParameters `json:"initProvider,omitempty"`
}

// VPCIpamPoolCidrStatus defines the observed state of VPCIpamPoolCidr.
type VPCIpamPoolCidrStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCIpamPoolCidrObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIpamPoolCidr is the Schema for the VPCIpamPoolCidrs API. Provisions a CIDR from an IPAM address pool.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCIpamPoolCidr struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCIpamPoolCidrSpec   `json:"spec"`
	Status            VPCIpamPoolCidrStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIpamPoolCidrList contains a list of VPCIpamPoolCidrs
type VPCIpamPoolCidrList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCIpamPoolCidr `json:"items"`
}

// Repository type metadata.
var (
	VPCIpamPoolCidr_Kind             = "VPCIpamPoolCidr"
	VPCIpamPoolCidr_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCIpamPoolCidr_Kind}.String()
	VPCIpamPoolCidr_KindAPIVersion   = VPCIpamPoolCidr_Kind + "." + CRDGroupVersion.String()
	VPCIpamPoolCidr_GroupVersionKind = CRDGroupVersion.WithKind(VPCIpamPoolCidr_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCIpamPoolCidr{}, &VPCIpamPoolCidrList{})
}
