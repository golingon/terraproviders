// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package firewallpolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Dns struct {
	// ProxyEnabled: bool, optional
	ProxyEnabled terra.BoolValue `hcl:"proxy_enabled,attr"`
	// Servers: list of string, optional
	Servers terra.ListValue[terra.StringValue] `hcl:"servers,attr"`
}

type ExplicitProxy struct {
	// EnablePacFile: bool, optional
	EnablePacFile terra.BoolValue `hcl:"enable_pac_file,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// HttpPort: number, optional
	HttpPort terra.NumberValue `hcl:"http_port,attr"`
	// HttpsPort: number, optional
	HttpsPort terra.NumberValue `hcl:"https_port,attr"`
	// PacFile: string, optional
	PacFile terra.StringValue `hcl:"pac_file,attr"`
	// PacFilePort: number, optional
	PacFilePort terra.NumberValue `hcl:"pac_file_port,attr"`
}

type Identity struct {
	// IdentityIds: set of string, required
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Insights struct {
	// DefaultLogAnalyticsWorkspaceId: string, required
	DefaultLogAnalyticsWorkspaceId terra.StringValue `hcl:"default_log_analytics_workspace_id,attr" validate:"required"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// RetentionInDays: number, optional
	RetentionInDays terra.NumberValue `hcl:"retention_in_days,attr"`
	// LogAnalyticsWorkspace: min=0
	LogAnalyticsWorkspace []LogAnalyticsWorkspace `hcl:"log_analytics_workspace,block" validate:"min=0"`
}

type LogAnalyticsWorkspace struct {
	// FirewallLocation: string, required
	FirewallLocation terra.StringValue `hcl:"firewall_location,attr" validate:"required"`
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
}

type IntrusionDetection struct {
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// PrivateRanges: list of string, optional
	PrivateRanges terra.ListValue[terra.StringValue] `hcl:"private_ranges,attr"`
	// SignatureOverrides: min=0
	SignatureOverrides []SignatureOverrides `hcl:"signature_overrides,block" validate:"min=0"`
	// TrafficBypass: min=0
	TrafficBypass []TrafficBypass `hcl:"traffic_bypass,block" validate:"min=0"`
}

type SignatureOverrides struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
}

type TrafficBypass struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DestinationAddresses: set of string, optional
	DestinationAddresses terra.SetValue[terra.StringValue] `hcl:"destination_addresses,attr"`
	// DestinationIpGroups: set of string, optional
	DestinationIpGroups terra.SetValue[terra.StringValue] `hcl:"destination_ip_groups,attr"`
	// DestinationPorts: set of string, optional
	DestinationPorts terra.SetValue[terra.StringValue] `hcl:"destination_ports,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// SourceAddresses: set of string, optional
	SourceAddresses terra.SetValue[terra.StringValue] `hcl:"source_addresses,attr"`
	// SourceIpGroups: set of string, optional
	SourceIpGroups terra.SetValue[terra.StringValue] `hcl:"source_ip_groups,attr"`
}

type ThreatIntelligenceAllowlist struct {
	// Fqdns: set of string, optional
	Fqdns terra.SetValue[terra.StringValue] `hcl:"fqdns,attr"`
	// IpAddresses: set of string, optional
	IpAddresses terra.SetValue[terra.StringValue] `hcl:"ip_addresses,attr"`
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

type TlsCertificate struct {
	// KeyVaultSecretId: string, required
	KeyVaultSecretId terra.StringValue `hcl:"key_vault_secret_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type DnsAttributes struct {
	ref terra.Reference
}

func (d DnsAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DnsAttributes) InternalWithRef(ref terra.Reference) DnsAttributes {
	return DnsAttributes{ref: ref}
}

func (d DnsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DnsAttributes) ProxyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("proxy_enabled"))
}

func (d DnsAttributes) Servers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("servers"))
}

type ExplicitProxyAttributes struct {
	ref terra.Reference
}

func (ep ExplicitProxyAttributes) InternalRef() (terra.Reference, error) {
	return ep.ref, nil
}

func (ep ExplicitProxyAttributes) InternalWithRef(ref terra.Reference) ExplicitProxyAttributes {
	return ExplicitProxyAttributes{ref: ref}
}

func (ep ExplicitProxyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ep.ref.InternalTokens()
}

