// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package iamworkloadidentitypoolprovider

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Aws struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
}

type Oidc struct {
	// AllowedAudiences: list of string, optional
	AllowedAudiences terra.ListValue[terra.StringValue] `hcl:"allowed_audiences,attr"`
	// IssuerUri: string, required
	IssuerUri terra.StringValue `hcl:"issuer_uri,attr" validate:"required"`
	// JwksJson: string, optional
	JwksJson terra.StringValue `hcl:"jwks_json,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AwsAttributes struct {
	ref terra.Reference
}

func (a AwsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AwsAttributes) InternalWithRef(ref terra.Reference) AwsAttributes {
	return AwsAttributes{ref: ref}
}

func (a AwsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AwsAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("account_id"))
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

func (o OidcAttributes) AllowedAudiences() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](o.ref.Append("allowed_audiences"))
}

func (o OidcAttributes) IssuerUri() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("issuer_uri"))
}

func (o OidcAttributes) JwksJson() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("jwks_json"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AwsState struct {
	AccountId string `json:"account_id"`
}

type OidcState struct {
	AllowedAudiences []string `json:"allowed_audiences"`
	IssuerUri        string   `json:"issuer_uri"`
	JwksJson         string   `json:"jwks_json"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}