// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_dms_s3_endpoint

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_dms_s3_endpoint.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDmsS3EndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adse *Resource) Type() string {
	return "aws_dms_s3_endpoint"
}

// LocalName returns the local name for [Resource].
func (adse *Resource) LocalName() string {
	return adse.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adse *Resource) Configuration() interface{} {
	return adse.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adse *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adse)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adse *Resource) Dependencies() terra.Dependencies {
	return adse.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adse *Resource) LifecycleManagement() *terra.Lifecycle {
	return adse.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adse *Resource) Attributes() awsDmsS3EndpointAttributes {
	return awsDmsS3EndpointAttributes{ref: terra.ReferenceResource(adse)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adse *Resource) ImportState(state io.Reader) error {
	adse.state = &awsDmsS3EndpointState{}
	if err := json.NewDecoder(state).Decode(adse.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adse.Type(), adse.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adse *Resource) State() (*awsDmsS3EndpointState, bool) {
	return adse.state, adse.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adse *Resource) StateMust() *awsDmsS3EndpointState {
	if adse.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adse.Type(), adse.LocalName()))
	}
	return adse.state
}

// Args contains the configurations for aws_dms_s3_endpoint.
type Args struct {
	// AddColumnName: bool, optional
	AddColumnName terra.BoolValue `hcl:"add_column_name,attr"`
	// AddTrailingPaddingCharacter: bool, optional
	AddTrailingPaddingCharacter terra.BoolValue `hcl:"add_trailing_padding_character,attr"`
	// BucketFolder: string, optional
	BucketFolder terra.StringValue `hcl:"bucket_folder,attr"`
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// CannedAclForObjects: string, optional
	CannedAclForObjects terra.StringValue `hcl:"canned_acl_for_objects,attr"`
	// CdcInsertsAndUpdates: bool, optional
	CdcInsertsAndUpdates terra.BoolValue `hcl:"cdc_inserts_and_updates,attr"`
	// CdcInsertsOnly: bool, optional
	CdcInsertsOnly terra.BoolValue `hcl:"cdc_inserts_only,attr"`
	// CdcMaxBatchInterval: number, optional
	CdcMaxBatchInterval terra.NumberValue `hcl:"cdc_max_batch_interval,attr"`
	// CdcMinFileSize: number, optional
	CdcMinFileSize terra.NumberValue `hcl:"cdc_min_file_size,attr"`
	// CdcPath: string, optional
	CdcPath terra.StringValue `hcl:"cdc_path,attr"`
	// CertificateArn: string, optional
	CertificateArn terra.StringValue `hcl:"certificate_arn,attr"`
	// CompressionType: string, optional
	CompressionType terra.StringValue `hcl:"compression_type,attr"`
	// CsvDelimiter: string, optional
	CsvDelimiter terra.StringValue `hcl:"csv_delimiter,attr"`
	// CsvNoSupValue: string, optional
	CsvNoSupValue terra.StringValue `hcl:"csv_no_sup_value,attr"`
	// CsvNullValue: string, optional
	CsvNullValue terra.StringValue `hcl:"csv_null_value,attr"`
	// CsvRowDelimiter: string, optional
	CsvRowDelimiter terra.StringValue `hcl:"csv_row_delimiter,attr"`
	// DataFormat: string, optional
	DataFormat terra.StringValue `hcl:"data_format,attr"`
	// DataPageSize: number, optional
	DataPageSize terra.NumberValue `hcl:"data_page_size,attr"`
	// DatePartitionDelimiter: string, optional
	DatePartitionDelimiter terra.StringValue `hcl:"date_partition_delimiter,attr"`
	// DatePartitionEnabled: bool, optional
	DatePartitionEnabled terra.BoolValue `hcl:"date_partition_enabled,attr"`
	// DatePartitionSequence: string, optional
	DatePartitionSequence terra.StringValue `hcl:"date_partition_sequence,attr"`
	// DatePartitionTimezone: string, optional
	DatePartitionTimezone terra.StringValue `hcl:"date_partition_timezone,attr"`
	// DetachTargetOnLobLookupFailureParquet: bool, optional
	DetachTargetOnLobLookupFailureParquet terra.BoolValue `hcl:"detach_target_on_lob_lookup_failure_parquet,attr"`
	// DictPageSizeLimit: number, optional
	DictPageSizeLimit terra.NumberValue `hcl:"dict_page_size_limit,attr"`
	// EnableStatistics: bool, optional
	EnableStatistics terra.BoolValue `hcl:"enable_statistics,attr"`
	// EncodingType: string, optional
	EncodingType terra.StringValue `hcl:"encoding_type,attr"`
	// EncryptionMode: string, optional
	EncryptionMode terra.StringValue `hcl:"encryption_mode,attr"`
	// EndpointId: string, required
	EndpointId terra.StringValue `hcl:"endpoint_id,attr" validate:"required"`
	// EndpointType: string, required
	EndpointType terra.StringValue `hcl:"endpoint_type,attr" validate:"required"`
	// ExpectedBucketOwner: string, optional
	ExpectedBucketOwner terra.StringValue `hcl:"expected_bucket_owner,attr"`
	// ExternalTableDefinition: string, optional
	ExternalTableDefinition terra.StringValue `hcl:"external_table_definition,attr"`
	// GlueCatalogGeneration: bool, optional
	GlueCatalogGeneration terra.BoolValue `hcl:"glue_catalog_generation,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreHeaderRows: number, optional
	IgnoreHeaderRows terra.NumberValue `hcl:"ignore_header_rows,attr"`
	// IncludeOpForFullLoad: bool, optional
	IncludeOpForFullLoad terra.BoolValue `hcl:"include_op_for_full_load,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// MaxFileSize: number, optional
	MaxFileSize terra.NumberValue `hcl:"max_file_size,attr"`
	// ParquetTimestampInMillisecond: bool, optional
	ParquetTimestampInMillisecond terra.BoolValue `hcl:"parquet_timestamp_in_millisecond,attr"`
	// ParquetVersion: string, optional
	ParquetVersion terra.StringValue `hcl:"parquet_version,attr"`
	// PreserveTransactions: bool, optional
	PreserveTransactions terra.BoolValue `hcl:"preserve_transactions,attr"`
	// Rfc4180: bool, optional
	Rfc4180 terra.BoolValue `hcl:"rfc_4180,attr"`
	// RowGroupLength: number, optional
	RowGroupLength terra.NumberValue `hcl:"row_group_length,attr"`
	// ServerSideEncryptionKmsKeyId: string, optional
	ServerSideEncryptionKmsKeyId terra.StringValue `hcl:"server_side_encryption_kms_key_id,attr"`
	// ServiceAccessRoleArn: string, required
	ServiceAccessRoleArn terra.StringValue `hcl:"service_access_role_arn,attr" validate:"required"`
	// SslMode: string, optional
	SslMode terra.StringValue `hcl:"ssl_mode,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TimestampColumnName: string, optional
	TimestampColumnName terra.StringValue `hcl:"timestamp_column_name,attr"`
	// UseCsvNoSupValue: bool, optional
	UseCsvNoSupValue terra.BoolValue `hcl:"use_csv_no_sup_value,attr"`
	// UseTaskStartTimeForFullLoadTimestamp: bool, optional
	UseTaskStartTimeForFullLoadTimestamp terra.BoolValue `hcl:"use_task_start_time_for_full_load_timestamp,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsDmsS3EndpointAttributes struct {
	ref terra.Reference
}

