// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storageaccountcustomermanagedkey "github.com/golingon/terraproviders/azurerm/3.49.0/storageaccountcustomermanagedkey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageAccountCustomerManagedKey creates a new instance of [StorageAccountCustomerManagedKey].
func NewStorageAccountCustomerManagedKey(name string, args StorageAccountCustomerManagedKeyArgs) *StorageAccountCustomerManagedKey {
	return &StorageAccountCustomerManagedKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageAccountCustomerManagedKey)(nil)

// StorageAccountCustomerManagedKey represents the Terraform resource azurerm_storage_account_customer_managed_key.
type StorageAccountCustomerManagedKey struct {
	Name      string
	Args      StorageAccountCustomerManagedKeyArgs
	state     *storageAccountCustomerManagedKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageAccountCustomerManagedKey].
func (sacmk *StorageAccountCustomerManagedKey) Type() string {
	return "azurerm_storage_account_customer_managed_key"
}

// LocalName returns the local name for [StorageAccountCustomerManagedKey].
func (sacmk *StorageAccountCustomerManagedKey) LocalName() string {
	return sacmk.Name
}

// Configuration returns the configuration (args) for [StorageAccountCustomerManagedKey].
func (sacmk *StorageAccountCustomerManagedKey) Configuration() interface{} {
	return sacmk.Args
}

// DependOn is used for other resources to depend on [StorageAccountCustomerManagedKey].
func (sacmk *StorageAccountCustomerManagedKey) DependOn() terra.Reference {
	return terra.ReferenceResource(sacmk)
}

// Dependencies returns the list of resources [StorageAccountCustomerManagedKey] depends_on.
func (sacmk *StorageAccountCustomerManagedKey) Dependencies() terra.Dependencies {
	return sacmk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageAccountCustomerManagedKey].
func (sacmk *StorageAccountCustomerManagedKey) LifecycleManagement() *terra.Lifecycle {
	return sacmk.Lifecycle
}

// Attributes returns the attributes for [StorageAccountCustomerManagedKey].
func (sacmk *StorageAccountCustomerManagedKey) Attributes() storageAccountCustomerManagedKeyAttributes {
	return storageAccountCustomerManagedKeyAttributes{ref: terra.ReferenceResource(sacmk)}
}

// ImportState imports the given attribute values into [StorageAccountCustomerManagedKey]'s state.
func (sacmk *StorageAccountCustomerManagedKey) ImportState(av io.Reader) error {
	sacmk.state = &storageAccountCustomerManagedKeyState{}
	if err := json.NewDecoder(av).Decode(sacmk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sacmk.Type(), sacmk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageAccountCustomerManagedKey] has state.
func (sacmk *StorageAccountCustomerManagedKey) State() (*storageAccountCustomerManagedKeyState, bool) {
	return sacmk.state, sacmk.state != nil
}

// StateMust returns the state for [StorageAccountCustomerManagedKey]. Panics if the state is nil.
func (sacmk *StorageAccountCustomerManagedKey) StateMust() *storageAccountCustomerManagedKeyState {
	if sacmk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sacmk.Type(), sacmk.LocalName()))
	}
	return sacmk.state
}

// StorageAccountCustomerManagedKeyArgs contains the configurations for azurerm_storage_account_customer_managed_key.
type StorageAccountCustomerManagedKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyName: string, required
	KeyName terra.StringValue `hcl:"key_name,attr" validate:"required"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// KeyVersion: string, optional
	KeyVersion terra.StringValue `hcl:"key_version,attr"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// UserAssignedIdentityId: string, optional
	UserAssignedIdentityId terra.StringValue `hcl:"user_assigned_identity_id,attr"`
	// Timeouts: optional
	Timeouts *storageaccountcustomermanagedkey.Timeouts `hcl:"timeouts,block"`
}
type storageAccountCustomerManagedKeyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_account_customer_managed_key.
func (sacmk storageAccountCustomerManagedKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sacmk.ref.Append("id"))
}

// KeyName returns a reference to field key_name of azurerm_storage_account_customer_managed_key.
func (sacmk storageAccountCustomerManagedKeyAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(sacmk.ref.Append("key_name"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_storage_account_customer_managed_key.
func (sacmk storageAccountCustomerManagedKeyAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(sacmk.ref.Append("key_vault_id"))
}

// KeyVersion returns a reference to field key_version of azurerm_storage_account_customer_managed_key.
func (sacmk storageAccountCustomerManagedKeyAttributes) KeyVersion() terra.StringValue {
	return terra.ReferenceAsString(sacmk.ref.Append("key_version"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_storage_account_customer_managed_key.
func (sacmk storageAccountCustomerManagedKeyAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(sacmk.ref.Append("storage_account_id"))
}

// UserAssignedIdentityId returns a reference to field user_assigned_identity_id of azurerm_storage_account_customer_managed_key.
func (sacmk storageAccountCustomerManagedKeyAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(sacmk.ref.Append("user_assigned_identity_id"))
}

func (sacmk storageAccountCustomerManagedKeyAttributes) Timeouts() storageaccountcustomermanagedkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storageaccountcustomermanagedkey.TimeoutsAttributes](sacmk.ref.Append("timeouts"))
}

type storageAccountCustomerManagedKeyState struct {
	Id                     string                                          `json:"id"`
	KeyName                string                                          `json:"key_name"`
	KeyVaultId             string                                          `json:"key_vault_id"`
	KeyVersion             string                                          `json:"key_version"`
	StorageAccountId       string                                          `json:"storage_account_id"`
	UserAssignedIdentityId string                                          `json:"user_assigned_identity_id"`
	Timeouts               *storageaccountcustomermanagedkey.TimeoutsState `json:"timeouts"`
}