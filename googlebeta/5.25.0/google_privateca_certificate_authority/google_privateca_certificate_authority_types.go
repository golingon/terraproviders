// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_privateca_certificate_authority

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Config struct {
	// ConfigSubjectConfig: required
	SubjectConfig *ConfigSubjectConfig `hcl:"subject_config,block" validate:"required"`
	// ConfigX509Config: required
	X509Config *ConfigX509Config `hcl:"x509_config,block" validate:"required"`
}

type ConfigSubjectConfig struct {
	// ConfigSubjectConfigSubject: required
	Subject *ConfigSubjectConfigSubject `hcl:"subject,block" validate:"required"`
	// ConfigSubjectConfigSubjectAltName: optional
	SubjectAltName *ConfigSubjectConfigSubjectAltName `hcl:"subject_alt_name,block"`
}

type ConfigSubjectConfigSubject struct {
	// CommonName: string, required
	CommonName terra.StringValue `hcl:"common_name,attr" validate:"required"`
	// CountryCode: string, optional
	CountryCode terra.StringValue `hcl:"country_code,attr"`
	// Locality: string, optional
	Locality terra.StringValue `hcl:"locality,attr"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// OrganizationalUnit: string, optional
	OrganizationalUnit terra.StringValue `hcl:"organizational_unit,attr"`
	// PostalCode: string, optional
	PostalCode terra.StringValue `hcl:"postal_code,attr"`
	// Province: string, optional
	Province terra.StringValue `hcl:"province,attr"`
	// StreetAddress: string, optional
	StreetAddress terra.StringValue `hcl:"street_address,attr"`
}

type ConfigSubjectConfigSubjectAltName struct {
	// DnsNames: list of string, optional
	DnsNames terra.ListValue[terra.StringValue] `hcl:"dns_names,attr"`
	// EmailAddresses: list of string, optional
	EmailAddresses terra.ListValue[terra.StringValue] `hcl:"email_addresses,attr"`
	// IpAddresses: list of string, optional
	IpAddresses terra.ListValue[terra.StringValue] `hcl:"ip_addresses,attr"`
	// Uris: list of string, optional
	Uris terra.ListValue[terra.StringValue] `hcl:"uris,attr"`
}

type ConfigX509Config struct {
	// AiaOcspServers: list of string, optional
	AiaOcspServers terra.ListValue[terra.StringValue] `hcl:"aia_ocsp_servers,attr"`
	// ConfigX509ConfigAdditionalExtensions: min=0
	AdditionalExtensions []ConfigX509ConfigAdditionalExtensions `hcl:"additional_extensions,block" validate:"min=0"`
	// ConfigX509ConfigCaOptions: required
	CaOptions *ConfigX509ConfigCaOptions `hcl:"ca_options,block" validate:"required"`
	// ConfigX509ConfigKeyUsage: required
	KeyUsage *ConfigX509ConfigKeyUsage `hcl:"key_usage,block" validate:"required"`
	// ConfigX509ConfigNameConstraints: optional
	NameConstraints *ConfigX509ConfigNameConstraints `hcl:"name_constraints,block"`
	// ConfigX509ConfigPolicyIds: min=0
	PolicyIds []ConfigX509ConfigPolicyIds `hcl:"policy_ids,block" validate:"min=0"`
}

type ConfigX509ConfigAdditionalExtensions struct {
	// Critical: bool, required
	Critical terra.BoolValue `hcl:"critical,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// ConfigX509ConfigAdditionalExtensionsObjectId: required
	ObjectId *ConfigX509ConfigAdditionalExtensionsObjectId `hcl:"object_id,block" validate:"required"`
}

type ConfigX509ConfigAdditionalExtensionsObjectId struct {
	// ObjectIdPath: list of number, required
	ObjectIdPath terra.ListValue[terra.NumberValue] `hcl:"object_id_path,attr" validate:"required"`
}

