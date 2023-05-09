// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacomputebackendservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Backend struct{}

type CdnPolicy struct {
	// CacheKeyPolicy: min=0
	CacheKeyPolicy []CacheKeyPolicy `hcl:"cache_key_policy,block" validate:"min=0"`
	// NegativeCachingPolicy: min=0
	NegativeCachingPolicy []NegativeCachingPolicy `hcl:"negative_caching_policy,block" validate:"min=0"`
}

type CacheKeyPolicy struct{}

type NegativeCachingPolicy struct{}

type CircuitBreakers struct{}

type ConsistentHash struct {
	// HttpCookie: min=0
	HttpCookie []HttpCookie `hcl:"http_cookie,block" validate:"min=0"`
}

type HttpCookie struct {
	// Ttl: min=0
	Ttl []Ttl `hcl:"ttl,block" validate:"min=0"`
}

type Ttl struct{}

type Iap struct{}

type LocalityLbPolicies struct {
	// CustomPolicy: min=0
	CustomPolicy []CustomPolicy `hcl:"custom_policy,block" validate:"min=0"`
	// Policy: min=0
	Policy []Policy `hcl:"policy,block" validate:"min=0"`
}

type CustomPolicy struct{}

type Policy struct{}

type LogConfig struct{}

type OutlierDetection struct {
	// BaseEjectionTime: min=0
	BaseEjectionTime []BaseEjectionTime `hcl:"base_ejection_time,block" validate:"min=0"`
	// Interval: min=0
	Interval []Interval `hcl:"interval,block" validate:"min=0"`
}

type BaseEjectionTime struct{}

type Interval struct{}

type SecuritySettings struct{}

type BackendAttributes struct {
	ref terra.Reference
}

func (b BackendAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b BackendAttributes) InternalWithRef(ref terra.Reference) BackendAttributes {
	return BackendAttributes{ref: ref}
}

func (b BackendAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b BackendAttributes) BalancingMode() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("balancing_mode"))
}

func (b BackendAttributes) CapacityScaler() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("capacity_scaler"))
}

func (b BackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("description"))
}

func (b BackendAttributes) Group() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("group"))
}

func (b BackendAttributes) MaxConnections() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("max_connections"))
}

func (b BackendAttributes) MaxConnectionsPerEndpoint() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("max_connections_per_endpoint"))
}

func (b BackendAttributes) MaxConnectionsPerInstance() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("max_connections_per_instance"))
}

func (b BackendAttributes) MaxRate() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("max_rate"))
}

func (b BackendAttributes) MaxRatePerEndpoint() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("max_rate_per_endpoint"))
}

func (b BackendAttributes) MaxRatePerInstance() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("max_rate_per_instance"))
}

func (b BackendAttributes) MaxUtilization() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("max_utilization"))
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

func (cp CdnPolicyAttributes) ClientTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("client_ttl"))
}

func (cp CdnPolicyAttributes) DefaultTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("default_ttl"))
}

func (cp CdnPolicyAttributes) MaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("max_ttl"))
}

func (cp CdnPolicyAttributes) NegativeCaching() terra.BoolValue {
	return terra.ReferenceAsBool(cp.ref.Append("negative_caching"))
}

func (cp CdnPolicyAttributes) ServeWhileStale() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("serve_while_stale"))
}

func (cp CdnPolicyAttributes) SignedUrlCacheMaxAgeSec() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("signed_url_cache_max_age_sec"))
}

func (cp CdnPolicyAttributes) CacheKeyPolicy() terra.ListValue[CacheKeyPolicyAttributes] {
	return terra.ReferenceAsList[CacheKeyPolicyAttributes](cp.ref.Append("cache_key_policy"))
}

func (cp CdnPolicyAttributes) NegativeCachingPolicy() terra.ListValue[NegativeCachingPolicyAttributes] {
	return terra.ReferenceAsList[NegativeCachingPolicyAttributes](cp.ref.Append("negative_caching_policy"))
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

func (ckp CacheKeyPolicyAttributes) IncludeHost() terra.BoolValue {
	return terra.ReferenceAsBool(ckp.ref.Append("include_host"))
}

func (ckp CacheKeyPolicyAttributes) IncludeHttpHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ckp.ref.Append("include_http_headers"))
}

