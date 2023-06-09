// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apigatewaymethodsettings

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Settings struct {
	// CacheDataEncrypted: bool, optional
	CacheDataEncrypted terra.BoolValue `hcl:"cache_data_encrypted,attr"`
	// CacheTtlInSeconds: number, optional
	CacheTtlInSeconds terra.NumberValue `hcl:"cache_ttl_in_seconds,attr"`
	// CachingEnabled: bool, optional
	CachingEnabled terra.BoolValue `hcl:"caching_enabled,attr"`
	// DataTraceEnabled: bool, optional
	DataTraceEnabled terra.BoolValue `hcl:"data_trace_enabled,attr"`
	// LoggingLevel: string, optional
	LoggingLevel terra.StringValue `hcl:"logging_level,attr"`
	// MetricsEnabled: bool, optional
	MetricsEnabled terra.BoolValue `hcl:"metrics_enabled,attr"`
	// RequireAuthorizationForCacheControl: bool, optional
	RequireAuthorizationForCacheControl terra.BoolValue `hcl:"require_authorization_for_cache_control,attr"`
	// ThrottlingBurstLimit: number, optional
	ThrottlingBurstLimit terra.NumberValue `hcl:"throttling_burst_limit,attr"`
	// ThrottlingRateLimit: number, optional
	ThrottlingRateLimit terra.NumberValue `hcl:"throttling_rate_limit,attr"`
	// UnauthorizedCacheControlHeaderStrategy: string, optional
	UnauthorizedCacheControlHeaderStrategy terra.StringValue `hcl:"unauthorized_cache_control_header_strategy,attr"`
}

type SettingsAttributes struct {
	ref terra.Reference
}

func (s SettingsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SettingsAttributes) InternalWithRef(ref terra.Reference) SettingsAttributes {
	return SettingsAttributes{ref: ref}
}

func (s SettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SettingsAttributes) CacheDataEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("cache_data_encrypted"))
}

func (s SettingsAttributes) CacheTtlInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("cache_ttl_in_seconds"))
}

func (s SettingsAttributes) CachingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("caching_enabled"))
}

func (s SettingsAttributes) DataTraceEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("data_trace_enabled"))
}

func (s SettingsAttributes) LoggingLevel() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("logging_level"))
}

func (s SettingsAttributes) MetricsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("metrics_enabled"))
}

func (s SettingsAttributes) RequireAuthorizationForCacheControl() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("require_authorization_for_cache_control"))
}

func (s SettingsAttributes) ThrottlingBurstLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("throttling_burst_limit"))
}

func (s SettingsAttributes) ThrottlingRateLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("throttling_rate_limit"))
}

func (s SettingsAttributes) UnauthorizedCacheControlHeaderStrategy() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("unauthorized_cache_control_header_strategy"))
}

type SettingsState struct {
	CacheDataEncrypted                     bool    `json:"cache_data_encrypted"`
	CacheTtlInSeconds                      float64 `json:"cache_ttl_in_seconds"`
	CachingEnabled                         bool    `json:"caching_enabled"`
	DataTraceEnabled                       bool    `json:"data_trace_enabled"`
	LoggingLevel                           string  `json:"logging_level"`
	MetricsEnabled                         bool    `json:"metrics_enabled"`
	RequireAuthorizationForCacheControl    bool    `json:"require_authorization_for_cache_control"`
	ThrottlingBurstLimit                   float64 `json:"throttling_burst_limit"`
	ThrottlingRateLimit                    float64 `json:"throttling_rate_limit"`
	UnauthorizedCacheControlHeaderStrategy string  `json:"unauthorized_cache_control_header_strategy"`
}
