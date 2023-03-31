// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataapimanagement

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AdditionalLocation struct{}

type HostnameConfiguration struct {
	// DeveloperPortal: min=0
	DeveloperPortal []DeveloperPortal `hcl:"developer_portal,block" validate:"min=0"`
	// Management: min=0
	Management []Management `hcl:"management,block" validate:"min=0"`
	// Portal: min=0
	Portal []Portal `hcl:"portal,block" validate:"min=0"`
	// Proxy: min=0
	Proxy []Proxy `hcl:"proxy,block" validate:"min=0"`
	// Scm: min=0
	Scm []Scm `hcl:"scm,block" validate:"min=0"`
}

type DeveloperPortal struct{}

type Management struct{}

type Portal struct{}

type Proxy struct{}

type Scm struct{}

type Identity struct{}

type TenantAccess struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type AdditionalLocationAttributes struct {
	ref terra.Reference
}

func (al AdditionalLocationAttributes) InternalRef() terra.Reference {
	return al.ref
}

func (al AdditionalLocationAttributes) InternalWithRef(ref terra.Reference) AdditionalLocationAttributes {
	return AdditionalLocationAttributes{ref: ref}
}

func (al AdditionalLocationAttributes) InternalTokens() hclwrite.Tokens {
	return al.ref.InternalTokens()
}

func (al AdditionalLocationAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceNumber(al.ref.Append("capacity"))
}

func (al AdditionalLocationAttributes) GatewayRegionalUrl() terra.StringValue {
	return terra.ReferenceString(al.ref.Append("gateway_regional_url"))
}

func (al AdditionalLocationAttributes) Location() terra.StringValue {
	return terra.ReferenceString(al.ref.Append("location"))
}

func (al AdditionalLocationAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](al.ref.Append("private_ip_addresses"))
}

func (al AdditionalLocationAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceString(al.ref.Append("public_ip_address_id"))
}

func (al AdditionalLocationAttributes) PublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](al.ref.Append("public_ip_addresses"))
}

func (al AdditionalLocationAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](al.ref.Append("zones"))
}

type HostnameConfigurationAttributes struct {
	ref terra.Reference
}

func (hc HostnameConfigurationAttributes) InternalRef() terra.Reference {
	return hc.ref
}

func (hc HostnameConfigurationAttributes) InternalWithRef(ref terra.Reference) HostnameConfigurationAttributes {
	return HostnameConfigurationAttributes{ref: ref}
}

func (hc HostnameConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return hc.ref.InternalTokens()
}

func (hc HostnameConfigurationAttributes) DeveloperPortal() terra.ListValue[DeveloperPortalAttributes] {
	return terra.ReferenceList[DeveloperPortalAttributes](hc.ref.Append("developer_portal"))
}

func (hc HostnameConfigurationAttributes) Management() terra.ListValue[ManagementAttributes] {
	return terra.ReferenceList[ManagementAttributes](hc.ref.Append("management"))
}

func (hc HostnameConfigurationAttributes) Portal() terra.ListValue[PortalAttributes] {
	return terra.ReferenceList[PortalAttributes](hc.ref.Append("portal"))
}

func (hc HostnameConfigurationAttributes) Proxy() terra.ListValue[ProxyAttributes] {
	return terra.ReferenceList[ProxyAttributes](hc.ref.Append("proxy"))
}

func (hc HostnameConfigurationAttributes) Scm() terra.ListValue[ScmAttributes] {
	return terra.ReferenceList[ScmAttributes](hc.ref.Append("scm"))
}

type DeveloperPortalAttributes struct {
	ref terra.Reference
}

func (dp DeveloperPortalAttributes) InternalRef() terra.Reference {
	return dp.ref
}

func (dp DeveloperPortalAttributes) InternalWithRef(ref terra.Reference) DeveloperPortalAttributes {
	return DeveloperPortalAttributes{ref: ref}
}

func (dp DeveloperPortalAttributes) InternalTokens() hclwrite.Tokens {
	return dp.ref.InternalTokens()
}

func (dp DeveloperPortalAttributes) HostName() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("host_name"))
}

func (dp DeveloperPortalAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("key_vault_id"))
}

func (dp DeveloperPortalAttributes) NegotiateClientCertificate() terra.BoolValue {
	return terra.ReferenceBool(dp.ref.Append("negotiate_client_certificate"))
}