// AddColumnName returns a reference to field add_column_name of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) AddColumnName() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("add_column_name"))
}

// AddTrailingPaddingCharacter returns a reference to field add_trailing_padding_character of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) AddTrailingPaddingCharacter() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("add_trailing_padding_character"))
}

// BucketFolder returns a reference to field bucket_folder of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) BucketFolder() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("bucket_folder"))
}

// BucketName returns a reference to field bucket_name of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("bucket_name"))
}

// CannedAclForObjects returns a reference to field canned_acl_for_objects of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CannedAclForObjects() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("canned_acl_for_objects"))
}

// CdcInsertsAndUpdates returns a reference to field cdc_inserts_and_updates of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CdcInsertsAndUpdates() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("cdc_inserts_and_updates"))
}

// CdcInsertsOnly returns a reference to field cdc_inserts_only of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CdcInsertsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("cdc_inserts_only"))
}

// CdcMaxBatchInterval returns a reference to field cdc_max_batch_interval of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CdcMaxBatchInterval() terra.NumberValue {
	return terra.ReferenceAsNumber(adse.ref.Append("cdc_max_batch_interval"))
}

// CdcMinFileSize returns a reference to field cdc_min_file_size of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CdcMinFileSize() terra.NumberValue {
	return terra.ReferenceAsNumber(adse.ref.Append("cdc_min_file_size"))
}

// CdcPath returns a reference to field cdc_path of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CdcPath() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("cdc_path"))
}

// CertificateArn returns a reference to field certificate_arn of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("certificate_arn"))
}

// CompressionType returns a reference to field compression_type of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CompressionType() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("compression_type"))
}

// CsvDelimiter returns a reference to field csv_delimiter of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CsvDelimiter() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("csv_delimiter"))
}

// CsvNoSupValue returns a reference to field csv_no_sup_value of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CsvNoSupValue() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("csv_no_sup_value"))
}

