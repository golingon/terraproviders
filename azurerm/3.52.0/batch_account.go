// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	batchaccount "github.com/golingon/terraproviders/azurerm/3.52.0/batchaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBatchAccount creates a new instance of [BatchAccount].
func NewBatchAccount(name string, args BatchAccountArgs) *BatchAccount {
	return &BatchAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BatchAccount)(nil)

// BatchAccount represents the Terraform resource azurerm_batch_account.
type BatchAccount struct {
	Name      string
	Args      BatchAccountArgs
	state     *batchAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BatchAccount].
func (ba *BatchAccount) Type() string {
	return "azurerm_batch_account"
}

// LocalName returns the local name for [BatchAccount].
func (ba *BatchAccount) LocalName() string {
	return ba.Name
}

// Configuration returns the configuration (args) for [BatchAccount].
func (ba *BatchAccount) Configuration() interface{} {
	return ba.Args
}

// DependOn is used for other resources to depend on [BatchAccount].
func (ba *BatchAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(ba)
}

// Dependencies returns the list of resources [BatchAccount] depends_on.
func (ba *BatchAccount) Dependencies() terra.Dependencies {
	return ba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BatchAccount].
func (ba *BatchAccount) LifecycleManagement() *terra.Lifecycle {
	return ba.Lifecycle
}

// Attributes returns the attributes for [BatchAccount].
func (ba *BatchAccount) Attributes() batchAccountAttributes {
	return batchAccountAttributes{ref: terra.ReferenceResource(ba)}
}

// ImportState imports the given attribute values into [BatchAccount]'s state.
func (ba *BatchAccount) ImportState(av io.Reader) error {
	ba.state = &batchAccountState{}
	if err := json.NewDecoder(av).Decode(ba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ba.Type(), ba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BatchAccount] has state.
func (ba *BatchAccount) State() (*batchAccountState, bool) {
	return ba.state, ba.state != nil
}

// StateMust returns the state for [BatchAccount]. Panics if the state is nil.
func (ba *BatchAccount) StateMust() *batchAccountState {
	if ba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ba.Type(), ba.LocalName()))
	}
	return ba.state
}

// BatchAccountArgs contains the configurations for azurerm_batch_account.
type BatchAccountArgs struct {
	// AllowedAuthenticationModes: set of string, optional
	AllowedAuthenticationModes terra.SetValue[terra.StringValue] `hcl:"allowed_authentication_modes,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PoolAllocationMode: string, optional
	PoolAllocationMode terra.StringValue `hcl:"pool_allocation_mode,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageAccountAuthenticationMode: string, optional
	StorageAccountAuthenticationMode terra.StringValue `hcl:"storage_account_authentication_mode,attr"`
	// StorageAccountId: string, optional
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr"`
	// StorageAccountNodeIdentity: string, optional
	StorageAccountNodeIdentity terra.StringValue `hcl:"storage_account_node_identity,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Encryption: min=0
	Encryption []batchaccount.Encryption `hcl:"encryption,block" validate:"min=0"`
	// Identity: optional
	Identity *batchaccount.Identity `hcl:"identity,block"`
	// KeyVaultReference: optional
	KeyVaultReference *batchaccount.KeyVaultReference `hcl:"key_vault_reference,block"`
	// Timeouts: optional
	Timeouts *batchaccount.Timeouts `hcl:"timeouts,block"`
}
type batchAccountAttributes struct {
	ref terra.Reference
}

// AccountEndpoint returns a reference to field account_endpoint of azurerm_batch_account.
func (ba batchAccountAttributes) AccountEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("account_endpoint"))
}

// AllowedAuthenticationModes returns a reference to field allowed_authentication_modes of azurerm_batch_account.
func (ba batchAccountAttributes) AllowedAuthenticationModes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ba.ref.Append("allowed_authentication_modes"))
}

// Id returns a reference to field id of azurerm_batch_account.
func (ba batchAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_batch_account.
func (ba batchAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_batch_account.
func (ba batchAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("name"))
}

// PoolAllocationMode returns a reference to field pool_allocation_mode of azurerm_batch_account.
func (ba batchAccountAttributes) PoolAllocationMode() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("pool_allocation_mode"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_batch_account.
func (ba batchAccountAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("primary_access_key"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_batch_account.
func (ba batchAccountAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ba.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_batch_account.
func (ba batchAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_batch_account.
func (ba batchAccountAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("secondary_access_key"))
}

// StorageAccountAuthenticationMode returns a reference to field storage_account_authentication_mode of azurerm_batch_account.
func (ba batchAccountAttributes) StorageAccountAuthenticationMode() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("storage_account_authentication_mode"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_batch_account.
func (ba batchAccountAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("storage_account_id"))
}

// StorageAccountNodeIdentity returns a reference to field storage_account_node_identity of azurerm_batch_account.
func (ba batchAccountAttributes) StorageAccountNodeIdentity() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("storage_account_node_identity"))
}

// Tags returns a reference to field tags of azurerm_batch_account.
func (ba batchAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ba.ref.Append("tags"))
}

func (ba batchAccountAttributes) Encryption() terra.ListValue[batchaccount.EncryptionAttributes] {
	return terra.ReferenceAsList[batchaccount.EncryptionAttributes](ba.ref.Append("encryption"))
}

func (ba batchAccountAttributes) Identity() terra.ListValue[batchaccount.IdentityAttributes] {
	return terra.ReferenceAsList[batchaccount.IdentityAttributes](ba.ref.Append("identity"))
}

func (ba batchAccountAttributes) KeyVaultReference() terra.ListValue[batchaccount.KeyVaultReferenceAttributes] {
	return terra.ReferenceAsList[batchaccount.KeyVaultReferenceAttributes](ba.ref.Append("key_vault_reference"))
}

func (ba batchAccountAttributes) Timeouts() batchaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[batchaccount.TimeoutsAttributes](ba.ref.Append("timeouts"))
}

type batchAccountState struct {
	AccountEndpoint                  string                                `json:"account_endpoint"`
	AllowedAuthenticationModes       []string                              `json:"allowed_authentication_modes"`
	Id                               string                                `json:"id"`
	Location                         string                                `json:"location"`
	Name                             string                                `json:"name"`
	PoolAllocationMode               string                                `json:"pool_allocation_mode"`
	PrimaryAccessKey                 string                                `json:"primary_access_key"`
	PublicNetworkAccessEnabled       bool                                  `json:"public_network_access_enabled"`
	ResourceGroupName                string                                `json:"resource_group_name"`
	SecondaryAccessKey               string                                `json:"secondary_access_key"`
	StorageAccountAuthenticationMode string                                `json:"storage_account_authentication_mode"`
	StorageAccountId                 string                                `json:"storage_account_id"`
	StorageAccountNodeIdentity       string                                `json:"storage_account_node_identity"`
	Tags                             map[string]string                     `json:"tags"`
	Encryption                       []batchaccount.EncryptionState        `json:"encryption"`
	Identity                         []batchaccount.IdentityState          `json:"identity"`
	KeyVaultReference                []batchaccount.KeyVaultReferenceState `json:"key_vault_reference"`
	Timeouts                         *batchaccount.TimeoutsState           `json:"timeouts"`
}
