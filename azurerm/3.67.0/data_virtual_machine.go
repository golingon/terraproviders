// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datavirtualmachine "github.com/golingon/terraproviders/azurerm/3.67.0/datavirtualmachine"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVirtualMachine creates a new instance of [DataVirtualMachine].
func NewDataVirtualMachine(name string, args DataVirtualMachineArgs) *DataVirtualMachine {
	return &DataVirtualMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVirtualMachine)(nil)

// DataVirtualMachine represents the Terraform data resource azurerm_virtual_machine.
type DataVirtualMachine struct {
	Name string
	Args DataVirtualMachineArgs
}

// DataSource returns the Terraform object type for [DataVirtualMachine].
func (vm *DataVirtualMachine) DataSource() string {
	return "azurerm_virtual_machine"
}

// LocalName returns the local name for [DataVirtualMachine].
func (vm *DataVirtualMachine) LocalName() string {
	return vm.Name
}

// Configuration returns the configuration (args) for [DataVirtualMachine].
func (vm *DataVirtualMachine) Configuration() interface{} {
	return vm.Args
}

// Attributes returns the attributes for [DataVirtualMachine].
func (vm *DataVirtualMachine) Attributes() dataVirtualMachineAttributes {
	return dataVirtualMachineAttributes{ref: terra.ReferenceDataResource(vm)}
}

// DataVirtualMachineArgs contains the configurations for azurerm_virtual_machine.
type DataVirtualMachineArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datavirtualmachine.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavirtualmachine.Timeouts `hcl:"timeouts,block"`
}
type dataVirtualMachineAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_virtual_machine.
func (vm dataVirtualMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_virtual_machine.
func (vm dataVirtualMachineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_machine.
func (vm dataVirtualMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("name"))
}

// PowerState returns a reference to field power_state of azurerm_virtual_machine.
func (vm dataVirtualMachineAttributes) PowerState() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("power_state"))
}

// PrivateIpAddress returns a reference to field private_ip_address of azurerm_virtual_machine.
func (vm dataVirtualMachineAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("private_ip_address"))
}

// PrivateIpAddresses returns a reference to field private_ip_addresses of azurerm_virtual_machine.
func (vm dataVirtualMachineAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vm.ref.Append("private_ip_addresses"))
}

// PublicIpAddress returns a reference to field public_ip_address of azurerm_virtual_machine.
func (vm dataVirtualMachineAttributes) PublicIpAddress() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("public_ip_address"))
}

// PublicIpAddresses returns a reference to field public_ip_addresses of azurerm_virtual_machine.
func (vm dataVirtualMachineAttributes) PublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vm.ref.Append("public_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_machine.
func (vm dataVirtualMachineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("resource_group_name"))
}

func (vm dataVirtualMachineAttributes) Identity() terra.ListValue[datavirtualmachine.IdentityAttributes] {
	return terra.ReferenceAsList[datavirtualmachine.IdentityAttributes](vm.ref.Append("identity"))
}

func (vm dataVirtualMachineAttributes) Timeouts() datavirtualmachine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavirtualmachine.TimeoutsAttributes](vm.ref.Append("timeouts"))
}