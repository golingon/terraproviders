// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datafirewallpolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Dns struct{}

type ThreatIntelligenceAllowlist struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
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

func (d DnsAttributes) NetworkRuleFqdnEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("network_rule_fqdn_enabled"))
}

func (d DnsAttributes) ProxyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("proxy_enabled"))
}

func (d DnsAttributes) Servers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("servers"))
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

func (tia ThreatIntelligenceAllowlistAttributes) Fqdns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tia.ref.Append("fqdns"))
}

func (tia ThreatIntelligenceAllowlistAttributes) IpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tia.ref.Append("ip_addresses"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DnsState struct {
	NetworkRuleFqdnEnabled bool     `json:"network_rule_fqdn_enabled"`
	ProxyEnabled           bool     `json:"proxy_enabled"`
	Servers                []string `json:"servers"`
}

type ThreatIntelligenceAllowlistState struct {
	Fqdns       []string `json:"fqdns"`
	IpAddresses []string `json:"ip_addresses"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
