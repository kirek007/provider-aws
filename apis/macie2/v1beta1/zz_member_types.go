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

type MemberInitParameters struct {

	// The AWS account ID for the account.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The email address for the account.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to true.
	InvitationDisableEmailNotification *bool `json:"invitationDisableEmailNotification,omitempty" tf:"invitation_disable_email_notification,omitempty"`

	// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
	InvitationMessage *string `json:"invitationMessage,omitempty" tf:"invitation_message,omitempty"`

	// Send an invitation to a member
	Invite *bool `json:"invite,omitempty" tf:"invite,omitempty"`

	// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to ENABLED. Valid values are ENABLED or PAUSED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MemberObservation struct {

	// The AWS account ID for the account.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The AWS account ID for the administrator account.
	AdministratorAccountID *string `json:"administratorAccountId,omitempty" tf:"administrator_account_id,omitempty"`

	// The Amazon Resource Name (ARN) of the account.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The email address for the account.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The unique identifier (ID) of the macie Member.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to true.
	InvitationDisableEmailNotification *bool `json:"invitationDisableEmailNotification,omitempty" tf:"invitation_disable_email_notification,omitempty"`

	// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
	InvitationMessage *string `json:"invitationMessage,omitempty" tf:"invitation_message,omitempty"`

	// Send an invitation to a member
	Invite *bool `json:"invite,omitempty" tf:"invite,omitempty"`

	// The date and time, in UTC and extended RFC 3339 format, when an Amazon Macie membership invitation was last sent to the account. This value is null if a Macie invitation hasn't been sent to the account.
	InvitedAt *string `json:"invitedAt,omitempty" tf:"invited_at,omitempty"`

	// The AWS account ID for the account.
	MasterAccountID *string `json:"masterAccountId,omitempty" tf:"master_account_id,omitempty"`

	// The current status of the relationship between the account and the administrator account.
	RelationshipStatus *string `json:"relationshipStatus,omitempty" tf:"relationship_status,omitempty"`

	// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to ENABLED. Valid values are ENABLED or PAUSED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The date and time, in UTC and extended RFC 3339 format, of the most recent change to the status of the relationship between the account and the administrator account.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type MemberParameters struct {

	// The AWS account ID for the account.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The email address for the account.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to true.
	// +kubebuilder:validation:Optional
	InvitationDisableEmailNotification *bool `json:"invitationDisableEmailNotification,omitempty" tf:"invitation_disable_email_notification,omitempty"`

	// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
	// +kubebuilder:validation:Optional
	InvitationMessage *string `json:"invitationMessage,omitempty" tf:"invitation_message,omitempty"`

	// Send an invitation to a member
	// +kubebuilder:validation:Optional
	Invite *bool `json:"invite,omitempty" tf:"invite,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to ENABLED. Valid values are ENABLED or PAUSED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// MemberSpec defines the desired state of Member
type MemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberParameters `json:"forProvider"`
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
	InitProvider MemberInitParameters `json:"initProvider,omitempty"`
}

// MemberStatus defines the observed state of Member.
type MemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Member is the Schema for the Members API. Provides a resource to manage an Amazon Macie Member.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Member struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || has(self.initProvider.accountId)",message="accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || has(self.initProvider.email)",message="email is a required parameter"
	Spec   MemberSpec   `json:"spec"`
	Status MemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberList contains a list of Members
type MemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Member `json:"items"`
}

// Repository type metadata.
var (
	Member_Kind             = "Member"
	Member_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Member_Kind}.String()
	Member_KindAPIVersion   = Member_Kind + "." + CRDGroupVersion.String()
	Member_GroupVersionKind = CRDGroupVersion.WithKind(Member_Kind)
)

func init() {
	SchemeBuilder.Register(&Member{}, &MemberList{})
}
