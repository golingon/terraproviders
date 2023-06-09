// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataNetworkmanagerDevices creates a new instance of [DataNetworkmanagerDevices].
func NewDataNetworkmanagerDevices(name string, args DataNetworkmanagerDevicesArgs) *DataNetworkmanagerDevices {
	return &DataNetworkmanagerDevices{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkmanagerDevices)(nil)

// DataNetworkmanagerDevices represents the Terraform data resource aws_networkmanager_devices.
type DataNetworkmanagerDevices struct {
	Name string
	Args DataNetworkmanagerDevicesArgs
}

// DataSource returns the Terraform object type for [DataNetworkmanagerDevices].
func (nd *DataNetworkmanagerDevices) DataSource() string {
	return "aws_networkmanager_devices"
}

// LocalName returns the local name for [DataNetworkmanagerDevices].
func (nd *DataNetworkmanagerDevices) LocalName() string {
	return nd.Name
}

// Configuration returns the configuration (args) for [DataNetworkmanagerDevices].
func (nd *DataNetworkmanagerDevices) Configuration() interface{} {
	return nd.Args
}

// Attributes returns the attributes for [DataNetworkmanagerDevices].
func (nd *DataNetworkmanagerDevices) Attributes() dataNetworkmanagerDevicesAttributes {
	return dataNetworkmanagerDevicesAttributes{ref: terra.ReferenceDataResource(nd)}
}

// DataNetworkmanagerDevicesArgs contains the configurations for aws_networkmanager_devices.
type DataNetworkmanagerDevicesArgs struct {
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SiteId: string, optional
	SiteId terra.StringValue `hcl:"site_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataNetworkmanagerDevicesAttributes struct {
	ref terra.Reference
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_devices.
func (nd dataNetworkmanagerDevicesAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_devices.
func (nd dataNetworkmanagerDevicesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_networkmanager_devices.
func (nd dataNetworkmanagerDevicesAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nd.ref.Append("ids"))
}

// SiteId returns a reference to field site_id of aws_networkmanager_devices.
func (nd dataNetworkmanagerDevicesAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("site_id"))
}

// Tags returns a reference to field tags of aws_networkmanager_devices.
func (nd dataNetworkmanagerDevicesAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nd.ref.Append("tags"))
}