// CsvNullValue returns a reference to field csv_null_value of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CsvNullValue() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("csv_null_value"))
}

// CsvRowDelimiter returns a reference to field csv_row_delimiter of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) CsvRowDelimiter() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("csv_row_delimiter"))
}

// DataFormat returns a reference to field data_format of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) DataFormat() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("data_format"))
}

// DataPageSize returns a reference to field data_page_size of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) DataPageSize() terra.NumberValue {
	return terra.ReferenceAsNumber(adse.ref.Append("data_page_size"))
}

// DatePartitionDelimiter returns a reference to field date_partition_delimiter of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) DatePartitionDelimiter() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("date_partition_delimiter"))
}

// DatePartitionEnabled returns a reference to field date_partition_enabled of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) DatePartitionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("date_partition_enabled"))
}

// DatePartitionSequence returns a reference to field date_partition_sequence of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) DatePartitionSequence() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("date_partition_sequence"))
}

// DatePartitionTimezone returns a reference to field date_partition_timezone of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) DatePartitionTimezone() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("date_partition_timezone"))
}

// DetachTargetOnLobLookupFailureParquet returns a reference to field detach_target_on_lob_lookup_failure_parquet of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) DetachTargetOnLobLookupFailureParquet() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("detach_target_on_lob_lookup_failure_parquet"))
}

// DictPageSizeLimit returns a reference to field dict_page_size_limit of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) DictPageSizeLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(adse.ref.Append("dict_page_size_limit"))
}

// EnableStatistics returns a reference to field enable_statistics of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) EnableStatistics() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("enable_statistics"))
}

// EncodingType returns a reference to field encoding_type of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) EncodingType() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("encoding_type"))
}

// EncryptionMode returns a reference to field encryption_mode of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) EncryptionMode() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("encryption_mode"))
}

// EndpointArn returns a reference to field endpoint_arn of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) EndpointArn() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("endpoint_arn"))
}

// EndpointId returns a reference to field endpoint_id of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) EndpointId() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("endpoint_id"))
}

// EndpointType returns a reference to field endpoint_type of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("endpoint_type"))
}

// EngineDisplayName returns a reference to field engine_display_name of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) EngineDisplayName() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("engine_display_name"))
}

// ExpectedBucketOwner returns a reference to field expected_bucket_owner of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) ExpectedBucketOwner() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("expected_bucket_owner"))
}

// ExternalId returns a reference to field external_id of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) ExternalId() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("external_id"))
}

// ExternalTableDefinition returns a reference to field external_table_definition of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) ExternalTableDefinition() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("external_table_definition"))
}

// GlueCatalogGeneration returns a reference to field glue_catalog_generation of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) GlueCatalogGeneration() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("glue_catalog_generation"))
}

// Id returns a reference to field id of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("id"))
}

// IgnoreHeaderRows returns a reference to field ignore_header_rows of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) IgnoreHeaderRows() terra.NumberValue {
	return terra.ReferenceAsNumber(adse.ref.Append("ignore_header_rows"))
}

// IncludeOpForFullLoad returns a reference to field include_op_for_full_load of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) IncludeOpForFullLoad() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("include_op_for_full_load"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("kms_key_arn"))
}

// MaxFileSize returns a reference to field max_file_size of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) MaxFileSize() terra.NumberValue {
	return terra.ReferenceAsNumber(adse.ref.Append("max_file_size"))
}

// ParquetTimestampInMillisecond returns a reference to field parquet_timestamp_in_millisecond of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) ParquetTimestampInMillisecond() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("parquet_timestamp_in_millisecond"))
}

// ParquetVersion returns a reference to field parquet_version of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) ParquetVersion() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("parquet_version"))
}

// PreserveTransactions returns a reference to field preserve_transactions of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) PreserveTransactions() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("preserve_transactions"))
}

// Rfc4180 returns a reference to field rfc_4180 of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) Rfc4180() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("rfc_4180"))
}

// RowGroupLength returns a reference to field row_group_length of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) RowGroupLength() terra.NumberValue {
	return terra.ReferenceAsNumber(adse.ref.Append("row_group_length"))
}

// ServerSideEncryptionKmsKeyId returns a reference to field server_side_encryption_kms_key_id of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) ServerSideEncryptionKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("server_side_encryption_kms_key_id"))
}

// ServiceAccessRoleArn returns a reference to field service_access_role_arn of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) ServiceAccessRoleArn() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("service_access_role_arn"))
}

// SslMode returns a reference to field ssl_mode of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) SslMode() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("ssl_mode"))
}

// Status returns a reference to field status of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adse.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adse.ref.Append("tags_all"))
}

// TimestampColumnName returns a reference to field timestamp_column_name of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) TimestampColumnName() terra.StringValue {
	return terra.ReferenceAsString(adse.ref.Append("timestamp_column_name"))
}

