// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appautoscalingpolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type StepScalingPolicyConfiguration struct {
	// AdjustmentType: string, optional
	AdjustmentType terra.StringValue `hcl:"adjustment_type,attr"`
	// Cooldown: number, optional
	Cooldown terra.NumberValue `hcl:"cooldown,attr"`
	// MetricAggregationType: string, optional
	MetricAggregationType terra.StringValue `hcl:"metric_aggregation_type,attr"`
	// MinAdjustmentMagnitude: number, optional
	MinAdjustmentMagnitude terra.NumberValue `hcl:"min_adjustment_magnitude,attr"`
	// StepAdjustment: min=0
	StepAdjustment []StepAdjustment `hcl:"step_adjustment,block" validate:"min=0"`
}

type StepAdjustment struct {
	// MetricIntervalLowerBound: string, optional
	MetricIntervalLowerBound terra.StringValue `hcl:"metric_interval_lower_bound,attr"`
	// MetricIntervalUpperBound: string, optional
	MetricIntervalUpperBound terra.StringValue `hcl:"metric_interval_upper_bound,attr"`
	// ScalingAdjustment: number, required
	ScalingAdjustment terra.NumberValue `hcl:"scaling_adjustment,attr" validate:"required"`
}

type TargetTrackingScalingPolicyConfiguration struct {
	// DisableScaleIn: bool, optional
	DisableScaleIn terra.BoolValue `hcl:"disable_scale_in,attr"`
	// ScaleInCooldown: number, optional
	ScaleInCooldown terra.NumberValue `hcl:"scale_in_cooldown,attr"`
	// ScaleOutCooldown: number, optional
	ScaleOutCooldown terra.NumberValue `hcl:"scale_out_cooldown,attr"`
	// TargetValue: number, required
	TargetValue terra.NumberValue `hcl:"target_value,attr" validate:"required"`
	// CustomizedMetricSpecification: optional
	CustomizedMetricSpecification *CustomizedMetricSpecification `hcl:"customized_metric_specification,block"`
	// PredefinedMetricSpecification: optional
	PredefinedMetricSpecification *PredefinedMetricSpecification `hcl:"predefined_metric_specification,block"`
}

