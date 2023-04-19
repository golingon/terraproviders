// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package virtualnetworkgateway

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BgpSettings struct {
	// Asn: number, optional
	Asn terra.NumberValue `hcl:"asn,attr"`
	// PeerWeight: number, optional
	PeerWeight terra.NumberValue `hcl:"peer_weight,attr"`
	// PeeringAddresses: min=0,max=2
	PeeringAddresses []PeeringAddresses `hcl:"peering_addresses,block" validate:"min=0,max=2"`
}

type PeeringAddresses struct {
	// ApipaAddresses: list of string, optional
	ApipaAddresses terra.ListValue[terra.StringValue] `hcl:"apipa_addresses,attr"`
	// IpConfigurationName: string, optional
	IpConfigurationName terra.StringValue `hcl:"ip_configuration_name,attr"`
}

type CustomRoute struct {
	// AddressPrefixes: set of string, optional
	AddressPrefixes terra.SetValue[terra.StringValue] `hcl:"address_prefixes,attr"`
}

type IpConfiguration struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// PrivateIpAddressAllocation: string, optional
	PrivateIpAddressAllocation terra.StringValue `hcl:"private_ip_address_allocation,attr"`
	// PublicIpAddressId: string, required
	PublicIpAddressId terra.StringValue `hcl:"public_ip_address_id,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
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

type VpnClientConfiguration struct {
	// AadAudience: string, optional
	AadAudience terra.StringValue `hcl:"aad_audience,attr"`
	// AadIssuer: string, optional
	AadIssuer terra.StringValue `hcl:"aad_issuer,attr"`
	// AadTenant: string, optional
	AadTenant terra.StringValue `hcl:"aad_tenant,attr"`
	// AddressSpace: list of string, required
	AddressSpace terra.ListValue[terra.StringValue] `hcl:"address_space,attr" validate:"required"`
	// RadiusServerAddress: string, optional
	RadiusServerAddress terra.StringValue `hcl:"radius_server_address,attr"`
	// RadiusServerSecret: string, optional
	RadiusServerSecret terra.StringValue `hcl:"radius_server_secret,attr"`
	// VpnAuthTypes: set of string, optional
	VpnAuthTypes terra.SetValue[terra.StringValue] `hcl:"vpn_auth_types,attr"`
	// VpnClientProtocols: set of string, optional
	VpnClientProtocols terra.SetValue[terra.StringValue] `hcl:"vpn_client_protocols,attr"`
	// RevokedCertificate: min=0
	RevokedCertificate []RevokedCertificate `hcl:"revoked_certificate,block" validate:"min=0"`
	// RootCertificate: min=0
	RootCertificate []RootCertificate `hcl:"root_certificate,block" validate:"min=0"`
}

type RevokedCertificate struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Thumbprint: string, required
	Thumbprint terra.StringValue `hcl:"thumbprint,attr" validate:"required"`
}

type RootCertificate struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicCertData: string, required
	PublicCertData terra.StringValue `hcl:"public_cert_data,attr" validate:"required"`
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

func (bs BgpSettingsAttributes) PeeringAddresses() terra.ListValue[PeeringAddressesAttributes] {
	return terra.ReferenceAsList[PeeringAddressesAttributes](bs.ref.Append("peering_addresses"))
}

type PeeringAddressesAttributes struct {
	ref terra.Reference
}

func (pa PeeringAddressesAttributes) InternalRef() (terra.Reference, error) {
	return pa.ref, nil
}

func (pa PeeringAddressesAttributes) InternalWithRef(ref terra.Reference) PeeringAddressesAttributes {
	return PeeringAddressesAttributes{ref: ref}
}

func (pa PeeringAddressesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pa.ref.InternalTokens()
}

func (pa PeeringAddressesAttributes) ApipaAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pa.ref.Append("apipa_addresses"))
}

func (pa PeeringAddressesAttributes) DefaultAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pa.ref.Append("default_addresses"))
}

func (pa PeeringAddressesAttributes) IpConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("ip_configuration_name"))
}

func (pa PeeringAddressesAttributes) TunnelIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pa.ref.Append("tunnel_ip_addresses"))
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

func (vcc VpnClientConfigurationAttributes) VpnAuthTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vcc.ref.Append("vpn_auth_types"))
}

func (vcc VpnClientConfigurationAttributes) VpnClientProtocols() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vcc.ref.Append("vpn_client_protocols"))
}

func (vcc VpnClientConfigurationAttributes) RevokedCertificate() terra.SetValue[RevokedCertificateAttributes] {
	return terra.ReferenceAsSet[RevokedCertificateAttributes](vcc.ref.Append("revoked_certificate"))
}

func (vcc VpnClientConfigurationAttributes) RootCertificate() terra.SetValue[RootCertificateAttributes] {
	return terra.ReferenceAsSet[RootCertificateAttributes](vcc.ref.Append("root_certificate"))
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

type BgpSettingsState struct {
	Asn              float64                 `json:"asn"`
	PeerWeight       float64                 `json:"peer_weight"`
	PeeringAddresses []PeeringAddressesState `json:"peering_addresses"`
}

type PeeringAddressesState struct {
	ApipaAddresses      []string `json:"apipa_addresses"`
	DefaultAddresses    []string `json:"default_addresses"`
	IpConfigurationName string   `json:"ip_configuration_name"`
	TunnelIpAddresses   []string `json:"tunnel_ip_addresses"`
}

type CustomRouteState struct {
	AddressPrefixes []string `json:"address_prefixes"`
}

type IpConfigurationState struct {
	Name                       string `json:"name"`
	PrivateIpAddressAllocation string `json:"private_ip_address_allocation"`
	PublicIpAddressId          string `json:"public_ip_address_id"`
	SubnetId                   string `json:"subnet_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type VpnClientConfigurationState struct {
	AadAudience         string                    `json:"aad_audience"`
	AadIssuer           string                    `json:"aad_issuer"`
	AadTenant           string                    `json:"aad_tenant"`
	AddressSpace        []string                  `json:"address_space"`
	RadiusServerAddress string                    `json:"radius_server_address"`
	RadiusServerSecret  string                    `json:"radius_server_secret"`
	VpnAuthTypes        []string                  `json:"vpn_auth_types"`
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