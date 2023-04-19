// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurestack

import (
	"encoding/json"
	"fmt"
	windowsvirtualmachine "github.com/golingon/terraproviders/azurestack/1.0.0/windowsvirtualmachine"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWindowsVirtualMachine creates a new instance of [WindowsVirtualMachine].
func NewWindowsVirtualMachine(name string, args WindowsVirtualMachineArgs) *WindowsVirtualMachine {
	return &WindowsVirtualMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WindowsVirtualMachine)(nil)

// WindowsVirtualMachine represents the Terraform resource azurestack_windows_virtual_machine.
type WindowsVirtualMachine struct {
	Name      string
	Args      WindowsVirtualMachineArgs
	state     *windowsVirtualMachineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WindowsVirtualMachine].
func (wvm *WindowsVirtualMachine) Type() string {
	return "azurestack_windows_virtual_machine"
}

// LocalName returns the local name for [WindowsVirtualMachine].
func (wvm *WindowsVirtualMachine) LocalName() string {
	return wvm.Name
}

// Configuration returns the configuration (args) for [WindowsVirtualMachine].
func (wvm *WindowsVirtualMachine) Configuration() interface{} {
	return wvm.Args
}

// DependOn is used for other resources to depend on [WindowsVirtualMachine].
func (wvm *WindowsVirtualMachine) DependOn() terra.Reference {
	return terra.ReferenceResource(wvm)
}

// Dependencies returns the list of resources [WindowsVirtualMachine] depends_on.
func (wvm *WindowsVirtualMachine) Dependencies() terra.Dependencies {
	return wvm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WindowsVirtualMachine].
func (wvm *WindowsVirtualMachine) LifecycleManagement() *terra.Lifecycle {
	return wvm.Lifecycle
}

// Attributes returns the attributes for [WindowsVirtualMachine].
func (wvm *WindowsVirtualMachine) Attributes() windowsVirtualMachineAttributes {
	return windowsVirtualMachineAttributes{ref: terra.ReferenceResource(wvm)}
}

// ImportState imports the given attribute values into [WindowsVirtualMachine]'s state.
func (wvm *WindowsVirtualMachine) ImportState(av io.Reader) error {
	wvm.state = &windowsVirtualMachineState{}
	if err := json.NewDecoder(av).Decode(wvm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wvm.Type(), wvm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WindowsVirtualMachine] has state.
func (wvm *WindowsVirtualMachine) State() (*windowsVirtualMachineState, bool) {
	return wvm.state, wvm.state != nil
}

// StateMust returns the state for [WindowsVirtualMachine]. Panics if the state is nil.
func (wvm *WindowsVirtualMachine) StateMust() *windowsVirtualMachineState {
	if wvm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wvm.Type(), wvm.LocalName()))
	}
	return wvm.state
}