type CustomizedMetricSpecification struct {
	// MetricName: string, required
	MetricName terra.StringValue `hcl:"metric_name,attr" validate:"required"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
	// Statistic: string, required
	Statistic terra.StringValue `hcl:"statistic,attr" validate:"required"`
	// Unit: string, optional
	Unit terra.StringValue `hcl:"unit,attr"`
	// Dimensions: min=0
	Dimensions []Dimensions `hcl:"dimensions,block" validate:"min=0"`
}

type Dimensions struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type PredefinedMetricSpecification struct {
	// PredefinedMetricType: string, required
	PredefinedMetricType terra.StringValue `hcl:"predefined_metric_type,attr" validate:"required"`
	// ResourceLabel: string, optional
	ResourceLabel terra.StringValue `hcl:"resource_label,attr"`
}

type StepScalingPolicyConfigurationAttributes struct {
	ref terra.Reference
}

func (sspc StepScalingPolicyConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return sspc.ref, nil
}

func (sspc StepScalingPolicyConfigurationAttributes) InternalWithRef(ref terra.Reference) StepScalingPolicyConfigurationAttributes {
	return StepScalingPolicyConfigurationAttributes{ref: ref}
}

func (sspc StepScalingPolicyConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return sspc.ref.InternalTokens()
}

func (sspc StepScalingPolicyConfigurationAttributes) AdjustmentType() terra.StringValue {
	return terra.ReferenceAsString(sspc.ref.Append("adjustment_type"))
}

func (sspc StepScalingPolicyConfigurationAttributes) Cooldown() terra.NumberValue {
	return terra.ReferenceAsNumber(sspc.ref.Append("cooldown"))
}

func (sspc StepScalingPolicyConfigurationAttributes) MetricAggregationType() terra.StringValue {
	return terra.ReferenceAsString(sspc.ref.Append("metric_aggregation_type"))
}

func (sspc StepScalingPolicyConfigurationAttributes) MinAdjustmentMagnitude() terra.NumberValue {
	return terra.ReferenceAsNumber(sspc.ref.Append("min_adjustment_magnitude"))
}

func (sspc StepScalingPolicyConfigurationAttributes) StepAdjustment() terra.SetValue[StepAdjustmentAttributes] {
	return terra.ReferenceAsSet[StepAdjustmentAttributes](sspc.ref.Append("step_adjustment"))
}

type StepAdjustmentAttributes struct {
	ref terra.Reference
}

func (sa StepAdjustmentAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa StepAdjustmentAttributes) InternalWithRef(ref terra.Reference) StepAdjustmentAttributes {
	return StepAdjustmentAttributes{ref: ref}
}

func (sa StepAdjustmentAttributes) InternalTokens() hclwrite.Tokens {
	return sa.ref.InternalTokens()
}

func (sa StepAdjustmentAttributes) MetricIntervalLowerBound() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("metric_interval_lower_bound"))
}

func (sa StepAdjustmentAttributes) MetricIntervalUpperBound() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("metric_interval_upper_bound"))
}

func (sa StepAdjustmentAttributes) ScalingAdjustment() terra.NumberValue {
	return terra.ReferenceAsNumber(sa.ref.Append("scaling_adjustment"))
}

type TargetTrackingScalingPolicyConfigurationAttributes struct {
	ref terra.Reference
}

func (ttspc TargetTrackingScalingPolicyConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ttspc.ref, nil
}

func (ttspc TargetTrackingScalingPolicyConfigurationAttributes) InternalWithRef(ref terra.Reference) TargetTrackingScalingPolicyConfigurationAttributes {
	return TargetTrackingScalingPolicyConfigurationAttributes{ref: ref}
}

func (ttspc TargetTrackingScalingPolicyConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return ttspc.ref.InternalTokens()
}

func (ttspc TargetTrackingScalingPolicyConfigurationAttributes) DisableScaleIn() terra.BoolValue {
	return terra.ReferenceAsBool(ttspc.ref.Append("disable_scale_in"))
}

func (ttspc TargetTrackingScalingPolicyConfigurationAttributes) ScaleInCooldown() terra.NumberValue {
	return terra.ReferenceAsNumber(ttspc.ref.Append("scale_in_cooldown"))
}

func (ttspc TargetTrackingScalingPolicyConfigurationAttributes) ScaleOutCooldown() terra.NumberValue {
	return terra.ReferenceAsNumber(ttspc.ref.Append("scale_out_cooldown"))
}

func (ttspc TargetTrackingScalingPolicyConfigurationAttributes) TargetValue() terra.NumberValue {
	return terra.ReferenceAsNumber(ttspc.ref.Append("target_value"))
}

func (ttspc TargetTrackingScalingPolicyConfigurationAttributes) CustomizedMetricSpecification() terra.ListValue[CustomizedMetricSpecificationAttributes] {
	return terra.ReferenceAsList[CustomizedMetricSpecificationAttributes](ttspc.ref.Append("customized_metric_specification"))
}

func (ttspc TargetTrackingScalingPolicyConfigurationAttributes) PredefinedMetricSpecification() terra.ListValue[PredefinedMetricSpecificationAttributes] {
	return terra.ReferenceAsList[PredefinedMetricSpecificationAttributes](ttspc.ref.Append("predefined_metric_specification"))
}

type CustomizedMetricSpecificationAttributes struct {
	ref terra.Reference
}

func (cms CustomizedMetricSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return cms.ref, nil
}

func (cms CustomizedMetricSpecificationAttributes) InternalWithRef(ref terra.Reference) CustomizedMetricSpecificationAttributes {
	return CustomizedMetricSpecificationAttributes{ref: ref}
}

func (cms CustomizedMetricSpecificationAttributes) InternalTokens() hclwrite.Tokens {
	return cms.ref.InternalTokens()
}

func (cms CustomizedMetricSpecificationAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("metric_name"))
}

func (cms CustomizedMetricSpecificationAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("namespace"))
}

func (cms CustomizedMetricSpecificationAttributes) Statistic() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("statistic"))
}

func (cms CustomizedMetricSpecificationAttributes) Unit() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("unit"))
}

func (cms CustomizedMetricSpecificationAttributes) Dimensions() terra.SetValue[DimensionsAttributes] {
	return terra.ReferenceAsSet[DimensionsAttributes](cms.ref.Append("dimensions"))
}

type DimensionsAttributes struct {
	ref terra.Reference
}

func (d DimensionsAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DimensionsAttributes) InternalWithRef(ref terra.Reference) DimensionsAttributes {
	return DimensionsAttributes{ref: ref}
}

func (d DimensionsAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d DimensionsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("name"))
}

func (d DimensionsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("value"))
}

type PredefinedMetricSpecificationAttributes struct {
	ref terra.Reference
}

func (pms PredefinedMetricSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return pms.ref, nil
}

func (pms PredefinedMetricSpecificationAttributes) InternalWithRef(ref terra.Reference) PredefinedMetricSpecificationAttributes {
	return PredefinedMetricSpecificationAttributes{ref: ref}
}

func (pms PredefinedMetricSpecificationAttributes) InternalTokens() hclwrite.Tokens {
	return pms.ref.InternalTokens()
}

func (pms PredefinedMetricSpecificationAttributes) PredefinedMetricType() terra.StringValue {
	return terra.ReferenceAsString(pms.ref.Append("predefined_metric_type"))
}

func (pms PredefinedMetricSpecificationAttributes) ResourceLabel() terra.StringValue {
	return terra.ReferenceAsString(pms.ref.Append("resource_label"))
}

type StepScalingPolicyConfigurationState struct {
	AdjustmentType         string                `json:"adjustment_type"`
	Cooldown               float64               `json:"cooldown"`
	MetricAggregationType  string                `json:"metric_aggregation_type"`
	MinAdjustmentMagnitude float64               `json:"min_adjustment_magnitude"`
	StepAdjustment         []StepAdjustmentState `json:"step_adjustment"`
}

type StepAdjustmentState struct {
	MetricIntervalLowerBound string  `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string  `json:"metric_interval_upper_bound"`
	ScalingAdjustment        float64 `json:"scaling_adjustment"`
}

type TargetTrackingScalingPolicyConfigurationState struct {
	DisableScaleIn                bool                                 `json:"disable_scale_in"`
	ScaleInCooldown               float64                              `json:"scale_in_cooldown"`
	ScaleOutCooldown              float64                              `json:"scale_out_cooldown"`
	TargetValue                   float64                              `json:"target_value"`
	CustomizedMetricSpecification []CustomizedMetricSpecificationState `json:"customized_metric_specification"`
	PredefinedMetricSpecification []PredefinedMetricSpecificationState `json:"predefined_metric_specification"`
}

type CustomizedMetricSpecificationState struct {
	MetricName string            `json:"metric_name"`
	Namespace  string            `json:"namespace"`
	Statistic  string            `json:"statistic"`
	Unit       string            `json:"unit"`
	Dimensions []DimensionsState `json:"dimensions"`
}

type DimensionsState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PredefinedMetricSpecificationState struct {
	PredefinedMetricType string `json:"predefined_metric_type"`
	ResourceLabel        string `json:"resource_label"`
}
