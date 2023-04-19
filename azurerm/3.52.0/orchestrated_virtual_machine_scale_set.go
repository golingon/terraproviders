// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	orchestratedvirtualmachinescaleset "github.com/golingon/terraproviders/azurerm/3.52.0/orchestratedvirtualmachinescaleset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrchestratedVirtualMachineScaleSet creates a new instance of [OrchestratedVirtualMachineScaleSet].
func NewOrchestratedVirtualMachineScaleSet(name string, args OrchestratedVirtualMachineScaleSetArgs) *OrchestratedVirtualMachineScaleSet {
	return &OrchestratedVirtualMachineScaleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrchestratedVirtualMachineScaleSet)(nil)

// OrchestratedVirtualMachineScaleSet represents the Terraform resource azurerm_orchestrated_virtual_machine_scale_set.
type OrchestratedVirtualMachineScaleSet struct {
	Name      string
	Args      OrchestratedVirtualMachineScaleSetArgs
	state     *orchestratedVirtualMachineScaleSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrchestratedVirtualMachineScaleSet].
func (ovmss *OrchestratedVirtualMachineScaleSet) Type() string {
	return "azurerm_orchestrated_virtual_machine_scale_set"
}

// LocalName returns the local name for [OrchestratedVirtualMachineScaleSet].
func (ovmss *OrchestratedVirtualMachineScaleSet) LocalName() string {
	return ovmss.Name
}

// Configuration returns the configuration (args) for [OrchestratedVirtualMachineScaleSet].
func (ovmss *OrchestratedVirtualMachineScaleSet) Configuration() interface{} {
	return ovmss.Args
}

// DependOn is used for other resources to depend on [OrchestratedVirtualMachineScaleSet].
func (ovmss *OrchestratedVirtualMachineScaleSet) DependOn() terra.Reference {
	return terra.ReferenceResource(ovmss)
}

// Dependencies returns the list of resources [OrchestratedVirtualMachineScaleSet] depends_on.
func (ovmss *OrchestratedVirtualMachineScaleSet) Dependencies() terra.Dependencies {
	return ovmss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrchestratedVirtualMachineScaleSet].
func (ovmss *OrchestratedVirtualMachineScaleSet) LifecycleManagement() *terra.Lifecycle {
	return ovmss.Lifecycle
}

// Attributes returns the attributes for [OrchestratedVirtualMachineScaleSet].
func (ovmss *OrchestratedVirtualMachineScaleSet) Attributes() orchestratedVirtualMachineScaleSetAttributes {
	return orchestratedVirtualMachineScaleSetAttributes{ref: terra.ReferenceResource(ovmss)}
}

// ImportState imports the given attribute values into [OrchestratedVirtualMachineScaleSet]'s state.
func (ovmss *OrchestratedVirtualMachineScaleSet) ImportState(av io.Reader) error {
	ovmss.state = &orchestratedVirtualMachineScaleSetState{}
	if err := json.NewDecoder(av).Decode(ovmss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ovmss.Type(), ovmss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrchestratedVirtualMachineScaleSet] has state.
func (ovmss *OrchestratedVirtualMachineScaleSet) State() (*orchestratedVirtualMachineScaleSetState, bool) {
	return ovmss.state, ovmss.state != nil
}

// StateMust returns the state for [OrchestratedVirtualMachineScaleSet]. Panics if the state is nil.
func (ovmss *OrchestratedVirtualMachineScaleSet) StateMust() *orchestratedVirtualMachineScaleSetState {
	if ovmss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ovmss.Type(), ovmss.LocalName()))
	}
	return ovmss.state
}

