// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkfirewallrulegroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EncryptionConfiguration struct {
	// KeyId: string, optional
	KeyId terra.StringValue `hcl:"key_id,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type RuleGroup struct {
	// ReferenceSets: optional
	ReferenceSets *ReferenceSets `hcl:"reference_sets,block"`
	// RuleVariables: optional
	RuleVariables *RuleVariables `hcl:"rule_variables,block"`
	// RulesSource: required
	RulesSource *RulesSource `hcl:"rules_source,block" validate:"required"`
	// StatefulRuleOptions: optional
	StatefulRuleOptions *StatefulRuleOptions `hcl:"stateful_rule_options,block"`
}

type ReferenceSets struct {
	// IpSetReferences: min=0,max=5
	IpSetReferences []IpSetReferences `hcl:"ip_set_references,block" validate:"min=0,max=5"`
}

type IpSetReferences struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// IpSetReference: min=1
	IpSetReference []IpSetReference `hcl:"ip_set_reference,block" validate:"min=1"`
}

type IpSetReference struct {
	// ReferenceArn: string, required
	ReferenceArn terra.StringValue `hcl:"reference_arn,attr" validate:"required"`
}

type RuleVariables struct {
	// IpSets: min=0
	IpSets []IpSets `hcl:"ip_sets,block" validate:"min=0"`
	// PortSets: min=0
	PortSets []PortSets `hcl:"port_sets,block" validate:"min=0"`
}

type IpSets struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// IpSet: required
	IpSet *IpSet `hcl:"ip_set,block" validate:"required"`
}

type IpSet struct {
	// Definition: set of string, required
	Definition terra.SetValue[terra.StringValue] `hcl:"definition,attr" validate:"required"`
}

type PortSets struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// PortSet: required
	PortSet *PortSet `hcl:"port_set,block" validate:"required"`
}

type PortSet struct {
	// Definition: set of string, required
	Definition terra.SetValue[terra.StringValue] `hcl:"definition,attr" validate:"required"`
}

type RulesSource struct {
	// RulesString: string, optional
	RulesString terra.StringValue `hcl:"rules_string,attr"`
	// RulesSourceList: optional
	RulesSourceList *RulesSourceList `hcl:"rules_source_list,block"`
	// StatefulRule: min=0
	StatefulRule []StatefulRule `hcl:"stateful_rule,block" validate:"min=0"`
	// StatelessRulesAndCustomActions: optional
	StatelessRulesAndCustomActions *StatelessRulesAndCustomActions `hcl:"stateless_rules_and_custom_actions,block"`
}

type RulesSourceList struct {
	// GeneratedRulesType: string, required
	GeneratedRulesType terra.StringValue `hcl:"generated_rules_type,attr" validate:"required"`
	// TargetTypes: set of string, required
	TargetTypes terra.SetValue[terra.StringValue] `hcl:"target_types,attr" validate:"required"`
	// Targets: set of string, required
	Targets terra.SetValue[terra.StringValue] `hcl:"targets,attr" validate:"required"`
}

type StatefulRule struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Header: required
	Header *Header `hcl:"header,block" validate:"required"`
	// RuleOption: min=1
	RuleOption []RuleOption `hcl:"rule_option,block" validate:"min=1"`
}

type Header struct {
	// Destination: string, required
	Destination terra.StringValue `hcl:"destination,attr" validate:"required"`
	// DestinationPort: string, required
	DestinationPort terra.StringValue `hcl:"destination_port,attr" validate:"required"`
	// Direction: string, required
	Direction terra.StringValue `hcl:"direction,attr" validate:"required"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
	// SourcePort: string, required
	SourcePort terra.StringValue `hcl:"source_port,attr" validate:"required"`
}

type RuleOption struct {
	// Keyword: string, required
	Keyword terra.StringValue `hcl:"keyword,attr" validate:"required"`
	// Settings: set of string, optional
	Settings terra.SetValue[terra.StringValue] `hcl:"settings,attr"`
}

