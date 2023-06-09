// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataPrivatecaCertificateTemplateIamPolicy creates a new instance of [DataPrivatecaCertificateTemplateIamPolicy].
func NewDataPrivatecaCertificateTemplateIamPolicy(name string, args DataPrivatecaCertificateTemplateIamPolicyArgs) *DataPrivatecaCertificateTemplateIamPolicy {
	return &DataPrivatecaCertificateTemplateIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivatecaCertificateTemplateIamPolicy)(nil)

// DataPrivatecaCertificateTemplateIamPolicy represents the Terraform data resource google_privateca_certificate_template_iam_policy.
type DataPrivatecaCertificateTemplateIamPolicy struct {
	Name string
	Args DataPrivatecaCertificateTemplateIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataPrivatecaCertificateTemplateIamPolicy].
func (pctip *DataPrivatecaCertificateTemplateIamPolicy) DataSource() string {
	return "google_privateca_certificate_template_iam_policy"
}

// LocalName returns the local name for [DataPrivatecaCertificateTemplateIamPolicy].
func (pctip *DataPrivatecaCertificateTemplateIamPolicy) LocalName() string {
	return pctip.Name
}

// Configuration returns the configuration (args) for [DataPrivatecaCertificateTemplateIamPolicy].
func (pctip *DataPrivatecaCertificateTemplateIamPolicy) Configuration() interface{} {
	return pctip.Args
}

// Attributes returns the attributes for [DataPrivatecaCertificateTemplateIamPolicy].
func (pctip *DataPrivatecaCertificateTemplateIamPolicy) Attributes() dataPrivatecaCertificateTemplateIamPolicyAttributes {
	return dataPrivatecaCertificateTemplateIamPolicyAttributes{ref: terra.ReferenceDataResource(pctip)}
}

// DataPrivatecaCertificateTemplateIamPolicyArgs contains the configurations for google_privateca_certificate_template_iam_policy.
type DataPrivatecaCertificateTemplateIamPolicyArgs struct {
	// CertificateTemplate: string, required
	CertificateTemplate terra.StringValue `hcl:"certificate_template,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataPrivatecaCertificateTemplateIamPolicyAttributes struct {
	ref terra.Reference
}

// CertificateTemplate returns a reference to field certificate_template of google_privateca_certificate_template_iam_policy.
func (pctip dataPrivatecaCertificateTemplateIamPolicyAttributes) CertificateTemplate() terra.StringValue {
	return terra.ReferenceAsString(pctip.ref.Append("certificate_template"))
}

// Etag returns a reference to field etag of google_privateca_certificate_template_iam_policy.
func (pctip dataPrivatecaCertificateTemplateIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(pctip.ref.Append("etag"))
}

// Id returns a reference to field id of google_privateca_certificate_template_iam_policy.
func (pctip dataPrivatecaCertificateTemplateIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pctip.ref.Append("id"))
}

// Location returns a reference to field location of google_privateca_certificate_template_iam_policy.
func (pctip dataPrivatecaCertificateTemplateIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pctip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_privateca_certificate_template_iam_policy.
func (pctip dataPrivatecaCertificateTemplateIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(pctip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_privateca_certificate_template_iam_policy.
func (pctip dataPrivatecaCertificateTemplateIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pctip.ref.Append("project"))
}
