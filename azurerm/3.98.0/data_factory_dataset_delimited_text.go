// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	datafactorydatasetdelimitedtext "github.com/golingon/terraproviders/azurerm/3.98.0/datafactorydatasetdelimitedtext"
	"io"
)

// NewDataFactoryDatasetDelimitedText creates a new instance of [DataFactoryDatasetDelimitedText].
func NewDataFactoryDatasetDelimitedText(name string, args DataFactoryDatasetDelimitedTextArgs) *DataFactoryDatasetDelimitedText {
	return &DataFactoryDatasetDelimitedText{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetDelimitedText)(nil)

// DataFactoryDatasetDelimitedText represents the Terraform resource azurerm_data_factory_dataset_delimited_text.
type DataFactoryDatasetDelimitedText struct {
	Name      string
	Args      DataFactoryDatasetDelimitedTextArgs
	state     *dataFactoryDatasetDelimitedTextState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryDatasetDelimitedText].
func (dfddt *DataFactoryDatasetDelimitedText) Type() string {
	return "azurerm_data_factory_dataset_delimited_text"
}

// LocalName returns the local name for [DataFactoryDatasetDelimitedText].
func (dfddt *DataFactoryDatasetDelimitedText) LocalName() string {
	return dfddt.Name
}

// Configuration returns the configuration (args) for [DataFactoryDatasetDelimitedText].
func (dfddt *DataFactoryDatasetDelimitedText) Configuration() interface{} {
	return dfddt.Args
}

// DependOn is used for other resources to depend on [DataFactoryDatasetDelimitedText].
func (dfddt *DataFactoryDatasetDelimitedText) DependOn() terra.Reference {
	return terra.ReferenceResource(dfddt)
}

// Dependencies returns the list of resources [DataFactoryDatasetDelimitedText] depends_on.
func (dfddt *DataFactoryDatasetDelimitedText) Dependencies() terra.Dependencies {
	return dfddt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryDatasetDelimitedText].
func (dfddt *DataFactoryDatasetDelimitedText) LifecycleManagement() *terra.Lifecycle {
	return dfddt.Lifecycle
}

// Attributes returns the attributes for [DataFactoryDatasetDelimitedText].
func (dfddt *DataFactoryDatasetDelimitedText) Attributes() dataFactoryDatasetDelimitedTextAttributes {
	return dataFactoryDatasetDelimitedTextAttributes{ref: terra.ReferenceResource(dfddt)}
}

// ImportState imports the given attribute values into [DataFactoryDatasetDelimitedText]'s state.
func (dfddt *DataFactoryDatasetDelimitedText) ImportState(av io.Reader) error {
	dfddt.state = &dataFactoryDatasetDelimitedTextState{}
	if err := json.NewDecoder(av).Decode(dfddt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfddt.Type(), dfddt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryDatasetDelimitedText] has state.
func (dfddt *DataFactoryDatasetDelimitedText) State() (*dataFactoryDatasetDelimitedTextState, bool) {
	return dfddt.state, dfddt.state != nil
}

// StateMust returns the state for [DataFactoryDatasetDelimitedText]. Panics if the state is nil.
func (dfddt *DataFactoryDatasetDelimitedText) StateMust() *dataFactoryDatasetDelimitedTextState {
	if dfddt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfddt.Type(), dfddt.LocalName()))
	}
	return dfddt.state
}

