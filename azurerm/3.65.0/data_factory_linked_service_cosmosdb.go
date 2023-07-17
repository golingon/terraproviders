// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedservicecosmosdb "github.com/golingon/terraproviders/azurerm/3.65.0/datafactorylinkedservicecosmosdb"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceCosmosdb creates a new instance of [DataFactoryLinkedServiceCosmosdb].
func NewDataFactoryLinkedServiceCosmosdb(name string, args DataFactoryLinkedServiceCosmosdbArgs) *DataFactoryLinkedServiceCosmosdb {
	return &DataFactoryLinkedServiceCosmosdb{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceCosmosdb)(nil)

// DataFactoryLinkedServiceCosmosdb represents the Terraform resource azurerm_data_factory_linked_service_cosmosdb.
type DataFactoryLinkedServiceCosmosdb struct {
	Name      string
	Args      DataFactoryLinkedServiceCosmosdbArgs
	state     *dataFactoryLinkedServiceCosmosdbState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceCosmosdb].
func (dflsc *DataFactoryLinkedServiceCosmosdb) Type() string {
	return "azurerm_data_factory_linked_service_cosmosdb"
}

// LocalName returns the local name for [DataFactoryLinkedServiceCosmosdb].
func (dflsc *DataFactoryLinkedServiceCosmosdb) LocalName() string {
	return dflsc.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceCosmosdb].
func (dflsc *DataFactoryLinkedServiceCosmosdb) Configuration() interface{} {
	return dflsc.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceCosmosdb].
func (dflsc *DataFactoryLinkedServiceCosmosdb) DependOn() terra.Reference {
	return terra.ReferenceResource(dflsc)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceCosmosdb] depends_on.
func (dflsc *DataFactoryLinkedServiceCosmosdb) Dependencies() terra.Dependencies {
	return dflsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceCosmosdb].
func (dflsc *DataFactoryLinkedServiceCosmosdb) LifecycleManagement() *terra.Lifecycle {
	return dflsc.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceCosmosdb].
func (dflsc *DataFactoryLinkedServiceCosmosdb) Attributes() dataFactoryLinkedServiceCosmosdbAttributes {
	return dataFactoryLinkedServiceCosmosdbAttributes{ref: terra.ReferenceResource(dflsc)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceCosmosdb]'s state.
func (dflsc *DataFactoryLinkedServiceCosmosdb) ImportState(av io.Reader) error {
	dflsc.state = &dataFactoryLinkedServiceCosmosdbState{}
	if err := json.NewDecoder(av).Decode(dflsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflsc.Type(), dflsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceCosmosdb] has state.
func (dflsc *DataFactoryLinkedServiceCosmosdb) State() (*dataFactoryLinkedServiceCosmosdbState, bool) {
	return dflsc.state, dflsc.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceCosmosdb]. Panics if the state is nil.
func (dflsc *DataFactoryLinkedServiceCosmosdb) StateMust() *dataFactoryLinkedServiceCosmosdbState {
	if dflsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflsc.Type(), dflsc.LocalName()))
	}
	return dflsc.state
}

// DataFactoryLinkedServiceCosmosdbArgs contains the configurations for azurerm_data_factory_linked_service_cosmosdb.
type DataFactoryLinkedServiceCosmosdbArgs struct {
	// AccountEndpoint: string, optional
	AccountEndpoint terra.StringValue `hcl:"account_endpoint,attr"`
	// AccountKey: string, optional
	AccountKey terra.StringValue `hcl:"account_key,attr"`
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
	// Timeouts: optional
	Timeouts *datafactorylinkedservicecosmosdb.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceCosmosdbAttributes struct {
	ref terra.Reference
}

// AccountEndpoint returns a reference to field account_endpoint of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) AccountEndpoint() terra.StringValue {
	return terra.ReferenceAsString(dflsc.ref.Append("account_endpoint"))
}

// AccountKey returns a reference to field account_key of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) AccountKey() terra.StringValue {
	return terra.ReferenceAsString(dflsc.ref.Append("account_key"))
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsc.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflsc.ref.Append("annotations"))
}

// ConnectionString returns a reference to field connection_string of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(dflsc.ref.Append("connection_string"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflsc.ref.Append("data_factory_id"))
}

// Database returns a reference to field database of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(dflsc.ref.Append("database"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflsc.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflsc.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflsc.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflsc.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_cosmosdb.
func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsc.ref.Append("parameters"))
}

func (dflsc dataFactoryLinkedServiceCosmosdbAttributes) Timeouts() datafactorylinkedservicecosmosdb.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedservicecosmosdb.TimeoutsAttributes](dflsc.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceCosmosdbState struct {
	AccountEndpoint        string                                          `json:"account_endpoint"`
	AccountKey             string                                          `json:"account_key"`
	AdditionalProperties   map[string]string                               `json:"additional_properties"`
	Annotations            []string                                        `json:"annotations"`
	ConnectionString       string                                          `json:"connection_string"`
	DataFactoryId          string                                          `json:"data_factory_id"`
	Database               string                                          `json:"database"`
	Description            string                                          `json:"description"`
	Id                     string                                          `json:"id"`
	IntegrationRuntimeName string                                          `json:"integration_runtime_name"`
	Name                   string                                          `json:"name"`
	Parameters             map[string]string                               `json:"parameters"`
	Timeouts               *datafactorylinkedservicecosmosdb.TimeoutsState `json:"timeouts"`
}
