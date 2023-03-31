// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datalocalnetworkgateway

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BgpSettings struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type BgpSettingsAttributes struct {
	ref terra.Reference
}

func (bs BgpSettingsAttributes) InternalRef() terra.Reference {
	return bs.ref
}

func (bs BgpSettingsAttributes) InternalWithRef(ref terra.Reference) BgpSettingsAttributes {
	return BgpSettingsAttributes{ref: ref}
}

func (bs BgpSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return bs.ref.InternalTokens()
}

func (bs BgpSettingsAttributes) Asn() terra.NumberValue {
	return terra.ReferenceNumber(bs.ref.Append("asn"))
}

func (bs BgpSettingsAttributes) BgpPeeringAddress() terra.StringValue {
	return terra.ReferenceString(bs.ref.Append("bgp_peering_address"))
}

func (bs BgpSettingsAttributes) PeerWeight() terra.NumberValue {
	return terra.ReferenceNumber(bs.ref.Append("peer_weight"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type BgpSettingsState struct {
	Asn               float64 `json:"asn"`
	BgpPeeringAddress string  `json:"bgp_peering_address"`
	PeerWeight        float64 `json:"peer_weight"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
