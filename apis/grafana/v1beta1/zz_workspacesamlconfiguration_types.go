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

type WorkspaceSAMLConfigurationInitParameters struct {

	// The admin role values.
	AdminRoleValues []*string `json:"adminRoleValues,omitempty" tf:"admin_role_values,omitempty"`

	// The allowed organizations.
	AllowedOrganizations []*string `json:"allowedOrganizations,omitempty" tf:"allowed_organizations,omitempty"`

	// The editor role values.
	EditorRoleValues []*string `json:"editorRoleValues,omitempty" tf:"editor_role_values,omitempty"`

	// The email assertion.
	EmailAssertion *string `json:"emailAssertion,omitempty" tf:"email_assertion,omitempty"`

	// The groups assertion.
	GroupsAssertion *string `json:"groupsAssertion,omitempty" tf:"groups_assertion,omitempty"`

	// The IDP Metadata URL. Note that either idp_metadata_url or idp_metadata_xml (but not both) must be specified.
	IdpMetadataURL *string `json:"idpMetadataUrl,omitempty" tf:"idp_metadata_url,omitempty"`

	// The IDP Metadata XML. Note that either idp_metadata_url or idp_metadata_xml (but not both) must be specified.
	IdpMetadataXML *string `json:"idpMetadataXml,omitempty" tf:"idp_metadata_xml,omitempty"`

	// The login assertion.
	LoginAssertion *string `json:"loginAssertion,omitempty" tf:"login_assertion,omitempty"`

	// The login validity duration.
	LoginValidityDuration *int64 `json:"loginValidityDuration,omitempty" tf:"login_validity_duration,omitempty"`

	// The name assertion.
	NameAssertion *string `json:"nameAssertion,omitempty" tf:"name_assertion,omitempty"`

	// The org assertion.
	OrgAssertion *string `json:"orgAssertion,omitempty" tf:"org_assertion,omitempty"`

	// The role assertion.
	RoleAssertion *string `json:"roleAssertion,omitempty" tf:"role_assertion,omitempty"`
}

