// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydatasetbinary "github.com/golingon/terraproviders/azurerm/3.67.0/datafactorydatasetbinary"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryDatasetBinary creates a new instance of [DataFactoryDatasetBinary].
func NewDataFactoryDatasetBinary(name string, args DataFactoryDatasetBinaryArgs) *DataFactoryDatasetBinary {
	return &DataFactoryDatasetBinary{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDatasetBinary)(nil)

// DataFactoryDatasetBinary represents the Terraform resource azurerm_data_factory_dataset_binary.
type DataFactoryDatasetBinary struct {
	Name      string
	Args      DataFactoryDatasetBinaryArgs
	state     *dataFactoryDatasetBinaryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryDatasetBinary].
func (dfdb *DataFactoryDatasetBinary) Type() string {
	return "azurerm_data_factory_dataset_binary"
}

// LocalName returns the local name for [DataFactoryDatasetBinary].
func (dfdb *DataFactoryDatasetBinary) LocalName() string {
	return dfdb.Name
}

// Configuration returns the configuration (args) for [DataFactoryDatasetBinary].
func (dfdb *DataFactoryDatasetBinary) Configuration() interface{} {
	return dfdb.Args
}

// DependOn is used for other resources to depend on [DataFactoryDatasetBinary].
func (dfdb *DataFactoryDatasetBinary) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdb)
}

// Dependencies returns the list of resources [DataFactoryDatasetBinary] depends_on.
func (dfdb *DataFactoryDatasetBinary) Dependencies() terra.Dependencies {
	return dfdb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryDatasetBinary].
func (dfdb *DataFactoryDatasetBinary) LifecycleManagement() *terra.Lifecycle {
	return dfdb.Lifecycle
}

// Attributes returns the attributes for [DataFactoryDatasetBinary].
func (dfdb *DataFactoryDatasetBinary) Attributes() dataFactoryDatasetBinaryAttributes {
	return dataFactoryDatasetBinaryAttributes{ref: terra.ReferenceResource(dfdb)}
}

// ImportState imports the given attribute values into [DataFactoryDatasetBinary]'s state.
func (dfdb *DataFactoryDatasetBinary) ImportState(av io.Reader) error {
	dfdb.state = &dataFactoryDatasetBinaryState{}
	if err := json.NewDecoder(av).Decode(dfdb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdb.Type(), dfdb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryDatasetBinary] has state.
func (dfdb *DataFactoryDatasetBinary) State() (*dataFactoryDatasetBinaryState, bool) {
	return dfdb.state, dfdb.state != nil
}

// StateMust returns the state for [DataFactoryDatasetBinary]. Panics if the state is nil.
func (dfdb *DataFactoryDatasetBinary) StateMust() *dataFactoryDatasetBinaryState {
	if dfdb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdb.Type(), dfdb.LocalName()))
	}
	return dfdb.state
}

// DataFactoryDatasetBinaryArgs contains the configurations for azurerm_data_factory_dataset_binary.
type DataFactoryDatasetBinaryArgs struct {
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
	// AzureBlobStorageLocation: optional
	AzureBlobStorageLocation *datafactorydatasetbinary.AzureBlobStorageLocation `hcl:"azure_blob_storage_location,block"`
	// Compression: optional
	Compression *datafactorydatasetbinary.Compression `hcl:"compression,block"`
	// HttpServerLocation: optional
	HttpServerLocation *datafactorydatasetbinary.HttpServerLocation `hcl:"http_server_location,block"`
	// SftpServerLocation: optional
	SftpServerLocation *datafactorydatasetbinary.SftpServerLocation `hcl:"sftp_server_location,block"`
	// Timeouts: optional
	Timeouts *datafactorydatasetbinary.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryDatasetBinaryAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_dataset_binary.
func (dfdb dataFactoryDatasetBinaryAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdb.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_dataset_binary.
func (dfdb dataFactoryDatasetBinaryAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfdb.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_dataset_binary.
func (dfdb dataFactoryDatasetBinaryAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfdb.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_dataset_binary.
func (dfdb dataFactoryDatasetBinaryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfdb.ref.Append("description"))
}

// Folder returns a reference to field folder of azurerm_data_factory_dataset_binary.
func (dfdb dataFactoryDatasetBinaryAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(dfdb.ref.Append("folder"))
}

// Id returns a reference to field id of azurerm_data_factory_dataset_binary.
func (dfdb dataFactoryDatasetBinaryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfdb.ref.Append("id"))
}

// LinkedServiceName returns a reference to field linked_service_name of azurerm_data_factory_dataset_binary.
func (dfdb dataFactoryDatasetBinaryAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(dfdb.ref.Append("linked_service_name"))
}

// Name returns a reference to field name of azurerm_data_factory_dataset_binary.
func (dfdb dataFactoryDatasetBinaryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfdb.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_dataset_binary.
func (dfdb dataFactoryDatasetBinaryAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfdb.ref.Append("parameters"))
}

func (dfdb dataFactoryDatasetBinaryAttributes) AzureBlobStorageLocation() terra.ListValue[datafactorydatasetbinary.AzureBlobStorageLocationAttributes] {
	return terra.ReferenceAsList[datafactorydatasetbinary.AzureBlobStorageLocationAttributes](dfdb.ref.Append("azure_blob_storage_location"))
}

func (dfdb dataFactoryDatasetBinaryAttributes) Compression() terra.ListValue[datafactorydatasetbinary.CompressionAttributes] {
	return terra.ReferenceAsList[datafactorydatasetbinary.CompressionAttributes](dfdb.ref.Append("compression"))
}

func (dfdb dataFactoryDatasetBinaryAttributes) HttpServerLocation() terra.ListValue[datafactorydatasetbinary.HttpServerLocationAttributes] {
	return terra.ReferenceAsList[datafactorydatasetbinary.HttpServerLocationAttributes](dfdb.ref.Append("http_server_location"))
}

func (dfdb dataFactoryDatasetBinaryAttributes) SftpServerLocation() terra.ListValue[datafactorydatasetbinary.SftpServerLocationAttributes] {
	return terra.ReferenceAsList[datafactorydatasetbinary.SftpServerLocationAttributes](dfdb.ref.Append("sftp_server_location"))
}

func (dfdb dataFactoryDatasetBinaryAttributes) Timeouts() datafactorydatasetbinary.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorydatasetbinary.TimeoutsAttributes](dfdb.ref.Append("timeouts"))
}

type dataFactoryDatasetBinaryState struct {
	AdditionalProperties     map[string]string                                        `json:"additional_properties"`
	Annotations              []string                                                 `json:"annotations"`
	DataFactoryId            string                                                   `json:"data_factory_id"`
	Description              string                                                   `json:"description"`
	Folder                   string                                                   `json:"folder"`
	Id                       string                                                   `json:"id"`
	LinkedServiceName        string                                                   `json:"linked_service_name"`
	Name                     string                                                   `json:"name"`
	Parameters               map[string]string                                        `json:"parameters"`
	AzureBlobStorageLocation []datafactorydatasetbinary.AzureBlobStorageLocationState `json:"azure_blob_storage_location"`
	Compression              []datafactorydatasetbinary.CompressionState              `json:"compression"`
	HttpServerLocation       []datafactorydatasetbinary.HttpServerLocationState       `json:"http_server_location"`
	SftpServerLocation       []datafactorydatasetbinary.SftpServerLocationState       `json:"sftp_server_location"`
	Timeouts                 *datafactorydatasetbinary.TimeoutsState                  `json:"timeouts"`
}
