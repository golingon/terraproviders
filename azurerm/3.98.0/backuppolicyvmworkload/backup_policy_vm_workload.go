// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package backuppolicyvmworkload

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ProtectionPolicy struct {
	// PolicyType: string, required
	PolicyType terra.StringValue `hcl:"policy_type,attr" validate:"required"`
	// Backup: required
	Backup *Backup `hcl:"backup,block" validate:"required"`
	// RetentionDaily: optional
	RetentionDaily *RetentionDaily `hcl:"retention_daily,block"`
	// RetentionMonthly: optional
	RetentionMonthly *RetentionMonthly `hcl:"retention_monthly,block"`
	// RetentionWeekly: optional
	RetentionWeekly *RetentionWeekly `hcl:"retention_weekly,block"`
	// RetentionYearly: optional
	RetentionYearly *RetentionYearly `hcl:"retention_yearly,block"`
	// SimpleRetention: optional
	SimpleRetention *SimpleRetention `hcl:"simple_retention,block"`
}

type Backup struct {
	// Frequency: string, optional
	Frequency terra.StringValue `hcl:"frequency,attr"`
	// FrequencyInMinutes: number, optional
	FrequencyInMinutes terra.NumberValue `hcl:"frequency_in_minutes,attr"`
	// Time: string, optional
	Time terra.StringValue `hcl:"time,attr"`
	// Weekdays: set of string, optional
	Weekdays terra.SetValue[terra.StringValue] `hcl:"weekdays,attr"`
}

type RetentionDaily struct {
	// Count: number, required
	Count terra.NumberValue `hcl:"count,attr" validate:"required"`
}

type RetentionMonthly struct {
	// Count: number, required
	Count terra.NumberValue `hcl:"count,attr" validate:"required"`
	// FormatType: string, required
	FormatType terra.StringValue `hcl:"format_type,attr" validate:"required"`
	// Monthdays: set of number, optional
	Monthdays terra.SetValue[terra.NumberValue] `hcl:"monthdays,attr"`
	// Weekdays: set of string, optional
	Weekdays terra.SetValue[terra.StringValue] `hcl:"weekdays,attr"`
	// Weeks: set of string, optional
	Weeks terra.SetValue[terra.StringValue] `hcl:"weeks,attr"`
}

type RetentionWeekly struct {
	// Count: number, required
	Count terra.NumberValue `hcl:"count,attr" validate:"required"`
	// Weekdays: set of string, required
	Weekdays terra.SetValue[terra.StringValue] `hcl:"weekdays,attr" validate:"required"`
}

type RetentionYearly struct {
	// Count: number, required
	Count terra.NumberValue `hcl:"count,attr" validate:"required"`
	// FormatType: string, required
	FormatType terra.StringValue `hcl:"format_type,attr" validate:"required"`
	// Monthdays: set of number, optional
	Monthdays terra.SetValue[terra.NumberValue] `hcl:"monthdays,attr"`
	// Months: set of string, required
	Months terra.SetValue[terra.StringValue] `hcl:"months,attr" validate:"required"`
	// Weekdays: set of string, optional
	Weekdays terra.SetValue[terra.StringValue] `hcl:"weekdays,attr"`
	// Weeks: set of string, optional
	Weeks terra.SetValue[terra.StringValue] `hcl:"weeks,attr"`
}

type SimpleRetention struct {
	// Count: number, required
	Count terra.NumberValue `hcl:"count,attr" validate:"required"`
}

type Settings struct {
	// CompressionEnabled: bool, optional
	CompressionEnabled terra.BoolValue `hcl:"compression_enabled,attr"`
	// TimeZone: string, required
	TimeZone terra.StringValue `hcl:"time_zone,attr" validate:"required"`
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

type ProtectionPolicyAttributes struct {
	ref terra.Reference
}

func (pp ProtectionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return pp.ref, nil
}

func (pp ProtectionPolicyAttributes) InternalWithRef(ref terra.Reference) ProtectionPolicyAttributes {
	return ProtectionPolicyAttributes{ref: ref}
}

func (pp ProtectionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pp.ref.InternalTokens()
}

func (pp ProtectionPolicyAttributes) PolicyType() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("policy_type"))
}

