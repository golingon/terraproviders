// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_site_recovery_vmware_replicated_vm

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

// Resource represents the Terraform resource azurerm_site_recovery_vmware_replicated_vm.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSiteRecoveryVmwareReplicatedVmState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asrvrv *Resource) Type() string {
	return "azurerm_site_recovery_vmware_replicated_vm"
}

// LocalName returns the local name for [Resource].
func (asrvrv *Resource) LocalName() string {
	return asrvrv.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asrvrv *Resource) Configuration() interface{} {
	return asrvrv.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asrvrv *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asrvrv)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asrvrv *Resource) Dependencies() terra.Dependencies {
	return asrvrv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asrvrv *Resource) LifecycleManagement() *terra.Lifecycle {
	return asrvrv.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asrvrv *Resource) Attributes() azurermSiteRecoveryVmwareReplicatedVmAttributes {
	return azurermSiteRecoveryVmwareReplicatedVmAttributes{ref: terra.ReferenceResource(asrvrv)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asrvrv *Resource) ImportState(state io.Reader) error {
	asrvrv.state = &azurermSiteRecoveryVmwareReplicatedVmState{}
	if err := json.NewDecoder(state).Decode(asrvrv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asrvrv.Type(), asrvrv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asrvrv *Resource) State() (*azurermSiteRecoveryVmwareReplicatedVmState, bool) {
	return asrvrv.state, asrvrv.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asrvrv *Resource) StateMust() *azurermSiteRecoveryVmwareReplicatedVmState {
	if asrvrv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asrvrv.Type(), asrvrv.LocalName()))
	}
	return asrvrv.state
}

