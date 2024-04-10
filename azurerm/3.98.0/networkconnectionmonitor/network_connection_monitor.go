// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package networkconnectionmonitor

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Endpoint struct {
	// Address: string, optional
	Address terra.StringValue `hcl:"address,attr"`
	// CoverageLevel: string, optional
	CoverageLevel terra.StringValue `hcl:"coverage_level,attr"`
	// ExcludedIpAddresses: set of string, optional
	ExcludedIpAddresses terra.SetValue[terra.StringValue] `hcl:"excluded_ip_addresses,attr"`
	// IncludedIpAddresses: set of string, optional
	IncludedIpAddresses terra.SetValue[terra.StringValue] `hcl:"included_ip_addresses,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TargetResourceId: string, optional
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr"`
	// TargetResourceType: string, optional
	TargetResourceType terra.StringValue `hcl:"target_resource_type,attr"`
	// Filter: optional
	Filter *Filter `hcl:"filter,block"`
}

type Filter struct {
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Item: min=0
	Item []Item `hcl:"item,block" validate:"min=0"`
}

type Item struct {
	// Address: string, optional
	Address terra.StringValue `hcl:"address,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type TestConfiguration struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PreferredIpVersion: string, optional
	PreferredIpVersion terra.StringValue `hcl:"preferred_ip_version,attr"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// TestFrequencyInSeconds: number, optional
	TestFrequencyInSeconds terra.NumberValue `hcl:"test_frequency_in_seconds,attr"`
	// HttpConfiguration: optional
	HttpConfiguration *HttpConfiguration `hcl:"http_configuration,block"`
	// IcmpConfiguration: optional
	IcmpConfiguration *IcmpConfiguration `hcl:"icmp_configuration,block"`
	// SuccessThreshold: optional
	SuccessThreshold *SuccessThreshold `hcl:"success_threshold,block"`
	// TcpConfiguration: optional
	TcpConfiguration *TcpConfiguration `hcl:"tcp_configuration,block"`
}

type HttpConfiguration struct {
	// Method: string, optional
	Method terra.StringValue `hcl:"method,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PreferHttps: bool, optional
	PreferHttps terra.BoolValue `hcl:"prefer_https,attr"`
	// ValidStatusCodeRanges: set of string, optional
	ValidStatusCodeRanges terra.SetValue[terra.StringValue] `hcl:"valid_status_code_ranges,attr"`
	// RequestHeader: min=0
	RequestHeader []RequestHeader `hcl:"request_header,block" validate:"min=0"`
}

type RequestHeader struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type IcmpConfiguration struct {
	// TraceRouteEnabled: bool, optional
	TraceRouteEnabled terra.BoolValue `hcl:"trace_route_enabled,attr"`
}

type SuccessThreshold struct {
	// ChecksFailedPercent: number, optional
	ChecksFailedPercent terra.NumberValue `hcl:"checks_failed_percent,attr"`
	// RoundTripTimeMs: number, optional
	RoundTripTimeMs terra.NumberValue `hcl:"round_trip_time_ms,attr"`
}

type TcpConfiguration struct {
	// DestinationPortBehavior: string, optional
	DestinationPortBehavior terra.StringValue `hcl:"destination_port_behavior,attr"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// TraceRouteEnabled: bool, optional
	TraceRouteEnabled terra.BoolValue `hcl:"trace_route_enabled,attr"`
}

