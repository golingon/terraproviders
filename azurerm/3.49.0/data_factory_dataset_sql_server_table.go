// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydatasetsqlservertable "github.com/golingon/terraproviders/azurerm/3.49.0/datafactorydatasetsqlservertable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataFactoryDatasetSqlServerTable(name string, args DataFactoryDatasetSqlServerTableArgs) *DataFactoryDatasetSqlServerTable {
	return &DataFactoryDatasetSqlServerTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetSqlServerTable)(nil)

type DataFactoryDatasetSqlServerTable struct {
	Name  string
	Args  DataFactoryDatasetSqlServerTableArgs
	state *dataFactoryDatasetSqlServerTableState
}

func (dfdsst *DataFactoryDatasetSqlServerTable) Type() string {
	return "azurerm_data_factory_dataset_sql_server_table"
}

func (dfdsst *DataFactoryDatasetSqlServerTable) LocalName() string {
	return dfdsst.Name
}

func (dfdsst *DataFactoryDatasetSqlServerTable) Configuration() interface{} {
	return dfdsst.Args
}

func (dfdsst *DataFactoryDatasetSqlServerTable) Attributes() dataFactoryDatasetSqlServerTableAttributes {
	return dataFactoryDatasetSqlServerTableAttributes{ref: terra.ReferenceResource(dfdsst)}
}

func (dfdsst *DataFactoryDatasetSqlServerTable) ImportState(av io.Reader) error {
	dfdsst.state = &dataFactoryDatasetSqlServerTableState{}
	if err := json.NewDecoder(av).Decode(dfdsst.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdsst.Type(), dfdsst.LocalName(), err)
	}
	return nil
}

func (dfdsst *DataFactoryDatasetSqlServerTable) State() (*dataFactoryDatasetSqlServerTableState, bool) {
	return dfdsst.state, dfdsst.state != nil
}

func (dfdsst *DataFactoryDatasetSqlServerTable) StateMust() *dataFactoryDatasetSqlServerTableState {
	if dfdsst.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdsst.Type(), dfdsst.LocalName()))
	}
	return dfdsst.state
}

func (dfdsst *DataFactoryDatasetSqlServerTable) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdsst)
}

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
	// DependsOn contains resources that DataFactoryDatasetSqlServerTable depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataFactoryDatasetSqlServerTableAttributes struct {
	ref terra.Reference
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dfdsst.ref.Append("additional_properties"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dfdsst.ref.Append("annotations"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceString(dfdsst.ref.Append("data_factory_id"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dfdsst.ref.Append("description"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Folder() terra.StringValue {
	return terra.ReferenceString(dfdsst.ref.Append("folder"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dfdsst.ref.Append("id"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceString(dfdsst.ref.Append("linked_service_name"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dfdsst.ref.Append("name"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dfdsst.ref.Append("parameters"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) TableName() terra.StringValue {
	return terra.ReferenceString(dfdsst.ref.Append("table_name"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) SchemaColumn() terra.ListValue[datafactorydatasetsqlservertable.SchemaColumnAttributes] {
	return terra.ReferenceList[datafactorydatasetsqlservertable.SchemaColumnAttributes](dfdsst.ref.Append("schema_column"))
}

func (dfdsst dataFactoryDatasetSqlServerTableAttributes) Timeouts() datafactorydatasetsqlservertable.TimeoutsAttributes {
	return terra.ReferenceSingle[datafactorydatasetsqlservertable.TimeoutsAttributes](dfdsst.ref.Append("timeouts"))
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
