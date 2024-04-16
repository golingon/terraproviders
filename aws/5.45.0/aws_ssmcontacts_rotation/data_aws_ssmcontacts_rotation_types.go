// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ssmcontacts_rotation

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataRecurrenceAttributes struct {
	ref terra.Reference
}

func (r DataRecurrenceAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r DataRecurrenceAttributes) InternalWithRef(ref terra.Reference) DataRecurrenceAttributes {
	return DataRecurrenceAttributes{ref: ref}
}

func (r DataRecurrenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r DataRecurrenceAttributes) NumberOfOnCalls() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("number_of_on_calls"))
}

func (r DataRecurrenceAttributes) RecurrenceMultiplier() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("recurrence_multiplier"))
}

func (r DataRecurrenceAttributes) DailySettings() terra.ListValue[DataRecurrenceDailySettingsAttributes] {
	return terra.ReferenceAsList[DataRecurrenceDailySettingsAttributes](r.ref.Append("daily_settings"))
}

func (r DataRecurrenceAttributes) MonthlySettings() terra.ListValue[DataRecurrenceMonthlySettingsAttributes] {
	return terra.ReferenceAsList[DataRecurrenceMonthlySettingsAttributes](r.ref.Append("monthly_settings"))
}

func (r DataRecurrenceAttributes) ShiftCoverages() terra.ListValue[DataRecurrenceShiftCoveragesAttributes] {
	return terra.ReferenceAsList[DataRecurrenceShiftCoveragesAttributes](r.ref.Append("shift_coverages"))
}

func (r DataRecurrenceAttributes) WeeklySettings() terra.ListValue[DataRecurrenceWeeklySettingsAttributes] {
	return terra.ReferenceAsList[DataRecurrenceWeeklySettingsAttributes](r.ref.Append("weekly_settings"))
}

type DataRecurrenceDailySettingsAttributes struct {
	ref terra.Reference
}

func (ds DataRecurrenceDailySettingsAttributes) InternalRef() (terra.Reference, error) {
	return ds.ref, nil
}

func (ds DataRecurrenceDailySettingsAttributes) InternalWithRef(ref terra.Reference) DataRecurrenceDailySettingsAttributes {
	return DataRecurrenceDailySettingsAttributes{ref: ref}
}

func (ds DataRecurrenceDailySettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ds.ref.InternalTokens()
}

func (ds DataRecurrenceDailySettingsAttributes) HourOfDay() terra.NumberValue {
	return terra.ReferenceAsNumber(ds.ref.Append("hour_of_day"))
}

func (ds DataRecurrenceDailySettingsAttributes) MinuteOfHour() terra.NumberValue {
	return terra.ReferenceAsNumber(ds.ref.Append("minute_of_hour"))
}

type DataRecurrenceMonthlySettingsAttributes struct {
	ref terra.Reference
}

func (ms DataRecurrenceMonthlySettingsAttributes) InternalRef() (terra.Reference, error) {
	return ms.ref, nil
}

func (ms DataRecurrenceMonthlySettingsAttributes) InternalWithRef(ref terra.Reference) DataRecurrenceMonthlySettingsAttributes {
	return DataRecurrenceMonthlySettingsAttributes{ref: ref}
}

func (ms DataRecurrenceMonthlySettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ms.ref.InternalTokens()
}

func (ms DataRecurrenceMonthlySettingsAttributes) DayOfMonth() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("day_of_month"))
}

func (ms DataRecurrenceMonthlySettingsAttributes) HandOffTime() terra.ListValue[DataRecurrenceMonthlySettingsHandOffTimeAttributes] {
	return terra.ReferenceAsList[DataRecurrenceMonthlySettingsHandOffTimeAttributes](ms.ref.Append("hand_off_time"))
}

type DataRecurrenceMonthlySettingsHandOffTimeAttributes struct {
	ref terra.Reference
}

func (hot DataRecurrenceMonthlySettingsHandOffTimeAttributes) InternalRef() (terra.Reference, error) {
	return hot.ref, nil
}

func (hot DataRecurrenceMonthlySettingsHandOffTimeAttributes) InternalWithRef(ref terra.Reference) DataRecurrenceMonthlySettingsHandOffTimeAttributes {
	return DataRecurrenceMonthlySettingsHandOffTimeAttributes{ref: ref}
}

func (hot DataRecurrenceMonthlySettingsHandOffTimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hot.ref.InternalTokens()
}

