// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_windows_virtual_machine_scale_set

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_windows_virtual_machine_scale_set.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermWindowsVirtualMachineScaleSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (awvmss *Resource) Type() string {
	return "azurerm_windows_virtual_machine_scale_set"
}

// LocalName returns the local name for [Resource].
func (awvmss *Resource) LocalName() string {
	return awvmss.Name
}

// Configuration returns the configuration (args) for [Resource].
func (awvmss *Resource) Configuration() interface{} {
	return awvmss.Args
}

// DependOn is used for other resources to depend on [Resource].
func (awvmss *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(awvmss)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (awvmss *Resource) Dependencies() terra.Dependencies {
	return awvmss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (awvmss *Resource) LifecycleManagement() *terra.Lifecycle {
	return awvmss.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (awvmss *Resource) Attributes() azurermWindowsVirtualMachineScaleSetAttributes {
	return azurermWindowsVirtualMachineScaleSetAttributes{ref: terra.ReferenceResource(awvmss)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (awvmss *Resource) ImportState(state io.Reader) error {
	awvmss.state = &azurermWindowsVirtualMachineScaleSetState{}
	if err := json.NewDecoder(state).Decode(awvmss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", awvmss.Type(), awvmss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (awvmss *Resource) State() (*azurermWindowsVirtualMachineScaleSetState, bool) {
	return awvmss.state, awvmss.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (awvmss *Resource) StateMust() *azurermWindowsVirtualMachineScaleSetState {
	if awvmss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", awvmss.Type(), awvmss.LocalName()))
	}
	return awvmss.state
}

// Args contains the configurations for azurerm_windows_virtual_machine_scale_set.
type Args struct {
	// AdminPassword: string, required
	AdminPassword terra.StringValue `hcl:"admin_password,attr" validate:"required"`
	// AdminUsername: string, required
	AdminUsername terra.StringValue `hcl:"admin_username,attr" validate:"required"`
	// CapacityReservationGroupId: string, optional
	CapacityReservationGroupId terra.StringValue `hcl:"capacity_reservation_group_id,attr"`
	// ComputerNamePrefix: string, optional
	ComputerNamePrefix terra.StringValue `hcl:"computer_name_prefix,attr"`
	// CustomData: string, optional
	CustomData terra.StringValue `hcl:"custom_data,attr"`
	// DoNotRunExtensionsOnOverprovisionedMachines: bool, optional
	DoNotRunExtensionsOnOverprovisionedMachines terra.BoolValue `hcl:"do_not_run_extensions_on_overprovisioned_machines,attr"`
	// EdgeZone: string, optional
	EdgeZone terra.StringValue `hcl:"edge_zone,attr"`
	// EnableAutomaticUpdates: bool, optional
	EnableAutomaticUpdates terra.BoolValue `hcl:"enable_automatic_updates,attr"`
	// EncryptionAtHostEnabled: bool, optional
	EncryptionAtHostEnabled terra.BoolValue `hcl:"encryption_at_host_enabled,attr"`
	// EvictionPolicy: string, optional
	EvictionPolicy terra.StringValue `hcl:"eviction_policy,attr"`
	// ExtensionOperationsEnabled: bool, optional
	ExtensionOperationsEnabled terra.BoolValue `hcl:"extension_operations_enabled,attr"`
	// ExtensionsTimeBudget: string, optional
	ExtensionsTimeBudget terra.StringValue `hcl:"extensions_time_budget,attr"`
	// HealthProbeId: string, optional
	HealthProbeId terra.StringValue `hcl:"health_probe_id,attr"`
	// HostGroupId: string, optional
	HostGroupId terra.StringValue `hcl:"host_group_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instances: number, required
	Instances terra.NumberValue `hcl:"instances,attr" validate:"required"`
	// LicenseType: string, optional
	LicenseType terra.StringValue `hcl:"license_type,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaxBidPrice: number, optional
	MaxBidPrice terra.NumberValue `hcl:"max_bid_price,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Overprovision: bool, optional
	Overprovision terra.BoolValue `hcl:"overprovision,attr"`
	// PlatformFaultDomainCount: number, optional
	PlatformFaultDomainCount terra.NumberValue `hcl:"platform_fault_domain_count,attr"`
	// Priority: string, optional
	Priority terra.StringValue `hcl:"priority,attr"`
	// ProvisionVmAgent: bool, optional
	ProvisionVmAgent terra.BoolValue `hcl:"provision_vm_agent,attr"`
	// ProximityPlacementGroupId: string, optional
	ProximityPlacementGroupId terra.StringValue `hcl:"proximity_placement_group_id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ScaleInPolicy: string, optional
	ScaleInPolicy terra.StringValue `hcl:"scale_in_policy,attr"`
	// SecureBootEnabled: bool, optional
	SecureBootEnabled terra.BoolValue `hcl:"secure_boot_enabled,attr"`
	// SinglePlacementGroup: bool, optional
	SinglePlacementGroup terra.BoolValue `hcl:"single_placement_group,attr"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// SourceImageId: string, optional
	SourceImageId terra.StringValue `hcl:"source_image_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timezone: string, optional
	Timezone terra.StringValue `hcl:"timezone,attr"`
	// UpgradeMode: string, optional
	UpgradeMode terra.StringValue `hcl:"upgrade_mode,attr"`
	// UserData: string, optional
	UserData terra.StringValue `hcl:"user_data,attr"`
	// VtpmEnabled: bool, optional
	VtpmEnabled terra.BoolValue `hcl:"vtpm_enabled,attr"`
	// ZoneBalance: bool, optional
	ZoneBalance terra.BoolValue `hcl:"zone_balance,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// AdditionalCapabilities: optional
	AdditionalCapabilities *AdditionalCapabilities `hcl:"additional_capabilities,block"`
	// AdditionalUnattendContent: min=0
	AdditionalUnattendContent []AdditionalUnattendContent `hcl:"additional_unattend_content,block" validate:"min=0"`
	// AutomaticInstanceRepair: optional
	AutomaticInstanceRepair *AutomaticInstanceRepair `hcl:"automatic_instance_repair,block"`
	// AutomaticOsUpgradePolicy: optional
	AutomaticOsUpgradePolicy *AutomaticOsUpgradePolicy `hcl:"automatic_os_upgrade_policy,block"`
	// BootDiagnostics: optional
	BootDiagnostics *BootDiagnostics `hcl:"boot_diagnostics,block"`
	// DataDisk: min=0
	DataDisk []DataDisk `hcl:"data_disk,block" validate:"min=0"`
	// Extension: min=0
	Extension []Extension `hcl:"extension,block" validate:"min=0"`
	// GalleryApplication: min=0,max=100
	GalleryApplication []GalleryApplication `hcl:"gallery_application,block" validate:"min=0,max=100"`
	// GalleryApplications: min=0,max=100
	GalleryApplications []GalleryApplications `hcl:"gallery_applications,block" validate:"min=0,max=100"`
	// Identity: optional
	Identity *Identity `hcl:"identity,block"`
	// NetworkInterface: min=1
	NetworkInterface []NetworkInterface `hcl:"network_interface,block" validate:"min=1"`
	// OsDisk: required
	OsDisk *OsDisk `hcl:"os_disk,block" validate:"required"`
	// Plan: optional
	Plan *Plan `hcl:"plan,block"`
	// RollingUpgradePolicy: optional
	RollingUpgradePolicy *RollingUpgradePolicy `hcl:"rolling_upgrade_policy,block"`
	// ScaleIn: optional
	ScaleIn *ScaleIn `hcl:"scale_in,block"`
	// Secret: min=0
	Secret []Secret `hcl:"secret,block" validate:"min=0"`
	// SourceImageReference: optional
	SourceImageReference *SourceImageReference `hcl:"source_image_reference,block"`
	// SpotRestore: optional
	SpotRestore *SpotRestore `hcl:"spot_restore,block"`
	// TerminateNotification: optional
	TerminateNotification *TerminateNotification `hcl:"terminate_notification,block"`
	// TerminationNotification: optional
	TerminationNotification *TerminationNotification `hcl:"termination_notification,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// WinrmListener: min=0
	WinrmListener []WinrmListener `hcl:"winrm_listener,block" validate:"min=0"`
}

type azurermWindowsVirtualMachineScaleSetAttributes struct {
	ref terra.Reference
}

// AdminPassword returns a reference to field admin_password of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) AdminPassword() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("admin_password"))
}

// AdminUsername returns a reference to field admin_username of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("admin_username"))
}

// CapacityReservationGroupId returns a reference to field capacity_reservation_group_id of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) CapacityReservationGroupId() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("capacity_reservation_group_id"))
}

// ComputerNamePrefix returns a reference to field computer_name_prefix of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) ComputerNamePrefix() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("computer_name_prefix"))
}

// CustomData returns a reference to field custom_data of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) CustomData() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("custom_data"))
}

// DoNotRunExtensionsOnOverprovisionedMachines returns a reference to field do_not_run_extensions_on_overprovisioned_machines of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) DoNotRunExtensionsOnOverprovisionedMachines() terra.BoolValue {
	return terra.ReferenceAsBool(awvmss.ref.Append("do_not_run_extensions_on_overprovisioned_machines"))
}