// Args contains the configurations for azurerm_site_recovery_vmware_replicated_vm.
type Args struct {
	// ApplianceName: string, required
	ApplianceName terra.StringValue `hcl:"appliance_name,attr" validate:"required"`
	// DefaultLogStorageAccountId: string, optional
	DefaultLogStorageAccountId terra.StringValue `hcl:"default_log_storage_account_id,attr"`
	// DefaultRecoveryDiskType: string, optional
	DefaultRecoveryDiskType terra.StringValue `hcl:"default_recovery_disk_type,attr"`
	// DefaultTargetDiskEncryptionSetId: string, optional
	DefaultTargetDiskEncryptionSetId terra.StringValue `hcl:"default_target_disk_encryption_set_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LicenseType: string, optional
	LicenseType terra.StringValue `hcl:"license_type,attr"`
	// MultiVmGroupName: string, optional
	MultiVmGroupName terra.StringValue `hcl:"multi_vm_group_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PhysicalServerCredentialName: string, required
	PhysicalServerCredentialName terra.StringValue `hcl:"physical_server_credential_name,attr" validate:"required"`
	// RecoveryReplicationPolicyId: string, required
	RecoveryReplicationPolicyId terra.StringValue `hcl:"recovery_replication_policy_id,attr" validate:"required"`
	// RecoveryVaultId: string, required
	RecoveryVaultId terra.StringValue `hcl:"recovery_vault_id,attr" validate:"required"`
	// SourceVmName: string, required
	SourceVmName terra.StringValue `hcl:"source_vm_name,attr" validate:"required"`
	// TargetAvailabilitySetId: string, optional
	TargetAvailabilitySetId terra.StringValue `hcl:"target_availability_set_id,attr"`
	// TargetBootDiagnosticsStorageAccountId: string, optional
	TargetBootDiagnosticsStorageAccountId terra.StringValue `hcl:"target_boot_diagnostics_storage_account_id,attr"`
	// TargetNetworkId: string, optional
	TargetNetworkId terra.StringValue `hcl:"target_network_id,attr"`
	// TargetProximityPlacementGroupId: string, optional
	TargetProximityPlacementGroupId terra.StringValue `hcl:"target_proximity_placement_group_id,attr"`
	// TargetResourceGroupId: string, required
	TargetResourceGroupId terra.StringValue `hcl:"target_resource_group_id,attr" validate:"required"`
	// TargetVmName: string, required
	TargetVmName terra.StringValue `hcl:"target_vm_name,attr" validate:"required"`
	// TargetVmSize: string, optional
	TargetVmSize terra.StringValue `hcl:"target_vm_size,attr"`
	// TargetZone: string, optional
	TargetZone terra.StringValue `hcl:"target_zone,attr"`
	// TestNetworkId: string, optional
	TestNetworkId terra.StringValue `hcl:"test_network_id,attr"`
	// ManagedDisk: min=0
	ManagedDisk []ManagedDisk `hcl:"managed_disk,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSiteRecoveryVmwareReplicatedVmAttributes struct {
	ref terra.Reference
}

// ApplianceName returns a reference to field appliance_name of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) ApplianceName() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("appliance_name"))
}

// DefaultLogStorageAccountId returns a reference to field default_log_storage_account_id of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) DefaultLogStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("default_log_storage_account_id"))
}

// DefaultRecoveryDiskType returns a reference to field default_recovery_disk_type of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) DefaultRecoveryDiskType() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("default_recovery_disk_type"))
}

// DefaultTargetDiskEncryptionSetId returns a reference to field default_target_disk_encryption_set_id of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) DefaultTargetDiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("default_target_disk_encryption_set_id"))
}

// Id returns a reference to field id of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("license_type"))
}

// MultiVmGroupName returns a reference to field multi_vm_group_name of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) MultiVmGroupName() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("multi_vm_group_name"))
}

// Name returns a reference to field name of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("name"))
}

// PhysicalServerCredentialName returns a reference to field physical_server_credential_name of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) PhysicalServerCredentialName() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("physical_server_credential_name"))
}

// RecoveryReplicationPolicyId returns a reference to field recovery_replication_policy_id of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) RecoveryReplicationPolicyId() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("recovery_replication_policy_id"))
}

// RecoveryVaultId returns a reference to field recovery_vault_id of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) RecoveryVaultId() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("recovery_vault_id"))
}

// SourceVmName returns a reference to field source_vm_name of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) SourceVmName() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("source_vm_name"))
}

// TargetAvailabilitySetId returns a reference to field target_availability_set_id of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) TargetAvailabilitySetId() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("target_availability_set_id"))
}

// TargetBootDiagnosticsStorageAccountId returns a reference to field target_boot_diagnostics_storage_account_id of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) TargetBootDiagnosticsStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("target_boot_diagnostics_storage_account_id"))
}

// TargetNetworkId returns a reference to field target_network_id of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) TargetNetworkId() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("target_network_id"))
}

// TargetProximityPlacementGroupId returns a reference to field target_proximity_placement_group_id of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) TargetProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("target_proximity_placement_group_id"))
}

// TargetResourceGroupId returns a reference to field target_resource_group_id of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) TargetResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("target_resource_group_id"))
}

// TargetVmName returns a reference to field target_vm_name of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) TargetVmName() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("target_vm_name"))
}

// TargetVmSize returns a reference to field target_vm_size of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) TargetVmSize() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("target_vm_size"))
}

// TargetZone returns a reference to field target_zone of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) TargetZone() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("target_zone"))
}

// TestNetworkId returns a reference to field test_network_id of azurerm_site_recovery_vmware_replicated_vm.
func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) TestNetworkId() terra.StringValue {
	return terra.ReferenceAsString(asrvrv.ref.Append("test_network_id"))
}

func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) ManagedDisk() terra.ListValue[ManagedDiskAttributes] {
	return terra.ReferenceAsList[ManagedDiskAttributes](asrvrv.ref.Append("managed_disk"))
}

func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) NetworkInterface() terra.ListValue[NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[NetworkInterfaceAttributes](asrvrv.ref.Append("network_interface"))
}

func (asrvrv azurermSiteRecoveryVmwareReplicatedVmAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asrvrv.ref.Append("timeouts"))
}

type azurermSiteRecoveryVmwareReplicatedVmState struct {
	ApplianceName                         string                  `json:"appliance_name"`
	DefaultLogStorageAccountId            string                  `json:"default_log_storage_account_id"`
	DefaultRecoveryDiskType               string                  `json:"default_recovery_disk_type"`
	DefaultTargetDiskEncryptionSetId      string                  `json:"default_target_disk_encryption_set_id"`
	Id                                    string                  `json:"id"`
	LicenseType                           string                  `json:"license_type"`
	MultiVmGroupName                      string                  `json:"multi_vm_group_name"`
	Name                                  string                  `json:"name"`
	PhysicalServerCredentialName          string                  `json:"physical_server_credential_name"`
	RecoveryReplicationPolicyId           string                  `json:"recovery_replication_policy_id"`
	RecoveryVaultId                       string                  `json:"recovery_vault_id"`
	SourceVmName                          string                  `json:"source_vm_name"`
	TargetAvailabilitySetId               string                  `json:"target_availability_set_id"`
	TargetBootDiagnosticsStorageAccountId string                  `json:"target_boot_diagnostics_storage_account_id"`
	TargetNetworkId                       string                  `json:"target_network_id"`
	TargetProximityPlacementGroupId       string                  `json:"target_proximity_placement_group_id"`
	TargetResourceGroupId                 string                  `json:"target_resource_group_id"`
	TargetVmName                          string                  `json:"target_vm_name"`
	TargetVmSize                          string                  `json:"target_vm_size"`
	TargetZone                            string                  `json:"target_zone"`
	TestNetworkId                         string                  `json:"test_network_id"`
	ManagedDisk                           []ManagedDiskState      `json:"managed_disk"`
	NetworkInterface                      []NetworkInterfaceState `json:"network_interface"`
	Timeouts                              *TimeoutsState          `json:"timeouts"`
}
