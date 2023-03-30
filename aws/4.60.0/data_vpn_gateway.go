// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpngateway "github.com/golingon/terraproviders/aws/4.60.0/datavpngateway"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataVpnGateway(name string, args DataVpnGatewayArgs) *DataVpnGateway {
	return &DataVpnGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpnGateway)(nil)

type DataVpnGateway struct {
	Name string
	Args DataVpnGatewayArgs
}

func (vg *DataVpnGateway) DataSource() string {
	return "aws_vpn_gateway"
}

func (vg *DataVpnGateway) LocalName() string {
	return vg.Name
}

func (vg *DataVpnGateway) Configuration() interface{} {
	return vg.Args
}

func (vg *DataVpnGateway) Attributes() dataVpnGatewayAttributes {
	return dataVpnGatewayAttributes{ref: terra.ReferenceDataResource(vg)}
}

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

func (vg dataVpnGatewayAttributes) AmazonSideAsn() terra.StringValue {
	return terra.ReferenceString(vg.ref.Append("amazon_side_asn"))
}

func (vg dataVpnGatewayAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(vg.ref.Append("arn"))
}

func (vg dataVpnGatewayAttributes) AttachedVpcId() terra.StringValue {
	return terra.ReferenceString(vg.ref.Append("attached_vpc_id"))
}

func (vg dataVpnGatewayAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceString(vg.ref.Append("availability_zone"))
}

func (vg dataVpnGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceString(vg.ref.Append("id"))
}

func (vg dataVpnGatewayAttributes) State() terra.StringValue {
	return terra.ReferenceString(vg.ref.Append("state"))
}

func (vg dataVpnGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](vg.ref.Append("tags"))
}

func (vg dataVpnGatewayAttributes) Filter() terra.SetValue[datavpngateway.FilterAttributes] {
	return terra.ReferenceSet[datavpngateway.FilterAttributes](vg.ref.Append("filter"))
}

func (vg dataVpnGatewayAttributes) Timeouts() datavpngateway.TimeoutsAttributes {
	return terra.ReferenceSingle[datavpngateway.TimeoutsAttributes](vg.ref.Append("timeouts"))
}
