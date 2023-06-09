// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apprunnerservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EncryptionConfiguration struct {
	// KmsKey: string, required
	KmsKey terra.StringValue `hcl:"kms_key,attr" validate:"required"`
}

type HealthCheckConfiguration struct {
	// HealthyThreshold: number, optional
	HealthyThreshold terra.NumberValue `hcl:"healthy_threshold,attr"`
	// Interval: number, optional
	Interval terra.NumberValue `hcl:"interval,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// Timeout: number, optional
	Timeout terra.NumberValue `hcl:"timeout,attr"`
	// UnhealthyThreshold: number, optional
	UnhealthyThreshold terra.NumberValue `hcl:"unhealthy_threshold,attr"`
}

type InstanceConfiguration struct {
	// Cpu: string, optional
	Cpu terra.StringValue `hcl:"cpu,attr"`
	// InstanceRoleArn: string, optional
	InstanceRoleArn terra.StringValue `hcl:"instance_role_arn,attr"`
	// Memory: string, optional
	Memory terra.StringValue `hcl:"memory,attr"`
}

type NetworkConfiguration struct {
	// EgressConfiguration: optional
	EgressConfiguration *EgressConfiguration `hcl:"egress_configuration,block"`
	// IngressConfiguration: optional
	IngressConfiguration *IngressConfiguration `hcl:"ingress_configuration,block"`
}

type EgressConfiguration struct {
	// EgressType: string, optional
	EgressType terra.StringValue `hcl:"egress_type,attr"`
	// VpcConnectorArn: string, optional
	VpcConnectorArn terra.StringValue `hcl:"vpc_connector_arn,attr"`
}

type IngressConfiguration struct {
	// IsPubliclyAccessible: bool, optional
	IsPubliclyAccessible terra.BoolValue `hcl:"is_publicly_accessible,attr"`
}

type ObservabilityConfiguration struct {
	// ObservabilityConfigurationArn: string, optional
	ObservabilityConfigurationArn terra.StringValue `hcl:"observability_configuration_arn,attr"`
	// ObservabilityEnabled: bool, required
	ObservabilityEnabled terra.BoolValue `hcl:"observability_enabled,attr" validate:"required"`
}

type SourceConfiguration struct {
	// AutoDeploymentsEnabled: bool, optional
	AutoDeploymentsEnabled terra.BoolValue `hcl:"auto_deployments_enabled,attr"`
	// AuthenticationConfiguration: optional
	AuthenticationConfiguration *AuthenticationConfiguration `hcl:"authentication_configuration,block"`
	// CodeRepository: optional
	CodeRepository *CodeRepository `hcl:"code_repository,block"`
	// ImageRepository: optional
	ImageRepository *ImageRepository `hcl:"image_repository,block"`
}

type AuthenticationConfiguration struct {
	// AccessRoleArn: string, optional
	AccessRoleArn terra.StringValue `hcl:"access_role_arn,attr"`
	// ConnectionArn: string, optional
	ConnectionArn terra.StringValue `hcl:"connection_arn,attr"`
}

type CodeRepository struct {
	// RepositoryUrl: string, required
	RepositoryUrl terra.StringValue `hcl:"repository_url,attr" validate:"required"`
	// CodeConfiguration: optional
	CodeConfiguration *CodeConfiguration `hcl:"code_configuration,block"`
	// SourceCodeVersion: required
	SourceCodeVersion *SourceCodeVersion `hcl:"source_code_version,block" validate:"required"`
}

type CodeConfiguration struct {
	// ConfigurationSource: string, required
	ConfigurationSource terra.StringValue `hcl:"configuration_source,attr" validate:"required"`
	// CodeConfigurationValues: optional
	CodeConfigurationValues *CodeConfigurationValues `hcl:"code_configuration_values,block"`
}

type CodeConfigurationValues struct {
	// BuildCommand: string, optional
	BuildCommand terra.StringValue `hcl:"build_command,attr"`
	// Port: string, optional
	Port terra.StringValue `hcl:"port,attr"`
	// Runtime: string, required
	Runtime terra.StringValue `hcl:"runtime,attr" validate:"required"`
	// RuntimeEnvironmentSecrets: map of string, optional
	RuntimeEnvironmentSecrets terra.MapValue[terra.StringValue] `hcl:"runtime_environment_secrets,attr"`
	// RuntimeEnvironmentVariables: map of string, optional
	RuntimeEnvironmentVariables terra.MapValue[terra.StringValue] `hcl:"runtime_environment_variables,attr"`
	// StartCommand: string, optional
	StartCommand terra.StringValue `hcl:"start_command,attr"`
}

type SourceCodeVersion struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type ImageRepository struct {
	// ImageIdentifier: string, required
	ImageIdentifier terra.StringValue `hcl:"image_identifier,attr" validate:"required"`
	// ImageRepositoryType: string, required
	ImageRepositoryType terra.StringValue `hcl:"image_repository_type,attr" validate:"required"`
	// ImageConfiguration: optional
	ImageConfiguration *ImageConfiguration `hcl:"image_configuration,block"`
}

type ImageConfiguration struct {
	// Port: string, optional
	Port terra.StringValue `hcl:"port,attr"`
	// RuntimeEnvironmentSecrets: map of string, optional
	RuntimeEnvironmentSecrets terra.MapValue[terra.StringValue] `hcl:"runtime_environment_secrets,attr"`
	// RuntimeEnvironmentVariables: map of string, optional
	RuntimeEnvironmentVariables terra.MapValue[terra.StringValue] `hcl:"runtime_environment_variables,attr"`
	// StartCommand: string, optional
	StartCommand terra.StringValue `hcl:"start_command,attr"`
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

func (ec EncryptionConfigurationAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key"))
}

type HealthCheckConfigurationAttributes struct {
	ref terra.Reference
}

func (hcc HealthCheckConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return hcc.ref, nil
}

func (hcc HealthCheckConfigurationAttributes) InternalWithRef(ref terra.Reference) HealthCheckConfigurationAttributes {
	return HealthCheckConfigurationAttributes{ref: ref}
}

func (hcc HealthCheckConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hcc.ref.InternalTokens()
}

func (hcc HealthCheckConfigurationAttributes) HealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hcc.ref.Append("healthy_threshold"))
}

func (hcc HealthCheckConfigurationAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(hcc.ref.Append("interval"))
}

func (hcc HealthCheckConfigurationAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(hcc.ref.Append("path"))
}

func (hcc HealthCheckConfigurationAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(hcc.ref.Append("protocol"))
}

func (hcc HealthCheckConfigurationAttributes) Timeout() terra.NumberValue {
	return terra.ReferenceAsNumber(hcc.ref.Append("timeout"))
}

func (hcc HealthCheckConfigurationAttributes) UnhealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hcc.ref.Append("unhealthy_threshold"))
}

type InstanceConfigurationAttributes struct {
	ref terra.Reference
}

func (ic InstanceConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic InstanceConfigurationAttributes) InternalWithRef(ref terra.Reference) InstanceConfigurationAttributes {
	return InstanceConfigurationAttributes{ref: ref}
}

func (ic InstanceConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic InstanceConfigurationAttributes) Cpu() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("cpu"))
}

func (ic InstanceConfigurationAttributes) InstanceRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("instance_role_arn"))
}

func (ic InstanceConfigurationAttributes) Memory() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("memory"))
}

type NetworkConfigurationAttributes struct {
	ref terra.Reference
}

func (nc NetworkConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NetworkConfigurationAttributes) InternalWithRef(ref terra.Reference) NetworkConfigurationAttributes {
	return NetworkConfigurationAttributes{ref: ref}
}

func (nc NetworkConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NetworkConfigurationAttributes) EgressConfiguration() terra.ListValue[EgressConfigurationAttributes] {
	return terra.ReferenceAsList[EgressConfigurationAttributes](nc.ref.Append("egress_configuration"))
}

func (nc NetworkConfigurationAttributes) IngressConfiguration() terra.ListValue[IngressConfigurationAttributes] {
	return terra.ReferenceAsList[IngressConfigurationAttributes](nc.ref.Append("ingress_configuration"))
}

type EgressConfigurationAttributes struct {
	ref terra.Reference
}

func (ec EgressConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EgressConfigurationAttributes) InternalWithRef(ref terra.Reference) EgressConfigurationAttributes {
	return EgressConfigurationAttributes{ref: ref}
}

func (ec EgressConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EgressConfigurationAttributes) EgressType() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("egress_type"))
}

func (ec EgressConfigurationAttributes) VpcConnectorArn() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("vpc_connector_arn"))
}

type IngressConfigurationAttributes struct {
	ref terra.Reference
}

func (ic IngressConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic IngressConfigurationAttributes) InternalWithRef(ref terra.Reference) IngressConfigurationAttributes {
	return IngressConfigurationAttributes{ref: ref}
}

func (ic IngressConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic IngressConfigurationAttributes) IsPubliclyAccessible() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("is_publicly_accessible"))
}

type ObservabilityConfigurationAttributes struct {
	ref terra.Reference
}

func (oc ObservabilityConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return oc.ref, nil
}

func (oc ObservabilityConfigurationAttributes) InternalWithRef(ref terra.Reference) ObservabilityConfigurationAttributes {
	return ObservabilityConfigurationAttributes{ref: ref}
}

func (oc ObservabilityConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oc.ref.InternalTokens()
}

func (oc ObservabilityConfigurationAttributes) ObservabilityConfigurationArn() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("observability_configuration_arn"))
}

func (oc ObservabilityConfigurationAttributes) ObservabilityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(oc.ref.Append("observability_enabled"))
}

type SourceConfigurationAttributes struct {
	ref terra.Reference
}

func (sc SourceConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SourceConfigurationAttributes) InternalWithRef(ref terra.Reference) SourceConfigurationAttributes {
	return SourceConfigurationAttributes{ref: ref}
}

func (sc SourceConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SourceConfigurationAttributes) AutoDeploymentsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("auto_deployments_enabled"))
}

func (sc SourceConfigurationAttributes) AuthenticationConfiguration() terra.ListValue[AuthenticationConfigurationAttributes] {
	return terra.ReferenceAsList[AuthenticationConfigurationAttributes](sc.ref.Append("authentication_configuration"))
}

func (sc SourceConfigurationAttributes) CodeRepository() terra.ListValue[CodeRepositoryAttributes] {
	return terra.ReferenceAsList[CodeRepositoryAttributes](sc.ref.Append("code_repository"))
}

func (sc SourceConfigurationAttributes) ImageRepository() terra.ListValue[ImageRepositoryAttributes] {
	return terra.ReferenceAsList[ImageRepositoryAttributes](sc.ref.Append("image_repository"))
}

type AuthenticationConfigurationAttributes struct {
	ref terra.Reference
}

func (ac AuthenticationConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AuthenticationConfigurationAttributes) InternalWithRef(ref terra.Reference) AuthenticationConfigurationAttributes {
	return AuthenticationConfigurationAttributes{ref: ref}
}

func (ac AuthenticationConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AuthenticationConfigurationAttributes) AccessRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("access_role_arn"))
}

func (ac AuthenticationConfigurationAttributes) ConnectionArn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("connection_arn"))
}

type CodeRepositoryAttributes struct {
	ref terra.Reference
}

func (cr CodeRepositoryAttributes) InternalRef() (terra.Reference, error) {
	return cr.ref, nil
}

func (cr CodeRepositoryAttributes) InternalWithRef(ref terra.Reference) CodeRepositoryAttributes {
	return CodeRepositoryAttributes{ref: ref}
}

func (cr CodeRepositoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cr.ref.InternalTokens()
}

func (cr CodeRepositoryAttributes) RepositoryUrl() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("repository_url"))
}

func (cr CodeRepositoryAttributes) CodeConfiguration() terra.ListValue[CodeConfigurationAttributes] {
	return terra.ReferenceAsList[CodeConfigurationAttributes](cr.ref.Append("code_configuration"))
}

func (cr CodeRepositoryAttributes) SourceCodeVersion() terra.ListValue[SourceCodeVersionAttributes] {
	return terra.ReferenceAsList[SourceCodeVersionAttributes](cr.ref.Append("source_code_version"))
}

type CodeConfigurationAttributes struct {
	ref terra.Reference
}

func (cc CodeConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc CodeConfigurationAttributes) InternalWithRef(ref terra.Reference) CodeConfigurationAttributes {
	return CodeConfigurationAttributes{ref: ref}
}

func (cc CodeConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc CodeConfigurationAttributes) ConfigurationSource() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("configuration_source"))
}

func (cc CodeConfigurationAttributes) CodeConfigurationValues() terra.ListValue[CodeConfigurationValuesAttributes] {
	return terra.ReferenceAsList[CodeConfigurationValuesAttributes](cc.ref.Append("code_configuration_values"))
}

type CodeConfigurationValuesAttributes struct {
	ref terra.Reference
}

func (ccv CodeConfigurationValuesAttributes) InternalRef() (terra.Reference, error) {
	return ccv.ref, nil
}

func (ccv CodeConfigurationValuesAttributes) InternalWithRef(ref terra.Reference) CodeConfigurationValuesAttributes {
	return CodeConfigurationValuesAttributes{ref: ref}
}

func (ccv CodeConfigurationValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ccv.ref.InternalTokens()
}

func (ccv CodeConfigurationValuesAttributes) BuildCommand() terra.StringValue {
	return terra.ReferenceAsString(ccv.ref.Append("build_command"))
}

func (ccv CodeConfigurationValuesAttributes) Port() terra.StringValue {
	return terra.ReferenceAsString(ccv.ref.Append("port"))
}

func (ccv CodeConfigurationValuesAttributes) Runtime() terra.StringValue {
	return terra.ReferenceAsString(ccv.ref.Append("runtime"))
}

func (ccv CodeConfigurationValuesAttributes) RuntimeEnvironmentSecrets() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ccv.ref.Append("runtime_environment_secrets"))
}

func (ccv CodeConfigurationValuesAttributes) RuntimeEnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ccv.ref.Append("runtime_environment_variables"))
}

func (ccv CodeConfigurationValuesAttributes) StartCommand() terra.StringValue {
	return terra.ReferenceAsString(ccv.ref.Append("start_command"))
}

type SourceCodeVersionAttributes struct {
	ref terra.Reference
}

func (scv SourceCodeVersionAttributes) InternalRef() (terra.Reference, error) {
	return scv.ref, nil
}

func (scv SourceCodeVersionAttributes) InternalWithRef(ref terra.Reference) SourceCodeVersionAttributes {
	return SourceCodeVersionAttributes{ref: ref}
}

func (scv SourceCodeVersionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return scv.ref.InternalTokens()
}

func (scv SourceCodeVersionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(scv.ref.Append("type"))
}

func (scv SourceCodeVersionAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(scv.ref.Append("value"))
}

type ImageRepositoryAttributes struct {
	ref terra.Reference
}

func (ir ImageRepositoryAttributes) InternalRef() (terra.Reference, error) {
	return ir.ref, nil
}

func (ir ImageRepositoryAttributes) InternalWithRef(ref terra.Reference) ImageRepositoryAttributes {
	return ImageRepositoryAttributes{ref: ref}
}

func (ir ImageRepositoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ir.ref.InternalTokens()
}

func (ir ImageRepositoryAttributes) ImageIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("image_identifier"))
}

func (ir ImageRepositoryAttributes) ImageRepositoryType() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("image_repository_type"))
}

func (ir ImageRepositoryAttributes) ImageConfiguration() terra.ListValue[ImageConfigurationAttributes] {
	return terra.ReferenceAsList[ImageConfigurationAttributes](ir.ref.Append("image_configuration"))
}

type ImageConfigurationAttributes struct {
	ref terra.Reference
}

func (ic ImageConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic ImageConfigurationAttributes) InternalWithRef(ref terra.Reference) ImageConfigurationAttributes {
	return ImageConfigurationAttributes{ref: ref}
}

func (ic ImageConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic ImageConfigurationAttributes) Port() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("port"))
}

func (ic ImageConfigurationAttributes) RuntimeEnvironmentSecrets() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ic.ref.Append("runtime_environment_secrets"))
}

func (ic ImageConfigurationAttributes) RuntimeEnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ic.ref.Append("runtime_environment_variables"))
}

func (ic ImageConfigurationAttributes) StartCommand() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("start_command"))
}

type EncryptionConfigurationState struct {
	KmsKey string `json:"kms_key"`
}

type HealthCheckConfigurationState struct {
	HealthyThreshold   float64 `json:"healthy_threshold"`
	Interval           float64 `json:"interval"`
	Path               string  `json:"path"`
	Protocol           string  `json:"protocol"`
	Timeout            float64 `json:"timeout"`
	UnhealthyThreshold float64 `json:"unhealthy_threshold"`
}

type InstanceConfigurationState struct {
	Cpu             string `json:"cpu"`
	InstanceRoleArn string `json:"instance_role_arn"`
	Memory          string `json:"memory"`
}

type NetworkConfigurationState struct {
	EgressConfiguration  []EgressConfigurationState  `json:"egress_configuration"`
	IngressConfiguration []IngressConfigurationState `json:"ingress_configuration"`
}

type EgressConfigurationState struct {
	EgressType      string `json:"egress_type"`
	VpcConnectorArn string `json:"vpc_connector_arn"`
}

type IngressConfigurationState struct {
	IsPubliclyAccessible bool `json:"is_publicly_accessible"`
}

type ObservabilityConfigurationState struct {
	ObservabilityConfigurationArn string `json:"observability_configuration_arn"`
	ObservabilityEnabled          bool   `json:"observability_enabled"`
}

type SourceConfigurationState struct {
	AutoDeploymentsEnabled      bool                               `json:"auto_deployments_enabled"`
	AuthenticationConfiguration []AuthenticationConfigurationState `json:"authentication_configuration"`
	CodeRepository              []CodeRepositoryState              `json:"code_repository"`
	ImageRepository             []ImageRepositoryState             `json:"image_repository"`
}

type AuthenticationConfigurationState struct {
	AccessRoleArn string `json:"access_role_arn"`
	ConnectionArn string `json:"connection_arn"`
}

type CodeRepositoryState struct {
	RepositoryUrl     string                   `json:"repository_url"`
	CodeConfiguration []CodeConfigurationState `json:"code_configuration"`
	SourceCodeVersion []SourceCodeVersionState `json:"source_code_version"`
}

type CodeConfigurationState struct {
	ConfigurationSource     string                         `json:"configuration_source"`
	CodeConfigurationValues []CodeConfigurationValuesState `json:"code_configuration_values"`
}

type CodeConfigurationValuesState struct {
	BuildCommand                string            `json:"build_command"`
	Port                        string            `json:"port"`
	Runtime                     string            `json:"runtime"`
	RuntimeEnvironmentSecrets   map[string]string `json:"runtime_environment_secrets"`
	RuntimeEnvironmentVariables map[string]string `json:"runtime_environment_variables"`
	StartCommand                string            `json:"start_command"`
}

type SourceCodeVersionState struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type ImageRepositoryState struct {
	ImageIdentifier     string                    `json:"image_identifier"`
	ImageRepositoryType string                    `json:"image_repository_type"`
	ImageConfiguration  []ImageConfigurationState `json:"image_configuration"`
}

type ImageConfigurationState struct {
	Port                        string            `json:"port"`
	RuntimeEnvironmentSecrets   map[string]string `json:"runtime_environment_secrets"`
	RuntimeEnvironmentVariables map[string]string `json:"runtime_environment_variables"`
	StartCommand                string            `json:"start_command"`
}
