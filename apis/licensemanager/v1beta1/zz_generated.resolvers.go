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
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Association) ResolveReferences( // ResolveReferences of this Association.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("licensemanager.aws.upbound.io",

			"v1beta1", "LicenseConfiguration", "LicenseConfigurationList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LicenseConfigurationArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.LicenseConfigurationArnRef,
			Selector:     mg.Spec.ForProvider.LicenseConfigurationArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LicenseConfigurationArn")
	}
	mg.Spec.ForProvider.LicenseConfigurationArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LicenseConfigurationArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io",

			"v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ResourceArnRef,
			Selector:     mg.Spec.ForProvider.ResourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceArn")
	}
	mg.Spec.ForProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("licensemanager.aws.upbound.io",

			"v1beta1", "LicenseConfiguration", "LicenseConfigurationList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LicenseConfigurationArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.LicenseConfigurationArnRef,
			Selector:     mg.Spec.InitProvider.LicenseConfigurationArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LicenseConfigurationArn")
	}
	mg.Spec.InitProvider.LicenseConfigurationArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LicenseConfigurationArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io",

			"v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.ResourceArnRef,
			Selector:     mg.Spec.InitProvider.ResourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceArn")
	}
	mg.Spec.InitProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceArnRef = rsp.ResolvedReference

	return nil
}
