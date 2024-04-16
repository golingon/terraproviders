// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_data_loss_prevention_stored_info_type

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Dictionary struct {
	// DictionaryCloudStoragePath: optional
	CloudStoragePath *DictionaryCloudStoragePath `hcl:"cloud_storage_path,block"`
	// DictionaryWordList: optional
	WordList *DictionaryWordList `hcl:"word_list,block"`
}

type DictionaryCloudStoragePath struct {
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
}

type DictionaryWordList struct {
	// Words: list of string, required
	Words terra.ListValue[terra.StringValue] `hcl:"words,attr" validate:"required"`
}

type LargeCustomDictionary struct {
	// LargeCustomDictionaryBigQueryField: optional
	BigQueryField *LargeCustomDictionaryBigQueryField `hcl:"big_query_field,block"`
	// LargeCustomDictionaryCloudStorageFileSet: optional
	CloudStorageFileSet *LargeCustomDictionaryCloudStorageFileSet `hcl:"cloud_storage_file_set,block"`
	// LargeCustomDictionaryOutputPath: required
	OutputPath *LargeCustomDictionaryOutputPath `hcl:"output_path,block" validate:"required"`
}

type LargeCustomDictionaryBigQueryField struct {
	// LargeCustomDictionaryBigQueryFieldField: required
	Field *LargeCustomDictionaryBigQueryFieldField `hcl:"field,block" validate:"required"`
	// LargeCustomDictionaryBigQueryFieldTable: required
	Table *LargeCustomDictionaryBigQueryFieldTable `hcl:"table,block" validate:"required"`
}

type LargeCustomDictionaryBigQueryFieldField struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type LargeCustomDictionaryBigQueryFieldTable struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// ProjectId: string, required
	ProjectId terra.StringValue `hcl:"project_id,attr" validate:"required"`
	// TableId: string, required
	TableId terra.StringValue `hcl:"table_id,attr" validate:"required"`
}

type LargeCustomDictionaryCloudStorageFileSet struct {
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
}

type LargeCustomDictionaryOutputPath struct {
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
}

