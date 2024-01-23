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

	// ResolveReferences of this BudgetResourceAssociation.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *BudgetResourceAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("budgets.aws.upbound.io",

			"v1beta1",
			"Budget", "BudgetList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BudgetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.BudgetNameRef,
			Selector:     mg.Spec.ForProvider.BudgetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BudgetName")
	}
	mg.Spec.ForProvider.BudgetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BudgetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Product",

			"ProductList",
		)
		if err != nil {
			return errors.Wrap(
				err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceIDRef,
			Selector:     mg.Spec.ForProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("budgets.aws.upbound.io",

			"v1beta1",
			"Budget", "BudgetList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BudgetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.BudgetNameRef,
			Selector:     mg.Spec.InitProvider.BudgetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BudgetName")
	}
	mg.Spec.InitProvider.BudgetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BudgetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Product",

			"ProductList",
		)
		if err != nil {
			return errors.Wrap(
				err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ResourceIDRef,
			Selector:     mg.Spec.InitProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceID")
	}
	mg.Spec.InitProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Constraint.
func (mg *Constraint) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Portfolio",

			"PortfolioList",
		)
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortfolioID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PortfolioIDRef,
			Selector:     mg.Spec.ForProvider.PortfolioIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortfolioID")
	}
	mg.Spec.ForProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortfolioIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Product",

			"ProductList",
		)
		if err != nil {
			return errors.Wrap(
				err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProductID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ProductIDRef,
			Selector:     mg.Spec.ForProvider.ProductIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProductID")
	}
	mg.Spec.ForProvider.ProductID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProductIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Portfolio",

			"PortfolioList",
		)
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PortfolioID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PortfolioIDRef,
			Selector:     mg.Spec.InitProvider.PortfolioIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PortfolioID")
	}
	mg.Spec.InitProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PortfolioIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Product",

			"ProductList",
		)
		if err != nil {
			return errors.Wrap(
				err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProductID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ProductIDRef,
			Selector:     mg.Spec.InitProvider.ProductIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProductID")
	}
	mg.Spec.InitProvider.ProductID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProductIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PortfolioShare.
func (mg *PortfolioShare) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Portfolio",

			"PortfolioList",
		)
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortfolioID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PortfolioIDRef,
			Selector:     mg.Spec.ForProvider.PortfolioIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortfolioID")
	}
	mg.Spec.ForProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortfolioIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Portfolio",

			"PortfolioList",
		)
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PortfolioID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PortfolioIDRef,
			Selector:     mg.Spec.InitProvider.PortfolioIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PortfolioID")
	}
	mg.Spec.InitProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PortfolioIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PrincipalPortfolioAssociation.
func (mg *PrincipalPortfolioAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Portfolio",

			"PortfolioList",
		)
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortfolioID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.PortfolioIDRef,
			Selector:     mg.Spec.ForProvider.PortfolioIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortfolioID")
	}
	mg.Spec.ForProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortfolioIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1",
			"User", "UserList",
		)
		if err !=

			nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrincipalArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.PrincipalArnRef,
			Selector:     mg.Spec.ForProvider.PrincipalArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrincipalArn")
	}
	mg.Spec.ForProvider.PrincipalArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrincipalArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Portfolio",

			"PortfolioList",
		)
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PortfolioID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.PortfolioIDRef,
			Selector:     mg.Spec.InitProvider.PortfolioIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PortfolioID")
	}
	mg.Spec.InitProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PortfolioIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1",
			"User", "UserList",
		)
		if err !=

			nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrincipalArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.PrincipalArnRef,
			Selector:     mg.Spec.InitProvider.PrincipalArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PrincipalArn")
	}
	mg.Spec.InitProvider.PrincipalArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PrincipalArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ProductPortfolioAssociation.
func (mg *ProductPortfolioAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Portfolio",

			"PortfolioList",
		)
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortfolioID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.PortfolioIDRef,
			Selector:     mg.Spec.ForProvider.PortfolioIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortfolioID")
	}
	mg.Spec.ForProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortfolioIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Product",

			"ProductList",
		)
		if err != nil {
			return errors.Wrap(
				err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProductID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ProductIDRef,
			Selector:     mg.Spec.ForProvider.ProductIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProductID")
	}
	mg.Spec.ForProvider.ProductID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProductIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Portfolio",

			"PortfolioList",
		)
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PortfolioID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.PortfolioIDRef,
			Selector:     mg.Spec.InitProvider.PortfolioIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PortfolioID")
	}
	mg.Spec.InitProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PortfolioIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Product",

			"ProductList",
		)
		if err != nil {
			return errors.Wrap(
				err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProductID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ProductIDRef,
			Selector:     mg.Spec.InitProvider.ProductIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProductID")
	}
	mg.Spec.InitProvider.ProductID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProductIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ProvisioningArtifact.
func (mg *ProvisioningArtifact) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Product",

			"ProductList",
		)
		if err != nil {
			return errors.Wrap(
				err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProductID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ProductIDRef,
			Selector:     mg.Spec.ForProvider.ProductIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProductID")
	}
	mg.Spec.ForProvider.ProductID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProductIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Product",

			"ProductList",
		)
		if err != nil {
			return errors.Wrap(
				err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProductID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ProductIDRef,
			Selector:     mg.Spec.InitProvider.ProductIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProductID")
	}
	mg.Spec.InitProvider.ProductID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProductIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TagOptionResourceAssociation.
func (mg *TagOptionResourceAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Product",

			"ProductList",
		)
		if err != nil {
			return errors.Wrap(
				err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceIDRef,
			Selector:     mg.Spec.ForProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "TagOption",

			"TagOptionList",
		)
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TagOptionID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.TagOptionIDRef,
			Selector:     mg.Spec.ForProvider.TagOptionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TagOptionID")
	}
	mg.Spec.ForProvider.TagOptionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TagOptionIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "Product",

			"ProductList",
		)
		if err != nil {
			return errors.Wrap(
				err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ResourceIDRef,
			Selector:     mg.Spec.InitProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceID")
	}
	mg.Spec.InitProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicecatalog.aws.upbound.io",

			"v1beta1", "TagOption",

			"TagOptionList",
		)
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TagOptionID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.TagOptionIDRef,
			Selector:     mg.Spec.InitProvider.TagOptionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TagOptionID")
	}
	mg.Spec.InitProvider.TagOptionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TagOptionIDRef = rsp.ResolvedReference

	return nil
}
