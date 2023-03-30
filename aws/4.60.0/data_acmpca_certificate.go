// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataAcmpcaCertificate(name string, args DataAcmpcaCertificateArgs) *DataAcmpcaCertificate {
	return &DataAcmpcaCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAcmpcaCertificate)(nil)

type DataAcmpcaCertificate struct {
	Name string
	Args DataAcmpcaCertificateArgs
}

func (ac *DataAcmpcaCertificate) DataSource() string {
	return "aws_acmpca_certificate"
}

func (ac *DataAcmpcaCertificate) LocalName() string {
	return ac.Name
}

func (ac *DataAcmpcaCertificate) Configuration() interface{} {
	return ac.Args
}

func (ac *DataAcmpcaCertificate) Attributes() dataAcmpcaCertificateAttributes {
	return dataAcmpcaCertificateAttributes{ref: terra.ReferenceDataResource(ac)}
}

type DataAcmpcaCertificateArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// CertificateAuthorityArn: string, required
	CertificateAuthorityArn terra.StringValue `hcl:"certificate_authority_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataAcmpcaCertificateAttributes struct {
	ref terra.Reference
}

func (ac dataAcmpcaCertificateAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("arn"))
}

func (ac dataAcmpcaCertificateAttributes) Certificate() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("certificate"))
}

func (ac dataAcmpcaCertificateAttributes) CertificateAuthorityArn() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("certificate_authority_arn"))
}

func (ac dataAcmpcaCertificateAttributes) CertificateChain() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("certificate_chain"))
}

func (ac dataAcmpcaCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("id"))
}
