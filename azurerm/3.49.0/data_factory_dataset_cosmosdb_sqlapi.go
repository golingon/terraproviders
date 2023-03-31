// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydatasetcosmosdbsqlapi "github.com/golingon/terraproviders/azurerm/3.49.0/datafactorydatasetcosmosdbsqlapi"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataFactoryDatasetCosmosdbSqlapi(name string, args DataFactoryDatasetCosmosdbSqlapiArgs) *DataFactoryDatasetCosmosdbSqlapi {
	return &DataFactoryDatasetCosmosdbSqlapi{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetCosmosdbSqlapi)(nil)

type DataFactoryDatasetCosmosdbSqlapi struct {
	Name  string
	Args  DataFactoryDatasetCosmosdbSqlapiArgs
	state *dataFactoryDatasetCosmosdbSqlapiState
}

func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) Type() string {
	return "azurerm_data_factory_dataset_cosmosdb_sqlapi"
}

func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) LocalName() string {
	return dfdcs.Name
}

func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) Configuration() interface{} {
	return dfdcs.Args
}

func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) Attributes() dataFactoryDatasetCosmosdbSqlapiAttributes {
	return dataFactoryDatasetCosmosdbSqlapiAttributes{ref: terra.ReferenceResource(dfdcs)}
}

func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) ImportState(av io.Reader) error {
	dfdcs.state = &dataFactoryDatasetCosmosdbSqlapiState{}
	if err := json.NewDecoder(av).Decode(dfdcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdcs.Type(), dfdcs.LocalName(), err)
	}
	return nil
}

func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) State() (*dataFactoryDatasetCosmosdbSqlapiState, bool) {
	return dfdcs.state, dfdcs.state != nil
}

func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) StateMust() *dataFactoryDatasetCosmosdbSqlapiState {
	if dfdcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdcs.Type(), dfdcs.LocalName()))
	}
	return dfdcs.state
}

func (dfdcs *DataFactoryDatasetCosmosdbSqlapi) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdcs)
}

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
	// DependsOn contains resources that DataFactoryDatasetCosmosdbSqlapi depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataFactoryDatasetCosmosdbSqlapiAttributes struct {
	ref terra.Reference
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dfdcs.ref.Append("additional_properties"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dfdcs.ref.Append("annotations"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) CollectionName() terra.StringValue {
	return terra.ReferenceString(dfdcs.ref.Append("collection_name"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceString(dfdcs.ref.Append("data_factory_id"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dfdcs.ref.Append("description"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Folder() terra.StringValue {
	return terra.ReferenceString(dfdcs.ref.Append("folder"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dfdcs.ref.Append("id"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceString(dfdcs.ref.Append("linked_service_name"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dfdcs.ref.Append("name"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dfdcs.ref.Append("parameters"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) SchemaColumn() terra.ListValue[datafactorydatasetcosmosdbsqlapi.SchemaColumnAttributes] {
	return terra.ReferenceList[datafactorydatasetcosmosdbsqlapi.SchemaColumnAttributes](dfdcs.ref.Append("schema_column"))
}

func (dfdcs dataFactoryDatasetCosmosdbSqlapiAttributes) Timeouts() datafactorydatasetcosmosdbsqlapi.TimeoutsAttributes {
	return terra.ReferenceSingle[datafactorydatasetcosmosdbsqlapi.TimeoutsAttributes](dfdcs.ref.Append("timeouts"))
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
