// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package activedirectorydomainservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type InitialReplicaSet struct {
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
}

type Notifications struct {
	// AdditionalRecipients: set of string, optional
	AdditionalRecipients terra.SetValue[terra.StringValue] `hcl:"additional_recipients,attr"`
	// NotifyDcAdmins: bool, optional
	NotifyDcAdmins terra.BoolValue `hcl:"notify_dc_admins,attr"`
	// NotifyGlobalAdmins: bool, optional
	NotifyGlobalAdmins terra.BoolValue `hcl:"notify_global_admins,attr"`
}

type SecureLdap struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// ExternalAccessEnabled: bool, optional
	ExternalAccessEnabled terra.BoolValue `hcl:"external_access_enabled,attr"`
	// PfxCertificate: string, required
	PfxCertificate terra.StringValue `hcl:"pfx_certificate,attr" validate:"required"`
	// PfxCertificatePassword: string, required
	PfxCertificatePassword terra.StringValue `hcl:"pfx_certificate_password,attr" validate:"required"`
}

type Security struct {
	// KerberosArmoringEnabled: bool, optional
	KerberosArmoringEnabled terra.BoolValue `hcl:"kerberos_armoring_enabled,attr"`
	// KerberosRc4EncryptionEnabled: bool, optional
	KerberosRc4EncryptionEnabled terra.BoolValue `hcl:"kerberos_rc4_encryption_enabled,attr"`
	// NtlmV1Enabled: bool, optional
	NtlmV1Enabled terra.BoolValue `hcl:"ntlm_v1_enabled,attr"`
	// SyncKerberosPasswords: bool, optional
	SyncKerberosPasswords terra.BoolValue `hcl:"sync_kerberos_passwords,attr"`
	// SyncNtlmPasswords: bool, optional
	SyncNtlmPasswords terra.BoolValue `hcl:"sync_ntlm_passwords,attr"`
	// SyncOnPremPasswords: bool, optional
	SyncOnPremPasswords terra.BoolValue `hcl:"sync_on_prem_passwords,attr"`
	// TlsV1Enabled: bool, optional
	TlsV1Enabled terra.BoolValue `hcl:"tls_v1_enabled,attr"`
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

type InitialReplicaSetAttributes struct {
	ref terra.Reference
}

func (irs InitialReplicaSetAttributes) InternalRef() (terra.Reference, error) {
	return irs.ref, nil
}

func (irs InitialReplicaSetAttributes) InternalWithRef(ref terra.Reference) InitialReplicaSetAttributes {
	return InitialReplicaSetAttributes{ref: ref}
}

func (irs InitialReplicaSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return irs.ref.InternalTokens()
}

func (irs InitialReplicaSetAttributes) DomainControllerIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](irs.ref.Append("domain_controller_ip_addresses"))
}

func (irs InitialReplicaSetAttributes) ExternalAccessIpAddress() terra.StringValue {
	return terra.ReferenceAsString(irs.ref.Append("external_access_ip_address"))
}

func (irs InitialReplicaSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(irs.ref.Append("id"))
}

func (irs InitialReplicaSetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(irs.ref.Append("location"))
}

func (irs InitialReplicaSetAttributes) ServiceStatus() terra.StringValue {
	return terra.ReferenceAsString(irs.ref.Append("service_status"))
}

func (irs InitialReplicaSetAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(irs.ref.Append("subnet_id"))
}

type NotificationsAttributes struct {
	ref terra.Reference
}

func (n NotificationsAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NotificationsAttributes) InternalWithRef(ref terra.Reference) NotificationsAttributes {
	return NotificationsAttributes{ref: ref}
}

func (n NotificationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NotificationsAttributes) AdditionalRecipients() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](n.ref.Append("additional_recipients"))
}

func (n NotificationsAttributes) NotifyDcAdmins() terra.BoolValue {
	return terra.ReferenceAsBool(n.ref.Append("notify_dc_admins"))
}

