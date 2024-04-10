// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	dataiothubdps "github.com/golingon/terraproviders/azurerm/3.98.0/dataiothubdps"
)

// NewDataIothubDps creates a new instance of [DataIothubDps].
func NewDataIothubDps(name string, args DataIothubDpsArgs) *DataIothubDps {
	return &DataIothubDps{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIothubDps)(nil)

// DataIothubDps represents the Terraform data resource azurerm_iothub_dps.
type DataIothubDps struct {
	Name string
	Args DataIothubDpsArgs
}

// DataSource returns the Terraform object type for [DataIothubDps].
func (id *DataIothubDps) DataSource() string {
	return "azurerm_iothub_dps"
}

// LocalName returns the local name for [DataIothubDps].
func (id *DataIothubDps) LocalName() string {
	return id.Name
}

// Configuration returns the configuration (args) for [DataIothubDps].
func (id *DataIothubDps) Configuration() interface{} {
	return id.Args
}

// Attributes returns the attributes for [DataIothubDps].
func (id *DataIothubDps) Attributes() dataIothubDpsAttributes {
	return dataIothubDpsAttributes{ref: terra.ReferenceDataResource(id)}
}

// DataIothubDpsArgs contains the configurations for azurerm_iothub_dps.
type DataIothubDpsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *dataiothubdps.Timeouts `hcl:"timeouts,block"`
}
type dataIothubDpsAttributes struct {
	ref terra.Reference
}

// AllocationPolicy returns a reference to field allocation_policy of azurerm_iothub_dps.
func (id dataIothubDpsAttributes) AllocationPolicy() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("allocation_policy"))
}

// DeviceProvisioningHostName returns a reference to field device_provisioning_host_name of azurerm_iothub_dps.
func (id dataIothubDpsAttributes) DeviceProvisioningHostName() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("device_provisioning_host_name"))
}

// Id returns a reference to field id of azurerm_iothub_dps.
func (id dataIothubDpsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("id"))
}

// IdScope returns a reference to field id_scope of azurerm_iothub_dps.
func (id dataIothubDpsAttributes) IdScope() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("id_scope"))
}

// Location returns a reference to field location of azurerm_iothub_dps.
func (id dataIothubDpsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_iothub_dps.
func (id dataIothubDpsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_dps.
func (id dataIothubDpsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("resource_group_name"))
}

// ServiceOperationsHostName returns a reference to field service_operations_host_name of azurerm_iothub_dps.
func (id dataIothubDpsAttributes) ServiceOperationsHostName() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("service_operations_host_name"))
}

// Tags returns a reference to field tags of azurerm_iothub_dps.
func (id dataIothubDpsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](id.ref.Append("tags"))
}

func (id dataIothubDpsAttributes) Timeouts() dataiothubdps.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataiothubdps.TimeoutsAttributes](id.ref.Append("timeouts"))
}
