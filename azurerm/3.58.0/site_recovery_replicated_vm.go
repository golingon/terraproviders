// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryreplicatedvm "github.com/golingon/terraproviders/azurerm/3.58.0/siterecoveryreplicatedvm"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSiteRecoveryReplicatedVm creates a new instance of [SiteRecoveryReplicatedVm].
func NewSiteRecoveryReplicatedVm(name string, args SiteRecoveryReplicatedVmArgs) *SiteRecoveryReplicatedVm {
	return &SiteRecoveryReplicatedVm{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryReplicatedVm)(nil)

// SiteRecoveryReplicatedVm represents the Terraform resource azurerm_site_recovery_replicated_vm.
type SiteRecoveryReplicatedVm struct {
	Name      string
	Args      SiteRecoveryReplicatedVmArgs
	state     *siteRecoveryReplicatedVmState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SiteRecoveryReplicatedVm].
func (srrv *SiteRecoveryReplicatedVm) Type() string {
	return "azurerm_site_recovery_replicated_vm"
}

// LocalName returns the local name for [SiteRecoveryReplicatedVm].
func (srrv *SiteRecoveryReplicatedVm) LocalName() string {
	return srrv.Name
}

// Configuration returns the configuration (args) for [SiteRecoveryReplicatedVm].
func (srrv *SiteRecoveryReplicatedVm) Configuration() interface{} {
	return srrv.Args
}

// DependOn is used for other resources to depend on [SiteRecoveryReplicatedVm].
func (srrv *SiteRecoveryReplicatedVm) DependOn() terra.Reference {
	return terra.ReferenceResource(srrv)
}

// Dependencies returns the list of resources [SiteRecoveryReplicatedVm] depends_on.
func (srrv *SiteRecoveryReplicatedVm) Dependencies() terra.Dependencies {
	return srrv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SiteRecoveryReplicatedVm].
func (srrv *SiteRecoveryReplicatedVm) LifecycleManagement() *terra.Lifecycle {
	return srrv.Lifecycle
}

// Attributes returns the attributes for [SiteRecoveryReplicatedVm].
func (srrv *SiteRecoveryReplicatedVm) Attributes() siteRecoveryReplicatedVmAttributes {
	return siteRecoveryReplicatedVmAttributes{ref: terra.ReferenceResource(srrv)}
}

// ImportState imports the given attribute values into [SiteRecoveryReplicatedVm]'s state.
func (srrv *SiteRecoveryReplicatedVm) ImportState(av io.Reader) error {
	srrv.state = &siteRecoveryReplicatedVmState{}
	if err := json.NewDecoder(av).Decode(srrv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srrv.Type(), srrv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SiteRecoveryReplicatedVm] has state.
func (srrv *SiteRecoveryReplicatedVm) State() (*siteRecoveryReplicatedVmState, bool) {
	return srrv.state, srrv.state != nil
}

// StateMust returns the state for [SiteRecoveryReplicatedVm]. Panics if the state is nil.
func (srrv *SiteRecoveryReplicatedVm) StateMust() *siteRecoveryReplicatedVmState {
	if srrv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srrv.Type(), srrv.LocalName()))
	}
	return srrv.state
}

