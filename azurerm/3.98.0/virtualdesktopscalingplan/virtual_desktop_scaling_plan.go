// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package virtualdesktopscalingplan

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type HostPool struct {
	// HostpoolId: string, required
	HostpoolId terra.StringValue `hcl:"hostpool_id,attr" validate:"required"`
	// ScalingPlanEnabled: bool, required
	ScalingPlanEnabled terra.BoolValue `hcl:"scaling_plan_enabled,attr" validate:"required"`
}

type Schedule struct {
	// DaysOfWeek: set of string, required
	DaysOfWeek terra.SetValue[terra.StringValue] `hcl:"days_of_week,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OffPeakLoadBalancingAlgorithm: string, required
	OffPeakLoadBalancingAlgorithm terra.StringValue `hcl:"off_peak_load_balancing_algorithm,attr" validate:"required"`
	// OffPeakStartTime: string, required
	OffPeakStartTime terra.StringValue `hcl:"off_peak_start_time,attr" validate:"required"`
	// PeakLoadBalancingAlgorithm: string, required
	PeakLoadBalancingAlgorithm terra.StringValue `hcl:"peak_load_balancing_algorithm,attr" validate:"required"`
	// PeakStartTime: string, required
	PeakStartTime terra.StringValue `hcl:"peak_start_time,attr" validate:"required"`
	// RampDownCapacityThresholdPercent: number, required
	RampDownCapacityThresholdPercent terra.NumberValue `hcl:"ramp_down_capacity_threshold_percent,attr" validate:"required"`
	// RampDownForceLogoffUsers: bool, required
	RampDownForceLogoffUsers terra.BoolValue `hcl:"ramp_down_force_logoff_users,attr" validate:"required"`
	// RampDownLoadBalancingAlgorithm: string, required
	RampDownLoadBalancingAlgorithm terra.StringValue `hcl:"ramp_down_load_balancing_algorithm,attr" validate:"required"`
	// RampDownMinimumHostsPercent: number, required
	RampDownMinimumHostsPercent terra.NumberValue `hcl:"ramp_down_minimum_hosts_percent,attr" validate:"required"`
	// RampDownNotificationMessage: string, required
	RampDownNotificationMessage terra.StringValue `hcl:"ramp_down_notification_message,attr" validate:"required"`
	// RampDownStartTime: string, required
	RampDownStartTime terra.StringValue `hcl:"ramp_down_start_time,attr" validate:"required"`
	// RampDownStopHostsWhen: string, required
	RampDownStopHostsWhen terra.StringValue `hcl:"ramp_down_stop_hosts_when,attr" validate:"required"`
	// RampDownWaitTimeMinutes: number, required
	RampDownWaitTimeMinutes terra.NumberValue `hcl:"ramp_down_wait_time_minutes,attr" validate:"required"`
	// RampUpCapacityThresholdPercent: number, optional
	RampUpCapacityThresholdPercent terra.NumberValue `hcl:"ramp_up_capacity_threshold_percent,attr"`
	// RampUpLoadBalancingAlgorithm: string, required
	RampUpLoadBalancingAlgorithm terra.StringValue `hcl:"ramp_up_load_balancing_algorithm,attr" validate:"required"`
	// RampUpMinimumHostsPercent: number, optional
	RampUpMinimumHostsPercent terra.NumberValue `hcl:"ramp_up_minimum_hosts_percent,attr"`
	// RampUpStartTime: string, required
	RampUpStartTime terra.StringValue `hcl:"ramp_up_start_time,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type HostPoolAttributes struct {
	ref terra.Reference
}

func (hp HostPoolAttributes) InternalRef() (terra.Reference, error) {
	return hp.ref, nil
}

func (hp HostPoolAttributes) InternalWithRef(ref terra.Reference) HostPoolAttributes {
	return HostPoolAttributes{ref: ref}
}

func (hp HostPoolAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hp.ref.InternalTokens()
}

func (hp HostPoolAttributes) HostpoolId() terra.StringValue {
	return terra.ReferenceAsString(hp.ref.Append("hostpool_id"))
}

