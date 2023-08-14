// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package kinesisanalyticsapplication

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CloudwatchLoggingOptions struct {
	// LogStreamArn: string, required
	LogStreamArn terra.StringValue `hcl:"log_stream_arn,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type Inputs struct {
	// NamePrefix: string, required
	NamePrefix terra.StringValue `hcl:"name_prefix,attr" validate:"required"`
	// InputsKinesisFirehose: optional
	KinesisFirehose *InputsKinesisFirehose `hcl:"kinesis_firehose,block"`
	// InputsKinesisStream: optional
	KinesisStream *InputsKinesisStream `hcl:"kinesis_stream,block"`
	// Parallelism: optional
	Parallelism *Parallelism `hcl:"parallelism,block"`
	// ProcessingConfiguration: optional
	ProcessingConfiguration *ProcessingConfiguration `hcl:"processing_configuration,block"`
	// InputsSchema: required
	Schema *InputsSchema `hcl:"schema,block" validate:"required"`
	// StartingPositionConfiguration: min=0
	StartingPositionConfiguration []StartingPositionConfiguration `hcl:"starting_position_configuration,block" validate:"min=0"`
}

type InputsKinesisFirehose struct {
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type InputsKinesisStream struct {
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type Parallelism struct {
	// Count: number, optional
	Count terra.NumberValue `hcl:"count,attr"`
}

type ProcessingConfiguration struct {
	// ProcessingConfigurationLambda: required
	Lambda *ProcessingConfigurationLambda `hcl:"lambda,block" validate:"required"`
}

type ProcessingConfigurationLambda struct {
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type InputsSchema struct {
	// RecordEncoding: string, optional
	RecordEncoding terra.StringValue `hcl:"record_encoding,attr"`
	// InputsSchemaRecordColumns: min=1,max=1000
	RecordColumns []InputsSchemaRecordColumns `hcl:"record_columns,block" validate:"min=1,max=1000"`
	// InputsSchemaRecordFormat: required
	RecordFormat *InputsSchemaRecordFormat `hcl:"record_format,block" validate:"required"`
}

type InputsSchemaRecordColumns struct {
	// Mapping: string, optional
	Mapping terra.StringValue `hcl:"mapping,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SqlType: string, required
	SqlType terra.StringValue `hcl:"sql_type,attr" validate:"required"`
}

type InputsSchemaRecordFormat struct {
	// InputsSchemaRecordFormatMappingParameters: optional
	MappingParameters *InputsSchemaRecordFormatMappingParameters `hcl:"mapping_parameters,block"`
}

type InputsSchemaRecordFormatMappingParameters struct {
	// InputsSchemaRecordFormatMappingParametersCsv: optional
	Csv *InputsSchemaRecordFormatMappingParametersCsv `hcl:"csv,block"`
	// InputsSchemaRecordFormatMappingParametersJson: optional
	Json *InputsSchemaRecordFormatMappingParametersJson `hcl:"json,block"`
}

type InputsSchemaRecordFormatMappingParametersCsv struct {
	// RecordColumnDelimiter: string, required
	RecordColumnDelimiter terra.StringValue `hcl:"record_column_delimiter,attr" validate:"required"`
	// RecordRowDelimiter: string, required
	RecordRowDelimiter terra.StringValue `hcl:"record_row_delimiter,attr" validate:"required"`
}

type InputsSchemaRecordFormatMappingParametersJson struct {
	// RecordRowPath: string, required
	RecordRowPath terra.StringValue `hcl:"record_row_path,attr" validate:"required"`
}

type StartingPositionConfiguration struct {
	// StartingPosition: string, optional
	StartingPosition terra.StringValue `hcl:"starting_position,attr"`
}

type Outputs struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OutputsKinesisFirehose: optional
	KinesisFirehose *OutputsKinesisFirehose `hcl:"kinesis_firehose,block"`
	// OutputsKinesisStream: optional
	KinesisStream *OutputsKinesisStream `hcl:"kinesis_stream,block"`
	// OutputsLambda: optional
	Lambda *OutputsLambda `hcl:"lambda,block"`
	// OutputsSchema: required
	Schema *OutputsSchema `hcl:"schema,block" validate:"required"`
}

