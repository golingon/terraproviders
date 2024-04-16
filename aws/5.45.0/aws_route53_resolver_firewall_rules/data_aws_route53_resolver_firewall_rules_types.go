// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_route53_resolver_firewall_rules

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataFirewallRulesAttributes struct {
	ref terra.Reference
}

func (fr DataFirewallRulesAttributes) InternalRef() (terra.Reference, error) {
	return fr.ref, nil
}

func (fr DataFirewallRulesAttributes) InternalWithRef(ref terra.Reference) DataFirewallRulesAttributes {
	return DataFirewallRulesAttributes{ref: ref}
}

func (fr DataFirewallRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fr.ref.InternalTokens()
}

func (fr DataFirewallRulesAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("action"))
}

func (fr DataFirewallRulesAttributes) BlockOverrideDnsType() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("block_override_dns_type"))
}

func (fr DataFirewallRulesAttributes) BlockOverrideDomain() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("block_override_domain"))
}

func (fr DataFirewallRulesAttributes) BlockOverrideTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(fr.ref.Append("block_override_ttl"))
}

func (fr DataFirewallRulesAttributes) BlockResponse() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("block_response"))
}

func (fr DataFirewallRulesAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("creation_time"))
}

func (fr DataFirewallRulesAttributes) CreatorRequestId() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("creator_request_id"))
}

func (fr DataFirewallRulesAttributes) FirewallDomainListId() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("firewall_domain_list_id"))
}

func (fr DataFirewallRulesAttributes) FirewallRuleGroupId() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("firewall_rule_group_id"))
}

func (fr DataFirewallRulesAttributes) ModificationTime() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("modification_time"))
}

func (fr DataFirewallRulesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("name"))
}

func (fr DataFirewallRulesAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(fr.ref.Append("priority"))
}

type DataFirewallRulesState struct {
	Action               string  `json:"action"`
	BlockOverrideDnsType string  `json:"block_override_dns_type"`
	BlockOverrideDomain  string  `json:"block_override_domain"`
	BlockOverrideTtl     float64 `json:"block_override_ttl"`
	BlockResponse        string  `json:"block_response"`
	CreationTime         string  `json:"creation_time"`
	CreatorRequestId     string  `json:"creator_request_id"`
	FirewallDomainListId string  `json:"firewall_domain_list_id"`
	FirewallRuleGroupId  string  `json:"firewall_rule_group_id"`
	ModificationTime     string  `json:"modification_time"`
	Name                 string  `json:"name"`
	Priority             float64 `json:"priority"`
}
