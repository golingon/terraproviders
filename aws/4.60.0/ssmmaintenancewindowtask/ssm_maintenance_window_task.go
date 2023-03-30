// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ssmmaintenancewindowtask

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Targets struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type TaskInvocationParameters struct {
	// AutomationParameters: optional
	AutomationParameters *AutomationParameters `hcl:"automation_parameters,block"`
	// LambdaParameters: optional
	LambdaParameters *LambdaParameters `hcl:"lambda_parameters,block"`
	// RunCommandParameters: optional
	RunCommandParameters *RunCommandParameters `hcl:"run_command_parameters,block"`
	// StepFunctionsParameters: optional
	StepFunctionsParameters *StepFunctionsParameters `hcl:"step_functions_parameters,block"`
}

type AutomationParameters struct {
	// DocumentVersion: string, optional
	DocumentVersion terra.StringValue `hcl:"document_version,attr"`
	// AutomationParametersParameter: min=0
	Parameter []AutomationParametersParameter `hcl:"parameter,block" validate:"min=0"`
}

type AutomationParametersParameter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type LambdaParameters struct {
	// ClientContext: string, optional
	ClientContext terra.StringValue `hcl:"client_context,attr"`
	// Payload: string, optional
	Payload terra.StringValue `hcl:"payload,attr"`
	// Qualifier: string, optional
	Qualifier terra.StringValue `hcl:"qualifier,attr"`
}

type RunCommandParameters struct {
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// DocumentHash: string, optional
	DocumentHash terra.StringValue `hcl:"document_hash,attr"`
	// DocumentHashType: string, optional
	DocumentHashType terra.StringValue `hcl:"document_hash_type,attr"`
	// DocumentVersion: string, optional
	DocumentVersion terra.StringValue `hcl:"document_version,attr"`
	// OutputS3Bucket: string, optional
	OutputS3Bucket terra.StringValue `hcl:"output_s3_bucket,attr"`
	// OutputS3KeyPrefix: string, optional
	OutputS3KeyPrefix terra.StringValue `hcl:"output_s3_key_prefix,attr"`
	// ServiceRoleArn: string, optional
	ServiceRoleArn terra.StringValue `hcl:"service_role_arn,attr"`
	// TimeoutSeconds: number, optional
	TimeoutSeconds terra.NumberValue `hcl:"timeout_seconds,attr"`
	// CloudwatchConfig: optional
	CloudwatchConfig *CloudwatchConfig `hcl:"cloudwatch_config,block"`
	// NotificationConfig: optional
	NotificationConfig *NotificationConfig `hcl:"notification_config,block"`
	// RunCommandParametersParameter: min=0
	Parameter []RunCommandParametersParameter `hcl:"parameter,block" validate:"min=0"`
}

type CloudwatchConfig struct {
	// CloudwatchLogGroupName: string, optional
	CloudwatchLogGroupName terra.StringValue `hcl:"cloudwatch_log_group_name,attr"`
	// CloudwatchOutputEnabled: bool, optional
	CloudwatchOutputEnabled terra.BoolValue `hcl:"cloudwatch_output_enabled,attr"`
}

type NotificationConfig struct {
	// NotificationArn: string, optional
	NotificationArn terra.StringValue `hcl:"notification_arn,attr"`
	// NotificationEvents: list of string, optional
	NotificationEvents terra.ListValue[terra.StringValue] `hcl:"notification_events,attr"`
	// NotificationType: string, optional
	NotificationType terra.StringValue `hcl:"notification_type,attr"`
}

