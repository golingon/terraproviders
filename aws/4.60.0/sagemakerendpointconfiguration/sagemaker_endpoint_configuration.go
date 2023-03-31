// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sagemakerendpointconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AsyncInferenceConfig struct {
	// ClientConfig: optional
	ClientConfig *ClientConfig `hcl:"client_config,block"`
	// OutputConfig: required
	OutputConfig *OutputConfig `hcl:"output_config,block" validate:"required"`
}

type ClientConfig struct {
	// MaxConcurrentInvocationsPerInstance: number, optional
	MaxConcurrentInvocationsPerInstance terra.NumberValue `hcl:"max_concurrent_invocations_per_instance,attr"`
}

type OutputConfig struct {
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// S3OutputPath: string, required
	S3OutputPath terra.StringValue `hcl:"s3_output_path,attr" validate:"required"`
	// NotificationConfig: optional
	NotificationConfig *NotificationConfig `hcl:"notification_config,block"`
}

type NotificationConfig struct {
	// ErrorTopic: string, optional
	ErrorTopic terra.StringValue `hcl:"error_topic,attr"`
	// SuccessTopic: string, optional
	SuccessTopic terra.StringValue `hcl:"success_topic,attr"`
}

type DataCaptureConfig struct {
	// DestinationS3Uri: string, required
	DestinationS3Uri terra.StringValue `hcl:"destination_s3_uri,attr" validate:"required"`
	// EnableCapture: bool, optional
	EnableCapture terra.BoolValue `hcl:"enable_capture,attr"`
	// InitialSamplingPercentage: number, required
	InitialSamplingPercentage terra.NumberValue `hcl:"initial_sampling_percentage,attr" validate:"required"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// CaptureContentTypeHeader: optional
	CaptureContentTypeHeader *CaptureContentTypeHeader `hcl:"capture_content_type_header,block"`
	// CaptureOptions: min=1,max=2
	CaptureOptions []CaptureOptions `hcl:"capture_options,block" validate:"min=1,max=2"`
}

type CaptureContentTypeHeader struct {
	// CsvContentTypes: set of string, optional
	CsvContentTypes terra.SetValue[terra.StringValue] `hcl:"csv_content_types,attr"`
	// JsonContentTypes: set of string, optional
	JsonContentTypes terra.SetValue[terra.StringValue] `hcl:"json_content_types,attr"`
}

type CaptureOptions struct {
	// CaptureMode: string, required
	CaptureMode terra.StringValue `hcl:"capture_mode,attr" validate:"required"`
}

type ProductionVariants struct {
	// AcceleratorType: string, optional
	AcceleratorType terra.StringValue `hcl:"accelerator_type,attr"`
	// ContainerStartupHealthCheckTimeoutInSeconds: number, optional
	ContainerStartupHealthCheckTimeoutInSeconds terra.NumberValue `hcl:"container_startup_health_check_timeout_in_seconds,attr"`
	// InitialInstanceCount: number, optional
	InitialInstanceCount terra.NumberValue `hcl:"initial_instance_count,attr"`
	// InitialVariantWeight: number, optional
	InitialVariantWeight terra.NumberValue `hcl:"initial_variant_weight,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// ModelDataDownloadTimeoutInSeconds: number, optional
	ModelDataDownloadTimeoutInSeconds terra.NumberValue `hcl:"model_data_download_timeout_in_seconds,attr"`
	// ModelName: string, required
	ModelName terra.StringValue `hcl:"model_name,attr" validate:"required"`
	// VariantName: string, optional
	VariantName terra.StringValue `hcl:"variant_name,attr"`
	// VolumeSizeInGb: number, optional
	VolumeSizeInGb terra.NumberValue `hcl:"volume_size_in_gb,attr"`
	// ProductionVariantsCoreDumpConfig: optional
	CoreDumpConfig *ProductionVariantsCoreDumpConfig `hcl:"core_dump_config,block"`
	// ProductionVariantsServerlessConfig: optional
	ServerlessConfig *ProductionVariantsServerlessConfig `hcl:"serverless_config,block"`
}