func (ckp CacheKeyPolicyAttributes) IncludeNamedCookies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ckp.ref.Append("include_named_cookies"))
}

func (ckp CacheKeyPolicyAttributes) IncludeProtocol() terra.BoolValue {
	return terra.ReferenceAsBool(ckp.ref.Append("include_protocol"))
}

func (ckp CacheKeyPolicyAttributes) IncludeQueryString() terra.BoolValue {
	return terra.ReferenceAsBool(ckp.ref.Append("include_query_string"))
}

func (ckp CacheKeyPolicyAttributes) QueryStringBlacklist() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ckp.ref.Append("query_string_blacklist"))
}

func (ckp CacheKeyPolicyAttributes) QueryStringWhitelist() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ckp.ref.Append("query_string_whitelist"))
}

type NegativeCachingPolicyAttributes struct {
	ref terra.Reference
}

func (ncp NegativeCachingPolicyAttributes) InternalRef() (terra.Reference, error) {
	return ncp.ref, nil
}

func (ncp NegativeCachingPolicyAttributes) InternalWithRef(ref terra.Reference) NegativeCachingPolicyAttributes {
	return NegativeCachingPolicyAttributes{ref: ref}
}

func (ncp NegativeCachingPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ncp.ref.InternalTokens()
}

func (ncp NegativeCachingPolicyAttributes) Code() terra.NumberValue {
	return terra.ReferenceAsNumber(ncp.ref.Append("code"))
}

func (ncp NegativeCachingPolicyAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(ncp.ref.Append("ttl"))
}

type CircuitBreakersAttributes struct {
	ref terra.Reference
}

func (cb CircuitBreakersAttributes) InternalRef() (terra.Reference, error) {
	return cb.ref, nil
}

func (cb CircuitBreakersAttributes) InternalWithRef(ref terra.Reference) CircuitBreakersAttributes {
	return CircuitBreakersAttributes{ref: ref}
}

func (cb CircuitBreakersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cb.ref.InternalTokens()
}

func (cb CircuitBreakersAttributes) MaxConnections() terra.NumberValue {
	return terra.ReferenceAsNumber(cb.ref.Append("max_connections"))
}

func (cb CircuitBreakersAttributes) MaxPendingRequests() terra.NumberValue {
	return terra.ReferenceAsNumber(cb.ref.Append("max_pending_requests"))
}

func (cb CircuitBreakersAttributes) MaxRequests() terra.NumberValue {
	return terra.ReferenceAsNumber(cb.ref.Append("max_requests"))
}

func (cb CircuitBreakersAttributes) MaxRequestsPerConnection() terra.NumberValue {
	return terra.ReferenceAsNumber(cb.ref.Append("max_requests_per_connection"))
}

func (cb CircuitBreakersAttributes) MaxRetries() terra.NumberValue {
	return terra.ReferenceAsNumber(cb.ref.Append("max_retries"))
}

type ConsistentHashAttributes struct {
	ref terra.Reference
}

func (ch ConsistentHashAttributes) InternalRef() (terra.Reference, error) {
	return ch.ref, nil
}

func (ch ConsistentHashAttributes) InternalWithRef(ref terra.Reference) ConsistentHashAttributes {
	return ConsistentHashAttributes{ref: ref}
}

func (ch ConsistentHashAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ch.ref.InternalTokens()
}

func (ch ConsistentHashAttributes) HttpHeaderName() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("http_header_name"))
}

func (ch ConsistentHashAttributes) MinimumRingSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ch.ref.Append("minimum_ring_size"))
}

func (ch ConsistentHashAttributes) HttpCookie() terra.ListValue[HttpCookieAttributes] {
	return terra.ReferenceAsList[HttpCookieAttributes](ch.ref.Append("http_cookie"))
}

type HttpCookieAttributes struct {
	ref terra.Reference
}