// OrchestratedVirtualMachineScaleSetArgs contains the configurations for azurerm_orchestrated_virtual_machine_scale_set.
type OrchestratedVirtualMachineScaleSetArgs struct {
	// CapacityReservationGroupId: string, optional
	CapacityReservationGroupId terra.StringValue `hcl:"capacity_reservation_group_id,attr"`
	// EncryptionAtHostEnabled: bool, optional
	EncryptionAtHostEnabled terra.BoolValue `hcl:"encryption_at_host_enabled,attr"`
	// EvictionPolicy: string, optional
	EvictionPolicy terra.StringValue `hcl:"eviction_policy,attr"`
	// ExtensionOperationsEnabled: bool, optional
	ExtensionOperationsEnabled terra.BoolValue `hcl:"extension_operations_enabled,attr"`
	// ExtensionsTimeBudget: string, optional
	ExtensionsTimeBudget terra.StringValue `hcl:"extensions_time_budget,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instances: number, optional
	Instances terra.NumberValue `hcl:"instances,attr"`
	// LicenseType: string, optional
	LicenseType terra.StringValue `hcl:"license_type,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaxBidPrice: number, optional
	MaxBidPrice terra.NumberValue `hcl:"max_bid_price,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PlatformFaultDomainCount: number, required
	PlatformFaultDomainCount terra.NumberValue `hcl:"platform_fault_domain_count,attr" validate:"required"`
	// Priority: string, optional
	Priority terra.StringValue `hcl:"priority,attr"`
	// ProximityPlacementGroupId: string, optional
	ProximityPlacementGroupId terra.StringValue `hcl:"proximity_placement_group_id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SinglePlacementGroup: bool, optional
	SinglePlacementGroup terra.BoolValue `hcl:"single_placement_group,attr"`
	// SkuName: string, optional
	SkuName terra.StringValue `hcl:"sku_name,attr"`
	// SourceImageId: string, optional
	SourceImageId terra.StringValue `hcl:"source_image_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// UserDataBase64: string, optional
	UserDataBase64 terra.StringValue `hcl:"user_data_base64,attr"`
	// ZoneBalance: bool, optional
	ZoneBalance terra.BoolValue `hcl:"zone_balance,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// AdditionalCapabilities: optional
	AdditionalCapabilities *orchestratedvirtualmachinescaleset.AdditionalCapabilities `hcl:"additional_capabilities,block"`
	// AutomaticInstanceRepair: optional
	AutomaticInstanceRepair *orchestratedvirtualmachinescaleset.AutomaticInstanceRepair `hcl:"automatic_instance_repair,block"`
	// BootDiagnostics: optional
	BootDiagnostics *orchestratedvirtualmachinescaleset.BootDiagnostics `hcl:"boot_diagnostics,block"`
	// DataDisk: min=0
	DataDisk []orchestratedvirtualmachinescaleset.DataDisk `hcl:"data_disk,block" validate:"min=0"`
	// Extension: min=0
	Extension []orchestratedvirtualmachinescaleset.Extension `hcl:"extension,block" validate:"min=0"`
	// Identity: optional
	Identity *orchestratedvirtualmachinescaleset.Identity `hcl:"identity,block"`
	// NetworkInterface: min=0
	NetworkInterface []orchestratedvirtualmachinescaleset.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// OsDisk: optional
	OsDisk *orchestratedvirtualmachinescaleset.OsDisk `hcl:"os_disk,block"`
	// OsProfile: optional
	OsProfile *orchestratedvirtualmachinescaleset.OsProfile `hcl:"os_profile,block"`
	// Plan: optional
	Plan *orchestratedvirtualmachinescaleset.Plan `hcl:"plan,block"`
	// PriorityMix: optional
	PriorityMix *orchestratedvirtualmachinescaleset.PriorityMix `hcl:"priority_mix,block"`
	// SourceImageReference: optional
	SourceImageReference *orchestratedvirtualmachinescaleset.SourceImageReference `hcl:"source_image_reference,block"`
	// TerminationNotification: optional
	TerminationNotification *orchestratedvirtualmachinescaleset.TerminationNotification `hcl:"termination_notification,block"`
	// Timeouts: optional
	Timeouts *orchestratedvirtualmachinescaleset.Timeouts `hcl:"timeouts,block"`
}
type orchestratedVirtualMachineScaleSetAttributes struct {
	ref terra.Reference
}

// CapacityReservationGroupId returns a reference to field capacity_reservation_group_id of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) CapacityReservationGroupId() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("capacity_reservation_group_id"))
}

// EncryptionAtHostEnabled returns a reference to field encryption_at_host_enabled of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) EncryptionAtHostEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ovmss.ref.Append("encryption_at_host_enabled"))
}

// EvictionPolicy returns a reference to field eviction_policy of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) EvictionPolicy() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("eviction_policy"))
}

