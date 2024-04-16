// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_dev_test_global_vm_shutdown_schedule

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
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

func (ns NotificationSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ns.ref, nil
}

func (ns NotificationSettingsAttributes) InternalWithRef(ref terra.Reference) NotificationSettingsAttributes {
	return NotificationSettingsAttributes{ref: ref}
}

func (ns NotificationSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ns.ref.InternalTokens()
}

func (ns NotificationSettingsAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("email"))
}

func (ns NotificationSettingsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ns.ref.Append("enabled"))
}

func (ns NotificationSettingsAttributes) TimeInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(ns.ref.Append("time_in_minutes"))
}

func (ns NotificationSettingsAttributes) WebhookUrl() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("webhook_url"))
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
