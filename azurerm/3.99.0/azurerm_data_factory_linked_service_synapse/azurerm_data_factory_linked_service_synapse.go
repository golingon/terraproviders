// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_factory_linked_service_synapse

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

// Resource represents the Terraform resource azurerm_data_factory_linked_service_synapse.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermDataFactoryLinkedServiceSynapseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adflss *Resource) Type() string {
	return "azurerm_data_factory_linked_service_synapse"
}

// LocalName returns the local name for [Resource].
func (adflss *Resource) LocalName() string {
	return adflss.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adflss *Resource) Configuration() interface{} {
	return adflss.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adflss *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adflss)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adflss *Resource) Dependencies() terra.Dependencies {
	return adflss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adflss *Resource) LifecycleManagement() *terra.Lifecycle {
	return adflss.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adflss *Resource) Attributes() azurermDataFactoryLinkedServiceSynapseAttributes {
	return azurermDataFactoryLinkedServiceSynapseAttributes{ref: terra.ReferenceResource(adflss)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adflss *Resource) ImportState(state io.Reader) error {
	adflss.state = &azurermDataFactoryLinkedServiceSynapseState{}
	if err := json.NewDecoder(state).Decode(adflss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adflss.Type(), adflss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adflss *Resource) State() (*azurermDataFactoryLinkedServiceSynapseState, bool) {
	return adflss.state, adflss.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adflss *Resource) StateMust() *azurermDataFactoryLinkedServiceSynapseState {
	if adflss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adflss.Type(), adflss.LocalName()))
	}
	return adflss.state
}

// Args contains the configurations for azurerm_data_factory_linked_service_synapse.
type Args struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// ConnectionString: string, required
	ConnectionString terra.StringValue `hcl:"connection_string,attr" validate:"required"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationRuntimeName: string, optional
	IntegrationRuntimeName terra.StringValue `hcl:"integration_runtime_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// KeyVaultPassword: optional
	KeyVaultPassword *KeyVaultPassword `hcl:"key_vault_password,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermDataFactoryLinkedServiceSynapseAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_synapse.
func (adflss azurermDataFactoryLinkedServiceSynapseAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adflss.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_synapse.
func (adflss azurermDataFactoryLinkedServiceSynapseAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](adflss.ref.Append("annotations"))
}

// ConnectionString returns a reference to field connection_string of azurerm_data_factory_linked_service_synapse.
func (adflss azurermDataFactoryLinkedServiceSynapseAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(adflss.ref.Append("connection_string"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_synapse.
func (adflss azurermDataFactoryLinkedServiceSynapseAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(adflss.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_synapse.
func (adflss azurermDataFactoryLinkedServiceSynapseAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(adflss.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_synapse.
func (adflss azurermDataFactoryLinkedServiceSynapseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adflss.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_synapse.
func (adflss azurermDataFactoryLinkedServiceSynapseAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(adflss.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_synapse.
func (adflss azurermDataFactoryLinkedServiceSynapseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adflss.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_synapse.
func (adflss azurermDataFactoryLinkedServiceSynapseAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adflss.ref.Append("parameters"))
}

func (adflss azurermDataFactoryLinkedServiceSynapseAttributes) KeyVaultPassword() terra.ListValue[KeyVaultPasswordAttributes] {
	return terra.ReferenceAsList[KeyVaultPasswordAttributes](adflss.ref.Append("key_vault_password"))
}

func (adflss azurermDataFactoryLinkedServiceSynapseAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adflss.ref.Append("timeouts"))
}

type azurermDataFactoryLinkedServiceSynapseState struct {
	AdditionalProperties   map[string]string       `json:"additional_properties"`
	Annotations            []string                `json:"annotations"`
	ConnectionString       string                  `json:"connection_string"`
	DataFactoryId          string                  `json:"data_factory_id"`
	Description            string                  `json:"description"`
	Id                     string                  `json:"id"`
	IntegrationRuntimeName string                  `json:"integration_runtime_name"`
	Name                   string                  `json:"name"`
	Parameters             map[string]string       `json:"parameters"`
	KeyVaultPassword       []KeyVaultPasswordState `json:"key_vault_password"`
	Timeouts               *TimeoutsState          `json:"timeouts"`
}
