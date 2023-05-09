// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2transitgatewayconnectpeer "github.com/golingon/terraproviders/aws/4.66.1/dataec2transitgatewayconnectpeer"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2TransitGatewayConnectPeer creates a new instance of [DataEc2TransitGatewayConnectPeer].
func NewDataEc2TransitGatewayConnectPeer(name string, args DataEc2TransitGatewayConnectPeerArgs) *DataEc2TransitGatewayConnectPeer {
	return &DataEc2TransitGatewayConnectPeer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2TransitGatewayConnectPeer)(nil)

// DataEc2TransitGatewayConnectPeer represents the Terraform data resource aws_ec2_transit_gateway_connect_peer.
type DataEc2TransitGatewayConnectPeer struct {
	Name string
	Args DataEc2TransitGatewayConnectPeerArgs
}

// DataSource returns the Terraform object type for [DataEc2TransitGatewayConnectPeer].
func (etgcp *DataEc2TransitGatewayConnectPeer) DataSource() string {
	return "aws_ec2_transit_gateway_connect_peer"
}

// LocalName returns the local name for [DataEc2TransitGatewayConnectPeer].
func (etgcp *DataEc2TransitGatewayConnectPeer) LocalName() string {
	return etgcp.Name
}

// Configuration returns the configuration (args) for [DataEc2TransitGatewayConnectPeer].
func (etgcp *DataEc2TransitGatewayConnectPeer) Configuration() interface{} {
	return etgcp.Args
}

// Attributes returns the attributes for [DataEc2TransitGatewayConnectPeer].
func (etgcp *DataEc2TransitGatewayConnectPeer) Attributes() dataEc2TransitGatewayConnectPeerAttributes {
	return dataEc2TransitGatewayConnectPeerAttributes{ref: terra.ReferenceDataResource(etgcp)}
}

// DataEc2TransitGatewayConnectPeerArgs contains the configurations for aws_ec2_transit_gateway_connect_peer.
type DataEc2TransitGatewayConnectPeerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TransitGatewayConnectPeerId: string, optional
	TransitGatewayConnectPeerId terra.StringValue `hcl:"transit_gateway_connect_peer_id,attr"`
	// Filter: min=0
	Filter []dataec2transitgatewayconnectpeer.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2transitgatewayconnectpeer.Timeouts `hcl:"timeouts,block"`
}
type dataEc2TransitGatewayConnectPeerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_transit_gateway_connect_peer.
func (etgcp dataEc2TransitGatewayConnectPeerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("arn"))
}

// BgpAsn returns a reference to field bgp_asn of aws_ec2_transit_gateway_connect_peer.
func (etgcp dataEc2TransitGatewayConnectPeerAttributes) BgpAsn() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("bgp_asn"))
}

// Id returns a reference to field id of aws_ec2_transit_gateway_connect_peer.
func (etgcp dataEc2TransitGatewayConnectPeerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("id"))
}

// InsideCidrBlocks returns a reference to field inside_cidr_blocks of aws_ec2_transit_gateway_connect_peer.
func (etgcp dataEc2TransitGatewayConnectPeerAttributes) InsideCidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](etgcp.ref.Append("inside_cidr_blocks"))
}

// PeerAddress returns a reference to field peer_address of aws_ec2_transit_gateway_connect_peer.
func (etgcp dataEc2TransitGatewayConnectPeerAttributes) PeerAddress() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("peer_address"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_connect_peer.
func (etgcp dataEc2TransitGatewayConnectPeerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgcp.ref.Append("tags"))
}

// TransitGatewayAddress returns a reference to field transit_gateway_address of aws_ec2_transit_gateway_connect_peer.
func (etgcp dataEc2TransitGatewayConnectPeerAttributes) TransitGatewayAddress() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("transit_gateway_address"))
}

// TransitGatewayAttachmentId returns a reference to field transit_gateway_attachment_id of aws_ec2_transit_gateway_connect_peer.
func (etgcp dataEc2TransitGatewayConnectPeerAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("transit_gateway_attachment_id"))
}

// TransitGatewayConnectPeerId returns a reference to field transit_gateway_connect_peer_id of aws_ec2_transit_gateway_connect_peer.
func (etgcp dataEc2TransitGatewayConnectPeerAttributes) TransitGatewayConnectPeerId() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("transit_gateway_connect_peer_id"))
}

func (etgcp dataEc2TransitGatewayConnectPeerAttributes) Filter() terra.SetValue[dataec2transitgatewayconnectpeer.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2transitgatewayconnectpeer.FilterAttributes](etgcp.ref.Append("filter"))
}

func (etgcp dataEc2TransitGatewayConnectPeerAttributes) Timeouts() dataec2transitgatewayconnectpeer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2transitgatewayconnectpeer.TimeoutsAttributes](etgcp.ref.Append("timeouts"))
}