type WorkspaceSAMLConfigurationObservation struct {

	// The admin role values.
	AdminRoleValues []*string `json:"adminRoleValues,omitempty" tf:"admin_role_values,omitempty"`

	// The allowed organizations.
	AllowedOrganizations []*string `json:"allowedOrganizations,omitempty" tf:"allowed_organizations,omitempty"`

	// The editor role values.
	EditorRoleValues []*string `json:"editorRoleValues,omitempty" tf:"editor_role_values,omitempty"`

	// The email assertion.
	EmailAssertion *string `json:"emailAssertion,omitempty" tf:"email_assertion,omitempty"`

	// The groups assertion.
	GroupsAssertion *string `json:"groupsAssertion,omitempty" tf:"groups_assertion,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IDP Metadata URL. Note that either idp_metadata_url or idp_metadata_xml (but not both) must be specified.
	IdpMetadataURL *string `json:"idpMetadataUrl,omitempty" tf:"idp_metadata_url,omitempty"`

	// The IDP Metadata XML. Note that either idp_metadata_url or idp_metadata_xml (but not both) must be specified.
	IdpMetadataXML *string `json:"idpMetadataXml,omitempty" tf:"idp_metadata_xml,omitempty"`

	// The login assertion.
	LoginAssertion *string `json:"loginAssertion,omitempty" tf:"login_assertion,omitempty"`

	// The login validity duration.
	LoginValidityDuration *int64 `json:"loginValidityDuration,omitempty" tf:"login_validity_duration,omitempty"`

	// The name assertion.
	NameAssertion *string `json:"nameAssertion,omitempty" tf:"name_assertion,omitempty"`

	// The org assertion.
	OrgAssertion *string `json:"orgAssertion,omitempty" tf:"org_assertion,omitempty"`

	// The role assertion.
	RoleAssertion *string `json:"roleAssertion,omitempty" tf:"role_assertion,omitempty"`

	// The status of the SAML configuration.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The workspace id.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type WorkspaceSAMLConfigurationParameters struct {

	// The admin role values.
	// +kubebuilder:validation:Optional
	AdminRoleValues []*string `json:"adminRoleValues,omitempty" tf:"admin_role_values,omitempty"`

	// The allowed organizations.
	// +kubebuilder:validation:Optional
	AllowedOrganizations []*string `json:"allowedOrganizations,omitempty" tf:"allowed_organizations,omitempty"`

	// The editor role values.
	// +kubebuilder:validation:Optional
	EditorRoleValues []*string `json:"editorRoleValues,omitempty" tf:"editor_role_values,omitempty"`

	// The email assertion.
	// +kubebuilder:validation:Optional
	EmailAssertion *string `json:"emailAssertion,omitempty" tf:"email_assertion,omitempty"`

	// The groups assertion.
	// +kubebuilder:validation:Optional
	GroupsAssertion *string `json:"groupsAssertion,omitempty" tf:"groups_assertion,omitempty"`

	// The IDP Metadata URL. Note that either idp_metadata_url or idp_metadata_xml (but not both) must be specified.
	// +kubebuilder:validation:Optional
	IdpMetadataURL *string `json:"idpMetadataUrl,omitempty" tf:"idp_metadata_url,omitempty"`

	// The IDP Metadata XML. Note that either idp_metadata_url or idp_metadata_xml (but not both) must be specified.
	// +kubebuilder:validation:Optional
	IdpMetadataXML *string `json:"idpMetadataXml,omitempty" tf:"idp_metadata_xml,omitempty"`

	// The login assertion.
	// +kubebuilder:validation:Optional
	LoginAssertion *string `json:"loginAssertion,omitempty" tf:"login_assertion,omitempty"`

	// The login validity duration.
	// +kubebuilder:validation:Optional
	LoginValidityDuration *int64 `json:"loginValidityDuration,omitempty" tf:"login_validity_duration,omitempty"`

	// The name assertion.
	// +kubebuilder:validation:Optional
	NameAssertion *string `json:"nameAssertion,omitempty" tf:"name_assertion,omitempty"`

	// The org assertion.
	// +kubebuilder:validation:Optional
	OrgAssertion *string `json:"orgAssertion,omitempty" tf:"org_assertion,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The role assertion.
	// +kubebuilder:validation:Optional
	RoleAssertion *string `json:"roleAssertion,omitempty" tf:"role_assertion,omitempty"`

	// The workspace id.
	// +crossplane:generate:reference:type=Workspace
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

// WorkspaceSAMLConfigurationSpec defines the desired state of WorkspaceSAMLConfiguration
type WorkspaceSAMLConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceSAMLConfigurationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider WorkspaceSAMLConfigurationInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceSAMLConfigurationStatus defines the observed state of WorkspaceSAMLConfiguration.
type WorkspaceSAMLConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceSAMLConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceSAMLConfiguration is the Schema for the WorkspaceSAMLConfigurations API. Provides an Amazon Managed Grafana workspace SAML configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WorkspaceSAMLConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.editorRoleValues) || has(self.initProvider.editorRoleValues)",message="editorRoleValues is a required parameter"
	Spec   WorkspaceSAMLConfigurationSpec   `json:"spec"`
	Status WorkspaceSAMLConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceSAMLConfigurationList contains a list of WorkspaceSAMLConfigurations
type WorkspaceSAMLConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceSAMLConfiguration `json:"items"`
}

// Repository type metadata.
var (
	WorkspaceSAMLConfiguration_Kind             = "WorkspaceSAMLConfiguration"
	WorkspaceSAMLConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkspaceSAMLConfiguration_Kind}.String()
	WorkspaceSAMLConfiguration_KindAPIVersion   = WorkspaceSAMLConfiguration_Kind + "." + CRDGroupVersion.String()
	WorkspaceSAMLConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(WorkspaceSAMLConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkspaceSAMLConfiguration{}, &WorkspaceSAMLConfigurationList{})
}
