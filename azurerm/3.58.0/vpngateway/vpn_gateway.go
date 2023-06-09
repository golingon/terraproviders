// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vpngateway

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BgpSettings struct {
	// Asn: number, required
	Asn terra.NumberValue `hcl:"asn,attr" validate:"required"`
	// PeerWeight: number, required
	PeerWeight terra.NumberValue `hcl:"peer_weight,attr" validate:"required"`
	// Instance0BgpPeeringAddress: optional
	Instance0BgpPeeringAddress *Instance0BgpPeeringAddress `hcl:"instance_0_bgp_peering_address,block"`
	// Instance1BgpPeeringAddress: optional
	Instance1BgpPeeringAddress *Instance1BgpPeeringAddress `hcl:"instance_1_bgp_peering_address,block"`
}

type Instance0BgpPeeringAddress struct {
	// CustomIps: set of string, required
	CustomIps terra.SetValue[terra.StringValue] `hcl:"custom_ips,attr" validate:"required"`
}

type Instance1BgpPeeringAddress struct {
	// CustomIps: set of string, required
	CustomIps terra.SetValue[terra.StringValue] `hcl:"custom_ips,attr" validate:"required"`
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

type BgpSettingsAttributes struct {
	ref terra.Reference
}

func (bs BgpSettingsAttributes) InternalRef() (terra.Reference, error) {
	return bs.ref, nil
}

func (bs BgpSettingsAttributes) InternalWithRef(ref terra.Reference) BgpSettingsAttributes {
	return BgpSettingsAttributes{ref: ref}
}

func (bs BgpSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bs.ref.InternalTokens()
}

func (bs BgpSettingsAttributes) Asn() terra.NumberValue {
	return terra.ReferenceAsNumber(bs.ref.Append("asn"))
}

func (bs BgpSettingsAttributes) BgpPeeringAddress() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("bgp_peering_address"))
}

func (bs BgpSettingsAttributes) PeerWeight() terra.NumberValue {
	return terra.ReferenceAsNumber(bs.ref.Append("peer_weight"))
}

func (bs BgpSettingsAttributes) Instance0BgpPeeringAddress() terra.ListValue[Instance0BgpPeeringAddressAttributes] {
	return terra.ReferenceAsList[Instance0BgpPeeringAddressAttributes](bs.ref.Append("instance_0_bgp_peering_address"))
}

func (bs BgpSettingsAttributes) Instance1BgpPeeringAddress() terra.ListValue[Instance1BgpPeeringAddressAttributes] {
	return terra.ReferenceAsList[Instance1BgpPeeringAddressAttributes](bs.ref.Append("instance_1_bgp_peering_address"))
}

type Instance0BgpPeeringAddressAttributes struct {
	ref terra.Reference
}

func (i0bpa Instance0BgpPeeringAddressAttributes) InternalRef() (terra.Reference, error) {
	return i0bpa.ref, nil
}

func (i0bpa Instance0BgpPeeringAddressAttributes) InternalWithRef(ref terra.Reference) Instance0BgpPeeringAddressAttributes {
	return Instance0BgpPeeringAddressAttributes{ref: ref}
}

func (i0bpa Instance0BgpPeeringAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i0bpa.ref.InternalTokens()
}

func (i0bpa Instance0BgpPeeringAddressAttributes) CustomIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i0bpa.ref.Append("custom_ips"))
}

func (i0bpa Instance0BgpPeeringAddressAttributes) DefaultIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i0bpa.ref.Append("default_ips"))
}

func (i0bpa Instance0BgpPeeringAddressAttributes) IpConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(i0bpa.ref.Append("ip_configuration_id"))
}

func (i0bpa Instance0BgpPeeringAddressAttributes) TunnelIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i0bpa.ref.Append("tunnel_ips"))
}

type Instance1BgpPeeringAddressAttributes struct {
	ref terra.Reference
}

func (i1bpa Instance1BgpPeeringAddressAttributes) InternalRef() (terra.Reference, error) {
	return i1bpa.ref, nil
}

func (i1bpa Instance1BgpPeeringAddressAttributes) InternalWithRef(ref terra.Reference) Instance1BgpPeeringAddressAttributes {
	return Instance1BgpPeeringAddressAttributes{ref: ref}
}

func (i1bpa Instance1BgpPeeringAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i1bpa.ref.InternalTokens()
}

func (i1bpa Instance1BgpPeeringAddressAttributes) CustomIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i1bpa.ref.Append("custom_ips"))
}

func (i1bpa Instance1BgpPeeringAddressAttributes) DefaultIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i1bpa.ref.Append("default_ips"))
}

func (i1bpa Instance1BgpPeeringAddressAttributes) IpConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(i1bpa.ref.Append("ip_configuration_id"))
}

func (i1bpa Instance1BgpPeeringAddressAttributes) TunnelIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i1bpa.ref.Append("tunnel_ips"))
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

type BgpSettingsState struct {
	Asn                        float64                           `json:"asn"`
	BgpPeeringAddress          string                            `json:"bgp_peering_address"`
	PeerWeight                 float64                           `json:"peer_weight"`
	Instance0BgpPeeringAddress []Instance0BgpPeeringAddressState `json:"instance_0_bgp_peering_address"`
	Instance1BgpPeeringAddress []Instance1BgpPeeringAddressState `json:"instance_1_bgp_peering_address"`
}

type Instance0BgpPeeringAddressState struct {
	CustomIps         []string `json:"custom_ips"`
	DefaultIps        []string `json:"default_ips"`
	IpConfigurationId string   `json:"ip_configuration_id"`
	TunnelIps         []string `json:"tunnel_ips"`
}

type Instance1BgpPeeringAddressState struct {
	CustomIps         []string `json:"custom_ips"`
	DefaultIps        []string `json:"default_ips"`
	IpConfigurationId string   `json:"ip_configuration_id"`
	TunnelIps         []string `json:"tunnel_ips"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
