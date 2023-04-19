// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudfrontfieldlevelencryptionprofile

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EncryptionEntities struct {
	// Items: min=0
	Items []Items `hcl:"items,block" validate:"min=0"`
}

type Items struct {
	// ProviderId: string, required
	ProviderId terra.StringValue `hcl:"provider_id,attr" validate:"required"`
	// PublicKeyId: string, required
	PublicKeyId terra.StringValue `hcl:"public_key_id,attr" validate:"required"`
	// FieldPatterns: required
	FieldPatterns *FieldPatterns `hcl:"field_patterns,block" validate:"required"`
}

type FieldPatterns struct {
	// Items: set of string, optional
	Items terra.SetValue[terra.StringValue] `hcl:"items,attr"`
}

type EncryptionEntitiesAttributes struct {
	ref terra.Reference
}

func (ee EncryptionEntitiesAttributes) InternalRef() (terra.Reference, error) {
	return ee.ref, nil
}

func (ee EncryptionEntitiesAttributes) InternalWithRef(ref terra.Reference) EncryptionEntitiesAttributes {
	return EncryptionEntitiesAttributes{ref: ref}
}

func (ee EncryptionEntitiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ee.ref.InternalTokens()
}

func (ee EncryptionEntitiesAttributes) Items() terra.SetValue[ItemsAttributes] {
	return terra.ReferenceAsSet[ItemsAttributes](ee.ref.Append("items"))
}

type ItemsAttributes struct {
	ref terra.Reference
}

func (i ItemsAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i ItemsAttributes) InternalWithRef(ref terra.Reference) ItemsAttributes {
	return ItemsAttributes{ref: ref}
}

func (i ItemsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i ItemsAttributes) ProviderId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("provider_id"))
}

func (i ItemsAttributes) PublicKeyId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("public_key_id"))
}

func (i ItemsAttributes) FieldPatterns() terra.ListValue[FieldPatternsAttributes] {
	return terra.ReferenceAsList[FieldPatternsAttributes](i.ref.Append("field_patterns"))
}

type FieldPatternsAttributes struct {
	ref terra.Reference
}

func (fp FieldPatternsAttributes) InternalRef() (terra.Reference, error) {
	return fp.ref, nil
}

func (fp FieldPatternsAttributes) InternalWithRef(ref terra.Reference) FieldPatternsAttributes {
	return FieldPatternsAttributes{ref: ref}
}

func (fp FieldPatternsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fp.ref.InternalTokens()
}

func (fp FieldPatternsAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fp.ref.Append("items"))
}

type EncryptionEntitiesState struct {
	Items []ItemsState `json:"items"`
}

type ItemsState struct {
	ProviderId    string               `json:"provider_id"`
	PublicKeyId   string               `json:"public_key_id"`
	FieldPatterns []FieldPatternsState `json:"field_patterns"`
}

type FieldPatternsState struct {
	Items []string `json:"items"`
}
