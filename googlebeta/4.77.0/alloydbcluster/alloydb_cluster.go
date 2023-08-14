// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package alloydbcluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BackupSource struct{}

type ContinuousBackupInfo struct {
	// ContinuousBackupInfoEncryptionInfo: min=0
	EncryptionInfo []ContinuousBackupInfoEncryptionInfo `hcl:"encryption_info,block" validate:"min=0"`
}

type ContinuousBackupInfoEncryptionInfo struct{}

type EncryptionInfo struct{}

type MigrationSource struct{}

type AutomatedBackupPolicy struct {
	// BackupWindow: string, optional
	BackupWindow terra.StringValue `hcl:"backup_window,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// AutomatedBackupPolicyEncryptionConfig: optional
	EncryptionConfig *AutomatedBackupPolicyEncryptionConfig `hcl:"encryption_config,block"`
	// QuantityBasedRetention: optional
	QuantityBasedRetention *QuantityBasedRetention `hcl:"quantity_based_retention,block"`
	// TimeBasedRetention: optional
	TimeBasedRetention *TimeBasedRetention `hcl:"time_based_retention,block"`
	// WeeklySchedule: optional
	WeeklySchedule *WeeklySchedule `hcl:"weekly_schedule,block"`
}

type AutomatedBackupPolicyEncryptionConfig struct {
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
}

type QuantityBasedRetention struct {
	// Count: number, optional
	Count terra.NumberValue `hcl:"count,attr"`
}

type TimeBasedRetention struct {
	// RetentionPeriod: string, optional
	RetentionPeriod terra.StringValue `hcl:"retention_period,attr"`
}

type WeeklySchedule struct {
	// DaysOfWeek: list of string, optional
	DaysOfWeek terra.ListValue[terra.StringValue] `hcl:"days_of_week,attr"`
	// StartTimes: min=1
	StartTimes []StartTimes `hcl:"start_times,block" validate:"min=1"`
}

type StartTimes struct {
	// Hours: number, optional
	Hours terra.NumberValue `hcl:"hours,attr"`
	// Minutes: number, optional
	Minutes terra.NumberValue `hcl:"minutes,attr"`
	// Nanos: number, optional
	Nanos terra.NumberValue `hcl:"nanos,attr"`
	// Seconds: number, optional
	Seconds terra.NumberValue `hcl:"seconds,attr"`
}

type ContinuousBackupConfig struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// RecoveryWindowDays: number, optional
	RecoveryWindowDays terra.NumberValue `hcl:"recovery_window_days,attr"`
	// ContinuousBackupConfigEncryptionConfig: optional
	EncryptionConfig *ContinuousBackupConfigEncryptionConfig `hcl:"encryption_config,block"`
}

type ContinuousBackupConfigEncryptionConfig struct {
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
}

type EncryptionConfig struct {
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
}

type InitialUser struct {
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// User: string, optional
	User terra.StringValue `hcl:"user,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type BackupSourceAttributes struct {
	ref terra.Reference
}

func (bs BackupSourceAttributes) InternalRef() (terra.Reference, error) {
	return bs.ref, nil
}

func (bs BackupSourceAttributes) InternalWithRef(ref terra.Reference) BackupSourceAttributes {
	return BackupSourceAttributes{ref: ref}
}

func (bs BackupSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bs.ref.InternalTokens()
}

func (bs BackupSourceAttributes) BackupName() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("backup_name"))
}

type ContinuousBackupInfoAttributes struct {
	ref terra.Reference
}

func (cbi ContinuousBackupInfoAttributes) InternalRef() (terra.Reference, error) {
	return cbi.ref, nil
}

func (cbi ContinuousBackupInfoAttributes) InternalWithRef(ref terra.Reference) ContinuousBackupInfoAttributes {
	return ContinuousBackupInfoAttributes{ref: ref}
}

func (cbi ContinuousBackupInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cbi.ref.InternalTokens()
}

func (cbi ContinuousBackupInfoAttributes) EarliestRestorableTime() terra.StringValue {
	return terra.ReferenceAsString(cbi.ref.Append("earliest_restorable_time"))
}

func (cbi ContinuousBackupInfoAttributes) EnabledTime() terra.StringValue {
	return terra.ReferenceAsString(cbi.ref.Append("enabled_time"))
}

