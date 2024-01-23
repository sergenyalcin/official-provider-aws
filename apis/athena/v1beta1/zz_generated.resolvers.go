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

	// ResolveReferences of this Database.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Database) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io",

			"v1beta1", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(
				err, "failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.BucketRef,
			Selector:     mg.Spec.ForProvider.BucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io",

			"v1beta1", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(
				err, "failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.BucketRef,
			Selector:     mg.Spec.InitProvider.BucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NamedQuery.
func (mg *NamedQuery) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("athena.aws.upbound.io",

			"v1beta1", "Database", "DatabaseList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Database),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DatabaseRef,
			Selector:     mg.Spec.ForProvider.DatabaseSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Database")
	}
	mg.Spec.ForProvider.Database = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("athena.aws.upbound.io",

			"v1beta1", "Workgroup", "WorkgroupList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Workgroup),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WorkgroupRef,
			Selector:     mg.Spec.ForProvider.WorkgroupSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Workgroup")
	}
	mg.Spec.ForProvider.Workgroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkgroupRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("athena.aws.upbound.io",

			"v1beta1", "Database", "DatabaseList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Database),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DatabaseRef,
			Selector:     mg.Spec.InitProvider.DatabaseSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Database")
	}
	mg.Spec.InitProvider.Database = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DatabaseRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("athena.aws.upbound.io",

			"v1beta1", "Workgroup", "WorkgroupList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Workgroup),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.WorkgroupRef,
			Selector:     mg.Spec.InitProvider.WorkgroupSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Workgroup")
	}
	mg.Spec.InitProvider.Workgroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WorkgroupRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Workgroup.
func (mg *Workgroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Configuration[i3].ResultConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io",

						"v1beta1", "Key", "KeyList")
					if err != nil {
						return errors.Wrap(err,
							"failed to get the reference target managed resource and its list for reference resolution",
						)
					}

					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArn),
						Extract:      common.ARNExtractor(),
						Reference:    mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArnRef,
						Selector:     mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArn")
				}
				mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Configuration[i3].ResultConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io",

						"v1beta1", "Key", "KeyList")
					if err != nil {
						return errors.Wrap(err,
							"failed to get the reference target managed resource and its list for reference resolution",
						)
					}

					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArn),
						Extract:      common.ARNExtractor(),
						Reference:    mg.Spec.InitProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArnRef,
						Selector:     mg.Spec.InitProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArn")
				}
				mg.Spec.InitProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.Configuration[i3].ResultConfiguration[i4].EncryptionConfiguration[i5].KMSKeyArnRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}