// WindowsVirtualMachineArgs contains the configurations for azurestack_windows_virtual_machine.
type WindowsVirtualMachineArgs struct {
	// AdminPassword: string, required
	AdminPassword terra.StringValue `hcl:"admin_password,attr" validate:"required"`
	// AdminUsername: string, required
	AdminUsername terra.StringValue `hcl:"admin_username,attr" validate:"required"`
	// AllowExtensionOperations: bool, optional
	AllowExtensionOperations terra.BoolValue `hcl:"allow_extension_operations,attr"`
	// AvailabilitySetId: string, optional
	AvailabilitySetId terra.StringValue `hcl:"availability_set_id,attr"`
	// ComputerName: string, optional
	ComputerName terra.StringValue `hcl:"computer_name,attr"`
	// CustomData: string, optional
	CustomData terra.StringValue `hcl:"custom_data,attr"`
	// EnableAutomaticUpdates: bool, optional
	EnableAutomaticUpdates terra.BoolValue `hcl:"enable_automatic_updates,attr"`
	// EncryptionAtHostEnabled: bool, optional
	EncryptionAtHostEnabled terra.BoolValue `hcl:"encryption_at_host_enabled,attr"`
	// EvictionPolicy: string, optional
	EvictionPolicy terra.StringValue `hcl:"eviction_policy,attr"`
	// ExtensionsTimeBudget: string, optional
	ExtensionsTimeBudget terra.StringValue `hcl:"extensions_time_budget,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LicenseType: string, optional
	LicenseType terra.StringValue `hcl:"license_type,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaxBidPrice: number, optional
	MaxBidPrice terra.NumberValue `hcl:"max_bid_price,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkInterfaceIds: list of string, required
	NetworkInterfaceIds terra.ListValue[terra.StringValue] `hcl:"network_interface_ids,attr" validate:"required"`
	// PatchMode: string, optional
	PatchMode terra.StringValue `hcl:"patch_mode,attr"`
	// Priority: string, optional
	Priority terra.StringValue `hcl:"priority,attr"`
	// ProvisionVmAgent: bool, optional
	ProvisionVmAgent terra.BoolValue `hcl:"provision_vm_agent,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Size: string, required
	Size terra.StringValue `hcl:"size,attr" validate:"required"`
	// SourceImageId: string, optional
	SourceImageId terra.StringValue `hcl:"source_image_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timezone: string, optional
	Timezone terra.StringValue `hcl:"timezone,attr"`
	// VirtualMachineScaleSetId: string, optional
	VirtualMachineScaleSetId terra.StringValue `hcl:"virtual_machine_scale_set_id,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// AdditionalCapabilities: optional
	AdditionalCapabilities *windowsvirtualmachine.AdditionalCapabilities `hcl:"additional_capabilities,block"`
	// AdditionalUnattendContent: min=0
	AdditionalUnattendContent []windowsvirtualmachine.AdditionalUnattendContent `hcl:"additional_unattend_content,block" validate:"min=0"`
	// BootDiagnostics: optional
	BootDiagnostics *windowsvirtualmachine.BootDiagnostics `hcl:"boot_diagnostics,block"`
	// OsDisk: required
	OsDisk *windowsvirtualmachine.OsDisk `hcl:"os_disk,block" validate:"required"`
	// Plan: optional
	Plan *windowsvirtualmachine.Plan `hcl:"plan,block"`
	// Secret: min=0
	Secret []windowsvirtualmachine.Secret `hcl:"secret,block" validate:"min=0"`
	// SourceImageReference: optional
	SourceImageReference *windowsvirtualmachine.SourceImageReference `hcl:"source_image_reference,block"`
	// Timeouts: optional
	Timeouts *windowsvirtualmachine.Timeouts `hcl:"timeouts,block"`
	// WinrmListener: min=0
	WinrmListener []windowsvirtualmachine.WinrmListener `hcl:"winrm_listener,block" validate:"min=0"`
}
type windowsVirtualMachineAttributes struct {
	ref terra.Reference
}

// AdminPassword returns a reference to field admin_password of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) AdminPassword() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("admin_password"))
}

// AdminUsername returns a reference to field admin_username of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("admin_username"))
}

// AllowExtensionOperations returns a reference to field allow_extension_operations of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) AllowExtensionOperations() terra.BoolValue {
	return terra.ReferenceAsBool(wvm.ref.Append("allow_extension_operations"))
}

// AvailabilitySetId returns a reference to field availability_set_id of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) AvailabilitySetId() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("availability_set_id"))
}

// ComputerName returns a reference to field computer_name of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) ComputerName() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("computer_name"))
}

// CustomData returns a reference to field custom_data of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) CustomData() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("custom_data"))
}

// EnableAutomaticUpdates returns a reference to field enable_automatic_updates of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) EnableAutomaticUpdates() terra.BoolValue {
	return terra.ReferenceAsBool(wvm.ref.Append("enable_automatic_updates"))
}

// EncryptionAtHostEnabled returns a reference to field encryption_at_host_enabled of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) EncryptionAtHostEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wvm.ref.Append("encryption_at_host_enabled"))
}

// EvictionPolicy returns a reference to field eviction_policy of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) EvictionPolicy() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("eviction_policy"))
}

