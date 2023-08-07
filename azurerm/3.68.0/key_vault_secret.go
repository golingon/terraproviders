// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	keyvaultsecret "github.com/golingon/terraproviders/azurerm/3.68.0/keyvaultsecret"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKeyVaultSecret creates a new instance of [KeyVaultSecret].
func NewKeyVaultSecret(name string, args KeyVaultSecretArgs) *KeyVaultSecret {
	return &KeyVaultSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyVaultSecret)(nil)

// KeyVaultSecret represents the Terraform resource azurerm_key_vault_secret.
type KeyVaultSecret struct {
	Name      string
	Args      KeyVaultSecretArgs
	state     *keyVaultSecretState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KeyVaultSecret].
func (kvs *KeyVaultSecret) Type() string {
	return "azurerm_key_vault_secret"
}

// LocalName returns the local name for [KeyVaultSecret].
func (kvs *KeyVaultSecret) LocalName() string {
	return kvs.Name
}

// Configuration returns the configuration (args) for [KeyVaultSecret].
func (kvs *KeyVaultSecret) Configuration() interface{} {
	return kvs.Args
}

// DependOn is used for other resources to depend on [KeyVaultSecret].
func (kvs *KeyVaultSecret) DependOn() terra.Reference {
	return terra.ReferenceResource(kvs)
}

// Dependencies returns the list of resources [KeyVaultSecret] depends_on.
func (kvs *KeyVaultSecret) Dependencies() terra.Dependencies {
	return kvs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KeyVaultSecret].
func (kvs *KeyVaultSecret) LifecycleManagement() *terra.Lifecycle {
	return kvs.Lifecycle
}

// Attributes returns the attributes for [KeyVaultSecret].
func (kvs *KeyVaultSecret) Attributes() keyVaultSecretAttributes {
	return keyVaultSecretAttributes{ref: terra.ReferenceResource(kvs)}
}

// ImportState imports the given attribute values into [KeyVaultSecret]'s state.
func (kvs *KeyVaultSecret) ImportState(av io.Reader) error {
	kvs.state = &keyVaultSecretState{}
	if err := json.NewDecoder(av).Decode(kvs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kvs.Type(), kvs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KeyVaultSecret] has state.
func (kvs *KeyVaultSecret) State() (*keyVaultSecretState, bool) {
	return kvs.state, kvs.state != nil
}

// StateMust returns the state for [KeyVaultSecret]. Panics if the state is nil.
func (kvs *KeyVaultSecret) StateMust() *keyVaultSecretState {
	if kvs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kvs.Type(), kvs.LocalName()))
	}
	return kvs.state
}

// KeyVaultSecretArgs contains the configurations for azurerm_key_vault_secret.
type KeyVaultSecretArgs struct {
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// ExpirationDate: string, optional
	ExpirationDate terra.StringValue `hcl:"expiration_date,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotBeforeDate: string, optional
	NotBeforeDate terra.StringValue `hcl:"not_before_date,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *keyvaultsecret.Timeouts `hcl:"timeouts,block"`
}
type keyVaultSecretAttributes struct {
	ref terra.Reference
}

// ContentType returns a reference to field content_type of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("content_type"))
}

// ExpirationDate returns a reference to field expiration_date of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("expiration_date"))
}

// Id returns a reference to field id of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("id"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("key_vault_id"))
}

// Name returns a reference to field name of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("name"))
}

// NotBeforeDate returns a reference to field not_before_date of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) NotBeforeDate() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("not_before_date"))
}

// ResourceId returns a reference to field resource_id of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("resource_id"))
}

// ResourceVersionlessId returns a reference to field resource_versionless_id of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) ResourceVersionlessId() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("resource_versionless_id"))
}

// Tags returns a reference to field tags of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kvs.ref.Append("tags"))
}

// Value returns a reference to field value of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("value"))
}

// Version returns a reference to field version of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("version"))
}

// VersionlessId returns a reference to field versionless_id of azurerm_key_vault_secret.
func (kvs keyVaultSecretAttributes) VersionlessId() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("versionless_id"))
}

func (kvs keyVaultSecretAttributes) Timeouts() keyvaultsecret.TimeoutsAttributes {
	return terra.ReferenceAsSingle[keyvaultsecret.TimeoutsAttributes](kvs.ref.Append("timeouts"))
}

type keyVaultSecretState struct {
	ContentType           string                        `json:"content_type"`
	ExpirationDate        string                        `json:"expiration_date"`
	Id                    string                        `json:"id"`
	KeyVaultId            string                        `json:"key_vault_id"`
	Name                  string                        `json:"name"`
	NotBeforeDate         string                        `json:"not_before_date"`
	ResourceId            string                        `json:"resource_id"`
	ResourceVersionlessId string                        `json:"resource_versionless_id"`
	Tags                  map[string]string             `json:"tags"`
	Value                 string                        `json:"value"`
	Version               string                        `json:"version"`
	VersionlessId         string                        `json:"versionless_id"`
	Timeouts              *keyvaultsecret.TimeoutsState `json:"timeouts"`
}
