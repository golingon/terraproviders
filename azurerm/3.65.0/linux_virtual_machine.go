// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	linuxvirtualmachine "github.com/golingon/terraproviders/azurerm/3.65.0/linuxvirtualmachine"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLinuxVirtualMachine creates a new instance of [LinuxVirtualMachine].
func NewLinuxVirtualMachine(name string, args LinuxVirtualMachineArgs) *LinuxVirtualMachine {
	return &LinuxVirtualMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LinuxVirtualMachine)(nil)

// LinuxVirtualMachine represents the Terraform resource azurerm_linux_virtual_machine.
type LinuxVirtualMachine struct {
	Name      string
	Args      LinuxVirtualMachineArgs
	state     *linuxVirtualMachineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LinuxVirtualMachine].
func (lvm *LinuxVirtualMachine) Type() string {
	return "azurerm_linux_virtual_machine"
}

// LocalName returns the local name for [LinuxVirtualMachine].
func (lvm *LinuxVirtualMachine) LocalName() string {
	return lvm.Name
}

// Configuration returns the configuration (args) for [LinuxVirtualMachine].
func (lvm *LinuxVirtualMachine) Configuration() interface{} {
	return lvm.Args
}

// DependOn is used for other resources to depend on [LinuxVirtualMachine].
func (lvm *LinuxVirtualMachine) DependOn() terra.Reference {
	return terra.ReferenceResource(lvm)
}

// Dependencies returns the list of resources [LinuxVirtualMachine] depends_on.
func (lvm *LinuxVirtualMachine) Dependencies() terra.Dependencies {
	return lvm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LinuxVirtualMachine].
func (lvm *LinuxVirtualMachine) LifecycleManagement() *terra.Lifecycle {
	return lvm.Lifecycle
}

// Attributes returns the attributes for [LinuxVirtualMachine].
func (lvm *LinuxVirtualMachine) Attributes() linuxVirtualMachineAttributes {
	return linuxVirtualMachineAttributes{ref: terra.ReferenceResource(lvm)}
}

// ImportState imports the given attribute values into [LinuxVirtualMachine]'s state.
func (lvm *LinuxVirtualMachine) ImportState(av io.Reader) error {
	lvm.state = &linuxVirtualMachineState{}
	if err := json.NewDecoder(av).Decode(lvm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lvm.Type(), lvm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LinuxVirtualMachine] has state.
func (lvm *LinuxVirtualMachine) State() (*linuxVirtualMachineState, bool) {
	return lvm.state, lvm.state != nil
}

// StateMust returns the state for [LinuxVirtualMachine]. Panics if the state is nil.
func (lvm *LinuxVirtualMachine) StateMust() *linuxVirtualMachineState {
	if lvm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lvm.Type(), lvm.LocalName()))
	}
	return lvm.state
}