// ExtensionsTimeBudget returns a reference to field extensions_time_budget of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) ExtensionsTimeBudget() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("extensions_time_budget"))
}

// Id returns a reference to field id of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("license_type"))
}

// Location returns a reference to field location of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("location"))
}

// MaxBidPrice returns a reference to field max_bid_price of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) MaxBidPrice() terra.NumberValue {
	return terra.ReferenceAsNumber(wvm.ref.Append("max_bid_price"))
}

// Name returns a reference to field name of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("name"))
}

// NetworkInterfaceIds returns a reference to field network_interface_ids of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) NetworkInterfaceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wvm.ref.Append("network_interface_ids"))
}

// PatchMode returns a reference to field patch_mode of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) PatchMode() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("patch_mode"))
}

// Priority returns a reference to field priority of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) Priority() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("priority"))
}

// PrivateIpAddress returns a reference to field private_ip_address of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("private_ip_address"))
}

// PrivateIpAddresses returns a reference to field private_ip_addresses of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wvm.ref.Append("private_ip_addresses"))
}

// ProvisionVmAgent returns a reference to field provision_vm_agent of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) ProvisionVmAgent() terra.BoolValue {
	return terra.ReferenceAsBool(wvm.ref.Append("provision_vm_agent"))
}

// PublicIpAddress returns a reference to field public_ip_address of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) PublicIpAddress() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("public_ip_address"))
}

// PublicIpAddresses returns a reference to field public_ip_addresses of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) PublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wvm.ref.Append("public_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("resource_group_name"))
}

// Size returns a reference to field size of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) Size() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("size"))
}

// SourceImageId returns a reference to field source_image_id of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) SourceImageId() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("source_image_id"))
}

// Tags returns a reference to field tags of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wvm.ref.Append("tags"))
}

// Timezone returns a reference to field timezone of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("timezone"))
}

// VirtualMachineId returns a reference to field virtual_machine_id of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) VirtualMachineId() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("virtual_machine_id"))
}

// VirtualMachineScaleSetId returns a reference to field virtual_machine_scale_set_id of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) VirtualMachineScaleSetId() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("virtual_machine_scale_set_id"))
}

// Zone returns a reference to field zone of azurestack_windows_virtual_machine.
func (wvm windowsVirtualMachineAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(wvm.ref.Append("zone"))
}

func (wvm windowsVirtualMachineAttributes) AdditionalCapabilities() terra.ListValue[windowsvirtualmachine.AdditionalCapabilitiesAttributes] {
	return terra.ReferenceAsList[windowsvirtualmachine.AdditionalCapabilitiesAttributes](wvm.ref.Append("additional_capabilities"))
}

func (wvm windowsVirtualMachineAttributes) AdditionalUnattendContent() terra.ListValue[windowsvirtualmachine.AdditionalUnattendContentAttributes] {
	return terra.ReferenceAsList[windowsvirtualmachine.AdditionalUnattendContentAttributes](wvm.ref.Append("additional_unattend_content"))
}

func (wvm windowsVirtualMachineAttributes) BootDiagnostics() terra.ListValue[windowsvirtualmachine.BootDiagnosticsAttributes] {
	return terra.ReferenceAsList[windowsvirtualmachine.BootDiagnosticsAttributes](wvm.ref.Append("boot_diagnostics"))
}

func (wvm windowsVirtualMachineAttributes) OsDisk() terra.ListValue[windowsvirtualmachine.OsDiskAttributes] {
	return terra.ReferenceAsList[windowsvirtualmachine.OsDiskAttributes](wvm.ref.Append("os_disk"))
}

func (wvm windowsVirtualMachineAttributes) Plan() terra.ListValue[windowsvirtualmachine.PlanAttributes] {
	return terra.ReferenceAsList[windowsvirtualmachine.PlanAttributes](wvm.ref.Append("plan"))
}

func (wvm windowsVirtualMachineAttributes) Secret() terra.ListValue[windowsvirtualmachine.SecretAttributes] {
	return terra.ReferenceAsList[windowsvirtualmachine.SecretAttributes](wvm.ref.Append("secret"))
}

