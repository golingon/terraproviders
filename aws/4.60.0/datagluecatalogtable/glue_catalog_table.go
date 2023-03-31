// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datagluecatalogtable

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PartitionIndex struct{}

type PartitionKeys struct{}

type StorageDescriptor struct {
	// Columns: min=0
	Columns []Columns `hcl:"columns,block" validate:"min=0"`
	// SchemaReference: min=0
	SchemaReference []SchemaReference `hcl:"schema_reference,block" validate:"min=0"`
	// SerDeInfo: min=0
	SerDeInfo []SerDeInfo `hcl:"ser_de_info,block" validate:"min=0"`
	// SkewedInfo: min=0
	SkewedInfo []SkewedInfo `hcl:"skewed_info,block" validate:"min=0"`
	// SortColumns: min=0
	SortColumns []SortColumns `hcl:"sort_columns,block" validate:"min=0"`
}

type Columns struct{}

type SchemaReference struct {
	// SchemaId: min=0
	SchemaId []SchemaId `hcl:"schema_id,block" validate:"min=0"`
}

type SchemaId struct{}

type SerDeInfo struct{}

type SkewedInfo struct{}

type SortColumns struct{}

type TargetTable struct{}

type PartitionIndexAttributes struct {
	ref terra.Reference
}

func (pi PartitionIndexAttributes) InternalRef() terra.Reference {
	return pi.ref
}

func (pi PartitionIndexAttributes) InternalWithRef(ref terra.Reference) PartitionIndexAttributes {
	return PartitionIndexAttributes{ref: ref}
}

func (pi PartitionIndexAttributes) InternalTokens() hclwrite.Tokens {
	return pi.ref.InternalTokens()
}

func (pi PartitionIndexAttributes) IndexName() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("index_name"))
}

func (pi PartitionIndexAttributes) IndexStatus() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("index_status"))
}

func (pi PartitionIndexAttributes) Keys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pi.ref.Append("keys"))
}

type PartitionKeysAttributes struct {
	ref terra.Reference
}

func (pk PartitionKeysAttributes) InternalRef() terra.Reference {
	return pk.ref
}

func (pk PartitionKeysAttributes) InternalWithRef(ref terra.Reference) PartitionKeysAttributes {
	return PartitionKeysAttributes{ref: ref}
}

func (pk PartitionKeysAttributes) InternalTokens() hclwrite.Tokens {
	return pk.ref.InternalTokens()
}

func (pk PartitionKeysAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("comment"))
}

func (pk PartitionKeysAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("name"))
}

func (pk PartitionKeysAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("type"))
}

type StorageDescriptorAttributes struct {
	ref terra.Reference
}

func (sd StorageDescriptorAttributes) InternalRef() terra.Reference {
	return sd.ref
}

func (sd StorageDescriptorAttributes) InternalWithRef(ref terra.Reference) StorageDescriptorAttributes {
	return StorageDescriptorAttributes{ref: ref}
}

func (sd StorageDescriptorAttributes) InternalTokens() hclwrite.Tokens {
	return sd.ref.InternalTokens()
}

func (sd StorageDescriptorAttributes) BucketColumns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sd.ref.Append("bucket_columns"))
}

func (sd StorageDescriptorAttributes) Compressed() terra.BoolValue {
	return terra.ReferenceAsBool(sd.ref.Append("compressed"))
}

func (sd StorageDescriptorAttributes) InputFormat() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("input_format"))
}

func (sd StorageDescriptorAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("location"))
}

func (sd StorageDescriptorAttributes) NumberOfBuckets() terra.NumberValue {
	return terra.ReferenceAsNumber(sd.ref.Append("number_of_buckets"))
}

func (sd StorageDescriptorAttributes) OutputFormat() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("output_format"))
}

func (sd StorageDescriptorAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sd.ref.Append("parameters"))
}

func (sd StorageDescriptorAttributes) StoredAsSubDirectories() terra.BoolValue {
	return terra.ReferenceAsBool(sd.ref.Append("stored_as_sub_directories"))
}

