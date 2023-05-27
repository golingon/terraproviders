// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2localgatewayvirtualinterface "github.com/golingon/terraproviders/aws/5.0.1/dataec2localgatewayvirtualinterface"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2LocalGatewayVirtualInterface creates a new instance of [DataEc2LocalGatewayVirtualInterface].
func NewDataEc2LocalGatewayVirtualInterface(name string, args DataEc2LocalGatewayVirtualInterfaceArgs) *DataEc2LocalGatewayVirtualInterface {
	return &DataEc2LocalGatewayVirtualInterface{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2LocalGatewayVirtualInterface)(nil)

// DataEc2LocalGatewayVirtualInterface represents the Terraform data resource aws_ec2_local_gateway_virtual_interface.
type DataEc2LocalGatewayVirtualInterface struct {
	Name string
	Args DataEc2LocalGatewayVirtualInterfaceArgs
}

// DataSource returns the Terraform object type for [DataEc2LocalGatewayVirtualInterface].
func (elgvi *DataEc2LocalGatewayVirtualInterface) DataSource() string {
	return "aws_ec2_local_gateway_virtual_interface"
}

// LocalName returns the local name for [DataEc2LocalGatewayVirtualInterface].
func (elgvi *DataEc2LocalGatewayVirtualInterface) LocalName() string {
	return elgvi.Name
}

// Configuration returns the configuration (args) for [DataEc2LocalGatewayVirtualInterface].
func (elgvi *DataEc2LocalGatewayVirtualInterface) Configuration() interface{} {
	return elgvi.Args
}

// Attributes returns the attributes for [DataEc2LocalGatewayVirtualInterface].
func (elgvi *DataEc2LocalGatewayVirtualInterface) Attributes() dataEc2LocalGatewayVirtualInterfaceAttributes {
	return dataEc2LocalGatewayVirtualInterfaceAttributes{ref: terra.ReferenceDataResource(elgvi)}
}

// DataEc2LocalGatewayVirtualInterfaceArgs contains the configurations for aws_ec2_local_gateway_virtual_interface.
type DataEc2LocalGatewayVirtualInterfaceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2localgatewayvirtualinterface.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2localgatewayvirtualinterface.Timeouts `hcl:"timeouts,block"`
}
type dataEc2LocalGatewayVirtualInterfaceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_local_gateway_virtual_interface.
func (elgvi dataEc2LocalGatewayVirtualInterfaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(elgvi.ref.Append("id"))
}

// LocalAddress returns a reference to field local_address of aws_ec2_local_gateway_virtual_interface.
func (elgvi dataEc2LocalGatewayVirtualInterfaceAttributes) LocalAddress() terra.StringValue {
	return terra.ReferenceAsString(elgvi.ref.Append("local_address"))
}

// LocalBgpAsn returns a reference to field local_bgp_asn of aws_ec2_local_gateway_virtual_interface.
func (elgvi dataEc2LocalGatewayVirtualInterfaceAttributes) LocalBgpAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(elgvi.ref.Append("local_bgp_asn"))
}

// LocalGatewayId returns a reference to field local_gateway_id of aws_ec2_local_gateway_virtual_interface.
func (elgvi dataEc2LocalGatewayVirtualInterfaceAttributes) LocalGatewayId() terra.StringValue {
	return terra.ReferenceAsString(elgvi.ref.Append("local_gateway_id"))
}

// LocalGatewayVirtualInterfaceIds returns a reference to field local_gateway_virtual_interface_ids of aws_ec2_local_gateway_virtual_interface.
func (elgvi dataEc2LocalGatewayVirtualInterfaceAttributes) LocalGatewayVirtualInterfaceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](elgvi.ref.Append("local_gateway_virtual_interface_ids"))
}

// PeerAddress returns a reference to field peer_address of aws_ec2_local_gateway_virtual_interface.
func (elgvi dataEc2LocalGatewayVirtualInterfaceAttributes) PeerAddress() terra.StringValue {
	return terra.ReferenceAsString(elgvi.ref.Append("peer_address"))
}

// PeerBgpAsn returns a reference to field peer_bgp_asn of aws_ec2_local_gateway_virtual_interface.
func (elgvi dataEc2LocalGatewayVirtualInterfaceAttributes) PeerBgpAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(elgvi.ref.Append("peer_bgp_asn"))
}

// Tags returns a reference to field tags of aws_ec2_local_gateway_virtual_interface.
func (elgvi dataEc2LocalGatewayVirtualInterfaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](elgvi.ref.Append("tags"))
}

// Vlan returns a reference to field vlan of aws_ec2_local_gateway_virtual_interface.
func (elgvi dataEc2LocalGatewayVirtualInterfaceAttributes) Vlan() terra.NumberValue {
	return terra.ReferenceAsNumber(elgvi.ref.Append("vlan"))
}

func (elgvi dataEc2LocalGatewayVirtualInterfaceAttributes) Filter() terra.SetValue[dataec2localgatewayvirtualinterface.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2localgatewayvirtualinterface.FilterAttributes](elgvi.ref.Append("filter"))
}

func (elgvi dataEc2LocalGatewayVirtualInterfaceAttributes) Timeouts() dataec2localgatewayvirtualinterface.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2localgatewayvirtualinterface.TimeoutsAttributes](elgvi.ref.Append("timeouts"))
}
