// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2transitgatewayconnect "github.com/golingon/terraproviders/aws/5.8.0/dataec2transitgatewayconnect"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2TransitGatewayConnect creates a new instance of [DataEc2TransitGatewayConnect].
func NewDataEc2TransitGatewayConnect(name string, args DataEc2TransitGatewayConnectArgs) *DataEc2TransitGatewayConnect {
	return &DataEc2TransitGatewayConnect{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2TransitGatewayConnect)(nil)

// DataEc2TransitGatewayConnect represents the Terraform data resource aws_ec2_transit_gateway_connect.
type DataEc2TransitGatewayConnect struct {
	Name string
	Args DataEc2TransitGatewayConnectArgs
}

// DataSource returns the Terraform object type for [DataEc2TransitGatewayConnect].
func (etgc *DataEc2TransitGatewayConnect) DataSource() string {
	return "aws_ec2_transit_gateway_connect"
}

// LocalName returns the local name for [DataEc2TransitGatewayConnect].
func (etgc *DataEc2TransitGatewayConnect) LocalName() string {
	return etgc.Name
}

// Configuration returns the configuration (args) for [DataEc2TransitGatewayConnect].
func (etgc *DataEc2TransitGatewayConnect) Configuration() interface{} {
	return etgc.Args
}

// Attributes returns the attributes for [DataEc2TransitGatewayConnect].
func (etgc *DataEc2TransitGatewayConnect) Attributes() dataEc2TransitGatewayConnectAttributes {
	return dataEc2TransitGatewayConnectAttributes{ref: terra.ReferenceDataResource(etgc)}
}

// DataEc2TransitGatewayConnectArgs contains the configurations for aws_ec2_transit_gateway_connect.
type DataEc2TransitGatewayConnectArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TransitGatewayConnectId: string, optional
	TransitGatewayConnectId terra.StringValue `hcl:"transit_gateway_connect_id,attr"`
	// Filter: min=0
	Filter []dataec2transitgatewayconnect.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2transitgatewayconnect.Timeouts `hcl:"timeouts,block"`
}
type dataEc2TransitGatewayConnectAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_transit_gateway_connect.
func (etgc dataEc2TransitGatewayConnectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgc.ref.Append("id"))
}

// Protocol returns a reference to field protocol of aws_ec2_transit_gateway_connect.
func (etgc dataEc2TransitGatewayConnectAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(etgc.ref.Append("protocol"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_connect.
func (etgc dataEc2TransitGatewayConnectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgc.ref.Append("tags"))
}

// TransitGatewayConnectId returns a reference to field transit_gateway_connect_id of aws_ec2_transit_gateway_connect.
func (etgc dataEc2TransitGatewayConnectAttributes) TransitGatewayConnectId() terra.StringValue {
	return terra.ReferenceAsString(etgc.ref.Append("transit_gateway_connect_id"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_ec2_transit_gateway_connect.
func (etgc dataEc2TransitGatewayConnectAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgc.ref.Append("transit_gateway_id"))
}

// TransportAttachmentId returns a reference to field transport_attachment_id of aws_ec2_transit_gateway_connect.
func (etgc dataEc2TransitGatewayConnectAttributes) TransportAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(etgc.ref.Append("transport_attachment_id"))
}

func (etgc dataEc2TransitGatewayConnectAttributes) Filter() terra.SetValue[dataec2transitgatewayconnect.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2transitgatewayconnect.FilterAttributes](etgc.ref.Append("filter"))
}

func (etgc dataEc2TransitGatewayConnectAttributes) Timeouts() dataec2transitgatewayconnect.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2transitgatewayconnect.TimeoutsAttributes](etgc.ref.Append("timeouts"))
}
