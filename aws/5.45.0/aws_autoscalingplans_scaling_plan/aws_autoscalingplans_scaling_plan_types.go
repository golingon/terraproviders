// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_autoscalingplans_scaling_plan

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ApplicationSource struct {
	// CloudformationStackArn: string, optional
	CloudformationStackArn terra.StringValue `hcl:"cloudformation_stack_arn,attr"`
	// ApplicationSourceTagFilter: min=0,max=50
	TagFilter []ApplicationSourceTagFilter `hcl:"tag_filter,block" validate:"min=0,max=50"`
}

type ApplicationSourceTagFilter struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type ScalingInstruction struct {
	// DisableDynamicScaling: bool, optional
	DisableDynamicScaling terra.BoolValue `hcl:"disable_dynamic_scaling,attr"`
	// MaxCapacity: number, required
	MaxCapacity terra.NumberValue `hcl:"max_capacity,attr" validate:"required"`
	// MinCapacity: number, required
	MinCapacity terra.NumberValue `hcl:"min_capacity,attr" validate:"required"`
	// PredictiveScalingMaxCapacityBehavior: string, optional
	PredictiveScalingMaxCapacityBehavior terra.StringValue `hcl:"predictive_scaling_max_capacity_behavior,attr"`
	// PredictiveScalingMaxCapacityBuffer: number, optional
	PredictiveScalingMaxCapacityBuffer terra.NumberValue `hcl:"predictive_scaling_max_capacity_buffer,attr"`
	// PredictiveScalingMode: string, optional
	PredictiveScalingMode terra.StringValue `hcl:"predictive_scaling_mode,attr"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// ScalableDimension: string, required
	ScalableDimension terra.StringValue `hcl:"scalable_dimension,attr" validate:"required"`
	// ScalingPolicyUpdateBehavior: string, optional
	ScalingPolicyUpdateBehavior terra.StringValue `hcl:"scaling_policy_update_behavior,attr"`
	// ScheduledActionBufferTime: number, optional
	ScheduledActionBufferTime terra.NumberValue `hcl:"scheduled_action_buffer_time,attr"`
	// ServiceNamespace: string, required
	ServiceNamespace terra.StringValue `hcl:"service_namespace,attr" validate:"required"`
	// ScalingInstructionCustomizedLoadMetricSpecification: optional
	CustomizedLoadMetricSpecification *ScalingInstructionCustomizedLoadMetricSpecification `hcl:"customized_load_metric_specification,block"`
	// ScalingInstructionPredefinedLoadMetricSpecification: optional
	PredefinedLoadMetricSpecification *ScalingInstructionPredefinedLoadMetricSpecification `hcl:"predefined_load_metric_specification,block"`
	// ScalingInstructionTargetTrackingConfiguration: min=1,max=10
	TargetTrackingConfiguration []ScalingInstructionTargetTrackingConfiguration `hcl:"target_tracking_configuration,block" validate:"min=1,max=10"`
}

type ScalingInstructionCustomizedLoadMetricSpecification struct {
	// Dimensions: map of string, optional
	Dimensions terra.MapValue[terra.StringValue] `hcl:"dimensions,attr"`
	// MetricName: string, required
	MetricName terra.StringValue `hcl:"metric_name,attr" validate:"required"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
	// Statistic: string, required
	Statistic terra.StringValue `hcl:"statistic,attr" validate:"required"`
	// Unit: string, optional
	Unit terra.StringValue `hcl:"unit,attr"`
}

type ScalingInstructionPredefinedLoadMetricSpecification struct {
	// PredefinedLoadMetricType: string, required
	PredefinedLoadMetricType terra.StringValue `hcl:"predefined_load_metric_type,attr" validate:"required"`
	// ResourceLabel: string, optional
	ResourceLabel terra.StringValue `hcl:"resource_label,attr"`
}

