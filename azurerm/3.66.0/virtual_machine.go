// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualmachine "github.com/golingon/terraproviders/azurerm/3.66.0/virtualmachine"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualMachine creates a new instance of [VirtualMachine].
func NewVirtualMachine(name string, args VirtualMachineArgs) *VirtualMachine {
	return &VirtualMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualMachine)(nil)

// VirtualMachine represents the Terraform resource azurerm_virtual_machine.
type VirtualMachine struct {
	Name      string
	Args      VirtualMachineArgs
	state     *virtualMachineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualMachine].
func (vm *VirtualMachine) Type() string {
	return "azurerm_virtual_machine"
}

// LocalName returns the local name for [VirtualMachine].
func (vm *VirtualMachine) LocalName() string {
	return vm.Name
}

// Configuration returns the configuration (args) for [VirtualMachine].
func (vm *VirtualMachine) Configuration() interface{} {
	return vm.Args
}

// DependOn is used for other resources to depend on [VirtualMachine].
func (vm *VirtualMachine) DependOn() terra.Reference {
	return terra.ReferenceResource(vm)
}

// Dependencies returns the list of resources [VirtualMachine] depends_on.
func (vm *VirtualMachine) Dependencies() terra.Dependencies {
	return vm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualMachine].
func (vm *VirtualMachine) LifecycleManagement() *terra.Lifecycle {
	return vm.Lifecycle
}

// Attributes returns the attributes for [VirtualMachine].
func (vm *VirtualMachine) Attributes() virtualMachineAttributes {
	return virtualMachineAttributes{ref: terra.ReferenceResource(vm)}
}

// ImportState imports the given attribute values into [VirtualMachine]'s state.
func (vm *VirtualMachine) ImportState(av io.Reader) error {
	vm.state = &virtualMachineState{}
	if err := json.NewDecoder(av).Decode(vm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vm.Type(), vm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualMachine] has state.
func (vm *VirtualMachine) State() (*virtualMachineState, bool) {
	return vm.state, vm.state != nil
}

// StateMust returns the state for [VirtualMachine]. Panics if the state is nil.
func (vm *VirtualMachine) StateMust() *virtualMachineState {
	if vm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vm.Type(), vm.LocalName()))
	}
	return vm.state
}

