// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydatasetsqlservertable "github.com/golingon/terraproviders/azurerm/3.64.0/datafactorydatasetsqlservertable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryDatasetSqlServerTable creates a new instance of [DataFactoryDatasetSqlServerTable].
func NewDataFactoryDatasetSqlServerTable(name string, args DataFactoryDatasetSqlServerTableArgs) *DataFactoryDatasetSqlServerTable {
	return &DataFactoryDatasetSqlServerTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetSqlServerTable)(nil)

// DataFactoryDatasetSqlServerTable represents the Terraform resource azurerm_data_factory_dataset_sql_server_table.
type DataFactoryDatasetSqlServerTable struct {
	Name      string
	Args      DataFactoryDatasetSqlServerTableArgs
	state     *dataFactoryDatasetSqlServerTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryDatasetSqlServerTable].
func (dfdsst *DataFactoryDatasetSqlServerTable) Type() string {
	return "azurerm_data_factory_dataset_sql_server_table"
}

// LocalName returns the local name for [DataFactoryDatasetSqlServerTable].
func (dfdsst *DataFactoryDatasetSqlServerTable) LocalName() string {
	return dfdsst.Name
}

// Configuration returns the configuration (args) for [DataFactoryDatasetSqlServerTable].
func (dfdsst *DataFactoryDatasetSqlServerTable) Configuration() interface{} {
	return dfdsst.Args
}

// DependOn is used for other resources to depend on [DataFactoryDatasetSqlServerTable].
func (dfdsst *DataFactoryDatasetSqlServerTable) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdsst)
}

// Dependencies returns the list of resources [DataFactoryDatasetSqlServerTable] depends_on.
func (dfdsst *DataFactoryDatasetSqlServerTable) Dependencies() terra.Dependencies {
	return dfdsst.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryDatasetSqlServerTable].
func (dfdsst *DataFactoryDatasetSqlServerTable) LifecycleManagement() *terra.Lifecycle {
	return dfdsst.Lifecycle
}

// Attributes returns the attributes for [DataFactoryDatasetSqlServerTable].
func (dfdsst *DataFactoryDatasetSqlServerTable) Attributes() dataFactoryDatasetSqlServerTableAttributes {
	return dataFactoryDatasetSqlServerTableAttributes{ref: terra.ReferenceResource(dfdsst)}
}

// ImportState imports the given attribute values into [DataFactoryDatasetSqlServerTable]'s state.
func (dfdsst *DataFactoryDatasetSqlServerTable) ImportState(av io.Reader) error {
	dfdsst.state = &dataFactoryDatasetSqlServerTableState{}
	if err := json.NewDecoder(av).Decode(dfdsst.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdsst.Type(), dfdsst.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryDatasetSqlServerTable] has state.
func (dfdsst *DataFactoryDatasetSqlServerTable) State() (*dataFactoryDatasetSqlServerTableState, bool) {
	return dfdsst.state, dfdsst.state != nil
}

// StateMust returns the state for [DataFactoryDatasetSqlServerTable]. Panics if the state is nil.
func (dfdsst *DataFactoryDatasetSqlServerTable) StateMust() *dataFactoryDatasetSqlServerTableState {
	if dfdsst.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdsst.Type(), dfdsst.LocalName()))
	}
	return dfdsst.state
}

// DataFactoryDatasetSqlServerTableArgs contains the configurations for azurerm_data_factory_dataset_sql_server_table.
type DataFactoryDatasetSqlServerTableArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
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
	// TableName: string, optional
	TableName terra.StringValue `hcl:"table_name,attr"`
	// SchemaColumn: min=0
	SchemaColumn []datafactorydatasetsqlservertable.SchemaColumn `hcl:"schema_column,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datafactorydatasetsqlservertable.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryDatasetSqlServerTableAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_dataset_sql_server_table.
func (dfdsst dataFactoryDatasetSqlServerTableAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdsst.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_dataset_sql_server_table.
func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfdsst.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_dataset_sql_server_table.
func (dfdsst dataFactoryDatasetSqlServerTableAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfdsst.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_dataset_sql_server_table.
func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfdsst.ref.Append("description"))
}

// Folder returns a reference to field folder of azurerm_data_factory_dataset_sql_server_table.
func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(dfdsst.ref.Append("folder"))
}

// Id returns a reference to field id of azurerm_data_factory_dataset_sql_server_table.
func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfdsst.ref.Append("id"))
}

// LinkedServiceName returns a reference to field linked_service_name of azurerm_data_factory_dataset_sql_server_table.
func (dfdsst dataFactoryDatasetSqlServerTableAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(dfdsst.ref.Append("linked_service_name"))
}

// Name returns a reference to field name of azurerm_data_factory_dataset_sql_server_table.
func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfdsst.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_dataset_sql_server_table.
func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdsst.ref.Append("parameters"))
}

// TableName returns a reference to field table_name of azurerm_data_factory_dataset_sql_server_table.
func (dfdsst dataFactoryDatasetSqlServerTableAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(dfdsst.ref.Append("table_name"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) SchemaColumn() terra.ListValue[datafactorydatasetsqlservertable.SchemaColumnAttributes] {
	return terra.ReferenceAsList[datafactorydatasetsqlservertable.SchemaColumnAttributes](dfdsst.ref.Append("schema_column"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Timeouts() datafactorydatasetsqlservertable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorydatasetsqlservertable.TimeoutsAttributes](dfdsst.ref.Append("timeouts"))
}

type dataFactoryDatasetSqlServerTableState struct {
	AdditionalProperties map[string]string                                    `json:"additional_properties"`
	Annotations          []string                                             `json:"annotations"`
	DataFactoryId        string                                               `json:"data_factory_id"`
	Description          string                                               `json:"description"`
	Folder               string                                               `json:"folder"`
	Id                   string                                               `json:"id"`
	LinkedServiceName    string                                               `json:"linked_service_name"`
	Name                 string                                               `json:"name"`
	Parameters           map[string]string                                    `json:"parameters"`
	TableName            string                                               `json:"table_name"`
	SchemaColumn         []datafactorydatasetsqlservertable.SchemaColumnState `json:"schema_column"`
	Timeouts             *datafactorydatasetsqlservertable.TimeoutsState      `json:"timeouts"`
}