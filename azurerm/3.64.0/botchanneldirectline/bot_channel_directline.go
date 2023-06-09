// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package botchanneldirectline

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Site struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EnhancedAuthenticationEnabled: bool, optional
	EnhancedAuthenticationEnabled terra.BoolValue `hcl:"enhanced_authentication_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TrustedOrigins: set of string, optional
	TrustedOrigins terra.SetValue[terra.StringValue] `hcl:"trusted_origins,attr"`
	// V1Allowed: bool, optional
	V1Allowed terra.BoolValue `hcl:"v1_allowed,attr"`
	// V3Allowed: bool, optional
	V3Allowed terra.BoolValue `hcl:"v3_allowed,attr"`
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

type SiteAttributes struct {
	ref terra.Reference
}

func (s SiteAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SiteAttributes) InternalWithRef(ref terra.Reference) SiteAttributes {
	return SiteAttributes{ref: ref}
}

func (s SiteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SiteAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enabled"))
}

func (s SiteAttributes) EnhancedAuthenticationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enhanced_authentication_enabled"))
}

func (s SiteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

func (s SiteAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("key"))
}

func (s SiteAttributes) Key2() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("key2"))
}

func (s SiteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SiteAttributes) TrustedOrigins() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("trusted_origins"))
}

func (s SiteAttributes) V1Allowed() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("v1_allowed"))
}

func (s SiteAttributes) V3Allowed() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("v3_allowed"))
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

type SiteState struct {
	Enabled                       bool     `json:"enabled"`
	EnhancedAuthenticationEnabled bool     `json:"enhanced_authentication_enabled"`
	Id                            string   `json:"id"`
	Key                           string   `json:"key"`
	Key2                          string   `json:"key2"`
	Name                          string   `json:"name"`
	TrustedOrigins                []string `json:"trusted_origins"`
	V1Allowed                     bool     `json:"v1_allowed"`
	V3Allowed                     bool     `json:"v3_allowed"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
