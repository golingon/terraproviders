// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataNetworkmanagerConnection creates a new instance of [DataNetworkmanagerConnection].
func NewDataNetworkmanagerConnection(name string, args DataNetworkmanagerConnectionArgs) *DataNetworkmanagerConnection {
	return &DataNetworkmanagerConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkmanagerConnection)(nil)

// DataNetworkmanagerConnection represents the Terraform data resource aws_networkmanager_connection.
type DataNetworkmanagerConnection struct {
	Name string
	Args DataNetworkmanagerConnectionArgs
}

// DataSource returns the Terraform object type for [DataNetworkmanagerConnection].
func (nc *DataNetworkmanagerConnection) DataSource() string {
	return "aws_networkmanager_connection"
}

// LocalName returns the local name for [DataNetworkmanagerConnection].
func (nc *DataNetworkmanagerConnection) LocalName() string {
	return nc.Name
}

// Configuration returns the configuration (args) for [DataNetworkmanagerConnection].
func (nc *DataNetworkmanagerConnection) Configuration() interface{} {
	return nc.Args
}

// Attributes returns the attributes for [DataNetworkmanagerConnection].
func (nc *DataNetworkmanagerConnection) Attributes() dataNetworkmanagerConnectionAttributes {
	return dataNetworkmanagerConnectionAttributes{ref: terra.ReferenceDataResource(nc)}
}

// DataNetworkmanagerConnectionArgs contains the configurations for aws_networkmanager_connection.
type DataNetworkmanagerConnectionArgs struct {
	// ConnectionId: string, required
	ConnectionId terra.StringValue `hcl:"connection_id,attr" validate:"required"`
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataNetworkmanagerConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_connection.
func (nc dataNetworkmanagerConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("arn"))
}

// ConnectedDeviceId returns a reference to field connected_device_id of aws_networkmanager_connection.
func (nc dataNetworkmanagerConnectionAttributes) ConnectedDeviceId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("connected_device_id"))
}

// ConnectedLinkId returns a reference to field connected_link_id of aws_networkmanager_connection.
func (nc dataNetworkmanagerConnectionAttributes) ConnectedLinkId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("connected_link_id"))
}

// ConnectionId returns a reference to field connection_id of aws_networkmanager_connection.
func (nc dataNetworkmanagerConnectionAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("connection_id"))
}

// Description returns a reference to field description of aws_networkmanager_connection.
func (nc dataNetworkmanagerConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("description"))
}

// DeviceId returns a reference to field device_id of aws_networkmanager_connection.
func (nc dataNetworkmanagerConnectionAttributes) DeviceId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("device_id"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_connection.
func (nc dataNetworkmanagerConnectionAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_connection.
func (nc dataNetworkmanagerConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("id"))
}

// LinkId returns a reference to field link_id of aws_networkmanager_connection.
func (nc dataNetworkmanagerConnectionAttributes) LinkId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("link_id"))
}

// Tags returns a reference to field tags of aws_networkmanager_connection.
func (nc dataNetworkmanagerConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nc.ref.Append("tags"))
}
