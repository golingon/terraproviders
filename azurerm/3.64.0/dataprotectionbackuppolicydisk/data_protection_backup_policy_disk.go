// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataprotectionbackuppolicydisk

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RetentionRule struct {
	// Duration: string, required
	Duration terra.StringValue `hcl:"duration,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// Criteria: required
	Criteria *Criteria `hcl:"criteria,block" validate:"required"`
}

type Criteria struct {
	// AbsoluteCriteria: string, optional
	AbsoluteCriteria terra.StringValue `hcl:"absolute_criteria,attr"`
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

type RetentionRuleAttributes struct {
	ref terra.Reference
}

func (rr RetentionRuleAttributes) InternalRef() (terra.Reference, error) {
	return rr.ref, nil
}

func (rr RetentionRuleAttributes) InternalWithRef(ref terra.Reference) RetentionRuleAttributes {
	return RetentionRuleAttributes{ref: ref}
}

func (rr RetentionRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rr.ref.InternalTokens()
}

func (rr RetentionRuleAttributes) Duration() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("duration"))
}

func (rr RetentionRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("name"))
}

func (rr RetentionRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(rr.ref.Append("priority"))
}

func (rr RetentionRuleAttributes) Criteria() terra.ListValue[CriteriaAttributes] {
	return terra.ReferenceAsList[CriteriaAttributes](rr.ref.Append("criteria"))
}

type CriteriaAttributes struct {
	ref terra.Reference
}

func (c CriteriaAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CriteriaAttributes) InternalWithRef(ref terra.Reference) CriteriaAttributes {
	return CriteriaAttributes{ref: ref}
}

func (c CriteriaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CriteriaAttributes) AbsoluteCriteria() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("absolute_criteria"))
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

type RetentionRuleState struct {
	Duration string          `json:"duration"`
	Name     string          `json:"name"`
	Priority float64         `json:"priority"`
	Criteria []CriteriaState `json:"criteria"`
}

type CriteriaState struct {
	AbsoluteCriteria string `json:"absolute_criteria"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
