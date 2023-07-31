// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	backuppolicyvm "github.com/golingon/terraproviders/azurerm/3.67.0/backuppolicyvm"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBackupPolicyVm creates a new instance of [BackupPolicyVm].
func NewBackupPolicyVm(name string, args BackupPolicyVmArgs) *BackupPolicyVm {
	return &BackupPolicyVm{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupPolicyVm)(nil)

// BackupPolicyVm represents the Terraform resource azurerm_backup_policy_vm.
type BackupPolicyVm struct {
	Name      string
	Args      BackupPolicyVmArgs
	state     *backupPolicyVmState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BackupPolicyVm].
func (bpv *BackupPolicyVm) Type() string {
	return "azurerm_backup_policy_vm"
}

// LocalName returns the local name for [BackupPolicyVm].
func (bpv *BackupPolicyVm) LocalName() string {
	return bpv.Name
}

// Configuration returns the configuration (args) for [BackupPolicyVm].
func (bpv *BackupPolicyVm) Configuration() interface{} {
	return bpv.Args
}

// DependOn is used for other resources to depend on [BackupPolicyVm].
func (bpv *BackupPolicyVm) DependOn() terra.Reference {
	return terra.ReferenceResource(bpv)
}

// Dependencies returns the list of resources [BackupPolicyVm] depends_on.
func (bpv *BackupPolicyVm) Dependencies() terra.Dependencies {
	return bpv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BackupPolicyVm].
func (bpv *BackupPolicyVm) LifecycleManagement() *terra.Lifecycle {
	return bpv.Lifecycle
}

// Attributes returns the attributes for [BackupPolicyVm].
func (bpv *BackupPolicyVm) Attributes() backupPolicyVmAttributes {
	return backupPolicyVmAttributes{ref: terra.ReferenceResource(bpv)}
}

// ImportState imports the given attribute values into [BackupPolicyVm]'s state.
func (bpv *BackupPolicyVm) ImportState(av io.Reader) error {
	bpv.state = &backupPolicyVmState{}
	if err := json.NewDecoder(av).Decode(bpv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bpv.Type(), bpv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BackupPolicyVm] has state.
func (bpv *BackupPolicyVm) State() (*backupPolicyVmState, bool) {
	return bpv.state, bpv.state != nil
}

// StateMust returns the state for [BackupPolicyVm]. Panics if the state is nil.
func (bpv *BackupPolicyVm) StateMust() *backupPolicyVmState {
	if bpv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bpv.Type(), bpv.LocalName()))
	}
	return bpv.state
}

// BackupPolicyVmArgs contains the configurations for azurerm_backup_policy_vm.
type BackupPolicyVmArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstantRestoreRetentionDays: number, optional
	InstantRestoreRetentionDays terra.NumberValue `hcl:"instant_restore_retention_days,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyType: string, optional
	PolicyType terra.StringValue `hcl:"policy_type,attr"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timezone: string, optional
	Timezone terra.StringValue `hcl:"timezone,attr"`
	// Backup: required
	Backup *backuppolicyvm.Backup `hcl:"backup,block" validate:"required"`
	// InstantRestoreResourceGroup: optional
	InstantRestoreResourceGroup *backuppolicyvm.InstantRestoreResourceGroup `hcl:"instant_restore_resource_group,block"`
	// RetentionDaily: optional
	RetentionDaily *backuppolicyvm.RetentionDaily `hcl:"retention_daily,block"`
	// RetentionMonthly: optional
	RetentionMonthly *backuppolicyvm.RetentionMonthly `hcl:"retention_monthly,block"`
	// RetentionWeekly: optional
	RetentionWeekly *backuppolicyvm.RetentionWeekly `hcl:"retention_weekly,block"`
	// RetentionYearly: optional
	RetentionYearly *backuppolicyvm.RetentionYearly `hcl:"retention_yearly,block"`
	// Timeouts: optional
	Timeouts *backuppolicyvm.Timeouts `hcl:"timeouts,block"`
}
type backupPolicyVmAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_backup_policy_vm.
func (bpv backupPolicyVmAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("id"))
}

