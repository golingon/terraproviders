// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package devtestglobalvmshutdownschedule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type NotificationSettings struct {
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// TimeInMinutes: number, optional
	TimeInMinutes terra.NumberValue `hcl:"time_in_minutes,attr"`
	// WebhookUrl: string, optional
	WebhookUrl terra.StringValue `hcl:"webhook_url,attr"`
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

type NotificationSettingsAttributes struct {
	ref terra.Reference
}

func (ns NotificationSettingsAttributes) InternalRef() terra.Reference {
	return ns.ref
}

func (ns NotificationSettingsAttributes) InternalWithRef(ref terra.Reference) NotificationSettingsAttributes {
	return NotificationSettingsAttributes{ref: ref}
}

func (ns NotificationSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ns.ref.InternalTokens()
}

func (ns NotificationSettingsAttributes) Email() terra.StringValue {
	return terra.ReferenceString(ns.ref.Append("email"))
}

func (ns NotificationSettingsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(ns.ref.Append("enabled"))
}

func (ns NotificationSettingsAttributes) TimeInMinutes() terra.NumberValue {
	return terra.ReferenceNumber(ns.ref.Append("time_in_minutes"))
}

func (ns NotificationSettingsAttributes) WebhookUrl() terra.StringValue {
	return terra.ReferenceString(ns.ref.Append("webhook_url"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type NotificationSettingsState struct {
	Email         string  `json:"email"`
	Enabled       bool    `json:"enabled"`
	TimeInMinutes float64 `json:"time_in_minutes"`
	WebhookUrl    string  `json:"webhook_url"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
