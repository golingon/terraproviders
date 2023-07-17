// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedserviceazureblobstorage "github.com/golingon/terraproviders/azurerm/3.65.0/datafactorylinkedserviceazureblobstorage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceAzureBlobStorage creates a new instance of [DataFactoryLinkedServiceAzureBlobStorage].
func NewDataFactoryLinkedServiceAzureBlobStorage(name string, args DataFactoryLinkedServiceAzureBlobStorageArgs) *DataFactoryLinkedServiceAzureBlobStorage {
	return &DataFactoryLinkedServiceAzureBlobStorage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceAzureBlobStorage)(nil)

// DataFactoryLinkedServiceAzureBlobStorage represents the Terraform resource azurerm_data_factory_linked_service_azure_blob_storage.
type DataFactoryLinkedServiceAzureBlobStorage struct {
	Name      string
	Args      DataFactoryLinkedServiceAzureBlobStorageArgs
	state     *dataFactoryLinkedServiceAzureBlobStorageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceAzureBlobStorage].
func (dflsabs *DataFactoryLinkedServiceAzureBlobStorage) Type() string {
	return "azurerm_data_factory_linked_service_azure_blob_storage"
}

// LocalName returns the local name for [DataFactoryLinkedServiceAzureBlobStorage].
func (dflsabs *DataFactoryLinkedServiceAzureBlobStorage) LocalName() string {
	return dflsabs.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceAzureBlobStorage].
func (dflsabs *DataFactoryLinkedServiceAzureBlobStorage) Configuration() interface{} {
	return dflsabs.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceAzureBlobStorage].
func (dflsabs *DataFactoryLinkedServiceAzureBlobStorage) DependOn() terra.Reference {
	return terra.ReferenceResource(dflsabs)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceAzureBlobStorage] depends_on.
func (dflsabs *DataFactoryLinkedServiceAzureBlobStorage) Dependencies() terra.Dependencies {
	return dflsabs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceAzureBlobStorage].
func (dflsabs *DataFactoryLinkedServiceAzureBlobStorage) LifecycleManagement() *terra.Lifecycle {
	return dflsabs.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceAzureBlobStorage].
func (dflsabs *DataFactoryLinkedServiceAzureBlobStorage) Attributes() dataFactoryLinkedServiceAzureBlobStorageAttributes {
	return dataFactoryLinkedServiceAzureBlobStorageAttributes{ref: terra.ReferenceResource(dflsabs)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceAzureBlobStorage]'s state.
func (dflsabs *DataFactoryLinkedServiceAzureBlobStorage) ImportState(av io.Reader) error {
	dflsabs.state = &dataFactoryLinkedServiceAzureBlobStorageState{}
	if err := json.NewDecoder(av).Decode(dflsabs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflsabs.Type(), dflsabs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceAzureBlobStorage] has state.
func (dflsabs *DataFactoryLinkedServiceAzureBlobStorage) State() (*dataFactoryLinkedServiceAzureBlobStorageState, bool) {
	return dflsabs.state, dflsabs.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceAzureBlobStorage]. Panics if the state is nil.
func (dflsabs *DataFactoryLinkedServiceAzureBlobStorage) StateMust() *dataFactoryLinkedServiceAzureBlobStorageState {
	if dflsabs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflsabs.Type(), dflsabs.LocalName()))
	}
	return dflsabs.state
}

