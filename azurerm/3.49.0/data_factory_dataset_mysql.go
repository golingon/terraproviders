// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydatasetmysql "github.com/golingon/terraproviders/azurerm/3.49.0/datafactorydatasetmysql"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataFactoryDatasetMysql(name string, args DataFactoryDatasetMysqlArgs) *DataFactoryDatasetMysql {
	return &DataFactoryDatasetMysql{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetMysql)(nil)

type DataFactoryDatasetMysql struct {
	Name  string
	Args  DataFactoryDatasetMysqlArgs
	state *dataFactoryDatasetMysqlState
}

func (dfdm *DataFactoryDatasetMysql) Type() string {
	return "azurerm_data_factory_dataset_mysql"
}

func (dfdm *DataFactoryDatasetMysql) LocalName() string {
	return dfdm.Name
}

func (dfdm *DataFactoryDatasetMysql) Configuration() interface{} {
	return dfdm.Args
}

func (dfdm *DataFactoryDatasetMysql) Attributes() dataFactoryDatasetMysqlAttributes {
	return dataFactoryDatasetMysqlAttributes{ref: terra.ReferenceResource(dfdm)}
}

func (dfdm *DataFactoryDatasetMysql) ImportState(av io.Reader) error {
	dfdm.state = &dataFactoryDatasetMysqlState{}
	if err := json.NewDecoder(av).Decode(dfdm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdm.Type(), dfdm.LocalName(), err)
	}
	return nil
}

func (dfdm *DataFactoryDatasetMysql) State() (*dataFactoryDatasetMysqlState, bool) {
	return dfdm.state, dfdm.state != nil
}

func (dfdm *DataFactoryDatasetMysql) StateMust() *dataFactoryDatasetMysqlState {
	if dfdm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdm.Type(), dfdm.LocalName()))
	}
	return dfdm.state
}

func (dfdm *DataFactoryDatasetMysql) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdm)
}

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
	// DependsOn contains resources that DataFactoryDatasetMysql depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataFactoryDatasetMysqlAttributes struct {
	ref terra.Reference
}

func (dfdm dataFactoryDatasetMysqlAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dfdm.ref.Append("additional_properties"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dfdm.ref.Append("annotations"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceString(dfdm.ref.Append("data_factory_id"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dfdm.ref.Append("description"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) Folder() terra.StringValue {
	return terra.ReferenceString(dfdm.ref.Append("folder"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dfdm.ref.Append("id"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceString(dfdm.ref.Append("linked_service_name"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dfdm.ref.Append("name"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dfdm.ref.Append("parameters"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) TableName() terra.StringValue {
	return terra.ReferenceString(dfdm.ref.Append("table_name"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) SchemaColumn() terra.ListValue[datafactorydatasetmysql.SchemaColumnAttributes] {
	return terra.ReferenceList[datafactorydatasetmysql.SchemaColumnAttributes](dfdm.ref.Append("schema_column"))
}

func (dfdm dataFactoryDatasetMysqlAttributes) Timeouts() datafactorydatasetmysql.TimeoutsAttributes {
	return terra.ReferenceSingle[datafactorydatasetmysql.TimeoutsAttributes](dfdm.ref.Append("timeouts"))
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
