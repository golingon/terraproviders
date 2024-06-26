// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_api_management

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataAdditionalLocationAttributes struct {
	ref terra.Reference
}

func (al DataAdditionalLocationAttributes) InternalRef() (terra.Reference, error) {
	return al.ref, nil
}

func (al DataAdditionalLocationAttributes) InternalWithRef(ref terra.Reference) DataAdditionalLocationAttributes {
	return DataAdditionalLocationAttributes{ref: ref}
}

func (al DataAdditionalLocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return al.ref.InternalTokens()
}

func (al DataAdditionalLocationAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(al.ref.Append("capacity"))
}

func (al DataAdditionalLocationAttributes) GatewayRegionalUrl() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("gateway_regional_url"))
}

func (al DataAdditionalLocationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("location"))
}

func (al DataAdditionalLocationAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](al.ref.Append("private_ip_addresses"))
}

func (al DataAdditionalLocationAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("public_ip_address_id"))
}

func (al DataAdditionalLocationAttributes) PublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](al.ref.Append("public_ip_addresses"))
}

func (al DataAdditionalLocationAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](al.ref.Append("zones"))
}

type DataHostnameConfigurationAttributes struct {
	ref terra.Reference
}

func (hc DataHostnameConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return hc.ref, nil
}

func (hc DataHostnameConfigurationAttributes) InternalWithRef(ref terra.Reference) DataHostnameConfigurationAttributes {
	return DataHostnameConfigurationAttributes{ref: ref}
}

func (hc DataHostnameConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hc.ref.InternalTokens()
}

func (hc DataHostnameConfigurationAttributes) DeveloperPortal() terra.ListValue[DataHostnameConfigurationDeveloperPortalAttributes] {
	return terra.ReferenceAsList[DataHostnameConfigurationDeveloperPortalAttributes](hc.ref.Append("developer_portal"))
}

func (hc DataHostnameConfigurationAttributes) Management() terra.ListValue[DataHostnameConfigurationManagementAttributes] {
	return terra.ReferenceAsList[DataHostnameConfigurationManagementAttributes](hc.ref.Append("management"))
}

func (hc DataHostnameConfigurationAttributes) Portal() terra.ListValue[DataHostnameConfigurationPortalAttributes] {
	return terra.ReferenceAsList[DataHostnameConfigurationPortalAttributes](hc.ref.Append("portal"))
}

func (hc DataHostnameConfigurationAttributes) Proxy() terra.ListValue[DataHostnameConfigurationProxyAttributes] {
	return terra.ReferenceAsList[DataHostnameConfigurationProxyAttributes](hc.ref.Append("proxy"))
}

func (hc DataHostnameConfigurationAttributes) Scm() terra.ListValue[DataHostnameConfigurationScmAttributes] {
	return terra.ReferenceAsList[DataHostnameConfigurationScmAttributes](hc.ref.Append("scm"))
}

type DataHostnameConfigurationDeveloperPortalAttributes struct {
	ref terra.Reference
}

func (dp DataHostnameConfigurationDeveloperPortalAttributes) InternalRef() (terra.Reference, error) {
	return dp.ref, nil
}

func (dp DataHostnameConfigurationDeveloperPortalAttributes) InternalWithRef(ref terra.Reference) DataHostnameConfigurationDeveloperPortalAttributes {
	return DataHostnameConfigurationDeveloperPortalAttributes{ref: ref}
}

func (dp DataHostnameConfigurationDeveloperPortalAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dp.ref.InternalTokens()
}

func (dp DataHostnameConfigurationDeveloperPortalAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("host_name"))
}

func (dp DataHostnameConfigurationDeveloperPortalAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("key_vault_id"))
}

func (dp DataHostnameConfigurationDeveloperPortalAttributes) NegotiateClientCertificate() terra.BoolValue {
	return terra.ReferenceAsBool(dp.ref.Append("negotiate_client_certificate"))
}

type DataHostnameConfigurationManagementAttributes struct {
	ref terra.Reference
}

func (m DataHostnameConfigurationManagementAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m DataHostnameConfigurationManagementAttributes) InternalWithRef(ref terra.Reference) DataHostnameConfigurationManagementAttributes {
	return DataHostnameConfigurationManagementAttributes{ref: ref}
}

func (m DataHostnameConfigurationManagementAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m DataHostnameConfigurationManagementAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("host_name"))
}

func (m DataHostnameConfigurationManagementAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("key_vault_id"))
}

func (m DataHostnameConfigurationManagementAttributes) NegotiateClientCertificate() terra.BoolValue {
	return terra.ReferenceAsBool(m.ref.Append("negotiate_client_certificate"))
}

