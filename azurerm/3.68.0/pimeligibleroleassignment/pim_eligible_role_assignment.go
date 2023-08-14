// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package pimeligibleroleassignment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Schedule struct {
	// StartDateTime: string, optional
	StartDateTime terra.StringValue `hcl:"start_date_time,attr"`
	// Expiration: optional
	Expiration *Expiration `hcl:"expiration,block"`
}

type Expiration struct {
	// DurationDays: number, optional
	DurationDays terra.NumberValue `hcl:"duration_days,attr"`
	// DurationHours: number, optional
	DurationHours terra.NumberValue `hcl:"duration_hours,attr"`
	// EndDateTime: string, optional
	EndDateTime terra.StringValue `hcl:"end_date_time,attr"`
}

type Ticket struct {
	// Number: string, optional
	Number terra.StringValue `hcl:"number,attr"`
	// System: string, optional
	System terra.StringValue `hcl:"system,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
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

func (s ScheduleAttributes) StartDateTime() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("start_date_time"))
}

func (s ScheduleAttributes) Expiration() terra.ListValue[ExpirationAttributes] {
	return terra.ReferenceAsList[ExpirationAttributes](s.ref.Append("expiration"))
}

type ExpirationAttributes struct {
	ref terra.Reference
}

func (e ExpirationAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ExpirationAttributes) InternalWithRef(ref terra.Reference) ExpirationAttributes {
	return ExpirationAttributes{ref: ref}
}

func (e ExpirationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ExpirationAttributes) DurationDays() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("duration_days"))
}

func (e ExpirationAttributes) DurationHours() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("duration_hours"))
}

func (e ExpirationAttributes) EndDateTime() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("end_date_time"))
}

type TicketAttributes struct {
	ref terra.Reference
}

func (t TicketAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TicketAttributes) InternalWithRef(ref terra.Reference) TicketAttributes {
	return TicketAttributes{ref: ref}
}

func (t TicketAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TicketAttributes) Number() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("number"))
}

func (t TicketAttributes) System() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("system"))
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

type ScheduleState struct {
	StartDateTime string            `json:"start_date_time"`
	Expiration    []ExpirationState `json:"expiration"`
}

type ExpirationState struct {
	DurationDays  float64 `json:"duration_days"`
	DurationHours float64 `json:"duration_hours"`
	EndDateTime   string  `json:"end_date_time"`
}

type TicketState struct {
	Number string `json:"number"`
	System string `json:"system"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
}