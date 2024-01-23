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
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Domain.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Domain) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LogPublishingOptions); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudwatchlogs.aws.upbound.io",

				"v1beta1", "Group",

				"GroupList",
			)
			if err !=
				nil {

				return errors.Wrap(err,

					"failed to get the reference target managed resource and its list for reference resolution",
				)
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnRef,
				Selector:     mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn")
		}
		mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.LogPublishingOptions); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudwatchlogs.aws.upbound.io",

				"v1beta1", "Group",

				"GroupList",
			)
			if err !=
				nil {

				return errors.Wrap(err,

					"failed to get the reference target managed resource and its list for reference resolution",
				)
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnRef,
				Selector:     mg.Spec.InitProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn")
		}
		mg.Spec.InitProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this DomainPolicy.
func (mg *DomainPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("opensearch.aws.upbound.io",

			"v1beta1",
			"Domain",

			"DomainList")

		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DomainName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DomainNameRef,
			Selector:     mg.Spec.ForProvider.DomainNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DomainName")
	}
	mg.Spec.ForProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("opensearch.aws.upbound.io",

			"v1beta1",
			"Domain",

			"DomainList")

		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DomainName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DomainNameRef,
			Selector:     mg.Spec.InitProvider.DomainNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DomainName")
	}
	mg.Spec.InitProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DomainNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DomainSAMLOptions.
func (mg *DomainSAMLOptions) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("opensearch.aws.upbound.io",

			"v1beta1",
			"Domain",

			"DomainList")

		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DomainName),
			Extract:      resource.ExtractParamPath("domain_name", false),
			Reference:    mg.Spec.ForProvider.DomainNameRef,
			Selector:     mg.Spec.ForProvider.DomainNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DomainName")
	}
	mg.Spec.ForProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("opensearch.aws.upbound.io",

			"v1beta1",
			"Domain",

			"DomainList")

		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DomainName),
			Extract:      resource.ExtractParamPath("domain_name", false),
			Reference:    mg.Spec.InitProvider.DomainNameRef,
			Selector:     mg.Spec.InitProvider.DomainNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DomainName")
	}
	mg.Spec.InitProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DomainNameRef = rsp.ResolvedReference

	return nil
}
