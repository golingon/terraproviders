// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_logic_app_integration_account_batch_configuration

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ReleaseCriteria struct {
	// BatchSize: number, optional
	BatchSize terra.NumberValue `hcl:"batch_size,attr"`
	// MessageCount: number, optional
	MessageCount terra.NumberValue `hcl:"message_count,attr"`
	// ReleaseCriteriaRecurrence: optional
	Recurrence *ReleaseCriteriaRecurrence `hcl:"recurrence,block"`
}

type ReleaseCriteriaRecurrence struct {
	// EndTime: string, optional
	EndTime terra.StringValue `hcl:"end_time,attr"`
	// Frequency: string, required
	Frequency terra.StringValue `hcl:"frequency,attr" validate:"required"`
	// Interval: number, required
	Interval terra.NumberValue `hcl:"interval,attr" validate:"required"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
	// TimeZone: string, optional
	TimeZone terra.StringValue `hcl:"time_zone,attr"`
	// ReleaseCriteriaRecurrenceSchedule: optional
	Schedule *ReleaseCriteriaRecurrenceSchedule `hcl:"schedule,block"`
}

type ReleaseCriteriaRecurrenceSchedule struct {
	// Hours: set of number, optional
	Hours terra.SetValue[terra.NumberValue] `hcl:"hours,attr"`
	// Minutes: set of number, optional
	Minutes terra.SetValue[terra.NumberValue] `hcl:"minutes,attr"`
	// MonthDays: set of number, optional
	MonthDays terra.SetValue[terra.NumberValue] `hcl:"month_days,attr"`
	// WeekDays: set of string, optional
	WeekDays terra.SetValue[terra.StringValue] `hcl:"week_days,attr"`
	// ReleaseCriteriaRecurrenceScheduleMonthly: min=0
	Monthly []ReleaseCriteriaRecurrenceScheduleMonthly `hcl:"monthly,block" validate:"min=0"`
}

type ReleaseCriteriaRecurrenceScheduleMonthly struct {
	// Week: number, required
	Week terra.NumberValue `hcl:"week,attr" validate:"required"`
	// Weekday: string, required
	Weekday terra.StringValue `hcl:"weekday,attr" validate:"required"`
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

type ReleaseCriteriaAttributes struct {
	ref terra.Reference
}

func (rc ReleaseCriteriaAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc ReleaseCriteriaAttributes) InternalWithRef(ref terra.Reference) ReleaseCriteriaAttributes {
	return ReleaseCriteriaAttributes{ref: ref}
}

func (rc ReleaseCriteriaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc ReleaseCriteriaAttributes) BatchSize() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("batch_size"))
}

func (rc ReleaseCriteriaAttributes) MessageCount() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("message_count"))
}

func (rc ReleaseCriteriaAttributes) Recurrence() terra.ListValue[ReleaseCriteriaRecurrenceAttributes] {
	return terra.ReferenceAsList[ReleaseCriteriaRecurrenceAttributes](rc.ref.Append("recurrence"))
}

type ReleaseCriteriaRecurrenceAttributes struct {
	ref terra.Reference
}

func (r ReleaseCriteriaRecurrenceAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ReleaseCriteriaRecurrenceAttributes) InternalWithRef(ref terra.Reference) ReleaseCriteriaRecurrenceAttributes {
	return ReleaseCriteriaRecurrenceAttributes{ref: ref}
}

func (r ReleaseCriteriaRecurrenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ReleaseCriteriaRecurrenceAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("end_time"))
}

func (r ReleaseCriteriaRecurrenceAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("frequency"))
}

func (r ReleaseCriteriaRecurrenceAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("interval"))
}

func (r ReleaseCriteriaRecurrenceAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("start_time"))
}

func (r ReleaseCriteriaRecurrenceAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("time_zone"))
}

func (r ReleaseCriteriaRecurrenceAttributes) Schedule() terra.ListValue[ReleaseCriteriaRecurrenceScheduleAttributes] {
	return terra.ReferenceAsList[ReleaseCriteriaRecurrenceScheduleAttributes](r.ref.Append("schedule"))
}

type ReleaseCriteriaRecurrenceScheduleAttributes struct {
	ref terra.Reference
}

func (s ReleaseCriteriaRecurrenceScheduleAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ReleaseCriteriaRecurrenceScheduleAttributes) InternalWithRef(ref terra.Reference) ReleaseCriteriaRecurrenceScheduleAttributes {
	return ReleaseCriteriaRecurrenceScheduleAttributes{ref: ref}
}

func (s ReleaseCriteriaRecurrenceScheduleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ReleaseCriteriaRecurrenceScheduleAttributes) Hours() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](s.ref.Append("hours"))
}

func (s ReleaseCriteriaRecurrenceScheduleAttributes) Minutes() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](s.ref.Append("minutes"))
}

func (s ReleaseCriteriaRecurrenceScheduleAttributes) MonthDays() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](s.ref.Append("month_days"))
}

func (s ReleaseCriteriaRecurrenceScheduleAttributes) WeekDays() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("week_days"))
}

func (s ReleaseCriteriaRecurrenceScheduleAttributes) Monthly() terra.SetValue[ReleaseCriteriaRecurrenceScheduleMonthlyAttributes] {
	return terra.ReferenceAsSet[ReleaseCriteriaRecurrenceScheduleMonthlyAttributes](s.ref.Append("monthly"))
}

type ReleaseCriteriaRecurrenceScheduleMonthlyAttributes struct {
	ref terra.Reference
}

func (m ReleaseCriteriaRecurrenceScheduleMonthlyAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m ReleaseCriteriaRecurrenceScheduleMonthlyAttributes) InternalWithRef(ref terra.Reference) ReleaseCriteriaRecurrenceScheduleMonthlyAttributes {
	return ReleaseCriteriaRecurrenceScheduleMonthlyAttributes{ref: ref}
}

func (m ReleaseCriteriaRecurrenceScheduleMonthlyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m ReleaseCriteriaRecurrenceScheduleMonthlyAttributes) Week() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("week"))
}

func (m ReleaseCriteriaRecurrenceScheduleMonthlyAttributes) Weekday() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("weekday"))
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

type ReleaseCriteriaState struct {
	BatchSize    float64                          `json:"batch_size"`
	MessageCount float64                          `json:"message_count"`
	Recurrence   []ReleaseCriteriaRecurrenceState `json:"recurrence"`
}

type ReleaseCriteriaRecurrenceState struct {
	EndTime   string                                   `json:"end_time"`
	Frequency string                                   `json:"frequency"`
	Interval  float64                                  `json:"interval"`
	StartTime string                                   `json:"start_time"`
	TimeZone  string                                   `json:"time_zone"`
	Schedule  []ReleaseCriteriaRecurrenceScheduleState `json:"schedule"`
}

type ReleaseCriteriaRecurrenceScheduleState struct {
	Hours     []float64                                       `json:"hours"`
	Minutes   []float64                                       `json:"minutes"`
	MonthDays []float64                                       `json:"month_days"`
	WeekDays  []string                                        `json:"week_days"`
	Monthly   []ReleaseCriteriaRecurrenceScheduleMonthlyState `json:"monthly"`
}

type ReleaseCriteriaRecurrenceScheduleMonthlyState struct {
	Week    float64 `json:"week"`
	Weekday string  `json:"weekday"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
