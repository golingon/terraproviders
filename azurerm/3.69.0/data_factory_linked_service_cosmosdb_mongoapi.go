// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedservicecosmosdbmongoapi "github.com/golingon/terraproviders/azurerm/3.69.0/datafactorylinkedservicecosmosdbmongoapi"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceCosmosdbMongoapi creates a new instance of [DataFactoryLinkedServiceCosmosdbMongoapi].
func NewDataFactoryLinkedServiceCosmosdbMongoapi(name string, args DataFactoryLinkedServiceCosmosdbMongoapiArgs) *DataFactoryLinkedServiceCosmosdbMongoapi {
	return &DataFactoryLinkedServiceCosmosdbMongoapi{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceCosmosdbMongoapi)(nil)

// DataFactoryLinkedServiceCosmosdbMongoapi represents the Terraform resource azurerm_data_factory_linked_service_cosmosdb_mongoapi.
type DataFactoryLinkedServiceCosmosdbMongoapi struct {
	Name      string
	Args      DataFactoryLinkedServiceCosmosdbMongoapiArgs
	state     *dataFactoryLinkedServiceCosmosdbMongoapiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceCosmosdbMongoapi].
func (dflscm *DataFactoryLinkedServiceCosmosdbMongoapi) Type() string {
	return "azurerm_data_factory_linked_service_cosmosdb_mongoapi"
}

// LocalName returns the local name for [DataFactoryLinkedServiceCosmosdbMongoapi].
func (dflscm *DataFactoryLinkedServiceCosmosdbMongoapi) LocalName() string {
	return dflscm.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceCosmosdbMongoapi].
func (dflscm *DataFactoryLinkedServiceCosmosdbMongoapi) Configuration() interface{} {
	return dflscm.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceCosmosdbMongoapi].
func (dflscm *DataFactoryLinkedServiceCosmosdbMongoapi) DependOn() terra.Reference {
	return terra.ReferenceResource(dflscm)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceCosmosdbMongoapi] depends_on.
func (dflscm *DataFactoryLinkedServiceCosmosdbMongoapi) Dependencies() terra.Dependencies {
	return dflscm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceCosmosdbMongoapi].
func (dflscm *DataFactoryLinkedServiceCosmosdbMongoapi) LifecycleManagement() *terra.Lifecycle {
	return dflscm.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceCosmosdbMongoapi].
func (dflscm *DataFactoryLinkedServiceCosmosdbMongoapi) Attributes() dataFactoryLinkedServiceCosmosdbMongoapiAttributes {
	return dataFactoryLinkedServiceCosmosdbMongoapiAttributes{ref: terra.ReferenceResource(dflscm)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceCosmosdbMongoapi]'s state.
func (dflscm *DataFactoryLinkedServiceCosmosdbMongoapi) ImportState(av io.Reader) error {
	dflscm.state = &dataFactoryLinkedServiceCosmosdbMongoapiState{}
	if err := json.NewDecoder(av).Decode(dflscm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflscm.Type(), dflscm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceCosmosdbMongoapi] has state.
func (dflscm *DataFactoryLinkedServiceCosmosdbMongoapi) State() (*dataFactoryLinkedServiceCosmosdbMongoapiState, bool) {
	return dflscm.state, dflscm.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceCosmosdbMongoapi]. Panics if the state is nil.
func (dflscm *DataFactoryLinkedServiceCosmosdbMongoapi) StateMust() *dataFactoryLinkedServiceCosmosdbMongoapiState {
	if dflscm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflscm.Type(), dflscm.LocalName()))
	}
	return dflscm.state
}

// DataFactoryLinkedServiceCosmosdbMongoapiArgs contains the configurations for azurerm_data_factory_linked_service_cosmosdb_mongoapi.
type DataFactoryLinkedServiceCosmosdbMongoapiArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// ConnectionString: string, optional
	ConnectionString terra.StringValue `hcl:"connection_string,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Database: string, optional
	Database terra.StringValue `hcl:"database,attr"`
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
	// ServerVersionIs32OrHigher: bool, optional
	ServerVersionIs32OrHigher terra.BoolValue `hcl:"server_version_is_32_or_higher,attr"`
	// Timeouts: optional
	Timeouts *datafactorylinkedservicecosmosdbmongoapi.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceCosmosdbMongoapiAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_cosmosdb_mongoapi.
func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflscm.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_cosmosdb_mongoapi.
func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflscm.ref.Append("annotations"))
}

// ConnectionString returns a reference to field connection_string of azurerm_data_factory_linked_service_cosmosdb_mongoapi.
func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(dflscm.ref.Append("connection_string"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_cosmosdb_mongoapi.
func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflscm.ref.Append("data_factory_id"))
}

// Database returns a reference to field database of azurerm_data_factory_linked_service_cosmosdb_mongoapi.
func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(dflscm.ref.Append("database"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_cosmosdb_mongoapi.
func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflscm.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_cosmosdb_mongoapi.
func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflscm.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_cosmosdb_mongoapi.
func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflscm.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_cosmosdb_mongoapi.
func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflscm.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_cosmosdb_mongoapi.
func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflscm.ref.Append("parameters"))
}

// ServerVersionIs32OrHigher returns a reference to field server_version_is_32_or_higher of azurerm_data_factory_linked_service_cosmosdb_mongoapi.
func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) ServerVersionIs32OrHigher() terra.BoolValue {
	return terra.ReferenceAsBool(dflscm.ref.Append("server_version_is_32_or_higher"))
}

func (dflscm dataFactoryLinkedServiceCosmosdbMongoapiAttributes) Timeouts() datafactorylinkedservicecosmosdbmongoapi.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedservicecosmosdbmongoapi.TimeoutsAttributes](dflscm.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceCosmosdbMongoapiState struct {
	AdditionalProperties      map[string]string                                       `json:"additional_properties"`
	Annotations               []string                                                `json:"annotations"`
	ConnectionString          string                                                  `json:"connection_string"`
	DataFactoryId             string                                                  `json:"data_factory_id"`
	Database                  string                                                  `json:"database"`
	Description               string                                                  `json:"description"`
	Id                        string                                                  `json:"id"`
	IntegrationRuntimeName    string                                                  `json:"integration_runtime_name"`
	Name                      string                                                  `json:"name"`
	Parameters                map[string]string                                       `json:"parameters"`
	ServerVersionIs32OrHigher bool                                                    `json:"server_version_is_32_or_higher"`
	Timeouts                  *datafactorylinkedservicecosmosdbmongoapi.TimeoutsState `json:"timeouts"`
}