type TestGroup struct {
	// DestinationEndpoints: set of string, required
	DestinationEndpoints terra.SetValue[terra.StringValue] `hcl:"destination_endpoints,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SourceEndpoints: set of string, required
	SourceEndpoints terra.SetValue[terra.StringValue] `hcl:"source_endpoints,attr" validate:"required"`
	// TestConfigurationNames: set of string, required
	TestConfigurationNames terra.SetValue[terra.StringValue] `hcl:"test_configuration_names,attr" validate:"required"`
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

type EndpointAttributes struct {
	ref terra.Reference
}

func (e EndpointAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EndpointAttributes) InternalWithRef(ref terra.Reference) EndpointAttributes {
	return EndpointAttributes{ref: ref}
}

func (e EndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EndpointAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("address"))
}

func (e EndpointAttributes) CoverageLevel() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("coverage_level"))
}

func (e EndpointAttributes) ExcludedIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](e.ref.Append("excluded_ip_addresses"))
}

func (e EndpointAttributes) IncludedIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](e.ref.Append("included_ip_addresses"))
}

func (e EndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("name"))
}

func (e EndpointAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("target_resource_id"))
}

func (e EndpointAttributes) TargetResourceType() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("target_resource_type"))
}

func (e EndpointAttributes) Filter() terra.ListValue[FilterAttributes] {
	return terra.ReferenceAsList[FilterAttributes](e.ref.Append("filter"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("type"))
}

func (f FilterAttributes) Item() terra.SetValue[ItemAttributes] {
	return terra.ReferenceAsSet[ItemAttributes](f.ref.Append("item"))
}

type ItemAttributes struct {
	ref terra.Reference
}

func (i ItemAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i ItemAttributes) InternalWithRef(ref terra.Reference) ItemAttributes {
	return ItemAttributes{ref: ref}
}

func (i ItemAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i ItemAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("address"))
}

func (i ItemAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type TestConfigurationAttributes struct {
	ref terra.Reference
}

func (tc TestConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc TestConfigurationAttributes) InternalWithRef(ref terra.Reference) TestConfigurationAttributes {
	return TestConfigurationAttributes{ref: ref}
}

func (tc TestConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tc.ref.InternalTokens()
}

func (tc TestConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("name"))
}

func (tc TestConfigurationAttributes) PreferredIpVersion() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("preferred_ip_version"))
}

func (tc TestConfigurationAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("protocol"))
}

func (tc TestConfigurationAttributes) TestFrequencyInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(tc.ref.Append("test_frequency_in_seconds"))
}

func (tc TestConfigurationAttributes) HttpConfiguration() terra.ListValue[HttpConfigurationAttributes] {
	return terra.ReferenceAsList[HttpConfigurationAttributes](tc.ref.Append("http_configuration"))
}

func (tc TestConfigurationAttributes) IcmpConfiguration() terra.ListValue[IcmpConfigurationAttributes] {
	return terra.ReferenceAsList[IcmpConfigurationAttributes](tc.ref.Append("icmp_configuration"))
}

func (tc TestConfigurationAttributes) SuccessThreshold() terra.ListValue[SuccessThresholdAttributes] {
	return terra.ReferenceAsList[SuccessThresholdAttributes](tc.ref.Append("success_threshold"))
}

func (tc TestConfigurationAttributes) TcpConfiguration() terra.ListValue[TcpConfigurationAttributes] {
	return terra.ReferenceAsList[TcpConfigurationAttributes](tc.ref.Append("tcp_configuration"))
}

type HttpConfigurationAttributes struct {
	ref terra.Reference
}

func (hc HttpConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return hc.ref, nil
}

func (hc HttpConfigurationAttributes) InternalWithRef(ref terra.Reference) HttpConfigurationAttributes {
	return HttpConfigurationAttributes{ref: ref}
}

func (hc HttpConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hc.ref.InternalTokens()
}

func (hc HttpConfigurationAttributes) Method() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("method"))
}

func (hc HttpConfigurationAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("path"))
}

func (hc HttpConfigurationAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("port"))
}

func (hc HttpConfigurationAttributes) PreferHttps() terra.BoolValue {
	return terra.ReferenceAsBool(hc.ref.Append("prefer_https"))
}

func (hc HttpConfigurationAttributes) ValidStatusCodeRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hc.ref.Append("valid_status_code_ranges"))
}

func (hc HttpConfigurationAttributes) RequestHeader() terra.SetValue[RequestHeaderAttributes] {
	return terra.ReferenceAsSet[RequestHeaderAttributes](hc.ref.Append("request_header"))
}

type RequestHeaderAttributes struct {
	ref terra.Reference
}

func (rh RequestHeaderAttributes) InternalRef() (terra.Reference, error) {
	return rh.ref, nil
}

func (rh RequestHeaderAttributes) InternalWithRef(ref terra.Reference) RequestHeaderAttributes {
	return RequestHeaderAttributes{ref: ref}
}

func (rh RequestHeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rh.ref.InternalTokens()
}

func (rh RequestHeaderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rh.ref.Append("name"))
}

func (rh RequestHeaderAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(rh.ref.Append("value"))
}

type IcmpConfigurationAttributes struct {
	ref terra.Reference
}

func (ic IcmpConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic IcmpConfigurationAttributes) InternalWithRef(ref terra.Reference) IcmpConfigurationAttributes {
	return IcmpConfigurationAttributes{ref: ref}
}

func (ic IcmpConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic IcmpConfigurationAttributes) TraceRouteEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("trace_route_enabled"))
}

type SuccessThresholdAttributes struct {
	ref terra.Reference
}

func (st SuccessThresholdAttributes) InternalRef() (terra.Reference, error) {
	return st.ref, nil
}

func (st SuccessThresholdAttributes) InternalWithRef(ref terra.Reference) SuccessThresholdAttributes {
	return SuccessThresholdAttributes{ref: ref}
}

func (st SuccessThresholdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return st.ref.InternalTokens()
}

func (st SuccessThresholdAttributes) ChecksFailedPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("checks_failed_percent"))
}

func (st SuccessThresholdAttributes) RoundTripTimeMs() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("round_trip_time_ms"))
}

type TcpConfigurationAttributes struct {
	ref terra.Reference
}

func (tc TcpConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc TcpConfigurationAttributes) InternalWithRef(ref terra.Reference) TcpConfigurationAttributes {
	return TcpConfigurationAttributes{ref: ref}
}

func (tc TcpConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tc.ref.InternalTokens()
}

func (tc TcpConfigurationAttributes) DestinationPortBehavior() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("destination_port_behavior"))
}

func (tc TcpConfigurationAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(tc.ref.Append("port"))
}

func (tc TcpConfigurationAttributes) TraceRouteEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(tc.ref.Append("trace_route_enabled"))
}

type TestGroupAttributes struct {
	ref terra.Reference
}

func (tg TestGroupAttributes) InternalRef() (terra.Reference, error) {
	return tg.ref, nil
}

func (tg TestGroupAttributes) InternalWithRef(ref terra.Reference) TestGroupAttributes {
	return TestGroupAttributes{ref: ref}
}

func (tg TestGroupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tg.ref.InternalTokens()
}

func (tg TestGroupAttributes) DestinationEndpoints() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tg.ref.Append("destination_endpoints"))
}

func (tg TestGroupAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(tg.ref.Append("enabled"))
}

func (tg TestGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tg.ref.Append("name"))
}

func (tg TestGroupAttributes) SourceEndpoints() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tg.ref.Append("source_endpoints"))
}

func (tg TestGroupAttributes) TestConfigurationNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tg.ref.Append("test_configuration_names"))
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

type EndpointState struct {
	Address             string        `json:"address"`
	CoverageLevel       string        `json:"coverage_level"`
	ExcludedIpAddresses []string      `json:"excluded_ip_addresses"`
	IncludedIpAddresses []string      `json:"included_ip_addresses"`
	Name                string        `json:"name"`
	TargetResourceId    string        `json:"target_resource_id"`
	TargetResourceType  string        `json:"target_resource_type"`
	Filter              []FilterState `json:"filter"`
}

type FilterState struct {
	Type string      `json:"type"`
	Item []ItemState `json:"item"`
}

type ItemState struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}

type TestConfigurationState struct {
	Name                   string                   `json:"name"`
	PreferredIpVersion     string                   `json:"preferred_ip_version"`
	Protocol               string                   `json:"protocol"`
	TestFrequencyInSeconds float64                  `json:"test_frequency_in_seconds"`
	HttpConfiguration      []HttpConfigurationState `json:"http_configuration"`
	IcmpConfiguration      []IcmpConfigurationState `json:"icmp_configuration"`
	SuccessThreshold       []SuccessThresholdState  `json:"success_threshold"`
	TcpConfiguration       []TcpConfigurationState  `json:"tcp_configuration"`
}

type HttpConfigurationState struct {
	Method                string               `json:"method"`
	Path                  string               `json:"path"`
	Port                  float64              `json:"port"`
	PreferHttps           bool                 `json:"prefer_https"`
	ValidStatusCodeRanges []string             `json:"valid_status_code_ranges"`
	RequestHeader         []RequestHeaderState `json:"request_header"`
}

type RequestHeaderState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type IcmpConfigurationState struct {
	TraceRouteEnabled bool `json:"trace_route_enabled"`
}

type SuccessThresholdState struct {
	ChecksFailedPercent float64 `json:"checks_failed_percent"`
	RoundTripTimeMs     float64 `json:"round_trip_time_ms"`
}

type TcpConfigurationState struct {
	DestinationPortBehavior string  `json:"destination_port_behavior"`
	Port                    float64 `json:"port"`
	TraceRouteEnabled       bool    `json:"trace_route_enabled"`
}

type TestGroupState struct {
	DestinationEndpoints   []string `json:"destination_endpoints"`
	Enabled                bool     `json:"enabled"`
	Name                   string   `json:"name"`
	SourceEndpoints        []string `json:"source_endpoints"`
	TestConfigurationNames []string `json:"test_configuration_names"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
