// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_apigatewayv2_api

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CorsConfiguration struct {
	// AllowCredentials: bool, optional
	AllowCredentials terra.BoolValue `hcl:"allow_credentials,attr"`
	// AllowHeaders: set of string, optional
	AllowHeaders terra.SetValue[terra.StringValue] `hcl:"allow_headers,attr"`
	// AllowMethods: set of string, optional
	AllowMethods terra.SetValue[terra.StringValue] `hcl:"allow_methods,attr"`
	// AllowOrigins: set of string, optional
	AllowOrigins terra.SetValue[terra.StringValue] `hcl:"allow_origins,attr"`
	// ExposeHeaders: set of string, optional
	ExposeHeaders terra.SetValue[terra.StringValue] `hcl:"expose_headers,attr"`
	// MaxAge: number, optional
	MaxAge terra.NumberValue `hcl:"max_age,attr"`
}

type CorsConfigurationAttributes struct {
	ref terra.Reference
}

func (cc CorsConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc CorsConfigurationAttributes) InternalWithRef(ref terra.Reference) CorsConfigurationAttributes {
	return CorsConfigurationAttributes{ref: ref}
}

func (cc CorsConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc CorsConfigurationAttributes) AllowCredentials() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("allow_credentials"))
}

func (cc CorsConfigurationAttributes) AllowHeaders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cc.ref.Append("allow_headers"))
}

func (cc CorsConfigurationAttributes) AllowMethods() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cc.ref.Append("allow_methods"))
}

func (cc CorsConfigurationAttributes) AllowOrigins() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cc.ref.Append("allow_origins"))
}

func (cc CorsConfigurationAttributes) ExposeHeaders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cc.ref.Append("expose_headers"))
}

func (cc CorsConfigurationAttributes) MaxAge() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("max_age"))
}

type CorsConfigurationState struct {
	AllowCredentials bool     `json:"allow_credentials"`
	AllowHeaders     []string `json:"allow_headers"`
	AllowMethods     []string `json:"allow_methods"`
	AllowOrigins     []string `json:"allow_origins"`
	ExposeHeaders    []string `json:"expose_headers"`
	MaxAge           float64  `json:"max_age"`
}
