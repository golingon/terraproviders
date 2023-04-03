// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package privatecacapool

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type IssuancePolicy struct {
	// MaximumLifetime: string, optional
	MaximumLifetime terra.StringValue `hcl:"maximum_lifetime,attr"`
	// AllowedIssuanceModes: optional
	AllowedIssuanceModes *AllowedIssuanceModes `hcl:"allowed_issuance_modes,block"`
	// AllowedKeyTypes: min=0
	AllowedKeyTypes []AllowedKeyTypes `hcl:"allowed_key_types,block" validate:"min=0"`
	// BaselineValues: optional
	BaselineValues *BaselineValues `hcl:"baseline_values,block"`
	// IdentityConstraints: optional
	IdentityConstraints *IdentityConstraints `hcl:"identity_constraints,block"`
}

type AllowedIssuanceModes struct {
	// AllowConfigBasedIssuance: bool, required
	AllowConfigBasedIssuance terra.BoolValue `hcl:"allow_config_based_issuance,attr" validate:"required"`
	// AllowCsrBasedIssuance: bool, required
	AllowCsrBasedIssuance terra.BoolValue `hcl:"allow_csr_based_issuance,attr" validate:"required"`
}

type AllowedKeyTypes struct {
	// EllipticCurve: optional
	EllipticCurve *EllipticCurve `hcl:"elliptic_curve,block"`
	// Rsa: optional
	Rsa *Rsa `hcl:"rsa,block"`
}

type EllipticCurve struct {
	// SignatureAlgorithm: string, required
	SignatureAlgorithm terra.StringValue `hcl:"signature_algorithm,attr" validate:"required"`
}

type Rsa struct {
	// MaxModulusSize: string, optional
	MaxModulusSize terra.StringValue `hcl:"max_modulus_size,attr"`
	// MinModulusSize: string, optional
	MinModulusSize terra.StringValue `hcl:"min_modulus_size,attr"`
}

type BaselineValues struct {
	// AiaOcspServers: list of string, optional
	AiaOcspServers terra.ListValue[terra.StringValue] `hcl:"aia_ocsp_servers,attr"`
	// AdditionalExtensions: min=0
	AdditionalExtensions []AdditionalExtensions `hcl:"additional_extensions,block" validate:"min=0"`
	// CaOptions: required
	CaOptions *CaOptions `hcl:"ca_options,block" validate:"required"`
	// KeyUsage: required
	KeyUsage *KeyUsage `hcl:"key_usage,block" validate:"required"`
	// NameConstraints: optional
	NameConstraints *NameConstraints `hcl:"name_constraints,block"`
	// PolicyIds: min=0
	PolicyIds []PolicyIds `hcl:"policy_ids,block" validate:"min=0"`
}

type AdditionalExtensions struct {
	// Critical: bool, required
	Critical terra.BoolValue `hcl:"critical,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// ObjectId: required
	ObjectId *ObjectId `hcl:"object_id,block" validate:"required"`
}

type ObjectId struct {
	// ObjectIdPath: list of number, required
	ObjectIdPath terra.ListValue[terra.NumberValue] `hcl:"object_id_path,attr" validate:"required"`
}

type CaOptions struct {
	// IsCa: bool, optional
	IsCa terra.BoolValue `hcl:"is_ca,attr"`
	// MaxIssuerPathLength: number, optional
	MaxIssuerPathLength terra.NumberValue `hcl:"max_issuer_path_length,attr"`
	// NonCa: bool, optional
	NonCa terra.BoolValue `hcl:"non_ca,attr"`
	// ZeroMaxIssuerPathLength: bool, optional
	ZeroMaxIssuerPathLength terra.BoolValue `hcl:"zero_max_issuer_path_length,attr"`
}

type KeyUsage struct {
	// BaseKeyUsage: required
	BaseKeyUsage *BaseKeyUsage `hcl:"base_key_usage,block" validate:"required"`
	// ExtendedKeyUsage: required
	ExtendedKeyUsage *ExtendedKeyUsage `hcl:"extended_key_usage,block" validate:"required"`
	// UnknownExtendedKeyUsages: min=0
	UnknownExtendedKeyUsages []UnknownExtendedKeyUsages `hcl:"unknown_extended_key_usages,block" validate:"min=0"`
}

type BaseKeyUsage struct {
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

type ExtendedKeyUsage struct {
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

type UnknownExtendedKeyUsages struct {
	// ObjectIdPath: list of number, required
	ObjectIdPath terra.ListValue[terra.NumberValue] `hcl:"object_id_path,attr" validate:"required"`
}

type NameConstraints struct {
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

type PolicyIds struct {
	// ObjectIdPath: list of number, required
	ObjectIdPath terra.ListValue[terra.NumberValue] `hcl:"object_id_path,attr" validate:"required"`
}

type IdentityConstraints struct {
	// AllowSubjectAltNamesPassthrough: bool, required
	AllowSubjectAltNamesPassthrough terra.BoolValue `hcl:"allow_subject_alt_names_passthrough,attr" validate:"required"`
	// AllowSubjectPassthrough: bool, required
	AllowSubjectPassthrough terra.BoolValue `hcl:"allow_subject_passthrough,attr" validate:"required"`
	// CelExpression: optional
	CelExpression *CelExpression `hcl:"cel_expression,block"`
}

type CelExpression struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Expression: string, required
	Expression terra.StringValue `hcl:"expression,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Title: string, optional
	Title terra.StringValue `hcl:"title,attr"`
}

