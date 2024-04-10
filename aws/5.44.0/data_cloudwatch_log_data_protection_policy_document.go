// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	datacloudwatchlogdataprotectionpolicydocument "github.com/golingon/terraproviders/aws/5.44.0/datacloudwatchlogdataprotectionpolicydocument"
)

// NewDataCloudwatchLogDataProtectionPolicyDocument creates a new instance of [DataCloudwatchLogDataProtectionPolicyDocument].
func NewDataCloudwatchLogDataProtectionPolicyDocument(name string, args DataCloudwatchLogDataProtectionPolicyDocumentArgs) *DataCloudwatchLogDataProtectionPolicyDocument {
	return &DataCloudwatchLogDataProtectionPolicyDocument{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudwatchLogDataProtectionPolicyDocument)(nil)

// DataCloudwatchLogDataProtectionPolicyDocument represents the Terraform data resource aws_cloudwatch_log_data_protection_policy_document.
type DataCloudwatchLogDataProtectionPolicyDocument struct {
	Name string
	Args DataCloudwatchLogDataProtectionPolicyDocumentArgs
}

// DataSource returns the Terraform object type for [DataCloudwatchLogDataProtectionPolicyDocument].
func (cldppd *DataCloudwatchLogDataProtectionPolicyDocument) DataSource() string {
	return "aws_cloudwatch_log_data_protection_policy_document"
}

// LocalName returns the local name for [DataCloudwatchLogDataProtectionPolicyDocument].
func (cldppd *DataCloudwatchLogDataProtectionPolicyDocument) LocalName() string {
	return cldppd.Name
}

// Configuration returns the configuration (args) for [DataCloudwatchLogDataProtectionPolicyDocument].
func (cldppd *DataCloudwatchLogDataProtectionPolicyDocument) Configuration() interface{} {
	return cldppd.Args
}

// Attributes returns the attributes for [DataCloudwatchLogDataProtectionPolicyDocument].
func (cldppd *DataCloudwatchLogDataProtectionPolicyDocument) Attributes() dataCloudwatchLogDataProtectionPolicyDocumentAttributes {
	return dataCloudwatchLogDataProtectionPolicyDocumentAttributes{ref: terra.ReferenceDataResource(cldppd)}
}

// DataCloudwatchLogDataProtectionPolicyDocumentArgs contains the configurations for aws_cloudwatch_log_data_protection_policy_document.
type DataCloudwatchLogDataProtectionPolicyDocumentArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// Statement: min=2,max=2
	Statement []datacloudwatchlogdataprotectionpolicydocument.Statement `hcl:"statement,block" validate:"min=2,max=2"`
}
type dataCloudwatchLogDataProtectionPolicyDocumentAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_cloudwatch_log_data_protection_policy_document.
func (cldppd dataCloudwatchLogDataProtectionPolicyDocumentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cldppd.ref.Append("description"))
}

// Id returns a reference to field id of aws_cloudwatch_log_data_protection_policy_document.
func (cldppd dataCloudwatchLogDataProtectionPolicyDocumentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cldppd.ref.Append("id"))
}

// Json returns a reference to field json of aws_cloudwatch_log_data_protection_policy_document.
func (cldppd dataCloudwatchLogDataProtectionPolicyDocumentAttributes) Json() terra.StringValue {
	return terra.ReferenceAsString(cldppd.ref.Append("json"))
}

// Name returns a reference to field name of aws_cloudwatch_log_data_protection_policy_document.
func (cldppd dataCloudwatchLogDataProtectionPolicyDocumentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cldppd.ref.Append("name"))
}

// Version returns a reference to field version of aws_cloudwatch_log_data_protection_policy_document.
func (cldppd dataCloudwatchLogDataProtectionPolicyDocumentAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(cldppd.ref.Append("version"))
}

func (cldppd dataCloudwatchLogDataProtectionPolicyDocumentAttributes) Statement() terra.ListValue[datacloudwatchlogdataprotectionpolicydocument.StatementAttributes] {
	return terra.ReferenceAsList[datacloudwatchlogdataprotectionpolicydocument.StatementAttributes](cldppd.ref.Append("statement"))
}