func (hc HttpCookieAttributes) InternalRef() (terra.Reference, error) {
	return hc.ref, nil
}

func (hc HttpCookieAttributes) InternalWithRef(ref terra.Reference) HttpCookieAttributes {
	return HttpCookieAttributes{ref: ref}
}

func (hc HttpCookieAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hc.ref.InternalTokens()
}

func (hc HttpCookieAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("name"))
}

func (hc HttpCookieAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("path"))
}

func (hc HttpCookieAttributes) Ttl() terra.ListValue[TtlAttributes] {
	return terra.ReferenceAsList[TtlAttributes](hc.ref.Append("ttl"))
}

type TtlAttributes struct {
	ref terra.Reference
}

func (t TtlAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TtlAttributes) InternalWithRef(ref terra.Reference) TtlAttributes {
	return TtlAttributes{ref: ref}
}

func (t TtlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TtlAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("nanos"))
}

func (t TtlAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("seconds"))
}

type IapAttributes struct {
	ref terra.Reference
}

func (i IapAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IapAttributes) InternalWithRef(ref terra.Reference) IapAttributes {
	return IapAttributes{ref: ref}
}

func (i IapAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IapAttributes) Oauth2ClientId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("oauth2_client_id"))
}

func (i IapAttributes) Oauth2ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("oauth2_client_secret"))
}

func (i IapAttributes) Oauth2ClientSecretSha256() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("oauth2_client_secret_sha256"))
}

type LocalityLbPoliciesAttributes struct {
	ref terra.Reference
}

func (llp LocalityLbPoliciesAttributes) InternalRef() (terra.Reference, error) {
	return llp.ref, nil
}

func (llp LocalityLbPoliciesAttributes) InternalWithRef(ref terra.Reference) LocalityLbPoliciesAttributes {
	return LocalityLbPoliciesAttributes{ref: ref}
}

func (llp LocalityLbPoliciesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return llp.ref.InternalTokens()
}

func (llp LocalityLbPoliciesAttributes) CustomPolicy() terra.ListValue[CustomPolicyAttributes] {
	return terra.ReferenceAsList[CustomPolicyAttributes](llp.ref.Append("custom_policy"))
}

func (llp LocalityLbPoliciesAttributes) Policy() terra.ListValue[PolicyAttributes] {
	return terra.ReferenceAsList[PolicyAttributes](llp.ref.Append("policy"))
}

type CustomPolicyAttributes struct {
	ref terra.Reference
}

func (cp CustomPolicyAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp CustomPolicyAttributes) InternalWithRef(ref terra.Reference) CustomPolicyAttributes {
	return CustomPolicyAttributes{ref: ref}
}

func (cp CustomPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cp.ref.InternalTokens()
}

func (cp CustomPolicyAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("data"))
}

func (cp CustomPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("name"))
}

type PolicyAttributes struct {
	ref terra.Reference
}

func (p PolicyAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PolicyAttributes) InternalWithRef(ref terra.Reference) PolicyAttributes {
	return PolicyAttributes{ref: ref}
}

func (p PolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
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

type OutlierDetectionAttributes struct {
	ref terra.Reference
}

func (od OutlierDetectionAttributes) InternalRef() (terra.Reference, error) {
	return od.ref, nil
}

func (od OutlierDetectionAttributes) InternalWithRef(ref terra.Reference) OutlierDetectionAttributes {
	return OutlierDetectionAttributes{ref: ref}
}

func (od OutlierDetectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return od.ref.InternalTokens()
}

func (od OutlierDetectionAttributes) ConsecutiveErrors() terra.NumberValue {
	return terra.ReferenceAsNumber(od.ref.Append("consecutive_errors"))
}

func (od OutlierDetectionAttributes) ConsecutiveGatewayFailure() terra.NumberValue {
	return terra.ReferenceAsNumber(od.ref.Append("consecutive_gateway_failure"))
}

func (od OutlierDetectionAttributes) EnforcingConsecutiveErrors() terra.NumberValue {
	return terra.ReferenceAsNumber(od.ref.Append("enforcing_consecutive_errors"))
}

func (od OutlierDetectionAttributes) EnforcingConsecutiveGatewayFailure() terra.NumberValue {
	return terra.ReferenceAsNumber(od.ref.Append("enforcing_consecutive_gateway_failure"))
}

func (od OutlierDetectionAttributes) EnforcingSuccessRate() terra.NumberValue {
	return terra.ReferenceAsNumber(od.ref.Append("enforcing_success_rate"))
}

func (od OutlierDetectionAttributes) MaxEjectionPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(od.ref.Append("max_ejection_percent"))
}

