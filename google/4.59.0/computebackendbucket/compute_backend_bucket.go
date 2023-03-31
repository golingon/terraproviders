// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package computebackendbucket

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CdnPolicy struct {
	// CacheMode: string, optional
	CacheMode terra.StringValue `hcl:"cache_mode,attr"`
	// ClientTtl: number, optional
	ClientTtl terra.NumberValue `hcl:"client_ttl,attr"`
	// DefaultTtl: number, optional
	DefaultTtl terra.NumberValue `hcl:"default_ttl,attr"`
	// MaxTtl: number, optional
	MaxTtl terra.NumberValue `hcl:"max_ttl,attr"`
	// NegativeCaching: bool, optional
	NegativeCaching terra.BoolValue `hcl:"negative_caching,attr"`
	// RequestCoalescing: bool, optional
	RequestCoalescing terra.BoolValue `hcl:"request_coalescing,attr"`
	// ServeWhileStale: number, optional
	ServeWhileStale terra.NumberValue `hcl:"serve_while_stale,attr"`
	// SignedUrlCacheMaxAgeSec: number, optional
	SignedUrlCacheMaxAgeSec terra.NumberValue `hcl:"signed_url_cache_max_age_sec,attr"`
	// BypassCacheOnRequestHeaders: min=0,max=5
	BypassCacheOnRequestHeaders []BypassCacheOnRequestHeaders `hcl:"bypass_cache_on_request_headers,block" validate:"min=0,max=5"`
	// CacheKeyPolicy: optional
	CacheKeyPolicy *CacheKeyPolicy `hcl:"cache_key_policy,block"`
	// NegativeCachingPolicy: min=0
	NegativeCachingPolicy []NegativeCachingPolicy `hcl:"negative_caching_policy,block" validate:"min=0"`
}

type BypassCacheOnRequestHeaders struct {
	// HeaderName: string, optional
	HeaderName terra.StringValue `hcl:"header_name,attr"`
}

type CacheKeyPolicy struct {
	// IncludeHttpHeaders: list of string, optional
	IncludeHttpHeaders terra.ListValue[terra.StringValue] `hcl:"include_http_headers,attr"`
	// QueryStringWhitelist: list of string, optional
	QueryStringWhitelist terra.ListValue[terra.StringValue] `hcl:"query_string_whitelist,attr"`
}

