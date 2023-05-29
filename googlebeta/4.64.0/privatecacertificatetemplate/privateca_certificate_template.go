// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package privatecacertificatetemplate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

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
	// Expression: string, optional
	Expression terra.StringValue `hcl:"expression,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Title: string, optional
	Title terra.StringValue `hcl:"title,attr"`
}

type PassthroughExtensions struct {
	// KnownExtensions: list of string, optional
	KnownExtensions terra.ListValue[terra.StringValue] `hcl:"known_extensions,attr"`
	// PassthroughExtensionsAdditionalExtensions: min=0
	AdditionalExtensions []PassthroughExtensionsAdditionalExtensions `hcl:"additional_extensions,block" validate:"min=0"`
}

type PassthroughExtensionsAdditionalExtensions struct {
	// ObjectIdPath: list of number, required
	ObjectIdPath terra.ListValue[terra.NumberValue] `hcl:"object_id_path,attr" validate:"required"`
}

type PredefinedValues struct {
	// AiaOcspServers: list of string, optional
	AiaOcspServers terra.ListValue[terra.StringValue] `hcl:"aia_ocsp_servers,attr"`
	// PredefinedValuesAdditionalExtensions: min=0
	AdditionalExtensions []PredefinedValuesAdditionalExtensions `hcl:"additional_extensions,block" validate:"min=0"`
	// CaOptions: optional
	CaOptions *CaOptions `hcl:"ca_options,block"`
	// KeyUsage: optional
	KeyUsage *KeyUsage `hcl:"key_usage,block"`
	// PolicyIds: min=0
	PolicyIds []PolicyIds `hcl:"policy_ids,block" validate:"min=0"`
}

type PredefinedValuesAdditionalExtensions struct {
	// Critical: bool, optional
	Critical terra.BoolValue `hcl:"critical,attr"`
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
}

type KeyUsage struct {
	// BaseKeyUsage: optional
	BaseKeyUsage *BaseKeyUsage `hcl:"base_key_usage,block"`
	// ExtendedKeyUsage: optional
	ExtendedKeyUsage *ExtendedKeyUsage `hcl:"extended_key_usage,block"`
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

type PolicyIds struct {
	// ObjectIdPath: list of number, required
	ObjectIdPath terra.ListValue[terra.NumberValue] `hcl:"object_id_path,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
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

type PassthroughExtensionsAttributes struct {
	ref terra.Reference
}

func (pe PassthroughExtensionsAttributes) InternalRef() (terra.Reference, error) {
	return pe.ref, nil
}

func (pe PassthroughExtensionsAttributes) InternalWithRef(ref terra.Reference) PassthroughExtensionsAttributes {
	return PassthroughExtensionsAttributes{ref: ref}
}

func (pe PassthroughExtensionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pe.ref.InternalTokens()
}

func (pe PassthroughExtensionsAttributes) KnownExtensions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pe.ref.Append("known_extensions"))
}

func (pe PassthroughExtensionsAttributes) AdditionalExtensions() terra.ListValue[PassthroughExtensionsAdditionalExtensionsAttributes] {
	return terra.ReferenceAsList[PassthroughExtensionsAdditionalExtensionsAttributes](pe.ref.Append("additional_extensions"))
}

type PassthroughExtensionsAdditionalExtensionsAttributes struct {
	ref terra.Reference
}

func (ae PassthroughExtensionsAdditionalExtensionsAttributes) InternalRef() (terra.Reference, error) {
	return ae.ref, nil
}

func (ae PassthroughExtensionsAdditionalExtensionsAttributes) InternalWithRef(ref terra.Reference) PassthroughExtensionsAdditionalExtensionsAttributes {
	return PassthroughExtensionsAdditionalExtensionsAttributes{ref: ref}
}

func (ae PassthroughExtensionsAdditionalExtensionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ae.ref.InternalTokens()
}