type ProductionVariantsCoreDumpConfig struct {
	// DestinationS3Uri: string, required
	DestinationS3Uri terra.StringValue `hcl:"destination_s3_uri,attr" validate:"required"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
}

type ProductionVariantsServerlessConfig struct {
	// MaxConcurrency: number, required
	MaxConcurrency terra.NumberValue `hcl:"max_concurrency,attr" validate:"required"`
	// MemorySizeInMb: number, required
	MemorySizeInMb terra.NumberValue `hcl:"memory_size_in_mb,attr" validate:"required"`
}

type ShadowProductionVariants struct {
	// AcceleratorType: string, optional
	AcceleratorType terra.StringValue `hcl:"accelerator_type,attr"`
	// ContainerStartupHealthCheckTimeoutInSeconds: number, optional
	ContainerStartupHealthCheckTimeoutInSeconds terra.NumberValue `hcl:"container_startup_health_check_timeout_in_seconds,attr"`
	// InitialInstanceCount: number, optional
	InitialInstanceCount terra.NumberValue `hcl:"initial_instance_count,attr"`
	// InitialVariantWeight: number, optional
	InitialVariantWeight terra.NumberValue `hcl:"initial_variant_weight,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// ModelDataDownloadTimeoutInSeconds: number, optional
	ModelDataDownloadTimeoutInSeconds terra.NumberValue `hcl:"model_data_download_timeout_in_seconds,attr"`
	// ModelName: string, required
	ModelName terra.StringValue `hcl:"model_name,attr" validate:"required"`
	// VariantName: string, optional
	VariantName terra.StringValue `hcl:"variant_name,attr"`
	// VolumeSizeInGb: number, optional
	VolumeSizeInGb terra.NumberValue `hcl:"volume_size_in_gb,attr"`
	// ShadowProductionVariantsCoreDumpConfig: optional
	CoreDumpConfig *ShadowProductionVariantsCoreDumpConfig `hcl:"core_dump_config,block"`
	// ShadowProductionVariantsServerlessConfig: optional
	ServerlessConfig *ShadowProductionVariantsServerlessConfig `hcl:"serverless_config,block"`
}

type ShadowProductionVariantsCoreDumpConfig struct {
	// DestinationS3Uri: string, required
	DestinationS3Uri terra.StringValue `hcl:"destination_s3_uri,attr" validate:"required"`
	// KmsKeyId: string, required
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr" validate:"required"`
}

type ShadowProductionVariantsServerlessConfig struct {
	// MaxConcurrency: number, required
	MaxConcurrency terra.NumberValue `hcl:"max_concurrency,attr" validate:"required"`
	// MemorySizeInMb: number, required
	MemorySizeInMb terra.NumberValue `hcl:"memory_size_in_mb,attr" validate:"required"`
}

type AsyncInferenceConfigAttributes struct {
	ref terra.Reference
}

func (aic AsyncInferenceConfigAttributes) InternalRef() terra.Reference {
	return aic.ref
}

func (aic AsyncInferenceConfigAttributes) InternalWithRef(ref terra.Reference) AsyncInferenceConfigAttributes {
	return AsyncInferenceConfigAttributes{ref: ref}
}

func (aic AsyncInferenceConfigAttributes) InternalTokens() hclwrite.Tokens {
	return aic.ref.InternalTokens()
}

func (aic AsyncInferenceConfigAttributes) ClientConfig() terra.ListValue[ClientConfigAttributes] {
	return terra.ReferenceAsList[ClientConfigAttributes](aic.ref.Append("client_config"))
}

func (aic AsyncInferenceConfigAttributes) OutputConfig() terra.ListValue[OutputConfigAttributes] {
	return terra.ReferenceAsList[OutputConfigAttributes](aic.ref.Append("output_config"))
}

type ClientConfigAttributes struct {
	ref terra.Reference
}

func (cc ClientConfigAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc ClientConfigAttributes) InternalWithRef(ref terra.Reference) ClientConfigAttributes {
	return ClientConfigAttributes{ref: ref}
}

func (cc ClientConfigAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc ClientConfigAttributes) MaxConcurrentInvocationsPerInstance() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("max_concurrent_invocations_per_instance"))
}

