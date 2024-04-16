// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_api_gateway_stage

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AccessLogSettings struct {
	// DestinationArn: string, required
	DestinationArn terra.StringValue `hcl:"destination_arn,attr" validate:"required"`
	// Format: string, required
	Format terra.StringValue `hcl:"format,attr" validate:"required"`
}

type CanarySettings struct {
	// PercentTraffic: number, optional
	PercentTraffic terra.NumberValue `hcl:"percent_traffic,attr"`
	// StageVariableOverrides: map of string, optional
	StageVariableOverrides terra.MapValue[terra.StringValue] `hcl:"stage_variable_overrides,attr"`
	// UseStageCache: bool, optional
	UseStageCache terra.BoolValue `hcl:"use_stage_cache,attr"`
}

type AccessLogSettingsAttributes struct {
	ref terra.Reference
}

func (als AccessLogSettingsAttributes) InternalRef() (terra.Reference, error) {
	return als.ref, nil
}

func (als AccessLogSettingsAttributes) InternalWithRef(ref terra.Reference) AccessLogSettingsAttributes {
	return AccessLogSettingsAttributes{ref: ref}
}

func (als AccessLogSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return als.ref.InternalTokens()
}

func (als AccessLogSettingsAttributes) DestinationArn() terra.StringValue {
	return terra.ReferenceAsString(als.ref.Append("destination_arn"))
}

func (als AccessLogSettingsAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(als.ref.Append("format"))
}

type CanarySettingsAttributes struct {
	ref terra.Reference
}

func (cs CanarySettingsAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs CanarySettingsAttributes) InternalWithRef(ref terra.Reference) CanarySettingsAttributes {
	return CanarySettingsAttributes{ref: ref}
}

func (cs CanarySettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs CanarySettingsAttributes) PercentTraffic() terra.NumberValue {
	return terra.ReferenceAsNumber(cs.ref.Append("percent_traffic"))
}

func (cs CanarySettingsAttributes) StageVariableOverrides() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cs.ref.Append("stage_variable_overrides"))
}

func (cs CanarySettingsAttributes) UseStageCache() terra.BoolValue {
	return terra.ReferenceAsBool(cs.ref.Append("use_stage_cache"))
}

type AccessLogSettingsState struct {
	DestinationArn string `json:"destination_arn"`
	Format         string `json:"format"`
}

type CanarySettingsState struct {
	PercentTraffic         float64           `json:"percent_traffic"`
	StageVariableOverrides map[string]string `json:"stage_variable_overrides"`
	UseStageCache          bool              `json:"use_stage_cache"`
}