func (wvm windowsVirtualMachineAttributes) SourceImageReference() terra.ListValue[windowsvirtualmachine.SourceImageReferenceAttributes] {
	return terra.ReferenceAsList[windowsvirtualmachine.SourceImageReferenceAttributes](wvm.ref.Append("source_image_reference"))
}

func (wvm windowsVirtualMachineAttributes) Timeouts() windowsvirtualmachine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[windowsvirtualmachine.TimeoutsAttributes](wvm.ref.Append("timeouts"))
}

func (wvm windowsVirtualMachineAttributes) WinrmListener() terra.SetValue[windowsvirtualmachine.WinrmListenerAttributes] {
	return terra.ReferenceAsSet[windowsvirtualmachine.WinrmListenerAttributes](wvm.ref.Append("winrm_listener"))
}

type windowsVirtualMachineState struct {
	AdminPassword             string                                                 `json:"admin_password"`
	AdminUsername             string                                                 `json:"admin_username"`
	AllowExtensionOperations  bool                                                   `json:"allow_extension_operations"`
	AvailabilitySetId         string                                                 `json:"availability_set_id"`
	ComputerName              string                                                 `json:"computer_name"`
	CustomData                string                                                 `json:"custom_data"`
	EnableAutomaticUpdates    bool                                                   `json:"enable_automatic_updates"`
	EncryptionAtHostEnabled   bool                                                   `json:"encryption_at_host_enabled"`
	EvictionPolicy            string                                                 `json:"eviction_policy"`
	ExtensionsTimeBudget      string                                                 `json:"extensions_time_budget"`
	Id                        string                                                 `json:"id"`
	LicenseType               string                                                 `json:"license_type"`
	Location                  string                                                 `json:"location"`
	MaxBidPrice               float64                                                `json:"max_bid_price"`
	Name                      string                                                 `json:"name"`
	NetworkInterfaceIds       []string                                               `json:"network_interface_ids"`
	PatchMode                 string                                                 `json:"patch_mode"`
	Priority                  string                                                 `json:"priority"`
	PrivateIpAddress          string                                                 `json:"private_ip_address"`
	PrivateIpAddresses        []string                                               `json:"private_ip_addresses"`
	ProvisionVmAgent          bool                                                   `json:"provision_vm_agent"`
	PublicIpAddress           string                                                 `json:"public_ip_address"`
	PublicIpAddresses         []string                                               `json:"public_ip_addresses"`
	ResourceGroupName         string                                                 `json:"resource_group_name"`
	Size                      string                                                 `json:"size"`
	SourceImageId             string                                                 `json:"source_image_id"`
	Tags                      map[string]string                                      `json:"tags"`
	Timezone                  string                                                 `json:"timezone"`
	VirtualMachineId          string                                                 `json:"virtual_machine_id"`
	VirtualMachineScaleSetId  string                                                 `json:"virtual_machine_scale_set_id"`
	Zone                      string                                                 `json:"zone"`
	AdditionalCapabilities    []windowsvirtualmachine.AdditionalCapabilitiesState    `json:"additional_capabilities"`
	AdditionalUnattendContent []windowsvirtualmachine.AdditionalUnattendContentState `json:"additional_unattend_content"`
	BootDiagnostics           []windowsvirtualmachine.BootDiagnosticsState           `json:"boot_diagnostics"`
	OsDisk                    []windowsvirtualmachine.OsDiskState                    `json:"os_disk"`
	Plan                      []windowsvirtualmachine.PlanState                      `json:"plan"`
	Secret                    []windowsvirtualmachine.SecretState                    `json:"secret"`
	SourceImageReference      []windowsvirtualmachine.SourceImageReferenceState      `json:"source_image_reference"`
	Timeouts                  *windowsvirtualmachine.TimeoutsState                   `json:"timeouts"`
	WinrmListener             []windowsvirtualmachine.WinrmListenerState             `json:"winrm_listener"`
}
