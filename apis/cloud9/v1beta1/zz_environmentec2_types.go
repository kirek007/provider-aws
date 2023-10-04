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

type EnvironmentEC2InitParameters struct {

	// The number of minutes until the running instance is shut down after the environment has last been used.
	AutomaticStopTimeMinutes *int64 `json:"automaticStopTimeMinutes,omitempty" tf:"automatic_stop_time_minutes,omitempty"`

	// The connection type used for connecting to an Amazon EC2 environment. Valid values are CONNECT_SSH and CONNECT_SSM. For more information please refer AWS documentation for Cloud9.
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// The description of the environment.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance. Valid values are
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The type of instance to connect to the environment, e.g., t2.micro.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The name of the environment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	OwnerArn *string `json:"ownerArn,omitempty" tf:"owner_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EnvironmentEC2Observation struct {

	// The ARN of the environment.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The number of minutes until the running instance is shut down after the environment has last been used.
	AutomaticStopTimeMinutes *int64 `json:"automaticStopTimeMinutes,omitempty" tf:"automatic_stop_time_minutes,omitempty"`

	// The connection type used for connecting to an Amazon EC2 environment. Valid values are CONNECT_SSH and CONNECT_SSM. For more information please refer AWS documentation for Cloud9.
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// The description of the environment.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the environment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance. Valid values are
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The type of instance to connect to the environment, e.g., t2.micro.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The name of the environment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	OwnerArn *string `json:"ownerArn,omitempty" tf:"owner_arn,omitempty"`

	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The type of the environment (e.g., ssh or ec2)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EnvironmentEC2Parameters struct {

	// The number of minutes until the running instance is shut down after the environment has last been used.
	// +kubebuilder:validation:Optional
	AutomaticStopTimeMinutes *int64 `json:"automaticStopTimeMinutes,omitempty" tf:"automatic_stop_time_minutes,omitempty"`

	// The connection type used for connecting to an Amazon EC2 environment. Valid values are CONNECT_SSH and CONNECT_SSM. For more information please refer AWS documentation for Cloud9.
	// +kubebuilder:validation:Optional
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// The description of the environment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance. Valid values are
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The type of instance to connect to the environment, e.g., t2.micro.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The name of the environment.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	// +kubebuilder:validation:Optional
	OwnerArn *string `json:"ownerArn,omitempty" tf:"owner_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EnvironmentEC2Spec defines the desired state of EnvironmentEC2
type EnvironmentEC2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentEC2Parameters `json:"forProvider"`
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
	InitProvider EnvironmentEC2InitParameters `json:"initProvider,omitempty"`
}

// EnvironmentEC2Status defines the observed state of EnvironmentEC2.
type EnvironmentEC2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentEC2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentEC2 is the Schema for the EnvironmentEC2s API. Provides a Cloud9 EC2 Development Environment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EnvironmentEC2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceType) || has(self.initProvider.instanceType)",message="instanceType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   EnvironmentEC2Spec   `json:"spec"`
	Status EnvironmentEC2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentEC2List contains a list of EnvironmentEC2s
type EnvironmentEC2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvironmentEC2 `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentEC2_Kind             = "EnvironmentEC2"
	EnvironmentEC2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EnvironmentEC2_Kind}.String()
	EnvironmentEC2_KindAPIVersion   = EnvironmentEC2_Kind + "." + CRDGroupVersion.String()
	EnvironmentEC2_GroupVersionKind = CRDGroupVersion.WithKind(EnvironmentEC2_Kind)
)

func init() {
	SchemeBuilder.Register(&EnvironmentEC2{}, &EnvironmentEC2List{})
}
