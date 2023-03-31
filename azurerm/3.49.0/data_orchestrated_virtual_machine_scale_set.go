// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataorchestratedvirtualmachinescaleset "github.com/golingon/terraproviders/azurerm/3.49.0/dataorchestratedvirtualmachinescaleset"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataOrchestratedVirtualMachineScaleSet(name string, args DataOrchestratedVirtualMachineScaleSetArgs) *DataOrchestratedVirtualMachineScaleSet {
	return &DataOrchestratedVirtualMachineScaleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOrchestratedVirtualMachineScaleSet)(nil)

type DataOrchestratedVirtualMachineScaleSet struct {
	Name string
	Args DataOrchestratedVirtualMachineScaleSetArgs
}

func (ovmss *DataOrchestratedVirtualMachineScaleSet) DataSource() string {
	return "azurerm_orchestrated_virtual_machine_scale_set"
}

func (ovmss *DataOrchestratedVirtualMachineScaleSet) LocalName() string {
	return ovmss.Name
}

func (ovmss *DataOrchestratedVirtualMachineScaleSet) Configuration() interface{} {
	return ovmss.Args
}

func (ovmss *DataOrchestratedVirtualMachineScaleSet) Attributes() dataOrchestratedVirtualMachineScaleSetAttributes {
	return dataOrchestratedVirtualMachineScaleSetAttributes{ref: terra.ReferenceDataResource(ovmss)}
}

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

func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ovmss.ref.Append("id"))
}

func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) Location() terra.StringValue {
	return terra.ReferenceString(ovmss.ref.Append("location"))
}

func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ovmss.ref.Append("name"))
}

func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(ovmss.ref.Append("resource_group_name"))
}

func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) Identity() terra.ListValue[dataorchestratedvirtualmachinescaleset.IdentityAttributes] {
	return terra.ReferenceList[dataorchestratedvirtualmachinescaleset.IdentityAttributes](ovmss.ref.Append("identity"))
}

func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) NetworkInterface() terra.ListValue[dataorchestratedvirtualmachinescaleset.NetworkInterfaceAttributes] {
	return terra.ReferenceList[dataorchestratedvirtualmachinescaleset.NetworkInterfaceAttributes](ovmss.ref.Append("network_interface"))
}

func (ovmss dataOrchestratedVirtualMachineScaleSetAttributes) Timeouts() dataorchestratedvirtualmachinescaleset.TimeoutsAttributes {
	return terra.ReferenceSingle[dataorchestratedvirtualmachinescaleset.TimeoutsAttributes](ovmss.ref.Append("timeouts"))
}
