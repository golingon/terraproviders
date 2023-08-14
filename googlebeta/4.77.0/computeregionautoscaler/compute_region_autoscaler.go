// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package computeregionautoscaler

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AutoscalingPolicy struct {
	// CooldownPeriod: number, optional
	CooldownPeriod terra.NumberValue `hcl:"cooldown_period,attr"`
	// MaxReplicas: number, required
	MaxReplicas terra.NumberValue `hcl:"max_replicas,attr" validate:"required"`
	// MinReplicas: number, required
	MinReplicas terra.NumberValue `hcl:"min_replicas,attr" validate:"required"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// CpuUtilization: optional
	CpuUtilization *CpuUtilization `hcl:"cpu_utilization,block"`
	// LoadBalancingUtilization: optional
	LoadBalancingUtilization *LoadBalancingUtilization `hcl:"load_balancing_utilization,block"`
	// Metric: min=0
	Metric []Metric `hcl:"metric,block" validate:"min=0"`
	// ScaleDownControl: optional
	ScaleDownControl *ScaleDownControl `hcl:"scale_down_control,block"`
	// ScaleInControl: optional
	ScaleInControl *ScaleInControl `hcl:"scale_in_control,block"`
	// ScalingSchedules: min=0
	ScalingSchedules []ScalingSchedules `hcl:"scaling_schedules,block" validate:"min=0"`
}

type CpuUtilization struct {
	// PredictiveMethod: string, optional
	PredictiveMethod terra.StringValue `hcl:"predictive_method,attr"`
	// Target: number, required
	Target terra.NumberValue `hcl:"target,attr" validate:"required"`
}

type LoadBalancingUtilization struct {
	// Target: number, required
	Target terra.NumberValue `hcl:"target,attr" validate:"required"`
}

type Metric struct {
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SingleInstanceAssignment: number, optional
	SingleInstanceAssignment terra.NumberValue `hcl:"single_instance_assignment,attr"`
	// Target: number, optional
	Target terra.NumberValue `hcl:"target,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type ScaleDownControl struct {
	// TimeWindowSec: number, optional
	TimeWindowSec terra.NumberValue `hcl:"time_window_sec,attr"`
	// MaxScaledDownReplicas: optional
	MaxScaledDownReplicas *MaxScaledDownReplicas `hcl:"max_scaled_down_replicas,block"`
}

type MaxScaledDownReplicas struct {
	// Fixed: number, optional
	Fixed terra.NumberValue `hcl:"fixed,attr"`
	// Percent: number, optional
	Percent terra.NumberValue `hcl:"percent,attr"`
}

type ScaleInControl struct {
	// TimeWindowSec: number, optional
	TimeWindowSec terra.NumberValue `hcl:"time_window_sec,attr"`
	// MaxScaledInReplicas: optional
	MaxScaledInReplicas *MaxScaledInReplicas `hcl:"max_scaled_in_replicas,block"`
}

type MaxScaledInReplicas struct {
	// Fixed: number, optional
	Fixed terra.NumberValue `hcl:"fixed,attr"`
	// Percent: number, optional
	Percent terra.NumberValue `hcl:"percent,attr"`
}

type ScalingSchedules struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// DurationSec: number, required
	DurationSec terra.NumberValue `hcl:"duration_sec,attr" validate:"required"`
	// MinRequiredReplicas: number, required
	MinRequiredReplicas terra.NumberValue `hcl:"min_required_replicas,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Schedule: string, required
	Schedule terra.StringValue `hcl:"schedule,attr" validate:"required"`
	// TimeZone: string, optional
	TimeZone terra.StringValue `hcl:"time_zone,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AutoscalingPolicyAttributes struct {
	ref terra.Reference
}

func (ap AutoscalingPolicyAttributes) InternalRef() (terra.Reference, error) {
	return ap.ref, nil
}

func (ap AutoscalingPolicyAttributes) InternalWithRef(ref terra.Reference) AutoscalingPolicyAttributes {
	return AutoscalingPolicyAttributes{ref: ref}
}

func (ap AutoscalingPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ap.ref.InternalTokens()
}

func (ap AutoscalingPolicyAttributes) CooldownPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(ap.ref.Append("cooldown_period"))
}

func (ap AutoscalingPolicyAttributes) MaxReplicas() terra.NumberValue {
	return terra.ReferenceAsNumber(ap.ref.Append("max_replicas"))
}

func (ap AutoscalingPolicyAttributes) MinReplicas() terra.NumberValue {
	return terra.ReferenceAsNumber(ap.ref.Append("min_replicas"))
}

func (ap AutoscalingPolicyAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("mode"))
}

func (ap AutoscalingPolicyAttributes) CpuUtilization() terra.ListValue[CpuUtilizationAttributes] {
	return terra.ReferenceAsList[CpuUtilizationAttributes](ap.ref.Append("cpu_utilization"))
}

