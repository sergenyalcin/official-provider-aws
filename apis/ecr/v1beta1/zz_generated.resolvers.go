/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this LifecyclePolicy.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *LifecyclePolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ecr.aws.upbound.io",

			"v1beta1", "Repository", "RepositoryList")

		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RepositoryRef,
			Selector:     mg.Spec.ForProvider.RepositorySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ecr.aws.upbound.io",

			"v1beta1", "Repository", "RepositoryList")

		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RepositoryRef,
			Selector:     mg.Spec.InitProvider.RepositorySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Repository.
func (mg *Repository) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EncryptionConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io",

				"v1beta1", "Key", "KeyList")
			if err != nil {
				return errors.
					Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKey),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKeyRef,
				Selector:     mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKeySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKey")
		}
		mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKey = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKeyRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.EncryptionConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io",

				"v1beta1", "Key", "KeyList")
			if err != nil {
				return errors.
					Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EncryptionConfiguration[i3].KMSKey),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.InitProvider.EncryptionConfiguration[i3].KMSKeyRef,
				Selector:     mg.Spec.InitProvider.EncryptionConfiguration[i3].KMSKeySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EncryptionConfiguration[i3].KMSKey")
		}
		mg.Spec.InitProvider.EncryptionConfiguration[i3].KMSKey = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EncryptionConfiguration[i3].KMSKeyRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this RepositoryPolicy.
func (mg *RepositoryPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ecr.aws.upbound.io",

			"v1beta1", "Repository", "RepositoryList")

		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RepositoryRef,
			Selector:     mg.Spec.ForProvider.RepositorySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ecr.aws.upbound.io",

			"v1beta1", "Repository", "RepositoryList")

		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RepositoryRef,
			Selector:     mg.Spec.InitProvider.RepositorySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}