type ConfigX509ConfigCaOptions struct {
	// IsCa: bool, required
	IsCa terra.BoolValue `hcl:"is_ca,attr" validate:"required"`
	// MaxIssuerPathLength: number, optional
	MaxIssuerPathLength terra.NumberValue `hcl:"max_issuer_path_length,attr"`
	// NonCa: bool, optional
	NonCa terra.BoolValue `hcl:"non_ca,attr"`
	// ZeroMaxIssuerPathLength: bool, optional
	ZeroMaxIssuerPathLength terra.BoolValue `hcl:"zero_max_issuer_path_length,attr"`
}

type ConfigX509ConfigKeyUsage struct {
	// ConfigX509ConfigKeyUsageBaseKeyUsage: required
	BaseKeyUsage *ConfigX509ConfigKeyUsageBaseKeyUsage `hcl:"base_key_usage,block" validate:"required"`
	// ConfigX509ConfigKeyUsageExtendedKeyUsage: required
	ExtendedKeyUsage *ConfigX509ConfigKeyUsageExtendedKeyUsage `hcl:"extended_key_usage,block" validate:"required"`
	// ConfigX509ConfigKeyUsageUnknownExtendedKeyUsages: min=0
	UnknownExtendedKeyUsages []ConfigX509ConfigKeyUsageUnknownExtendedKeyUsages `hcl:"unknown_extended_key_usages,block" validate:"min=0"`
}

type ConfigX509ConfigKeyUsageBaseKeyUsage struct {
	// CertSign: bool, optional
	CertSign terra.BoolValue `hcl:"cert_sign,attr"`
	// ContentCommitment: bool, optional
	ContentCommitment terra.BoolValue `hcl:"content_commitment,attr"`
	// CrlSign: bool, optional
	CrlSign terra.BoolValue `hcl:"crl_sign,attr"`
	// DataEncipherment: bool, optional
	DataEncipherment terra.BoolValue `hcl:"data_encipherment,attr"`
	// DecipherOnly: bool, optional
	DecipherOnly terra.BoolValue `hcl:"decipher_only,attr"`
	// DigitalSignature: bool, optional
	DigitalSignature terra.BoolValue `hcl:"digital_signature,attr"`
	// EncipherOnly: bool, optional
	EncipherOnly terra.BoolValue `hcl:"encipher_only,attr"`
	// KeyAgreement: bool, optional
	KeyAgreement terra.BoolValue `hcl:"key_agreement,attr"`
	// KeyEncipherment: bool, optional
	KeyEncipherment terra.BoolValue `hcl:"key_encipherment,attr"`
}

type ConfigX509ConfigKeyUsageExtendedKeyUsage struct {
	// ClientAuth: bool, optional
	ClientAuth terra.BoolValue `hcl:"client_auth,attr"`
	// CodeSigning: bool, optional
	CodeSigning terra.BoolValue `hcl:"code_signing,attr"`
	// EmailProtection: bool, optional
	EmailProtection terra.BoolValue `hcl:"email_protection,attr"`
	// OcspSigning: bool, optional
	OcspSigning terra.BoolValue `hcl:"ocsp_signing,attr"`
	// ServerAuth: bool, optional
	ServerAuth terra.BoolValue `hcl:"server_auth,attr"`
	// TimeStamping: bool, optional
	TimeStamping terra.BoolValue `hcl:"time_stamping,attr"`
}

type ConfigX509ConfigKeyUsageUnknownExtendedKeyUsages struct {
	// ObjectIdPath: list of number, required
	ObjectIdPath terra.ListValue[terra.NumberValue] `hcl:"object_id_path,attr" validate:"required"`
}

