// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apimanagementcustomdomain

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DeveloperPortal struct {
	// Certificate: string, optional
	Certificate terra.StringValue `hcl:"certificate,attr"`
	// CertificatePassword: string, optional
	CertificatePassword terra.StringValue `hcl:"certificate_password,attr"`
	// HostName: string, required
	HostName terra.StringValue `hcl:"host_name,attr" validate:"required"`
	// KeyVaultId: string, optional
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr"`
	// NegotiateClientCertificate: bool, optional
	NegotiateClientCertificate terra.BoolValue `hcl:"negotiate_client_certificate,attr"`
	// SslKeyvaultIdentityClientId: string, optional
	SslKeyvaultIdentityClientId terra.StringValue `hcl:"ssl_keyvault_identity_client_id,attr"`
}

type Gateway struct {
	// Certificate: string, optional
	Certificate terra.StringValue `hcl:"certificate,attr"`
	// CertificatePassword: string, optional
	CertificatePassword terra.StringValue `hcl:"certificate_password,attr"`
	// DefaultSslBinding: bool, optional
	DefaultSslBinding terra.BoolValue `hcl:"default_ssl_binding,attr"`
	// HostName: string, required
	HostName terra.StringValue `hcl:"host_name,attr" validate:"required"`
	// KeyVaultId: string, optional
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr"`
	// NegotiateClientCertificate: bool, optional
	NegotiateClientCertificate terra.BoolValue `hcl:"negotiate_client_certificate,attr"`
	// SslKeyvaultIdentityClientId: string, optional
	SslKeyvaultIdentityClientId terra.StringValue `hcl:"ssl_keyvault_identity_client_id,attr"`
}

type Management struct {
	// Certificate: string, optional
	Certificate terra.StringValue `hcl:"certificate,attr"`
	// CertificatePassword: string, optional
	CertificatePassword terra.StringValue `hcl:"certificate_password,attr"`
	// HostName: string, required
	HostName terra.StringValue `hcl:"host_name,attr" validate:"required"`
	// KeyVaultId: string, optional
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr"`
	// NegotiateClientCertificate: bool, optional
	NegotiateClientCertificate terra.BoolValue `hcl:"negotiate_client_certificate,attr"`
	// SslKeyvaultIdentityClientId: string, optional
	SslKeyvaultIdentityClientId terra.StringValue `hcl:"ssl_keyvault_identity_client_id,attr"`
}

type Portal struct {
	// Certificate: string, optional
	Certificate terra.StringValue `hcl:"certificate,attr"`
	// CertificatePassword: string, optional
	CertificatePassword terra.StringValue `hcl:"certificate_password,attr"`
	// HostName: string, required
	HostName terra.StringValue `hcl:"host_name,attr" validate:"required"`
	// KeyVaultId: string, optional
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr"`
	// NegotiateClientCertificate: bool, optional
	NegotiateClientCertificate terra.BoolValue `hcl:"negotiate_client_certificate,attr"`
	// SslKeyvaultIdentityClientId: string, optional
	SslKeyvaultIdentityClientId terra.StringValue `hcl:"ssl_keyvault_identity_client_id,attr"`
}

type Scm struct {
	// Certificate: string, optional
	Certificate terra.StringValue `hcl:"certificate,attr"`
	// CertificatePassword: string, optional
	CertificatePassword terra.StringValue `hcl:"certificate_password,attr"`
	// HostName: string, required
	HostName terra.StringValue `hcl:"host_name,attr" validate:"required"`
	// KeyVaultId: string, optional
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr"`
	// NegotiateClientCertificate: bool, optional
	NegotiateClientCertificate terra.BoolValue `hcl:"negotiate_client_certificate,attr"`
	// SslKeyvaultIdentityClientId: string, optional
	SslKeyvaultIdentityClientId terra.StringValue `hcl:"ssl_keyvault_identity_client_id,attr"`
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

func (dp DeveloperPortalAttributes) Certificate() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("certificate"))
}

