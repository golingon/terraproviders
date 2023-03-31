// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package frontdoor

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ExplicitResourceOrder struct{}

type BackendPool struct {
	// HealthProbeName: string, required
	HealthProbeName terra.StringValue `hcl:"health_probe_name,attr" validate:"required"`
	// LoadBalancingName: string, required
	LoadBalancingName terra.StringValue `hcl:"load_balancing_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Backend: min=1,max=500
	Backend []Backend `hcl:"backend,block" validate:"min=1,max=500"`
}

type Backend struct {
	// Address: string, required
	Address terra.StringValue `hcl:"address,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// HostHeader: string, required
	HostHeader terra.StringValue `hcl:"host_header,attr" validate:"required"`
	// HttpPort: number, required
	HttpPort terra.NumberValue `hcl:"http_port,attr" validate:"required"`
	// HttpsPort: number, required
	HttpsPort terra.NumberValue `hcl:"https_port,attr" validate:"required"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
}

type BackendPoolHealthProbe struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// IntervalInSeconds: number, optional
	IntervalInSeconds terra.NumberValue `hcl:"interval_in_seconds,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// ProbeMethod: string, optional
	ProbeMethod terra.StringValue `hcl:"probe_method,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
}

type BackendPoolLoadBalancing struct {
	// AdditionalLatencyMilliseconds: number, optional
	AdditionalLatencyMilliseconds terra.NumberValue `hcl:"additional_latency_milliseconds,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SampleSize: number, optional
	SampleSize terra.NumberValue `hcl:"sample_size,attr"`
	// SuccessfulSamplesRequired: number, optional
	SuccessfulSamplesRequired terra.NumberValue `hcl:"successful_samples_required,attr"`
}

type BackendPoolSettings struct {
	// BackendPoolsSendReceiveTimeoutSeconds: number, optional
	BackendPoolsSendReceiveTimeoutSeconds terra.NumberValue `hcl:"backend_pools_send_receive_timeout_seconds,attr"`
	// EnforceBackendPoolsCertificateNameCheck: bool, required
	EnforceBackendPoolsCertificateNameCheck terra.BoolValue `hcl:"enforce_backend_pools_certificate_name_check,attr" validate:"required"`
}

