/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *GroupMembership) ResolveReferences( // ResolveReferences of this GroupMembership.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("identitystore.aws.upbound.io",

			"v1beta1", "Group", "GroupList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupID),
			Extract:      resource.ExtractParamPath("group_id", true),
			Reference:    mg.Spec.ForProvider.GroupIDRef,
			Selector:     mg.Spec.ForProvider.GroupIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupID")
	}
	mg.Spec.ForProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("identitystore.aws.upbound.io",

			"v1beta1", "User", "UserList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MemberID),
			Extract:      resource.ExtractParamPath("user_id", true),
			Reference:    mg.Spec.ForProvider.MemberIDRef,
			Selector:     mg.Spec.ForProvider.MemberIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MemberID")
	}
	mg.Spec.ForProvider.MemberID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MemberIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("identitystore.aws.upbound.io",

			"v1beta1", "Group", "GroupList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.GroupID),
			Extract:      resource.ExtractParamPath("group_id", true),
			Reference:    mg.Spec.InitProvider.GroupIDRef,
			Selector:     mg.Spec.InitProvider.GroupIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GroupID")
	}
	mg.Spec.InitProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GroupIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("identitystore.aws.upbound.io",

			"v1beta1", "User", "UserList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MemberID),
			Extract:      resource.ExtractParamPath("user_id", true),
			Reference:    mg.Spec.InitProvider.MemberIDRef,
			Selector:     mg.Spec.InitProvider.MemberIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MemberID")
	}
	mg.Spec.InitProvider.MemberID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MemberIDRef = rsp.ResolvedReference

	return nil
}
