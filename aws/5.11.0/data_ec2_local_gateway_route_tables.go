// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2localgatewayroutetables "github.com/golingon/terraproviders/aws/5.11.0/dataec2localgatewayroutetables"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2LocalGatewayRouteTables creates a new instance of [DataEc2LocalGatewayRouteTables].
func NewDataEc2LocalGatewayRouteTables(name string, args DataEc2LocalGatewayRouteTablesArgs) *DataEc2LocalGatewayRouteTables {
	return &DataEc2LocalGatewayRouteTables{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2LocalGatewayRouteTables)(nil)

// DataEc2LocalGatewayRouteTables represents the Terraform data resource aws_ec2_local_gateway_route_tables.
type DataEc2LocalGatewayRouteTables struct {
	Name string
	Args DataEc2LocalGatewayRouteTablesArgs
}

// DataSource returns the Terraform object type for [DataEc2LocalGatewayRouteTables].
func (elgrt *DataEc2LocalGatewayRouteTables) DataSource() string {
	return "aws_ec2_local_gateway_route_tables"
}

// LocalName returns the local name for [DataEc2LocalGatewayRouteTables].
func (elgrt *DataEc2LocalGatewayRouteTables) LocalName() string {
	return elgrt.Name
}

// Configuration returns the configuration (args) for [DataEc2LocalGatewayRouteTables].
func (elgrt *DataEc2LocalGatewayRouteTables) Configuration() interface{} {
	return elgrt.Args
}

// Attributes returns the attributes for [DataEc2LocalGatewayRouteTables].
func (elgrt *DataEc2LocalGatewayRouteTables) Attributes() dataEc2LocalGatewayRouteTablesAttributes {
	return dataEc2LocalGatewayRouteTablesAttributes{ref: terra.ReferenceDataResource(elgrt)}
}

// DataEc2LocalGatewayRouteTablesArgs contains the configurations for aws_ec2_local_gateway_route_tables.
type DataEc2LocalGatewayRouteTablesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2localgatewayroutetables.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2localgatewayroutetables.Timeouts `hcl:"timeouts,block"`
}
type dataEc2LocalGatewayRouteTablesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_local_gateway_route_tables.
func (elgrt dataEc2LocalGatewayRouteTablesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(elgrt.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ec2_local_gateway_route_tables.
func (elgrt dataEc2LocalGatewayRouteTablesAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](elgrt.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_ec2_local_gateway_route_tables.
func (elgrt dataEc2LocalGatewayRouteTablesAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](elgrt.ref.Append("tags"))
}

func (elgrt dataEc2LocalGatewayRouteTablesAttributes) Filter() terra.SetValue[dataec2localgatewayroutetables.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2localgatewayroutetables.FilterAttributes](elgrt.ref.Append("filter"))
}

func (elgrt dataEc2LocalGatewayRouteTablesAttributes) Timeouts() dataec2localgatewayroutetables.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2localgatewayroutetables.TimeoutsAttributes](elgrt.ref.Append("timeouts"))
}
