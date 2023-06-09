// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package rdsreservedinstance

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RecurringCharges struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type RecurringChargesAttributes struct {
	ref terra.Reference
}

func (rc RecurringChargesAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RecurringChargesAttributes) InternalWithRef(ref terra.Reference) RecurringChargesAttributes {
	return RecurringChargesAttributes{ref: ref}
}

func (rc RecurringChargesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RecurringChargesAttributes) RecurringChargeAmount() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("recurring_charge_amount"))
}

func (rc RecurringChargesAttributes) RecurringChargeFrequency() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("recurring_charge_frequency"))
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

type RecurringChargesState struct {
	RecurringChargeAmount    float64 `json:"recurring_charge_amount"`
	RecurringChargeFrequency string  `json:"recurring_charge_frequency"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