func (od OutlierDetectionAttributes) SuccessRateMinimumHosts() terra.NumberValue {
	return terra.ReferenceAsNumber(od.ref.Append("success_rate_minimum_hosts"))
}

func (od OutlierDetectionAttributes) SuccessRateRequestVolume() terra.NumberValue {
	return terra.ReferenceAsNumber(od.ref.Append("success_rate_request_volume"))
}

func (od OutlierDetectionAttributes) SuccessRateStdevFactor() terra.NumberValue {
	return terra.ReferenceAsNumber(od.ref.Append("success_rate_stdev_factor"))
}

func (od OutlierDetectionAttributes) BaseEjectionTime() terra.ListValue[BaseEjectionTimeAttributes] {
	return terra.ReferenceAsList[BaseEjectionTimeAttributes](od.ref.Append("base_ejection_time"))
}

func (od OutlierDetectionAttributes) Interval() terra.ListValue[IntervalAttributes] {
	return terra.ReferenceAsList[IntervalAttributes](od.ref.Append("interval"))
}

type BaseEjectionTimeAttributes struct {
	ref terra.Reference
}

func (bet BaseEjectionTimeAttributes) InternalRef() (terra.Reference, error) {
	return bet.ref, nil
}

func (bet BaseEjectionTimeAttributes) InternalWithRef(ref terra.Reference) BaseEjectionTimeAttributes {
	return BaseEjectionTimeAttributes{ref: ref}
}

func (bet BaseEjectionTimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bet.ref.InternalTokens()
}

func (bet BaseEjectionTimeAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceAsNumber(bet.ref.Append("nanos"))
}

func (bet BaseEjectionTimeAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceAsNumber(bet.ref.Append("seconds"))
}

type IntervalAttributes struct {
	ref terra.Reference
}

func (i IntervalAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IntervalAttributes) InternalWithRef(ref terra.Reference) IntervalAttributes {
	return IntervalAttributes{ref: ref}
}

func (i IntervalAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IntervalAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("nanos"))
}

func (i IntervalAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("seconds"))
}

type SecuritySettingsAttributes struct {
	ref terra.Reference
}

func (ss SecuritySettingsAttributes) InternalRef() (terra.Reference, error) {
	return ss.ref, nil
}

func (ss SecuritySettingsAttributes) InternalWithRef(ref terra.Reference) SecuritySettingsAttributes {
	return SecuritySettingsAttributes{ref: ref}
}

func (ss SecuritySettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ss.ref.InternalTokens()
}

func (ss SecuritySettingsAttributes) ClientTlsPolicy() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("client_tls_policy"))
}

func (ss SecuritySettingsAttributes) SubjectAltNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ss.ref.Append("subject_alt_names"))
}

type BackendState struct {
	BalancingMode             string  `json:"balancing_mode"`
	CapacityScaler            float64 `json:"capacity_scaler"`
	Description               string  `json:"description"`
	Group                     string  `json:"group"`
	MaxConnections            float64 `json:"max_connections"`
	MaxConnectionsPerEndpoint float64 `json:"max_connections_per_endpoint"`
	MaxConnectionsPerInstance float64 `json:"max_connections_per_instance"`
	MaxRate                   float64 `json:"max_rate"`
	MaxRatePerEndpoint        float64 `json:"max_rate_per_endpoint"`
	MaxRatePerInstance        float64 `json:"max_rate_per_instance"`
	MaxUtilization            float64 `json:"max_utilization"`
}

