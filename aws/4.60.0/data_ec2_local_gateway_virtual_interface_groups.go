// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2localgatewayvirtualinterfacegroups "github.com/golingon/terraproviders/aws/4.60.0/dataec2localgatewayvirtualinterfacegroups"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataEc2LocalGatewayVirtualInterfaceGroups(name string, args DataEc2LocalGatewayVirtualInterfaceGroupsArgs) *DataEc2LocalGatewayVirtualInterfaceGroups {
	return &DataEc2LocalGatewayVirtualInterfaceGroups{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2LocalGatewayVirtualInterfaceGroups)(nil)

type DataEc2LocalGatewayVirtualInterfaceGroups struct {
	Name string
	Args DataEc2LocalGatewayVirtualInterfaceGroupsArgs
}

func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroups) DataSource() string {
	return "aws_ec2_local_gateway_virtual_interface_groups"
}

func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroups) LocalName() string {
	return elgvig.Name
}

func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroups) Configuration() interface{} {
	return elgvig.Args
}

func (elgvig *DataEc2LocalGatewayVirtualInterfaceGroups) Attributes() dataEc2LocalGatewayVirtualInterfaceGroupsAttributes {
	return dataEc2LocalGatewayVirtualInterfaceGroupsAttributes{ref: terra.ReferenceDataResource(elgvig)}
}

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

func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(elgvig.ref.Append("id"))
}

func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](elgvig.ref.Append("ids"))
}

func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) LocalGatewayVirtualInterfaceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](elgvig.ref.Append("local_gateway_virtual_interface_ids"))
}

func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](elgvig.ref.Append("tags"))
}

func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) Filter() terra.SetValue[dataec2localgatewayvirtualinterfacegroups.FilterAttributes] {
	return terra.ReferenceSet[dataec2localgatewayvirtualinterfacegroups.FilterAttributes](elgvig.ref.Append("filter"))
}

func (elgvig dataEc2LocalGatewayVirtualInterfaceGroupsAttributes) Timeouts() dataec2localgatewayvirtualinterfacegroups.TimeoutsAttributes {
	return terra.ReferenceSingle[dataec2localgatewayvirtualinterfacegroups.TimeoutsAttributes](elgvig.ref.Append("timeouts"))
}