// ExtensionOperationsEnabled returns a reference to field extension_operations_enabled of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) ExtensionOperationsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ovmss.ref.Append("extension_operations_enabled"))
}

// ExtensionsTimeBudget returns a reference to field extensions_time_budget of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) ExtensionsTimeBudget() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("extensions_time_budget"))
}

// Id returns a reference to field id of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("id"))
}

// Instances returns a reference to field instances of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) Instances() terra.NumberValue {
	return terra.ReferenceAsNumber(ovmss.ref.Append("instances"))
}

// LicenseType returns a reference to field license_type of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("location"))
}

// MaxBidPrice returns a reference to field max_bid_price of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) MaxBidPrice() terra.NumberValue {
	return terra.ReferenceAsNumber(ovmss.ref.Append("max_bid_price"))
}

// Name returns a reference to field name of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("name"))
}

// PlatformFaultDomainCount returns a reference to field platform_fault_domain_count of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) PlatformFaultDomainCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ovmss.ref.Append("platform_fault_domain_count"))
}

// Priority returns a reference to field priority of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) Priority() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("priority"))
}

// ProximityPlacementGroupId returns a reference to field proximity_placement_group_id of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) ProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("proximity_placement_group_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("resource_group_name"))
}

// SinglePlacementGroup returns a reference to field single_placement_group of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) SinglePlacementGroup() terra.BoolValue {
	return terra.ReferenceAsBool(ovmss.ref.Append("single_placement_group"))
}

// SkuName returns a reference to field sku_name of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("sku_name"))
}

// SourceImageId returns a reference to field source_image_id of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) SourceImageId() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("source_image_id"))
}

// Tags returns a reference to field tags of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ovmss.ref.Append("tags"))
}

// UniqueId returns a reference to field unique_id of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) UniqueId() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("unique_id"))
}

// UserDataBase64 returns a reference to field user_data_base64 of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) UserDataBase64() terra.StringValue {
	return terra.ReferenceAsString(ovmss.ref.Append("user_data_base64"))
}

// ZoneBalance returns a reference to field zone_balance of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) ZoneBalance() terra.BoolValue {
	return terra.ReferenceAsBool(ovmss.ref.Append("zone_balance"))
}

