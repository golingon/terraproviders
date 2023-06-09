// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataIamServerCertificate creates a new instance of [DataIamServerCertificate].
func NewDataIamServerCertificate(name string, args DataIamServerCertificateArgs) *DataIamServerCertificate {
	return &DataIamServerCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamServerCertificate)(nil)

// DataIamServerCertificate represents the Terraform data resource aws_iam_server_certificate.
type DataIamServerCertificate struct {
	Name string
	Args DataIamServerCertificateArgs
}

// DataSource returns the Terraform object type for [DataIamServerCertificate].
func (isc *DataIamServerCertificate) DataSource() string {
	return "aws_iam_server_certificate"
}

// LocalName returns the local name for [DataIamServerCertificate].
func (isc *DataIamServerCertificate) LocalName() string {
	return isc.Name
}

// Configuration returns the configuration (args) for [DataIamServerCertificate].
func (isc *DataIamServerCertificate) Configuration() interface{} {
	return isc.Args
}

// Attributes returns the attributes for [DataIamServerCertificate].
func (isc *DataIamServerCertificate) Attributes() dataIamServerCertificateAttributes {
	return dataIamServerCertificateAttributes{ref: terra.ReferenceDataResource(isc)}
}

// DataIamServerCertificateArgs contains the configurations for aws_iam_server_certificate.
type DataIamServerCertificateArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Latest: bool, optional
	Latest terra.BoolValue `hcl:"latest,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// PathPrefix: string, optional
	PathPrefix terra.StringValue `hcl:"path_prefix,attr"`
}
type dataIamServerCertificateAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_server_certificate.
func (isc dataIamServerCertificateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(isc.ref.Append("arn"))
}

// CertificateBody returns a reference to field certificate_body of aws_iam_server_certificate.
func (isc dataIamServerCertificateAttributes) CertificateBody() terra.StringValue {
	return terra.ReferenceAsString(isc.ref.Append("certificate_body"))
}

// CertificateChain returns a reference to field certificate_chain of aws_iam_server_certificate.
func (isc dataIamServerCertificateAttributes) CertificateChain() terra.StringValue {
	return terra.ReferenceAsString(isc.ref.Append("certificate_chain"))
}

// ExpirationDate returns a reference to field expiration_date of aws_iam_server_certificate.
func (isc dataIamServerCertificateAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(isc.ref.Append("expiration_date"))
}

// Id returns a reference to field id of aws_iam_server_certificate.
func (isc dataIamServerCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(isc.ref.Append("id"))
}

// Latest returns a reference to field latest of aws_iam_server_certificate.
func (isc dataIamServerCertificateAttributes) Latest() terra.BoolValue {
	return terra.ReferenceAsBool(isc.ref.Append("latest"))
}

// Name returns a reference to field name of aws_iam_server_certificate.
func (isc dataIamServerCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(isc.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_iam_server_certificate.
func (isc dataIamServerCertificateAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(isc.ref.Append("name_prefix"))
}

// Path returns a reference to field path of aws_iam_server_certificate.
func (isc dataIamServerCertificateAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(isc.ref.Append("path"))
}

// PathPrefix returns a reference to field path_prefix of aws_iam_server_certificate.
func (isc dataIamServerCertificateAttributes) PathPrefix() terra.StringValue {
	return terra.ReferenceAsString(isc.ref.Append("path_prefix"))
}

// UploadDate returns a reference to field upload_date of aws_iam_server_certificate.
func (isc dataIamServerCertificateAttributes) UploadDate() terra.StringValue {
	return terra.ReferenceAsString(isc.ref.Append("upload_date"))
}
