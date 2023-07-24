// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataactivedirectorydomainservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Notifications struct{}

type ReplicaSets struct{}

type SecureLdap struct{}

type Security struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
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

func (n NotificationsAttributes) AdditionalRecipients() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("additional_recipients"))
}

func (n NotificationsAttributes) NotifyDcAdmins() terra.BoolValue {
	return terra.ReferenceAsBool(n.ref.Append("notify_dc_admins"))
}

func (n NotificationsAttributes) NotifyGlobalAdmins() terra.BoolValue {
	return terra.ReferenceAsBool(n.ref.Append("notify_global_admins"))
}

type ReplicaSetsAttributes struct {
	ref terra.Reference
}

func (rs ReplicaSetsAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs ReplicaSetsAttributes) InternalWithRef(ref terra.Reference) ReplicaSetsAttributes {
	return ReplicaSetsAttributes{ref: ref}
}

func (rs ReplicaSetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs ReplicaSetsAttributes) DomainControllerIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rs.ref.Append("domain_controller_ip_addresses"))
}

func (rs ReplicaSetsAttributes) ExternalAccessIpAddress() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("external_access_ip_address"))
}

func (rs ReplicaSetsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("id"))
}

func (rs ReplicaSetsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("location"))
}

func (rs ReplicaSetsAttributes) ServiceStatus() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("service_status"))
}

func (rs ReplicaSetsAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("subnet_id"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type NotificationsState struct {
	AdditionalRecipients []string `json:"additional_recipients"`
	NotifyDcAdmins       bool     `json:"notify_dc_admins"`
	NotifyGlobalAdmins   bool     `json:"notify_global_admins"`
}

type ReplicaSetsState struct {
	DomainControllerIpAddresses []string `json:"domain_controller_ip_addresses"`
	ExternalAccessIpAddress     string   `json:"external_access_ip_address"`
	Id                          string   `json:"id"`
	Location                    string   `json:"location"`
	ServiceStatus               string   `json:"service_status"`
	SubnetId                    string   `json:"subnet_id"`
}

type SecureLdapState struct {
	CertificateExpiry     string `json:"certificate_expiry"`
	CertificateThumbprint string `json:"certificate_thumbprint"`
	Enabled               bool   `json:"enabled"`
	ExternalAccessEnabled bool   `json:"external_access_enabled"`
	PublicCertificate     string `json:"public_certificate"`
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
	Read string `json:"read"`
}