type ConfigX509ConfigNameConstraints struct {
	// Critical: bool, required
	Critical terra.BoolValue `hcl:"critical,attr" validate:"required"`
	// ExcludedDnsNames: list of string, optional
	ExcludedDnsNames terra.ListValue[terra.StringValue] `hcl:"excluded_dns_names,attr"`
	// ExcludedEmailAddresses: list of string, optional
	ExcludedEmailAddresses terra.ListValue[terra.StringValue] `hcl:"excluded_email_addresses,attr"`
	// ExcludedIpRanges: list of string, optional
	ExcludedIpRanges terra.ListValue[terra.StringValue] `hcl:"excluded_ip_ranges,attr"`
	// ExcludedUris: list of string, optional
	ExcludedUris terra.ListValue[terra.StringValue] `hcl:"excluded_uris,attr"`
	// PermittedDnsNames: list of string, optional
	PermittedDnsNames terra.ListValue[terra.StringValue] `hcl:"permitted_dns_names,attr"`
	// PermittedEmailAddresses: list of string, optional
	PermittedEmailAddresses terra.ListValue[terra.StringValue] `hcl:"permitted_email_addresses,attr"`
	// PermittedIpRanges: list of string, optional
	PermittedIpRanges terra.ListValue[terra.StringValue] `hcl:"permitted_ip_ranges,attr"`
	// PermittedUris: list of string, optional
	PermittedUris terra.ListValue[terra.StringValue] `hcl:"permitted_uris,attr"`
}

type ConfigX509ConfigPolicyIds struct {
	// ObjectIdPath: list of number, required
	ObjectIdPath terra.ListValue[terra.NumberValue] `hcl:"object_id_path,attr" validate:"required"`
}

type KeySpec struct {
	// Algorithm: string, optional
	Algorithm terra.StringValue `hcl:"algorithm,attr"`
	// CloudKmsKeyVersion: string, optional
	CloudKmsKeyVersion terra.StringValue `hcl:"cloud_kms_key_version,attr"`
}

type SubordinateConfig struct {
	// CertificateAuthority: string, optional
	CertificateAuthority terra.StringValue `hcl:"certificate_authority,attr"`
	// SubordinateConfigPemIssuerChain: optional
	PemIssuerChain *SubordinateConfigPemIssuerChain `hcl:"pem_issuer_chain,block"`
}

