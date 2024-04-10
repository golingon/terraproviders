// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dataroute53resolverfirewallrules

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type FirewallRules struct{}

type FirewallRulesAttributes struct {
	ref terra.Reference
}

func (fr FirewallRulesAttributes) InternalRef() (terra.Reference, error) {
	return fr.ref, nil
}

func (fr FirewallRulesAttributes) InternalWithRef(ref terra.Reference) FirewallRulesAttributes {
	return FirewallRulesAttributes{ref: ref}
}

func (fr FirewallRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fr.ref.InternalTokens()
}

func (fr FirewallRulesAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("action"))
}

func (fr FirewallRulesAttributes) BlockOverrideDnsType() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("block_override_dns_type"))
}

func (fr FirewallRulesAttributes) BlockOverrideDomain() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("block_override_domain"))
}

func (fr FirewallRulesAttributes) BlockOverrideTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(fr.ref.Append("block_override_ttl"))
}

func (fr FirewallRulesAttributes) BlockResponse() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("block_response"))
}

func (fr FirewallRulesAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("creation_time"))
}

func (fr FirewallRulesAttributes) CreatorRequestId() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("creator_request_id"))
}

func (fr FirewallRulesAttributes) FirewallDomainListId() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("firewall_domain_list_id"))
}

func (fr FirewallRulesAttributes) FirewallRuleGroupId() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("firewall_rule_group_id"))
}

func (fr FirewallRulesAttributes) ModificationTime() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("modification_time"))
}

func (fr FirewallRulesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("name"))
}

func (fr FirewallRulesAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(fr.ref.Append("priority"))
}

type FirewallRulesState struct {
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