// InstantRestoreRetentionDays returns a reference to field instant_restore_retention_days of azurerm_backup_policy_vm.
func (bpv backupPolicyVmAttributes) InstantRestoreRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(bpv.ref.Append("instant_restore_retention_days"))
}

// Name returns a reference to field name of azurerm_backup_policy_vm.
func (bpv backupPolicyVmAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("name"))
}

// PolicyType returns a reference to field policy_type of azurerm_backup_policy_vm.
func (bpv backupPolicyVmAttributes) PolicyType() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("policy_type"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_backup_policy_vm.
func (bpv backupPolicyVmAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_backup_policy_vm.
func (bpv backupPolicyVmAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("resource_group_name"))
}

// Timezone returns a reference to field timezone of azurerm_backup_policy_vm.
func (bpv backupPolicyVmAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("timezone"))
}

func (bpv backupPolicyVmAttributes) Backup() terra.ListValue[backuppolicyvm.BackupAttributes] {
	return terra.ReferenceAsList[backuppolicyvm.BackupAttributes](bpv.ref.Append("backup"))
}

func (bpv backupPolicyVmAttributes) InstantRestoreResourceGroup() terra.ListValue[backuppolicyvm.InstantRestoreResourceGroupAttributes] {
	return terra.ReferenceAsList[backuppolicyvm.InstantRestoreResourceGroupAttributes](bpv.ref.Append("instant_restore_resource_group"))
}

func (bpv backupPolicyVmAttributes) RetentionDaily() terra.ListValue[backuppolicyvm.RetentionDailyAttributes] {
	return terra.ReferenceAsList[backuppolicyvm.RetentionDailyAttributes](bpv.ref.Append("retention_daily"))
}

func (bpv backupPolicyVmAttributes) RetentionMonthly() terra.ListValue[backuppolicyvm.RetentionMonthlyAttributes] {
	return terra.ReferenceAsList[backuppolicyvm.RetentionMonthlyAttributes](bpv.ref.Append("retention_monthly"))
}

func (bpv backupPolicyVmAttributes) RetentionWeekly() terra.ListValue[backuppolicyvm.RetentionWeeklyAttributes] {
	return terra.ReferenceAsList[backuppolicyvm.RetentionWeeklyAttributes](bpv.ref.Append("retention_weekly"))
}

func (bpv backupPolicyVmAttributes) RetentionYearly() terra.ListValue[backuppolicyvm.RetentionYearlyAttributes] {
	return terra.ReferenceAsList[backuppolicyvm.RetentionYearlyAttributes](bpv.ref.Append("retention_yearly"))
}

func (bpv backupPolicyVmAttributes) Timeouts() backuppolicyvm.TimeoutsAttributes {
	return terra.ReferenceAsSingle[backuppolicyvm.TimeoutsAttributes](bpv.ref.Append("timeouts"))
}

type backupPolicyVmState struct {
	Id                          string                                            `json:"id"`
	InstantRestoreRetentionDays float64                                           `json:"instant_restore_retention_days"`
	Name                        string                                            `json:"name"`
	PolicyType                  string                                            `json:"policy_type"`
	RecoveryVaultName           string                                            `json:"recovery_vault_name"`
	ResourceGroupName           string                                            `json:"resource_group_name"`
	Timezone                    string                                            `json:"timezone"`
	Backup                      []backuppolicyvm.BackupState                      `json:"backup"`
	InstantRestoreResourceGroup []backuppolicyvm.InstantRestoreResourceGroupState `json:"instant_restore_resource_group"`
	RetentionDaily              []backuppolicyvm.RetentionDailyState              `json:"retention_daily"`
	RetentionMonthly            []backuppolicyvm.RetentionMonthlyState            `json:"retention_monthly"`
	RetentionWeekly             []backuppolicyvm.RetentionWeeklyState             `json:"retention_weekly"`
	RetentionYearly             []backuppolicyvm.RetentionYearlyState             `json:"retention_yearly"`
	Timeouts                    *backuppolicyvm.TimeoutsState                     `json:"timeouts"`
}
