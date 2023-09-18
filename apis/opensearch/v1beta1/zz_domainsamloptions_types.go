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

type DomainSAMLOptionsInitParameters struct {

	// SAML authentication options for an AWS OpenSearch Domain.
	SAMLOptions []SAMLOptionsInitParameters `json:"samlOptions,omitempty" tf:"saml_options,omitempty"`
}

type DomainSAMLOptionsObservation struct {

	// Name of the domain.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Name of the domain the SAML options are associated with.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// SAML authentication options for an AWS OpenSearch Domain.
	SAMLOptions []SAMLOptionsObservation `json:"samlOptions,omitempty" tf:"saml_options,omitempty"`
}

type DomainSAMLOptionsParameters struct {

	// Name of the domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/opensearch/v1beta1.Domain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("domain_name",false)
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Reference to a Domain in opensearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameRef *v1.Reference `json:"domainNameRef,omitempty" tf:"-"`

	// Selector for a Domain in opensearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameSelector *v1.Selector `json:"domainNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// SAML authentication options for an AWS OpenSearch Domain.
	// +kubebuilder:validation:Optional
	SAMLOptions []SAMLOptionsParameters `json:"samlOptions,omitempty" tf:"saml_options,omitempty"`
}

type IdpInitParameters struct {

	// Unique Entity ID of the application in SAML Identity Provider.
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// Metadata of the SAML application in xml format.
	MetadataContent *string `json:"metadataContent,omitempty" tf:"metadata_content,omitempty"`
}

type IdpObservation struct {

	// Unique Entity ID of the application in SAML Identity Provider.
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// Metadata of the SAML application in xml format.
	MetadataContent *string `json:"metadataContent,omitempty" tf:"metadata_content,omitempty"`
}

type IdpParameters struct {

	// Unique Entity ID of the application in SAML Identity Provider.
	// +kubebuilder:validation:Optional
	EntityID *string `json:"entityId" tf:"entity_id,omitempty"`

	// Metadata of the SAML application in xml format.
	// +kubebuilder:validation:Optional
	MetadataContent *string `json:"metadataContent" tf:"metadata_content,omitempty"`
}

type SAMLOptionsInitParameters struct {

	// Whether SAML authentication is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Information from your identity provider.
	Idp []IdpInitParameters `json:"idp,omitempty" tf:"idp,omitempty"`

	// This backend role from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
	MasterBackendRole *string `json:"masterBackendRole,omitempty" tf:"master_backend_role,omitempty"`

	// Element of the SAML assertion to use for backend roles. Default is roles.
	RolesKey *string `json:"rolesKey,omitempty" tf:"roles_key,omitempty"`

	// Duration of a session in minutes after a user logs in. Default is 60. Maximum value is 1,440.
	SessionTimeoutMinutes *int64 `json:"sessionTimeoutMinutes,omitempty" tf:"session_timeout_minutes,omitempty"`

	// Element of the SAML assertion to use for username. Default is NameID.
	SubjectKey *string `json:"subjectKey,omitempty" tf:"subject_key,omitempty"`
}

type SAMLOptionsObservation struct {

	// Whether SAML authentication is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Information from your identity provider.
	Idp []IdpObservation `json:"idp,omitempty" tf:"idp,omitempty"`

	// This backend role from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
	MasterBackendRole *string `json:"masterBackendRole,omitempty" tf:"master_backend_role,omitempty"`

	// Element of the SAML assertion to use for backend roles. Default is roles.
	RolesKey *string `json:"rolesKey,omitempty" tf:"roles_key,omitempty"`

	// Duration of a session in minutes after a user logs in. Default is 60. Maximum value is 1,440.
	SessionTimeoutMinutes *int64 `json:"sessionTimeoutMinutes,omitempty" tf:"session_timeout_minutes,omitempty"`

	// Element of the SAML assertion to use for username. Default is NameID.
	SubjectKey *string `json:"subjectKey,omitempty" tf:"subject_key,omitempty"`
}

type SAMLOptionsParameters struct {

	// Whether SAML authentication is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Information from your identity provider.
	// +kubebuilder:validation:Optional
	Idp []IdpParameters `json:"idp,omitempty" tf:"idp,omitempty"`

	// This backend role from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
	// +kubebuilder:validation:Optional
	MasterBackendRole *string `json:"masterBackendRole,omitempty" tf:"master_backend_role,omitempty"`

	// This username from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
	// +kubebuilder:validation:Optional
	MasterUserNameSecretRef *v1.SecretKeySelector `json:"masterUserNameSecretRef,omitempty" tf:"-"`

	// Element of the SAML assertion to use for backend roles. Default is roles.
	// +kubebuilder:validation:Optional
	RolesKey *string `json:"rolesKey,omitempty" tf:"roles_key,omitempty"`

	// Duration of a session in minutes after a user logs in. Default is 60. Maximum value is 1,440.
	// +kubebuilder:validation:Optional
	SessionTimeoutMinutes *int64 `json:"sessionTimeoutMinutes,omitempty" tf:"session_timeout_minutes,omitempty"`

	// Element of the SAML assertion to use for username. Default is NameID.
	// +kubebuilder:validation:Optional
	SubjectKey *string `json:"subjectKey,omitempty" tf:"subject_key,omitempty"`
}

// DomainSAMLOptionsSpec defines the desired state of DomainSAMLOptions
type DomainSAMLOptionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainSAMLOptionsParameters `json:"forProvider"`
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
	InitProvider DomainSAMLOptionsInitParameters `json:"initProvider,omitempty"`
}

// DomainSAMLOptionsStatus defines the observed state of DomainSAMLOptions.
type DomainSAMLOptionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainSAMLOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainSAMLOptions is the Schema for the DomainSAMLOptionss API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DomainSAMLOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSAMLOptionsSpec   `json:"spec"`
	Status            DomainSAMLOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainSAMLOptionsList contains a list of DomainSAMLOptionss
type DomainSAMLOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainSAMLOptions `json:"items"`
}

// Repository type metadata.
var (
	DomainSAMLOptions_Kind             = "DomainSAMLOptions"
	DomainSAMLOptions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainSAMLOptions_Kind}.String()
	DomainSAMLOptions_KindAPIVersion   = DomainSAMLOptions_Kind + "." + CRDGroupVersion.String()
	DomainSAMLOptions_GroupVersionKind = CRDGroupVersion.WithKind(DomainSAMLOptions_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainSAMLOptions{}, &DomainSAMLOptionsList{})
}
