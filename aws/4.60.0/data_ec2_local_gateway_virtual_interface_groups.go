// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2localgatewayvirtualinterfacegroups "github.com/golingon/terraproviders/aws/4.60.0/dataec2localgatewayvirtualinterfacegroups"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2LocalGatewayVirtualInterfaceGroups creates a new instance of [DataEc2LocalGatewayVirtualInterfaceGroups].
func NewDataEc2LocalGatewayVirtualInterfaceGroups(name string, args DataEc2LocalGatewayVirtualInterfaceGroupsArgs) *DataEc2LocalGatewayVirtualInterfaceGroups {
	return &DataEc2LocalGatewayVirtualInterfaceGroups{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2LocalGatewayVirtualInterfaceGroups)(nil)

// DataEc2LocalGatewayVirtualInterfaceGroups represents the Terraform data resource aws_ec2_local_gateway_virtual_interface_groups.
type DataEc2LocalGatewayVirtualInterfaceGroups struct {
	Name string
	Args DataEc2LocalGatewayVirtualInterfaceGroupsArgs
}

// DataSource returns the Terraform object type for [DataEc2LocalGatewayVirtualInterfaceGroups].
func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroups) DataSource() string {
	return "aws_ec2_local_gateway_virtual_interface_groups"
}

// LocalName returns the local name for [DataEc2LocalGatewayVirtualInterfaceGroups].
func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroups) LocalName() string {
	return elgvig.Name
}

// Configuration returns the configuration (args) for [DataEc2LocalGatewayVirtualInterfaceGroups].
func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroups) Configuration() interface{} {
	return elgvig.Args
}

// Attributes returns the attributes for [DataEc2LocalGatewayVirtualInterfaceGroups].
func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroups) Attributes() dataEc2LocalGatewayVirtualInterfaceGroupsAttributes {
	return dataEc2LocalGatewayVirtualInterfaceGroupsAttributes{ref: terra.ReferenceDataResource(elgvig)}
}

// DataEc2LocalGatewayVirtualInterfaceGroupsArgs contains the configurations for aws_ec2_local_gateway_virtual_interface_groups.
type DataEc2LocalGatewayVirtualInterfaceGroupsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2localgatewayvirtualinterfacegroups.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2localgatewayvirtualinterfacegroups.Timeouts `hcl:"timeouts,block"`
}
type dataEc2LocalGatewayVirtualInterfaceGroupsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_local_gateway_virtual_interface_groups.
func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(elgvig.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ec2_local_gateway_virtual_interface_groups.
func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](elgvig.ref.Append("ids"))
}

// LocalGatewayVirtualInterfaceIds returns a reference to field local_gateway_virtual_interface_ids of aws_ec2_local_gateway_virtual_interface_groups.
func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) LocalGatewayVirtualInterfaceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](elgvig.ref.Append("local_gateway_virtual_interface_ids"))
}

// Tags returns a reference to field tags of aws_ec2_local_gateway_virtual_interface_groups.
func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](elgvig.ref.Append("tags"))
}

func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) Filter() terra.SetValue[dataec2localgatewayvirtualinterfacegroups.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2localgatewayvirtualinterfacegroups.FilterAttributes](elgvig.ref.Append("filter"))
}

func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) Timeouts() dataec2localgatewayvirtualinterfacegroups.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2localgatewayvirtualinterfacegroups.TimeoutsAttributes](elgvig.ref.Append("timeouts"))
}