func (pp ProtectionPolicyAttributes) Backup() terra.ListValue[BackupAttributes] {
	return terra.ReferenceAsList[BackupAttributes](pp.ref.Append("backup"))
}

func (pp ProtectionPolicyAttributes) RetentionDaily() terra.ListValue[RetentionDailyAttributes] {
	return terra.ReferenceAsList[RetentionDailyAttributes](pp.ref.Append("retention_daily"))
}

func (pp ProtectionPolicyAttributes) RetentionMonthly() terra.ListValue[RetentionMonthlyAttributes] {
	return terra.ReferenceAsList[RetentionMonthlyAttributes](pp.ref.Append("retention_monthly"))
}

func (pp ProtectionPolicyAttributes) RetentionWeekly() terra.ListValue[RetentionWeeklyAttributes] {
	return terra.ReferenceAsList[RetentionWeeklyAttributes](pp.ref.Append("retention_weekly"))
}

func (pp ProtectionPolicyAttributes) RetentionYearly() terra.ListValue[RetentionYearlyAttributes] {
	return terra.ReferenceAsList[RetentionYearlyAttributes](pp.ref.Append("retention_yearly"))
}

func (pp ProtectionPolicyAttributes) SimpleRetention() terra.ListValue[SimpleRetentionAttributes] {
	return terra.ReferenceAsList[SimpleRetentionAttributes](pp.ref.Append("simple_retention"))
}

type BackupAttributes struct {
	ref terra.Reference
}

func (b BackupAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b BackupAttributes) InternalWithRef(ref terra.Reference) BackupAttributes {
	return BackupAttributes{ref: ref}
}

func (b BackupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b BackupAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("frequency"))
}

func (b BackupAttributes) FrequencyInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("frequency_in_minutes"))
}

func (b BackupAttributes) Time() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("time"))
}

func (b BackupAttributes) Weekdays() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](b.ref.Append("weekdays"))
}

type RetentionDailyAttributes struct {
	ref terra.Reference
}

func (rd RetentionDailyAttributes) InternalRef() (terra.Reference, error) {
	return rd.ref, nil
}

func (rd RetentionDailyAttributes) InternalWithRef(ref terra.Reference) RetentionDailyAttributes {
	return RetentionDailyAttributes{ref: ref}
}

func (rd RetentionDailyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rd.ref.InternalTokens()
}

func (rd RetentionDailyAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(rd.ref.Append("count"))
}

type RetentionMonthlyAttributes struct {
	ref terra.Reference
}

func (rm RetentionMonthlyAttributes) InternalRef() (terra.Reference, error) {
	return rm.ref, nil
}

func (rm RetentionMonthlyAttributes) InternalWithRef(ref terra.Reference) RetentionMonthlyAttributes {
	return RetentionMonthlyAttributes{ref: ref}
}

func (rm RetentionMonthlyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rm.ref.InternalTokens()
}

func (rm RetentionMonthlyAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(rm.ref.Append("count"))
}

func (rm RetentionMonthlyAttributes) FormatType() terra.StringValue {
	return terra.ReferenceAsString(rm.ref.Append("format_type"))
}

func (rm RetentionMonthlyAttributes) Monthdays() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](rm.ref.Append("monthdays"))
}

func (rm RetentionMonthlyAttributes) Weekdays() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rm.ref.Append("weekdays"))
}

func (rm RetentionMonthlyAttributes) Weeks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rm.ref.Append("weeks"))
}

type RetentionWeeklyAttributes struct {
	ref terra.Reference
}

func (rw RetentionWeeklyAttributes) InternalRef() (terra.Reference, error) {
	return rw.ref, nil
}

func (rw RetentionWeeklyAttributes) InternalWithRef(ref terra.Reference) RetentionWeeklyAttributes {
	return RetentionWeeklyAttributes{ref: ref}
}

