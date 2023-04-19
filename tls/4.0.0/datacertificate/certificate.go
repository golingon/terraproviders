// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacertificate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Certificates struct{}

type CertificatesAttributes struct {
	ref terra.Reference
}

func (c CertificatesAttributes) InternalRef() terra.Reference {
	return c.ref
}

func (c CertificatesAttributes) InternalWithRef(ref terra.Reference) CertificatesAttributes {
	return CertificatesAttributes{ref: ref}
}

func (c CertificatesAttributes) InternalTokens() hclwrite.Tokens {
	return c.ref.InternalTokens()
}

func (c CertificatesAttributes) CertPem() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("cert_pem"))
}

func (c CertificatesAttributes) IsCa() terra.BoolValue {
	return terra.ReferenceBool(c.ref.Append("is_ca"))
}

func (c CertificatesAttributes) Issuer() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("issuer"))
}

func (c CertificatesAttributes) NotAfter() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("not_after"))
}

func (c CertificatesAttributes) NotBefore() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("not_before"))
}

func (c CertificatesAttributes) PublicKeyAlgorithm() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("public_key_algorithm"))
}

func (c CertificatesAttributes) SerialNumber() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("serial_number"))
}

func (c CertificatesAttributes) Sha1Fingerprint() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("sha1_fingerprint"))
}

func (c CertificatesAttributes) SignatureAlgorithm() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("signature_algorithm"))
}

func (c CertificatesAttributes) Subject() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("subject"))
}

func (c CertificatesAttributes) Version() terra.NumberValue {
	return terra.ReferenceNumber(c.ref.Append("version"))
}

type CertificatesState struct {
	CertPem            string  `json:"cert_pem"`
	IsCa               bool    `json:"is_ca"`
	Issuer             string  `json:"issuer"`
	NotAfter           string  `json:"not_after"`
	NotBefore          string  `json:"not_before"`
	PublicKeyAlgorithm string  `json:"public_key_algorithm"`
	SerialNumber       string  `json:"serial_number"`
	Sha1Fingerprint    string  `json:"sha1_fingerprint"`
	SignatureAlgorithm string  `json:"signature_algorithm"`
	Subject            string  `json:"subject"`
	Version            float64 `json:"version"`
}