type StatelessRulesAndCustomActions struct {
	// CustomAction: min=0
	CustomAction []CustomAction `hcl:"custom_action,block" validate:"min=0"`
	// StatelessRule: min=1
	StatelessRule []StatelessRule `hcl:"stateless_rule,block" validate:"min=1"`
}

type CustomAction struct {
	// ActionName: string, required
	ActionName terra.StringValue `hcl:"action_name,attr" validate:"required"`
	// ActionDefinition: required
	ActionDefinition *ActionDefinition `hcl:"action_definition,block" validate:"required"`
}

type ActionDefinition struct {
	// PublishMetricAction: required
	PublishMetricAction *PublishMetricAction `hcl:"publish_metric_action,block" validate:"required"`
}

type PublishMetricAction struct {
	// Dimension: min=1
	Dimension []Dimension `hcl:"dimension,block" validate:"min=1"`
}

type Dimension struct {
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type StatelessRule struct {
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// RuleDefinition: required
	RuleDefinition *RuleDefinition `hcl:"rule_definition,block" validate:"required"`
}

type RuleDefinition struct {
	// Actions: set of string, required
	Actions terra.SetValue[terra.StringValue] `hcl:"actions,attr" validate:"required"`
	// MatchAttributes: required
	MatchAttributes *MatchAttributes `hcl:"match_attributes,block" validate:"required"`
}

type MatchAttributes struct {
	// Protocols: set of number, optional
	Protocols terra.SetValue[terra.NumberValue] `hcl:"protocols,attr"`
	// Destination: min=0
	Destination []Destination `hcl:"destination,block" validate:"min=0"`
	// DestinationPort: min=0
	DestinationPort []DestinationPort `hcl:"destination_port,block" validate:"min=0"`
	// Source: min=0
	Source []Source `hcl:"source,block" validate:"min=0"`
	// SourcePort: min=0
	SourcePort []SourcePort `hcl:"source_port,block" validate:"min=0"`
	// TcpFlag: min=0
	TcpFlag []TcpFlag `hcl:"tcp_flag,block" validate:"min=0"`
}

type Destination struct {
	// AddressDefinition: string, required
	AddressDefinition terra.StringValue `hcl:"address_definition,attr" validate:"required"`
}

type DestinationPort struct {
	// FromPort: number, required
	FromPort terra.NumberValue `hcl:"from_port,attr" validate:"required"`
	// ToPort: number, optional
	ToPort terra.NumberValue `hcl:"to_port,attr"`
}

type Source struct {
	// AddressDefinition: string, required
	AddressDefinition terra.StringValue `hcl:"address_definition,attr" validate:"required"`
}

type SourcePort struct {
	// FromPort: number, required
	FromPort terra.NumberValue `hcl:"from_port,attr" validate:"required"`
	// ToPort: number, optional
	ToPort terra.NumberValue `hcl:"to_port,attr"`
}

type TcpFlag struct {
	// Flags: set of string, required
	Flags terra.SetValue[terra.StringValue] `hcl:"flags,attr" validate:"required"`
	// Masks: set of string, optional
	Masks terra.SetValue[terra.StringValue] `hcl:"masks,attr"`
}

type StatefulRuleOptions struct {
	// RuleOrder: string, required
	RuleOrder terra.StringValue `hcl:"rule_order,attr" validate:"required"`
}

type EncryptionConfigurationAttributes struct {
	ref terra.Reference
}

func (ec EncryptionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EncryptionConfigurationAttributes) InternalWithRef(ref terra.Reference) EncryptionConfigurationAttributes {
	return EncryptionConfigurationAttributes{ref: ref}
}

func (ec EncryptionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EncryptionConfigurationAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("key_id"))
}

func (ec EncryptionConfigurationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("type"))
}

type RuleGroupAttributes struct {
	ref terra.Reference
}

func (rg RuleGroupAttributes) InternalRef() (terra.Reference, error) {
	return rg.ref, nil
}

