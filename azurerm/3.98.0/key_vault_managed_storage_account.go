// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	keyvaultmanagedstorageaccount "github.com/golingon/terraproviders/azurerm/3.98.0/keyvaultmanagedstorageaccount"
	"io"
)

// NewKeyVaultManagedStorageAccount creates a new instance of [KeyVaultManagedStorageAccount].
func NewKeyVaultManagedStorageAccount(name string, args KeyVaultManagedStorageAccountArgs) *KeyVaultManagedStorageAccount {
	return &KeyVaultManagedStorageAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyVaultManagedStorageAccount)(nil)

// KeyVaultManagedStorageAccount represents the Terraform resource azurerm_key_vault_managed_storage_account.
type KeyVaultManagedStorageAccount struct {
	Name      string
	Args      KeyVaultManagedStorageAccountArgs
	state     *keyVaultManagedStorageAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KeyVaultManagedStorageAccount].
func (kvmsa *KeyVaultManagedStorageAccount) Type() string {
	return "azurerm_key_vault_managed_storage_account"
}

// LocalName returns the local name for [KeyVaultManagedStorageAccount].
func (kvmsa *KeyVaultManagedStorageAccount) LocalName() string {
	return kvmsa.Name
}

// Configuration returns the configuration (args) for [KeyVaultManagedStorageAccount].
func (kvmsa *KeyVaultManagedStorageAccount) Configuration() interface{} {
	return kvmsa.Args
}

// DependOn is used for other resources to depend on [KeyVaultManagedStorageAccount].
func (kvmsa *KeyVaultManagedStorageAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(kvmsa)
}

// Dependencies returns the list of resources [KeyVaultManagedStorageAccount] depends_on.
func (kvmsa *KeyVaultManagedStorageAccount) Dependencies() terra.Dependencies {
	return kvmsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KeyVaultManagedStorageAccount].
func (kvmsa *KeyVaultManagedStorageAccount) LifecycleManagement() *terra.Lifecycle {
	return kvmsa.Lifecycle
}

// Attributes returns the attributes for [KeyVaultManagedStorageAccount].
func (kvmsa *KeyVaultManagedStorageAccount) Attributes() keyVaultManagedStorageAccountAttributes {
	return keyVaultManagedStorageAccountAttributes{ref: terra.ReferenceResource(kvmsa)}
}

// ImportState imports the given attribute values into [KeyVaultManagedStorageAccount]'s state.
func (kvmsa *KeyVaultManagedStorageAccount) ImportState(av io.Reader) error {
	kvmsa.state = &keyVaultManagedStorageAccountState{}
	if err := json.NewDecoder(av).Decode(kvmsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kvmsa.Type(), kvmsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KeyVaultManagedStorageAccount] has state.
func (kvmsa *KeyVaultManagedStorageAccount) State() (*keyVaultManagedStorageAccountState, bool) {
	return kvmsa.state, kvmsa.state != nil
}

// StateMust returns the state for [KeyVaultManagedStorageAccount]. Panics if the state is nil.
func (kvmsa *KeyVaultManagedStorageAccount) StateMust() *keyVaultManagedStorageAccountState {
	if kvmsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kvmsa.Type(), kvmsa.LocalName()))
	}
	return kvmsa.state
}

// KeyVaultManagedStorageAccountArgs contains the configurations for azurerm_key_vault_managed_storage_account.
type KeyVaultManagedStorageAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RegenerateKeyAutomatically: bool, optional
	RegenerateKeyAutomatically terra.BoolValue `hcl:"regenerate_key_automatically,attr"`
	// RegenerationPeriod: string, optional
	RegenerationPeriod terra.StringValue `hcl:"regeneration_period,attr"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// StorageAccountKey: string, required
	StorageAccountKey terra.StringValue `hcl:"storage_account_key,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *keyvaultmanagedstorageaccount.Timeouts `hcl:"timeouts,block"`
}
type keyVaultManagedStorageAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_key_vault_managed_storage_account.
func (kvmsa keyVaultManagedStorageAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvmsa.ref.Append("id"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_key_vault_managed_storage_account.
func (kvmsa keyVaultManagedStorageAccountAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kvmsa.ref.Append("key_vault_id"))
}

// Name returns a reference to field name of azurerm_key_vault_managed_storage_account.
func (kvmsa keyVaultManagedStorageAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kvmsa.ref.Append("name"))
}

// RegenerateKeyAutomatically returns a reference to field regenerate_key_automatically of azurerm_key_vault_managed_storage_account.
func (kvmsa keyVaultManagedStorageAccountAttributes) RegenerateKeyAutomatically() terra.BoolValue {
	return terra.ReferenceAsBool(kvmsa.ref.Append("regenerate_key_automatically"))
}

// RegenerationPeriod returns a reference to field regeneration_period of azurerm_key_vault_managed_storage_account.
func (kvmsa keyVaultManagedStorageAccountAttributes) RegenerationPeriod() terra.StringValue {
	return terra.ReferenceAsString(kvmsa.ref.Append("regeneration_period"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_key_vault_managed_storage_account.
func (kvmsa keyVaultManagedStorageAccountAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(kvmsa.ref.Append("storage_account_id"))
}

// StorageAccountKey returns a reference to field storage_account_key of azurerm_key_vault_managed_storage_account.
func (kvmsa keyVaultManagedStorageAccountAttributes) StorageAccountKey() terra.StringValue {
	return terra.ReferenceAsString(kvmsa.ref.Append("storage_account_key"))
}

// Tags returns a reference to field tags of azurerm_key_vault_managed_storage_account.
func (kvmsa keyVaultManagedStorageAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kvmsa.ref.Append("tags"))
}

func (kvmsa keyVaultManagedStorageAccountAttributes) Timeouts() keyvaultmanagedstorageaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[keyvaultmanagedstorageaccount.TimeoutsAttributes](kvmsa.ref.Append("timeouts"))
}

type keyVaultManagedStorageAccountState struct {
	Id                         string                                       `json:"id"`
	KeyVaultId                 string                                       `json:"key_vault_id"`
	Name                       string                                       `json:"name"`
	RegenerateKeyAutomatically bool                                         `json:"regenerate_key_automatically"`
	RegenerationPeriod         string                                       `json:"regeneration_period"`
	StorageAccountId           string                                       `json:"storage_account_id"`
	StorageAccountKey          string                                       `json:"storage_account_key"`
	Tags                       map[string]string                            `json:"tags"`
	Timeouts                   *keyvaultmanagedstorageaccount.TimeoutsState `json:"timeouts"`
}
