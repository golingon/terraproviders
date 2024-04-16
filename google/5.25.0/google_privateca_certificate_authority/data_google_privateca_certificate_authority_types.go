// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_privateca_certificate_authority

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataAccessUrlsAttributes struct {
	ref terra.Reference
}

func (au DataAccessUrlsAttributes) InternalRef() (terra.Reference, error) {
	return au.ref, nil
}

func (au DataAccessUrlsAttributes) InternalWithRef(ref terra.Reference) DataAccessUrlsAttributes {
	return DataAccessUrlsAttributes{ref: ref}
}

func (au DataAccessUrlsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return au.ref.InternalTokens()
}

func (au DataAccessUrlsAttributes) CaCertificateAccessUrl() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("ca_certificate_access_url"))
}

func (au DataAccessUrlsAttributes) CrlAccessUrls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](au.ref.Append("crl_access_urls"))
}

type DataConfigAttributes struct {
	ref terra.Reference
}

func (c DataConfigAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c DataConfigAttributes) InternalWithRef(ref terra.Reference) DataConfigAttributes {
	return DataConfigAttributes{ref: ref}
}

func (c DataConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c DataConfigAttributes) SubjectConfig() terra.ListValue[DataConfigSubjectConfigAttributes] {
	return terra.ReferenceAsList[DataConfigSubjectConfigAttributes](c.ref.Append("subject_config"))
}

func (c DataConfigAttributes) X509Config() terra.ListValue[DataConfigX509ConfigAttributes] {
	return terra.ReferenceAsList[DataConfigX509ConfigAttributes](c.ref.Append("x509_config"))
}

type DataConfigSubjectConfigAttributes struct {
	ref terra.Reference
}

func (sc DataConfigSubjectConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc DataConfigSubjectConfigAttributes) InternalWithRef(ref terra.Reference) DataConfigSubjectConfigAttributes {
	return DataConfigSubjectConfigAttributes{ref: ref}
}

func (sc DataConfigSubjectConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc DataConfigSubjectConfigAttributes) Subject() terra.ListValue[DataConfigSubjectConfigSubjectAttributes] {
	return terra.ReferenceAsList[DataConfigSubjectConfigSubjectAttributes](sc.ref.Append("subject"))
}

func (sc DataConfigSubjectConfigAttributes) SubjectAltName() terra.ListValue[DataConfigSubjectConfigSubjectAltNameAttributes] {
	return terra.ReferenceAsList[DataConfigSubjectConfigSubjectAltNameAttributes](sc.ref.Append("subject_alt_name"))
}

type DataConfigSubjectConfigSubjectAttributes struct {
	ref terra.Reference
}

func (s DataConfigSubjectConfigSubjectAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s DataConfigSubjectConfigSubjectAttributes) InternalWithRef(ref terra.Reference) DataConfigSubjectConfigSubjectAttributes {
	return DataConfigSubjectConfigSubjectAttributes{ref: ref}
}

func (s DataConfigSubjectConfigSubjectAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s DataConfigSubjectConfigSubjectAttributes) CommonName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("common_name"))
}

func (s DataConfigSubjectConfigSubjectAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("country_code"))
}

func (s DataConfigSubjectConfigSubjectAttributes) Locality() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("locality"))
}

func (s DataConfigSubjectConfigSubjectAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("organization"))
}

func (s DataConfigSubjectConfigSubjectAttributes) OrganizationalUnit() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("organizational_unit"))
}

func (s DataConfigSubjectConfigSubjectAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("postal_code"))
}

func (s DataConfigSubjectConfigSubjectAttributes) Province() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("province"))
}

func (s DataConfigSubjectConfigSubjectAttributes) StreetAddress() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("street_address"))
}

type DataConfigSubjectConfigSubjectAltNameAttributes struct {
	ref terra.Reference
}

func (san DataConfigSubjectConfigSubjectAltNameAttributes) InternalRef() (terra.Reference, error) {
	return san.ref, nil
}

func (san DataConfigSubjectConfigSubjectAltNameAttributes) InternalWithRef(ref terra.Reference) DataConfigSubjectConfigSubjectAltNameAttributes {
	return DataConfigSubjectConfigSubjectAltNameAttributes{ref: ref}
}

func (san DataConfigSubjectConfigSubjectAltNameAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return san.ref.InternalTokens()
}

