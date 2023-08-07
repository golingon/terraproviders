// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataorchestratedvirtualmachinescaleset "github.com/golingon/terraproviders/azurerm/3.68.0/dataorchestratedvirtualmachinescaleset"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataOrchestratedVirtualMachineScaleSet creates a new instance of [DataOrchestratedVirtualMachineScaleSet].
func NewDataOrchestratedVirtualMachineScaleSet(name string, args DataOrchestratedVirtualMachineScaleSetArgs) *DataOrchestratedVirtualMachineScaleSet {
	return &DataOrchestratedVirtualMachineScaleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOrchestratedVirtualMachineScaleSet)(nil)

// DataOrchestratedVirtualMachineScaleSet represents the Terraform data resource azurerm_orchestrated_virtual_machine_scale_set.
type DataOrchestratedVirtualMachineScaleSet struct {
	Name string
	Args DataOrchestratedVirtualMachineScaleSetArgs
}

// DataSource returns the Terraform object type for [DataOrchestratedVirtualMachineScaleSet].
func (ovmss *DataOrchestratedVirtualMachineScaleSet) DataSource() string {
	return "azurerm_orchestrated_virtual_machine_scale_set"
}

// LocalName returns the local name for [DataOrchestratedVirtualMachineScaleSet].
func (ovmss *DataOrchestratedVirtualMachineScaleSet) LocalName() string {
	return ovmss.Name
}

// Configuration returns the configuration (args) for [DataOrchestratedVirtualMachineScaleSet].
func (ovmss *DataOrchestratedVirtualMachineScaleSet) Configuration() interface{} {
	return ovmss.Args
}

// Attributes returns the attributes for [DataOrchestratedVirtualMachineScaleSet].
func (ovmss *DataOrchestratedVirtualMachineScaleSet) Attributes() dataOrchestratedVirtualMachineScaleSetAttributes {
	return dataOrchestratedVirtualMachineScaleSetAttributes{ref: terra.ReferenceDataResource(ovmss)}
}

// DataOrchestratedVirtualMachineScaleSetArgs contains the configurations for azurerm_orchestrated_virtual_machine_scale_set.
type DataOrchestratedVirtualMachineScaleSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []dataorchestratedvirtualmachinescaleset.Identity `hcl:"identity,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []dataorchestratedvirtualmachinescaleset.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataorchestratedvirtualmachinescaleset.Timeouts `hcl:"timeouts,block"`
}
type dataOrchestratedVirtualMachineScaleSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("resource_group_name"))
}

func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) Identity() terra.ListValue[dataorchestratedvirtualmachinescaleset.IdentityAttributes] {
	return terra.ReferenceAsList[dataorchestratedvirtualmachinescaleset.IdentityAttributes](ovmss.ref.Append("identity"))
}

func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) NetworkInterface() terra.ListValue[dataorchestratedvirtualmachinescaleset.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[dataorchestratedvirtualmachinescaleset.NetworkInterfaceAttributes](ovmss.ref.Append("network_interface"))
}

func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) Timeouts() dataorchestratedvirtualmachinescaleset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataorchestratedvirtualmachinescaleset.TimeoutsAttributes](ovmss.ref.Append("timeouts"))
}
