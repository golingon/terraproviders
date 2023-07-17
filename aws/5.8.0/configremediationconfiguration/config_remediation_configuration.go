// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package configremediationconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ExecutionControls struct {
	// SsmControls: optional
	SsmControls *SsmControls `hcl:"ssm_controls,block"`
}

type SsmControls struct {
	// ConcurrentExecutionRatePercentage: number, optional
	ConcurrentExecutionRatePercentage terra.NumberValue `hcl:"concurrent_execution_rate_percentage,attr"`
	// ErrorPercentage: number, optional
	ErrorPercentage terra.NumberValue `hcl:"error_percentage,attr"`
}

type Parameter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceValue: string, optional
	ResourceValue terra.StringValue `hcl:"resource_value,attr"`
	// StaticValue: string, optional
	StaticValue terra.StringValue `hcl:"static_value,attr"`
	// StaticValues: list of string, optional
	StaticValues terra.ListValue[terra.StringValue] `hcl:"static_values,attr"`
}

type ExecutionControlsAttributes struct {
	ref terra.Reference
}

func (ec ExecutionControlsAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec ExecutionControlsAttributes) InternalWithRef(ref terra.Reference) ExecutionControlsAttributes {
	return ExecutionControlsAttributes{ref: ref}
}

func (ec ExecutionControlsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec ExecutionControlsAttributes) SsmControls() terra.ListValue[SsmControlsAttributes] {
	return terra.ReferenceAsList[SsmControlsAttributes](ec.ref.Append("ssm_controls"))
}

type SsmControlsAttributes struct {
	ref terra.Reference
}

func (sc SsmControlsAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SsmControlsAttributes) InternalWithRef(ref terra.Reference) SsmControlsAttributes {
	return SsmControlsAttributes{ref: ref}
}

func (sc SsmControlsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SsmControlsAttributes) ConcurrentExecutionRatePercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("concurrent_execution_rate_percentage"))
}

func (sc SsmControlsAttributes) ErrorPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("error_percentage"))
}

type ParameterAttributes struct {
	ref terra.Reference
}

func (p ParameterAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParameterAttributes) InternalWithRef(ref terra.Reference) ParameterAttributes {
	return ParameterAttributes{ref: ref}
}

func (p ParameterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ParameterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p ParameterAttributes) ResourceValue() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("resource_value"))
}

func (p ParameterAttributes) StaticValue() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("static_value"))
}

func (p ParameterAttributes) StaticValues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](p.ref.Append("static_values"))
}

type ExecutionControlsState struct {
	SsmControls []SsmControlsState `json:"ssm_controls"`
}

type SsmControlsState struct {
	ConcurrentExecutionRatePercentage float64 `json:"concurrent_execution_rate_percentage"`
	ErrorPercentage                   float64 `json:"error_percentage"`
}

type ParameterState struct {
	Name          string   `json:"name"`
	ResourceValue string   `json:"resource_value"`
	StaticValue   string   `json:"static_value"`
	StaticValues  []string `json:"static_values"`
}
