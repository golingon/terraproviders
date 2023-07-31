// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datasharedatasetblobstorage "github.com/golingon/terraproviders/azurerm/3.67.0/datasharedatasetblobstorage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataShareDatasetBlobStorage creates a new instance of [DataShareDatasetBlobStorage].
func NewDataShareDatasetBlobStorage(name string, args DataShareDatasetBlobStorageArgs) *DataShareDatasetBlobStorage {
	return &DataShareDatasetBlobStorage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataShareDatasetBlobStorage)(nil)

// DataShareDatasetBlobStorage represents the Terraform resource azurerm_data_share_dataset_blob_storage.
type DataShareDatasetBlobStorage struct {
	Name      string
	Args      DataShareDatasetBlobStorageArgs
	state     *dataShareDatasetBlobStorageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataShareDatasetBlobStorage].
func (dsdbs *DataShareDatasetBlobStorage) Type() string {
	return "azurerm_data_share_dataset_blob_storage"
}

// LocalName returns the local name for [DataShareDatasetBlobStorage].
func (dsdbs *DataShareDatasetBlobStorage) LocalName() string {
	return dsdbs.Name
}

// Configuration returns the configuration (args) for [DataShareDatasetBlobStorage].
func (dsdbs *DataShareDatasetBlobStorage) Configuration() interface{} {
	return dsdbs.Args
}

// DependOn is used for other resources to depend on [DataShareDatasetBlobStorage].
func (dsdbs *DataShareDatasetBlobStorage) DependOn() terra.Reference {
	return terra.ReferenceResource(dsdbs)
}

// Dependencies returns the list of resources [DataShareDatasetBlobStorage] depends_on.
func (dsdbs *DataShareDatasetBlobStorage) Dependencies() terra.Dependencies {
	return dsdbs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataShareDatasetBlobStorage].
func (dsdbs *DataShareDatasetBlobStorage) LifecycleManagement() *terra.Lifecycle {
	return dsdbs.Lifecycle
}

// Attributes returns the attributes for [DataShareDatasetBlobStorage].
func (dsdbs *DataShareDatasetBlobStorage) Attributes() dataShareDatasetBlobStorageAttributes {
	return dataShareDatasetBlobStorageAttributes{ref: terra.ReferenceResource(dsdbs)}
}

// ImportState imports the given attribute values into [DataShareDatasetBlobStorage]'s state.
func (dsdbs *DataShareDatasetBlobStorage) ImportState(av io.Reader) error {
	dsdbs.state = &dataShareDatasetBlobStorageState{}
	if err := json.NewDecoder(av).Decode(dsdbs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dsdbs.Type(), dsdbs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataShareDatasetBlobStorage] has state.
func (dsdbs *DataShareDatasetBlobStorage) State() (*dataShareDatasetBlobStorageState, bool) {
	return dsdbs.state, dsdbs.state != nil
}

// StateMust returns the state for [DataShareDatasetBlobStorage]. Panics if the state is nil.
func (dsdbs *DataShareDatasetBlobStorage) StateMust() *dataShareDatasetBlobStorageState {
	if dsdbs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dsdbs.Type(), dsdbs.LocalName()))
	}
	return dsdbs.state
}

// DataShareDatasetBlobStorageArgs contains the configurations for azurerm_data_share_dataset_blob_storage.
type DataShareDatasetBlobStorageArgs struct {
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// DataShareId: string, required
	DataShareId terra.StringValue `hcl:"data_share_id,attr" validate:"required"`
	// FilePath: string, optional
	FilePath terra.StringValue `hcl:"file_path,attr"`
	// FolderPath: string, optional
	FolderPath terra.StringValue `hcl:"folder_path,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageAccount: required
	StorageAccount *datasharedatasetblobstorage.StorageAccount `hcl:"storage_account,block" validate:"required"`
	// Timeouts: optional
	Timeouts *datasharedatasetblobstorage.Timeouts `hcl:"timeouts,block"`
}
type dataShareDatasetBlobStorageAttributes struct {
	ref terra.Reference
}

// ContainerName returns a reference to field container_name of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataShareDatasetBlobStorageAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("container_name"))
}

// DataShareId returns a reference to field data_share_id of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataShareDatasetBlobStorageAttributes) DataShareId() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("data_share_id"))
}

// DisplayName returns a reference to field display_name of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataShareDatasetBlobStorageAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("display_name"))
}

// FilePath returns a reference to field file_path of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataShareDatasetBlobStorageAttributes) FilePath() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("file_path"))
}

// FolderPath returns a reference to field folder_path of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataShareDatasetBlobStorageAttributes) FolderPath() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("folder_path"))
}

// Id returns a reference to field id of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataShareDatasetBlobStorageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataShareDatasetBlobStorageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("name"))
}

func (dsdbs dataShareDatasetBlobStorageAttributes) StorageAccount() terra.ListValue[datasharedatasetblobstorage.StorageAccountAttributes] {
	return terra.ReferenceAsList[datasharedatasetblobstorage.StorageAccountAttributes](dsdbs.ref.Append("storage_account"))
}

func (dsdbs dataShareDatasetBlobStorageAttributes) Timeouts() datasharedatasetblobstorage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasharedatasetblobstorage.TimeoutsAttributes](dsdbs.ref.Append("timeouts"))
}

type dataShareDatasetBlobStorageState struct {
	ContainerName  string                                            `json:"container_name"`
	DataShareId    string                                            `json:"data_share_id"`
	DisplayName    string                                            `json:"display_name"`
	FilePath       string                                            `json:"file_path"`
	FolderPath     string                                            `json:"folder_path"`
	Id             string                                            `json:"id"`
	Name           string                                            `json:"name"`
	StorageAccount []datasharedatasetblobstorage.StorageAccountState `json:"storage_account"`
	Timeouts       *datasharedatasetblobstorage.TimeoutsState        `json:"timeouts"`
}
