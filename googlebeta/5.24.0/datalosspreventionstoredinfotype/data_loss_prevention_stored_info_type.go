// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datalosspreventionstoredinfotype

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Dictionary struct {
	// CloudStoragePath: optional
	CloudStoragePath *CloudStoragePath `hcl:"cloud_storage_path,block"`
	// WordList: optional
	WordList *WordList `hcl:"word_list,block"`
}

type CloudStoragePath struct {
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
}

type WordList struct {
	// Words: list of string, required
	Words terra.ListValue[terra.StringValue] `hcl:"words,attr" validate:"required"`
}

type LargeCustomDictionary struct {
	// BigQueryField: optional
	BigQueryField *BigQueryField `hcl:"big_query_field,block"`
	// CloudStorageFileSet: optional
	CloudStorageFileSet *CloudStorageFileSet `hcl:"cloud_storage_file_set,block"`
	// OutputPath: required
	OutputPath *OutputPath `hcl:"output_path,block" validate:"required"`
}

type BigQueryField struct {
	// Field: required
	Field *Field `hcl:"field,block" validate:"required"`
	// Table: required
	Table *Table `hcl:"table,block" validate:"required"`
}

type Field struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type Table struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// ProjectId: string, required
	ProjectId terra.StringValue `hcl:"project_id,attr" validate:"required"`
	// TableId: string, required
	TableId terra.StringValue `hcl:"table_id,attr" validate:"required"`
}

type CloudStorageFileSet struct {
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
}

type OutputPath struct {
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

func (d DictionaryAttributes) CloudStoragePath() terra.ListValue[CloudStoragePathAttributes] {
	return terra.ReferenceAsList[CloudStoragePathAttributes](d.ref.Append("cloud_storage_path"))
}

func (d DictionaryAttributes) WordList() terra.ListValue[WordListAttributes] {
	return terra.ReferenceAsList[WordListAttributes](d.ref.Append("word_list"))
}

type CloudStoragePathAttributes struct {
	ref terra.Reference
}

func (csp CloudStoragePathAttributes) InternalRef() (terra.Reference, error) {
	return csp.ref, nil
}

func (csp CloudStoragePathAttributes) InternalWithRef(ref terra.Reference) CloudStoragePathAttributes {
	return CloudStoragePathAttributes{ref: ref}
}

func (csp CloudStoragePathAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return csp.ref.InternalTokens()
}

func (csp CloudStoragePathAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("path"))
}

type WordListAttributes struct {
	ref terra.Reference
}

func (wl WordListAttributes) InternalRef() (terra.Reference, error) {
	return wl.ref, nil
}

func (wl WordListAttributes) InternalWithRef(ref terra.Reference) WordListAttributes {
	return WordListAttributes{ref: ref}
}

func (wl WordListAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wl.ref.InternalTokens()
}

func (wl WordListAttributes) Words() terra.ListValue[terra.StringValue] {
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

func (lcd LargeCustomDictionaryAttributes) BigQueryField() terra.ListValue[BigQueryFieldAttributes] {
	return terra.ReferenceAsList[BigQueryFieldAttributes](lcd.ref.Append("big_query_field"))
}

func (lcd LargeCustomDictionaryAttributes) CloudStorageFileSet() terra.ListValue[CloudStorageFileSetAttributes] {
	return terra.ReferenceAsList[CloudStorageFileSetAttributes](lcd.ref.Append("cloud_storage_file_set"))
}

func (lcd LargeCustomDictionaryAttributes) OutputPath() terra.ListValue[OutputPathAttributes] {
	return terra.ReferenceAsList[OutputPathAttributes](lcd.ref.Append("output_path"))
}

type BigQueryFieldAttributes struct {
	ref terra.Reference
}

func (bqf BigQueryFieldAttributes) InternalRef() (terra.Reference, error) {
	return bqf.ref, nil
}

func (bqf BigQueryFieldAttributes) InternalWithRef(ref terra.Reference) BigQueryFieldAttributes {
	return BigQueryFieldAttributes{ref: ref}
}

func (bqf BigQueryFieldAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bqf.ref.InternalTokens()
}

func (bqf BigQueryFieldAttributes) Field() terra.ListValue[FieldAttributes] {
	return terra.ReferenceAsList[FieldAttributes](bqf.ref.Append("field"))
}

func (bqf BigQueryFieldAttributes) Table() terra.ListValue[TableAttributes] {
	return terra.ReferenceAsList[TableAttributes](bqf.ref.Append("table"))
}

type FieldAttributes struct {
	ref terra.Reference
}

func (f FieldAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FieldAttributes) InternalWithRef(ref terra.Reference) FieldAttributes {
	return FieldAttributes{ref: ref}
}

func (f FieldAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FieldAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

type TableAttributes struct {
	ref terra.Reference
}

func (t TableAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TableAttributes) InternalWithRef(ref terra.Reference) TableAttributes {
	return TableAttributes{ref: ref}
}

func (t TableAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TableAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("dataset_id"))
}

func (t TableAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("project_id"))
}

func (t TableAttributes) TableId() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("table_id"))
}

type CloudStorageFileSetAttributes struct {
	ref terra.Reference
}

func (csfs CloudStorageFileSetAttributes) InternalRef() (terra.Reference, error) {
	return csfs.ref, nil
}

func (csfs CloudStorageFileSetAttributes) InternalWithRef(ref terra.Reference) CloudStorageFileSetAttributes {
	return CloudStorageFileSetAttributes{ref: ref}
}

func (csfs CloudStorageFileSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return csfs.ref.InternalTokens()
}

func (csfs CloudStorageFileSetAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(csfs.ref.Append("url"))
}

type OutputPathAttributes struct {
	ref terra.Reference
}

func (op OutputPathAttributes) InternalRef() (terra.Reference, error) {
	return op.ref, nil
}

func (op OutputPathAttributes) InternalWithRef(ref terra.Reference) OutputPathAttributes {
	return OutputPathAttributes{ref: ref}
}

func (op OutputPathAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return op.ref.InternalTokens()
}

func (op OutputPathAttributes) Path() terra.StringValue {
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
	CloudStoragePath []CloudStoragePathState `json:"cloud_storage_path"`
	WordList         []WordListState         `json:"word_list"`
}

type CloudStoragePathState struct {
	Path string `json:"path"`
}

type WordListState struct {
	Words []string `json:"words"`
}

type LargeCustomDictionaryState struct {
	BigQueryField       []BigQueryFieldState       `json:"big_query_field"`
	CloudStorageFileSet []CloudStorageFileSetState `json:"cloud_storage_file_set"`
	OutputPath          []OutputPathState          `json:"output_path"`
}

type BigQueryFieldState struct {
	Field []FieldState `json:"field"`
	Table []TableState `json:"table"`
}

type FieldState struct {
	Name string `json:"name"`
}

type TableState struct {
	DatasetId string `json:"dataset_id"`
	ProjectId string `json:"project_id"`
	TableId   string `json:"table_id"`
}

type CloudStorageFileSetState struct {
	Url string `json:"url"`
}

type OutputPathState struct {
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