type ManagementAttributes struct {
	ref terra.Reference
}

func (m ManagementAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m ManagementAttributes) InternalWithRef(ref terra.Reference) ManagementAttributes {
	return ManagementAttributes{ref: ref}
}

func (m ManagementAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m ManagementAttributes) HostName() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("host_name"))
}

func (m ManagementAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("key_vault_id"))
}

func (m ManagementAttributes) NegotiateClientCertificate() terra.BoolValue {
	return terra.ReferenceBool(m.ref.Append("negotiate_client_certificate"))
}

type PortalAttributes struct {
	ref terra.Reference
}

func (p PortalAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p PortalAttributes) InternalWithRef(ref terra.Reference) PortalAttributes {
	return PortalAttributes{ref: ref}
}

func (p PortalAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p PortalAttributes) HostName() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("host_name"))
}

func (p PortalAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("key_vault_id"))
}

func (p PortalAttributes) NegotiateClientCertificate() terra.BoolValue {
	return terra.ReferenceBool(p.ref.Append("negotiate_client_certificate"))
}

type ProxyAttributes struct {
	ref terra.Reference
}

func (p ProxyAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p ProxyAttributes) InternalWithRef(ref terra.Reference) ProxyAttributes {
	return ProxyAttributes{ref: ref}
}

func (p ProxyAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p ProxyAttributes) DefaultSslBinding() terra.BoolValue {
	return terra.ReferenceBool(p.ref.Append("default_ssl_binding"))
}

func (p ProxyAttributes) HostName() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("host_name"))
}

func (p ProxyAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("key_vault_id"))
}

func (p ProxyAttributes) NegotiateClientCertificate() terra.BoolValue {
	return terra.ReferenceBool(p.ref.Append("negotiate_client_certificate"))
}

type ScmAttributes struct {
	ref terra.Reference
}

func (s ScmAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s ScmAttributes) InternalWithRef(ref terra.Reference) ScmAttributes {
	return ScmAttributes{ref: ref}
}

func (s ScmAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s ScmAttributes) HostName() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("host_name"))
}

func (s ScmAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("key_vault_id"))
}

func (s ScmAttributes) NegotiateClientCertificate() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("negotiate_client_certificate"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() terra.Reference {
	return i.ref
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() hclwrite.Tokens {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("type"))
}

type TenantAccessAttributes struct {
	ref terra.Reference
}

func (ta TenantAccessAttributes) InternalRef() terra.Reference {
	return ta.ref
}

func (ta TenantAccessAttributes) InternalWithRef(ref terra.Reference) TenantAccessAttributes {
	return TenantAccessAttributes{ref: ref}
}

func (ta TenantAccessAttributes) InternalTokens() hclwrite.Tokens {
	return ta.ref.InternalTokens()
}

func (ta TenantAccessAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(ta.ref.Append("enabled"))
}

func (ta TenantAccessAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceString(ta.ref.Append("primary_key"))
}

func (ta TenantAccessAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceString(ta.ref.Append("secondary_key"))
}

func (ta TenantAccessAttributes) TenantId() terra.StringValue {
	return terra.ReferenceString(ta.ref.Append("tenant_id"))
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

type AdditionalLocationState struct {
	Capacity           float64  `json:"capacity"`
	GatewayRegionalUrl string   `json:"gateway_regional_url"`
	Location           string   `json:"location"`
	PrivateIpAddresses []string `json:"private_ip_addresses"`
	PublicIpAddressId  string   `json:"public_ip_address_id"`
	PublicIpAddresses  []string `json:"public_ip_addresses"`
	Zones              []string `json:"zones"`
}

type HostnameConfigurationState struct {
	DeveloperPortal []DeveloperPortalState `json:"developer_portal"`
	Management      []ManagementState      `json:"management"`
	Portal          []PortalState          `json:"portal"`
	Proxy           []ProxyState           `json:"proxy"`
	Scm             []ScmState             `json:"scm"`
}

type DeveloperPortalState struct {
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type ManagementState struct {
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type PortalState struct {
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type ProxyState struct {
	DefaultSslBinding          bool   `json:"default_ssl_binding"`
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type ScmState struct {
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type TenantAccessState struct {
	Enabled      bool   `json:"enabled"`
	PrimaryKey   string `json:"primary_key"`
	SecondaryKey string `json:"secondary_key"`
	TenantId     string `json:"tenant_id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
