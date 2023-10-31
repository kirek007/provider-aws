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

type CoreNetworkInitParameters struct {

	// The base policy created by setting the create_base_policy argument to true requires a region to be set in the edge-locations, location key. If base_policy_region is not specified, the region used in the base policy defaults to the region specified in the provider block.
	BasePolicyRegion *string `json:"basePolicyRegion,omitempty" tf:"base_policy_region,omitempty"`

	// A list of regions to add to the base policy. The base policy created by setting the create_base_policy argument to true requires one or more regions to be set in the edge-locations, location key. If base_policy_regions is not specified, the region used in the base policy defaults to the region specified in the provider block.
	BasePolicyRegions []*string `json:"basePolicyRegions,omitempty" tf:"base_policy_regions,omitempty"`

	// Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to LIVE to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the aws_networkmanager_core_network_policy_attachment resource. This base policy is needed if your core network does not have any LIVE policies (e.g. a core network resource created without the policy_document argument) and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are true or false. Conflicts with policy_document. An example base policy is shown below. This base policy is overridden with the policy that you specify in the aws_networkmanager_core_network_policy_attachment resource.
	CreateBasePolicy *bool `json:"createBasePolicy,omitempty" tf:"create_base_policy,omitempty"`

	// Description of the Core Network.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the LATEST and LIVE policy document. Refer to the Core network policies documentation for more information. Conflicts with create_base_policy.
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CoreNetworkObservation struct {

	// Core Network Amazon Resource Name (ARN).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The base policy created by setting the create_base_policy argument to true requires a region to be set in the edge-locations, location key. If base_policy_region is not specified, the region used in the base policy defaults to the region specified in the provider block.
	BasePolicyRegion *string `json:"basePolicyRegion,omitempty" tf:"base_policy_region,omitempty"`

	// A list of regions to add to the base policy. The base policy created by setting the create_base_policy argument to true requires one or more regions to be set in the edge-locations, location key. If base_policy_regions is not specified, the region used in the base policy defaults to the region specified in the provider block.
	BasePolicyRegions []*string `json:"basePolicyRegions,omitempty" tf:"base_policy_regions,omitempty"`

	// Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to LIVE to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the aws_networkmanager_core_network_policy_attachment resource. This base policy is needed if your core network does not have any LIVE policies (e.g. a core network resource created without the policy_document argument) and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are true or false. Conflicts with policy_document. An example base policy is shown below. This base policy is overridden with the policy that you specify in the aws_networkmanager_core_network_policy_attachment resource.
	CreateBasePolicy *bool `json:"createBasePolicy,omitempty" tf:"create_base_policy,omitempty"`

	// Timestamp when a core network was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Description of the Core Network.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more blocks detailing the edges within a core network. Detailed below.
	Edges []EdgesObservation `json:"edges,omitempty" tf:"edges,omitempty"`

	// The ID of the global network that a core network will be a part of.
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	// Core Network ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the LATEST and LIVE policy document. Refer to the Core network policies documentation for more information. Conflicts with create_base_policy.
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// One or more blocks detailing the segments within a core network. Detailed below.
	Segments []SegmentsObservation `json:"segments,omitempty" tf:"segments,omitempty"`

	// Current state of a core network.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CoreNetworkParameters struct {

	// The base policy created by setting the create_base_policy argument to true requires a region to be set in the edge-locations, location key. If base_policy_region is not specified, the region used in the base policy defaults to the region specified in the provider block.
	// +kubebuilder:validation:Optional
	BasePolicyRegion *string `json:"basePolicyRegion,omitempty" tf:"base_policy_region,omitempty"`

	// A list of regions to add to the base policy. The base policy created by setting the create_base_policy argument to true requires one or more regions to be set in the edge-locations, location key. If base_policy_regions is not specified, the region used in the base policy defaults to the region specified in the provider block.
	// +kubebuilder:validation:Optional
	BasePolicyRegions []*string `json:"basePolicyRegions,omitempty" tf:"base_policy_regions,omitempty"`

	// Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to LIVE to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the aws_networkmanager_core_network_policy_attachment resource. This base policy is needed if your core network does not have any LIVE policies (e.g. a core network resource created without the policy_document argument) and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are true or false. Conflicts with policy_document. An example base policy is shown below. This base policy is overridden with the policy that you specify in the aws_networkmanager_core_network_policy_attachment resource.
	// +kubebuilder:validation:Optional
	CreateBasePolicy *bool `json:"createBasePolicy,omitempty" tf:"create_base_policy,omitempty"`

	// Description of the Core Network.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the global network that a core network will be a part of.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.GlobalNetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	// Reference to a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDRef *v1.Reference `json:"globalNetworkIdRef,omitempty" tf:"-"`

	// Selector for a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDSelector *v1.Selector `json:"globalNetworkIdSelector,omitempty" tf:"-"`

	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the LATEST and LIVE policy document. Refer to the Core network policies documentation for more information. Conflicts with create_base_policy.
	// +kubebuilder:validation:Optional
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EdgesInitParameters struct {
}

type EdgesObservation struct {

	// ASN of a core network edge.
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// Region where a core network edge is located.
	EdgeLocation *string `json:"edgeLocation,omitempty" tf:"edge_location,omitempty"`

	// Inside IP addresses used for core network edges.
	InsideCidrBlocks []*string `json:"insideCidrBlocks,omitempty" tf:"inside_cidr_blocks,omitempty"`
}

type EdgesParameters struct {
}

type SegmentsInitParameters struct {
}

type SegmentsObservation struct {

	// Regions where the edges are located.
	EdgeLocations []*string `json:"edgeLocations,omitempty" tf:"edge_locations,omitempty"`

	// Name of a core network segment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Shared segments of a core network.
	SharedSegments []*string `json:"sharedSegments,omitempty" tf:"shared_segments,omitempty"`
}

type SegmentsParameters struct {
}

// CoreNetworkSpec defines the desired state of CoreNetwork
type CoreNetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CoreNetworkParameters `json:"forProvider"`
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
	InitProvider CoreNetworkInitParameters `json:"initProvider,omitempty"`
}

// CoreNetworkStatus defines the observed state of CoreNetwork.
type CoreNetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CoreNetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CoreNetwork is the Schema for the CoreNetworks API. Provides a core network resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CoreNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CoreNetworkSpec   `json:"spec"`
	Status            CoreNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CoreNetworkList contains a list of CoreNetworks
type CoreNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CoreNetwork `json:"items"`
}

// Repository type metadata.
var (
	CoreNetwork_Kind             = "CoreNetwork"
	CoreNetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CoreNetwork_Kind}.String()
	CoreNetwork_KindAPIVersion   = CoreNetwork_Kind + "." + CRDGroupVersion.String()
	CoreNetwork_GroupVersionKind = CRDGroupVersion.WithKind(CoreNetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&CoreNetwork{}, &CoreNetworkList{})
}
