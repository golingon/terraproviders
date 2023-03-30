// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dmsendpoint

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ElasticsearchSettings struct {
	// EndpointUri: string, required
	EndpointUri terra.StringValue `hcl:"endpoint_uri,attr" validate:"required"`
	// ErrorRetryDuration: number, optional
	ErrorRetryDuration terra.NumberValue `hcl:"error_retry_duration,attr"`
	// FullLoadErrorPercentage: number, optional
	FullLoadErrorPercentage terra.NumberValue `hcl:"full_load_error_percentage,attr"`
	// ServiceAccessRoleArn: string, required
	ServiceAccessRoleArn terra.StringValue `hcl:"service_access_role_arn,attr" validate:"required"`
}

type KafkaSettings struct {
	// Broker: string, required
	Broker terra.StringValue `hcl:"broker,attr" validate:"required"`
	// IncludeControlDetails: bool, optional
	IncludeControlDetails terra.BoolValue `hcl:"include_control_details,attr"`
	// IncludeNullAndEmpty: bool, optional
	IncludeNullAndEmpty terra.BoolValue `hcl:"include_null_and_empty,attr"`
	// IncludePartitionValue: bool, optional
	IncludePartitionValue terra.BoolValue `hcl:"include_partition_value,attr"`
	// IncludeTableAlterOperations: bool, optional
	IncludeTableAlterOperations terra.BoolValue `hcl:"include_table_alter_operations,attr"`
	// IncludeTransactionDetails: bool, optional
	IncludeTransactionDetails terra.BoolValue `hcl:"include_transaction_details,attr"`
	// MessageFormat: string, optional
	MessageFormat terra.StringValue `hcl:"message_format,attr"`
	// MessageMaxBytes: number, optional
	MessageMaxBytes terra.NumberValue `hcl:"message_max_bytes,attr"`
	// NoHexPrefix: bool, optional
	NoHexPrefix terra.BoolValue `hcl:"no_hex_prefix,attr"`
	// PartitionIncludeSchemaTable: bool, optional
	PartitionIncludeSchemaTable terra.BoolValue `hcl:"partition_include_schema_table,attr"`
	// SaslPassword: string, optional
	SaslPassword terra.StringValue `hcl:"sasl_password,attr"`
	// SaslUsername: string, optional
	SaslUsername terra.StringValue `hcl:"sasl_username,attr"`
	// SecurityProtocol: string, optional
	SecurityProtocol terra.StringValue `hcl:"security_protocol,attr"`
	// SslCaCertificateArn: string, optional
	SslCaCertificateArn terra.StringValue `hcl:"ssl_ca_certificate_arn,attr"`
	// SslClientCertificateArn: string, optional
	SslClientCertificateArn terra.StringValue `hcl:"ssl_client_certificate_arn,attr"`
	// SslClientKeyArn: string, optional
	SslClientKeyArn terra.StringValue `hcl:"ssl_client_key_arn,attr"`
	// SslClientKeyPassword: string, optional
	SslClientKeyPassword terra.StringValue `hcl:"ssl_client_key_password,attr"`
	// Topic: string, optional
	Topic terra.StringValue `hcl:"topic,attr"`
}

type KinesisSettings struct {
	// IncludeControlDetails: bool, optional
	IncludeControlDetails terra.BoolValue `hcl:"include_control_details,attr"`
	// IncludeNullAndEmpty: bool, optional
	IncludeNullAndEmpty terra.BoolValue `hcl:"include_null_and_empty,attr"`
	// IncludePartitionValue: bool, optional
	IncludePartitionValue terra.BoolValue `hcl:"include_partition_value,attr"`
	// IncludeTableAlterOperations: bool, optional
	IncludeTableAlterOperations terra.BoolValue `hcl:"include_table_alter_operations,attr"`
	// IncludeTransactionDetails: bool, optional
	IncludeTransactionDetails terra.BoolValue `hcl:"include_transaction_details,attr"`
	// MessageFormat: string, optional
	MessageFormat terra.StringValue `hcl:"message_format,attr"`
	// PartitionIncludeSchemaTable: bool, optional
	PartitionIncludeSchemaTable terra.BoolValue `hcl:"partition_include_schema_table,attr"`
	// ServiceAccessRoleArn: string, optional
	ServiceAccessRoleArn terra.StringValue `hcl:"service_access_role_arn,attr"`
	// StreamArn: string, optional
	StreamArn terra.StringValue `hcl:"stream_arn,attr"`
}