// SiteRecoveryReplicatedVmArgs contains the configurations for azurerm_site_recovery_replicated_vm.
type SiteRecoveryReplicatedVmArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MultiVmGroupName: string, optional
	MultiVmGroupName terra.StringValue `hcl:"multi_vm_group_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryReplicationPolicyId: string, required
	RecoveryReplicationPolicyId terra.StringValue `hcl:"recovery_replication_policy_id,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SourceRecoveryFabricName: string, required
	SourceRecoveryFabricName terra.StringValue `hcl:"source_recovery_fabric_name,attr" validate:"required"`
	// SourceRecoveryProtectionContainerName: string, required
	SourceRecoveryProtectionContainerName terra.StringValue `hcl:"source_recovery_protection_container_name,attr" validate:"required"`
	// SourceVmId: string, required
	SourceVmId terra.StringValue `hcl:"source_vm_id,attr" validate:"required"`
	// TargetAvailabilitySetId: string, optional
	TargetAvailabilitySetId terra.StringValue `hcl:"target_availability_set_id,attr"`
	// TargetBootDiagnosticStorageAccountId: string, optional
	TargetBootDiagnosticStorageAccountId terra.StringValue `hcl:"target_boot_diagnostic_storage_account_id,attr"`
	// TargetCapacityReservationGroupId: string, optional
	TargetCapacityReservationGroupId terra.StringValue `hcl:"target_capacity_reservation_group_id,attr"`
	// TargetEdgeZone: string, optional
	TargetEdgeZone terra.StringValue `hcl:"target_edge_zone,attr"`
	// TargetNetworkId: string, optional
	TargetNetworkId terra.StringValue `hcl:"target_network_id,attr"`
	// TargetProximityPlacementGroupId: string, optional
	TargetProximityPlacementGroupId terra.StringValue `hcl:"target_proximity_placement_group_id,attr"`
	// TargetRecoveryFabricId: string, required
	TargetRecoveryFabricId terra.StringValue `hcl:"target_recovery_fabric_id,attr" validate:"required"`
	// TargetRecoveryProtectionContainerId: string, required
	TargetRecoveryProtectionContainerId terra.StringValue `hcl:"target_recovery_protection_container_id,attr" validate:"required"`
	// TargetResourceGroupId: string, required
	TargetResourceGroupId terra.StringValue `hcl:"target_resource_group_id,attr" validate:"required"`
	// TargetVirtualMachineScaleSetId: string, optional
	TargetVirtualMachineScaleSetId terra.StringValue `hcl:"target_virtual_machine_scale_set_id,attr"`
	// TargetZone: string, optional
	TargetZone terra.StringValue `hcl:"target_zone,attr"`
	// TestNetworkId: string, optional
	TestNetworkId terra.StringValue `hcl:"test_network_id,attr"`
	// ManagedDisk: min=0
	ManagedDisk []siterecoveryreplicatedvm.ManagedDisk `hcl:"managed_disk,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []siterecoveryreplicatedvm.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// UnmanagedDisk: min=0
	UnmanagedDisk []siterecoveryreplicatedvm.UnmanagedDisk `hcl:"unmanaged_disk,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *siterecoveryreplicatedvm.Timeouts `hcl:"timeouts,block"`
}
type siteRecoveryReplicatedVmAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("id"))
}

// MultiVmGroupName returns a reference to field multi_vm_group_name of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) MultiVmGroupName() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("multi_vm_group_name"))
}

// Name returns a reference to field name of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("name"))
}

// RecoveryReplicationPolicyId returns a reference to field recovery_replication_policy_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) RecoveryReplicationPolicyId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("recovery_replication_policy_id"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("resource_group_name"))
}

// SourceRecoveryFabricName returns a reference to field source_recovery_fabric_name of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) SourceRecoveryFabricName() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("source_recovery_fabric_name"))
}

// SourceRecoveryProtectionContainerName returns a reference to field source_recovery_protection_container_name of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) SourceRecoveryProtectionContainerName() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("source_recovery_protection_container_name"))
}

// SourceVmId returns a reference to field source_vm_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) SourceVmId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("source_vm_id"))
}

// TargetAvailabilitySetId returns a reference to field target_availability_set_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TargetAvailabilitySetId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("target_availability_set_id"))
}

// TargetBootDiagnosticStorageAccountId returns a reference to field target_boot_diagnostic_storage_account_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TargetBootDiagnosticStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("target_boot_diagnostic_storage_account_id"))
}

// TargetCapacityReservationGroupId returns a reference to field target_capacity_reservation_group_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TargetCapacityReservationGroupId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("target_capacity_reservation_group_id"))
}

// TargetEdgeZone returns a reference to field target_edge_zone of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TargetEdgeZone() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("target_edge_zone"))
}

// TargetNetworkId returns a reference to field target_network_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TargetNetworkId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("target_network_id"))
}

// TargetProximityPlacementGroupId returns a reference to field target_proximity_placement_group_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TargetProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("target_proximity_placement_group_id"))
}

// TargetRecoveryFabricId returns a reference to field target_recovery_fabric_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TargetRecoveryFabricId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("target_recovery_fabric_id"))
}

// TargetRecoveryProtectionContainerId returns a reference to field target_recovery_protection_container_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TargetRecoveryProtectionContainerId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("target_recovery_protection_container_id"))
}

