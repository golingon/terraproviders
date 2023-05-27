// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydatasetpostgresql "github.com/golingon/terraproviders/azurerm/3.58.0/datafactorydatasetpostgresql"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryDatasetPostgresql creates a new instance of [DataFactoryDatasetPostgresql].
func NewDataFactoryDatasetPostgresql(name string, args DataFactoryDatasetPostgresqlArgs) *DataFactoryDatasetPostgresql {
	return &DataFactoryDatasetPostgresql{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetPostgresql)(nil)

// DataFactoryDatasetPostgresql represents the Terraform resource azurerm_data_factory_dataset_postgresql.
type DataFactoryDatasetPostgresql struct {
	Name      string
	Args      DataFactoryDatasetPostgresqlArgs
	state     *dataFactoryDatasetPostgresqlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryDatasetPostgresql].
func (dfdp *DataFactoryDatasetPostgresql) Type() string {
	return "azurerm_data_factory_dataset_postgresql"
}

// LocalName returns the local name for [DataFactoryDatasetPostgresql].
func (dfdp *DataFactoryDatasetPostgresql) LocalName() string {
	return dfdp.Name
}

// Configuration returns the configuration (args) for [DataFactoryDatasetPostgresql].
func (dfdp *DataFactoryDatasetPostgresql) Configuration() interface{} {
	return dfdp.Args
}

// DependOn is used for other resources to depend on [DataFactoryDatasetPostgresql].
func (dfdp *DataFactoryDatasetPostgresql) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdp)
}

// Dependencies returns the list of resources [DataFactoryDatasetPostgresql] depends_on.
func (dfdp *DataFactoryDatasetPostgresql) Dependencies() terra.Dependencies {
	return dfdp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryDatasetPostgresql].
func (dfdp *DataFactoryDatasetPostgresql) LifecycleManagement() *terra.Lifecycle {
	return dfdp.Lifecycle
}

// Attributes returns the attributes for [DataFactoryDatasetPostgresql].
func (dfdp *DataFactoryDatasetPostgresql) Attributes() dataFactoryDatasetPostgresqlAttributes {
	return dataFactoryDatasetPostgresqlAttributes{ref: terra.ReferenceResource(dfdp)}
}

// ImportState imports the given attribute values into [DataFactoryDatasetPostgresql]'s state.
func (dfdp *DataFactoryDatasetPostgresql) ImportState(av io.Reader) error {
	dfdp.state = &dataFactoryDatasetPostgresqlState{}
	if err := json.NewDecoder(av).Decode(dfdp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdp.Type(), dfdp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryDatasetPostgresql] has state.
func (dfdp *DataFactoryDatasetPostgresql) State() (*dataFactoryDatasetPostgresqlState, bool) {
	return dfdp.state, dfdp.state != nil
}

// StateMust returns the state for [DataFactoryDatasetPostgresql]. Panics if the state is nil.
func (dfdp *DataFactoryDatasetPostgresql) StateMust() *dataFactoryDatasetPostgresqlState {
	if dfdp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdp.Type(), dfdp.LocalName()))
	}
	return dfdp.state
}

// DataFactoryDatasetPostgresqlArgs contains the configurations for azurerm_data_factory_dataset_postgresql.
type DataFactoryDatasetPostgresqlArgs struct {
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
	SchemaColumn []datafactorydatasetpostgresql.SchemaColumn `hcl:"schema_column,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datafactorydatasetpostgresql.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryDatasetPostgresqlAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_dataset_postgresql.
func (dfdp dataFactoryDatasetPostgresqlAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdp.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_dataset_postgresql.
func (dfdp dataFactoryDatasetPostgresqlAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfdp.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_dataset_postgresql.
func (dfdp dataFactoryDatasetPostgresqlAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfdp.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_dataset_postgresql.
func (dfdp dataFactoryDatasetPostgresqlAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfdp.ref.Append("description"))
}

// Folder returns a reference to field folder of azurerm_data_factory_dataset_postgresql.
func (dfdp dataFactoryDatasetPostgresqlAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(dfdp.ref.Append("folder"))
}

// Id returns a reference to field id of azurerm_data_factory_dataset_postgresql.
func (dfdp dataFactoryDatasetPostgresqlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfdp.ref.Append("id"))
}

// LinkedServiceName returns a reference to field linked_service_name of azurerm_data_factory_dataset_postgresql.
func (dfdp dataFactoryDatasetPostgresqlAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(dfdp.ref.Append("linked_service_name"))
}

// Name returns a reference to field name of azurerm_data_factory_dataset_postgresql.
func (dfdp dataFactoryDatasetPostgresqlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfdp.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_dataset_postgresql.
func (dfdp dataFactoryDatasetPostgresqlAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdp.ref.Append("parameters"))
}

// TableName returns a reference to field table_name of azurerm_data_factory_dataset_postgresql.
func (dfdp dataFactoryDatasetPostgresqlAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(dfdp.ref.Append("table_name"))
}

func (dfdp dataFactoryDatasetPostgresqlAttributes) SchemaColumn() terra.ListValue[datafactorydatasetpostgresql.SchemaColumnAttributes] {
	return terra.ReferenceAsList[datafactorydatasetpostgresql.SchemaColumnAttributes](dfdp.ref.Append("schema_column"))
}

func (dfdp dataFactoryDatasetPostgresqlAttributes) Timeouts() datafactorydatasetpostgresql.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorydatasetpostgresql.TimeoutsAttributes](dfdp.ref.Append("timeouts"))
}

type dataFactoryDatasetPostgresqlState struct {
	AdditionalProperties map[string]string                                `json:"additional_properties"`
	Annotations          []string                                         `json:"annotations"`
	DataFactoryId        string                                           `json:"data_factory_id"`
	Description          string                                           `json:"description"`
	Folder               string                                           `json:"folder"`
	Id                   string                                           `json:"id"`
	LinkedServiceName    string                                           `json:"linked_service_name"`
	Name                 string                                           `json:"name"`
	Parameters           map[string]string                                `json:"parameters"`
	TableName            string                                           `json:"table_name"`
	SchemaColumn         []datafactorydatasetpostgresql.SchemaColumnState `json:"schema_column"`
	Timeouts             *datafactorydatasetpostgresql.TimeoutsState      `json:"timeouts"`
}
