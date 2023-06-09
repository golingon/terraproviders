// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudfrontoriginrequestpolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CookiesConfig struct {
	// CookieBehavior: string, required
	CookieBehavior terra.StringValue `hcl:"cookie_behavior,attr" validate:"required"`
	// Cookies: optional
	Cookies *Cookies `hcl:"cookies,block"`
}

type Cookies struct {
	// Items: set of string, optional
	Items terra.SetValue[terra.StringValue] `hcl:"items,attr"`
}

type HeadersConfig struct {
	// HeaderBehavior: string, optional
	HeaderBehavior terra.StringValue `hcl:"header_behavior,attr"`
	// Headers: optional
	Headers *Headers `hcl:"headers,block"`
}

type Headers struct {
	// Items: set of string, optional
	Items terra.SetValue[terra.StringValue] `hcl:"items,attr"`
}

type QueryStringsConfig struct {
	// QueryStringBehavior: string, required
	QueryStringBehavior terra.StringValue `hcl:"query_string_behavior,attr" validate:"required"`
	// QueryStrings: optional
	QueryStrings *QueryStrings `hcl:"query_strings,block"`
}

type QueryStrings struct {
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

func (cc CookiesConfigAttributes) Cookies() terra.ListValue[CookiesAttributes] {
	return terra.ReferenceAsList[CookiesAttributes](cc.ref.Append("cookies"))
}

type CookiesAttributes struct {
	ref terra.Reference
}

func (c CookiesAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CookiesAttributes) InternalWithRef(ref terra.Reference) CookiesAttributes {
	return CookiesAttributes{ref: ref}
}

func (c CookiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CookiesAttributes) Items() terra.SetValue[terra.StringValue] {
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

func (hc HeadersConfigAttributes) Headers() terra.ListValue[HeadersAttributes] {
	return terra.ReferenceAsList[HeadersAttributes](hc.ref.Append("headers"))
}

type HeadersAttributes struct {
	ref terra.Reference
}

func (h HeadersAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h HeadersAttributes) InternalWithRef(ref terra.Reference) HeadersAttributes {
	return HeadersAttributes{ref: ref}
}

func (h HeadersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h HeadersAttributes) Items() terra.SetValue[terra.StringValue] {
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

func (qsc QueryStringsConfigAttributes) QueryStrings() terra.ListValue[QueryStringsAttributes] {
	return terra.ReferenceAsList[QueryStringsAttributes](qsc.ref.Append("query_strings"))
}

type QueryStringsAttributes struct {
	ref terra.Reference
}

func (qs QueryStringsAttributes) InternalRef() (terra.Reference, error) {
	return qs.ref, nil
}

func (qs QueryStringsAttributes) InternalWithRef(ref terra.Reference) QueryStringsAttributes {
	return QueryStringsAttributes{ref: ref}
}

func (qs QueryStringsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qs.ref.InternalTokens()
}

func (qs QueryStringsAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](qs.ref.Append("items"))
}

type CookiesConfigState struct {
	CookieBehavior string         `json:"cookie_behavior"`
	Cookies        []CookiesState `json:"cookies"`
}

type CookiesState struct {
	Items []string `json:"items"`
}

type HeadersConfigState struct {
	HeaderBehavior string         `json:"header_behavior"`
	Headers        []HeadersState `json:"headers"`
}

type HeadersState struct {
	Items []string `json:"items"`
}

type QueryStringsConfigState struct {
	QueryStringBehavior string              `json:"query_string_behavior"`
	QueryStrings        []QueryStringsState `json:"query_strings"`
}

type QueryStringsState struct {
	Items []string `json:"items"`
}
