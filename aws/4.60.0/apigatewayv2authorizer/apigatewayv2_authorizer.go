// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apigatewayv2authorizer

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type JwtConfiguration struct {
	// Audience: set of string, optional
	Audience terra.SetValue[terra.StringValue] `hcl:"audience,attr"`
	// Issuer: string, optional
	Issuer terra.StringValue `hcl:"issuer,attr"`
}

type JwtConfigurationAttributes struct {
	ref terra.Reference
}

func (jc JwtConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return jc.ref, nil
}

func (jc JwtConfigurationAttributes) InternalWithRef(ref terra.Reference) JwtConfigurationAttributes {
	return JwtConfigurationAttributes{ref: ref}
}

func (jc JwtConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return jc.ref.InternalTokens()
}

func (jc JwtConfigurationAttributes) Audience() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](jc.ref.Append("audience"))
}

func (jc JwtConfigurationAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(jc.ref.Append("issuer"))
}

type JwtConfigurationState struct {
	Audience []string `json:"audience"`
	Issuer   string   `json:"issuer"`
}
