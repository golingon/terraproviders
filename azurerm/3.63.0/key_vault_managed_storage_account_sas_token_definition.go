// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	keyvaultmanagedstorageaccountsastokendefinition "github.com/golingon/terraproviders/azurerm/3.63.0/keyvaultmanagedstorageaccountsastokendefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKeyVaultManagedStorageAccountSasTokenDefinition creates a new instance of [KeyVaultManagedStorageAccountSasTokenDefinition].
func NewKeyVaultManagedStorageAccountSasTokenDefinition(name string, args KeyVaultManagedStorageAccountSasTokenDefinitionArgs) *KeyVaultManagedStorageAccountSasTokenDefinition {
	return &KeyVaultManagedStorageAccountSasTokenDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyVaultManagedStorageAccountSasTokenDefinition)(nil)

// KeyVaultManagedStorageAccountSasTokenDefinition represents the Terraform resource azurerm_key_vault_managed_storage_account_sas_token_definition.
type KeyVaultManagedStorageAccountSasTokenDefinition struct {
	Name      string
	Args      KeyVaultManagedStorageAccountSasTokenDefinitionArgs
	state     *keyVaultManagedStorageAccountSasTokenDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KeyVaultManagedStorageAccountSasTokenDefinition].
func (kvmsastd *KeyVaultManagedStorageAccountSasTokenDefinition) Type() string {
	return "azurerm_key_vault_managed_storage_account_sas_token_definition"
}

// LocalName returns the local name for [KeyVaultManagedStorageAccountSasTokenDefinition].
func (kvmsastd *KeyVaultManagedStorageAccountSasTokenDefinition) LocalName() string {
	return kvmsastd.Name
}

// Configuration returns the configuration (args) for [KeyVaultManagedStorageAccountSasTokenDefinition].
func (kvmsastd *KeyVaultManagedStorageAccountSasTokenDefinition) Configuration() interface{} {
	return kvmsastd.Args
}

// DependOn is used for other resources to depend on [KeyVaultManagedStorageAccountSasTokenDefinition].
func (kvmsastd *KeyVaultManagedStorageAccountSasTokenDefinition) DependOn() terra.Reference {
	return terra.ReferenceResource(kvmsastd)
}

// Dependencies returns the list of resources [KeyVaultManagedStorageAccountSasTokenDefinition] depends_on.
func (kvmsastd *KeyVaultManagedStorageAccountSasTokenDefinition) Dependencies() terra.Dependencies {
	return kvmsastd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KeyVaultManagedStorageAccountSasTokenDefinition].
func (kvmsastd *KeyVaultManagedStorageAccountSasTokenDefinition) LifecycleManagement() *terra.Lifecycle {
	return kvmsastd.Lifecycle
}

// Attributes returns the attributes for [KeyVaultManagedStorageAccountSasTokenDefinition].
func (kvmsastd *KeyVaultManagedStorageAccountSasTokenDefinition) Attributes() keyVaultManagedStorageAccountSasTokenDefinitionAttributes {
	return keyVaultManagedStorageAccountSasTokenDefinitionAttributes{ref: terra.ReferenceResource(kvmsastd)}
}

// ImportState imports the given attribute values into [KeyVaultManagedStorageAccountSasTokenDefinition]'s state.
func (kvmsastd *KeyVaultManagedStorageAccountSasTokenDefinition) ImportState(av io.Reader) error {
	kvmsastd.state = &keyVaultManagedStorageAccountSasTokenDefinitionState{}
	if err := json.NewDecoder(av).Decode(kvmsastd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kvmsastd.Type(), kvmsastd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KeyVaultManagedStorageAccountSasTokenDefinition] has state.
func (kvmsastd *KeyVaultManagedStorageAccountSasTokenDefinition) State() (*keyVaultManagedStorageAccountSasTokenDefinitionState, bool) {
	return kvmsastd.state, kvmsastd.state != nil
}

// StateMust returns the state for [KeyVaultManagedStorageAccountSasTokenDefinition]. Panics if the state is nil.
func (kvmsastd *KeyVaultManagedStorageAccountSasTokenDefinition) StateMust() *keyVaultManagedStorageAccountSasTokenDefinitionState {
	if kvmsastd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kvmsastd.Type(), kvmsastd.LocalName()))
	}
	return kvmsastd.state
}

