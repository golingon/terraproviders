// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_key_vault_managed_storage_account_sas_token_definition

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

// Resource represents the Terraform resource azurerm_key_vault_managed_storage_account_sas_token_definition.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermKeyVaultManagedStorageAccountSasTokenDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (akvmsastd *Resource) Type() string {
	return "azurerm_key_vault_managed_storage_account_sas_token_definition"
}

// LocalName returns the local name for [Resource].
func (akvmsastd *Resource) LocalName() string {
	return akvmsastd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (akvmsastd *Resource) Configuration() interface{} {
	return akvmsastd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (akvmsastd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(akvmsastd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (akvmsastd *Resource) Dependencies() terra.Dependencies {
	return akvmsastd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (akvmsastd *Resource) LifecycleManagement() *terra.Lifecycle {
	return akvmsastd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (akvmsastd *Resource) Attributes() azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes {
	return azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes{ref: terra.ReferenceResource(akvmsastd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (akvmsastd *Resource) ImportState(state io.Reader) error {
	akvmsastd.state = &azurermKeyVaultManagedStorageAccountSasTokenDefinitionState{}
	if err := json.NewDecoder(state).Decode(akvmsastd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akvmsastd.Type(), akvmsastd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (akvmsastd *Resource) State() (*azurermKeyVaultManagedStorageAccountSasTokenDefinitionState, bool) {
	return akvmsastd.state, akvmsastd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (akvmsastd *Resource) StateMust() *azurermKeyVaultManagedStorageAccountSasTokenDefinitionState {
	if akvmsastd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akvmsastd.Type(), akvmsastd.LocalName()))
	}
	return akvmsastd.state
}

// Args contains the configurations for azurerm_key_vault_managed_storage_account_sas_token_definition.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (akvmsastd azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akvmsastd.ref.Append("id"))
}

// ManagedStorageAccountId returns a reference to field managed_storage_account_id of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (akvmsastd azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes) ManagedStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(akvmsastd.ref.Append("managed_storage_account_id"))
}

// Name returns a reference to field name of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (akvmsastd azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akvmsastd.ref.Append("name"))
}

// SasTemplateUri returns a reference to field sas_template_uri of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (akvmsastd azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes) SasTemplateUri() terra.StringValue {
	return terra.ReferenceAsString(akvmsastd.ref.Append("sas_template_uri"))
}

// SasType returns a reference to field sas_type of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (akvmsastd azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes) SasType() terra.StringValue {
	return terra.ReferenceAsString(akvmsastd.ref.Append("sas_type"))
}

// SecretId returns a reference to field secret_id of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (akvmsastd azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(akvmsastd.ref.Append("secret_id"))
}

// Tags returns a reference to field tags of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (akvmsastd azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akvmsastd.ref.Append("tags"))
}

// ValidityPeriod returns a reference to field validity_period of azurerm_key_vault_managed_storage_account_sas_token_definition.
func (akvmsastd azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes) ValidityPeriod() terra.StringValue {
	return terra.ReferenceAsString(akvmsastd.ref.Append("validity_period"))
}

func (akvmsastd azurermKeyVaultManagedStorageAccountSasTokenDefinitionAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](akvmsastd.ref.Append("timeouts"))
}

type azurermKeyVaultManagedStorageAccountSasTokenDefinitionState struct {
	Id                      string            `json:"id"`
	ManagedStorageAccountId string            `json:"managed_storage_account_id"`
	Name                    string            `json:"name"`
	SasTemplateUri          string            `json:"sas_template_uri"`
	SasType                 string            `json:"sas_type"`
	SecretId                string            `json:"secret_id"`
	Tags                    map[string]string `json:"tags"`
	ValidityPeriod          string            `json:"validity_period"`
	Timeouts                *TimeoutsState    `json:"timeouts"`
}
