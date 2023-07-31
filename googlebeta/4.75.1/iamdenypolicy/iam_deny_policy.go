// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package iamdenypolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Rules struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DenyRule: optional
	DenyRule *DenyRule `hcl:"deny_rule,block"`
}

type DenyRule struct {
	// DeniedPermissions: list of string, optional
	DeniedPermissions terra.ListValue[terra.StringValue] `hcl:"denied_permissions,attr"`
	// DeniedPrincipals: list of string, optional
	DeniedPrincipals terra.ListValue[terra.StringValue] `hcl:"denied_principals,attr"`
	// ExceptionPermissions: list of string, optional
	ExceptionPermissions terra.ListValue[terra.StringValue] `hcl:"exception_permissions,attr"`
	// ExceptionPrincipals: list of string, optional
	ExceptionPrincipals terra.ListValue[terra.StringValue] `hcl:"exception_principals,attr"`
	// DenialCondition: optional
	DenialCondition *DenialCondition `hcl:"denial_condition,block"`
}

type DenialCondition struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Expression: string, required
	Expression terra.StringValue `hcl:"expression,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Title: string, optional
	Title terra.StringValue `hcl:"title,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
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

func (r RulesAttributes) DenyRule() terra.ListValue[DenyRuleAttributes] {
	return terra.ReferenceAsList[DenyRuleAttributes](r.ref.Append("deny_rule"))
}

type DenyRuleAttributes struct {
	ref terra.Reference
}

func (dr DenyRuleAttributes) InternalRef() (terra.Reference, error) {
	return dr.ref, nil
}

func (dr DenyRuleAttributes) InternalWithRef(ref terra.Reference) DenyRuleAttributes {
	return DenyRuleAttributes{ref: ref}
}

func (dr DenyRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dr.ref.InternalTokens()
}

func (dr DenyRuleAttributes) DeniedPermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dr.ref.Append("denied_permissions"))
}

func (dr DenyRuleAttributes) DeniedPrincipals() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dr.ref.Append("denied_principals"))
}

func (dr DenyRuleAttributes) ExceptionPermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dr.ref.Append("exception_permissions"))
}

func (dr DenyRuleAttributes) ExceptionPrincipals() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dr.ref.Append("exception_principals"))
}

func (dr DenyRuleAttributes) DenialCondition() terra.ListValue[DenialConditionAttributes] {
	return terra.ReferenceAsList[DenialConditionAttributes](dr.ref.Append("denial_condition"))
}

type DenialConditionAttributes struct {
	ref terra.Reference
}

func (dc DenialConditionAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DenialConditionAttributes) InternalWithRef(ref terra.Reference) DenialConditionAttributes {
	return DenialConditionAttributes{ref: ref}
}

func (dc DenialConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DenialConditionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("description"))
}

func (dc DenialConditionAttributes) Expression() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("expression"))
}

func (dc DenialConditionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("location"))
}

func (dc DenialConditionAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("title"))
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

type RulesState struct {
	Description string          `json:"description"`
	DenyRule    []DenyRuleState `json:"deny_rule"`
}

type DenyRuleState struct {
	DeniedPermissions    []string               `json:"denied_permissions"`
	DeniedPrincipals     []string               `json:"denied_principals"`
	ExceptionPermissions []string               `json:"exception_permissions"`
	ExceptionPrincipals  []string               `json:"exception_principals"`
	DenialCondition      []DenialConditionState `json:"denial_condition"`
}

type DenialConditionState struct {
	Description string `json:"description"`
	Expression  string `json:"expression"`
	Location    string `json:"location"`
	Title       string `json:"title"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
