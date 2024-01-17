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
	v1beta11 "github.com/upbound/provider-aws/apis/acm/v1beta1"
	v1beta13 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/lambda/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	apis "github.com/upbound/provider-aws/config/common/apis"
	lambda "github.com/upbound/provider-aws/config/common/apis/lambda"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this APIMapping.
func (mg *APIMapping) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DomainName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DomainNameRef,
		Selector:     mg.Spec.ForProvider.DomainNameSelector,
		To: reference.To{
			List:    &DomainNameList{},
			Managed: &DomainName{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DomainName")
	}
	mg.Spec.ForProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Stage),
		Extract:      common.TerraformID(),
		Reference:    mg.Spec.ForProvider.StageRef,
		Selector:     mg.Spec.ForProvider.StageSelector,
		To: reference.To{
			List:    &StageList{},
			Managed: &Stage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Stage")
	}
	mg.Spec.ForProvider.Stage = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StageRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.APIIDRef,
		Selector:     mg.Spec.InitProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DomainName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.DomainNameRef,
		Selector:     mg.Spec.InitProvider.DomainNameSelector,
		To: reference.To{
			List:    &DomainNameList{},
			Managed: &DomainName{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DomainName")
	}
	mg.Spec.InitProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DomainNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Stage),
		Extract:      common.TerraformID(),
		Reference:    mg.Spec.InitProvider.StageRef,
		Selector:     mg.Spec.InitProvider.StageSelector,
		To: reference.To{
			List:    &StageList{},
			Managed: &Stage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Stage")
	}
	mg.Spec.InitProvider.Stage = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StageRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Authorizer.
func (mg *Authorizer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthorizerURI),
		Extract:      lambda.FunctionInvokeARN(),
		Reference:    mg.Spec.ForProvider.AuthorizerURIRef,
		Selector:     mg.Spec.ForProvider.AuthorizerURISelector,
		To: reference.To{
			List:    &v1beta1.FunctionList{},
			Managed: &v1beta1.Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthorizerURI")
	}
	mg.Spec.ForProvider.AuthorizerURI = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthorizerURIRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.APIIDRef,
		Selector:     mg.Spec.InitProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AuthorizerURI),
		Extract:      lambda.FunctionInvokeARN(),
		Reference:    mg.Spec.InitProvider.AuthorizerURIRef,
		Selector:     mg.Spec.InitProvider.AuthorizerURISelector,
		To: reference.To{
			List:    &v1beta1.FunctionList{},
			Managed: &v1beta1.Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AuthorizerURI")
	}
	mg.Spec.InitProvider.AuthorizerURI = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AuthorizerURIRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Deployment.
