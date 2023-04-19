// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	backupcontainerstorageaccount "github.com/golingon/terraproviders/azurerm/3.52.0/backupcontainerstorageaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBackupContainerStorageAccount creates a new instance of [BackupContainerStorageAccount].
func NewBackupContainerStorageAccount(name string, args BackupContainerStorageAccountArgs) *BackupContainerStorageAccount {
	return &BackupContainerStorageAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupContainerStorageAccount)(nil)

// BackupContainerStorageAccount represents the Terraform resource azurerm_backup_container_storage_account.
type BackupContainerStorageAccount struct {
	Name      string
	Args      BackupContainerStorageAccountArgs
	state     *backupContainerStorageAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BackupContainerStorageAccount].
func (bcsa *BackupContainerStorageAccount) Type() string {
	return "azurerm_backup_container_storage_account"
}

// LocalName returns the local name for [BackupContainerStorageAccount].
func (bcsa *BackupContainerStorageAccount) LocalName() string {
	return bcsa.Name
}

// Configuration returns the configuration (args) for [BackupContainerStorageAccount].
func (bcsa *BackupContainerStorageAccount) Configuration() interface{} {
	return bcsa.Args
}

// DependOn is used for other resources to depend on [BackupContainerStorageAccount].
func (bcsa *BackupContainerStorageAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(bcsa)
}

// Dependencies returns the list of resources [BackupContainerStorageAccount] depends_on.
func (bcsa *BackupContainerStorageAccount) Dependencies() terra.Dependencies {
	return bcsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BackupContainerStorageAccount].
func (bcsa *BackupContainerStorageAccount) LifecycleManagement() *terra.Lifecycle {
	return bcsa.Lifecycle
}

// Attributes returns the attributes for [BackupContainerStorageAccount].
func (bcsa *BackupContainerStorageAccount) Attributes() backupContainerStorageAccountAttributes {
	return backupContainerStorageAccountAttributes{ref: terra.ReferenceResource(bcsa)}
}

// ImportState imports the given attribute values into [BackupContainerStorageAccount]'s state.
func (bcsa *BackupContainerStorageAccount) ImportState(av io.Reader) error {
	bcsa.state = &backupContainerStorageAccountState{}
	if err := json.NewDecoder(av).Decode(bcsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcsa.Type(), bcsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BackupContainerStorageAccount] has state.
func (bcsa *BackupContainerStorageAccount) State() (*backupContainerStorageAccountState, bool) {
	return bcsa.state, bcsa.state != nil
}

// StateMust returns the state for [BackupContainerStorageAccount]. Panics if the state is nil.
func (bcsa *BackupContainerStorageAccount) StateMust() *backupContainerStorageAccountState {
	if bcsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcsa.Type(), bcsa.LocalName()))
	}
	return bcsa.state
}

// BackupContainerStorageAccountArgs contains the configurations for azurerm_backup_container_storage_account.
type BackupContainerStorageAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *backupcontainerstorageaccount.Timeouts `hcl:"timeouts,block"`
}
type backupContainerStorageAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_backup_container_storage_account.
func (bcsa backupContainerStorageAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcsa.ref.Append("id"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_backup_container_storage_account.
func (bcsa backupContainerStorageAccountAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(bcsa.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_backup_container_storage_account.
func (bcsa backupContainerStorageAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bcsa.ref.Append("resource_group_name"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_backup_container_storage_account.
func (bcsa backupContainerStorageAccountAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(bcsa.ref.Append("storage_account_id"))
}

func (bcsa backupContainerStorageAccountAttributes) Timeouts() backupcontainerstorageaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[backupcontainerstorageaccount.TimeoutsAttributes](bcsa.ref.Append("timeouts"))
}

type backupContainerStorageAccountState struct {
	Id                string                                       `json:"id"`
	RecoveryVaultName string                                       `json:"recovery_vault_name"`
	ResourceGroupName string                                       `json:"resource_group_name"`
	StorageAccountId  string                                       `json:"storage_account_id"`
	Timeouts          *backupcontainerstorageaccount.TimeoutsState `json:"timeouts"`
}
