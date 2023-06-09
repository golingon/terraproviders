// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedservicesynapse "github.com/golingon/terraproviders/azurerm/3.49.0/datafactorylinkedservicesynapse"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceSynapse creates a new instance of [DataFactoryLinkedServiceSynapse].
func NewDataFactoryLinkedServiceSynapse(name string, args DataFactoryLinkedServiceSynapseArgs) *DataFactoryLinkedServiceSynapse {
	return &DataFactoryLinkedServiceSynapse{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceSynapse)(nil)

// DataFactoryLinkedServiceSynapse represents the Terraform resource azurerm_data_factory_linked_service_synapse.
type DataFactoryLinkedServiceSynapse struct {
	Name      string
	Args      DataFactoryLinkedServiceSynapseArgs
	state     *dataFactoryLinkedServiceSynapseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceSynapse].
func (dflss *DataFactoryLinkedServiceSynapse) Type() string {
	return "azurerm_data_factory_linked_service_synapse"
}

// LocalName returns the local name for [DataFactoryLinkedServiceSynapse].
func (dflss *DataFactoryLinkedServiceSynapse) LocalName() string {
	return dflss.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceSynapse].
func (dflss *DataFactoryLinkedServiceSynapse) Configuration() interface{} {
	return dflss.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceSynapse].
func (dflss *DataFactoryLinkedServiceSynapse) DependOn() terra.Reference {
	return terra.ReferenceResource(dflss)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceSynapse] depends_on.
func (dflss *DataFactoryLinkedServiceSynapse) Dependencies() terra.Dependencies {
	return dflss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceSynapse].
func (dflss *DataFactoryLinkedServiceSynapse) LifecycleManagement() *terra.Lifecycle {
	return dflss.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceSynapse].
func (dflss *DataFactoryLinkedServiceSynapse) Attributes() dataFactoryLinkedServiceSynapseAttributes {
	return dataFactoryLinkedServiceSynapseAttributes{ref: terra.ReferenceResource(dflss)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceSynapse]'s state.
func (dflss *DataFactoryLinkedServiceSynapse) ImportState(av io.Reader) error {
	dflss.state = &dataFactoryLinkedServiceSynapseState{}
	if err := json.NewDecoder(av).Decode(dflss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflss.Type(), dflss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceSynapse] has state.
func (dflss *DataFactoryLinkedServiceSynapse) State() (*dataFactoryLinkedServiceSynapseState, bool) {
	return dflss.state, dflss.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceSynapse]. Panics if the state is nil.
func (dflss *DataFactoryLinkedServiceSynapse) StateMust() *dataFactoryLinkedServiceSynapseState {
	if dflss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflss.Type(), dflss.LocalName()))
	}
	return dflss.state
}

// DataFactoryLinkedServiceSynapseArgs contains the configurations for azurerm_data_factory_linked_service_synapse.
type DataFactoryLinkedServiceSynapseArgs struct {
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
	KeyVaultPassword *datafactorylinkedservicesynapse.KeyVaultPassword `hcl:"key_vault_password,block"`
	// Timeouts: optional
	Timeouts *datafactorylinkedservicesynapse.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceSynapseAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_synapse.
func (dflss dataFactoryLinkedServiceSynapseAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflss.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_synapse.
func (dflss dataFactoryLinkedServiceSynapseAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflss.ref.Append("annotations"))
}

// ConnectionString returns a reference to field connection_string of azurerm_data_factory_linked_service_synapse.
func (dflss dataFactoryLinkedServiceSynapseAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("connection_string"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_synapse.
func (dflss dataFactoryLinkedServiceSynapseAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_synapse.
func (dflss dataFactoryLinkedServiceSynapseAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_synapse.
func (dflss dataFactoryLinkedServiceSynapseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_synapse.
func (dflss dataFactoryLinkedServiceSynapseAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_synapse.
func (dflss dataFactoryLinkedServiceSynapseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_synapse.
func (dflss dataFactoryLinkedServiceSynapseAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflss.ref.Append("parameters"))
}

func (dflss dataFactoryLinkedServiceSynapseAttributes) KeyVaultPassword() terra.ListValue[datafactorylinkedservicesynapse.KeyVaultPasswordAttributes] {
	return terra.ReferenceAsList[datafactorylinkedservicesynapse.KeyVaultPasswordAttributes](dflss.ref.Append("key_vault_password"))
}

func (dflss dataFactoryLinkedServiceSynapseAttributes) Timeouts() datafactorylinkedservicesynapse.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedservicesynapse.TimeoutsAttributes](dflss.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceSynapseState struct {
	AdditionalProperties   map[string]string                                       `json:"additional_properties"`
	Annotations            []string                                                `json:"annotations"`
	ConnectionString       string                                                  `json:"connection_string"`
	DataFactoryId          string                                                  `json:"data_factory_id"`
	Description            string                                                  `json:"description"`
	Id                     string                                                  `json:"id"`
	IntegrationRuntimeName string                                                  `json:"integration_runtime_name"`
	Name                   string                                                  `json:"name"`
	Parameters             map[string]string                                       `json:"parameters"`
	KeyVaultPassword       []datafactorylinkedservicesynapse.KeyVaultPasswordState `json:"key_vault_password"`
	Timeouts               *datafactorylinkedservicesynapse.TimeoutsState          `json:"timeouts"`
}