// TargetResourceGroupId returns a reference to field target_resource_group_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TargetResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("target_resource_group_id"))
}

// TargetVirtualMachineScaleSetId returns a reference to field target_virtual_machine_scale_set_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TargetVirtualMachineScaleSetId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("target_virtual_machine_scale_set_id"))
}

// TargetZone returns a reference to field target_zone of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TargetZone() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("target_zone"))
}

// TestNetworkId returns a reference to field test_network_id of azurerm_site_recovery_replicated_vm.
func (srrv siteRecoveryReplicatedVmAttributes) TestNetworkId() terra.StringValue {
	return terra.ReferenceAsString(srrv.ref.Append("test_network_id"))
}

func (srrv siteRecoveryReplicatedVmAttributes) ManagedDisk() terra.SetValue[siterecoveryreplicatedvm.ManagedDiskAttributes] {
	return terra.ReferenceAsSet[siterecoveryreplicatedvm.ManagedDiskAttributes](srrv.ref.Append("managed_disk"))
}

func (srrv siteRecoveryReplicatedVmAttributes) NetworkInterface() terra.SetValue[siterecoveryreplicatedvm.NetworkInterfaceAttributes] {
	return terra.ReferenceAsSet[siterecoveryreplicatedvm.NetworkInterfaceAttributes](srrv.ref.Append("network_interface"))
}

func (srrv siteRecoveryReplicatedVmAttributes) UnmanagedDisk() terra.SetValue[siterecoveryreplicatedvm.UnmanagedDiskAttributes] {
	return terra.ReferenceAsSet[siterecoveryreplicatedvm.UnmanagedDiskAttributes](srrv.ref.Append("unmanaged_disk"))
}

func (srrv siteRecoveryReplicatedVmAttributes) Timeouts() siterecoveryreplicatedvm.TimeoutsAttributes {
	return terra.ReferenceAsSingle[siterecoveryreplicatedvm.TimeoutsAttributes](srrv.ref.Append("timeouts"))
}

type siteRecoveryReplicatedVmState struct {
	Id                                    string                                           `json:"id"`
	MultiVmGroupName                      string                                           `json:"multi_vm_group_name"`
	Name                                  string                                           `json:"name"`
	RecoveryReplicationPolicyId           string                                           `json:"recovery_replication_policy_id"`
	RecoveryVaultName                     string                                           `json:"recovery_vault_name"`
	ResourceGroupName                     string                                           `json:"resource_group_name"`
	SourceRecoveryFabricName              string                                           `json:"source_recovery_fabric_name"`
	SourceRecoveryProtectionContainerName string                                           `json:"source_recovery_protection_container_name"`
	SourceVmId                            string                                           `json:"source_vm_id"`
	TargetAvailabilitySetId               string                                           `json:"target_availability_set_id"`
	TargetBootDiagnosticStorageAccountId  string                                           `json:"target_boot_diagnostic_storage_account_id"`
	TargetCapacityReservationGroupId      string                                           `json:"target_capacity_reservation_group_id"`
	TargetEdgeZone                        string                                           `json:"target_edge_zone"`
	TargetNetworkId                       string                                           `json:"target_network_id"`
	TargetProximityPlacementGroupId       string                                           `json:"target_proximity_placement_group_id"`
	TargetRecoveryFabricId                string                                           `json:"target_recovery_fabric_id"`
	TargetRecoveryProtectionContainerId   string                                           `json:"target_recovery_protection_container_id"`
	TargetResourceGroupId                 string                                           `json:"target_resource_group_id"`
	TargetVirtualMachineScaleSetId        string                                           `json:"target_virtual_machine_scale_set_id"`
	TargetZone                            string                                           `json:"target_zone"`
	TestNetworkId                         string                                           `json:"test_network_id"`
	ManagedDisk                           []siterecoveryreplicatedvm.ManagedDiskState      `json:"managed_disk"`
	NetworkInterface                      []siterecoveryreplicatedvm.NetworkInterfaceState `json:"network_interface"`
	UnmanagedDisk                         []siterecoveryreplicatedvm.UnmanagedDiskState    `json:"unmanaged_disk"`
	Timeouts                              *siterecoveryreplicatedvm.TimeoutsState          `json:"timeouts"`
}
