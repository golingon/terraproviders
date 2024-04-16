// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_lambda_function

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataDeadLetterConfigAttributes struct {
	ref terra.Reference
}

func (dlc DataDeadLetterConfigAttributes) InternalRef() (terra.Reference, error) {
	return dlc.ref, nil
}

func (dlc DataDeadLetterConfigAttributes) InternalWithRef(ref terra.Reference) DataDeadLetterConfigAttributes {
	return DataDeadLetterConfigAttributes{ref: ref}
}

func (dlc DataDeadLetterConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dlc.ref.InternalTokens()
}

func (dlc DataDeadLetterConfigAttributes) TargetArn() terra.StringValue {
	return terra.ReferenceAsString(dlc.ref.Append("target_arn"))
}

type DataEnvironmentAttributes struct {
	ref terra.Reference
}

func (e DataEnvironmentAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e DataEnvironmentAttributes) InternalWithRef(ref terra.Reference) DataEnvironmentAttributes {
	return DataEnvironmentAttributes{ref: ref}
}

func (e DataEnvironmentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e DataEnvironmentAttributes) Variables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](e.ref.Append("variables"))
}

type DataEphemeralStorageAttributes struct {
	ref terra.Reference
}

func (es DataEphemeralStorageAttributes) InternalRef() (terra.Reference, error) {
	return es.ref, nil
}

func (es DataEphemeralStorageAttributes) InternalWithRef(ref terra.Reference) DataEphemeralStorageAttributes {
	return DataEphemeralStorageAttributes{ref: ref}
}

func (es DataEphemeralStorageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return es.ref.InternalTokens()
}

func (es DataEphemeralStorageAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("size"))
}

type DataFileSystemConfigAttributes struct {
	ref terra.Reference
}

func (fsc DataFileSystemConfigAttributes) InternalRef() (terra.Reference, error) {
	return fsc.ref, nil
}

func (fsc DataFileSystemConfigAttributes) InternalWithRef(ref terra.Reference) DataFileSystemConfigAttributes {
	return DataFileSystemConfigAttributes{ref: ref}
}

func (fsc DataFileSystemConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fsc.ref.InternalTokens()
}

func (fsc DataFileSystemConfigAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fsc.ref.Append("arn"))
}

func (fsc DataFileSystemConfigAttributes) LocalMountPath() terra.StringValue {
	return terra.ReferenceAsString(fsc.ref.Append("local_mount_path"))
}

type DataLoggingConfigAttributes struct {
	ref terra.Reference
}

func (lc DataLoggingConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc DataLoggingConfigAttributes) InternalWithRef(ref terra.Reference) DataLoggingConfigAttributes {
	return DataLoggingConfigAttributes{ref: ref}
}

func (lc DataLoggingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc DataLoggingConfigAttributes) ApplicationLogLevel() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("application_log_level"))
}

func (lc DataLoggingConfigAttributes) LogFormat() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("log_format"))
}

func (lc DataLoggingConfigAttributes) LogGroup() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("log_group"))
}

func (lc DataLoggingConfigAttributes) SystemLogLevel() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("system_log_level"))
}

type DataTracingConfigAttributes struct {
	ref terra.Reference
}

func (tc DataTracingConfigAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc DataTracingConfigAttributes) InternalWithRef(ref terra.Reference) DataTracingConfigAttributes {
	return DataTracingConfigAttributes{ref: ref}
}

func (tc DataTracingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tc.ref.InternalTokens()
}

func (tc DataTracingConfigAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("mode"))
}

type DataVpcConfigAttributes struct {
	ref terra.Reference
}

func (vc DataVpcConfigAttributes) InternalRef() (terra.Reference, error) {
	return vc.ref, nil
}

func (vc DataVpcConfigAttributes) InternalWithRef(ref terra.Reference) DataVpcConfigAttributes {
	return DataVpcConfigAttributes{ref: ref}
}

func (vc DataVpcConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vc.ref.InternalTokens()
}

func (vc DataVpcConfigAttributes) Ipv6AllowedForDualStack() terra.BoolValue {
	return terra.ReferenceAsBool(vc.ref.Append("ipv6_allowed_for_dual_stack"))
}

func (vc DataVpcConfigAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("security_group_ids"))
}

func (vc DataVpcConfigAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("subnet_ids"))
}

func (vc DataVpcConfigAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("vpc_id"))
}

type DataDeadLetterConfigState struct {
	TargetArn string `json:"target_arn"`
}

type DataEnvironmentState struct {
	Variables map[string]string `json:"variables"`
}

type DataEphemeralStorageState struct {
	Size float64 `json:"size"`
}

type DataFileSystemConfigState struct {
	Arn            string `json:"arn"`
	LocalMountPath string `json:"local_mount_path"`
}

type DataLoggingConfigState struct {
	ApplicationLogLevel string `json:"application_log_level"`
	LogFormat           string `json:"log_format"`
	LogGroup            string `json:"log_group"`
	SystemLogLevel      string `json:"system_log_level"`
}

type DataTracingConfigState struct {
	Mode string `json:"mode"`
}

type DataVpcConfigState struct {
	Ipv6AllowedForDualStack bool     `json:"ipv6_allowed_for_dual_stack"`
	SecurityGroupIds        []string `json:"security_group_ids"`
	SubnetIds               []string `json:"subnet_ids"`
	VpcId                   string   `json:"vpc_id"`
}