func (san DataConfigSubjectConfigSubjectAltNameAttributes) DnsNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](san.ref.Append("dns_names"))
}

func (san DataConfigSubjectConfigSubjectAltNameAttributes) EmailAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](san.ref.Append("email_addresses"))
}

func (san DataConfigSubjectConfigSubjectAltNameAttributes) IpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](san.ref.Append("ip_addresses"))
}

func (san DataConfigSubjectConfigSubjectAltNameAttributes) Uris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](san.ref.Append("uris"))
}

type DataConfigX509ConfigAttributes struct {
	ref terra.Reference
}

func (xc DataConfigX509ConfigAttributes) InternalRef() (terra.Reference, error) {
	return xc.ref, nil
}

func (xc DataConfigX509ConfigAttributes) InternalWithRef(ref terra.Reference) DataConfigX509ConfigAttributes {
	return DataConfigX509ConfigAttributes{ref: ref}
}

func (xc DataConfigX509ConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return xc.ref.InternalTokens()
}

func (xc DataConfigX509ConfigAttributes) AiaOcspServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](xc.ref.Append("aia_ocsp_servers"))
}

func (xc DataConfigX509ConfigAttributes) AdditionalExtensions() terra.ListValue[DataConfigX509ConfigAdditionalExtensionsAttributes] {
	return terra.ReferenceAsList[DataConfigX509ConfigAdditionalExtensionsAttributes](xc.ref.Append("additional_extensions"))
}

func (xc DataConfigX509ConfigAttributes) CaOptions() terra.ListValue[DataConfigX509ConfigCaOptionsAttributes] {
	return terra.ReferenceAsList[DataConfigX509ConfigCaOptionsAttributes](xc.ref.Append("ca_options"))
}

func (xc DataConfigX509ConfigAttributes) KeyUsage() terra.ListValue[DataConfigX509ConfigKeyUsageAttributes] {
	return terra.ReferenceAsList[DataConfigX509ConfigKeyUsageAttributes](xc.ref.Append("key_usage"))
}

func (xc DataConfigX509ConfigAttributes) NameConstraints() terra.ListValue[DataConfigX509ConfigNameConstraintsAttributes] {
	return terra.ReferenceAsList[DataConfigX509ConfigNameConstraintsAttributes](xc.ref.Append("name_constraints"))
}

func (xc DataConfigX509ConfigAttributes) PolicyIds() terra.ListValue[DataConfigX509ConfigPolicyIdsAttributes] {
	return terra.ReferenceAsList[DataConfigX509ConfigPolicyIdsAttributes](xc.ref.Append("policy_ids"))
}

type DataConfigX509ConfigAdditionalExtensionsAttributes struct {
	ref terra.Reference
}

func (ae DataConfigX509ConfigAdditionalExtensionsAttributes) InternalRef() (terra.Reference, error) {
	return ae.ref, nil
}

func (ae DataConfigX509ConfigAdditionalExtensionsAttributes) InternalWithRef(ref terra.Reference) DataConfigX509ConfigAdditionalExtensionsAttributes {
	return DataConfigX509ConfigAdditionalExtensionsAttributes{ref: ref}
}

func (ae DataConfigX509ConfigAdditionalExtensionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ae.ref.InternalTokens()
}

func (ae DataConfigX509ConfigAdditionalExtensionsAttributes) Critical() terra.BoolValue {
	return terra.ReferenceAsBool(ae.ref.Append("critical"))
}

func (ae DataConfigX509ConfigAdditionalExtensionsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("value"))
}

func (ae DataConfigX509ConfigAdditionalExtensionsAttributes) ObjectId() terra.ListValue[DataConfigX509ConfigAdditionalExtensionsObjectIdAttributes] {
	return terra.ReferenceAsList[DataConfigX509ConfigAdditionalExtensionsObjectIdAttributes](ae.ref.Append("object_id"))
}

type DataConfigX509ConfigAdditionalExtensionsObjectIdAttributes struct {
	ref terra.Reference
}

func (oi DataConfigX509ConfigAdditionalExtensionsObjectIdAttributes) InternalRef() (terra.Reference, error) {
	return oi.ref, nil
}

func (oi DataConfigX509ConfigAdditionalExtensionsObjectIdAttributes) InternalWithRef(ref terra.Reference) DataConfigX509ConfigAdditionalExtensionsObjectIdAttributes {
	return DataConfigX509ConfigAdditionalExtensionsObjectIdAttributes{ref: ref}
}

