// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package s3bucketwebsiteconfiguration

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ErrorDocument struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
}

type IndexDocument struct {
	// Suffix: string, required
	Suffix terra.StringValue `hcl:"suffix,attr" validate:"required"`
}

type RedirectAllRequestsTo struct {
	// HostName: string, required
	HostName terra.StringValue `hcl:"host_name,attr" validate:"required"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
}

type RoutingRule struct {
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
	// Redirect: required
	Redirect *Redirect `hcl:"redirect,block" validate:"required"`
}

type Condition struct {
	// HttpErrorCodeReturnedEquals: string, optional
	HttpErrorCodeReturnedEquals terra.StringValue `hcl:"http_error_code_returned_equals,attr"`
	// KeyPrefixEquals: string, optional
	KeyPrefixEquals terra.StringValue `hcl:"key_prefix_equals,attr"`
}

type Redirect struct {
	// HostName: string, optional
	HostName terra.StringValue `hcl:"host_name,attr"`
	// HttpRedirectCode: string, optional
	HttpRedirectCode terra.StringValue `hcl:"http_redirect_code,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// ReplaceKeyPrefixWith: string, optional
	ReplaceKeyPrefixWith terra.StringValue `hcl:"replace_key_prefix_with,attr"`
	// ReplaceKeyWith: string, optional
	ReplaceKeyWith terra.StringValue `hcl:"replace_key_with,attr"`
}

type ErrorDocumentAttributes struct {
	ref terra.Reference
}

func (ed ErrorDocumentAttributes) InternalRef() (terra.Reference, error) {
	return ed.ref, nil
}

func (ed ErrorDocumentAttributes) InternalWithRef(ref terra.Reference) ErrorDocumentAttributes {
	return ErrorDocumentAttributes{ref: ref}
}

func (ed ErrorDocumentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ed.ref.InternalTokens()
}

func (ed ErrorDocumentAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("key"))
}

type IndexDocumentAttributes struct {
	ref terra.Reference
}

func (id IndexDocumentAttributes) InternalRef() (terra.Reference, error) {
	return id.ref, nil
}

func (id IndexDocumentAttributes) InternalWithRef(ref terra.Reference) IndexDocumentAttributes {
	return IndexDocumentAttributes{ref: ref}
}

func (id IndexDocumentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return id.ref.InternalTokens()
}

func (id IndexDocumentAttributes) Suffix() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("suffix"))
}

type RedirectAllRequestsToAttributes struct {
	ref terra.Reference
}

func (rart RedirectAllRequestsToAttributes) InternalRef() (terra.Reference, error) {
	return rart.ref, nil
}

func (rart RedirectAllRequestsToAttributes) InternalWithRef(ref terra.Reference) RedirectAllRequestsToAttributes {
	return RedirectAllRequestsToAttributes{ref: ref}
}

func (rart RedirectAllRequestsToAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rart.ref.InternalTokens()
}

func (rart RedirectAllRequestsToAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(rart.ref.Append("host_name"))
}

func (rart RedirectAllRequestsToAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(rart.ref.Append("protocol"))
}

type RoutingRuleAttributes struct {
	ref terra.Reference
}

func (rr RoutingRuleAttributes) InternalRef() (terra.Reference, error) {
	return rr.ref, nil
}

func (rr RoutingRuleAttributes) InternalWithRef(ref terra.Reference) RoutingRuleAttributes {
	return RoutingRuleAttributes{ref: ref}
}

func (rr RoutingRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rr.ref.InternalTokens()
}

func (rr RoutingRuleAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](rr.ref.Append("condition"))
}

func (rr RoutingRuleAttributes) Redirect() terra.ListValue[RedirectAttributes] {
	return terra.ReferenceAsList[RedirectAttributes](rr.ref.Append("redirect"))
}

type ConditionAttributes struct {
	ref terra.Reference
}

func (c ConditionAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionAttributes) InternalWithRef(ref terra.Reference) ConditionAttributes {
	return ConditionAttributes{ref: ref}
}

func (c ConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionAttributes) HttpErrorCodeReturnedEquals() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("http_error_code_returned_equals"))
}

func (c ConditionAttributes) KeyPrefixEquals() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("key_prefix_equals"))
}

type RedirectAttributes struct {
	ref terra.Reference
}

func (r RedirectAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RedirectAttributes) InternalWithRef(ref terra.Reference) RedirectAttributes {
	return RedirectAttributes{ref: ref}
}

func (r RedirectAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RedirectAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("host_name"))
}

func (r RedirectAttributes) HttpRedirectCode() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("http_redirect_code"))
}

func (r RedirectAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("protocol"))
}

func (r RedirectAttributes) ReplaceKeyPrefixWith() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("replace_key_prefix_with"))
}

func (r RedirectAttributes) ReplaceKeyWith() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("replace_key_with"))
}

type ErrorDocumentState struct {
	Key string `json:"key"`
}

type IndexDocumentState struct {
	Suffix string `json:"suffix"`
}

type RedirectAllRequestsToState struct {
	HostName string `json:"host_name"`
	Protocol string `json:"protocol"`
}

type RoutingRuleState struct {
	Condition []ConditionState `json:"condition"`
	Redirect  []RedirectState  `json:"redirect"`
}

type ConditionState struct {
	HttpErrorCodeReturnedEquals string `json:"http_error_code_returned_equals"`
	KeyPrefixEquals             string `json:"key_prefix_equals"`
}

type RedirectState struct {
	HostName             string `json:"host_name"`
	HttpRedirectCode     string `json:"http_redirect_code"`
	Protocol             string `json:"protocol"`
	ReplaceKeyPrefixWith string `json:"replace_key_prefix_with"`
	ReplaceKeyWith       string `json:"replace_key_with"`
}