func (hp HostPoolAttributes) ScalingPlanEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(hp.ref.Append("scaling_plan_enabled"))
}

type ScheduleAttributes struct {
	ref terra.Reference
}

func (s ScheduleAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ScheduleAttributes) InternalWithRef(ref terra.Reference) ScheduleAttributes {
	return ScheduleAttributes{ref: ref}
}

func (s ScheduleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ScheduleAttributes) DaysOfWeek() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("days_of_week"))
}

func (s ScheduleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s ScheduleAttributes) OffPeakLoadBalancingAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("off_peak_load_balancing_algorithm"))
}

func (s ScheduleAttributes) OffPeakStartTime() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("off_peak_start_time"))
}

func (s ScheduleAttributes) PeakLoadBalancingAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("peak_load_balancing_algorithm"))
}

func (s ScheduleAttributes) PeakStartTime() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("peak_start_time"))
}

func (s ScheduleAttributes) RampDownCapacityThresholdPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("ramp_down_capacity_threshold_percent"))
}

func (s ScheduleAttributes) RampDownForceLogoffUsers() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("ramp_down_force_logoff_users"))
}

func (s ScheduleAttributes) RampDownLoadBalancingAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("ramp_down_load_balancing_algorithm"))
}

func (s ScheduleAttributes) RampDownMinimumHostsPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("ramp_down_minimum_hosts_percent"))
}

func (s ScheduleAttributes) RampDownNotificationMessage() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("ramp_down_notification_message"))
}

func (s ScheduleAttributes) RampDownStartTime() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("ramp_down_start_time"))
}

func (s ScheduleAttributes) RampDownStopHostsWhen() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("ramp_down_stop_hosts_when"))
}

func (s ScheduleAttributes) RampDownWaitTimeMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("ramp_down_wait_time_minutes"))
}

func (s ScheduleAttributes) RampUpCapacityThresholdPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("ramp_up_capacity_threshold_percent"))
}

func (s ScheduleAttributes) RampUpLoadBalancingAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("ramp_up_load_balancing_algorithm"))
}

func (s ScheduleAttributes) RampUpMinimumHostsPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("ramp_up_minimum_hosts_percent"))
}

func (s ScheduleAttributes) RampUpStartTime() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("ramp_up_start_time"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type HostPoolState struct {
	HostpoolId         string `json:"hostpool_id"`
	ScalingPlanEnabled bool   `json:"scaling_plan_enabled"`
}

type ScheduleState struct {
	DaysOfWeek                       []string `json:"days_of_week"`
	Name                             string   `json:"name"`
	OffPeakLoadBalancingAlgorithm    string   `json:"off_peak_load_balancing_algorithm"`
	OffPeakStartTime                 string   `json:"off_peak_start_time"`
	PeakLoadBalancingAlgorithm       string   `json:"peak_load_balancing_algorithm"`
	PeakStartTime                    string   `json:"peak_start_time"`
	RampDownCapacityThresholdPercent float64  `json:"ramp_down_capacity_threshold_percent"`
	RampDownForceLogoffUsers         bool     `json:"ramp_down_force_logoff_users"`
	RampDownLoadBalancingAlgorithm   string   `json:"ramp_down_load_balancing_algorithm"`
	RampDownMinimumHostsPercent      float64  `json:"ramp_down_minimum_hosts_percent"`
	RampDownNotificationMessage      string   `json:"ramp_down_notification_message"`
	RampDownStartTime                string   `json:"ramp_down_start_time"`
	RampDownStopHostsWhen            string   `json:"ramp_down_stop_hosts_when"`
	RampDownWaitTimeMinutes          float64  `json:"ramp_down_wait_time_minutes"`
	RampUpCapacityThresholdPercent   float64  `json:"ramp_up_capacity_threshold_percent"`
	RampUpLoadBalancingAlgorithm     string   `json:"ramp_up_load_balancing_algorithm"`
	RampUpMinimumHostsPercent        float64  `json:"ramp_up_minimum_hosts_percent"`
	RampUpStartTime                  string   `json:"ramp_up_start_time"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
