// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkservicesedgecacheservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type LogConfig struct {
	// Enable: bool, optional
	Enable terra.BoolValue `hcl:"enable,attr"`
	// SampleRate: number, optional
	SampleRate terra.NumberValue `hcl:"sample_rate,attr"`
}

type Routing struct {
	// HostRule: min=1,max=10
	HostRule []HostRule `hcl:"host_rule,block" validate:"min=1,max=10"`
	// PathMatcher: min=1,max=10
	PathMatcher []PathMatcher `hcl:"path_matcher,block" validate:"min=1,max=10"`
}

type HostRule struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Hosts: list of string, required
	Hosts terra.ListValue[terra.StringValue] `hcl:"hosts,attr" validate:"required"`
	// PathMatcher: string, required
	PathMatcher terra.StringValue `hcl:"path_matcher,attr" validate:"required"`
}

type PathMatcher struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RouteRule: min=1,max=200
	RouteRule []RouteRule `hcl:"route_rule,block" validate:"min=1,max=200"`
}

type RouteRule struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Origin: string, optional
	Origin terra.StringValue `hcl:"origin,attr"`
	// Priority: string, required
	Priority terra.StringValue `hcl:"priority,attr" validate:"required"`
	// HeaderAction: optional
	HeaderAction *HeaderAction `hcl:"header_action,block"`
	// MatchRule: min=1,max=5
	MatchRule []MatchRule `hcl:"match_rule,block" validate:"min=1,max=5"`
	// RouteAction: optional
	RouteAction *RouteAction `hcl:"route_action,block"`
	// UrlRedirect: optional
	UrlRedirect *UrlRedirect `hcl:"url_redirect,block"`
}

type HeaderAction struct {
	// RequestHeaderToAdd: min=0,max=25
	RequestHeaderToAdd []RequestHeaderToAdd `hcl:"request_header_to_add,block" validate:"min=0,max=25"`
	// RequestHeaderToRemove: min=0,max=25
	RequestHeaderToRemove []RequestHeaderToRemove `hcl:"request_header_to_remove,block" validate:"min=0,max=25"`
	// ResponseHeaderToAdd: min=0,max=25
	ResponseHeaderToAdd []ResponseHeaderToAdd `hcl:"response_header_to_add,block" validate:"min=0,max=25"`
	// ResponseHeaderToRemove: min=0,max=25
	ResponseHeaderToRemove []ResponseHeaderToRemove `hcl:"response_header_to_remove,block" validate:"min=0,max=25"`
}

type RequestHeaderToAdd struct {
	// HeaderName: string, required
	HeaderName terra.StringValue `hcl:"header_name,attr" validate:"required"`
	// HeaderValue: string, required
	HeaderValue terra.StringValue `hcl:"header_value,attr" validate:"required"`
	// Replace: bool, optional
	Replace terra.BoolValue `hcl:"replace,attr"`
}

type RequestHeaderToRemove struct {
	// HeaderName: string, required
	HeaderName terra.StringValue `hcl:"header_name,attr" validate:"required"`
}

type ResponseHeaderToAdd struct {
	// HeaderName: string, required
	HeaderName terra.StringValue `hcl:"header_name,attr" validate:"required"`
	// HeaderValue: string, required
	HeaderValue terra.StringValue `hcl:"header_value,attr" validate:"required"`
	// Replace: bool, optional
	Replace terra.BoolValue `hcl:"replace,attr"`
}

type ResponseHeaderToRemove struct {
	// HeaderName: string, required
	HeaderName terra.StringValue `hcl:"header_name,attr" validate:"required"`
}

type MatchRule struct {
	// FullPathMatch: string, optional
	FullPathMatch terra.StringValue `hcl:"full_path_match,attr"`
	// IgnoreCase: bool, optional
	IgnoreCase terra.BoolValue `hcl:"ignore_case,attr"`
	// PathTemplateMatch: string, optional
	PathTemplateMatch terra.StringValue `hcl:"path_template_match,attr"`
	// PrefixMatch: string, optional
	PrefixMatch terra.StringValue `hcl:"prefix_match,attr"`
	// HeaderMatch: min=0,max=3
	HeaderMatch []HeaderMatch `hcl:"header_match,block" validate:"min=0,max=3"`
	// QueryParameterMatch: min=0,max=5
	QueryParameterMatch []QueryParameterMatch `hcl:"query_parameter_match,block" validate:"min=0,max=5"`
}

type HeaderMatch struct {
	// ExactMatch: string, optional
	ExactMatch terra.StringValue `hcl:"exact_match,attr"`
	// HeaderName: string, required
	HeaderName terra.StringValue `hcl:"header_name,attr" validate:"required"`
	// InvertMatch: bool, optional
	InvertMatch terra.BoolValue `hcl:"invert_match,attr"`
	// PrefixMatch: string, optional
	PrefixMatch terra.StringValue `hcl:"prefix_match,attr"`
	// PresentMatch: bool, optional
	PresentMatch terra.BoolValue `hcl:"present_match,attr"`
	// SuffixMatch: string, optional
	SuffixMatch terra.StringValue `hcl:"suffix_match,attr"`
}

