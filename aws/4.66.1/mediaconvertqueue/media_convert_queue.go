// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mediaconvertqueue

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ReservationPlanSettings struct {
	// Commitment: string, required
	Commitment terra.StringValue `hcl:"commitment,attr" validate:"required"`
	// RenewalType: string, required
	RenewalType terra.StringValue `hcl:"renewal_type,attr" validate:"required"`
	// ReservedSlots: number, required
	ReservedSlots terra.NumberValue `hcl:"reserved_slots,attr" validate:"required"`
}

type ReservationPlanSettingsAttributes struct {
	ref terra.Reference
}

func (rps ReservationPlanSettingsAttributes) InternalRef() (terra.Reference, error) {
	return rps.ref, nil
}

func (rps ReservationPlanSettingsAttributes) InternalWithRef(ref terra.Reference) ReservationPlanSettingsAttributes {
	return ReservationPlanSettingsAttributes{ref: ref}
}

func (rps ReservationPlanSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rps.ref.InternalTokens()
}

func (rps ReservationPlanSettingsAttributes) Commitment() terra.StringValue {
	return terra.ReferenceAsString(rps.ref.Append("commitment"))
}

func (rps ReservationPlanSettingsAttributes) RenewalType() terra.StringValue {
	return terra.ReferenceAsString(rps.ref.Append("renewal_type"))
}

func (rps ReservationPlanSettingsAttributes) ReservedSlots() terra.NumberValue {
	return terra.ReferenceAsNumber(rps.ref.Append("reserved_slots"))
}

type ReservationPlanSettingsState struct {
	Commitment    string  `json:"commitment"`
	RenewalType   string  `json:"renewal_type"`
	ReservedSlots float64 `json:"reserved_slots"`
}
