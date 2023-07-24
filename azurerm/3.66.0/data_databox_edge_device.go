// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadataboxedgedevice "github.com/golingon/terraproviders/azurerm/3.66.0/datadataboxedgedevice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDataboxEdgeDevice creates a new instance of [DataDataboxEdgeDevice].
func NewDataDataboxEdgeDevice(name string, args DataDataboxEdgeDeviceArgs) *DataDataboxEdgeDevice {
	return &DataDataboxEdgeDevice{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataboxEdgeDevice)(nil)

// DataDataboxEdgeDevice represents the Terraform data resource azurerm_databox_edge_device.
type DataDataboxEdgeDevice struct {
	Name string
	Args DataDataboxEdgeDeviceArgs
}

// DataSource returns the Terraform object type for [DataDataboxEdgeDevice].
func (ded *DataDataboxEdgeDevice) DataSource() string {
	return "azurerm_databox_edge_device"
}

// LocalName returns the local name for [DataDataboxEdgeDevice].
func (ded *DataDataboxEdgeDevice) LocalName() string {
	return ded.Name
}

// Configuration returns the configuration (args) for [DataDataboxEdgeDevice].
func (ded *DataDataboxEdgeDevice) Configuration() interface{} {
	return ded.Args
}

// Attributes returns the attributes for [DataDataboxEdgeDevice].
func (ded *DataDataboxEdgeDevice) Attributes() dataDataboxEdgeDeviceAttributes {
	return dataDataboxEdgeDeviceAttributes{ref: terra.ReferenceDataResource(ded)}
}

// DataDataboxEdgeDeviceArgs contains the configurations for azurerm_databox_edge_device.
type DataDataboxEdgeDeviceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// DeviceProperties: min=0
	DeviceProperties []datadataboxedgedevice.DeviceProperties `hcl:"device_properties,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadataboxedgedevice.Timeouts `hcl:"timeouts,block"`
}
type dataDataboxEdgeDeviceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_databox_edge_device.
func (ded dataDataboxEdgeDeviceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ded.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_databox_edge_device.
func (ded dataDataboxEdgeDeviceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ded.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_databox_edge_device.
func (ded dataDataboxEdgeDeviceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ded.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_databox_edge_device.
func (ded dataDataboxEdgeDeviceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ded.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_databox_edge_device.
func (ded dataDataboxEdgeDeviceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ded.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_databox_edge_device.
func (ded dataDataboxEdgeDeviceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ded.ref.Append("tags"))
}

func (ded dataDataboxEdgeDeviceAttributes) DeviceProperties() terra.ListValue[datadataboxedgedevice.DevicePropertiesAttributes] {
	return terra.ReferenceAsList[datadataboxedgedevice.DevicePropertiesAttributes](ded.ref.Append("device_properties"))
}

func (ded dataDataboxEdgeDeviceAttributes) Timeouts() datadataboxedgedevice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadataboxedgedevice.TimeoutsAttributes](ded.ref.Append("timeouts"))
}
