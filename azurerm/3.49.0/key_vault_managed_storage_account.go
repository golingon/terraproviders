// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	keyvaultmanagedstorageaccount "github.com/golingon/terraproviders/azurerm/3.49.0/keyvaultmanagedstorageaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewKeyVaultManagedStorageAccount(name string, args KeyVaultManagedStorageAccountArgs) *KeyVaultManagedStorageAccount {
	return &KeyVaultManagedStorageAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyVaultManagedStorageAccount)(nil)

type KeyVaultManagedStorageAccount struct {
	Name  string
	Args  KeyVaultManagedStorageAccountArgs
	state *keyVaultManagedStorageAccountState
}

func (kvmsa *KeyVaultManagedStorageAccount) Type() string {
	return "azurerm_key_vault_managed_storage_account"
}

func (kvmsa *KeyVaultManagedStorageAccount) LocalName() string {
	return kvmsa.Name
}

func (kvmsa *KeyVaultManagedStorageAccount) Configuration() interface{} {
	return kvmsa.Args
}

func (kvmsa *KeyVaultManagedStorageAccount) Attributes() keyVaultManagedStorageAccountAttributes {
	return keyVaultManagedStorageAccountAttributes{ref: terra.ReferenceResource(kvmsa)}
}

func (kvmsa *KeyVaultManagedStorageAccount) ImportState(av io.Reader) error {
	kvmsa.state = &keyVaultManagedStorageAccountState{}
	if err := json.NewDecoder(av).Decode(kvmsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kvmsa.Type(), kvmsa.LocalName(), err)
	}
	return nil
}

func (kvmsa *KeyVaultManagedStorageAccount) State() (*keyVaultManagedStorageAccountState, bool) {
	return kvmsa.state, kvmsa.state != nil
}

func (kvmsa *KeyVaultManagedStorageAccount) StateMust() *keyVaultManagedStorageAccountState {
	if kvmsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kvmsa.Type(), kvmsa.LocalName()))
	}
	return kvmsa.state
}

func (kvmsa *KeyVaultManagedStorageAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(kvmsa)
}

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
	// DependsOn contains resources that KeyVaultManagedStorageAccount depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type keyVaultManagedStorageAccountAttributes struct {
	ref terra.Reference
}

func (kvmsa keyVaultManagedStorageAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceString(kvmsa.ref.Append("id"))
}

func (kvmsa keyVaultManagedStorageAccountAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(kvmsa.ref.Append("key_vault_id"))
}

func (kvmsa keyVaultManagedStorageAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceString(kvmsa.ref.Append("name"))
}

func (kvmsa keyVaultManagedStorageAccountAttributes) RegenerateKeyAutomatically() terra.BoolValue {
	return terra.ReferenceBool(kvmsa.ref.Append("regenerate_key_automatically"))
}

func (kvmsa keyVaultManagedStorageAccountAttributes) RegenerationPeriod() terra.StringValue {
	return terra.ReferenceString(kvmsa.ref.Append("regeneration_period"))
}

func (kvmsa keyVaultManagedStorageAccountAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceString(kvmsa.ref.Append("storage_account_id"))
}

func (kvmsa keyVaultManagedStorageAccountAttributes) StorageAccountKey() terra.StringValue {
	return terra.ReferenceString(kvmsa.ref.Append("storage_account_key"))
}

func (kvmsa keyVaultManagedStorageAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](kvmsa.ref.Append("tags"))
}

func (kvmsa keyVaultManagedStorageAccountAttributes) Timeouts() keyvaultmanagedstorageaccount.TimeoutsAttributes {
	return terra.ReferenceSingle[keyvaultmanagedstorageaccount.TimeoutsAttributes](kvmsa.ref.Append("timeouts"))
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
