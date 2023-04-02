// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package autoscalingplansscalingplan

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ApplicationSource struct {
	// CloudformationStackArn: string, optional
	CloudformationStackArn terra.StringValue `hcl:"cloudformation_stack_arn,attr"`
	// TagFilter: min=0,max=50
	TagFilter []TagFilter `hcl:"tag_filter,block" validate:"min=0,max=50"`
}

type TagFilter struct {
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
	// CustomizedLoadMetricSpecification: optional
	CustomizedLoadMetricSpecification *CustomizedLoadMetricSpecification `hcl:"customized_load_metric_specification,block"`
	// PredefinedLoadMetricSpecification: optional
	PredefinedLoadMetricSpecification *PredefinedLoadMetricSpecification `hcl:"predefined_load_metric_specification,block"`
	// TargetTrackingConfiguration: min=1,max=10
	TargetTrackingConfiguration []TargetTrackingConfiguration `hcl:"target_tracking_configuration,block" validate:"min=1,max=10"`
}

type CustomizedLoadMetricSpecification struct {
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

type PredefinedLoadMetricSpecification struct {
	// PredefinedLoadMetricType: string, required
	PredefinedLoadMetricType terra.StringValue `hcl:"predefined_load_metric_type,attr" validate:"required"`
	// ResourceLabel: string, optional
	ResourceLabel terra.StringValue `hcl:"resource_label,attr"`
}

type TargetTrackingConfiguration struct {
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
	// CustomizedScalingMetricSpecification: optional
	CustomizedScalingMetricSpecification *CustomizedScalingMetricSpecification `hcl:"customized_scaling_metric_specification,block"`
	// PredefinedScalingMetricSpecification: optional
	PredefinedScalingMetricSpecification *PredefinedScalingMetricSpecification `hcl:"predefined_scaling_metric_specification,block"`
}

type CustomizedScalingMetricSpecification struct {
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

type PredefinedScalingMetricSpecification struct {
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

func (as ApplicationSourceAttributes) InternalTokens() hclwrite.Tokens {
	return as.ref.InternalTokens()
}

func (as ApplicationSourceAttributes) CloudformationStackArn() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("cloudformation_stack_arn"))
}

func (as ApplicationSourceAttributes) TagFilter() terra.SetValue[TagFilterAttributes] {
	return terra.ReferenceAsSet[TagFilterAttributes](as.ref.Append("tag_filter"))
}

type TagFilterAttributes struct {
	ref terra.Reference
}

func (tf TagFilterAttributes) InternalRef() (terra.Reference, error) {
	return tf.ref, nil
}

func (tf TagFilterAttributes) InternalWithRef(ref terra.Reference) TagFilterAttributes {
	return TagFilterAttributes{ref: ref}
}

func (tf TagFilterAttributes) InternalTokens() hclwrite.Tokens {
	return tf.ref.InternalTokens()
}

func (tf TagFilterAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(tf.ref.Append("key"))
}

func (tf TagFilterAttributes) Values() terra.SetValue[terra.StringValue] {
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

func (si ScalingInstructionAttributes) InternalTokens() hclwrite.Tokens {
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

func (si ScalingInstructionAttributes) CustomizedLoadMetricSpecification() terra.ListValue[CustomizedLoadMetricSpecificationAttributes] {
	return terra.ReferenceAsList[CustomizedLoadMetricSpecificationAttributes](si.ref.Append("customized_load_metric_specification"))
}

func (si ScalingInstructionAttributes) PredefinedLoadMetricSpecification() terra.ListValue[PredefinedLoadMetricSpecificationAttributes] {
	return terra.ReferenceAsList[PredefinedLoadMetricSpecificationAttributes](si.ref.Append("predefined_load_metric_specification"))
}

func (si ScalingInstructionAttributes) TargetTrackingConfiguration() terra.SetValue[TargetTrackingConfigurationAttributes] {
	return terra.ReferenceAsSet[TargetTrackingConfigurationAttributes](si.ref.Append("target_tracking_configuration"))
}

type CustomizedLoadMetricSpecificationAttributes struct {
	ref terra.Reference
}

func (clms CustomizedLoadMetricSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return clms.ref, nil
}

func (clms CustomizedLoadMetricSpecificationAttributes) InternalWithRef(ref terra.Reference) CustomizedLoadMetricSpecificationAttributes {
	return CustomizedLoadMetricSpecificationAttributes{ref: ref}
}

func (clms CustomizedLoadMetricSpecificationAttributes) InternalTokens() hclwrite.Tokens {
	return clms.ref.InternalTokens()
}

func (clms CustomizedLoadMetricSpecificationAttributes) Dimensions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](clms.ref.Append("dimensions"))
}

func (clms CustomizedLoadMetricSpecificationAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(clms.ref.Append("metric_name"))
}

func (clms CustomizedLoadMetricSpecificationAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(clms.ref.Append("namespace"))
}

func (clms CustomizedLoadMetricSpecificationAttributes) Statistic() terra.StringValue {
	return terra.ReferenceAsString(clms.ref.Append("statistic"))
}

func (clms CustomizedLoadMetricSpecificationAttributes) Unit() terra.StringValue {
	return terra.ReferenceAsString(clms.ref.Append("unit"))
}

type PredefinedLoadMetricSpecificationAttributes struct {
	ref terra.Reference
}

func (plms PredefinedLoadMetricSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return plms.ref, nil
}

func (plms PredefinedLoadMetricSpecificationAttributes) InternalWithRef(ref terra.Reference) PredefinedLoadMetricSpecificationAttributes {
	return PredefinedLoadMetricSpecificationAttributes{ref: ref}
}

func (plms PredefinedLoadMetricSpecificationAttributes) InternalTokens() hclwrite.Tokens {
	return plms.ref.InternalTokens()
}

func (plms PredefinedLoadMetricSpecificationAttributes) PredefinedLoadMetricType() terra.StringValue {
	return terra.ReferenceAsString(plms.ref.Append("predefined_load_metric_type"))
}

func (plms PredefinedLoadMetricSpecificationAttributes) ResourceLabel() terra.StringValue {
	return terra.ReferenceAsString(plms.ref.Append("resource_label"))
}

type TargetTrackingConfigurationAttributes struct {
	ref terra.Reference
}

func (ttc TargetTrackingConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ttc.ref, nil
}

func (ttc TargetTrackingConfigurationAttributes) InternalWithRef(ref terra.Reference) TargetTrackingConfigurationAttributes {
	return TargetTrackingConfigurationAttributes{ref: ref}
}

func (ttc TargetTrackingConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return ttc.ref.InternalTokens()
}

func (ttc TargetTrackingConfigurationAttributes) DisableScaleIn() terra.BoolValue {
	return terra.ReferenceAsBool(ttc.ref.Append("disable_scale_in"))
}

func (ttc TargetTrackingConfigurationAttributes) EstimatedInstanceWarmup() terra.NumberValue {
	return terra.ReferenceAsNumber(ttc.ref.Append("estimated_instance_warmup"))
}

func (ttc TargetTrackingConfigurationAttributes) ScaleInCooldown() terra.NumberValue {
	return terra.ReferenceAsNumber(ttc.ref.Append("scale_in_cooldown"))
}

func (ttc TargetTrackingConfigurationAttributes) ScaleOutCooldown() terra.NumberValue {
	return terra.ReferenceAsNumber(ttc.ref.Append("scale_out_cooldown"))
}

func (ttc TargetTrackingConfigurationAttributes) TargetValue() terra.NumberValue {
	return terra.ReferenceAsNumber(ttc.ref.Append("target_value"))
}

func (ttc TargetTrackingConfigurationAttributes) CustomizedScalingMetricSpecification() terra.ListValue[CustomizedScalingMetricSpecificationAttributes] {
	return terra.ReferenceAsList[CustomizedScalingMetricSpecificationAttributes](ttc.ref.Append("customized_scaling_metric_specification"))
}

func (ttc TargetTrackingConfigurationAttributes) PredefinedScalingMetricSpecification() terra.ListValue[PredefinedScalingMetricSpecificationAttributes] {
	return terra.ReferenceAsList[PredefinedScalingMetricSpecificationAttributes](ttc.ref.Append("predefined_scaling_metric_specification"))
}

type CustomizedScalingMetricSpecificationAttributes struct {
	ref terra.Reference
}

func (csms CustomizedScalingMetricSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return csms.ref, nil
}

func (csms CustomizedScalingMetricSpecificationAttributes) InternalWithRef(ref terra.Reference) CustomizedScalingMetricSpecificationAttributes {
	return CustomizedScalingMetricSpecificationAttributes{ref: ref}
}

func (csms CustomizedScalingMetricSpecificationAttributes) InternalTokens() hclwrite.Tokens {
	return csms.ref.InternalTokens()
}

func (csms CustomizedScalingMetricSpecificationAttributes) Dimensions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](csms.ref.Append("dimensions"))
}

func (csms CustomizedScalingMetricSpecificationAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(csms.ref.Append("metric_name"))
}

func (csms CustomizedScalingMetricSpecificationAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(csms.ref.Append("namespace"))
}

func (csms CustomizedScalingMetricSpecificationAttributes) Statistic() terra.StringValue {
	return terra.ReferenceAsString(csms.ref.Append("statistic"))
}

func (csms CustomizedScalingMetricSpecificationAttributes) Unit() terra.StringValue {
	return terra.ReferenceAsString(csms.ref.Append("unit"))
}

type PredefinedScalingMetricSpecificationAttributes struct {
	ref terra.Reference
}

func (psms PredefinedScalingMetricSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return psms.ref, nil
}

func (psms PredefinedScalingMetricSpecificationAttributes) InternalWithRef(ref terra.Reference) PredefinedScalingMetricSpecificationAttributes {
	return PredefinedScalingMetricSpecificationAttributes{ref: ref}
}

func (psms PredefinedScalingMetricSpecificationAttributes) InternalTokens() hclwrite.Tokens {
	return psms.ref.InternalTokens()
}

func (psms PredefinedScalingMetricSpecificationAttributes) PredefinedScalingMetricType() terra.StringValue {
	return terra.ReferenceAsString(psms.ref.Append("predefined_scaling_metric_type"))
}

