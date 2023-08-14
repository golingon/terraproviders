// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2localgatewayroutetable "github.com/golingon/terraproviders/aws/5.10.0/dataec2localgatewayroutetable"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2LocalGatewayRouteTable creates a new instance of [DataEc2LocalGatewayRouteTable].
func NewDataEc2LocalGatewayRouteTable(name string, args DataEc2LocalGatewayRouteTableArgs) *DataEc2LocalGatewayRouteTable {
	return &DataEc2LocalGatewayRouteTable{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2LocalGatewayRouteTable)(nil)

// DataEc2LocalGatewayRouteTable represents the Terraform data resource aws_ec2_local_gateway_route_table.
type DataEc2LocalGatewayRouteTable struct {
	Name string
	Args DataEc2LocalGatewayRouteTableArgs
}

// DataSource returns the Terraform object type for [DataEc2LocalGatewayRouteTable].
func (elgrt *DataEc2LocalGatewayRouteTable) DataSource() string {
	return "aws_ec2_local_gateway_route_table"
}

// LocalName returns the local name for [DataEc2LocalGatewayRouteTable].
func (elgrt *DataEc2LocalGatewayRouteTable) LocalName() string {
	return elgrt.Name
}

// Configuration returns the configuration (args) for [DataEc2LocalGatewayRouteTable].
func (elgrt *DataEc2LocalGatewayRouteTable) Configuration() interface{} {
	return elgrt.Args
}

// Attributes returns the attributes for [DataEc2LocalGatewayRouteTable].
func (elgrt *DataEc2LocalGatewayRouteTable) Attributes() dataEc2LocalGatewayRouteTableAttributes {
	return dataEc2LocalGatewayRouteTableAttributes{ref: terra.ReferenceDataResource(elgrt)}
}

// DataEc2LocalGatewayRouteTableArgs contains the configurations for aws_ec2_local_gateway_route_table.
type DataEc2LocalGatewayRouteTableArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalGatewayId: string, optional
	LocalGatewayId terra.StringValue `hcl:"local_gateway_id,attr"`
	// LocalGatewayRouteTableId: string, optional
	LocalGatewayRouteTableId terra.StringValue `hcl:"local_gateway_route_table_id,attr"`
	// OutpostArn: string, optional
	OutpostArn terra.StringValue `hcl:"outpost_arn,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2localgatewayroutetable.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2localgatewayroutetable.Timeouts `hcl:"timeouts,block"`
}
type dataEc2LocalGatewayRouteTableAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_local_gateway_route_table.
func (elgrt dataEc2LocalGatewayRouteTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(elgrt.ref.Append("id"))
}

// LocalGatewayId returns a reference to field local_gateway_id of aws_ec2_local_gateway_route_table.
func (elgrt dataEc2LocalGatewayRouteTableAttributes) LocalGatewayId() terra.StringValue {
	return terra.ReferenceAsString(elgrt.ref.Append("local_gateway_id"))
}

// LocalGatewayRouteTableId returns a reference to field local_gateway_route_table_id of aws_ec2_local_gateway_route_table.
func (elgrt dataEc2LocalGatewayRouteTableAttributes) LocalGatewayRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(elgrt.ref.Append("local_gateway_route_table_id"))
}

// OutpostArn returns a reference to field outpost_arn of aws_ec2_local_gateway_route_table.
func (elgrt dataEc2LocalGatewayRouteTableAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(elgrt.ref.Append("outpost_arn"))
}

// State returns a reference to field state of aws_ec2_local_gateway_route_table.
func (elgrt dataEc2LocalGatewayRouteTableAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(elgrt.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_ec2_local_gateway_route_table.
func (elgrt dataEc2LocalGatewayRouteTableAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](elgrt.ref.Append("tags"))
}

func (elgrt dataEc2LocalGatewayRouteTableAttributes) Filter() terra.SetValue[dataec2localgatewayroutetable.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2localgatewayroutetable.FilterAttributes](elgrt.ref.Append("filter"))
}

func (elgrt dataEc2LocalGatewayRouteTableAttributes) Timeouts() dataec2localgatewayroutetable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2localgatewayroutetable.TimeoutsAttributes](elgrt.ref.Append("timeouts"))
}