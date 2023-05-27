// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataidentitystoregroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ExternalIds struct{}

type AlternateIdentifier struct {
	// ExternalId: optional
	ExternalId *ExternalId `hcl:"external_id,block"`
	// UniqueAttribute: optional
	UniqueAttribute *UniqueAttribute `hcl:"unique_attribute,block"`
}

type ExternalId struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Issuer: string, required
	Issuer terra.StringValue `hcl:"issuer,attr" validate:"required"`
}

type UniqueAttribute struct {
	// AttributePath: string, required
	AttributePath terra.StringValue `hcl:"attribute_path,attr" validate:"required"`
	// AttributeValue: string, required
	AttributeValue terra.StringValue `hcl:"attribute_value,attr" validate:"required"`
}

type ExternalIdsAttributes struct {
	ref terra.Reference
}

func (ei ExternalIdsAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei ExternalIdsAttributes) InternalWithRef(ref terra.Reference) ExternalIdsAttributes {
	return ExternalIdsAttributes{ref: ref}
}

func (ei ExternalIdsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ei.ref.InternalTokens()
}

func (ei ExternalIdsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("id"))
}

func (ei ExternalIdsAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("issuer"))
}

type AlternateIdentifierAttributes struct {
	ref terra.Reference
}

func (ai AlternateIdentifierAttributes) InternalRef() (terra.Reference, error) {
	return ai.ref, nil
}

func (ai AlternateIdentifierAttributes) InternalWithRef(ref terra.Reference) AlternateIdentifierAttributes {
	return AlternateIdentifierAttributes{ref: ref}
}

func (ai AlternateIdentifierAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ai.ref.InternalTokens()
}

func (ai AlternateIdentifierAttributes) ExternalId() terra.ListValue[ExternalIdAttributes] {
	return terra.ReferenceAsList[ExternalIdAttributes](ai.ref.Append("external_id"))
}

func (ai AlternateIdentifierAttributes) UniqueAttribute() terra.ListValue[UniqueAttributeAttributes] {
	return terra.ReferenceAsList[UniqueAttributeAttributes](ai.ref.Append("unique_attribute"))
}

type ExternalIdAttributes struct {
	ref terra.Reference
}

func (ei ExternalIdAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei ExternalIdAttributes) InternalWithRef(ref terra.Reference) ExternalIdAttributes {
	return ExternalIdAttributes{ref: ref}
}

func (ei ExternalIdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ei.ref.InternalTokens()
}

func (ei ExternalIdAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("id"))
}

func (ei ExternalIdAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("issuer"))
}

type UniqueAttributeAttributes struct {
	ref terra.Reference
}

func (ua UniqueAttributeAttributes) InternalRef() (terra.Reference, error) {
	return ua.ref, nil
}

func (ua UniqueAttributeAttributes) InternalWithRef(ref terra.Reference) UniqueAttributeAttributes {
	return UniqueAttributeAttributes{ref: ref}
}

func (ua UniqueAttributeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ua.ref.InternalTokens()
}

func (ua UniqueAttributeAttributes) AttributePath() terra.StringValue {
	return terra.ReferenceAsString(ua.ref.Append("attribute_path"))
}

func (ua UniqueAttributeAttributes) AttributeValue() terra.StringValue {
	return terra.ReferenceAsString(ua.ref.Append("attribute_value"))
}

type ExternalIdsState struct {
	Id     string `json:"id"`
	Issuer string `json:"issuer"`
}

type AlternateIdentifierState struct {
	ExternalId      []ExternalIdState      `json:"external_id"`
	UniqueAttribute []UniqueAttributeState `json:"unique_attribute"`
}

type ExternalIdState struct {
	Id     string `json:"id"`
	Issuer string `json:"issuer"`
}

type UniqueAttributeState struct {
	AttributePath  string `json:"attribute_path"`
	AttributeValue string `json:"attribute_value"`
}