type MongodbSettings struct {
	// AuthMechanism: string, optional
	AuthMechanism terra.StringValue `hcl:"auth_mechanism,attr"`
	// AuthSource: string, optional
	AuthSource terra.StringValue `hcl:"auth_source,attr"`
	// AuthType: string, optional
	AuthType terra.StringValue `hcl:"auth_type,attr"`
	// DocsToInvestigate: string, optional
	DocsToInvestigate terra.StringValue `hcl:"docs_to_investigate,attr"`
	// ExtractDocId: string, optional
	ExtractDocId terra.StringValue `hcl:"extract_doc_id,attr"`
	// NestingLevel: string, optional
	NestingLevel terra.StringValue `hcl:"nesting_level,attr"`
}

type RedisSettings struct {
	// AuthPassword: string, optional
	AuthPassword terra.StringValue `hcl:"auth_password,attr"`
	// AuthType: string, required
	AuthType terra.StringValue `hcl:"auth_type,attr" validate:"required"`
	// AuthUserName: string, optional
	AuthUserName terra.StringValue `hcl:"auth_user_name,attr"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// SslCaCertificateArn: string, optional
	SslCaCertificateArn terra.StringValue `hcl:"ssl_ca_certificate_arn,attr"`
	// SslSecurityProtocol: string, optional
	SslSecurityProtocol terra.StringValue `hcl:"ssl_security_protocol,attr"`
}

type RedshiftSettings struct {
	// BucketFolder: string, optional
	BucketFolder terra.StringValue `hcl:"bucket_folder,attr"`
	// BucketName: string, optional
	BucketName terra.StringValue `hcl:"bucket_name,attr"`
	// EncryptionMode: string, optional
	EncryptionMode terra.StringValue `hcl:"encryption_mode,attr"`
	// ServerSideEncryptionKmsKeyId: string, optional
	ServerSideEncryptionKmsKeyId terra.StringValue `hcl:"server_side_encryption_kms_key_id,attr"`
	// ServiceAccessRoleArn: string, optional
	ServiceAccessRoleArn terra.StringValue `hcl:"service_access_role_arn,attr"`
}

type S3Settings struct {
	// AddColumnName: bool, optional
	AddColumnName terra.BoolValue `hcl:"add_column_name,attr"`
	// BucketFolder: string, optional
	BucketFolder terra.StringValue `hcl:"bucket_folder,attr"`
	// BucketName: string, optional
	BucketName terra.StringValue `hcl:"bucket_name,attr"`
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
	// DictPageSizeLimit: number, optional
	DictPageSizeLimit terra.NumberValue `hcl:"dict_page_size_limit,attr"`
	// EnableStatistics: bool, optional
	EnableStatistics terra.BoolValue `hcl:"enable_statistics,attr"`
	// EncodingType: string, optional
	EncodingType terra.StringValue `hcl:"encoding_type,attr"`
	// EncryptionMode: string, optional
	EncryptionMode terra.StringValue `hcl:"encryption_mode,attr"`
	// ExternalTableDefinition: string, optional
	ExternalTableDefinition terra.StringValue `hcl:"external_table_definition,attr"`
	// IgnoreHeaderRows: number, optional
	IgnoreHeaderRows terra.NumberValue `hcl:"ignore_header_rows,attr"`
	// IgnoreHeadersRow: number, optional
	IgnoreHeadersRow terra.NumberValue `hcl:"ignore_headers_row,attr"`
	// IncludeOpForFullLoad: bool, optional
	IncludeOpForFullLoad terra.BoolValue `hcl:"include_op_for_full_load,attr"`
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
	// ServiceAccessRoleArn: string, optional
	ServiceAccessRoleArn terra.StringValue `hcl:"service_access_role_arn,attr"`
	// TimestampColumnName: string, optional
	TimestampColumnName terra.StringValue `hcl:"timestamp_column_name,attr"`
	// UseCsvNoSupValue: bool, optional
	UseCsvNoSupValue terra.BoolValue `hcl:"use_csv_no_sup_value,attr"`
	// UseTaskStartTimeForFullLoadTimestamp: bool, optional
	UseTaskStartTimeForFullLoadTimestamp terra.BoolValue `hcl:"use_task_start_time_for_full_load_timestamp,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type ElasticsearchSettingsAttributes struct {
	ref terra.Reference
}

