// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2localgatewayvirtualinterfacegroup "github.com/golingon/terraproviders/aws/5.10.0/dataec2localgatewayvirtualinterfacegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2LocalGatewayVirtualInterfaceGroup creates a new instance of [DataEc2LocalGatewayVirtualInterfaceGroup].
func NewDataEc2LocalGatewayVirtualInterfaceGroup(name string, args DataEc2LocalGatewayVirtualInterfaceGroupArgs) *DataEc2LocalGatewayVirtualInterfaceGroup {
	return &DataEc2LocalGatewayVirtualInterfaceGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2LocalGatewayVirtualInterfaceGroup)(nil)

// DataEc2LocalGatewayVirtualInterfaceGroup represents the Terraform data resource aws_ec2_local_gateway_virtual_interface_group.
type DataEc2LocalGatewayVirtualInterfaceGroup struct {
	Name string
	Args DataEc2LocalGatewayVirtualInterfaceGroupArgs
}

// DataSource returns the Terraform object type for [DataEc2LocalGatewayVirtualInterfaceGroup].
func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroup) DataSource() string {
	return "aws_ec2_local_gateway_virtual_interface_group"
}

// LocalName returns the local name for [DataEc2LocalGatewayVirtualInterfaceGroup].
func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroup) LocalName() string {
	return elgvig.Name
}

// Configuration returns the configuration (args) for [DataEc2LocalGatewayVirtualInterfaceGroup].
func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroup) Configuration() interface{} {
	return elgvig.Args
}

// Attributes returns the attributes for [DataEc2LocalGatewayVirtualInterfaceGroup].
func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroup) Attributes() dataEc2LocalGatewayVirtualInterfaceGroupAttributes {
	return dataEc2LocalGatewayVirtualInterfaceGroupAttributes{ref: terra.ReferenceDataResource(elgvig)}
}

// DataEc2LocalGatewayVirtualInterfaceGroupArgs contains the configurations for aws_ec2_local_gateway_virtual_interface_group.
type DataEc2LocalGatewayVirtualInterfaceGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalGatewayId: string, optional
	LocalGatewayId terra.StringValue `hcl:"local_gateway_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2localgatewayvirtualinterfacegroup.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2localgatewayvirtualinterfacegroup.Timeouts `hcl:"timeouts,block"`
}
type dataEc2LocalGatewayVirtualInterfaceGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_local_gateway_virtual_interface_group.
func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(elgvig.ref.Append("id"))
}

// LocalGatewayId returns a reference to field local_gateway_id of aws_ec2_local_gateway_virtual_interface_group.
func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupAttributes) LocalGatewayId() terra.StringValue {
	return terra.ReferenceAsString(elgvig.ref.Append("local_gateway_id"))
}

// LocalGatewayVirtualInterfaceIds returns a reference to field local_gateway_virtual_interface_ids of aws_ec2_local_gateway_virtual_interface_group.
func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupAttributes) LocalGatewayVirtualInterfaceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](elgvig.ref.Append("local_gateway_virtual_interface_ids"))
}

// Tags returns a reference to field tags of aws_ec2_local_gateway_virtual_interface_group.
func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](elgvig.ref.Append("tags"))
}

func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupAttributes) Filter() terra.SetValue[dataec2localgatewayvirtualinterfacegroup.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2localgatewayvirtualinterfacegroup.FilterAttributes](elgvig.ref.Append("filter"))
}

func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupAttributes) Timeouts() dataec2localgatewayvirtualinterfacegroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2localgatewayvirtualinterfacegroup.TimeoutsAttributes](elgvig.ref.Append("timeouts"))
}
