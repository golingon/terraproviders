// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package gluemltransform

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Schema struct{}

type InputRecordTables struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// ConnectionName: string, optional
	ConnectionName terra.StringValue `hcl:"connection_name,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// TableName: string, required
	TableName terra.StringValue `hcl:"table_name,attr" validate:"required"`
}

type Parameters struct {
	// TransformType: string, required
	TransformType terra.StringValue `hcl:"transform_type,attr" validate:"required"`
	// FindMatchesParameters: required
	FindMatchesParameters *FindMatchesParameters `hcl:"find_matches_parameters,block" validate:"required"`
}

type FindMatchesParameters struct {
	// AccuracyCostTradeOff: number, optional
	AccuracyCostTradeOff terra.NumberValue `hcl:"accuracy_cost_trade_off,attr"`
	// EnforceProvidedLabels: bool, optional
	EnforceProvidedLabels terra.BoolValue `hcl:"enforce_provided_labels,attr"`
	// PrecisionRecallTradeOff: number, optional
	PrecisionRecallTradeOff terra.NumberValue `hcl:"precision_recall_trade_off,attr"`
	// PrimaryKeyColumnName: string, optional
	PrimaryKeyColumnName terra.StringValue `hcl:"primary_key_column_name,attr"`
}

type SchemaAttributes struct {
	ref terra.Reference
}

func (s SchemaAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SchemaAttributes) InternalWithRef(ref terra.Reference) SchemaAttributes {
	return SchemaAttributes{ref: ref}
}

func (s SchemaAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SchemaAttributes) DataType() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("data_type"))
}

func (s SchemaAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

type InputRecordTablesAttributes struct {
	ref terra.Reference
}

func (irt InputRecordTablesAttributes) InternalRef() (terra.Reference, error) {
	return irt.ref, nil
}

func (irt InputRecordTablesAttributes) InternalWithRef(ref terra.Reference) InputRecordTablesAttributes {
	return InputRecordTablesAttributes{ref: ref}
}

func (irt InputRecordTablesAttributes) InternalTokens() hclwrite.Tokens {
	return irt.ref.InternalTokens()
}

func (irt InputRecordTablesAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(irt.ref.Append("catalog_id"))
}

func (irt InputRecordTablesAttributes) ConnectionName() terra.StringValue {
	return terra.ReferenceAsString(irt.ref.Append("connection_name"))
}

func (irt InputRecordTablesAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(irt.ref.Append("database_name"))
}

func (irt InputRecordTablesAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(irt.ref.Append("table_name"))
}

type ParametersAttributes struct {
	ref terra.Reference
}

func (p ParametersAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParametersAttributes) InternalWithRef(ref terra.Reference) ParametersAttributes {
	return ParametersAttributes{ref: ref}
}

func (p ParametersAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p ParametersAttributes) TransformType() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("transform_type"))
}

func (p ParametersAttributes) FindMatchesParameters() terra.ListValue[FindMatchesParametersAttributes] {
	return terra.ReferenceAsList[FindMatchesParametersAttributes](p.ref.Append("find_matches_parameters"))
}

type FindMatchesParametersAttributes struct {
	ref terra.Reference
}

func (fmp FindMatchesParametersAttributes) InternalRef() (terra.Reference, error) {
	return fmp.ref, nil
}

func (fmp FindMatchesParametersAttributes) InternalWithRef(ref terra.Reference) FindMatchesParametersAttributes {
	return FindMatchesParametersAttributes{ref: ref}
}

func (fmp FindMatchesParametersAttributes) InternalTokens() hclwrite.Tokens {
	return fmp.ref.InternalTokens()
}

func (fmp FindMatchesParametersAttributes) AccuracyCostTradeOff() terra.NumberValue {
	return terra.ReferenceAsNumber(fmp.ref.Append("accuracy_cost_trade_off"))
}

func (fmp FindMatchesParametersAttributes) EnforceProvidedLabels() terra.BoolValue {
	return terra.ReferenceAsBool(fmp.ref.Append("enforce_provided_labels"))
}

func (fmp FindMatchesParametersAttributes) PrecisionRecallTradeOff() terra.NumberValue {
	return terra.ReferenceAsNumber(fmp.ref.Append("precision_recall_trade_off"))
}

func (fmp FindMatchesParametersAttributes) PrimaryKeyColumnName() terra.StringValue {
	return terra.ReferenceAsString(fmp.ref.Append("primary_key_column_name"))
}

type SchemaState struct {
	DataType string `json:"data_type"`
	Name     string `json:"name"`
}

type InputRecordTablesState struct {
	CatalogId      string `json:"catalog_id"`
	ConnectionName string `json:"connection_name"`
	DatabaseName   string `json:"database_name"`
	TableName      string `json:"table_name"`
}

type ParametersState struct {
	TransformType         string                       `json:"transform_type"`
	FindMatchesParameters []FindMatchesParametersState `json:"find_matches_parameters"`
}

type FindMatchesParametersState struct {
	AccuracyCostTradeOff    float64 `json:"accuracy_cost_trade_off"`
	EnforceProvidedLabels   bool    `json:"enforce_provided_labels"`
	PrecisionRecallTradeOff float64 `json:"precision_recall_trade_off"`
	PrimaryKeyColumnName    string  `json:"primary_key_column_name"`
}
