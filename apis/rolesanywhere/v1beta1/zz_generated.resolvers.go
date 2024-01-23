/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Profile.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Profile) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.RoleArns),
			Extract:       common.ARNExtractor(),
			References:    mg.Spec.ForProvider.RoleArnsRefs,
			Selector:      mg.Spec.ForProvider.RoleArnsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArns")
	}
	mg.Spec.ForProvider.RoleArns = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.RoleArnsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.RoleArns),
			Extract:       common.ARNExtractor(),
			References:    mg.Spec.InitProvider.RoleArnsRefs,
			Selector:      mg.Spec.InitProvider.RoleArnsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArns")
	}
	mg.Spec.InitProvider.RoleArns = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.RoleArnsRefs = mrsp.ResolvedReferences

	return nil
}
