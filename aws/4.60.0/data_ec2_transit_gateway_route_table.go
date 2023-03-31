// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2transitgatewayroutetable "github.com/golingon/terraproviders/aws/4.60.0/dataec2transitgatewayroutetable"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2TransitGatewayRouteTable creates a new instance of [DataEc2TransitGatewayRouteTable].
func NewDataEc2TransitGatewayRouteTable(name string, args DataEc2TransitGatewayRouteTableArgs) *DataEc2TransitGatewayRouteTable {
	return &DataEc2TransitGatewayRouteTable{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2TransitGatewayRouteTable)(nil)

// DataEc2TransitGatewayRouteTable represents the Terraform data resource aws_ec2_transit_gateway_route_table.
type DataEc2TransitGatewayRouteTable struct {
	Name string
	Args DataEc2TransitGatewayRouteTableArgs
}

// DataSource returns the Terraform object type for [DataEc2TransitGatewayRouteTable].
func (etgrt *DataEc2TransitGatewayRouteTable) DataSource() string {
	return "aws_ec2_transit_gateway_route_table"
}

// LocalName returns the local name for [DataEc2TransitGatewayRouteTable].
func (etgrt *DataEc2TransitGatewayRouteTable) LocalName() string {
	return etgrt.Name
}

// Configuration returns the configuration (args) for [DataEc2TransitGatewayRouteTable].
func (etgrt *DataEc2TransitGatewayRouteTable) Configuration() interface{} {
	return etgrt.Args
}

// Attributes returns the attributes for [DataEc2TransitGatewayRouteTable].
func (etgrt *DataEc2TransitGatewayRouteTable) Attributes() dataEc2TransitGatewayRouteTableAttributes {
	return dataEc2TransitGatewayRouteTableAttributes{ref: terra.ReferenceDataResource(etgrt)}
}

// DataEc2TransitGatewayRouteTableArgs contains the configurations for aws_ec2_transit_gateway_route_table.
type DataEc2TransitGatewayRouteTableArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2transitgatewayroutetable.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2transitgatewayroutetable.Timeouts `hcl:"timeouts,block"`
}
type dataEc2TransitGatewayRouteTableAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_transit_gateway_route_table.
func (etgrt dataEc2TransitGatewayRouteTableAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(etgrt.ref.Append("arn"))
}

// DefaultAssociationRouteTable returns a reference to field default_association_route_table of aws_ec2_transit_gateway_route_table.
func (etgrt dataEc2TransitGatewayRouteTableAttributes) DefaultAssociationRouteTable() terra.BoolValue {
	return terra.ReferenceAsBool(etgrt.ref.Append("default_association_route_table"))
}

// DefaultPropagationRouteTable returns a reference to field default_propagation_route_table of aws_ec2_transit_gateway_route_table.
func (etgrt dataEc2TransitGatewayRouteTableAttributes) DefaultPropagationRouteTable() terra.BoolValue {
	return terra.ReferenceAsBool(etgrt.ref.Append("default_propagation_route_table"))
}

// Id returns a reference to field id of aws_ec2_transit_gateway_route_table.
func (etgrt dataEc2TransitGatewayRouteTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgrt.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_route_table.
func (etgrt dataEc2TransitGatewayRouteTableAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgrt.ref.Append("tags"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_ec2_transit_gateway_route_table.
func (etgrt dataEc2TransitGatewayRouteTableAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgrt.ref.Append("transit_gateway_id"))
}

func (etgrt dataEc2TransitGatewayRouteTableAttributes) Filter() terra.SetValue[dataec2transitgatewayroutetable.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2transitgatewayroutetable.FilterAttributes](etgrt.ref.Append("filter"))
}

func (etgrt dataEc2TransitGatewayRouteTableAttributes) Timeouts() dataec2transitgatewayroutetable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2transitgatewayroutetable.TimeoutsAttributes](etgrt.ref.Append("timeouts"))
}