type PublishingOptions struct {
	// PublishCaCert: bool, required
	PublishCaCert terra.BoolValue `hcl:"publish_ca_cert,attr" validate:"required"`
	// PublishCrl: bool, required
	PublishCrl terra.BoolValue `hcl:"publish_crl,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type IssuancePolicyAttributes struct {
	ref terra.Reference
}

func (ip IssuancePolicyAttributes) InternalRef() (terra.Reference, error) {
	return ip.ref, nil
}

func (ip IssuancePolicyAttributes) InternalWithRef(ref terra.Reference) IssuancePolicyAttributes {
	return IssuancePolicyAttributes{ref: ref}
}

func (ip IssuancePolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ip.ref.InternalTokens()
}

func (ip IssuancePolicyAttributes) MaximumLifetime() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("maximum_lifetime"))
}

func (ip IssuancePolicyAttributes) AllowedIssuanceModes() terra.ListValue[AllowedIssuanceModesAttributes] {
	return terra.ReferenceAsList[AllowedIssuanceModesAttributes](ip.ref.Append("allowed_issuance_modes"))
}

func (ip IssuancePolicyAttributes) AllowedKeyTypes() terra.ListValue[AllowedKeyTypesAttributes] {
	return terra.ReferenceAsList[AllowedKeyTypesAttributes](ip.ref.Append("allowed_key_types"))
}

func (ip IssuancePolicyAttributes) BaselineValues() terra.ListValue[BaselineValuesAttributes] {
	return terra.ReferenceAsList[BaselineValuesAttributes](ip.ref.Append("baseline_values"))
}

func (ip IssuancePolicyAttributes) IdentityConstraints() terra.ListValue[IdentityConstraintsAttributes] {
	return terra.ReferenceAsList[IdentityConstraintsAttributes](ip.ref.Append("identity_constraints"))
}

type AllowedIssuanceModesAttributes struct {
	ref terra.Reference
}

func (aim AllowedIssuanceModesAttributes) InternalRef() (terra.Reference, error) {
	return aim.ref, nil
}

func (aim AllowedIssuanceModesAttributes) InternalWithRef(ref terra.Reference) AllowedIssuanceModesAttributes {
	return AllowedIssuanceModesAttributes{ref: ref}
}

func (aim AllowedIssuanceModesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aim.ref.InternalTokens()
}

func (aim AllowedIssuanceModesAttributes) AllowConfigBasedIssuance() terra.BoolValue {
	return terra.ReferenceAsBool(aim.ref.Append("allow_config_based_issuance"))
}

func (aim AllowedIssuanceModesAttributes) AllowCsrBasedIssuance() terra.BoolValue {
	return terra.ReferenceAsBool(aim.ref.Append("allow_csr_based_issuance"))
}

type AllowedKeyTypesAttributes struct {
	ref terra.Reference
}

func (akt AllowedKeyTypesAttributes) InternalRef() (terra.Reference, error) {
	return akt.ref, nil
}

func (akt AllowedKeyTypesAttributes) InternalWithRef(ref terra.Reference) AllowedKeyTypesAttributes {
	return AllowedKeyTypesAttributes{ref: ref}
}

func (akt AllowedKeyTypesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return akt.ref.InternalTokens()
}

func (akt AllowedKeyTypesAttributes) EllipticCurve() terra.ListValue[EllipticCurveAttributes] {
	return terra.ReferenceAsList[EllipticCurveAttributes](akt.ref.Append("elliptic_curve"))
}

func (akt AllowedKeyTypesAttributes) Rsa() terra.ListValue[RsaAttributes] {
	return terra.ReferenceAsList[RsaAttributes](akt.ref.Append("rsa"))
}

type EllipticCurveAttributes struct {
	ref terra.Reference
}

func (ec EllipticCurveAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EllipticCurveAttributes) InternalWithRef(ref terra.Reference) EllipticCurveAttributes {
	return EllipticCurveAttributes{ref: ref}
}

func (ec EllipticCurveAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EllipticCurveAttributes) SignatureAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("signature_algorithm"))
}

type RsaAttributes struct {
	ref terra.Reference
}

func (r RsaAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RsaAttributes) InternalWithRef(ref terra.Reference) RsaAttributes {
	return RsaAttributes{ref: ref}
}

func (r RsaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RsaAttributes) MaxModulusSize() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("max_modulus_size"))
}

func (r RsaAttributes) MinModulusSize() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("min_modulus_size"))
}

type BaselineValuesAttributes struct {
	ref terra.Reference
}

func (bv BaselineValuesAttributes) InternalRef() (terra.Reference, error) {
	return bv.ref, nil
}

func (bv BaselineValuesAttributes) InternalWithRef(ref terra.Reference) BaselineValuesAttributes {
	return BaselineValuesAttributes{ref: ref}
}

func (bv BaselineValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bv.ref.InternalTokens()
}

func (bv BaselineValuesAttributes) AiaOcspServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bv.ref.Append("aia_ocsp_servers"))
}

func (bv BaselineValuesAttributes) AdditionalExtensions() terra.ListValue[AdditionalExtensionsAttributes] {
	return terra.ReferenceAsList[AdditionalExtensionsAttributes](bv.ref.Append("additional_extensions"))
}

func (bv BaselineValuesAttributes) CaOptions() terra.ListValue[CaOptionsAttributes] {
	return terra.ReferenceAsList[CaOptionsAttributes](bv.ref.Append("ca_options"))
}

func (bv BaselineValuesAttributes) KeyUsage() terra.ListValue[KeyUsageAttributes] {
	return terra.ReferenceAsList[KeyUsageAttributes](bv.ref.Append("key_usage"))
}

func (bv BaselineValuesAttributes) NameConstraints() terra.ListValue[NameConstraintsAttributes] {
	return terra.ReferenceAsList[NameConstraintsAttributes](bv.ref.Append("name_constraints"))
}

func (bv BaselineValuesAttributes) PolicyIds() terra.ListValue[PolicyIdsAttributes] {
	return terra.ReferenceAsList[PolicyIdsAttributes](bv.ref.Append("policy_ids"))
}

type AdditionalExtensionsAttributes struct {
	ref terra.Reference
}

func (ae AdditionalExtensionsAttributes) InternalRef() (terra.Reference, error) {
	return ae.ref, nil
}

func (ae AdditionalExtensionsAttributes) InternalWithRef(ref terra.Reference) AdditionalExtensionsAttributes {
	return AdditionalExtensionsAttributes{ref: ref}
}

func (ae AdditionalExtensionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ae.ref.InternalTokens()
}

func (ae AdditionalExtensionsAttributes) Critical() terra.BoolValue {
	return terra.ReferenceAsBool(ae.ref.Append("critical"))
}

func (ae AdditionalExtensionsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("value"))
}

func (ae AdditionalExtensionsAttributes) ObjectId() terra.ListValue[ObjectIdAttributes] {
	return terra.ReferenceAsList[ObjectIdAttributes](ae.ref.Append("object_id"))
}

type ObjectIdAttributes struct {
	ref terra.Reference
}

func (oi ObjectIdAttributes) InternalRef() (terra.Reference, error) {
	return oi.ref, nil
}

func (oi ObjectIdAttributes) InternalWithRef(ref terra.Reference) ObjectIdAttributes {
	return ObjectIdAttributes{ref: ref}
}

func (oi ObjectIdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oi.ref.InternalTokens()
}

func (oi ObjectIdAttributes) ObjectIdPath() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](oi.ref.Append("object_id_path"))
}

type CaOptionsAttributes struct {
	ref terra.Reference
}

func (co CaOptionsAttributes) InternalRef() (terra.Reference, error) {
	return co.ref, nil
}

func (co CaOptionsAttributes) InternalWithRef(ref terra.Reference) CaOptionsAttributes {
	return CaOptionsAttributes{ref: ref}
}

func (co CaOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return co.ref.InternalTokens()
}

func (co CaOptionsAttributes) IsCa() terra.BoolValue {
	return terra.ReferenceAsBool(co.ref.Append("is_ca"))
}

func (co CaOptionsAttributes) MaxIssuerPathLength() terra.NumberValue {
	return terra.ReferenceAsNumber(co.ref.Append("max_issuer_path_length"))
}

func (co CaOptionsAttributes) NonCa() terra.BoolValue {
	return terra.ReferenceAsBool(co.ref.Append("non_ca"))
}

func (co CaOptionsAttributes) ZeroMaxIssuerPathLength() terra.BoolValue {
	return terra.ReferenceAsBool(co.ref.Append("zero_max_issuer_path_length"))
}

type KeyUsageAttributes struct {
	ref terra.Reference
}

func (ku KeyUsageAttributes) InternalRef() (terra.Reference, error) {
	return ku.ref, nil
}

func (ku KeyUsageAttributes) InternalWithRef(ref terra.Reference) KeyUsageAttributes {
	return KeyUsageAttributes{ref: ref}
}

func (ku KeyUsageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ku.ref.InternalTokens()
}

func (ku KeyUsageAttributes) BaseKeyUsage() terra.ListValue[BaseKeyUsageAttributes] {
	return terra.ReferenceAsList[BaseKeyUsageAttributes](ku.ref.Append("base_key_usage"))
}

func (ku KeyUsageAttributes) ExtendedKeyUsage() terra.ListValue[ExtendedKeyUsageAttributes] {
	return terra.ReferenceAsList[ExtendedKeyUsageAttributes](ku.ref.Append("extended_key_usage"))
}

func (ku KeyUsageAttributes) UnknownExtendedKeyUsages() terra.ListValue[UnknownExtendedKeyUsagesAttributes] {
	return terra.ReferenceAsList[UnknownExtendedKeyUsagesAttributes](ku.ref.Append("unknown_extended_key_usages"))
}

type BaseKeyUsageAttributes struct {
	ref terra.Reference
}

func (bku BaseKeyUsageAttributes) InternalRef() (terra.Reference, error) {
	return bku.ref, nil
}

func (bku BaseKeyUsageAttributes) InternalWithRef(ref terra.Reference) BaseKeyUsageAttributes {
	return BaseKeyUsageAttributes{ref: ref}
}

func (bku BaseKeyUsageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bku.ref.InternalTokens()
}

func (bku BaseKeyUsageAttributes) CertSign() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("cert_sign"))
}

func (bku BaseKeyUsageAttributes) ContentCommitment() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("content_commitment"))
}

func (bku BaseKeyUsageAttributes) CrlSign() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("crl_sign"))
}

func (bku BaseKeyUsageAttributes) DataEncipherment() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("data_encipherment"))
}

func (bku BaseKeyUsageAttributes) DecipherOnly() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("decipher_only"))
}

func (bku BaseKeyUsageAttributes) DigitalSignature() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("digital_signature"))
}

func (bku BaseKeyUsageAttributes) EncipherOnly() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("encipher_only"))
}

func (bku BaseKeyUsageAttributes) KeyAgreement() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("key_agreement"))
}

func (bku BaseKeyUsageAttributes) KeyEncipherment() terra.BoolValue {
	return terra.ReferenceAsBool(bku.ref.Append("key_encipherment"))
}

type ExtendedKeyUsageAttributes struct {
	ref terra.Reference
}

func (eku ExtendedKeyUsageAttributes) InternalRef() (terra.Reference, error) {
	return eku.ref, nil
}

func (eku ExtendedKeyUsageAttributes) InternalWithRef(ref terra.Reference) ExtendedKeyUsageAttributes {
	return ExtendedKeyUsageAttributes{ref: ref}
}

func (eku ExtendedKeyUsageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eku.ref.InternalTokens()
}

func (eku ExtendedKeyUsageAttributes) ClientAuth() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("client_auth"))
}

func (eku ExtendedKeyUsageAttributes) CodeSigning() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("code_signing"))
}

func (eku ExtendedKeyUsageAttributes) EmailProtection() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("email_protection"))
}

func (eku ExtendedKeyUsageAttributes) OcspSigning() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("ocsp_signing"))
}

func (eku ExtendedKeyUsageAttributes) ServerAuth() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("server_auth"))
}

func (eku ExtendedKeyUsageAttributes) TimeStamping() terra.BoolValue {
	return terra.ReferenceAsBool(eku.ref.Append("time_stamping"))
}

type UnknownExtendedKeyUsagesAttributes struct {
	ref terra.Reference
}

func (ueku UnknownExtendedKeyUsagesAttributes) InternalRef() (terra.Reference, error) {
	return ueku.ref, nil
}

func (ueku UnknownExtendedKeyUsagesAttributes) InternalWithRef(ref terra.Reference) UnknownExtendedKeyUsagesAttributes {
	return UnknownExtendedKeyUsagesAttributes{ref: ref}
}

func (ueku UnknownExtendedKeyUsagesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ueku.ref.InternalTokens()
}

func (ueku UnknownExtendedKeyUsagesAttributes) ObjectIdPath() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](ueku.ref.Append("object_id_path"))
}

type NameConstraintsAttributes struct {
	ref terra.Reference
}

func (nc NameConstraintsAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NameConstraintsAttributes) InternalWithRef(ref terra.Reference) NameConstraintsAttributes {
	return NameConstraintsAttributes{ref: ref}
}

func (nc NameConstraintsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NameConstraintsAttributes) Critical() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("critical"))
}

func (nc NameConstraintsAttributes) ExcludedDnsNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_dns_names"))
}

func (nc NameConstraintsAttributes) ExcludedEmailAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_email_addresses"))
}

func (nc NameConstraintsAttributes) ExcludedIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_ip_ranges"))
}

func (nc NameConstraintsAttributes) ExcludedUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("excluded_uris"))
}

func (nc NameConstraintsAttributes) PermittedDnsNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_dns_names"))
}

func (nc NameConstraintsAttributes) PermittedEmailAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_email_addresses"))
}

func (nc NameConstraintsAttributes) PermittedIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_ip_ranges"))
}

func (nc NameConstraintsAttributes) PermittedUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("permitted_uris"))
}

type PolicyIdsAttributes struct {
	ref terra.Reference
}

func (pi PolicyIdsAttributes) InternalRef() (terra.Reference, error) {
	return pi.ref, nil
}

func (pi PolicyIdsAttributes) InternalWithRef(ref terra.Reference) PolicyIdsAttributes {
	return PolicyIdsAttributes{ref: ref}
}

func (pi PolicyIdsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pi.ref.InternalTokens()
}

func (pi PolicyIdsAttributes) ObjectIdPath() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](pi.ref.Append("object_id_path"))
}

type IdentityConstraintsAttributes struct {
	ref terra.Reference
}

func (ic IdentityConstraintsAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic IdentityConstraintsAttributes) InternalWithRef(ref terra.Reference) IdentityConstraintsAttributes {
	return IdentityConstraintsAttributes{ref: ref}
}

func (ic IdentityConstraintsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic IdentityConstraintsAttributes) AllowSubjectAltNamesPassthrough() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("allow_subject_alt_names_passthrough"))
}

func (ic IdentityConstraintsAttributes) AllowSubjectPassthrough() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("allow_subject_passthrough"))
}

func (ic IdentityConstraintsAttributes) CelExpression() terra.ListValue[CelExpressionAttributes] {
	return terra.ReferenceAsList[CelExpressionAttributes](ic.ref.Append("cel_expression"))
}

type CelExpressionAttributes struct {
	ref terra.Reference
}

func (ce CelExpressionAttributes) InternalRef() (terra.Reference, error) {
	return ce.ref, nil
}

func (ce CelExpressionAttributes) InternalWithRef(ref terra.Reference) CelExpressionAttributes {
	return CelExpressionAttributes{ref: ref}
}

func (ce CelExpressionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ce.ref.InternalTokens()
}

func (ce CelExpressionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("description"))
}

func (ce CelExpressionAttributes) Expression() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("expression"))
}

func (ce CelExpressionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("location"))
}

func (ce CelExpressionAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("title"))
}

type PublishingOptionsAttributes struct {
	ref terra.Reference
}

func (po PublishingOptionsAttributes) InternalRef() (terra.Reference, error) {
	return po.ref, nil
}

func (po PublishingOptionsAttributes) InternalWithRef(ref terra.Reference) PublishingOptionsAttributes {
	return PublishingOptionsAttributes{ref: ref}
}

func (po PublishingOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return po.ref.InternalTokens()
}

func (po PublishingOptionsAttributes) PublishCaCert() terra.BoolValue {
	return terra.ReferenceAsBool(po.ref.Append("publish_ca_cert"))
}

func (po PublishingOptionsAttributes) PublishCrl() terra.BoolValue {
	return terra.ReferenceAsBool(po.ref.Append("publish_crl"))
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

type IssuancePolicyState struct {
	MaximumLifetime      string                      `json:"maximum_lifetime"`
	AllowedIssuanceModes []AllowedIssuanceModesState `json:"allowed_issuance_modes"`
	AllowedKeyTypes      []AllowedKeyTypesState      `json:"allowed_key_types"`
	BaselineValues       []BaselineValuesState       `json:"baseline_values"`
	IdentityConstraints  []IdentityConstraintsState  `json:"identity_constraints"`
}

type AllowedIssuanceModesState struct {
	AllowConfigBasedIssuance bool `json:"allow_config_based_issuance"`
	AllowCsrBasedIssuance    bool `json:"allow_csr_based_issuance"`
}

type AllowedKeyTypesState struct {
	EllipticCurve []EllipticCurveState `json:"elliptic_curve"`
	Rsa           []RsaState           `json:"rsa"`
}

type EllipticCurveState struct {
	SignatureAlgorithm string `json:"signature_algorithm"`
}

type RsaState struct {
	MaxModulusSize string `json:"max_modulus_size"`
	MinModulusSize string `json:"min_modulus_size"`
}

type BaselineValuesState struct {
	AiaOcspServers       []string                    `json:"aia_ocsp_servers"`
	AdditionalExtensions []AdditionalExtensionsState `json:"additional_extensions"`
	CaOptions            []CaOptionsState            `json:"ca_options"`
	KeyUsage             []KeyUsageState             `json:"key_usage"`
	NameConstraints      []NameConstraintsState      `json:"name_constraints"`
	PolicyIds            []PolicyIdsState            `json:"policy_ids"`
}

type AdditionalExtensionsState struct {
	Critical bool            `json:"critical"`
	Value    string          `json:"value"`
	ObjectId []ObjectIdState `json:"object_id"`
}

type ObjectIdState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type CaOptionsState struct {
	IsCa                    bool    `json:"is_ca"`
	MaxIssuerPathLength     float64 `json:"max_issuer_path_length"`
	NonCa                   bool    `json:"non_ca"`
	ZeroMaxIssuerPathLength bool    `json:"zero_max_issuer_path_length"`
}

type KeyUsageState struct {
	BaseKeyUsage             []BaseKeyUsageState             `json:"base_key_usage"`
	ExtendedKeyUsage         []ExtendedKeyUsageState         `json:"extended_key_usage"`
	UnknownExtendedKeyUsages []UnknownExtendedKeyUsagesState `json:"unknown_extended_key_usages"`
}

type BaseKeyUsageState struct {
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

type ExtendedKeyUsageState struct {
	ClientAuth      bool `json:"client_auth"`
	CodeSigning     bool `json:"code_signing"`
	EmailProtection bool `json:"email_protection"`
	OcspSigning     bool `json:"ocsp_signing"`
	ServerAuth      bool `json:"server_auth"`
	TimeStamping    bool `json:"time_stamping"`
}

type UnknownExtendedKeyUsagesState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type NameConstraintsState struct {
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

type PolicyIdsState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type IdentityConstraintsState struct {
	AllowSubjectAltNamesPassthrough bool                 `json:"allow_subject_alt_names_passthrough"`
	AllowSubjectPassthrough         bool                 `json:"allow_subject_passthrough"`
	CelExpression                   []CelExpressionState `json:"cel_expression"`
}

type CelExpressionState struct {
	Description string `json:"description"`
	Expression  string `json:"expression"`
	Location    string `json:"location"`
	Title       string `json:"title"`
}

type PublishingOptionsState struct {
	PublishCaCert bool `json:"publish_ca_cert"`
	PublishCrl    bool `json:"publish_crl"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
