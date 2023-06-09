// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkmanagerconnectpeer

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Configuration struct {
	// BgpConfigurations: min=0
	BgpConfigurations []BgpConfigurations `hcl:"bgp_configurations,block" validate:"min=0"`
}

type BgpConfigurations struct{}

type BgpOptions struct {
	// PeerAsn: number, optional
	PeerAsn terra.NumberValue `hcl:"peer_asn,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type ConfigurationAttributes struct {
	ref terra.Reference
}

func (c ConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigurationAttributes) InternalWithRef(ref terra.Reference) ConfigurationAttributes {
	return ConfigurationAttributes{ref: ref}
}

func (c ConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigurationAttributes) CoreNetworkAddress() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("core_network_address"))
}

func (c ConfigurationAttributes) InsideCidrBlocks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("inside_cidr_blocks"))
}

func (c ConfigurationAttributes) PeerAddress() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("peer_address"))
}

func (c ConfigurationAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("protocol"))
}

func (c ConfigurationAttributes) BgpConfigurations() terra.ListValue[BgpConfigurationsAttributes] {
	return terra.ReferenceAsList[BgpConfigurationsAttributes](c.ref.Append("bgp_configurations"))
}

type BgpConfigurationsAttributes struct {
	ref terra.Reference
}

func (bc BgpConfigurationsAttributes) InternalRef() (terra.Reference, error) {
	return bc.ref, nil
}

func (bc BgpConfigurationsAttributes) InternalWithRef(ref terra.Reference) BgpConfigurationsAttributes {
	return BgpConfigurationsAttributes{ref: ref}
}

func (bc BgpConfigurationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bc.ref.InternalTokens()
}

func (bc BgpConfigurationsAttributes) CoreNetworkAddress() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("core_network_address"))
}

func (bc BgpConfigurationsAttributes) CoreNetworkAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(bc.ref.Append("core_network_asn"))
}

func (bc BgpConfigurationsAttributes) PeerAddress() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("peer_address"))
}

func (bc BgpConfigurationsAttributes) PeerAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(bc.ref.Append("peer_asn"))
}

type BgpOptionsAttributes struct {
	ref terra.Reference
}

func (bo BgpOptionsAttributes) InternalRef() (terra.Reference, error) {
	return bo.ref, nil
}

func (bo BgpOptionsAttributes) InternalWithRef(ref terra.Reference) BgpOptionsAttributes {
	return BgpOptionsAttributes{ref: ref}
}

func (bo BgpOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bo.ref.InternalTokens()
}

func (bo BgpOptionsAttributes) PeerAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(bo.ref.Append("peer_asn"))
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

type ConfigurationState struct {
	CoreNetworkAddress string                   `json:"core_network_address"`
	InsideCidrBlocks   []string                 `json:"inside_cidr_blocks"`
	PeerAddress        string                   `json:"peer_address"`
	Protocol           string                   `json:"protocol"`
	BgpConfigurations  []BgpConfigurationsState `json:"bgp_configurations"`
}

type BgpConfigurationsState struct {
	CoreNetworkAddress string  `json:"core_network_address"`
	CoreNetworkAsn     float64 `json:"core_network_asn"`
	PeerAddress        string  `json:"peer_address"`
	PeerAsn            float64 `json:"peer_asn"`
}

type BgpOptionsState struct {
	PeerAsn float64 `json:"peer_asn"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