func (ae PassthroughExtensionsAdditionalExtensionsAttributes) ObjectIdPath() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](ae.ref.Append("object_id_path"))
}

type PredefinedValuesAttributes struct {
	ref terra.Reference
}

func (pv PredefinedValuesAttributes) InternalRef() (terra.Reference, error) {
	return pv.ref, nil
}

func (pv PredefinedValuesAttributes) InternalWithRef(ref terra.Reference) PredefinedValuesAttributes {
	return PredefinedValuesAttributes{ref: ref}
}

func (pv PredefinedValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pv.ref.InternalTokens()
}

func (pv PredefinedValuesAttributes) AiaOcspServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pv.ref.Append("aia_ocsp_servers"))
}

func (pv PredefinedValuesAttributes) AdditionalExtensions() terra.ListValue[PredefinedValuesAdditionalExtensionsAttributes] {
	return terra.ReferenceAsList[PredefinedValuesAdditionalExtensionsAttributes](pv.ref.Append("additional_extensions"))
}

func (pv PredefinedValuesAttributes) CaOptions() terra.ListValue[CaOptionsAttributes] {
	return terra.ReferenceAsList[CaOptionsAttributes](pv.ref.Append("ca_options"))
}

func (pv PredefinedValuesAttributes) KeyUsage() terra.ListValue[KeyUsageAttributes] {
	return terra.ReferenceAsList[KeyUsageAttributes](pv.ref.Append("key_usage"))
}

func (pv PredefinedValuesAttributes) PolicyIds() terra.ListValue[PolicyIdsAttributes] {
	return terra.ReferenceAsList[PolicyIdsAttributes](pv.ref.Append("policy_ids"))
}

type PredefinedValuesAdditionalExtensionsAttributes struct {
	ref terra.Reference
}

func (ae PredefinedValuesAdditionalExtensionsAttributes) InternalRef() (terra.Reference, error) {
	return ae.ref, nil
}

func (ae PredefinedValuesAdditionalExtensionsAttributes) InternalWithRef(ref terra.Reference) PredefinedValuesAdditionalExtensionsAttributes {
	return PredefinedValuesAdditionalExtensionsAttributes{ref: ref}
}

func (ae PredefinedValuesAdditionalExtensionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ae.ref.InternalTokens()
}

func (ae PredefinedValuesAdditionalExtensionsAttributes) Critical() terra.BoolValue {
	return terra.ReferenceAsBool(ae.ref.Append("critical"))
}

func (ae PredefinedValuesAdditionalExtensionsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("value"))
}

func (ae PredefinedValuesAdditionalExtensionsAttributes) ObjectId() terra.ListValue[ObjectIdAttributes] {
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

type PassthroughExtensionsState struct {
	KnownExtensions      []string                                         `json:"known_extensions"`
	AdditionalExtensions []PassthroughExtensionsAdditionalExtensionsState `json:"additional_extensions"`
}

type PassthroughExtensionsAdditionalExtensionsState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type PredefinedValuesState struct {
	AiaOcspServers       []string                                    `json:"aia_ocsp_servers"`
	AdditionalExtensions []PredefinedValuesAdditionalExtensionsState `json:"additional_extensions"`
	CaOptions            []CaOptionsState                            `json:"ca_options"`
	KeyUsage             []KeyUsageState                             `json:"key_usage"`
	PolicyIds            []PolicyIdsState                            `json:"policy_ids"`
}

type PredefinedValuesAdditionalExtensionsState struct {
	Critical bool            `json:"critical"`
	Value    string          `json:"value"`
	ObjectId []ObjectIdState `json:"object_id"`
}

type ObjectIdState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type CaOptionsState struct {
	IsCa                bool    `json:"is_ca"`
	MaxIssuerPathLength float64 `json:"max_issuer_path_length"`
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

type PolicyIdsState struct {
	ObjectIdPath []float64 `json:"object_id_path"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}