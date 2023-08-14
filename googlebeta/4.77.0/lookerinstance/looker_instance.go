// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lookerinstance

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AdminSettings struct {
	// AllowedEmailDomains: list of string, optional
	AllowedEmailDomains terra.ListValue[terra.StringValue] `hcl:"allowed_email_domains,attr"`
}

type DenyMaintenancePeriod struct {
	// EndDate: required
	EndDate *EndDate `hcl:"end_date,block" validate:"required"`
	// StartDate: required
	StartDate *StartDate `hcl:"start_date,block" validate:"required"`
	// Time: required
	Time *Time `hcl:"time,block" validate:"required"`
}

type EndDate struct {
	// Day: number, optional
	Day terra.NumberValue `hcl:"day,attr"`
	// Month: number, optional
	Month terra.NumberValue `hcl:"month,attr"`
	// Year: number, optional
	Year terra.NumberValue `hcl:"year,attr"`
}

type StartDate struct {
	// Day: number, optional
	Day terra.NumberValue `hcl:"day,attr"`
	// Month: number, optional
	Month terra.NumberValue `hcl:"month,attr"`
	// Year: number, optional
	Year terra.NumberValue `hcl:"year,attr"`
}

type Time struct {
	// Hours: number, optional
	Hours terra.NumberValue `hcl:"hours,attr"`
	// Minutes: number, optional
	Minutes terra.NumberValue `hcl:"minutes,attr"`
	// Nanos: number, optional
	Nanos terra.NumberValue `hcl:"nanos,attr"`
	// Seconds: number, optional
	Seconds terra.NumberValue `hcl:"seconds,attr"`
}

type EncryptionConfig struct {
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
}

type MaintenanceWindow struct {
	// DayOfWeek: string, required
	DayOfWeek terra.StringValue `hcl:"day_of_week,attr" validate:"required"`
	// StartTime: required
	StartTime *StartTime `hcl:"start_time,block" validate:"required"`
}

type StartTime struct {
	// Hours: number, optional
	Hours terra.NumberValue `hcl:"hours,attr"`
	// Minutes: number, optional
	Minutes terra.NumberValue `hcl:"minutes,attr"`
	// Nanos: number, optional
	Nanos terra.NumberValue `hcl:"nanos,attr"`
	// Seconds: number, optional
	Seconds terra.NumberValue `hcl:"seconds,attr"`
}

type OauthConfig struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type UserMetadata struct {
	// AdditionalDeveloperUserCount: number, optional
	AdditionalDeveloperUserCount terra.NumberValue `hcl:"additional_developer_user_count,attr"`
	// AdditionalStandardUserCount: number, optional
	AdditionalStandardUserCount terra.NumberValue `hcl:"additional_standard_user_count,attr"`
	// AdditionalViewerUserCount: number, optional
	AdditionalViewerUserCount terra.NumberValue `hcl:"additional_viewer_user_count,attr"`
}

type AdminSettingsAttributes struct {
	ref terra.Reference
}

func (as AdminSettingsAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as AdminSettingsAttributes) InternalWithRef(ref terra.Reference) AdminSettingsAttributes {
	return AdminSettingsAttributes{ref: ref}
}

func (as AdminSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as AdminSettingsAttributes) AllowedEmailDomains() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](as.ref.Append("allowed_email_domains"))
}

type DenyMaintenancePeriodAttributes struct {
	ref terra.Reference
}

func (dmp DenyMaintenancePeriodAttributes) InternalRef() (terra.Reference, error) {
	return dmp.ref, nil
}

func (dmp DenyMaintenancePeriodAttributes) InternalWithRef(ref terra.Reference) DenyMaintenancePeriodAttributes {
	return DenyMaintenancePeriodAttributes{ref: ref}
}

func (dmp DenyMaintenancePeriodAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dmp.ref.InternalTokens()
}

func (dmp DenyMaintenancePeriodAttributes) EndDate() terra.ListValue[EndDateAttributes] {
	return terra.ReferenceAsList[EndDateAttributes](dmp.ref.Append("end_date"))
}

func (dmp DenyMaintenancePeriodAttributes) StartDate() terra.ListValue[StartDateAttributes] {
	return terra.ReferenceAsList[StartDateAttributes](dmp.ref.Append("start_date"))
}

func (dmp DenyMaintenancePeriodAttributes) Time() terra.ListValue[TimeAttributes] {
	return terra.ReferenceAsList[TimeAttributes](dmp.ref.Append("time"))
}

type EndDateAttributes struct {
	ref terra.Reference
}

func (ed EndDateAttributes) InternalRef() (terra.Reference, error) {
	return ed.ref, nil
}

func (ed EndDateAttributes) InternalWithRef(ref terra.Reference) EndDateAttributes {
	return EndDateAttributes{ref: ref}
}

func (ed EndDateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ed.ref.InternalTokens()
}

func (ed EndDateAttributes) Day() terra.NumberValue {
	return terra.ReferenceAsNumber(ed.ref.Append("day"))
}

func (ed EndDateAttributes) Month() terra.NumberValue {
	return terra.ReferenceAsNumber(ed.ref.Append("month"))
}

func (ed EndDateAttributes) Year() terra.NumberValue {
	return terra.ReferenceAsNumber(ed.ref.Append("year"))
}

type StartDateAttributes struct {
	ref terra.Reference
}

func (sd StartDateAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd StartDateAttributes) InternalWithRef(ref terra.Reference) StartDateAttributes {
	return StartDateAttributes{ref: ref}
}

