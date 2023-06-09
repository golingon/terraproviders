// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydatasetsnowflake "github.com/golingon/terraproviders/azurerm/3.63.0/datafactorydatasetsnowflake"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryDatasetSnowflake creates a new instance of [DataFactoryDatasetSnowflake].
func NewDataFactoryDatasetSnowflake(name string, args DataFactoryDatasetSnowflakeArgs) *DataFactoryDatasetSnowflake {
	return &DataFactoryDatasetSnowflake{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetSnowflake)(nil)

// DataFactoryDatasetSnowflake represents the Terraform resource azurerm_data_factory_dataset_snowflake.
type DataFactoryDatasetSnowflake struct {
	Name      string
	Args      DataFactoryDatasetSnowflakeArgs
	state     *dataFactoryDatasetSnowflakeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryDatasetSnowflake].
func (dfds *DataFactoryDatasetSnowflake) Type() string {
	return "azurerm_data_factory_dataset_snowflake"
}

// LocalName returns the local name for [DataFactoryDatasetSnowflake].
func (dfds *DataFactoryDatasetSnowflake) LocalName() string {
	return dfds.Name
}

// Configuration returns the configuration (args) for [DataFactoryDatasetSnowflake].
func (dfds *DataFactoryDatasetSnowflake) Configuration() interface{} {
	return dfds.Args
}

// DependOn is used for other resources to depend on [DataFactoryDatasetSnowflake].
func (dfds *DataFactoryDatasetSnowflake) DependOn() terra.Reference {
	return terra.ReferenceResource(dfds)
}

// Dependencies returns the list of resources [DataFactoryDatasetSnowflake] depends_on.
func (dfds *DataFactoryDatasetSnowflake) Dependencies() terra.Dependencies {
	return dfds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryDatasetSnowflake].
func (dfds *DataFactoryDatasetSnowflake) LifecycleManagement() *terra.Lifecycle {
	return dfds.Lifecycle
}

// Attributes returns the attributes for [DataFactoryDatasetSnowflake].
func (dfds *DataFactoryDatasetSnowflake) Attributes() dataFactoryDatasetSnowflakeAttributes {
	return dataFactoryDatasetSnowflakeAttributes{ref: terra.ReferenceResource(dfds)}
}

// ImportState imports the given attribute values into [DataFactoryDatasetSnowflake]'s state.
func (dfds *DataFactoryDatasetSnowflake) ImportState(av io.Reader) error {
	dfds.state = &dataFactoryDatasetSnowflakeState{}
	if err := json.NewDecoder(av).Decode(dfds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfds.Type(), dfds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryDatasetSnowflake] has state.
func (dfds *DataFactoryDatasetSnowflake) State() (*dataFactoryDatasetSnowflakeState, bool) {
	return dfds.state, dfds.state != nil
}

// StateMust returns the state for [DataFactoryDatasetSnowflake]. Panics if the state is nil.
func (dfds *DataFactoryDatasetSnowflake) StateMust() *dataFactoryDatasetSnowflakeState {
	if dfds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfds.Type(), dfds.LocalName()))
	}
	return dfds.state
}

// DataFactoryDatasetSnowflakeArgs contains the configurations for azurerm_data_factory_dataset_snowflake.
type DataFactoryDatasetSnowflakeArgs struct {
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
	// SchemaName: string, optional
	SchemaName terra.StringValue `hcl:"schema_name,attr"`
	// TableName: string, optional
	TableName terra.StringValue `hcl:"table_name,attr"`
	// SchemaColumn: min=0
	SchemaColumn []datafactorydatasetsnowflake.SchemaColumn `hcl:"schema_column,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datafactorydatasetsnowflake.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryDatasetSnowflakeAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_dataset_snowflake.
func (dfds dataFactoryDatasetSnowflakeAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfds.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_dataset_snowflake.
func (dfds dataFactoryDatasetSnowflakeAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfds.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_dataset_snowflake.
func (dfds dataFactoryDatasetSnowflakeAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfds.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_dataset_snowflake.
func (dfds dataFactoryDatasetSnowflakeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfds.ref.Append("description"))
}

// Folder returns a reference to field folder of azurerm_data_factory_dataset_snowflake.
func (dfds dataFactoryDatasetSnowflakeAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(dfds.ref.Append("folder"))
}

// Id returns a reference to field id of azurerm_data_factory_dataset_snowflake.
func (dfds dataFactoryDatasetSnowflakeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfds.ref.Append("id"))
}

// LinkedServiceName returns a reference to field linked_service_name of azurerm_data_factory_dataset_snowflake.
func (dfds dataFactoryDatasetSnowflakeAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(dfds.ref.Append("linked_service_name"))
}

// Name returns a reference to field name of azurerm_data_factory_dataset_snowflake.
func (dfds dataFactoryDatasetSnowflakeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfds.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_dataset_snowflake.
func (dfds dataFactoryDatasetSnowflakeAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfds.ref.Append("parameters"))
}

// SchemaName returns a reference to field schema_name of azurerm_data_factory_dataset_snowflake.
func (dfds dataFactoryDatasetSnowflakeAttributes) SchemaName() terra.StringValue {
	return terra.ReferenceAsString(dfds.ref.Append("schema_name"))
}

// TableName returns a reference to field table_name of azurerm_data_factory_dataset_snowflake.
func (dfds dataFactoryDatasetSnowflakeAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(dfds.ref.Append("table_name"))
}

func (dfds dataFactoryDatasetSnowflakeAttributes) SchemaColumn() terra.ListValue[datafactorydatasetsnowflake.SchemaColumnAttributes] {
	return terra.ReferenceAsList[datafactorydatasetsnowflake.SchemaColumnAttributes](dfds.ref.Append("schema_column"))
}

func (dfds dataFactoryDatasetSnowflakeAttributes) Timeouts() datafactorydatasetsnowflake.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorydatasetsnowflake.TimeoutsAttributes](dfds.ref.Append("timeouts"))
}

type dataFactoryDatasetSnowflakeState struct {
	AdditionalProperties map[string]string                               `json:"additional_properties"`
	Annotations          []string                                        `json:"annotations"`
	DataFactoryId        string                                          `json:"data_factory_id"`
	Description          string                                          `json:"description"`
	Folder               string                                          `json:"folder"`
	Id                   string                                          `json:"id"`
	LinkedServiceName    string                                          `json:"linked_service_name"`
	Name                 string                                          `json:"name"`
	Parameters           map[string]string                               `json:"parameters"`
	SchemaName           string                                          `json:"schema_name"`
	TableName            string                                          `json:"table_name"`
	SchemaColumn         []datafactorydatasetsnowflake.SchemaColumnState `json:"schema_column"`
	Timeouts             *datafactorydatasetsnowflake.TimeoutsState      `json:"timeouts"`
}
