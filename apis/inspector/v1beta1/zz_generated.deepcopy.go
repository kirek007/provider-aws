//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTarget) DeepCopyInto(out *AssessmentTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTarget.
func (in *AssessmentTarget) DeepCopy() *AssessmentTarget {
	if in == nil {
		return nil
	}
	out := new(AssessmentTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssessmentTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTargetInitParameters) DeepCopyInto(out *AssessmentTargetInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTargetInitParameters.
func (in *AssessmentTargetInitParameters) DeepCopy() *AssessmentTargetInitParameters {
	if in == nil {
		return nil
	}
	out := new(AssessmentTargetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTargetList) DeepCopyInto(out *AssessmentTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AssessmentTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTargetList.
func (in *AssessmentTargetList) DeepCopy() *AssessmentTargetList {
	if in == nil {
		return nil
	}
	out := new(AssessmentTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssessmentTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTargetObservation) DeepCopyInto(out *AssessmentTargetObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupArn != nil {
		in, out := &in.ResourceGroupArn, &out.ResourceGroupArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTargetObservation.
func (in *AssessmentTargetObservation) DeepCopy() *AssessmentTargetObservation {
	if in == nil {
		return nil
	}
	out := new(AssessmentTargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTargetParameters) DeepCopyInto(out *AssessmentTargetParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupArn != nil {
		in, out := &in.ResourceGroupArn, &out.ResourceGroupArn
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupArnRef != nil {
		in, out := &in.ResourceGroupArnRef, &out.ResourceGroupArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupArnSelector != nil {
		in, out := &in.ResourceGroupArnSelector, &out.ResourceGroupArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTargetParameters.
func (in *AssessmentTargetParameters) DeepCopy() *AssessmentTargetParameters {
	if in == nil {
		return nil
	}
	out := new(AssessmentTargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTargetSpec) DeepCopyInto(out *AssessmentTargetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTargetSpec.
func (in *AssessmentTargetSpec) DeepCopy() *AssessmentTargetSpec {
	if in == nil {
		return nil
	}
	out := new(AssessmentTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTargetStatus) DeepCopyInto(out *AssessmentTargetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTargetStatus.
func (in *AssessmentTargetStatus) DeepCopy() *AssessmentTargetStatus {
	if in == nil {
		return nil
	}
	out := new(AssessmentTargetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTemplate) DeepCopyInto(out *AssessmentTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTemplate.
func (in *AssessmentTemplate) DeepCopy() *AssessmentTemplate {
	if in == nil {
		return nil
	}
	out := new(AssessmentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssessmentTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTemplateInitParameters) DeepCopyInto(out *AssessmentTemplateInitParameters) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(float64)
		**out = **in
	}
	if in.EventSubscription != nil {
		in, out := &in.EventSubscription, &out.EventSubscription
		*out = make([]EventSubscriptionInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RulesPackageArns != nil {
		in, out := &in.RulesPackageArns, &out.RulesPackageArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTemplateInitParameters.
func (in *AssessmentTemplateInitParameters) DeepCopy() *AssessmentTemplateInitParameters {
	if in == nil {
		return nil
	}
	out := new(AssessmentTemplateInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTemplateList) DeepCopyInto(out *AssessmentTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AssessmentTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTemplateList.
func (in *AssessmentTemplateList) DeepCopy() *AssessmentTemplateList {
	if in == nil {
		return nil
	}
	out := new(AssessmentTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssessmentTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTemplateObservation) DeepCopyInto(out *AssessmentTemplateObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(float64)
		**out = **in
	}
	if in.EventSubscription != nil {
		in, out := &in.EventSubscription, &out.EventSubscription
		*out = make([]EventSubscriptionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RulesPackageArns != nil {
		in, out := &in.RulesPackageArns, &out.RulesPackageArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TargetArn != nil {
		in, out := &in.TargetArn, &out.TargetArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTemplateObservation.
func (in *AssessmentTemplateObservation) DeepCopy() *AssessmentTemplateObservation {
	if in == nil {
		return nil
	}
	out := new(AssessmentTemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTemplateParameters) DeepCopyInto(out *AssessmentTemplateParameters) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(float64)
		**out = **in
	}
	if in.EventSubscription != nil {
		in, out := &in.EventSubscription, &out.EventSubscription
		*out = make([]EventSubscriptionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RulesPackageArns != nil {
		in, out := &in.RulesPackageArns, &out.RulesPackageArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TargetArn != nil {
		in, out := &in.TargetArn, &out.TargetArn
		*out = new(string)
		**out = **in
	}
	if in.TargetArnRef != nil {
		in, out := &in.TargetArnRef, &out.TargetArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetArnSelector != nil {
		in, out := &in.TargetArnSelector, &out.TargetArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTemplateParameters.
func (in *AssessmentTemplateParameters) DeepCopy() *AssessmentTemplateParameters {
	if in == nil {
		return nil
	}
	out := new(AssessmentTemplateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTemplateSpec) DeepCopyInto(out *AssessmentTemplateSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTemplateSpec.
func (in *AssessmentTemplateSpec) DeepCopy() *AssessmentTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(AssessmentTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentTemplateStatus) DeepCopyInto(out *AssessmentTemplateStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentTemplateStatus.
func (in *AssessmentTemplateStatus) DeepCopy() *AssessmentTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(AssessmentTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionInitParameters) DeepCopyInto(out *EventSubscriptionInitParameters) {
	*out = *in
	if in.Event != nil {
		in, out := &in.Event, &out.Event
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionInitParameters.
func (in *EventSubscriptionInitParameters) DeepCopy() *EventSubscriptionInitParameters {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionObservation) DeepCopyInto(out *EventSubscriptionObservation) {
	*out = *in
	if in.Event != nil {
		in, out := &in.Event, &out.Event
		*out = new(string)
		**out = **in
	}
	if in.TopicArn != nil {
		in, out := &in.TopicArn, &out.TopicArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionObservation.
func (in *EventSubscriptionObservation) DeepCopy() *EventSubscriptionObservation {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionParameters) DeepCopyInto(out *EventSubscriptionParameters) {
	*out = *in
	if in.Event != nil {
		in, out := &in.Event, &out.Event
		*out = new(string)
		**out = **in
	}
	if in.TopicArn != nil {
		in, out := &in.TopicArn, &out.TopicArn
		*out = new(string)
		**out = **in
	}
	if in.TopicArnRef != nil {
		in, out := &in.TopicArnRef, &out.TopicArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TopicArnSelector != nil {
		in, out := &in.TopicArnSelector, &out.TopicArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionParameters.
func (in *EventSubscriptionParameters) DeepCopy() *EventSubscriptionParameters {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroup) DeepCopyInto(out *ResourceGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroup.
func (in *ResourceGroup) DeepCopy() *ResourceGroup {
	if in == nil {
		return nil
	}
	out := new(ResourceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupInitParameters) DeepCopyInto(out *ResourceGroupInitParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupInitParameters.
func (in *ResourceGroupInitParameters) DeepCopy() *ResourceGroupInitParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupList) DeepCopyInto(out *ResourceGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupList.
func (in *ResourceGroupList) DeepCopy() *ResourceGroupList {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupObservation) DeepCopyInto(out *ResourceGroupObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupObservation.
func (in *ResourceGroupObservation) DeepCopy() *ResourceGroupObservation {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupParameters) DeepCopyInto(out *ResourceGroupParameters) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupParameters.
func (in *ResourceGroupParameters) DeepCopy() *ResourceGroupParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupSpec) DeepCopyInto(out *ResourceGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupSpec.
func (in *ResourceGroupSpec) DeepCopy() *ResourceGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupStatus) DeepCopyInto(out *ResourceGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupStatus.
func (in *ResourceGroupStatus) DeepCopy() *ResourceGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupStatus)
	in.DeepCopyInto(out)
	return out
}