// VirtualMachineArgs contains the configurations for azurerm_virtual_machine.
type VirtualMachineArgs struct {
	// AvailabilitySetId: string, optional
	AvailabilitySetId terra.StringValue `hcl:"availability_set_id,attr"`
	// DeleteDataDisksOnTermination: bool, optional
	DeleteDataDisksOnTermination terra.BoolValue `hcl:"delete_data_disks_on_termination,attr"`
	// DeleteOsDiskOnTermination: bool, optional
	DeleteOsDiskOnTermination terra.BoolValue `hcl:"delete_os_disk_on_termination,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LicenseType: string, optional
	LicenseType terra.StringValue `hcl:"license_type,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkInterfaceIds: list of string, required
	NetworkInterfaceIds terra.ListValue[terra.StringValue] `hcl:"network_interface_ids,attr" validate:"required"`
	// PrimaryNetworkInterfaceId: string, optional
	PrimaryNetworkInterfaceId terra.StringValue `hcl:"primary_network_interface_id,attr"`
	// ProximityPlacementGroupId: string, optional
	ProximityPlacementGroupId terra.StringValue `hcl:"proximity_placement_group_id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VmSize: string, required
	VmSize terra.StringValue `hcl:"vm_size,attr" validate:"required"`
	// Zones: list of string, optional
	Zones terra.ListValue[terra.StringValue] `hcl:"zones,attr"`
	// AdditionalCapabilities: optional
	AdditionalCapabilities *virtualmachine.AdditionalCapabilities `hcl:"additional_capabilities,block"`
	// BootDiagnostics: optional
	BootDiagnostics *virtualmachine.BootDiagnostics `hcl:"boot_diagnostics,block"`
	// Identity: optional
	Identity *virtualmachine.Identity `hcl:"identity,block"`
	// OsProfile: optional
	OsProfile *virtualmachine.OsProfile `hcl:"os_profile,block"`
	// OsProfileLinuxConfig: optional
	OsProfileLinuxConfig *virtualmachine.OsProfileLinuxConfig `hcl:"os_profile_linux_config,block"`
	// OsProfileSecrets: min=0
	OsProfileSecrets []virtualmachine.OsProfileSecrets `hcl:"os_profile_secrets,block" validate:"min=0"`
	// OsProfileWindowsConfig: optional
	OsProfileWindowsConfig *virtualmachine.OsProfileWindowsConfig `hcl:"os_profile_windows_config,block"`
	// Plan: optional
	Plan *virtualmachine.Plan `hcl:"plan,block"`
	// StorageDataDisk: min=0
	StorageDataDisk []virtualmachine.StorageDataDisk `hcl:"storage_data_disk,block" validate:"min=0"`
	// StorageImageReference: optional
	StorageImageReference *virtualmachine.StorageImageReference `hcl:"storage_image_reference,block"`
	// StorageOsDisk: required
	StorageOsDisk *virtualmachine.StorageOsDisk `hcl:"storage_os_disk,block" validate:"required"`
	// Timeouts: optional
	Timeouts *virtualmachine.Timeouts `hcl:"timeouts,block"`
}
type virtualMachineAttributes struct {
	ref terra.Reference
}

// AvailabilitySetId returns a reference to field availability_set_id of azurerm_virtual_machine.
func (vm virtualMachineAttributes) AvailabilitySetId() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("availability_set_id"))
}

// DeleteDataDisksOnTermination returns a reference to field delete_data_disks_on_termination of azurerm_virtual_machine.
func (vm virtualMachineAttributes) DeleteDataDisksOnTermination() terra.BoolValue {
	return terra.ReferenceAsBool(vm.ref.Append("delete_data_disks_on_termination"))
}

// DeleteOsDiskOnTermination returns a reference to field delete_os_disk_on_termination of azurerm_virtual_machine.
func (vm virtualMachineAttributes) DeleteOsDiskOnTermination() terra.BoolValue {
	return terra.ReferenceAsBool(vm.ref.Append("delete_os_disk_on_termination"))
}

// Id returns a reference to field id of azurerm_virtual_machine.
func (vm virtualMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_virtual_machine.
func (vm virtualMachineAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_virtual_machine.
func (vm virtualMachineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_machine.
func (vm virtualMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("name"))
}

// NetworkInterfaceIds returns a reference to field network_interface_ids of azurerm_virtual_machine.
func (vm virtualMachineAttributes) NetworkInterfaceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vm.ref.Append("network_interface_ids"))
}

// PrimaryNetworkInterfaceId returns a reference to field primary_network_interface_id of azurerm_virtual_machine.
func (vm virtualMachineAttributes) PrimaryNetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("primary_network_interface_id"))
}

// ProximityPlacementGroupId returns a reference to field proximity_placement_group_id of azurerm_virtual_machine.
func (vm virtualMachineAttributes) ProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("proximity_placement_group_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_machine.
func (vm virtualMachineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_virtual_machine.
func (vm virtualMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vm.ref.Append("tags"))
}

// VmSize returns a reference to field vm_size of azurerm_virtual_machine.
func (vm virtualMachineAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("vm_size"))
}

// Zones returns a reference to field zones of azurerm_virtual_machine.
func (vm virtualMachineAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vm.ref.Append("zones"))
}

func (vm virtualMachineAttributes) AdditionalCapabilities() terra.ListValue[virtualmachine.AdditionalCapabilitiesAttributes] {
	return terra.ReferenceAsList[virtualmachine.AdditionalCapabilitiesAttributes](vm.ref.Append("additional_capabilities"))
}

func (vm virtualMachineAttributes) BootDiagnostics() terra.ListValue[virtualmachine.BootDiagnosticsAttributes] {
	return terra.ReferenceAsList[virtualmachine.BootDiagnosticsAttributes](vm.ref.Append("boot_diagnostics"))
}

func (vm virtualMachineAttributes) Identity() terra.ListValue[virtualmachine.IdentityAttributes] {
	return terra.ReferenceAsList[virtualmachine.IdentityAttributes](vm.ref.Append("identity"))
}

func (vm virtualMachineAttributes) OsProfile() terra.SetValue[virtualmachine.OsProfileAttributes] {
	return terra.ReferenceAsSet[virtualmachine.OsProfileAttributes](vm.ref.Append("os_profile"))
}

func (vm virtualMachineAttributes) OsProfileLinuxConfig() terra.SetValue[virtualmachine.OsProfileLinuxConfigAttributes] {
	return terra.ReferenceAsSet[virtualmachine.OsProfileLinuxConfigAttributes](vm.ref.Append("os_profile_linux_config"))
}

func (vm virtualMachineAttributes) OsProfileSecrets() terra.ListValue[virtualmachine.OsProfileSecretsAttributes] {
	return terra.ReferenceAsList[virtualmachine.OsProfileSecretsAttributes](vm.ref.Append("os_profile_secrets"))
}

func (vm virtualMachineAttributes) OsProfileWindowsConfig() terra.SetValue[virtualmachine.OsProfileWindowsConfigAttributes] {
	return terra.ReferenceAsSet[virtualmachine.OsProfileWindowsConfigAttributes](vm.ref.Append("os_profile_windows_config"))
}

func (vm virtualMachineAttributes) Plan() terra.ListValue[virtualmachine.PlanAttributes] {
	return terra.ReferenceAsList[virtualmachine.PlanAttributes](vm.ref.Append("plan"))
}

func (vm virtualMachineAttributes) StorageDataDisk() terra.ListValue[virtualmachine.StorageDataDiskAttributes] {
	return terra.ReferenceAsList[virtualmachine.StorageDataDiskAttributes](vm.ref.Append("storage_data_disk"))
}

func (vm virtualMachineAttributes) StorageImageReference() terra.SetValue[virtualmachine.StorageImageReferenceAttributes] {
	return terra.ReferenceAsSet[virtualmachine.StorageImageReferenceAttributes](vm.ref.Append("storage_image_reference"))
}

func (vm virtualMachineAttributes) StorageOsDisk() terra.ListValue[virtualmachine.StorageOsDiskAttributes] {
	return terra.ReferenceAsList[virtualmachine.StorageOsDiskAttributes](vm.ref.Append("storage_os_disk"))
}

func (vm virtualMachineAttributes) Timeouts() virtualmachine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualmachine.TimeoutsAttributes](vm.ref.Append("timeouts"))
}

type virtualMachineState struct {
	AvailabilitySetId            string                                       `json:"availability_set_id"`
	DeleteDataDisksOnTermination bool                                         `json:"delete_data_disks_on_termination"`
	DeleteOsDiskOnTermination    bool                                         `json:"delete_os_disk_on_termination"`
	Id                           string                                       `json:"id"`
	LicenseType                  string                                       `json:"license_type"`
	Location                     string                                       `json:"location"`
	Name                         string                                       `json:"name"`
	NetworkInterfaceIds          []string                                     `json:"network_interface_ids"`
	PrimaryNetworkInterfaceId    string                                       `json:"primary_network_interface_id"`
	ProximityPlacementGroupId    string                                       `json:"proximity_placement_group_id"`
	ResourceGroupName            string                                       `json:"resource_group_name"`
	Tags                         map[string]string                            `json:"tags"`
	VmSize                       string                                       `json:"vm_size"`
	Zones                        []string                                     `json:"zones"`
	AdditionalCapabilities       []virtualmachine.AdditionalCapabilitiesState `json:"additional_capabilities"`
	BootDiagnostics              []virtualmachine.BootDiagnosticsState        `json:"boot_diagnostics"`
	Identity                     []virtualmachine.IdentityState               `json:"identity"`
	OsProfile                    []virtualmachine.OsProfileState              `json:"os_profile"`
	OsProfileLinuxConfig         []virtualmachine.OsProfileLinuxConfigState   `json:"os_profile_linux_config"`
	OsProfileSecrets             []virtualmachine.OsProfileSecretsState       `json:"os_profile_secrets"`
	OsProfileWindowsConfig       []virtualmachine.OsProfileWindowsConfigState `json:"os_profile_windows_config"`
	Plan                         []virtualmachine.PlanState                   `json:"plan"`
	StorageDataDisk              []virtualmachine.StorageDataDiskState        `json:"storage_data_disk"`
	StorageImageReference        []virtualmachine.StorageImageReferenceState  `json:"storage_image_reference"`
	StorageOsDisk                []virtualmachine.StorageOsDiskState          `json:"storage_os_disk"`
	Timeouts                     *virtualmachine.TimeoutsState                `json:"timeouts"`
}
