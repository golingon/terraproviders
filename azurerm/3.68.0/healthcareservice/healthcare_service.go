// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package healthcareservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AuthenticationConfiguration struct {
	// Audience: string, optional
	Audience terra.StringValue `hcl:"audience,attr"`
	// Authority: string, optional
	Authority terra.StringValue `hcl:"authority,attr"`
	// SmartProxyEnabled: bool, optional
	SmartProxyEnabled terra.BoolValue `hcl:"smart_proxy_enabled,attr"`
}

type CorsConfiguration struct {
	// AllowCredentials: bool, optional
	AllowCredentials terra.BoolValue `hcl:"allow_credentials,attr"`
	// AllowedHeaders: set of string, optional
	AllowedHeaders terra.SetValue[terra.StringValue] `hcl:"allowed_headers,attr"`
	// AllowedMethods: list of string, optional
	AllowedMethods terra.ListValue[terra.StringValue] `hcl:"allowed_methods,attr"`
	// AllowedOrigins: set of string, optional
	AllowedOrigins terra.SetValue[terra.StringValue] `hcl:"allowed_origins,attr"`
	// MaxAgeInSeconds: number, optional
	MaxAgeInSeconds terra.NumberValue `hcl:"max_age_in_seconds,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AuthenticationConfigurationAttributes struct {
	ref terra.Reference
}

func (ac AuthenticationConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AuthenticationConfigurationAttributes) InternalWithRef(ref terra.Reference) AuthenticationConfigurationAttributes {
	return AuthenticationConfigurationAttributes{ref: ref}
}

func (ac AuthenticationConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AuthenticationConfigurationAttributes) Audience() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("audience"))
}

func (ac AuthenticationConfigurationAttributes) Authority() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("authority"))
}

func (ac AuthenticationConfigurationAttributes) SmartProxyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("smart_proxy_enabled"))
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

func (cc CorsConfigurationAttributes) AllowedHeaders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cc.ref.Append("allowed_headers"))
}

func (cc CorsConfigurationAttributes) AllowedMethods() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("allowed_methods"))
}

func (cc CorsConfigurationAttributes) AllowedOrigins() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cc.ref.Append("allowed_origins"))
}

func (cc CorsConfigurationAttributes) MaxAgeInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("max_age_in_seconds"))
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

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AuthenticationConfigurationState struct {
	Audience          string `json:"audience"`
	Authority         string `json:"authority"`
	SmartProxyEnabled bool   `json:"smart_proxy_enabled"`
}

type CorsConfigurationState struct {
	AllowCredentials bool     `json:"allow_credentials"`
	AllowedHeaders   []string `json:"allowed_headers"`
	AllowedMethods   []string `json:"allowed_methods"`
	AllowedOrigins   []string `json:"allowed_origins"`
	MaxAgeInSeconds  float64  `json:"max_age_in_seconds"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}