func (rg RuleGroupAttributes) InternalWithRef(ref terra.Reference) RuleGroupAttributes {
	return RuleGroupAttributes{ref: ref}
}

func (rg RuleGroupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rg.ref.InternalTokens()
}

func (rg RuleGroupAttributes) ReferenceSets() terra.ListValue[ReferenceSetsAttributes] {
	return terra.ReferenceAsList[ReferenceSetsAttributes](rg.ref.Append("reference_sets"))
}

func (rg RuleGroupAttributes) RuleVariables() terra.ListValue[RuleVariablesAttributes] {
	return terra.ReferenceAsList[RuleVariablesAttributes](rg.ref.Append("rule_variables"))
}

func (rg RuleGroupAttributes) RulesSource() terra.ListValue[RulesSourceAttributes] {
	return terra.ReferenceAsList[RulesSourceAttributes](rg.ref.Append("rules_source"))
}

func (rg RuleGroupAttributes) StatefulRuleOptions() terra.ListValue[StatefulRuleOptionsAttributes] {
	return terra.ReferenceAsList[StatefulRuleOptionsAttributes](rg.ref.Append("stateful_rule_options"))
}

type ReferenceSetsAttributes struct {
	ref terra.Reference
}

func (rs ReferenceSetsAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs ReferenceSetsAttributes) InternalWithRef(ref terra.Reference) ReferenceSetsAttributes {
	return ReferenceSetsAttributes{ref: ref}
}

func (rs ReferenceSetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs ReferenceSetsAttributes) IpSetReferences() terra.SetValue[IpSetReferencesAttributes] {
	return terra.ReferenceAsSet[IpSetReferencesAttributes](rs.ref.Append("ip_set_references"))
}

type IpSetReferencesAttributes struct {
	ref terra.Reference
}

func (isr IpSetReferencesAttributes) InternalRef() (terra.Reference, error) {
	return isr.ref, nil
}

func (isr IpSetReferencesAttributes) InternalWithRef(ref terra.Reference) IpSetReferencesAttributes {
	return IpSetReferencesAttributes{ref: ref}
}

func (isr IpSetReferencesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return isr.ref.InternalTokens()
}

func (isr IpSetReferencesAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(isr.ref.Append("key"))
}

func (isr IpSetReferencesAttributes) IpSetReference() terra.ListValue[IpSetReferenceAttributes] {
	return terra.ReferenceAsList[IpSetReferenceAttributes](isr.ref.Append("ip_set_reference"))
}

type IpSetReferenceAttributes struct {
	ref terra.Reference
}

func (isr IpSetReferenceAttributes) InternalRef() (terra.Reference, error) {
	return isr.ref, nil
}

func (isr IpSetReferenceAttributes) InternalWithRef(ref terra.Reference) IpSetReferenceAttributes {
	return IpSetReferenceAttributes{ref: ref}
}

func (isr IpSetReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return isr.ref.InternalTokens()
}

func (isr IpSetReferenceAttributes) ReferenceArn() terra.StringValue {
	return terra.ReferenceAsString(isr.ref.Append("reference_arn"))
}

type RuleVariablesAttributes struct {
	ref terra.Reference
}

func (rv RuleVariablesAttributes) InternalRef() (terra.Reference, error) {
	return rv.ref, nil
}

func (rv RuleVariablesAttributes) InternalWithRef(ref terra.Reference) RuleVariablesAttributes {
	return RuleVariablesAttributes{ref: ref}
}

func (rv RuleVariablesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rv.ref.InternalTokens()
}

func (rv RuleVariablesAttributes) IpSets() terra.SetValue[IpSetsAttributes] {
	return terra.ReferenceAsSet[IpSetsAttributes](rv.ref.Append("ip_sets"))
}

func (rv RuleVariablesAttributes) PortSets() terra.SetValue[PortSetsAttributes] {
	return terra.ReferenceAsSet[PortSetsAttributes](rv.ref.Append("port_sets"))
}