type ScalingInstructionTargetTrackingConfiguration struct {
	// DisableScaleIn: bool, optional
	DisableScaleIn terra.BoolValue `hcl:"disable_scale_in,attr"`
	// EstimatedInstanceWarmup: number, optional
	EstimatedInstanceWarmup terra.NumberValue `hcl:"estimated_instance_warmup,attr"`
	// ScaleInCooldown: number, optional
	ScaleInCooldown terra.NumberValue `hcl:"scale_in_cooldown,attr"`
	// ScaleOutCooldown: number, optional
	ScaleOutCooldown terra.NumberValue `hcl:"scale_out_cooldown,attr"`
	// TargetValue: number, required
	TargetValue terra.NumberValue `hcl:"target_value,attr" validate:"required"`
	// ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification: optional
	CustomizedScalingMetricSpecification *ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification `hcl:"customized_scaling_metric_specification,block"`
	// ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification: optional
	PredefinedScalingMetricSpecification *ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification `hcl:"predefined_scaling_metric_specification,block"`
}

type ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification struct {
	// Dimensions: map of string, optional
	Dimensions terra.MapValue[terra.StringValue] `hcl:"dimensions,attr"`
	// MetricName: string, required
	MetricName terra.StringValue `hcl:"metric_name,attr" validate:"required"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
	// Statistic: string, required
	Statistic terra.StringValue `hcl:"statistic,attr" validate:"required"`
	// Unit: string, optional
	Unit terra.StringValue `hcl:"unit,attr"`
}

type ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification struct {
	// PredefinedScalingMetricType: string, required
	PredefinedScalingMetricType terra.StringValue `hcl:"predefined_scaling_metric_type,attr" validate:"required"`
	// ResourceLabel: string, optional
	ResourceLabel terra.StringValue `hcl:"resource_label,attr"`
}

type ApplicationSourceAttributes struct {
	ref terra.Reference
}

func (as ApplicationSourceAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as ApplicationSourceAttributes) InternalWithRef(ref terra.Reference) ApplicationSourceAttributes {
	return ApplicationSourceAttributes{ref: ref}
}

func (as ApplicationSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as ApplicationSourceAttributes) CloudformationStackArn() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("cloudformation_stack_arn"))
}

func (as ApplicationSourceAttributes) TagFilter() terra.SetValue[ApplicationSourceTagFilterAttributes] {
	return terra.ReferenceAsSet[ApplicationSourceTagFilterAttributes](as.ref.Append("tag_filter"))
}

type ApplicationSourceTagFilterAttributes struct {
	ref terra.Reference
}

func (tf ApplicationSourceTagFilterAttributes) InternalRef() (terra.Reference, error) {
	return tf.ref, nil
}

func (tf ApplicationSourceTagFilterAttributes) InternalWithRef(ref terra.Reference) ApplicationSourceTagFilterAttributes {
	return ApplicationSourceTagFilterAttributes{ref: ref}
}

func (tf ApplicationSourceTagFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tf.ref.InternalTokens()
}

func (tf ApplicationSourceTagFilterAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(tf.ref.Append("key"))
}

func (tf ApplicationSourceTagFilterAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tf.ref.Append("values"))
}

type ScalingInstructionAttributes struct {
	ref terra.Reference
}

func (si ScalingInstructionAttributes) InternalRef() (terra.Reference, error) {
	return si.ref, nil
}

func (si ScalingInstructionAttributes) InternalWithRef(ref terra.Reference) ScalingInstructionAttributes {
	return ScalingInstructionAttributes{ref: ref}
}

func (si ScalingInstructionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return si.ref.InternalTokens()
}

func (si ScalingInstructionAttributes) DisableDynamicScaling() terra.BoolValue {
	return terra.ReferenceAsBool(si.ref.Append("disable_dynamic_scaling"))
}

func (si ScalingInstructionAttributes) MaxCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(si.ref.Append("max_capacity"))
}

func (si ScalingInstructionAttributes) MinCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(si.ref.Append("min_capacity"))
}

func (si ScalingInstructionAttributes) PredictiveScalingMaxCapacityBehavior() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("predictive_scaling_max_capacity_behavior"))
}

func (si ScalingInstructionAttributes) PredictiveScalingMaxCapacityBuffer() terra.NumberValue {
	return terra.ReferenceAsNumber(si.ref.Append("predictive_scaling_max_capacity_buffer"))
}

func (si ScalingInstructionAttributes) PredictiveScalingMode() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("predictive_scaling_mode"))
}

func (si ScalingInstructionAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("resource_id"))
}

func (si ScalingInstructionAttributes) ScalableDimension() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("scalable_dimension"))
}

func (si ScalingInstructionAttributes) ScalingPolicyUpdateBehavior() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("scaling_policy_update_behavior"))
}

func (si ScalingInstructionAttributes) ScheduledActionBufferTime() terra.NumberValue {
	return terra.ReferenceAsNumber(si.ref.Append("scheduled_action_buffer_time"))
}

func (si ScalingInstructionAttributes) ServiceNamespace() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("service_namespace"))
}

func (si ScalingInstructionAttributes) CustomizedLoadMetricSpecification() terra.ListValue[ScalingInstructionCustomizedLoadMetricSpecificationAttributes] {
	return terra.ReferenceAsList[ScalingInstructionCustomizedLoadMetricSpecificationAttributes](si.ref.Append("customized_load_metric_specification"))
}

func (si ScalingInstructionAttributes) PredefinedLoadMetricSpecification() terra.ListValue[ScalingInstructionPredefinedLoadMetricSpecificationAttributes] {
	return terra.ReferenceAsList[ScalingInstructionPredefinedLoadMetricSpecificationAttributes](si.ref.Append("predefined_load_metric_specification"))
}

func (si ScalingInstructionAttributes) TargetTrackingConfiguration() terra.SetValue[ScalingInstructionTargetTrackingConfigurationAttributes] {
	return terra.ReferenceAsSet[ScalingInstructionTargetTrackingConfigurationAttributes](si.ref.Append("target_tracking_configuration"))
}

type ScalingInstructionCustomizedLoadMetricSpecificationAttributes struct {
	ref terra.Reference
}

func (clms ScalingInstructionCustomizedLoadMetricSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return clms.ref, nil
}

func (clms ScalingInstructionCustomizedLoadMetricSpecificationAttributes) InternalWithRef(ref terra.Reference) ScalingInstructionCustomizedLoadMetricSpecificationAttributes {
	return ScalingInstructionCustomizedLoadMetricSpecificationAttributes{ref: ref}
}

func (clms ScalingInstructionCustomizedLoadMetricSpecificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return clms.ref.InternalTokens()
}

func (clms ScalingInstructionCustomizedLoadMetricSpecificationAttributes) Dimensions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](clms.ref.Append("dimensions"))
}

func (clms ScalingInstructionCustomizedLoadMetricSpecificationAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(clms.ref.Append("metric_name"))
}

func (clms ScalingInstructionCustomizedLoadMetricSpecificationAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(clms.ref.Append("namespace"))
}

func (clms ScalingInstructionCustomizedLoadMetricSpecificationAttributes) Statistic() terra.StringValue {
	return terra.ReferenceAsString(clms.ref.Append("statistic"))
}

func (clms ScalingInstructionCustomizedLoadMetricSpecificationAttributes) Unit() terra.StringValue {
	return terra.ReferenceAsString(clms.ref.Append("unit"))
}

type ScalingInstructionPredefinedLoadMetricSpecificationAttributes struct {
	ref terra.Reference
}

func (plms ScalingInstructionPredefinedLoadMetricSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return plms.ref, nil
}

func (plms ScalingInstructionPredefinedLoadMetricSpecificationAttributes) InternalWithRef(ref terra.Reference) ScalingInstructionPredefinedLoadMetricSpecificationAttributes {
	return ScalingInstructionPredefinedLoadMetricSpecificationAttributes{ref: ref}
}

func (plms ScalingInstructionPredefinedLoadMetricSpecificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return plms.ref.InternalTokens()
}

func (plms ScalingInstructionPredefinedLoadMetricSpecificationAttributes) PredefinedLoadMetricType() terra.StringValue {
	return terra.ReferenceAsString(plms.ref.Append("predefined_load_metric_type"))
}

func (plms ScalingInstructionPredefinedLoadMetricSpecificationAttributes) ResourceLabel() terra.StringValue {
	return terra.ReferenceAsString(plms.ref.Append("resource_label"))
}

type ScalingInstructionTargetTrackingConfigurationAttributes struct {
	ref terra.Reference
}

func (ttc ScalingInstructionTargetTrackingConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ttc.ref, nil
}

func (ttc ScalingInstructionTargetTrackingConfigurationAttributes) InternalWithRef(ref terra.Reference) ScalingInstructionTargetTrackingConfigurationAttributes {
	return ScalingInstructionTargetTrackingConfigurationAttributes{ref: ref}
}

func (ttc ScalingInstructionTargetTrackingConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ttc.ref.InternalTokens()
}

func (ttc ScalingInstructionTargetTrackingConfigurationAttributes) DisableScaleIn() terra.BoolValue {
	return terra.ReferenceAsBool(ttc.ref.Append("disable_scale_in"))
}

func (ttc ScalingInstructionTargetTrackingConfigurationAttributes) EstimatedInstanceWarmup() terra.NumberValue {
	return terra.ReferenceAsNumber(ttc.ref.Append("estimated_instance_warmup"))
}

func (ttc ScalingInstructionTargetTrackingConfigurationAttributes) ScaleInCooldown() terra.NumberValue {
	return terra.ReferenceAsNumber(ttc.ref.Append("scale_in_cooldown"))
}

func (ttc ScalingInstructionTargetTrackingConfigurationAttributes) ScaleOutCooldown() terra.NumberValue {
	return terra.ReferenceAsNumber(ttc.ref.Append("scale_out_cooldown"))
}

func (ttc ScalingInstructionTargetTrackingConfigurationAttributes) TargetValue() terra.NumberValue {
	return terra.ReferenceAsNumber(ttc.ref.Append("target_value"))
}

func (ttc ScalingInstructionTargetTrackingConfigurationAttributes) CustomizedScalingMetricSpecification() terra.ListValue[ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes] {
	return terra.ReferenceAsList[ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes](ttc.ref.Append("customized_scaling_metric_specification"))
}

func (ttc ScalingInstructionTargetTrackingConfigurationAttributes) PredefinedScalingMetricSpecification() terra.ListValue[ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationAttributes] {
	return terra.ReferenceAsList[ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationAttributes](ttc.ref.Append("predefined_scaling_metric_specification"))
}

type ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes struct {
	ref terra.Reference
}

func (csms ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return csms.ref, nil
}

func (csms ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes) InternalWithRef(ref terra.Reference) ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes {
	return ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes{ref: ref}
}

func (csms ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return csms.ref.InternalTokens()
}

func (csms ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes) Dimensions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](csms.ref.Append("dimensions"))
}

func (csms ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(csms.ref.Append("metric_name"))
}

func (csms ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(csms.ref.Append("namespace"))
}

func (csms ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes) Statistic() terra.StringValue {
	return terra.ReferenceAsString(csms.ref.Append("statistic"))
}

func (csms ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationAttributes) Unit() terra.StringValue {
	return terra.ReferenceAsString(csms.ref.Append("unit"))
}

type ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationAttributes struct {
	ref terra.Reference
}

func (psms ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return psms.ref, nil
}

func (psms ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationAttributes) InternalWithRef(ref terra.Reference) ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationAttributes {
	return ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationAttributes{ref: ref}
}

func (psms ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return psms.ref.InternalTokens()
}

func (psms ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationAttributes) PredefinedScalingMetricType() terra.StringValue {
	return terra.ReferenceAsString(psms.ref.Append("predefined_scaling_metric_type"))
}

func (psms ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationAttributes) ResourceLabel() terra.StringValue {
	return terra.ReferenceAsString(psms.ref.Append("resource_label"))
}

type ApplicationSourceState struct {
	CloudformationStackArn string                            `json:"cloudformation_stack_arn"`
	TagFilter              []ApplicationSourceTagFilterState `json:"tag_filter"`
}

type ApplicationSourceTagFilterState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type ScalingInstructionState struct {
	DisableDynamicScaling                bool                                                       `json:"disable_dynamic_scaling"`
	MaxCapacity                          float64                                                    `json:"max_capacity"`
	MinCapacity                          float64                                                    `json:"min_capacity"`
	PredictiveScalingMaxCapacityBehavior string                                                     `json:"predictive_scaling_max_capacity_behavior"`
	PredictiveScalingMaxCapacityBuffer   float64                                                    `json:"predictive_scaling_max_capacity_buffer"`
	PredictiveScalingMode                string                                                     `json:"predictive_scaling_mode"`
	ResourceId                           string                                                     `json:"resource_id"`
	ScalableDimension                    string                                                     `json:"scalable_dimension"`
	ScalingPolicyUpdateBehavior          string                                                     `json:"scaling_policy_update_behavior"`
	ScheduledActionBufferTime            float64                                                    `json:"scheduled_action_buffer_time"`
	ServiceNamespace                     string                                                     `json:"service_namespace"`
	CustomizedLoadMetricSpecification    []ScalingInstructionCustomizedLoadMetricSpecificationState `json:"customized_load_metric_specification"`
	PredefinedLoadMetricSpecification    []ScalingInstructionPredefinedLoadMetricSpecificationState `json:"predefined_load_metric_specification"`
	TargetTrackingConfiguration          []ScalingInstructionTargetTrackingConfigurationState       `json:"target_tracking_configuration"`
}

type ScalingInstructionCustomizedLoadMetricSpecificationState struct {
	Dimensions map[string]string `json:"dimensions"`
	MetricName string            `json:"metric_name"`
	Namespace  string            `json:"namespace"`
	Statistic  string            `json:"statistic"`
	Unit       string            `json:"unit"`
}

type ScalingInstructionPredefinedLoadMetricSpecificationState struct {
	PredefinedLoadMetricType string `json:"predefined_load_metric_type"`
	ResourceLabel            string `json:"resource_label"`
}

type ScalingInstructionTargetTrackingConfigurationState struct {
	DisableScaleIn                       bool                                                                                     `json:"disable_scale_in"`
	EstimatedInstanceWarmup              float64                                                                                  `json:"estimated_instance_warmup"`
	ScaleInCooldown                      float64                                                                                  `json:"scale_in_cooldown"`
	ScaleOutCooldown                     float64                                                                                  `json:"scale_out_cooldown"`
	TargetValue                          float64                                                                                  `json:"target_value"`
	CustomizedScalingMetricSpecification []ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationState `json:"customized_scaling_metric_specification"`
	PredefinedScalingMetricSpecification []ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationState `json:"predefined_scaling_metric_specification"`
}

type ScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationState struct {
	Dimensions map[string]string `json:"dimensions"`
	MetricName string            `json:"metric_name"`
	Namespace  string            `json:"namespace"`
	Statistic  string            `json:"statistic"`
	Unit       string            `json:"unit"`
}

type ScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationState struct {
	PredefinedScalingMetricType string `json:"predefined_scaling_metric_type"`
	ResourceLabel               string `json:"resource_label"`
}