type NegativeCachingPolicy struct {
	// Code: number, optional
	Code terra.NumberValue `hcl:"code,attr"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type CdnPolicyAttributes struct {
	ref terra.Reference
}

func (cp CdnPolicyAttributes) InternalRef() terra.Reference {
	return cp.ref
}

func (cp CdnPolicyAttributes) InternalWithRef(ref terra.Reference) CdnPolicyAttributes {
	return CdnPolicyAttributes{ref: ref}
}

func (cp CdnPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return cp.ref.InternalTokens()
}

func (cp CdnPolicyAttributes) CacheMode() terra.StringValue {
	return terra.ReferenceString(cp.ref.Append("cache_mode"))
}

func (cp CdnPolicyAttributes) ClientTtl() terra.NumberValue {
	return terra.ReferenceNumber(cp.ref.Append("client_ttl"))
}

func (cp CdnPolicyAttributes) DefaultTtl() terra.NumberValue {
	return terra.ReferenceNumber(cp.ref.Append("default_ttl"))
}

func (cp CdnPolicyAttributes) MaxTtl() terra.NumberValue {
	return terra.ReferenceNumber(cp.ref.Append("max_ttl"))
}

func (cp CdnPolicyAttributes) NegativeCaching() terra.BoolValue {
	return terra.ReferenceBool(cp.ref.Append("negative_caching"))
}

func (cp CdnPolicyAttributes) RequestCoalescing() terra.BoolValue {
	return terra.ReferenceBool(cp.ref.Append("request_coalescing"))
}

func (cp CdnPolicyAttributes) ServeWhileStale() terra.NumberValue {
	return terra.ReferenceNumber(cp.ref.Append("serve_while_stale"))
}

func (cp CdnPolicyAttributes) SignedUrlCacheMaxAgeSec() terra.NumberValue {
	return terra.ReferenceNumber(cp.ref.Append("signed_url_cache_max_age_sec"))
}

func (cp CdnPolicyAttributes) BypassCacheOnRequestHeaders() terra.ListValue[BypassCacheOnRequestHeadersAttributes] {
	return terra.ReferenceList[BypassCacheOnRequestHeadersAttributes](cp.ref.Append("bypass_cache_on_request_headers"))
}

func (cp CdnPolicyAttributes) CacheKeyPolicy() terra.ListValue[CacheKeyPolicyAttributes] {
	return terra.ReferenceList[CacheKeyPolicyAttributes](cp.ref.Append("cache_key_policy"))
}

func (cp CdnPolicyAttributes) NegativeCachingPolicy() terra.ListValue[NegativeCachingPolicyAttributes] {
	return terra.ReferenceList[NegativeCachingPolicyAttributes](cp.ref.Append("negative_caching_policy"))
}

type BypassCacheOnRequestHeadersAttributes struct {
	ref terra.Reference
}

func (bcorh BypassCacheOnRequestHeadersAttributes) InternalRef() terra.Reference {
	return bcorh.ref
}

func (bcorh BypassCacheOnRequestHeadersAttributes) InternalWithRef(ref terra.Reference) BypassCacheOnRequestHeadersAttributes {
	return BypassCacheOnRequestHeadersAttributes{ref: ref}
}

func (bcorh BypassCacheOnRequestHeadersAttributes) InternalTokens() hclwrite.Tokens {
	return bcorh.ref.InternalTokens()
}

func (bcorh BypassCacheOnRequestHeadersAttributes) HeaderName() terra.StringValue {
	return terra.ReferenceString(bcorh.ref.Append("header_name"))
}

type CacheKeyPolicyAttributes struct {
	ref terra.Reference
}

func (ckp CacheKeyPolicyAttributes) InternalRef() terra.Reference {
	return ckp.ref
}

func (ckp CacheKeyPolicyAttributes) InternalWithRef(ref terra.Reference) CacheKeyPolicyAttributes {
	return CacheKeyPolicyAttributes{ref: ref}
}

func (ckp CacheKeyPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return ckp.ref.InternalTokens()
}

func (ckp CacheKeyPolicyAttributes) IncludeHttpHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ckp.ref.Append("include_http_headers"))
}

func (ckp CacheKeyPolicyAttributes) QueryStringWhitelist() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ckp.ref.Append("query_string_whitelist"))
}

type NegativeCachingPolicyAttributes struct {
	ref terra.Reference
}

func (ncp NegativeCachingPolicyAttributes) InternalRef() terra.Reference {
	return ncp.ref
}

func (ncp NegativeCachingPolicyAttributes) InternalWithRef(ref terra.Reference) NegativeCachingPolicyAttributes {
	return NegativeCachingPolicyAttributes{ref: ref}
}

func (ncp NegativeCachingPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return ncp.ref.InternalTokens()
}

func (ncp NegativeCachingPolicyAttributes) Code() terra.NumberValue {
	return terra.ReferenceNumber(ncp.ref.Append("code"))
}

func (ncp NegativeCachingPolicyAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceNumber(ncp.ref.Append("ttl"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type CdnPolicyState struct {
	CacheMode                   string                             `json:"cache_mode"`
	ClientTtl                   float64                            `json:"client_ttl"`
	DefaultTtl                  float64                            `json:"default_ttl"`
	MaxTtl                      float64                            `json:"max_ttl"`
	NegativeCaching             bool                               `json:"negative_caching"`
	RequestCoalescing           bool                               `json:"request_coalescing"`
	ServeWhileStale             float64                            `json:"serve_while_stale"`
	SignedUrlCacheMaxAgeSec     float64                            `json:"signed_url_cache_max_age_sec"`
	BypassCacheOnRequestHeaders []BypassCacheOnRequestHeadersState `json:"bypass_cache_on_request_headers"`
	CacheKeyPolicy              []CacheKeyPolicyState              `json:"cache_key_policy"`
	NegativeCachingPolicy       []NegativeCachingPolicyState       `json:"negative_caching_policy"`
}

type BypassCacheOnRequestHeadersState struct {
	HeaderName string `json:"header_name"`
}

type CacheKeyPolicyState struct {
	IncludeHttpHeaders   []string `json:"include_http_headers"`
	QueryStringWhitelist []string `json:"query_string_whitelist"`
}

type NegativeCachingPolicyState struct {
	Code float64 `json:"code"`
	Ttl  float64 `json:"ttl"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