func (ap AutoscalingPolicyAttributes) LoadBalancingUtilization() terra.ListValue[LoadBalancingUtilizationAttributes] {
	return terra.ReferenceAsList[LoadBalancingUtilizationAttributes](ap.ref.Append("load_balancing_utilization"))
}

func (ap AutoscalingPolicyAttributes) Metric() terra.ListValue[MetricAttributes] {
	return terra.ReferenceAsList[MetricAttributes](ap.ref.Append("metric"))
}

func (ap AutoscalingPolicyAttributes) ScaleDownControl() terra.ListValue[ScaleDownControlAttributes] {
	return terra.ReferenceAsList[ScaleDownControlAttributes](ap.ref.Append("scale_down_control"))
}

func (ap AutoscalingPolicyAttributes) ScaleInControl() terra.ListValue[ScaleInControlAttributes] {
	return terra.ReferenceAsList[ScaleInControlAttributes](ap.ref.Append("scale_in_control"))
}

func (ap AutoscalingPolicyAttributes) ScalingSchedules() terra.SetValue[ScalingSchedulesAttributes] {
	return terra.ReferenceAsSet[ScalingSchedulesAttributes](ap.ref.Append("scaling_schedules"))
}

type CpuUtilizationAttributes struct {
	ref terra.Reference
}

func (cu CpuUtilizationAttributes) InternalRef() (terra.Reference, error) {
	return cu.ref, nil
}

func (cu CpuUtilizationAttributes) InternalWithRef(ref terra.Reference) CpuUtilizationAttributes {
	return CpuUtilizationAttributes{ref: ref}
}

func (cu CpuUtilizationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cu.ref.InternalTokens()
}

func (cu CpuUtilizationAttributes) PredictiveMethod() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("predictive_method"))
}

func (cu CpuUtilizationAttributes) Target() terra.NumberValue {
	return terra.ReferenceAsNumber(cu.ref.Append("target"))
}

type LoadBalancingUtilizationAttributes struct {
	ref terra.Reference
}

func (lbu LoadBalancingUtilizationAttributes) InternalRef() (terra.Reference, error) {
	return lbu.ref, nil
}

func (lbu LoadBalancingUtilizationAttributes) InternalWithRef(ref terra.Reference) LoadBalancingUtilizationAttributes {
	return LoadBalancingUtilizationAttributes{ref: ref}
}

func (lbu LoadBalancingUtilizationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lbu.ref.InternalTokens()
}

func (lbu LoadBalancingUtilizationAttributes) Target() terra.NumberValue {
	return terra.ReferenceAsNumber(lbu.ref.Append("target"))
}

type MetricAttributes struct {
	ref terra.Reference
}

func (m MetricAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MetricAttributes) InternalWithRef(ref terra.Reference) MetricAttributes {
	return MetricAttributes{ref: ref}
}

func (m MetricAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MetricAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("filter"))
}

func (m MetricAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("name"))
}

func (m MetricAttributes) SingleInstanceAssignment() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("single_instance_assignment"))
}

func (m MetricAttributes) Target() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("target"))
}

func (m MetricAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("type"))
}

type ScaleDownControlAttributes struct {
	ref terra.Reference
}

func (sdc ScaleDownControlAttributes) InternalRef() (terra.Reference, error) {
	return sdc.ref, nil
}

func (sdc ScaleDownControlAttributes) InternalWithRef(ref terra.Reference) ScaleDownControlAttributes {
	return ScaleDownControlAttributes{ref: ref}
}

func (sdc ScaleDownControlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sdc.ref.InternalTokens()
}

func (sdc ScaleDownControlAttributes) TimeWindowSec() terra.NumberValue {
	return terra.ReferenceAsNumber(sdc.ref.Append("time_window_sec"))
}

func (sdc ScaleDownControlAttributes) MaxScaledDownReplicas() terra.ListValue[MaxScaledDownReplicasAttributes] {
	return terra.ReferenceAsList[MaxScaledDownReplicasAttributes](sdc.ref.Append("max_scaled_down_replicas"))
}

type MaxScaledDownReplicasAttributes struct {
	ref terra.Reference
}

func (msdr MaxScaledDownReplicasAttributes) InternalRef() (terra.Reference, error) {
	return msdr.ref, nil
}

func (msdr MaxScaledDownReplicasAttributes) InternalWithRef(ref terra.Reference) MaxScaledDownReplicasAttributes {
	return MaxScaledDownReplicasAttributes{ref: ref}
}

func (msdr MaxScaledDownReplicasAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return msdr.ref.InternalTokens()
}

func (msdr MaxScaledDownReplicasAttributes) Fixed() terra.NumberValue {
	return terra.ReferenceAsNumber(msdr.ref.Append("fixed"))
}

func (msdr MaxScaledDownReplicasAttributes) Percent() terra.NumberValue {
	return terra.ReferenceAsNumber(msdr.ref.Append("percent"))
}

type ScaleInControlAttributes struct {
	ref terra.Reference
}

