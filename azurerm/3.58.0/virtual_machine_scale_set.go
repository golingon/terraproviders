// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualmachinescaleset "github.com/golingon/terraproviders/azurerm/3.58.0/virtualmachinescaleset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualMachineScaleSet creates a new instance of [VirtualMachineScaleSet].
func NewVirtualMachineScaleSet(name string, args VirtualMachineScaleSetArgs) *VirtualMachineScaleSet {
	return &VirtualMachineScaleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualMachineScaleSet)(nil)

// VirtualMachineScaleSet represents the Terraform resource azurerm_virtual_machine_scale_set.
type VirtualMachineScaleSet struct {
	Name      string
	Args      VirtualMachineScaleSetArgs
	state     *virtualMachineScaleSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualMachineScaleSet].
func (vmss *VirtualMachineScaleSet) Type() string {
	return "azurerm_virtual_machine_scale_set"
}

// LocalName returns the local name for [VirtualMachineScaleSet].
func (vmss *VirtualMachineScaleSet) LocalName() string {
	return vmss.Name
}

// Configuration returns the configuration (args) for [VirtualMachineScaleSet].
func (vmss *VirtualMachineScaleSet) Configuration() interface{} {
	return vmss.Args
}

// DependOn is used for other resources to depend on [VirtualMachineScaleSet].
func (vmss *VirtualMachineScaleSet) DependOn() terra.Reference {
	return terra.ReferenceResource(vmss)
}

// Dependencies returns the list of resources [VirtualMachineScaleSet] depends_on.
func (vmss *VirtualMachineScaleSet) Dependencies() terra.Dependencies {
	return vmss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualMachineScaleSet].
func (vmss *VirtualMachineScaleSet) LifecycleManagement() *terra.Lifecycle {
	return vmss.Lifecycle
}

// Attributes returns the attributes for [VirtualMachineScaleSet].
func (vmss *VirtualMachineScaleSet) Attributes() virtualMachineScaleSetAttributes {
	return virtualMachineScaleSetAttributes{ref: terra.ReferenceResource(vmss)}
}

// ImportState imports the given attribute values into [VirtualMachineScaleSet]'s state.
func (vmss *VirtualMachineScaleSet) ImportState(av io.Reader) error {
	vmss.state = &virtualMachineScaleSetState{}
	if err := json.NewDecoder(av).Decode(vmss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vmss.Type(), vmss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualMachineScaleSet] has state.
func (vmss *VirtualMachineScaleSet) State() (*virtualMachineScaleSetState, bool) {
	return vmss.state, vmss.state != nil
}

// StateMust returns the state for [VirtualMachineScaleSet]. Panics if the state is nil.
func (vmss *VirtualMachineScaleSet) StateMust() *virtualMachineScaleSetState {
	if vmss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vmss.Type(), vmss.LocalName()))
	}
	return vmss.state
}