type OutputsKinesisFirehose struct {
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type OutputsKinesisStream struct {
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type OutputsLambda struct {
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type OutputsSchema struct {
	// RecordFormatType: string, required
	RecordFormatType terra.StringValue `hcl:"record_format_type,attr" validate:"required"`
}

type ReferenceDataSources struct {
	// TableName: string, required
	TableName terra.StringValue `hcl:"table_name,attr" validate:"required"`
	// S3: required
	S3 *S3 `hcl:"s3,block" validate:"required"`
	// ReferenceDataSourcesSchema: required
	Schema *ReferenceDataSourcesSchema `hcl:"schema,block" validate:"required"`
}

type S3 struct {
	// BucketArn: string, required
	BucketArn terra.StringValue `hcl:"bucket_arn,attr" validate:"required"`
	// FileKey: string, required
	FileKey terra.StringValue `hcl:"file_key,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type ReferenceDataSourcesSchema struct {
	// RecordEncoding: string, optional
	RecordEncoding terra.StringValue `hcl:"record_encoding,attr"`
	// ReferenceDataSourcesSchemaRecordColumns: min=1,max=1000
	RecordColumns []ReferenceDataSourcesSchemaRecordColumns `hcl:"record_columns,block" validate:"min=1,max=1000"`
	// ReferenceDataSourcesSchemaRecordFormat: required
	RecordFormat *ReferenceDataSourcesSchemaRecordFormat `hcl:"record_format,block" validate:"required"`
}

type ReferenceDataSourcesSchemaRecordColumns struct {
	// Mapping: string, optional
	Mapping terra.StringValue `hcl:"mapping,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SqlType: string, required
	SqlType terra.StringValue `hcl:"sql_type,attr" validate:"required"`
}

type ReferenceDataSourcesSchemaRecordFormat struct {
	// ReferenceDataSourcesSchemaRecordFormatMappingParameters: optional
	MappingParameters *ReferenceDataSourcesSchemaRecordFormatMappingParameters `hcl:"mapping_parameters,block"`
}

type ReferenceDataSourcesSchemaRecordFormatMappingParameters struct {
	// ReferenceDataSourcesSchemaRecordFormatMappingParametersCsv: optional
	Csv *ReferenceDataSourcesSchemaRecordFormatMappingParametersCsv `hcl:"csv,block"`
	// ReferenceDataSourcesSchemaRecordFormatMappingParametersJson: optional
	Json *ReferenceDataSourcesSchemaRecordFormatMappingParametersJson `hcl:"json,block"`
}

type ReferenceDataSourcesSchemaRecordFormatMappingParametersCsv struct {
	// RecordColumnDelimiter: string, required
	RecordColumnDelimiter terra.StringValue `hcl:"record_column_delimiter,attr" validate:"required"`
	// RecordRowDelimiter: string, required
	RecordRowDelimiter terra.StringValue `hcl:"record_row_delimiter,attr" validate:"required"`
}

type ReferenceDataSourcesSchemaRecordFormatMappingParametersJson struct {
	// RecordRowPath: string, required
	RecordRowPath terra.StringValue `hcl:"record_row_path,attr" validate:"required"`
}

type CloudwatchLoggingOptionsAttributes struct {
	ref terra.Reference
}

func (clo CloudwatchLoggingOptionsAttributes) InternalRef() (terra.Reference, error) {
	return clo.ref, nil
}

func (clo CloudwatchLoggingOptionsAttributes) InternalWithRef(ref terra.Reference) CloudwatchLoggingOptionsAttributes {
	return CloudwatchLoggingOptionsAttributes{ref: ref}
}

func (clo CloudwatchLoggingOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return clo.ref.InternalTokens()
}

func (clo CloudwatchLoggingOptionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(clo.ref.Append("id"))
}

func (clo CloudwatchLoggingOptionsAttributes) LogStreamArn() terra.StringValue {
	return terra.ReferenceAsString(clo.ref.Append("log_stream_arn"))
}

func (clo CloudwatchLoggingOptionsAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(clo.ref.Append("role_arn"))
}

type InputsAttributes struct {
	ref terra.Reference
}

func (i InputsAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i InputsAttributes) InternalWithRef(ref terra.Reference) InputsAttributes {
	return InputsAttributes{ref: ref}
}

func (i InputsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i InputsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("id"))
}

func (i InputsAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("name_prefix"))
}

func (i InputsAttributes) StreamNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("stream_names"))
}