func (psms PredefinedScalingMetricSpecificationAttributes) ResourceLabel() terra.StringValue {
	return terra.ReferenceAsString(psms.ref.Append("resource_label"))
}

type ApplicationSourceState struct {
	CloudformationStackArn string           `json:"cloudformation_stack_arn"`
	TagFilter              []TagFilterState `json:"tag_filter"`
}

type TagFilterState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type ScalingInstructionState struct {
	DisableDynamicScaling                bool                                     `json:"disable_dynamic_scaling"`
	MaxCapacity                          float64                                  `json:"max_capacity"`
	MinCapacity                          float64                                  `json:"min_capacity"`
	PredictiveScalingMaxCapacityBehavior string                                   `json:"predictive_scaling_max_capacity_behavior"`
	PredictiveScalingMaxCapacityBuffer   float64                                  `json:"predictive_scaling_max_capacity_buffer"`
	PredictiveScalingMode                string                                   `json:"predictive_scaling_mode"`
	ResourceId                           string                                   `json:"resource_id"`
	ScalableDimension                    string                                   `json:"scalable_dimension"`
	ScalingPolicyUpdateBehavior          string                                   `json:"scaling_policy_update_behavior"`
	ScheduledActionBufferTime            float64                                  `json:"scheduled_action_buffer_time"`
	ServiceNamespace                     string                                   `json:"service_namespace"`
	CustomizedLoadMetricSpecification    []CustomizedLoadMetricSpecificationState `json:"customized_load_metric_specification"`
	PredefinedLoadMetricSpecification    []PredefinedLoadMetricSpecificationState `json:"predefined_load_metric_specification"`
	TargetTrackingConfiguration          []TargetTrackingConfigurationState       `json:"target_tracking_configuration"`
}

type CustomizedLoadMetricSpecificationState struct {
	Dimensions map[string]string `json:"dimensions"`
	MetricName string            `json:"metric_name"`
	Namespace  string            `json:"namespace"`
	Statistic  string            `json:"statistic"`
	Unit       string            `json:"unit"`
}

type PredefinedLoadMetricSpecificationState struct {
	PredefinedLoadMetricType string `json:"predefined_load_metric_type"`
	ResourceLabel            string `json:"resource_label"`
}

type TargetTrackingConfigurationState struct {
	DisableScaleIn                       bool                                        `json:"disable_scale_in"`
	EstimatedInstanceWarmup              float64                                     `json:"estimated_instance_warmup"`
	ScaleInCooldown                      float64                                     `json:"scale_in_cooldown"`
	ScaleOutCooldown                     float64                                     `json:"scale_out_cooldown"`
	TargetValue                          float64                                     `json:"target_value"`
	CustomizedScalingMetricSpecification []CustomizedScalingMetricSpecificationState `json:"customized_scaling_metric_specification"`
	PredefinedScalingMetricSpecification []PredefinedScalingMetricSpecificationState `json:"predefined_scaling_metric_specification"`
}

type CustomizedScalingMetricSpecificationState struct {
	Dimensions map[string]string `json:"dimensions"`
	MetricName string            `json:"metric_name"`
	Namespace  string            `json:"namespace"`
	Statistic  string            `json:"statistic"`
	Unit       string            `json:"unit"`
}

type PredefinedScalingMetricSpecificationState struct {
	PredefinedScalingMetricType string `json:"predefined_scaling_metric_type"`
	ResourceLabel               string `json:"resource_label"`
}
