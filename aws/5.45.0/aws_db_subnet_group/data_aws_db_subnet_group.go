// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_db_subnet_group

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_db_subnet_group.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (adsg *DataSource) DataSource() string {
	return "aws_db_subnet_group"
}

// LocalName returns the local name for [DataSource].
func (adsg *DataSource) LocalName() string {
	return adsg.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (adsg *DataSource) Configuration() interface{} {
	return adsg.Args
}

// Attributes returns the attributes for [DataSource].
func (adsg *DataSource) Attributes() dataAwsDbSubnetGroupAttributes {
	return dataAwsDbSubnetGroupAttributes{ref: terra.ReferenceDataSource(adsg)}
}

// DataArgs contains the configurations for aws_db_subnet_group.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type dataAwsDbSubnetGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_db_subnet_group.
func (adsg dataAwsDbSubnetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_db_subnet_group.
func (adsg dataAwsDbSubnetGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("description"))
}

// Id returns a reference to field id of aws_db_subnet_group.
func (adsg dataAwsDbSubnetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("id"))
}

// Name returns a reference to field name of aws_db_subnet_group.
func (adsg dataAwsDbSubnetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("name"))
}

// Status returns a reference to field status of aws_db_subnet_group.
func (adsg dataAwsDbSubnetGroupAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("status"))
}

// SubnetIds returns a reference to field subnet_ids of aws_db_subnet_group.
func (adsg dataAwsDbSubnetGroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adsg.ref.Append("subnet_ids"))
}

// SupportedNetworkTypes returns a reference to field supported_network_types of aws_db_subnet_group.
func (adsg dataAwsDbSubnetGroupAttributes) SupportedNetworkTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adsg.ref.Append("supported_network_types"))
}

// VpcId returns a reference to field vpc_id of aws_db_subnet_group.
func (adsg dataAwsDbSubnetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("vpc_id"))
}
