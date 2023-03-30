// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package s3bucketanalyticsconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Filter struct {
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type StorageClassAnalysis struct {
	// DataExport: required
	DataExport *DataExport `hcl:"data_export,block" validate:"required"`
}

type DataExport struct {
	// OutputSchemaVersion: string, optional
	OutputSchemaVersion terra.StringValue `hcl:"output_schema_version,attr"`
	// Destination: required
	Destination *Destination `hcl:"destination,block" validate:"required"`
}

type Destination struct {
	// S3BucketDestination: required
	S3BucketDestination *S3BucketDestination `hcl:"s3_bucket_destination,block" validate:"required"`
}

type S3BucketDestination struct {
	// BucketAccountId: string, optional
	BucketAccountId terra.StringValue `hcl:"bucket_account_id,attr"`
	// BucketArn: string, required
	BucketArn terra.StringValue `hcl:"bucket_arn,attr" validate:"required"`
	// Format: string, optional
	Format terra.StringValue `hcl:"format,attr"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() terra.Reference {
	return f.ref
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Prefix() terra.StringValue {
	return terra.ReferenceString(f.ref.Append("prefix"))
}

func (f FilterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](f.ref.Append("tags"))
}

type StorageClassAnalysisAttributes struct {
	ref terra.Reference
}

func (sca StorageClassAnalysisAttributes) InternalRef() terra.Reference {
	return sca.ref
}

func (sca StorageClassAnalysisAttributes) InternalWithRef(ref terra.Reference) StorageClassAnalysisAttributes {
	return StorageClassAnalysisAttributes{ref: ref}
}

func (sca StorageClassAnalysisAttributes) InternalTokens() hclwrite.Tokens {
	return sca.ref.InternalTokens()
}

func (sca StorageClassAnalysisAttributes) DataExport() terra.ListValue[DataExportAttributes] {
	return terra.ReferenceList[DataExportAttributes](sca.ref.Append("data_export"))
}

type DataExportAttributes struct {
	ref terra.Reference
}

func (de DataExportAttributes) InternalRef() terra.Reference {
	return de.ref
}

func (de DataExportAttributes) InternalWithRef(ref terra.Reference) DataExportAttributes {
	return DataExportAttributes{ref: ref}
}

func (de DataExportAttributes) InternalTokens() hclwrite.Tokens {
	return de.ref.InternalTokens()
}

func (de DataExportAttributes) OutputSchemaVersion() terra.StringValue {
	return terra.ReferenceString(de.ref.Append("output_schema_version"))
}

func (de DataExportAttributes) Destination() terra.ListValue[DestinationAttributes] {
	return terra.ReferenceList[DestinationAttributes](de.ref.Append("destination"))
}

type DestinationAttributes struct {
	ref terra.Reference
}

func (d DestinationAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d DestinationAttributes) InternalWithRef(ref terra.Reference) DestinationAttributes {
	return DestinationAttributes{ref: ref}
}

func (d DestinationAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d DestinationAttributes) S3BucketDestination() terra.ListValue[S3BucketDestinationAttributes] {
	return terra.ReferenceList[S3BucketDestinationAttributes](d.ref.Append("s3_bucket_destination"))
}

type S3BucketDestinationAttributes struct {
	ref terra.Reference
}

func (sbd S3BucketDestinationAttributes) InternalRef() terra.Reference {
	return sbd.ref
}

func (sbd S3BucketDestinationAttributes) InternalWithRef(ref terra.Reference) S3BucketDestinationAttributes {
	return S3BucketDestinationAttributes{ref: ref}
}

func (sbd S3BucketDestinationAttributes) InternalTokens() hclwrite.Tokens {
	return sbd.ref.InternalTokens()
}

func (sbd S3BucketDestinationAttributes) BucketAccountId() terra.StringValue {
	return terra.ReferenceString(sbd.ref.Append("bucket_account_id"))
}

func (sbd S3BucketDestinationAttributes) BucketArn() terra.StringValue {
	return terra.ReferenceString(sbd.ref.Append("bucket_arn"))
}

func (sbd S3BucketDestinationAttributes) Format() terra.StringValue {
	return terra.ReferenceString(sbd.ref.Append("format"))
}

func (sbd S3BucketDestinationAttributes) Prefix() terra.StringValue {
	return terra.ReferenceString(sbd.ref.Append("prefix"))
}

type FilterState struct {
	Prefix string            `json:"prefix"`
	Tags   map[string]string `json:"tags"`
}

type StorageClassAnalysisState struct {
	DataExport []DataExportState `json:"data_export"`
}

type DataExportState struct {
	OutputSchemaVersion string             `json:"output_schema_version"`
	Destination         []DestinationState `json:"destination"`
}

type DestinationState struct {
	S3BucketDestination []S3BucketDestinationState `json:"s3_bucket_destination"`
}

type S3BucketDestinationState struct {
	BucketAccountId string `json:"bucket_account_id"`
	BucketArn       string `json:"bucket_arn"`
	Format          string `json:"format"`
	Prefix          string `json:"prefix"`
}