func (cbi ContinuousBackupInfoAttributes) Schedule() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cbi.ref.Append("schedule"))
}

func (cbi ContinuousBackupInfoAttributes) EncryptionInfo() terra.ListValue[ContinuousBackupInfoEncryptionInfoAttributes] {
	return terra.ReferenceAsList[ContinuousBackupInfoEncryptionInfoAttributes](cbi.ref.Append("encryption_info"))
}

type ContinuousBackupInfoEncryptionInfoAttributes struct {
	ref terra.Reference
}

func (ei ContinuousBackupInfoEncryptionInfoAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei ContinuousBackupInfoEncryptionInfoAttributes) InternalWithRef(ref terra.Reference) ContinuousBackupInfoEncryptionInfoAttributes {
	return ContinuousBackupInfoEncryptionInfoAttributes{ref: ref}
}

func (ei ContinuousBackupInfoEncryptionInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ei.ref.InternalTokens()
}

func (ei ContinuousBackupInfoEncryptionInfoAttributes) EncryptionType() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("encryption_type"))
}

func (ei ContinuousBackupInfoEncryptionInfoAttributes) KmsKeyVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ei.ref.Append("kms_key_versions"))
}

type EncryptionInfoAttributes struct {
	ref terra.Reference
}

func (ei EncryptionInfoAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei EncryptionInfoAttributes) InternalWithRef(ref terra.Reference) EncryptionInfoAttributes {
	return EncryptionInfoAttributes{ref: ref}
}

func (ei EncryptionInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ei.ref.InternalTokens()
}

func (ei EncryptionInfoAttributes) EncryptionType() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("encryption_type"))
}

func (ei EncryptionInfoAttributes) KmsKeyVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ei.ref.Append("kms_key_versions"))
}

type MigrationSourceAttributes struct {
	ref terra.Reference
}

func (ms MigrationSourceAttributes) InternalRef() (terra.Reference, error) {
	return ms.ref, nil
}

func (ms MigrationSourceAttributes) InternalWithRef(ref terra.Reference) MigrationSourceAttributes {
	return MigrationSourceAttributes{ref: ref}
}

func (ms MigrationSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ms.ref.InternalTokens()
}

func (ms MigrationSourceAttributes) HostPort() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("host_port"))
}

func (ms MigrationSourceAttributes) ReferenceId() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("reference_id"))
}

func (ms MigrationSourceAttributes) SourceType() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("source_type"))
}

type AutomatedBackupPolicyAttributes struct {
	ref terra.Reference
}

func (abp AutomatedBackupPolicyAttributes) InternalRef() (terra.Reference, error) {
	return abp.ref, nil
}

func (abp AutomatedBackupPolicyAttributes) InternalWithRef(ref terra.Reference) AutomatedBackupPolicyAttributes {
	return AutomatedBackupPolicyAttributes{ref: ref}
}

func (abp AutomatedBackupPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return abp.ref.InternalTokens()
}

func (abp AutomatedBackupPolicyAttributes) BackupWindow() terra.StringValue {
	return terra.ReferenceAsString(abp.ref.Append("backup_window"))
}

func (abp AutomatedBackupPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(abp.ref.Append("enabled"))
}

func (abp AutomatedBackupPolicyAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](abp.ref.Append("labels"))
}

func (abp AutomatedBackupPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(abp.ref.Append("location"))
}

func (abp AutomatedBackupPolicyAttributes) EncryptionConfig() terra.ListValue[AutomatedBackupPolicyEncryptionConfigAttributes] {
	return terra.ReferenceAsList[AutomatedBackupPolicyEncryptionConfigAttributes](abp.ref.Append("encryption_config"))
}

func (abp AutomatedBackupPolicyAttributes) QuantityBasedRetention() terra.ListValue[QuantityBasedRetentionAttributes] {
	return terra.ReferenceAsList[QuantityBasedRetentionAttributes](abp.ref.Append("quantity_based_retention"))
}

func (abp AutomatedBackupPolicyAttributes) TimeBasedRetention() terra.ListValue[TimeBasedRetentionAttributes] {
	return terra.ReferenceAsList[TimeBasedRetentionAttributes](abp.ref.Append("time_based_retention"))
}

func (abp AutomatedBackupPolicyAttributes) WeeklySchedule() terra.ListValue[WeeklyScheduleAttributes] {
	return terra.ReferenceAsList[WeeklyScheduleAttributes](abp.ref.Append("weekly_schedule"))
}

