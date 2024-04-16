// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_storage_insights_report_config

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CsvOptions struct {
	// Delimiter: string, optional
	Delimiter terra.StringValue `hcl:"delimiter,attr"`
	// HeaderRequired: bool, optional
	HeaderRequired terra.BoolValue `hcl:"header_required,attr"`
	// RecordSeparator: string, optional
	RecordSeparator terra.StringValue `hcl:"record_separator,attr"`
}

type FrequencyOptions struct {
	// Frequency: string, required
	Frequency terra.StringValue `hcl:"frequency,attr" validate:"required"`
	// FrequencyOptionsEndDate: required
	EndDate *FrequencyOptionsEndDate `hcl:"end_date,block" validate:"required"`
	// FrequencyOptionsStartDate: required
	StartDate *FrequencyOptionsStartDate `hcl:"start_date,block" validate:"required"`
}

type FrequencyOptionsEndDate struct {
	// Day: number, required
	Day terra.NumberValue `hcl:"day,attr" validate:"required"`
	// Month: number, required
	Month terra.NumberValue `hcl:"month,attr" validate:"required"`
	// Year: number, required
	Year terra.NumberValue `hcl:"year,attr" validate:"required"`
}

type FrequencyOptionsStartDate struct {
	// Day: number, required
	Day terra.NumberValue `hcl:"day,attr" validate:"required"`
	// Month: number, required
	Month terra.NumberValue `hcl:"month,attr" validate:"required"`
	// Year: number, required
	Year terra.NumberValue `hcl:"year,attr" validate:"required"`
}

type ObjectMetadataReportOptions struct {
	// MetadataFields: list of string, required
	MetadataFields terra.ListValue[terra.StringValue] `hcl:"metadata_fields,attr" validate:"required"`
	// ObjectMetadataReportOptionsStorageDestinationOptions: required
	StorageDestinationOptions *ObjectMetadataReportOptionsStorageDestinationOptions `hcl:"storage_destination_options,block" validate:"required"`
	// ObjectMetadataReportOptionsStorageFilters: optional
	StorageFilters *ObjectMetadataReportOptionsStorageFilters `hcl:"storage_filters,block"`
}

type ObjectMetadataReportOptionsStorageDestinationOptions struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// DestinationPath: string, optional
	DestinationPath terra.StringValue `hcl:"destination_path,attr"`
}

