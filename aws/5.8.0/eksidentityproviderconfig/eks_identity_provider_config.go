// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package eksidentityproviderconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Oidc struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// GroupsClaim: string, optional
	GroupsClaim terra.StringValue `hcl:"groups_claim,attr"`
	// GroupsPrefix: string, optional
	GroupsPrefix terra.StringValue `hcl:"groups_prefix,attr"`
	// IdentityProviderConfigName: string, required
	IdentityProviderConfigName terra.StringValue `hcl:"identity_provider_config_name,attr" validate:"required"`
	// IssuerUrl: string, required
	IssuerUrl terra.StringValue `hcl:"issuer_url,attr" validate:"required"`
	// RequiredClaims: map of string, optional
	RequiredClaims terra.MapValue[terra.StringValue] `hcl:"required_claims,attr"`
	// UsernameClaim: string, optional
	UsernameClaim terra.StringValue `hcl:"username_claim,attr"`
	// UsernamePrefix: string, optional
	UsernamePrefix terra.StringValue `hcl:"username_prefix,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type OidcAttributes struct {
	ref terra.Reference
}

func (o OidcAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OidcAttributes) InternalWithRef(ref terra.Reference) OidcAttributes {
	return OidcAttributes{ref: ref}
}

func (o OidcAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OidcAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("client_id"))
}

func (o OidcAttributes) GroupsClaim() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("groups_claim"))
}

func (o OidcAttributes) GroupsPrefix() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("groups_prefix"))
}

func (o OidcAttributes) IdentityProviderConfigName() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("identity_provider_config_name"))
}

func (o OidcAttributes) IssuerUrl() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("issuer_url"))
}

func (o OidcAttributes) RequiredClaims() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](o.ref.Append("required_claims"))
}

func (o OidcAttributes) UsernameClaim() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("username_claim"))
}

func (o OidcAttributes) UsernamePrefix() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("username_prefix"))
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

type OidcState struct {
	ClientId                   string            `json:"client_id"`
	GroupsClaim                string            `json:"groups_claim"`
	GroupsPrefix               string            `json:"groups_prefix"`
	IdentityProviderConfigName string            `json:"identity_provider_config_name"`
	IssuerUrl                  string            `json:"issuer_url"`
	RequiredClaims             map[string]string `json:"required_claims"`
	UsernameClaim              string            `json:"username_claim"`
	UsernamePrefix             string            `json:"username_prefix"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