func (sd StorageDescriptorAttributes) Columns() terra.ListValue[ColumnsAttributes] {
	return terra.ReferenceAsList[ColumnsAttributes](sd.ref.Append("columns"))
}

func (sd StorageDescriptorAttributes) SchemaReference() terra.ListValue[SchemaReferenceAttributes] {
	return terra.ReferenceAsList[SchemaReferenceAttributes](sd.ref.Append("schema_reference"))
}

func (sd StorageDescriptorAttributes) SerDeInfo() terra.ListValue[SerDeInfoAttributes] {
	return terra.ReferenceAsList[SerDeInfoAttributes](sd.ref.Append("ser_de_info"))
}

func (sd StorageDescriptorAttributes) SkewedInfo() terra.ListValue[SkewedInfoAttributes] {
	return terra.ReferenceAsList[SkewedInfoAttributes](sd.ref.Append("skewed_info"))
}

func (sd StorageDescriptorAttributes) SortColumns() terra.ListValue[SortColumnsAttributes] {
	return terra.ReferenceAsList[SortColumnsAttributes](sd.ref.Append("sort_columns"))
}

type ColumnsAttributes struct {
	ref terra.Reference
}

func (c ColumnsAttributes) InternalRef() terra.Reference {
	return c.ref
}

func (c ColumnsAttributes) InternalWithRef(ref terra.Reference) ColumnsAttributes {
	return ColumnsAttributes{ref: ref}
}

func (c ColumnsAttributes) InternalTokens() hclwrite.Tokens {
	return c.ref.InternalTokens()
}

func (c ColumnsAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("comment"))
}

func (c ColumnsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c ColumnsAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("parameters"))
}

func (c ColumnsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("type"))
}

type SchemaReferenceAttributes struct {
	ref terra.Reference
}

func (sr SchemaReferenceAttributes) InternalRef() terra.Reference {
	return sr.ref
}

func (sr SchemaReferenceAttributes) InternalWithRef(ref terra.Reference) SchemaReferenceAttributes {
	return SchemaReferenceAttributes{ref: ref}
}

func (sr SchemaReferenceAttributes) InternalTokens() hclwrite.Tokens {
	return sr.ref.InternalTokens()
}

func (sr SchemaReferenceAttributes) SchemaVersionId() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("schema_version_id"))
}

func (sr SchemaReferenceAttributes) SchemaVersionNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(sr.ref.Append("schema_version_number"))
}

func (sr SchemaReferenceAttributes) SchemaId() terra.ListValue[SchemaIdAttributes] {
	return terra.ReferenceAsList[SchemaIdAttributes](sr.ref.Append("schema_id"))
}

type SchemaIdAttributes struct {
	ref terra.Reference
}

func (si SchemaIdAttributes) InternalRef() terra.Reference {
	return si.ref
}

func (si SchemaIdAttributes) InternalWithRef(ref terra.Reference) SchemaIdAttributes {
	return SchemaIdAttributes{ref: ref}
}

func (si SchemaIdAttributes) InternalTokens() hclwrite.Tokens {
	return si.ref.InternalTokens()
}

func (si SchemaIdAttributes) RegistryName() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("registry_name"))
}

func (si SchemaIdAttributes) SchemaArn() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("schema_arn"))
}

func (si SchemaIdAttributes) SchemaName() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("schema_name"))
}

type SerDeInfoAttributes struct {
	ref terra.Reference
}

func (sdi SerDeInfoAttributes) InternalRef() terra.Reference {
	return sdi.ref
}

func (sdi SerDeInfoAttributes) InternalWithRef(ref terra.Reference) SerDeInfoAttributes {
	return SerDeInfoAttributes{ref: ref}
}

func (sdi SerDeInfoAttributes) InternalTokens() hclwrite.Tokens {
	return sdi.ref.InternalTokens()
}

func (sdi SerDeInfoAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("name"))
}

func (sdi SerDeInfoAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sdi.ref.Append("parameters"))
}

func (sdi SerDeInfoAttributes) SerializationLibrary() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("serialization_library"))
}

type SkewedInfoAttributes struct {
	ref terra.Reference
}

func (si SkewedInfoAttributes) InternalRef() terra.Reference {
	return si.ref
}