// VirtualMachineScaleSetArgs contains the configurations for azurerm_virtual_machine_scale_set.
type VirtualMachineScaleSetArgs struct {
	// AutomaticOsUpgrade: bool, optional
	AutomaticOsUpgrade terra.BoolValue `hcl:"automatic_os_upgrade,attr"`
	// EvictionPolicy: string, optional
	EvictionPolicy terra.StringValue `hcl:"eviction_policy,attr"`
	// HealthProbeId: string, optional
	HealthProbeId terra.StringValue `hcl:"health_probe_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LicenseType: string, optional
	LicenseType terra.StringValue `hcl:"license_type,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Overprovision: bool, optional
	Overprovision terra.BoolValue `hcl:"overprovision,attr"`
	// Priority: string, optional
	Priority terra.StringValue `hcl:"priority,attr"`
	// ProximityPlacementGroupId: string, optional
	ProximityPlacementGroupId terra.StringValue `hcl:"proximity_placement_group_id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SinglePlacementGroup: bool, optional
	SinglePlacementGroup terra.BoolValue `hcl:"single_placement_group,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// UpgradePolicyMode: string, required
	UpgradePolicyMode terra.StringValue `hcl:"upgrade_policy_mode,attr" validate:"required"`
	// Zones: list of string, optional
	Zones terra.ListValue[terra.StringValue] `hcl:"zones,attr"`
	// BootDiagnostics: optional
	BootDiagnostics *virtualmachinescaleset.BootDiagnostics `hcl:"boot_diagnostics,block"`
	// Extension: min=0
	Extension []virtualmachinescaleset.Extension `hcl:"extension,block" validate:"min=0"`
	// Identity: optional
	Identity *virtualmachinescaleset.Identity `hcl:"identity,block"`
	// NetworkProfile: min=1
	NetworkProfile []virtualmachinescaleset.NetworkProfile `hcl:"network_profile,block" validate:"min=1"`
	// OsProfile: required
	OsProfile *virtualmachinescaleset.OsProfile `hcl:"os_profile,block" validate:"required"`
	// OsProfileLinuxConfig: optional
	OsProfileLinuxConfig *virtualmachinescaleset.OsProfileLinuxConfig `hcl:"os_profile_linux_config,block"`
	// OsProfileSecrets: min=0
	OsProfileSecrets []virtualmachinescaleset.OsProfileSecrets `hcl:"os_profile_secrets,block" validate:"min=0"`
	// OsProfileWindowsConfig: optional
	OsProfileWindowsConfig *virtualmachinescaleset.OsProfileWindowsConfig `hcl:"os_profile_windows_config,block"`
	// Plan: optional
	Plan *virtualmachinescaleset.Plan `hcl:"plan,block"`
	// RollingUpgradePolicy: optional
	RollingUpgradePolicy *virtualmachinescaleset.RollingUpgradePolicy `hcl:"rolling_upgrade_policy,block"`
	// Sku: required
	Sku *virtualmachinescaleset.Sku `hcl:"sku,block" validate:"required"`
	// StorageProfileDataDisk: min=0
	StorageProfileDataDisk []virtualmachinescaleset.StorageProfileDataDisk `hcl:"storage_profile_data_disk,block" validate:"min=0"`
	// StorageProfileImageReference: optional
	StorageProfileImageReference *virtualmachinescaleset.StorageProfileImageReference `hcl:"storage_profile_image_reference,block"`
	// StorageProfileOsDisk: required
	StorageProfileOsDisk *virtualmachinescaleset.StorageProfileOsDisk `hcl:"storage_profile_os_disk,block" validate:"required"`
	// Timeouts: optional
	Timeouts *virtualmachinescaleset.Timeouts `hcl:"timeouts,block"`
}
type virtualMachineScaleSetAttributes struct {
	ref terra.Reference
}

// AutomaticOsUpgrade returns a reference to field automatic_os_upgrade of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) AutomaticOsUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(vmss.ref.Append("automatic_os_upgrade"))
}

// EvictionPolicy returns a reference to field eviction_policy of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) EvictionPolicy() terra.StringValue {
	return terra.ReferenceAsString(vmss.ref.Append("eviction_policy"))
}

// HealthProbeId returns a reference to field health_probe_id of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) HealthProbeId() terra.StringValue {
	return terra.ReferenceAsString(vmss.ref.Append("health_probe_id"))
}

// Id returns a reference to field id of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vmss.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(vmss.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vmss.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vmss.ref.Append("name"))
}

// Overprovision returns a reference to field overprovision of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) Overprovision() terra.BoolValue {
	return terra.ReferenceAsBool(vmss.ref.Append("overprovision"))
}

// Priority returns a reference to field priority of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) Priority() terra.StringValue {
	return terra.ReferenceAsString(vmss.ref.Append("priority"))
}

// ProximityPlacementGroupId returns a reference to field proximity_placement_group_id of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) ProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceAsString(vmss.ref.Append("proximity_placement_group_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vmss.ref.Append("resource_group_name"))
}

