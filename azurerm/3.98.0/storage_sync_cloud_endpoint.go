// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	storagesynccloudendpoint "github.com/golingon/terraproviders/azurerm/3.98.0/storagesynccloudendpoint"
	"io"
)

// NewStorageSyncCloudEndpoint creates a new instance of [StorageSyncCloudEndpoint].
func NewStorageSyncCloudEndpoint(name string, args StorageSyncCloudEndpointArgs) *StorageSyncCloudEndpoint {
	return &StorageSyncCloudEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageSyncCloudEndpoint)(nil)

// StorageSyncCloudEndpoint represents the Terraform resource azurerm_storage_sync_cloud_endpoint.
type StorageSyncCloudEndpoint struct {
	Name      string
	Args      StorageSyncCloudEndpointArgs
	state     *storageSyncCloudEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageSyncCloudEndpoint].
func (ssce *StorageSyncCloudEndpoint) Type() string {
	return "azurerm_storage_sync_cloud_endpoint"
}

// LocalName returns the local name for [StorageSyncCloudEndpoint].
func (ssce *StorageSyncCloudEndpoint) LocalName() string {
	return ssce.Name
}

// Configuration returns the configuration (args) for [StorageSyncCloudEndpoint].
func (ssce *StorageSyncCloudEndpoint) Configuration() interface{} {
	return ssce.Args
}

// DependOn is used for other resources to depend on [StorageSyncCloudEndpoint].
func (ssce *StorageSyncCloudEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(ssce)
}

// Dependencies returns the list of resources [StorageSyncCloudEndpoint] depends_on.
func (ssce *StorageSyncCloudEndpoint) Dependencies() terra.Dependencies {
	return ssce.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageSyncCloudEndpoint].
func (ssce *StorageSyncCloudEndpoint) LifecycleManagement() *terra.Lifecycle {
	return ssce.Lifecycle
}

// Attributes returns the attributes for [StorageSyncCloudEndpoint].
func (ssce *StorageSyncCloudEndpoint) Attributes() storageSyncCloudEndpointAttributes {
	return storageSyncCloudEndpointAttributes{ref: terra.ReferenceResource(ssce)}
}

// ImportState imports the given attribute values into [StorageSyncCloudEndpoint]'s state.
func (ssce *StorageSyncCloudEndpoint) ImportState(av io.Reader) error {
	ssce.state = &storageSyncCloudEndpointState{}
	if err := json.NewDecoder(av).Decode(ssce.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssce.Type(), ssce.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageSyncCloudEndpoint] has state.
func (ssce *StorageSyncCloudEndpoint) State() (*storageSyncCloudEndpointState, bool) {
	return ssce.state, ssce.state != nil
}

// StateMust returns the state for [StorageSyncCloudEndpoint]. Panics if the state is nil.
func (ssce *StorageSyncCloudEndpoint) StateMust() *storageSyncCloudEndpointState {
	if ssce.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssce.Type(), ssce.LocalName()))
	}
	return ssce.state
}

// StorageSyncCloudEndpointArgs contains the configurations for azurerm_storage_sync_cloud_endpoint.
type StorageSyncCloudEndpointArgs struct {
	// FileShareName: string, required
	FileShareName terra.StringValue `hcl:"file_share_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// StorageAccountTenantId: string, optional
	StorageAccountTenantId terra.StringValue `hcl:"storage_account_tenant_id,attr"`
	// StorageSyncGroupId: string, required
	StorageSyncGroupId terra.StringValue `hcl:"storage_sync_group_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *storagesynccloudendpoint.Timeouts `hcl:"timeouts,block"`
}
type storageSyncCloudEndpointAttributes struct {
	ref terra.Reference
}

// FileShareName returns a reference to field file_share_name of azurerm_storage_sync_cloud_endpoint.
func (ssce storageSyncCloudEndpointAttributes) FileShareName() terra.StringValue {
	return terra.ReferenceAsString(ssce.ref.Append("file_share_name"))
}

// Id returns a reference to field id of azurerm_storage_sync_cloud_endpoint.
func (ssce storageSyncCloudEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssce.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_storage_sync_cloud_endpoint.
func (ssce storageSyncCloudEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssce.ref.Append("name"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_storage_sync_cloud_endpoint.
func (ssce storageSyncCloudEndpointAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(ssce.ref.Append("storage_account_id"))
}

// StorageAccountTenantId returns a reference to field storage_account_tenant_id of azurerm_storage_sync_cloud_endpoint.
func (ssce storageSyncCloudEndpointAttributes) StorageAccountTenantId() terra.StringValue {
	return terra.ReferenceAsString(ssce.ref.Append("storage_account_tenant_id"))
}

// StorageSyncGroupId returns a reference to field storage_sync_group_id of azurerm_storage_sync_cloud_endpoint.
func (ssce storageSyncCloudEndpointAttributes) StorageSyncGroupId() terra.StringValue {
	return terra.ReferenceAsString(ssce.ref.Append("storage_sync_group_id"))
}

func (ssce storageSyncCloudEndpointAttributes) Timeouts() storagesynccloudendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagesynccloudendpoint.TimeoutsAttributes](ssce.ref.Append("timeouts"))
}

type storageSyncCloudEndpointState struct {
	FileShareName          string                                  `json:"file_share_name"`
	Id                     string                                  `json:"id"`
	Name                   string                                  `json:"name"`
	StorageAccountId       string                                  `json:"storage_account_id"`
	StorageAccountTenantId string                                  `json:"storage_account_tenant_id"`
	StorageSyncGroupId     string                                  `json:"storage_sync_group_id"`
	Timeouts               *storagesynccloudendpoint.TimeoutsState `json:"timeouts"`
}