type IpSetsAttributes struct {
	ref terra.Reference
}

func (is IpSetsAttributes) InternalRef() (terra.Reference, error) {
	return is.ref, nil
}

func (is IpSetsAttributes) InternalWithRef(ref terra.Reference) IpSetsAttributes {
	return IpSetsAttributes{ref: ref}
}

func (is IpSetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return is.ref.InternalTokens()
}

func (is IpSetsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(is.ref.Append("key"))
}

func (is IpSetsAttributes) IpSet() terra.ListValue[IpSetAttributes] {
	return terra.ReferenceAsList[IpSetAttributes](is.ref.Append("ip_set"))
}

type IpSetAttributes struct {
	ref terra.Reference
}

func (is IpSetAttributes) InternalRef() (terra.Reference, error) {
	return is.ref, nil
}

func (is IpSetAttributes) InternalWithRef(ref terra.Reference) IpSetAttributes {
	return IpSetAttributes{ref: ref}
}

func (is IpSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return is.ref.InternalTokens()
}

func (is IpSetAttributes) Definition() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](is.ref.Append("definition"))
}

type PortSetsAttributes struct {
	ref terra.Reference
}

func (ps PortSetsAttributes) InternalRef() (terra.Reference, error) {
	return ps.ref, nil
}

func (ps PortSetsAttributes) InternalWithRef(ref terra.Reference) PortSetsAttributes {
	return PortSetsAttributes{ref: ref}
}

func (ps PortSetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ps.ref.InternalTokens()
}

func (ps PortSetsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("key"))
}

func (ps PortSetsAttributes) PortSet() terra.ListValue[PortSetAttributes] {
	return terra.ReferenceAsList[PortSetAttributes](ps.ref.Append("port_set"))
}

type PortSetAttributes struct {
	ref terra.Reference
}

func (ps PortSetAttributes) InternalRef() (terra.Reference, error) {
	return ps.ref, nil
}

func (ps PortSetAttributes) InternalWithRef(ref terra.Reference) PortSetAttributes {
	return PortSetAttributes{ref: ref}
}

func (ps PortSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ps.ref.InternalTokens()
}

func (ps PortSetAttributes) Definition() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ps.ref.Append("definition"))
}

type RulesSourceAttributes struct {
	ref terra.Reference
}

func (rs RulesSourceAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs RulesSourceAttributes) InternalWithRef(ref terra.Reference) RulesSourceAttributes {
	return RulesSourceAttributes{ref: ref}
}

func (rs RulesSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs RulesSourceAttributes) RulesString() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("rules_string"))
}

func (rs RulesSourceAttributes) RulesSourceList() terra.ListValue[RulesSourceListAttributes] {
	return terra.ReferenceAsList[RulesSourceListAttributes](rs.ref.Append("rules_source_list"))
}

func (rs RulesSourceAttributes) StatefulRule() terra.ListValue[StatefulRuleAttributes] {
	return terra.ReferenceAsList[StatefulRuleAttributes](rs.ref.Append("stateful_rule"))
}

func (rs RulesSourceAttributes) StatelessRulesAndCustomActions() terra.ListValue[StatelessRulesAndCustomActionsAttributes] {
	return terra.ReferenceAsList[StatelessRulesAndCustomActionsAttributes](rs.ref.Append("stateless_rules_and_custom_actions"))
}

type RulesSourceListAttributes struct {
	ref terra.Reference
}

func (rsl RulesSourceListAttributes) InternalRef() (terra.Reference, error) {
	return rsl.ref, nil
}

func (rsl RulesSourceListAttributes) InternalWithRef(ref terra.Reference) RulesSourceListAttributes {
	return RulesSourceListAttributes{ref: ref}
}

func (rsl RulesSourceListAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rsl.ref.InternalTokens()
}

func (rsl RulesSourceListAttributes) GeneratedRulesType() terra.StringValue {
	return terra.ReferenceAsString(rsl.ref.Append("generated_rules_type"))
}