func (oi DataConfigX509ConfigAdditionalExtensionsObjectIdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oi.ref.InternalTokens()
}

func (oi DataConfigX509ConfigAdditionalExtensionsObjectIdAttributes) ObjectIdPath() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](oi.ref.Append("object_id_path"))
}

type DataConfigX509ConfigCaOptionsAttributes struct {
	ref terra.Reference
}

func (co DataConfigX509ConfigCaOptionsAttributes) InternalRef() (terra.Reference, error) {
	return co.ref, nil
}

func (co DataConfigX509ConfigCaOptionsAttributes) InternalWithRef(ref terra.Reference) DataConfigX509ConfigCaOptionsAttributes {
	return DataConfigX509ConfigCaOptionsAttributes{ref: ref}
}

func (co DataConfigX509ConfigCaOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return co.ref.InternalTokens()
}

func (co DataConfigX509ConfigCaOptionsAttributes) IsCa() terra.BoolValue {
	return terra.ReferenceAsBool(co.ref.Append("is_ca"))
}

func (co DataConfigX509ConfigCaOptionsAttributes) MaxIssuerPathLength() terra.NumberValue {
	return terra.ReferenceAsNumber(co.ref.Append("max_issuer_path_length"))
}

func (co DataConfigX509ConfigCaOptionsAttributes) NonCa() terra.BoolValue {
	return terra.ReferenceAsBool(co.ref.Append("non_ca"))
}

func (co DataConfigX509ConfigCaOptionsAttributes) ZeroMaxIssuerPathLength() terra.BoolValue {
	return terra.ReferenceAsBool(co.ref.Append("zero_max_issuer_path_length"))
}

type DataConfigX509ConfigKeyUsageAttributes struct {
	ref terra.Reference
}

func (ku DataConfigX509ConfigKeyUsageAttributes) InternalRef() (terra.Reference, error) {
	return ku.ref, nil
}

func (ku DataConfigX509ConfigKeyUsageAttributes) InternalWithRef(ref terra.Reference) DataConfigX509ConfigKeyUsageAttributes {
	return DataConfigX509ConfigKeyUsageAttributes{ref: ref}
}

func (ku DataConfigX509ConfigKeyUsageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ku.ref.InternalTokens()
}

func (ku DataConfigX509ConfigKeyUsageAttributes) BaseKeyUsage() terra.ListValue[DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes] {
	return terra.ReferenceAsList[DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes](ku.ref.Append("base_key_usage"))
}

func (ku DataConfigX509ConfigKeyUsageAttributes) ExtendedKeyUsage() terra.ListValue[DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes] {
	return terra.ReferenceAsList[DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes](ku.ref.Append("extended_key_usage"))
}

func (ku DataConfigX509ConfigKeyUsageAttributes) UnknownExtendedKeyUsages() terra.ListValue[DataConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes] {
	return terra.ReferenceAsList[DataConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes](ku.ref.Append("unknown_extended_key_usages"))
}

type DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes struct {
	ref terra.Reference
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) InternalRef() (terra.Reference, error) {
	return bku.ref, nil
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) InternalWithRef(ref terra.Reference) DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes {
	return DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes{ref: ref}
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bku.ref.InternalTokens()
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) CertSign() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("cert_sign"))
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) ContentCommitment() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("content_commitment"))
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) CrlSign() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("crl_sign"))
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) DataEncipherment() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("data_encipherment"))
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) DecipherOnly() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("decipher_only"))
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) DigitalSignature() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("digital_signature"))
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) EncipherOnly() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("encipher_only"))
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) KeyAgreement() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("key_agreement"))
}

func (bku DataConfigX509ConfigKeyUsageBaseKeyUsageAttributes) KeyEncipherment() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("key_encipherment"))
}

type DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes struct {
	ref terra.Reference
}

func (eku DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) InternalRef() (terra.Reference, error) {
	return eku.ref, nil
}

func (eku DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) InternalWithRef(ref terra.Reference) DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes {
	return DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes{ref: ref}
}

func (eku DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eku.ref.InternalTokens()
}

func (eku DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) ClientAuth() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("client_auth"))
}

func (eku DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) CodeSigning() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("code_signing"))
}