func (mg *Deployment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.APIIDRef,
		Selector:     mg.Spec.InitProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DomainName.
func (mg *DomainName) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DomainNameConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DomainNameConfiguration[i3].CertificateArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.DomainNameConfiguration[i3].CertificateArnRef,
			Selector:     mg.Spec.ForProvider.DomainNameConfiguration[i3].CertificateArnSelector,
			To: reference.To{
				List:    &v1beta11.CertificateList{},
				Managed: &v1beta11.Certificate{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DomainNameConfiguration[i3].CertificateArn")
		}
		mg.Spec.ForProvider.DomainNameConfiguration[i3].CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DomainNameConfiguration[i3].CertificateArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.DomainNameConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DomainNameConfiguration[i3].CertificateArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.DomainNameConfiguration[i3].CertificateArnRef,
			Selector:     mg.Spec.InitProvider.DomainNameConfiguration[i3].CertificateArnSelector,
			To: reference.To{
				List:    &v1beta11.CertificateList{},
				Managed: &v1beta11.Certificate{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DomainNameConfiguration[i3].CertificateArn")
		}
		mg.Spec.InitProvider.DomainNameConfiguration[i3].CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DomainNameConfiguration[i3].CertificateArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Integration.
func (mg *Integration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConnectionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ConnectionIDRef,
		Selector:     mg.Spec.ForProvider.ConnectionIDSelector,
		To: reference.To{
			List:    &VPCLinkList{},
			Managed: &VPCLink{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConnectionID")
	}
	mg.Spec.ForProvider.ConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConnectionIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CredentialsArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.CredentialsArnRef,
		Selector:     mg.Spec.ForProvider.CredentialsArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CredentialsArn")
	}
	mg.Spec.ForProvider.CredentialsArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CredentialsArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IntegrationURI),
		Extract:      resource.ExtractParamPath("invoke_arn", true),
		Reference:    mg.Spec.ForProvider.IntegrationURIRef,
		Selector:     mg.Spec.ForProvider.IntegrationURISelector,
		To: reference.To{
			List:    &v1beta1.FunctionList{},
			Managed: &v1beta1.Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IntegrationURI")
	}
	mg.Spec.ForProvider.IntegrationURI = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IntegrationURIRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.APIIDRef,
		Selector:     mg.Spec.InitProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConnectionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.ConnectionIDRef,
		Selector:     mg.Spec.InitProvider.ConnectionIDSelector,
		To: reference.To{
			List:    &VPCLinkList{},
			Managed: &VPCLink{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConnectionID")
	}
	mg.Spec.InitProvider.ConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConnectionIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CredentialsArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.InitProvider.CredentialsArnRef,
		Selector:     mg.Spec.InitProvider.CredentialsArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CredentialsArn")
	}
	mg.Spec.InitProvider.CredentialsArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CredentialsArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IntegrationURI),
		Extract:      resource.ExtractParamPath("invoke_arn", true),
		Reference:    mg.Spec.InitProvider.IntegrationURIRef,
		Selector:     mg.Spec.InitProvider.IntegrationURISelector,
		To: reference.To{
			List:    &v1beta1.FunctionList{},
			Managed: &v1beta1.Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IntegrationURI")
	}
	mg.Spec.InitProvider.IntegrationURI = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IntegrationURIRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this IntegrationResponse.
func (mg *IntegrationResponse) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IntegrationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.IntegrationIDRef,
		Selector:     mg.Spec.ForProvider.IntegrationIDSelector,
		To: reference.To{
			List:    &IntegrationList{},
			Managed: &Integration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IntegrationID")
	}
	mg.Spec.ForProvider.IntegrationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IntegrationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.APIIDRef,
		Selector:     mg.Spec.InitProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IntegrationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.IntegrationIDRef,
		Selector:     mg.Spec.InitProvider.IntegrationIDSelector,
		To: reference.To{
			List:    &IntegrationList{},
			Managed: &Integration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IntegrationID")
	}
	mg.Spec.InitProvider.IntegrationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IntegrationIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Model.
func (mg *Model) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.APIIDRef,
		Selector:     mg.Spec.InitProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Route.
func (mg *Route) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthorizerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AuthorizerIDRef,
		Selector:     mg.Spec.ForProvider.AuthorizerIDSelector,
		To: reference.To{
			List:    &AuthorizerList{},
			Managed: &Authorizer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthorizerID")
	}
	mg.Spec.ForProvider.AuthorizerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthorizerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Target),
		Extract:      apis.IntegrationIDPrefixed(),
		Reference:    mg.Spec.ForProvider.TargetRef,
		Selector:     mg.Spec.ForProvider.TargetSelector,
		To: reference.To{
			List:    &IntegrationList{},
			Managed: &Integration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Target")
	}
	mg.Spec.ForProvider.Target = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.APIIDRef,
		Selector:     mg.Spec.InitProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AuthorizerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.AuthorizerIDRef,
		Selector:     mg.Spec.InitProvider.AuthorizerIDSelector,
		To: reference.To{
			List:    &AuthorizerList{},
			Managed: &Authorizer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AuthorizerID")
	}
	mg.Spec.InitProvider.AuthorizerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AuthorizerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Target),
		Extract:      apis.IntegrationIDPrefixed(),
		Reference:    mg.Spec.InitProvider.TargetRef,
		Selector:     mg.Spec.InitProvider.TargetSelector,
		To: reference.To{
			List:    &IntegrationList{},
			Managed: &Integration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Target")
	}
	mg.Spec.InitProvider.Target = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TargetRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RouteResponse.
func (mg *RouteResponse) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RouteID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RouteIDRef,
		Selector:     mg.Spec.ForProvider.RouteIDSelector,
		To: reference.To{
			List:    &RouteList{},
			Managed: &Route{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RouteID")
	}
	mg.Spec.ForProvider.RouteID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RouteIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.APIIDRef,
		Selector:     mg.Spec.InitProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RouteID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RouteIDRef,
		Selector:     mg.Spec.InitProvider.RouteIDSelector,
		To: reference.To{
			List:    &RouteList{},
			Managed: &Route{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RouteID")
	}
	mg.Spec.InitProvider.RouteID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RouteIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Stage.
func (mg *Stage) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeploymentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DeploymentIDRef,
		Selector:     mg.Spec.ForProvider.DeploymentIDSelector,
		To: reference.To{
			List:    &DeploymentList{},
			Managed: &Deployment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeploymentID")
	}
	mg.Spec.ForProvider.DeploymentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeploymentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.APIIDRef,
		Selector:     mg.Spec.InitProvider.APIIDSelector,
		To: reference.To{
			List:    &APIList{},
			Managed: &API{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DeploymentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.DeploymentIDRef,
		Selector:     mg.Spec.InitProvider.DeploymentIDSelector,
		To: reference.To{
			List:    &DeploymentList{},
			Managed: &Deployment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DeploymentID")
	}
	mg.Spec.InitProvider.DeploymentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DeploymentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VPCLink.
func (mg *VPCLink) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta13.SecurityGroupList{},
			Managed: &v1beta13.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIDRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIDRefs,
		Selector:      mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta13.SubnetList{},
			Managed: &v1beta13.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIDRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.SecurityGroupIDRefs,
		Selector:      mg.Spec.InitProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta13.SecurityGroupList{},
			Managed: &v1beta13.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroupIds")
	}
	mg.Spec.InitProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SecurityGroupIDRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.SubnetIDRefs,
		Selector:      mg.Spec.InitProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta13.SubnetList{},
			Managed: &v1beta13.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetIds")
	}
	mg.Spec.InitProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetIDRefs = mrsp.ResolvedReferences

	return nil
}
