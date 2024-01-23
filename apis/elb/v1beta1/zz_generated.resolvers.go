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

	// ResolveReferences of this AppCookieStickinessPolicy.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *AppCookieStickinessPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancer),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LoadBalancerRef,
			Selector:     mg.Spec.ForProvider.LoadBalancerSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancer")
	}
	mg.Spec.ForProvider.LoadBalancer = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Attachment.
func (mg *Attachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ELB),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ELBRef,
			Selector:     mg.Spec.ForProvider.ELBSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ELB")
	}
	mg.Spec.ForProvider.ELB = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ELBRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io",

			"v1beta1",
			"Instance", "InstanceList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.InstanceRef,
			Selector:     mg.Spec.ForProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ELB),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ELBRef,
			Selector:     mg.Spec.InitProvider.ELBSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ELB")
	}
	mg.Spec.InitProvider.ELB = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ELBRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io",

			"v1beta1",
			"Instance", "InstanceList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Instance),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.InstanceRef,
			Selector:     mg.Spec.InitProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Instance")
	}
	mg.Spec.InitProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackendServerPolicy.
func (mg *BackendServerPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancerName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LoadBalancerNameRef,
			Selector:     mg.Spec.ForProvider.LoadBalancerNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancerName")
	}
	mg.Spec.ForProvider.LoadBalancerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancerName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LoadBalancerNameRef,
			Selector:     mg.Spec.InitProvider.LoadBalancerNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancerName")
	}
	mg.Spec.InitProvider.LoadBalancerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ELB.
func (mg *ELB) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io",

			"v1beta1",
			"Instance", "InstanceList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Instances),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.InstancesRefs,
			Selector:      mg.Spec.ForProvider.InstancesSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instances")
	}
	mg.Spec.ForProvider.Instances = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.InstancesRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io",

			"v1beta1",
			"Subnet", "SubnetList",
		)
		if err !=
			nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Subnets),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SubnetsRefs,
			Selector:      mg.Spec.ForProvider.SubnetsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnets")
	}
	mg.Spec.ForProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io",

			"v1beta1",
			"Instance", "InstanceList",
		)
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Instances),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.InstancesRefs,
			Selector:      mg.Spec.InitProvider.InstancesSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Instances")
	}
	mg.Spec.InitProvider.Instances = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.InstancesRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io",

			"v1beta1",
			"Subnet", "SubnetList",
		)
		if err !=
			nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Subnets),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SubnetsRefs,
			Selector:      mg.Spec.InitProvider.SubnetsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Subnets")
	}
	mg.Spec.InitProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this LBCookieStickinessPolicy.
func (mg *LBCookieStickinessPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancer),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.LoadBalancerRef,
			Selector:     mg.Spec.ForProvider.LoadBalancerSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancer")
	}
	mg.Spec.ForProvider.LoadBalancer = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancer),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.LoadBalancerRef,
			Selector:     mg.Spec.InitProvider.LoadBalancerSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancer")
	}
	mg.Spec.InitProvider.LoadBalancer = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBSSLNegotiationPolicy.
func (mg *LBSSLNegotiationPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancer),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.LoadBalancerRef,
			Selector:     mg.Spec.ForProvider.LoadBalancerSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancer")
	}
	mg.Spec.ForProvider.LoadBalancer = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancer),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.LoadBalancerRef,
			Selector:     mg.Spec.InitProvider.LoadBalancerSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancer")
	}
	mg.Spec.InitProvider.LoadBalancer = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ListenerPolicy.
func (mg *ListenerPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancerName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LoadBalancerNameRef,
			Selector:     mg.Spec.ForProvider.LoadBalancerNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancerName")
	}
	mg.Spec.ForProvider.LoadBalancerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancerName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LoadBalancerNameRef,
			Selector:     mg.Spec.InitProvider.LoadBalancerNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancerName")
	}
	mg.Spec.InitProvider.LoadBalancerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancerName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LoadBalancerNameRef,
			Selector:     mg.Spec.ForProvider.LoadBalancerNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancerName")
	}
	mg.Spec.ForProvider.LoadBalancerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PolicyAttribute); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

				"v1beta1",
				"Policy", "PolicyList",
			)
			if err !=
				nil {
				return errors.
					Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyAttribute[i3].Value),
				Extract:      resource.ExtractParamPath("policy_name", false),
				Reference:    mg.Spec.ForProvider.PolicyAttribute[i3].ValueRef,
				Selector:     mg.Spec.ForProvider.PolicyAttribute[i3].ValueSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PolicyAttribute[i3].Value")
		}
		mg.Spec.ForProvider.PolicyAttribute[i3].Value = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PolicyAttribute[i3].ValueRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancerName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LoadBalancerNameRef,
			Selector:     mg.Spec.InitProvider.LoadBalancerNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancerName")
	}
	mg.Spec.InitProvider.LoadBalancerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.PolicyAttribute); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

				"v1beta1",
				"Policy", "PolicyList",
			)
			if err !=
				nil {
				return errors.
					Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PolicyAttribute[i3].Value),
				Extract:      resource.ExtractParamPath("policy_name", false),
				Reference:    mg.Spec.InitProvider.PolicyAttribute[i3].ValueRef,
				Selector:     mg.Spec.InitProvider.PolicyAttribute[i3].ValueSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PolicyAttribute[i3].Value")
		}
		mg.Spec.InitProvider.PolicyAttribute[i3].Value = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PolicyAttribute[i3].ValueRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ProxyProtocolPolicy.
func (mg *ProxyProtocolPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancer),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LoadBalancerRef,
			Selector:     mg.Spec.ForProvider.LoadBalancerSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancer")
	}
	mg.Spec.ForProvider.LoadBalancer = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io",

			"v1beta1",
			"ELB", "ELBList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancer),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LoadBalancerRef,
			Selector:     mg.Spec.InitProvider.LoadBalancerSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancer")
	}
	mg.Spec.InitProvider.LoadBalancer = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerRef = rsp.ResolvedReference

	return nil
}
