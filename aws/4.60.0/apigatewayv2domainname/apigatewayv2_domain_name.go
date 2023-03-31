// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apigatewayv2domainname

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DomainNameConfiguration struct {
	// CertificateArn: string, required
	CertificateArn terra.StringValue `hcl:"certificate_arn,attr" validate:"required"`
	// EndpointType: string, required
	EndpointType terra.StringValue `hcl:"endpoint_type,attr" validate:"required"`
	// OwnershipVerificationCertificateArn: string, optional
	OwnershipVerificationCertificateArn terra.StringValue `hcl:"ownership_verification_certificate_arn,attr"`
	// SecurityPolicy: string, required
	SecurityPolicy terra.StringValue `hcl:"security_policy,attr" validate:"required"`
}

type MutualTlsAuthentication struct {
	// TruststoreUri: string, required
	TruststoreUri terra.StringValue `hcl:"truststore_uri,attr" validate:"required"`
	// TruststoreVersion: string, optional
	TruststoreVersion terra.StringValue `hcl:"truststore_version,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DomainNameConfigurationAttributes struct {
	ref terra.Reference
}

func (dnc DomainNameConfigurationAttributes) InternalRef() terra.Reference {
	return dnc.ref
}

func (dnc DomainNameConfigurationAttributes) InternalWithRef(ref terra.Reference) DomainNameConfigurationAttributes {
	return DomainNameConfigurationAttributes{ref: ref}
}

func (dnc DomainNameConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return dnc.ref.InternalTokens()
}

func (dnc DomainNameConfigurationAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(dnc.ref.Append("certificate_arn"))
}

func (dnc DomainNameConfigurationAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(dnc.ref.Append("endpoint_type"))
}

func (dnc DomainNameConfigurationAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(dnc.ref.Append("hosted_zone_id"))
}

func (dnc DomainNameConfigurationAttributes) OwnershipVerificationCertificateArn() terra.StringValue {
	return terra.ReferenceAsString(dnc.ref.Append("ownership_verification_certificate_arn"))
}

func (dnc DomainNameConfigurationAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(dnc.ref.Append("security_policy"))
}

func (dnc DomainNameConfigurationAttributes) TargetDomainName() terra.StringValue {
	return terra.ReferenceAsString(dnc.ref.Append("target_domain_name"))
}

type MutualTlsAuthenticationAttributes struct {
	ref terra.Reference
}

func (mta MutualTlsAuthenticationAttributes) InternalRef() terra.Reference {
	return mta.ref
}

func (mta MutualTlsAuthenticationAttributes) InternalWithRef(ref terra.Reference) MutualTlsAuthenticationAttributes {
	return MutualTlsAuthenticationAttributes{ref: ref}
}

func (mta MutualTlsAuthenticationAttributes) InternalTokens() hclwrite.Tokens {
	return mta.ref.InternalTokens()
}

func (mta MutualTlsAuthenticationAttributes) TruststoreUri() terra.StringValue {
	return terra.ReferenceAsString(mta.ref.Append("truststore_uri"))
}

func (mta MutualTlsAuthenticationAttributes) TruststoreVersion() terra.StringValue {
	return terra.ReferenceAsString(mta.ref.Append("truststore_version"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type DomainNameConfigurationState struct {
	CertificateArn                      string `json:"certificate_arn"`
	EndpointType                        string `json:"endpoint_type"`
	HostedZoneId                        string `json:"hosted_zone_id"`
	OwnershipVerificationCertificateArn string `json:"ownership_verification_certificate_arn"`
	SecurityPolicy                      string `json:"security_policy"`
	TargetDomainName                    string `json:"target_domain_name"`
}

type MutualTlsAuthenticationState struct {
	TruststoreUri     string `json:"truststore_uri"`
	TruststoreVersion string `json:"truststore_version"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Update string `json:"update"`
}
