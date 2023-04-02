// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lambdafunctionurl

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Cors struct {
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

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
}

type CorsAttributes struct {
	ref terra.Reference
}

func (c CorsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CorsAttributes) InternalWithRef(ref terra.Reference) CorsAttributes {
	return CorsAttributes{ref: ref}
}

func (c CorsAttributes) InternalTokens() hclwrite.Tokens {
	return c.ref.InternalTokens()
}

func (c CorsAttributes) AllowCredentials() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("allow_credentials"))
}

func (c CorsAttributes) AllowHeaders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("allow_headers"))
}

func (c CorsAttributes) AllowMethods() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("allow_methods"))
}

func (c CorsAttributes) AllowOrigins() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("allow_origins"))
}

func (c CorsAttributes) ExposeHeaders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("expose_headers"))
}

func (c CorsAttributes) MaxAge() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("max_age"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

type CorsState struct {
	AllowCredentials bool     `json:"allow_credentials"`
	AllowHeaders     []string `json:"allow_headers"`
	AllowMethods     []string `json:"allow_methods"`
	AllowOrigins     []string `json:"allow_origins"`
	ExposeHeaders    []string `json:"expose_headers"`
	MaxAge           float64  `json:"max_age"`
}

type TimeoutsState struct {
	Create string `json:"create"`
}