func (rsl RulesSourceListAttributes) TargetTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rsl.ref.Append("target_types"))
}

func (rsl RulesSourceListAttributes) Targets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rsl.ref.Append("targets"))
}

type StatefulRuleAttributes struct {
	ref terra.Reference
}

func (sr StatefulRuleAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr StatefulRuleAttributes) InternalWithRef(ref terra.Reference) StatefulRuleAttributes {
	return StatefulRuleAttributes{ref: ref}
}

func (sr StatefulRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr StatefulRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("action"))
}

func (sr StatefulRuleAttributes) Header() terra.ListValue[HeaderAttributes] {
	return terra.ReferenceAsList[HeaderAttributes](sr.ref.Append("header"))
}

func (sr StatefulRuleAttributes) RuleOption() terra.SetValue[RuleOptionAttributes] {
	return terra.ReferenceAsSet[RuleOptionAttributes](sr.ref.Append("rule_option"))
}

type HeaderAttributes struct {
	ref terra.Reference
}

func (h HeaderAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h HeaderAttributes) InternalWithRef(ref terra.Reference) HeaderAttributes {
	return HeaderAttributes{ref: ref}
}

func (h HeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h HeaderAttributes) Destination() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("destination"))
}

func (h HeaderAttributes) DestinationPort() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("destination_port"))
}

func (h HeaderAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("direction"))
}

func (h HeaderAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("protocol"))
}

func (h HeaderAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("source"))
}

func (h HeaderAttributes) SourcePort() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("source_port"))
}

type RuleOptionAttributes struct {
	ref terra.Reference
}

func (ro RuleOptionAttributes) InternalRef() (terra.Reference, error) {
	return ro.ref, nil
}

func (ro RuleOptionAttributes) InternalWithRef(ref terra.Reference) RuleOptionAttributes {
	return RuleOptionAttributes{ref: ref}
}

func (ro RuleOptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ro.ref.InternalTokens()
}

func (ro RuleOptionAttributes) Keyword() terra.StringValue {
	return terra.ReferenceAsString(ro.ref.Append("keyword"))
}

func (ro RuleOptionAttributes) Settings() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ro.ref.Append("settings"))
}

type StatelessRulesAndCustomActionsAttributes struct {
	ref terra.Reference
}

func (sraca StatelessRulesAndCustomActionsAttributes) InternalRef() (terra.Reference, error) {
	return sraca.ref, nil
}

func (sraca StatelessRulesAndCustomActionsAttributes) InternalWithRef(ref terra.Reference) StatelessRulesAndCustomActionsAttributes {
	return StatelessRulesAndCustomActionsAttributes{ref: ref}
}

func (sraca StatelessRulesAndCustomActionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sraca.ref.InternalTokens()
}

func (sraca StatelessRulesAndCustomActionsAttributes) CustomAction() terra.SetValue[CustomActionAttributes] {
	return terra.ReferenceAsSet[CustomActionAttributes](sraca.ref.Append("custom_action"))
}

func (sraca StatelessRulesAndCustomActionsAttributes) StatelessRule() terra.SetValue[StatelessRuleAttributes] {
	return terra.ReferenceAsSet[StatelessRuleAttributes](sraca.ref.Append("stateless_rule"))
}

type CustomActionAttributes struct {
	ref terra.Reference
}

func (ca CustomActionAttributes) InternalRef() (terra.Reference, error) {
	return ca.ref, nil
}

func (ca CustomActionAttributes) InternalWithRef(ref terra.Reference) CustomActionAttributes {
	return CustomActionAttributes{ref: ref}
}

func (ca CustomActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ca.ref.InternalTokens()
}

func (ca CustomActionAttributes) ActionName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("action_name"))
}

func (ca CustomActionAttributes) ActionDefinition() terra.ListValue[ActionDefinitionAttributes] {
	return terra.ReferenceAsList[ActionDefinitionAttributes](ca.ref.Append("action_definition"))
}