type DataHostnameConfigurationPortalAttributes struct {
	ref terra.Reference
}

func (p DataHostnameConfigurationPortalAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p DataHostnameConfigurationPortalAttributes) InternalWithRef(ref terra.Reference) DataHostnameConfigurationPortalAttributes {
	return DataHostnameConfigurationPortalAttributes{ref: ref}
}

func (p DataHostnameConfigurationPortalAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p DataHostnameConfigurationPortalAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("host_name"))
}

func (p DataHostnameConfigurationPortalAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("key_vault_id"))
}

func (p DataHostnameConfigurationPortalAttributes) NegotiateClientCertificate() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("negotiate_client_certificate"))
}

type DataHostnameConfigurationProxyAttributes struct {
	ref terra.Reference
}

func (p DataHostnameConfigurationProxyAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p DataHostnameConfigurationProxyAttributes) InternalWithRef(ref terra.Reference) DataHostnameConfigurationProxyAttributes {
	return DataHostnameConfigurationProxyAttributes{ref: ref}
}

func (p DataHostnameConfigurationProxyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p DataHostnameConfigurationProxyAttributes) DefaultSslBinding() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("default_ssl_binding"))
}

func (p DataHostnameConfigurationProxyAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("host_name"))
}

func (p DataHostnameConfigurationProxyAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("key_vault_id"))
}

func (p DataHostnameConfigurationProxyAttributes) NegotiateClientCertificate() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("negotiate_client_certificate"))
}

type DataHostnameConfigurationScmAttributes struct {
	ref terra.Reference
}

func (s DataHostnameConfigurationScmAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s DataHostnameConfigurationScmAttributes) InternalWithRef(ref terra.Reference) DataHostnameConfigurationScmAttributes {
	return DataHostnameConfigurationScmAttributes{ref: ref}
}

func (s DataHostnameConfigurationScmAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s DataHostnameConfigurationScmAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("host_name"))
}

func (s DataHostnameConfigurationScmAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("key_vault_id"))
}

func (s DataHostnameConfigurationScmAttributes) NegotiateClientCertificate() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("negotiate_client_certificate"))
}

type DataIdentityAttributes struct {
	ref terra.Reference
}

func (i DataIdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i DataIdentityAttributes) InternalWithRef(ref terra.Reference) DataIdentityAttributes {
	return DataIdentityAttributes{ref: ref}
}

func (i DataIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i DataIdentityAttributes) IdentityIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i DataIdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i DataIdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i DataIdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type DataTenantAccessAttributes struct {
	ref terra.Reference
}

func (ta DataTenantAccessAttributes) InternalRef() (terra.Reference, error) {
	return ta.ref, nil
}

func (ta DataTenantAccessAttributes) InternalWithRef(ref terra.Reference) DataTenantAccessAttributes {
	return DataTenantAccessAttributes{ref: ref}
}

func (ta DataTenantAccessAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ta.ref.InternalTokens()
}

func (ta DataTenantAccessAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ta.ref.Append("enabled"))
}

func (ta DataTenantAccessAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("primary_key"))
}

func (ta DataTenantAccessAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("secondary_key"))
}

func (ta DataTenantAccessAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("tenant_id"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataAdditionalLocationState struct {
	Capacity           float64  `json:"capacity"`
	GatewayRegionalUrl string   `json:"gateway_regional_url"`
	Location           string   `json:"location"`
	PrivateIpAddresses []string `json:"private_ip_addresses"`
	PublicIpAddressId  string   `json:"public_ip_address_id"`
	PublicIpAddresses  []string `json:"public_ip_addresses"`
	Zones              []string `json:"zones"`
}

type DataHostnameConfigurationState struct {
	DeveloperPortal []DataHostnameConfigurationDeveloperPortalState `json:"developer_portal"`
	Management      []DataHostnameConfigurationManagementState      `json:"management"`
	Portal          []DataHostnameConfigurationPortalState          `json:"portal"`
	Proxy           []DataHostnameConfigurationProxyState           `json:"proxy"`
	Scm             []DataHostnameConfigurationScmState             `json:"scm"`
}

type DataHostnameConfigurationDeveloperPortalState struct {
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type DataHostnameConfigurationManagementState struct {
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type DataHostnameConfigurationPortalState struct {
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type DataHostnameConfigurationProxyState struct {
	DefaultSslBinding          bool   `json:"default_ssl_binding"`
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type DataHostnameConfigurationScmState struct {
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type DataIdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type DataTenantAccessState struct {
	Enabled      bool   `json:"enabled"`
	PrimaryKey   string `json:"primary_key"`
	SecondaryKey string `json:"secondary_key"`
	TenantId     string `json:"tenant_id"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
