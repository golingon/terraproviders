// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataRdsCertificate creates a new instance of [DataRdsCertificate].
func NewDataRdsCertificate(name string, args DataRdsCertificateArgs) *DataRdsCertificate {
	return &DataRdsCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRdsCertificate)(nil)

// DataRdsCertificate represents the Terraform data resource aws_rds_certificate.
type DataRdsCertificate struct {
	Name string
	Args DataRdsCertificateArgs
}

// DataSource returns the Terraform object type for [DataRdsCertificate].
func (rc *DataRdsCertificate) DataSource() string {
	return "aws_rds_certificate"
}

// LocalName returns the local name for [DataRdsCertificate].
func (rc *DataRdsCertificate) LocalName() string {
	return rc.Name
}

// Configuration returns the configuration (args) for [DataRdsCertificate].
func (rc *DataRdsCertificate) Configuration() interface{} {
	return rc.Args
}

// Attributes returns the attributes for [DataRdsCertificate].
func (rc *DataRdsCertificate) Attributes() dataRdsCertificateAttributes {
	return dataRdsCertificateAttributes{ref: terra.ReferenceDataResource(rc)}
}

// DataRdsCertificateArgs contains the configurations for aws_rds_certificate.
type DataRdsCertificateArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LatestValidTill: bool, optional
	LatestValidTill terra.BoolValue `hcl:"latest_valid_till,attr"`
}
type dataRdsCertificateAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_rds_certificate.
func (rc dataRdsCertificateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("arn"))
}

// CertificateType returns a reference to field certificate_type of aws_rds_certificate.
func (rc dataRdsCertificateAttributes) CertificateType() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("certificate_type"))
}

// CustomerOverride returns a reference to field customer_override of aws_rds_certificate.
func (rc dataRdsCertificateAttributes) CustomerOverride() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("customer_override"))
}

// CustomerOverrideValidTill returns a reference to field customer_override_valid_till of aws_rds_certificate.
func (rc dataRdsCertificateAttributes) CustomerOverrideValidTill() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("customer_override_valid_till"))
}

// Id returns a reference to field id of aws_rds_certificate.
func (rc dataRdsCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("id"))
}

// LatestValidTill returns a reference to field latest_valid_till of aws_rds_certificate.
func (rc dataRdsCertificateAttributes) LatestValidTill() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("latest_valid_till"))
}

// Thumbprint returns a reference to field thumbprint of aws_rds_certificate.
func (rc dataRdsCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("thumbprint"))
}

// ValidFrom returns a reference to field valid_from of aws_rds_certificate.
func (rc dataRdsCertificateAttributes) ValidFrom() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("valid_from"))
}

// ValidTill returns a reference to field valid_till of aws_rds_certificate.
func (rc dataRdsCertificateAttributes) ValidTill() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("valid_till"))
}
