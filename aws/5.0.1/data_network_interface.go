// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datanetworkinterface "github.com/golingon/terraproviders/aws/5.0.1/datanetworkinterface"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetworkInterface creates a new instance of [DataNetworkInterface].
func NewDataNetworkInterface(name string, args DataNetworkInterfaceArgs) *DataNetworkInterface {
	return &DataNetworkInterface{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkInterface)(nil)

// DataNetworkInterface represents the Terraform data resource aws_network_interface.
type DataNetworkInterface struct {
	Name string
	Args DataNetworkInterfaceArgs
}

// DataSource returns the Terraform object type for [DataNetworkInterface].
func (ni *DataNetworkInterface) DataSource() string {
	return "aws_network_interface"
}

// LocalName returns the local name for [DataNetworkInterface].
func (ni *DataNetworkInterface) LocalName() string {
	return ni.Name
}

// Configuration returns the configuration (args) for [DataNetworkInterface].
func (ni *DataNetworkInterface) Configuration() interface{} {
	return ni.Args
}

// Attributes returns the attributes for [DataNetworkInterface].
func (ni *DataNetworkInterface) Attributes() dataNetworkInterfaceAttributes {
	return dataNetworkInterfaceAttributes{ref: terra.ReferenceDataResource(ni)}
}

// DataNetworkInterfaceArgs contains the configurations for aws_network_interface.
type DataNetworkInterfaceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Association: min=0
	Association []datanetworkinterface.Association `hcl:"association,block" validate:"min=0"`
	// Attachment: min=0
	Attachment []datanetworkinterface.Attachment `hcl:"attachment,block" validate:"min=0"`
	// Filter: min=0
	Filter []datanetworkinterface.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datanetworkinterface.Timeouts `hcl:"timeouts,block"`
}
type dataNetworkInterfaceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("arn"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("availability_zone"))
}

// Description returns a reference to field description of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("description"))
}

// Id returns a reference to field id of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("id"))
}

// InterfaceType returns a reference to field interface_type of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) InterfaceType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("interface_type"))
}

// Ipv6Addresses returns a reference to field ipv6_addresses of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) Ipv6Addresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ni.ref.Append("ipv6_addresses"))
}

// MacAddress returns a reference to field mac_address of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) MacAddress() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("mac_address"))
}

// OutpostArn returns a reference to field outpost_arn of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("outpost_arn"))
}

// OwnerId returns a reference to field owner_id of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("owner_id"))
}

// PrivateDnsName returns a reference to field private_dns_name of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) PrivateDnsName() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("private_dns_name"))
}

// PrivateIp returns a reference to field private_ip of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) PrivateIp() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("private_ip"))
}

// PrivateIps returns a reference to field private_ips of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) PrivateIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ni.ref.Append("private_ips"))
}

// RequesterId returns a reference to field requester_id of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) RequesterId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("requester_id"))
}

// SecurityGroups returns a reference to field security_groups of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ni.ref.Append("security_groups"))
}

// SubnetId returns a reference to field subnet_id of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ni.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_network_interface.
func (ni dataNetworkInterfaceAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("vpc_id"))
}

func (ni dataNetworkInterfaceAttributes) Association() terra.ListValue[datanetworkinterface.AssociationAttributes] {
	return terra.ReferenceAsList[datanetworkinterface.AssociationAttributes](ni.ref.Append("association"))
}

func (ni dataNetworkInterfaceAttributes) Attachment() terra.ListValue[datanetworkinterface.AttachmentAttributes] {
	return terra.ReferenceAsList[datanetworkinterface.AttachmentAttributes](ni.ref.Append("attachment"))
}

func (ni dataNetworkInterfaceAttributes) Filter() terra.SetValue[datanetworkinterface.FilterAttributes] {
	return terra.ReferenceAsSet[datanetworkinterface.FilterAttributes](ni.ref.Append("filter"))
}

func (ni dataNetworkInterfaceAttributes) Timeouts() datanetworkinterface.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanetworkinterface.TimeoutsAttributes](ni.ref.Append("timeouts"))
}