// SinglePlacementGroup returns a reference to field single_placement_group of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) SinglePlacementGroup() terra.BoolValue {
	return terra.ReferenceAsBool(vmss.ref.Append("single_placement_group"))
}

// Tags returns a reference to field tags of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vmss.ref.Append("tags"))
}

// UpgradePolicyMode returns a reference to field upgrade_policy_mode of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) UpgradePolicyMode() terra.StringValue {
	return terra.ReferenceAsString(vmss.ref.Append("upgrade_policy_mode"))
}

// Zones returns a reference to field zones of azurerm_virtual_machine_scale_set.
func (vmss virtualMachineScaleSetAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vmss.ref.Append("zones"))
}

func (vmss virtualMachineScaleSetAttributes) BootDiagnostics() terra.ListValue[virtualmachinescaleset.BootDiagnosticsAttributes] {
	return terra.ReferenceAsList[virtualmachinescaleset.BootDiagnosticsAttributes](vmss.ref.Append("boot_diagnostics"))
}

func (vmss virtualMachineScaleSetAttributes) Extension() terra.SetValue[virtualmachinescaleset.ExtensionAttributes] {
	return terra.ReferenceAsSet[virtualmachinescaleset.ExtensionAttributes](vmss.ref.Append("extension"))
}

func (vmss virtualMachineScaleSetAttributes) Identity() terra.ListValue[virtualmachinescaleset.IdentityAttributes] {
	return terra.ReferenceAsList[virtualmachinescaleset.IdentityAttributes](vmss.ref.Append("identity"))
}

func (vmss virtualMachineScaleSetAttributes) NetworkProfile() terra.SetValue[virtualmachinescaleset.NetworkProfileAttributes] {
	return terra.ReferenceAsSet[virtualmachinescaleset.NetworkProfileAttributes](vmss.ref.Append("network_profile"))
}

func (vmss virtualMachineScaleSetAttributes) OsProfile() terra.ListValue[virtualmachinescaleset.OsProfileAttributes] {
	return terra.ReferenceAsList[virtualmachinescaleset.OsProfileAttributes](vmss.ref.Append("os_profile"))
}

func (vmss virtualMachineScaleSetAttributes) OsProfileLinuxConfig() terra.SetValue[virtualmachinescaleset.OsProfileLinuxConfigAttributes] {
	return terra.ReferenceAsSet[virtualmachinescaleset.OsProfileLinuxConfigAttributes](vmss.ref.Append("os_profile_linux_config"))
}

func (vmss virtualMachineScaleSetAttributes) OsProfileSecrets() terra.SetValue[virtualmachinescaleset.OsProfileSecretsAttributes] {
	return terra.ReferenceAsSet[virtualmachinescaleset.OsProfileSecretsAttributes](vmss.ref.Append("os_profile_secrets"))
}

func (vmss virtualMachineScaleSetAttributes) OsProfileWindowsConfig() terra.SetValue[virtualmachinescaleset.OsProfileWindowsConfigAttributes] {
	return terra.ReferenceAsSet[virtualmachinescaleset.OsProfileWindowsConfigAttributes](vmss.ref.Append("os_profile_windows_config"))
}

func (vmss virtualMachineScaleSetAttributes) Plan() terra.SetValue[virtualmachinescaleset.PlanAttributes] {
	return terra.ReferenceAsSet[virtualmachinescaleset.PlanAttributes](vmss.ref.Append("plan"))
}

func (vmss virtualMachineScaleSetAttributes) RollingUpgradePolicy() terra.ListValue[virtualmachinescaleset.RollingUpgradePolicyAttributes] {
	return terra.ReferenceAsList[virtualmachinescaleset.RollingUpgradePolicyAttributes](vmss.ref.Append("rolling_upgrade_policy"))
}

func (vmss virtualMachineScaleSetAttributes) Sku() terra.ListValue[virtualmachinescaleset.SkuAttributes] {
	return terra.ReferenceAsList[virtualmachinescaleset.SkuAttributes](vmss.ref.Append("sku"))
}