// UseCsvNoSupValue returns a reference to field use_csv_no_sup_value of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) UseCsvNoSupValue() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("use_csv_no_sup_value"))
}

// UseTaskStartTimeForFullLoadTimestamp returns a reference to field use_task_start_time_for_full_load_timestamp of aws_dms_s3_endpoint.
func (adse awsDmsS3EndpointAttributes) UseTaskStartTimeForFullLoadTimestamp() terra.BoolValue {
	return terra.ReferenceAsBool(adse.ref.Append("use_task_start_time_for_full_load_timestamp"))
}

func (adse awsDmsS3EndpointAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adse.ref.Append("timeouts"))
}

type awsDmsS3EndpointState struct {
	AddColumnName                         bool              `json:"add_column_name"`
	AddTrailingPaddingCharacter           bool              `json:"add_trailing_padding_character"`
	BucketFolder                          string            `json:"bucket_folder"`
	BucketName                            string            `json:"bucket_name"`
	CannedAclForObjects                   string            `json:"canned_acl_for_objects"`
	CdcInsertsAndUpdates                  bool              `json:"cdc_inserts_and_updates"`
	CdcInsertsOnly                        bool              `json:"cdc_inserts_only"`
	CdcMaxBatchInterval                   float64           `json:"cdc_max_batch_interval"`
	CdcMinFileSize                        float64           `json:"cdc_min_file_size"`
	CdcPath                               string            `json:"cdc_path"`
	CertificateArn                        string            `json:"certificate_arn"`
	CompressionType                       string            `json:"compression_type"`
	CsvDelimiter                          string            `json:"csv_delimiter"`
	CsvNoSupValue                         string            `json:"csv_no_sup_value"`
	CsvNullValue                          string            `json:"csv_null_value"`
	CsvRowDelimiter                       string            `json:"csv_row_delimiter"`
	DataFormat                            string            `json:"data_format"`
	DataPageSize                          float64           `json:"data_page_size"`
	DatePartitionDelimiter                string            `json:"date_partition_delimiter"`
	DatePartitionEnabled                  bool              `json:"date_partition_enabled"`
	DatePartitionSequence                 string            `json:"date_partition_sequence"`
	DatePartitionTimezone                 string            `json:"date_partition_timezone"`
	DetachTargetOnLobLookupFailureParquet bool              `json:"detach_target_on_lob_lookup_failure_parquet"`
	DictPageSizeLimit                     float64           `json:"dict_page_size_limit"`
	EnableStatistics                      bool              `json:"enable_statistics"`
	EncodingType                          string            `json:"encoding_type"`
	EncryptionMode                        string            `json:"encryption_mode"`
	EndpointArn                           string            `json:"endpoint_arn"`
	EndpointId                            string            `json:"endpoint_id"`
	EndpointType                          string            `json:"endpoint_type"`
	EngineDisplayName                     string            `json:"engine_display_name"`
	ExpectedBucketOwner                   string            `json:"expected_bucket_owner"`
	ExternalId                            string            `json:"external_id"`
	ExternalTableDefinition               string            `json:"external_table_definition"`
	GlueCatalogGeneration                 bool              `json:"glue_catalog_generation"`
	Id                                    string            `json:"id"`
	IgnoreHeaderRows                      float64           `json:"ignore_header_rows"`
	IncludeOpForFullLoad                  bool              `json:"include_op_for_full_load"`
	KmsKeyArn                             string            `json:"kms_key_arn"`
	MaxFileSize                           float64           `json:"max_file_size"`
	ParquetTimestampInMillisecond         bool              `json:"parquet_timestamp_in_millisecond"`
	ParquetVersion                        string            `json:"parquet_version"`
	PreserveTransactions                  bool              `json:"preserve_transactions"`
	Rfc4180                               bool              `json:"rfc_4180"`
	RowGroupLength                        float64           `json:"row_group_length"`
	ServerSideEncryptionKmsKeyId          string            `json:"server_side_encryption_kms_key_id"`
	ServiceAccessRoleArn                  string            `json:"service_access_role_arn"`
	SslMode                               string            `json:"ssl_mode"`
	Status                                string            `json:"status"`
	Tags                                  map[string]string `json:"tags"`
	TagsAll                               map[string]string `json:"tags_all"`
	TimestampColumnName                   string            `json:"timestamp_column_name"`
	UseCsvNoSupValue                      bool              `json:"use_csv_no_sup_value"`
	UseTaskStartTimeForFullLoadTimestamp  bool              `json:"use_task_start_time_for_full_load_timestamp"`
	Timeouts                              *TimeoutsState    `json:"timeouts"`
}