func (dp DeveloperPortalAttributes) CertificatePassword() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("certificate_password"))
}

func (dp DeveloperPortalAttributes) CertificateSource() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("certificate_source"))
}

func (dp DeveloperPortalAttributes) CertificateStatus() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("certificate_status"))
}

func (dp DeveloperPortalAttributes) Expiry() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("expiry"))
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

func (dp DeveloperPortalAttributes) SslKeyvaultIdentityClientId() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("ssl_keyvault_identity_client_id"))
}

func (dp DeveloperPortalAttributes) Subject() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("subject"))
}

func (dp DeveloperPortalAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("thumbprint"))
}

type GatewayAttributes struct {
	ref terra.Reference
}

func (g GatewayAttributes) InternalRef() terra.Reference {
	return g.ref
}

func (g GatewayAttributes) InternalWithRef(ref terra.Reference) GatewayAttributes {
	return GatewayAttributes{ref: ref}
}

func (g GatewayAttributes) InternalTokens() hclwrite.Tokens {
	return g.ref.InternalTokens()
}

func (g GatewayAttributes) Certificate() terra.StringValue {
	return terra.ReferenceString(g.ref.Append("certificate"))
}

func (g GatewayAttributes) CertificatePassword() terra.StringValue {
	return terra.ReferenceString(g.ref.Append("certificate_password"))
}

func (g GatewayAttributes) CertificateSource() terra.StringValue {
	return terra.ReferenceString(g.ref.Append("certificate_source"))
}

func (g GatewayAttributes) CertificateStatus() terra.StringValue {
	return terra.ReferenceString(g.ref.Append("certificate_status"))
}

func (g GatewayAttributes) DefaultSslBinding() terra.BoolValue {
	return terra.ReferenceBool(g.ref.Append("default_ssl_binding"))
}

func (g GatewayAttributes) Expiry() terra.StringValue {
	return terra.ReferenceString(g.ref.Append("expiry"))
}

func (g GatewayAttributes) HostName() terra.StringValue {
	return terra.ReferenceString(g.ref.Append("host_name"))
}

func (g GatewayAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(g.ref.Append("key_vault_id"))
}

func (g GatewayAttributes) NegotiateClientCertificate() terra.BoolValue {
	return terra.ReferenceBool(g.ref.Append("negotiate_client_certificate"))
}

func (g GatewayAttributes) SslKeyvaultIdentityClientId() terra.StringValue {
	return terra.ReferenceString(g.ref.Append("ssl_keyvault_identity_client_id"))
}

func (g GatewayAttributes) Subject() terra.StringValue {
	return terra.ReferenceString(g.ref.Append("subject"))
}

func (g GatewayAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceString(g.ref.Append("thumbprint"))
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

func (m ManagementAttributes) Certificate() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("certificate"))
}

func (m ManagementAttributes) CertificatePassword() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("certificate_password"))
}

func (m ManagementAttributes) CertificateSource() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("certificate_source"))
}

func (m ManagementAttributes) CertificateStatus() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("certificate_status"))
}

func (m ManagementAttributes) Expiry() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("expiry"))
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

func (m ManagementAttributes) SslKeyvaultIdentityClientId() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("ssl_keyvault_identity_client_id"))
}

func (m ManagementAttributes) Subject() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("subject"))
}

func (m ManagementAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("thumbprint"))
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

func (p PortalAttributes) Certificate() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("certificate"))
}

func (p PortalAttributes) CertificatePassword() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("certificate_password"))
}

func (p PortalAttributes) CertificateSource() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("certificate_source"))
}

func (p PortalAttributes) CertificateStatus() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("certificate_status"))
}

func (p PortalAttributes) Expiry() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("expiry"))
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

func (p PortalAttributes) SslKeyvaultIdentityClientId() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("ssl_keyvault_identity_client_id"))
}

func (p PortalAttributes) Subject() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("subject"))
}

func (p PortalAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("thumbprint"))
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