type CdnPolicyState struct {
	CacheMode               string                       `json:"cache_mode"`
	ClientTtl               float64                      `json:"client_ttl"`
	DefaultTtl              float64                      `json:"default_ttl"`
	MaxTtl                  float64                      `json:"max_ttl"`
	NegativeCaching         bool                         `json:"negative_caching"`
	ServeWhileStale         float64                      `json:"serve_while_stale"`
	SignedUrlCacheMaxAgeSec float64                      `json:"signed_url_cache_max_age_sec"`
	CacheKeyPolicy          []CacheKeyPolicyState        `json:"cache_key_policy"`
	NegativeCachingPolicy   []NegativeCachingPolicyState `json:"negative_caching_policy"`
}

type CacheKeyPolicyState struct {
	IncludeHost          bool     `json:"include_host"`
	IncludeHttpHeaders   []string `json:"include_http_headers"`
	IncludeNamedCookies  []string `json:"include_named_cookies"`
	IncludeProtocol      bool     `json:"include_protocol"`
	IncludeQueryString   bool     `json:"include_query_string"`
	QueryStringBlacklist []string `json:"query_string_blacklist"`
	QueryStringWhitelist []string `json:"query_string_whitelist"`
}

type NegativeCachingPolicyState struct {
	Code float64 `json:"code"`
	Ttl  float64 `json:"ttl"`
}

type CircuitBreakersState struct {
	MaxConnections           float64 `json:"max_connections"`
	MaxPendingRequests       float64 `json:"max_pending_requests"`
	MaxRequests              float64 `json:"max_requests"`
	MaxRequestsPerConnection float64 `json:"max_requests_per_connection"`
	MaxRetries               float64 `json:"max_retries"`
}

type ConsistentHashState struct {
	HttpHeaderName  string            `json:"http_header_name"`
	MinimumRingSize float64           `json:"minimum_ring_size"`
	HttpCookie      []HttpCookieState `json:"http_cookie"`
}

type HttpCookieState struct {
	Name string     `json:"name"`
	Path string     `json:"path"`
	Ttl  []TtlState `json:"ttl"`
}

type TtlState struct {
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type IapState struct {
	Oauth2ClientId           string `json:"oauth2_client_id"`
	Oauth2ClientSecret       string `json:"oauth2_client_secret"`
	Oauth2ClientSecretSha256 string `json:"oauth2_client_secret_sha256"`
}

type LocalityLbPoliciesState struct {
	CustomPolicy []CustomPolicyState `json:"custom_policy"`
	Policy       []PolicyState       `json:"policy"`
}

type CustomPolicyState struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

type PolicyState struct {
	Name string `json:"name"`
}

type LogConfigState struct {
	Enable     bool    `json:"enable"`
	SampleRate float64 `json:"sample_rate"`
}

type OutlierDetectionState struct {
	ConsecutiveErrors                  float64                 `json:"consecutive_errors"`
	ConsecutiveGatewayFailure          float64                 `json:"consecutive_gateway_failure"`
	EnforcingConsecutiveErrors         float64                 `json:"enforcing_consecutive_errors"`
	EnforcingConsecutiveGatewayFailure float64                 `json:"enforcing_consecutive_gateway_failure"`
	EnforcingSuccessRate               float64                 `json:"enforcing_success_rate"`
	MaxEjectionPercent                 float64                 `json:"max_ejection_percent"`
	SuccessRateMinimumHosts            float64                 `json:"success_rate_minimum_hosts"`
	SuccessRateRequestVolume           float64                 `json:"success_rate_request_volume"`
	SuccessRateStdevFactor             float64                 `json:"success_rate_stdev_factor"`
	BaseEjectionTime                   []BaseEjectionTimeState `json:"base_ejection_time"`
	Interval                           []IntervalState         `json:"interval"`
}

type BaseEjectionTimeState struct {
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type IntervalState struct {
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type SecuritySettingsState struct {
	ClientTlsPolicy string   `json:"client_tls_policy"`
	SubjectAltNames []string `json:"subject_alt_names"`
}