func (es ElasticsearchSettingsAttributes) InternalRef() terra.Reference {
	return es.ref
}

func (es ElasticsearchSettingsAttributes) InternalWithRef(ref terra.Reference) ElasticsearchSettingsAttributes {
	return ElasticsearchSettingsAttributes{ref: ref}
}

func (es ElasticsearchSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return es.ref.InternalTokens()
}

func (es ElasticsearchSettingsAttributes) EndpointUri() terra.StringValue {
	return terra.ReferenceString(es.ref.Append("endpoint_uri"))
}

func (es ElasticsearchSettingsAttributes) ErrorRetryDuration() terra.NumberValue {
	return terra.ReferenceNumber(es.ref.Append("error_retry_duration"))
}

func (es ElasticsearchSettingsAttributes) FullLoadErrorPercentage() terra.NumberValue {
	return terra.ReferenceNumber(es.ref.Append("full_load_error_percentage"))
}

func (es ElasticsearchSettingsAttributes) ServiceAccessRoleArn() terra.StringValue {
	return terra.ReferenceString(es.ref.Append("service_access_role_arn"))
}

type KafkaSettingsAttributes struct {
	ref terra.Reference
}

func (ks KafkaSettingsAttributes) InternalRef() terra.Reference {
	return ks.ref
}

func (ks KafkaSettingsAttributes) InternalWithRef(ref terra.Reference) KafkaSettingsAttributes {
	return KafkaSettingsAttributes{ref: ref}
}

func (ks KafkaSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ks.ref.InternalTokens()
}

func (ks KafkaSettingsAttributes) Broker() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("broker"))
}

func (ks KafkaSettingsAttributes) IncludeControlDetails() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("include_control_details"))
}

func (ks KafkaSettingsAttributes) IncludeNullAndEmpty() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("include_null_and_empty"))
}

func (ks KafkaSettingsAttributes) IncludePartitionValue() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("include_partition_value"))
}

func (ks KafkaSettingsAttributes) IncludeTableAlterOperations() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("include_table_alter_operations"))
}

func (ks KafkaSettingsAttributes) IncludeTransactionDetails() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("include_transaction_details"))
}

func (ks KafkaSettingsAttributes) MessageFormat() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("message_format"))
}

func (ks KafkaSettingsAttributes) MessageMaxBytes() terra.NumberValue {
	return terra.ReferenceNumber(ks.ref.Append("message_max_bytes"))
}

func (ks KafkaSettingsAttributes) NoHexPrefix() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("no_hex_prefix"))
}

func (ks KafkaSettingsAttributes) PartitionIncludeSchemaTable() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("partition_include_schema_table"))
}

func (ks KafkaSettingsAttributes) SaslPassword() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("sasl_password"))
}

func (ks KafkaSettingsAttributes) SaslUsername() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("sasl_username"))
}

func (ks KafkaSettingsAttributes) SecurityProtocol() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("security_protocol"))
}

func (ks KafkaSettingsAttributes) SslCaCertificateArn() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("ssl_ca_certificate_arn"))
}

func (ks KafkaSettingsAttributes) SslClientCertificateArn() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("ssl_client_certificate_arn"))
}

func (ks KafkaSettingsAttributes) SslClientKeyArn() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("ssl_client_key_arn"))
}

func (ks KafkaSettingsAttributes) SslClientKeyPassword() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("ssl_client_key_password"))
}

func (ks KafkaSettingsAttributes) Topic() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("topic"))
}

type KinesisSettingsAttributes struct {
	ref terra.Reference
}

func (ks KinesisSettingsAttributes) InternalRef() terra.Reference {
	return ks.ref
}

func (ks KinesisSettingsAttributes) InternalWithRef(ref terra.Reference) KinesisSettingsAttributes {
	return KinesisSettingsAttributes{ref: ref}
}