// Zones returns a reference to field zones of azurerm_orchestrated_virtual_machine_scale_set.
func (ovmss orchestratedVirtualMachineScaleSetAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ovmss.ref.Append("zones"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) AdditionalCapabilities() terra.ListValue[orchestratedvirtualmachinescaleset.AdditionalCapabilitiesAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.AdditionalCapabilitiesAttributes](ovmss.ref.Append("additional_capabilities"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) AutomaticInstanceRepair() terra.ListValue[orchestratedvirtualmachinescaleset.AutomaticInstanceRepairAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.AutomaticInstanceRepairAttributes](ovmss.ref.Append("automatic_instance_repair"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) BootDiagnostics() terra.ListValue[orchestratedvirtualmachinescaleset.BootDiagnosticsAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.BootDiagnosticsAttributes](ovmss.ref.Append("boot_diagnostics"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) DataDisk() terra.ListValue[orchestratedvirtualmachinescaleset.DataDiskAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.DataDiskAttributes](ovmss.ref.Append("data_disk"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) Extension() terra.SetValue[orchestratedvirtualmachinescaleset.ExtensionAttributes] {
	return terra.ReferenceAsSet[orchestratedvirtualmachinescaleset.ExtensionAttributes](ovmss.ref.Append("extension"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) Identity() terra.ListValue[orchestratedvirtualmachinescaleset.IdentityAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.IdentityAttributes](ovmss.ref.Append("identity"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) NetworkInterface() terra.ListValue[orchestratedvirtualmachinescaleset.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.NetworkInterfaceAttributes](ovmss.ref.Append("network_interface"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) OsDisk() terra.ListValue[orchestratedvirtualmachinescaleset.OsDiskAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.OsDiskAttributes](ovmss.ref.Append("os_disk"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) OsProfile() terra.ListValue[orchestratedvirtualmachinescaleset.OsProfileAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.OsProfileAttributes](ovmss.ref.Append("os_profile"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) Plan() terra.ListValue[orchestratedvirtualmachinescaleset.PlanAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.PlanAttributes](ovmss.ref.Append("plan"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) PriorityMix() terra.ListValue[orchestratedvirtualmachinescaleset.PriorityMixAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.PriorityMixAttributes](ovmss.ref.Append("priority_mix"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) SourceImageReference() terra.ListValue[orchestratedvirtualmachinescaleset.SourceImageReferenceAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.SourceImageReferenceAttributes](ovmss.ref.Append("source_image_reference"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) TerminationNotification() terra.ListValue[orchestratedvirtualmachinescaleset.TerminationNotificationAttributes] {
	return terra.ReferenceAsList[orchestratedvirtualmachinescaleset.TerminationNotificationAttributes](ovmss.ref.Append("termination_notification"))
}

func (ovmss orchestratedVirtualMachineScaleSetAttributes) Timeouts() orchestratedvirtualmachinescaleset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[orchestratedvirtualmachinescaleset.TimeoutsAttributes](ovmss.ref.Append("timeouts"))
}

type orchestratedVirtualMachineScaleSetState struct {
	CapacityReservationGroupId string                                                            `json:"capacity_reservation_group_id"`
	EncryptionAtHostEnabled    bool                                                              `json:"encryption_at_host_enabled"`
	EvictionPolicy             string                                                            `json:"eviction_policy"`
	ExtensionOperationsEnabled bool                                                              `json:"extension_operations_enabled"`
	ExtensionsTimeBudget       string                                                            `json:"extensions_time_budget"`
	Id                         string                                                            `json:"id"`
	Instances                  float64                                                           `json:"instances"`
	LicenseType                string                                                            `json:"license_type"`
	Location                   string                                                            `json:"location"`
	MaxBidPrice                float64                                                           `json:"max_bid_price"`
	Name                       string                                                            `json:"name"`
	PlatformFaultDomainCount   float64                                                           `json:"platform_fault_domain_count"`
	Priority                   string                                                            `json:"priority"`
	ProximityPlacementGroupId  string                                                            `json:"proximity_placement_group_id"`
	ResourceGroupName          string                                                            `json:"resource_group_name"`
	SinglePlacementGroup       bool                                                              `json:"single_placement_group"`
	SkuName                    string                                                            `json:"sku_name"`
	SourceImageId              string                                                            `json:"source_image_id"`
	Tags                       map[string]string                                                 `json:"tags"`
	UniqueId                   string                                                            `json:"unique_id"`
	UserDataBase64             string                                                            `json:"user_data_base64"`
	ZoneBalance                bool                                                              `json:"zone_balance"`
	Zones                      []string                                                          `json:"zones"`
	AdditionalCapabilities     []orchestratedvirtualmachinescaleset.AdditionalCapabilitiesState  `json:"additional_capabilities"`
	AutomaticInstanceRepair    []orchestratedvirtualmachinescaleset.AutomaticInstanceRepairState `json:"automatic_instance_repair"`
	BootDiagnostics            []orchestratedvirtualmachinescaleset.BootDiagnosticsState         `json:"boot_diagnostics"`
	DataDisk                   []orchestratedvirtualmachinescaleset.DataDiskState                `json:"data_disk"`
	Extension                  []orchestratedvirtualmachinescaleset.ExtensionState               `json:"extension"`
	Identity                   []orchestratedvirtualmachinescaleset.IdentityState                `json:"identity"`
	NetworkInterface           []orchestratedvirtualmachinescaleset.NetworkInterfaceState        `json:"network_interface"`
	OsDisk                     []orchestratedvirtualmachinescaleset.OsDiskState                  `json:"os_disk"`
	OsProfile                  []orchestratedvirtualmachinescaleset.OsProfileState               `json:"os_profile"`
	Plan                       []orchestratedvirtualmachinescaleset.PlanState                    `json:"plan"`
	PriorityMix                []orchestratedvirtualmachinescaleset.PriorityMixState             `json:"priority_mix"`
	SourceImageReference       []orchestratedvirtualmachinescaleset.SourceImageReferenceState    `json:"source_image_reference"`
	TerminationNotification    []orchestratedvirtualmachinescaleset.TerminationNotificationState `json:"termination_notification"`
	Timeouts                   *orchestratedvirtualmachinescaleset.TimeoutsState                 `json:"timeouts"`
}
