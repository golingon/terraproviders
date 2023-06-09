// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package computerouterpeer

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AdvertisedIpRanges struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Range: string, required
	Range terra.StringValue `hcl:"range,attr" validate:"required"`
}

type Bfd struct {
	// MinReceiveInterval: number, optional
	MinReceiveInterval terra.NumberValue `hcl:"min_receive_interval,attr"`
	// MinTransmitInterval: number, optional
	MinTransmitInterval terra.NumberValue `hcl:"min_transmit_interval,attr"`
	// Multiplier: number, optional
	Multiplier terra.NumberValue `hcl:"multiplier,attr"`
	// SessionInitializationMode: string, required
	SessionInitializationMode terra.StringValue `hcl:"session_initialization_mode,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AdvertisedIpRangesAttributes struct {
	ref terra.Reference
}

func (air AdvertisedIpRangesAttributes) InternalRef() (terra.Reference, error) {
	return air.ref, nil
}

func (air AdvertisedIpRangesAttributes) InternalWithRef(ref terra.Reference) AdvertisedIpRangesAttributes {
	return AdvertisedIpRangesAttributes{ref: ref}
}

func (air AdvertisedIpRangesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return air.ref.InternalTokens()
}

func (air AdvertisedIpRangesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("description"))
}

func (air AdvertisedIpRangesAttributes) Range() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("range"))
}

type BfdAttributes struct {
	ref terra.Reference
}

func (b BfdAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b BfdAttributes) InternalWithRef(ref terra.Reference) BfdAttributes {
	return BfdAttributes{ref: ref}
}

func (b BfdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b BfdAttributes) MinReceiveInterval() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("min_receive_interval"))
}

func (b BfdAttributes) MinTransmitInterval() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("min_transmit_interval"))
}

func (b BfdAttributes) Multiplier() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("multiplier"))
}

func (b BfdAttributes) SessionInitializationMode() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("session_initialization_mode"))
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

type AdvertisedIpRangesState struct {
	Description string `json:"description"`
	Range       string `json:"range"`
}

type BfdState struct {
	MinReceiveInterval        float64 `json:"min_receive_interval"`
	MinTransmitInterval       float64 `json:"min_transmit_interval"`
	Multiplier                float64 `json:"multiplier"`
	SessionInitializationMode string  `json:"session_initialization_mode"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