func (eku DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) EmailProtection() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("email_protection"))
}

func (eku DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) OcspSigning() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("ocsp_signing"))
}

func (eku DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) ServerAuth() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("server_auth"))
}

func (eku DataConfigX509ConfigKeyUsageExtendedKeyUsageAttributes) TimeStamping() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("time_stamping"))
}

type DataConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes struct {
	ref terra.Reference
}

func (ueku DataConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes) InternalRef() (terra.Reference, error) {
	return ueku.ref, nil
}

func (ueku DataConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes) InternalWithRef(ref terra.Reference) DataConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes {
	return DataConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes{ref: ref}
}

func (ueku DataConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ueku.ref.InternalTokens()
}

func (ueku DataConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesAttributes) ObjectIdPath() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](ueku.ref.Append("object_id_path"))
}

type DataConfigX509ConfigNameConstraintsAttributes struct {
	ref terra.Reference
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) InternalWithRef(ref terra.Reference) DataConfigX509ConfigNameConstraintsAttributes {
	return DataConfigX509ConfigNameConstraintsAttributes{ref: ref}
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) Critical() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("critical"))
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) ExcludedDnsNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_dns_names"))
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) ExcludedEmailAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_email_addresses"))
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) ExcludedIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_ip_ranges"))
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) ExcludedUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_uris"))
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) PermittedDnsNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_dns_names"))
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) PermittedEmailAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_email_addresses"))
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) PermittedIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_ip_ranges"))
}

func (nc DataConfigX509ConfigNameConstraintsAttributes) PermittedUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_uris"))
}

type DataConfigX509ConfigPolicyIdsAttributes struct {
	ref terra.Reference
}

func (pi DataConfigX509ConfigPolicyIdsAttributes) InternalRef() (terra.Reference, error) {
	return pi.ref, nil
}

func (pi DataConfigX509ConfigPolicyIdsAttributes) InternalWithRef(ref terra.Reference) DataConfigX509ConfigPolicyIdsAttributes {
	return DataConfigX509ConfigPolicyIdsAttributes{ref: ref}
}

func (pi DataConfigX509ConfigPolicyIdsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pi.ref.InternalTokens()
}

func (pi DataConfigX509ConfigPolicyIdsAttributes) ObjectIdPath() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](pi.ref.Append("object_id_path"))
}

type DataKeySpecAttributes struct {
	ref terra.Reference
}

func (ks DataKeySpecAttributes) InternalRef() (terra.Reference, error) {
	return ks.ref, nil
}

func (ks DataKeySpecAttributes) InternalWithRef(ref terra.Reference) DataKeySpecAttributes {
	return DataKeySpecAttributes{ref: ref}
}

func (ks DataKeySpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ks.ref.InternalTokens()
}

func (ks DataKeySpecAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("algorithm"))
}

func (ks DataKeySpecAttributes) CloudKmsKeyVersion() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("cloud_kms_key_version"))
}

type DataSubordinateConfigAttributes struct {
	ref terra.Reference
}

func (sc DataSubordinateConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc DataSubordinateConfigAttributes) InternalWithRef(ref terra.Reference) DataSubordinateConfigAttributes {
	return DataSubordinateConfigAttributes{ref: ref}
}

func (sc DataSubordinateConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc DataSubordinateConfigAttributes) CertificateAuthority() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("certificate_authority"))
}

func (sc DataSubordinateConfigAttributes) PemIssuerChain() terra.ListValue[DataSubordinateConfigPemIssuerChainAttributes] {
	return terra.ReferenceAsList[DataSubordinateConfigPemIssuerChainAttributes](sc.ref.Append("pem_issuer_chain"))
}

type DataSubordinateConfigPemIssuerChainAttributes struct {
	ref terra.Reference
}

func (pic DataSubordinateConfigPemIssuerChainAttributes) InternalRef() (terra.Reference, error) {
	return pic.ref, nil
}

func (pic DataSubordinateConfigPemIssuerChainAttributes) InternalWithRef(ref terra.Reference) DataSubordinateConfigPemIssuerChainAttributes {
	return DataSubordinateConfigPemIssuerChainAttributes{ref: ref}
}

func (pic DataSubordinateConfigPemIssuerChainAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pic.ref.InternalTokens()
}

