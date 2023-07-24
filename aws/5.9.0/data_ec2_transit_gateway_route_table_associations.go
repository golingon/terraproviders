// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2transitgatewayroutetableassociations "github.com/golingon/terraproviders/aws/5.9.0/dataec2transitgatewayroutetableassociations"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2TransitGatewayRouteTableAssociations creates a new instance of [DataEc2TransitGatewayRouteTableAssociations].
func NewDataEc2TransitGatewayRouteTableAssociations(name string, args DataEc2TransitGatewayRouteTableAssociationsArgs) *DataEc2TransitGatewayRouteTableAssociations {
	return &DataEc2TransitGatewayRouteTableAssociations{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2TransitGatewayRouteTableAssociations)(nil)

// DataEc2TransitGatewayRouteTableAssociations represents the Terraform data resource aws_ec2_transit_gateway_route_table_associations.
type DataEc2TransitGatewayRouteTableAssociations struct {
	Name string
	Args DataEc2TransitGatewayRouteTableAssociationsArgs
}

// DataSource returns the Terraform object type for [DataEc2TransitGatewayRouteTableAssociations].
func (etgrta *DataEc2TransitGatewayRouteTableAssociations) DataSource() string {
	return "aws_ec2_transit_gateway_route_table_associations"
}

// LocalName returns the local name for [DataEc2TransitGatewayRouteTableAssociations].
func (etgrta *DataEc2TransitGatewayRouteTableAssociations) LocalName() string {
	return etgrta.Name
}

// Configuration returns the configuration (args) for [DataEc2TransitGatewayRouteTableAssociations].
func (etgrta *DataEc2TransitGatewayRouteTableAssociations) Configuration() interface{} {
	return etgrta.Args
}

// Attributes returns the attributes for [DataEc2TransitGatewayRouteTableAssociations].
func (etgrta *DataEc2TransitGatewayRouteTableAssociations) Attributes() dataEc2TransitGatewayRouteTableAssociationsAttributes {
	return dataEc2TransitGatewayRouteTableAssociationsAttributes{ref: terra.ReferenceDataResource(etgrta)}
}

// DataEc2TransitGatewayRouteTableAssociationsArgs contains the configurations for aws_ec2_transit_gateway_route_table_associations.
type DataEc2TransitGatewayRouteTableAssociationsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TransitGatewayRouteTableId: string, required
	TransitGatewayRouteTableId terra.StringValue `hcl:"transit_gateway_route_table_id,attr" validate:"required"`
	// Filter: min=0
	Filter []dataec2transitgatewayroutetableassociations.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2transitgatewayroutetableassociations.Timeouts `hcl:"timeouts,block"`
}
type dataEc2TransitGatewayRouteTableAssociationsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_transit_gateway_route_table_associations.
func (etgrta dataEc2TransitGatewayRouteTableAssociationsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgrta.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ec2_transit_gateway_route_table_associations.
func (etgrta dataEc2TransitGatewayRouteTableAssociationsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](etgrta.ref.Append("ids"))
}

// TransitGatewayRouteTableId returns a reference to field transit_gateway_route_table_id of aws_ec2_transit_gateway_route_table_associations.
func (etgrta dataEc2TransitGatewayRouteTableAssociationsAttributes) TransitGatewayRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(etgrta.ref.Append("transit_gateway_route_table_id"))
}

func (etgrta dataEc2TransitGatewayRouteTableAssociationsAttributes) Filter() terra.SetValue[dataec2transitgatewayroutetableassociations.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2transitgatewayroutetableassociations.FilterAttributes](etgrta.ref.Append("filter"))
}

func (etgrta dataEc2TransitGatewayRouteTableAssociationsAttributes) Timeouts() dataec2transitgatewayroutetableassociations.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2transitgatewayroutetableassociations.TimeoutsAttributes](etgrta.ref.Append("timeouts"))
}
