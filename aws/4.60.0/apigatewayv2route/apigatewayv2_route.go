// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apigatewayv2route

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RequestParameter struct {
	// RequestParameterKey: string, required
	RequestParameterKey terra.StringValue `hcl:"request_parameter_key,attr" validate:"required"`
	// Required: bool, required
	Required terra.BoolValue `hcl:"required,attr" validate:"required"`
}

type RequestParameterAttributes struct {
	ref terra.Reference
}

func (rp RequestParameterAttributes) InternalRef() terra.Reference {
	return rp.ref
}

func (rp RequestParameterAttributes) InternalWithRef(ref terra.Reference) RequestParameterAttributes {
	return RequestParameterAttributes{ref: ref}
}

func (rp RequestParameterAttributes) InternalTokens() hclwrite.Tokens {
	return rp.ref.InternalTokens()
}

func (rp RequestParameterAttributes) RequestParameterKey() terra.StringValue {
	return terra.ReferenceString(rp.ref.Append("request_parameter_key"))
}

func (rp RequestParameterAttributes) Required() terra.BoolValue {
	return terra.ReferenceBool(rp.ref.Append("required"))
}

type RequestParameterState struct {
	RequestParameterKey string `json:"request_parameter_key"`
	Required            bool   `json:"required"`
}