func (sic ScaleInControlAttributes) InternalRef() (terra.Reference, error) {
	return sic.ref, nil
}

func (sic ScaleInControlAttributes) InternalWithRef(ref terra.Reference) ScaleInControlAttributes {
	return ScaleInControlAttributes{ref: ref}
}

func (sic ScaleInControlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sic.ref.InternalTokens()
}

func (sic ScaleInControlAttributes) TimeWindowSec() terra.NumberValue {
	return terra.ReferenceAsNumber(sic.ref.Append("time_window_sec"))
}

func (sic ScaleInControlAttributes) MaxScaledInReplicas() terra.ListValue[MaxScaledInReplicasAttributes] {
	return terra.ReferenceAsList[MaxScaledInReplicasAttributes](sic.ref.Append("max_scaled_in_replicas"))
}

type MaxScaledInReplicasAttributes struct {
	ref terra.Reference
}

func (msir MaxScaledInReplicasAttributes) InternalRef() (terra.Reference, error) {
	return msir.ref, nil
}

func (msir MaxScaledInReplicasAttributes) InternalWithRef(ref terra.Reference) MaxScaledInReplicasAttributes {
	return MaxScaledInReplicasAttributes{ref: ref}
}

func (msir MaxScaledInReplicasAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return msir.ref.InternalTokens()
}

func (msir MaxScaledInReplicasAttributes) Fixed() terra.NumberValue {
	return terra.ReferenceAsNumber(msir.ref.Append("fixed"))
}

func (msir MaxScaledInReplicasAttributes) Percent() terra.NumberValue {
	return terra.ReferenceAsNumber(msir.ref.Append("percent"))
}

type ScalingSchedulesAttributes struct {
	ref terra.Reference
}

func (ss ScalingSchedulesAttributes) InternalRef() (terra.Reference, error) {
	return ss.ref, nil
}

func (ss ScalingSchedulesAttributes) InternalWithRef(ref terra.Reference) ScalingSchedulesAttributes {
	return ScalingSchedulesAttributes{ref: ref}
}

func (ss ScalingSchedulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ss.ref.InternalTokens()
}

func (ss ScalingSchedulesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("description"))
}

func (ss ScalingSchedulesAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("disabled"))
}

func (ss ScalingSchedulesAttributes) DurationSec() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("duration_sec"))
}

func (ss ScalingSchedulesAttributes) MinRequiredReplicas() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("min_required_replicas"))
}

func (ss ScalingSchedulesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

func (ss ScalingSchedulesAttributes) Schedule() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("schedule"))
}

func (ss ScalingSchedulesAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("time_zone"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AutoscalingPolicyState struct {
	CooldownPeriod           float64                         `json:"cooldown_period"`
	MaxReplicas              float64                         `json:"max_replicas"`
	MinReplicas              float64                         `json:"min_replicas"`
	Mode                     string                          `json:"mode"`
	CpuUtilization           []CpuUtilizationState           `json:"cpu_utilization"`
	LoadBalancingUtilization []LoadBalancingUtilizationState `json:"load_balancing_utilization"`
	Metric                   []MetricState                   `json:"metric"`
	ScaleDownControl         []ScaleDownControlState         `json:"scale_down_control"`
	ScaleInControl           []ScaleInControlState           `json:"scale_in_control"`
	ScalingSchedules         []ScalingSchedulesState         `json:"scaling_schedules"`
}

type CpuUtilizationState struct {
	PredictiveMethod string  `json:"predictive_method"`
	Target           float64 `json:"target"`
}

type LoadBalancingUtilizationState struct {
	Target float64 `json:"target"`
}

type MetricState struct {
	Filter                   string  `json:"filter"`
	Name                     string  `json:"name"`
	SingleInstanceAssignment float64 `json:"single_instance_assignment"`
	Target                   float64 `json:"target"`
	Type                     string  `json:"type"`
}

type ScaleDownControlState struct {
	TimeWindowSec         float64                      `json:"time_window_sec"`
	MaxScaledDownReplicas []MaxScaledDownReplicasState `json:"max_scaled_down_replicas"`
}

type MaxScaledDownReplicasState struct {
	Fixed   float64 `json:"fixed"`
	Percent float64 `json:"percent"`
}

type ScaleInControlState struct {
	TimeWindowSec       float64                    `json:"time_window_sec"`
	MaxScaledInReplicas []MaxScaledInReplicasState `json:"max_scaled_in_replicas"`
}

type MaxScaledInReplicasState struct {
	Fixed   float64 `json:"fixed"`
	Percent float64 `json:"percent"`
}

type ScalingSchedulesState struct {
	Description         string  `json:"description"`
	Disabled            bool    `json:"disabled"`
	DurationSec         float64 `json:"duration_sec"`
	MinRequiredReplicas float64 `json:"min_required_replicas"`
	Name                string  `json:"name"`
	Schedule            string  `json:"schedule"`
	TimeZone            string  `json:"time_zone"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
