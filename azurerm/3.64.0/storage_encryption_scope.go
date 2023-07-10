// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storageencryptionscope "github.com/golingon/terraproviders/azurerm/3.64.0/storageencryptionscope"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageEncryptionScope creates a new instance of [StorageEncryptionScope].
func NewStorageEncryptionScope(name string, args StorageEncryptionScopeArgs) *StorageEncryptionScope {
	return &StorageEncryptionScope{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageEncryptionScope)(nil)

// StorageEncryptionScope represents the Terraform resource azurerm_storage_encryption_scope.
type StorageEncryptionScope struct {
	Name      string
	Args      StorageEncryptionScopeArgs
	state     *storageEncryptionScopeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageEncryptionScope].
func (ses *StorageEncryptionScope) Type() string {
	return "azurerm_storage_encryption_scope"
}

// LocalName returns the local name for [StorageEncryptionScope].
func (ses *StorageEncryptionScope) LocalName() string {
	return ses.Name
}

// Configuration returns the configuration (args) for [StorageEncryptionScope].
func (ses *StorageEncryptionScope) Configuration() interface{} {
	return ses.Args
}

// DependOn is used for other resources to depend on [StorageEncryptionScope].
func (ses *StorageEncryptionScope) DependOn() terra.Reference {
	return terra.ReferenceResource(ses)
}

// Dependencies returns the list of resources [StorageEncryptionScope] depends_on.
func (ses *StorageEncryptionScope) Dependencies() terra.Dependencies {
	return ses.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageEncryptionScope].
func (ses *StorageEncryptionScope) LifecycleManagement() *terra.Lifecycle {
	return ses.Lifecycle
}

// Attributes returns the attributes for [StorageEncryptionScope].
func (ses *StorageEncryptionScope) Attributes() storageEncryptionScopeAttributes {
	return storageEncryptionScopeAttributes{ref: terra.ReferenceResource(ses)}
}

// ImportState imports the given attribute values into [StorageEncryptionScope]'s state.
func (ses *StorageEncryptionScope) ImportState(av io.Reader) error {
	ses.state = &storageEncryptionScopeState{}
	if err := json.NewDecoder(av).Decode(ses.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ses.Type(), ses.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageEncryptionScope] has state.
func (ses *StorageEncryptionScope) State() (*storageEncryptionScopeState, bool) {
	return ses.state, ses.state != nil
}

// StateMust returns the state for [StorageEncryptionScope]. Panics if the state is nil.
func (ses *StorageEncryptionScope) StateMust() *storageEncryptionScopeState {
	if ses.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ses.Type(), ses.LocalName()))
	}
	return ses.state
}

// StorageEncryptionScopeArgs contains the configurations for azurerm_storage_encryption_scope.
type StorageEncryptionScopeArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InfrastructureEncryptionRequired: bool, optional
	InfrastructureEncryptionRequired terra.BoolValue `hcl:"infrastructure_encryption_required,attr"`
	// KeyVaultKeyId: string, optional
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *storageencryptionscope.Timeouts `hcl:"timeouts,block"`
}
type storageEncryptionScopeAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_encryption_scope.
func (ses storageEncryptionScopeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ses.ref.Append("id"))
}

// InfrastructureEncryptionRequired returns a reference to field infrastructure_encryption_required of azurerm_storage_encryption_scope.
func (ses storageEncryptionScopeAttributes) InfrastructureEncryptionRequired() terra.BoolValue {
	return terra.ReferenceAsBool(ses.ref.Append("infrastructure_encryption_required"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_storage_encryption_scope.
func (ses storageEncryptionScopeAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(ses.ref.Append("key_vault_key_id"))
}

// Name returns a reference to field name of azurerm_storage_encryption_scope.
func (ses storageEncryptionScopeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ses.ref.Append("name"))
}

// Source returns a reference to field source of azurerm_storage_encryption_scope.
func (ses storageEncryptionScopeAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(ses.ref.Append("source"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_storage_encryption_scope.
func (ses storageEncryptionScopeAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(ses.ref.Append("storage_account_id"))
}

func (ses storageEncryptionScopeAttributes) Timeouts() storageencryptionscope.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storageencryptionscope.TimeoutsAttributes](ses.ref.Append("timeouts"))
}

type storageEncryptionScopeState struct {
	Id                               string                                `json:"id"`
	InfrastructureEncryptionRequired bool                                  `json:"infrastructure_encryption_required"`
	KeyVaultKeyId                    string                                `json:"key_vault_key_id"`
	Name                             string                                `json:"name"`
	Source                           string                                `json:"source"`
	StorageAccountId                 string                                `json:"storage_account_id"`
	Timeouts                         *storageencryptionscope.TimeoutsState `json:"timeouts"`
}