func (ks KinesisSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ks.ref.InternalTokens()
}

func (ks KinesisSettingsAttributes) IncludeControlDetails() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("include_control_details"))
}

func (ks KinesisSettingsAttributes) IncludeNullAndEmpty() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("include_null_and_empty"))
}

func (ks KinesisSettingsAttributes) IncludePartitionValue() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("include_partition_value"))
}

func (ks KinesisSettingsAttributes) IncludeTableAlterOperations() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("include_table_alter_operations"))
}

func (ks KinesisSettingsAttributes) IncludeTransactionDetails() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("include_transaction_details"))
}

func (ks KinesisSettingsAttributes) MessageFormat() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("message_format"))
}

func (ks KinesisSettingsAttributes) PartitionIncludeSchemaTable() terra.BoolValue {
	return terra.ReferenceBool(ks.ref.Append("partition_include_schema_table"))
}

func (ks KinesisSettingsAttributes) ServiceAccessRoleArn() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("service_access_role_arn"))
}

func (ks KinesisSettingsAttributes) StreamArn() terra.StringValue {
	return terra.ReferenceString(ks.ref.Append("stream_arn"))
}

type MongodbSettingsAttributes struct {
	ref terra.Reference
}

func (ms MongodbSettingsAttributes) InternalRef() terra.Reference {
	return ms.ref
}

func (ms MongodbSettingsAttributes) InternalWithRef(ref terra.Reference) MongodbSettingsAttributes {
	return MongodbSettingsAttributes{ref: ref}
}

func (ms MongodbSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ms.ref.InternalTokens()
}

func (ms MongodbSettingsAttributes) AuthMechanism() terra.StringValue {
	return terra.ReferenceString(ms.ref.Append("auth_mechanism"))
}

func (ms MongodbSettingsAttributes) AuthSource() terra.StringValue {
	return terra.ReferenceString(ms.ref.Append("auth_source"))
}

func (ms MongodbSettingsAttributes) AuthType() terra.StringValue {
	return terra.ReferenceString(ms.ref.Append("auth_type"))
}

func (ms MongodbSettingsAttributes) DocsToInvestigate() terra.StringValue {
	return terra.ReferenceString(ms.ref.Append("docs_to_investigate"))
}

func (ms MongodbSettingsAttributes) ExtractDocId() terra.StringValue {
	return terra.ReferenceString(ms.ref.Append("extract_doc_id"))
}

func (ms MongodbSettingsAttributes) NestingLevel() terra.StringValue {
	return terra.ReferenceString(ms.ref.Append("nesting_level"))
}

type RedisSettingsAttributes struct {
	ref terra.Reference
}

func (rs RedisSettingsAttributes) InternalRef() terra.Reference {
	return rs.ref
}

func (rs RedisSettingsAttributes) InternalWithRef(ref terra.Reference) RedisSettingsAttributes {
	return RedisSettingsAttributes{ref: ref}
}

func (rs RedisSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return rs.ref.InternalTokens()
}

func (rs RedisSettingsAttributes) AuthPassword() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("auth_password"))
}

func (rs RedisSettingsAttributes) AuthType() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("auth_type"))
}

func (rs RedisSettingsAttributes) AuthUserName() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("auth_user_name"))
}

func (rs RedisSettingsAttributes) Port() terra.NumberValue {
	return terra.ReferenceNumber(rs.ref.Append("port"))
}

func (rs RedisSettingsAttributes) ServerName() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("server_name"))
}

func (rs RedisSettingsAttributes) SslCaCertificateArn() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("ssl_ca_certificate_arn"))
}

func (rs RedisSettingsAttributes) SslSecurityProtocol() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("ssl_security_protocol"))
}

type RedshiftSettingsAttributes struct {
	ref terra.Reference
}

func (rs RedshiftSettingsAttributes) InternalRef() terra.Reference {
	return rs.ref
}

func (rs RedshiftSettingsAttributes) InternalWithRef(ref terra.Reference) RedshiftSettingsAttributes {
	return RedshiftSettingsAttributes{ref: ref}
}

func (rs RedshiftSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return rs.ref.InternalTokens()
}