type ActionDefinitionAttributes struct {
	ref terra.Reference
}

func (ad ActionDefinitionAttributes) InternalRef() (terra.Reference, error) {
	return ad.ref, nil
}

func (ad ActionDefinitionAttributes) InternalWithRef(ref terra.Reference) ActionDefinitionAttributes {
	return ActionDefinitionAttributes{ref: ref}
}

func (ad ActionDefinitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ad.ref.InternalTokens()
}

func (ad ActionDefinitionAttributes) PublishMetricAction() terra.ListValue[PublishMetricActionAttributes] {
	return terra.ReferenceAsList[PublishMetricActionAttributes](ad.ref.Append("publish_metric_action"))
}

type PublishMetricActionAttributes struct {
	ref terra.Reference
}

func (pma PublishMetricActionAttributes) InternalRef() (terra.Reference, error) {
	return pma.ref, nil
}

func (pma PublishMetricActionAttributes) InternalWithRef(ref terra.Reference) PublishMetricActionAttributes {
	return PublishMetricActionAttributes{ref: ref}
}

func (pma PublishMetricActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pma.ref.InternalTokens()
}

func (pma PublishMetricActionAttributes) Dimension() terra.SetValue[DimensionAttributes] {
	return terra.ReferenceAsSet[DimensionAttributes](pma.ref.Append("dimension"))
}

type DimensionAttributes struct {
	ref terra.Reference
}

func (d DimensionAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DimensionAttributes) InternalWithRef(ref terra.Reference) DimensionAttributes {
	return DimensionAttributes{ref: ref}
}

func (d DimensionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DimensionAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("value"))
}

type StatelessRuleAttributes struct {
	ref terra.Reference
}

func (sr StatelessRuleAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr StatelessRuleAttributes) InternalWithRef(ref terra.Reference) StatelessRuleAttributes {
	return StatelessRuleAttributes{ref: ref}
}

func (sr StatelessRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr StatelessRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(sr.ref.Append("priority"))
}

func (sr StatelessRuleAttributes) RuleDefinition() terra.ListValue[RuleDefinitionAttributes] {
	return terra.ReferenceAsList[RuleDefinitionAttributes](sr.ref.Append("rule_definition"))
}

type RuleDefinitionAttributes struct {
	ref terra.Reference
}

func (rd RuleDefinitionAttributes) InternalRef() (terra.Reference, error) {
	return rd.ref, nil
}

func (rd RuleDefinitionAttributes) InternalWithRef(ref terra.Reference) RuleDefinitionAttributes {
	return RuleDefinitionAttributes{ref: ref}
}

func (rd RuleDefinitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rd.ref.InternalTokens()
}

func (rd RuleDefinitionAttributes) Actions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rd.ref.Append("actions"))
}

func (rd RuleDefinitionAttributes) MatchAttributes() terra.ListValue[MatchAttributesAttributes] {
	return terra.ReferenceAsList[MatchAttributesAttributes](rd.ref.Append("match_attributes"))
}

type MatchAttributesAttributes struct {
	ref terra.Reference
}

func (ma MatchAttributesAttributes) InternalRef() (terra.Reference, error) {
	return ma.ref, nil
}

func (ma MatchAttributesAttributes) InternalWithRef(ref terra.Reference) MatchAttributesAttributes {
	return MatchAttributesAttributes{ref: ref}
}

func (ma MatchAttributesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ma.ref.InternalTokens()
}

func (ma MatchAttributesAttributes) Protocols() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](ma.ref.Append("protocols"))
}

func (ma MatchAttributesAttributes) Destination() terra.SetValue[DestinationAttributes] {
	return terra.ReferenceAsSet[DestinationAttributes](ma.ref.Append("destination"))
}

func (ma MatchAttributesAttributes) DestinationPort() terra.SetValue[DestinationPortAttributes] {
	return terra.ReferenceAsSet[DestinationPortAttributes](ma.ref.Append("destination_port"))
}

