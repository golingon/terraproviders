// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googleworkspace_schema

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataFieldsAttributes struct {
	ref terra.Reference
}

func (f DataFieldsAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f DataFieldsAttributes) InternalWithRef(ref terra.Reference) DataFieldsAttributes {
	return DataFieldsAttributes{ref: ref}
}

func (f DataFieldsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f DataFieldsAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("display_name"))
}

func (f DataFieldsAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("etag"))
}

func (f DataFieldsAttributes) FieldId() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("field_id"))
}

func (f DataFieldsAttributes) FieldName() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("field_name"))
}

func (f DataFieldsAttributes) FieldType() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("field_type"))
}

func (f DataFieldsAttributes) Indexed() terra.BoolValue {
	return terra.ReferenceAsBool(f.ref.Append("indexed"))
}

func (f DataFieldsAttributes) MultiValued() terra.BoolValue {
	return terra.ReferenceAsBool(f.ref.Append("multi_valued"))
}

func (f DataFieldsAttributes) ReadAccessType() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("read_access_type"))
}

func (f DataFieldsAttributes) NumericIndexingSpec() terra.ListValue[DataFieldsNumericIndexingSpecAttributes] {
	return terra.ReferenceAsList[DataFieldsNumericIndexingSpecAttributes](f.ref.Append("numeric_indexing_spec"))
}

type DataFieldsNumericIndexingSpecAttributes struct {
	ref terra.Reference
}

func (nis DataFieldsNumericIndexingSpecAttributes) InternalRef() (terra.Reference, error) {
	return nis.ref, nil
}

func (nis DataFieldsNumericIndexingSpecAttributes) InternalWithRef(ref terra.Reference) DataFieldsNumericIndexingSpecAttributes {
	return DataFieldsNumericIndexingSpecAttributes{ref: ref}
}

func (nis DataFieldsNumericIndexingSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nis.ref.InternalTokens()
}

func (nis DataFieldsNumericIndexingSpecAttributes) MaxValue() terra.NumberValue {
	return terra.ReferenceAsNumber(nis.ref.Append("max_value"))
}

func (nis DataFieldsNumericIndexingSpecAttributes) MinValue() terra.NumberValue {
	return terra.ReferenceAsNumber(nis.ref.Append("min_value"))
}

type DataFieldsState struct {
	DisplayName         string                               `json:"display_name"`
	Etag                string                               `json:"etag"`
	FieldId             string                               `json:"field_id"`
	FieldName           string                               `json:"field_name"`
	FieldType           string                               `json:"field_type"`
	Indexed             bool                                 `json:"indexed"`
	MultiValued         bool                                 `json:"multi_valued"`
	ReadAccessType      string                               `json:"read_access_type"`
	NumericIndexingSpec []DataFieldsNumericIndexingSpecState `json:"numeric_indexing_spec"`
}

type DataFieldsNumericIndexingSpecState struct {
	MaxValue float64 `json:"max_value"`
	MinValue float64 `json:"min_value"`
}
