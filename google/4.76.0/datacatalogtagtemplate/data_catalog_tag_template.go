// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacatalogtagtemplate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Fields struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// FieldId: string, required
	FieldId terra.StringValue `hcl:"field_id,attr" validate:"required"`
	// IsRequired: bool, optional
	IsRequired terra.BoolValue `hcl:"is_required,attr"`
	// Order: number, optional
	Order terra.NumberValue `hcl:"order,attr"`
	// Type: required
	Type *Type `hcl:"type,block" validate:"required"`
}

type Type struct {
	// PrimitiveType: string, optional
	PrimitiveType terra.StringValue `hcl:"primitive_type,attr"`
	// EnumType: optional
	EnumType *EnumType `hcl:"enum_type,block"`
}

type EnumType struct {
	// AllowedValues: min=1
	AllowedValues []AllowedValues `hcl:"allowed_values,block" validate:"min=1"`
}

type AllowedValues struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type FieldsAttributes struct {
	ref terra.Reference
}

func (f FieldsAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FieldsAttributes) InternalWithRef(ref terra.Reference) FieldsAttributes {
	return FieldsAttributes{ref: ref}
}

func (f FieldsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FieldsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("description"))
}

func (f FieldsAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("display_name"))
}

func (f FieldsAttributes) FieldId() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("field_id"))
}

func (f FieldsAttributes) IsRequired() terra.BoolValue {
	return terra.ReferenceAsBool(f.ref.Append("is_required"))
}

func (f FieldsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f FieldsAttributes) Order() terra.NumberValue {
	return terra.ReferenceAsNumber(f.ref.Append("order"))
}

func (f FieldsAttributes) Type() terra.ListValue[TypeAttributes] {
	return terra.ReferenceAsList[TypeAttributes](f.ref.Append("type"))
}

type TypeAttributes struct {
	ref terra.Reference
}

func (t TypeAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TypeAttributes) InternalWithRef(ref terra.Reference) TypeAttributes {
	return TypeAttributes{ref: ref}
}

func (t TypeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TypeAttributes) PrimitiveType() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("primitive_type"))
}

func (t TypeAttributes) EnumType() terra.ListValue[EnumTypeAttributes] {
	return terra.ReferenceAsList[EnumTypeAttributes](t.ref.Append("enum_type"))
}

type EnumTypeAttributes struct {
	ref terra.Reference
}

func (et EnumTypeAttributes) InternalRef() (terra.Reference, error) {
	return et.ref, nil
}

func (et EnumTypeAttributes) InternalWithRef(ref terra.Reference) EnumTypeAttributes {
	return EnumTypeAttributes{ref: ref}
}

func (et EnumTypeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return et.ref.InternalTokens()
}

func (et EnumTypeAttributes) AllowedValues() terra.SetValue[AllowedValuesAttributes] {
	return terra.ReferenceAsSet[AllowedValuesAttributes](et.ref.Append("allowed_values"))
}

type AllowedValuesAttributes struct {
	ref terra.Reference
}

func (av AllowedValuesAttributes) InternalRef() (terra.Reference, error) {
	return av.ref, nil
}

func (av AllowedValuesAttributes) InternalWithRef(ref terra.Reference) AllowedValuesAttributes {
	return AllowedValuesAttributes{ref: ref}
}

func (av AllowedValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return av.ref.InternalTokens()
}

func (av AllowedValuesAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(av.ref.Append("display_name"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type FieldsState struct {
	Description string      `json:"description"`
	DisplayName string      `json:"display_name"`
	FieldId     string      `json:"field_id"`
	IsRequired  bool        `json:"is_required"`
	Name        string      `json:"name"`
	Order       float64     `json:"order"`
	Type        []TypeState `json:"type"`
}

type TypeState struct {
	PrimitiveType string          `json:"primitive_type"`
	EnumType      []EnumTypeState `json:"enum_type"`
}

type EnumTypeState struct {
	AllowedValues []AllowedValuesState `json:"allowed_values"`
}

type AllowedValuesState struct {
	DisplayName string `json:"display_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
