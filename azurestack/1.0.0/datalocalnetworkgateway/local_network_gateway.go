// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datalocalnetworkgateway

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type BgpSettings struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type BgpSettingsState struct {
	Asn               float64 `json:"asn"`
	BgpPeeringAddress string  `json:"bgp_peering_address"`
	PeerWeight        float64 `json:"peer_weight"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
