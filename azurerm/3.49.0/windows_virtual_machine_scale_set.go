// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	windowsvirtualmachinescaleset "github.com/golingon/terraproviders/azurerm/3.49.0/windowsvirtualmachinescaleset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewWindowsVirtualMachineScaleSet(name string, args WindowsVirtualMachineScaleSetArgs) *WindowsVirtualMachineScaleSet {
	return &WindowsVirtualMachineScaleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WindowsVirtualMachineScaleSet)(nil)

type WindowsVirtualMachineScaleSet struct {
	Name  string
	Args  WindowsVirtualMachineScaleSetArgs
	state *windowsVirtualMachineScaleSetState
}

func (wvmss *WindowsVirtualMachineScaleSet) Type() string {
	return "azurerm_windows_virtual_machine_scale_set"
}

func (wvmss *WindowsVirtualMachineScaleSet) LocalName() string {
	return wvmss.Name
}

func (wvmss *WindowsVirtualMachineScaleSet) Configuration() interface{} {
	return wvmss.Args
}

func (wvmss *WindowsVirtualMachineScaleSet) Attributes() windowsVirtualMachineScaleSetAttributes {
	return windowsVirtualMachineScaleSetAttributes{ref: terra.ReferenceResource(wvmss)}
}

func (wvmss *WindowsVirtualMachineScaleSet) ImportState(av io.Reader) error {
	wvmss.state = &windowsVirtualMachineScaleSetState{}
	if err := json.NewDecoder(av).Decode(wvmss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wvmss.Type(), wvmss.LocalName(), err)
	}
	return nil
}

func (wvmss *WindowsVirtualMachineScaleSet) State() (*windowsVirtualMachineScaleSetState, bool) {
	return wvmss.state, wvmss.state != nil
}

func (wvmss *WindowsVirtualMachineScaleSet) StateMust() *windowsVirtualMachineScaleSetState {
	if wvmss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wvmss.Type(), wvmss.LocalName()))
	}
	return wvmss.state
}

func (wvmss *WindowsVirtualMachineScaleSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wvmss)
}

type WindowsVirtualMachineScaleSetArgs struct {
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
	AdditionalCapabilities *windowsvirtualmachinescaleset.AdditionalCapabilities `hcl:"additional_capabilities,block"`
	// AdditionalUnattendContent: min=0
	AdditionalUnattendContent []windowsvirtualmachinescaleset.AdditionalUnattendContent `hcl:"additional_unattend_content,block" validate:"min=0"`
	// AutomaticInstanceRepair: optional
	AutomaticInstanceRepair *windowsvirtualmachinescaleset.AutomaticInstanceRepair `hcl:"automatic_instance_repair,block"`
	// AutomaticOsUpgradePolicy: optional
	AutomaticOsUpgradePolicy *windowsvirtualmachinescaleset.AutomaticOsUpgradePolicy `hcl:"automatic_os_upgrade_policy,block"`
	// BootDiagnostics: optional
	BootDiagnostics *windowsvirtualmachinescaleset.BootDiagnostics `hcl:"boot_diagnostics,block"`
	// DataDisk: min=0
	DataDisk []windowsvirtualmachinescaleset.DataDisk `hcl:"data_disk,block" validate:"min=0"`
	// Extension: min=0
	Extension []windowsvirtualmachinescaleset.Extension `hcl:"extension,block" validate:"min=0"`
	// GalleryApplication: min=0,max=100
	GalleryApplication []windowsvirtualmachinescaleset.GalleryApplication `hcl:"gallery_application,block" validate:"min=0,max=100"`
	// GalleryApplications: min=0,max=100
	GalleryApplications []windowsvirtualmachinescaleset.GalleryApplications `hcl:"gallery_applications,block" validate:"min=0,max=100"`
	// Identity: optional
	Identity *windowsvirtualmachinescaleset.Identity `hcl:"identity,block"`
	// NetworkInterface: min=1
	NetworkInterface []windowsvirtualmachinescaleset.NetworkInterface `hcl:"network_interface,block" validate:"min=1"`
	// OsDisk: required
	OsDisk *windowsvirtualmachinescaleset.OsDisk `hcl:"os_disk,block" validate:"required"`
	// Plan: optional
	Plan *windowsvirtualmachinescaleset.Plan `hcl:"plan,block"`
	// RollingUpgradePolicy: optional
	RollingUpgradePolicy *windowsvirtualmachinescaleset.RollingUpgradePolicy `hcl:"rolling_upgrade_policy,block"`
	// ScaleIn: optional
	ScaleIn *windowsvirtualmachinescaleset.ScaleIn `hcl:"scale_in,block"`
	// Secret: min=0
	Secret []windowsvirtualmachinescaleset.Secret `hcl:"secret,block" validate:"min=0"`
	// SourceImageReference: optional
	SourceImageReference *windowsvirtualmachinescaleset.SourceImageReference `hcl:"source_image_reference,block"`
	// SpotRestore: optional
	SpotRestore *windowsvirtualmachinescaleset.SpotRestore `hcl:"spot_restore,block"`
	// TerminateNotification: optional
	TerminateNotification *windowsvirtualmachinescaleset.TerminateNotification `hcl:"terminate_notification,block"`
	// TerminationNotification: optional
	TerminationNotification *windowsvirtualmachinescaleset.TerminationNotification `hcl:"termination_notification,block"`
	// Timeouts: optional
	Timeouts *windowsvirtualmachinescaleset.Timeouts `hcl:"timeouts,block"`
	// WinrmListener: min=0
	WinrmListener []windowsvirtualmachinescaleset.WinrmListener `hcl:"winrm_listener,block" validate:"min=0"`
	// DependsOn contains resources that WindowsVirtualMachineScaleSet depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type windowsVirtualMachineScaleSetAttributes struct {
	ref terra.Reference
}