func (pic DataSubordinateConfigPemIssuerChainAttributes) PemCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pic.ref.Append("pem_certificates"))
}

type DataAccessUrlsState struct {
	CaCertificateAccessUrl string   `json:"ca_certificate_access_url"`
	CrlAccessUrls          []string `json:"crl_access_urls"`
}

type DataConfigState struct {
	SubjectConfig []DataConfigSubjectConfigState `json:"subject_config"`
	X509Config    []DataConfigX509ConfigState    `json:"x509_config"`
}

type DataConfigSubjectConfigState struct {
	Subject        []DataConfigSubjectConfigSubjectState        `json:"subject"`
	SubjectAltName []DataConfigSubjectConfigSubjectAltNameState `json:"subject_alt_name"`
}

type DataConfigSubjectConfigSubjectState struct {
	CommonName         string `json:"common_name"`
	CountryCode        string `json:"country_code"`
	Locality           string `json:"locality"`
	Organization       string `json:"organization"`
	OrganizationalUnit string `json:"organizational_unit"`
	PostalCode         string `json:"postal_code"`
	Province           string `json:"province"`
	StreetAddress      string `json:"street_address"`
}

type DataConfigSubjectConfigSubjectAltNameState struct {
	DnsNames       []string `json:"dns_names"`
	EmailAddresses []string `json:"email_addresses"`
	IpAddresses    []string `json:"ip_addresses"`
	Uris           []string `json:"uris"`
}

type DataConfigX509ConfigState struct {
	AiaOcspServers       []string                                        `json:"aia_ocsp_servers"`
	AdditionalExtensions []DataConfigX509ConfigAdditionalExtensionsState `json:"additional_extensions"`
	CaOptions            []DataConfigX509ConfigCaOptionsState            `json:"ca_options"`
	KeyUsage             []DataConfigX509ConfigKeyUsageState             `json:"key_usage"`
	NameConstraints      []DataConfigX509ConfigNameConstraintsState      `json:"name_constraints"`
	PolicyIds            []DataConfigX509ConfigPolicyIdsState            `json:"policy_ids"`
}

type DataConfigX509ConfigAdditionalExtensionsState struct {
	Critical bool                                                    `json:"critical"`
	Value    string                                                  `json:"value"`
	ObjectId []DataConfigX509ConfigAdditionalExtensionsObjectIdState `json:"object_id"`
}

type DataConfigX509ConfigAdditionalExtensionsObjectIdState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type DataConfigX509ConfigCaOptionsState struct {
	IsCa                    bool    `json:"is_ca"`
	MaxIssuerPathLength     float64 `json:"max_issuer_path_length"`
	NonCa                   bool    `json:"non_ca"`
	ZeroMaxIssuerPathLength bool    `json:"zero_max_issuer_path_length"`
}

type DataConfigX509ConfigKeyUsageState struct {
	BaseKeyUsage             []DataConfigX509ConfigKeyUsageBaseKeyUsageState             `json:"base_key_usage"`
	ExtendedKeyUsage         []DataConfigX509ConfigKeyUsageExtendedKeyUsageState         `json:"extended_key_usage"`
	UnknownExtendedKeyUsages []DataConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesState `json:"unknown_extended_key_usages"`
}

type DataConfigX509ConfigKeyUsageBaseKeyUsageState struct {
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

type DataConfigX509ConfigKeyUsageExtendedKeyUsageState struct {
	ClientAuth      bool `json:"client_auth"`
	CodeSigning     bool `json:"code_signing"`
	EmailProtection bool `json:"email_protection"`
	OcspSigning     bool `json:"ocsp_signing"`
	ServerAuth      bool `json:"server_auth"`
	TimeStamping    bool `json:"time_stamping"`
}

type DataConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type DataConfigX509ConfigNameConstraintsState struct {
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

type DataConfigX509ConfigPolicyIdsState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type DataKeySpecState struct {
	Algorithm          string `json:"algorithm"`
	CloudKmsKeyVersion string `json:"cloud_kms_key_version"`
}

type DataSubordinateConfigState struct {
	CertificateAuthority string                                     `json:"certificate_authority"`
	PemIssuerChain       []DataSubordinateConfigPemIssuerChainState `json:"pem_issuer_chain"`
}

type DataSubordinateConfigPemIssuerChainState struct {
	PemCertificates []string `json:"pem_certificates"`
}
