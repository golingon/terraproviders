// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydatasetcosmosdbsqlapi "github.com/golingon/terraproviders/azurerm/3.69.0/datafactorydatasetcosmosdbsqlapi"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryDatasetCosmosdbSqlapi creates a new instance of [DataFactoryDatasetCosmosdbSqlapi].
func NewDataFactoryDatasetCosmosdbSqlapi(name string, args DataFactoryDatasetCosmosdbSqlapiArgs) *DataFactoryDatasetCosmosdbSqlapi {
	return &DataFactoryDatasetCosmosdbSqlapi{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetCosmosdbSqlapi)(nil)

// DataFactoryDatasetCosmosdbSqlapi represents the Terraform resource azurerm_data_factory_dataset_cosmosdb_sqlapi.
type DataFactoryDatasetCosmosdbSqlapi struct {
	Name      string
	Args      DataFactoryDatasetCosmosdbSqlapiArgs
	state     *dataFactoryDatasetCosmosdbSqlapiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryDatasetCosmosdbSqlapi].
func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) Type() string {
	return "azurerm_data_factory_dataset_cosmosdb_sqlapi"
}

// LocalName returns the local name for [DataFactoryDatasetCosmosdbSqlapi].
func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) LocalName() string {
	return dfdcs.Name
}

// Configuration returns the configuration (args) for [DataFactoryDatasetCosmosdbSqlapi].
func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) Configuration() interface{} {
	return dfdcs.Args
}

// DependOn is used for other resources to depend on [DataFactoryDatasetCosmosdbSqlapi].
func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdcs)
}

// Dependencies returns the list of resources [DataFactoryDatasetCosmosdbSqlapi] depends_on.
func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) Dependencies() terra.Dependencies {
	return dfdcs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryDatasetCosmosdbSqlapi].
func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) LifecycleManagement() *terra.Lifecycle {
	return dfdcs.Lifecycle
}

// Attributes returns the attributes for [DataFactoryDatasetCosmosdbSqlapi].
func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) Attributes() dataFactoryDatasetCosmosdbSqlapiAttributes {
	return dataFactoryDatasetCosmosdbSqlapiAttributes{ref: terra.ReferenceResource(dfdcs)}
}

// ImportState imports the given attribute values into [DataFactoryDatasetCosmosdbSqlapi]'s state.
func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) ImportState(av io.Reader) error {
	dfdcs.state = &dataFactoryDatasetCosmosdbSqlapiState{}
	if err := json.NewDecoder(av).Decode(dfdcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdcs.Type(), dfdcs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryDatasetCosmosdbSqlapi] has state.
func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) State() (*dataFactoryDatasetCosmosdbSqlapiState, bool) {
	return dfdcs.state, dfdcs.state != nil
}

// StateMust returns the state for [DataFactoryDatasetCosmosdbSqlapi]. Panics if the state is nil.
func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) StateMust() *dataFactoryDatasetCosmosdbSqlapiState {
	if dfdcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdcs.Type(), dfdcs.LocalName()))
	}
	return dfdcs.state
}

// DataFactoryDatasetCosmosdbSqlapiArgs contains the configurations for azurerm_data_factory_dataset_cosmosdb_sqlapi.
type DataFactoryDatasetCosmosdbSqlapiArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// CollectionName: string, optional
	CollectionName terra.StringValue `hcl:"collection_name,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Folder: string, optional
	Folder terra.StringValue `hcl:"folder,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkedServiceName: string, required
	LinkedServiceName terra.StringValue `hcl:"linked_service_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// SchemaColumn: min=0
	SchemaColumn []datafactorydatasetcosmosdbsqlapi.SchemaColumn `hcl:"schema_column,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datafactorydatasetcosmosdbsqlapi.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryDatasetCosmosdbSqlapiAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_dataset_cosmosdb_sqlapi.
func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdcs.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_dataset_cosmosdb_sqlapi.
func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfdcs.ref.Append("annotations"))
}

// CollectionName returns a reference to field collection_name of azurerm_data_factory_dataset_cosmosdb_sqlapi.
func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) CollectionName() terra.StringValue {
	return terra.ReferenceAsString(dfdcs.ref.Append("collection_name"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_dataset_cosmosdb_sqlapi.
func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfdcs.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_dataset_cosmosdb_sqlapi.
func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfdcs.ref.Append("description"))
}

// Folder returns a reference to field folder of azurerm_data_factory_dataset_cosmosdb_sqlapi.
func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(dfdcs.ref.Append("folder"))
}

// Id returns a reference to field id of azurerm_data_factory_dataset_cosmosdb_sqlapi.
func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfdcs.ref.Append("id"))
}

// LinkedServiceName returns a reference to field linked_service_name of azurerm_data_factory_dataset_cosmosdb_sqlapi.
func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(dfdcs.ref.Append("linked_service_name"))
}

// Name returns a reference to field name of azurerm_data_factory_dataset_cosmosdb_sqlapi.
func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfdcs.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_dataset_cosmosdb_sqlapi.
func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdcs.ref.Append("parameters"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) SchemaColumn() terra.ListValue[datafactorydatasetcosmosdbsqlapi.SchemaColumnAttributes] {
	return terra.ReferenceAsList[datafactorydatasetcosmosdbsqlapi.SchemaColumnAttributes](dfdcs.ref.Append("schema_column"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Timeouts() datafactorydatasetcosmosdbsqlapi.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorydatasetcosmosdbsqlapi.TimeoutsAttributes](dfdcs.ref.Append("timeouts"))
}

type dataFactoryDatasetCosmosdbSqlapiState struct {
	AdditionalProperties map[string]string                                    `json:"additional_properties"`
	Annotations          []string                                             `json:"annotations"`
	CollectionName       string                                               `json:"collection_name"`
	DataFactoryId        string                                               `json:"data_factory_id"`
	Description          string                                               `json:"description"`
	Folder               string                                               `json:"folder"`
	Id                   string                                               `json:"id"`
	LinkedServiceName    string                                               `json:"linked_service_name"`
	Name                 string                                               `json:"name"`
	Parameters           map[string]string                                    `json:"parameters"`
	SchemaColumn         []datafactorydatasetcosmosdbsqlapi.SchemaColumnState `json:"schema_column"`
	Timeouts             *datafactorydatasetcosmosdbsqlapi.TimeoutsState      `json:"timeouts"`
}