func (si SkewedInfoAttributes) InternalWithRef(ref terra.Reference) SkewedInfoAttributes {
	return SkewedInfoAttributes{ref: ref}
}

func (si SkewedInfoAttributes) InternalTokens() hclwrite.Tokens {
	return si.ref.InternalTokens()
}

func (si SkewedInfoAttributes) SkewedColumnNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](si.ref.Append("skewed_column_names"))
}

func (si SkewedInfoAttributes) SkewedColumnValueLocationMaps() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](si.ref.Append("skewed_column_value_location_maps"))
}

func (si SkewedInfoAttributes) SkewedColumnValues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](si.ref.Append("skewed_column_values"))
}

type SortColumnsAttributes struct {
	ref terra.Reference
}

func (sc SortColumnsAttributes) InternalRef() terra.Reference {
	return sc.ref
}

func (sc SortColumnsAttributes) InternalWithRef(ref terra.Reference) SortColumnsAttributes {
	return SortColumnsAttributes{ref: ref}
}

func (sc SortColumnsAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc SortColumnsAttributes) Column() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("column"))
}

func (sc SortColumnsAttributes) SortOrder() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("sort_order"))
}

type TargetTableAttributes struct {
	ref terra.Reference
}

func (tt TargetTableAttributes) InternalRef() terra.Reference {
	return tt.ref
}

func (tt TargetTableAttributes) InternalWithRef(ref terra.Reference) TargetTableAttributes {
	return TargetTableAttributes{ref: ref}
}

func (tt TargetTableAttributes) InternalTokens() hclwrite.Tokens {
	return tt.ref.InternalTokens()
}

func (tt TargetTableAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("catalog_id"))
}

func (tt TargetTableAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("database_name"))
}

func (tt TargetTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("name"))
}

type PartitionIndexState struct {
	IndexName   string   `json:"index_name"`
	IndexStatus string   `json:"index_status"`
	Keys        []string `json:"keys"`
}

type PartitionKeysState struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type StorageDescriptorState struct {
	BucketColumns          []string               `json:"bucket_columns"`
	Compressed             bool                   `json:"compressed"`
	InputFormat            string                 `json:"input_format"`
	Location               string                 `json:"location"`
	NumberOfBuckets        float64                `json:"number_of_buckets"`
	OutputFormat           string                 `json:"output_format"`
	Parameters             map[string]string      `json:"parameters"`
	StoredAsSubDirectories bool                   `json:"stored_as_sub_directories"`
	Columns                []ColumnsState         `json:"columns"`
	SchemaReference        []SchemaReferenceState `json:"schema_reference"`
	SerDeInfo              []SerDeInfoState       `json:"ser_de_info"`
	SkewedInfo             []SkewedInfoState      `json:"skewed_info"`
	SortColumns            []SortColumnsState     `json:"sort_columns"`
}

type ColumnsState struct {
	Comment    string            `json:"comment"`
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
	Type       string            `json:"type"`
}

type SchemaReferenceState struct {
	SchemaVersionId     string          `json:"schema_version_id"`
	SchemaVersionNumber float64         `json:"schema_version_number"`
	SchemaId            []SchemaIdState `json:"schema_id"`
}

type SchemaIdState struct {
	RegistryName string `json:"registry_name"`
	SchemaArn    string `json:"schema_arn"`
	SchemaName   string `json:"schema_name"`
}

type SerDeInfoState struct {
	Name                 string            `json:"name"`
	Parameters           map[string]string `json:"parameters"`
	SerializationLibrary string            `json:"serialization_library"`
}

type SkewedInfoState struct {
	SkewedColumnNames             []string          `json:"skewed_column_names"`
	SkewedColumnValueLocationMaps map[string]string `json:"skewed_column_value_location_maps"`
	SkewedColumnValues            []string          `json:"skewed_column_values"`
}

type SortColumnsState struct {
	Column    string  `json:"column"`
	SortOrder float64 `json:"sort_order"`
}

type TargetTableState struct {
	CatalogId    string `json:"catalog_id"`
	DatabaseName string `json:"database_name"`
	Name         string `json:"name"`
}
