// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package logzmonitor

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Plan struct {
	// BillingCycle: string, required
	BillingCycle terra.StringValue `hcl:"billing_cycle,attr" validate:"required"`
	// EffectiveDate: string, required
	EffectiveDate terra.StringValue `hcl:"effective_date,attr" validate:"required"`
	// PlanId: string, optional
	PlanId terra.StringValue `hcl:"plan_id,attr"`
	// UsageType: string, required
	UsageType terra.StringValue `hcl:"usage_type,attr" validate:"required"`
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

type User struct {
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// FirstName: string, required
	FirstName terra.StringValue `hcl:"first_name,attr" validate:"required"`
	// LastName: string, required
	LastName terra.StringValue `hcl:"last_name,attr" validate:"required"`
	// PhoneNumber: string, required
	PhoneNumber terra.StringValue `hcl:"phone_number,attr" validate:"required"`
}

type PlanAttributes struct {
	ref terra.Reference
}

func (p PlanAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PlanAttributes) InternalWithRef(ref terra.Reference) PlanAttributes {
	return PlanAttributes{ref: ref}
}

func (p PlanAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PlanAttributes) BillingCycle() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("billing_cycle"))
}

func (p PlanAttributes) EffectiveDate() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("effective_date"))
}

func (p PlanAttributes) PlanId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("plan_id"))
}

func (p PlanAttributes) UsageType() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("usage_type"))
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

type UserAttributes struct {
	ref terra.Reference
}

func (u UserAttributes) InternalRef() (terra.Reference, error) {
	return u.ref, nil
}

func (u UserAttributes) InternalWithRef(ref terra.Reference) UserAttributes {
	return UserAttributes{ref: ref}
}

func (u UserAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return u.ref.InternalTokens()
}

func (u UserAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("email"))
}

func (u UserAttributes) FirstName() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("first_name"))
}

func (u UserAttributes) LastName() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("last_name"))
}

func (u UserAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("phone_number"))
}

type PlanState struct {
	BillingCycle  string `json:"billing_cycle"`
	EffectiveDate string `json:"effective_date"`
	PlanId        string `json:"plan_id"`
	UsageType     string `json:"usage_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type UserState struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}
