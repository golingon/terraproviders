// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataNetworkmanagerConnections creates a new instance of [DataNetworkmanagerConnections].
func NewDataNetworkmanagerConnections(name string, args DataNetworkmanagerConnectionsArgs) *DataNetworkmanagerConnections {
	return &DataNetworkmanagerConnections{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkmanagerConnections)(nil)

// DataNetworkmanagerConnections represents the Terraform data resource aws_networkmanager_connections.
type DataNetworkmanagerConnections struct {
	Name string
	Args DataNetworkmanagerConnectionsArgs
}

// DataSource returns the Terraform object type for [DataNetworkmanagerConnections].
func (nc *DataNetworkmanagerConnections) DataSource() string {
	return "aws_networkmanager_connections"
}

// LocalName returns the local name for [DataNetworkmanagerConnections].
func (nc *DataNetworkmanagerConnections) LocalName() string {
	return nc.Name
}

// Configuration returns the configuration (args) for [DataNetworkmanagerConnections].
func (nc *DataNetworkmanagerConnections) Configuration() interface{} {
	return nc.Args
}

// Attributes returns the attributes for [DataNetworkmanagerConnections].
func (nc *DataNetworkmanagerConnections) Attributes() dataNetworkmanagerConnectionsAttributes {
	return dataNetworkmanagerConnectionsAttributes{ref: terra.ReferenceDataResource(nc)}
}

// DataNetworkmanagerConnectionsArgs contains the configurations for aws_networkmanager_connections.
type DataNetworkmanagerConnectionsArgs struct {
	// DeviceId: string, optional
	DeviceId terra.StringValue `hcl:"device_id,attr"`
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataNetworkmanagerConnectionsAttributes struct {
	ref terra.Reference
}

// DeviceId returns a reference to field device_id of aws_networkmanager_connections.
func (nc dataNetworkmanagerConnectionsAttributes) DeviceId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("device_id"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_connections.
func (nc dataNetworkmanagerConnectionsAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_connections.
func (nc dataNetworkmanagerConnectionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_networkmanager_connections.
func (nc dataNetworkmanagerConnectionsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nc.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_networkmanager_connections.
func (nc dataNetworkmanagerConnectionsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nc.ref.Append("tags"))
}