func (ma MatchAttributesAttributes) Source() terra.SetValue[SourceAttributes] {
	return terra.ReferenceAsSet[SourceAttributes](ma.ref.Append("source"))
}

func (ma MatchAttributesAttributes) SourcePort() terra.SetValue[SourcePortAttributes] {
	return terra.ReferenceAsSet[SourcePortAttributes](ma.ref.Append("source_port"))
}

func (ma MatchAttributesAttributes) TcpFlag() terra.SetValue[TcpFlagAttributes] {
	return terra.ReferenceAsSet[TcpFlagAttributes](ma.ref.Append("tcp_flag"))
}

type DestinationAttributes struct {
	ref terra.Reference
}

func (d DestinationAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DestinationAttributes) InternalWithRef(ref terra.Reference) DestinationAttributes {
	return DestinationAttributes{ref: ref}
}

func (d DestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DestinationAttributes) AddressDefinition() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("address_definition"))
}

type DestinationPortAttributes struct {
	ref terra.Reference
}

func (dp DestinationPortAttributes) InternalRef() (terra.Reference, error) {
	return dp.ref, nil
}

func (dp DestinationPortAttributes) InternalWithRef(ref terra.Reference) DestinationPortAttributes {
	return DestinationPortAttributes{ref: ref}
}

func (dp DestinationPortAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dp.ref.InternalTokens()
}

func (dp DestinationPortAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(dp.ref.Append("from_port"))
}

func (dp DestinationPortAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(dp.ref.Append("to_port"))
}

type SourceAttributes struct {
	ref terra.Reference
}

func (s SourceAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SourceAttributes) InternalWithRef(ref terra.Reference) SourceAttributes {
	return SourceAttributes{ref: ref}
}

func (s SourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SourceAttributes) AddressDefinition() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("address_definition"))
}

type SourcePortAttributes struct {
	ref terra.Reference
}

func (sp SourcePortAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp SourcePortAttributes) InternalWithRef(ref terra.Reference) SourcePortAttributes {
	return SourcePortAttributes{ref: ref}
}

func (sp SourcePortAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp SourcePortAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("from_port"))
}

func (sp SourcePortAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("to_port"))
}

type TcpFlagAttributes struct {
	ref terra.Reference
}

func (tf TcpFlagAttributes) InternalRef() (terra.Reference, error) {
	return tf.ref, nil
}

func (tf TcpFlagAttributes) InternalWithRef(ref terra.Reference) TcpFlagAttributes {
	return TcpFlagAttributes{ref: ref}
}

func (tf TcpFlagAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tf.ref.InternalTokens()
}

func (tf TcpFlagAttributes) Flags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tf.ref.Append("flags"))
}

func (tf TcpFlagAttributes) Masks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tf.ref.Append("masks"))
}

type StatefulRuleOptionsAttributes struct {
	ref terra.Reference
}

func (sro StatefulRuleOptionsAttributes) InternalRef() (terra.Reference, error) {
	return sro.ref, nil
}

func (sro StatefulRuleOptionsAttributes) InternalWithRef(ref terra.Reference) StatefulRuleOptionsAttributes {
	return StatefulRuleOptionsAttributes{ref: ref}
}

func (sro StatefulRuleOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sro.ref.InternalTokens()
}

func (sro StatefulRuleOptionsAttributes) RuleOrder() terra.StringValue {
	return terra.ReferenceAsString(sro.ref.Append("rule_order"))
}

type EncryptionConfigurationState struct {
	KeyId string `json:"key_id"`
	Type  string `json:"type"`
}

type RuleGroupState struct {
	ReferenceSets       []ReferenceSetsState       `json:"reference_sets"`
	RuleVariables       []RuleVariablesState       `json:"rule_variables"`
	RulesSource         []RulesSourceState         `json:"rules_source"`
	StatefulRuleOptions []StatefulRuleOptionsState `json:"stateful_rule_options"`
}

type ReferenceSetsState struct {
	IpSetReferences []IpSetReferencesState `json:"ip_set_references"`
}

