// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mssqldatabase

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Import struct {
	// AdministratorLogin: string, required
	AdministratorLogin terra.StringValue `hcl:"administrator_login,attr" validate:"required"`
	// AdministratorLoginPassword: string, required
	AdministratorLoginPassword terra.StringValue `hcl:"administrator_login_password,attr" validate:"required"`
	// AuthenticationType: string, required
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr" validate:"required"`
	// StorageAccountId: string, optional
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr"`
	// StorageKey: string, required
	StorageKey terra.StringValue `hcl:"storage_key,attr" validate:"required"`
	// StorageKeyType: string, required
	StorageKeyType terra.StringValue `hcl:"storage_key_type,attr" validate:"required"`
	// StorageUri: string, required
	StorageUri terra.StringValue `hcl:"storage_uri,attr" validate:"required"`
}

type LongTermRetentionPolicy struct {
	// MonthlyRetention: string, optional
	MonthlyRetention terra.StringValue `hcl:"monthly_retention,attr"`
	// WeekOfYear: number, optional
	WeekOfYear terra.NumberValue `hcl:"week_of_year,attr"`
	// WeeklyRetention: string, optional
	WeeklyRetention terra.StringValue `hcl:"weekly_retention,attr"`
	// YearlyRetention: string, optional
	YearlyRetention terra.StringValue `hcl:"yearly_retention,attr"`
}

type ShortTermRetentionPolicy struct {
	// BackupIntervalInHours: number, optional
	BackupIntervalInHours terra.NumberValue `hcl:"backup_interval_in_hours,attr"`
	// RetentionDays: number, required
	RetentionDays terra.NumberValue `hcl:"retention_days,attr" validate:"required"`
}

type ThreatDetectionPolicy struct {
	// DisabledAlerts: set of string, optional
	DisabledAlerts terra.SetValue[terra.StringValue] `hcl:"disabled_alerts,attr"`
	// EmailAccountAdmins: string, optional
	EmailAccountAdmins terra.StringValue `hcl:"email_account_admins,attr"`
	// EmailAddresses: set of string, optional
	EmailAddresses terra.SetValue[terra.StringValue] `hcl:"email_addresses,attr"`
	// RetentionDays: number, optional
	RetentionDays terra.NumberValue `hcl:"retention_days,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// StorageAccountAccessKey: string, optional
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr"`
	// StorageEndpoint: string, optional
	StorageEndpoint terra.StringValue `hcl:"storage_endpoint,attr"`
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

type ImportAttributes struct {
	ref terra.Reference
}

func (i ImportAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i ImportAttributes) InternalWithRef(ref terra.Reference) ImportAttributes {
	return ImportAttributes{ref: ref}
}

func (i ImportAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i ImportAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("administrator_login"))
}

func (i ImportAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("administrator_login_password"))
}

func (i ImportAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("authentication_type"))
}

func (i ImportAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("storage_account_id"))
}

func (i ImportAttributes) StorageKey() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("storage_key"))
}

func (i ImportAttributes) StorageKeyType() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("storage_key_type"))
}

func (i ImportAttributes) StorageUri() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("storage_uri"))
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

type ShortTermRetentionPolicyAttributes struct {
	ref terra.Reference
}

func (strp ShortTermRetentionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return strp.ref, nil
}

func (strp ShortTermRetentionPolicyAttributes) InternalWithRef(ref terra.Reference) ShortTermRetentionPolicyAttributes {
	return ShortTermRetentionPolicyAttributes{ref: ref}
}

func (strp ShortTermRetentionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return strp.ref.InternalTokens()
}

func (strp ShortTermRetentionPolicyAttributes) BackupIntervalInHours() terra.NumberValue {
	return terra.ReferenceAsNumber(strp.ref.Append("backup_interval_in_hours"))
}

func (strp ShortTermRetentionPolicyAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(strp.ref.Append("retention_days"))
}

type ThreatDetectionPolicyAttributes struct {
	ref terra.Reference
}

func (tdp ThreatDetectionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return tdp.ref, nil
}

func (tdp ThreatDetectionPolicyAttributes) InternalWithRef(ref terra.Reference) ThreatDetectionPolicyAttributes {
	return ThreatDetectionPolicyAttributes{ref: ref}
}

func (tdp ThreatDetectionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tdp.ref.InternalTokens()
}

func (tdp ThreatDetectionPolicyAttributes) DisabledAlerts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tdp.ref.Append("disabled_alerts"))
}

func (tdp ThreatDetectionPolicyAttributes) EmailAccountAdmins() terra.StringValue {
	return terra.ReferenceAsString(tdp.ref.Append("email_account_admins"))
}

func (tdp ThreatDetectionPolicyAttributes) EmailAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tdp.ref.Append("email_addresses"))
}

func (tdp ThreatDetectionPolicyAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(tdp.ref.Append("retention_days"))
}

func (tdp ThreatDetectionPolicyAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(tdp.ref.Append("state"))
}

func (tdp ThreatDetectionPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(tdp.ref.Append("storage_account_access_key"))
}

func (tdp ThreatDetectionPolicyAttributes) StorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(tdp.ref.Append("storage_endpoint"))
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

type ImportState struct {
	AdministratorLogin         string `json:"administrator_login"`
	AdministratorLoginPassword string `json:"administrator_login_password"`
	AuthenticationType         string `json:"authentication_type"`
	StorageAccountId           string `json:"storage_account_id"`
	StorageKey                 string `json:"storage_key"`
	StorageKeyType             string `json:"storage_key_type"`
	StorageUri                 string `json:"storage_uri"`
}

type LongTermRetentionPolicyState struct {
	MonthlyRetention string  `json:"monthly_retention"`
	WeekOfYear       float64 `json:"week_of_year"`
	WeeklyRetention  string  `json:"weekly_retention"`
	YearlyRetention  string  `json:"yearly_retention"`
}

type ShortTermRetentionPolicyState struct {
	BackupIntervalInHours float64 `json:"backup_interval_in_hours"`
	RetentionDays         float64 `json:"retention_days"`
}

type ThreatDetectionPolicyState struct {
	DisabledAlerts          []string `json:"disabled_alerts"`
	EmailAccountAdmins      string   `json:"email_account_admins"`
	EmailAddresses          []string `json:"email_addresses"`
	RetentionDays           float64  `json:"retention_days"`
	State                   string   `json:"state"`
	StorageAccountAccessKey string   `json:"storage_account_access_key"`
	StorageEndpoint         string   `json:"storage_endpoint"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}