func (vmss virtualMachineScaleSetAttributes) StorageProfileDataDisk() terra.ListValue[virtualmachinescaleset.StorageProfileDataDiskAttributes] {
	return terra.ReferenceAsList[virtualmachinescaleset.StorageProfileDataDiskAttributes](vmss.ref.Append("storage_profile_data_disk"))
}

func (vmss virtualMachineScaleSetAttributes) StorageProfileImageReference() terra.SetValue[virtualmachinescaleset.StorageProfileImageReferenceAttributes] {
	return terra.ReferenceAsSet[virtualmachinescaleset.StorageProfileImageReferenceAttributes](vmss.ref.Append("storage_profile_image_reference"))
}

func (vmss virtualMachineScaleSetAttributes) StorageProfileOsDisk() terra.SetValue[virtualmachinescaleset.StorageProfileOsDiskAttributes] {
	return terra.ReferenceAsSet[virtualmachinescaleset.StorageProfileOsDiskAttributes](vmss.ref.Append("storage_profile_os_disk"))
}

func (vmss virtualMachineScaleSetAttributes) Timeouts() virtualmachinescaleset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualmachinescaleset.TimeoutsAttributes](vmss.ref.Append("timeouts"))
}

type virtualMachineScaleSetState struct {
	AutomaticOsUpgrade           bool                                                       `json:"automatic_os_upgrade"`
	EvictionPolicy               string                                                     `json:"eviction_policy"`
	HealthProbeId                string                                                     `json:"health_probe_id"`
	Id                           string                                                     `json:"id"`
	LicenseType                  string                                                     `json:"license_type"`
	Location                     string                                                     `json:"location"`
	Name                         string                                                     `json:"name"`
	Overprovision                bool                                                       `json:"overprovision"`
	Priority                     string                                                     `json:"priority"`
	ProximityPlacementGroupId    string                                                     `json:"proximity_placement_group_id"`
	ResourceGroupName            string                                                     `json:"resource_group_name"`
	SinglePlacementGroup         bool                                                       `json:"single_placement_group"`
	Tags                         map[string]string                                          `json:"tags"`
	UpgradePolicyMode            string                                                     `json:"upgrade_policy_mode"`
	Zones                        []string                                                   `json:"zones"`
	BootDiagnostics              []virtualmachinescaleset.BootDiagnosticsState              `json:"boot_diagnostics"`
	Extension                    []virtualmachinescaleset.ExtensionState                    `json:"extension"`
	Identity                     []virtualmachinescaleset.IdentityState                     `json:"identity"`
	NetworkProfile               []virtualmachinescaleset.NetworkProfileState               `json:"network_profile"`
	OsProfile                    []virtualmachinescaleset.OsProfileState                    `json:"os_profile"`
	OsProfileLinuxConfig         []virtualmachinescaleset.OsProfileLinuxConfigState         `json:"os_profile_linux_config"`
	OsProfileSecrets             []virtualmachinescaleset.OsProfileSecretsState             `json:"os_profile_secrets"`
	OsProfileWindowsConfig       []virtualmachinescaleset.OsProfileWindowsConfigState       `json:"os_profile_windows_config"`
	Plan                         []virtualmachinescaleset.PlanState                         `json:"plan"`
	RollingUpgradePolicy         []virtualmachinescaleset.RollingUpgradePolicyState         `json:"rolling_upgrade_policy"`
	Sku                          []virtualmachinescaleset.SkuState                          `json:"sku"`
	StorageProfileDataDisk       []virtualmachinescaleset.StorageProfileDataDiskState       `json:"storage_profile_data_disk"`
	StorageProfileImageReference []virtualmachinescaleset.StorageProfileImageReferenceState `json:"storage_profile_image_reference"`
	StorageProfileOsDisk         []virtualmachinescaleset.StorageProfileOsDiskState         `json:"storage_profile_os_disk"`
	Timeouts                     *virtualmachinescaleset.TimeoutsState                      `json:"timeouts"`
}