func (ep ExplicitProxyAttributes) EnablePacFile() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("enable_pac_file"))
}

func (ep ExplicitProxyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("enabled"))
}

func (ep ExplicitProxyAttributes) HttpPort() terra.NumberValue {
	return terra.ReferenceAsNumber(ep.ref.Append("http_port"))
}

func (ep ExplicitProxyAttributes) HttpsPort() terra.NumberValue {
	return terra.ReferenceAsNumber(ep.ref.Append("https_port"))
}

func (ep ExplicitProxyAttributes) PacFile() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("pac_file"))
}

func (ep ExplicitProxyAttributes) PacFilePort() terra.NumberValue {
	return terra.ReferenceAsNumber(ep.ref.Append("pac_file_port"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type InsightsAttributes struct {
	ref terra.Reference
}

func (i InsightsAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i InsightsAttributes) InternalWithRef(ref terra.Reference) InsightsAttributes {
	return InsightsAttributes{ref: ref}
}

func (i InsightsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i InsightsAttributes) DefaultLogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("default_log_analytics_workspace_id"))
}

func (i InsightsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("enabled"))
}

func (i InsightsAttributes) RetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("retention_in_days"))
}

func (i InsightsAttributes) LogAnalyticsWorkspace() terra.ListValue[LogAnalyticsWorkspaceAttributes] {
	return terra.ReferenceAsList[LogAnalyticsWorkspaceAttributes](i.ref.Append("log_analytics_workspace"))
}

type LogAnalyticsWorkspaceAttributes struct {
	ref terra.Reference
}

func (law LogAnalyticsWorkspaceAttributes) InternalRef() (terra.Reference, error) {
	return law.ref, nil
}

func (law LogAnalyticsWorkspaceAttributes) InternalWithRef(ref terra.Reference) LogAnalyticsWorkspaceAttributes {
	return LogAnalyticsWorkspaceAttributes{ref: ref}
}

func (law LogAnalyticsWorkspaceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return law.ref.InternalTokens()
}

func (law LogAnalyticsWorkspaceAttributes) FirewallLocation() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("firewall_location"))
}

func (law LogAnalyticsWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("id"))
}

type IntrusionDetectionAttributes struct {
	ref terra.Reference
}

func (id IntrusionDetectionAttributes) InternalRef() (terra.Reference, error) {
	return id.ref, nil
}

func (id IntrusionDetectionAttributes) InternalWithRef(ref terra.Reference) IntrusionDetectionAttributes {
	return IntrusionDetectionAttributes{ref: ref}
}

func (id IntrusionDetectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return id.ref.InternalTokens()
}

func (id IntrusionDetectionAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("mode"))
}

func (id IntrusionDetectionAttributes) PrivateRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](id.ref.Append("private_ranges"))
}

func (id IntrusionDetectionAttributes) SignatureOverrides() terra.ListValue[SignatureOverridesAttributes] {
	return terra.ReferenceAsList[SignatureOverridesAttributes](id.ref.Append("signature_overrides"))
}

func (id IntrusionDetectionAttributes) TrafficBypass() terra.ListValue[TrafficBypassAttributes] {
	return terra.ReferenceAsList[TrafficBypassAttributes](id.ref.Append("traffic_bypass"))
}

type SignatureOverridesAttributes struct {
	ref terra.Reference
}

func (so SignatureOverridesAttributes) InternalRef() (terra.Reference, error) {
	return so.ref, nil
}

func (so SignatureOverridesAttributes) InternalWithRef(ref terra.Reference) SignatureOverridesAttributes {
	return SignatureOverridesAttributes{ref: ref}
}

func (so SignatureOverridesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return so.ref.InternalTokens()
}

func (so SignatureOverridesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("id"))
}

func (so SignatureOverridesAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("state"))
}

type TrafficBypassAttributes struct {
	ref terra.Reference
}

func (tb TrafficBypassAttributes) InternalRef() (terra.Reference, error) {
	return tb.ref, nil
}

func (tb TrafficBypassAttributes) InternalWithRef(ref terra.Reference) TrafficBypassAttributes {
	return TrafficBypassAttributes{ref: ref}
}

func (tb TrafficBypassAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tb.ref.InternalTokens()
}

