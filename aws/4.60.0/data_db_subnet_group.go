// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataDbSubnetGroup(name string, args DataDbSubnetGroupArgs) *DataDbSubnetGroup {
	return &DataDbSubnetGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDbSubnetGroup)(nil)

type DataDbSubnetGroup struct {
	Name string
	Args DataDbSubnetGroupArgs
}

func (dsg *DataDbSubnetGroup) DataSource() string {
	return "aws_db_subnet_group"
}

func (dsg *DataDbSubnetGroup) LocalName() string {
	return dsg.Name
}

func (dsg *DataDbSubnetGroup) Configuration() interface{} {
	return dsg.Args
}

func (dsg *DataDbSubnetGroup) Attributes() dataDbSubnetGroupAttributes {
	return dataDbSubnetGroupAttributes{ref: terra.ReferenceDataResource(dsg)}
}

type DataDbSubnetGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataDbSubnetGroupAttributes struct {
	ref terra.Reference
}

func (dsg dataDbSubnetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(dsg.ref.Append("arn"))
}

func (dsg dataDbSubnetGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dsg.ref.Append("description"))
}

func (dsg dataDbSubnetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dsg.ref.Append("id"))
}

func (dsg dataDbSubnetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dsg.ref.Append("name"))
}

func (dsg dataDbSubnetGroupAttributes) Status() terra.StringValue {
	return terra.ReferenceString(dsg.ref.Append("status"))
}

func (dsg dataDbSubnetGroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](dsg.ref.Append("subnet_ids"))
}

func (dsg dataDbSubnetGroupAttributes) SupportedNetworkTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](dsg.ref.Append("supported_network_types"))
}

func (dsg dataDbSubnetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(dsg.ref.Append("vpc_id"))
}
