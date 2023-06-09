// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package labserviceschedule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Recurrence struct {
	// ExpirationDate: string, required
	ExpirationDate terra.StringValue `hcl:"expiration_date,attr" validate:"required"`
	// Frequency: string, required
	Frequency terra.StringValue `hcl:"frequency,attr" validate:"required"`
	// Interval: number, optional
	Interval terra.NumberValue `hcl:"interval,attr"`
	// WeekDays: list of string, optional
	WeekDays terra.ListValue[terra.StringValue] `hcl:"week_days,attr"`
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

type RecurrenceAttributes struct {
	ref terra.Reference
}

func (r RecurrenceAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RecurrenceAttributes) InternalWithRef(ref terra.Reference) RecurrenceAttributes {
	return RecurrenceAttributes{ref: ref}
}

func (r RecurrenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RecurrenceAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("expiration_date"))
}

func (r RecurrenceAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("frequency"))
}

func (r RecurrenceAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("interval"))
}

func (r RecurrenceAttributes) WeekDays() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("week_days"))
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

type RecurrenceState struct {
	ExpirationDate string   `json:"expiration_date"`
	Frequency      string   `json:"frequency"`
	Interval       float64  `json:"interval"`
	WeekDays       []string `json:"week_days"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
