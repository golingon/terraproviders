// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datanetworkmanagerdevice "github.com/golingon/terraproviders/aws/5.0.1/datanetworkmanagerdevice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetworkmanagerDevice creates a new instance of [DataNetworkmanagerDevice].
func NewDataNetworkmanagerDevice(name string, args DataNetworkmanagerDeviceArgs) *DataNetworkmanagerDevice {
	return &DataNetworkmanagerDevice{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkmanagerDevice)(nil)

// DataNetworkmanagerDevice represents the Terraform data resource aws_networkmanager_device.
type DataNetworkmanagerDevice struct {
	Name string
	Args DataNetworkmanagerDeviceArgs
}

// DataSource returns the Terraform object type for [DataNetworkmanagerDevice].
func (nd *DataNetworkmanagerDevice) DataSource() string {
	return "aws_networkmanager_device"
}

// LocalName returns the local name for [DataNetworkmanagerDevice].
func (nd *DataNetworkmanagerDevice) LocalName() string {
	return nd.Name
}

// Configuration returns the configuration (args) for [DataNetworkmanagerDevice].
func (nd *DataNetworkmanagerDevice) Configuration() interface{} {
	return nd.Args
}

// Attributes returns the attributes for [DataNetworkmanagerDevice].
func (nd *DataNetworkmanagerDevice) Attributes() dataNetworkmanagerDeviceAttributes {
	return dataNetworkmanagerDeviceAttributes{ref: terra.ReferenceDataResource(nd)}
}

// DataNetworkmanagerDeviceArgs contains the configurations for aws_networkmanager_device.
type DataNetworkmanagerDeviceArgs struct {
	// DeviceId: string, required
	DeviceId terra.StringValue `hcl:"device_id,attr" validate:"required"`
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AwsLocation: min=0
	AwsLocation []datanetworkmanagerdevice.AwsLocation `hcl:"aws_location,block" validate:"min=0"`
	// Location: min=0
	Location []datanetworkmanagerdevice.Location `hcl:"location,block" validate:"min=0"`
}
type dataNetworkmanagerDeviceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_device.
func (nd dataNetworkmanagerDeviceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("arn"))
}

// Description returns a reference to field description of aws_networkmanager_device.
func (nd dataNetworkmanagerDeviceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("description"))
}

// DeviceId returns a reference to field device_id of aws_networkmanager_device.
func (nd dataNetworkmanagerDeviceAttributes) DeviceId() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("device_id"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_device.
func (nd dataNetworkmanagerDeviceAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_device.
func (nd dataNetworkmanagerDeviceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("id"))
}

// Model returns a reference to field model of aws_networkmanager_device.
func (nd dataNetworkmanagerDeviceAttributes) Model() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("model"))
}

// SerialNumber returns a reference to field serial_number of aws_networkmanager_device.
func (nd dataNetworkmanagerDeviceAttributes) SerialNumber() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("serial_number"))
}

// SiteId returns a reference to field site_id of aws_networkmanager_device.
func (nd dataNetworkmanagerDeviceAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("site_id"))
}

// Tags returns a reference to field tags of aws_networkmanager_device.
func (nd dataNetworkmanagerDeviceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nd.ref.Append("tags"))
}

// Type returns a reference to field type of aws_networkmanager_device.
func (nd dataNetworkmanagerDeviceAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("type"))
}

// Vendor returns a reference to field vendor of aws_networkmanager_device.
func (nd dataNetworkmanagerDeviceAttributes) Vendor() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("vendor"))
}

func (nd dataNetworkmanagerDeviceAttributes) AwsLocation() terra.ListValue[datanetworkmanagerdevice.AwsLocationAttributes] {
	return terra.ReferenceAsList[datanetworkmanagerdevice.AwsLocationAttributes](nd.ref.Append("aws_location"))
}

func (nd dataNetworkmanagerDeviceAttributes) Location() terra.ListValue[datanetworkmanagerdevice.LocationAttributes] {
	return terra.ReferenceAsList[datanetworkmanagerdevice.LocationAttributes](nd.ref.Append("location"))
}