type QueryParameterMatch struct {
	// ExactMatch: string, optional
	ExactMatch terra.StringValue `hcl:"exact_match,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PresentMatch: bool, optional
	PresentMatch terra.BoolValue `hcl:"present_match,attr"`
}

type RouteAction struct {
	// CdnPolicy: optional
	CdnPolicy *CdnPolicy `hcl:"cdn_policy,block"`
	// CorsPolicy: optional
	CorsPolicy *CorsPolicy `hcl:"cors_policy,block"`
	// UrlRewrite: optional
	UrlRewrite *UrlRewrite `hcl:"url_rewrite,block"`
}

type CdnPolicy struct {
	// CacheMode: string, optional
	CacheMode terra.StringValue `hcl:"cache_mode,attr"`
	// ClientTtl: string, optional
	ClientTtl terra.StringValue `hcl:"client_ttl,attr"`
	// DefaultTtl: string, optional
	DefaultTtl terra.StringValue `hcl:"default_ttl,attr"`
	// MaxTtl: string, optional
	MaxTtl terra.StringValue `hcl:"max_ttl,attr"`
	// NegativeCaching: bool, optional
	NegativeCaching terra.BoolValue `hcl:"negative_caching,attr"`
	// NegativeCachingPolicy: map of string, optional
	NegativeCachingPolicy terra.MapValue[terra.StringValue] `hcl:"negative_caching_policy,attr"`
	// SignedRequestKeyset: string, optional
	SignedRequestKeyset terra.StringValue `hcl:"signed_request_keyset,attr"`
	// SignedRequestMaximumExpirationTtl: string, optional
	SignedRequestMaximumExpirationTtl terra.StringValue `hcl:"signed_request_maximum_expiration_ttl,attr"`
	// SignedRequestMode: string, optional
	SignedRequestMode terra.StringValue `hcl:"signed_request_mode,attr"`
	// AddSignatures: optional
	AddSignatures *AddSignatures `hcl:"add_signatures,block"`
	// CacheKeyPolicy: optional
	CacheKeyPolicy *CacheKeyPolicy `hcl:"cache_key_policy,block"`
	// SignedTokenOptions: optional
	SignedTokenOptions *SignedTokenOptions `hcl:"signed_token_options,block"`
}

type AddSignatures struct {
	// Actions: list of string, required
	Actions terra.ListValue[terra.StringValue] `hcl:"actions,attr" validate:"required"`
	// CopiedParameters: list of string, optional
	CopiedParameters terra.ListValue[terra.StringValue] `hcl:"copied_parameters,attr"`
	// Keyset: string, optional
	Keyset terra.StringValue `hcl:"keyset,attr"`
	// TokenQueryParameter: string, optional
	TokenQueryParameter terra.StringValue `hcl:"token_query_parameter,attr"`
	// TokenTtl: string, optional
	TokenTtl terra.StringValue `hcl:"token_ttl,attr"`
}

type CacheKeyPolicy struct {
	// ExcludeHost: bool, optional
	ExcludeHost terra.BoolValue `hcl:"exclude_host,attr"`
	// ExcludeQueryString: bool, optional
	ExcludeQueryString terra.BoolValue `hcl:"exclude_query_string,attr"`
	// ExcludedQueryParameters: list of string, optional
	ExcludedQueryParameters terra.ListValue[terra.StringValue] `hcl:"excluded_query_parameters,attr"`
	// IncludeProtocol: bool, optional
	IncludeProtocol terra.BoolValue `hcl:"include_protocol,attr"`
	// IncludedCookieNames: list of string, optional
	IncludedCookieNames terra.ListValue[terra.StringValue] `hcl:"included_cookie_names,attr"`
	// IncludedHeaderNames: list of string, optional
	IncludedHeaderNames terra.ListValue[terra.StringValue] `hcl:"included_header_names,attr"`
	// IncludedQueryParameters: list of string, optional
	IncludedQueryParameters terra.ListValue[terra.StringValue] `hcl:"included_query_parameters,attr"`
}

type SignedTokenOptions struct {
	// AllowedSignatureAlgorithms: list of string, optional
	AllowedSignatureAlgorithms terra.ListValue[terra.StringValue] `hcl:"allowed_signature_algorithms,attr"`
	// TokenQueryParameter: string, optional
	TokenQueryParameter terra.StringValue `hcl:"token_query_parameter,attr"`
}

type CorsPolicy struct {
	// AllowCredentials: bool, optional
	AllowCredentials terra.BoolValue `hcl:"allow_credentials,attr"`
	// AllowHeaders: list of string, optional
	AllowHeaders terra.ListValue[terra.StringValue] `hcl:"allow_headers,attr"`
	// AllowMethods: list of string, optional
	AllowMethods terra.ListValue[terra.StringValue] `hcl:"allow_methods,attr"`
	// AllowOrigins: list of string, optional
	AllowOrigins terra.ListValue[terra.StringValue] `hcl:"allow_origins,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// ExposeHeaders: list of string, optional
	ExposeHeaders terra.ListValue[terra.StringValue] `hcl:"expose_headers,attr"`
	// MaxAge: string, required
	MaxAge terra.StringValue `hcl:"max_age,attr" validate:"required"`
}