func (hot DataRecurrenceMonthlySettingsHandOffTimeAttributes) HourOfDay() terra.NumberValue {
	return terra.ReferenceAsNumber(hot.ref.Append("hour_of_day"))
}

func (hot DataRecurrenceMonthlySettingsHandOffTimeAttributes) MinuteOfHour() terra.NumberValue {
	return terra.ReferenceAsNumber(hot.ref.Append("minute_of_hour"))
}

type DataRecurrenceShiftCoveragesAttributes struct {
	ref terra.Reference
}

func (sc DataRecurrenceShiftCoveragesAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc DataRecurrenceShiftCoveragesAttributes) InternalWithRef(ref terra.Reference) DataRecurrenceShiftCoveragesAttributes {
	return DataRecurrenceShiftCoveragesAttributes{ref: ref}
}

func (sc DataRecurrenceShiftCoveragesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc DataRecurrenceShiftCoveragesAttributes) MapBlockKey() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("map_block_key"))
}

func (sc DataRecurrenceShiftCoveragesAttributes) CoverageTimes() terra.ListValue[DataRecurrenceShiftCoveragesCoverageTimesAttributes] {
	return terra.ReferenceAsList[DataRecurrenceShiftCoveragesCoverageTimesAttributes](sc.ref.Append("coverage_times"))
}

type DataRecurrenceShiftCoveragesCoverageTimesAttributes struct {
	ref terra.Reference
}

func (ct DataRecurrenceShiftCoveragesCoverageTimesAttributes) InternalRef() (terra.Reference, error) {
	return ct.ref, nil
}

func (ct DataRecurrenceShiftCoveragesCoverageTimesAttributes) InternalWithRef(ref terra.Reference) DataRecurrenceShiftCoveragesCoverageTimesAttributes {
	return DataRecurrenceShiftCoveragesCoverageTimesAttributes{ref: ref}
}

func (ct DataRecurrenceShiftCoveragesCoverageTimesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ct.ref.InternalTokens()
}

func (ct DataRecurrenceShiftCoveragesCoverageTimesAttributes) End() terra.ListValue[DataRecurrenceShiftCoveragesCoverageTimesEndAttributes] {
	return terra.ReferenceAsList[DataRecurrenceShiftCoveragesCoverageTimesEndAttributes](ct.ref.Append("end"))
}

func (ct DataRecurrenceShiftCoveragesCoverageTimesAttributes) Start() terra.ListValue[DataRecurrenceShiftCoveragesCoverageTimesStartAttributes] {
	return terra.ReferenceAsList[DataRecurrenceShiftCoveragesCoverageTimesStartAttributes](ct.ref.Append("start"))
}

type DataRecurrenceShiftCoveragesCoverageTimesEndAttributes struct {
	ref terra.Reference
}

func (e DataRecurrenceShiftCoveragesCoverageTimesEndAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e DataRecurrenceShiftCoveragesCoverageTimesEndAttributes) InternalWithRef(ref terra.Reference) DataRecurrenceShiftCoveragesCoverageTimesEndAttributes {
	return DataRecurrenceShiftCoveragesCoverageTimesEndAttributes{ref: ref}
}

func (e DataRecurrenceShiftCoveragesCoverageTimesEndAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e DataRecurrenceShiftCoveragesCoverageTimesEndAttributes) HourOfDay() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("hour_of_day"))
}

func (e DataRecurrenceShiftCoveragesCoverageTimesEndAttributes) MinuteOfHour() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("minute_of_hour"))
}

type DataRecurrenceShiftCoveragesCoverageTimesStartAttributes struct {
	ref terra.Reference
}

func (s DataRecurrenceShiftCoveragesCoverageTimesStartAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s DataRecurrenceShiftCoveragesCoverageTimesStartAttributes) InternalWithRef(ref terra.Reference) DataRecurrenceShiftCoveragesCoverageTimesStartAttributes {
	return DataRecurrenceShiftCoveragesCoverageTimesStartAttributes{ref: ref}
}

func (s DataRecurrenceShiftCoveragesCoverageTimesStartAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s DataRecurrenceShiftCoveragesCoverageTimesStartAttributes) HourOfDay() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("hour_of_day"))
}

func (s DataRecurrenceShiftCoveragesCoverageTimesStartAttributes) MinuteOfHour() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("minute_of_hour"))
}

type DataRecurrenceWeeklySettingsAttributes struct {
	ref terra.Reference
}

func (ws DataRecurrenceWeeklySettingsAttributes) InternalRef() (terra.Reference, error) {
	return ws.ref, nil
}