func (sd StartDateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd StartDateAttributes) Day() terra.NumberValue {
	return terra.ReferenceAsNumber(sd.ref.Append("day"))
}

func (sd StartDateAttributes) Month() terra.NumberValue {
	return terra.ReferenceAsNumber(sd.ref.Append("month"))
}

func (sd StartDateAttributes) Year() terra.NumberValue {
	return terra.ReferenceAsNumber(sd.ref.Append("year"))
}

type TimeAttributes struct {
	ref terra.Reference
}

func (t TimeAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeAttributes) InternalWithRef(ref terra.Reference) TimeAttributes {
	return TimeAttributes{ref: ref}
}

func (t TimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeAttributes) Hours() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("hours"))
}

func (t TimeAttributes) Minutes() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("minutes"))
}

func (t TimeAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("nanos"))
}

func (t TimeAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("seconds"))
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

func (ec EncryptionConfigAttributes) KmsKeyNameVersion() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key_name_version"))
}

func (ec EncryptionConfigAttributes) KmsKeyState() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key_state"))
}

type MaintenanceWindowAttributes struct {
	ref terra.Reference
}

func (mw MaintenanceWindowAttributes) InternalRef() (terra.Reference, error) {
	return mw.ref, nil
}

func (mw MaintenanceWindowAttributes) InternalWithRef(ref terra.Reference) MaintenanceWindowAttributes {
	return MaintenanceWindowAttributes{ref: ref}
}

func (mw MaintenanceWindowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mw.ref.InternalTokens()
}

func (mw MaintenanceWindowAttributes) DayOfWeek() terra.StringValue {
	return terra.ReferenceAsString(mw.ref.Append("day_of_week"))
}

func (mw MaintenanceWindowAttributes) StartTime() terra.ListValue[StartTimeAttributes] {
	return terra.ReferenceAsList[StartTimeAttributes](mw.ref.Append("start_time"))
}

type StartTimeAttributes struct {
	ref terra.Reference
}

func (st StartTimeAttributes) InternalRef() (terra.Reference, error) {
	return st.ref, nil
}

func (st StartTimeAttributes) InternalWithRef(ref terra.Reference) StartTimeAttributes {
	return StartTimeAttributes{ref: ref}
}

func (st StartTimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return st.ref.InternalTokens()
}

func (st StartTimeAttributes) Hours() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("hours"))
}

func (st StartTimeAttributes) Minutes() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("minutes"))
}

func (st StartTimeAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("nanos"))
}

func (st StartTimeAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("seconds"))
}

type OauthConfigAttributes struct {
	ref terra.Reference
}

func (oc OauthConfigAttributes) InternalRef() (terra.Reference, error) {
	return oc.ref, nil
}

func (oc OauthConfigAttributes) InternalWithRef(ref terra.Reference) OauthConfigAttributes {
	return OauthConfigAttributes{ref: ref}
}

func (oc OauthConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oc.ref.InternalTokens()
}

func (oc OauthConfigAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("client_id"))
}

func (oc OauthConfigAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("client_secret"))
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

type UserMetadataAttributes struct {
	ref terra.Reference
}

func (um UserMetadataAttributes) InternalRef() (terra.Reference, error) {
	return um.ref, nil
}

func (um UserMetadataAttributes) InternalWithRef(ref terra.Reference) UserMetadataAttributes {
	return UserMetadataAttributes{ref: ref}
}

func (um UserMetadataAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return um.ref.InternalTokens()
}

func (um UserMetadataAttributes) AdditionalDeveloperUserCount() terra.NumberValue {
	return terra.ReferenceAsNumber(um.ref.Append("additional_developer_user_count"))
}

func (um UserMetadataAttributes) AdditionalStandardUserCount() terra.NumberValue {
	return terra.ReferenceAsNumber(um.ref.Append("additional_standard_user_count"))
}

func (um UserMetadataAttributes) AdditionalViewerUserCount() terra.NumberValue {
	return terra.ReferenceAsNumber(um.ref.Append("additional_viewer_user_count"))
}

type AdminSettingsState struct {
	AllowedEmailDomains []string `json:"allowed_email_domains"`
}

type DenyMaintenancePeriodState struct {
	EndDate   []EndDateState   `json:"end_date"`
	StartDate []StartDateState `json:"start_date"`
	Time      []TimeState      `json:"time"`
}

type EndDateState struct {
	Day   float64 `json:"day"`
	Month float64 `json:"month"`
	Year  float64 `json:"year"`
}

type StartDateState struct {
	Day   float64 `json:"day"`
	Month float64 `json:"month"`
	Year  float64 `json:"year"`
}

type TimeState struct {
	Hours   float64 `json:"hours"`
	Minutes float64 `json:"minutes"`
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type EncryptionConfigState struct {
	KmsKeyName        string `json:"kms_key_name"`
	KmsKeyNameVersion string `json:"kms_key_name_version"`
	KmsKeyState       string `json:"kms_key_state"`
}

type MaintenanceWindowState struct {
	DayOfWeek string           `json:"day_of_week"`
	StartTime []StartTimeState `json:"start_time"`
}

type StartTimeState struct {
	Hours   float64 `json:"hours"`
	Minutes float64 `json:"minutes"`
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type OauthConfigState struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type UserMetadataState struct {
	AdditionalDeveloperUserCount float64 `json:"additional_developer_user_count"`
	AdditionalStandardUserCount  float64 `json:"additional_standard_user_count"`
	AdditionalViewerUserCount    float64 `json:"additional_viewer_user_count"`
}