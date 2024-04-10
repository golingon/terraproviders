// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataSfnActivity creates a new instance of [DataSfnActivity].
func NewDataSfnActivity(name string, args DataSfnActivityArgs) *DataSfnActivity {
	return &DataSfnActivity{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSfnActivity)(nil)

// DataSfnActivity represents the Terraform data resource aws_sfn_activity.
type DataSfnActivity struct {
	Name string
	Args DataSfnActivityArgs
}

// DataSource returns the Terraform object type for [DataSfnActivity].
func (sa *DataSfnActivity) DataSource() string {
	return "aws_sfn_activity"
}

// LocalName returns the local name for [DataSfnActivity].
func (sa *DataSfnActivity) LocalName() string {
	return sa.Name
}

// Configuration returns the configuration (args) for [DataSfnActivity].
func (sa *DataSfnActivity) Configuration() interface{} {
	return sa.Args
}

// Attributes returns the attributes for [DataSfnActivity].
func (sa *DataSfnActivity) Attributes() dataSfnActivityAttributes {
	return dataSfnActivityAttributes{ref: terra.ReferenceDataResource(sa)}
}

// DataSfnActivityArgs contains the configurations for aws_sfn_activity.
type DataSfnActivityArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}
type dataSfnActivityAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sfn_activity.
func (sa dataSfnActivityAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("arn"))
}

// CreationDate returns a reference to field creation_date of aws_sfn_activity.
func (sa dataSfnActivityAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("creation_date"))
}

// Id returns a reference to field id of aws_sfn_activity.
func (sa dataSfnActivityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("id"))
}

// Name returns a reference to field name of aws_sfn_activity.
func (sa dataSfnActivityAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
}
