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

type PreProvisioningHookInitParameters struct {

	// The version of the payload that was sent to the target function. The only valid (and the default) payload version is "2020-04-01".
	PayloadVersion *string `json:"payloadVersion,omitempty" tf:"payload_version,omitempty"`

	// The ARN of the target function.
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`
}

type PreProvisioningHookObservation struct {

	// The version of the payload that was sent to the target function. The only valid (and the default) payload version is "2020-04-01".
	PayloadVersion *string `json:"payloadVersion,omitempty" tf:"payload_version,omitempty"`

	// The ARN of the target function.
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`
}

type PreProvisioningHookParameters struct {

	// The version of the payload that was sent to the target function. The only valid (and the default) payload version is "2020-04-01".
	// +kubebuilder:validation:Optional
	PayloadVersion *string `json:"payloadVersion,omitempty" tf:"payload_version,omitempty"`

	// The ARN of the target function.
	// +kubebuilder:validation:Optional
	TargetArn *string `json:"targetArn" tf:"target_arn,omitempty"`
}

type ProvisioningTemplateInitParameters struct {

	// The description of the fleet provisioning template.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// True to enable the fleet provisioning template, otherwise false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Creates a pre-provisioning hook template. Details below.
	PreProvisioningHook []PreProvisioningHookInitParameters `json:"preProvisioningHook,omitempty" tf:"pre_provisioning_hook,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The JSON formatted contents of the fleet provisioning template.
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`
}

type ProvisioningTemplateObservation struct {

	// The ARN that identifies the provisioning template.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The default version of the fleet provisioning template.
	DefaultVersionID *int64 `json:"defaultVersionId,omitempty" tf:"default_version_id,omitempty"`

	// The description of the fleet provisioning template.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// True to enable the fleet provisioning template, otherwise false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Creates a pre-provisioning hook template. Details below.
	PreProvisioningHook []PreProvisioningHookObservation `json:"preProvisioningHook,omitempty" tf:"pre_provisioning_hook,omitempty"`

	// The role ARN for the role associated with the fleet provisioning template. This IoT role grants permission to provision a device.
	ProvisioningRoleArn *string `json:"provisioningRoleArn,omitempty" tf:"provisioning_role_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The JSON formatted contents of the fleet provisioning template.
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`
}

type ProvisioningTemplateParameters struct {

	// The description of the fleet provisioning template.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// True to enable the fleet provisioning template, otherwise false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Creates a pre-provisioning hook template. Details below.
	// +kubebuilder:validation:Optional
	PreProvisioningHook []PreProvisioningHookParameters `json:"preProvisioningHook,omitempty" tf:"pre_provisioning_hook,omitempty"`

	// The role ARN for the role associated with the fleet provisioning template. This IoT role grants permission to provision a device.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ProvisioningRoleArn *string `json:"provisioningRoleArn,omitempty" tf:"provisioning_role_arn,omitempty"`

	// Reference to a Role in iam to populate provisioningRoleArn.
	// +kubebuilder:validation:Optional
	ProvisioningRoleArnRef *v1.Reference `json:"provisioningRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate provisioningRoleArn.
	// +kubebuilder:validation:Optional
	ProvisioningRoleArnSelector *v1.Selector `json:"provisioningRoleArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The JSON formatted contents of the fleet provisioning template.
	// +kubebuilder:validation:Optional
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`
}

// ProvisioningTemplateSpec defines the desired state of ProvisioningTemplate
type ProvisioningTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProvisioningTemplateParameters `json:"forProvider"`
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
	InitProvider ProvisioningTemplateInitParameters `json:"initProvider,omitempty"`
}

// ProvisioningTemplateStatus defines the observed state of ProvisioningTemplate.
type ProvisioningTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProvisioningTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProvisioningTemplate is the Schema for the ProvisioningTemplates API. Manages an IoT fleet provisioning template.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ProvisioningTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.templateBody) || has(self.initProvider.templateBody)",message="templateBody is a required parameter"
	Spec   ProvisioningTemplateSpec   `json:"spec"`
	Status ProvisioningTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProvisioningTemplateList contains a list of ProvisioningTemplates
type ProvisioningTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProvisioningTemplate `json:"items"`
}

// Repository type metadata.
var (
	ProvisioningTemplate_Kind             = "ProvisioningTemplate"
	ProvisioningTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProvisioningTemplate_Kind}.String()
	ProvisioningTemplate_KindAPIVersion   = ProvisioningTemplate_Kind + "." + CRDGroupVersion.String()
	ProvisioningTemplate_GroupVersionKind = CRDGroupVersion.WithKind(ProvisioningTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&ProvisioningTemplate{}, &ProvisioningTemplateList{})
}