// DataFactoryDatasetDelimitedTextArgs contains the configurations for azurerm_data_factory_dataset_delimited_text.
type DataFactoryDatasetDelimitedTextArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// ColumnDelimiter: string, optional
	ColumnDelimiter terra.StringValue `hcl:"column_delimiter,attr"`
	// CompressionCodec: string, optional
	CompressionCodec terra.StringValue `hcl:"compression_codec,attr"`
	// CompressionLevel: string, optional
	CompressionLevel terra.StringValue `hcl:"compression_level,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Encoding: string, optional
	Encoding terra.StringValue `hcl:"encoding,attr"`
	// EscapeCharacter: string, optional
	EscapeCharacter terra.StringValue `hcl:"escape_character,attr"`
	// FirstRowAsHeader: bool, optional
	FirstRowAsHeader terra.BoolValue `hcl:"first_row_as_header,attr"`
	// Folder: string, optional
	Folder terra.StringValue `hcl:"folder,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkedServiceName: string, required
	LinkedServiceName terra.StringValue `hcl:"linked_service_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NullValue: string, optional
	NullValue terra.StringValue `hcl:"null_value,attr"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// QuoteCharacter: string, optional
	QuoteCharacter terra.StringValue `hcl:"quote_character,attr"`
	// RowDelimiter: string, optional
	RowDelimiter terra.StringValue `hcl:"row_delimiter,attr"`
	// AzureBlobFsLocation: optional
	AzureBlobFsLocation *datafactorydatasetdelimitedtext.AzureBlobFsLocation `hcl:"azure_blob_fs_location,block"`
	// AzureBlobStorageLocation: optional
	AzureBlobStorageLocation *datafactorydatasetdelimitedtext.AzureBlobStorageLocation `hcl:"azure_blob_storage_location,block"`
	// HttpServerLocation: optional
	HttpServerLocation *datafactorydatasetdelimitedtext.HttpServerLocation `hcl:"http_server_location,block"`
	// SchemaColumn: min=0
	SchemaColumn []datafactorydatasetdelimitedtext.SchemaColumn `hcl:"schema_column,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datafactorydatasetdelimitedtext.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryDatasetDelimitedTextAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfddt.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfddt.ref.Append("annotations"))
}

// ColumnDelimiter returns a reference to field column_delimiter of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) ColumnDelimiter() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("column_delimiter"))
}

// CompressionCodec returns a reference to field compression_codec of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) CompressionCodec() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("compression_codec"))
}

// CompressionLevel returns a reference to field compression_level of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) CompressionLevel() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("compression_level"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("description"))
}

// Encoding returns a reference to field encoding of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) Encoding() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("encoding"))
}

// EscapeCharacter returns a reference to field escape_character of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) EscapeCharacter() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("escape_character"))
}

// FirstRowAsHeader returns a reference to field first_row_as_header of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) FirstRowAsHeader() terra.BoolValue {
	return terra.ReferenceAsBool(dfddt.ref.Append("first_row_as_header"))
}

// Folder returns a reference to field folder of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("folder"))
}

// Id returns a reference to field id of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("id"))
}

// LinkedServiceName returns a reference to field linked_service_name of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("linked_service_name"))
}

// Name returns a reference to field name of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("name"))
}

// NullValue returns a reference to field null_value of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) NullValue() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("null_value"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfddt.ref.Append("parameters"))
}

// QuoteCharacter returns a reference to field quote_character of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) QuoteCharacter() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("quote_character"))
}

// RowDelimiter returns a reference to field row_delimiter of azurerm_data_factory_dataset_delimited_text.
func (dfddt dataFactoryDatasetDelimitedTextAttributes) RowDelimiter() terra.StringValue {
	return terra.ReferenceAsString(dfddt.ref.Append("row_delimiter"))
}

func (dfddt dataFactoryDatasetDelimitedTextAttributes) AzureBlobFsLocation() terra.ListValue[datafactorydatasetdelimitedtext.AzureBlobFsLocationAttributes] {
	return terra.ReferenceAsList[datafactorydatasetdelimitedtext.AzureBlobFsLocationAttributes](dfddt.ref.Append("azure_blob_fs_location"))
}

func (dfddt dataFactoryDatasetDelimitedTextAttributes) AzureBlobStorageLocation() terra.ListValue[datafactorydatasetdelimitedtext.AzureBlobStorageLocationAttributes] {
	return terra.ReferenceAsList[datafactorydatasetdelimitedtext.AzureBlobStorageLocationAttributes](dfddt.ref.Append("azure_blob_storage_location"))
}

func (dfddt dataFactoryDatasetDelimitedTextAttributes) HttpServerLocation() terra.ListValue[datafactorydatasetdelimitedtext.HttpServerLocationAttributes] {
	return terra.ReferenceAsList[datafactorydatasetdelimitedtext.HttpServerLocationAttributes](dfddt.ref.Append("http_server_location"))
}

func (dfddt dataFactoryDatasetDelimitedTextAttributes) SchemaColumn() terra.ListValue[datafactorydatasetdelimitedtext.SchemaColumnAttributes] {
	return terra.ReferenceAsList[datafactorydatasetdelimitedtext.SchemaColumnAttributes](dfddt.ref.Append("schema_column"))
}

func (dfddt dataFactoryDatasetDelimitedTextAttributes) Timeouts() datafactorydatasetdelimitedtext.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorydatasetdelimitedtext.TimeoutsAttributes](dfddt.ref.Append("timeouts"))
}

type dataFactoryDatasetDelimitedTextState struct {
	AdditionalProperties     map[string]string                                               `json:"additional_properties"`
	Annotations              []string                                                        `json:"annotations"`
	ColumnDelimiter          string                                                          `json:"column_delimiter"`
	CompressionCodec         string                                                          `json:"compression_codec"`
	CompressionLevel         string                                                          `json:"compression_level"`
	DataFactoryId            string                                                          `json:"data_factory_id"`
	Description              string                                                          `json:"description"`
	Encoding                 string                                                          `json:"encoding"`
	EscapeCharacter          string                                                          `json:"escape_character"`
	FirstRowAsHeader         bool                                                            `json:"first_row_as_header"`
	Folder                   string                                                          `json:"folder"`
	Id                       string                                                          `json:"id"`
	LinkedServiceName        string                                                          `json:"linked_service_name"`
	Name                     string                                                          `json:"name"`
	NullValue                string                                                          `json:"null_value"`
	Parameters               map[string]string                                               `json:"parameters"`
	QuoteCharacter           string                                                          `json:"quote_character"`
	RowDelimiter             string                                                          `json:"row_delimiter"`
	AzureBlobFsLocation      []datafactorydatasetdelimitedtext.AzureBlobFsLocationState      `json:"azure_blob_fs_location"`
	AzureBlobStorageLocation []datafactorydatasetdelimitedtext.AzureBlobStorageLocationState `json:"azure_blob_storage_location"`
	HttpServerLocation       []datafactorydatasetdelimitedtext.HttpServerLocationState       `json:"http_server_location"`
	SchemaColumn             []datafactorydatasetdelimitedtext.SchemaColumnState             `json:"schema_column"`
	Timeouts                 *datafactorydatasetdelimitedtext.TimeoutsState                  `json:"timeouts"`
}
