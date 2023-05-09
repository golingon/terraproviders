// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package codepipelinewebhook

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AuthenticationConfiguration struct {
	// AllowedIpRange: string, optional
	AllowedIpRange terra.StringValue `hcl:"allowed_ip_range,attr"`
	// SecretToken: string, optional
	SecretToken terra.StringValue `hcl:"secret_token,attr"`
}

type Filter struct {
	// JsonPath: string, required
	JsonPath terra.StringValue `hcl:"json_path,attr" validate:"required"`
	// MatchEquals: string, required
	MatchEquals terra.StringValue `hcl:"match_equals,attr" validate:"required"`
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

func (ac AuthenticationConfigurationAttributes) AllowedIpRange() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("allowed_ip_range"))
}

func (ac AuthenticationConfigurationAttributes) SecretToken() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("secret_token"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) JsonPath() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("json_path"))
}

func (f FilterAttributes) MatchEquals() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("match_equals"))
}

type AuthenticationConfigurationState struct {
	AllowedIpRange string `json:"allowed_ip_range"`
	SecretToken    string `json:"secret_token"`
}

type FilterState struct {
	JsonPath    string `json:"json_path"`
	MatchEquals string `json:"match_equals"`
}
