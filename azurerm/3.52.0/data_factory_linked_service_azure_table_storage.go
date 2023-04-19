// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedserviceazuretablestorage "github.com/golingon/terraproviders/azurerm/3.52.0/datafactorylinkedserviceazuretablestorage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceAzureTableStorage creates a new instance of [DataFactoryLinkedServiceAzureTableStorage].
func NewDataFactoryLinkedServiceAzureTableStorage(name string, args DataFactoryLinkedServiceAzureTableStorageArgs) *DataFactoryLinkedServiceAzureTableStorage {
	return &DataFactoryLinkedServiceAzureTableStorage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceAzureTableStorage)(nil)

// DataFactoryLinkedServiceAzureTableStorage represents the Terraform resource azurerm_data_factory_linked_service_azure_table_storage.
type DataFactoryLinkedServiceAzureTableStorage struct {
	Name      string
	Args      DataFactoryLinkedServiceAzureTableStorageArgs
	state     *dataFactoryLinkedServiceAzureTableStorageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceAzureTableStorage].
func (dflsats *DataFactoryLinkedServiceAzureTableStorage) Type() string {
	return "azurerm_data_factory_linked_service_azure_table_storage"
}

// LocalName returns the local name for [DataFactoryLinkedServiceAzureTableStorage].
func (dflsats *DataFactoryLinkedServiceAzureTableStorage) LocalName() string {
	return dflsats.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceAzureTableStorage].
func (dflsats *DataFactoryLinkedServiceAzureTableStorage) Configuration() interface{} {
	return dflsats.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceAzureTableStorage].
func (dflsats *DataFactoryLinkedServiceAzureTableStorage) DependOn() terra.Reference {
	return terra.ReferenceResource(dflsats)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceAzureTableStorage] depends_on.
func (dflsats *DataFactoryLinkedServiceAzureTableStorage) Dependencies() terra.Dependencies {
	return dflsats.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceAzureTableStorage].
func (dflsats *DataFactoryLinkedServiceAzureTableStorage) LifecycleManagement() *terra.Lifecycle {
	return dflsats.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceAzureTableStorage].
func (dflsats *DataFactoryLinkedServiceAzureTableStorage) Attributes() dataFactoryLinkedServiceAzureTableStorageAttributes {
	return dataFactoryLinkedServiceAzureTableStorageAttributes{ref: terra.ReferenceResource(dflsats)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceAzureTableStorage]'s state.
func (dflsats *DataFactoryLinkedServiceAzureTableStorage) ImportState(av io.Reader) error {
	dflsats.state = &dataFactoryLinkedServiceAzureTableStorageState{}
	if err := json.NewDecoder(av).Decode(dflsats.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflsats.Type(), dflsats.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceAzureTableStorage] has state.
func (dflsats *DataFactoryLinkedServiceAzureTableStorage) State() (*dataFactoryLinkedServiceAzureTableStorageState, bool) {
	return dflsats.state, dflsats.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceAzureTableStorage]. Panics if the state is nil.
func (dflsats *DataFactoryLinkedServiceAzureTableStorage) StateMust() *dataFactoryLinkedServiceAzureTableStorageState {
	if dflsats.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflsats.Type(), dflsats.LocalName()))
	}
	return dflsats.state
}

// DataFactoryLinkedServiceAzureTableStorageArgs contains the configurations for azurerm_data_factory_linked_service_azure_table_storage.
type DataFactoryLinkedServiceAzureTableStorageArgs struct {
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
	// Timeouts: optional
	Timeouts *datafactorylinkedserviceazuretablestorage.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceAzureTableStorageAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_azure_table_storage.
func (dflsats dataFactoryLinkedServiceAzureTableStorageAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsats.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_azure_table_storage.
func (dflsats dataFactoryLinkedServiceAzureTableStorageAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflsats.ref.Append("annotations"))
}

// ConnectionString returns a reference to field connection_string of azurerm_data_factory_linked_service_azure_table_storage.
func (dflsats dataFactoryLinkedServiceAzureTableStorageAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(dflsats.ref.Append("connection_string"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_azure_table_storage.
func (dflsats dataFactoryLinkedServiceAzureTableStorageAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflsats.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_azure_table_storage.
func (dflsats dataFactoryLinkedServiceAzureTableStorageAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflsats.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_azure_table_storage.
func (dflsats dataFactoryLinkedServiceAzureTableStorageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflsats.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_azure_table_storage.
func (dflsats dataFactoryLinkedServiceAzureTableStorageAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflsats.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_azure_table_storage.
func (dflsats dataFactoryLinkedServiceAzureTableStorageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflsats.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_azure_table_storage.
func (dflsats dataFactoryLinkedServiceAzureTableStorageAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsats.ref.Append("parameters"))
}

func (dflsats dataFactoryLinkedServiceAzureTableStorageAttributes) Timeouts() datafactorylinkedserviceazuretablestorage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedserviceazuretablestorage.TimeoutsAttributes](dflsats.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceAzureTableStorageState struct {
	AdditionalProperties   map[string]string                                        `json:"additional_properties"`
	Annotations            []string                                                 `json:"annotations"`
	ConnectionString       string                                                   `json:"connection_string"`
	DataFactoryId          string                                                   `json:"data_factory_id"`
	Description            string                                                   `json:"description"`
	Id                     string                                                   `json:"id"`
	IntegrationRuntimeName string                                                   `json:"integration_runtime_name"`
	Name                   string                                                   `json:"name"`
	Parameters             map[string]string                                        `json:"parameters"`
	Timeouts               *datafactorylinkedserviceazuretablestorage.TimeoutsState `json:"timeouts"`
}