func (rw RetentionWeeklyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rw.ref.InternalTokens()
}

func (rw RetentionWeeklyAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(rw.ref.Append("count"))
}

func (rw RetentionWeeklyAttributes) Weekdays() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rw.ref.Append("weekdays"))
}

type RetentionYearlyAttributes struct {
	ref terra.Reference
}

func (ry RetentionYearlyAttributes) InternalRef() (terra.Reference, error) {
	return ry.ref, nil
}

func (ry RetentionYearlyAttributes) InternalWithRef(ref terra.Reference) RetentionYearlyAttributes {
	return RetentionYearlyAttributes{ref: ref}
}

func (ry RetentionYearlyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ry.ref.InternalTokens()
}

func (ry RetentionYearlyAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(ry.ref.Append("count"))
}

func (ry RetentionYearlyAttributes) FormatType() terra.StringValue {
	return terra.ReferenceAsString(ry.ref.Append("format_type"))
}

func (ry RetentionYearlyAttributes) Monthdays() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](ry.ref.Append("monthdays"))
}

func (ry RetentionYearlyAttributes) Months() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ry.ref.Append("months"))
}

func (ry RetentionYearlyAttributes) Weekdays() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ry.ref.Append("weekdays"))
}

func (ry RetentionYearlyAttributes) Weeks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ry.ref.Append("weeks"))
}

type SimpleRetentionAttributes struct {
	ref terra.Reference
}

func (sr SimpleRetentionAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr SimpleRetentionAttributes) InternalWithRef(ref terra.Reference) SimpleRetentionAttributes {
	return SimpleRetentionAttributes{ref: ref}
}

func (sr SimpleRetentionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr SimpleRetentionAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(sr.ref.Append("count"))
}

type SettingsAttributes struct {
	ref terra.Reference
}

func (s SettingsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SettingsAttributes) InternalWithRef(ref terra.Reference) SettingsAttributes {
	return SettingsAttributes{ref: ref}
}

func (s SettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SettingsAttributes) CompressionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("compression_enabled"))
}

func (s SettingsAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("time_zone"))
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

type ProtectionPolicyState struct {
	PolicyType       string                  `json:"policy_type"`
	Backup           []BackupState           `json:"backup"`
	RetentionDaily   []RetentionDailyState   `json:"retention_daily"`
	RetentionMonthly []RetentionMonthlyState `json:"retention_monthly"`
	RetentionWeekly  []RetentionWeeklyState  `json:"retention_weekly"`
	RetentionYearly  []RetentionYearlyState  `json:"retention_yearly"`
	SimpleRetention  []SimpleRetentionState  `json:"simple_retention"`
}

type BackupState struct {
	Frequency          string   `json:"frequency"`
	FrequencyInMinutes float64  `json:"frequency_in_minutes"`
	Time               string   `json:"time"`
	Weekdays           []string `json:"weekdays"`
}

type RetentionDailyState struct {
	Count float64 `json:"count"`
}

type RetentionMonthlyState struct {
	Count      float64   `json:"count"`
	FormatType string    `json:"format_type"`
	Monthdays  []float64 `json:"monthdays"`
	Weekdays   []string  `json:"weekdays"`
	Weeks      []string  `json:"weeks"`
}

type RetentionWeeklyState struct {
	Count    float64  `json:"count"`
	Weekdays []string `json:"weekdays"`
}

type RetentionYearlyState struct {
	Count      float64   `json:"count"`
	FormatType string    `json:"format_type"`
	Monthdays  []float64 `json:"monthdays"`
	Months     []string  `json:"months"`
	Weekdays   []string  `json:"weekdays"`
	Weeks      []string  `json:"weeks"`
}

type SimpleRetentionState struct {
	Count float64 `json:"count"`
}

type SettingsState struct {
	CompressionEnabled bool   `json:"compression_enabled"`
	TimeZone           string `json:"time_zone"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
