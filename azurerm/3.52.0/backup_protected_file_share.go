// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	backupprotectedfileshare "github.com/golingon/terraproviders/azurerm/3.52.0/backupprotectedfileshare"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBackupProtectedFileShare creates a new instance of [BackupProtectedFileShare].
func NewBackupProtectedFileShare(name string, args BackupProtectedFileShareArgs) *BackupProtectedFileShare {
	return &BackupProtectedFileShare{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupProtectedFileShare)(nil)

// BackupProtectedFileShare represents the Terraform resource azurerm_backup_protected_file_share.
type BackupProtectedFileShare struct {
	Name      string
	Args      BackupProtectedFileShareArgs
	state     *backupProtectedFileShareState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BackupProtectedFileShare].
func (bpfs *BackupProtectedFileShare) Type() string {
	return "azurerm_backup_protected_file_share"
}

// LocalName returns the local name for [BackupProtectedFileShare].
func (bpfs *BackupProtectedFileShare) LocalName() string {
	return bpfs.Name
}

// Configuration returns the configuration (args) for [BackupProtectedFileShare].
func (bpfs *BackupProtectedFileShare) Configuration() interface{} {
	return bpfs.Args
}

// DependOn is used for other resources to depend on [BackupProtectedFileShare].
func (bpfs *BackupProtectedFileShare) DependOn() terra.Reference {
	return terra.ReferenceResource(bpfs)
}

// Dependencies returns the list of resources [BackupProtectedFileShare] depends_on.
func (bpfs *BackupProtectedFileShare) Dependencies() terra.Dependencies {
	return bpfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BackupProtectedFileShare].
func (bpfs *BackupProtectedFileShare) LifecycleManagement() *terra.Lifecycle {
	return bpfs.Lifecycle
}

// Attributes returns the attributes for [BackupProtectedFileShare].
func (bpfs *BackupProtectedFileShare) Attributes() backupProtectedFileShareAttributes {
	return backupProtectedFileShareAttributes{ref: terra.ReferenceResource(bpfs)}
}

// ImportState imports the given attribute values into [BackupProtectedFileShare]'s state.
func (bpfs *BackupProtectedFileShare) ImportState(av io.Reader) error {
	bpfs.state = &backupProtectedFileShareState{}
	if err := json.NewDecoder(av).Decode(bpfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bpfs.Type(), bpfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BackupProtectedFileShare] has state.
func (bpfs *BackupProtectedFileShare) State() (*backupProtectedFileShareState, bool) {
	return bpfs.state, bpfs.state != nil
}

// StateMust returns the state for [BackupProtectedFileShare]. Panics if the state is nil.
func (bpfs *BackupProtectedFileShare) StateMust() *backupProtectedFileShareState {
	if bpfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bpfs.Type(), bpfs.LocalName()))
	}
	return bpfs.state
}

// BackupProtectedFileShareArgs contains the configurations for azurerm_backup_protected_file_share.
type BackupProtectedFileShareArgs struct {
	// BackupPolicyId: string, required
	BackupPolicyId terra.StringValue `hcl:"backup_policy_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SourceFileShareName: string, required
	SourceFileShareName terra.StringValue `hcl:"source_file_share_name,attr" validate:"required"`
	// SourceStorageAccountId: string, required
	SourceStorageAccountId terra.StringValue `hcl:"source_storage_account_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *backupprotectedfileshare.Timeouts `hcl:"timeouts,block"`
}
type backupProtectedFileShareAttributes struct {
	ref terra.Reference
}

// BackupPolicyId returns a reference to field backup_policy_id of azurerm_backup_protected_file_share.
func (bpfs backupProtectedFileShareAttributes) BackupPolicyId() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("backup_policy_id"))
}

// Id returns a reference to field id of azurerm_backup_protected_file_share.
func (bpfs backupProtectedFileShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("id"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_backup_protected_file_share.
func (bpfs backupProtectedFileShareAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_backup_protected_file_share.
func (bpfs backupProtectedFileShareAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("resource_group_name"))
}

// SourceFileShareName returns a reference to field source_file_share_name of azurerm_backup_protected_file_share.
func (bpfs backupProtectedFileShareAttributes) SourceFileShareName() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("source_file_share_name"))
}

// SourceStorageAccountId returns a reference to field source_storage_account_id of azurerm_backup_protected_file_share.
func (bpfs backupProtectedFileShareAttributes) SourceStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("source_storage_account_id"))
}

func (bpfs backupProtectedFileShareAttributes) Timeouts() backupprotectedfileshare.TimeoutsAttributes {
	return terra.ReferenceAsSingle[backupprotectedfileshare.TimeoutsAttributes](bpfs.ref.Append("timeouts"))
}

type backupProtectedFileShareState struct {
	BackupPolicyId         string                                  `json:"backup_policy_id"`
	Id                     string                                  `json:"id"`
	RecoveryVaultName      string                                  `json:"recovery_vault_name"`
	ResourceGroupName      string                                  `json:"resource_group_name"`
	SourceFileShareName    string                                  `json:"source_file_share_name"`
	SourceStorageAccountId string                                  `json:"source_storage_account_id"`
	Timeouts               *backupprotectedfileshare.TimeoutsState `json:"timeouts"`
}