// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_apigee_keystores_aliases_key_cert_file

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CertsInfo struct {
	// CertsInfoCertInfo: min=0
	CertInfo []CertsInfoCertInfo `hcl:"cert_info,block" validate:"min=0"`
}

type CertsInfoCertInfo struct {
	// BasicConstraints: string, optional
	BasicConstraints terra.StringValue `hcl:"basic_constraints,attr"`
	// ExpiryDate: string, optional
	ExpiryDate terra.StringValue `hcl:"expiry_date,attr"`
	// IsValid: string, optional
	IsValid terra.StringValue `hcl:"is_valid,attr"`
	// Issuer: string, optional
	Issuer terra.StringValue `hcl:"issuer,attr"`
	// PublicKey: string, optional
	PublicKey terra.StringValue `hcl:"public_key,attr"`
	// SerialNumber: string, optional
	SerialNumber terra.StringValue `hcl:"serial_number,attr"`
	// SigAlgName: string, optional
	SigAlgName terra.StringValue `hcl:"sig_alg_name,attr"`
	// Subject: string, optional
	Subject terra.StringValue `hcl:"subject,attr"`
	// SubjectAlternativeNames: list of string, optional
	SubjectAlternativeNames terra.ListValue[terra.StringValue] `hcl:"subject_alternative_names,attr"`
	// ValidFrom: string, optional
	ValidFrom terra.StringValue `hcl:"valid_from,attr"`
	// Version: number, optional
	Version terra.NumberValue `hcl:"version,attr"`
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

type CertsInfoAttributes struct {
	ref terra.Reference
}

func (ci CertsInfoAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci CertsInfoAttributes) InternalWithRef(ref terra.Reference) CertsInfoAttributes {
	return CertsInfoAttributes{ref: ref}
}

func (ci CertsInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ci.ref.InternalTokens()
}

func (ci CertsInfoAttributes) CertInfo() terra.ListValue[CertsInfoCertInfoAttributes] {
	return terra.ReferenceAsList[CertsInfoCertInfoAttributes](ci.ref.Append("cert_info"))
}

type CertsInfoCertInfoAttributes struct {
	ref terra.Reference
}

func (ci CertsInfoCertInfoAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci CertsInfoCertInfoAttributes) InternalWithRef(ref terra.Reference) CertsInfoCertInfoAttributes {
	return CertsInfoCertInfoAttributes{ref: ref}
}

func (ci CertsInfoCertInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ci.ref.InternalTokens()
}

func (ci CertsInfoCertInfoAttributes) BasicConstraints() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("basic_constraints"))
}

func (ci CertsInfoCertInfoAttributes) ExpiryDate() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("expiry_date"))
}

func (ci CertsInfoCertInfoAttributes) IsValid() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("is_valid"))
}

func (ci CertsInfoCertInfoAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("issuer"))
}

func (ci CertsInfoCertInfoAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("public_key"))
}

func (ci CertsInfoCertInfoAttributes) SerialNumber() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("serial_number"))
}

func (ci CertsInfoCertInfoAttributes) SigAlgName() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("sig_alg_name"))
}

func (ci CertsInfoCertInfoAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("subject"))
}

func (ci CertsInfoCertInfoAttributes) SubjectAlternativeNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ci.ref.Append("subject_alternative_names"))
}

func (ci CertsInfoCertInfoAttributes) ValidFrom() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("valid_from"))
}

func (ci CertsInfoCertInfoAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(ci.ref.Append("version"))
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

type CertsInfoState struct {
	CertInfo []CertsInfoCertInfoState `json:"cert_info"`
}

type CertsInfoCertInfoState struct {
	BasicConstraints        string   `json:"basic_constraints"`
	ExpiryDate              string   `json:"expiry_date"`
	IsValid                 string   `json:"is_valid"`
	Issuer                  string   `json:"issuer"`
	PublicKey               string   `json:"public_key"`
	SerialNumber            string   `json:"serial_number"`
	SigAlgName              string   `json:"sig_alg_name"`
	Subject                 string   `json:"subject"`
	SubjectAlternativeNames []string `json:"subject_alternative_names"`
	ValidFrom               string   `json:"valid_from"`
	Version                 float64  `json:"version"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
