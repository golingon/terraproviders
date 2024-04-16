// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_gcp_auth_backend

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Tune struct {
	// AllowedResponseHeaders: list of string, optional
	AllowedResponseHeaders terra.ListValue[terra.StringValue] `hcl:"allowed_response_headers,attr"`
	// AuditNonHmacRequestKeys: list of string, optional
	AuditNonHmacRequestKeys terra.ListValue[terra.StringValue] `hcl:"audit_non_hmac_request_keys,attr"`
	// AuditNonHmacResponseKeys: list of string, optional
	AuditNonHmacResponseKeys terra.ListValue[terra.StringValue] `hcl:"audit_non_hmac_response_keys,attr"`
	// DefaultLeaseTtl: string, optional
	DefaultLeaseTtl terra.StringValue `hcl:"default_lease_ttl,attr"`
	// ListingVisibility: string, optional
	ListingVisibility terra.StringValue `hcl:"listing_visibility,attr"`
	// MaxLeaseTtl: string, optional
	MaxLeaseTtl terra.StringValue `hcl:"max_lease_ttl,attr"`
	// PassthroughRequestHeaders: list of string, optional
	PassthroughRequestHeaders terra.ListValue[terra.StringValue] `hcl:"passthrough_request_headers,attr"`
	// TokenType: string, optional
	TokenType terra.StringValue `hcl:"token_type,attr"`
}

type CustomEndpoint struct {
	// Api: string, optional
	Api terra.StringValue `hcl:"api,attr"`
	// Compute: string, optional
	Compute terra.StringValue `hcl:"compute,attr"`
	// Crm: string, optional
	Crm terra.StringValue `hcl:"crm,attr"`
	// Iam: string, optional
	Iam terra.StringValue `hcl:"iam,attr"`
}

type TuneAttributes struct {
	ref terra.Reference
}

func (t TuneAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TuneAttributes) InternalWithRef(ref terra.Reference) TuneAttributes {
	return TuneAttributes{ref: ref}
}

func (t TuneAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TuneAttributes) AllowedResponseHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](t.ref.Append("allowed_response_headers"))
}

func (t TuneAttributes) AuditNonHmacRequestKeys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](t.ref.Append("audit_non_hmac_request_keys"))
}

func (t TuneAttributes) AuditNonHmacResponseKeys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](t.ref.Append("audit_non_hmac_response_keys"))
}

func (t TuneAttributes) DefaultLeaseTtl() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("default_lease_ttl"))
}

func (t TuneAttributes) ListingVisibility() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("listing_visibility"))
}

func (t TuneAttributes) MaxLeaseTtl() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("max_lease_ttl"))
}

func (t TuneAttributes) PassthroughRequestHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](t.ref.Append("passthrough_request_headers"))
}

func (t TuneAttributes) TokenType() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("token_type"))
}

type CustomEndpointAttributes struct {
	ref terra.Reference
}

func (ce CustomEndpointAttributes) InternalRef() (terra.Reference, error) {
	return ce.ref, nil
}

func (ce CustomEndpointAttributes) InternalWithRef(ref terra.Reference) CustomEndpointAttributes {
	return CustomEndpointAttributes{ref: ref}
}

func (ce CustomEndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ce.ref.InternalTokens()
}

func (ce CustomEndpointAttributes) Api() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("api"))
}

func (ce CustomEndpointAttributes) Compute() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("compute"))
}

func (ce CustomEndpointAttributes) Crm() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("crm"))
}

func (ce CustomEndpointAttributes) Iam() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("iam"))
}

type TuneState struct {
	AllowedResponseHeaders    []string `json:"allowed_response_headers"`
	AuditNonHmacRequestKeys   []string `json:"audit_non_hmac_request_keys"`
	AuditNonHmacResponseKeys  []string `json:"audit_non_hmac_response_keys"`
	DefaultLeaseTtl           string   `json:"default_lease_ttl"`
	ListingVisibility         string   `json:"listing_visibility"`
	MaxLeaseTtl               string   `json:"max_lease_ttl"`
	PassthroughRequestHeaders []string `json:"passthrough_request_headers"`
	TokenType                 string   `json:"token_type"`
}

type CustomEndpointState struct {
	Api     string `json:"api"`
	Compute string `json:"compute"`
	Crm     string `json:"crm"`
	Iam     string `json:"iam"`
}
