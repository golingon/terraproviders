// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apigatewayv2stage

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AccessLogSettings struct {
	// DestinationArn: string, required
	DestinationArn terra.StringValue `hcl:"destination_arn,attr" validate:"required"`
	// Format: string, required
	Format terra.StringValue `hcl:"format,attr" validate:"required"`
}

type DefaultRouteSettings struct {
	// DataTraceEnabled: bool, optional
	DataTraceEnabled terra.BoolValue `hcl:"data_trace_enabled,attr"`
	// DetailedMetricsEnabled: bool, optional
	DetailedMetricsEnabled terra.BoolValue `hcl:"detailed_metrics_enabled,attr"`
	// LoggingLevel: string, optional
	LoggingLevel terra.StringValue `hcl:"logging_level,attr"`
	// ThrottlingBurstLimit: number, optional
	ThrottlingBurstLimit terra.NumberValue `hcl:"throttling_burst_limit,attr"`
	// ThrottlingRateLimit: number, optional
	ThrottlingRateLimit terra.NumberValue `hcl:"throttling_rate_limit,attr"`
}

type RouteSettings struct {
	// DataTraceEnabled: bool, optional
	DataTraceEnabled terra.BoolValue `hcl:"data_trace_enabled,attr"`
	// DetailedMetricsEnabled: bool, optional
	DetailedMetricsEnabled terra.BoolValue `hcl:"detailed_metrics_enabled,attr"`
	// LoggingLevel: string, optional
	LoggingLevel terra.StringValue `hcl:"logging_level,attr"`
	// RouteKey: string, required
	RouteKey terra.StringValue `hcl:"route_key,attr" validate:"required"`
	// ThrottlingBurstLimit: number, optional
	ThrottlingBurstLimit terra.NumberValue `hcl:"throttling_burst_limit,attr"`
	// ThrottlingRateLimit: number, optional
	ThrottlingRateLimit terra.NumberValue `hcl:"throttling_rate_limit,attr"`
}

type AccessLogSettingsAttributes struct {
	ref terra.Reference
}

func (als AccessLogSettingsAttributes) InternalRef() terra.Reference {
	return als.ref
}

func (als AccessLogSettingsAttributes) InternalWithRef(ref terra.Reference) AccessLogSettingsAttributes {
	return AccessLogSettingsAttributes{ref: ref}
}

func (als AccessLogSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return als.ref.InternalTokens()
}

func (als AccessLogSettingsAttributes) DestinationArn() terra.StringValue {
	return terra.ReferenceAsString(als.ref.Append("destination_arn"))
}

func (als AccessLogSettingsAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(als.ref.Append("format"))
}

type DefaultRouteSettingsAttributes struct {
	ref terra.Reference
}

func (drs DefaultRouteSettingsAttributes) InternalRef() terra.Reference {
	return drs.ref
}

func (drs DefaultRouteSettingsAttributes) InternalWithRef(ref terra.Reference) DefaultRouteSettingsAttributes {
	return DefaultRouteSettingsAttributes{ref: ref}
}

func (drs DefaultRouteSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return drs.ref.InternalTokens()
}

func (drs DefaultRouteSettingsAttributes) DataTraceEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(drs.ref.Append("data_trace_enabled"))
}

func (drs DefaultRouteSettingsAttributes) DetailedMetricsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(drs.ref.Append("detailed_metrics_enabled"))
}

func (drs DefaultRouteSettingsAttributes) LoggingLevel() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("logging_level"))
}

func (drs DefaultRouteSettingsAttributes) ThrottlingBurstLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(drs.ref.Append("throttling_burst_limit"))
}

func (drs DefaultRouteSettingsAttributes) ThrottlingRateLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(drs.ref.Append("throttling_rate_limit"))
}

type RouteSettingsAttributes struct {
	ref terra.Reference
}

func (rs RouteSettingsAttributes) InternalRef() terra.Reference {
	return rs.ref
}

func (rs RouteSettingsAttributes) InternalWithRef(ref terra.Reference) RouteSettingsAttributes {
	return RouteSettingsAttributes{ref: ref}
}

func (rs RouteSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return rs.ref.InternalTokens()
}

func (rs RouteSettingsAttributes) DataTraceEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rs.ref.Append("data_trace_enabled"))
}

func (rs RouteSettingsAttributes) DetailedMetricsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rs.ref.Append("detailed_metrics_enabled"))
}

func (rs RouteSettingsAttributes) LoggingLevel() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("logging_level"))
}

func (rs RouteSettingsAttributes) RouteKey() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("route_key"))
}

func (rs RouteSettingsAttributes) ThrottlingBurstLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(rs.ref.Append("throttling_burst_limit"))
}

func (rs RouteSettingsAttributes) ThrottlingRateLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(rs.ref.Append("throttling_rate_limit"))
}

type AccessLogSettingsState struct {
	DestinationArn string `json:"destination_arn"`
	Format         string `json:"format"`
}

type DefaultRouteSettingsState struct {
	DataTraceEnabled       bool    `json:"data_trace_enabled"`
	DetailedMetricsEnabled bool    `json:"detailed_metrics_enabled"`
	LoggingLevel           string  `json:"logging_level"`
	ThrottlingBurstLimit   float64 `json:"throttling_burst_limit"`
	ThrottlingRateLimit    float64 `json:"throttling_rate_limit"`
}

type RouteSettingsState struct {
	DataTraceEnabled       bool    `json:"data_trace_enabled"`
	DetailedMetricsEnabled bool    `json:"detailed_metrics_enabled"`
	LoggingLevel           string  `json:"logging_level"`
	RouteKey               string  `json:"route_key"`
	ThrottlingBurstLimit   float64 `json:"throttling_burst_limit"`
	ThrottlingRateLimit    float64 `json:"throttling_rate_limit"`
}
