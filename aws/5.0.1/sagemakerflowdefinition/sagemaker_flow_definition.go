// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sagemakerflowdefinition

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type HumanLoopActivationConfig struct {
	// HumanLoopActivationConditionsConfig: optional
	HumanLoopActivationConditionsConfig *HumanLoopActivationConditionsConfig `hcl:"human_loop_activation_conditions_config,block"`
}

type HumanLoopActivationConditionsConfig struct {
	// HumanLoopActivationConditions: string, required
	HumanLoopActivationConditions terra.StringValue `hcl:"human_loop_activation_conditions,attr" validate:"required"`
}

type HumanLoopConfig struct {
	// HumanTaskUiArn: string, required
	HumanTaskUiArn terra.StringValue `hcl:"human_task_ui_arn,attr" validate:"required"`
	// TaskAvailabilityLifetimeInSeconds: number, optional
	TaskAvailabilityLifetimeInSeconds terra.NumberValue `hcl:"task_availability_lifetime_in_seconds,attr"`
	// TaskCount: number, required
	TaskCount terra.NumberValue `hcl:"task_count,attr" validate:"required"`
	// TaskDescription: string, required
	TaskDescription terra.StringValue `hcl:"task_description,attr" validate:"required"`
	// TaskKeywords: set of string, optional
	TaskKeywords terra.SetValue[terra.StringValue] `hcl:"task_keywords,attr"`
	// TaskTimeLimitInSeconds: number, optional
	TaskTimeLimitInSeconds terra.NumberValue `hcl:"task_time_limit_in_seconds,attr"`
	// TaskTitle: string, required
	TaskTitle terra.StringValue `hcl:"task_title,attr" validate:"required"`
	// WorkteamArn: string, required
	WorkteamArn terra.StringValue `hcl:"workteam_arn,attr" validate:"required"`
	// PublicWorkforceTaskPrice: optional
	PublicWorkforceTaskPrice *PublicWorkforceTaskPrice `hcl:"public_workforce_task_price,block"`
}

type PublicWorkforceTaskPrice struct {
	// AmountInUsd: optional
	AmountInUsd *AmountInUsd `hcl:"amount_in_usd,block"`
}

type AmountInUsd struct {
	// Cents: number, optional
	Cents terra.NumberValue `hcl:"cents,attr"`
	// Dollars: number, optional
	Dollars terra.NumberValue `hcl:"dollars,attr"`
	// TenthFractionsOfACent: number, optional
	TenthFractionsOfACent terra.NumberValue `hcl:"tenth_fractions_of_a_cent,attr"`
}

type HumanLoopRequestSource struct {
	// AwsManagedHumanLoopRequestSource: string, required
	AwsManagedHumanLoopRequestSource terra.StringValue `hcl:"aws_managed_human_loop_request_source,attr" validate:"required"`
}

type OutputConfig struct {
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// S3OutputPath: string, required
	S3OutputPath terra.StringValue `hcl:"s3_output_path,attr" validate:"required"`
}

type HumanLoopActivationConfigAttributes struct {
	ref terra.Reference
}

func (hlac HumanLoopActivationConfigAttributes) InternalRef() (terra.Reference, error) {
	return hlac.ref, nil
}

func (hlac HumanLoopActivationConfigAttributes) InternalWithRef(ref terra.Reference) HumanLoopActivationConfigAttributes {
	return HumanLoopActivationConfigAttributes{ref: ref}
}

func (hlac HumanLoopActivationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hlac.ref.InternalTokens()
}

func (hlac HumanLoopActivationConfigAttributes) HumanLoopActivationConditionsConfig() terra.ListValue[HumanLoopActivationConditionsConfigAttributes] {
	return terra.ReferenceAsList[HumanLoopActivationConditionsConfigAttributes](hlac.ref.Append("human_loop_activation_conditions_config"))
}

