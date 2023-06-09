// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package athenaworkgroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Configuration struct {
	// BytesScannedCutoffPerQuery: number, optional
	BytesScannedCutoffPerQuery terra.NumberValue `hcl:"bytes_scanned_cutoff_per_query,attr"`
	// EnforceWorkgroupConfiguration: bool, optional
	EnforceWorkgroupConfiguration terra.BoolValue `hcl:"enforce_workgroup_configuration,attr"`
	// ExecutionRole: string, optional
	ExecutionRole terra.StringValue `hcl:"execution_role,attr"`
	// PublishCloudwatchMetricsEnabled: bool, optional
	PublishCloudwatchMetricsEnabled terra.BoolValue `hcl:"publish_cloudwatch_metrics_enabled,attr"`
	// RequesterPaysEnabled: bool, optional
	RequesterPaysEnabled terra.BoolValue `hcl:"requester_pays_enabled,attr"`
	// EngineVersion: optional
	EngineVersion *EngineVersion `hcl:"engine_version,block"`
	// ResultConfiguration: optional
	ResultConfiguration *ResultConfiguration `hcl:"result_configuration,block"`
}

type EngineVersion struct {
	// SelectedEngineVersion: string, optional
	SelectedEngineVersion terra.StringValue `hcl:"selected_engine_version,attr"`
}

type ResultConfiguration struct {
	// ExpectedBucketOwner: string, optional
	ExpectedBucketOwner terra.StringValue `hcl:"expected_bucket_owner,attr"`
	// OutputLocation: string, optional
	OutputLocation terra.StringValue `hcl:"output_location,attr"`
	// AclConfiguration: optional
	AclConfiguration *AclConfiguration `hcl:"acl_configuration,block"`
	// EncryptionConfiguration: optional
	EncryptionConfiguration *EncryptionConfiguration `hcl:"encryption_configuration,block"`
}

type AclConfiguration struct {
	// S3AclOption: string, required
	S3AclOption terra.StringValue `hcl:"s3_acl_option,attr" validate:"required"`
}

type EncryptionConfiguration struct {
	// EncryptionOption: string, optional
	EncryptionOption terra.StringValue `hcl:"encryption_option,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
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

func (c ConfigurationAttributes) BytesScannedCutoffPerQuery() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("bytes_scanned_cutoff_per_query"))
}

func (c ConfigurationAttributes) EnforceWorkgroupConfiguration() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("enforce_workgroup_configuration"))
}

func (c ConfigurationAttributes) ExecutionRole() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("execution_role"))
}

func (c ConfigurationAttributes) PublishCloudwatchMetricsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("publish_cloudwatch_metrics_enabled"))
}

func (c ConfigurationAttributes) RequesterPaysEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("requester_pays_enabled"))
}

func (c ConfigurationAttributes) EngineVersion() terra.ListValue[EngineVersionAttributes] {
	return terra.ReferenceAsList[EngineVersionAttributes](c.ref.Append("engine_version"))
}

func (c ConfigurationAttributes) ResultConfiguration() terra.ListValue[ResultConfigurationAttributes] {
	return terra.ReferenceAsList[ResultConfigurationAttributes](c.ref.Append("result_configuration"))
}

type EngineVersionAttributes struct {
	ref terra.Reference
}

func (ev EngineVersionAttributes) InternalRef() (terra.Reference, error) {
	return ev.ref, nil
}

func (ev EngineVersionAttributes) InternalWithRef(ref terra.Reference) EngineVersionAttributes {
	return EngineVersionAttributes{ref: ref}
}

func (ev EngineVersionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ev.ref.InternalTokens()
}

func (ev EngineVersionAttributes) EffectiveEngineVersion() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("effective_engine_version"))
}

func (ev EngineVersionAttributes) SelectedEngineVersion() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("selected_engine_version"))
}

type ResultConfigurationAttributes struct {
	ref terra.Reference
}

func (rc ResultConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc ResultConfigurationAttributes) InternalWithRef(ref terra.Reference) ResultConfigurationAttributes {
	return ResultConfigurationAttributes{ref: ref}
}

func (rc ResultConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc ResultConfigurationAttributes) ExpectedBucketOwner() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("expected_bucket_owner"))
}

func (rc ResultConfigurationAttributes) OutputLocation() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("output_location"))
}

func (rc ResultConfigurationAttributes) AclConfiguration() terra.ListValue[AclConfigurationAttributes] {
	return terra.ReferenceAsList[AclConfigurationAttributes](rc.ref.Append("acl_configuration"))
}

func (rc ResultConfigurationAttributes) EncryptionConfiguration() terra.ListValue[EncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[EncryptionConfigurationAttributes](rc.ref.Append("encryption_configuration"))
}

type AclConfigurationAttributes struct {
	ref terra.Reference
}

func (ac AclConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AclConfigurationAttributes) InternalWithRef(ref terra.Reference) AclConfigurationAttributes {
	return AclConfigurationAttributes{ref: ref}
}

func (ac AclConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AclConfigurationAttributes) S3AclOption() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("s3_acl_option"))
}

type EncryptionConfigurationAttributes struct {
	ref terra.Reference
}

func (ec EncryptionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EncryptionConfigurationAttributes) InternalWithRef(ref terra.Reference) EncryptionConfigurationAttributes {
	return EncryptionConfigurationAttributes{ref: ref}
}

func (ec EncryptionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EncryptionConfigurationAttributes) EncryptionOption() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("encryption_option"))
}

func (ec EncryptionConfigurationAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key_arn"))
}

type ConfigurationState struct {
	BytesScannedCutoffPerQuery      float64                    `json:"bytes_scanned_cutoff_per_query"`
	EnforceWorkgroupConfiguration   bool                       `json:"enforce_workgroup_configuration"`
	ExecutionRole                   string                     `json:"execution_role"`
	PublishCloudwatchMetricsEnabled bool                       `json:"publish_cloudwatch_metrics_enabled"`
	RequesterPaysEnabled            bool                       `json:"requester_pays_enabled"`
	EngineVersion                   []EngineVersionState       `json:"engine_version"`
	ResultConfiguration             []ResultConfigurationState `json:"result_configuration"`
}

type EngineVersionState struct {
	EffectiveEngineVersion string `json:"effective_engine_version"`
	SelectedEngineVersion  string `json:"selected_engine_version"`
}

type ResultConfigurationState struct {
	ExpectedBucketOwner     string                         `json:"expected_bucket_owner"`
	OutputLocation          string                         `json:"output_location"`
	AclConfiguration        []AclConfigurationState        `json:"acl_configuration"`
	EncryptionConfiguration []EncryptionConfigurationState `json:"encryption_configuration"`
}

type AclConfigurationState struct {
	S3AclOption string `json:"s3_acl_option"`
}

type EncryptionConfigurationState struct {
	EncryptionOption string `json:"encryption_option"`
	KmsKeyArn        string `json:"kms_key_arn"`
}
