// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package orbitalspacecraft

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Links struct {
	// BandwidthMhz: number, required
	BandwidthMhz terra.NumberValue `hcl:"bandwidth_mhz,attr" validate:"required"`
	// CenterFrequencyMhz: number, required
	CenterFrequencyMhz terra.NumberValue `hcl:"center_frequency_mhz,attr" validate:"required"`
	// Direction: string, required
	Direction terra.StringValue `hcl:"direction,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Polarization: string, required
	Polarization terra.StringValue `hcl:"polarization,attr" validate:"required"`
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

type LinksAttributes struct {
	ref terra.Reference
}

func (l LinksAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LinksAttributes) InternalWithRef(ref terra.Reference) LinksAttributes {
	return LinksAttributes{ref: ref}
}

func (l LinksAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LinksAttributes) BandwidthMhz() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("bandwidth_mhz"))
}

func (l LinksAttributes) CenterFrequencyMhz() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("center_frequency_mhz"))
}

func (l LinksAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("direction"))
}

func (l LinksAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("name"))
}

func (l LinksAttributes) Polarization() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("polarization"))
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

type LinksState struct {
	BandwidthMhz       float64 `json:"bandwidth_mhz"`
	CenterFrequencyMhz float64 `json:"center_frequency_mhz"`
	Direction          string  `json:"direction"`
	Name               string  `json:"name"`
	Polarization       string  `json:"polarization"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}