// EdgeZone returns a reference to field edge_zone of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) EdgeZone() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("edge_zone"))
}

// EnableAutomaticUpdates returns a reference to field enable_automatic_updates of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) EnableAutomaticUpdates() terra.BoolValue {
	return terra.ReferenceAsBool(awvmss.ref.Append("enable_automatic_updates"))
}

// EncryptionAtHostEnabled returns a reference to field encryption_at_host_enabled of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) EncryptionAtHostEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awvmss.ref.Append("encryption_at_host_enabled"))
}

// EvictionPolicy returns a reference to field eviction_policy of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) EvictionPolicy() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("eviction_policy"))
}

// ExtensionOperationsEnabled returns a reference to field extension_operations_enabled of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) ExtensionOperationsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awvmss.ref.Append("extension_operations_enabled"))
}

// ExtensionsTimeBudget returns a reference to field extensions_time_budget of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) ExtensionsTimeBudget() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("extensions_time_budget"))
}

// HealthProbeId returns a reference to field health_probe_id of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) HealthProbeId() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("health_probe_id"))
}

// HostGroupId returns a reference to field host_group_id of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) HostGroupId() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("host_group_id"))
}

// Id returns a reference to field id of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("id"))
}

// Instances returns a reference to field instances of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Instances() terra.NumberValue {
	return terra.ReferenceAsNumber(awvmss.ref.Append("instances"))
}

