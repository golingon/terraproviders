// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package wafbytematchset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ByteMatchTuples struct {
	// PositionalConstraint: string, required
	PositionalConstraint terra.StringValue `hcl:"positional_constraint,attr" validate:"required"`
	// TargetString: string, optional
	TargetString terra.StringValue `hcl:"target_string,attr"`
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

type ByteMatchTuplesAttributes struct {
	ref terra.Reference
}

func (bmt ByteMatchTuplesAttributes) InternalRef() terra.Reference {
	return bmt.ref
}

func (bmt ByteMatchTuplesAttributes) InternalWithRef(ref terra.Reference) ByteMatchTuplesAttributes {
	return ByteMatchTuplesAttributes{ref: ref}
}

func (bmt ByteMatchTuplesAttributes) InternalTokens() hclwrite.Tokens {
	return bmt.ref.InternalTokens()
}

func (bmt ByteMatchTuplesAttributes) PositionalConstraint() terra.StringValue {
	return terra.ReferenceAsString(bmt.ref.Append("positional_constraint"))
}

func (bmt ByteMatchTuplesAttributes) TargetString() terra.StringValue {
	return terra.ReferenceAsString(bmt.ref.Append("target_string"))
}

func (bmt ByteMatchTuplesAttributes) TextTransformation() terra.StringValue {
	return terra.ReferenceAsString(bmt.ref.Append("text_transformation"))
}

func (bmt ByteMatchTuplesAttributes) FieldToMatch() terra.ListValue[FieldToMatchAttributes] {
	return terra.ReferenceAsList[FieldToMatchAttributes](bmt.ref.Append("field_to_match"))
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

type ByteMatchTuplesState struct {
	PositionalConstraint string              `json:"positional_constraint"`
	TargetString         string              `json:"target_string"`
	TextTransformation   string              `json:"text_transformation"`
	FieldToMatch         []FieldToMatchState `json:"field_to_match"`
}

type FieldToMatchState struct {
	Data string `json:"data"`
	Type string `json:"type"`
}
