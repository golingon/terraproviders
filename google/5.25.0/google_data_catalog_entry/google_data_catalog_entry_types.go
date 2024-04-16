// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_data_catalog_entry

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type GcsFilesetSpec struct {
	// FilePatterns: list of string, required
	FilePatterns terra.ListValue[terra.StringValue] `hcl:"file_patterns,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type BigqueryDateShardedSpecAttributes struct {
	ref terra.Reference
}

func (bdss BigqueryDateShardedSpecAttributes) InternalRef() (terra.Reference, error) {
	return bdss.ref, nil
}

func (bdss BigqueryDateShardedSpecAttributes) InternalWithRef(ref terra.Reference) BigqueryDateShardedSpecAttributes {
	return BigqueryDateShardedSpecAttributes{ref: ref}
}

func (bdss BigqueryDateShardedSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bdss.ref.InternalTokens()
}

func (bdss BigqueryDateShardedSpecAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(bdss.ref.Append("dataset"))
}

func (bdss BigqueryDateShardedSpecAttributes) ShardCount() terra.NumberValue {
	return terra.ReferenceAsNumber(bdss.ref.Append("shard_count"))
}

func (bdss BigqueryDateShardedSpecAttributes) TablePrefix() terra.StringValue {
	return terra.ReferenceAsString(bdss.ref.Append("table_prefix"))
}

type BigqueryTableSpecAttributes struct {
	ref terra.Reference
}

func (bts BigqueryTableSpecAttributes) InternalRef() (terra.Reference, error) {
	return bts.ref, nil
}

func (bts BigqueryTableSpecAttributes) InternalWithRef(ref terra.Reference) BigqueryTableSpecAttributes {
	return BigqueryTableSpecAttributes{ref: ref}
}

func (bts BigqueryTableSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bts.ref.InternalTokens()
}

func (bts BigqueryTableSpecAttributes) TableSourceType() terra.StringValue {
	return terra.ReferenceAsString(bts.ref.Append("table_source_type"))
}

func (bts BigqueryTableSpecAttributes) TableSpec() terra.ListValue[BigqueryTableSpecTableSpecAttributes] {
	return terra.ReferenceAsList[BigqueryTableSpecTableSpecAttributes](bts.ref.Append("table_spec"))
}

func (bts BigqueryTableSpecAttributes) ViewSpec() terra.ListValue[BigqueryTableSpecViewSpecAttributes] {
	return terra.ReferenceAsList[BigqueryTableSpecViewSpecAttributes](bts.ref.Append("view_spec"))
}

type BigqueryTableSpecTableSpecAttributes struct {
	ref terra.Reference
}

func (ts BigqueryTableSpecTableSpecAttributes) InternalRef() (terra.Reference, error) {
	return ts.ref, nil
}

func (ts BigqueryTableSpecTableSpecAttributes) InternalWithRef(ref terra.Reference) BigqueryTableSpecTableSpecAttributes {
	return BigqueryTableSpecTableSpecAttributes{ref: ref}
}

func (ts BigqueryTableSpecTableSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ts.ref.InternalTokens()
}

func (ts BigqueryTableSpecTableSpecAttributes) GroupedEntry() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("grouped_entry"))
}

type BigqueryTableSpecViewSpecAttributes struct {
	ref terra.Reference
}

func (vs BigqueryTableSpecViewSpecAttributes) InternalRef() (terra.Reference, error) {
	return vs.ref, nil
}

func (vs BigqueryTableSpecViewSpecAttributes) InternalWithRef(ref terra.Reference) BigqueryTableSpecViewSpecAttributes {
	return BigqueryTableSpecViewSpecAttributes{ref: ref}
}

func (vs BigqueryTableSpecViewSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vs.ref.InternalTokens()
}

func (vs BigqueryTableSpecViewSpecAttributes) ViewQuery() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("view_query"))
}

type GcsFilesetSpecAttributes struct {
	ref terra.Reference
}

func (gfs GcsFilesetSpecAttributes) InternalRef() (terra.Reference, error) {
	return gfs.ref, nil
}

func (gfs GcsFilesetSpecAttributes) InternalWithRef(ref terra.Reference) GcsFilesetSpecAttributes {
	return GcsFilesetSpecAttributes{ref: ref}
}

func (gfs GcsFilesetSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gfs.ref.InternalTokens()
}

func (gfs GcsFilesetSpecAttributes) FilePatterns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gfs.ref.Append("file_patterns"))
}

func (gfs GcsFilesetSpecAttributes) SampleGcsFileSpecs() terra.ListValue[GcsFilesetSpecSampleGcsFileSpecsAttributes] {
	return terra.ReferenceAsList[GcsFilesetSpecSampleGcsFileSpecsAttributes](gfs.ref.Append("sample_gcs_file_specs"))
}

type GcsFilesetSpecSampleGcsFileSpecsAttributes struct {
	ref terra.Reference
}

func (sgfs GcsFilesetSpecSampleGcsFileSpecsAttributes) InternalRef() (terra.Reference, error) {
	return sgfs.ref, nil
}

func (sgfs GcsFilesetSpecSampleGcsFileSpecsAttributes) InternalWithRef(ref terra.Reference) GcsFilesetSpecSampleGcsFileSpecsAttributes {
	return GcsFilesetSpecSampleGcsFileSpecsAttributes{ref: ref}
}

func (sgfs GcsFilesetSpecSampleGcsFileSpecsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sgfs.ref.InternalTokens()
}

func (sgfs GcsFilesetSpecSampleGcsFileSpecsAttributes) FilePath() terra.StringValue {
	return terra.ReferenceAsString(sgfs.ref.Append("file_path"))
}

func (sgfs GcsFilesetSpecSampleGcsFileSpecsAttributes) SizeBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(sgfs.ref.Append("size_bytes"))
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

type BigqueryDateShardedSpecState struct {
	Dataset     string  `json:"dataset"`
	ShardCount  float64 `json:"shard_count"`
	TablePrefix string  `json:"table_prefix"`
}

type BigqueryTableSpecState struct {
	TableSourceType string                            `json:"table_source_type"`
	TableSpec       []BigqueryTableSpecTableSpecState `json:"table_spec"`
	ViewSpec        []BigqueryTableSpecViewSpecState  `json:"view_spec"`
}

type BigqueryTableSpecTableSpecState struct {
	GroupedEntry string `json:"grouped_entry"`
}

type BigqueryTableSpecViewSpecState struct {
	ViewQuery string `json:"view_query"`
}

type GcsFilesetSpecState struct {
	FilePatterns       []string                                `json:"file_patterns"`
	SampleGcsFileSpecs []GcsFilesetSpecSampleGcsFileSpecsState `json:"sample_gcs_file_specs"`
}

type GcsFilesetSpecSampleGcsFileSpecsState struct {
	FilePath  string  `json:"file_path"`
	SizeBytes float64 `json:"size_bytes"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