type Regex struct {
	// GroupIndexes: list of number, optional
	GroupIndexes terra.ListValue[terra.NumberValue] `hcl:"group_indexes,attr"`
	// Pattern: string, required
	Pattern terra.StringValue `hcl:"pattern,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DictionaryAttributes struct {
	ref terra.Reference
}

func (d DictionaryAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DictionaryAttributes) InternalWithRef(ref terra.Reference) DictionaryAttributes {
	return DictionaryAttributes{ref: ref}
}

func (d DictionaryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DictionaryAttributes) CloudStoragePath() terra.ListValue[DictionaryCloudStoragePathAttributes] {
	return terra.ReferenceAsList[DictionaryCloudStoragePathAttributes](d.ref.Append("cloud_storage_path"))
}

func (d DictionaryAttributes) WordList() terra.ListValue[DictionaryWordListAttributes] {
	return terra.ReferenceAsList[DictionaryWordListAttributes](d.ref.Append("word_list"))
}

type DictionaryCloudStoragePathAttributes struct {
	ref terra.Reference
}

func (csp DictionaryCloudStoragePathAttributes) InternalRef() (terra.Reference, error) {
	return csp.ref, nil
}

func (csp DictionaryCloudStoragePathAttributes) InternalWithRef(ref terra.Reference) DictionaryCloudStoragePathAttributes {
	return DictionaryCloudStoragePathAttributes{ref: ref}
}

func (csp DictionaryCloudStoragePathAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return csp.ref.InternalTokens()
}

func (csp DictionaryCloudStoragePathAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("path"))
}

type DictionaryWordListAttributes struct {
	ref terra.Reference
}

func (wl DictionaryWordListAttributes) InternalRef() (terra.Reference, error) {
	return wl.ref, nil
}

func (wl DictionaryWordListAttributes) InternalWithRef(ref terra.Reference) DictionaryWordListAttributes {
	return DictionaryWordListAttributes{ref: ref}
}

func (wl DictionaryWordListAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wl.ref.InternalTokens()
}

func (wl DictionaryWordListAttributes) Words() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wl.ref.Append("words"))
}

type LargeCustomDictionaryAttributes struct {
	ref terra.Reference
}

func (lcd LargeCustomDictionaryAttributes) InternalRef() (terra.Reference, error) {
	return lcd.ref, nil
}

func (lcd LargeCustomDictionaryAttributes) InternalWithRef(ref terra.Reference) LargeCustomDictionaryAttributes {
	return LargeCustomDictionaryAttributes{ref: ref}
}

func (lcd LargeCustomDictionaryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lcd.ref.InternalTokens()
}

func (lcd LargeCustomDictionaryAttributes) BigQueryField() terra.ListValue[LargeCustomDictionaryBigQueryFieldAttributes] {
	return terra.ReferenceAsList[LargeCustomDictionaryBigQueryFieldAttributes](lcd.ref.Append("big_query_field"))
}

func (lcd LargeCustomDictionaryAttributes) CloudStorageFileSet() terra.ListValue[LargeCustomDictionaryCloudStorageFileSetAttributes] {
	return terra.ReferenceAsList[LargeCustomDictionaryCloudStorageFileSetAttributes](lcd.ref.Append("cloud_storage_file_set"))
}

func (lcd LargeCustomDictionaryAttributes) OutputPath() terra.ListValue[LargeCustomDictionaryOutputPathAttributes] {
	return terra.ReferenceAsList[LargeCustomDictionaryOutputPathAttributes](lcd.ref.Append("output_path"))
}

type LargeCustomDictionaryBigQueryFieldAttributes struct {
	ref terra.Reference
}

func (bqf LargeCustomDictionaryBigQueryFieldAttributes) InternalRef() (terra.Reference, error) {
	return bqf.ref, nil
}

func (bqf LargeCustomDictionaryBigQueryFieldAttributes) InternalWithRef(ref terra.Reference) LargeCustomDictionaryBigQueryFieldAttributes {
	return LargeCustomDictionaryBigQueryFieldAttributes{ref: ref}
}

func (bqf LargeCustomDictionaryBigQueryFieldAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bqf.ref.InternalTokens()
}

func (bqf LargeCustomDictionaryBigQueryFieldAttributes) Field() terra.ListValue[LargeCustomDictionaryBigQueryFieldFieldAttributes] {
	return terra.ReferenceAsList[LargeCustomDictionaryBigQueryFieldFieldAttributes](bqf.ref.Append("field"))
}

func (bqf LargeCustomDictionaryBigQueryFieldAttributes) Table() terra.ListValue[LargeCustomDictionaryBigQueryFieldTableAttributes] {
	return terra.ReferenceAsList[LargeCustomDictionaryBigQueryFieldTableAttributes](bqf.ref.Append("table"))
}

type LargeCustomDictionaryBigQueryFieldFieldAttributes struct {
	ref terra.Reference
}

func (f LargeCustomDictionaryBigQueryFieldFieldAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f LargeCustomDictionaryBigQueryFieldFieldAttributes) InternalWithRef(ref terra.Reference) LargeCustomDictionaryBigQueryFieldFieldAttributes {
	return LargeCustomDictionaryBigQueryFieldFieldAttributes{ref: ref}
}

func (f LargeCustomDictionaryBigQueryFieldFieldAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f LargeCustomDictionaryBigQueryFieldFieldAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

type LargeCustomDictionaryBigQueryFieldTableAttributes struct {
	ref terra.Reference
}

func (t LargeCustomDictionaryBigQueryFieldTableAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t LargeCustomDictionaryBigQueryFieldTableAttributes) InternalWithRef(ref terra.Reference) LargeCustomDictionaryBigQueryFieldTableAttributes {
	return LargeCustomDictionaryBigQueryFieldTableAttributes{ref: ref}
}

func (t LargeCustomDictionaryBigQueryFieldTableAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t LargeCustomDictionaryBigQueryFieldTableAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("dataset_id"))
}

func (t LargeCustomDictionaryBigQueryFieldTableAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("project_id"))
}

func (t LargeCustomDictionaryBigQueryFieldTableAttributes) TableId() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("table_id"))
}

type LargeCustomDictionaryCloudStorageFileSetAttributes struct {
	ref terra.Reference
}

func (csfs LargeCustomDictionaryCloudStorageFileSetAttributes) InternalRef() (terra.Reference, error) {
	return csfs.ref, nil
}

func (csfs LargeCustomDictionaryCloudStorageFileSetAttributes) InternalWithRef(ref terra.Reference) LargeCustomDictionaryCloudStorageFileSetAttributes {
	return LargeCustomDictionaryCloudStorageFileSetAttributes{ref: ref}
}

func (csfs LargeCustomDictionaryCloudStorageFileSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return csfs.ref.InternalTokens()
}

func (csfs LargeCustomDictionaryCloudStorageFileSetAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(csfs.ref.Append("url"))
}

type LargeCustomDictionaryOutputPathAttributes struct {
	ref terra.Reference
}

func (op LargeCustomDictionaryOutputPathAttributes) InternalRef() (terra.Reference, error) {
	return op.ref, nil
}

func (op LargeCustomDictionaryOutputPathAttributes) InternalWithRef(ref terra.Reference) LargeCustomDictionaryOutputPathAttributes {
	return LargeCustomDictionaryOutputPathAttributes{ref: ref}
}

func (op LargeCustomDictionaryOutputPathAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return op.ref.InternalTokens()
}

func (op LargeCustomDictionaryOutputPathAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(op.ref.Append("path"))
}

type RegexAttributes struct {
	ref terra.Reference
}

func (r RegexAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RegexAttributes) InternalWithRef(ref terra.Reference) RegexAttributes {
	return RegexAttributes{ref: ref}
}

func (r RegexAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RegexAttributes) GroupIndexes() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](r.ref.Append("group_indexes"))
}

func (r RegexAttributes) Pattern() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("pattern"))
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

type DictionaryState struct {
	CloudStoragePath []DictionaryCloudStoragePathState `json:"cloud_storage_path"`
	WordList         []DictionaryWordListState         `json:"word_list"`
}

type DictionaryCloudStoragePathState struct {
	Path string `json:"path"`
}

type DictionaryWordListState struct {
	Words []string `json:"words"`
}

type LargeCustomDictionaryState struct {
	BigQueryField       []LargeCustomDictionaryBigQueryFieldState       `json:"big_query_field"`
	CloudStorageFileSet []LargeCustomDictionaryCloudStorageFileSetState `json:"cloud_storage_file_set"`
	OutputPath          []LargeCustomDictionaryOutputPathState          `json:"output_path"`
}

type LargeCustomDictionaryBigQueryFieldState struct {
	Field []LargeCustomDictionaryBigQueryFieldFieldState `json:"field"`
	Table []LargeCustomDictionaryBigQueryFieldTableState `json:"table"`
}

type LargeCustomDictionaryBigQueryFieldFieldState struct {
	Name string `json:"name"`
}

type LargeCustomDictionaryBigQueryFieldTableState struct {
	DatasetId string `json:"dataset_id"`
	ProjectId string `json:"project_id"`
	TableId   string `json:"table_id"`
}

type LargeCustomDictionaryCloudStorageFileSetState struct {
	Url string `json:"url"`
}

type LargeCustomDictionaryOutputPathState struct {
	Path string `json:"path"`
}

type RegexState struct {
	GroupIndexes []float64 `json:"group_indexes"`
	Pattern      string    `json:"pattern"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
