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

type AwsLocationInitParameters struct {

	// The Amazon Resource Name (ARN) of the subnet that the device is located in.
	SubnetArn *string `json:"subnetArn,omitempty" tf:"subnet_arn,omitempty"`

	// The Zone that the device is located in. Specify the ID of an Availability Zone, Local Zone, Wavelength Zone, or an Outpost.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type AwsLocationObservation struct {

	// The Amazon Resource Name (ARN) of the subnet that the device is located in.
	SubnetArn *string `json:"subnetArn,omitempty" tf:"subnet_arn,omitempty"`

	// The Zone that the device is located in. Specify the ID of an Availability Zone, Local Zone, Wavelength Zone, or an Outpost.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type AwsLocationParameters struct {

	// The Amazon Resource Name (ARN) of the subnet that the device is located in.
	// +kubebuilder:validation:Optional
	SubnetArn *string `json:"subnetArn,omitempty" tf:"subnet_arn,omitempty"`

	// The Zone that the device is located in. Specify the ID of an Availability Zone, Local Zone, Wavelength Zone, or an Outpost.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DeviceInitParameters struct {

	// The AWS location of the device. Documented below.
	AwsLocation []AwsLocationInitParameters `json:"awsLocation,omitempty" tf:"aws_location,omitempty"`

	// A description of the device.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the global network.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.GlobalNetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	// Reference to a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDRef *v1.Reference `json:"globalNetworkIdRef,omitempty" tf:"-"`

	// Selector for a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDSelector *v1.Selector `json:"globalNetworkIdSelector,omitempty" tf:"-"`

	// The location of the device. Documented below.
	Location []LocationInitParameters `json:"location,omitempty" tf:"location,omitempty"`

	// The model of device.
	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	// The serial number of the device.
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	// The ID of the site.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.Site
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SiteID *string `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// Reference to a Site in networkmanager to populate siteId.
	// +kubebuilder:validation:Optional
	SiteIDRef *v1.Reference `json:"siteIdRef,omitempty" tf:"-"`

	// Selector for a Site in networkmanager to populate siteId.
	// +kubebuilder:validation:Optional
	SiteIDSelector *v1.Selector `json:"siteIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of device.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The vendor of the device.
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`
}

type DeviceObservation struct {

	// The Amazon Resource Name (ARN) of the device.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The AWS location of the device. Documented below.
	AwsLocation []AwsLocationObservation `json:"awsLocation,omitempty" tf:"aws_location,omitempty"`

	// A description of the device.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the global network.
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the device. Documented below.
	Location []LocationObservation `json:"location,omitempty" tf:"location,omitempty"`

	// The model of device.
	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	// The serial number of the device.
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	// The ID of the site.
	SiteID *string `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The type of device.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The vendor of the device.
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`
}

type DeviceParameters struct {

	// The AWS location of the device. Documented below.
	// +kubebuilder:validation:Optional
	AwsLocation []AwsLocationParameters `json:"awsLocation,omitempty" tf:"aws_location,omitempty"`

	// A description of the device.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the global network.
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

	// The location of the device. Documented below.
	// +kubebuilder:validation:Optional
	Location []LocationParameters `json:"location,omitempty" tf:"location,omitempty"`

	// The model of device.
	// +kubebuilder:validation:Optional
	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The serial number of the device.
	// +kubebuilder:validation:Optional
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	// The ID of the site.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.Site
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SiteID *string `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// Reference to a Site in networkmanager to populate siteId.
	// +kubebuilder:validation:Optional
	SiteIDRef *v1.Reference `json:"siteIdRef,omitempty" tf:"-"`

	// Selector for a Site in networkmanager to populate siteId.
	// +kubebuilder:validation:Optional
	SiteIDSelector *v1.Selector `json:"siteIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of device.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The vendor of the device.
	// +kubebuilder:validation:Optional
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`
}

type LocationInitParameters struct {

	// The physical address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The latitude.
	Latitude *string `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// The longitude.
	Longitude *string `json:"longitude,omitempty" tf:"longitude,omitempty"`
}

type LocationObservation struct {

	// The physical address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The latitude.
	Latitude *string `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// The longitude.
	Longitude *string `json:"longitude,omitempty" tf:"longitude,omitempty"`
}

type LocationParameters struct {

	// The physical address.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The latitude.
	// +kubebuilder:validation:Optional
	Latitude *string `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// The longitude.
	// +kubebuilder:validation:Optional
	Longitude *string `json:"longitude,omitempty" tf:"longitude,omitempty"`
}

// DeviceSpec defines the desired state of Device
type DeviceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeviceParameters `json:"forProvider"`
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
	InitProvider DeviceInitParameters `json:"initProvider,omitempty"`
}

// DeviceStatus defines the observed state of Device.
type DeviceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeviceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Device is the Schema for the Devices API. Creates a device in a global network.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Device struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeviceSpec   `json:"spec"`
	Status            DeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceList contains a list of Devices
type DeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Device `json:"items"`
}

// Repository type metadata.
var (
	Device_Kind             = "Device"
	Device_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Device_Kind}.String()
	Device_KindAPIVersion   = Device_Kind + "." + CRDGroupVersion.String()
	Device_GroupVersionKind = CRDGroupVersion.WithKind(Device_Kind)
)

func init() {
	SchemeBuilder.Register(&Device{}, &DeviceList{})
}