// LicenseType returns a reference to field license_type of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("location"))
}

// MaxBidPrice returns a reference to field max_bid_price of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) MaxBidPrice() terra.NumberValue {
	return terra.ReferenceAsNumber(awvmss.ref.Append("max_bid_price"))
}

// Name returns a reference to field name of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("name"))
}

// Overprovision returns a reference to field overprovision of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Overprovision() terra.BoolValue {
	return terra.ReferenceAsBool(awvmss.ref.Append("overprovision"))
}

// PlatformFaultDomainCount returns a reference to field platform_fault_domain_count of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) PlatformFaultDomainCount() terra.NumberValue {
	return terra.ReferenceAsNumber(awvmss.ref.Append("platform_fault_domain_count"))
}

// Priority returns a reference to field priority of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Priority() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("priority"))
}

// ProvisionVmAgent returns a reference to field provision_vm_agent of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) ProvisionVmAgent() terra.BoolValue {
	return terra.ReferenceAsBool(awvmss.ref.Append("provision_vm_agent"))
}

// ProximityPlacementGroupId returns a reference to field proximity_placement_group_id of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) ProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("proximity_placement_group_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("resource_group_name"))
}

// ScaleInPolicy returns a reference to field scale_in_policy of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) ScaleInPolicy() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("scale_in_policy"))
}

// SecureBootEnabled returns a reference to field secure_boot_enabled of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) SecureBootEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awvmss.ref.Append("secure_boot_enabled"))
}

// SinglePlacementGroup returns a reference to field single_placement_group of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) SinglePlacementGroup() terra.BoolValue {
	return terra.ReferenceAsBool(awvmss.ref.Append("single_placement_group"))
}

// Sku returns a reference to field sku of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("sku"))
}

// SourceImageId returns a reference to field source_image_id of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) SourceImageId() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("source_image_id"))
}

// Tags returns a reference to field tags of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](awvmss.ref.Append("tags"))
}

// Timezone returns a reference to field timezone of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("timezone"))
}

// UniqueId returns a reference to field unique_id of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) UniqueId() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("unique_id"))
}

// UpgradeMode returns a reference to field upgrade_mode of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) UpgradeMode() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("upgrade_mode"))
}

// UserData returns a reference to field user_data of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) UserData() terra.StringValue {
	return terra.ReferenceAsString(awvmss.ref.Append("user_data"))
}

// VtpmEnabled returns a reference to field vtpm_enabled of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) VtpmEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awvmss.ref.Append("vtpm_enabled"))
}

// ZoneBalance returns a reference to field zone_balance of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) ZoneBalance() terra.BoolValue {
	return terra.ReferenceAsBool(awvmss.ref.Append("zone_balance"))
}