func (i InputsAttributes) KinesisFirehose() terra.ListValue[InputsKinesisFirehoseAttributes] {
	return terra.ReferenceAsList[InputsKinesisFirehoseAttributes](i.ref.Append("kinesis_firehose"))
}

func (i InputsAttributes) KinesisStream() terra.ListValue[InputsKinesisStreamAttributes] {
	return terra.ReferenceAsList[InputsKinesisStreamAttributes](i.ref.Append("kinesis_stream"))
}

func (i InputsAttributes) Parallelism() terra.ListValue[ParallelismAttributes] {
	return terra.ReferenceAsList[ParallelismAttributes](i.ref.Append("parallelism"))
}

func (i InputsAttributes) ProcessingConfiguration() terra.ListValue[ProcessingConfigurationAttributes] {
	return terra.ReferenceAsList[ProcessingConfigurationAttributes](i.ref.Append("processing_configuration"))
}

func (i InputsAttributes) Schema() terra.ListValue[InputsSchemaAttributes] {
	return terra.ReferenceAsList[InputsSchemaAttributes](i.ref.Append("schema"))
}

func (i InputsAttributes) StartingPositionConfiguration() terra.ListValue[StartingPositionConfigurationAttributes] {
	return terra.ReferenceAsList[StartingPositionConfigurationAttributes](i.ref.Append("starting_position_configuration"))
}

type InputsKinesisFirehoseAttributes struct {
	ref terra.Reference
}

func (kf InputsKinesisFirehoseAttributes) InternalRef() (terra.Reference, error) {
	return kf.ref, nil
}

func (kf InputsKinesisFirehoseAttributes) InternalWithRef(ref terra.Reference) InputsKinesisFirehoseAttributes {
	return InputsKinesisFirehoseAttributes{ref: ref}
}

func (kf InputsKinesisFirehoseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kf.ref.InternalTokens()
}

func (kf InputsKinesisFirehoseAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("resource_arn"))
}

func (kf InputsKinesisFirehoseAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("role_arn"))
}

type InputsKinesisStreamAttributes struct {
	ref terra.Reference
}

func (ks InputsKinesisStreamAttributes) InternalRef() (terra.Reference, error) {
	return ks.ref, nil
}

func (ks InputsKinesisStreamAttributes) InternalWithRef(ref terra.Reference) InputsKinesisStreamAttributes {
	return InputsKinesisStreamAttributes{ref: ref}
}

func (ks InputsKinesisStreamAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ks.ref.InternalTokens()
}

func (ks InputsKinesisStreamAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("resource_arn"))
}

func (ks InputsKinesisStreamAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("role_arn"))
}

type ParallelismAttributes struct {
	ref terra.Reference
}

func (p ParallelismAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParallelismAttributes) InternalWithRef(ref terra.Reference) ParallelismAttributes {
	return ParallelismAttributes{ref: ref}
}

func (p ParallelismAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ParallelismAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("count"))
}

type ProcessingConfigurationAttributes struct {
	ref terra.Reference
}

func (pc ProcessingConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc ProcessingConfigurationAttributes) InternalWithRef(ref terra.Reference) ProcessingConfigurationAttributes {
	return ProcessingConfigurationAttributes{ref: ref}
}

func (pc ProcessingConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc ProcessingConfigurationAttributes) Lambda() terra.ListValue[ProcessingConfigurationLambdaAttributes] {
	return terra.ReferenceAsList[ProcessingConfigurationLambdaAttributes](pc.ref.Append("lambda"))
}

type ProcessingConfigurationLambdaAttributes struct {
	ref terra.Reference
}

func (l ProcessingConfigurationLambdaAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l ProcessingConfigurationLambdaAttributes) InternalWithRef(ref terra.Reference) ProcessingConfigurationLambdaAttributes {
	return ProcessingConfigurationLambdaAttributes{ref: ref}
}

func (l ProcessingConfigurationLambdaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l ProcessingConfigurationLambdaAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("resource_arn"))
}

func (l ProcessingConfigurationLambdaAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("role_arn"))
}

type InputsSchemaAttributes struct {
	ref terra.Reference
}

func (s InputsSchemaAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s InputsSchemaAttributes) InternalWithRef(ref terra.Reference) InputsSchemaAttributes {
	return InputsSchemaAttributes{ref: ref}
}

func (s InputsSchemaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s InputsSchemaAttributes) RecordEncoding() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("record_encoding"))
}

func (s InputsSchemaAttributes) RecordColumns() terra.ListValue[InputsSchemaRecordColumnsAttributes] {
	return terra.ReferenceAsList[InputsSchemaRecordColumnsAttributes](s.ref.Append("record_columns"))
}

func (s InputsSchemaAttributes) RecordFormat() terra.ListValue[InputsSchemaRecordFormatAttributes] {
	return terra.ReferenceAsList[InputsSchemaRecordFormatAttributes](s.ref.Append("record_format"))
}

type InputsSchemaRecordColumnsAttributes struct {
	ref terra.Reference
}

func (rc InputsSchemaRecordColumnsAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc InputsSchemaRecordColumnsAttributes) InternalWithRef(ref terra.Reference) InputsSchemaRecordColumnsAttributes {
	return InputsSchemaRecordColumnsAttributes{ref: ref}
}

func (rc InputsSchemaRecordColumnsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc InputsSchemaRecordColumnsAttributes) Mapping() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("mapping"))
}

func (rc InputsSchemaRecordColumnsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("name"))
}

func (rc InputsSchemaRecordColumnsAttributes) SqlType() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("sql_type"))
}

type InputsSchemaRecordFormatAttributes struct {
	ref terra.Reference
}

func (rf InputsSchemaRecordFormatAttributes) InternalRef() (terra.Reference, error) {
	return rf.ref, nil
}

func (rf InputsSchemaRecordFormatAttributes) InternalWithRef(ref terra.Reference) InputsSchemaRecordFormatAttributes {
	return InputsSchemaRecordFormatAttributes{ref: ref}
}

func (rf InputsSchemaRecordFormatAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rf.ref.InternalTokens()
}

func (rf InputsSchemaRecordFormatAttributes) RecordFormatType() terra.StringValue {
	return terra.ReferenceAsString(rf.ref.Append("record_format_type"))
}

func (rf InputsSchemaRecordFormatAttributes) MappingParameters() terra.ListValue[InputsSchemaRecordFormatMappingParametersAttributes] {
	return terra.ReferenceAsList[InputsSchemaRecordFormatMappingParametersAttributes](rf.ref.Append("mapping_parameters"))
}

type InputsSchemaRecordFormatMappingParametersAttributes struct {
	ref terra.Reference
}

func (mp InputsSchemaRecordFormatMappingParametersAttributes) InternalRef() (terra.Reference, error) {
	return mp.ref, nil
}

func (mp InputsSchemaRecordFormatMappingParametersAttributes) InternalWithRef(ref terra.Reference) InputsSchemaRecordFormatMappingParametersAttributes {
	return InputsSchemaRecordFormatMappingParametersAttributes{ref: ref}
}

func (mp InputsSchemaRecordFormatMappingParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mp.ref.InternalTokens()
}

func (mp InputsSchemaRecordFormatMappingParametersAttributes) Csv() terra.ListValue[InputsSchemaRecordFormatMappingParametersCsvAttributes] {
	return terra.ReferenceAsList[InputsSchemaRecordFormatMappingParametersCsvAttributes](mp.ref.Append("csv"))
}

func (mp InputsSchemaRecordFormatMappingParametersAttributes) Json() terra.ListValue[InputsSchemaRecordFormatMappingParametersJsonAttributes] {
	return terra.ReferenceAsList[InputsSchemaRecordFormatMappingParametersJsonAttributes](mp.ref.Append("json"))
}

type InputsSchemaRecordFormatMappingParametersCsvAttributes struct {
	ref terra.Reference
}

func (c InputsSchemaRecordFormatMappingParametersCsvAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c InputsSchemaRecordFormatMappingParametersCsvAttributes) InternalWithRef(ref terra.Reference) InputsSchemaRecordFormatMappingParametersCsvAttributes {
	return InputsSchemaRecordFormatMappingParametersCsvAttributes{ref: ref}
}

