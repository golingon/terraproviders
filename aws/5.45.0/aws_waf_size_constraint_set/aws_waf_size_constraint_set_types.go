// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_waf_size_constraint_set

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type SizeConstraints struct {
	// ComparisonOperator: string, required
	ComparisonOperator terra.StringValue `hcl:"comparison_operator,attr" validate:"required"`
	// Size: number, required
	Size terra.NumberValue `hcl:"size,attr" validate:"required"`
	// TextTransformation: string, required
	TextTransformation terra.StringValue `hcl:"text_transformation,attr" validate:"required"`
	// SizeConstraintsFieldToMatch: required
	FieldToMatch *SizeConstraintsFieldToMatch `hcl:"field_to_match,block" validate:"required"`
}

type SizeConstraintsFieldToMatch struct {
	// Data: string, optional
	Data terra.StringValue `hcl:"data,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type SizeConstraintsAttributes struct {
	ref terra.Reference
}

func (sc SizeConstraintsAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SizeConstraintsAttributes) InternalWithRef(ref terra.Reference) SizeConstraintsAttributes {
	return SizeConstraintsAttributes{ref: ref}
}

func (sc SizeConstraintsAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

func (sc SizeConstraintsAttributes) FieldToMatch() terra.ListValue[SizeConstraintsFieldToMatchAttributes] {
	return terra.ReferenceAsList[SizeConstraintsFieldToMatchAttributes](sc.ref.Append("field_to_match"))
}

type SizeConstraintsFieldToMatchAttributes struct {
	ref terra.Reference
}

func (ftm SizeConstraintsFieldToMatchAttributes) InternalRef() (terra.Reference, error) {
	return ftm.ref, nil
}

func (ftm SizeConstraintsFieldToMatchAttributes) InternalWithRef(ref terra.Reference) SizeConstraintsFieldToMatchAttributes {
	return SizeConstraintsFieldToMatchAttributes{ref: ref}
}

func (ftm SizeConstraintsFieldToMatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ftm.ref.InternalTokens()
}

func (ftm SizeConstraintsFieldToMatchAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(ftm.ref.Append("data"))
}

func (ftm SizeConstraintsFieldToMatchAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ftm.ref.Append("type"))
}

type SizeConstraintsState struct {
	ComparisonOperator string                             `json:"comparison_operator"`
	Size               float64                            `json:"size"`
	TextTransformation string                             `json:"text_transformation"`
	FieldToMatch       []SizeConstraintsFieldToMatchState `json:"field_to_match"`
}

type SizeConstraintsFieldToMatchState struct {
	Data string `json:"data"`
	Type string `json:"type"`
}
