// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package orgpolicypolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Spec struct {
	// InheritFromParent: bool, optional
	InheritFromParent terra.BoolValue `hcl:"inherit_from_parent,attr"`
	// Reset: bool, optional
	Reset terra.BoolValue `hcl:"reset,attr"`
	// Rules: min=0
	Rules []Rules `hcl:"rules,block" validate:"min=0"`
}

type Rules struct {
	// AllowAll: string, optional
	AllowAll terra.StringValue `hcl:"allow_all,attr"`
	// DenyAll: string, optional
	DenyAll terra.StringValue `hcl:"deny_all,attr"`
	// Enforce: string, optional
	Enforce terra.StringValue `hcl:"enforce,attr"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
	// Values: optional
	Values *Values `hcl:"values,block"`
}

type Condition struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Expression: string, optional
	Expression terra.StringValue `hcl:"expression,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Title: string, optional
	Title terra.StringValue `hcl:"title,attr"`
}

type Values struct {
	// AllowedValues: list of string, optional
	AllowedValues terra.ListValue[terra.StringValue] `hcl:"allowed_values,attr"`
	// DeniedValues: list of string, optional
	DeniedValues terra.ListValue[terra.StringValue] `hcl:"denied_values,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type SpecAttributes struct {
	ref terra.Reference
}

func (s SpecAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SpecAttributes) InternalWithRef(ref terra.Reference) SpecAttributes {
	return SpecAttributes{ref: ref}
}

func (s SpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SpecAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("etag"))
}

func (s SpecAttributes) InheritFromParent() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("inherit_from_parent"))
}

func (s SpecAttributes) Reset() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("reset"))
}

func (s SpecAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("update_time"))
}

func (s SpecAttributes) Rules() terra.ListValue[RulesAttributes] {
	return terra.ReferenceAsList[RulesAttributes](s.ref.Append("rules"))
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

func (r RulesAttributes) AllowAll() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("allow_all"))
}

func (r RulesAttributes) DenyAll() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("deny_all"))
}

func (r RulesAttributes) Enforce() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("enforce"))
}

func (r RulesAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](r.ref.Append("condition"))
}

func (r RulesAttributes) Values() terra.ListValue[ValuesAttributes] {
	return terra.ReferenceAsList[ValuesAttributes](r.ref.Append("values"))
}

type ConditionAttributes struct {
	ref terra.Reference
}

func (c ConditionAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionAttributes) InternalWithRef(ref terra.Reference) ConditionAttributes {
	return ConditionAttributes{ref: ref}
}

func (c ConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("description"))
}

func (c ConditionAttributes) Expression() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("expression"))
}

func (c ConditionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("location"))
}

func (c ConditionAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("title"))
}

type ValuesAttributes struct {
	ref terra.Reference
}

func (v ValuesAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v ValuesAttributes) InternalWithRef(ref terra.Reference) ValuesAttributes {
	return ValuesAttributes{ref: ref}
}

func (v ValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v ValuesAttributes) AllowedValues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](v.ref.Append("allowed_values"))
}

func (v ValuesAttributes) DeniedValues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](v.ref.Append("denied_values"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type SpecState struct {
	Etag              string       `json:"etag"`
	InheritFromParent bool         `json:"inherit_from_parent"`
	Reset             bool         `json:"reset"`
	UpdateTime        string       `json:"update_time"`
	Rules             []RulesState `json:"rules"`
}

type RulesState struct {
	AllowAll  string           `json:"allow_all"`
	DenyAll   string           `json:"deny_all"`
	Enforce   string           `json:"enforce"`
	Condition []ConditionState `json:"condition"`
	Values    []ValuesState    `json:"values"`
}

type ConditionState struct {
	Description string `json:"description"`
	Expression  string `json:"expression"`
	Location    string `json:"location"`
	Title       string `json:"title"`
}

type ValuesState struct {
	AllowedValues []string `json:"allowed_values"`
	DeniedValues  []string `json:"denied_values"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