type AutomatedBackupPolicyEncryptionConfigAttributes struct {
	ref terra.Reference
}

func (ec AutomatedBackupPolicyEncryptionConfigAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec AutomatedBackupPolicyEncryptionConfigAttributes) InternalWithRef(ref terra.Reference) AutomatedBackupPolicyEncryptionConfigAttributes {
	return AutomatedBackupPolicyEncryptionConfigAttributes{ref: ref}
}

func (ec AutomatedBackupPolicyEncryptionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec AutomatedBackupPolicyEncryptionConfigAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key_name"))
}

type QuantityBasedRetentionAttributes struct {
	ref terra.Reference
}

func (qbr QuantityBasedRetentionAttributes) InternalRef() (terra.Reference, error) {
	return qbr.ref, nil
}

func (qbr QuantityBasedRetentionAttributes) InternalWithRef(ref terra.Reference) QuantityBasedRetentionAttributes {
	return QuantityBasedRetentionAttributes{ref: ref}
}

func (qbr QuantityBasedRetentionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qbr.ref.InternalTokens()
}

func (qbr QuantityBasedRetentionAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(qbr.ref.Append("count"))
}

type TimeBasedRetentionAttributes struct {
	ref terra.Reference
}

func (tbr TimeBasedRetentionAttributes) InternalRef() (terra.Reference, error) {
	return tbr.ref, nil
}

func (tbr TimeBasedRetentionAttributes) InternalWithRef(ref terra.Reference) TimeBasedRetentionAttributes {
	return TimeBasedRetentionAttributes{ref: ref}
}

func (tbr TimeBasedRetentionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tbr.ref.InternalTokens()
}

func (tbr TimeBasedRetentionAttributes) RetentionPeriod() terra.StringValue {
	return terra.ReferenceAsString(tbr.ref.Append("retention_period"))
}

type WeeklyScheduleAttributes struct {
	ref terra.Reference
}

func (ws WeeklyScheduleAttributes) InternalRef() (terra.Reference, error) {
	return ws.ref, nil
}

func (ws WeeklyScheduleAttributes) InternalWithRef(ref terra.Reference) WeeklyScheduleAttributes {
	return WeeklyScheduleAttributes{ref: ref}
}

func (ws WeeklyScheduleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ws.ref.InternalTokens()
}

func (ws WeeklyScheduleAttributes) DaysOfWeek() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ws.ref.Append("days_of_week"))
}

func (ws WeeklyScheduleAttributes) StartTimes() terra.ListValue[StartTimesAttributes] {
	return terra.ReferenceAsList[StartTimesAttributes](ws.ref.Append("start_times"))
}

type StartTimesAttributes struct {
	ref terra.Reference
}

func (st StartTimesAttributes) InternalRef() (terra.Reference, error) {
	return st.ref, nil
}

func (st StartTimesAttributes) InternalWithRef(ref terra.Reference) StartTimesAttributes {
	return StartTimesAttributes{ref: ref}
}

func (st StartTimesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return st.ref.InternalTokens()
}

func (st StartTimesAttributes) Hours() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("hours"))
}

func (st StartTimesAttributes) Minutes() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("minutes"))
}

func (st StartTimesAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("nanos"))
}

func (st StartTimesAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("seconds"))
}

type ContinuousBackupConfigAttributes struct {
	ref terra.Reference
}

func (cbc ContinuousBackupConfigAttributes) InternalRef() (terra.Reference, error) {
	return cbc.ref, nil
}

func (cbc ContinuousBackupConfigAttributes) InternalWithRef(ref terra.Reference) ContinuousBackupConfigAttributes {
	return ContinuousBackupConfigAttributes{ref: ref}
}

func (cbc ContinuousBackupConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cbc.ref.InternalTokens()
}

func (cbc ContinuousBackupConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cbc.ref.Append("enabled"))
}

func (cbc ContinuousBackupConfigAttributes) RecoveryWindowDays() terra.NumberValue {
	return terra.ReferenceAsNumber(cbc.ref.Append("recovery_window_days"))
}

func (cbc ContinuousBackupConfigAttributes) EncryptionConfig() terra.ListValue[ContinuousBackupConfigEncryptionConfigAttributes] {
	return terra.ReferenceAsList[ContinuousBackupConfigEncryptionConfigAttributes](cbc.ref.Append("encryption_config"))
}

