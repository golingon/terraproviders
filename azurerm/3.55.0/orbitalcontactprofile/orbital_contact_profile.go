// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package orbitalcontactprofile

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Links struct {
	// Direction: string, required
	Direction terra.StringValue `hcl:"direction,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Polarization: string, required
	Polarization terra.StringValue `hcl:"polarization,attr" validate:"required"`
	// Channels: min=1
	Channels []Channels `hcl:"channels,block" validate:"min=1"`
}

type Channels struct {
	// BandwidthMhz: number, required
	BandwidthMhz terra.NumberValue `hcl:"bandwidth_mhz,attr" validate:"required"`
	// CenterFrequencyMhz: number, required
	CenterFrequencyMhz terra.NumberValue `hcl:"center_frequency_mhz,attr" validate:"required"`
	// DemodulationConfiguration: string, optional
	DemodulationConfiguration terra.StringValue `hcl:"demodulation_configuration,attr"`
	// ModulationConfiguration: string, optional
	ModulationConfiguration terra.StringValue `hcl:"modulation_configuration,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// EndPoint: min=1
	EndPoint []EndPoint `hcl:"end_point,block" validate:"min=1"`
}

type EndPoint struct {
	// EndPointName: string, required
	EndPointName terra.StringValue `hcl:"end_point_name,attr" validate:"required"`
	// IpAddress: string, required
	IpAddress terra.StringValue `hcl:"ip_address,attr" validate:"required"`
	// Port: string, required
	Port terra.StringValue `hcl:"port,attr" validate:"required"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
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

func (l LinksAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("direction"))
}

func (l LinksAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("name"))
}

func (l LinksAttributes) Polarization() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("polarization"))
}

func (l LinksAttributes) Channels() terra.ListValue[ChannelsAttributes] {
	return terra.ReferenceAsList[ChannelsAttributes](l.ref.Append("channels"))
}

type ChannelsAttributes struct {
	ref terra.Reference
}

func (c ChannelsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ChannelsAttributes) InternalWithRef(ref terra.Reference) ChannelsAttributes {
	return ChannelsAttributes{ref: ref}
}

func (c ChannelsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ChannelsAttributes) BandwidthMhz() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("bandwidth_mhz"))
}

func (c ChannelsAttributes) CenterFrequencyMhz() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("center_frequency_mhz"))
}

func (c ChannelsAttributes) DemodulationConfiguration() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("demodulation_configuration"))
}

func (c ChannelsAttributes) ModulationConfiguration() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("modulation_configuration"))
}

func (c ChannelsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c ChannelsAttributes) EndPoint() terra.SetValue[EndPointAttributes] {
	return terra.ReferenceAsSet[EndPointAttributes](c.ref.Append("end_point"))
}

type EndPointAttributes struct {
	ref terra.Reference
}

func (ep EndPointAttributes) InternalRef() (terra.Reference, error) {
	return ep.ref, nil
}

func (ep EndPointAttributes) InternalWithRef(ref terra.Reference) EndPointAttributes {
	return EndPointAttributes{ref: ref}
}

func (ep EndPointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ep.ref.InternalTokens()
}

func (ep EndPointAttributes) EndPointName() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("end_point_name"))
}

func (ep EndPointAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("ip_address"))
}

func (ep EndPointAttributes) Port() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("port"))
}

func (ep EndPointAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("protocol"))
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
	Direction    string          `json:"direction"`
	Name         string          `json:"name"`
	Polarization string          `json:"polarization"`
	Channels     []ChannelsState `json:"channels"`
}

type ChannelsState struct {
	BandwidthMhz              float64         `json:"bandwidth_mhz"`
	CenterFrequencyMhz        float64         `json:"center_frequency_mhz"`
	DemodulationConfiguration string          `json:"demodulation_configuration"`
	ModulationConfiguration   string          `json:"modulation_configuration"`
	Name                      string          `json:"name"`
	EndPoint                  []EndPointState `json:"end_point"`
}

type EndPointState struct {
	EndPointName string `json:"end_point_name"`
	IpAddress    string `json:"ip_address"`
	Port         string `json:"port"`
	Protocol     string `json:"protocol"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}