func (rs RedshiftSettingsAttributes) BucketFolder() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("bucket_folder"))
}

func (rs RedshiftSettingsAttributes) BucketName() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("bucket_name"))
}

func (rs RedshiftSettingsAttributes) EncryptionMode() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("encryption_mode"))
}

func (rs RedshiftSettingsAttributes) ServerSideEncryptionKmsKeyId() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("server_side_encryption_kms_key_id"))
}

func (rs RedshiftSettingsAttributes) ServiceAccessRoleArn() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("service_access_role_arn"))
}

type S3SettingsAttributes struct {
	ref terra.Reference
}

func (ss S3SettingsAttributes) InternalRef() terra.Reference {
	return ss.ref
}

func (ss S3SettingsAttributes) InternalWithRef(ref terra.Reference) S3SettingsAttributes {
	return S3SettingsAttributes{ref: ref}
}

func (ss S3SettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ss.ref.InternalTokens()
}

func (ss S3SettingsAttributes) AddColumnName() terra.BoolValue {
	return terra.ReferenceBool(ss.ref.Append("add_column_name"))
}

func (ss S3SettingsAttributes) BucketFolder() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("bucket_folder"))
}

func (ss S3SettingsAttributes) BucketName() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("bucket_name"))
}

func (ss S3SettingsAttributes) CannedAclForObjects() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("canned_acl_for_objects"))
}

func (ss S3SettingsAttributes) CdcInsertsAndUpdates() terra.BoolValue {
	return terra.ReferenceBool(ss.ref.Append("cdc_inserts_and_updates"))
}

func (ss S3SettingsAttributes) CdcInsertsOnly() terra.BoolValue {
	return terra.ReferenceBool(ss.ref.Append("cdc_inserts_only"))
}

func (ss S3SettingsAttributes) CdcMaxBatchInterval() terra.NumberValue {
	return terra.ReferenceNumber(ss.ref.Append("cdc_max_batch_interval"))
}

func (ss S3SettingsAttributes) CdcMinFileSize() terra.NumberValue {
	return terra.ReferenceNumber(ss.ref.Append("cdc_min_file_size"))
}

func (ss S3SettingsAttributes) CdcPath() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("cdc_path"))
}

func (ss S3SettingsAttributes) CompressionType() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("compression_type"))
}

func (ss S3SettingsAttributes) CsvDelimiter() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("csv_delimiter"))
}

func (ss S3SettingsAttributes) CsvNoSupValue() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("csv_no_sup_value"))
}

func (ss S3SettingsAttributes) CsvNullValue() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("csv_null_value"))
}

func (ss S3SettingsAttributes) CsvRowDelimiter() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("csv_row_delimiter"))
}

func (ss S3SettingsAttributes) DataFormat() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("data_format"))
}

func (ss S3SettingsAttributes) DataPageSize() terra.NumberValue {
	return terra.ReferenceNumber(ss.ref.Append("data_page_size"))
}

func (ss S3SettingsAttributes) DatePartitionDelimiter() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("date_partition_delimiter"))
}

func (ss S3SettingsAttributes) DatePartitionEnabled() terra.BoolValue {
	return terra.ReferenceBool(ss.ref.Append("date_partition_enabled"))
}

func (ss S3SettingsAttributes) DatePartitionSequence() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("date_partition_sequence"))
}

func (ss S3SettingsAttributes) DictPageSizeLimit() terra.NumberValue {
	return terra.ReferenceNumber(ss.ref.Append("dict_page_size_limit"))
}

func (ss S3SettingsAttributes) EnableStatistics() terra.BoolValue {
	return terra.ReferenceBool(ss.ref.Append("enable_statistics"))
}

func (ss S3SettingsAttributes) EncodingType() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("encoding_type"))
}

func (ss S3SettingsAttributes) EncryptionMode() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("encryption_mode"))
}

func (ss S3SettingsAttributes) ExternalTableDefinition() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("external_table_definition"))
}

func (ss S3SettingsAttributes) IgnoreHeaderRows() terra.NumberValue {
	return terra.ReferenceNumber(ss.ref.Append("ignore_header_rows"))
}

