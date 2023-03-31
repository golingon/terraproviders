// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataimagebuilderimagepipeline

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ImageTestsConfiguration struct{}

type Schedule struct{}

type ImageTestsConfigurationAttributes struct {
	ref terra.Reference
}

func (itc ImageTestsConfigurationAttributes) InternalRef() terra.Reference {
	return itc.ref
}

func (itc ImageTestsConfigurationAttributes) InternalWithRef(ref terra.Reference) ImageTestsConfigurationAttributes {
	return ImageTestsConfigurationAttributes{ref: ref}
}

func (itc ImageTestsConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return itc.ref.InternalTokens()
}

func (itc ImageTestsConfigurationAttributes) ImageTestsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(itc.ref.Append("image_tests_enabled"))
}

func (itc ImageTestsConfigurationAttributes) TimeoutMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(itc.ref.Append("timeout_minutes"))
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

func (s ScheduleAttributes) PipelineExecutionStartCondition() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("pipeline_execution_start_condition"))
}

func (s ScheduleAttributes) ScheduleExpression() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("schedule_expression"))
}

type ImageTestsConfigurationState struct {
	ImageTestsEnabled bool    `json:"image_tests_enabled"`
	TimeoutMinutes    float64 `json:"timeout_minutes"`
}

type ScheduleState struct {
	PipelineExecutionStartCondition string `json:"pipeline_execution_start_condition"`
	ScheduleExpression              string `json:"schedule_expression"`
}