type OutputConfigAttributes struct {
	ref terra.Reference
}

func (oc OutputConfigAttributes) InternalRef() terra.Reference {
	return oc.ref
}

func (oc OutputConfigAttributes) InternalWithRef(ref terra.Reference) OutputConfigAttributes {
	return OutputConfigAttributes{ref: ref}
}

func (oc OutputConfigAttributes) InternalTokens() hclwrite.Tokens {
	return oc.ref.InternalTokens()
}

func (oc OutputConfigAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("kms_key_id"))
}

func (oc OutputConfigAttributes) S3OutputPath() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("s3_output_path"))
}

func (oc OutputConfigAttributes) NotificationConfig() terra.ListValue[NotificationConfigAttributes] {
	return terra.ReferenceAsList[NotificationConfigAttributes](oc.ref.Append("notification_config"))
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

func (nc NotificationConfigAttributes) ErrorTopic() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("error_topic"))
}

func (nc NotificationConfigAttributes) SuccessTopic() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("success_topic"))
}

type DataCaptureConfigAttributes struct {
	ref terra.Reference
}

func (dcc DataCaptureConfigAttributes) InternalRef() terra.Reference {
	return dcc.ref
}

func (dcc DataCaptureConfigAttributes) InternalWithRef(ref terra.Reference) DataCaptureConfigAttributes {
	return DataCaptureConfigAttributes{ref: ref}
}

func (dcc DataCaptureConfigAttributes) InternalTokens() hclwrite.Tokens {
	return dcc.ref.InternalTokens()
}

func (dcc DataCaptureConfigAttributes) DestinationS3Uri() terra.StringValue {
	return terra.ReferenceAsString(dcc.ref.Append("destination_s3_uri"))
}

func (dcc DataCaptureConfigAttributes) EnableCapture() terra.BoolValue {
	return terra.ReferenceAsBool(dcc.ref.Append("enable_capture"))
}

func (dcc DataCaptureConfigAttributes) InitialSamplingPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(dcc.ref.Append("initial_sampling_percentage"))
}

func (dcc DataCaptureConfigAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(dcc.ref.Append("kms_key_id"))
}

func (dcc DataCaptureConfigAttributes) CaptureContentTypeHeader() terra.ListValue[CaptureContentTypeHeaderAttributes] {
	return terra.ReferenceAsList[CaptureContentTypeHeaderAttributes](dcc.ref.Append("capture_content_type_header"))
}

func (dcc DataCaptureConfigAttributes) CaptureOptions() terra.ListValue[CaptureOptionsAttributes] {
	return terra.ReferenceAsList[CaptureOptionsAttributes](dcc.ref.Append("capture_options"))
}

type CaptureContentTypeHeaderAttributes struct {
	ref terra.Reference
}

func (ccth CaptureContentTypeHeaderAttributes) InternalRef() terra.Reference {
	return ccth.ref
}

func (ccth CaptureContentTypeHeaderAttributes) InternalWithRef(ref terra.Reference) CaptureContentTypeHeaderAttributes {
	return CaptureContentTypeHeaderAttributes{ref: ref}
}

func (ccth CaptureContentTypeHeaderAttributes) InternalTokens() hclwrite.Tokens {
	return ccth.ref.InternalTokens()
}

func (ccth CaptureContentTypeHeaderAttributes) CsvContentTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ccth.ref.Append("csv_content_types"))
}

func (ccth CaptureContentTypeHeaderAttributes) JsonContentTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ccth.ref.Append("json_content_types"))
}

type CaptureOptionsAttributes struct {
	ref terra.Reference
}

func (co CaptureOptionsAttributes) InternalRef() terra.Reference {
	return co.ref
}

func (co CaptureOptionsAttributes) InternalWithRef(ref terra.Reference) CaptureOptionsAttributes {
	return CaptureOptionsAttributes{ref: ref}
}

func (co CaptureOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return co.ref.InternalTokens()
}

func (co CaptureOptionsAttributes) CaptureMode() terra.StringValue {
	return terra.ReferenceAsString(co.ref.Append("capture_mode"))
}

type ProductionVariantsAttributes struct {
	ref terra.Reference
}

