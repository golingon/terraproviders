// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydatasetazureblob "github.com/golingon/terraproviders/azurerm/3.55.0/datafactorydatasetazureblob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryDatasetAzureBlob creates a new instance of [DataFactoryDatasetAzureBlob].
func NewDataFactoryDatasetAzureBlob(name string, args DataFactoryDatasetAzureBlobArgs) *DataFactoryDatasetAzureBlob {
	return &DataFactoryDatasetAzureBlob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetAzureBlob)(nil)

// DataFactoryDatasetAzureBlob represents the Terraform resource azurerm_data_factory_dataset_azure_blob.
type DataFactoryDatasetAzureBlob struct {
	Name      string
	Args      DataFactoryDatasetAzureBlobArgs
	state     *dataFactoryDatasetAzureBlobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryDatasetAzureBlob].
func (dfdab *DataFactoryDatasetAzureBlob) Type() string {
	return "azurerm_data_factory_dataset_azure_blob"
}

// LocalName returns the local name for [DataFactoryDatasetAzureBlob].
func (dfdab *DataFactoryDatasetAzureBlob) LocalName() string {
	return dfdab.Name
}

// Configuration returns the configuration (args) for [DataFactoryDatasetAzureBlob].
func (dfdab *DataFactoryDatasetAzureBlob) Configuration() interface{} {
	return dfdab.Args
}

// DependOn is used for other resources to depend on [DataFactoryDatasetAzureBlob].
func (dfdab *DataFactoryDatasetAzureBlob) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdab)
}

// Dependencies returns the list of resources [DataFactoryDatasetAzureBlob] depends_on.
func (dfdab *DataFactoryDatasetAzureBlob) Dependencies() terra.Dependencies {
	return dfdab.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryDatasetAzureBlob].
func (dfdab *DataFactoryDatasetAzureBlob) LifecycleManagement() *terra.Lifecycle {
	return dfdab.Lifecycle
}

// Attributes returns the attributes for [DataFactoryDatasetAzureBlob].
func (dfdab *DataFactoryDatasetAzureBlob) Attributes() dataFactoryDatasetAzureBlobAttributes {
	return dataFactoryDatasetAzureBlobAttributes{ref: terra.ReferenceResource(dfdab)}
}

// ImportState imports the given attribute values into [DataFactoryDatasetAzureBlob]'s state.
func (dfdab *DataFactoryDatasetAzureBlob) ImportState(av io.Reader) error {
	dfdab.state = &dataFactoryDatasetAzureBlobState{}
	if err := json.NewDecoder(av).Decode(dfdab.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdab.Type(), dfdab.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryDatasetAzureBlob] has state.
func (dfdab *DataFactoryDatasetAzureBlob) State() (*dataFactoryDatasetAzureBlobState, bool) {
	return dfdab.state, dfdab.state != nil
}

// StateMust returns the state for [DataFactoryDatasetAzureBlob]. Panics if the state is nil.
func (dfdab *DataFactoryDatasetAzureBlob) StateMust() *dataFactoryDatasetAzureBlobState {
	if dfdab.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdab.Type(), dfdab.LocalName()))
	}
	return dfdab.state
}

// DataFactoryDatasetAzureBlobArgs contains the configurations for azurerm_data_factory_dataset_azure_blob.
type DataFactoryDatasetAzureBlobArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DynamicFilenameEnabled: bool, optional
	DynamicFilenameEnabled terra.BoolValue `hcl:"dynamic_filename_enabled,attr"`
	// DynamicPathEnabled: bool, optional
	DynamicPathEnabled terra.BoolValue `hcl:"dynamic_path_enabled,attr"`
	// Filename: string, optional
	Filename terra.StringValue `hcl:"filename,attr"`
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
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// SchemaColumn: min=0
	SchemaColumn []datafactorydatasetazureblob.SchemaColumn `hcl:"schema_column,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datafactorydatasetazureblob.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryDatasetAzureBlobAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdab.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfdab.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfdab.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfdab.ref.Append("description"))
}

// DynamicFilenameEnabled returns a reference to field dynamic_filename_enabled of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) DynamicFilenameEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dfdab.ref.Append("dynamic_filename_enabled"))
}

// DynamicPathEnabled returns a reference to field dynamic_path_enabled of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) DynamicPathEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dfdab.ref.Append("dynamic_path_enabled"))
}

// Filename returns a reference to field filename of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) Filename() terra.StringValue {
	return terra.ReferenceAsString(dfdab.ref.Append("filename"))
}

// Folder returns a reference to field folder of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(dfdab.ref.Append("folder"))
}

// Id returns a reference to field id of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfdab.ref.Append("id"))
}

// LinkedServiceName returns a reference to field linked_service_name of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(dfdab.ref.Append("linked_service_name"))
}

// Name returns a reference to field name of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfdab.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdab.ref.Append("parameters"))
}

// Path returns a reference to field path of azurerm_data_factory_dataset_azure_blob.
func (dfdab dataFactoryDatasetAzureBlobAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(dfdab.ref.Append("path"))
}

func (dfdab dataFactoryDatasetAzureBlobAttributes) SchemaColumn() terra.ListValue[datafactorydatasetazureblob.SchemaColumnAttributes] {
	return terra.ReferenceAsList[datafactorydatasetazureblob.SchemaColumnAttributes](dfdab.ref.Append("schema_column"))
}

func (dfdab dataFactoryDatasetAzureBlobAttributes) Timeouts() datafactorydatasetazureblob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorydatasetazureblob.TimeoutsAttributes](dfdab.ref.Append("timeouts"))
}

type dataFactoryDatasetAzureBlobState struct {
	AdditionalProperties   map[string]string                               `json:"additional_properties"`
	Annotations            []string                                        `json:"annotations"`
	DataFactoryId          string                                          `json:"data_factory_id"`
	Description            string                                          `json:"description"`
	DynamicFilenameEnabled bool                                            `json:"dynamic_filename_enabled"`
	DynamicPathEnabled     bool                                            `json:"dynamic_path_enabled"`
	Filename               string                                          `json:"filename"`
	Folder                 string                                          `json:"folder"`
	Id                     string                                          `json:"id"`
	LinkedServiceName      string                                          `json:"linked_service_name"`
	Name                   string                                          `json:"name"`
	Parameters             map[string]string                               `json:"parameters"`
	Path                   string                                          `json:"path"`
	SchemaColumn           []datafactorydatasetazureblob.SchemaColumnState `json:"schema_column"`
	Timeouts               *datafactorydatasetazureblob.TimeoutsState      `json:"timeouts"`
}