type SubordinateConfigPemIssuerChain struct {
	// PemCertificates: list of string, optional
	PemCertificates terra.ListValue[terra.StringValue] `hcl:"pem_certificates,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AccessUrlsAttributes struct {
	ref terra.Reference
}

func (au AccessUrlsAttributes) InternalRef() (terra.Reference, error) {
	return au.ref, nil
}

func (au AccessUrlsAttributes) InternalWithRef(ref terra.Reference) AccessUrlsAttributes {
	return AccessUrlsAttributes{ref: ref}
}

func (au AccessUrlsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return au.ref.InternalTokens()
}

func (au AccessUrlsAttributes) CaCertificateAccessUrl() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("ca_certificate_access_url"))
}

func (au AccessUrlsAttributes) CrlAccessUrls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](au.ref.Append("crl_access_urls"))
}

type ConfigAttributes struct {
	ref terra.Reference
}

func (c ConfigAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigAttributes) InternalWithRef(ref terra.Reference) ConfigAttributes {
	return ConfigAttributes{ref: ref}
}

func (c ConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigAttributes) SubjectConfig() terra.ListValue[ConfigSubjectConfigAttributes] {
	return terra.ReferenceAsList[ConfigSubjectConfigAttributes](c.ref.Append("subject_config"))
}

func (c ConfigAttributes) X509Config() terra.ListValue[ConfigX509ConfigAttributes] {
	return terra.ReferenceAsList[ConfigX509ConfigAttributes](c.ref.Append("x509_config"))
}

type ConfigSubjectConfigAttributes struct {
	ref terra.Reference
}

func (sc ConfigSubjectConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc ConfigSubjectConfigAttributes) InternalWithRef(ref terra.Reference) ConfigSubjectConfigAttributes {
	return ConfigSubjectConfigAttributes{ref: ref}
}

func (sc ConfigSubjectConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc ConfigSubjectConfigAttributes) Subject() terra.ListValue[ConfigSubjectConfigSubjectAttributes] {
	return terra.ReferenceAsList[ConfigSubjectConfigSubjectAttributes](sc.ref.Append("subject"))
}

func (sc ConfigSubjectConfigAttributes) SubjectAltName() terra.ListValue[ConfigSubjectConfigSubjectAltNameAttributes] {
	return terra.ReferenceAsList[ConfigSubjectConfigSubjectAltNameAttributes](sc.ref.Append("subject_alt_name"))
}

type ConfigSubjectConfigSubjectAttributes struct {
	ref terra.Reference
}

func (s ConfigSubjectConfigSubjectAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ConfigSubjectConfigSubjectAttributes) InternalWithRef(ref terra.Reference) ConfigSubjectConfigSubjectAttributes {
	return ConfigSubjectConfigSubjectAttributes{ref: ref}
}

func (s ConfigSubjectConfigSubjectAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ConfigSubjectConfigSubjectAttributes) CommonName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("common_name"))
}

func (s ConfigSubjectConfigSubjectAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("country_code"))
}

func (s ConfigSubjectConfigSubjectAttributes) Locality() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("locality"))
}

func (s ConfigSubjectConfigSubjectAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("organization"))
}

func (s ConfigSubjectConfigSubjectAttributes) OrganizationalUnit() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("organizational_unit"))
}

func (s ConfigSubjectConfigSubjectAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("postal_code"))
}

func (s ConfigSubjectConfigSubjectAttributes) Province() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("province"))
}

func (s ConfigSubjectConfigSubjectAttributes) StreetAddress() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("street_address"))
}

type ConfigSubjectConfigSubjectAltNameAttributes struct {
	ref terra.Reference
}

func (san ConfigSubjectConfigSubjectAltNameAttributes) InternalRef() (terra.Reference, error) {
	return san.ref, nil
}

func (san ConfigSubjectConfigSubjectAltNameAttributes) InternalWithRef(ref terra.Reference) ConfigSubjectConfigSubjectAltNameAttributes {
	return ConfigSubjectConfigSubjectAltNameAttributes{ref: ref}
}

func (san ConfigSubjectConfigSubjectAltNameAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return san.ref.InternalTokens()
}

func (san ConfigSubjectConfigSubjectAltNameAttributes) DnsNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](san.ref.Append("dns_names"))
}

func (san ConfigSubjectConfigSubjectAltNameAttributes) EmailAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](san.ref.Append("email_addresses"))
}

func (san ConfigSubjectConfigSubjectAltNameAttributes) IpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](san.ref.Append("ip_addresses"))
}

func (san ConfigSubjectConfigSubjectAltNameAttributes) Uris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](san.ref.Append("uris"))
}

type ConfigX509ConfigAttributes struct {
	ref terra.Reference
}

func (xc ConfigX509ConfigAttributes) InternalRef() (terra.Reference, error) {
	return xc.ref, nil
}

func (xc ConfigX509ConfigAttributes) InternalWithRef(ref terra.Reference) ConfigX509ConfigAttributes {
	return ConfigX509ConfigAttributes{ref: ref}
}

func (xc ConfigX509ConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return xc.ref.InternalTokens()
}

func (xc ConfigX509ConfigAttributes) AiaOcspServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](xc.ref.Append("aia_ocsp_servers"))
}

func (xc ConfigX509ConfigAttributes) AdditionalExtensions() terra.ListValue[ConfigX509ConfigAdditionalExtensionsAttributes] {
	return terra.ReferenceAsList[ConfigX509ConfigAdditionalExtensionsAttributes](xc.ref.Append("additional_extensions"))
}

func (xc ConfigX509ConfigAttributes) CaOptions() terra.ListValue[ConfigX509ConfigCaOptionsAttributes] {
	return terra.ReferenceAsList[ConfigX509ConfigCaOptionsAttributes](xc.ref.Append("ca_options"))
}

func (xc ConfigX509ConfigAttributes) KeyUsage() terra.ListValue[ConfigX509ConfigKeyUsageAttributes] {
	return terra.ReferenceAsList[ConfigX509ConfigKeyUsageAttributes](xc.ref.Append("key_usage"))
}

func (xc ConfigX509ConfigAttributes) NameConstraints() terra.ListValue[ConfigX509ConfigNameConstraintsAttributes] {
	return terra.ReferenceAsList[ConfigX509ConfigNameConstraintsAttributes](xc.ref.Append("name_constraints"))
}

func (xc ConfigX509ConfigAttributes) PolicyIds() terra.ListValue[ConfigX509ConfigPolicyIdsAttributes] {
	return terra.ReferenceAsList[ConfigX509ConfigPolicyIdsAttributes](xc.ref.Append("policy_ids"))
}

type ConfigX509ConfigAdditionalExtensionsAttributes struct {
	ref terra.Reference
}

func (ae ConfigX509ConfigAdditionalExtensionsAttributes) InternalRef() (terra.Reference, error) {
	return ae.ref, nil
}

func (ae ConfigX509ConfigAdditionalExtensionsAttributes) InternalWithRef(ref terra.Reference) ConfigX509ConfigAdditionalExtensionsAttributes {
	return ConfigX509ConfigAdditionalExtensionsAttributes{ref: ref}
}

func (ae ConfigX509ConfigAdditionalExtensionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ae.ref.InternalTokens()
}

func (ae ConfigX509ConfigAdditionalExtensionsAttributes) Critical() terra.BoolValue {
	return terra.ReferenceAsBool(ae.ref.Append("critical"))
}

func (ae ConfigX509ConfigAdditionalExtensionsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("value"))
}

func (ae ConfigX509ConfigAdditionalExtensionsAttributes) ObjectId() terra.ListValue[ConfigX509ConfigAdditionalExtensionsObjectIdAttributes] {
	return terra.ReferenceAsList[ConfigX509ConfigAdditionalExtensionsObjectIdAttributes](ae.ref.Append("object_id"))
}

type ConfigX509ConfigAdditionalExtensionsObjectIdAttributes struct {
	ref terra.Reference
}

func (oi ConfigX509ConfigAdditionalExtensionsObjectIdAttributes) InternalRef() (terra.Reference, error) {
	return oi.ref, nil
}

func (oi ConfigX509ConfigAdditionalExtensionsObjectIdAttributes) InternalWithRef(ref terra.Reference) ConfigX509ConfigAdditionalExtensionsObjectIdAttributes {
	return ConfigX509ConfigAdditionalExtensionsObjectIdAttributes{ref: ref}
}

func (oi ConfigX509ConfigAdditionalExtensionsObjectIdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oi.ref.InternalTokens()
}

func (oi ConfigX509ConfigAdditionalExtensionsObjectIdAttributes) ObjectIdPath() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](oi.ref.Append("object_id_path"))
}

type ConfigX509ConfigCaOptionsAttributes struct {
	ref terra.Reference
}

func (co ConfigX509ConfigCaOptionsAttributes) InternalRef() (terra.Reference, error) {
	return co.ref, nil
}

func (co ConfigX509ConfigCaOptionsAttributes) InternalWithRef(ref terra.Reference) ConfigX509ConfigCaOptionsAttributes {
	return ConfigX509ConfigCaOptionsAttributes{ref: ref}
}

func (co ConfigX509ConfigCaOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return co.ref.InternalTokens()
}

func (co ConfigX509ConfigCaOptionsAttributes) IsCa() terra.BoolValue {
	return terra.ReferenceAsBool(co.ref.Append("is_ca"))
}

func (co ConfigX509ConfigCaOptionsAttributes) MaxIssuerPathLength() terra.NumberValue {
	return terra.ReferenceAsNumber(co.ref.Append("max_issuer_path_length"))
}

func (co ConfigX509ConfigCaOptionsAttributes) NonCa() terra.BoolValue {
	return terra.ReferenceAsBool(co.ref.Append("non_ca"))
}

func (co ConfigX509ConfigCaOptionsAttributes) ZeroMaxIssuerPathLength() terra.BoolValue {
	return terra.ReferenceAsBool(co.ref.Append("zero_max_issuer_path_length"))
}

type ConfigX509ConfigKeyUsageAttributes struct {
	ref terra.Reference
}

func (ku ConfigX509ConfigKeyUsageAttributes) InternalRef() (terra.Reference, error) {
	return ku.ref, nil
}

func (ku ConfigX509ConfigKeyUsageAttributes) InternalWithRef(ref terra.Reference) ConfigX509ConfigKeyUsageAttributes {
	return ConfigX509ConfigKeyUsageAttributes{ref: ref}
}

func (ku ConfigX509ConfigKeyUsageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ku.ref.InternalTokens()
}

func (ku ConfigX509ConfigKeyUsageAttributes) BaseKeyUsage() terra.ListValue[ConfigX509ConfigKeyUsageBaseKeyUsageAttributes] {
	return terra.ReferenceAsList[ConfigX509ConfigKeyUsageBaseKeyUsageAttributes](ku.ref.Append("base_key_usage"))
}

func (ku ConfigX509ConfigKeyUsageAttributes) ExtendedKeyUsage() terra.ListValue[ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes] {
	return terra.ReferenceAsList[ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes](ku.ref.Append("extended_key_usage"))
}

func (ku ConfigX509ConfigKeyUsageAttributes) UnknownExtendedKeyUsages() terra.ListValue[ConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes] {
	return terra.ReferenceAsList[ConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes](ku.ref.Append("unknown_extended_key_usages"))
}

type ConfigX509ConfigKeyUsageBaseKeyUsageAttributes struct {
	ref terra.Reference
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) InternalRef() (terra.Reference, error) {
	return bku.ref, nil
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) InternalWithRef(ref terra.Reference) ConfigX509ConfigKeyUsageBaseKeyUsageAttributes {
	return ConfigX509ConfigKeyUsageBaseKeyUsageAttributes{ref: ref}
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bku.ref.InternalTokens()
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) CertSign() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("cert_sign"))
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) ContentCommitment() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("content_commitment"))
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) CrlSign() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("crl_sign"))
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) DataEncipherment() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("data_encipherment"))
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) DecipherOnly() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("decipher_only"))
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) DigitalSignature() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("digital_signature"))
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) EncipherOnly() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("encipher_only"))
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) KeyAgreement() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("key_agreement"))
}

func (bku ConfigX509ConfigKeyUsageBaseKeyUsageAttributes) KeyEncipherment() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("key_encipherment"))
}

type ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes struct {
	ref terra.Reference
}

func (eku ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) InternalRef() (terra.Reference, error) {
	return eku.ref, nil
}

func (eku ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) InternalWithRef(ref terra.Reference) ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes {
	return ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes{ref: ref}
}

func (eku ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eku.ref.InternalTokens()
}

func (eku ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) ClientAuth() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("client_auth"))
}

func (eku ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) CodeSigning() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("code_signing"))
}

func (eku ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) EmailProtection() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("email_protection"))
}

func (eku ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) OcspSigning() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("ocsp_signing"))
}

func (eku ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) ServerAuth() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("server_auth"))
}

func (eku ConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) TimeStamping() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("time_stamping"))
}

type ConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes struct {
	ref terra.Reference
}

func (ueku ConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes) InternalRef() (terra.Reference, error) {
	return ueku.ref, nil
}

func (ueku ConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes) InternalWithRef(ref terra.Reference) ConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes {
	return ConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes{ref: ref}
}

func (ueku ConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ueku.ref.InternalTokens()
}

func (ueku ConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes) ObjectIdPath() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](ueku.ref.Append("object_id_path"))
}

type ConfigX509ConfigNameConstraintsAttributes struct {
	ref terra.Reference
}

func (nc ConfigX509ConfigNameConstraintsAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc ConfigX509ConfigNameConstraintsAttributes) InternalWithRef(ref terra.Reference) ConfigX509ConfigNameConstraintsAttributes {
	return ConfigX509ConfigNameConstraintsAttributes{ref: ref}
}

func (nc ConfigX509ConfigNameConstraintsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc ConfigX509ConfigNameConstraintsAttributes) Critical() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("critical"))
}

func (nc ConfigX509ConfigNameConstraintsAttributes) ExcludedDnsNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_dns_names"))
}

func (nc ConfigX509ConfigNameConstraintsAttributes) ExcludedEmailAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_email_addresses"))
}

func (nc ConfigX509ConfigNameConstraintsAttributes) ExcludedIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_ip_ranges"))
}

func (nc ConfigX509ConfigNameConstraintsAttributes) ExcludedUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_uris"))
}

func (nc ConfigX509ConfigNameConstraintsAttributes) PermittedDnsNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_dns_names"))
}

func (nc ConfigX509ConfigNameConstraintsAttributes) PermittedEmailAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_email_addresses"))
}

func (nc ConfigX509ConfigNameConstraintsAttributes) PermittedIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_ip_ranges"))
}

func (nc ConfigX509ConfigNameConstraintsAttributes) PermittedUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_uris"))
}

type ConfigX509ConfigPolicyIdsAttributes struct {
	ref terra.Reference
}

func (pi ConfigX509ConfigPolicyIdsAttributes) InternalRef() (terra.Reference, error) {
	return pi.ref, nil
}

func (pi ConfigX509ConfigPolicyIdsAttributes) InternalWithRef(ref terra.Reference) ConfigX509ConfigPolicyIdsAttributes {
	return ConfigX509ConfigPolicyIdsAttributes{ref: ref}
}

func (pi ConfigX509ConfigPolicyIdsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pi.ref.InternalTokens()
}

func (pi ConfigX509ConfigPolicyIdsAttributes) ObjectIdPath() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](pi.ref.Append("object_id_path"))
}

type KeySpecAttributes struct {
	ref terra.Reference
}

func (ks KeySpecAttributes) InternalRef() (terra.Reference, error) {
	return ks.ref, nil
}

func (ks KeySpecAttributes) InternalWithRef(ref terra.Reference) KeySpecAttributes {
	return KeySpecAttributes{ref: ref}
}

func (ks KeySpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ks.ref.InternalTokens()
}

func (ks KeySpecAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("algorithm"))
}

func (ks KeySpecAttributes) CloudKmsKeyVersion() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("cloud_kms_key_version"))
}

type SubordinateConfigAttributes struct {
	ref terra.Reference
}

func (sc SubordinateConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SubordinateConfigAttributes) InternalWithRef(ref terra.Reference) SubordinateConfigAttributes {
	return SubordinateConfigAttributes{ref: ref}
}

func (sc SubordinateConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SubordinateConfigAttributes) CertificateAuthority() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("certificate_authority"))
}

func (sc SubordinateConfigAttributes) PemIssuerChain() terra.ListValue[SubordinateConfigPemIssuerChainAttributes] {
	return terra.ReferenceAsList[SubordinateConfigPemIssuerChainAttributes](sc.ref.Append("pem_issuer_chain"))
}

type SubordinateConfigPemIssuerChainAttributes struct {
	ref terra.Reference
}

func (pic SubordinateConfigPemIssuerChainAttributes) InternalRef() (terra.Reference, error) {
	return pic.ref, nil
}

func (pic SubordinateConfigPemIssuerChainAttributes) InternalWithRef(ref terra.Reference) SubordinateConfigPemIssuerChainAttributes {
	return SubordinateConfigPemIssuerChainAttributes{ref: ref}
}

func (pic SubordinateConfigPemIssuerChainAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pic.ref.InternalTokens()
}

func (pic SubordinateConfigPemIssuerChainAttributes) PemCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pic.ref.Append("pem_certificates"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AccessUrlsState struct {
	CaCertificateAccessUrl string   `json:"ca_certificate_access_url"`
	CrlAccessUrls          []string `json:"crl_access_urls"`
}

type ConfigState struct {
	SubjectConfig []ConfigSubjectConfigState `json:"subject_config"`
	X509Config    []ConfigX509ConfigState    `json:"x509_config"`
}

type ConfigSubjectConfigState struct {
	Subject        []ConfigSubjectConfigSubjectState        `json:"subject"`
	SubjectAltName []ConfigSubjectConfigSubjectAltNameState `json:"subject_alt_name"`
}

type ConfigSubjectConfigSubjectState struct {
	CommonName         string `json:"common_name"`
	CountryCode        string `json:"country_code"`
	Locality           string `json:"locality"`
	Organization       string `json:"organization"`
	OrganizationalUnit string `json:"organizational_unit"`
	PostalCode         string `json:"postal_code"`
	Province           string `json:"province"`
	StreetAddress      string `json:"street_address"`
}

type ConfigSubjectConfigSubjectAltNameState struct {
	DnsNames       []string `json:"dns_names"`
	EmailAddresses []string `json:"email_addresses"`
	IpAddresses    []string `json:"ip_addresses"`
	Uris           []string `json:"uris"`
}

type ConfigX509ConfigState struct {
	AiaOcspServers       []string                                    `json:"aia_ocsp_servers"`
	AdditionalExtensions []ConfigX509ConfigAdditionalExtensionsState `json:"additional_extensions"`
	CaOptions            []ConfigX509ConfigCaOptionsState            `json:"ca_options"`
	KeyUsage             []ConfigX509ConfigKeyUsageState             `json:"key_usage"`
	NameConstraints      []ConfigX509ConfigNameConstraintsState      `json:"name_constraints"`
	PolicyIds            []ConfigX509ConfigPolicyIdsState            `json:"policy_ids"`
}

type ConfigX509ConfigAdditionalExtensionsState struct {
	Critical bool                                                `json:"critical"`
	Value    string                                              `json:"value"`
	ObjectId []ConfigX509ConfigAdditionalExtensionsObjectIdState `json:"object_id"`
}

type ConfigX509ConfigAdditionalExtensionsObjectIdState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type ConfigX509ConfigCaOptionsState struct {
	IsCa                    bool    `json:"is_ca"`
	MaxIssuerPathLength     float64 `json:"max_issuer_path_length"`
	NonCa                   bool    `json:"non_ca"`
	ZeroMaxIssuerPathLength bool    `json:"zero_max_issuer_path_length"`
}

type ConfigX509ConfigKeyUsageState struct {
	BaseKeyUsage             []ConfigX509ConfigKeyUsageBaseKeyUsageState             `json:"base_key_usage"`
	ExtendedKeyUsage         []ConfigX509ConfigKeyUsageExtendedKeyUsageState         `json:"extended_key_usage"`
	UnknownExtendedKeyUsages []ConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesState `json:"unknown_extended_key_usages"`
}

type ConfigX509ConfigKeyUsageBaseKeyUsageState struct {
	CertSign          bool `json:"cert_sign"`
	ContentCommitment bool `json:"content_commitment"`
	CrlSign           bool `json:"crl_sign"`
	DataEncipherment  bool `json:"data_encipherment"`
	DecipherOnly      bool `json:"decipher_only"`
	DigitalSignature  bool `json:"digital_signature"`
	EncipherOnly      bool `json:"encipher_only"`
	KeyAgreement      bool `json:"key_agreement"`
	KeyEncipherment   bool `json:"key_encipherment"`
}

type ConfigX509ConfigKeyUsageExtendedKeyUsageState struct {
	ClientAuth      bool `json:"client_auth"`
	CodeSigning     bool `json:"code_signing"`
	EmailProtection bool `json:"email_protection"`
	OcspSigning     bool `json:"ocsp_signing"`
	ServerAuth      bool `json:"server_auth"`
	TimeStamping    bool `json:"time_stamping"`
}

type ConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type ConfigX509ConfigNameConstraintsState struct {
	Critical                bool     `json:"critical"`
	ExcludedDnsNames        []string `json:"excluded_dns_names"`
	ExcludedEmailAddresses  []string `json:"excluded_email_addresses"`
	ExcludedIpRanges        []string `json:"excluded_ip_ranges"`
	ExcludedUris            []string `json:"excluded_uris"`
	PermittedDnsNames       []string `json:"permitted_dns_names"`
	PermittedEmailAddresses []string `json:"permitted_email_addresses"`
	PermittedIpRanges       []string `json:"permitted_ip_ranges"`
	PermittedUris           []string `json:"permitted_uris"`
}

type ConfigX509ConfigPolicyIdsState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type KeySpecState struct {
	Algorithm          string `json:"algorithm"`
	CloudKmsKeyVersion string `json:"cloud_kms_key_version"`
}

type SubordinateConfigState struct {
	CertificateAuthority string                                 `json:"certificate_authority"`
	PemIssuerChain       []SubordinateConfigPemIssuerChainState `json:"pem_issuer_chain"`
}

type SubordinateConfigPemIssuerChainState struct {
	PemCertificates []string `json:"pem_certificates"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
