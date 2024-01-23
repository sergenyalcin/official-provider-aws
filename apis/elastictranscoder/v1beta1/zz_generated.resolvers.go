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
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Pipeline.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Pipeline) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ContentConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io",

				"v1beta1", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(
					err, "failed to get the reference target managed resource and its list for reference resolution",
				)
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContentConfig[i3].Bucket),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.ContentConfig[i3].BucketRef,
				Selector:     mg.Spec.ForProvider.ContentConfig[i3].BucketSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ContentConfig[i3].Bucket")
		}
		mg.Spec.ForProvider.ContentConfig[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ContentConfig[i3].BucketRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io",

			"v1beta1", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(
				err, "failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InputBucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InputBucketRef,
			Selector:     mg.Spec.ForProvider.InputBucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InputBucket")
	}
	mg.Spec.ForProvider.InputBucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InputBucketRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.RoleRef,
			Selector:     mg.Spec.ForProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ThumbnailConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io",

				"v1beta1", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(
					err, "failed to get the reference target managed resource and its list for reference resolution",
				)
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ThumbnailConfig[i3].Bucket),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.ThumbnailConfig[i3].BucketRef,
				Selector:     mg.Spec.ForProvider.ThumbnailConfig[i3].BucketSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ThumbnailConfig[i3].Bucket")
		}
		mg.Spec.ForProvider.ThumbnailConfig[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ThumbnailConfig[i3].BucketRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ContentConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io",

				"v1beta1", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(
					err, "failed to get the reference target managed resource and its list for reference resolution",
				)
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ContentConfig[i3].Bucket),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.ContentConfig[i3].BucketRef,
				Selector:     mg.Spec.InitProvider.ContentConfig[i3].BucketSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ContentConfig[i3].Bucket")
		}
		mg.Spec.InitProvider.ContentConfig[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ContentConfig[i3].BucketRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io",

			"v1beta1", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(
				err, "failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InputBucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InputBucketRef,
			Selector:     mg.Spec.InitProvider.InputBucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InputBucket")
	}
	mg.Spec.InitProvider.InputBucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InputBucketRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.RoleRef,
			Selector:     mg.Spec.InitProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.ThumbnailConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io",

				"v1beta1", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(
					err, "failed to get the reference target managed resource and its list for reference resolution",
				)
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ThumbnailConfig[i3].Bucket),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.ThumbnailConfig[i3].BucketRef,
				Selector:     mg.Spec.InitProvider.ThumbnailConfig[i3].BucketSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ThumbnailConfig[i3].Bucket")
		}
		mg.Spec.InitProvider.ThumbnailConfig[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ThumbnailConfig[i3].BucketRef = rsp.ResolvedReference

	}

	return nil
}
