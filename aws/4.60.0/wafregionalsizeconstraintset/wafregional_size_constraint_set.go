// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package wafregionalsizeconstraintset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SizeConstraints struct {
	// ComparisonOperator: string, required
	ComparisonOperator terra.StringValue `hcl:"comparison_operator,attr" validate:"required"`
	// Size: number, required
	Size terra.NumberValue `hcl:"size,attr" validate:"required"`
	// TextTransformation: string, required
	TextTransformation terra.StringValue `hcl:"text_transformation,attr" validate:"required"`
	// FieldToMatch: required
	FieldToMatch *FieldToMatch `hcl:"field_to_match,block" validate:"required"`
}

type FieldToMatch struct {
	// Data: string, optional
	Data terra.StringValue `hcl:"data,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type SizeConstraintsAttributes struct {
	ref terra.Reference
}

func (sc SizeConstraintsAttributes) InternalRef() terra.Reference {
	return sc.ref
}

func (sc SizeConstraintsAttributes) InternalWithRef(ref terra.Reference) SizeConstraintsAttributes {
	return SizeConstraintsAttributes{ref: ref}
}

func (sc SizeConstraintsAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc SizeConstraintsAttributes) ComparisonOperator() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("comparison_operator"))
}

func (sc SizeConstraintsAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("size"))
}

func (sc SizeConstraintsAttributes) TextTransformation() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("text_transformation"))
}

func (sc SizeConstraintsAttributes) FieldToMatch() terra.ListValue[FieldToMatchAttributes] {
	return terra.ReferenceAsList[FieldToMatchAttributes](sc.ref.Append("field_to_match"))
}

type FieldToMatchAttributes struct {
	ref terra.Reference
}

func (ftm FieldToMatchAttributes) InternalRef() terra.Reference {
	return ftm.ref
}

func (ftm FieldToMatchAttributes) InternalWithRef(ref terra.Reference) FieldToMatchAttributes {
	return FieldToMatchAttributes{ref: ref}
}

func (ftm FieldToMatchAttributes) InternalTokens() hclwrite.Tokens {
	return ftm.ref.InternalTokens()
}

func (ftm FieldToMatchAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(ftm.ref.Append("data"))
}

func (ftm FieldToMatchAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ftm.ref.Append("type"))
}

type SizeConstraintsState struct {
	ComparisonOperator string              `json:"comparison_operator"`
	Size               float64             `json:"size"`
	TextTransformation string              `json:"text_transformation"`
	FieldToMatch       []FieldToMatchState `json:"field_to_match"`
}

type FieldToMatchState struct {
	Data string `json:"data"`
	Type string `json:"type"`
}