// LinuxVirtualMachineArgs contains the configurations for azurerm_linux_virtual_machine.
type LinuxVirtualMachineArgs struct {
	// AdminPassword: string, optional
	AdminPassword terra.StringValue `hcl:"admin_password,attr"`
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
	// DisablePasswordAuthentication: bool, optional
	DisablePasswordAuthentication terra.BoolValue `hcl:"disable_password_authentication,attr"`
	// EdgeZone: string, optional
	EdgeZone terra.StringValue `hcl:"edge_zone,attr"`
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
	// UserData: string, optional
	UserData terra.StringValue `hcl:"user_data,attr"`
	// VirtualMachineScaleSetId: string, optional
	VirtualMachineScaleSetId terra.StringValue `hcl:"virtual_machine_scale_set_id,attr"`
	// VtpmEnabled: bool, optional
	VtpmEnabled terra.BoolValue `hcl:"vtpm_enabled,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// AdditionalCapabilities: optional
	AdditionalCapabilities *linuxvirtualmachine.AdditionalCapabilities `hcl:"additional_capabilities,block"`
	// AdminSshKey: min=0
	AdminSshKey []linuxvirtualmachine.AdminSshKey `hcl:"admin_ssh_key,block" validate:"min=0"`
	// BootDiagnostics: optional
	BootDiagnostics *linuxvirtualmachine.BootDiagnostics `hcl:"boot_diagnostics,block"`
	// GalleryApplication: min=0,max=100
	GalleryApplication []linuxvirtualmachine.GalleryApplication `hcl:"gallery_application,block" validate:"min=0,max=100"`
	// Identity: optional
	Identity *linuxvirtualmachine.Identity `hcl:"identity,block"`
	// OsDisk: required
	OsDisk *linuxvirtualmachine.OsDisk `hcl:"os_disk,block" validate:"required"`
	// Plan: optional
	Plan *linuxvirtualmachine.Plan `hcl:"plan,block"`
	// Secret: min=0
	Secret []linuxvirtualmachine.Secret `hcl:"secret,block" validate:"min=0"`
	// SourceImageReference: optional
	SourceImageReference *linuxvirtualmachine.SourceImageReference `hcl:"source_image_reference,block"`
	// TerminationNotification: optional
	TerminationNotification *linuxvirtualmachine.TerminationNotification `hcl:"termination_notification,block"`
	// Timeouts: optional
	Timeouts *linuxvirtualmachine.Timeouts `hcl:"timeouts,block"`
}
type linuxVirtualMachineAttributes struct {
	ref terra.Reference
}

// AdminPassword returns a reference to field admin_password of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) AdminPassword() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("admin_password"))
}

// AdminUsername returns a reference to field admin_username of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("admin_username"))
}

// AllowExtensionOperations returns a reference to field allow_extension_operations of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) AllowExtensionOperations() terra.BoolValue {
	return terra.ReferenceAsBool(lvm.ref.Append("allow_extension_operations"))
}

// AvailabilitySetId returns a reference to field availability_set_id of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) AvailabilitySetId() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("availability_set_id"))
}

// CapacityReservationGroupId returns a reference to field capacity_reservation_group_id of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) CapacityReservationGroupId() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("capacity_reservation_group_id"))
}

// ComputerName returns a reference to field computer_name of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) ComputerName() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("computer_name"))
}

// CustomData returns a reference to field custom_data of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) CustomData() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("custom_data"))
}

// DedicatedHostGroupId returns a reference to field dedicated_host_group_id of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) DedicatedHostGroupId() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("dedicated_host_group_id"))
}

// DedicatedHostId returns a reference to field dedicated_host_id of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) DedicatedHostId() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("dedicated_host_id"))
}

// DisablePasswordAuthentication returns a reference to field disable_password_authentication of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) DisablePasswordAuthentication() terra.BoolValue {
	return terra.ReferenceAsBool(lvm.ref.Append("disable_password_authentication"))
}

// EdgeZone returns a reference to field edge_zone of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) EdgeZone() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("edge_zone"))
}

// EncryptionAtHostEnabled returns a reference to field encryption_at_host_enabled of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) EncryptionAtHostEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lvm.ref.Append("encryption_at_host_enabled"))
}

// EvictionPolicy returns a reference to field eviction_policy of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) EvictionPolicy() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("eviction_policy"))
}

// ExtensionsTimeBudget returns a reference to field extensions_time_budget of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) ExtensionsTimeBudget() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("extensions_time_budget"))
}

// Id returns a reference to field id of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("location"))
}

// MaxBidPrice returns a reference to field max_bid_price of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) MaxBidPrice() terra.NumberValue {
	return terra.ReferenceAsNumber(lvm.ref.Append("max_bid_price"))
}

// Name returns a reference to field name of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("name"))
}

// NetworkInterfaceIds returns a reference to field network_interface_ids of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) NetworkInterfaceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lvm.ref.Append("network_interface_ids"))
}

// PatchAssessmentMode returns a reference to field patch_assessment_mode of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) PatchAssessmentMode() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("patch_assessment_mode"))
}

// PatchMode returns a reference to field patch_mode of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) PatchMode() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("patch_mode"))
}

// PlatformFaultDomain returns a reference to field platform_fault_domain of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) PlatformFaultDomain() terra.NumberValue {
	return terra.ReferenceAsNumber(lvm.ref.Append("platform_fault_domain"))
}

// Priority returns a reference to field priority of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) Priority() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("priority"))
}

// PrivateIpAddress returns a reference to field private_ip_address of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("private_ip_address"))
}

// PrivateIpAddresses returns a reference to field private_ip_addresses of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lvm.ref.Append("private_ip_addresses"))
}

// ProvisionVmAgent returns a reference to field provision_vm_agent of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) ProvisionVmAgent() terra.BoolValue {
	return terra.ReferenceAsBool(lvm.ref.Append("provision_vm_agent"))
}

// ProximityPlacementGroupId returns a reference to field proximity_placement_group_id of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) ProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("proximity_placement_group_id"))
}

// PublicIpAddress returns a reference to field public_ip_address of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) PublicIpAddress() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("public_ip_address"))
}

// PublicIpAddresses returns a reference to field public_ip_addresses of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) PublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lvm.ref.Append("public_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("resource_group_name"))
}

// SecureBootEnabled returns a reference to field secure_boot_enabled of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) SecureBootEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lvm.ref.Append("secure_boot_enabled"))
}

// Size returns a reference to field size of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) Size() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("size"))
}

// SourceImageId returns a reference to field source_image_id of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) SourceImageId() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("source_image_id"))
}

// Tags returns a reference to field tags of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lvm.ref.Append("tags"))
}

// UserData returns a reference to field user_data of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) UserData() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("user_data"))
}

// VirtualMachineId returns a reference to field virtual_machine_id of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) VirtualMachineId() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("virtual_machine_id"))
}

// VirtualMachineScaleSetId returns a reference to field virtual_machine_scale_set_id of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) VirtualMachineScaleSetId() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("virtual_machine_scale_set_id"))
}

// VtpmEnabled returns a reference to field vtpm_enabled of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) VtpmEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lvm.ref.Append("vtpm_enabled"))
}

// Zone returns a reference to field zone of azurerm_linux_virtual_machine.
func (lvm linuxVirtualMachineAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(lvm.ref.Append("zone"))
}

func (lvm linuxVirtualMachineAttributes) AdditionalCapabilities() terra.ListValue[linuxvirtualmachine.AdditionalCapabilitiesAttributes] {
	return terra.ReferenceAsList[linuxvirtualmachine.AdditionalCapabilitiesAttributes](lvm.ref.Append("additional_capabilities"))
}

func (lvm linuxVirtualMachineAttributes) AdminSshKey() terra.SetValue[linuxvirtualmachine.AdminSshKeyAttributes] {
	return terra.ReferenceAsSet[linuxvirtualmachine.AdminSshKeyAttributes](lvm.ref.Append("admin_ssh_key"))
}

func (lvm linuxVirtualMachineAttributes) BootDiagnostics() terra.ListValue[linuxvirtualmachine.BootDiagnosticsAttributes] {
	return terra.ReferenceAsList[linuxvirtualmachine.BootDiagnosticsAttributes](lvm.ref.Append("boot_diagnostics"))
}

func (lvm linuxVirtualMachineAttributes) GalleryApplication() terra.ListValue[linuxvirtualmachine.GalleryApplicationAttributes] {
	return terra.ReferenceAsList[linuxvirtualmachine.GalleryApplicationAttributes](lvm.ref.Append("gallery_application"))
}

func (lvm linuxVirtualMachineAttributes) Identity() terra.ListValue[linuxvirtualmachine.IdentityAttributes] {
	return terra.ReferenceAsList[linuxvirtualmachine.IdentityAttributes](lvm.ref.Append("identity"))
}

func (lvm linuxVirtualMachineAttributes) OsDisk() terra.ListValue[linuxvirtualmachine.OsDiskAttributes] {
	return terra.ReferenceAsList[linuxvirtualmachine.OsDiskAttributes](lvm.ref.Append("os_disk"))
}

func (lvm linuxVirtualMachineAttributes) Plan() terra.ListValue[linuxvirtualmachine.PlanAttributes] {
	return terra.ReferenceAsList[linuxvirtualmachine.PlanAttributes](lvm.ref.Append("plan"))
}

func (lvm linuxVirtualMachineAttributes) Secret() terra.ListValue[linuxvirtualmachine.SecretAttributes] {
	return terra.ReferenceAsList[linuxvirtualmachine.SecretAttributes](lvm.ref.Append("secret"))
}

func (lvm linuxVirtualMachineAttributes) SourceImageReference() terra.ListValue[linuxvirtualmachine.SourceImageReferenceAttributes] {
	return terra.ReferenceAsList[linuxvirtualmachine.SourceImageReferenceAttributes](lvm.ref.Append("source_image_reference"))
}

func (lvm linuxVirtualMachineAttributes) TerminationNotification() terra.ListValue[linuxvirtualmachine.TerminationNotificationAttributes] {
	return terra.ReferenceAsList[linuxvirtualmachine.TerminationNotificationAttributes](lvm.ref.Append("termination_notification"))
}

func (lvm linuxVirtualMachineAttributes) Timeouts() linuxvirtualmachine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[linuxvirtualmachine.TimeoutsAttributes](lvm.ref.Append("timeouts"))
}

type linuxVirtualMachineState struct {
	AdminPassword                 string                                             `json:"admin_password"`
	AdminUsername                 string                                             `json:"admin_username"`
	AllowExtensionOperations      bool                                               `json:"allow_extension_operations"`
	AvailabilitySetId             string                                             `json:"availability_set_id"`
	CapacityReservationGroupId    string                                             `json:"capacity_reservation_group_id"`
	ComputerName                  string                                             `json:"computer_name"`
	CustomData                    string                                             `json:"custom_data"`
	DedicatedHostGroupId          string                                             `json:"dedicated_host_group_id"`
	DedicatedHostId               string                                             `json:"dedicated_host_id"`
	DisablePasswordAuthentication bool                                               `json:"disable_password_authentication"`
	EdgeZone                      string                                             `json:"edge_zone"`
	EncryptionAtHostEnabled       bool                                               `json:"encryption_at_host_enabled"`
	EvictionPolicy                string                                             `json:"eviction_policy"`
	ExtensionsTimeBudget          string                                             `json:"extensions_time_budget"`
	Id                            string                                             `json:"id"`
	LicenseType                   string                                             `json:"license_type"`
	Location                      string                                             `json:"location"`
	MaxBidPrice                   float64                                            `json:"max_bid_price"`
	Name                          string                                             `json:"name"`
	NetworkInterfaceIds           []string                                           `json:"network_interface_ids"`
	PatchAssessmentMode           string                                             `json:"patch_assessment_mode"`
	PatchMode                     string                                             `json:"patch_mode"`
	PlatformFaultDomain           float64                                            `json:"platform_fault_domain"`
	Priority                      string                                             `json:"priority"`
	PrivateIpAddress              string                                             `json:"private_ip_address"`
	PrivateIpAddresses            []string                                           `json:"private_ip_addresses"`
	ProvisionVmAgent              bool                                               `json:"provision_vm_agent"`
	ProximityPlacementGroupId     string                                             `json:"proximity_placement_group_id"`
	PublicIpAddress               string                                             `json:"public_ip_address"`
	PublicIpAddresses             []string                                           `json:"public_ip_addresses"`
	ResourceGroupName             string                                             `json:"resource_group_name"`
	SecureBootEnabled             bool                                               `json:"secure_boot_enabled"`
	Size                          string                                             `json:"size"`
	SourceImageId                 string                                             `json:"source_image_id"`
	Tags                          map[string]string                                  `json:"tags"`
	UserData                      string                                             `json:"user_data"`
	VirtualMachineId              string                                             `json:"virtual_machine_id"`
	VirtualMachineScaleSetId      string                                             `json:"virtual_machine_scale_set_id"`
	VtpmEnabled                   bool                                               `json:"vtpm_enabled"`
	Zone                          string                                             `json:"zone"`
	AdditionalCapabilities        []linuxvirtualmachine.AdditionalCapabilitiesState  `json:"additional_capabilities"`
	AdminSshKey                   []linuxvirtualmachine.AdminSshKeyState             `json:"admin_ssh_key"`
	BootDiagnostics               []linuxvirtualmachine.BootDiagnosticsState         `json:"boot_diagnostics"`
	GalleryApplication            []linuxvirtualmachine.GalleryApplicationState      `json:"gallery_application"`
	Identity                      []linuxvirtualmachine.IdentityState                `json:"identity"`
	OsDisk                        []linuxvirtualmachine.OsDiskState                  `json:"os_disk"`
	Plan                          []linuxvirtualmachine.PlanState                    `json:"plan"`
	Secret                        []linuxvirtualmachine.SecretState                  `json:"secret"`
	SourceImageReference          []linuxvirtualmachine.SourceImageReferenceState    `json:"source_image_reference"`
	TerminationNotification       []linuxvirtualmachine.TerminationNotificationState `json:"termination_notification"`
	Timeouts                      *linuxvirtualmachine.TimeoutsState                 `json:"timeouts"`
}