func (c InputsSchemaRecordFormatMappingParametersCsvAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c InputsSchemaRecordFormatMappingParametersCsvAttributes) RecordColumnDelimiter() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("record_column_delimiter"))
}

func (c InputsSchemaRecordFormatMappingParametersCsvAttributes) RecordRowDelimiter() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("record_row_delimiter"))
}

type InputsSchemaRecordFormatMappingParametersJsonAttributes struct {
	ref terra.Reference
}

func (j InputsSchemaRecordFormatMappingParametersJsonAttributes) InternalRef() (terra.Reference, error) {
	return j.ref, nil
}

func (j InputsSchemaRecordFormatMappingParametersJsonAttributes) InternalWithRef(ref terra.Reference) InputsSchemaRecordFormatMappingParametersJsonAttributes {
	return InputsSchemaRecordFormatMappingParametersJsonAttributes{ref: ref}
}

func (j InputsSchemaRecordFormatMappingParametersJsonAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return j.ref.InternalTokens()
}

func (j InputsSchemaRecordFormatMappingParametersJsonAttributes) RecordRowPath() terra.StringValue {
	return terra.ReferenceAsString(j.ref.Append("record_row_path"))
}

type StartingPositionConfigurationAttributes struct {
	ref terra.Reference
}

func (spc StartingPositionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return spc.ref, nil
}

func (spc StartingPositionConfigurationAttributes) InternalWithRef(ref terra.Reference) StartingPositionConfigurationAttributes {
	return StartingPositionConfigurationAttributes{ref: ref}
}

func (spc StartingPositionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return spc.ref.InternalTokens()
}

func (spc StartingPositionConfigurationAttributes) StartingPosition() terra.StringValue {
	return terra.ReferenceAsString(spc.ref.Append("starting_position"))
}

type OutputsAttributes struct {
	ref terra.Reference
}

func (o OutputsAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OutputsAttributes) InternalWithRef(ref terra.Reference) OutputsAttributes {
	return OutputsAttributes{ref: ref}
}

func (o OutputsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OutputsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("id"))
}

func (o OutputsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("name"))
}

func (o OutputsAttributes) KinesisFirehose() terra.ListValue[OutputsKinesisFirehoseAttributes] {
	return terra.ReferenceAsList[OutputsKinesisFirehoseAttributes](o.ref.Append("kinesis_firehose"))
}

func (o OutputsAttributes) KinesisStream() terra.ListValue[OutputsKinesisStreamAttributes] {
	return terra.ReferenceAsList[OutputsKinesisStreamAttributes](o.ref.Append("kinesis_stream"))
}

func (o OutputsAttributes) Lambda() terra.ListValue[OutputsLambdaAttributes] {
	return terra.ReferenceAsList[OutputsLambdaAttributes](o.ref.Append("lambda"))
}

func (o OutputsAttributes) Schema() terra.ListValue[OutputsSchemaAttributes] {
	return terra.ReferenceAsList[OutputsSchemaAttributes](o.ref.Append("schema"))
}

type OutputsKinesisFirehoseAttributes struct {
	ref terra.Reference
}

func (kf OutputsKinesisFirehoseAttributes) InternalRef() (terra.Reference, error) {
	return kf.ref, nil
}

func (kf OutputsKinesisFirehoseAttributes) InternalWithRef(ref terra.Reference) OutputsKinesisFirehoseAttributes {
	return OutputsKinesisFirehoseAttributes{ref: ref}
}

func (kf OutputsKinesisFirehoseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kf.ref.InternalTokens()
}

func (kf OutputsKinesisFirehoseAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("resource_arn"))
}

func (kf OutputsKinesisFirehoseAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("role_arn"))
}

type OutputsKinesisStreamAttributes struct {
	ref terra.Reference
}

func (ks OutputsKinesisStreamAttributes) InternalRef() (terra.Reference, error) {
	return ks.ref, nil
}

func (ks OutputsKinesisStreamAttributes) InternalWithRef(ref terra.Reference) OutputsKinesisStreamAttributes {
	return OutputsKinesisStreamAttributes{ref: ref}
}

func (ks OutputsKinesisStreamAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ks.ref.InternalTokens()
}

