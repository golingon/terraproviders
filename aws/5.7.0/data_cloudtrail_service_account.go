// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataCloudtrailServiceAccount creates a new instance of [DataCloudtrailServiceAccount].
func NewDataCloudtrailServiceAccount(name string, args DataCloudtrailServiceAccountArgs) *DataCloudtrailServiceAccount {
	return &DataCloudtrailServiceAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudtrailServiceAccount)(nil)

// DataCloudtrailServiceAccount represents the Terraform data resource aws_cloudtrail_service_account.
type DataCloudtrailServiceAccount struct {
	Name string
	Args DataCloudtrailServiceAccountArgs
}

// DataSource returns the Terraform object type for [DataCloudtrailServiceAccount].
func (csa *DataCloudtrailServiceAccount) DataSource() string {
	return "aws_cloudtrail_service_account"
}

// LocalName returns the local name for [DataCloudtrailServiceAccount].
func (csa *DataCloudtrailServiceAccount) LocalName() string {
	return csa.Name
}

// Configuration returns the configuration (args) for [DataCloudtrailServiceAccount].
func (csa *DataCloudtrailServiceAccount) Configuration() interface{} {
	return csa.Args
}

// Attributes returns the attributes for [DataCloudtrailServiceAccount].
func (csa *DataCloudtrailServiceAccount) Attributes() dataCloudtrailServiceAccountAttributes {
	return dataCloudtrailServiceAccountAttributes{ref: terra.ReferenceDataResource(csa)}
}

// DataCloudtrailServiceAccountArgs contains the configurations for aws_cloudtrail_service_account.
type DataCloudtrailServiceAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataCloudtrailServiceAccountAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudtrail_service_account.
func (csa dataCloudtrailServiceAccountAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("arn"))
}

// Id returns a reference to field id of aws_cloudtrail_service_account.
func (csa dataCloudtrailServiceAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("id"))
}

// Region returns a reference to field region of aws_cloudtrail_service_account.
func (csa dataCloudtrailServiceAccountAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("region"))
}