type HumanLoopActivationConditionsConfigAttributes struct {
	ref terra.Reference
}

func (hlacc HumanLoopActivationConditionsConfigAttributes) InternalRef() (terra.Reference, error) {
	return hlacc.ref, nil
}

func (hlacc HumanLoopActivationConditionsConfigAttributes) InternalWithRef(ref terra.Reference) HumanLoopActivationConditionsConfigAttributes {
	return HumanLoopActivationConditionsConfigAttributes{ref: ref}
}

func (hlacc HumanLoopActivationConditionsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hlacc.ref.InternalTokens()
}

func (hlacc HumanLoopActivationConditionsConfigAttributes) HumanLoopActivationConditions() terra.StringValue {
	return terra.ReferenceAsString(hlacc.ref.Append("human_loop_activation_conditions"))
}

type HumanLoopConfigAttributes struct {
	ref terra.Reference
}

func (hlc HumanLoopConfigAttributes) InternalRef() (terra.Reference, error) {
	return hlc.ref, nil
}

func (hlc HumanLoopConfigAttributes) InternalWithRef(ref terra.Reference) HumanLoopConfigAttributes {
	return HumanLoopConfigAttributes{ref: ref}
}

func (hlc HumanLoopConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hlc.ref.InternalTokens()
}

func (hlc HumanLoopConfigAttributes) HumanTaskUiArn() terra.StringValue {
	return terra.ReferenceAsString(hlc.ref.Append("human_task_ui_arn"))
}

func (hlc HumanLoopConfigAttributes) TaskAvailabilityLifetimeInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(hlc.ref.Append("task_availability_lifetime_in_seconds"))
}

func (hlc HumanLoopConfigAttributes) TaskCount() terra.NumberValue {
	return terra.ReferenceAsNumber(hlc.ref.Append("task_count"))
}

func (hlc HumanLoopConfigAttributes) TaskDescription() terra.StringValue {
	return terra.ReferenceAsString(hlc.ref.Append("task_description"))
}

func (hlc HumanLoopConfigAttributes) TaskKeywords() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hlc.ref.Append("task_keywords"))
}

func (hlc HumanLoopConfigAttributes) TaskTimeLimitInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(hlc.ref.Append("task_time_limit_in_seconds"))
}

func (hlc HumanLoopConfigAttributes) TaskTitle() terra.StringValue {
	return terra.ReferenceAsString(hlc.ref.Append("task_title"))
}

func (hlc HumanLoopConfigAttributes) WorkteamArn() terra.StringValue {
	return terra.ReferenceAsString(hlc.ref.Append("workteam_arn"))
}

func (hlc HumanLoopConfigAttributes) PublicWorkforceTaskPrice() terra.ListValue[PublicWorkforceTaskPriceAttributes] {
	return terra.ReferenceAsList[PublicWorkforceTaskPriceAttributes](hlc.ref.Append("public_workforce_task_price"))
}

type PublicWorkforceTaskPriceAttributes struct {
	ref terra.Reference
}

func (pwtp PublicWorkforceTaskPriceAttributes) InternalRef() (terra.Reference, error) {
	return pwtp.ref, nil
}

func (pwtp PublicWorkforceTaskPriceAttributes) InternalWithRef(ref terra.Reference) PublicWorkforceTaskPriceAttributes {
	return PublicWorkforceTaskPriceAttributes{ref: ref}
}

func (pwtp PublicWorkforceTaskPriceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pwtp.ref.InternalTokens()
}

func (pwtp PublicWorkforceTaskPriceAttributes) AmountInUsd() terra.ListValue[AmountInUsdAttributes] {
	return terra.ReferenceAsList[AmountInUsdAttributes](pwtp.ref.Append("amount_in_usd"))
}

type AmountInUsdAttributes struct {
	ref terra.Reference
}

func (aiu AmountInUsdAttributes) InternalRef() (terra.Reference, error) {
	return aiu.ref, nil
}

