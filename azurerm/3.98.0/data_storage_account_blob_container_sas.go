// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datastorageaccountblobcontainersas "github.com/golingon/terraproviders/azurerm/3.98.0/datastorageaccountblobcontainersas"
)

// NewDataStorageAccountBlobContainerSas creates a new instance of [DataStorageAccountBlobContainerSas].
func NewDataStorageAccountBlobContainerSas(name string, args DataStorageAccountBlobContainerSasArgs) *DataStorageAccountBlobContainerSas {
	return &DataStorageAccountBlobContainerSas{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageAccountBlobContainerSas)(nil)

// DataStorageAccountBlobContainerSas represents the Terraform data resource azurerm_storage_account_blob_container_sas.
type DataStorageAccountBlobContainerSas struct {
	Name string
	Args DataStorageAccountBlobContainerSasArgs
}

// DataSource returns the Terraform object type for [DataStorageAccountBlobContainerSas].
func (sabcs *DataStorageAccountBlobContainerSas) DataSource() string {
	return "azurerm_storage_account_blob_container_sas"
}

// LocalName returns the local name for [DataStorageAccountBlobContainerSas].
func (sabcs *DataStorageAccountBlobContainerSas) LocalName() string {
	return sabcs.Name
}

// Configuration returns the configuration (args) for [DataStorageAccountBlobContainerSas].
func (sabcs *DataStorageAccountBlobContainerSas) Configuration() interface{} {
	return sabcs.Args
}

// Attributes returns the attributes for [DataStorageAccountBlobContainerSas].
func (sabcs *DataStorageAccountBlobContainerSas) Attributes() dataStorageAccountBlobContainerSasAttributes {
	return dataStorageAccountBlobContainerSasAttributes{ref: terra.ReferenceDataResource(sabcs)}
}

// DataStorageAccountBlobContainerSasArgs contains the configurations for azurerm_storage_account_blob_container_sas.
type DataStorageAccountBlobContainerSasArgs struct {
	// CacheControl: string, optional
	CacheControl terra.StringValue `hcl:"cache_control,attr"`
	// ConnectionString: string, required
	ConnectionString terra.StringValue `hcl:"connection_string,attr" validate:"required"`
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// ContentDisposition: string, optional
	ContentDisposition terra.StringValue `hcl:"content_disposition,attr"`
	// ContentEncoding: string, optional
	ContentEncoding terra.StringValue `hcl:"content_encoding,attr"`
	// ContentLanguage: string, optional
	ContentLanguage terra.StringValue `hcl:"content_language,attr"`
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// Expiry: string, required
	Expiry terra.StringValue `hcl:"expiry,attr" validate:"required"`
	// HttpsOnly: bool, optional
	HttpsOnly terra.BoolValue `hcl:"https_only,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// Start: string, required
	Start terra.StringValue `hcl:"start,attr" validate:"required"`
	// Permissions: required
	Permissions *datastorageaccountblobcontainersas.Permissions `hcl:"permissions,block" validate:"required"`
	// Timeouts: optional
	Timeouts *datastorageaccountblobcontainersas.Timeouts `hcl:"timeouts,block"`
}
type dataStorageAccountBlobContainerSasAttributes struct {
	ref terra.Reference
}

// CacheControl returns a reference to field cache_control of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("cache_control"))
}

// ConnectionString returns a reference to field connection_string of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("connection_string"))
}

// ContainerName returns a reference to field container_name of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("container_name"))
}

// ContentDisposition returns a reference to field content_disposition of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("content_disposition"))
}

// ContentEncoding returns a reference to field content_encoding of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("content_encoding"))
}

// ContentLanguage returns a reference to field content_language of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) ContentLanguage() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("content_language"))
}

// ContentType returns a reference to field content_type of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("content_type"))
}

// Expiry returns a reference to field expiry of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) Expiry() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("expiry"))
}

// HttpsOnly returns a reference to field https_only of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(sabcs.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("ip_address"))
}

// Sas returns a reference to field sas of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) Sas() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("sas"))
}

// Start returns a reference to field start of azurerm_storage_account_blob_container_sas.
func (sabcs dataStorageAccountBlobContainerSasAttributes) Start() terra.StringValue {
	return terra.ReferenceAsString(sabcs.ref.Append("start"))
}

func (sabcs dataStorageAccountBlobContainerSasAttributes) Permissions() terra.ListValue[datastorageaccountblobcontainersas.PermissionsAttributes] {
	return terra.ReferenceAsList[datastorageaccountblobcontainersas.PermissionsAttributes](sabcs.ref.Append("permissions"))
}

func (sabcs dataStorageAccountBlobContainerSasAttributes) Timeouts() datastorageaccountblobcontainersas.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastorageaccountblobcontainersas.TimeoutsAttributes](sabcs.ref.Append("timeouts"))
}
