// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package evidentlyfeature

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EvaluationRules struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type Variations struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: required
	Value *Value `hcl:"value,block" validate:"required"`
}

type Value struct {
	// BoolValue: string, optional
	BoolValue terra.StringValue `hcl:"bool_value,attr"`
	// DoubleValue: string, optional
	DoubleValue terra.StringValue `hcl:"double_value,attr"`
	// LongValue: string, optional
	LongValue terra.StringValue `hcl:"long_value,attr"`
	// StringValue: string, optional
	StringValue terra.StringValue `hcl:"string_value,attr"`
}

type EvaluationRulesAttributes struct {
	ref terra.Reference
}

func (er EvaluationRulesAttributes) InternalRef() (terra.Reference, error) {
	return er.ref, nil
}

func (er EvaluationRulesAttributes) InternalWithRef(ref terra.Reference) EvaluationRulesAttributes {
	return EvaluationRulesAttributes{ref: ref}
}

func (er EvaluationRulesAttributes) InternalTokens() hclwrite.Tokens {
	return er.ref.InternalTokens()
}

func (er EvaluationRulesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("name"))
}

func (er EvaluationRulesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("type"))
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

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
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

type VariationsAttributes struct {
	ref terra.Reference
}

func (v VariationsAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v VariationsAttributes) InternalWithRef(ref terra.Reference) VariationsAttributes {
	return VariationsAttributes{ref: ref}
}

func (v VariationsAttributes) InternalTokens() hclwrite.Tokens {
	return v.ref.InternalTokens()
}

func (v VariationsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("name"))
}

func (v VariationsAttributes) Value() terra.ListValue[ValueAttributes] {
	return terra.ReferenceAsList[ValueAttributes](v.ref.Append("value"))
}

type ValueAttributes struct {
	ref terra.Reference
}

func (v ValueAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v ValueAttributes) InternalWithRef(ref terra.Reference) ValueAttributes {
	return ValueAttributes{ref: ref}
}

func (v ValueAttributes) InternalTokens() hclwrite.Tokens {
	return v.ref.InternalTokens()
}

func (v ValueAttributes) BoolValue() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("bool_value"))
}

func (v ValueAttributes) DoubleValue() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("double_value"))
}

func (v ValueAttributes) LongValue() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("long_value"))
}

func (v ValueAttributes) StringValue() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("string_value"))
}

type EvaluationRulesState struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type VariationsState struct {
	Name  string       `json:"name"`
	Value []ValueState `json:"value"`
}

type ValueState struct {
	BoolValue   string `json:"bool_value"`
	DoubleValue string `json:"double_value"`
	LongValue   string `json:"long_value"`
	StringValue string `json:"string_value"`
}
