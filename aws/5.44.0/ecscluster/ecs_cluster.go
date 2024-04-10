// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package ecscluster

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Configuration struct {
	// ExecuteCommandConfiguration: optional
	ExecuteCommandConfiguration *ExecuteCommandConfiguration `hcl:"execute_command_configuration,block"`
}

type ExecuteCommandConfiguration struct {
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Logging: string, optional
	Logging terra.StringValue `hcl:"logging,attr"`
	// LogConfiguration: optional
	LogConfiguration *LogConfiguration `hcl:"log_configuration,block"`
}

type LogConfiguration struct {
	// CloudWatchEncryptionEnabled: bool, optional
	CloudWatchEncryptionEnabled terra.BoolValue `hcl:"cloud_watch_encryption_enabled,attr"`
	// CloudWatchLogGroupName: string, optional
	CloudWatchLogGroupName terra.StringValue `hcl:"cloud_watch_log_group_name,attr"`
	// S3BucketEncryptionEnabled: bool, optional
	S3BucketEncryptionEnabled terra.BoolValue `hcl:"s3_bucket_encryption_enabled,attr"`
	// S3BucketName: string, optional
	S3BucketName terra.StringValue `hcl:"s3_bucket_name,attr"`
	// S3KeyPrefix: string, optional
	S3KeyPrefix terra.StringValue `hcl:"s3_key_prefix,attr"`
}

type ServiceConnectDefaults struct {
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
}

type Setting struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type ConfigurationAttributes struct {
	ref terra.Reference
}

func (c ConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigurationAttributes) InternalWithRef(ref terra.Reference) ConfigurationAttributes {
	return ConfigurationAttributes{ref: ref}
}

func (c ConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigurationAttributes) ExecuteCommandConfiguration() terra.ListValue[ExecuteCommandConfigurationAttributes] {
	return terra.ReferenceAsList[ExecuteCommandConfigurationAttributes](c.ref.Append("execute_command_configuration"))
}

type ExecuteCommandConfigurationAttributes struct {
	ref terra.Reference
}

func (ecc ExecuteCommandConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ecc.ref, nil
}

func (ecc ExecuteCommandConfigurationAttributes) InternalWithRef(ref terra.Reference) ExecuteCommandConfigurationAttributes {
	return ExecuteCommandConfigurationAttributes{ref: ref}
}

func (ecc ExecuteCommandConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ecc.ref.InternalTokens()
}

func (ecc ExecuteCommandConfigurationAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ecc.ref.Append("kms_key_id"))
}

func (ecc ExecuteCommandConfigurationAttributes) Logging() terra.StringValue {
	return terra.ReferenceAsString(ecc.ref.Append("logging"))
}

func (ecc ExecuteCommandConfigurationAttributes) LogConfiguration() terra.ListValue[LogConfigurationAttributes] {
	return terra.ReferenceAsList[LogConfigurationAttributes](ecc.ref.Append("log_configuration"))
}

type LogConfigurationAttributes struct {
	ref terra.Reference
}

func (lc LogConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc LogConfigurationAttributes) InternalWithRef(ref terra.Reference) LogConfigurationAttributes {
	return LogConfigurationAttributes{ref: ref}
}

func (lc LogConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc LogConfigurationAttributes) CloudWatchEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lc.ref.Append("cloud_watch_encryption_enabled"))
}

func (lc LogConfigurationAttributes) CloudWatchLogGroupName() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("cloud_watch_log_group_name"))
}

func (lc LogConfigurationAttributes) S3BucketEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lc.ref.Append("s3_bucket_encryption_enabled"))
}

func (lc LogConfigurationAttributes) S3BucketName() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("s3_bucket_name"))
}

func (lc LogConfigurationAttributes) S3KeyPrefix() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("s3_key_prefix"))
}

type ServiceConnectDefaultsAttributes struct {
	ref terra.Reference
}

func (scd ServiceConnectDefaultsAttributes) InternalRef() (terra.Reference, error) {
	return scd.ref, nil
}

func (scd ServiceConnectDefaultsAttributes) InternalWithRef(ref terra.Reference) ServiceConnectDefaultsAttributes {
	return ServiceConnectDefaultsAttributes{ref: ref}
}

func (scd ServiceConnectDefaultsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return scd.ref.InternalTokens()
}

func (scd ServiceConnectDefaultsAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(scd.ref.Append("namespace"))
}

type SettingAttributes struct {
	ref terra.Reference
}

func (s SettingAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SettingAttributes) InternalWithRef(ref terra.Reference) SettingAttributes {
	return SettingAttributes{ref: ref}
}

func (s SettingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SettingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SettingAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("value"))
}

type ConfigurationState struct {
	ExecuteCommandConfiguration []ExecuteCommandConfigurationState `json:"execute_command_configuration"`
}

type ExecuteCommandConfigurationState struct {
	KmsKeyId         string                  `json:"kms_key_id"`
	Logging          string                  `json:"logging"`
	LogConfiguration []LogConfigurationState `json:"log_configuration"`
}

type LogConfigurationState struct {
	CloudWatchEncryptionEnabled bool   `json:"cloud_watch_encryption_enabled"`
	CloudWatchLogGroupName      string `json:"cloud_watch_log_group_name"`
	S3BucketEncryptionEnabled   bool   `json:"s3_bucket_encryption_enabled"`
	S3BucketName                string `json:"s3_bucket_name"`
	S3KeyPrefix                 string `json:"s3_key_prefix"`
}

type ServiceConnectDefaultsState struct {
	Namespace string `json:"namespace"`
}

type SettingState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