type FrontendEndpoint struct {
	// HostName: string, required
	HostName terra.StringValue `hcl:"host_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SessionAffinityEnabled: bool, optional
	SessionAffinityEnabled terra.BoolValue `hcl:"session_affinity_enabled,attr"`
	// SessionAffinityTtlSeconds: number, optional
	SessionAffinityTtlSeconds terra.NumberValue `hcl:"session_affinity_ttl_seconds,attr"`
	// WebApplicationFirewallPolicyLinkId: string, optional
	WebApplicationFirewallPolicyLinkId terra.StringValue `hcl:"web_application_firewall_policy_link_id,attr"`
}

type RoutingRule struct {
	// AcceptedProtocols: list of string, required
	AcceptedProtocols terra.ListValue[terra.StringValue] `hcl:"accepted_protocols,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// FrontendEndpoints: list of string, required
	FrontendEndpoints terra.ListValue[terra.StringValue] `hcl:"frontend_endpoints,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PatternsToMatch: list of string, required
	PatternsToMatch terra.ListValue[terra.StringValue] `hcl:"patterns_to_match,attr" validate:"required"`
	// ForwardingConfiguration: optional
	ForwardingConfiguration *ForwardingConfiguration `hcl:"forwarding_configuration,block"`
	// RedirectConfiguration: optional
	RedirectConfiguration *RedirectConfiguration `hcl:"redirect_configuration,block"`
}

type ForwardingConfiguration struct {
	// BackendPoolName: string, required
	BackendPoolName terra.StringValue `hcl:"backend_pool_name,attr" validate:"required"`
	// CacheDuration: string, optional
	CacheDuration terra.StringValue `hcl:"cache_duration,attr"`
	// CacheEnabled: bool, optional
	CacheEnabled terra.BoolValue `hcl:"cache_enabled,attr"`
	// CacheQueryParameterStripDirective: string, optional
	CacheQueryParameterStripDirective terra.StringValue `hcl:"cache_query_parameter_strip_directive,attr"`
	// CacheQueryParameters: list of string, optional
	CacheQueryParameters terra.ListValue[terra.StringValue] `hcl:"cache_query_parameters,attr"`
	// CacheUseDynamicCompression: bool, optional
	CacheUseDynamicCompression terra.BoolValue `hcl:"cache_use_dynamic_compression,attr"`
	// CustomForwardingPath: string, optional
	CustomForwardingPath terra.StringValue `hcl:"custom_forwarding_path,attr"`
	// ForwardingProtocol: string, optional
	ForwardingProtocol terra.StringValue `hcl:"forwarding_protocol,attr"`
}

type RedirectConfiguration struct {
	// CustomFragment: string, optional
	CustomFragment terra.StringValue `hcl:"custom_fragment,attr"`
	// CustomHost: string, optional
	CustomHost terra.StringValue `hcl:"custom_host,attr"`
	// CustomPath: string, optional
	CustomPath terra.StringValue `hcl:"custom_path,attr"`
	// CustomQueryString: string, optional
	CustomQueryString terra.StringValue `hcl:"custom_query_string,attr"`
	// RedirectProtocol: string, required
	RedirectProtocol terra.StringValue `hcl:"redirect_protocol,attr" validate:"required"`
	// RedirectType: string, required
	RedirectType terra.StringValue `hcl:"redirect_type,attr" validate:"required"`
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

type ExplicitResourceOrderAttributes struct {
	ref terra.Reference
}

func (ero ExplicitResourceOrderAttributes) InternalRef() terra.Reference {
	return ero.ref
}

func (ero ExplicitResourceOrderAttributes) InternalWithRef(ref terra.Reference) ExplicitResourceOrderAttributes {
	return ExplicitResourceOrderAttributes{ref: ref}
}

func (ero ExplicitResourceOrderAttributes) InternalTokens() hclwrite.Tokens {
	return ero.ref.InternalTokens()
}

func (ero ExplicitResourceOrderAttributes) BackendPoolHealthProbeIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ero.ref.Append("backend_pool_health_probe_ids"))
}

func (ero ExplicitResourceOrderAttributes) BackendPoolIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ero.ref.Append("backend_pool_ids"))
}

func (ero ExplicitResourceOrderAttributes) BackendPoolLoadBalancingIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ero.ref.Append("backend_pool_load_balancing_ids"))
}

func (ero ExplicitResourceOrderAttributes) FrontendEndpointIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ero.ref.Append("frontend_endpoint_ids"))
}

func (ero ExplicitResourceOrderAttributes) RoutingRuleIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ero.ref.Append("routing_rule_ids"))
}

type BackendPoolAttributes struct {
	ref terra.Reference
}

func (bp BackendPoolAttributes) InternalRef() terra.Reference {
	return bp.ref
}

func (bp BackendPoolAttributes) InternalWithRef(ref terra.Reference) BackendPoolAttributes {
	return BackendPoolAttributes{ref: ref}
}

func (bp BackendPoolAttributes) InternalTokens() hclwrite.Tokens {
	return bp.ref.InternalTokens()
}

func (bp BackendPoolAttributes) HealthProbeName() terra.StringValue {
	return terra.ReferenceString(bp.ref.Append("health_probe_name"))
}

func (bp BackendPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceString(bp.ref.Append("id"))
}

func (bp BackendPoolAttributes) LoadBalancingName() terra.StringValue {
	return terra.ReferenceString(bp.ref.Append("load_balancing_name"))
}

func (bp BackendPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceString(bp.ref.Append("name"))
}

func (bp BackendPoolAttributes) Backend() terra.ListValue[BackendAttributes] {
	return terra.ReferenceList[BackendAttributes](bp.ref.Append("backend"))
}

type BackendAttributes struct {
	ref terra.Reference
}

func (b BackendAttributes) InternalRef() terra.Reference {
	return b.ref
}

func (b BackendAttributes) InternalWithRef(ref terra.Reference) BackendAttributes {
	return BackendAttributes{ref: ref}
}

func (b BackendAttributes) InternalTokens() hclwrite.Tokens {
	return b.ref.InternalTokens()
}

func (b BackendAttributes) Address() terra.StringValue {
	return terra.ReferenceString(b.ref.Append("address"))
}

func (b BackendAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(b.ref.Append("enabled"))
}

func (b BackendAttributes) HostHeader() terra.StringValue {
	return terra.ReferenceString(b.ref.Append("host_header"))
}

func (b BackendAttributes) HttpPort() terra.NumberValue {
	return terra.ReferenceNumber(b.ref.Append("http_port"))
}

func (b BackendAttributes) HttpsPort() terra.NumberValue {
	return terra.ReferenceNumber(b.ref.Append("https_port"))
}

func (b BackendAttributes) Priority() terra.NumberValue {
	return terra.ReferenceNumber(b.ref.Append("priority"))
}

func (b BackendAttributes) Weight() terra.NumberValue {
	return terra.ReferenceNumber(b.ref.Append("weight"))
}

type BackendPoolHealthProbeAttributes struct {
	ref terra.Reference
}

func (bphp BackendPoolHealthProbeAttributes) InternalRef() terra.Reference {
	return bphp.ref
}

func (bphp BackendPoolHealthProbeAttributes) InternalWithRef(ref terra.Reference) BackendPoolHealthProbeAttributes {
	return BackendPoolHealthProbeAttributes{ref: ref}
}

func (bphp BackendPoolHealthProbeAttributes) InternalTokens() hclwrite.Tokens {
	return bphp.ref.InternalTokens()
}

func (bphp BackendPoolHealthProbeAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(bphp.ref.Append("enabled"))
}

func (bphp BackendPoolHealthProbeAttributes) Id() terra.StringValue {
	return terra.ReferenceString(bphp.ref.Append("id"))
}

func (bphp BackendPoolHealthProbeAttributes) IntervalInSeconds() terra.NumberValue {
	return terra.ReferenceNumber(bphp.ref.Append("interval_in_seconds"))
}

func (bphp BackendPoolHealthProbeAttributes) Name() terra.StringValue {
	return terra.ReferenceString(bphp.ref.Append("name"))
}

func (bphp BackendPoolHealthProbeAttributes) Path() terra.StringValue {
	return terra.ReferenceString(bphp.ref.Append("path"))
}

func (bphp BackendPoolHealthProbeAttributes) ProbeMethod() terra.StringValue {
	return terra.ReferenceString(bphp.ref.Append("probe_method"))
}

func (bphp BackendPoolHealthProbeAttributes) Protocol() terra.StringValue {
	return terra.ReferenceString(bphp.ref.Append("protocol"))
}

type BackendPoolLoadBalancingAttributes struct {
	ref terra.Reference
}

func (bplb BackendPoolLoadBalancingAttributes) InternalRef() terra.Reference {
	return bplb.ref
}

func (bplb BackendPoolLoadBalancingAttributes) InternalWithRef(ref terra.Reference) BackendPoolLoadBalancingAttributes {
	return BackendPoolLoadBalancingAttributes{ref: ref}
}

func (bplb BackendPoolLoadBalancingAttributes) InternalTokens() hclwrite.Tokens {
	return bplb.ref.InternalTokens()
}

func (bplb BackendPoolLoadBalancingAttributes) AdditionalLatencyMilliseconds() terra.NumberValue {
	return terra.ReferenceNumber(bplb.ref.Append("additional_latency_milliseconds"))
}

func (bplb BackendPoolLoadBalancingAttributes) Id() terra.StringValue {
	return terra.ReferenceString(bplb.ref.Append("id"))
}

func (bplb BackendPoolLoadBalancingAttributes) Name() terra.StringValue {
	return terra.ReferenceString(bplb.ref.Append("name"))
}

func (bplb BackendPoolLoadBalancingAttributes) SampleSize() terra.NumberValue {
	return terra.ReferenceNumber(bplb.ref.Append("sample_size"))
}

func (bplb BackendPoolLoadBalancingAttributes) SuccessfulSamplesRequired() terra.NumberValue {
	return terra.ReferenceNumber(bplb.ref.Append("successful_samples_required"))
}

type BackendPoolSettingsAttributes struct {
	ref terra.Reference
}

func (bps BackendPoolSettingsAttributes) InternalRef() terra.Reference {
	return bps.ref
}

func (bps BackendPoolSettingsAttributes) InternalWithRef(ref terra.Reference) BackendPoolSettingsAttributes {
	return BackendPoolSettingsAttributes{ref: ref}
}

func (bps BackendPoolSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return bps.ref.InternalTokens()
}

func (bps BackendPoolSettingsAttributes) BackendPoolsSendReceiveTimeoutSeconds() terra.NumberValue {
	return terra.ReferenceNumber(bps.ref.Append("backend_pools_send_receive_timeout_seconds"))
}

func (bps BackendPoolSettingsAttributes) EnforceBackendPoolsCertificateNameCheck() terra.BoolValue {
	return terra.ReferenceBool(bps.ref.Append("enforce_backend_pools_certificate_name_check"))
}

type FrontendEndpointAttributes struct {
	ref terra.Reference
}

func (fe FrontendEndpointAttributes) InternalRef() terra.Reference {
	return fe.ref
}

func (fe FrontendEndpointAttributes) InternalWithRef(ref terra.Reference) FrontendEndpointAttributes {
	return FrontendEndpointAttributes{ref: ref}
}

func (fe FrontendEndpointAttributes) InternalTokens() hclwrite.Tokens {
	return fe.ref.InternalTokens()
}

func (fe FrontendEndpointAttributes) HostName() terra.StringValue {
	return terra.ReferenceString(fe.ref.Append("host_name"))
}

func (fe FrontendEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceString(fe.ref.Append("id"))
}

func (fe FrontendEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceString(fe.ref.Append("name"))
}

func (fe FrontendEndpointAttributes) SessionAffinityEnabled() terra.BoolValue {
	return terra.ReferenceBool(fe.ref.Append("session_affinity_enabled"))
}

func (fe FrontendEndpointAttributes) SessionAffinityTtlSeconds() terra.NumberValue {
	return terra.ReferenceNumber(fe.ref.Append("session_affinity_ttl_seconds"))
}

func (fe FrontendEndpointAttributes) WebApplicationFirewallPolicyLinkId() terra.StringValue {
	return terra.ReferenceString(fe.ref.Append("web_application_firewall_policy_link_id"))
}

type RoutingRuleAttributes struct {
	ref terra.Reference
}

func (rr RoutingRuleAttributes) InternalRef() terra.Reference {
	return rr.ref
}

func (rr RoutingRuleAttributes) InternalWithRef(ref terra.Reference) RoutingRuleAttributes {
	return RoutingRuleAttributes{ref: ref}
}

func (rr RoutingRuleAttributes) InternalTokens() hclwrite.Tokens {
	return rr.ref.InternalTokens()
}

func (rr RoutingRuleAttributes) AcceptedProtocols() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](rr.ref.Append("accepted_protocols"))
}

func (rr RoutingRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(rr.ref.Append("enabled"))
}

func (rr RoutingRuleAttributes) FrontendEndpoints() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](rr.ref.Append("frontend_endpoints"))
}

func (rr RoutingRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rr.ref.Append("id"))
}

func (rr RoutingRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceString(rr.ref.Append("name"))
}

func (rr RoutingRuleAttributes) PatternsToMatch() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](rr.ref.Append("patterns_to_match"))
}

func (rr RoutingRuleAttributes) ForwardingConfiguration() terra.ListValue[ForwardingConfigurationAttributes] {
	return terra.ReferenceList[ForwardingConfigurationAttributes](rr.ref.Append("forwarding_configuration"))
}

func (rr RoutingRuleAttributes) RedirectConfiguration() terra.ListValue[RedirectConfigurationAttributes] {
	return terra.ReferenceList[RedirectConfigurationAttributes](rr.ref.Append("redirect_configuration"))
}

type ForwardingConfigurationAttributes struct {
	ref terra.Reference
}

func (fc ForwardingConfigurationAttributes) InternalRef() terra.Reference {
	return fc.ref
}

func (fc ForwardingConfigurationAttributes) InternalWithRef(ref terra.Reference) ForwardingConfigurationAttributes {
	return ForwardingConfigurationAttributes{ref: ref}
}

func (fc ForwardingConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return fc.ref.InternalTokens()
}

func (fc ForwardingConfigurationAttributes) BackendPoolName() terra.StringValue {
	return terra.ReferenceString(fc.ref.Append("backend_pool_name"))
}

func (fc ForwardingConfigurationAttributes) CacheDuration() terra.StringValue {
	return terra.ReferenceString(fc.ref.Append("cache_duration"))
}

func (fc ForwardingConfigurationAttributes) CacheEnabled() terra.BoolValue {
	return terra.ReferenceBool(fc.ref.Append("cache_enabled"))
}

func (fc ForwardingConfigurationAttributes) CacheQueryParameterStripDirective() terra.StringValue {
	return terra.ReferenceString(fc.ref.Append("cache_query_parameter_strip_directive"))
}

func (fc ForwardingConfigurationAttributes) CacheQueryParameters() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](fc.ref.Append("cache_query_parameters"))
}

func (fc ForwardingConfigurationAttributes) CacheUseDynamicCompression() terra.BoolValue {
	return terra.ReferenceBool(fc.ref.Append("cache_use_dynamic_compression"))
}

func (fc ForwardingConfigurationAttributes) CustomForwardingPath() terra.StringValue {
	return terra.ReferenceString(fc.ref.Append("custom_forwarding_path"))
}

func (fc ForwardingConfigurationAttributes) ForwardingProtocol() terra.StringValue {
	return terra.ReferenceString(fc.ref.Append("forwarding_protocol"))
}

type RedirectConfigurationAttributes struct {
	ref terra.Reference
}

func (rc RedirectConfigurationAttributes) InternalRef() terra.Reference {
	return rc.ref
}

func (rc RedirectConfigurationAttributes) InternalWithRef(ref terra.Reference) RedirectConfigurationAttributes {
	return RedirectConfigurationAttributes{ref: ref}
}

func (rc RedirectConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return rc.ref.InternalTokens()
}

func (rc RedirectConfigurationAttributes) CustomFragment() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("custom_fragment"))
}

func (rc RedirectConfigurationAttributes) CustomHost() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("custom_host"))
}

func (rc RedirectConfigurationAttributes) CustomPath() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("custom_path"))
}

func (rc RedirectConfigurationAttributes) CustomQueryString() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("custom_query_string"))
}

func (rc RedirectConfigurationAttributes) RedirectProtocol() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("redirect_protocol"))
}

func (rc RedirectConfigurationAttributes) RedirectType() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("redirect_type"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type ExplicitResourceOrderState struct {
	BackendPoolHealthProbeIds   []string `json:"backend_pool_health_probe_ids"`
	BackendPoolIds              []string `json:"backend_pool_ids"`
	BackendPoolLoadBalancingIds []string `json:"backend_pool_load_balancing_ids"`
	FrontendEndpointIds         []string `json:"frontend_endpoint_ids"`
	RoutingRuleIds              []string `json:"routing_rule_ids"`
}

type BackendPoolState struct {
	HealthProbeName   string         `json:"health_probe_name"`
	Id                string         `json:"id"`
	LoadBalancingName string         `json:"load_balancing_name"`
	Name              string         `json:"name"`
	Backend           []BackendState `json:"backend"`
}

type BackendState struct {
	Address    string  `json:"address"`
	Enabled    bool    `json:"enabled"`
	HostHeader string  `json:"host_header"`
	HttpPort   float64 `json:"http_port"`
	HttpsPort  float64 `json:"https_port"`
	Priority   float64 `json:"priority"`
	Weight     float64 `json:"weight"`
}

type BackendPoolHealthProbeState struct {
	Enabled           bool    `json:"enabled"`
	Id                string  `json:"id"`
	IntervalInSeconds float64 `json:"interval_in_seconds"`
	Name              string  `json:"name"`
	Path              string  `json:"path"`
	ProbeMethod       string  `json:"probe_method"`
	Protocol          string  `json:"protocol"`
}

type BackendPoolLoadBalancingState struct {
	AdditionalLatencyMilliseconds float64 `json:"additional_latency_milliseconds"`
	Id                            string  `json:"id"`
	Name                          string  `json:"name"`
	SampleSize                    float64 `json:"sample_size"`
	SuccessfulSamplesRequired     float64 `json:"successful_samples_required"`
}

type BackendPoolSettingsState struct {
	BackendPoolsSendReceiveTimeoutSeconds   float64 `json:"backend_pools_send_receive_timeout_seconds"`
	EnforceBackendPoolsCertificateNameCheck bool    `json:"enforce_backend_pools_certificate_name_check"`
}

type FrontendEndpointState struct {
	HostName                           string  `json:"host_name"`
	Id                                 string  `json:"id"`
	Name                               string  `json:"name"`
	SessionAffinityEnabled             bool    `json:"session_affinity_enabled"`
	SessionAffinityTtlSeconds          float64 `json:"session_affinity_ttl_seconds"`
	WebApplicationFirewallPolicyLinkId string  `json:"web_application_firewall_policy_link_id"`
}

type RoutingRuleState struct {
	AcceptedProtocols       []string                       `json:"accepted_protocols"`
	Enabled                 bool                           `json:"enabled"`
	FrontendEndpoints       []string                       `json:"frontend_endpoints"`
	Id                      string                         `json:"id"`
	Name                    string                         `json:"name"`
	PatternsToMatch         []string                       `json:"patterns_to_match"`
	ForwardingConfiguration []ForwardingConfigurationState `json:"forwarding_configuration"`
	RedirectConfiguration   []RedirectConfigurationState   `json:"redirect_configuration"`
}

type ForwardingConfigurationState struct {
	BackendPoolName                   string   `json:"backend_pool_name"`
	CacheDuration                     string   `json:"cache_duration"`
	CacheEnabled                      bool     `json:"cache_enabled"`
	CacheQueryParameterStripDirective string   `json:"cache_query_parameter_strip_directive"`
	CacheQueryParameters              []string `json:"cache_query_parameters"`
	CacheUseDynamicCompression        bool     `json:"cache_use_dynamic_compression"`
	CustomForwardingPath              string   `json:"custom_forwarding_path"`
	ForwardingProtocol                string   `json:"forwarding_protocol"`
}

type RedirectConfigurationState struct {
	CustomFragment    string `json:"custom_fragment"`
	CustomHost        string `json:"custom_host"`
	CustomPath        string `json:"custom_path"`
	CustomQueryString string `json:"custom_query_string"`
	RedirectProtocol  string `json:"redirect_protocol"`
	RedirectType      string `json:"redirect_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