// KeyVaultManagedStorageAccountSasTokenDefinitionArgs contains the configurations for azurerm_key_vault_managed_storage_account_sas_token_definition.
type KeyVaultManagedStorageAccountSasTokenDefinitionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagedStorageAccountId: string, required
	ManagedStorageAccountId terra.StringValue `hcl:"managed_storage_account_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SasTemplateUri: string, required
	SasTemplateUri terra.StringValue `hcl:"sas_template_uri,attr" validate:"required"`
	// SasType: string, required
	SasType terra.StringValue `hcl:"sas_type,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ValidityPeriod: string, required
	ValidityPeriod terra.StringValue `hcl:"validity_period,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *keyvaultmanagedstorageaccountsastokendefinition.Timeouts `hcl:"timeouts,block"`
}
type keyVaultManagedStorageAccountSasTokenDefinitionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (kvmsastd keyVaultManagedStorageAccountSasTokenDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvmsastd.ref.Append("id"))
}

// ManagedStorageAccountId returns a reference to field managed_storage_account_id of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (kvmsastd keyVaultManagedStorageAccountSasTokenDefinitionAttributes) ManagedStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(kvmsastd.ref.Append("managed_storage_account_id"))
}

// Name returns a reference to field name of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (kvmsastd keyVaultManagedStorageAccountSasTokenDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kvmsastd.ref.Append("name"))
}

// SasTemplateUri returns a reference to field sas_template_uri of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (kvmsastd keyVaultManagedStorageAccountSasTokenDefinitionAttributes) SasTemplateUri() terra.StringValue {
	return terra.ReferenceAsString(kvmsastd.ref.Append("sas_template_uri"))
}

// SasType returns a reference to field sas_type of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (kvmsastd keyVaultManagedStorageAccountSasTokenDefinitionAttributes) SasType() terra.StringValue {
	return terra.ReferenceAsString(kvmsastd.ref.Append("sas_type"))
}

// SecretId returns a reference to field secret_id of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (kvmsastd keyVaultManagedStorageAccountSasTokenDefinitionAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(kvmsastd.ref.Append("secret_id"))
}

// Tags returns a reference to field tags of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (kvmsastd keyVaultManagedStorageAccountSasTokenDefinitionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kvmsastd.ref.Append("tags"))
}

// ValidityPeriod returns a reference to field validity_period of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (kvmsastd keyVaultManagedStorageAccountSasTokenDefinitionAttributes) ValidityPeriod() terra.StringValue {
	return terra.ReferenceAsString(kvmsastd.ref.Append("validity_period"))
}

func (kvmsastd keyVaultManagedStorageAccountSasTokenDefinitionAttributes) Timeouts() keyvaultmanagedstorageaccountsastokendefinition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[keyvaultmanagedstorageaccountsastokendefinition.TimeoutsAttributes](kvmsastd.ref.Append("timeouts"))
}

type keyVaultManagedStorageAccountSasTokenDefinitionState struct {
	Id                      string                                                         `json:"id"`
	ManagedStorageAccountId string                                                         `json:"managed_storage_account_id"`
	Name                    string                                                         `json:"name"`
	SasTemplateUri          string                                                         `json:"sas_template_uri"`
	SasType                 string                                                         `json:"sas_type"`
	SecretId                string                                                         `json:"secret_id"`
	Tags                    map[string]string                                              `json:"tags"`
	ValidityPeriod          string                                                         `json:"validity_period"`
	Timeouts                *keyvaultmanagedstorageaccountsastokendefinition.TimeoutsState `json:"timeouts"`
}