func (ss S3SettingsAttributes) IgnoreHeadersRow() terra.NumberValue {
	return terra.ReferenceNumber(ss.ref.Append("ignore_headers_row"))
}

func (ss S3SettingsAttributes) IncludeOpForFullLoad() terra.BoolValue {
	return terra.ReferenceBool(ss.ref.Append("include_op_for_full_load"))
}

func (ss S3SettingsAttributes) MaxFileSize() terra.NumberValue {
	return terra.ReferenceNumber(ss.ref.Append("max_file_size"))
}

func (ss S3SettingsAttributes) ParquetTimestampInMillisecond() terra.BoolValue {
	return terra.ReferenceBool(ss.ref.Append("parquet_timestamp_in_millisecond"))
}

func (ss S3SettingsAttributes) ParquetVersion() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("parquet_version"))
}

func (ss S3SettingsAttributes) PreserveTransactions() terra.BoolValue {
	return terra.ReferenceBool(ss.ref.Append("preserve_transactions"))
}

func (ss S3SettingsAttributes) Rfc4180() terra.BoolValue {
	return terra.ReferenceBool(ss.ref.Append("rfc_4180"))
}

func (ss S3SettingsAttributes) RowGroupLength() terra.NumberValue {
	return terra.ReferenceNumber(ss.ref.Append("row_group_length"))
}

func (ss S3SettingsAttributes) ServerSideEncryptionKmsKeyId() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("server_side_encryption_kms_key_id"))
}

func (ss S3SettingsAttributes) ServiceAccessRoleArn() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("service_access_role_arn"))
}

func (ss S3SettingsAttributes) TimestampColumnName() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("timestamp_column_name"))
}

func (ss S3SettingsAttributes) UseCsvNoSupValue() terra.BoolValue {
	return terra.ReferenceBool(ss.ref.Append("use_csv_no_sup_value"))
}

