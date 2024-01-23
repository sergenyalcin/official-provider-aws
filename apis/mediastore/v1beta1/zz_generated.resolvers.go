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
	errors "github.com/pkg/errors"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ContainerPolicy) ResolveReferences( // ResolveReferences of this ContainerPolicy.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("mediastore.aws.upbound.io", "v1beta1", "Container", "ContainerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ContainerNameRef,
			Selector:     mg.Spec.ForProvider.ContainerNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContainerName")
	}
	mg.Spec.ForProvider.ContainerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContainerNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("mediastore.aws.upbound.io", "v1beta1", "Container", "ContainerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ContainerName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ContainerNameRef,
			Selector:     mg.Spec.InitProvider.ContainerNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ContainerName")
	}
	mg.Spec.InitProvider.ContainerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ContainerNameRef = rsp.ResolvedReference

	return nil
}
