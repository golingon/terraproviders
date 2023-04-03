// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package webapplicationfirewallpolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CustomRules struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// RuleType: string, required
	RuleType terra.StringValue `hcl:"rule_type,attr" validate:"required"`
	// MatchConditions: min=1
	MatchConditions []MatchConditions `hcl:"match_conditions,block" validate:"min=1"`
}

type MatchConditions struct {
	// MatchValues: list of string, required
	MatchValues terra.ListValue[terra.StringValue] `hcl:"match_values,attr" validate:"required"`
	// NegationCondition: bool, optional
	NegationCondition terra.BoolValue `hcl:"negation_condition,attr"`
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Transforms: set of string, optional
	Transforms terra.SetValue[terra.StringValue] `hcl:"transforms,attr"`
	// MatchVariables: min=1
	MatchVariables []MatchVariables `hcl:"match_variables,block" validate:"min=1"`
}

type MatchVariables struct {
	// Selector: string, optional
	Selector terra.StringValue `hcl:"selector,attr"`
	// VariableName: string, required
	VariableName terra.StringValue `hcl:"variable_name,attr" validate:"required"`
}

type ManagedRules struct {
	// Exclusion: min=0
	Exclusion []Exclusion `hcl:"exclusion,block" validate:"min=0"`
	// ManagedRuleSet: min=1
	ManagedRuleSet []ManagedRuleSet `hcl:"managed_rule_set,block" validate:"min=1"`
}

type Exclusion struct {
	// MatchVariable: string, required
	MatchVariable terra.StringValue `hcl:"match_variable,attr" validate:"required"`
	// Selector: string, required
	Selector terra.StringValue `hcl:"selector,attr" validate:"required"`
	// SelectorMatchOperator: string, required
	SelectorMatchOperator terra.StringValue `hcl:"selector_match_operator,attr" validate:"required"`
	// ExcludedRuleSet: optional
	ExcludedRuleSet *ExcludedRuleSet `hcl:"excluded_rule_set,block"`
}

type ExcludedRuleSet struct {
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// RuleGroup: min=0
	RuleGroup []RuleGroup `hcl:"rule_group,block" validate:"min=0"`
}

type RuleGroup struct {
	// ExcludedRules: list of string, optional
	ExcludedRules terra.ListValue[terra.StringValue] `hcl:"excluded_rules,attr"`
	// RuleGroupName: string, required
	RuleGroupName terra.StringValue `hcl:"rule_group_name,attr" validate:"required"`
}

type ManagedRuleSet struct {
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// RuleGroupOverride: min=0
	RuleGroupOverride []RuleGroupOverride `hcl:"rule_group_override,block" validate:"min=0"`
}

type RuleGroupOverride struct {
	// DisabledRules: list of string, optional
	DisabledRules terra.ListValue[terra.StringValue] `hcl:"disabled_rules,attr"`
	// RuleGroupName: string, required
	RuleGroupName terra.StringValue `hcl:"rule_group_name,attr" validate:"required"`
	// Rule: min=0
	Rule []Rule `hcl:"rule,block" validate:"min=0"`
}

type Rule struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
}

