// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacomputerouternat

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type LogConfig struct{}

type Rules struct {
	// Action: min=0
	Action []Action `hcl:"action,block" validate:"min=0"`
}

type Action struct{}

type Subnetwork struct{}

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

func (lc LogConfigAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("filter"))
}

type RulesAttributes struct {
	ref terra.Reference
}

func (r RulesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RulesAttributes) InternalWithRef(ref terra.Reference) RulesAttributes {
	return RulesAttributes{ref: ref}
}

func (r RulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RulesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("description"))
}

func (r RulesAttributes) Match() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("match"))
}

func (r RulesAttributes) RuleNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("rule_number"))
}

func (r RulesAttributes) Action() terra.ListValue[ActionAttributes] {
	return terra.ReferenceAsList[ActionAttributes](r.ref.Append("action"))
}

type ActionAttributes struct {
	ref terra.Reference
}

func (a ActionAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ActionAttributes) InternalWithRef(ref terra.Reference) ActionAttributes {
	return ActionAttributes{ref: ref}
}

func (a ActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ActionAttributes) SourceNatActiveIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("source_nat_active_ips"))
}

func (a ActionAttributes) SourceNatDrainIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("source_nat_drain_ips"))
}

type SubnetworkAttributes struct {
	ref terra.Reference
}

func (s SubnetworkAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SubnetworkAttributes) InternalWithRef(ref terra.Reference) SubnetworkAttributes {
	return SubnetworkAttributes{ref: ref}
}

func (s SubnetworkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SubnetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SubnetworkAttributes) SecondaryIpRangeNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("secondary_ip_range_names"))
}

func (s SubnetworkAttributes) SourceIpRangesToNat() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("source_ip_ranges_to_nat"))
}

type LogConfigState struct {
	Enable bool   `json:"enable"`
	Filter string `json:"filter"`
}

type RulesState struct {
	Description string        `json:"description"`
	Match       string        `json:"match"`
	RuleNumber  float64       `json:"rule_number"`
	Action      []ActionState `json:"action"`
}

type ActionState struct {
	SourceNatActiveIps []string `json:"source_nat_active_ips"`
	SourceNatDrainIps  []string `json:"source_nat_drain_ips"`
}

type SubnetworkState struct {
	Name                  string   `json:"name"`
	SecondaryIpRangeNames []string `json:"secondary_ip_range_names"`
	SourceIpRangesToNat   []string `json:"source_ip_ranges_to_nat"`
}
