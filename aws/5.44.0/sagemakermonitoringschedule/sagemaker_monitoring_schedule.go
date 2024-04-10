// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package sagemakermonitoringschedule

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type MonitoringScheduleConfig struct {
	// MonitoringJobDefinitionName: string, required
	MonitoringJobDefinitionName terra.StringValue `hcl:"monitoring_job_definition_name,attr" validate:"required"`
	// MonitoringType: string, required
	MonitoringType terra.StringValue `hcl:"monitoring_type,attr" validate:"required"`
	// ScheduleConfig: optional
	ScheduleConfig *ScheduleConfig `hcl:"schedule_config,block"`
}

type ScheduleConfig struct {
	// ScheduleExpression: string, required
	ScheduleExpression terra.StringValue `hcl:"schedule_expression,attr" validate:"required"`
}

type MonitoringScheduleConfigAttributes struct {
	ref terra.Reference
}

func (msc MonitoringScheduleConfigAttributes) InternalRef() (terra.Reference, error) {
	return msc.ref, nil
}

func (msc MonitoringScheduleConfigAttributes) InternalWithRef(ref terra.Reference) MonitoringScheduleConfigAttributes {
	return MonitoringScheduleConfigAttributes{ref: ref}
}

func (msc MonitoringScheduleConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return msc.ref.InternalTokens()
}

func (msc MonitoringScheduleConfigAttributes) MonitoringJobDefinitionName() terra.StringValue {
	return terra.ReferenceAsString(msc.ref.Append("monitoring_job_definition_name"))
}

func (msc MonitoringScheduleConfigAttributes) MonitoringType() terra.StringValue {
	return terra.ReferenceAsString(msc.ref.Append("monitoring_type"))
}

func (msc MonitoringScheduleConfigAttributes) ScheduleConfig() terra.ListValue[ScheduleConfigAttributes] {
	return terra.ReferenceAsList[ScheduleConfigAttributes](msc.ref.Append("schedule_config"))
}

type ScheduleConfigAttributes struct {
	ref terra.Reference
}

func (sc ScheduleConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc ScheduleConfigAttributes) InternalWithRef(ref terra.Reference) ScheduleConfigAttributes {
	return ScheduleConfigAttributes{ref: ref}
}

func (sc ScheduleConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc ScheduleConfigAttributes) ScheduleExpression() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("schedule_expression"))
}

type MonitoringScheduleConfigState struct {
	MonitoringJobDefinitionName string                `json:"monitoring_job_definition_name"`
	MonitoringType              string                `json:"monitoring_type"`
	ScheduleConfig              []ScheduleConfigState `json:"schedule_config"`
}

type ScheduleConfigState struct {
	ScheduleExpression string `json:"schedule_expression"`
}