func (ks OutputsKinesisStreamAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("resource_arn"))
}

func (ks OutputsKinesisStreamAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("role_arn"))
}

type OutputsLambdaAttributes struct {
	ref terra.Reference
}

func (l OutputsLambdaAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l OutputsLambdaAttributes) InternalWithRef(ref terra.Reference) OutputsLambdaAttributes {
	return OutputsLambdaAttributes{ref: ref}
}

func (l OutputsLambdaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l OutputsLambdaAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("resource_arn"))
}

func (l OutputsLambdaAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("role_arn"))
}

type OutputsSchemaAttributes struct {
	ref terra.Reference
}

func (s OutputsSchemaAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s OutputsSchemaAttributes) InternalWithRef(ref terra.Reference) OutputsSchemaAttributes {
	return OutputsSchemaAttributes{ref: ref}
}

func (s OutputsSchemaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s OutputsSchemaAttributes) RecordFormatType() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("record_format_type"))
}

type ReferenceDataSourcesAttributes struct {
	ref terra.Reference
}

func (rds ReferenceDataSourcesAttributes) InternalRef() (terra.Reference, error) {
	return rds.ref, nil
}

func (rds ReferenceDataSourcesAttributes) InternalWithRef(ref terra.Reference) ReferenceDataSourcesAttributes {
	return ReferenceDataSourcesAttributes{ref: ref}
}

func (rds ReferenceDataSourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rds.ref.InternalTokens()
}

func (rds ReferenceDataSourcesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rds.ref.Append("id"))
}

func (rds ReferenceDataSourcesAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(rds.ref.Append("table_name"))
}

func (rds ReferenceDataSourcesAttributes) S3() terra.ListValue[S3Attributes] {
	return terra.ReferenceAsList[S3Attributes](rds.ref.Append("s3"))
}

func (rds ReferenceDataSourcesAttributes) Schema() terra.ListValue[ReferenceDataSourcesSchemaAttributes] {
	return terra.ReferenceAsList[ReferenceDataSourcesSchemaAttributes](rds.ref.Append("schema"))
}

type S3Attributes struct {
	ref terra.Reference
}

func (s S3Attributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s S3Attributes) InternalWithRef(ref terra.Reference) S3Attributes {
	return S3Attributes{ref: ref}
}

func (s S3Attributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s S3Attributes) BucketArn() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("bucket_arn"))
}

func (s S3Attributes) FileKey() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("file_key"))
}

func (s S3Attributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("role_arn"))
}

type ReferenceDataSourcesSchemaAttributes struct {
	ref terra.Reference
}

func (s ReferenceDataSourcesSchemaAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ReferenceDataSourcesSchemaAttributes) InternalWithRef(ref terra.Reference) ReferenceDataSourcesSchemaAttributes {
	return ReferenceDataSourcesSchemaAttributes{ref: ref}
}

func (s ReferenceDataSourcesSchemaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ReferenceDataSourcesSchemaAttributes) RecordEncoding() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("record_encoding"))
}

func (s ReferenceDataSourcesSchemaAttributes) RecordColumns() terra.ListValue[ReferenceDataSourcesSchemaRecordColumnsAttributes] {
	return terra.ReferenceAsList[ReferenceDataSourcesSchemaRecordColumnsAttributes](s.ref.Append("record_columns"))
}

func (s ReferenceDataSourcesSchemaAttributes) RecordFormat() terra.ListValue[ReferenceDataSourcesSchemaRecordFormatAttributes] {
	return terra.ReferenceAsList[ReferenceDataSourcesSchemaRecordFormatAttributes](s.ref.Append("record_format"))
}

type ReferenceDataSourcesSchemaRecordColumnsAttributes struct {
	ref terra.Reference
}

func (rc ReferenceDataSourcesSchemaRecordColumnsAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc ReferenceDataSourcesSchemaRecordColumnsAttributes) InternalWithRef(ref terra.Reference) ReferenceDataSourcesSchemaRecordColumnsAttributes {
	return ReferenceDataSourcesSchemaRecordColumnsAttributes{ref: ref}
}

func (rc ReferenceDataSourcesSchemaRecordColumnsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc ReferenceDataSourcesSchemaRecordColumnsAttributes) Mapping() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("mapping"))
}