func (n NotificationsAttributes) NotifyGlobalAdmins() terra.BoolValue {
	return terra.ReferenceAsBool(n.ref.Append("notify_global_admins"))
}

type SecureLdapAttributes struct {
	ref terra.Reference
}

func (sl SecureLdapAttributes) InternalRef() (terra.Reference, error) {
	return sl.ref, nil
}

func (sl SecureLdapAttributes) InternalWithRef(ref terra.Reference) SecureLdapAttributes {
	return SecureLdapAttributes{ref: ref}
}

func (sl SecureLdapAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sl.ref.InternalTokens()
}

func (sl SecureLdapAttributes) CertificateExpiry() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("certificate_expiry"))
}

func (sl SecureLdapAttributes) CertificateThumbprint() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("certificate_thumbprint"))
}

func (sl SecureLdapAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sl.ref.Append("enabled"))
}

func (sl SecureLdapAttributes) ExternalAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sl.ref.Append("external_access_enabled"))
}

func (sl SecureLdapAttributes) PfxCertificate() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("pfx_certificate"))
}

func (sl SecureLdapAttributes) PfxCertificatePassword() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("pfx_certificate_password"))
}

func (sl SecureLdapAttributes) PublicCertificate() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("public_certificate"))
}

type SecurityAttributes struct {
	ref terra.Reference
}

func (s SecurityAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SecurityAttributes) InternalWithRef(ref terra.Reference) SecurityAttributes {
	return SecurityAttributes{ref: ref}
}

func (s SecurityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SecurityAttributes) KerberosArmoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("kerberos_armoring_enabled"))
}

func (s SecurityAttributes) KerberosRc4EncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("kerberos_rc4_encryption_enabled"))
}

func (s SecurityAttributes) NtlmV1Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("ntlm_v1_enabled"))
}

func (s SecurityAttributes) SyncKerberosPasswords() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("sync_kerberos_passwords"))
}

func (s SecurityAttributes) SyncNtlmPasswords() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("sync_ntlm_passwords"))
}

func (s SecurityAttributes) SyncOnPremPasswords() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("sync_on_prem_passwords"))
}

func (s SecurityAttributes) TlsV1Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("tls_v1_enabled"))
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

type InitialReplicaSetState struct {
	DomainControllerIpAddresses []string `json:"domain_controller_ip_addresses"`
	ExternalAccessIpAddress     string   `json:"external_access_ip_address"`
	Id                          string   `json:"id"`
	Location                    string   `json:"location"`
	ServiceStatus               string   `json:"service_status"`
	SubnetId                    string   `json:"subnet_id"`
}

type NotificationsState struct {
	AdditionalRecipients []string `json:"additional_recipients"`
	NotifyDcAdmins       bool     `json:"notify_dc_admins"`
	NotifyGlobalAdmins   bool     `json:"notify_global_admins"`
}

type SecureLdapState struct {
	CertificateExpiry      string `json:"certificate_expiry"`
	CertificateThumbprint  string `json:"certificate_thumbprint"`
	Enabled                bool   `json:"enabled"`
	ExternalAccessEnabled  bool   `json:"external_access_enabled"`
	PfxCertificate         string `json:"pfx_certificate"`
	PfxCertificatePassword string `json:"pfx_certificate_password"`
	PublicCertificate      string `json:"public_certificate"`
}

type SecurityState struct {
	KerberosArmoringEnabled      bool `json:"kerberos_armoring_enabled"`
	KerberosRc4EncryptionEnabled bool `json:"kerberos_rc4_encryption_enabled"`
	NtlmV1Enabled                bool `json:"ntlm_v1_enabled"`
	SyncKerberosPasswords        bool `json:"sync_kerberos_passwords"`
	SyncNtlmPasswords            bool `json:"sync_ntlm_passwords"`
	SyncOnPremPasswords          bool `json:"sync_on_prem_passwords"`
	TlsV1Enabled                 bool `json:"tls_v1_enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
