// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datacatalogentry

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type BigqueryDateShardedSpec struct{}

type BigqueryTableSpec struct {
	// TableSpec: min=0
	TableSpec []TableSpec `hcl:"table_spec,block" validate:"min=0"`
	// ViewSpec: min=0
	ViewSpec []ViewSpec `hcl:"view_spec,block" validate:"min=0"`
}

type TableSpec struct{}

type ViewSpec struct{}

type GcsFilesetSpec struct {
	// FilePatterns: list of string, required
	FilePatterns terra.ListValue[terra.StringValue] `hcl:"file_patterns,attr" validate:"required"`
	// SampleGcsFileSpecs: min=0
	SampleGcsFileSpecs []SampleGcsFileSpecs `hcl:"sample_gcs_file_specs,block" validate:"min=0"`
}

type SampleGcsFileSpecs struct{}

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

func (bts BigqueryTableSpecAttributes) TableSpec() terra.ListValue[TableSpecAttributes] {
	return terra.ReferenceAsList[TableSpecAttributes](bts.ref.Append("table_spec"))
}

func (bts BigqueryTableSpecAttributes) ViewSpec() terra.ListValue[ViewSpecAttributes] {
	return terra.ReferenceAsList[ViewSpecAttributes](bts.ref.Append("view_spec"))
}

type TableSpecAttributes struct {
	ref terra.Reference
}

func (ts TableSpecAttributes) InternalRef() (terra.Reference, error) {
	return ts.ref, nil
}

func (ts TableSpecAttributes) InternalWithRef(ref terra.Reference) TableSpecAttributes {
	return TableSpecAttributes{ref: ref}
}

func (ts TableSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ts.ref.InternalTokens()
}

func (ts TableSpecAttributes) GroupedEntry() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("grouped_entry"))
}

type ViewSpecAttributes struct {
	ref terra.Reference
}

func (vs ViewSpecAttributes) InternalRef() (terra.Reference, error) {
	return vs.ref, nil
}

func (vs ViewSpecAttributes) InternalWithRef(ref terra.Reference) ViewSpecAttributes {
	return ViewSpecAttributes{ref: ref}
}

func (vs ViewSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vs.ref.InternalTokens()
}

func (vs ViewSpecAttributes) ViewQuery() terra.StringValue {
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

func (gfs GcsFilesetSpecAttributes) SampleGcsFileSpecs() terra.ListValue[SampleGcsFileSpecsAttributes] {
	return terra.ReferenceAsList[SampleGcsFileSpecsAttributes](gfs.ref.Append("sample_gcs_file_specs"))
}

type SampleGcsFileSpecsAttributes struct {
	ref terra.Reference
}

func (sgfs SampleGcsFileSpecsAttributes) InternalRef() (terra.Reference, error) {
	return sgfs.ref, nil
}

func (sgfs SampleGcsFileSpecsAttributes) InternalWithRef(ref terra.Reference) SampleGcsFileSpecsAttributes {
	return SampleGcsFileSpecsAttributes{ref: ref}
}

func (sgfs SampleGcsFileSpecsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sgfs.ref.InternalTokens()
}

func (sgfs SampleGcsFileSpecsAttributes) FilePath() terra.StringValue {
	return terra.ReferenceAsString(sgfs.ref.Append("file_path"))
}

func (sgfs SampleGcsFileSpecsAttributes) SizeBytes() terra.NumberValue {
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
	TableSourceType string           `json:"table_source_type"`
	TableSpec       []TableSpecState `json:"table_spec"`
	ViewSpec        []ViewSpecState  `json:"view_spec"`
}

type TableSpecState struct {
	GroupedEntry string `json:"grouped_entry"`
}

type ViewSpecState struct {
	ViewQuery string `json:"view_query"`
}

type GcsFilesetSpecState struct {
	FilePatterns       []string                  `json:"file_patterns"`
	SampleGcsFileSpecs []SampleGcsFileSpecsState `json:"sample_gcs_file_specs"`
}

type SampleGcsFileSpecsState struct {
	FilePath  string  `json:"file_path"`
	SizeBytes float64 `json:"size_bytes"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