func (wvmss windowsVirtualMachineScaleSetAttributes) AdminPassword() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("admin_password"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("admin_username"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) CapacityReservationGroupId() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("capacity_reservation_group_id"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) ComputerNamePrefix() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("computer_name_prefix"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) CustomData() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("custom_data"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) DoNotRunExtensionsOnOverprovisionedMachines() terra.BoolValue {
	return terra.ReferenceBool(wvmss.ref.Append("do_not_run_extensions_on_overprovisioned_machines"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) EdgeZone() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("edge_zone"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) EnableAutomaticUpdates() terra.BoolValue {
	return terra.ReferenceBool(wvmss.ref.Append("enable_automatic_updates"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) EncryptionAtHostEnabled() terra.BoolValue {
	return terra.ReferenceBool(wvmss.ref.Append("encryption_at_host_enabled"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) EvictionPolicy() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("eviction_policy"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) ExtensionOperationsEnabled() terra.BoolValue {
	return terra.ReferenceBool(wvmss.ref.Append("extension_operations_enabled"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) ExtensionsTimeBudget() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("extensions_time_budget"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) HealthProbeId() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("health_probe_id"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) HostGroupId() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("host_group_id"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("id"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Instances() terra.NumberValue {
	return terra.ReferenceNumber(wvmss.ref.Append("instances"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("license_type"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Location() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("location"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) MaxBidPrice() terra.NumberValue {
	return terra.ReferenceNumber(wvmss.ref.Append("max_bid_price"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Name() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("name"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Overprovision() terra.BoolValue {
	return terra.ReferenceBool(wvmss.ref.Append("overprovision"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) PlatformFaultDomainCount() terra.NumberValue {
	return terra.ReferenceNumber(wvmss.ref.Append("platform_fault_domain_count"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Priority() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("priority"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) ProvisionVmAgent() terra.BoolValue {
	return terra.ReferenceBool(wvmss.ref.Append("provision_vm_agent"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) ProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("proximity_placement_group_id"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("resource_group_name"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) ScaleInPolicy() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("scale_in_policy"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) SecureBootEnabled() terra.BoolValue {
	return terra.ReferenceBool(wvmss.ref.Append("secure_boot_enabled"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) SinglePlacementGroup() terra.BoolValue {
	return terra.ReferenceBool(wvmss.ref.Append("single_placement_group"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Sku() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("sku"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) SourceImageId() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("source_image_id"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](wvmss.ref.Append("tags"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Timezone() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("timezone"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) UniqueId() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("unique_id"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) UpgradeMode() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("upgrade_mode"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) UserData() terra.StringValue {
	return terra.ReferenceString(wvmss.ref.Append("user_data"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) VtpmEnabled() terra.BoolValue {
	return terra.ReferenceBool(wvmss.ref.Append("vtpm_enabled"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) ZoneBalance() terra.BoolValue {
	return terra.ReferenceBool(wvmss.ref.Append("zone_balance"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](wvmss.ref.Append("zones"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) AdditionalCapabilities() terra.ListValue[windowsvirtualmachinescaleset.AdditionalCapabilitiesAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.AdditionalCapabilitiesAttributes](wvmss.ref.Append("additional_capabilities"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) AdditionalUnattendContent() terra.ListValue[windowsvirtualmachinescaleset.AdditionalUnattendContentAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.AdditionalUnattendContentAttributes](wvmss.ref.Append("additional_unattend_content"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) AutomaticInstanceRepair() terra.ListValue[windowsvirtualmachinescaleset.AutomaticInstanceRepairAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.AutomaticInstanceRepairAttributes](wvmss.ref.Append("automatic_instance_repair"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) AutomaticOsUpgradePolicy() terra.ListValue[windowsvirtualmachinescaleset.AutomaticOsUpgradePolicyAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.AutomaticOsUpgradePolicyAttributes](wvmss.ref.Append("automatic_os_upgrade_policy"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) BootDiagnostics() terra.ListValue[windowsvirtualmachinescaleset.BootDiagnosticsAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.BootDiagnosticsAttributes](wvmss.ref.Append("boot_diagnostics"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) DataDisk() terra.ListValue[windowsvirtualmachinescaleset.DataDiskAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.DataDiskAttributes](wvmss.ref.Append("data_disk"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Extension() terra.SetValue[windowsvirtualmachinescaleset.ExtensionAttributes] {
	return terra.ReferenceSet[windowsvirtualmachinescaleset.ExtensionAttributes](wvmss.ref.Append("extension"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) GalleryApplication() terra.ListValue[windowsvirtualmachinescaleset.GalleryApplicationAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.GalleryApplicationAttributes](wvmss.ref.Append("gallery_application"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) GalleryApplications() terra.ListValue[windowsvirtualmachinescaleset.GalleryApplicationsAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.GalleryApplicationsAttributes](wvmss.ref.Append("gallery_applications"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Identity() terra.ListValue[windowsvirtualmachinescaleset.IdentityAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.IdentityAttributes](wvmss.ref.Append("identity"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) NetworkInterface() terra.ListValue[windowsvirtualmachinescaleset.NetworkInterfaceAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.NetworkInterfaceAttributes](wvmss.ref.Append("network_interface"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) OsDisk() terra.ListValue[windowsvirtualmachinescaleset.OsDiskAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.OsDiskAttributes](wvmss.ref.Append("os_disk"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Plan() terra.ListValue[windowsvirtualmachinescaleset.PlanAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.PlanAttributes](wvmss.ref.Append("plan"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) RollingUpgradePolicy() terra.ListValue[windowsvirtualmachinescaleset.RollingUpgradePolicyAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.RollingUpgradePolicyAttributes](wvmss.ref.Append("rolling_upgrade_policy"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) ScaleIn() terra.ListValue[windowsvirtualmachinescaleset.ScaleInAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.ScaleInAttributes](wvmss.ref.Append("scale_in"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Secret() terra.ListValue[windowsvirtualmachinescaleset.SecretAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.SecretAttributes](wvmss.ref.Append("secret"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) SourceImageReference() terra.ListValue[windowsvirtualmachinescaleset.SourceImageReferenceAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.SourceImageReferenceAttributes](wvmss.ref.Append("source_image_reference"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) SpotRestore() terra.ListValue[windowsvirtualmachinescaleset.SpotRestoreAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.SpotRestoreAttributes](wvmss.ref.Append("spot_restore"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) TerminateNotification() terra.ListValue[windowsvirtualmachinescaleset.TerminateNotificationAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.TerminateNotificationAttributes](wvmss.ref.Append("terminate_notification"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) TerminationNotification() terra.ListValue[windowsvirtualmachinescaleset.TerminationNotificationAttributes] {
	return terra.ReferenceList[windowsvirtualmachinescaleset.TerminationNotificationAttributes](wvmss.ref.Append("termination_notification"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) Timeouts() windowsvirtualmachinescaleset.TimeoutsAttributes {
	return terra.ReferenceSingle[windowsvirtualmachinescaleset.TimeoutsAttributes](wvmss.ref.Append("timeouts"))
}

func (wvmss windowsVirtualMachineScaleSetAttributes) WinrmListener() terra.SetValue[windowsvirtualmachinescaleset.WinrmListenerAttributes] {
	return terra.ReferenceSet[windowsvirtualmachinescaleset.WinrmListenerAttributes](wvmss.ref.Append("winrm_listener"))
}

type windowsVirtualMachineScaleSetState struct {
	AdminPassword                               string                                                         `json:"admin_password"`
	AdminUsername                               string                                                         `json:"admin_username"`
	CapacityReservationGroupId                  string                                                         `json:"capacity_reservation_group_id"`
	ComputerNamePrefix                          string                                                         `json:"computer_name_prefix"`
	CustomData                                  string                                                         `json:"custom_data"`
	DoNotRunExtensionsOnOverprovisionedMachines bool                                                           `json:"do_not_run_extensions_on_overprovisioned_machines"`
	EdgeZone                                    string                                                         `json:"edge_zone"`
	EnableAutomaticUpdates                      bool                                                           `json:"enable_automatic_updates"`
	EncryptionAtHostEnabled                     bool                                                           `json:"encryption_at_host_enabled"`
	EvictionPolicy                              string                                                         `json:"eviction_policy"`
	ExtensionOperationsEnabled                  bool                                                           `json:"extension_operations_enabled"`
	ExtensionsTimeBudget                        string                                                         `json:"extensions_time_budget"`
	HealthProbeId                               string                                                         `json:"health_probe_id"`
	HostGroupId                                 string                                                         `json:"host_group_id"`
	Id                                          string                                                         `json:"id"`
	Instances                                   float64                                                        `json:"instances"`
	LicenseType                                 string                                                         `json:"license_type"`
	Location                                    string                                                         `json:"location"`
	MaxBidPrice                                 float64                                                        `json:"max_bid_price"`
	Name                                        string                                                         `json:"name"`
	Overprovision                               bool                                                           `json:"overprovision"`
	PlatformFaultDomainCount                    float64                                                        `json:"platform_fault_domain_count"`
	Priority                                    string                                                         `json:"priority"`
	ProvisionVmAgent                            bool                                                           `json:"provision_vm_agent"`
	ProximityPlacementGroupId                   string                                                         `json:"proximity_placement_group_id"`
	ResourceGroupName                           string                                                         `json:"resource_group_name"`
	ScaleInPolicy                               string                                                         `json:"scale_in_policy"`
	SecureBootEnabled                           bool                                                           `json:"secure_boot_enabled"`
	SinglePlacementGroup                        bool                                                           `json:"single_placement_group"`
	Sku                                         string                                                         `json:"sku"`
	SourceImageId                               string                                                         `json:"source_image_id"`
	Tags                                        map[string]string                                              `json:"tags"`
	Timezone                                    string                                                         `json:"timezone"`
	UniqueId                                    string                                                         `json:"unique_id"`
	UpgradeMode                                 string                                                         `json:"upgrade_mode"`
	UserData                                    string                                                         `json:"user_data"`
	VtpmEnabled                                 bool                                                           `json:"vtpm_enabled"`
	ZoneBalance                                 bool                                                           `json:"zone_balance"`
	Zones                                       []string                                                       `json:"zones"`
	AdditionalCapabilities                      []windowsvirtualmachinescaleset.AdditionalCapabilitiesState    `json:"additional_capabilities"`
	AdditionalUnattendContent                   []windowsvirtualmachinescaleset.AdditionalUnattendContentState `json:"additional_unattend_content"`
	AutomaticInstanceRepair                     []windowsvirtualmachinescaleset.AutomaticInstanceRepairState   `json:"automatic_instance_repair"`
	AutomaticOsUpgradePolicy                    []windowsvirtualmachinescaleset.AutomaticOsUpgradePolicyState  `json:"automatic_os_upgrade_policy"`
	BootDiagnostics                             []windowsvirtualmachinescaleset.BootDiagnosticsState           `json:"boot_diagnostics"`
	DataDisk                                    []windowsvirtualmachinescaleset.DataDiskState                  `json:"data_disk"`
	Extension                                   []windowsvirtualmachinescaleset.ExtensionState                 `json:"extension"`
	GalleryApplication                          []windowsvirtualmachinescaleset.GalleryApplicationState        `json:"gallery_application"`
	GalleryApplications                         []windowsvirtualmachinescaleset.GalleryApplicationsState       `json:"gallery_applications"`
	Identity                                    []windowsvirtualmachinescaleset.IdentityState                  `json:"identity"`
	NetworkInterface                            []windowsvirtualmachinescaleset.NetworkInterfaceState          `json:"network_interface"`
	OsDisk                                      []windowsvirtualmachinescaleset.OsDiskState                    `json:"os_disk"`
	Plan                                        []windowsvirtualmachinescaleset.PlanState                      `json:"plan"`
	RollingUpgradePolicy                        []windowsvirtualmachinescaleset.RollingUpgradePolicyState      `json:"rolling_upgrade_policy"`
	ScaleIn                                     []windowsvirtualmachinescaleset.ScaleInState                   `json:"scale_in"`
	Secret                                      []windowsvirtualmachinescaleset.SecretState                    `json:"secret"`
	SourceImageReference                        []windowsvirtualmachinescaleset.SourceImageReferenceState      `json:"source_image_reference"`
	SpotRestore                                 []windowsvirtualmachinescaleset.SpotRestoreState               `json:"spot_restore"`
	TerminateNotification                       []windowsvirtualmachinescaleset.TerminateNotificationState     `json:"terminate_notification"`
	TerminationNotification                     []windowsvirtualmachinescaleset.TerminationNotificationState   `json:"termination_notification"`
	Timeouts                                    *windowsvirtualmachinescaleset.TimeoutsState                   `json:"timeouts"`
	WinrmListener                               []windowsvirtualmachinescaleset.WinrmListenerState             `json:"winrm_listener"`
}