type ContinuousBackupConfigEncryptionConfigAttributes struct {
	ref terra.Reference
}

func (ec ContinuousBackupConfigEncryptionConfigAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec ContinuousBackupConfigEncryptionConfigAttributes) InternalWithRef(ref terra.Reference) ContinuousBackupConfigEncryptionConfigAttributes {
	return ContinuousBackupConfigEncryptionConfigAttributes{ref: ref}
}

func (ec ContinuousBackupConfigEncryptionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec ContinuousBackupConfigEncryptionConfigAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key_name"))
}

type EncryptionConfigAttributes struct {
	ref terra.Reference
}

func (ec EncryptionConfigAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EncryptionConfigAttributes) InternalWithRef(ref terra.Reference) EncryptionConfigAttributes {
	return EncryptionConfigAttributes{ref: ref}
}

func (ec EncryptionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EncryptionConfigAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key_name"))
}

type InitialUserAttributes struct {
	ref terra.Reference
}

func (iu InitialUserAttributes) InternalRef() (terra.Reference, error) {
	return iu.ref, nil
}

func (iu InitialUserAttributes) InternalWithRef(ref terra.Reference) InitialUserAttributes {
	return InitialUserAttributes{ref: ref}
}

func (iu InitialUserAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iu.ref.InternalTokens()
}

func (iu InitialUserAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("password"))
}

func (iu InitialUserAttributes) User() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("user"))
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

type BackupSourceState struct {
	BackupName string `json:"backup_name"`
}

type ContinuousBackupInfoState struct {
	EarliestRestorableTime string                                    `json:"earliest_restorable_time"`
	EnabledTime            string                                    `json:"enabled_time"`
	Schedule               []string                                  `json:"schedule"`
	EncryptionInfo         []ContinuousBackupInfoEncryptionInfoState `json:"encryption_info"`
}

type ContinuousBackupInfoEncryptionInfoState struct {
	EncryptionType string   `json:"encryption_type"`
	KmsKeyVersions []string `json:"kms_key_versions"`
}

type EncryptionInfoState struct {
	EncryptionType string   `json:"encryption_type"`
	KmsKeyVersions []string `json:"kms_key_versions"`
}

type MigrationSourceState struct {
	HostPort    string `json:"host_port"`
	ReferenceId string `json:"reference_id"`
	SourceType  string `json:"source_type"`
}

type AutomatedBackupPolicyState struct {
	BackupWindow           string                                       `json:"backup_window"`
	Enabled                bool                                         `json:"enabled"`
	Labels                 map[string]string                            `json:"labels"`
	Location               string                                       `json:"location"`
	EncryptionConfig       []AutomatedBackupPolicyEncryptionConfigState `json:"encryption_config"`
	QuantityBasedRetention []QuantityBasedRetentionState                `json:"quantity_based_retention"`
	TimeBasedRetention     []TimeBasedRetentionState                    `json:"time_based_retention"`
	WeeklySchedule         []WeeklyScheduleState                        `json:"weekly_schedule"`
}

type AutomatedBackupPolicyEncryptionConfigState struct {
	KmsKeyName string `json:"kms_key_name"`
}

type QuantityBasedRetentionState struct {
	Count float64 `json:"count"`
}

type TimeBasedRetentionState struct {
	RetentionPeriod string `json:"retention_period"`
}

type WeeklyScheduleState struct {
	DaysOfWeek []string          `json:"days_of_week"`
	StartTimes []StartTimesState `json:"start_times"`
}

type StartTimesState struct {
	Hours   float64 `json:"hours"`
	Minutes float64 `json:"minutes"`
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type ContinuousBackupConfigState struct {
	Enabled            bool                                          `json:"enabled"`
	RecoveryWindowDays float64                                       `json:"recovery_window_days"`
	EncryptionConfig   []ContinuousBackupConfigEncryptionConfigState `json:"encryption_config"`
}

type ContinuousBackupConfigEncryptionConfigState struct {
	KmsKeyName string `json:"kms_key_name"`
}

type EncryptionConfigState struct {
	KmsKeyName string `json:"kms_key_name"`
}

type InitialUserState struct {
	Password string `json:"password"`
	User     string `json:"user"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
