// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package wafxssmatchset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type XssMatchTuples struct {
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

type XssMatchTuplesAttributes struct {
	ref terra.Reference
}

func (xmt XssMatchTuplesAttributes) InternalRef() terra.Reference {
	return xmt.ref
}

func (xmt XssMatchTuplesAttributes) InternalWithRef(ref terra.Reference) XssMatchTuplesAttributes {
	return XssMatchTuplesAttributes{ref: ref}
}

func (xmt XssMatchTuplesAttributes) InternalTokens() hclwrite.Tokens {
	return xmt.ref.InternalTokens()
}

func (xmt XssMatchTuplesAttributes) TextTransformation() terra.StringValue {
	return terra.ReferenceAsString(xmt.ref.Append("text_transformation"))
}

func (xmt XssMatchTuplesAttributes) FieldToMatch() terra.ListValue[FieldToMatchAttributes] {
	return terra.ReferenceAsList[FieldToMatchAttributes](xmt.ref.Append("field_to_match"))
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

type XssMatchTuplesState struct {
	TextTransformation string              `json:"text_transformation"`
	FieldToMatch       []FieldToMatchState `json:"field_to_match"`
}

type FieldToMatchState struct {
	Data string `json:"data"`
	Type string `json:"type"`
}
