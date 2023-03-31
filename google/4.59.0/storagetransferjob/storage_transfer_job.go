// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package storagetransferjob

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type NotificationConfig struct {
	// EventTypes: set of string, optional
	EventTypes terra.SetValue[terra.StringValue] `hcl:"event_types,attr"`
	// PayloadFormat: string, required
	PayloadFormat terra.StringValue `hcl:"payload_format,attr" validate:"required"`
	// PubsubTopic: string, required
	PubsubTopic terra.StringValue `hcl:"pubsub_topic,attr" validate:"required"`
}

type Schedule struct {
	// RepeatInterval: string, optional
	RepeatInterval terra.StringValue `hcl:"repeat_interval,attr"`
	// ScheduleEndDate: optional
	ScheduleEndDate *ScheduleEndDate `hcl:"schedule_end_date,block"`
	// ScheduleStartDate: required
	ScheduleStartDate *ScheduleStartDate `hcl:"schedule_start_date,block" validate:"required"`
	// StartTimeOfDay: optional
	StartTimeOfDay *StartTimeOfDay `hcl:"start_time_of_day,block"`
}

type ScheduleEndDate struct {
	// Day: number, required
	Day terra.NumberValue `hcl:"day,attr" validate:"required"`
	// Month: number, required
	Month terra.NumberValue `hcl:"month,attr" validate:"required"`
	// Year: number, required
	Year terra.NumberValue `hcl:"year,attr" validate:"required"`
}

type ScheduleStartDate struct {
	// Day: number, required
	Day terra.NumberValue `hcl:"day,attr" validate:"required"`
	// Month: number, required
	Month terra.NumberValue `hcl:"month,attr" validate:"required"`
	// Year: number, required
	Year terra.NumberValue `hcl:"year,attr" validate:"required"`
}

type StartTimeOfDay struct {
	// Hours: number, required
	Hours terra.NumberValue `hcl:"hours,attr" validate:"required"`
	// Minutes: number, required
	Minutes terra.NumberValue `hcl:"minutes,attr" validate:"required"`
	// Nanos: number, required
	Nanos terra.NumberValue `hcl:"nanos,attr" validate:"required"`
	// Seconds: number, required
	Seconds terra.NumberValue `hcl:"seconds,attr" validate:"required"`
}

type TransferSpec struct {
	// SinkAgentPoolName: string, optional
	SinkAgentPoolName terra.StringValue `hcl:"sink_agent_pool_name,attr"`
	// SourceAgentPoolName: string, optional
	SourceAgentPoolName terra.StringValue `hcl:"source_agent_pool_name,attr"`
	// AwsS3DataSource: optional
	AwsS3DataSource *AwsS3DataSource `hcl:"aws_s3_data_source,block"`
	// AzureBlobStorageDataSource: optional
	AzureBlobStorageDataSource *AzureBlobStorageDataSource `hcl:"azure_blob_storage_data_source,block"`
	// GcsDataSink: optional
	GcsDataSink *GcsDataSink `hcl:"gcs_data_sink,block"`
	// GcsDataSource: optional
	GcsDataSource *GcsDataSource `hcl:"gcs_data_source,block"`
	// HttpDataSource: optional
	HttpDataSource *HttpDataSource `hcl:"http_data_source,block"`
	// ObjectConditions: optional
	ObjectConditions *ObjectConditions `hcl:"object_conditions,block"`
	// PosixDataSink: optional
	PosixDataSink *PosixDataSink `hcl:"posix_data_sink,block"`
	// PosixDataSource: optional
	PosixDataSource *PosixDataSource `hcl:"posix_data_source,block"`
	// TransferOptions: optional
	TransferOptions *TransferOptions `hcl:"transfer_options,block"`
}

type AwsS3DataSource struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// AwsAccessKey: optional
	AwsAccessKey *AwsAccessKey `hcl:"aws_access_key,block"`
}

type AwsAccessKey struct {
	// AccessKeyId: string, required
	AccessKeyId terra.StringValue `hcl:"access_key_id,attr" validate:"required"`
	// SecretAccessKey: string, required
	SecretAccessKey terra.StringValue `hcl:"secret_access_key,attr" validate:"required"`
}

