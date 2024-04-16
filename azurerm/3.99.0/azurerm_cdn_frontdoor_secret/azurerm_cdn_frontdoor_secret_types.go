// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_cdn_frontdoor_secret

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Secret struct {
	// SecretCustomerCertificate: min=1
	CustomerCertificate []SecretCustomerCertificate `hcl:"customer_certificate,block" validate:"min=1"`
}

type SecretCustomerCertificate struct {
	// KeyVaultCertificateId: string, required
	KeyVaultCertificateId terra.StringValue `hcl:"key_vault_certificate_id,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type SecretAttributes struct {
	ref terra.Reference
}

func (s SecretAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SecretAttributes) InternalWithRef(ref terra.Reference) SecretAttributes {
	return SecretAttributes{ref: ref}
}

func (s SecretAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SecretAttributes) CustomerCertificate() terra.ListValue[SecretCustomerCertificateAttributes] {
	return terra.ReferenceAsList[SecretCustomerCertificateAttributes](s.ref.Append("customer_certificate"))
}

type SecretCustomerCertificateAttributes struct {
	ref terra.Reference
}

func (cc SecretCustomerCertificateAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc SecretCustomerCertificateAttributes) InternalWithRef(ref terra.Reference) SecretCustomerCertificateAttributes {
	return SecretCustomerCertificateAttributes{ref: ref}
}

func (cc SecretCustomerCertificateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc SecretCustomerCertificateAttributes) KeyVaultCertificateId() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("key_vault_certificate_id"))
}

func (cc SecretCustomerCertificateAttributes) SubjectAlternativeNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("subject_alternative_names"))
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

type SecretState struct {
	CustomerCertificate []SecretCustomerCertificateState `json:"customer_certificate"`
}

type SecretCustomerCertificateState struct {
	KeyVaultCertificateId   string   `json:"key_vault_certificate_id"`
	SubjectAlternativeNames []string `json:"subject_alternative_names"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
}
