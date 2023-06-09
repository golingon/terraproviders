// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package securityscannerscanconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Authentication struct {
	// CustomAccount: optional
	CustomAccount *CustomAccount `hcl:"custom_account,block"`
	// GoogleAccount: optional
	GoogleAccount *GoogleAccount `hcl:"google_account,block"`
}

type CustomAccount struct {
	// LoginUrl: string, required
	LoginUrl terra.StringValue `hcl:"login_url,attr" validate:"required"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type GoogleAccount struct {
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type Schedule struct {
	// IntervalDurationDays: number, required
	IntervalDurationDays terra.NumberValue `hcl:"interval_duration_days,attr" validate:"required"`
	// ScheduleTime: string, optional
	ScheduleTime terra.StringValue `hcl:"schedule_time,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AuthenticationAttributes struct {
	ref terra.Reference
}

func (a AuthenticationAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AuthenticationAttributes) InternalWithRef(ref terra.Reference) AuthenticationAttributes {
	return AuthenticationAttributes{ref: ref}
}

func (a AuthenticationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AuthenticationAttributes) CustomAccount() terra.ListValue[CustomAccountAttributes] {
	return terra.ReferenceAsList[CustomAccountAttributes](a.ref.Append("custom_account"))
}

func (a AuthenticationAttributes) GoogleAccount() terra.ListValue[GoogleAccountAttributes] {
	return terra.ReferenceAsList[GoogleAccountAttributes](a.ref.Append("google_account"))
}

type CustomAccountAttributes struct {
	ref terra.Reference
}

func (ca CustomAccountAttributes) InternalRef() (terra.Reference, error) {
	return ca.ref, nil
}

func (ca CustomAccountAttributes) InternalWithRef(ref terra.Reference) CustomAccountAttributes {
	return CustomAccountAttributes{ref: ref}
}

func (ca CustomAccountAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ca.ref.InternalTokens()
}

func (ca CustomAccountAttributes) LoginUrl() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("login_url"))
}

func (ca CustomAccountAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("password"))
}

func (ca CustomAccountAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("username"))
}

type GoogleAccountAttributes struct {
	ref terra.Reference
}

func (ga GoogleAccountAttributes) InternalRef() (terra.Reference, error) {
	return ga.ref, nil
}

func (ga GoogleAccountAttributes) InternalWithRef(ref terra.Reference) GoogleAccountAttributes {
	return GoogleAccountAttributes{ref: ref}
}

func (ga GoogleAccountAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ga.ref.InternalTokens()
}

func (ga GoogleAccountAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("password"))
}

func (ga GoogleAccountAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("username"))
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

func (s ScheduleAttributes) IntervalDurationDays() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("interval_duration_days"))
}

func (s ScheduleAttributes) ScheduleTime() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("schedule_time"))
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

type AuthenticationState struct {
	CustomAccount []CustomAccountState `json:"custom_account"`
	GoogleAccount []GoogleAccountState `json:"google_account"`
}

type CustomAccountState struct {
	LoginUrl string `json:"login_url"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type GoogleAccountState struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type ScheduleState struct {
	IntervalDurationDays float64 `json:"interval_duration_days"`
	ScheduleTime         string  `json:"schedule_time"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
