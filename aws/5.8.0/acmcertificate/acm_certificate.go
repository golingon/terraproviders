// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package acmcertificate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DomainValidationOptions struct{}

type RenewalSummary struct{}

type Options struct {
	// CertificateTransparencyLoggingPreference: string, optional
	CertificateTransparencyLoggingPreference terra.StringValue `hcl:"certificate_transparency_logging_preference,attr"`
}

type ValidationOption struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// ValidationDomain: string, required
	ValidationDomain terra.StringValue `hcl:"validation_domain,attr" validate:"required"`
}

type DomainValidationOptionsAttributes struct {
	ref terra.Reference
}

func (dvo DomainValidationOptionsAttributes) InternalRef() (terra.Reference, error) {
	return dvo.ref, nil
}

func (dvo DomainValidationOptionsAttributes) InternalWithRef(ref terra.Reference) DomainValidationOptionsAttributes {
	return DomainValidationOptionsAttributes{ref: ref}
}

func (dvo DomainValidationOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dvo.ref.InternalTokens()
}

func (dvo DomainValidationOptionsAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(dvo.ref.Append("domain_name"))
}

func (dvo DomainValidationOptionsAttributes) ResourceRecordName() terra.StringValue {
	return terra.ReferenceAsString(dvo.ref.Append("resource_record_name"))
}

func (dvo DomainValidationOptionsAttributes) ResourceRecordType() terra.StringValue {
	return terra.ReferenceAsString(dvo.ref.Append("resource_record_type"))
}

func (dvo DomainValidationOptionsAttributes) ResourceRecordValue() terra.StringValue {
	return terra.ReferenceAsString(dvo.ref.Append("resource_record_value"))
}

type RenewalSummaryAttributes struct {
	ref terra.Reference
}

func (rs RenewalSummaryAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs RenewalSummaryAttributes) InternalWithRef(ref terra.Reference) RenewalSummaryAttributes {
	return RenewalSummaryAttributes{ref: ref}
}

func (rs RenewalSummaryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs RenewalSummaryAttributes) RenewalStatus() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("renewal_status"))
}

func (rs RenewalSummaryAttributes) RenewalStatusReason() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("renewal_status_reason"))
}

func (rs RenewalSummaryAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("updated_at"))
}

type OptionsAttributes struct {
	ref terra.Reference
}

func (o OptionsAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OptionsAttributes) InternalWithRef(ref terra.Reference) OptionsAttributes {
	return OptionsAttributes{ref: ref}
}

func (o OptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OptionsAttributes) CertificateTransparencyLoggingPreference() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("certificate_transparency_logging_preference"))
}

type ValidationOptionAttributes struct {
	ref terra.Reference
}

func (vo ValidationOptionAttributes) InternalRef() (terra.Reference, error) {
	return vo.ref, nil
}

func (vo ValidationOptionAttributes) InternalWithRef(ref terra.Reference) ValidationOptionAttributes {
	return ValidationOptionAttributes{ref: ref}
}

func (vo ValidationOptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vo.ref.InternalTokens()
}

func (vo ValidationOptionAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(vo.ref.Append("domain_name"))
}

func (vo ValidationOptionAttributes) ValidationDomain() terra.StringValue {
	return terra.ReferenceAsString(vo.ref.Append("validation_domain"))
}

type DomainValidationOptionsState struct {
	DomainName          string `json:"domain_name"`
	ResourceRecordName  string `json:"resource_record_name"`
	ResourceRecordType  string `json:"resource_record_type"`
	ResourceRecordValue string `json:"resource_record_value"`
}

type RenewalSummaryState struct {
	RenewalStatus       string `json:"renewal_status"`
	RenewalStatusReason string `json:"renewal_status_reason"`
	UpdatedAt           string `json:"updated_at"`
}

type OptionsState struct {
	CertificateTransparencyLoggingPreference string `json:"certificate_transparency_logging_preference"`
}

type ValidationOptionState struct {
	DomainName       string `json:"domain_name"`
	ValidationDomain string `json:"validation_domain"`
}
