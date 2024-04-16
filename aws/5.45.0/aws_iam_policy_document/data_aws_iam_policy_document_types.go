// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_iam_policy_document

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataStatement struct {
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
	// StatementCondition: min=0
	Condition []DataStatementCondition `hcl:"condition,block" validate:"min=0"`
	// StatementNotPrincipals: min=0
	NotPrincipals []DataStatementNotPrincipals `hcl:"not_principals,block" validate:"min=0"`
	// StatementPrincipals: min=0
	Principals []DataStatementPrincipals `hcl:"principals,block" validate:"min=0"`
}

type DataStatementCondition struct {
	// Test: string, required
	Test terra.StringValue `hcl:"test,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
	// Variable: string, required
	Variable terra.StringValue `hcl:"variable,attr" validate:"required"`
}

type DataStatementNotPrincipals struct {
	// Identifiers: set of string, required
	Identifiers terra.SetValue[terra.StringValue] `hcl:"identifiers,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type DataStatementPrincipals struct {
	// Identifiers: set of string, required
	Identifiers terra.SetValue[terra.StringValue] `hcl:"identifiers,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type DataStatementAttributes struct {
	ref terra.Reference
}

func (s DataStatementAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s DataStatementAttributes) InternalWithRef(ref terra.Reference) DataStatementAttributes {
	return DataStatementAttributes{ref: ref}
}

func (s DataStatementAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s DataStatementAttributes) Actions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("actions"))
}

func (s DataStatementAttributes) Effect() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("effect"))
}

func (s DataStatementAttributes) NotActions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("not_actions"))
}

func (s DataStatementAttributes) NotResources() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("not_resources"))
}

func (s DataStatementAttributes) Resources() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("resources"))
}

func (s DataStatementAttributes) Sid() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("sid"))
}

func (s DataStatementAttributes) Condition() terra.SetValue[DataStatementConditionAttributes] {
	return terra.ReferenceAsSet[DataStatementConditionAttributes](s.ref.Append("condition"))
}

func (s DataStatementAttributes) NotPrincipals() terra.SetValue[DataStatementNotPrincipalsAttributes] {
	return terra.ReferenceAsSet[DataStatementNotPrincipalsAttributes](s.ref.Append("not_principals"))
}

func (s DataStatementAttributes) Principals() terra.SetValue[DataStatementPrincipalsAttributes] {
	return terra.ReferenceAsSet[DataStatementPrincipalsAttributes](s.ref.Append("principals"))
}

type DataStatementConditionAttributes struct {
	ref terra.Reference
}

func (c DataStatementConditionAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c DataStatementConditionAttributes) InternalWithRef(ref terra.Reference) DataStatementConditionAttributes {
	return DataStatementConditionAttributes{ref: ref}
}

func (c DataStatementConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c DataStatementConditionAttributes) Test() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("test"))
}

func (c DataStatementConditionAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("values"))
}

func (c DataStatementConditionAttributes) Variable() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("variable"))
}

type DataStatementNotPrincipalsAttributes struct {
	ref terra.Reference
}

func (np DataStatementNotPrincipalsAttributes) InternalRef() (terra.Reference, error) {
	return np.ref, nil
}

func (np DataStatementNotPrincipalsAttributes) InternalWithRef(ref terra.Reference) DataStatementNotPrincipalsAttributes {
	return DataStatementNotPrincipalsAttributes{ref: ref}
}

func (np DataStatementNotPrincipalsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return np.ref.InternalTokens()
}

func (np DataStatementNotPrincipalsAttributes) Identifiers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](np.ref.Append("identifiers"))
}

func (np DataStatementNotPrincipalsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("type"))
}

type DataStatementPrincipalsAttributes struct {
	ref terra.Reference
}

func (p DataStatementPrincipalsAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p DataStatementPrincipalsAttributes) InternalWithRef(ref terra.Reference) DataStatementPrincipalsAttributes {
	return DataStatementPrincipalsAttributes{ref: ref}
}

func (p DataStatementPrincipalsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p DataStatementPrincipalsAttributes) Identifiers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](p.ref.Append("identifiers"))
}

func (p DataStatementPrincipalsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("type"))
}

type DataStatementState struct {
	Actions       []string                          `json:"actions"`
	Effect        string                            `json:"effect"`
	NotActions    []string                          `json:"not_actions"`
	NotResources  []string                          `json:"not_resources"`
	Resources     []string                          `json:"resources"`
	Sid           string                            `json:"sid"`
	Condition     []DataStatementConditionState     `json:"condition"`
	NotPrincipals []DataStatementNotPrincipalsState `json:"not_principals"`
	Principals    []DataStatementPrincipalsState    `json:"principals"`
}

type DataStatementConditionState struct {
	Test     string   `json:"test"`
	Values   []string `json:"values"`
	Variable string   `json:"variable"`
}

type DataStatementNotPrincipalsState struct {
	Identifiers []string `json:"identifiers"`
	Type        string   `json:"type"`
}

type DataStatementPrincipalsState struct {
	Identifiers []string `json:"identifiers"`
	Type        string   `json:"type"`
}
