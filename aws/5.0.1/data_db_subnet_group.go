// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDbSubnetGroup creates a new instance of [DataDbSubnetGroup].
func NewDataDbSubnetGroup(name string, args DataDbSubnetGroupArgs) *DataDbSubnetGroup {
	return &DataDbSubnetGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDbSubnetGroup)(nil)

// DataDbSubnetGroup represents the Terraform data resource aws_db_subnet_group.
type DataDbSubnetGroup struct {
	Name string
	Args DataDbSubnetGroupArgs
}

// DataSource returns the Terraform object type for [DataDbSubnetGroup].
func (dsg *DataDbSubnetGroup) DataSource() string {
	return "aws_db_subnet_group"
}

// LocalName returns the local name for [DataDbSubnetGroup].
func (dsg *DataDbSubnetGroup) LocalName() string {
	return dsg.Name
}

// Configuration returns the configuration (args) for [DataDbSubnetGroup].
func (dsg *DataDbSubnetGroup) Configuration() interface{} {
	return dsg.Args
}

// Attributes returns the attributes for [DataDbSubnetGroup].
func (dsg *DataDbSubnetGroup) Attributes() dataDbSubnetGroupAttributes {
	return dataDbSubnetGroupAttributes{ref: terra.ReferenceDataResource(dsg)}
}

// DataDbSubnetGroupArgs contains the configurations for aws_db_subnet_group.
type DataDbSubnetGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataDbSubnetGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_db_subnet_group.
func (dsg dataDbSubnetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_db_subnet_group.
func (dsg dataDbSubnetGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("description"))
}

// Id returns a reference to field id of aws_db_subnet_group.
func (dsg dataDbSubnetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("id"))
}

// Name returns a reference to field name of aws_db_subnet_group.
func (dsg dataDbSubnetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("name"))
}

// Status returns a reference to field status of aws_db_subnet_group.
func (dsg dataDbSubnetGroupAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("status"))
}

// SubnetIds returns a reference to field subnet_ids of aws_db_subnet_group.
func (dsg dataDbSubnetGroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dsg.ref.Append("subnet_ids"))
}

// SupportedNetworkTypes returns a reference to field supported_network_types of aws_db_subnet_group.
func (dsg dataDbSubnetGroupAttributes) SupportedNetworkTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dsg.ref.Append("supported_network_types"))
}

// VpcId returns a reference to field vpc_id of aws_db_subnet_group.
func (dsg dataDbSubnetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("vpc_id"))
}
