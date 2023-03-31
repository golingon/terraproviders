// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataacmpcacertificateauthority "github.com/golingon/terraproviders/aws/4.60.0/dataacmpcacertificateauthority"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAcmpcaCertificateAuthority creates a new instance of [DataAcmpcaCertificateAuthority].
func NewDataAcmpcaCertificateAuthority(name string, args DataAcmpcaCertificateAuthorityArgs) *DataAcmpcaCertificateAuthority {
	return &DataAcmpcaCertificateAuthority{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAcmpcaCertificateAuthority)(nil)

// DataAcmpcaCertificateAuthority represents the Terraform data resource aws_acmpca_certificate_authority.
type DataAcmpcaCertificateAuthority struct {
	Name string
	Args DataAcmpcaCertificateAuthorityArgs
}

// DataSource returns the Terraform object type for [DataAcmpcaCertificateAuthority].
func (aca *DataAcmpcaCertificateAuthority) DataSource() string {
	return "aws_acmpca_certificate_authority"
}

// LocalName returns the local name for [DataAcmpcaCertificateAuthority].
func (aca *DataAcmpcaCertificateAuthority) LocalName() string {
	return aca.Name
}

// Configuration returns the configuration (args) for [DataAcmpcaCertificateAuthority].
func (aca *DataAcmpcaCertificateAuthority) Configuration() interface{} {
	return aca.Args
}

// Attributes returns the attributes for [DataAcmpcaCertificateAuthority].
func (aca *DataAcmpcaCertificateAuthority) Attributes() dataAcmpcaCertificateAuthorityAttributes {
	return dataAcmpcaCertificateAuthorityAttributes{ref: terra.ReferenceDataResource(aca)}
}

// DataAcmpcaCertificateAuthorityArgs contains the configurations for aws_acmpca_certificate_authority.
type DataAcmpcaCertificateAuthorityArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// RevocationConfiguration: min=0
	RevocationConfiguration []dataacmpcacertificateauthority.RevocationConfiguration `hcl:"revocation_configuration,block" validate:"min=0"`
}
type dataAcmpcaCertificateAuthorityAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("arn"))
}

// Certificate returns a reference to field certificate of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("certificate"))
}

// CertificateChain returns a reference to field certificate_chain of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) CertificateChain() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("certificate_chain"))
}

// CertificateSigningRequest returns a reference to field certificate_signing_request of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) CertificateSigningRequest() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("certificate_signing_request"))
}

// Id returns a reference to field id of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("id"))
}

// NotAfter returns a reference to field not_after of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) NotAfter() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("not_after"))
}

// NotBefore returns a reference to field not_before of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) NotBefore() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("not_before"))
}

// Serial returns a reference to field serial of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) Serial() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("serial"))
}

// Status returns a reference to field status of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aca.ref.Append("tags"))
}

// Type returns a reference to field type of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("type"))
}

// UsageMode returns a reference to field usage_mode of aws_acmpca_certificate_authority.
func (aca dataAcmpcaCertificateAuthorityAttributes) UsageMode() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("usage_mode"))
}

func (aca dataAcmpcaCertificateAuthorityAttributes) RevocationConfiguration() terra.ListValue[dataacmpcacertificateauthority.RevocationConfigurationAttributes] {
	return terra.ReferenceAsList[dataacmpcacertificateauthority.RevocationConfigurationAttributes](aca.ref.Append("revocation_configuration"))
}