func (tb TrafficBypassAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(tb.ref.Append("description"))
}

func (tb TrafficBypassAttributes) DestinationAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tb.ref.Append("destination_addresses"))
}

func (tb TrafficBypassAttributes) DestinationIpGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tb.ref.Append("destination_ip_groups"))
}

func (tb TrafficBypassAttributes) DestinationPorts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tb.ref.Append("destination_ports"))
}

func (tb TrafficBypassAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tb.ref.Append("name"))
}

func (tb TrafficBypassAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(tb.ref.Append("protocol"))
}

func (tb TrafficBypassAttributes) SourceAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tb.ref.Append("source_addresses"))
}

func (tb TrafficBypassAttributes) SourceIpGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tb.ref.Append("source_ip_groups"))
}

type ThreatIntelligenceAllowlistAttributes struct {
	ref terra.Reference
}

func (tia ThreatIntelligenceAllowlistAttributes) InternalRef() (terra.Reference, error) {
	return tia.ref, nil
}

func (tia ThreatIntelligenceAllowlistAttributes) InternalWithRef(ref terra.Reference) ThreatIntelligenceAllowlistAttributes {
	return ThreatIntelligenceAllowlistAttributes{ref: ref}
}

func (tia ThreatIntelligenceAllowlistAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tia.ref.InternalTokens()
}

func (tia ThreatIntelligenceAllowlistAttributes) Fqdns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tia.ref.Append("fqdns"))
}

func (tia ThreatIntelligenceAllowlistAttributes) IpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tia.ref.Append("ip_addresses"))
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

type TlsCertificateAttributes struct {
	ref terra.Reference
}

func (tc TlsCertificateAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc TlsCertificateAttributes) InternalWithRef(ref terra.Reference) TlsCertificateAttributes {
	return TlsCertificateAttributes{ref: ref}
}

func (tc TlsCertificateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tc.ref.InternalTokens()
}

func (tc TlsCertificateAttributes) KeyVaultSecretId() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("key_vault_secret_id"))
}

func (tc TlsCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("name"))
}

type DnsState struct {
	ProxyEnabled bool     `json:"proxy_enabled"`
	Servers      []string `json:"servers"`
}

type ExplicitProxyState struct {
	EnablePacFile bool    `json:"enable_pac_file"`
	Enabled       bool    `json:"enabled"`
	HttpPort      float64 `json:"http_port"`
	HttpsPort     float64 `json:"https_port"`
	PacFile       string  `json:"pac_file"`
	PacFilePort   float64 `json:"pac_file_port"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	Type        string   `json:"type"`
}

type InsightsState struct {
	DefaultLogAnalyticsWorkspaceId string                       `json:"default_log_analytics_workspace_id"`
	Enabled                        bool                         `json:"enabled"`
	RetentionInDays                float64                      `json:"retention_in_days"`
	LogAnalyticsWorkspace          []LogAnalyticsWorkspaceState `json:"log_analytics_workspace"`
}

type LogAnalyticsWorkspaceState struct {
	FirewallLocation string `json:"firewall_location"`
	Id               string `json:"id"`
}

type IntrusionDetectionState struct {
	Mode               string                    `json:"mode"`
	PrivateRanges      []string                  `json:"private_ranges"`
	SignatureOverrides []SignatureOverridesState `json:"signature_overrides"`
	TrafficBypass      []TrafficBypassState      `json:"traffic_bypass"`
}

type SignatureOverridesState struct {
	Id    string `json:"id"`
	State string `json:"state"`
}

type TrafficBypassState struct {
	Description          string   `json:"description"`
	DestinationAddresses []string `json:"destination_addresses"`
	DestinationIpGroups  []string `json:"destination_ip_groups"`
	DestinationPorts     []string `json:"destination_ports"`
	Name                 string   `json:"name"`
	Protocol             string   `json:"protocol"`
	SourceAddresses      []string `json:"source_addresses"`
	SourceIpGroups       []string `json:"source_ip_groups"`
}

type ThreatIntelligenceAllowlistState struct {
	Fqdns       []string `json:"fqdns"`
	IpAddresses []string `json:"ip_addresses"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type TlsCertificateState struct {
	KeyVaultSecretId string `json:"key_vault_secret_id"`
	Name             string `json:"name"`
}
