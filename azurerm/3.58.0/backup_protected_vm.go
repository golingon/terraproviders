// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	backupprotectedvm "github.com/golingon/terraproviders/azurerm/3.58.0/backupprotectedvm"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBackupProtectedVm creates a new instance of [BackupProtectedVm].
func NewBackupProtectedVm(name string, args BackupProtectedVmArgs) *BackupProtectedVm {
	return &BackupProtectedVm{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupProtectedVm)(nil)

// BackupProtectedVm represents the Terraform resource azurerm_backup_protected_vm.
type BackupProtectedVm struct {
	Name      string
	Args      BackupProtectedVmArgs
	state     *backupProtectedVmState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BackupProtectedVm].
func (bpv *BackupProtectedVm) Type() string {
	return "azurerm_backup_protected_vm"
}

// LocalName returns the local name for [BackupProtectedVm].
func (bpv *BackupProtectedVm) LocalName() string {
	return bpv.Name
}

// Configuration returns the configuration (args) for [BackupProtectedVm].
func (bpv *BackupProtectedVm) Configuration() interface{} {
	return bpv.Args
}

// DependOn is used for other resources to depend on [BackupProtectedVm].
func (bpv *BackupProtectedVm) DependOn() terra.Reference {
	return terra.ReferenceResource(bpv)
}

// Dependencies returns the list of resources [BackupProtectedVm] depends_on.
func (bpv *BackupProtectedVm) Dependencies() terra.Dependencies {
	return bpv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BackupProtectedVm].
func (bpv *BackupProtectedVm) LifecycleManagement() *terra.Lifecycle {
	return bpv.Lifecycle
}

// Attributes returns the attributes for [BackupProtectedVm].
func (bpv *BackupProtectedVm) Attributes() backupProtectedVmAttributes {
	return backupProtectedVmAttributes{ref: terra.ReferenceResource(bpv)}
}

// ImportState imports the given attribute values into [BackupProtectedVm]'s state.
func (bpv *BackupProtectedVm) ImportState(av io.Reader) error {
	bpv.state = &backupProtectedVmState{}
	if err := json.NewDecoder(av).Decode(bpv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bpv.Type(), bpv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BackupProtectedVm] has state.
func (bpv *BackupProtectedVm) State() (*backupProtectedVmState, bool) {
	return bpv.state, bpv.state != nil
}

// StateMust returns the state for [BackupProtectedVm]. Panics if the state is nil.
func (bpv *BackupProtectedVm) StateMust() *backupProtectedVmState {
	if bpv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bpv.Type(), bpv.LocalName()))
	}
	return bpv.state
}

// BackupProtectedVmArgs contains the configurations for azurerm_backup_protected_vm.
type BackupProtectedVmArgs struct {
	// BackupPolicyId: string, required
	BackupPolicyId terra.StringValue `hcl:"backup_policy_id,attr" validate:"required"`
	// ExcludeDiskLuns: set of number, optional
	ExcludeDiskLuns terra.SetValue[terra.NumberValue] `hcl:"exclude_disk_luns,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludeDiskLuns: set of number, optional
	IncludeDiskLuns terra.SetValue[terra.NumberValue] `hcl:"include_disk_luns,attr"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SourceVmId: string, optional
	SourceVmId terra.StringValue `hcl:"source_vm_id,attr"`
	// Timeouts: optional
	Timeouts *backupprotectedvm.Timeouts `hcl:"timeouts,block"`
}
type backupProtectedVmAttributes struct {
	ref terra.Reference
}

// BackupPolicyId returns a reference to field backup_policy_id of azurerm_backup_protected_vm.
func (bpv backupProtectedVmAttributes) BackupPolicyId() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("backup_policy_id"))
}

// ExcludeDiskLuns returns a reference to field exclude_disk_luns of azurerm_backup_protected_vm.
func (bpv backupProtectedVmAttributes) ExcludeDiskLuns() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](bpv.ref.Append("exclude_disk_luns"))
}

// Id returns a reference to field id of azurerm_backup_protected_vm.
func (bpv backupProtectedVmAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("id"))
}

// IncludeDiskLuns returns a reference to field include_disk_luns of azurerm_backup_protected_vm.
func (bpv backupProtectedVmAttributes) IncludeDiskLuns() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](bpv.ref.Append("include_disk_luns"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_backup_protected_vm.
func (bpv backupProtectedVmAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_backup_protected_vm.
func (bpv backupProtectedVmAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("resource_group_name"))
}

// SourceVmId returns a reference to field source_vm_id of azurerm_backup_protected_vm.
func (bpv backupProtectedVmAttributes) SourceVmId() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("source_vm_id"))
}

func (bpv backupProtectedVmAttributes) Timeouts() backupprotectedvm.TimeoutsAttributes {
	return terra.ReferenceAsSingle[backupprotectedvm.TimeoutsAttributes](bpv.ref.Append("timeouts"))
}

type backupProtectedVmState struct {
	BackupPolicyId    string                           `json:"backup_policy_id"`
	ExcludeDiskLuns   []float64                        `json:"exclude_disk_luns"`
	Id                string                           `json:"id"`
	IncludeDiskLuns   []float64                        `json:"include_disk_luns"`
	RecoveryVaultName string                           `json:"recovery_vault_name"`
	ResourceGroupName string                           `json:"resource_group_name"`
	SourceVmId        string                           `json:"source_vm_id"`
	Timeouts          *backupprotectedvm.TimeoutsState `json:"timeouts"`
}