func (ss S3SettingsAttributes) UseTaskStartTimeForFullLoadTimestamp() terra.BoolValue {
	return terra.ReferenceBool(ss.ref.Append("use_task_start_time_for_full_load_timestamp"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

type ElasticsearchSettingsState struct {
	EndpointUri             string  `json:"endpoint_uri"`
	ErrorRetryDuration      float64 `json:"error_retry_duration"`
	FullLoadErrorPercentage float64 `json:"full_load_error_percentage"`
	ServiceAccessRoleArn    string  `json:"service_access_role_arn"`
}

type KafkaSettingsState struct {
	Broker                      string  `json:"broker"`
	IncludeControlDetails       bool    `json:"include_control_details"`
	IncludeNullAndEmpty         bool    `json:"include_null_and_empty"`
	IncludePartitionValue       bool    `json:"include_partition_value"`
	IncludeTableAlterOperations bool    `json:"include_table_alter_operations"`
	IncludeTransactionDetails   bool    `json:"include_transaction_details"`
	MessageFormat               string  `json:"message_format"`
	MessageMaxBytes             float64 `json:"message_max_bytes"`
	NoHexPrefix                 bool    `json:"no_hex_prefix"`
	PartitionIncludeSchemaTable bool    `json:"partition_include_schema_table"`
	SaslPassword                string  `json:"sasl_password"`
	SaslUsername                string  `json:"sasl_username"`
	SecurityProtocol            string  `json:"security_protocol"`
	SslCaCertificateArn         string  `json:"ssl_ca_certificate_arn"`
	SslClientCertificateArn     string  `json:"ssl_client_certificate_arn"`
	SslClientKeyArn             string  `json:"ssl_client_key_arn"`
	SslClientKeyPassword        string  `json:"ssl_client_key_password"`
	Topic                       string  `json:"topic"`
}

type KinesisSettingsState struct {
	IncludeControlDetails       bool   `json:"include_control_details"`
	IncludeNullAndEmpty         bool   `json:"include_null_and_empty"`
	IncludePartitionValue       bool   `json:"include_partition_value"`
	IncludeTableAlterOperations bool   `json:"include_table_alter_operations"`
	IncludeTransactionDetails   bool   `json:"include_transaction_details"`
	MessageFormat               string `json:"message_format"`
	PartitionIncludeSchemaTable bool   `json:"partition_include_schema_table"`
	ServiceAccessRoleArn        string `json:"service_access_role_arn"`
	StreamArn                   string `json:"stream_arn"`
}

type MongodbSettingsState struct {
	AuthMechanism     string `json:"auth_mechanism"`
	AuthSource        string `json:"auth_source"`
	AuthType          string `json:"auth_type"`
	DocsToInvestigate string `json:"docs_to_investigate"`
	ExtractDocId      string `json:"extract_doc_id"`
	NestingLevel      string `json:"nesting_level"`
}

type RedisSettingsState struct {
	AuthPassword        string  `json:"auth_password"`
	AuthType            string  `json:"auth_type"`
	AuthUserName        string  `json:"auth_user_name"`
	Port                float64 `json:"port"`
	ServerName          string  `json:"server_name"`
	SslCaCertificateArn string  `json:"ssl_ca_certificate_arn"`
	SslSecurityProtocol string  `json:"ssl_security_protocol"`
}

type RedshiftSettingsState struct {
	BucketFolder                 string `json:"bucket_folder"`
	BucketName                   string `json:"bucket_name"`
	EncryptionMode               string `json:"encryption_mode"`
	ServerSideEncryptionKmsKeyId string `json:"server_side_encryption_kms_key_id"`
	ServiceAccessRoleArn         string `json:"service_access_role_arn"`
}

type S3SettingsState struct {
	AddColumnName                        bool    `json:"add_column_name"`
	BucketFolder                         string  `json:"bucket_folder"`
	BucketName                           string  `json:"bucket_name"`
	CannedAclForObjects                  string  `json:"canned_acl_for_objects"`
	CdcInsertsAndUpdates                 bool    `json:"cdc_inserts_and_updates"`
	CdcInsertsOnly                       bool    `json:"cdc_inserts_only"`
	CdcMaxBatchInterval                  float64 `json:"cdc_max_batch_interval"`
	CdcMinFileSize                       float64 `json:"cdc_min_file_size"`
	CdcPath                              string  `json:"cdc_path"`
	CompressionType                      string  `json:"compression_type"`
	CsvDelimiter                         string  `json:"csv_delimiter"`
	CsvNoSupValue                        string  `json:"csv_no_sup_value"`
	CsvNullValue                         string  `json:"csv_null_value"`
	CsvRowDelimiter                      string  `json:"csv_row_delimiter"`
	DataFormat                           string  `json:"data_format"`
	DataPageSize                         float64 `json:"data_page_size"`
	DatePartitionDelimiter               string  `json:"date_partition_delimiter"`
	DatePartitionEnabled                 bool    `json:"date_partition_enabled"`
	DatePartitionSequence                string  `json:"date_partition_sequence"`
	DictPageSizeLimit                    float64 `json:"dict_page_size_limit"`
	EnableStatistics                     bool    `json:"enable_statistics"`
	EncodingType                         string  `json:"encoding_type"`
	EncryptionMode                       string  `json:"encryption_mode"`
	ExternalTableDefinition              string  `json:"external_table_definition"`
	IgnoreHeaderRows                     float64 `json:"ignore_header_rows"`
	IgnoreHeadersRow                     float64 `json:"ignore_headers_row"`
	IncludeOpForFullLoad                 bool    `json:"include_op_for_full_load"`
	MaxFileSize                          float64 `json:"max_file_size"`
	ParquetTimestampInMillisecond        bool    `json:"parquet_timestamp_in_millisecond"`
	ParquetVersion                       string  `json:"parquet_version"`
	PreserveTransactions                 bool    `json:"preserve_transactions"`
	Rfc4180                              bool    `json:"rfc_4180"`
	RowGroupLength                       float64 `json:"row_group_length"`
	ServerSideEncryptionKmsKeyId         string  `json:"server_side_encryption_kms_key_id"`
	ServiceAccessRoleArn                 string  `json:"service_access_role_arn"`
	TimestampColumnName                  string  `json:"timestamp_column_name"`
	UseCsvNoSupValue                     bool    `json:"use_csv_no_sup_value"`
	UseTaskStartTimeForFullLoadTimestamp bool    `json:"use_task_start_time_for_full_load_timestamp"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
