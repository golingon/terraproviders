// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datahealthcareservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AuthenticationConfiguration struct{}

type CorsConfiguration struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
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
	Read string `json:"read"`
}