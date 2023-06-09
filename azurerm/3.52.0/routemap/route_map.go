// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package routemap

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Rule struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NextStepIfMatched: string, optional
	NextStepIfMatched terra.StringValue `hcl:"next_step_if_matched,attr"`
	// Action: min=0
	Action []Action `hcl:"action,block" validate:"min=0"`
	// MatchCriterion: min=0
	MatchCriterion []MatchCriterion `hcl:"match_criterion,block" validate:"min=0"`
}

type Action struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Parameter: min=1
	Parameter []Parameter `hcl:"parameter,block" validate:"min=1"`
}

type Parameter struct {
	// AsPath: list of string, optional
	AsPath terra.ListValue[terra.StringValue] `hcl:"as_path,attr"`
	// Community: list of string, optional
	Community terra.ListValue[terra.StringValue] `hcl:"community,attr"`
	// RoutePrefix: list of string, optional
	RoutePrefix terra.ListValue[terra.StringValue] `hcl:"route_prefix,attr"`
}

type MatchCriterion struct {
	// AsPath: list of string, optional
	AsPath terra.ListValue[terra.StringValue] `hcl:"as_path,attr"`
	// Community: list of string, optional
	Community terra.ListValue[terra.StringValue] `hcl:"community,attr"`
	// MatchCondition: string, required
	MatchCondition terra.StringValue `hcl:"match_condition,attr" validate:"required"`
	// RoutePrefix: list of string, optional
	RoutePrefix terra.ListValue[terra.StringValue] `hcl:"route_prefix,attr"`
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

func (r RuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r RuleAttributes) NextStepIfMatched() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("next_step_if_matched"))
}

func (r RuleAttributes) Action() terra.ListValue[ActionAttributes] {
	return terra.ReferenceAsList[ActionAttributes](r.ref.Append("action"))
}

func (r RuleAttributes) MatchCriterion() terra.ListValue[MatchCriterionAttributes] {
	return terra.ReferenceAsList[MatchCriterionAttributes](r.ref.Append("match_criterion"))
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

func (a ActionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("type"))
}

func (a ActionAttributes) Parameter() terra.ListValue[ParameterAttributes] {
	return terra.ReferenceAsList[ParameterAttributes](a.ref.Append("parameter"))
}

type ParameterAttributes struct {
	ref terra.Reference
}

func (p ParameterAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParameterAttributes) InternalWithRef(ref terra.Reference) ParameterAttributes {
	return ParameterAttributes{ref: ref}
}

func (p ParameterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ParameterAttributes) AsPath() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](p.ref.Append("as_path"))
}

func (p ParameterAttributes) Community() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](p.ref.Append("community"))
}

func (p ParameterAttributes) RoutePrefix() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](p.ref.Append("route_prefix"))
}

type MatchCriterionAttributes struct {
	ref terra.Reference
}

func (mc MatchCriterionAttributes) InternalRef() (terra.Reference, error) {
	return mc.ref, nil
}

func (mc MatchCriterionAttributes) InternalWithRef(ref terra.Reference) MatchCriterionAttributes {
	return MatchCriterionAttributes{ref: ref}
}

func (mc MatchCriterionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mc.ref.InternalTokens()
}

func (mc MatchCriterionAttributes) AsPath() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mc.ref.Append("as_path"))
}

func (mc MatchCriterionAttributes) Community() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mc.ref.Append("community"))
}

func (mc MatchCriterionAttributes) MatchCondition() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("match_condition"))
}

func (mc MatchCriterionAttributes) RoutePrefix() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mc.ref.Append("route_prefix"))
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

type RuleState struct {
	Name              string                `json:"name"`
	NextStepIfMatched string                `json:"next_step_if_matched"`
	Action            []ActionState         `json:"action"`
	MatchCriterion    []MatchCriterionState `json:"match_criterion"`
}

type ActionState struct {
	Type      string           `json:"type"`
	Parameter []ParameterState `json:"parameter"`
}

type ParameterState struct {
	AsPath      []string `json:"as_path"`
	Community   []string `json:"community"`
	RoutePrefix []string `json:"route_prefix"`
}

type MatchCriterionState struct {
	AsPath         []string `json:"as_path"`
	Community      []string `json:"community"`
	MatchCondition string   `json:"match_condition"`
	RoutePrefix    []string `json:"route_prefix"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
