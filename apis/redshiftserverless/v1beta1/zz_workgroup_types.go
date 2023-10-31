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

type ConfigParameterInitParameters struct {

	// The key of the parameter. The options are datestyle, enable_user_activity_logging, query_group, search_path, and max_query_execution_time.
	ParameterKey *string `json:"parameterKey,omitempty" tf:"parameter_key,omitempty"`

	// The value of the parameter to set.
	ParameterValue *string `json:"parameterValue,omitempty" tf:"parameter_value,omitempty"`
}

type ConfigParameterObservation struct {

	// The key of the parameter. The options are datestyle, enable_user_activity_logging, query_group, search_path, and max_query_execution_time.
	ParameterKey *string `json:"parameterKey,omitempty" tf:"parameter_key,omitempty"`

	// The value of the parameter to set.
	ParameterValue *string `json:"parameterValue,omitempty" tf:"parameter_value,omitempty"`
}

type ConfigParameterParameters struct {

	// The key of the parameter. The options are datestyle, enable_user_activity_logging, query_group, search_path, and max_query_execution_time.
	// +kubebuilder:validation:Optional
	ParameterKey *string `json:"parameterKey" tf:"parameter_key,omitempty"`

	// The value of the parameter to set.
	// +kubebuilder:validation:Optional
	ParameterValue *string `json:"parameterValue" tf:"parameter_value,omitempty"`
}

type EndpointInitParameters struct {
}

type EndpointObservation struct {

	// The DNS address of the VPC endpoint.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The port that Amazon Redshift Serverless listens on.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The VPC endpoint or the Redshift Serverless workgroup. See VPC Endpoint below.
	VPCEndpoint []EndpointVPCEndpointObservation `json:"vpcEndpoint,omitempty" tf:"vpc_endpoint,omitempty"`
}

type EndpointParameters struct {
}

type EndpointVPCEndpointInitParameters struct {
}

type EndpointVPCEndpointObservation struct {

	// The network interfaces of the endpoint.. See Network Interface below.
	NetworkInterface []VPCEndpointNetworkInterfaceObservation `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// The DNS address of the VPC endpoint.
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The port that Amazon Redshift Serverless listens on.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type EndpointVPCEndpointParameters struct {
}

type VPCEndpointNetworkInterfaceInitParameters struct {
}

type VPCEndpointNetworkInterfaceObservation struct {

	// The availability Zone.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The unique identifier of the network interface.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The IPv4 address of the network interface within the subnet.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The unique identifier of the subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type VPCEndpointNetworkInterfaceParameters struct {
}

type WorkgroupInitParameters struct {

	// The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
	BaseCapacity *float64 `json:"baseCapacity,omitempty" tf:"base_capacity,omitempty"`

	// An array of parameters to set for more control over a serverless database. See Config Parameter below.
	ConfigParameter []ConfigParameterInitParameters `json:"configParameter,omitempty" tf:"config_parameter,omitempty"`

	// The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.
	EnhancedVPCRouting *bool `json:"enhancedVpcRouting,omitempty" tf:"enhanced_vpc_routing,omitempty"`

	// The name of the namespace.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// A value that specifies whether the workgroup can be accessed from a public network.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WorkgroupObservation struct {

	// Amazon Resource Name (ARN) of the Redshift Serverless Workgroup.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
	BaseCapacity *float64 `json:"baseCapacity,omitempty" tf:"base_capacity,omitempty"`

	// An array of parameters to set for more control over a serverless database. See Config Parameter below.
	ConfigParameter []ConfigParameterObservation `json:"configParameter,omitempty" tf:"config_parameter,omitempty"`

	// The endpoint that is created from the workgroup. See Endpoint below.
	Endpoint []EndpointObservation `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.
	EnhancedVPCRouting *bool `json:"enhancedVpcRouting,omitempty" tf:"enhanced_vpc_routing,omitempty"`

	// The Redshift Workgroup Name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the namespace.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// A value that specifies whether the workgroup can be accessed from a public network.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// An array of security group IDs to associate with the workgroup.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// An array of VPC subnet IDs to associate with the workgroup. When set, must contain at least three subnets spanning three Availability Zones. A minimum number of IP addresses is required and scales with the Base Capacity. For more information, see the following AWS document.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The Redshift Workgroup ID.
	WorkgroupID *string `json:"workgroupId,omitempty" tf:"workgroup_id,omitempty"`
}

type WorkgroupParameters struct {

	// The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
	// +kubebuilder:validation:Optional
	BaseCapacity *float64 `json:"baseCapacity,omitempty" tf:"base_capacity,omitempty"`

	// An array of parameters to set for more control over a serverless database. See Config Parameter below.
	// +kubebuilder:validation:Optional
	ConfigParameter []ConfigParameterParameters `json:"configParameter,omitempty" tf:"config_parameter,omitempty"`

	// The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.
	// +kubebuilder:validation:Optional
	EnhancedVPCRouting *bool `json:"enhancedVpcRouting,omitempty" tf:"enhanced_vpc_routing,omitempty"`

	// The name of the namespace.
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// A value that specifies whether the workgroup can be accessed from a public network.
	// +kubebuilder:validation:Optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// An array of security group IDs to associate with the workgroup.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// An array of VPC subnet IDs to associate with the workgroup. When set, must contain at least three subnets spanning three Availability Zones. A minimum number of IP addresses is required and scales with the Base Capacity. For more information, see the following AWS document.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkgroupSpec defines the desired state of Workgroup
type WorkgroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkgroupParameters `json:"forProvider"`
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
	InitProvider WorkgroupInitParameters `json:"initProvider,omitempty"`
}

// WorkgroupStatus defines the observed state of Workgroup.
type WorkgroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkgroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workgroup is the Schema for the Workgroups API. Provides a Redshift Serverless Workgroup resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Workgroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.namespaceName) || (has(self.initProvider) && has(self.initProvider.namespaceName))",message="spec.forProvider.namespaceName is a required parameter"
	Spec   WorkgroupSpec   `json:"spec"`
	Status WorkgroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkgroupList contains a list of Workgroups
type WorkgroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workgroup `json:"items"`
}

// Repository type metadata.
var (
	Workgroup_Kind             = "Workgroup"
	Workgroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workgroup_Kind}.String()
	Workgroup_KindAPIVersion   = Workgroup_Kind + "." + CRDGroupVersion.String()
	Workgroup_GroupVersionKind = CRDGroupVersion.WithKind(Workgroup_Kind)
)

func init() {
	SchemeBuilder.Register(&Workgroup{}, &WorkgroupList{})
}