func (rc ReferenceDataSourcesSchemaRecordColumnsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("name"))
}

func (rc ReferenceDataSourcesSchemaRecordColumnsAttributes) SqlType() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("sql_type"))
}

type ReferenceDataSourcesSchemaRecordFormatAttributes struct {
	ref terra.Reference
}

func (rf ReferenceDataSourcesSchemaRecordFormatAttributes) InternalRef() (terra.Reference, error) {
	return rf.ref, nil
}

func (rf ReferenceDataSourcesSchemaRecordFormatAttributes) InternalWithRef(ref terra.Reference) ReferenceDataSourcesSchemaRecordFormatAttributes {
	return ReferenceDataSourcesSchemaRecordFormatAttributes{ref: ref}
}

func (rf ReferenceDataSourcesSchemaRecordFormatAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rf.ref.InternalTokens()
}

func (rf ReferenceDataSourcesSchemaRecordFormatAttributes) RecordFormatType() terra.StringValue {
	return terra.ReferenceAsString(rf.ref.Append("record_format_type"))
}

func (rf ReferenceDataSourcesSchemaRecordFormatAttributes) MappingParameters() terra.ListValue[ReferenceDataSourcesSchemaRecordFormatMappingParametersAttributes] {
	return terra.ReferenceAsList[ReferenceDataSourcesSchemaRecordFormatMappingParametersAttributes](rf.ref.Append("mapping_parameters"))
}

type ReferenceDataSourcesSchemaRecordFormatMappingParametersAttributes struct {
	ref terra.Reference
}

func (mp ReferenceDataSourcesSchemaRecordFormatMappingParametersAttributes) InternalRef() (terra.Reference, error) {
	return mp.ref, nil
}

func (mp ReferenceDataSourcesSchemaRecordFormatMappingParametersAttributes) InternalWithRef(ref terra.Reference) ReferenceDataSourcesSchemaRecordFormatMappingParametersAttributes {
	return ReferenceDataSourcesSchemaRecordFormatMappingParametersAttributes{ref: ref}
}

func (mp ReferenceDataSourcesSchemaRecordFormatMappingParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mp.ref.InternalTokens()
}

func (mp ReferenceDataSourcesSchemaRecordFormatMappingParametersAttributes) Csv() terra.ListValue[ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvAttributes] {
	return terra.ReferenceAsList[ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvAttributes](mp.ref.Append("csv"))
}

func (mp ReferenceDataSourcesSchemaRecordFormatMappingParametersAttributes) Json() terra.ListValue[ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonAttributes] {
	return terra.ReferenceAsList[ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonAttributes](mp.ref.Append("json"))
}

type ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvAttributes struct {
	ref terra.Reference
}

func (c ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvAttributes) InternalWithRef(ref terra.Reference) ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvAttributes {
	return ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvAttributes{ref: ref}
}

func (c ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvAttributes) RecordColumnDelimiter() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("record_column_delimiter"))
}

func (c ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvAttributes) RecordRowDelimiter() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("record_row_delimiter"))
}

type ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonAttributes struct {
	ref terra.Reference
}

func (j ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonAttributes) InternalRef() (terra.Reference, error) {
	return j.ref, nil
}

func (j ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonAttributes) InternalWithRef(ref terra.Reference) ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonAttributes {
	return ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonAttributes{ref: ref}
}

func (j ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return j.ref.InternalTokens()
}

func (j ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonAttributes) RecordRowPath() terra.StringValue {
	return terra.ReferenceAsString(j.ref.Append("record_row_path"))
}

type CloudwatchLoggingOptionsState struct {
	Id           string `json:"id"`
	LogStreamArn string `json:"log_stream_arn"`
	RoleArn      string `json:"role_arn"`
}

type InputsState struct {
	Id                            string                               `json:"id"`
	NamePrefix                    string                               `json:"name_prefix"`
	StreamNames                   []string                             `json:"stream_names"`
	KinesisFirehose               []InputsKinesisFirehoseState         `json:"kinesis_firehose"`
	KinesisStream                 []InputsKinesisStreamState           `json:"kinesis_stream"`
	Parallelism                   []ParallelismState                   `json:"parallelism"`
	ProcessingConfiguration       []ProcessingConfigurationState       `json:"processing_configuration"`
	Schema                        []InputsSchemaState                  `json:"schema"`
	StartingPositionConfiguration []StartingPositionConfigurationState `json:"starting_position_configuration"`
}