type UrlRewrite struct {
	// HostRewrite: string, optional
	HostRewrite terra.StringValue `hcl:"host_rewrite,attr"`
	// PathPrefixRewrite: string, optional
	PathPrefixRewrite terra.StringValue `hcl:"path_prefix_rewrite,attr"`
	// PathTemplateRewrite: string, optional
	PathTemplateRewrite terra.StringValue `hcl:"path_template_rewrite,attr"`
}

type UrlRedirect struct {
	// HostRedirect: string, optional
	HostRedirect terra.StringValue `hcl:"host_redirect,attr"`
	// HttpsRedirect: bool, optional
	HttpsRedirect terra.BoolValue `hcl:"https_redirect,attr"`
	// PathRedirect: string, optional
	PathRedirect terra.StringValue `hcl:"path_redirect,attr"`
	// PrefixRedirect: string, optional
	PrefixRedirect terra.StringValue `hcl:"prefix_redirect,attr"`
	// RedirectResponseCode: string, optional
	RedirectResponseCode terra.StringValue `hcl:"redirect_response_code,attr"`
	// StripQuery: bool, optional
	StripQuery terra.BoolValue `hcl:"strip_query,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type LogConfigAttributes struct {
	ref terra.Reference
}

func (lc LogConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc LogConfigAttributes) InternalWithRef(ref terra.Reference) LogConfigAttributes {
	return LogConfigAttributes{ref: ref}
}

func (lc LogConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc LogConfigAttributes) Enable() terra.BoolValue {
	return terra.ReferenceAsBool(lc.ref.Append("enable"))
}

func (lc LogConfigAttributes) SampleRate() terra.NumberValue {
	return terra.ReferenceAsNumber(lc.ref.Append("sample_rate"))
}

type RoutingAttributes struct {
	ref terra.Reference
}

func (r RoutingAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RoutingAttributes) InternalWithRef(ref terra.Reference) RoutingAttributes {
	return RoutingAttributes{ref: ref}
}

func (r RoutingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RoutingAttributes) HostRule() terra.ListValue[HostRuleAttributes] {
	return terra.ReferenceAsList[HostRuleAttributes](r.ref.Append("host_rule"))
}

func (r RoutingAttributes) PathMatcher() terra.ListValue[PathMatcherAttributes] {
	return terra.ReferenceAsList[PathMatcherAttributes](r.ref.Append("path_matcher"))
}

type HostRuleAttributes struct {
	ref terra.Reference
}

func (hr HostRuleAttributes) InternalRef() (terra.Reference, error) {
	return hr.ref, nil
}

func (hr HostRuleAttributes) InternalWithRef(ref terra.Reference) HostRuleAttributes {
	return HostRuleAttributes{ref: ref}
}

func (hr HostRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hr.ref.InternalTokens()
}

func (hr HostRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(hr.ref.Append("description"))
}

func (hr HostRuleAttributes) Hosts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](hr.ref.Append("hosts"))
}

func (hr HostRuleAttributes) PathMatcher() terra.StringValue {
	return terra.ReferenceAsString(hr.ref.Append("path_matcher"))
}

type PathMatcherAttributes struct {
	ref terra.Reference
}

func (pm PathMatcherAttributes) InternalRef() (terra.Reference, error) {
	return pm.ref, nil
}

func (pm PathMatcherAttributes) InternalWithRef(ref terra.Reference) PathMatcherAttributes {
	return PathMatcherAttributes{ref: ref}
}

func (pm PathMatcherAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pm.ref.InternalTokens()
}

func (pm PathMatcherAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(pm.ref.Append("description"))
}

func (pm PathMatcherAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pm.ref.Append("name"))
}

func (pm PathMatcherAttributes) RouteRule() terra.ListValue[RouteRuleAttributes] {
	return terra.ReferenceAsList[RouteRuleAttributes](pm.ref.Append("route_rule"))
}

type RouteRuleAttributes struct {
	ref terra.Reference
}

func (rr RouteRuleAttributes) InternalRef() (terra.Reference, error) {
	return rr.ref, nil
}

func (rr RouteRuleAttributes) InternalWithRef(ref terra.Reference) RouteRuleAttributes {
	return RouteRuleAttributes{ref: ref}
}

func (rr RouteRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rr.ref.InternalTokens()
}

func (rr RouteRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("description"))
}

func (rr RouteRuleAttributes) Origin() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("origin"))
}

func (rr RouteRuleAttributes) Priority() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("priority"))
}

func (rr RouteRuleAttributes) HeaderAction() terra.ListValue[HeaderActionAttributes] {
	return terra.ReferenceAsList[HeaderActionAttributes](rr.ref.Append("header_action"))
}

func (rr RouteRuleAttributes) MatchRule() terra.ListValue[MatchRuleAttributes] {
	return terra.ReferenceAsList[MatchRuleAttributes](rr.ref.Append("match_rule"))
}

func (rr RouteRuleAttributes) RouteAction() terra.ListValue[RouteActionAttributes] {
	return terra.ReferenceAsList[RouteActionAttributes](rr.ref.Append("route_action"))
}

func (rr RouteRuleAttributes) UrlRedirect() terra.ListValue[UrlRedirectAttributes] {
	return terra.ReferenceAsList[UrlRedirectAttributes](rr.ref.Append("url_redirect"))
}

type HeaderActionAttributes struct {
	ref terra.Reference
}

func (ha HeaderActionAttributes) InternalRef() (terra.Reference, error) {
	return ha.ref, nil
}

func (ha HeaderActionAttributes) InternalWithRef(ref terra.Reference) HeaderActionAttributes {
	return HeaderActionAttributes{ref: ref}
}

func (ha HeaderActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ha.ref.InternalTokens()
}

func (ha HeaderActionAttributes) RequestHeaderToAdd() terra.ListValue[RequestHeaderToAddAttributes] {
	return terra.ReferenceAsList[RequestHeaderToAddAttributes](ha.ref.Append("request_header_to_add"))
}

func (ha HeaderActionAttributes) RequestHeaderToRemove() terra.ListValue[RequestHeaderToRemoveAttributes] {
	return terra.ReferenceAsList[RequestHeaderToRemoveAttributes](ha.ref.Append("request_header_to_remove"))
}

func (ha HeaderActionAttributes) ResponseHeaderToAdd() terra.ListValue[ResponseHeaderToAddAttributes] {
	return terra.ReferenceAsList[ResponseHeaderToAddAttributes](ha.ref.Append("response_header_to_add"))
}

func (ha HeaderActionAttributes) ResponseHeaderToRemove() terra.ListValue[ResponseHeaderToRemoveAttributes] {
	return terra.ReferenceAsList[ResponseHeaderToRemoveAttributes](ha.ref.Append("response_header_to_remove"))
}

type RequestHeaderToAddAttributes struct {
	ref terra.Reference
}

func (rhta RequestHeaderToAddAttributes) InternalRef() (terra.Reference, error) {
	return rhta.ref, nil
}

func (rhta RequestHeaderToAddAttributes) InternalWithRef(ref terra.Reference) RequestHeaderToAddAttributes {
	return RequestHeaderToAddAttributes{ref: ref}
}

func (rhta RequestHeaderToAddAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rhta.ref.InternalTokens()
}

func (rhta RequestHeaderToAddAttributes) HeaderName() terra.StringValue {
	return terra.ReferenceAsString(rhta.ref.Append("header_name"))
}

func (rhta RequestHeaderToAddAttributes) HeaderValue() terra.StringValue {
	return terra.ReferenceAsString(rhta.ref.Append("header_value"))
}

func (rhta RequestHeaderToAddAttributes) Replace() terra.BoolValue {
	return terra.ReferenceAsBool(rhta.ref.Append("replace"))
}

type RequestHeaderToRemoveAttributes struct {
	ref terra.Reference
}

func (rhtr RequestHeaderToRemoveAttributes) InternalRef() (terra.Reference, error) {
	return rhtr.ref, nil
}

func (rhtr RequestHeaderToRemoveAttributes) InternalWithRef(ref terra.Reference) RequestHeaderToRemoveAttributes {
	return RequestHeaderToRemoveAttributes{ref: ref}
}

func (rhtr RequestHeaderToRemoveAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rhtr.ref.InternalTokens()
}

func (rhtr RequestHeaderToRemoveAttributes) HeaderName() terra.StringValue {
	return terra.ReferenceAsString(rhtr.ref.Append("header_name"))
}

type ResponseHeaderToAddAttributes struct {
	ref terra.Reference
}

func (rhta ResponseHeaderToAddAttributes) InternalRef() (terra.Reference, error) {
	return rhta.ref, nil
}

func (rhta ResponseHeaderToAddAttributes) InternalWithRef(ref terra.Reference) ResponseHeaderToAddAttributes {
	return ResponseHeaderToAddAttributes{ref: ref}
}

func (rhta ResponseHeaderToAddAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rhta.ref.InternalTokens()
}

func (rhta ResponseHeaderToAddAttributes) HeaderName() terra.StringValue {
	return terra.ReferenceAsString(rhta.ref.Append("header_name"))
}

func (rhta ResponseHeaderToAddAttributes) HeaderValue() terra.StringValue {
	return terra.ReferenceAsString(rhta.ref.Append("header_value"))
}

func (rhta ResponseHeaderToAddAttributes) Replace() terra.BoolValue {
	return terra.ReferenceAsBool(rhta.ref.Append("replace"))
}

type ResponseHeaderToRemoveAttributes struct {
	ref terra.Reference
}

func (rhtr ResponseHeaderToRemoveAttributes) InternalRef() (terra.Reference, error) {
	return rhtr.ref, nil
}

func (rhtr ResponseHeaderToRemoveAttributes) InternalWithRef(ref terra.Reference) ResponseHeaderToRemoveAttributes {
	return ResponseHeaderToRemoveAttributes{ref: ref}
}

func (rhtr ResponseHeaderToRemoveAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rhtr.ref.InternalTokens()
}

func (rhtr ResponseHeaderToRemoveAttributes) HeaderName() terra.StringValue {
	return terra.ReferenceAsString(rhtr.ref.Append("header_name"))
}

type MatchRuleAttributes struct {
	ref terra.Reference
}

func (mr MatchRuleAttributes) InternalRef() (terra.Reference, error) {
	return mr.ref, nil
}

func (mr MatchRuleAttributes) InternalWithRef(ref terra.Reference) MatchRuleAttributes {
	return MatchRuleAttributes{ref: ref}
}

func (mr MatchRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mr.ref.InternalTokens()
}

func (mr MatchRuleAttributes) FullPathMatch() terra.StringValue {
	return terra.ReferenceAsString(mr.ref.Append("full_path_match"))
}

func (mr MatchRuleAttributes) IgnoreCase() terra.BoolValue {
	return terra.ReferenceAsBool(mr.ref.Append("ignore_case"))
}

func (mr MatchRuleAttributes) PathTemplateMatch() terra.StringValue {
	return terra.ReferenceAsString(mr.ref.Append("path_template_match"))
}

func (mr MatchRuleAttributes) PrefixMatch() terra.StringValue {
	return terra.ReferenceAsString(mr.ref.Append("prefix_match"))
}

func (mr MatchRuleAttributes) HeaderMatch() terra.ListValue[HeaderMatchAttributes] {
	return terra.ReferenceAsList[HeaderMatchAttributes](mr.ref.Append("header_match"))
}

func (mr MatchRuleAttributes) QueryParameterMatch() terra.ListValue[QueryParameterMatchAttributes] {
	return terra.ReferenceAsList[QueryParameterMatchAttributes](mr.ref.Append("query_parameter_match"))
}

type HeaderMatchAttributes struct {
	ref terra.Reference
}

func (hm HeaderMatchAttributes) InternalRef() (terra.Reference, error) {
	return hm.ref, nil
}

func (hm HeaderMatchAttributes) InternalWithRef(ref terra.Reference) HeaderMatchAttributes {
	return HeaderMatchAttributes{ref: ref}
}

func (hm HeaderMatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hm.ref.InternalTokens()
}

func (hm HeaderMatchAttributes) ExactMatch() terra.StringValue {
	return terra.ReferenceAsString(hm.ref.Append("exact_match"))
}

func (hm HeaderMatchAttributes) HeaderName() terra.StringValue {
	return terra.ReferenceAsString(hm.ref.Append("header_name"))
}

func (hm HeaderMatchAttributes) InvertMatch() terra.BoolValue {
	return terra.ReferenceAsBool(hm.ref.Append("invert_match"))
}

func (hm HeaderMatchAttributes) PrefixMatch() terra.StringValue {
	return terra.ReferenceAsString(hm.ref.Append("prefix_match"))
}

func (hm HeaderMatchAttributes) PresentMatch() terra.BoolValue {
	return terra.ReferenceAsBool(hm.ref.Append("present_match"))
}

func (hm HeaderMatchAttributes) SuffixMatch() terra.StringValue {
	return terra.ReferenceAsString(hm.ref.Append("suffix_match"))
}

type QueryParameterMatchAttributes struct {
	ref terra.Reference
}

func (qpm QueryParameterMatchAttributes) InternalRef() (terra.Reference, error) {
	return qpm.ref, nil
}

func (qpm QueryParameterMatchAttributes) InternalWithRef(ref terra.Reference) QueryParameterMatchAttributes {
	return QueryParameterMatchAttributes{ref: ref}
}

func (qpm QueryParameterMatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qpm.ref.InternalTokens()
}

func (qpm QueryParameterMatchAttributes) ExactMatch() terra.StringValue {
	return terra.ReferenceAsString(qpm.ref.Append("exact_match"))
}

func (qpm QueryParameterMatchAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(qpm.ref.Append("name"))
}

func (qpm QueryParameterMatchAttributes) PresentMatch() terra.BoolValue {
	return terra.ReferenceAsBool(qpm.ref.Append("present_match"))
}

type RouteActionAttributes struct {
	ref terra.Reference
}

func (ra RouteActionAttributes) InternalRef() (terra.Reference, error) {
	return ra.ref, nil
}

func (ra RouteActionAttributes) InternalWithRef(ref terra.Reference) RouteActionAttributes {
	return RouteActionAttributes{ref: ref}
}

func (ra RouteActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ra.ref.InternalTokens()
}

func (ra RouteActionAttributes) CdnPolicy() terra.ListValue[CdnPolicyAttributes] {
	return terra.ReferenceAsList[CdnPolicyAttributes](ra.ref.Append("cdn_policy"))
}

func (ra RouteActionAttributes) CorsPolicy() terra.ListValue[CorsPolicyAttributes] {
	return terra.ReferenceAsList[CorsPolicyAttributes](ra.ref.Append("cors_policy"))
}

func (ra RouteActionAttributes) UrlRewrite() terra.ListValue[UrlRewriteAttributes] {
	return terra.ReferenceAsList[UrlRewriteAttributes](ra.ref.Append("url_rewrite"))
}

type CdnPolicyAttributes struct {
	ref terra.Reference
}

func (cp CdnPolicyAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp CdnPolicyAttributes) InternalWithRef(ref terra.Reference) CdnPolicyAttributes {
	return CdnPolicyAttributes{ref: ref}
}

func (cp CdnPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cp.ref.InternalTokens()
}

func (cp CdnPolicyAttributes) CacheMode() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("cache_mode"))
}

func (cp CdnPolicyAttributes) ClientTtl() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("client_ttl"))
}

func (cp CdnPolicyAttributes) DefaultTtl() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("default_ttl"))
}

func (cp CdnPolicyAttributes) MaxTtl() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("max_ttl"))
}

func (cp CdnPolicyAttributes) NegativeCaching() terra.BoolValue {
	return terra.ReferenceAsBool(cp.ref.Append("negative_caching"))
}

func (cp CdnPolicyAttributes) NegativeCachingPolicy() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cp.ref.Append("negative_caching_policy"))
}

func (cp CdnPolicyAttributes) SignedRequestKeyset() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("signed_request_keyset"))
}

func (cp CdnPolicyAttributes) SignedRequestMaximumExpirationTtl() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("signed_request_maximum_expiration_ttl"))
}

func (cp CdnPolicyAttributes) SignedRequestMode() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("signed_request_mode"))
}

func (cp CdnPolicyAttributes) AddSignatures() terra.ListValue[AddSignaturesAttributes] {
	return terra.ReferenceAsList[AddSignaturesAttributes](cp.ref.Append("add_signatures"))
}

func (cp CdnPolicyAttributes) CacheKeyPolicy() terra.ListValue[CacheKeyPolicyAttributes] {
	return terra.ReferenceAsList[CacheKeyPolicyAttributes](cp.ref.Append("cache_key_policy"))
}

func (cp CdnPolicyAttributes) SignedTokenOptions() terra.ListValue[SignedTokenOptionsAttributes] {
	return terra.ReferenceAsList[SignedTokenOptionsAttributes](cp.ref.Append("signed_token_options"))
}

type AddSignaturesAttributes struct {
	ref terra.Reference
}

func (as AddSignaturesAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as AddSignaturesAttributes) InternalWithRef(ref terra.Reference) AddSignaturesAttributes {
	return AddSignaturesAttributes{ref: ref}
}

func (as AddSignaturesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as AddSignaturesAttributes) Actions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](as.ref.Append("actions"))
}

func (as AddSignaturesAttributes) CopiedParameters() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](as.ref.Append("copied_parameters"))
}

func (as AddSignaturesAttributes) Keyset() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("keyset"))
}

func (as AddSignaturesAttributes) TokenQueryParameter() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("token_query_parameter"))
}

func (as AddSignaturesAttributes) TokenTtl() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("token_ttl"))
}

type CacheKeyPolicyAttributes struct {
	ref terra.Reference
}

func (ckp CacheKeyPolicyAttributes) InternalRef() (terra.Reference, error) {
	return ckp.ref, nil
}

func (ckp CacheKeyPolicyAttributes) InternalWithRef(ref terra.Reference) CacheKeyPolicyAttributes {
	return CacheKeyPolicyAttributes{ref: ref}
}

func (ckp CacheKeyPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ckp.ref.InternalTokens()
}

func (ckp CacheKeyPolicyAttributes) ExcludeHost() terra.BoolValue {
	return terra.ReferenceAsBool(ckp.ref.Append("exclude_host"))
}

func (ckp CacheKeyPolicyAttributes) ExcludeQueryString() terra.BoolValue {
	return terra.ReferenceAsBool(ckp.ref.Append("exclude_query_string"))
}

func (ckp CacheKeyPolicyAttributes) ExcludedQueryParameters() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ckp.ref.Append("excluded_query_parameters"))
}

func (ckp CacheKeyPolicyAttributes) IncludeProtocol() terra.BoolValue {
	return terra.ReferenceAsBool(ckp.ref.Append("include_protocol"))
}

func (ckp CacheKeyPolicyAttributes) IncludedCookieNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ckp.ref.Append("included_cookie_names"))
}

func (ckp CacheKeyPolicyAttributes) IncludedHeaderNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ckp.ref.Append("included_header_names"))
}

func (ckp CacheKeyPolicyAttributes) IncludedQueryParameters() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ckp.ref.Append("included_query_parameters"))
}

type SignedTokenOptionsAttributes struct {
	ref terra.Reference
}

func (sto SignedTokenOptionsAttributes) InternalRef() (terra.Reference, error) {
	return sto.ref, nil
}

func (sto SignedTokenOptionsAttributes) InternalWithRef(ref terra.Reference) SignedTokenOptionsAttributes {
	return SignedTokenOptionsAttributes{ref: ref}
}

func (sto SignedTokenOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sto.ref.InternalTokens()
}

func (sto SignedTokenOptionsAttributes) AllowedSignatureAlgorithms() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sto.ref.Append("allowed_signature_algorithms"))
}

func (sto SignedTokenOptionsAttributes) TokenQueryParameter() terra.StringValue {
	return terra.ReferenceAsString(sto.ref.Append("token_query_parameter"))
}

type CorsPolicyAttributes struct {
	ref terra.Reference
}

func (cp CorsPolicyAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp CorsPolicyAttributes) InternalWithRef(ref terra.Reference) CorsPolicyAttributes {
	return CorsPolicyAttributes{ref: ref}
}

func (cp CorsPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cp.ref.InternalTokens()
}

func (cp CorsPolicyAttributes) AllowCredentials() terra.BoolValue {
	return terra.ReferenceAsBool(cp.ref.Append("allow_credentials"))
}

func (cp CorsPolicyAttributes) AllowHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cp.ref.Append("allow_headers"))
}

func (cp CorsPolicyAttributes) AllowMethods() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cp.ref.Append("allow_methods"))
}

func (cp CorsPolicyAttributes) AllowOrigins() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cp.ref.Append("allow_origins"))
}

func (cp CorsPolicyAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(cp.ref.Append("disabled"))
}

func (cp CorsPolicyAttributes) ExposeHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cp.ref.Append("expose_headers"))
}

func (cp CorsPolicyAttributes) MaxAge() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("max_age"))
}

type UrlRewriteAttributes struct {
	ref terra.Reference
}

func (ur UrlRewriteAttributes) InternalRef() (terra.Reference, error) {
	return ur.ref, nil
}

func (ur UrlRewriteAttributes) InternalWithRef(ref terra.Reference) UrlRewriteAttributes {
	return UrlRewriteAttributes{ref: ref}
}

func (ur UrlRewriteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ur.ref.InternalTokens()
}

func (ur UrlRewriteAttributes) HostRewrite() terra.StringValue {
	return terra.ReferenceAsString(ur.ref.Append("host_rewrite"))
}

func (ur UrlRewriteAttributes) PathPrefixRewrite() terra.StringValue {
	return terra.ReferenceAsString(ur.ref.Append("path_prefix_rewrite"))
}

func (ur UrlRewriteAttributes) PathTemplateRewrite() terra.StringValue {
	return terra.ReferenceAsString(ur.ref.Append("path_template_rewrite"))
}

type UrlRedirectAttributes struct {
	ref terra.Reference
}

func (ur UrlRedirectAttributes) InternalRef() (terra.Reference, error) {
	return ur.ref, nil
}

func (ur UrlRedirectAttributes) InternalWithRef(ref terra.Reference) UrlRedirectAttributes {
	return UrlRedirectAttributes{ref: ref}
}

func (ur UrlRedirectAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ur.ref.InternalTokens()
}

func (ur UrlRedirectAttributes) HostRedirect() terra.StringValue {
	return terra.ReferenceAsString(ur.ref.Append("host_redirect"))
}

func (ur UrlRedirectAttributes) HttpsRedirect() terra.BoolValue {
	return terra.ReferenceAsBool(ur.ref.Append("https_redirect"))
}

func (ur UrlRedirectAttributes) PathRedirect() terra.StringValue {
	return terra.ReferenceAsString(ur.ref.Append("path_redirect"))
}

func (ur UrlRedirectAttributes) PrefixRedirect() terra.StringValue {
	return terra.ReferenceAsString(ur.ref.Append("prefix_redirect"))
}

func (ur UrlRedirectAttributes) RedirectResponseCode() terra.StringValue {
	return terra.ReferenceAsString(ur.ref.Append("redirect_response_code"))
}

func (ur UrlRedirectAttributes) StripQuery() terra.BoolValue {
	return terra.ReferenceAsBool(ur.ref.Append("strip_query"))
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

type LogConfigState struct {
	Enable     bool    `json:"enable"`
	SampleRate float64 `json:"sample_rate"`
}

type RoutingState struct {
	HostRule    []HostRuleState    `json:"host_rule"`
	PathMatcher []PathMatcherState `json:"path_matcher"`
}

type HostRuleState struct {
	Description string   `json:"description"`
	Hosts       []string `json:"hosts"`
	PathMatcher string   `json:"path_matcher"`
}

type PathMatcherState struct {
	Description string           `json:"description"`
	Name        string           `json:"name"`
	RouteRule   []RouteRuleState `json:"route_rule"`
}

type RouteRuleState struct {
	Description  string              `json:"description"`
	Origin       string              `json:"origin"`
	Priority     string              `json:"priority"`
	HeaderAction []HeaderActionState `json:"header_action"`
	MatchRule    []MatchRuleState    `json:"match_rule"`
	RouteAction  []RouteActionState  `json:"route_action"`
	UrlRedirect  []UrlRedirectState  `json:"url_redirect"`
}

type HeaderActionState struct {
	RequestHeaderToAdd     []RequestHeaderToAddState     `json:"request_header_to_add"`
	RequestHeaderToRemove  []RequestHeaderToRemoveState  `json:"request_header_to_remove"`
	ResponseHeaderToAdd    []ResponseHeaderToAddState    `json:"response_header_to_add"`
	ResponseHeaderToRemove []ResponseHeaderToRemoveState `json:"response_header_to_remove"`
}

type RequestHeaderToAddState struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
	Replace     bool   `json:"replace"`
}

type RequestHeaderToRemoveState struct {
	HeaderName string `json:"header_name"`
}

type ResponseHeaderToAddState struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
	Replace     bool   `json:"replace"`
}

type ResponseHeaderToRemoveState struct {
	HeaderName string `json:"header_name"`
}

type MatchRuleState struct {
	FullPathMatch       string                     `json:"full_path_match"`
	IgnoreCase          bool                       `json:"ignore_case"`
	PathTemplateMatch   string                     `json:"path_template_match"`
	PrefixMatch         string                     `json:"prefix_match"`
	HeaderMatch         []HeaderMatchState         `json:"header_match"`
	QueryParameterMatch []QueryParameterMatchState `json:"query_parameter_match"`
}

type HeaderMatchState struct {
	ExactMatch   string `json:"exact_match"`
	HeaderName   string `json:"header_name"`
	InvertMatch  bool   `json:"invert_match"`
	PrefixMatch  string `json:"prefix_match"`
	PresentMatch bool   `json:"present_match"`
	SuffixMatch  string `json:"suffix_match"`
}

type QueryParameterMatchState struct {
	ExactMatch   string `json:"exact_match"`
	Name         string `json:"name"`
	PresentMatch bool   `json:"present_match"`
}

type RouteActionState struct {
	CdnPolicy  []CdnPolicyState  `json:"cdn_policy"`
	CorsPolicy []CorsPolicyState `json:"cors_policy"`
	UrlRewrite []UrlRewriteState `json:"url_rewrite"`
}

type CdnPolicyState struct {
	CacheMode                         string                    `json:"cache_mode"`
	ClientTtl                         string                    `json:"client_ttl"`
	DefaultTtl                        string                    `json:"default_ttl"`
	MaxTtl                            string                    `json:"max_ttl"`
	NegativeCaching                   bool                      `json:"negative_caching"`
	NegativeCachingPolicy             map[string]string         `json:"negative_caching_policy"`
	SignedRequestKeyset               string                    `json:"signed_request_keyset"`
	SignedRequestMaximumExpirationTtl string                    `json:"signed_request_maximum_expiration_ttl"`
	SignedRequestMode                 string                    `json:"signed_request_mode"`
	AddSignatures                     []AddSignaturesState      `json:"add_signatures"`
	CacheKeyPolicy                    []CacheKeyPolicyState     `json:"cache_key_policy"`
	SignedTokenOptions                []SignedTokenOptionsState `json:"signed_token_options"`
}

type AddSignaturesState struct {
	Actions             []string `json:"actions"`
	CopiedParameters    []string `json:"copied_parameters"`
	Keyset              string   `json:"keyset"`
	TokenQueryParameter string   `json:"token_query_parameter"`
	TokenTtl            string   `json:"token_ttl"`
}

type CacheKeyPolicyState struct {
	ExcludeHost             bool     `json:"exclude_host"`
	ExcludeQueryString      bool     `json:"exclude_query_string"`
	ExcludedQueryParameters []string `json:"excluded_query_parameters"`
	IncludeProtocol         bool     `json:"include_protocol"`
	IncludedCookieNames     []string `json:"included_cookie_names"`
	IncludedHeaderNames     []string `json:"included_header_names"`
	IncludedQueryParameters []string `json:"included_query_parameters"`
}

type SignedTokenOptionsState struct {
	AllowedSignatureAlgorithms []string `json:"allowed_signature_algorithms"`
	TokenQueryParameter        string   `json:"token_query_parameter"`
}

type CorsPolicyState struct {
	AllowCredentials bool     `json:"allow_credentials"`
	AllowHeaders     []string `json:"allow_headers"`
	AllowMethods     []string `json:"allow_methods"`
	AllowOrigins     []string `json:"allow_origins"`
	Disabled         bool     `json:"disabled"`
	ExposeHeaders    []string `json:"expose_headers"`
	MaxAge           string   `json:"max_age"`
}

type UrlRewriteState struct {
	HostRewrite         string `json:"host_rewrite"`
	PathPrefixRewrite   string `json:"path_prefix_rewrite"`
	PathTemplateRewrite string `json:"path_template_rewrite"`
}

type UrlRedirectState struct {
	HostRedirect         string `json:"host_redirect"`
	HttpsRedirect        bool   `json:"https_redirect"`
	PathRedirect         string `json:"path_redirect"`
	PrefixRedirect       string `json:"prefix_redirect"`
	RedirectResponseCode string `json:"redirect_response_code"`
	StripQuery           bool   `json:"strip_query"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