func (pv ProductionVariantsAttributes) InternalRef() terra.Reference {
	return pv.ref
}

func (pv ProductionVariantsAttributes) InternalWithRef(ref terra.Reference) ProductionVariantsAttributes {
	return ProductionVariantsAttributes{ref: ref}
}

func (pv ProductionVariantsAttributes) InternalTokens() hclwrite.Tokens {
	return pv.ref.InternalTokens()
}

func (pv ProductionVariantsAttributes) AcceleratorType() terra.StringValue {
	return terra.ReferenceAsString(pv.ref.Append("accelerator_type"))
}

func (pv ProductionVariantsAttributes) ContainerStartupHealthCheckTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(pv.ref.Append("container_startup_health_check_timeout_in_seconds"))
}

func (pv ProductionVariantsAttributes) InitialInstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(pv.ref.Append("initial_instance_count"))
}

func (pv ProductionVariantsAttributes) InitialVariantWeight() terra.NumberValue {
	return terra.ReferenceAsNumber(pv.ref.Append("initial_variant_weight"))
}

func (pv ProductionVariantsAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(pv.ref.Append("instance_type"))
}

func (pv ProductionVariantsAttributes) ModelDataDownloadTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(pv.ref.Append("model_data_download_timeout_in_seconds"))
}

func (pv ProductionVariantsAttributes) ModelName() terra.StringValue {
	return terra.ReferenceAsString(pv.ref.Append("model_name"))
}

func (pv ProductionVariantsAttributes) VariantName() terra.StringValue {
	return terra.ReferenceAsString(pv.ref.Append("variant_name"))
}

func (pv ProductionVariantsAttributes) VolumeSizeInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(pv.ref.Append("volume_size_in_gb"))
}

func (pv ProductionVariantsAttributes) CoreDumpConfig() terra.ListValue[ProductionVariantsCoreDumpConfigAttributes] {
	return terra.ReferenceAsList[ProductionVariantsCoreDumpConfigAttributes](pv.ref.Append("core_dump_config"))
}

func (pv ProductionVariantsAttributes) ServerlessConfig() terra.ListValue[ProductionVariantsServerlessConfigAttributes] {
	return terra.ReferenceAsList[ProductionVariantsServerlessConfigAttributes](pv.ref.Append("serverless_config"))
}

type ProductionVariantsCoreDumpConfigAttributes struct {
	ref terra.Reference
}

func (cdc ProductionVariantsCoreDumpConfigAttributes) InternalRef() terra.Reference {
	return cdc.ref
}

func (cdc ProductionVariantsCoreDumpConfigAttributes) InternalWithRef(ref terra.Reference) ProductionVariantsCoreDumpConfigAttributes {
	return ProductionVariantsCoreDumpConfigAttributes{ref: ref}
}

func (cdc ProductionVariantsCoreDumpConfigAttributes) InternalTokens() hclwrite.Tokens {
	return cdc.ref.InternalTokens()
}

func (cdc ProductionVariantsCoreDumpConfigAttributes) DestinationS3Uri() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("destination_s3_uri"))
}

func (cdc ProductionVariantsCoreDumpConfigAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("kms_key_id"))
}

type ProductionVariantsServerlessConfigAttributes struct {
	ref terra.Reference
}

func (sc ProductionVariantsServerlessConfigAttributes) InternalRef() terra.Reference {
	return sc.ref
}

func (sc ProductionVariantsServerlessConfigAttributes) InternalWithRef(ref terra.Reference) ProductionVariantsServerlessConfigAttributes {
	return ProductionVariantsServerlessConfigAttributes{ref: ref}
}

func (sc ProductionVariantsServerlessConfigAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc ProductionVariantsServerlessConfigAttributes) MaxConcurrency() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("max_concurrency"))
}

func (sc ProductionVariantsServerlessConfigAttributes) MemorySizeInMb() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("memory_size_in_mb"))
}

type ShadowProductionVariantsAttributes struct {
	ref terra.Reference
}

func (spv ShadowProductionVariantsAttributes) InternalRef() terra.Reference {
	return spv.ref
}

