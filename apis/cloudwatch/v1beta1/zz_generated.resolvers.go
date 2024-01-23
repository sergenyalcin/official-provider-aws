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

func (mg *CompositeAlarm) ResolveReferences( // ResolveReferences of this CompositeAlarm.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io",

			"v1beta1", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.AlarmActions),
			Extract:       resource.ExtractParamPath("arn", true),
			References:    mg.Spec.ForProvider.AlarmActionsRefs,
			Selector:      mg.Spec.ForProvider.AlarmActionsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AlarmActions")
	}
	mg.Spec.ForProvider.AlarmActions = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.AlarmActionsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io",

			"v1beta1", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.OkActions),
			Extract:       resource.ExtractParamPath("arn", true),
			References:    mg.Spec.ForProvider.OkActionsRefs,
			Selector:      mg.Spec.ForProvider.OkActionsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OkActions")
	}
	mg.Spec.ForProvider.OkActions = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.OkActionsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io",

			"v1beta1", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.AlarmActions),
			Extract:       resource.ExtractParamPath("arn", true),
			References:    mg.Spec.InitProvider.AlarmActionsRefs,
			Selector:      mg.Spec.InitProvider.AlarmActionsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AlarmActions")
	}
	mg.Spec.InitProvider.AlarmActions = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.AlarmActionsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io",

			"v1beta1", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.OkActions),
			Extract:       resource.ExtractParamPath("arn", true),
			References:    mg.Spec.InitProvider.OkActionsRefs,
			Selector:      mg.Spec.InitProvider.OkActionsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.OkActions")
	}
	mg.Spec.InitProvider.OkActions = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.OkActionsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this MetricStream.
func (mg *MetricStream) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io",

			"v1beta1", "DeliveryStream", "DeliveryStreamList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FirehoseArn),
			Extract:      resource.ExtractParamPath("arn", false),
			Reference:    mg.Spec.ForProvider.FirehoseArnRef,
			Selector:     mg.Spec.ForProvider.FirehoseArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FirehoseArn")
	}
	mg.Spec.ForProvider.FirehoseArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FirehoseArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RoleArnRef,
			Selector:     mg.Spec.ForProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io",

			"v1beta1", "DeliveryStream", "DeliveryStreamList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FirehoseArn),
			Extract:      resource.ExtractParamPath("arn", false),
			Reference:    mg.Spec.InitProvider.FirehoseArnRef,
			Selector:     mg.Spec.InitProvider.FirehoseArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FirehoseArn")
	}
	mg.Spec.InitProvider.FirehoseArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FirehoseArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err,
				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RoleArnRef,
			Selector:     mg.Spec.InitProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}
