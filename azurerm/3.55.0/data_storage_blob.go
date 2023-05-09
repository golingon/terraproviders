// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datastorageblob "github.com/golingon/terraproviders/azurerm/3.55.0/datastorageblob"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataStorageBlob creates a new instance of [DataStorageBlob].
func NewDataStorageBlob(name string, args DataStorageBlobArgs) *DataStorageBlob {
	return &DataStorageBlob{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageBlob)(nil)

// DataStorageBlob represents the Terraform data resource azurerm_storage_blob.
type DataStorageBlob struct {
	Name string
	Args DataStorageBlobArgs
}

// DataSource returns the Terraform object type for [DataStorageBlob].
func (sb *DataStorageBlob) DataSource() string {
	return "azurerm_storage_blob"
}

// LocalName returns the local name for [DataStorageBlob].
func (sb *DataStorageBlob) LocalName() string {
	return sb.Name
}

// Configuration returns the configuration (args) for [DataStorageBlob].
func (sb *DataStorageBlob) Configuration() interface{} {
	return sb.Args
}

// Attributes returns the attributes for [DataStorageBlob].
func (sb *DataStorageBlob) Attributes() dataStorageBlobAttributes {
	return dataStorageBlobAttributes{ref: terra.ReferenceDataResource(sb)}
}

// DataStorageBlobArgs contains the configurations for azurerm_storage_blob.
type DataStorageBlobArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// StorageContainerName: string, required
	StorageContainerName terra.StringValue `hcl:"storage_container_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datastorageblob.Timeouts `hcl:"timeouts,block"`
}
type dataStorageBlobAttributes struct {
	ref terra.Reference
}

// AccessTier returns a reference to field access_tier of azurerm_storage_blob.
func (sb dataStorageBlobAttributes) AccessTier() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("access_tier"))
}

// ContentMd5 returns a reference to field content_md5 of azurerm_storage_blob.
func (sb dataStorageBlobAttributes) ContentMd5() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("content_md5"))
}

// ContentType returns a reference to field content_type of azurerm_storage_blob.
func (sb dataStorageBlobAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("content_type"))
}

// Id returns a reference to field id of azurerm_storage_blob.
func (sb dataStorageBlobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_storage_blob.
func (sb dataStorageBlobAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sb.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_storage_blob.
func (sb dataStorageBlobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("name"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_storage_blob.
func (sb dataStorageBlobAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("storage_account_name"))
}

// StorageContainerName returns a reference to field storage_container_name of azurerm_storage_blob.
func (sb dataStorageBlobAttributes) StorageContainerName() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("storage_container_name"))
}

// Type returns a reference to field type of azurerm_storage_blob.
func (sb dataStorageBlobAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("type"))
}

// Url returns a reference to field url of azurerm_storage_blob.
func (sb dataStorageBlobAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("url"))
}

func (sb dataStorageBlobAttributes) Timeouts() datastorageblob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastorageblob.TimeoutsAttributes](sb.ref.Append("timeouts"))
}