type ObjectMetadataReportOptionsStorageFilters struct {
	// Bucket: string, optional
	Bucket terra.StringValue `hcl:"bucket,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type CsvOptionsAttributes struct {
	ref terra.Reference
}

func (co CsvOptionsAttributes) InternalRef() (terra.Reference, error) {
	return co.ref, nil
}

func (co CsvOptionsAttributes) InternalWithRef(ref terra.Reference) CsvOptionsAttributes {
	return CsvOptionsAttributes{ref: ref}
}

func (co CsvOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return co.ref.InternalTokens()
}

func (co CsvOptionsAttributes) Delimiter() terra.StringValue {
	return terra.ReferenceAsString(co.ref.Append("delimiter"))
}

func (co CsvOptionsAttributes) HeaderRequired() terra.BoolValue {
	return terra.ReferenceAsBool(co.ref.Append("header_required"))
}

func (co CsvOptionsAttributes) RecordSeparator() terra.StringValue {
	return terra.ReferenceAsString(co.ref.Append("record_separator"))
}

type FrequencyOptionsAttributes struct {
	ref terra.Reference
}

func (fo FrequencyOptionsAttributes) InternalRef() (terra.Reference, error) {
	return fo.ref, nil
}

func (fo FrequencyOptionsAttributes) InternalWithRef(ref terra.Reference) FrequencyOptionsAttributes {
	return FrequencyOptionsAttributes{ref: ref}
}

func (fo FrequencyOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fo.ref.InternalTokens()
}

func (fo FrequencyOptionsAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(fo.ref.Append("frequency"))
}

func (fo FrequencyOptionsAttributes) EndDate() terra.ListValue[FrequencyOptionsEndDateAttributes] {
	return terra.ReferenceAsList[FrequencyOptionsEndDateAttributes](fo.ref.Append("end_date"))
}

func (fo FrequencyOptionsAttributes) StartDate() terra.ListValue[FrequencyOptionsStartDateAttributes] {
	return terra.ReferenceAsList[FrequencyOptionsStartDateAttributes](fo.ref.Append("start_date"))
}

type FrequencyOptionsEndDateAttributes struct {
	ref terra.Reference
}

func (ed FrequencyOptionsEndDateAttributes) InternalRef() (terra.Reference, error) {
	return ed.ref, nil
}

func (ed FrequencyOptionsEndDateAttributes) InternalWithRef(ref terra.Reference) FrequencyOptionsEndDateAttributes {
	return FrequencyOptionsEndDateAttributes{ref: ref}
}

func (ed FrequencyOptionsEndDateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ed.ref.InternalTokens()
}

func (ed FrequencyOptionsEndDateAttributes) Day() terra.NumberValue {
	return terra.ReferenceAsNumber(ed.ref.Append("day"))
}

func (ed FrequencyOptionsEndDateAttributes) Month() terra.NumberValue {
	return terra.ReferenceAsNumber(ed.ref.Append("month"))
}

func (ed FrequencyOptionsEndDateAttributes) Year() terra.NumberValue {
	return terra.ReferenceAsNumber(ed.ref.Append("year"))
}

type FrequencyOptionsStartDateAttributes struct {
	ref terra.Reference
}

func (sd FrequencyOptionsStartDateAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd FrequencyOptionsStartDateAttributes) InternalWithRef(ref terra.Reference) FrequencyOptionsStartDateAttributes {
	return FrequencyOptionsStartDateAttributes{ref: ref}
}

func (sd FrequencyOptionsStartDateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd FrequencyOptionsStartDateAttributes) Day() terra.NumberValue {
	return terra.ReferenceAsNumber(sd.ref.Append("day"))
}

func (sd FrequencyOptionsStartDateAttributes) Month() terra.NumberValue {
	return terra.ReferenceAsNumber(sd.ref.Append("month"))
}

func (sd FrequencyOptionsStartDateAttributes) Year() terra.NumberValue {
	return terra.ReferenceAsNumber(sd.ref.Append("year"))
}

type ObjectMetadataReportOptionsAttributes struct {
	ref terra.Reference
}

func (omro ObjectMetadataReportOptionsAttributes) InternalRef() (terra.Reference, error) {
	return omro.ref, nil
}

func (omro ObjectMetadataReportOptionsAttributes) InternalWithRef(ref terra.Reference) ObjectMetadataReportOptionsAttributes {
	return ObjectMetadataReportOptionsAttributes{ref: ref}
}

func (omro ObjectMetadataReportOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return omro.ref.InternalTokens()
}

func (omro ObjectMetadataReportOptionsAttributes) MetadataFields() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](omro.ref.Append("metadata_fields"))
}

func (omro ObjectMetadataReportOptionsAttributes) StorageDestinationOptions() terra.ListValue[ObjectMetadataReportOptionsStorageDestinationOptionsAttributes] {
	return terra.ReferenceAsList[ObjectMetadataReportOptionsStorageDestinationOptionsAttributes](omro.ref.Append("storage_destination_options"))
}

func (omro ObjectMetadataReportOptionsAttributes) StorageFilters() terra.ListValue[ObjectMetadataReportOptionsStorageFiltersAttributes] {
	return terra.ReferenceAsList[ObjectMetadataReportOptionsStorageFiltersAttributes](omro.ref.Append("storage_filters"))
}

type ObjectMetadataReportOptionsStorageDestinationOptionsAttributes struct {
	ref terra.Reference
}

func (sdo ObjectMetadataReportOptionsStorageDestinationOptionsAttributes) InternalRef() (terra.Reference, error) {
	return sdo.ref, nil
}

func (sdo ObjectMetadataReportOptionsStorageDestinationOptionsAttributes) InternalWithRef(ref terra.Reference) ObjectMetadataReportOptionsStorageDestinationOptionsAttributes {
	return ObjectMetadataReportOptionsStorageDestinationOptionsAttributes{ref: ref}
}

func (sdo ObjectMetadataReportOptionsStorageDestinationOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sdo.ref.InternalTokens()
}

func (sdo ObjectMetadataReportOptionsStorageDestinationOptionsAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sdo.ref.Append("bucket"))
}

func (sdo ObjectMetadataReportOptionsStorageDestinationOptionsAttributes) DestinationPath() terra.StringValue {
	return terra.ReferenceAsString(sdo.ref.Append("destination_path"))
}

type ObjectMetadataReportOptionsStorageFiltersAttributes struct {
	ref terra.Reference
}

func (sf ObjectMetadataReportOptionsStorageFiltersAttributes) InternalRef() (terra.Reference, error) {
	return sf.ref, nil
}

func (sf ObjectMetadataReportOptionsStorageFiltersAttributes) InternalWithRef(ref terra.Reference) ObjectMetadataReportOptionsStorageFiltersAttributes {
	return ObjectMetadataReportOptionsStorageFiltersAttributes{ref: ref}
}

func (sf ObjectMetadataReportOptionsStorageFiltersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sf.ref.InternalTokens()
}

func (sf ObjectMetadataReportOptionsStorageFiltersAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sf.ref.Append("bucket"))
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

type CsvOptionsState struct {
	Delimiter       string `json:"delimiter"`
	HeaderRequired  bool   `json:"header_required"`
	RecordSeparator string `json:"record_separator"`
}

type FrequencyOptionsState struct {
	Frequency string                           `json:"frequency"`
	EndDate   []FrequencyOptionsEndDateState   `json:"end_date"`
	StartDate []FrequencyOptionsStartDateState `json:"start_date"`
}

type FrequencyOptionsEndDateState struct {
	Day   float64 `json:"day"`
	Month float64 `json:"month"`
	Year  float64 `json:"year"`
}

type FrequencyOptionsStartDateState struct {
	Day   float64 `json:"day"`
	Month float64 `json:"month"`
	Year  float64 `json:"year"`
}

type ObjectMetadataReportOptionsState struct {
	MetadataFields            []string                                                    `json:"metadata_fields"`
	StorageDestinationOptions []ObjectMetadataReportOptionsStorageDestinationOptionsState `json:"storage_destination_options"`
	StorageFilters            []ObjectMetadataReportOptionsStorageFiltersState            `json:"storage_filters"`
}

type ObjectMetadataReportOptionsStorageDestinationOptionsState struct {
	Bucket          string `json:"bucket"`
	DestinationPath string `json:"destination_path"`
}

type ObjectMetadataReportOptionsStorageFiltersState struct {
	Bucket string `json:"bucket"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
