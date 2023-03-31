// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package macie2findingsfilter

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type FindingCriteria struct {
	// Criterion: min=0
	Criterion []Criterion `hcl:"criterion,block" validate:"min=0"`
}

type Criterion struct {
	// Eq: set of string, optional
	Eq terra.SetValue[terra.StringValue] `hcl:"eq,attr"`
	// EqExactMatch: set of string, optional
	EqExactMatch terra.SetValue[terra.StringValue] `hcl:"eq_exact_match,attr"`
	// Field: string, required
	Field terra.StringValue `hcl:"field,attr" validate:"required"`
	// Gt: string, optional
	Gt terra.StringValue `hcl:"gt,attr"`
	// Gte: string, optional
	Gte terra.StringValue `hcl:"gte,attr"`
	// Lt: string, optional
	Lt terra.StringValue `hcl:"lt,attr"`
	// Lte: string, optional
	Lte terra.StringValue `hcl:"lte,attr"`
	// Neq: set of string, optional
	Neq terra.SetValue[terra.StringValue] `hcl:"neq,attr"`
}

type FindingCriteriaAttributes struct {
	ref terra.Reference
}

func (fc FindingCriteriaAttributes) InternalRef() terra.Reference {
	return fc.ref
}

func (fc FindingCriteriaAttributes) InternalWithRef(ref terra.Reference) FindingCriteriaAttributes {
	return FindingCriteriaAttributes{ref: ref}
}

func (fc FindingCriteriaAttributes) InternalTokens() hclwrite.Tokens {
	return fc.ref.InternalTokens()
}

func (fc FindingCriteriaAttributes) Criterion() terra.SetValue[CriterionAttributes] {
	return terra.ReferenceAsSet[CriterionAttributes](fc.ref.Append("criterion"))
}

type CriterionAttributes struct {
	ref terra.Reference
}

func (c CriterionAttributes) InternalRef() terra.Reference {
	return c.ref
}

func (c CriterionAttributes) InternalWithRef(ref terra.Reference) CriterionAttributes {
	return CriterionAttributes{ref: ref}
}

func (c CriterionAttributes) InternalTokens() hclwrite.Tokens {
	return c.ref.InternalTokens()
}

func (c CriterionAttributes) Eq() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("eq"))
}

func (c CriterionAttributes) EqExactMatch() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("eq_exact_match"))
}

func (c CriterionAttributes) Field() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("field"))
}

func (c CriterionAttributes) Gt() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("gt"))
}

func (c CriterionAttributes) Gte() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("gte"))
}

func (c CriterionAttributes) Lt() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("lt"))
}

func (c CriterionAttributes) Lte() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("lte"))
}

func (c CriterionAttributes) Neq() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("neq"))
}

type FindingCriteriaState struct {
	Criterion []CriterionState `json:"criterion"`
}

type CriterionState struct {
	Eq           []string `json:"eq"`
	EqExactMatch []string `json:"eq_exact_match"`
	Field        string   `json:"field"`
	Gt           string   `json:"gt"`
	Gte          string   `json:"gte"`
	Lt           string   `json:"lt"`
	Lte          string   `json:"lte"`
	Neq          []string `json:"neq"`
}
