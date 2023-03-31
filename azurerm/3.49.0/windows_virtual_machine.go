// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	windowsvirtualmachine "github.com/golingon/terraproviders/azurerm/3.49.0/windowsvirtualmachine"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewWindowsVirtualMachine(name string, args WindowsVirtualMachineArgs) *WindowsVirtualMachine {
	return &WindowsVirtualMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WindowsVirtualMachine)(nil)

type WindowsVirtualMachine struct {
	Name  string
	Args  WindowsVirtualMachineArgs
	state *windowsVirtualMachineState
}

func (wvm *WindowsVirtualMachine) Type() string {
	return "azurerm_windows_virtual_machine"
}

func (wvm *WindowsVirtualMachine) LocalName() string {
	return wvm.Name
}

func (wvm *WindowsVirtualMachine) Configuration() interface{} {
	return wvm.Args
}

func (wvm *WindowsVirtualMachine) Attributes() windowsVirtualMachineAttributes {
	return windowsVirtualMachineAttributes{ref: terra.ReferenceResource(wvm)}
}

func (wvm *WindowsVirtualMachine) ImportState(av io.Reader) error {
	wvm.state = &windowsVirtualMachineState{}
	if err := json.NewDecoder(av).Decode(wvm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wvm.Type(), wvm.LocalName(), err)
	}
	return nil
}

func (wvm *WindowsVirtualMachine) State() (*windowsVirtualMachineState, bool) {
	return wvm.state, wvm.state != nil
}

func (wvm *WindowsVirtualMachine) StateMust() *windowsVirtualMachineState {
	if wvm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wvm.Type(), wvm.LocalName()))
	}
	return wvm.state
}

func (wvm *WindowsVirtualMachine) DependOn() terra.Reference {
	return terra.ReferenceResource(wvm)
}

