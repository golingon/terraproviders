// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudfront_origin_request_policy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CookiesConfig struct {
	// CookieBehavior: string, required
	CookieBehavior terra.StringValue `hcl:"cookie_behavior,attr" validate:"required"`
	// CookiesConfigCookies: optional
	Cookies *CookiesConfigCookies `hcl:"cookies,block"`
}

type CookiesConfigCookies struct {
	// Items: set of string, optional
	Items terra.SetValue[terra.StringValue] `hcl:"items,attr"`
}

type HeadersConfig struct {
	// HeaderBehavior: string, optional
	HeaderBehavior terra.StringValue `hcl:"header_behavior,attr"`
	// HeadersConfigHeaders: optional
	Headers *HeadersConfigHeaders `hcl:"headers,block"`
}

type HeadersConfigHeaders struct {
	// Items: set of string, optional
	Items terra.SetValue[terra.StringValue] `hcl:"items,attr"`
}

type QueryStringsConfig struct {
	// QueryStringBehavior: string, required
	QueryStringBehavior terra.StringValue `hcl:"query_string_behavior,attr" validate:"required"`
	// QueryStringsConfigQueryStrings: optional
	QueryStrings *QueryStringsConfigQueryStrings `hcl:"query_strings,block"`
}

type QueryStringsConfigQueryStrings struct {
	// Items: set of string, optional
	Items terra.SetValue[terra.StringValue] `hcl:"items,attr"`
}

type CookiesConfigAttributes struct {
	ref terra.Reference
}

func (cc CookiesConfigAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc CookiesConfigAttributes) InternalWithRef(ref terra.Reference) CookiesConfigAttributes {
	return CookiesConfigAttributes{ref: ref}
}

func (cc CookiesConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc CookiesConfigAttributes) CookieBehavior() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("cookie_behavior"))
}

func (cc CookiesConfigAttributes) Cookies() terra.ListValue[CookiesConfigCookiesAttributes] {
	return terra.ReferenceAsList[CookiesConfigCookiesAttributes](cc.ref.Append("cookies"))
}

type CookiesConfigCookiesAttributes struct {
	ref terra.Reference
}

func (c CookiesConfigCookiesAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CookiesConfigCookiesAttributes) InternalWithRef(ref terra.Reference) CookiesConfigCookiesAttributes {
	return CookiesConfigCookiesAttributes{ref: ref}
}

func (c CookiesConfigCookiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CookiesConfigCookiesAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("items"))
}

type HeadersConfigAttributes struct {
	ref terra.Reference
}

func (hc HeadersConfigAttributes) InternalRef() (terra.Reference, error) {
	return hc.ref, nil
}

func (hc HeadersConfigAttributes) InternalWithRef(ref terra.Reference) HeadersConfigAttributes {
	return HeadersConfigAttributes{ref: ref}
}

func (hc HeadersConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hc.ref.InternalTokens()
}

func (hc HeadersConfigAttributes) HeaderBehavior() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("header_behavior"))
}

func (hc HeadersConfigAttributes) Headers() terra.ListValue[HeadersConfigHeadersAttributes] {
	return terra.ReferenceAsList[HeadersConfigHeadersAttributes](hc.ref.Append("headers"))
}

type HeadersConfigHeadersAttributes struct {
	ref terra.Reference
}

func (h HeadersConfigHeadersAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h HeadersConfigHeadersAttributes) InternalWithRef(ref terra.Reference) HeadersConfigHeadersAttributes {
	return HeadersConfigHeadersAttributes{ref: ref}
}

func (h HeadersConfigHeadersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h HeadersConfigHeadersAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](h.ref.Append("items"))
}

type QueryStringsConfigAttributes struct {
	ref terra.Reference
}

func (qsc QueryStringsConfigAttributes) InternalRef() (terra.Reference, error) {
	return qsc.ref, nil
}

func (qsc QueryStringsConfigAttributes) InternalWithRef(ref terra.Reference) QueryStringsConfigAttributes {
	return QueryStringsConfigAttributes{ref: ref}
}

func (qsc QueryStringsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qsc.ref.InternalTokens()
}

func (qsc QueryStringsConfigAttributes) QueryStringBehavior() terra.StringValue {
	return terra.ReferenceAsString(qsc.ref.Append("query_string_behavior"))
}

func (qsc QueryStringsConfigAttributes) QueryStrings() terra.ListValue[QueryStringsConfigQueryStringsAttributes] {
	return terra.ReferenceAsList[QueryStringsConfigQueryStringsAttributes](qsc.ref.Append("query_strings"))
}

type QueryStringsConfigQueryStringsAttributes struct {
	ref terra.Reference
}

func (qs QueryStringsConfigQueryStringsAttributes) InternalRef() (terra.Reference, error) {
	return qs.ref, nil
}

func (qs QueryStringsConfigQueryStringsAttributes) InternalWithRef(ref terra.Reference) QueryStringsConfigQueryStringsAttributes {
	return QueryStringsConfigQueryStringsAttributes{ref: ref}
}

func (qs QueryStringsConfigQueryStringsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qs.ref.InternalTokens()
}

func (qs QueryStringsConfigQueryStringsAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](qs.ref.Append("items"))
}

type CookiesConfigState struct {
	CookieBehavior string                      `json:"cookie_behavior"`
	Cookies        []CookiesConfigCookiesState `json:"cookies"`
}

type CookiesConfigCookiesState struct {
	Items []string `json:"items"`
}

type HeadersConfigState struct {
	HeaderBehavior string                      `json:"header_behavior"`
	Headers        []HeadersConfigHeadersState `json:"headers"`
}

type HeadersConfigHeadersState struct {
	Items []string `json:"items"`
}

type QueryStringsConfigState struct {
	QueryStringBehavior string                                `json:"query_string_behavior"`
	QueryStrings        []QueryStringsConfigQueryStringsState `json:"query_strings"`
}

type QueryStringsConfigQueryStringsState struct {
	Items []string `json:"items"`
}