type IpSetReferencesState struct {
	Key            string                `json:"key"`
	IpSetReference []IpSetReferenceState `json:"ip_set_reference"`
}

type IpSetReferenceState struct {
	ReferenceArn string `json:"reference_arn"`
}

type RuleVariablesState struct {
	IpSets   []IpSetsState   `json:"ip_sets"`
	PortSets []PortSetsState `json:"port_sets"`
}

type IpSetsState struct {
	Key   string       `json:"key"`
	IpSet []IpSetState `json:"ip_set"`
}

type IpSetState struct {
	Definition []string `json:"definition"`
}

type PortSetsState struct {
	Key     string         `json:"key"`
	PortSet []PortSetState `json:"port_set"`
}

type PortSetState struct {
	Definition []string `json:"definition"`
}

type RulesSourceState struct {
	RulesString                    string                                `json:"rules_string"`
	RulesSourceList                []RulesSourceListState                `json:"rules_source_list"`
	StatefulRule                   []StatefulRuleState                   `json:"stateful_rule"`
	StatelessRulesAndCustomActions []StatelessRulesAndCustomActionsState `json:"stateless_rules_and_custom_actions"`
}

type RulesSourceListState struct {
	GeneratedRulesType string   `json:"generated_rules_type"`
	TargetTypes        []string `json:"target_types"`
	Targets            []string `json:"targets"`
}

type StatefulRuleState struct {
	Action     string            `json:"action"`
	Header     []HeaderState     `json:"header"`
	RuleOption []RuleOptionState `json:"rule_option"`
}

type HeaderState struct {
	Destination     string `json:"destination"`
	DestinationPort string `json:"destination_port"`
	Direction       string `json:"direction"`
	Protocol        string `json:"protocol"`
	Source          string `json:"source"`
	SourcePort      string `json:"source_port"`
}

type RuleOptionState struct {
	Keyword  string   `json:"keyword"`
	Settings []string `json:"settings"`
}

type StatelessRulesAndCustomActionsState struct {
	CustomAction  []CustomActionState  `json:"custom_action"`
	StatelessRule []StatelessRuleState `json:"stateless_rule"`
}

type CustomActionState struct {
	ActionName       string                  `json:"action_name"`
	ActionDefinition []ActionDefinitionState `json:"action_definition"`
}

type ActionDefinitionState struct {
	PublishMetricAction []PublishMetricActionState `json:"publish_metric_action"`
}

type PublishMetricActionState struct {
	Dimension []DimensionState `json:"dimension"`
}

type DimensionState struct {
	Value string `json:"value"`
}

type StatelessRuleState struct {
	Priority       float64               `json:"priority"`
	RuleDefinition []RuleDefinitionState `json:"rule_definition"`
}

type RuleDefinitionState struct {
	Actions         []string               `json:"actions"`
	MatchAttributes []MatchAttributesState `json:"match_attributes"`
}

type MatchAttributesState struct {
	Protocols       []float64              `json:"protocols"`
	Destination     []DestinationState     `json:"destination"`
	DestinationPort []DestinationPortState `json:"destination_port"`
	Source          []SourceState          `json:"source"`
	SourcePort      []SourcePortState      `json:"source_port"`
	TcpFlag         []TcpFlagState         `json:"tcp_flag"`
}

type DestinationState struct {
	AddressDefinition string `json:"address_definition"`
}

type DestinationPortState struct {
	FromPort float64 `json:"from_port"`
	ToPort   float64 `json:"to_port"`
}

type SourceState struct {
	AddressDefinition string `json:"address_definition"`
}

type SourcePortState struct {
	FromPort float64 `json:"from_port"`
	ToPort   float64 `json:"to_port"`
}

type TcpFlagState struct {
	Flags []string `json:"flags"`
	Masks []string `json:"masks"`
}

type StatefulRuleOptionsState struct {
	RuleOrder string `json:"rule_order"`
}
