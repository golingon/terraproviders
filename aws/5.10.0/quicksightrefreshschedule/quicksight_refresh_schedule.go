// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package quicksightrefreshschedule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Schedule struct {
	// RefreshType: string, required
	RefreshType terra.StringValue `hcl:"refresh_type,attr" validate:"required"`
	// StartAfterDateTime: string, optional
	StartAfterDateTime terra.StringValue `hcl:"start_after_date_time,attr"`
	// ScheduleFrequency: min=0
	ScheduleFrequency []ScheduleFrequency `hcl:"schedule_frequency,block" validate:"min=0"`
}

type ScheduleFrequency struct {
	// Interval: string, required
	Interval terra.StringValue `hcl:"interval,attr" validate:"required"`
	// TimeOfTheDay: string, optional
	TimeOfTheDay terra.StringValue `hcl:"time_of_the_day,attr"`
	// Timezone: string, optional
	Timezone terra.StringValue `hcl:"timezone,attr"`
	// RefreshOnDay: min=0
	RefreshOnDay []RefreshOnDay `hcl:"refresh_on_day,block" validate:"min=0"`
}

type RefreshOnDay struct {
	// DayOfMonth: string, optional
	DayOfMonth terra.StringValue `hcl:"day_of_month,attr"`
	// DayOfWeek: string, optional
	DayOfWeek terra.StringValue `hcl:"day_of_week,attr"`
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

func (s ScheduleAttributes) RefreshType() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("refresh_type"))
}

func (s ScheduleAttributes) StartAfterDateTime() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("start_after_date_time"))
}

func (s ScheduleAttributes) ScheduleFrequency() terra.ListValue[ScheduleFrequencyAttributes] {
	return terra.ReferenceAsList[ScheduleFrequencyAttributes](s.ref.Append("schedule_frequency"))
}

type ScheduleFrequencyAttributes struct {
	ref terra.Reference
}

func (sf ScheduleFrequencyAttributes) InternalRef() (terra.Reference, error) {
	return sf.ref, nil
}

func (sf ScheduleFrequencyAttributes) InternalWithRef(ref terra.Reference) ScheduleFrequencyAttributes {
	return ScheduleFrequencyAttributes{ref: ref}
}

func (sf ScheduleFrequencyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sf.ref.InternalTokens()
}

func (sf ScheduleFrequencyAttributes) Interval() terra.StringValue {
	return terra.ReferenceAsString(sf.ref.Append("interval"))
}

func (sf ScheduleFrequencyAttributes) TimeOfTheDay() terra.StringValue {
	return terra.ReferenceAsString(sf.ref.Append("time_of_the_day"))
}

func (sf ScheduleFrequencyAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(sf.ref.Append("timezone"))
}

func (sf ScheduleFrequencyAttributes) RefreshOnDay() terra.ListValue[RefreshOnDayAttributes] {
	return terra.ReferenceAsList[RefreshOnDayAttributes](sf.ref.Append("refresh_on_day"))
}

type RefreshOnDayAttributes struct {
	ref terra.Reference
}

func (rod RefreshOnDayAttributes) InternalRef() (terra.Reference, error) {
	return rod.ref, nil
}

func (rod RefreshOnDayAttributes) InternalWithRef(ref terra.Reference) RefreshOnDayAttributes {
	return RefreshOnDayAttributes{ref: ref}
}

func (rod RefreshOnDayAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rod.ref.InternalTokens()
}

func (rod RefreshOnDayAttributes) DayOfMonth() terra.StringValue {
	return terra.ReferenceAsString(rod.ref.Append("day_of_month"))
}

func (rod RefreshOnDayAttributes) DayOfWeek() terra.StringValue {
	return terra.ReferenceAsString(rod.ref.Append("day_of_week"))
}

type ScheduleState struct {
	RefreshType        string                   `json:"refresh_type"`
	StartAfterDateTime string                   `json:"start_after_date_time"`
	ScheduleFrequency  []ScheduleFrequencyState `json:"schedule_frequency"`
}

type ScheduleFrequencyState struct {
	Interval     string              `json:"interval"`
	TimeOfTheDay string              `json:"time_of_the_day"`
	Timezone     string              `json:"timezone"`
	RefreshOnDay []RefreshOnDayState `json:"refresh_on_day"`
}

type RefreshOnDayState struct {
	DayOfMonth string `json:"day_of_month"`
	DayOfWeek  string `json:"day_of_week"`
}
