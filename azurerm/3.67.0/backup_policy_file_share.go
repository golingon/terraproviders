// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	backuppolicyfileshare "github.com/golingon/terraproviders/azurerm/3.67.0/backuppolicyfileshare"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBackupPolicyFileShare creates a new instance of [BackupPolicyFileShare].
func NewBackupPolicyFileShare(name string, args BackupPolicyFileShareArgs) *BackupPolicyFileShare {
	return &BackupPolicyFileShare{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupPolicyFileShare)(nil)

// BackupPolicyFileShare represents the Terraform resource azurerm_backup_policy_file_share.
type BackupPolicyFileShare struct {
	Name      string
	Args      BackupPolicyFileShareArgs
	state     *backupPolicyFileShareState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BackupPolicyFileShare].
func (bpfs *BackupPolicyFileShare) Type() string {
	return "azurerm_backup_policy_file_share"
}

// LocalName returns the local name for [BackupPolicyFileShare].
func (bpfs *BackupPolicyFileShare) LocalName() string {
	return bpfs.Name
}

// Configuration returns the configuration (args) for [BackupPolicyFileShare].
func (bpfs *BackupPolicyFileShare) Configuration() interface{} {
	return bpfs.Args
}

// DependOn is used for other resources to depend on [BackupPolicyFileShare].
func (bpfs *BackupPolicyFileShare) DependOn() terra.Reference {
	return terra.ReferenceResource(bpfs)
}

// Dependencies returns the list of resources [BackupPolicyFileShare] depends_on.
func (bpfs *BackupPolicyFileShare) Dependencies() terra.Dependencies {
	return bpfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BackupPolicyFileShare].
func (bpfs *BackupPolicyFileShare) LifecycleManagement() *terra.Lifecycle {
	return bpfs.Lifecycle
}

// Attributes returns the attributes for [BackupPolicyFileShare].
func (bpfs *BackupPolicyFileShare) Attributes() backupPolicyFileShareAttributes {
	return backupPolicyFileShareAttributes{ref: terra.ReferenceResource(bpfs)}
}

// ImportState imports the given attribute values into [BackupPolicyFileShare]'s state.
func (bpfs *BackupPolicyFileShare) ImportState(av io.Reader) error {
	bpfs.state = &backupPolicyFileShareState{}
	if err := json.NewDecoder(av).Decode(bpfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bpfs.Type(), bpfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BackupPolicyFileShare] has state.
func (bpfs *BackupPolicyFileShare) State() (*backupPolicyFileShareState, bool) {
	return bpfs.state, bpfs.state != nil
}

// StateMust returns the state for [BackupPolicyFileShare]. Panics if the state is nil.
func (bpfs *BackupPolicyFileShare) StateMust() *backupPolicyFileShareState {
	if bpfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bpfs.Type(), bpfs.LocalName()))
	}
	return bpfs.state
}

// BackupPolicyFileShareArgs contains the configurations for azurerm_backup_policy_file_share.
type BackupPolicyFileShareArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timezone: string, optional
	Timezone terra.StringValue `hcl:"timezone,attr"`
	// Backup: required
	Backup *backuppolicyfileshare.Backup `hcl:"backup,block" validate:"required"`
	// RetentionDaily: required
	RetentionDaily *backuppolicyfileshare.RetentionDaily `hcl:"retention_daily,block" validate:"required"`
	// RetentionMonthly: optional
	RetentionMonthly *backuppolicyfileshare.RetentionMonthly `hcl:"retention_monthly,block"`
	// RetentionWeekly: optional
	RetentionWeekly *backuppolicyfileshare.RetentionWeekly `hcl:"retention_weekly,block"`
	// RetentionYearly: optional
	RetentionYearly *backuppolicyfileshare.RetentionYearly `hcl:"retention_yearly,block"`
	// Timeouts: optional
	Timeouts *backuppolicyfileshare.Timeouts `hcl:"timeouts,block"`
}
type backupPolicyFileShareAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_backup_policy_file_share.
func (bpfs backupPolicyFileShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_backup_policy_file_share.
func (bpfs backupPolicyFileShareAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("name"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_backup_policy_file_share.
func (bpfs backupPolicyFileShareAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_backup_policy_file_share.
func (bpfs backupPolicyFileShareAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("resource_group_name"))
}

// Timezone returns a reference to field timezone of azurerm_backup_policy_file_share.
func (bpfs backupPolicyFileShareAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("timezone"))
}

func (bpfs backupPolicyFileShareAttributes) Backup() terra.ListValue[backuppolicyfileshare.BackupAttributes] {
	return terra.ReferenceAsList[backuppolicyfileshare.BackupAttributes](bpfs.ref.Append("backup"))
}

func (bpfs backupPolicyFileShareAttributes) RetentionDaily() terra.ListValue[backuppolicyfileshare.RetentionDailyAttributes] {
	return terra.ReferenceAsList[backuppolicyfileshare.RetentionDailyAttributes](bpfs.ref.Append("retention_daily"))
}

func (bpfs backupPolicyFileShareAttributes) RetentionMonthly() terra.ListValue[backuppolicyfileshare.RetentionMonthlyAttributes] {
	return terra.ReferenceAsList[backuppolicyfileshare.RetentionMonthlyAttributes](bpfs.ref.Append("retention_monthly"))
}

func (bpfs backupPolicyFileShareAttributes) RetentionWeekly() terra.ListValue[backuppolicyfileshare.RetentionWeeklyAttributes] {
	return terra.ReferenceAsList[backuppolicyfileshare.RetentionWeeklyAttributes](bpfs.ref.Append("retention_weekly"))
}

func (bpfs backupPolicyFileShareAttributes) RetentionYearly() terra.ListValue[backuppolicyfileshare.RetentionYearlyAttributes] {
	return terra.ReferenceAsList[backuppolicyfileshare.RetentionYearlyAttributes](bpfs.ref.Append("retention_yearly"))
}

func (bpfs backupPolicyFileShareAttributes) Timeouts() backuppolicyfileshare.TimeoutsAttributes {
	return terra.ReferenceAsSingle[backuppolicyfileshare.TimeoutsAttributes](bpfs.ref.Append("timeouts"))
}

type backupPolicyFileShareState struct {
	Id                string                                        `json:"id"`
	Name              string                                        `json:"name"`
	RecoveryVaultName string                                        `json:"recovery_vault_name"`
	ResourceGroupName string                                        `json:"resource_group_name"`
	Timezone          string                                        `json:"timezone"`
	Backup            []backuppolicyfileshare.BackupState           `json:"backup"`
	RetentionDaily    []backuppolicyfileshare.RetentionDailyState   `json:"retention_daily"`
	RetentionMonthly  []backuppolicyfileshare.RetentionMonthlyState `json:"retention_monthly"`
	RetentionWeekly   []backuppolicyfileshare.RetentionWeeklyState  `json:"retention_weekly"`
	RetentionYearly   []backuppolicyfileshare.RetentionYearlyState  `json:"retention_yearly"`
	Timeouts          *backuppolicyfileshare.TimeoutsState          `json:"timeouts"`
}