type InputsKinesisFirehoseState struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type InputsKinesisStreamState struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type ParallelismState struct {
	Count float64 `json:"count"`
}

type ProcessingConfigurationState struct {
	Lambda []ProcessingConfigurationLambdaState `json:"lambda"`
}

type ProcessingConfigurationLambdaState struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type InputsSchemaState struct {
	RecordEncoding string                           `json:"record_encoding"`
	RecordColumns  []InputsSchemaRecordColumnsState `json:"record_columns"`
	RecordFormat   []InputsSchemaRecordFormatState  `json:"record_format"`
}

type InputsSchemaRecordColumnsState struct {
	Mapping string `json:"mapping"`
	Name    string `json:"name"`
	SqlType string `json:"sql_type"`
}

type InputsSchemaRecordFormatState struct {
	RecordFormatType  string                                           `json:"record_format_type"`
	MappingParameters []InputsSchemaRecordFormatMappingParametersState `json:"mapping_parameters"`
}

type InputsSchemaRecordFormatMappingParametersState struct {
	Csv  []InputsSchemaRecordFormatMappingParametersCsvState  `json:"csv"`
	Json []InputsSchemaRecordFormatMappingParametersJsonState `json:"json"`
}

type InputsSchemaRecordFormatMappingParametersCsvState struct {
	RecordColumnDelimiter string `json:"record_column_delimiter"`
	RecordRowDelimiter    string `json:"record_row_delimiter"`
}

type InputsSchemaRecordFormatMappingParametersJsonState struct {
	RecordRowPath string `json:"record_row_path"`
}

type StartingPositionConfigurationState struct {
	StartingPosition string `json:"starting_position"`
}

type OutputsState struct {
	Id              string                        `json:"id"`
	Name            string                        `json:"name"`
	KinesisFirehose []OutputsKinesisFirehoseState `json:"kinesis_firehose"`
	KinesisStream   []OutputsKinesisStreamState   `json:"kinesis_stream"`
	Lambda          []OutputsLambdaState          `json:"lambda"`
	Schema          []OutputsSchemaState          `json:"schema"`
}

type OutputsKinesisFirehoseState struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type OutputsKinesisStreamState struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type OutputsLambdaState struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type OutputsSchemaState struct {
	RecordFormatType string `json:"record_format_type"`
}

type ReferenceDataSourcesState struct {
	Id        string                            `json:"id"`
	TableName string                            `json:"table_name"`
	S3        []S3State                         `json:"s3"`
	Schema    []ReferenceDataSourcesSchemaState `json:"schema"`
}

type S3State struct {
	BucketArn string `json:"bucket_arn"`
	FileKey   string `json:"file_key"`
	RoleArn   string `json:"role_arn"`
}

type ReferenceDataSourcesSchemaState struct {
	RecordEncoding string                                         `json:"record_encoding"`
	RecordColumns  []ReferenceDataSourcesSchemaRecordColumnsState `json:"record_columns"`
	RecordFormat   []ReferenceDataSourcesSchemaRecordFormatState  `json:"record_format"`
}

type ReferenceDataSourcesSchemaRecordColumnsState struct {
	Mapping string `json:"mapping"`
	Name    string `json:"name"`
	SqlType string `json:"sql_type"`
}

type ReferenceDataSourcesSchemaRecordFormatState struct {
	RecordFormatType  string                                                         `json:"record_format_type"`
	MappingParameters []ReferenceDataSourcesSchemaRecordFormatMappingParametersState `json:"mapping_parameters"`
}

type ReferenceDataSourcesSchemaRecordFormatMappingParametersState struct {
	Csv  []ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvState  `json:"csv"`
	Json []ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonState `json:"json"`
}

type ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvState struct {
	RecordColumnDelimiter string `json:"record_column_delimiter"`
	RecordRowDelimiter    string `json:"record_row_delimiter"`
}

type ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonState struct {
	RecordRowPath string `json:"record_row_path"`
}