func (spv ShadowProductionVariantsAttributes) InternalWithRef(ref terra.Reference) ShadowProductionVariantsAttributes {
	return ShadowProductionVariantsAttributes{ref: ref}
}

func (spv ShadowProductionVariantsAttributes) InternalTokens() hclwrite.Tokens {
	return spv.ref.InternalTokens()
}

func (spv ShadowProductionVariantsAttributes) AcceleratorType() terra.StringValue {
	return terra.ReferenceAsString(spv.ref.Append("accelerator_type"))
}

func (spv ShadowProductionVariantsAttributes) ContainerStartupHealthCheckTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(spv.ref.Append("container_startup_health_check_timeout_in_seconds"))
}

func (spv ShadowProductionVariantsAttributes) InitialInstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(spv.ref.Append("initial_instance_count"))
}

func (spv ShadowProductionVariantsAttributes) InitialVariantWeight() terra.NumberValue {
	return terra.ReferenceAsNumber(spv.ref.Append("initial_variant_weight"))
}

func (spv ShadowProductionVariantsAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(spv.ref.Append("instance_type"))
}

func (spv ShadowProductionVariantsAttributes) ModelDataDownloadTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(spv.ref.Append("model_data_download_timeout_in_seconds"))
}

func (spv ShadowProductionVariantsAttributes) ModelName() terra.StringValue {
	return terra.ReferenceAsString(spv.ref.Append("model_name"))
}

func (spv ShadowProductionVariantsAttributes) VariantName() terra.StringValue {
	return terra.ReferenceAsString(spv.ref.Append("variant_name"))
}

func (spv ShadowProductionVariantsAttributes) VolumeSizeInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(spv.ref.Append("volume_size_in_gb"))
}

func (spv ShadowProductionVariantsAttributes) CoreDumpConfig() terra.ListValue[ShadowProductionVariantsCoreDumpConfigAttributes] {
	return terra.ReferenceAsList[ShadowProductionVariantsCoreDumpConfigAttributes](spv.ref.Append("core_dump_config"))
}

func (spv ShadowProductionVariantsAttributes) ServerlessConfig() terra.ListValue[ShadowProductionVariantsServerlessConfigAttributes] {
	return terra.ReferenceAsList[ShadowProductionVariantsServerlessConfigAttributes](spv.ref.Append("serverless_config"))
}

type ShadowProductionVariantsCoreDumpConfigAttributes struct {
	ref terra.Reference
}

func (cdc ShadowProductionVariantsCoreDumpConfigAttributes) InternalRef() terra.Reference {
	return cdc.ref
}

func (cdc ShadowProductionVariantsCoreDumpConfigAttributes) InternalWithRef(ref terra.Reference) ShadowProductionVariantsCoreDumpConfigAttributes {
	return ShadowProductionVariantsCoreDumpConfigAttributes{ref: ref}
}

func (cdc ShadowProductionVariantsCoreDumpConfigAttributes) InternalTokens() hclwrite.Tokens {
	return cdc.ref.InternalTokens()
}

func (cdc ShadowProductionVariantsCoreDumpConfigAttributes) DestinationS3Uri() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("destination_s3_uri"))
}

func (cdc ShadowProductionVariantsCoreDumpConfigAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("kms_key_id"))
}

type ShadowProductionVariantsServerlessConfigAttributes struct {
	ref terra.Reference
}

func (sc ShadowProductionVariantsServerlessConfigAttributes) InternalRef() terra.Reference {
	return sc.ref
}

func (sc ShadowProductionVariantsServerlessConfigAttributes) InternalWithRef(ref terra.Reference) ShadowProductionVariantsServerlessConfigAttributes {
	return ShadowProductionVariantsServerlessConfigAttributes{ref: ref}
}

func (sc ShadowProductionVariantsServerlessConfigAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc ShadowProductionVariantsServerlessConfigAttributes) MaxConcurrency() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("max_concurrency"))
}

func (sc ShadowProductionVariantsServerlessConfigAttributes) MemorySizeInMb() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("memory_size_in_mb"))
}

type AsyncInferenceConfigState struct {
	ClientConfig []ClientConfigState `json:"client_config"`
	OutputConfig []OutputConfigState `json:"output_config"`
}