func (s ScmAttributes) Certificate() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("certificate"))
}

func (s ScmAttributes) CertificatePassword() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("certificate_password"))
}

func (s ScmAttributes) CertificateSource() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("certificate_source"))
}

func (s ScmAttributes) CertificateStatus() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("certificate_status"))
}

func (s ScmAttributes) Expiry() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("expiry"))
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

func (s ScmAttributes) SslKeyvaultIdentityClientId() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("ssl_keyvault_identity_client_id"))
}

func (s ScmAttributes) Subject() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("subject"))
}

func (s ScmAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("thumbprint"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type DeveloperPortalState struct {
	Certificate                 string `json:"certificate"`
	CertificatePassword         string `json:"certificate_password"`
	CertificateSource           string `json:"certificate_source"`
	CertificateStatus           string `json:"certificate_status"`
	Expiry                      string `json:"expiry"`
	HostName                    string `json:"host_name"`
	KeyVaultId                  string `json:"key_vault_id"`
	NegotiateClientCertificate  bool   `json:"negotiate_client_certificate"`
	SslKeyvaultIdentityClientId string `json:"ssl_keyvault_identity_client_id"`
	Subject                     string `json:"subject"`
	Thumbprint                  string `json:"thumbprint"`
}

type GatewayState struct {
	Certificate                 string `json:"certificate"`
	CertificatePassword         string `json:"certificate_password"`
	CertificateSource           string `json:"certificate_source"`
	CertificateStatus           string `json:"certificate_status"`
	DefaultSslBinding           bool   `json:"default_ssl_binding"`
	Expiry                      string `json:"expiry"`
	HostName                    string `json:"host_name"`
	KeyVaultId                  string `json:"key_vault_id"`
	NegotiateClientCertificate  bool   `json:"negotiate_client_certificate"`
	SslKeyvaultIdentityClientId string `json:"ssl_keyvault_identity_client_id"`
	Subject                     string `json:"subject"`
	Thumbprint                  string `json:"thumbprint"`
}

type ManagementState struct {
	Certificate                 string `json:"certificate"`
	CertificatePassword         string `json:"certificate_password"`
	CertificateSource           string `json:"certificate_source"`
	CertificateStatus           string `json:"certificate_status"`
	Expiry                      string `json:"expiry"`
	HostName                    string `json:"host_name"`
	KeyVaultId                  string `json:"key_vault_id"`
	NegotiateClientCertificate  bool   `json:"negotiate_client_certificate"`
	SslKeyvaultIdentityClientId string `json:"ssl_keyvault_identity_client_id"`
	Subject                     string `json:"subject"`
	Thumbprint                  string `json:"thumbprint"`
}

type PortalState struct {
	Certificate                 string `json:"certificate"`
	CertificatePassword         string `json:"certificate_password"`
	CertificateSource           string `json:"certificate_source"`
	CertificateStatus           string `json:"certificate_status"`
	Expiry                      string `json:"expiry"`
	HostName                    string `json:"host_name"`
	KeyVaultId                  string `json:"key_vault_id"`
	NegotiateClientCertificate  bool   `json:"negotiate_client_certificate"`
	SslKeyvaultIdentityClientId string `json:"ssl_keyvault_identity_client_id"`
	Subject                     string `json:"subject"`
	Thumbprint                  string `json:"thumbprint"`
}

type ScmState struct {
	Certificate                 string `json:"certificate"`
	CertificatePassword         string `json:"certificate_password"`
	CertificateSource           string `json:"certificate_source"`
	CertificateStatus           string `json:"certificate_status"`
	Expiry                      string `json:"expiry"`
	HostName                    string `json:"host_name"`
	KeyVaultId                  string `json:"key_vault_id"`
	NegotiateClientCertificate  bool   `json:"negotiate_client_certificate"`
	SslKeyvaultIdentityClientId string `json:"ssl_keyvault_identity_client_id"`
	Subject                     string `json:"subject"`
	Thumbprint                  string `json:"thumbprint"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
