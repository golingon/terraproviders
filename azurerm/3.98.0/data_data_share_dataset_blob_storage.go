// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datadatasharedatasetblobstorage "github.com/golingon/terraproviders/azurerm/3.98.0/datadatasharedatasetblobstorage"
)

// NewDataDataShareDatasetBlobStorage creates a new instance of [DataDataShareDatasetBlobStorage].
func NewDataDataShareDatasetBlobStorage(name string, args DataDataShareDatasetBlobStorageArgs) *DataDataShareDatasetBlobStorage {
	return &DataDataShareDatasetBlobStorage{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataShareDatasetBlobStorage)(nil)

// DataDataShareDatasetBlobStorage represents the Terraform data resource azurerm_data_share_dataset_blob_storage.
type DataDataShareDatasetBlobStorage struct {
	Name string
	Args DataDataShareDatasetBlobStorageArgs
}

// DataSource returns the Terraform object type for [DataDataShareDatasetBlobStorage].
func (dsdbs *DataDataShareDatasetBlobStorage) DataSource() string {
	return "azurerm_data_share_dataset_blob_storage"
}

// LocalName returns the local name for [DataDataShareDatasetBlobStorage].
func (dsdbs *DataDataShareDatasetBlobStorage) LocalName() string {
	return dsdbs.Name
}

// Configuration returns the configuration (args) for [DataDataShareDatasetBlobStorage].
func (dsdbs *DataDataShareDatasetBlobStorage) Configuration() interface{} {
	return dsdbs.Args
}

// Attributes returns the attributes for [DataDataShareDatasetBlobStorage].
func (dsdbs *DataDataShareDatasetBlobStorage) Attributes() dataDataShareDatasetBlobStorageAttributes {
	return dataDataShareDatasetBlobStorageAttributes{ref: terra.ReferenceDataResource(dsdbs)}
}

// DataDataShareDatasetBlobStorageArgs contains the configurations for azurerm_data_share_dataset_blob_storage.
type DataDataShareDatasetBlobStorageArgs struct {
	// DataShareId: string, required
	DataShareId terra.StringValue `hcl:"data_share_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageAccount: min=0
	StorageAccount []datadatasharedatasetblobstorage.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadatasharedatasetblobstorage.Timeouts `hcl:"timeouts,block"`
}
type dataDataShareDatasetBlobStorageAttributes struct {
	ref terra.Reference
}

// ContainerName returns a reference to field container_name of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataDataShareDatasetBlobStorageAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("container_name"))
}

// DataShareId returns a reference to field data_share_id of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataDataShareDatasetBlobStorageAttributes) DataShareId() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("data_share_id"))
}

// DisplayName returns a reference to field display_name of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataDataShareDatasetBlobStorageAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("display_name"))
}

// FilePath returns a reference to field file_path of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataDataShareDatasetBlobStorageAttributes) FilePath() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("file_path"))
}

// FolderPath returns a reference to field folder_path of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataDataShareDatasetBlobStorageAttributes) FolderPath() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("folder_path"))
}

// Id returns a reference to field id of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataDataShareDatasetBlobStorageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_share_dataset_blob_storage.
func (dsdbs dataDataShareDatasetBlobStorageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsdbs.ref.Append("name"))
}

func (dsdbs dataDataShareDatasetBlobStorageAttributes) StorageAccount() terra.ListValue[datadatasharedatasetblobstorage.StorageAccountAttributes] {
	return terra.ReferenceAsList[datadatasharedatasetblobstorage.StorageAccountAttributes](dsdbs.ref.Append("storage_account"))
}

func (dsdbs dataDataShareDatasetBlobStorageAttributes) Timeouts() datadatasharedatasetblobstorage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadatasharedatasetblobstorage.TimeoutsAttributes](dsdbs.ref.Append("timeouts"))
}