type ClientConfigState struct {
	MaxConcurrentInvocationsPerInstance float64 `json:"max_concurrent_invocations_per_instance"`
}

type OutputConfigState struct {
	KmsKeyId           string                    `json:"kms_key_id"`
	S3OutputPath       string                    `json:"s3_output_path"`
	NotificationConfig []NotificationConfigState `json:"notification_config"`
}

type NotificationConfigState struct {
	ErrorTopic   string `json:"error_topic"`
	SuccessTopic string `json:"success_topic"`
}

type DataCaptureConfigState struct {
	DestinationS3Uri          string                          `json:"destination_s3_uri"`
	EnableCapture             bool                            `json:"enable_capture"`
	InitialSamplingPercentage float64                         `json:"initial_sampling_percentage"`
	KmsKeyId                  string                          `json:"kms_key_id"`
	CaptureContentTypeHeader  []CaptureContentTypeHeaderState `json:"capture_content_type_header"`
	CaptureOptions            []CaptureOptionsState           `json:"capture_options"`
}

type CaptureContentTypeHeaderState struct {
	CsvContentTypes  []string `json:"csv_content_types"`
	JsonContentTypes []string `json:"json_content_types"`
}

type CaptureOptionsState struct {
	CaptureMode string `json:"capture_mode"`
}

type ProductionVariantsState struct {
	AcceleratorType                             string                                    `json:"accelerator_type"`
	ContainerStartupHealthCheckTimeoutInSeconds float64                                   `json:"container_startup_health_check_timeout_in_seconds"`
	InitialInstanceCount                        float64                                   `json:"initial_instance_count"`
	InitialVariantWeight                        float64                                   `json:"initial_variant_weight"`
	InstanceType                                string                                    `json:"instance_type"`
	ModelDataDownloadTimeoutInSeconds           float64                                   `json:"model_data_download_timeout_in_seconds"`
	ModelName                                   string                                    `json:"model_name"`
	VariantName                                 string                                    `json:"variant_name"`
	VolumeSizeInGb                              float64                                   `json:"volume_size_in_gb"`
	CoreDumpConfig                              []ProductionVariantsCoreDumpConfigState   `json:"core_dump_config"`
	ServerlessConfig                            []ProductionVariantsServerlessConfigState `json:"serverless_config"`
}

type ProductionVariantsCoreDumpConfigState struct {
	DestinationS3Uri string `json:"destination_s3_uri"`
	KmsKeyId         string `json:"kms_key_id"`
}

type ProductionVariantsServerlessConfigState struct {
	MaxConcurrency float64 `json:"max_concurrency"`
	MemorySizeInMb float64 `json:"memory_size_in_mb"`
}

type ShadowProductionVariantsState struct {
	AcceleratorType                             string                                          `json:"accelerator_type"`
	ContainerStartupHealthCheckTimeoutInSeconds float64                                         `json:"container_startup_health_check_timeout_in_seconds"`
	InitialInstanceCount                        float64                                         `json:"initial_instance_count"`
	InitialVariantWeight                        float64                                         `json:"initial_variant_weight"`
	InstanceType                                string                                          `json:"instance_type"`
	ModelDataDownloadTimeoutInSeconds           float64                                         `json:"model_data_download_timeout_in_seconds"`
	ModelName                                   string                                          `json:"model_name"`
	VariantName                                 string                                          `json:"variant_name"`
	VolumeSizeInGb                              float64                                         `json:"volume_size_in_gb"`
	CoreDumpConfig                              []ShadowProductionVariantsCoreDumpConfigState   `json:"core_dump_config"`
	ServerlessConfig                            []ShadowProductionVariantsServerlessConfigState `json:"serverless_config"`
}

type ShadowProductionVariantsCoreDumpConfigState struct {
	DestinationS3Uri string `json:"destination_s3_uri"`
	KmsKeyId         string `json:"kms_key_id"`
}

type ShadowProductionVariantsServerlessConfigState struct {
	MaxConcurrency float64 `json:"max_concurrency"`
	MemorySizeInMb float64 `json:"memory_size_in_mb"`
}
