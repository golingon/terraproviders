// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydatasetjson "github.com/golingon/terraproviders/azurerm/3.68.0/datafactorydatasetjson"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryDatasetJson creates a new instance of [DataFactoryDatasetJson].
func NewDataFactoryDatasetJson(name string, args DataFactoryDatasetJsonArgs) *DataFactoryDatasetJson {
	return &DataFactoryDatasetJson{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetJson)(nil)

// DataFactoryDatasetJson represents the Terraform resource azurerm_data_factory_dataset_json.
type DataFactoryDatasetJson struct {
	Name      string
	Args      DataFactoryDatasetJsonArgs
	state     *dataFactoryDatasetJsonState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryDatasetJson].
func (dfdj *DataFactoryDatasetJson) Type() string {
	return "azurerm_data_factory_dataset_json"
}

// LocalName returns the local name for [DataFactoryDatasetJson].
func (dfdj *DataFactoryDatasetJson) LocalName() string {
	return dfdj.Name
}

// Configuration returns the configuration (args) for [DataFactoryDatasetJson].
func (dfdj *DataFactoryDatasetJson) Configuration() interface{} {
	return dfdj.Args
}

// DependOn is used for other resources to depend on [DataFactoryDatasetJson].
func (dfdj *DataFactoryDatasetJson) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdj)
}

// Dependencies returns the list of resources [DataFactoryDatasetJson] depends_on.
func (dfdj *DataFactoryDatasetJson) Dependencies() terra.Dependencies {
	return dfdj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryDatasetJson].
func (dfdj *DataFactoryDatasetJson) LifecycleManagement() *terra.Lifecycle {
	return dfdj.Lifecycle
}

// Attributes returns the attributes for [DataFactoryDatasetJson].
func (dfdj *DataFactoryDatasetJson) Attributes() dataFactoryDatasetJsonAttributes {
	return dataFactoryDatasetJsonAttributes{ref: terra.ReferenceResource(dfdj)}
}

// ImportState imports the given attribute values into [DataFactoryDatasetJson]'s state.
func (dfdj *DataFactoryDatasetJson) ImportState(av io.Reader) error {
	dfdj.state = &dataFactoryDatasetJsonState{}
	if err := json.NewDecoder(av).Decode(dfdj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdj.Type(), dfdj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryDatasetJson] has state.
func (dfdj *DataFactoryDatasetJson) State() (*dataFactoryDatasetJsonState, bool) {
	return dfdj.state, dfdj.state != nil
}

// StateMust returns the state for [DataFactoryDatasetJson]. Panics if the state is nil.
func (dfdj *DataFactoryDatasetJson) StateMust() *dataFactoryDatasetJsonState {
	if dfdj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdj.Type(), dfdj.LocalName()))
	}
	return dfdj.state
}

// DataFactoryDatasetJsonArgs contains the configurations for azurerm_data_factory_dataset_json.
type DataFactoryDatasetJsonArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Encoding: string, optional
	Encoding terra.StringValue `hcl:"encoding,attr"`
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
	// AzureBlobStorageLocation: optional
	AzureBlobStorageLocation *datafactorydatasetjson.AzureBlobStorageLocation `hcl:"azure_blob_storage_location,block"`
	// HttpServerLocation: optional
	HttpServerLocation *datafactorydatasetjson.HttpServerLocation `hcl:"http_server_location,block"`
	// SchemaColumn: min=0
	SchemaColumn []datafactorydatasetjson.SchemaColumn `hcl:"schema_column,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datafactorydatasetjson.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryDatasetJsonAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_dataset_json.
func (dfdj dataFactoryDatasetJsonAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdj.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_dataset_json.
func (dfdj dataFactoryDatasetJsonAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfdj.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_dataset_json.
func (dfdj dataFactoryDatasetJsonAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfdj.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_dataset_json.
func (dfdj dataFactoryDatasetJsonAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfdj.ref.Append("description"))
}

// Encoding returns a reference to field encoding of azurerm_data_factory_dataset_json.
func (dfdj dataFactoryDatasetJsonAttributes) Encoding() terra.StringValue {
	return terra.ReferenceAsString(dfdj.ref.Append("encoding"))
}

// Folder returns a reference to field folder of azurerm_data_factory_dataset_json.
func (dfdj dataFactoryDatasetJsonAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(dfdj.ref.Append("folder"))
}

// Id returns a reference to field id of azurerm_data_factory_dataset_json.
func (dfdj dataFactoryDatasetJsonAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfdj.ref.Append("id"))
}

// LinkedServiceName returns a reference to field linked_service_name of azurerm_data_factory_dataset_json.
func (dfdj dataFactoryDatasetJsonAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(dfdj.ref.Append("linked_service_name"))
}

// Name returns a reference to field name of azurerm_data_factory_dataset_json.
func (dfdj dataFactoryDatasetJsonAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfdj.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_dataset_json.
func (dfdj dataFactoryDatasetJsonAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdj.ref.Append("parameters"))
}

func (dfdj dataFactoryDatasetJsonAttributes) AzureBlobStorageLocation() terra.ListValue[datafactorydatasetjson.AzureBlobStorageLocationAttributes] {
	return terra.ReferenceAsList[datafactorydatasetjson.AzureBlobStorageLocationAttributes](dfdj.ref.Append("azure_blob_storage_location"))
}

func (dfdj dataFactoryDatasetJsonAttributes) HttpServerLocation() terra.ListValue[datafactorydatasetjson.HttpServerLocationAttributes] {
	return terra.ReferenceAsList[datafactorydatasetjson.HttpServerLocationAttributes](dfdj.ref.Append("http_server_location"))
}

func (dfdj dataFactoryDatasetJsonAttributes) SchemaColumn() terra.ListValue[datafactorydatasetjson.SchemaColumnAttributes] {
	return terra.ReferenceAsList[datafactorydatasetjson.SchemaColumnAttributes](dfdj.ref.Append("schema_column"))
}

func (dfdj dataFactoryDatasetJsonAttributes) Timeouts() datafactorydatasetjson.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorydatasetjson.TimeoutsAttributes](dfdj.ref.Append("timeouts"))
}

type dataFactoryDatasetJsonState struct {
	AdditionalProperties     map[string]string                                      `json:"additional_properties"`
	Annotations              []string                                               `json:"annotations"`
	DataFactoryId            string                                                 `json:"data_factory_id"`
	Description              string                                                 `json:"description"`
	Encoding                 string                                                 `json:"encoding"`
	Folder                   string                                                 `json:"folder"`
	Id                       string                                                 `json:"id"`
	LinkedServiceName        string                                                 `json:"linked_service_name"`
	Name                     string                                                 `json:"name"`
	Parameters               map[string]string                                      `json:"parameters"`
	AzureBlobStorageLocation []datafactorydatasetjson.AzureBlobStorageLocationState `json:"azure_blob_storage_location"`
	HttpServerLocation       []datafactorydatasetjson.HttpServerLocationState       `json:"http_server_location"`
	SchemaColumn             []datafactorydatasetjson.SchemaColumnState             `json:"schema_column"`
	Timeouts                 *datafactorydatasetjson.TimeoutsState                  `json:"timeouts"`
}
