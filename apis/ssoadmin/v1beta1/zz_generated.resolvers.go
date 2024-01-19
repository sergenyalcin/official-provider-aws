/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this AccountAssignment.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *AccountAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ssoadmin.aws.upbound.io",

			"v1beta1", "PermissionSet", "PermissionSetList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PermissionSetArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.PermissionSetArnRef,
			Selector:     mg.Spec.ForProvider.PermissionSetArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PermissionSetArn")
	}
	mg.Spec.ForProvider.PermissionSetArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PermissionSetArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("identitystore.aws.upbound.io",

			"v1beta1", "Group", "GroupList")

		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrincipalID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.PrincipalIDFromGroupRef,
			Selector:     mg.Spec.ForProvider.PrincipalIDFromGroupSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrincipalID")
	}
	mg.Spec.ForProvider.PrincipalID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrincipalIDFromGroupRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CustomerManagedPolicyAttachment.
func (mg *CustomerManagedPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CustomerManagedPolicyReference); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

				"v1beta1", "Policy", "PolicyList")
			if err !=
				nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomerManagedPolicyReference[i3].Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.CustomerManagedPolicyReference[i3].PolicyNameRef,
				Selector:     mg.Spec.ForProvider.CustomerManagedPolicyReference[i3].PolicyNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomerManagedPolicyReference[i3].Name")
		}
		mg.Spec.ForProvider.CustomerManagedPolicyReference[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CustomerManagedPolicyReference[i3].PolicyNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("ssoadmin.aws.upbound.io",

			"v1beta1", "PermissionSet", "PermissionSetList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PermissionSetArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.PermissionSetArnRef,
			Selector:     mg.Spec.ForProvider.PermissionSetArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PermissionSetArn")
	}
	mg.Spec.ForProvider.PermissionSetArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PermissionSetArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.CustomerManagedPolicyReference); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

				"v1beta1", "Policy", "PolicyList")
			if err !=
				nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CustomerManagedPolicyReference[i3].Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.CustomerManagedPolicyReference[i3].PolicyNameRef,
				Selector:     mg.Spec.InitProvider.CustomerManagedPolicyReference[i3].PolicyNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.CustomerManagedPolicyReference[i3].Name")
		}
		mg.Spec.InitProvider.CustomerManagedPolicyReference[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.CustomerManagedPolicyReference[i3].PolicyNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ManagedPolicyAttachment.
func (mg *ManagedPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ssoadmin.aws.upbound.io",

			"v1beta1", "PermissionSet", "PermissionSetList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PermissionSetArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.PermissionSetArnRef,
			Selector:     mg.Spec.ForProvider.PermissionSetArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PermissionSetArn")
	}
	mg.Spec.ForProvider.PermissionSetArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PermissionSetArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PermissionSetInlinePolicy.
func (mg *PermissionSetInlinePolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ssoadmin.aws.upbound.io",

			"v1beta1", "PermissionSet", "PermissionSetList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PermissionSetArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.PermissionSetArnRef,
			Selector:     mg.Spec.ForProvider.PermissionSetArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PermissionSetArn")
	}
	mg.Spec.ForProvider.PermissionSetArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PermissionSetArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PermissionsBoundaryAttachment.
func (mg *PermissionsBoundaryAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ssoadmin.aws.upbound.io",

			"v1beta1", "PermissionSet", "PermissionSetList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PermissionSetArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.PermissionSetArnRef,
			Selector:     mg.Spec.ForProvider.PermissionSetArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PermissionSetArn")
	}
	mg.Spec.ForProvider.PermissionSetArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PermissionSetArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PermissionsBoundary); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

					"v1beta1", "Policy", "PolicyList")
				if err !=
					nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}

				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].Name),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].NameRef,
					Selector:     mg.Spec.ForProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].NameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].Name")
			}
			mg.Spec.ForProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].Name = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].NameRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.PermissionsBoundary); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

					"v1beta1", "Policy", "PolicyList")
				if err !=
					nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}

				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].Name),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].NameRef,
					Selector:     mg.Spec.InitProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].NameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].Name")
			}
			mg.Spec.InitProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].Name = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.PermissionsBoundary[i3].CustomerManagedPolicyReference[i4].NameRef = rsp.ResolvedReference

		}
	}

	return nil
}