type PolicySettings struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// FileUploadLimitInMb: number, optional
	FileUploadLimitInMb terra.NumberValue `hcl:"file_upload_limit_in_mb,attr"`
	// MaxRequestBodySizeInKb: number, optional
	MaxRequestBodySizeInKb terra.NumberValue `hcl:"max_request_body_size_in_kb,attr"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// RequestBodyCheck: bool, optional
	RequestBodyCheck terra.BoolValue `hcl:"request_body_check,attr"`
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

type CustomRulesAttributes struct {
	ref terra.Reference
}

func (cr CustomRulesAttributes) InternalRef() (terra.Reference, error) {
	return cr.ref, nil
}

func (cr CustomRulesAttributes) InternalWithRef(ref terra.Reference) CustomRulesAttributes {
	return CustomRulesAttributes{ref: ref}
}

func (cr CustomRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cr.ref.InternalTokens()
}

func (cr CustomRulesAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("action"))
}

func (cr CustomRulesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("name"))
}

func (cr CustomRulesAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(cr.ref.Append("priority"))
}

func (cr CustomRulesAttributes) RuleType() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("rule_type"))
}

func (cr CustomRulesAttributes) MatchConditions() terra.ListValue[MatchConditionsAttributes] {
	return terra.ReferenceAsList[MatchConditionsAttributes](cr.ref.Append("match_conditions"))
}

type MatchConditionsAttributes struct {
	ref terra.Reference
}

func (mc MatchConditionsAttributes) InternalRef() (terra.Reference, error) {
	return mc.ref, nil
}

func (mc MatchConditionsAttributes) InternalWithRef(ref terra.Reference) MatchConditionsAttributes {
	return MatchConditionsAttributes{ref: ref}
}

func (mc MatchConditionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mc.ref.InternalTokens()
}

func (mc MatchConditionsAttributes) MatchValues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mc.ref.Append("match_values"))
}

func (mc MatchConditionsAttributes) NegationCondition() terra.BoolValue {
	return terra.ReferenceAsBool(mc.ref.Append("negation_condition"))
}

func (mc MatchConditionsAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("operator"))
}

func (mc MatchConditionsAttributes) Transforms() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mc.ref.Append("transforms"))
}

func (mc MatchConditionsAttributes) MatchVariables() terra.ListValue[MatchVariablesAttributes] {
	return terra.ReferenceAsList[MatchVariablesAttributes](mc.ref.Append("match_variables"))
}

type MatchVariablesAttributes struct {
	ref terra.Reference
}

func (mv MatchVariablesAttributes) InternalRef() (terra.Reference, error) {
	return mv.ref, nil
}

func (mv MatchVariablesAttributes) InternalWithRef(ref terra.Reference) MatchVariablesAttributes {
	return MatchVariablesAttributes{ref: ref}
}

func (mv MatchVariablesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mv.ref.InternalTokens()
}

func (mv MatchVariablesAttributes) Selector() terra.StringValue {
	return terra.ReferenceAsString(mv.ref.Append("selector"))
}

func (mv MatchVariablesAttributes) VariableName() terra.StringValue {
	return terra.ReferenceAsString(mv.ref.Append("variable_name"))
}

type ManagedRulesAttributes struct {
	ref terra.Reference
}

func (mr ManagedRulesAttributes) InternalRef() (terra.Reference, error) {
	return mr.ref, nil
}

func (mr ManagedRulesAttributes) InternalWithRef(ref terra.Reference) ManagedRulesAttributes {
	return ManagedRulesAttributes{ref: ref}
}

func (mr ManagedRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mr.ref.InternalTokens()
}

func (mr ManagedRulesAttributes) Exclusion() terra.ListValue[ExclusionAttributes] {
	return terra.ReferenceAsList[ExclusionAttributes](mr.ref.Append("exclusion"))
}

func (mr ManagedRulesAttributes) ManagedRuleSet() terra.ListValue[ManagedRuleSetAttributes] {
	return terra.ReferenceAsList[ManagedRuleSetAttributes](mr.ref.Append("managed_rule_set"))
}

type ExclusionAttributes struct {
	ref terra.Reference
}

func (e ExclusionAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ExclusionAttributes) InternalWithRef(ref terra.Reference) ExclusionAttributes {
	return ExclusionAttributes{ref: ref}
}

func (e ExclusionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ExclusionAttributes) MatchVariable() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("match_variable"))
}

func (e ExclusionAttributes) Selector() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("selector"))
}

func (e ExclusionAttributes) SelectorMatchOperator() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("selector_match_operator"))
}

func (e ExclusionAttributes) ExcludedRuleSet() terra.ListValue[ExcludedRuleSetAttributes] {
	return terra.ReferenceAsList[ExcludedRuleSetAttributes](e.ref.Append("excluded_rule_set"))
}

type ExcludedRuleSetAttributes struct {
	ref terra.Reference
}

func (ers ExcludedRuleSetAttributes) InternalRef() (terra.Reference, error) {
	return ers.ref, nil
}

func (ers ExcludedRuleSetAttributes) InternalWithRef(ref terra.Reference) ExcludedRuleSetAttributes {
	return ExcludedRuleSetAttributes{ref: ref}
}

func (ers ExcludedRuleSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ers.ref.InternalTokens()
}

func (ers ExcludedRuleSetAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ers.ref.Append("type"))
}

func (ers ExcludedRuleSetAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ers.ref.Append("version"))
}

func (ers ExcludedRuleSetAttributes) RuleGroup() terra.ListValue[RuleGroupAttributes] {
	return terra.ReferenceAsList[RuleGroupAttributes](ers.ref.Append("rule_group"))
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

func (rg RuleGroupAttributes) ExcludedRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rg.ref.Append("excluded_rules"))
}

func (rg RuleGroupAttributes) RuleGroupName() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("rule_group_name"))
}

type ManagedRuleSetAttributes struct {
	ref terra.Reference
}

func (mrs ManagedRuleSetAttributes) InternalRef() (terra.Reference, error) {
	return mrs.ref, nil
}

func (mrs ManagedRuleSetAttributes) InternalWithRef(ref terra.Reference) ManagedRuleSetAttributes {
	return ManagedRuleSetAttributes{ref: ref}
}

func (mrs ManagedRuleSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mrs.ref.InternalTokens()
}

func (mrs ManagedRuleSetAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(mrs.ref.Append("type"))
}

func (mrs ManagedRuleSetAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(mrs.ref.Append("version"))
}

func (mrs ManagedRuleSetAttributes) RuleGroupOverride() terra.ListValue[RuleGroupOverrideAttributes] {
	return terra.ReferenceAsList[RuleGroupOverrideAttributes](mrs.ref.Append("rule_group_override"))
}

type RuleGroupOverrideAttributes struct {
	ref terra.Reference
}

func (rgo RuleGroupOverrideAttributes) InternalRef() (terra.Reference, error) {
	return rgo.ref, nil
}

func (rgo RuleGroupOverrideAttributes) InternalWithRef(ref terra.Reference) RuleGroupOverrideAttributes {
	return RuleGroupOverrideAttributes{ref: ref}
}

func (rgo RuleGroupOverrideAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rgo.ref.InternalTokens()
}

func (rgo RuleGroupOverrideAttributes) DisabledRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rgo.ref.Append("disabled_rules"))
}

func (rgo RuleGroupOverrideAttributes) RuleGroupName() terra.StringValue {
	return terra.ReferenceAsString(rgo.ref.Append("rule_group_name"))
}

func (rgo RuleGroupOverrideAttributes) Rule() terra.ListValue[RuleAttributes] {
	return terra.ReferenceAsList[RuleAttributes](rgo.ref.Append("rule"))
}

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("action"))
}

func (r RuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("enabled"))
}

func (r RuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("id"))
}

type PolicySettingsAttributes struct {
	ref terra.Reference
}

func (ps PolicySettingsAttributes) InternalRef() (terra.Reference, error) {
	return ps.ref, nil
}

func (ps PolicySettingsAttributes) InternalWithRef(ref terra.Reference) PolicySettingsAttributes {
	return PolicySettingsAttributes{ref: ref}
}

func (ps PolicySettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ps.ref.InternalTokens()
}

func (ps PolicySettingsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("enabled"))
}

func (ps PolicySettingsAttributes) FileUploadLimitInMb() terra.NumberValue {
	return terra.ReferenceAsNumber(ps.ref.Append("file_upload_limit_in_mb"))
}

func (ps PolicySettingsAttributes) MaxRequestBodySizeInKb() terra.NumberValue {
	return terra.ReferenceAsNumber(ps.ref.Append("max_request_body_size_in_kb"))
}

func (ps PolicySettingsAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("mode"))
}

func (ps PolicySettingsAttributes) RequestBodyCheck() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("request_body_check"))
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

type CustomRulesState struct {
	Action          string                 `json:"action"`
	Name            string                 `json:"name"`
	Priority        float64                `json:"priority"`
	RuleType        string                 `json:"rule_type"`
	MatchConditions []MatchConditionsState `json:"match_conditions"`
}

type MatchConditionsState struct {
	MatchValues       []string              `json:"match_values"`
	NegationCondition bool                  `json:"negation_condition"`
	Operator          string                `json:"operator"`
	Transforms        []string              `json:"transforms"`
	MatchVariables    []MatchVariablesState `json:"match_variables"`
}

type MatchVariablesState struct {
	Selector     string `json:"selector"`
	VariableName string `json:"variable_name"`
}

type ManagedRulesState struct {
	Exclusion      []ExclusionState      `json:"exclusion"`
	ManagedRuleSet []ManagedRuleSetState `json:"managed_rule_set"`
}

type ExclusionState struct {
	MatchVariable         string                 `json:"match_variable"`
	Selector              string                 `json:"selector"`
	SelectorMatchOperator string                 `json:"selector_match_operator"`
	ExcludedRuleSet       []ExcludedRuleSetState `json:"excluded_rule_set"`
}

type ExcludedRuleSetState struct {
	Type      string           `json:"type"`
	Version   string           `json:"version"`
	RuleGroup []RuleGroupState `json:"rule_group"`
}

type RuleGroupState struct {
	ExcludedRules []string `json:"excluded_rules"`
	RuleGroupName string   `json:"rule_group_name"`
}

type ManagedRuleSetState struct {
	Type              string                   `json:"type"`
	Version           string                   `json:"version"`
	RuleGroupOverride []RuleGroupOverrideState `json:"rule_group_override"`
}

type RuleGroupOverrideState struct {
	DisabledRules []string    `json:"disabled_rules"`
	RuleGroupName string      `json:"rule_group_name"`
	Rule          []RuleState `json:"rule"`
}

type RuleState struct {
	Action  string `json:"action"`
	Enabled bool   `json:"enabled"`
	Id      string `json:"id"`
}

type PolicySettingsState struct {
	Enabled                bool    `json:"enabled"`
	FileUploadLimitInMb    float64 `json:"file_upload_limit_in_mb"`
	MaxRequestBodySizeInKb float64 `json:"max_request_body_size_in_kb"`
	Mode                   string  `json:"mode"`
	RequestBodyCheck       bool    `json:"request_body_check"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
