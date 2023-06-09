// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package wafregionalxssmatchset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type XssMatchTuple struct {
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

type XssMatchTupleAttributes struct {
	ref terra.Reference
}

func (xmt XssMatchTupleAttributes) InternalRef() (terra.Reference, error) {
	return xmt.ref, nil
}

func (xmt XssMatchTupleAttributes) InternalWithRef(ref terra.Reference) XssMatchTupleAttributes {
	return XssMatchTupleAttributes{ref: ref}
}

func (xmt XssMatchTupleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return xmt.ref.InternalTokens()
}

func (xmt XssMatchTupleAttributes) TextTransformation() terra.StringValue {
	return terra.ReferenceAsString(xmt.ref.Append("text_transformation"))
}

func (xmt XssMatchTupleAttributes) FieldToMatch() terra.ListValue[FieldToMatchAttributes] {
	return terra.ReferenceAsList[FieldToMatchAttributes](xmt.ref.Append("field_to_match"))
}

type FieldToMatchAttributes struct {
	ref terra.Reference
}

func (ftm FieldToMatchAttributes) InternalRef() (terra.Reference, error) {
	return ftm.ref, nil
}

func (ftm FieldToMatchAttributes) InternalWithRef(ref terra.Reference) FieldToMatchAttributes {
	return FieldToMatchAttributes{ref: ref}
}

func (ftm FieldToMatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ftm.ref.InternalTokens()
}

func (ftm FieldToMatchAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(ftm.ref.Append("data"))
}

func (ftm FieldToMatchAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ftm.ref.Append("type"))
}

type XssMatchTupleState struct {
	TextTransformation string              `json:"text_transformation"`
	FieldToMatch       []FieldToMatchState `json:"field_to_match"`
}

type FieldToMatchState struct {
	Data string `json:"data"`
	Type string `json:"type"`
}
