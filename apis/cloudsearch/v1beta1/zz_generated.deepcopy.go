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
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Domain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainInitParameters) DeepCopyInto(out *DomainInitParameters) {
	*out = *in
	if in.EndpointOptions != nil {
		in, out := &in.EndpointOptions, &out.EndpointOptions
		*out = make([]EndpointOptionsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IndexField != nil {
		in, out := &in.IndexField, &out.IndexField
		*out = make([]IndexFieldInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MultiAz != nil {
		in, out := &in.MultiAz, &out.MultiAz
		*out = new(bool)
		**out = **in
	}
	if in.ScalingParameters != nil {
		in, out := &in.ScalingParameters, &out.ScalingParameters
		*out = make([]ScalingParametersInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainInitParameters.
func (in *DomainInitParameters) DeepCopy() *DomainInitParameters {
	if in == nil {
		return nil
	}
	out := new(DomainInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainList) DeepCopyInto(out *DomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Domain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainList.
func (in *DomainList) DeepCopy() *DomainList {
	if in == nil {
		return nil
	}
	out := new(DomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainObservation) DeepCopyInto(out *DomainObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.DocumentServiceEndpoint != nil {
		in, out := &in.DocumentServiceEndpoint, &out.DocumentServiceEndpoint
		*out = new(string)
		**out = **in
	}
	if in.DomainID != nil {
		in, out := &in.DomainID, &out.DomainID
		*out = new(string)
		**out = **in
	}
	if in.EndpointOptions != nil {
		in, out := &in.EndpointOptions, &out.EndpointOptions
		*out = make([]EndpointOptionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IndexField != nil {
		in, out := &in.IndexField, &out.IndexField
		*out = make([]IndexFieldObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MultiAz != nil {
		in, out := &in.MultiAz, &out.MultiAz
		*out = new(bool)
		**out = **in
	}
	if in.ScalingParameters != nil {
		in, out := &in.ScalingParameters, &out.ScalingParameters
		*out = make([]ScalingParametersObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SearchServiceEndpoint != nil {
		in, out := &in.SearchServiceEndpoint, &out.SearchServiceEndpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainObservation.
func (in *DomainObservation) DeepCopy() *DomainObservation {
	if in == nil {
		return nil
	}
	out := new(DomainObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainParameters) DeepCopyInto(out *DomainParameters) {
	*out = *in
	if in.EndpointOptions != nil {
		in, out := &in.EndpointOptions, &out.EndpointOptions
		*out = make([]EndpointOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IndexField != nil {
		in, out := &in.IndexField, &out.IndexField
		*out = make([]IndexFieldParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MultiAz != nil {
		in, out := &in.MultiAz, &out.MultiAz
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ScalingParameters != nil {
		in, out := &in.ScalingParameters, &out.ScalingParameters
		*out = make([]ScalingParametersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainParameters.
func (in *DomainParameters) DeepCopy() *DomainParameters {
	if in == nil {
		return nil
	}
	out := new(DomainParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainServiceAccessPolicy) DeepCopyInto(out *DomainServiceAccessPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainServiceAccessPolicy.
func (in *DomainServiceAccessPolicy) DeepCopy() *DomainServiceAccessPolicy {
	if in == nil {
		return nil
	}
	out := new(DomainServiceAccessPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainServiceAccessPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainServiceAccessPolicyInitParameters) DeepCopyInto(out *DomainServiceAccessPolicyInitParameters) {
	*out = *in
	if in.AccessPolicy != nil {
		in, out := &in.AccessPolicy, &out.AccessPolicy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainServiceAccessPolicyInitParameters.
func (in *DomainServiceAccessPolicyInitParameters) DeepCopy() *DomainServiceAccessPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(DomainServiceAccessPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainServiceAccessPolicyList) DeepCopyInto(out *DomainServiceAccessPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DomainServiceAccessPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainServiceAccessPolicyList.
func (in *DomainServiceAccessPolicyList) DeepCopy() *DomainServiceAccessPolicyList {
	if in == nil {
		return nil
	}
	out := new(DomainServiceAccessPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainServiceAccessPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainServiceAccessPolicyObservation) DeepCopyInto(out *DomainServiceAccessPolicyObservation) {
	*out = *in
	if in.AccessPolicy != nil {
		in, out := &in.AccessPolicy, &out.AccessPolicy
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainServiceAccessPolicyObservation.
func (in *DomainServiceAccessPolicyObservation) DeepCopy() *DomainServiceAccessPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(DomainServiceAccessPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainServiceAccessPolicyParameters) DeepCopyInto(out *DomainServiceAccessPolicyParameters) {
	*out = *in
	if in.AccessPolicy != nil {
		in, out := &in.AccessPolicy, &out.AccessPolicy
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.DomainNameRef != nil {
		in, out := &in.DomainNameRef, &out.DomainNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DomainNameSelector != nil {
		in, out := &in.DomainNameSelector, &out.DomainNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainServiceAccessPolicyParameters.
func (in *DomainServiceAccessPolicyParameters) DeepCopy() *DomainServiceAccessPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(DomainServiceAccessPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainServiceAccessPolicySpec) DeepCopyInto(out *DomainServiceAccessPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainServiceAccessPolicySpec.
func (in *DomainServiceAccessPolicySpec) DeepCopy() *DomainServiceAccessPolicySpec {
	if in == nil {
		return nil
	}
	out := new(DomainServiceAccessPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainServiceAccessPolicyStatus) DeepCopyInto(out *DomainServiceAccessPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainServiceAccessPolicyStatus.
func (in *DomainServiceAccessPolicyStatus) DeepCopy() *DomainServiceAccessPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(DomainServiceAccessPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpec) DeepCopyInto(out *DomainSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpec.
func (in *DomainSpec) DeepCopy() *DomainSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainStatus) DeepCopyInto(out *DomainStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainStatus.
func (in *DomainStatus) DeepCopy() *DomainStatus {
	if in == nil {
		return nil
	}
	out := new(DomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointOptionsInitParameters) DeepCopyInto(out *EndpointOptionsInitParameters) {
	*out = *in
	if in.EnforceHTTPS != nil {
		in, out := &in.EnforceHTTPS, &out.EnforceHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.TLSSecurityPolicy != nil {
		in, out := &in.TLSSecurityPolicy, &out.TLSSecurityPolicy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointOptionsInitParameters.
func (in *EndpointOptionsInitParameters) DeepCopy() *EndpointOptionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(EndpointOptionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointOptionsObservation) DeepCopyInto(out *EndpointOptionsObservation) {
	*out = *in
	if in.EnforceHTTPS != nil {
		in, out := &in.EnforceHTTPS, &out.EnforceHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.TLSSecurityPolicy != nil {
		in, out := &in.TLSSecurityPolicy, &out.TLSSecurityPolicy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointOptionsObservation.
func (in *EndpointOptionsObservation) DeepCopy() *EndpointOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(EndpointOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointOptionsParameters) DeepCopyInto(out *EndpointOptionsParameters) {
	*out = *in
	if in.EnforceHTTPS != nil {
		in, out := &in.EnforceHTTPS, &out.EnforceHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.TLSSecurityPolicy != nil {
		in, out := &in.TLSSecurityPolicy, &out.TLSSecurityPolicy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointOptionsParameters.
func (in *EndpointOptionsParameters) DeepCopy() *EndpointOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(EndpointOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexFieldInitParameters) DeepCopyInto(out *IndexFieldInitParameters) {
	*out = *in
	if in.AnalysisScheme != nil {
		in, out := &in.AnalysisScheme, &out.AnalysisScheme
		*out = new(string)
		**out = **in
	}
	if in.DefaultValue != nil {
		in, out := &in.DefaultValue, &out.DefaultValue
		*out = new(string)
		**out = **in
	}
	if in.Facet != nil {
		in, out := &in.Facet, &out.Facet
		*out = new(bool)
		**out = **in
	}
	if in.Highlight != nil {
		in, out := &in.Highlight, &out.Highlight
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Return != nil {
		in, out := &in.Return, &out.Return
		*out = new(bool)
		**out = **in
	}
	if in.Search != nil {
		in, out := &in.Search, &out.Search
		*out = new(bool)
		**out = **in
	}
	if in.Sort != nil {
		in, out := &in.Sort, &out.Sort
		*out = new(bool)
		**out = **in
	}
	if in.SourceFields != nil {
		in, out := &in.SourceFields, &out.SourceFields
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexFieldInitParameters.
func (in *IndexFieldInitParameters) DeepCopy() *IndexFieldInitParameters {
	if in == nil {
		return nil
	}
	out := new(IndexFieldInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexFieldObservation) DeepCopyInto(out *IndexFieldObservation) {
	*out = *in
	if in.AnalysisScheme != nil {
		in, out := &in.AnalysisScheme, &out.AnalysisScheme
		*out = new(string)
		**out = **in
	}
	if in.DefaultValue != nil {
		in, out := &in.DefaultValue, &out.DefaultValue
		*out = new(string)
		**out = **in
	}
	if in.Facet != nil {
		in, out := &in.Facet, &out.Facet
		*out = new(bool)
		**out = **in
	}
	if in.Highlight != nil {
		in, out := &in.Highlight, &out.Highlight
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Return != nil {
		in, out := &in.Return, &out.Return
		*out = new(bool)
		**out = **in
	}
	if in.Search != nil {
		in, out := &in.Search, &out.Search
		*out = new(bool)
		**out = **in
	}
	if in.Sort != nil {
		in, out := &in.Sort, &out.Sort
		*out = new(bool)
		**out = **in
	}
	if in.SourceFields != nil {
		in, out := &in.SourceFields, &out.SourceFields
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexFieldObservation.
func (in *IndexFieldObservation) DeepCopy() *IndexFieldObservation {
	if in == nil {
		return nil
	}
	out := new(IndexFieldObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexFieldParameters) DeepCopyInto(out *IndexFieldParameters) {
	*out = *in
	if in.AnalysisScheme != nil {
		in, out := &in.AnalysisScheme, &out.AnalysisScheme
		*out = new(string)
		**out = **in
	}
	if in.DefaultValue != nil {
		in, out := &in.DefaultValue, &out.DefaultValue
		*out = new(string)
		**out = **in
	}
	if in.Facet != nil {
		in, out := &in.Facet, &out.Facet
		*out = new(bool)
		**out = **in
	}
	if in.Highlight != nil {
		in, out := &in.Highlight, &out.Highlight
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Return != nil {
		in, out := &in.Return, &out.Return
		*out = new(bool)
		**out = **in
	}
	if in.Search != nil {
		in, out := &in.Search, &out.Search
		*out = new(bool)
		**out = **in
	}
	if in.Sort != nil {
		in, out := &in.Sort, &out.Sort
		*out = new(bool)
		**out = **in
	}
	if in.SourceFields != nil {
		in, out := &in.SourceFields, &out.SourceFields
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexFieldParameters.
func (in *IndexFieldParameters) DeepCopy() *IndexFieldParameters {
	if in == nil {
		return nil
	}
	out := new(IndexFieldParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingParametersInitParameters) DeepCopyInto(out *ScalingParametersInitParameters) {
	*out = *in
	if in.DesiredInstanceType != nil {
		in, out := &in.DesiredInstanceType, &out.DesiredInstanceType
		*out = new(string)
		**out = **in
	}
	if in.DesiredPartitionCount != nil {
		in, out := &in.DesiredPartitionCount, &out.DesiredPartitionCount
		*out = new(float64)
		**out = **in
	}
	if in.DesiredReplicationCount != nil {
		in, out := &in.DesiredReplicationCount, &out.DesiredReplicationCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingParametersInitParameters.
func (in *ScalingParametersInitParameters) DeepCopy() *ScalingParametersInitParameters {
	if in == nil {
		return nil
	}
	out := new(ScalingParametersInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingParametersObservation) DeepCopyInto(out *ScalingParametersObservation) {
	*out = *in
	if in.DesiredInstanceType != nil {
		in, out := &in.DesiredInstanceType, &out.DesiredInstanceType
		*out = new(string)
		**out = **in
	}
	if in.DesiredPartitionCount != nil {
		in, out := &in.DesiredPartitionCount, &out.DesiredPartitionCount
		*out = new(float64)
		**out = **in
	}
	if in.DesiredReplicationCount != nil {
		in, out := &in.DesiredReplicationCount, &out.DesiredReplicationCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingParametersObservation.
func (in *ScalingParametersObservation) DeepCopy() *ScalingParametersObservation {
	if in == nil {
		return nil
	}
	out := new(ScalingParametersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingParametersParameters) DeepCopyInto(out *ScalingParametersParameters) {
	*out = *in
	if in.DesiredInstanceType != nil {
		in, out := &in.DesiredInstanceType, &out.DesiredInstanceType
		*out = new(string)
		**out = **in
	}
	if in.DesiredPartitionCount != nil {
		in, out := &in.DesiredPartitionCount, &out.DesiredPartitionCount
		*out = new(float64)
		**out = **in
	}
	if in.DesiredReplicationCount != nil {
		in, out := &in.DesiredReplicationCount, &out.DesiredReplicationCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingParametersParameters.
func (in *ScalingParametersParameters) DeepCopy() *ScalingParametersParameters {
	if in == nil {
		return nil
	}
	out := new(ScalingParametersParameters)
	in.DeepCopyInto(out)
	return out
}