type RunCommandParametersParameter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type StepFunctionsParameters struct {
	// Input: string, optional
	Input terra.StringValue `hcl:"input,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}

type TargetsAttributes struct {
	ref terra.Reference
}

func (t TargetsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TargetsAttributes) InternalWithRef(ref terra.Reference) TargetsAttributes {
	return TargetsAttributes{ref: ref}
}

func (t TargetsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TargetsAttributes) Key() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("key"))
}

func (t TargetsAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](t.ref.Append("values"))
}

type TaskInvocationParametersAttributes struct {
	ref terra.Reference
}

func (tip TaskInvocationParametersAttributes) InternalRef() terra.Reference {
	return tip.ref
}

func (tip TaskInvocationParametersAttributes) InternalWithRef(ref terra.Reference) TaskInvocationParametersAttributes {
	return TaskInvocationParametersAttributes{ref: ref}
}

func (tip TaskInvocationParametersAttributes) InternalTokens() hclwrite.Tokens {
	return tip.ref.InternalTokens()
}

func (tip TaskInvocationParametersAttributes) AutomationParameters() terra.ListValue[AutomationParametersAttributes] {
	return terra.ReferenceList[AutomationParametersAttributes](tip.ref.Append("automation_parameters"))
}

func (tip TaskInvocationParametersAttributes) LambdaParameters() terra.ListValue[LambdaParametersAttributes] {
	return terra.ReferenceList[LambdaParametersAttributes](tip.ref.Append("lambda_parameters"))
}

func (tip TaskInvocationParametersAttributes) RunCommandParameters() terra.ListValue[RunCommandParametersAttributes] {
	return terra.ReferenceList[RunCommandParametersAttributes](tip.ref.Append("run_command_parameters"))
}

func (tip TaskInvocationParametersAttributes) StepFunctionsParameters() terra.ListValue[StepFunctionsParametersAttributes] {
	return terra.ReferenceList[StepFunctionsParametersAttributes](tip.ref.Append("step_functions_parameters"))
}

type AutomationParametersAttributes struct {
	ref terra.Reference
}

func (ap AutomationParametersAttributes) InternalRef() terra.Reference {
	return ap.ref
}

func (ap AutomationParametersAttributes) InternalWithRef(ref terra.Reference) AutomationParametersAttributes {
	return AutomationParametersAttributes{ref: ref}
}

func (ap AutomationParametersAttributes) InternalTokens() hclwrite.Tokens {
	return ap.ref.InternalTokens()
}

func (ap AutomationParametersAttributes) DocumentVersion() terra.StringValue {
	return terra.ReferenceString(ap.ref.Append("document_version"))
}

func (ap AutomationParametersAttributes) Parameter() terra.SetValue[AutomationParametersParameterAttributes] {
	return terra.ReferenceSet[AutomationParametersParameterAttributes](ap.ref.Append("parameter"))
}

type AutomationParametersParameterAttributes struct {
	ref terra.Reference
}

func (p AutomationParametersParameterAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p AutomationParametersParameterAttributes) InternalWithRef(ref terra.Reference) AutomationParametersParameterAttributes {
	return AutomationParametersParameterAttributes{ref: ref}
}

func (p AutomationParametersParameterAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p AutomationParametersParameterAttributes) Name() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("name"))
}

func (p AutomationParametersParameterAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](p.ref.Append("values"))
}

type LambdaParametersAttributes struct {
	ref terra.Reference
}

func (lp LambdaParametersAttributes) InternalRef() terra.Reference {
	return lp.ref
}

func (lp LambdaParametersAttributes) InternalWithRef(ref terra.Reference) LambdaParametersAttributes {
	return LambdaParametersAttributes{ref: ref}
}

func (lp LambdaParametersAttributes) InternalTokens() hclwrite.Tokens {
	return lp.ref.InternalTokens()
}

func (lp LambdaParametersAttributes) ClientContext() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("client_context"))
}

func (lp LambdaParametersAttributes) Payload() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("payload"))
}

func (lp LambdaParametersAttributes) Qualifier() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("qualifier"))
}

type RunCommandParametersAttributes struct {
	ref terra.Reference
}

func (rcp RunCommandParametersAttributes) InternalRef() terra.Reference {
	return rcp.ref
}

func (rcp RunCommandParametersAttributes) InternalWithRef(ref terra.Reference) RunCommandParametersAttributes {
	return RunCommandParametersAttributes{ref: ref}
}

func (rcp RunCommandParametersAttributes) InternalTokens() hclwrite.Tokens {
	return rcp.ref.InternalTokens()
}

func (rcp RunCommandParametersAttributes) Comment() terra.StringValue {
	return terra.ReferenceString(rcp.ref.Append("comment"))
}

func (rcp RunCommandParametersAttributes) DocumentHash() terra.StringValue {
	return terra.ReferenceString(rcp.ref.Append("document_hash"))
}

func (rcp RunCommandParametersAttributes) DocumentHashType() terra.StringValue {
	return terra.ReferenceString(rcp.ref.Append("document_hash_type"))
}

func (rcp RunCommandParametersAttributes) DocumentVersion() terra.StringValue {
	return terra.ReferenceString(rcp.ref.Append("document_version"))
}

func (rcp RunCommandParametersAttributes) OutputS3Bucket() terra.StringValue {
	return terra.ReferenceString(rcp.ref.Append("output_s3_bucket"))
}

func (rcp RunCommandParametersAttributes) OutputS3KeyPrefix() terra.StringValue {
	return terra.ReferenceString(rcp.ref.Append("output_s3_key_prefix"))
}

func (rcp RunCommandParametersAttributes) ServiceRoleArn() terra.StringValue {
	return terra.ReferenceString(rcp.ref.Append("service_role_arn"))
}

func (rcp RunCommandParametersAttributes) TimeoutSeconds() terra.NumberValue {
	return terra.ReferenceNumber(rcp.ref.Append("timeout_seconds"))
}

func (rcp RunCommandParametersAttributes) CloudwatchConfig() terra.ListValue[CloudwatchConfigAttributes] {
	return terra.ReferenceList[CloudwatchConfigAttributes](rcp.ref.Append("cloudwatch_config"))
}

func (rcp RunCommandParametersAttributes) NotificationConfig() terra.ListValue[NotificationConfigAttributes] {
	return terra.ReferenceList[NotificationConfigAttributes](rcp.ref.Append("notification_config"))
}

func (rcp RunCommandParametersAttributes) Parameter() terra.SetValue[RunCommandParametersParameterAttributes] {
	return terra.ReferenceSet[RunCommandParametersParameterAttributes](rcp.ref.Append("parameter"))
}

type CloudwatchConfigAttributes struct {
	ref terra.Reference
}

func (cc CloudwatchConfigAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc CloudwatchConfigAttributes) InternalWithRef(ref terra.Reference) CloudwatchConfigAttributes {
	return CloudwatchConfigAttributes{ref: ref}
}

func (cc CloudwatchConfigAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc CloudwatchConfigAttributes) CloudwatchLogGroupName() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("cloudwatch_log_group_name"))
}

func (cc CloudwatchConfigAttributes) CloudwatchOutputEnabled() terra.BoolValue {
	return terra.ReferenceBool(cc.ref.Append("cloudwatch_output_enabled"))
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

func (nc NotificationConfigAttributes) NotificationArn() terra.StringValue {
	return terra.ReferenceString(nc.ref.Append("notification_arn"))
}

func (nc NotificationConfigAttributes) NotificationEvents() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](nc.ref.Append("notification_events"))
}

func (nc NotificationConfigAttributes) NotificationType() terra.StringValue {
	return terra.ReferenceString(nc.ref.Append("notification_type"))
}

type RunCommandParametersParameterAttributes struct {
	ref terra.Reference
}

func (p RunCommandParametersParameterAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p RunCommandParametersParameterAttributes) InternalWithRef(ref terra.Reference) RunCommandParametersParameterAttributes {
	return RunCommandParametersParameterAttributes{ref: ref}
}

func (p RunCommandParametersParameterAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p RunCommandParametersParameterAttributes) Name() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("name"))
}

func (p RunCommandParametersParameterAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](p.ref.Append("values"))
}

type StepFunctionsParametersAttributes struct {
	ref terra.Reference
}

func (sfp StepFunctionsParametersAttributes) InternalRef() terra.Reference {
	return sfp.ref
}

func (sfp StepFunctionsParametersAttributes) InternalWithRef(ref terra.Reference) StepFunctionsParametersAttributes {
	return StepFunctionsParametersAttributes{ref: ref}
}

func (sfp StepFunctionsParametersAttributes) InternalTokens() hclwrite.Tokens {
	return sfp.ref.InternalTokens()
}

func (sfp StepFunctionsParametersAttributes) Input() terra.StringValue {
	return terra.ReferenceString(sfp.ref.Append("input"))
}

func (sfp StepFunctionsParametersAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sfp.ref.Append("name"))
}

type TargetsState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type TaskInvocationParametersState struct {
	AutomationParameters    []AutomationParametersState    `json:"automation_parameters"`
	LambdaParameters        []LambdaParametersState        `json:"lambda_parameters"`
	RunCommandParameters    []RunCommandParametersState    `json:"run_command_parameters"`
	StepFunctionsParameters []StepFunctionsParametersState `json:"step_functions_parameters"`
}

type AutomationParametersState struct {
	DocumentVersion string                               `json:"document_version"`
	Parameter       []AutomationParametersParameterState `json:"parameter"`
}

type AutomationParametersParameterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type LambdaParametersState struct {
	ClientContext string `json:"client_context"`
	Payload       string `json:"payload"`
	Qualifier     string `json:"qualifier"`
}

type RunCommandParametersState struct {
	Comment            string                               `json:"comment"`
	DocumentHash       string                               `json:"document_hash"`
	DocumentHashType   string                               `json:"document_hash_type"`
	DocumentVersion    string                               `json:"document_version"`
	OutputS3Bucket     string                               `json:"output_s3_bucket"`
	OutputS3KeyPrefix  string                               `json:"output_s3_key_prefix"`
	ServiceRoleArn     string                               `json:"service_role_arn"`
	TimeoutSeconds     float64                              `json:"timeout_seconds"`
	CloudwatchConfig   []CloudwatchConfigState              `json:"cloudwatch_config"`
	NotificationConfig []NotificationConfigState            `json:"notification_config"`
	Parameter          []RunCommandParametersParameterState `json:"parameter"`
}

type CloudwatchConfigState struct {
	CloudwatchLogGroupName  string `json:"cloudwatch_log_group_name"`
	CloudwatchOutputEnabled bool   `json:"cloudwatch_output_enabled"`
}

type NotificationConfigState struct {
	NotificationArn    string   `json:"notification_arn"`
	NotificationEvents []string `json:"notification_events"`
	NotificationType   string   `json:"notification_type"`
}

type RunCommandParametersParameterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type StepFunctionsParametersState struct {
	Input string `json:"input"`
	Name  string `json:"name"`
}