type WindowsVirtualMachineArgs struct {
	// AdminPassword: string, required
	AdminPassword terra.StringValue `hcl:"admin_password,attr" validate:"required"`
	// AdminUsername: string, required
	AdminUsername terra.StringValue `hcl:"admin_username,attr" validate:"required"`
	// AllowExtensionOperations: bool, optional
	AllowExtensionOperations terra.BoolValue `hcl:"allow_extension_operations,attr"`
	// AvailabilitySetId: string, optional
	AvailabilitySetId terra.StringValue `hcl:"availability_set_id,attr"`
	// CapacityReservationGroupId: string, optional
	CapacityReservationGroupId terra.StringValue `hcl:"capacity_reservation_group_id,attr"`
	// ComputerName: string, optional
	ComputerName terra.StringValue `hcl:"computer_name,attr"`
	// CustomData: string, optional
	CustomData terra.StringValue `hcl:"custom_data,attr"`
	// DedicatedHostGroupId: string, optional
	DedicatedHostGroupId terra.StringValue `hcl:"dedicated_host_group_id,attr"`
	// DedicatedHostId: string, optional
	DedicatedHostId terra.StringValue `hcl:"dedicated_host_id,attr"`
	// EdgeZone: string, optional
	EdgeZone terra.StringValue `hcl:"edge_zone,attr"`
	// EnableAutomaticUpdates: bool, optional
	EnableAutomaticUpdates terra.BoolValue `hcl:"enable_automatic_updates,attr"`
	// EncryptionAtHostEnabled: bool, optional
	EncryptionAtHostEnabled terra.BoolValue `hcl:"encryption_at_host_enabled,attr"`
	// EvictionPolicy: string, optional
	EvictionPolicy terra.StringValue `hcl:"eviction_policy,attr"`
	// ExtensionsTimeBudget: string, optional
	ExtensionsTimeBudget terra.StringValue `hcl:"extensions_time_budget,attr"`
	// HotpatchingEnabled: bool, optional
	HotpatchingEnabled terra.BoolValue `hcl:"hotpatching_enabled,attr"`
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
	// PatchAssessmentMode: string, optional
	PatchAssessmentMode terra.StringValue `hcl:"patch_assessment_mode,attr"`
	// PatchMode: string, optional
	PatchMode terra.StringValue `hcl:"patch_mode,attr"`
	// PlatformFaultDomain: number, optional
	PlatformFaultDomain terra.NumberValue `hcl:"platform_fault_domain,attr"`
	// Priority: string, optional
	Priority terra.StringValue `hcl:"priority,attr"`
	// ProvisionVmAgent: bool, optional
	ProvisionVmAgent terra.BoolValue `hcl:"provision_vm_agent,attr"`
	// ProximityPlacementGroupId: string, optional
	ProximityPlacementGroupId terra.StringValue `hcl:"proximity_placement_group_id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SecureBootEnabled: bool, optional
	SecureBootEnabled terra.BoolValue `hcl:"secure_boot_enabled,attr"`
	// Size: string, required
	Size terra.StringValue `hcl:"size,attr" validate:"required"`
	// SourceImageId: string, optional
	SourceImageId terra.StringValue `hcl:"source_image_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timezone: string, optional
	Timezone terra.StringValue `hcl:"timezone,attr"`
	// UserData: string, optional
	UserData terra.StringValue `hcl:"user_data,attr"`
	// VirtualMachineScaleSetId: string, optional
	VirtualMachineScaleSetId terra.StringValue `hcl:"virtual_machine_scale_set_id,attr"`
	// VtpmEnabled: bool, optional
	VtpmEnabled terra.BoolValue `hcl:"vtpm_enabled,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// AdditionalCapabilities: optional
	AdditionalCapabilities *windowsvirtualmachine.AdditionalCapabilities `hcl:"additional_capabilities,block"`
	// AdditionalUnattendContent: min=0
	AdditionalUnattendContent []windowsvirtualmachine.AdditionalUnattendContent `hcl:"additional_unattend_content,block" validate:"min=0"`
	// BootDiagnostics: optional
	BootDiagnostics *windowsvirtualmachine.BootDiagnostics `hcl:"boot_diagnostics,block"`
	// GalleryApplication: min=0,max=100
	GalleryApplication []windowsvirtualmachine.GalleryApplication `hcl:"gallery_application,block" validate:"min=0,max=100"`
	// Identity: optional
	Identity *windowsvirtualmachine.Identity `hcl:"identity,block"`
	// OsDisk: required
	OsDisk *windowsvirtualmachine.OsDisk `hcl:"os_disk,block" validate:"required"`
	// Plan: optional
	Plan *windowsvirtualmachine.Plan `hcl:"plan,block"`
	// Secret: min=0
	Secret []windowsvirtualmachine.Secret `hcl:"secret,block" validate:"min=0"`
	// SourceImageReference: optional
	SourceImageReference *windowsvirtualmachine.SourceImageReference `hcl:"source_image_reference,block"`
	// TerminationNotification: optional
	TerminationNotification *windowsvirtualmachine.TerminationNotification `hcl:"termination_notification,block"`
	// Timeouts: optional
	Timeouts *windowsvirtualmachine.Timeouts `hcl:"timeouts,block"`
	// WinrmListener: min=0
	WinrmListener []windowsvirtualmachine.WinrmListener `hcl:"winrm_listener,block" validate:"min=0"`
	// DependsOn contains resources that WindowsVirtualMachine depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type windowsVirtualMachineAttributes struct {
	ref terra.Reference
}

func (wvm windowsVirtualMachineAttributes) AdminPassword() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("admin_password"))
}

func (wvm windowsVirtualMachineAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("admin_username"))
}

func (wvm windowsVirtualMachineAttributes) AllowExtensionOperations() terra.BoolValue {
	return terra.ReferenceBool(wvm.ref.Append("allow_extension_operations"))
}

func (wvm windowsVirtualMachineAttributes) AvailabilitySetId() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("availability_set_id"))
}

func (wvm windowsVirtualMachineAttributes) CapacityReservationGroupId() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("capacity_reservation_group_id"))
}

func (wvm windowsVirtualMachineAttributes) ComputerName() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("computer_name"))
}

func (wvm windowsVirtualMachineAttributes) CustomData() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("custom_data"))
}

func (wvm windowsVirtualMachineAttributes) DedicatedHostGroupId() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("dedicated_host_group_id"))
}

func (wvm windowsVirtualMachineAttributes) DedicatedHostId() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("dedicated_host_id"))
}

func (wvm windowsVirtualMachineAttributes) EdgeZone() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("edge_zone"))
}

func (wvm windowsVirtualMachineAttributes) EnableAutomaticUpdates() terra.BoolValue {
	return terra.ReferenceBool(wvm.ref.Append("enable_automatic_updates"))
}

func (wvm windowsVirtualMachineAttributes) EncryptionAtHostEnabled() terra.BoolValue {
	return terra.ReferenceBool(wvm.ref.Append("encryption_at_host_enabled"))
}

func (wvm windowsVirtualMachineAttributes) EvictionPolicy() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("eviction_policy"))
}

func (wvm windowsVirtualMachineAttributes) ExtensionsTimeBudget() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("extensions_time_budget"))
}

func (wvm windowsVirtualMachineAttributes) HotpatchingEnabled() terra.BoolValue {
	return terra.ReferenceBool(wvm.ref.Append("hotpatching_enabled"))
}

func (wvm windowsVirtualMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("id"))
}

func (wvm windowsVirtualMachineAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("license_type"))
}

func (wvm windowsVirtualMachineAttributes) Location() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("location"))
}

func (wvm windowsVirtualMachineAttributes) MaxBidPrice() terra.NumberValue {
	return terra.ReferenceNumber(wvm.ref.Append("max_bid_price"))
}

func (wvm windowsVirtualMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("name"))
}

func (wvm windowsVirtualMachineAttributes) NetworkInterfaceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](wvm.ref.Append("network_interface_ids"))
}

func (wvm windowsVirtualMachineAttributes) PatchAssessmentMode() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("patch_assessment_mode"))
}

func (wvm windowsVirtualMachineAttributes) PatchMode() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("patch_mode"))
}

func (wvm windowsVirtualMachineAttributes) PlatformFaultDomain() terra.NumberValue {
	return terra.ReferenceNumber(wvm.ref.Append("platform_fault_domain"))
}

func (wvm windowsVirtualMachineAttributes) Priority() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("priority"))
}

func (wvm windowsVirtualMachineAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("private_ip_address"))
}

func (wvm windowsVirtualMachineAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](wvm.ref.Append("private_ip_addresses"))
}

func (wvm windowsVirtualMachineAttributes) ProvisionVmAgent() terra.BoolValue {
	return terra.ReferenceBool(wvm.ref.Append("provision_vm_agent"))
}

func (wvm windowsVirtualMachineAttributes) ProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("proximity_placement_group_id"))
}

func (wvm windowsVirtualMachineAttributes) PublicIpAddress() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("public_ip_address"))
}

func (wvm windowsVirtualMachineAttributes) PublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](wvm.ref.Append("public_ip_addresses"))
}

func (wvm windowsVirtualMachineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("resource_group_name"))
}

func (wvm windowsVirtualMachineAttributes) SecureBootEnabled() terra.BoolValue {
	return terra.ReferenceBool(wvm.ref.Append("secure_boot_enabled"))
}

func (wvm windowsVirtualMachineAttributes) Size() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("size"))
}

func (wvm windowsVirtualMachineAttributes) SourceImageId() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("source_image_id"))
}

func (wvm windowsVirtualMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](wvm.ref.Append("tags"))
}

func (wvm windowsVirtualMachineAttributes) Timezone() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("timezone"))
}

func (wvm windowsVirtualMachineAttributes) UserData() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("user_data"))
}

func (wvm windowsVirtualMachineAttributes) VirtualMachineId() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("virtual_machine_id"))
}

func (wvm windowsVirtualMachineAttributes) VirtualMachineScaleSetId() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("virtual_machine_scale_set_id"))
}

func (wvm windowsVirtualMachineAttributes) VtpmEnabled() terra.BoolValue {
	return terra.ReferenceBool(wvm.ref.Append("vtpm_enabled"))
}

func (wvm windowsVirtualMachineAttributes) Zone() terra.StringValue {
	return terra.ReferenceString(wvm.ref.Append("zone"))
}

func (wvm windowsVirtualMachineAttributes) AdditionalCapabilities() terra.ListValue[windowsvirtualmachine.AdditionalCapabilitiesAttributes] {
	return terra.ReferenceList[windowsvirtualmachine.AdditionalCapabilitiesAttributes](wvm.ref.Append("additional_capabilities"))
}

func (wvm windowsVirtualMachineAttributes) AdditionalUnattendContent() terra.ListValue[windowsvirtualmachine.AdditionalUnattendContentAttributes] {
	return terra.ReferenceList[windowsvirtualmachine.AdditionalUnattendContentAttributes](wvm.ref.Append("additional_unattend_content"))
}

func (wvm windowsVirtualMachineAttributes) BootDiagnostics() terra.ListValue[windowsvirtualmachine.BootDiagnosticsAttributes] {
	return terra.ReferenceList[windowsvirtualmachine.BootDiagnosticsAttributes](wvm.ref.Append("boot_diagnostics"))
}

func (wvm windowsVirtualMachineAttributes) GalleryApplication() terra.ListValue[windowsvirtualmachine.GalleryApplicationAttributes] {
	return terra.ReferenceList[windowsvirtualmachine.GalleryApplicationAttributes](wvm.ref.Append("gallery_application"))
}

func (wvm windowsVirtualMachineAttributes) Identity() terra.ListValue[windowsvirtualmachine.IdentityAttributes] {
	return terra.ReferenceList[windowsvirtualmachine.IdentityAttributes](wvm.ref.Append("identity"))
}

func (wvm windowsVirtualMachineAttributes) OsDisk() terra.ListValue[windowsvirtualmachine.OsDiskAttributes] {
	return terra.ReferenceList[windowsvirtualmachine.OsDiskAttributes](wvm.ref.Append("os_disk"))
}

func (wvm windowsVirtualMachineAttributes) Plan() terra.ListValue[windowsvirtualmachine.PlanAttributes] {
	return terra.ReferenceList[windowsvirtualmachine.PlanAttributes](wvm.ref.Append("plan"))
}

func (wvm windowsVirtualMachineAttributes) Secret() terra.ListValue[windowsvirtualmachine.SecretAttributes] {
	return terra.ReferenceList[windowsvirtualmachine.SecretAttributes](wvm.ref.Append("secret"))
}

func (wvm windowsVirtualMachineAttributes) SourceImageReference() terra.ListValue[windowsvirtualmachine.SourceImageReferenceAttributes] {
	return terra.ReferenceList[windowsvirtualmachine.SourceImageReferenceAttributes](wvm.ref.Append("source_image_reference"))
}

func (wvm windowsVirtualMachineAttributes) TerminationNotification() terra.ListValue[windowsvirtualmachine.TerminationNotificationAttributes] {
	return terra.ReferenceList[windowsvirtualmachine.TerminationNotificationAttributes](wvm.ref.Append("termination_notification"))
}

func (wvm windowsVirtualMachineAttributes) Timeouts() windowsvirtualmachine.TimeoutsAttributes {
	return terra.ReferenceSingle[windowsvirtualmachine.TimeoutsAttributes](wvm.ref.Append("timeouts"))
}

func (wvm windowsVirtualMachineAttributes) WinrmListener() terra.SetValue[windowsvirtualmachine.WinrmListenerAttributes] {
	return terra.ReferenceSet[windowsvirtualmachine.WinrmListenerAttributes](wvm.ref.Append("winrm_listener"))
}

type windowsVirtualMachineState struct {
	AdminPassword              string                                                 `json:"admin_password"`
	AdminUsername              string                                                 `json:"admin_username"`
	AllowExtensionOperations   bool                                                   `json:"allow_extension_operations"`
	AvailabilitySetId          string                                                 `json:"availability_set_id"`
	CapacityReservationGroupId string                                                 `json:"capacity_reservation_group_id"`
	ComputerName               string                                                 `json:"computer_name"`
	CustomData                 string                                                 `json:"custom_data"`
	DedicatedHostGroupId       string                                                 `json:"dedicated_host_group_id"`
	DedicatedHostId            string                                                 `json:"dedicated_host_id"`
	EdgeZone                   string                                                 `json:"edge_zone"`
	EnableAutomaticUpdates     bool                                                   `json:"enable_automatic_updates"`
	EncryptionAtHostEnabled    bool                                                   `json:"encryption_at_host_enabled"`
	EvictionPolicy             string                                                 `json:"eviction_policy"`
	ExtensionsTimeBudget       string                                                 `json:"extensions_time_budget"`
	HotpatchingEnabled         bool                                                   `json:"hotpatching_enabled"`
	Id                         string                                                 `json:"id"`
	LicenseType                string                                                 `json:"license_type"`
	Location                   string                                                 `json:"location"`
	MaxBidPrice                float64                                                `json:"max_bid_price"`
	Name                       string                                                 `json:"name"`
	NetworkInterfaceIds        []string                                               `json:"network_interface_ids"`
	PatchAssessmentMode        string                                                 `json:"patch_assessment_mode"`
	PatchMode                  string                                                 `json:"patch_mode"`
	PlatformFaultDomain        float64                                                `json:"platform_fault_domain"`
	Priority                   string                                                 `json:"priority"`
	PrivateIpAddress           string                                                 `json:"private_ip_address"`
	PrivateIpAddresses         []string                                               `json:"private_ip_addresses"`
	ProvisionVmAgent           bool                                                   `json:"provision_vm_agent"`
	ProximityPlacementGroupId  string                                                 `json:"proximity_placement_group_id"`
	PublicIpAddress            string                                                 `json:"public_ip_address"`
	PublicIpAddresses          []string                                               `json:"public_ip_addresses"`
	ResourceGroupName          string                                                 `json:"resource_group_name"`
	SecureBootEnabled          bool                                                   `json:"secure_boot_enabled"`
	Size                       string                                                 `json:"size"`
	SourceImageId              string                                                 `json:"source_image_id"`
	Tags                       map[string]string                                      `json:"tags"`
	Timezone                   string                                                 `json:"timezone"`
	UserData                   string                                                 `json:"user_data"`
	VirtualMachineId           string                                                 `json:"virtual_machine_id"`
	VirtualMachineScaleSetId   string                                                 `json:"virtual_machine_scale_set_id"`
	VtpmEnabled                bool                                                   `json:"vtpm_enabled"`
	Zone                       string                                                 `json:"zone"`
	AdditionalCapabilities     []windowsvirtualmachine.AdditionalCapabilitiesState    `json:"additional_capabilities"`
	AdditionalUnattendContent  []windowsvirtualmachine.AdditionalUnattendContentState `json:"additional_unattend_content"`
	BootDiagnostics            []windowsvirtualmachine.BootDiagnosticsState           `json:"boot_diagnostics"`
	GalleryApplication         []windowsvirtualmachine.GalleryApplicationState        `json:"gallery_application"`
	Identity                   []windowsvirtualmachine.IdentityState                  `json:"identity"`
	OsDisk                     []windowsvirtualmachine.OsDiskState                    `json:"os_disk"`
	Plan                       []windowsvirtualmachine.PlanState                      `json:"plan"`
	Secret                     []windowsvirtualmachine.SecretState                    `json:"secret"`
	SourceImageReference       []windowsvirtualmachine.SourceImageReferenceState      `json:"source_image_reference"`
	TerminationNotification    []windowsvirtualmachine.TerminationNotificationState   `json:"termination_notification"`
	Timeouts                   *windowsvirtualmachine.TimeoutsState                   `json:"timeouts"`
	WinrmListener              []windowsvirtualmachine.WinrmListenerState             `json:"winrm_listener"`
}