type AzureBlobStorageDataSource struct {
	// Container: string, required
	Container terra.StringValue `hcl:"container,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// StorageAccount: string, required
	StorageAccount terra.StringValue `hcl:"storage_account,attr" validate:"required"`
	// AzureCredentials: required
	AzureCredentials *AzureCredentials `hcl:"azure_credentials,block" validate:"required"`
}

type AzureCredentials struct {
	// SasToken: string, required
	SasToken terra.StringValue `hcl:"sas_token,attr" validate:"required"`
}

type GcsDataSink struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
}

type GcsDataSource struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
}

type HttpDataSource struct {
	// ListUrl: string, required
	ListUrl terra.StringValue `hcl:"list_url,attr" validate:"required"`
}

type ObjectConditions struct {
	// ExcludePrefixes: list of string, optional
	ExcludePrefixes terra.ListValue[terra.StringValue] `hcl:"exclude_prefixes,attr"`
	// IncludePrefixes: list of string, optional
	IncludePrefixes terra.ListValue[terra.StringValue] `hcl:"include_prefixes,attr"`
	// MaxTimeElapsedSinceLastModification: string, optional
	MaxTimeElapsedSinceLastModification terra.StringValue `hcl:"max_time_elapsed_since_last_modification,attr"`
	// MinTimeElapsedSinceLastModification: string, optional
	MinTimeElapsedSinceLastModification terra.StringValue `hcl:"min_time_elapsed_since_last_modification,attr"`
}

type PosixDataSink struct {
	// RootDirectory: string, required
	RootDirectory terra.StringValue `hcl:"root_directory,attr" validate:"required"`
}

type PosixDataSource struct {
	// RootDirectory: string, required
	RootDirectory terra.StringValue `hcl:"root_directory,attr" validate:"required"`
}

type TransferOptions struct {
	// DeleteObjectsFromSourceAfterTransfer: bool, optional
	DeleteObjectsFromSourceAfterTransfer terra.BoolValue `hcl:"delete_objects_from_source_after_transfer,attr"`
	// DeleteObjectsUniqueInSink: bool, optional
	DeleteObjectsUniqueInSink terra.BoolValue `hcl:"delete_objects_unique_in_sink,attr"`
	// OverwriteObjectsAlreadyExistingInSink: bool, optional
	OverwriteObjectsAlreadyExistingInSink terra.BoolValue `hcl:"overwrite_objects_already_existing_in_sink,attr"`
	// OverwriteWhen: string, optional
	OverwriteWhen terra.StringValue `hcl:"overwrite_when,attr"`
}

type NotificationConfigAttributes struct {
	ref terra.Reference
}

func (nc NotificationConfigAttributes) InternalRef() terra.Reference {
	return nc.ref
}

func (nc NotificationConfigAttributes) InternalWithRef(ref terra.Reference) NotificationConfigAttributes {
	return NotificationConfigAttributes{ref: ref}
}

func (nc NotificationConfigAttributes) InternalTokens() hclwrite.Tokens {
	return nc.ref.InternalTokens()
}

func (nc NotificationConfigAttributes) EventTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](nc.ref.Append("event_types"))
}

func (nc NotificationConfigAttributes) PayloadFormat() terra.StringValue {
	return terra.ReferenceString(nc.ref.Append("payload_format"))
}

func (nc NotificationConfigAttributes) PubsubTopic() terra.StringValue {
	return terra.ReferenceString(nc.ref.Append("pubsub_topic"))
}

type ScheduleAttributes struct {
	ref terra.Reference
}

func (s ScheduleAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s ScheduleAttributes) InternalWithRef(ref terra.Reference) ScheduleAttributes {
	return ScheduleAttributes{ref: ref}
}

func (s ScheduleAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s ScheduleAttributes) RepeatInterval() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("repeat_interval"))
}

func (s ScheduleAttributes) ScheduleEndDate() terra.ListValue[ScheduleEndDateAttributes] {
	return terra.ReferenceList[ScheduleEndDateAttributes](s.ref.Append("schedule_end_date"))
}

func (s ScheduleAttributes) ScheduleStartDate() terra.ListValue[ScheduleStartDateAttributes] {
	return terra.ReferenceList[ScheduleStartDateAttributes](s.ref.Append("schedule_start_date"))
}

func (s ScheduleAttributes) StartTimeOfDay() terra.ListValue[StartTimeOfDayAttributes] {
	return terra.ReferenceList[StartTimeOfDayAttributes](s.ref.Append("start_time_of_day"))
}

type ScheduleEndDateAttributes struct {
	ref terra.Reference
}

func (sed ScheduleEndDateAttributes) InternalRef() terra.Reference {
	return sed.ref
}

func (sed ScheduleEndDateAttributes) InternalWithRef(ref terra.Reference) ScheduleEndDateAttributes {
	return ScheduleEndDateAttributes{ref: ref}
}

func (sed ScheduleEndDateAttributes) InternalTokens() hclwrite.Tokens {
	return sed.ref.InternalTokens()
}

func (sed ScheduleEndDateAttributes) Day() terra.NumberValue {
	return terra.ReferenceNumber(sed.ref.Append("day"))
}

func (sed ScheduleEndDateAttributes) Month() terra.NumberValue {
	return terra.ReferenceNumber(sed.ref.Append("month"))
}

func (sed ScheduleEndDateAttributes) Year() terra.NumberValue {
	return terra.ReferenceNumber(sed.ref.Append("year"))
}

type ScheduleStartDateAttributes struct {
	ref terra.Reference
}

func (ssd ScheduleStartDateAttributes) InternalRef() terra.Reference {
	return ssd.ref
}

func (ssd ScheduleStartDateAttributes) InternalWithRef(ref terra.Reference) ScheduleStartDateAttributes {
	return ScheduleStartDateAttributes{ref: ref}
}

func (ssd ScheduleStartDateAttributes) InternalTokens() hclwrite.Tokens {
	return ssd.ref.InternalTokens()
}

func (ssd ScheduleStartDateAttributes) Day() terra.NumberValue {
	return terra.ReferenceNumber(ssd.ref.Append("day"))
}

func (ssd ScheduleStartDateAttributes) Month() terra.NumberValue {
	return terra.ReferenceNumber(ssd.ref.Append("month"))
}

func (ssd ScheduleStartDateAttributes) Year() terra.NumberValue {
	return terra.ReferenceNumber(ssd.ref.Append("year"))
}

type StartTimeOfDayAttributes struct {
	ref terra.Reference
}

func (stod StartTimeOfDayAttributes) InternalRef() terra.Reference {
	return stod.ref
}

func (stod StartTimeOfDayAttributes) InternalWithRef(ref terra.Reference) StartTimeOfDayAttributes {
	return StartTimeOfDayAttributes{ref: ref}
}

func (stod StartTimeOfDayAttributes) InternalTokens() hclwrite.Tokens {
	return stod.ref.InternalTokens()
}

func (stod StartTimeOfDayAttributes) Hours() terra.NumberValue {
	return terra.ReferenceNumber(stod.ref.Append("hours"))
}

func (stod StartTimeOfDayAttributes) Minutes() terra.NumberValue {
	return terra.ReferenceNumber(stod.ref.Append("minutes"))
}

func (stod StartTimeOfDayAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceNumber(stod.ref.Append("nanos"))
}

func (stod StartTimeOfDayAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceNumber(stod.ref.Append("seconds"))
}

type TransferSpecAttributes struct {
	ref terra.Reference
}

func (ts TransferSpecAttributes) InternalRef() terra.Reference {
	return ts.ref
}

func (ts TransferSpecAttributes) InternalWithRef(ref terra.Reference) TransferSpecAttributes {
	return TransferSpecAttributes{ref: ref}
}

func (ts TransferSpecAttributes) InternalTokens() hclwrite.Tokens {
	return ts.ref.InternalTokens()
}

func (ts TransferSpecAttributes) SinkAgentPoolName() terra.StringValue {
	return terra.ReferenceString(ts.ref.Append("sink_agent_pool_name"))
}

func (ts TransferSpecAttributes) SourceAgentPoolName() terra.StringValue {
	return terra.ReferenceString(ts.ref.Append("source_agent_pool_name"))
}

func (ts TransferSpecAttributes) AwsS3DataSource() terra.ListValue[AwsS3DataSourceAttributes] {
	return terra.ReferenceList[AwsS3DataSourceAttributes](ts.ref.Append("aws_s3_data_source"))
}

func (ts TransferSpecAttributes) AzureBlobStorageDataSource() terra.ListValue[AzureBlobStorageDataSourceAttributes] {
	return terra.ReferenceList[AzureBlobStorageDataSourceAttributes](ts.ref.Append("azure_blob_storage_data_source"))
}

func (ts TransferSpecAttributes) GcsDataSink() terra.ListValue[GcsDataSinkAttributes] {
	return terra.ReferenceList[GcsDataSinkAttributes](ts.ref.Append("gcs_data_sink"))
}

func (ts TransferSpecAttributes) GcsDataSource() terra.ListValue[GcsDataSourceAttributes] {
	return terra.ReferenceList[GcsDataSourceAttributes](ts.ref.Append("gcs_data_source"))
}

func (ts TransferSpecAttributes) HttpDataSource() terra.ListValue[HttpDataSourceAttributes] {
	return terra.ReferenceList[HttpDataSourceAttributes](ts.ref.Append("http_data_source"))
}

func (ts TransferSpecAttributes) ObjectConditions() terra.ListValue[ObjectConditionsAttributes] {
	return terra.ReferenceList[ObjectConditionsAttributes](ts.ref.Append("object_conditions"))
}

func (ts TransferSpecAttributes) PosixDataSink() terra.ListValue[PosixDataSinkAttributes] {
	return terra.ReferenceList[PosixDataSinkAttributes](ts.ref.Append("posix_data_sink"))
}

func (ts TransferSpecAttributes) PosixDataSource() terra.ListValue[PosixDataSourceAttributes] {
	return terra.ReferenceList[PosixDataSourceAttributes](ts.ref.Append("posix_data_source"))
}

func (ts TransferSpecAttributes) TransferOptions() terra.ListValue[TransferOptionsAttributes] {
	return terra.ReferenceList[TransferOptionsAttributes](ts.ref.Append("transfer_options"))
}

type AwsS3DataSourceAttributes struct {
	ref terra.Reference
}

func (asds AwsS3DataSourceAttributes) InternalRef() terra.Reference {
	return asds.ref
}

func (asds AwsS3DataSourceAttributes) InternalWithRef(ref terra.Reference) AwsS3DataSourceAttributes {
	return AwsS3DataSourceAttributes{ref: ref}
}

func (asds AwsS3DataSourceAttributes) InternalTokens() hclwrite.Tokens {
	return asds.ref.InternalTokens()
}

func (asds AwsS3DataSourceAttributes) BucketName() terra.StringValue {
	return terra.ReferenceString(asds.ref.Append("bucket_name"))
}

func (asds AwsS3DataSourceAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceString(asds.ref.Append("role_arn"))
}

func (asds AwsS3DataSourceAttributes) AwsAccessKey() terra.ListValue[AwsAccessKeyAttributes] {
	return terra.ReferenceList[AwsAccessKeyAttributes](asds.ref.Append("aws_access_key"))
}

type AwsAccessKeyAttributes struct {
	ref terra.Reference
}

func (aak AwsAccessKeyAttributes) InternalRef() terra.Reference {
	return aak.ref
}

func (aak AwsAccessKeyAttributes) InternalWithRef(ref terra.Reference) AwsAccessKeyAttributes {
	return AwsAccessKeyAttributes{ref: ref}
}

func (aak AwsAccessKeyAttributes) InternalTokens() hclwrite.Tokens {
	return aak.ref.InternalTokens()
}

func (aak AwsAccessKeyAttributes) AccessKeyId() terra.StringValue {
	return terra.ReferenceString(aak.ref.Append("access_key_id"))
}

func (aak AwsAccessKeyAttributes) SecretAccessKey() terra.StringValue {
	return terra.ReferenceString(aak.ref.Append("secret_access_key"))
}

type AzureBlobStorageDataSourceAttributes struct {
	ref terra.Reference
}

func (absds AzureBlobStorageDataSourceAttributes) InternalRef() terra.Reference {
	return absds.ref
}

func (absds AzureBlobStorageDataSourceAttributes) InternalWithRef(ref terra.Reference) AzureBlobStorageDataSourceAttributes {
	return AzureBlobStorageDataSourceAttributes{ref: ref}
}

func (absds AzureBlobStorageDataSourceAttributes) InternalTokens() hclwrite.Tokens {
	return absds.ref.InternalTokens()
}

func (absds AzureBlobStorageDataSourceAttributes) Container() terra.StringValue {
	return terra.ReferenceString(absds.ref.Append("container"))
}

func (absds AzureBlobStorageDataSourceAttributes) Path() terra.StringValue {
	return terra.ReferenceString(absds.ref.Append("path"))
}

func (absds AzureBlobStorageDataSourceAttributes) StorageAccount() terra.StringValue {
	return terra.ReferenceString(absds.ref.Append("storage_account"))
}

func (absds AzureBlobStorageDataSourceAttributes) AzureCredentials() terra.ListValue[AzureCredentialsAttributes] {
	return terra.ReferenceList[AzureCredentialsAttributes](absds.ref.Append("azure_credentials"))
}

type AzureCredentialsAttributes struct {
	ref terra.Reference
}

func (ac AzureCredentialsAttributes) InternalRef() terra.Reference {
	return ac.ref
}

func (ac AzureCredentialsAttributes) InternalWithRef(ref terra.Reference) AzureCredentialsAttributes {
	return AzureCredentialsAttributes{ref: ref}
}

func (ac AzureCredentialsAttributes) InternalTokens() hclwrite.Tokens {
	return ac.ref.InternalTokens()
}

func (ac AzureCredentialsAttributes) SasToken() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("sas_token"))
}

type GcsDataSinkAttributes struct {
	ref terra.Reference
}

func (gds GcsDataSinkAttributes) InternalRef() terra.Reference {
	return gds.ref
}

func (gds GcsDataSinkAttributes) InternalWithRef(ref terra.Reference) GcsDataSinkAttributes {
	return GcsDataSinkAttributes{ref: ref}
}

func (gds GcsDataSinkAttributes) InternalTokens() hclwrite.Tokens {
	return gds.ref.InternalTokens()
}

func (gds GcsDataSinkAttributes) BucketName() terra.StringValue {
	return terra.ReferenceString(gds.ref.Append("bucket_name"))
}

func (gds GcsDataSinkAttributes) Path() terra.StringValue {
	return terra.ReferenceString(gds.ref.Append("path"))
}

type GcsDataSourceAttributes struct {
	ref terra.Reference
}

func (gds GcsDataSourceAttributes) InternalRef() terra.Reference {
	return gds.ref
}

func (gds GcsDataSourceAttributes) InternalWithRef(ref terra.Reference) GcsDataSourceAttributes {
	return GcsDataSourceAttributes{ref: ref}
}

func (gds GcsDataSourceAttributes) InternalTokens() hclwrite.Tokens {
	return gds.ref.InternalTokens()
}

func (gds GcsDataSourceAttributes) BucketName() terra.StringValue {
	return terra.ReferenceString(gds.ref.Append("bucket_name"))
}

func (gds GcsDataSourceAttributes) Path() terra.StringValue {
	return terra.ReferenceString(gds.ref.Append("path"))
}

type HttpDataSourceAttributes struct {
	ref terra.Reference
}

func (hds HttpDataSourceAttributes) InternalRef() terra.Reference {
	return hds.ref
}

func (hds HttpDataSourceAttributes) InternalWithRef(ref terra.Reference) HttpDataSourceAttributes {
	return HttpDataSourceAttributes{ref: ref}
}

func (hds HttpDataSourceAttributes) InternalTokens() hclwrite.Tokens {
	return hds.ref.InternalTokens()
}

func (hds HttpDataSourceAttributes) ListUrl() terra.StringValue {
	return terra.ReferenceString(hds.ref.Append("list_url"))
}

type ObjectConditionsAttributes struct {
	ref terra.Reference
}

func (oc ObjectConditionsAttributes) InternalRef() terra.Reference {
	return oc.ref
}

func (oc ObjectConditionsAttributes) InternalWithRef(ref terra.Reference) ObjectConditionsAttributes {
	return ObjectConditionsAttributes{ref: ref}
}

func (oc ObjectConditionsAttributes) InternalTokens() hclwrite.Tokens {
	return oc.ref.InternalTokens()
}

func (oc ObjectConditionsAttributes) ExcludePrefixes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](oc.ref.Append("exclude_prefixes"))
}

func (oc ObjectConditionsAttributes) IncludePrefixes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](oc.ref.Append("include_prefixes"))
}

func (oc ObjectConditionsAttributes) MaxTimeElapsedSinceLastModification() terra.StringValue {
	return terra.ReferenceString(oc.ref.Append("max_time_elapsed_since_last_modification"))
}

func (oc ObjectConditionsAttributes) MinTimeElapsedSinceLastModification() terra.StringValue {
	return terra.ReferenceString(oc.ref.Append("min_time_elapsed_since_last_modification"))
}

type PosixDataSinkAttributes struct {
	ref terra.Reference
}

func (pds PosixDataSinkAttributes) InternalRef() terra.Reference {
	return pds.ref
}

func (pds PosixDataSinkAttributes) InternalWithRef(ref terra.Reference) PosixDataSinkAttributes {
	return PosixDataSinkAttributes{ref: ref}
}

func (pds PosixDataSinkAttributes) InternalTokens() hclwrite.Tokens {
	return pds.ref.InternalTokens()
}

func (pds PosixDataSinkAttributes) RootDirectory() terra.StringValue {
	return terra.ReferenceString(pds.ref.Append("root_directory"))
}

type PosixDataSourceAttributes struct {
	ref terra.Reference
}

func (pds PosixDataSourceAttributes) InternalRef() terra.Reference {
	return pds.ref
}

func (pds PosixDataSourceAttributes) InternalWithRef(ref terra.Reference) PosixDataSourceAttributes {
	return PosixDataSourceAttributes{ref: ref}
}

func (pds PosixDataSourceAttributes) InternalTokens() hclwrite.Tokens {
	return pds.ref.InternalTokens()
}

func (pds PosixDataSourceAttributes) RootDirectory() terra.StringValue {
	return terra.ReferenceString(pds.ref.Append("root_directory"))
}

type TransferOptionsAttributes struct {
	ref terra.Reference
}

func (to TransferOptionsAttributes) InternalRef() terra.Reference {
	return to.ref
}

func (to TransferOptionsAttributes) InternalWithRef(ref terra.Reference) TransferOptionsAttributes {
	return TransferOptionsAttributes{ref: ref}
}

func (to TransferOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return to.ref.InternalTokens()
}

func (to TransferOptionsAttributes) DeleteObjectsFromSourceAfterTransfer() terra.BoolValue {
	return terra.ReferenceBool(to.ref.Append("delete_objects_from_source_after_transfer"))
}

func (to TransferOptionsAttributes) DeleteObjectsUniqueInSink() terra.BoolValue {
	return terra.ReferenceBool(to.ref.Append("delete_objects_unique_in_sink"))
}

func (to TransferOptionsAttributes) OverwriteObjectsAlreadyExistingInSink() terra.BoolValue {
	return terra.ReferenceBool(to.ref.Append("overwrite_objects_already_existing_in_sink"))
}

func (to TransferOptionsAttributes) OverwriteWhen() terra.StringValue {
	return terra.ReferenceString(to.ref.Append("overwrite_when"))
}

type NotificationConfigState struct {
	EventTypes    []string `json:"event_types"`
	PayloadFormat string   `json:"payload_format"`
	PubsubTopic   string   `json:"pubsub_topic"`
}

type ScheduleState struct {
	RepeatInterval    string                   `json:"repeat_interval"`
	ScheduleEndDate   []ScheduleEndDateState   `json:"schedule_end_date"`
	ScheduleStartDate []ScheduleStartDateState `json:"schedule_start_date"`
	StartTimeOfDay    []StartTimeOfDayState    `json:"start_time_of_day"`
}

type ScheduleEndDateState struct {
	Day   float64 `json:"day"`
	Month float64 `json:"month"`
	Year  float64 `json:"year"`
}

type ScheduleStartDateState struct {
	Day   float64 `json:"day"`
	Month float64 `json:"month"`
	Year  float64 `json:"year"`
}

type StartTimeOfDayState struct {
	Hours   float64 `json:"hours"`
	Minutes float64 `json:"minutes"`
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type TransferSpecState struct {
	SinkAgentPoolName          string                            `json:"sink_agent_pool_name"`
	SourceAgentPoolName        string                            `json:"source_agent_pool_name"`
	AwsS3DataSource            []AwsS3DataSourceState            `json:"aws_s3_data_source"`
	AzureBlobStorageDataSource []AzureBlobStorageDataSourceState `json:"azure_blob_storage_data_source"`
	GcsDataSink                []GcsDataSinkState                `json:"gcs_data_sink"`
	GcsDataSource              []GcsDataSourceState              `json:"gcs_data_source"`
	HttpDataSource             []HttpDataSourceState             `json:"http_data_source"`
	ObjectConditions           []ObjectConditionsState           `json:"object_conditions"`
	PosixDataSink              []PosixDataSinkState              `json:"posix_data_sink"`
	PosixDataSource            []PosixDataSourceState            `json:"posix_data_source"`
	TransferOptions            []TransferOptionsState            `json:"transfer_options"`
}

type AwsS3DataSourceState struct {
	BucketName   string              `json:"bucket_name"`
	RoleArn      string              `json:"role_arn"`
	AwsAccessKey []AwsAccessKeyState `json:"aws_access_key"`
}

type AwsAccessKeyState struct {
	AccessKeyId     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
}

type AzureBlobStorageDataSourceState struct {
	Container        string                  `json:"container"`
	Path             string                  `json:"path"`
	StorageAccount   string                  `json:"storage_account"`
	AzureCredentials []AzureCredentialsState `json:"azure_credentials"`
}

type AzureCredentialsState struct {
	SasToken string `json:"sas_token"`
}

type GcsDataSinkState struct {
	BucketName string `json:"bucket_name"`
	Path       string `json:"path"`
}

type GcsDataSourceState struct {
	BucketName string `json:"bucket_name"`
	Path       string `json:"path"`
}

type HttpDataSourceState struct {
	ListUrl string `json:"list_url"`
}

type ObjectConditionsState struct {
	ExcludePrefixes                     []string `json:"exclude_prefixes"`
	IncludePrefixes                     []string `json:"include_prefixes"`
	MaxTimeElapsedSinceLastModification string   `json:"max_time_elapsed_since_last_modification"`
	MinTimeElapsedSinceLastModification string   `json:"min_time_elapsed_since_last_modification"`
}

type PosixDataSinkState struct {
	RootDirectory string `json:"root_directory"`
}

type PosixDataSourceState struct {
	RootDirectory string `json:"root_directory"`
}

type TransferOptionsState struct {
	DeleteObjectsFromSourceAfterTransfer  bool   `json:"delete_objects_from_source_after_transfer"`
	DeleteObjectsUniqueInSink             bool   `json:"delete_objects_unique_in_sink"`
	OverwriteObjectsAlreadyExistingInSink bool   `json:"overwrite_objects_already_existing_in_sink"`
	OverwriteWhen                         string `json:"overwrite_when"`
}