func (ws DataRecurrenceWeeklySettingsAttributes) InternalWithRef(ref terra.Reference) DataRecurrenceWeeklySettingsAttributes {
	return DataRecurrenceWeeklySettingsAttributes{ref: ref}
}

func (ws DataRecurrenceWeeklySettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ws.ref.InternalTokens()
}

func (ws DataRecurrenceWeeklySettingsAttributes) DayOfWeek() terra.StringValue {
	return terra.ReferenceAsString(ws.ref.Append("day_of_week"))
}

func (ws DataRecurrenceWeeklySettingsAttributes) HandOffTime() terra.ListValue[DataRecurrenceWeeklySettingsHandOffTimeAttributes] {
	return terra.ReferenceAsList[DataRecurrenceWeeklySettingsHandOffTimeAttributes](ws.ref.Append("hand_off_time"))
}

type DataRecurrenceWeeklySettingsHandOffTimeAttributes struct {
	ref terra.Reference
}

func (hot DataRecurrenceWeeklySettingsHandOffTimeAttributes) InternalRef() (terra.Reference, error) {
	return hot.ref, nil
}

func (hot DataRecurrenceWeeklySettingsHandOffTimeAttributes) InternalWithRef(ref terra.Reference) DataRecurrenceWeeklySettingsHandOffTimeAttributes {
	return DataRecurrenceWeeklySettingsHandOffTimeAttributes{ref: ref}
}

func (hot DataRecurrenceWeeklySettingsHandOffTimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hot.ref.InternalTokens()
}

func (hot DataRecurrenceWeeklySettingsHandOffTimeAttributes) HourOfDay() terra.NumberValue {
	return terra.ReferenceAsNumber(hot.ref.Append("hour_of_day"))
}

func (hot DataRecurrenceWeeklySettingsHandOffTimeAttributes) MinuteOfHour() terra.NumberValue {
	return terra.ReferenceAsNumber(hot.ref.Append("minute_of_hour"))
}

type DataRecurrenceState struct {
	NumberOfOnCalls      float64                              `json:"number_of_on_calls"`
	RecurrenceMultiplier float64                              `json:"recurrence_multiplier"`
	DailySettings        []DataRecurrenceDailySettingsState   `json:"daily_settings"`
	MonthlySettings      []DataRecurrenceMonthlySettingsState `json:"monthly_settings"`
	ShiftCoverages       []DataRecurrenceShiftCoveragesState  `json:"shift_coverages"`
	WeeklySettings       []DataRecurrenceWeeklySettingsState  `json:"weekly_settings"`
}

type DataRecurrenceDailySettingsState struct {
	HourOfDay    float64 `json:"hour_of_day"`
	MinuteOfHour float64 `json:"minute_of_hour"`
}

type DataRecurrenceMonthlySettingsState struct {
	DayOfMonth  float64                                         `json:"day_of_month"`
	HandOffTime []DataRecurrenceMonthlySettingsHandOffTimeState `json:"hand_off_time"`
}

type DataRecurrenceMonthlySettingsHandOffTimeState struct {
	HourOfDay    float64 `json:"hour_of_day"`
	MinuteOfHour float64 `json:"minute_of_hour"`
}

type DataRecurrenceShiftCoveragesState struct {
	MapBlockKey   string                                           `json:"map_block_key"`
	CoverageTimes []DataRecurrenceShiftCoveragesCoverageTimesState `json:"coverage_times"`
}

type DataRecurrenceShiftCoveragesCoverageTimesState struct {
	End   []DataRecurrenceShiftCoveragesCoverageTimesEndState   `json:"end"`
	Start []DataRecurrenceShiftCoveragesCoverageTimesStartState `json:"start"`
}

type DataRecurrenceShiftCoveragesCoverageTimesEndState struct {
	HourOfDay    float64 `json:"hour_of_day"`
	MinuteOfHour float64 `json:"minute_of_hour"`
}

type DataRecurrenceShiftCoveragesCoverageTimesStartState struct {
	HourOfDay    float64 `json:"hour_of_day"`
	MinuteOfHour float64 `json:"minute_of_hour"`
}

type DataRecurrenceWeeklySettingsState struct {
	DayOfWeek   string                                         `json:"day_of_week"`
	HandOffTime []DataRecurrenceWeeklySettingsHandOffTimeState `json:"hand_off_time"`
}

type DataRecurrenceWeeklySettingsHandOffTimeState struct {
	HourOfDay    float64 `json:"hour_of_day"`
	MinuteOfHour float64 `json:"minute_of_hour"`
}
