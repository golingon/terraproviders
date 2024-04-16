// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_iam_account_alias

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_iam_account_alias.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aiaa *DataSource) DataSource() string {
	return "aws_iam_account_alias"
}

// LocalName returns the local name for [DataSource].
func (aiaa *DataSource) LocalName() string {
	return aiaa.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aiaa *DataSource) Configuration() interface{} {
	return aiaa.Args
}

// Attributes returns the attributes for [DataSource].
func (aiaa *DataSource) Attributes() dataAwsIamAccountAliasAttributes {
	return dataAwsIamAccountAliasAttributes{ref: terra.ReferenceDataSource(aiaa)}
}

// DataArgs contains the configurations for aws_iam_account_alias.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type dataAwsIamAccountAliasAttributes struct {
	ref terra.Reference
}

// AccountAlias returns a reference to field account_alias of aws_iam_account_alias.
func (aiaa dataAwsIamAccountAliasAttributes) AccountAlias() terra.StringValue {
	return terra.ReferenceAsString(aiaa.ref.Append("account_alias"))
}

// Id returns a reference to field id of aws_iam_account_alias.
func (aiaa dataAwsIamAccountAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aiaa.ref.Append("id"))
}
