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

type AbortIncompleteMultipartUploadInitParameters struct {

	// Number of days after which Amazon S3 aborts an incomplete multipart upload.
	DaysAfterInitiation *int64 `json:"daysAfterInitiation,omitempty" tf:"days_after_initiation,omitempty"`
}

type AbortIncompleteMultipartUploadObservation struct {

	// Number of days after which Amazon S3 aborts an incomplete multipart upload.
	DaysAfterInitiation *int64 `json:"daysAfterInitiation,omitempty" tf:"days_after_initiation,omitempty"`
}

type AbortIncompleteMultipartUploadParameters struct {

	// Number of days after which Amazon S3 aborts an incomplete multipart upload.
	// +kubebuilder:validation:Optional
	DaysAfterInitiation *int64 `json:"daysAfterInitiation,omitempty" tf:"days_after_initiation,omitempty"`
}

type AndInitParameters struct {

	// Minimum object size (in bytes) to which the rule applies.
	ObjectSizeGreaterThan *int64 `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// Maximum object size (in bytes) to which the rule applies.
	ObjectSizeLessThan *int64 `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags. All of these tags must exist in the object's tag set in order for the rule to apply.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AndObservation struct {

	// Minimum object size (in bytes) to which the rule applies.
	ObjectSizeGreaterThan *int64 `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// Maximum object size (in bytes) to which the rule applies.
	ObjectSizeLessThan *int64 `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags. All of these tags must exist in the object's tag set in order for the rule to apply.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AndParameters struct {

	// Minimum object size (in bytes) to which the rule applies.
	// +kubebuilder:validation:Optional
	ObjectSizeGreaterThan *int64 `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// Maximum object size (in bytes) to which the rule applies.
	// +kubebuilder:validation:Optional
	ObjectSizeLessThan *int64 `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags. All of these tags must exist in the object's tag set in order for the rule to apply.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BucketLifecycleConfigurationInitParameters struct {

	// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// List of configuration blocks describing the rules managing the replication. See below.
	Rule []BucketLifecycleConfigurationRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketLifecycleConfigurationObservation struct {

	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// and status)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of configuration blocks describing the rules managing the replication. See below.
	Rule []BucketLifecycleConfigurationRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketLifecycleConfigurationParameters struct {

	// Name of the source S3 bucket you want Amazon S3 to monitor.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of configuration blocks describing the rules managing the replication. See below.
	// +kubebuilder:validation:Optional
	Rule []BucketLifecycleConfigurationRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketLifecycleConfigurationRuleInitParameters struct {

	// Configuration block that specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload. See below.
	AbortIncompleteMultipartUpload []AbortIncompleteMultipartUploadInitParameters `json:"abortIncompleteMultipartUpload,omitempty" tf:"abort_incomplete_multipart_upload,omitempty"`

	// Configuration block that specifies the expiration for the lifecycle of the object in the form of date, days and, whether the object has a delete marker. See below.
	Expiration []RuleExpirationInitParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Configuration block used to identify objects that a Lifecycle Rule applies to. See below. If not specified, the rule will default to using prefix.
	Filter []RuleFilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block that specifies when noncurrent object versions expire. See below.
	NoncurrentVersionExpiration []RuleNoncurrentVersionExpirationInitParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Set of configuration blocks that specify the transition rule for the lifecycle rule that describes when noncurrent objects transition to a specific storage class. See below.
	NoncurrentVersionTransition []RuleNoncurrentVersionTransitionInitParameters `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Whether the rule is currently being applied. Valid values: Enabled or Disabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Set of configuration blocks that specify when an Amazon S3 object transitions to a specified storage class. See below.
	Transition []RuleTransitionInitParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type BucketLifecycleConfigurationRuleObservation struct {

	// Configuration block that specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload. See below.
	AbortIncompleteMultipartUpload []AbortIncompleteMultipartUploadObservation `json:"abortIncompleteMultipartUpload,omitempty" tf:"abort_incomplete_multipart_upload,omitempty"`

	// Configuration block that specifies the expiration for the lifecycle of the object in the form of date, days and, whether the object has a delete marker. See below.
	Expiration []RuleExpirationObservation `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Configuration block used to identify objects that a Lifecycle Rule applies to. See below. If not specified, the rule will default to using prefix.
	Filter []RuleFilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block that specifies when noncurrent object versions expire. See below.
	NoncurrentVersionExpiration []RuleNoncurrentVersionExpirationObservation `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Set of configuration blocks that specify the transition rule for the lifecycle rule that describes when noncurrent objects transition to a specific storage class. See below.
	NoncurrentVersionTransition []RuleNoncurrentVersionTransitionObservation `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Whether the rule is currently being applied. Valid values: Enabled or Disabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Set of configuration blocks that specify when an Amazon S3 object transitions to a specified storage class. See below.
	Transition []RuleTransitionObservation `json:"transition,omitempty" tf:"transition,omitempty"`
}

type BucketLifecycleConfigurationRuleParameters struct {

	// Configuration block that specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload. See below.
	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUpload []AbortIncompleteMultipartUploadParameters `json:"abortIncompleteMultipartUpload,omitempty" tf:"abort_incomplete_multipart_upload,omitempty"`

	// Configuration block that specifies the expiration for the lifecycle of the object in the form of date, days and, whether the object has a delete marker. See below.
	// +kubebuilder:validation:Optional
	Expiration []RuleExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Configuration block used to identify objects that a Lifecycle Rule applies to. See below. If not specified, the rule will default to using prefix.
	// +kubebuilder:validation:Optional
	Filter []RuleFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// Configuration block that specifies when noncurrent object versions expire. See below.
	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration []RuleNoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Set of configuration blocks that specify the transition rule for the lifecycle rule that describes when noncurrent objects transition to a specific storage class. See below.
	// +kubebuilder:validation:Optional
	NoncurrentVersionTransition []RuleNoncurrentVersionTransitionParameters `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Whether the rule is currently being applied. Valid values: Enabled or Disabled.
	// +kubebuilder:validation:Optional
	Status *string `json:"status" tf:"status,omitempty"`

	// Set of configuration blocks that specify when an Amazon S3 object transitions to a specified storage class. See below.
	// +kubebuilder:validation:Optional
	Transition []RuleTransitionParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type RuleExpirationInitParameters struct {

	// Date objects are transitioned to the specified storage class. The date value must be in RFC3339 full-date format e.g. 2023-08-22.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Number of days after creation when objects are transitioned to the specified storage class. The value must be a positive integer. If both days and date are not specified, defaults to 0. Valid values depend on storage_class, see Transition objects using Amazon S3 Lifecycle for more details.
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions. If set to true, the delete marker will be expired; if set to false the policy takes no action.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type RuleExpirationObservation struct {

	// Date objects are transitioned to the specified storage class. The date value must be in RFC3339 full-date format e.g. 2023-08-22.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Number of days after creation when objects are transitioned to the specified storage class. The value must be a positive integer. If both days and date are not specified, defaults to 0. Valid values depend on storage_class, see Transition objects using Amazon S3 Lifecycle for more details.
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions. If set to true, the delete marker will be expired; if set to false the policy takes no action.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type RuleExpirationParameters struct {

	// Date objects are transitioned to the specified storage class. The date value must be in RFC3339 full-date format e.g. 2023-08-22.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Number of days after creation when objects are transitioned to the specified storage class. The value must be a positive integer. If both days and date are not specified, defaults to 0. Valid values depend on storage_class, see Transition objects using Amazon S3 Lifecycle for more details.
	// +kubebuilder:validation:Optional
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions. If set to true, the delete marker will be expired; if set to false the policy takes no action.
	// +kubebuilder:validation:Optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type RuleFilterInitParameters struct {

	// Configuration block used to apply a logical AND to two or more predicates. See below. The Lifecycle Rule will apply to any object matching all the predicates configured inside the and block.
	And []AndInitParameters `json:"and,omitempty" tf:"and,omitempty"`

	// Minimum object size (in bytes) to which the rule applies.
	ObjectSizeGreaterThan *string `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// Maximum object size (in bytes) to which the rule applies.
	ObjectSizeLessThan *string `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Configuration block for specifying a tag key and value. See below.
	Tag []TagInitParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type RuleFilterObservation struct {

	// Configuration block used to apply a logical AND to two or more predicates. See below. The Lifecycle Rule will apply to any object matching all the predicates configured inside the and block.
	And []AndObservation `json:"and,omitempty" tf:"and,omitempty"`

	// Minimum object size (in bytes) to which the rule applies.
	ObjectSizeGreaterThan *string `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// Maximum object size (in bytes) to which the rule applies.
	ObjectSizeLessThan *string `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Configuration block for specifying a tag key and value. See below.
	Tag []TagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type RuleFilterParameters struct {

	// Configuration block used to apply a logical AND to two or more predicates. See below. The Lifecycle Rule will apply to any object matching all the predicates configured inside the and block.
	// +kubebuilder:validation:Optional
	And []AndParameters `json:"and,omitempty" tf:"and,omitempty"`

	// Minimum object size (in bytes) to which the rule applies.
	// +kubebuilder:validation:Optional
	ObjectSizeGreaterThan *string `json:"objectSizeGreaterThan,omitempty" tf:"object_size_greater_than,omitempty"`

	// Maximum object size (in bytes) to which the rule applies.
	// +kubebuilder:validation:Optional
	ObjectSizeLessThan *string `json:"objectSizeLessThan,omitempty" tf:"object_size_less_than,omitempty"`

	// DEPRECATED Use filter instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if filter is not specified.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Configuration block for specifying a tag key and value. See below.
	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type RuleNoncurrentVersionExpirationInitParameters struct {

	// Number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// Number of days an object is noncurrent before Amazon S3 can perform the associated action.
	NoncurrentDays *int64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`
}

type RuleNoncurrentVersionExpirationObservation struct {

	// Number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// Number of days an object is noncurrent before Amazon S3 can perform the associated action.
	NoncurrentDays *int64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`
}

type RuleNoncurrentVersionExpirationParameters struct {

	// Number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	// +kubebuilder:validation:Optional
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// Number of days an object is noncurrent before Amazon S3 can perform the associated action.
	// +kubebuilder:validation:Optional
	NoncurrentDays *int64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`
}

type RuleNoncurrentVersionTransitionInitParameters struct {

	// Number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// Number of days an object is noncurrent before Amazon S3 can perform the associated action.
	NoncurrentDays *int64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`

	// Class of storage used to store the object. Valid Values: GLACIER, STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, DEEP_ARCHIVE, GLACIER_IR.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleNoncurrentVersionTransitionObservation struct {

	// Number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// Number of days an object is noncurrent before Amazon S3 can perform the associated action.
	NoncurrentDays *int64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`

	// Class of storage used to store the object. Valid Values: GLACIER, STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, DEEP_ARCHIVE, GLACIER_IR.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleNoncurrentVersionTransitionParameters struct {

	// Number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	// +kubebuilder:validation:Optional
	NewerNoncurrentVersions *string `json:"newerNoncurrentVersions,omitempty" tf:"newer_noncurrent_versions,omitempty"`

	// Number of days an object is noncurrent before Amazon S3 can perform the associated action.
	// +kubebuilder:validation:Optional
	NoncurrentDays *int64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`

	// Class of storage used to store the object. Valid Values: GLACIER, STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, DEEP_ARCHIVE, GLACIER_IR.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type RuleTransitionInitParameters struct {

	// Date objects are transitioned to the specified storage class. The date value must be in RFC3339 full-date format e.g. 2023-08-22.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Number of days after creation when objects are transitioned to the specified storage class. The value must be a positive integer. If both days and date are not specified, defaults to 0. Valid values depend on storage_class, see Transition objects using Amazon S3 Lifecycle for more details.
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// Class of storage used to store the object. Valid Values: GLACIER, STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, DEEP_ARCHIVE, GLACIER_IR.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleTransitionObservation struct {

	// Date objects are transitioned to the specified storage class. The date value must be in RFC3339 full-date format e.g. 2023-08-22.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Number of days after creation when objects are transitioned to the specified storage class. The value must be a positive integer. If both days and date are not specified, defaults to 0. Valid values depend on storage_class, see Transition objects using Amazon S3 Lifecycle for more details.
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// Class of storage used to store the object. Valid Values: GLACIER, STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, DEEP_ARCHIVE, GLACIER_IR.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleTransitionParameters struct {

	// Date objects are transitioned to the specified storage class. The date value must be in RFC3339 full-date format e.g. 2023-08-22.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Number of days after creation when objects are transitioned to the specified storage class. The value must be a positive integer. If both days and date are not specified, defaults to 0. Valid values depend on storage_class, see Transition objects using Amazon S3 Lifecycle for more details.
	// +kubebuilder:validation:Optional
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// Class of storage used to store the object. Valid Values: GLACIER, STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, DEEP_ARCHIVE, GLACIER_IR.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type TagInitParameters struct {

	// Name of the object key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Value of the tag.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagObservation struct {

	// Name of the object key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Value of the tag.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagParameters struct {

	// Name of the object key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Value of the tag.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// BucketLifecycleConfigurationSpec defines the desired state of BucketLifecycleConfiguration
type BucketLifecycleConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketLifecycleConfigurationParameters `json:"forProvider"`
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
	InitProvider BucketLifecycleConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketLifecycleConfigurationStatus defines the observed state of BucketLifecycleConfiguration.
type BucketLifecycleConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketLifecycleConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLifecycleConfiguration is the Schema for the BucketLifecycleConfigurations API. Provides a S3 bucket lifecycle configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketLifecycleConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rule) || has(self.initProvider.rule)",message="rule is a required parameter"
	Spec   BucketLifecycleConfigurationSpec   `json:"spec"`
	Status BucketLifecycleConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLifecycleConfigurationList contains a list of BucketLifecycleConfigurations
type BucketLifecycleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketLifecycleConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketLifecycleConfiguration_Kind             = "BucketLifecycleConfiguration"
	BucketLifecycleConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketLifecycleConfiguration_Kind}.String()
	BucketLifecycleConfiguration_KindAPIVersion   = BucketLifecycleConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketLifecycleConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketLifecycleConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketLifecycleConfiguration{}, &BucketLifecycleConfigurationList{})
}
