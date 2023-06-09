// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataexpressroutecircuit

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Peerings struct{}

type ServiceProviderProperties struct{}

type Sku struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type PeeringsAttributes struct {
	ref terra.Reference
}

func (p PeeringsAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PeeringsAttributes) InternalWithRef(ref terra.Reference) PeeringsAttributes {
	return PeeringsAttributes{ref: ref}
}

func (p PeeringsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PeeringsAttributes) AzureAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("azure_asn"))
}

func (p PeeringsAttributes) PeerAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("peer_asn"))
}

func (p PeeringsAttributes) PeeringType() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("peering_type"))
}

func (p PeeringsAttributes) PrimaryPeerAddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("primary_peer_address_prefix"))
}

func (p PeeringsAttributes) SecondaryPeerAddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("secondary_peer_address_prefix"))
}

func (p PeeringsAttributes) SharedKey() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("shared_key"))
}

func (p PeeringsAttributes) VlanId() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("vlan_id"))
}

type ServiceProviderPropertiesAttributes struct {
	ref terra.Reference
}

func (spp ServiceProviderPropertiesAttributes) InternalRef() (terra.Reference, error) {
	return spp.ref, nil
}

func (spp ServiceProviderPropertiesAttributes) InternalWithRef(ref terra.Reference) ServiceProviderPropertiesAttributes {
	return ServiceProviderPropertiesAttributes{ref: ref}
}

func (spp ServiceProviderPropertiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return spp.ref.InternalTokens()
}

func (spp ServiceProviderPropertiesAttributes) BandwidthInMbps() terra.NumberValue {
	return terra.ReferenceAsNumber(spp.ref.Append("bandwidth_in_mbps"))
}

func (spp ServiceProviderPropertiesAttributes) PeeringLocation() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("peering_location"))
}

func (spp ServiceProviderPropertiesAttributes) ServiceProviderName() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("service_provider_name"))
}

type SkuAttributes struct {
	ref terra.Reference
}

func (s SkuAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SkuAttributes) InternalWithRef(ref terra.Reference) SkuAttributes {
	return SkuAttributes{ref: ref}
}

func (s SkuAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SkuAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("family"))
}

func (s SkuAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("tier"))
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

type PeeringsState struct {
	AzureAsn                   float64 `json:"azure_asn"`
	PeerAsn                    float64 `json:"peer_asn"`
	PeeringType                string  `json:"peering_type"`
	PrimaryPeerAddressPrefix   string  `json:"primary_peer_address_prefix"`
	SecondaryPeerAddressPrefix string  `json:"secondary_peer_address_prefix"`
	SharedKey                  string  `json:"shared_key"`
	VlanId                     float64 `json:"vlan_id"`
}

type ServiceProviderPropertiesState struct {
	BandwidthInMbps     float64 `json:"bandwidth_in_mbps"`
	PeeringLocation     string  `json:"peering_location"`
	ServiceProviderName string  `json:"service_provider_name"`
}

type SkuState struct {
	Family string `json:"family"`
	Tier   string `json:"tier"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
