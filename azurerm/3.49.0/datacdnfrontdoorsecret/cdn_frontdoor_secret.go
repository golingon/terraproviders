// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacdnfrontdoorsecret

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Secret struct {
	// CustomerCertificate: min=0
	CustomerCertificate []CustomerCertificate `hcl:"customer_certificate,block" validate:"min=0"`
}

type CustomerCertificate struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type SecretAttributes struct {
	ref terra.Reference
}

func (s SecretAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s SecretAttributes) InternalWithRef(ref terra.Reference) SecretAttributes {
	return SecretAttributes{ref: ref}
}

func (s SecretAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SecretAttributes) CustomerCertificate() terra.ListValue[CustomerCertificateAttributes] {
	return terra.ReferenceList[CustomerCertificateAttributes](s.ref.Append("customer_certificate"))
}

type CustomerCertificateAttributes struct {
	ref terra.Reference
}

func (cc CustomerCertificateAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc CustomerCertificateAttributes) InternalWithRef(ref terra.Reference) CustomerCertificateAttributes {
	return CustomerCertificateAttributes{ref: ref}
}

func (cc CustomerCertificateAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc CustomerCertificateAttributes) KeyVaultCertificateId() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("key_vault_certificate_id"))
}

func (cc CustomerCertificateAttributes) SubjectAlternativeNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](cc.ref.Append("subject_alternative_names"))
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

type SecretState struct {
	CustomerCertificate []CustomerCertificateState `json:"customer_certificate"`
}

type CustomerCertificateState struct {
	KeyVaultCertificateId   string   `json:"key_vault_certificate_id"`
	SubjectAlternativeNames []string `json:"subject_alternative_names"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