// DataFactoryLinkedServiceAzureBlobStorageArgs contains the configurations for azurerm_data_factory_linked_service_azure_blob_storage.
type DataFactoryLinkedServiceAzureBlobStorageArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// ConnectionString: string, optional
	ConnectionString terra.StringValue `hcl:"connection_string,attr"`
	// ConnectionStringInsecure: string, optional
	ConnectionStringInsecure terra.StringValue `hcl:"connection_string_insecure,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationRuntimeName: string, optional
	IntegrationRuntimeName terra.StringValue `hcl:"integration_runtime_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// SasUri: string, optional
	SasUri terra.StringValue `hcl:"sas_uri,attr"`
	// ServiceEndpoint: string, optional
	ServiceEndpoint terra.StringValue `hcl:"service_endpoint,attr"`
	// ServicePrincipalId: string, optional
	ServicePrincipalId terra.StringValue `hcl:"service_principal_id,attr"`
	// ServicePrincipalKey: string, optional
	ServicePrincipalKey terra.StringValue `hcl:"service_principal_key,attr"`
	// StorageKind: string, optional
	StorageKind terra.StringValue `hcl:"storage_kind,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// UseManagedIdentity: bool, optional
	UseManagedIdentity terra.BoolValue `hcl:"use_managed_identity,attr"`
	// KeyVaultSasToken: optional
	KeyVaultSasToken *datafactorylinkedserviceazureblobstorage.KeyVaultSasToken `hcl:"key_vault_sas_token,block"`
	// ServicePrincipalLinkedKeyVaultKey: optional
	ServicePrincipalLinkedKeyVaultKey *datafactorylinkedserviceazureblobstorage.ServicePrincipalLinkedKeyVaultKey `hcl:"service_principal_linked_key_vault_key,block"`
	// Timeouts: optional
	Timeouts *datafactorylinkedserviceazureblobstorage.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceAzureBlobStorageAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsabs.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflsabs.ref.Append("annotations"))
}

// ConnectionString returns a reference to field connection_string of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("connection_string"))
}

// ConnectionStringInsecure returns a reference to field connection_string_insecure of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) ConnectionStringInsecure() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("connection_string_insecure"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsabs.ref.Append("parameters"))
}

// SasUri returns a reference to field sas_uri of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) SasUri() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("sas_uri"))
}

// ServiceEndpoint returns a reference to field service_endpoint of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) ServiceEndpoint() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("service_endpoint"))
}

// ServicePrincipalId returns a reference to field service_principal_id of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) ServicePrincipalId() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("service_principal_id"))
}

// ServicePrincipalKey returns a reference to field service_principal_key of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) ServicePrincipalKey() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("service_principal_key"))
}

// StorageKind returns a reference to field storage_kind of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) StorageKind() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("storage_kind"))
}

// TenantId returns a reference to field tenant_id of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(dflsabs.ref.Append("tenant_id"))
}

// UseManagedIdentity returns a reference to field use_managed_identity of azurerm_data_factory_linked_service_azure_blob_storage.
func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) UseManagedIdentity() terra.BoolValue {
	return terra.ReferenceAsBool(dflsabs.ref.Append("use_managed_identity"))
}

func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) KeyVaultSasToken() terra.ListValue[datafactorylinkedserviceazureblobstorage.KeyVaultSasTokenAttributes] {
	return terra.ReferenceAsList[datafactorylinkedserviceazureblobstorage.KeyVaultSasTokenAttributes](dflsabs.ref.Append("key_vault_sas_token"))
}

func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) ServicePrincipalLinkedKeyVaultKey() terra.ListValue[datafactorylinkedserviceazureblobstorage.ServicePrincipalLinkedKeyVaultKeyAttributes] {
	return terra.ReferenceAsList[datafactorylinkedserviceazureblobstorage.ServicePrincipalLinkedKeyVaultKeyAttributes](dflsabs.ref.Append("service_principal_linked_key_vault_key"))
}

func (dflsabs dataFactoryLinkedServiceAzureBlobStorageAttributes) Timeouts() datafactorylinkedserviceazureblobstorage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedserviceazureblobstorage.TimeoutsAttributes](dflsabs.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceAzureBlobStorageState struct {
	AdditionalProperties              map[string]string                                                                 `json:"additional_properties"`
	Annotations                       []string                                                                          `json:"annotations"`
	ConnectionString                  string                                                                            `json:"connection_string"`
	ConnectionStringInsecure          string                                                                            `json:"connection_string_insecure"`
	DataFactoryId                     string                                                                            `json:"data_factory_id"`
	Description                       string                                                                            `json:"description"`
	Id                                string                                                                            `json:"id"`
	IntegrationRuntimeName            string                                                                            `json:"integration_runtime_name"`
	Name                              string                                                                            `json:"name"`
	Parameters                        map[string]string                                                                 `json:"parameters"`
	SasUri                            string                                                                            `json:"sas_uri"`
	ServiceEndpoint                   string                                                                            `json:"service_endpoint"`
	ServicePrincipalId                string                                                                            `json:"service_principal_id"`
	ServicePrincipalKey               string                                                                            `json:"service_principal_key"`
	StorageKind                       string                                                                            `json:"storage_kind"`
	TenantId                          string                                                                            `json:"tenant_id"`
	UseManagedIdentity                bool                                                                              `json:"use_managed_identity"`
	KeyVaultSasToken                  []datafactorylinkedserviceazureblobstorage.KeyVaultSasTokenState                  `json:"key_vault_sas_token"`
	ServicePrincipalLinkedKeyVaultKey []datafactorylinkedserviceazureblobstorage.ServicePrincipalLinkedKeyVaultKeyState `json:"service_principal_linked_key_vault_key"`
	Timeouts                          *datafactorylinkedserviceazureblobstorage.TimeoutsState                           `json:"timeouts"`
}
