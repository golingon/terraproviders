// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_backup_container_storage_account

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

// Resource represents the Terraform resource azurerm_backup_container_storage_account.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermBackupContainerStorageAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (abcsa *Resource) Type() string {
	return "azurerm_backup_container_storage_account"
}

// LocalName returns the local name for [Resource].
func (abcsa *Resource) LocalName() string {
	return abcsa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (abcsa *Resource) Configuration() interface{} {
	return abcsa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (abcsa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(abcsa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (abcsa *Resource) Dependencies() terra.Dependencies {
	return abcsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (abcsa *Resource) LifecycleManagement() *terra.Lifecycle {
	return abcsa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (abcsa *Resource) Attributes() azurermBackupContainerStorageAccountAttributes {
	return azurermBackupContainerStorageAccountAttributes{ref: terra.ReferenceResource(abcsa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (abcsa *Resource) ImportState(state io.Reader) error {
	abcsa.state = &azurermBackupContainerStorageAccountState{}
	if err := json.NewDecoder(state).Decode(abcsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", abcsa.Type(), abcsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (abcsa *Resource) State() (*azurermBackupContainerStorageAccountState, bool) {
	return abcsa.state, abcsa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (abcsa *Resource) StateMust() *azurermBackupContainerStorageAccountState {
	if abcsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", abcsa.Type(), abcsa.LocalName()))
	}
	return abcsa.state
}

// Args contains the configurations for azurerm_backup_container_storage_account.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermBackupContainerStorageAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_backup_container_storage_account.
func (abcsa azurermBackupContainerStorageAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(abcsa.ref.Append("id"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_backup_container_storage_account.
func (abcsa azurermBackupContainerStorageAccountAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(abcsa.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_backup_container_storage_account.
func (abcsa azurermBackupContainerStorageAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(abcsa.ref.Append("resource_group_name"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_backup_container_storage_account.
func (abcsa azurermBackupContainerStorageAccountAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(abcsa.ref.Append("storage_account_id"))
}

func (abcsa azurermBackupContainerStorageAccountAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](abcsa.ref.Append("timeouts"))
}

type azurermBackupContainerStorageAccountState struct {
	Id                string         `json:"id"`
	RecoveryVaultName string         `json:"recovery_vault_name"`
	ResourceGroupName string         `json:"resource_group_name"`
	StorageAccountId  string         `json:"storage_account_id"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
