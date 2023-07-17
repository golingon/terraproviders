// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydatasetmysql "github.com/golingon/terraproviders/azurerm/3.65.0/datafactorydatasetmysql"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryDatasetMysql creates a new instance of [DataFactoryDatasetMysql].
func NewDataFactoryDatasetMysql(name string, args DataFactoryDatasetMysqlArgs) *DataFactoryDatasetMysql {
	return &DataFactoryDatasetMysql{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetMysql)(nil)

// DataFactoryDatasetMysql represents the Terraform resource azurerm_data_factory_dataset_mysql.
type DataFactoryDatasetMysql struct {
	Name      string
	Args      DataFactoryDatasetMysqlArgs
	state     *dataFactoryDatasetMysqlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryDatasetMysql].
func (dfdm *DataFactoryDatasetMysql) Type() string {
	return "azurerm_data_factory_dataset_mysql"
}

// LocalName returns the local name for [DataFactoryDatasetMysql].
func (dfdm *DataFactoryDatasetMysql) LocalName() string {
	return dfdm.Name
}

// Configuration returns the configuration (args) for [DataFactoryDatasetMysql].
func (dfdm *DataFactoryDatasetMysql) Configuration() interface{} {
	return dfdm.Args
}

// DependOn is used for other resources to depend on [DataFactoryDatasetMysql].
func (dfdm *DataFactoryDatasetMysql) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdm)
}

// Dependencies returns the list of resources [DataFactoryDatasetMysql] depends_on.
func (dfdm *DataFactoryDatasetMysql) Dependencies() terra.Dependencies {
	return dfdm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryDatasetMysql].
func (dfdm *DataFactoryDatasetMysql) LifecycleManagement() *terra.Lifecycle {
	return dfdm.Lifecycle
}

// Attributes returns the attributes for [DataFactoryDatasetMysql].
func (dfdm *DataFactoryDatasetMysql) Attributes() dataFactoryDatasetMysqlAttributes {
	return dataFactoryDatasetMysqlAttributes{ref: terra.ReferenceResource(dfdm)}
}

// ImportState imports the given attribute values into [DataFactoryDatasetMysql]'s state.
func (dfdm *DataFactoryDatasetMysql) ImportState(av io.Reader) error {
	dfdm.state = &dataFactoryDatasetMysqlState{}
	if err := json.NewDecoder(av).Decode(dfdm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdm.Type(), dfdm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryDatasetMysql] has state.
func (dfdm *DataFactoryDatasetMysql) State() (*dataFactoryDatasetMysqlState, bool) {
	return dfdm.state, dfdm.state != nil
}

// StateMust returns the state for [DataFactoryDatasetMysql]. Panics if the state is nil.
func (dfdm *DataFactoryDatasetMysql) StateMust() *dataFactoryDatasetMysqlState {
	if dfdm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdm.Type(), dfdm.LocalName()))
	}
	return dfdm.state
}

// DataFactoryDatasetMysqlArgs contains the configurations for azurerm_data_factory_dataset_mysql.
type DataFactoryDatasetMysqlArgs struct {
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
	SchemaColumn []datafactorydatasetmysql.SchemaColumn `hcl:"schema_column,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datafactorydatasetmysql.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryDatasetMysqlAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_dataset_mysql.
func (dfdm dataFactoryDatasetMysqlAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdm.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_dataset_mysql.
func (dfdm dataFactoryDatasetMysqlAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfdm.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_dataset_mysql.
func (dfdm dataFactoryDatasetMysqlAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfdm.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_dataset_mysql.
func (dfdm dataFactoryDatasetMysqlAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfdm.ref.Append("description"))
}

// Folder returns a reference to field folder of azurerm_data_factory_dataset_mysql.
func (dfdm dataFactoryDatasetMysqlAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(dfdm.ref.Append("folder"))
}

// Id returns a reference to field id of azurerm_data_factory_dataset_mysql.
func (dfdm dataFactoryDatasetMysqlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfdm.ref.Append("id"))
}

// LinkedServiceName returns a reference to field linked_service_name of azurerm_data_factory_dataset_mysql.
func (dfdm dataFactoryDatasetMysqlAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(dfdm.ref.Append("linked_service_name"))
}

// Name returns a reference to field name of azurerm_data_factory_dataset_mysql.
func (dfdm dataFactoryDatasetMysqlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfdm.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_dataset_mysql.
func (dfdm dataFactoryDatasetMysqlAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdm.ref.Append("parameters"))
}

// TableName returns a reference to field table_name of azurerm_data_factory_dataset_mysql.
func (dfdm dataFactoryDatasetMysqlAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(dfdm.ref.Append("table_name"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) SchemaColumn() terra.ListValue[datafactorydatasetmysql.SchemaColumnAttributes] {
	return terra.ReferenceAsList[datafactorydatasetmysql.SchemaColumnAttributes](dfdm.ref.Append("schema_column"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) Timeouts() datafactorydatasetmysql.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorydatasetmysql.TimeoutsAttributes](dfdm.ref.Append("timeouts"))
}

type dataFactoryDatasetMysqlState struct {
	AdditionalProperties map[string]string                           `json:"additional_properties"`
	Annotations          []string                                    `json:"annotations"`
	DataFactoryId        string                                      `json:"data_factory_id"`
	Description          string                                      `json:"description"`
	Folder               string                                      `json:"folder"`
	Id                   string                                      `json:"id"`
	LinkedServiceName    string                                      `json:"linked_service_name"`
	Name                 string                                      `json:"name"`
	Parameters           map[string]string                           `json:"parameters"`
	TableName            string                                      `json:"table_name"`
	SchemaColumn         []datafactorydatasetmysql.SchemaColumnState `json:"schema_column"`
	Timeouts             *datafactorydatasetmysql.TimeoutsState      `json:"timeouts"`
}