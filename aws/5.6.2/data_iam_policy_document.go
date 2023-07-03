// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataiampolicydocument "github.com/golingon/terraproviders/aws/5.6.2/dataiampolicydocument"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIamPolicyDocument creates a new instance of [DataIamPolicyDocument].
func NewDataIamPolicyDocument(name string, args DataIamPolicyDocumentArgs) *DataIamPolicyDocument {
	return &DataIamPolicyDocument{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamPolicyDocument)(nil)

// DataIamPolicyDocument represents the Terraform data resource aws_iam_policy_document.
type DataIamPolicyDocument struct {
	Name string
	Args DataIamPolicyDocumentArgs
}

// DataSource returns the Terraform object type for [DataIamPolicyDocument].
func (ipd *DataIamPolicyDocument) DataSource() string {
	return "aws_iam_policy_document"
}

// LocalName returns the local name for [DataIamPolicyDocument].
func (ipd *DataIamPolicyDocument) LocalName() string {
	return ipd.Name
}

// Configuration returns the configuration (args) for [DataIamPolicyDocument].
func (ipd *DataIamPolicyDocument) Configuration() interface{} {
	return ipd.Args
}

// Attributes returns the attributes for [DataIamPolicyDocument].
func (ipd *DataIamPolicyDocument) Attributes() dataIamPolicyDocumentAttributes {
	return dataIamPolicyDocumentAttributes{ref: terra.ReferenceDataResource(ipd)}
}

// DataIamPolicyDocumentArgs contains the configurations for aws_iam_policy_document.
type DataIamPolicyDocumentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OverridePolicyDocuments: list of string, optional
	OverridePolicyDocuments terra.ListValue[terra.StringValue] `hcl:"override_policy_documents,attr"`
	// PolicyId: string, optional
	PolicyId terra.StringValue `hcl:"policy_id,attr"`
	// SourcePolicyDocuments: list of string, optional
	SourcePolicyDocuments terra.ListValue[terra.StringValue] `hcl:"source_policy_documents,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// Statement: min=0
	Statement []dataiampolicydocument.Statement `hcl:"statement,block" validate:"min=0"`
}
type dataIamPolicyDocumentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_iam_policy_document.
func (ipd dataIamPolicyDocumentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ipd.ref.Append("id"))
}

// Json returns a reference to field json of aws_iam_policy_document.
func (ipd dataIamPolicyDocumentAttributes) Json() terra.StringValue {
	return terra.ReferenceAsString(ipd.ref.Append("json"))
}

// OverridePolicyDocuments returns a reference to field override_policy_documents of aws_iam_policy_document.
func (ipd dataIamPolicyDocumentAttributes) OverridePolicyDocuments() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ipd.ref.Append("override_policy_documents"))
}

// PolicyId returns a reference to field policy_id of aws_iam_policy_document.
func (ipd dataIamPolicyDocumentAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(ipd.ref.Append("policy_id"))
}

// SourcePolicyDocuments returns a reference to field source_policy_documents of aws_iam_policy_document.
func (ipd dataIamPolicyDocumentAttributes) SourcePolicyDocuments() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ipd.ref.Append("source_policy_documents"))
}

// Version returns a reference to field version of aws_iam_policy_document.
func (ipd dataIamPolicyDocumentAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ipd.ref.Append("version"))
}

func (ipd dataIamPolicyDocumentAttributes) Statement() terra.ListValue[dataiampolicydocument.StatementAttributes] {
	return terra.ReferenceAsList[dataiampolicydocument.StatementAttributes](ipd.ref.Append("statement"))
}
