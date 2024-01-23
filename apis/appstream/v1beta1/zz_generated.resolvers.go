/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Fleet.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Fleet) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IAMRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.IAMRoleArnRef,
			Selector:     mg.Spec.ForProvider.IAMRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRoleArn")
	}
	mg.Spec.ForProvider.IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IAMRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCConfig[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.VPCConfig[i3].SubnetIDRefs,
				Selector:      mg.Spec.ForProvider.VPCConfig[i3].SubnetIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfig[i3].SubnetIds")
		}
		mg.Spec.ForProvider.VPCConfig[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCConfig[i3].SubnetIDRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IAMRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.IAMRoleArnRef,
			Selector:     mg.Spec.InitProvider.IAMRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IAMRoleArn")
	}
	mg.Spec.InitProvider.IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IAMRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPCConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCConfig[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.VPCConfig[i3].SubnetIDRefs,
				Selector:      mg.Spec.InitProvider.VPCConfig[i3].SubnetIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCConfig[i3].SubnetIds")
		}
		mg.Spec.InitProvider.VPCConfig[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.VPCConfig[i3].SubnetIDRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this FleetStackAssociation.
func (mg *FleetStackAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appstream.aws.upbound.io", "v1beta1", "Fleet", "FleetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FleetName),
			Extract:      resource.ExtractParamPath("name", false),
			Reference:    mg.Spec.ForProvider.FleetNameRef,
			Selector:     mg.Spec.ForProvider.FleetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FleetName")
	}
	mg.Spec.ForProvider.FleetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FleetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appstream.aws.upbound.io", "v1beta1", "Stack", "StackList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StackName),
			Extract:      resource.ExtractParamPath("name", false),
			Reference:    mg.Spec.ForProvider.StackNameRef,
			Selector:     mg.Spec.ForProvider.StackNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StackName")
	}
	mg.Spec.ForProvider.StackName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StackNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ImageBuilder.
func (mg *ImageBuilder) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IAMRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.IAMRoleArnRef,
			Selector:     mg.Spec.ForProvider.IAMRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRoleArn")
	}
	mg.Spec.ForProvider.IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IAMRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCConfig[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.VPCConfig[i3].SubnetIDRefs,
				Selector:      mg.Spec.ForProvider.VPCConfig[i3].SubnetIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfig[i3].SubnetIds")
		}
		mg.Spec.ForProvider.VPCConfig[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCConfig[i3].SubnetIDRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IAMRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.IAMRoleArnRef,
			Selector:     mg.Spec.InitProvider.IAMRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IAMRoleArn")
	}
	mg.Spec.InitProvider.IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IAMRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPCConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCConfig[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.VPCConfig[i3].SubnetIDRefs,
				Selector:      mg.Spec.InitProvider.VPCConfig[i3].SubnetIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCConfig[i3].SubnetIds")
		}
		mg.Spec.InitProvider.VPCConfig[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.VPCConfig[i3].SubnetIDRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this UserStackAssociation.
func (mg *UserStackAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appstream.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthenticationType),
			Extract:      resource.ExtractParamPath("authentication_type", false),
			Reference:    mg.Spec.ForProvider.AuthenticationTypeRef,
			Selector:     mg.Spec.ForProvider.AuthenticationTypeSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthenticationType")
	}
	mg.Spec.ForProvider.AuthenticationType = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthenticationTypeRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appstream.aws.upbound.io", "v1beta1", "Stack", "StackList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StackName),
			Extract:      resource.ExtractParamPath("name", false),
			Reference:    mg.Spec.ForProvider.StackNameRef,
			Selector:     mg.Spec.ForProvider.StackNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StackName")
	}
	mg.Spec.ForProvider.StackName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StackNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appstream.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.UserNameRef,
			Selector:     mg.Spec.ForProvider.UserNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserName")
	}
	mg.Spec.ForProvider.UserName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserNameRef = rsp.ResolvedReference

	return nil
}
