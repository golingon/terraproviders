// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datavirtualnetworkgateway

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BgpSettings struct{}

type CustomRoute struct{}

type IpConfiguration struct{}

type VpnClientConfiguration struct {
	// RevokedCertificate: min=0
	RevokedCertificate []RevokedCertificate `hcl:"revoked_certificate,block" validate:"min=0"`
	// RootCertificate: min=0
	RootCertificate []RootCertificate `hcl:"root_certificate,block" validate:"min=0"`
}

type RevokedCertificate struct{}

type RootCertificate struct{}

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

func (bs BgpSettingsAttributes) PeerWeight() terra.NumberValue {
	return terra.ReferenceAsNumber(bs.ref.Append("peer_weight"))
}

func (bs BgpSettingsAttributes) PeeringAddress() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("peering_address"))
}

type CustomRouteAttributes struct {
	ref terra.Reference
}

func (cr CustomRouteAttributes) InternalRef() (terra.Reference, error) {
	return cr.ref, nil
}

func (cr CustomRouteAttributes) InternalWithRef(ref terra.Reference) CustomRouteAttributes {
	return CustomRouteAttributes{ref: ref}
}

func (cr CustomRouteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cr.ref.InternalTokens()
}

func (cr CustomRouteAttributes) AddressPrefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cr.ref.Append("address_prefixes"))
}

type IpConfigurationAttributes struct {
	ref terra.Reference
}

func (ic IpConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic IpConfigurationAttributes) InternalWithRef(ref terra.Reference) IpConfigurationAttributes {
	return IpConfigurationAttributes{ref: ref}
}

func (ic IpConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic IpConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("id"))
}

func (ic IpConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("name"))
}

func (ic IpConfigurationAttributes) PrivateIpAddressAllocation() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("private_ip_address_allocation"))
}

func (ic IpConfigurationAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("public_ip_address_id"))
}

func (ic IpConfigurationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("subnet_id"))
}

type VpnClientConfigurationAttributes struct {
	ref terra.Reference
}

func (vcc VpnClientConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return vcc.ref, nil
}

func (vcc VpnClientConfigurationAttributes) InternalWithRef(ref terra.Reference) VpnClientConfigurationAttributes {
	return VpnClientConfigurationAttributes{ref: ref}
}

func (vcc VpnClientConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vcc.ref.InternalTokens()
}

func (vcc VpnClientConfigurationAttributes) AadAudience() terra.StringValue {
	return terra.ReferenceAsString(vcc.ref.Append("aad_audience"))
}

func (vcc VpnClientConfigurationAttributes) AadIssuer() terra.StringValue {
	return terra.ReferenceAsString(vcc.ref.Append("aad_issuer"))
}

func (vcc VpnClientConfigurationAttributes) AadTenant() terra.StringValue {
	return terra.ReferenceAsString(vcc.ref.Append("aad_tenant"))
}

func (vcc VpnClientConfigurationAttributes) AddressSpace() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vcc.ref.Append("address_space"))
}

func (vcc VpnClientConfigurationAttributes) RadiusServerAddress() terra.StringValue {
	return terra.ReferenceAsString(vcc.ref.Append("radius_server_address"))
}

func (vcc VpnClientConfigurationAttributes) RadiusServerSecret() terra.StringValue {
	return terra.ReferenceAsString(vcc.ref.Append("radius_server_secret"))
}

func (vcc VpnClientConfigurationAttributes) VpnClientProtocols() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vcc.ref.Append("vpn_client_protocols"))
}

func (vcc VpnClientConfigurationAttributes) RevokedCertificate() terra.ListValue[RevokedCertificateAttributes] {
	return terra.ReferenceAsList[RevokedCertificateAttributes](vcc.ref.Append("revoked_certificate"))
}

func (vcc VpnClientConfigurationAttributes) RootCertificate() terra.ListValue[RootCertificateAttributes] {
	return terra.ReferenceAsList[RootCertificateAttributes](vcc.ref.Append("root_certificate"))
}

type RevokedCertificateAttributes struct {
	ref terra.Reference
}

func (rc RevokedCertificateAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RevokedCertificateAttributes) InternalWithRef(ref terra.Reference) RevokedCertificateAttributes {
	return RevokedCertificateAttributes{ref: ref}
}

func (rc RevokedCertificateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RevokedCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("name"))
}

func (rc RevokedCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("thumbprint"))
}

type RootCertificateAttributes struct {
	ref terra.Reference
}

func (rc RootCertificateAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RootCertificateAttributes) InternalWithRef(ref terra.Reference) RootCertificateAttributes {
	return RootCertificateAttributes{ref: ref}
}

func (rc RootCertificateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RootCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("name"))
}

func (rc RootCertificateAttributes) PublicCertData() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("public_cert_data"))
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
	Asn            float64 `json:"asn"`
	PeerWeight     float64 `json:"peer_weight"`
	PeeringAddress string  `json:"peering_address"`
}

type CustomRouteState struct {
	AddressPrefixes []string `json:"address_prefixes"`
}

type IpConfigurationState struct {
	Id                         string `json:"id"`
	Name                       string `json:"name"`
	PrivateIpAddressAllocation string `json:"private_ip_address_allocation"`
	PublicIpAddressId          string `json:"public_ip_address_id"`
	SubnetId                   string `json:"subnet_id"`
}

type VpnClientConfigurationState struct {
	AadAudience         string                    `json:"aad_audience"`
	AadIssuer           string                    `json:"aad_issuer"`
	AadTenant           string                    `json:"aad_tenant"`
	AddressSpace        []string                  `json:"address_space"`
	RadiusServerAddress string                    `json:"radius_server_address"`
	RadiusServerSecret  string                    `json:"radius_server_secret"`
	VpnClientProtocols  []string                  `json:"vpn_client_protocols"`
	RevokedCertificate  []RevokedCertificateState `json:"revoked_certificate"`
	RootCertificate     []RootCertificateState    `json:"root_certificate"`
}

type RevokedCertificateState struct {
	Name       string `json:"name"`
	Thumbprint string `json:"thumbprint"`
}

type RootCertificateState struct {
	Name           string `json:"name"`
	PublicCertData string `json:"public_cert_data"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
