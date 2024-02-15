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

type SigningProfilePermissionInitParameters struct {

	// An AWS Signer action permitted as part of cross-account permissions. Valid values: signer:StartSigningJob, signer:GetSigningProfile, or signer:RevokeSignature.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The AWS principal to be granted a cross-account permission.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// The signing profile version that a permission applies to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/signer/v1beta1.SigningProfile
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("version",true)
	ProfileVersion *string `json:"profileVersion,omitempty" tf:"profile_version,omitempty"`

	// Reference to a SigningProfile in signer to populate profileVersion.
	// +kubebuilder:validation:Optional
	ProfileVersionRef *v1.Reference `json:"profileVersionRef,omitempty" tf:"-"`

	// Selector for a SigningProfile in signer to populate profileVersion.
	// +kubebuilder:validation:Optional
	ProfileVersionSelector *v1.Selector `json:"profileVersionSelector,omitempty" tf:"-"`

	// A statement identifier prefix. Conflicts with statement_id.
	StatementIDPrefix *string `json:"statementIdPrefix,omitempty" tf:"statement_id_prefix,omitempty"`
}

type SigningProfilePermissionObservation struct {

	// An AWS Signer action permitted as part of cross-account permissions. Valid values: signer:StartSigningJob, signer:GetSigningProfile, or signer:RevokeSignature.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The AWS principal to be granted a cross-account permission.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// Name of the signing profile to add the cross-account permissions.
	ProfileName *string `json:"profileName,omitempty" tf:"profile_name,omitempty"`

	// The signing profile version that a permission applies to.
	ProfileVersion *string `json:"profileVersion,omitempty" tf:"profile_version,omitempty"`

	// A unique statement identifier.
	StatementID *string `json:"statementId,omitempty" tf:"statement_id,omitempty"`

	// A statement identifier prefix. Conflicts with statement_id.
	StatementIDPrefix *string `json:"statementIdPrefix,omitempty" tf:"statement_id_prefix,omitempty"`
}

type SigningProfilePermissionParameters struct {

	// An AWS Signer action permitted as part of cross-account permissions. Valid values: signer:StartSigningJob, signer:GetSigningProfile, or signer:RevokeSignature.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The AWS principal to be granted a cross-account permission.
	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// Name of the signing profile to add the cross-account permissions.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/signer/v1beta1.SigningProfile
	// +kubebuilder:validation:Optional
	ProfileName *string `json:"profileName,omitempty" tf:"profile_name,omitempty"`

	// Reference to a SigningProfile in signer to populate profileName.
	// +kubebuilder:validation:Optional
	ProfileNameRef *v1.Reference `json:"profileNameRef,omitempty" tf:"-"`

	// Selector for a SigningProfile in signer to populate profileName.
	// +kubebuilder:validation:Optional
	ProfileNameSelector *v1.Selector `json:"profileNameSelector,omitempty" tf:"-"`

	// The signing profile version that a permission applies to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/signer/v1beta1.SigningProfile
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("version",true)
	// +kubebuilder:validation:Optional
	ProfileVersion *string `json:"profileVersion,omitempty" tf:"profile_version,omitempty"`

	// Reference to a SigningProfile in signer to populate profileVersion.
	// +kubebuilder:validation:Optional
	ProfileVersionRef *v1.Reference `json:"profileVersionRef,omitempty" tf:"-"`

	// Selector for a SigningProfile in signer to populate profileVersion.
	// +kubebuilder:validation:Optional
	ProfileVersionSelector *v1.Selector `json:"profileVersionSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A unique statement identifier.
	// +kubebuilder:validation:Optional
	StatementID *string `json:"statementId,omitempty" tf:"statement_id,omitempty"`

	// A statement identifier prefix. Conflicts with statement_id.
	// +kubebuilder:validation:Optional
	StatementIDPrefix *string `json:"statementIdPrefix,omitempty" tf:"statement_id_prefix,omitempty"`
}

// SigningProfilePermissionSpec defines the desired state of SigningProfilePermission
type SigningProfilePermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SigningProfilePermissionParameters `json:"forProvider"`
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
	InitProvider SigningProfilePermissionInitParameters `json:"initProvider,omitempty"`
}

// SigningProfilePermissionStatus defines the observed state of SigningProfilePermission.
type SigningProfilePermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SigningProfilePermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SigningProfilePermission is the Schema for the SigningProfilePermissions API. Creates a Signer Signing Profile Permission.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SigningProfilePermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principal) || (has(self.initProvider) && has(self.initProvider.principal))",message="spec.forProvider.principal is a required parameter"
	Spec   SigningProfilePermissionSpec   `json:"spec"`
	Status SigningProfilePermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SigningProfilePermissionList contains a list of SigningProfilePermissions
type SigningProfilePermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SigningProfilePermission `json:"items"`
}

// Repository type metadata.
var (
	SigningProfilePermission_Kind             = "SigningProfilePermission"
	SigningProfilePermission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SigningProfilePermission_Kind}.String()
	SigningProfilePermission_KindAPIVersion   = SigningProfilePermission_Kind + "." + CRDGroupVersion.String()
	SigningProfilePermission_GroupVersionKind = CRDGroupVersion.WithKind(SigningProfilePermission_Kind)
)

func init() {
	SchemeBuilder.Register(&SigningProfilePermission{}, &SigningProfilePermissionList{})
}
