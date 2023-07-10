// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataiampolicydocument

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Statement struct {
	// Actions: set of string, optional
	Actions terra.SetValue[terra.StringValue] `hcl:"actions,attr"`
	// Effect: string, optional
	Effect terra.StringValue `hcl:"effect,attr"`
	// NotActions: set of string, optional
	NotActions terra.SetValue[terra.StringValue] `hcl:"not_actions,attr"`
	// NotResources: set of string, optional
	NotResources terra.SetValue[terra.StringValue] `hcl:"not_resources,attr"`
	// Resources: set of string, optional
	Resources terra.SetValue[terra.StringValue] `hcl:"resources,attr"`
	// Sid: string, optional
	Sid terra.StringValue `hcl:"sid,attr"`
	// Condition: min=0
	Condition []Condition `hcl:"condition,block" validate:"min=0"`
	// NotPrincipals: min=0
	NotPrincipals []NotPrincipals `hcl:"not_principals,block" validate:"min=0"`
	// Principals: min=0
	Principals []Principals `hcl:"principals,block" validate:"min=0"`
}

type Condition struct {
	// Test: string, required
	Test terra.StringValue `hcl:"test,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
	// Variable: string, required
	Variable terra.StringValue `hcl:"variable,attr" validate:"required"`
}

type NotPrincipals struct {
	// Identifiers: set of string, required
	Identifiers terra.SetValue[terra.StringValue] `hcl:"identifiers,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Principals struct {
	// Identifiers: set of string, required
	Identifiers terra.SetValue[terra.StringValue] `hcl:"identifiers,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type StatementAttributes struct {
	ref terra.Reference
}

func (s StatementAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StatementAttributes) InternalWithRef(ref terra.Reference) StatementAttributes {
	return StatementAttributes{ref: ref}
}

func (s StatementAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StatementAttributes) Actions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("actions"))
}

func (s StatementAttributes) Effect() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("effect"))
}

func (s StatementAttributes) NotActions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("not_actions"))
}

func (s StatementAttributes) NotResources() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("not_resources"))
}

func (s StatementAttributes) Resources() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("resources"))
}

func (s StatementAttributes) Sid() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("sid"))
}

func (s StatementAttributes) Condition() terra.SetValue[ConditionAttributes] {
	return terra.ReferenceAsSet[ConditionAttributes](s.ref.Append("condition"))
}

func (s StatementAttributes) NotPrincipals() terra.SetValue[NotPrincipalsAttributes] {
	return terra.ReferenceAsSet[NotPrincipalsAttributes](s.ref.Append("not_principals"))
}

func (s StatementAttributes) Principals() terra.SetValue[PrincipalsAttributes] {
	return terra.ReferenceAsSet[PrincipalsAttributes](s.ref.Append("principals"))
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

func (c ConditionAttributes) Test() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("test"))
}

func (c ConditionAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("values"))
}

func (c ConditionAttributes) Variable() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("variable"))
}

type NotPrincipalsAttributes struct {
	ref terra.Reference
}

func (np NotPrincipalsAttributes) InternalRef() (terra.Reference, error) {
	return np.ref, nil
}

func (np NotPrincipalsAttributes) InternalWithRef(ref terra.Reference) NotPrincipalsAttributes {
	return NotPrincipalsAttributes{ref: ref}
}

func (np NotPrincipalsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return np.ref.InternalTokens()
}

func (np NotPrincipalsAttributes) Identifiers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](np.ref.Append("identifiers"))
}

func (np NotPrincipalsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("type"))
}

type PrincipalsAttributes struct {
	ref terra.Reference
}

func (p PrincipalsAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PrincipalsAttributes) InternalWithRef(ref terra.Reference) PrincipalsAttributes {
	return PrincipalsAttributes{ref: ref}
}

func (p PrincipalsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PrincipalsAttributes) Identifiers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](p.ref.Append("identifiers"))
}

func (p PrincipalsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("type"))
}

type StatementState struct {
	Actions       []string             `json:"actions"`
	Effect        string               `json:"effect"`
	NotActions    []string             `json:"not_actions"`
	NotResources  []string             `json:"not_resources"`
	Resources     []string             `json:"resources"`
	Sid           string               `json:"sid"`
	Condition     []ConditionState     `json:"condition"`
	NotPrincipals []NotPrincipalsState `json:"not_principals"`
	Principals    []PrincipalsState    `json:"principals"`
}

type ConditionState struct {
	Test     string   `json:"test"`
	Values   []string `json:"values"`
	Variable string   `json:"variable"`
}

type NotPrincipalsState struct {
	Identifiers []string `json:"identifiers"`
	Type        string   `json:"type"`
}

type PrincipalsState struct {
	Identifiers []string `json:"identifiers"`
	Type        string   `json:"type"`
}
