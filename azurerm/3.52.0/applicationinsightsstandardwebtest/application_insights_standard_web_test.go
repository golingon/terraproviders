// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package applicationinsightsstandardwebtest

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Request struct {
	// Body: string, optional
	Body terra.StringValue `hcl:"body,attr"`
	// FollowRedirectsEnabled: bool, optional
	FollowRedirectsEnabled terra.BoolValue `hcl:"follow_redirects_enabled,attr"`
	// HttpVerb: string, optional
	HttpVerb terra.StringValue `hcl:"http_verb,attr"`
	// ParseDependentRequestsEnabled: bool, optional
	ParseDependentRequestsEnabled terra.BoolValue `hcl:"parse_dependent_requests_enabled,attr"`
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
	// Header: min=0
	Header []Header `hcl:"header,block" validate:"min=0"`
}

type Header struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
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

type ValidationRules struct {
	// ExpectedStatusCode: number, optional
	ExpectedStatusCode terra.NumberValue `hcl:"expected_status_code,attr"`
	// SslCertRemainingLifetime: number, optional
	SslCertRemainingLifetime terra.NumberValue `hcl:"ssl_cert_remaining_lifetime,attr"`
	// SslCheckEnabled: bool, optional
	SslCheckEnabled terra.BoolValue `hcl:"ssl_check_enabled,attr"`
	// Content: optional
	Content *Content `hcl:"content,block"`
}

type Content struct {
	// ContentMatch: string, required
	ContentMatch terra.StringValue `hcl:"content_match,attr" validate:"required"`
	// IgnoreCase: bool, optional
	IgnoreCase terra.BoolValue `hcl:"ignore_case,attr"`
	// PassIfTextFound: bool, optional
	PassIfTextFound terra.BoolValue `hcl:"pass_if_text_found,attr"`
}

type RequestAttributes struct {
	ref terra.Reference
}

func (r RequestAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RequestAttributes) InternalWithRef(ref terra.Reference) RequestAttributes {
	return RequestAttributes{ref: ref}
}

func (r RequestAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RequestAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("body"))
}

func (r RequestAttributes) FollowRedirectsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("follow_redirects_enabled"))
}

func (r RequestAttributes) HttpVerb() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("http_verb"))
}

func (r RequestAttributes) ParseDependentRequestsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("parse_dependent_requests_enabled"))
}

func (r RequestAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("url"))
}

func (r RequestAttributes) Header() terra.ListValue[HeaderAttributes] {
	return terra.ReferenceAsList[HeaderAttributes](r.ref.Append("header"))
}

type HeaderAttributes struct {
	ref terra.Reference
}

func (h HeaderAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h HeaderAttributes) InternalWithRef(ref terra.Reference) HeaderAttributes {
	return HeaderAttributes{ref: ref}
}

func (h HeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h HeaderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("name"))
}

func (h HeaderAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("value"))
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

type ValidationRulesAttributes struct {
	ref terra.Reference
}

func (vr ValidationRulesAttributes) InternalRef() (terra.Reference, error) {
	return vr.ref, nil
}

func (vr ValidationRulesAttributes) InternalWithRef(ref terra.Reference) ValidationRulesAttributes {
	return ValidationRulesAttributes{ref: ref}
}

func (vr ValidationRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vr.ref.InternalTokens()
}

func (vr ValidationRulesAttributes) ExpectedStatusCode() terra.NumberValue {
	return terra.ReferenceAsNumber(vr.ref.Append("expected_status_code"))
}

func (vr ValidationRulesAttributes) SslCertRemainingLifetime() terra.NumberValue {
	return terra.ReferenceAsNumber(vr.ref.Append("ssl_cert_remaining_lifetime"))
}

func (vr ValidationRulesAttributes) SslCheckEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vr.ref.Append("ssl_check_enabled"))
}

func (vr ValidationRulesAttributes) Content() terra.ListValue[ContentAttributes] {
	return terra.ReferenceAsList[ContentAttributes](vr.ref.Append("content"))
}

type ContentAttributes struct {
	ref terra.Reference
}

func (c ContentAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ContentAttributes) InternalWithRef(ref terra.Reference) ContentAttributes {
	return ContentAttributes{ref: ref}
}

func (c ContentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ContentAttributes) ContentMatch() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("content_match"))
}

func (c ContentAttributes) IgnoreCase() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("ignore_case"))
}

func (c ContentAttributes) PassIfTextFound() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("pass_if_text_found"))
}

type RequestState struct {
	Body                          string        `json:"body"`
	FollowRedirectsEnabled        bool          `json:"follow_redirects_enabled"`
	HttpVerb                      string        `json:"http_verb"`
	ParseDependentRequestsEnabled bool          `json:"parse_dependent_requests_enabled"`
	Url                           string        `json:"url"`
	Header                        []HeaderState `json:"header"`
}

type HeaderState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type ValidationRulesState struct {
	ExpectedStatusCode       float64        `json:"expected_status_code"`
	SslCertRemainingLifetime float64        `json:"ssl_cert_remaining_lifetime"`
	SslCheckEnabled          bool           `json:"ssl_check_enabled"`
	Content                  []ContentState `json:"content"`
}

type ContentState struct {
	ContentMatch    string `json:"content_match"`
	IgnoreCase      bool   `json:"ignore_case"`
	PassIfTextFound bool   `json:"pass_if_text_found"`
}
