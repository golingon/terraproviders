// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedservicekeyvault "github.com/golingon/terraproviders/azurerm/3.52.0/datafactorylinkedservicekeyvault"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceKeyVault creates a new instance of [DataFactoryLinkedServiceKeyVault].
func NewDataFactoryLinkedServiceKeyVault(name string, args DataFactoryLinkedServiceKeyVaultArgs) *DataFactoryLinkedServiceKeyVault {
	return &DataFactoryLinkedServiceKeyVault{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceKeyVault)(nil)

// DataFactoryLinkedServiceKeyVault represents the Terraform resource azurerm_data_factory_linked_service_key_vault.
type DataFactoryLinkedServiceKeyVault struct {
	Name      string
	Args      DataFactoryLinkedServiceKeyVaultArgs
	state     *dataFactoryLinkedServiceKeyVaultState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceKeyVault].
func (dflskv *DataFactoryLinkedServiceKeyVault) Type() string {
	return "azurerm_data_factory_linked_service_key_vault"
}

// LocalName returns the local name for [DataFactoryLinkedServiceKeyVault].
func (dflskv *DataFactoryLinkedServiceKeyVault) LocalName() string {
	return dflskv.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceKeyVault].
func (dflskv *DataFactoryLinkedServiceKeyVault) Configuration() interface{} {
	return dflskv.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceKeyVault].
func (dflskv *DataFactoryLinkedServiceKeyVault) DependOn() terra.Reference {
	return terra.ReferenceResource(dflskv)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceKeyVault] depends_on.
func (dflskv *DataFactoryLinkedServiceKeyVault) Dependencies() terra.Dependencies {
	return dflskv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceKeyVault].
func (dflskv *DataFactoryLinkedServiceKeyVault) LifecycleManagement() *terra.Lifecycle {
	return dflskv.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceKeyVault].
func (dflskv *DataFactoryLinkedServiceKeyVault) Attributes() dataFactoryLinkedServiceKeyVaultAttributes {
	return dataFactoryLinkedServiceKeyVaultAttributes{ref: terra.ReferenceResource(dflskv)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceKeyVault]'s state.
func (dflskv *DataFactoryLinkedServiceKeyVault) ImportState(av io.Reader) error {
	dflskv.state = &dataFactoryLinkedServiceKeyVaultState{}
	if err := json.NewDecoder(av).Decode(dflskv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflskv.Type(), dflskv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceKeyVault] has state.
func (dflskv *DataFactoryLinkedServiceKeyVault) State() (*dataFactoryLinkedServiceKeyVaultState, bool) {
	return dflskv.state, dflskv.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceKeyVault]. Panics if the state is nil.
func (dflskv *DataFactoryLinkedServiceKeyVault) StateMust() *dataFactoryLinkedServiceKeyVaultState {
	if dflskv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflskv.Type(), dflskv.LocalName()))
	}
	return dflskv.state
}

// DataFactoryLinkedServiceKeyVaultArgs contains the configurations for azurerm_data_factory_linked_service_key_vault.
type DataFactoryLinkedServiceKeyVaultArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationRuntimeName: string, optional
	IntegrationRuntimeName terra.StringValue `hcl:"integration_runtime_name,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// Timeouts: optional
	Timeouts *datafactorylinkedservicekeyvault.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceKeyVaultAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_key_vault.
func (dflskv dataFactoryLinkedServiceKeyVaultAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflskv.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_key_vault.
func (dflskv dataFactoryLinkedServiceKeyVaultAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflskv.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_key_vault.
func (dflskv dataFactoryLinkedServiceKeyVaultAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflskv.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_key_vault.
func (dflskv dataFactoryLinkedServiceKeyVaultAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflskv.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_key_vault.
func (dflskv dataFactoryLinkedServiceKeyVaultAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflskv.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_key_vault.
func (dflskv dataFactoryLinkedServiceKeyVaultAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflskv.ref.Append("integration_runtime_name"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_data_factory_linked_service_key_vault.
func (dflskv dataFactoryLinkedServiceKeyVaultAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(dflskv.ref.Append("key_vault_id"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_key_vault.
func (dflskv dataFactoryLinkedServiceKeyVaultAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflskv.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_key_vault.
func (dflskv dataFactoryLinkedServiceKeyVaultAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflskv.ref.Append("parameters"))
}

func (dflskv dataFactoryLinkedServiceKeyVaultAttributes) Timeouts() datafactorylinkedservicekeyvault.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedservicekeyvault.TimeoutsAttributes](dflskv.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceKeyVaultState struct {
	AdditionalProperties   map[string]string                               `json:"additional_properties"`
	Annotations            []string                                        `json:"annotations"`
	DataFactoryId          string                                          `json:"data_factory_id"`
	Description            string                                          `json:"description"`
	Id                     string                                          `json:"id"`
	IntegrationRuntimeName string                                          `json:"integration_runtime_name"`
	KeyVaultId             string                                          `json:"key_vault_id"`
	Name                   string                                          `json:"name"`
	Parameters             map[string]string                               `json:"parameters"`
	Timeouts               *datafactorylinkedservicekeyvault.TimeoutsState `json:"timeouts"`
}