// Zones returns a reference to field zones of azurerm_windows_virtual_machine_scale_set.
func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](awvmss.ref.Append("zones"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) AdditionalCapabilities() terra.ListValue[AdditionalCapabilitiesAttributes] {
	return terra.ReferenceAsList[AdditionalCapabilitiesAttributes](awvmss.ref.Append("additional_capabilities"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) AdditionalUnattendContent() terra.ListValue[AdditionalUnattendContentAttributes] {
	return terra.ReferenceAsList[AdditionalUnattendContentAttributes](awvmss.ref.Append("additional_unattend_content"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) AutomaticInstanceRepair() terra.ListValue[AutomaticInstanceRepairAttributes] {
	return terra.ReferenceAsList[AutomaticInstanceRepairAttributes](awvmss.ref.Append("automatic_instance_repair"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) AutomaticOsUpgradePolicy() terra.ListValue[AutomaticOsUpgradePolicyAttributes] {
	return terra.ReferenceAsList[AutomaticOsUpgradePolicyAttributes](awvmss.ref.Append("automatic_os_upgrade_policy"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) BootDiagnostics() terra.ListValue[BootDiagnosticsAttributes] {
	return terra.ReferenceAsList[BootDiagnosticsAttributes](awvmss.ref.Append("boot_diagnostics"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) DataDisk() terra.ListValue[DataDiskAttributes] {
	return terra.ReferenceAsList[DataDiskAttributes](awvmss.ref.Append("data_disk"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Extension() terra.SetValue[ExtensionAttributes] {
	return terra.ReferenceAsSet[ExtensionAttributes](awvmss.ref.Append("extension"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) GalleryApplication() terra.ListValue[GalleryApplicationAttributes] {
	return terra.ReferenceAsList[GalleryApplicationAttributes](awvmss.ref.Append("gallery_application"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) GalleryApplications() terra.ListValue[GalleryApplicationsAttributes] {
	return terra.ReferenceAsList[GalleryApplicationsAttributes](awvmss.ref.Append("gallery_applications"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](awvmss.ref.Append("identity"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) NetworkInterface() terra.ListValue[NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[NetworkInterfaceAttributes](awvmss.ref.Append("network_interface"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) OsDisk() terra.ListValue[OsDiskAttributes] {
	return terra.ReferenceAsList[OsDiskAttributes](awvmss.ref.Append("os_disk"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Plan() terra.ListValue[PlanAttributes] {
	return terra.ReferenceAsList[PlanAttributes](awvmss.ref.Append("plan"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) RollingUpgradePolicy() terra.ListValue[RollingUpgradePolicyAttributes] {
	return terra.ReferenceAsList[RollingUpgradePolicyAttributes](awvmss.ref.Append("rolling_upgrade_policy"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) ScaleIn() terra.ListValue[ScaleInAttributes] {
	return terra.ReferenceAsList[ScaleInAttributes](awvmss.ref.Append("scale_in"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Secret() terra.ListValue[SecretAttributes] {
	return terra.ReferenceAsList[SecretAttributes](awvmss.ref.Append("secret"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) SourceImageReference() terra.ListValue[SourceImageReferenceAttributes] {
	return terra.ReferenceAsList[SourceImageReferenceAttributes](awvmss.ref.Append("source_image_reference"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) SpotRestore() terra.ListValue[SpotRestoreAttributes] {
	return terra.ReferenceAsList[SpotRestoreAttributes](awvmss.ref.Append("spot_restore"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) TerminateNotification() terra.ListValue[TerminateNotificationAttributes] {
	return terra.ReferenceAsList[TerminateNotificationAttributes](awvmss.ref.Append("terminate_notification"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) TerminationNotification() terra.ListValue[TerminationNotificationAttributes] {
	return terra.ReferenceAsList[TerminationNotificationAttributes](awvmss.ref.Append("termination_notification"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](awvmss.ref.Append("timeouts"))
}

func (awvmss azurermWindowsVirtualMachineScaleSetAttributes) WinrmListener() terra.SetValue[WinrmListenerAttributes] {
	return terra.ReferenceAsSet[WinrmListenerAttributes](awvmss.ref.Append("winrm_listener"))
}

type azurermWindowsVirtualMachineScaleSetState struct {
	AdminPassword                               string                           `json:"admin_password"`
	AdminUsername                               string                           `json:"admin_username"`
	CapacityReservationGroupId                  string                           `json:"capacity_reservation_group_id"`
	ComputerNamePrefix                          string                           `json:"computer_name_prefix"`
	CustomData                                  string                           `json:"custom_data"`
	DoNotRunExtensionsOnOverprovisionedMachines bool                             `json:"do_not_run_extensions_on_overprovisioned_machines"`
	EdgeZone                                    string                           `json:"edge_zone"`
	EnableAutomaticUpdates                      bool                             `json:"enable_automatic_updates"`
	EncryptionAtHostEnabled                     bool                             `json:"encryption_at_host_enabled"`
	EvictionPolicy                              string                           `json:"eviction_policy"`
	ExtensionOperationsEnabled                  bool                             `json:"extension_operations_enabled"`
	ExtensionsTimeBudget                        string                           `json:"extensions_time_budget"`
	HealthProbeId                               string                           `json:"health_probe_id"`
	HostGroupId                                 string                           `json:"host_group_id"`
	Id                                          string                           `json:"id"`
	Instances                                   float64                          `json:"instances"`
	LicenseType                                 string                           `json:"license_type"`
	Location                                    string                           `json:"location"`
	MaxBidPrice                                 float64                          `json:"max_bid_price"`
	Name                                        string                           `json:"name"`
	Overprovision                               bool                             `json:"overprovision"`
	PlatformFaultDomainCount                    float64                          `json:"platform_fault_domain_count"`
	Priority                                    string                           `json:"priority"`
	ProvisionVmAgent                            bool                             `json:"provision_vm_agent"`
	ProximityPlacementGroupId                   string                           `json:"proximity_placement_group_id"`
	ResourceGroupName                           string                           `json:"resource_group_name"`
	ScaleInPolicy                               string                           `json:"scale_in_policy"`
	SecureBootEnabled                           bool                             `json:"secure_boot_enabled"`
	SinglePlacementGroup                        bool                             `json:"single_placement_group"`
	Sku                                         string                           `json:"sku"`
	SourceImageId                               string                           `json:"source_image_id"`
	Tags                                        map[string]string                `json:"tags"`
	Timezone                                    string                           `json:"timezone"`
	UniqueId                                    string                           `json:"unique_id"`
	UpgradeMode                                 string                           `json:"upgrade_mode"`
	UserData                                    string                           `json:"user_data"`
	VtpmEnabled                                 bool                             `json:"vtpm_enabled"`
	ZoneBalance                                 bool                             `json:"zone_balance"`
	Zones                                       []string                         `json:"zones"`
	AdditionalCapabilities                      []AdditionalCapabilitiesState    `json:"additional_capabilities"`
	AdditionalUnattendContent                   []AdditionalUnattendContentState `json:"additional_unattend_content"`
	AutomaticInstanceRepair                     []AutomaticInstanceRepairState   `json:"automatic_instance_repair"`
	AutomaticOsUpgradePolicy                    []AutomaticOsUpgradePolicyState  `json:"automatic_os_upgrade_policy"`
	BootDiagnostics                             []BootDiagnosticsState           `json:"boot_diagnostics"`
	DataDisk                                    []DataDiskState                  `json:"data_disk"`
	Extension                                   []ExtensionState                 `json:"extension"`
	GalleryApplication                          []GalleryApplicationState        `json:"gallery_application"`
	GalleryApplications                         []GalleryApplicationsState       `json:"gallery_applications"`
	Identity                                    []IdentityState                  `json:"identity"`
	NetworkInterface                            []NetworkInterfaceState          `json:"network_interface"`
	OsDisk                                      []OsDiskState                    `json:"os_disk"`
	Plan                                        []PlanState                      `json:"plan"`
	RollingUpgradePolicy                        []RollingUpgradePolicyState      `json:"rolling_upgrade_policy"`
	ScaleIn                                     []ScaleInState                   `json:"scale_in"`
	Secret                                      []SecretState                    `json:"secret"`
	SourceImageReference                        []SourceImageReferenceState      `json:"source_image_reference"`
	SpotRestore                                 []SpotRestoreState               `json:"spot_restore"`
	TerminateNotification                       []TerminateNotificationState     `json:"terminate_notification"`
	TerminationNotification                     []TerminationNotificationState   `json:"termination_notification"`
	Timeouts                                    *TimeoutsState                   `json:"timeouts"`
	WinrmListener                               []WinrmListenerState             `json:"winrm_listener"`
}
