// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2transitgatewayroutetables "github.com/golingon/terraproviders/aws/4.60.0/dataec2transitgatewayroutetables"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2TransitGatewayRouteTables creates a new instance of [DataEc2TransitGatewayRouteTables].
func NewDataEc2TransitGatewayRouteTables(name string, args DataEc2TransitGatewayRouteTablesArgs) *DataEc2TransitGatewayRouteTables {
	return &DataEc2TransitGatewayRouteTables{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2TransitGatewayRouteTables)(nil)

// DataEc2TransitGatewayRouteTables represents the Terraform data resource aws_ec2_transit_gateway_route_tables.
type DataEc2TransitGatewayRouteTables struct {
	Name string
	Args DataEc2TransitGatewayRouteTablesArgs
}

// DataSource returns the Terraform object type for [DataEc2TransitGatewayRouteTables].
func (etgrt *DataEc2TransitGatewayRouteTables) DataSource() string {
	return "aws_ec2_transit_gateway_route_tables"
}

// LocalName returns the local name for [DataEc2TransitGatewayRouteTables].
func (etgrt *DataEc2TransitGatewayRouteTables) LocalName() string {
	return etgrt.Name
}

// Configuration returns the configuration (args) for [DataEc2TransitGatewayRouteTables].
func (etgrt *DataEc2TransitGatewayRouteTables) Configuration() interface{} {
	return etgrt.Args
}

// Attributes returns the attributes for [DataEc2TransitGatewayRouteTables].
func (etgrt *DataEc2TransitGatewayRouteTables) Attributes() dataEc2TransitGatewayRouteTablesAttributes {
	return dataEc2TransitGatewayRouteTablesAttributes{ref: terra.ReferenceDataResource(etgrt)}
}

// DataEc2TransitGatewayRouteTablesArgs contains the configurations for aws_ec2_transit_gateway_route_tables.
type DataEc2TransitGatewayRouteTablesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2transitgatewayroutetables.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2transitgatewayroutetables.Timeouts `hcl:"timeouts,block"`
}
type dataEc2TransitGatewayRouteTablesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_transit_gateway_route_tables.
func (etgrt dataEc2TransitGatewayRouteTablesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgrt.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ec2_transit_gateway_route_tables.
func (etgrt dataEc2TransitGatewayRouteTablesAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](etgrt.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_route_tables.
func (etgrt dataEc2TransitGatewayRouteTablesAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgrt.ref.Append("tags"))
}

func (etgrt dataEc2TransitGatewayRouteTablesAttributes) Filter() terra.SetValue[dataec2transitgatewayroutetables.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2transitgatewayroutetables.FilterAttributes](etgrt.ref.Append("filter"))
}

func (etgrt dataEc2TransitGatewayRouteTablesAttributes) Timeouts() dataec2transitgatewayroutetables.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2transitgatewayroutetables.TimeoutsAttributes](etgrt.ref.Append("timeouts"))
}
