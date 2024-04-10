// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package mssqlmanageddatabase

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type LongTermRetentionPolicy struct {
	// ImmutableBackupsEnabled: bool, optional
	ImmutableBackupsEnabled terra.BoolValue `hcl:"immutable_backups_enabled,attr"`
	// MonthlyRetention: string, optional
	MonthlyRetention terra.StringValue `hcl:"monthly_retention,attr"`
	// WeekOfYear: number, optional
	WeekOfYear terra.NumberValue `hcl:"week_of_year,attr"`
	// WeeklyRetention: string, optional
	WeeklyRetention terra.StringValue `hcl:"weekly_retention,attr"`
	// YearlyRetention: string, optional
	YearlyRetention terra.StringValue `hcl:"yearly_retention,attr"`
}

type PointInTimeRestore struct {
	// RestorePointInTime: string, required
	RestorePointInTime terra.StringValue `hcl:"restore_point_in_time,attr" validate:"required"`
	// SourceDatabaseId: string, required
	SourceDatabaseId terra.StringValue `hcl:"source_database_id,attr" validate:"required"`
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

type LongTermRetentionPolicyAttributes struct {
	ref terra.Reference
}

func (ltrp LongTermRetentionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return ltrp.ref, nil
}

func (ltrp LongTermRetentionPolicyAttributes) InternalWithRef(ref terra.Reference) LongTermRetentionPolicyAttributes {
	return LongTermRetentionPolicyAttributes{ref: ref}
}

func (ltrp LongTermRetentionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ltrp.ref.InternalTokens()
}

func (ltrp LongTermRetentionPolicyAttributes) ImmutableBackupsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ltrp.ref.Append("immutable_backups_enabled"))
}

func (ltrp LongTermRetentionPolicyAttributes) MonthlyRetention() terra.StringValue {
	return terra.ReferenceAsString(ltrp.ref.Append("monthly_retention"))
}

func (ltrp LongTermRetentionPolicyAttributes) WeekOfYear() terra.NumberValue {
	return terra.ReferenceAsNumber(ltrp.ref.Append("week_of_year"))
}

func (ltrp LongTermRetentionPolicyAttributes) WeeklyRetention() terra.StringValue {
	return terra.ReferenceAsString(ltrp.ref.Append("weekly_retention"))
}

func (ltrp LongTermRetentionPolicyAttributes) YearlyRetention() terra.StringValue {
	return terra.ReferenceAsString(ltrp.ref.Append("yearly_retention"))
}

type PointInTimeRestoreAttributes struct {
	ref terra.Reference
}

func (pitr PointInTimeRestoreAttributes) InternalRef() (terra.Reference, error) {
	return pitr.ref, nil
}

func (pitr PointInTimeRestoreAttributes) InternalWithRef(ref terra.Reference) PointInTimeRestoreAttributes {
	return PointInTimeRestoreAttributes{ref: ref}
}

func (pitr PointInTimeRestoreAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pitr.ref.InternalTokens()
}

func (pitr PointInTimeRestoreAttributes) RestorePointInTime() terra.StringValue {
	return terra.ReferenceAsString(pitr.ref.Append("restore_point_in_time"))
}

func (pitr PointInTimeRestoreAttributes) SourceDatabaseId() terra.StringValue {
	return terra.ReferenceAsString(pitr.ref.Append("source_database_id"))
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

type LongTermRetentionPolicyState struct {
	ImmutableBackupsEnabled bool    `json:"immutable_backups_enabled"`
	MonthlyRetention        string  `json:"monthly_retention"`
	WeekOfYear              float64 `json:"week_of_year"`
	WeeklyRetention         string  `json:"weekly_retention"`
	YearlyRetention         string  `json:"yearly_retention"`
}

type PointInTimeRestoreState struct {
	RestorePointInTime string `json:"restore_point_in_time"`
	SourceDatabaseId   string `json:"source_database_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
