// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataAcmCertificate creates a new instance of [DataAcmCertificate].
func NewDataAcmCertificate(name string, args DataAcmCertificateArgs) *DataAcmCertificate {
	return &DataAcmCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAcmCertificate)(nil)

// DataAcmCertificate represents the Terraform data resource aws_acm_certificate.
type DataAcmCertificate struct {
	Name string
	Args DataAcmCertificateArgs
}

// DataSource returns the Terraform object type for [DataAcmCertificate].
func (ac *DataAcmCertificate) DataSource() string {
	return "aws_acm_certificate"
}

// LocalName returns the local name for [DataAcmCertificate].
func (ac *DataAcmCertificate) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [DataAcmCertificate].
func (ac *DataAcmCertificate) Configuration() interface{} {
	return ac.Args
}

// Attributes returns the attributes for [DataAcmCertificate].
func (ac *DataAcmCertificate) Attributes() dataAcmCertificateAttributes {
	return dataAcmCertificateAttributes{ref: terra.ReferenceDataResource(ac)}
}

// DataAcmCertificateArgs contains the configurations for aws_acm_certificate.
type DataAcmCertificateArgs struct {
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyTypes: set of string, optional
	KeyTypes terra.SetValue[terra.StringValue] `hcl:"key_types,attr"`
	// MostRecent: bool, optional
	MostRecent terra.BoolValue `hcl:"most_recent,attr"`
	// Statuses: list of string, optional
	Statuses terra.ListValue[terra.StringValue] `hcl:"statuses,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Types: list of string, optional
	Types terra.ListValue[terra.StringValue] `hcl:"types,attr"`
}
type dataAcmCertificateAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_acm_certificate.
func (ac dataAcmCertificateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("arn"))
}

// Certificate returns a reference to field certificate of aws_acm_certificate.
func (ac dataAcmCertificateAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("certificate"))
}

// CertificateChain returns a reference to field certificate_chain of aws_acm_certificate.
func (ac dataAcmCertificateAttributes) CertificateChain() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("certificate_chain"))
}

// Domain returns a reference to field domain of aws_acm_certificate.
func (ac dataAcmCertificateAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("domain"))
}

// Id returns a reference to field id of aws_acm_certificate.
func (ac dataAcmCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// KeyTypes returns a reference to field key_types of aws_acm_certificate.
func (ac dataAcmCertificateAttributes) KeyTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ac.ref.Append("key_types"))
}

// MostRecent returns a reference to field most_recent of aws_acm_certificate.
func (ac dataAcmCertificateAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("most_recent"))
}

// Status returns a reference to field status of aws_acm_certificate.
func (ac dataAcmCertificateAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("status"))
}

// Statuses returns a reference to field statuses of aws_acm_certificate.
func (ac dataAcmCertificateAttributes) Statuses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ac.ref.Append("statuses"))
}

// Tags returns a reference to field tags of aws_acm_certificate.
func (ac dataAcmCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("tags"))
}

// Types returns a reference to field types of aws_acm_certificate.
func (ac dataAcmCertificateAttributes) Types() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ac.ref.Append("types"))
}
