// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	datavpngateway "github.com/golingon/terraproviders/aws/5.44.0/datavpngateway"
)

// NewDataVpnGateway creates a new instance of [DataVpnGateway].
func NewDataVpnGateway(name string, args DataVpnGatewayArgs) *DataVpnGateway {
	return &DataVpnGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpnGateway)(nil)

// DataVpnGateway represents the Terraform data resource aws_vpn_gateway.
type DataVpnGateway struct {
	Name string
	Args DataVpnGatewayArgs
}

// DataSource returns the Terraform object type for [DataVpnGateway].
func (vg *DataVpnGateway) DataSource() string {
	return "aws_vpn_gateway"
}

// LocalName returns the local name for [DataVpnGateway].
func (vg *DataVpnGateway) LocalName() string {
	return vg.Name
}

// Configuration returns the configuration (args) for [DataVpnGateway].
func (vg *DataVpnGateway) Configuration() interface{} {
	return vg.Args
}

// Attributes returns the attributes for [DataVpnGateway].
func (vg *DataVpnGateway) Attributes() dataVpnGatewayAttributes {
	return dataVpnGatewayAttributes{ref: terra.ReferenceDataResource(vg)}
}

// DataVpnGatewayArgs contains the configurations for aws_vpn_gateway.
type DataVpnGatewayArgs struct {
	// AmazonSideAsn: string, optional
	AmazonSideAsn terra.StringValue `hcl:"amazon_side_asn,attr"`
	// AttachedVpcId: string, optional
	AttachedVpcId terra.StringValue `hcl:"attached_vpc_id,attr"`
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []datavpngateway.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavpngateway.Timeouts `hcl:"timeouts,block"`
}
type dataVpnGatewayAttributes struct {
	ref terra.Reference
}

// AmazonSideAsn returns a reference to field amazon_side_asn of aws_vpn_gateway.
func (vg dataVpnGatewayAttributes) AmazonSideAsn() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("amazon_side_asn"))
}

// Arn returns a reference to field arn of aws_vpn_gateway.
func (vg dataVpnGatewayAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("arn"))
}

// AttachedVpcId returns a reference to field attached_vpc_id of aws_vpn_gateway.
func (vg dataVpnGatewayAttributes) AttachedVpcId() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("attached_vpc_id"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_vpn_gateway.
func (vg dataVpnGatewayAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("availability_zone"))
}

// Id returns a reference to field id of aws_vpn_gateway.
func (vg dataVpnGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("id"))
}

// State returns a reference to field state of aws_vpn_gateway.
func (vg dataVpnGatewayAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_vpn_gateway.
func (vg dataVpnGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vg.ref.Append("tags"))
}

func (vg dataVpnGatewayAttributes) Filter() terra.SetValue[datavpngateway.FilterAttributes] {
	return terra.ReferenceAsSet[datavpngateway.FilterAttributes](vg.ref.Append("filter"))
}

func (vg dataVpnGatewayAttributes) Timeouts() datavpngateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpngateway.TimeoutsAttributes](vg.ref.Append("timeouts"))
}
