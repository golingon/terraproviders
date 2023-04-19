// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package iamworkforcepoolprovider

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Oidc struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// IssuerUri: string, required
	IssuerUri terra.StringValue `hcl:"issuer_uri,attr" validate:"required"`
}

type Saml struct {
	// IdpMetadataXml: string, required
	IdpMetadataXml terra.StringValue `hcl:"idp_metadata_xml,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
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

func (o OidcAttributes) IssuerUri() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("issuer_uri"))
}

type SamlAttributes struct {
	ref terra.Reference
}

func (s SamlAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SamlAttributes) InternalWithRef(ref terra.Reference) SamlAttributes {
	return SamlAttributes{ref: ref}
}

func (s SamlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SamlAttributes) IdpMetadataXml() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("idp_metadata_xml"))
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

type OidcState struct {
	ClientId  string `json:"client_id"`
	IssuerUri string `json:"issuer_uri"`
}

type SamlState struct {
	IdpMetadataXml string `json:"idp_metadata_xml"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