func (aiu AmountInUsdAttributes) InternalWithRef(ref terra.Reference) AmountInUsdAttributes {
	return AmountInUsdAttributes{ref: ref}
}

func (aiu AmountInUsdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aiu.ref.InternalTokens()
}

func (aiu AmountInUsdAttributes) Cents() terra.NumberValue {
	return terra.ReferenceAsNumber(aiu.ref.Append("cents"))
}

func (aiu AmountInUsdAttributes) Dollars() terra.NumberValue {
	return terra.ReferenceAsNumber(aiu.ref.Append("dollars"))
}

func (aiu AmountInUsdAttributes) TenthFractionsOfACent() terra.NumberValue {
	return terra.ReferenceAsNumber(aiu.ref.Append("tenth_fractions_of_a_cent"))
}

type HumanLoopRequestSourceAttributes struct {
	ref terra.Reference
}

func (hlrs HumanLoopRequestSourceAttributes) InternalRef() (terra.Reference, error) {
	return hlrs.ref, nil
}

func (hlrs HumanLoopRequestSourceAttributes) InternalWithRef(ref terra.Reference) HumanLoopRequestSourceAttributes {
	return HumanLoopRequestSourceAttributes{ref: ref}
}

func (hlrs HumanLoopRequestSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hlrs.ref.InternalTokens()
}

func (hlrs HumanLoopRequestSourceAttributes) AwsManagedHumanLoopRequestSource() terra.StringValue {
	return terra.ReferenceAsString(hlrs.ref.Append("aws_managed_human_loop_request_source"))
}

type OutputConfigAttributes struct {
	ref terra.Reference
}

func (oc OutputConfigAttributes) InternalRef() (terra.Reference, error) {
	return oc.ref, nil
}

func (oc OutputConfigAttributes) InternalWithRef(ref terra.Reference) OutputConfigAttributes {
	return OutputConfigAttributes{ref: ref}
}

func (oc OutputConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oc.ref.InternalTokens()
}

func (oc OutputConfigAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("kms_key_id"))
}

func (oc OutputConfigAttributes) S3OutputPath() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("s3_output_path"))
}

type HumanLoopActivationConfigState struct {
	HumanLoopActivationConditionsConfig []HumanLoopActivationConditionsConfigState `json:"human_loop_activation_conditions_config"`
}

type HumanLoopActivationConditionsConfigState struct {
	HumanLoopActivationConditions string `json:"human_loop_activation_conditions"`
}

type HumanLoopConfigState struct {
	HumanTaskUiArn                    string                          `json:"human_task_ui_arn"`
	TaskAvailabilityLifetimeInSeconds float64                         `json:"task_availability_lifetime_in_seconds"`
	TaskCount                         float64                         `json:"task_count"`
	TaskDescription                   string                          `json:"task_description"`
	TaskKeywords                      []string                        `json:"task_keywords"`
	TaskTimeLimitInSeconds            float64                         `json:"task_time_limit_in_seconds"`
	TaskTitle                         string                          `json:"task_title"`
	WorkteamArn                       string                          `json:"workteam_arn"`
	PublicWorkforceTaskPrice          []PublicWorkforceTaskPriceState `json:"public_workforce_task_price"`
}

type PublicWorkforceTaskPriceState struct {
	AmountInUsd []AmountInUsdState `json:"amount_in_usd"`
}

type AmountInUsdState struct {
	Cents                 float64 `json:"cents"`
	Dollars               float64 `json:"dollars"`
	TenthFractionsOfACent float64 `json:"tenth_fractions_of_a_cent"`
}

type HumanLoopRequestSourceState struct {
	AwsManagedHumanLoopRequestSource string `json:"aws_managed_human_loop_request_source"`
}

type OutputConfigState struct {
	KmsKeyId     string `json:"kms_key_id"`
	S3OutputPath string `json:"s3_output_path"`
}
