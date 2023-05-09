// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storagemoversourceendpoint "github.com/golingon/terraproviders/azurerm/3.55.0/storagemoversourceendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageMoverSourceEndpoint creates a new instance of [StorageMoverSourceEndpoint].
func NewStorageMoverSourceEndpoint(name string, args StorageMoverSourceEndpointArgs) *StorageMoverSourceEndpoint {
	return &StorageMoverSourceEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageMoverSourceEndpoint)(nil)

// StorageMoverSourceEndpoint represents the Terraform resource azurerm_storage_mover_source_endpoint.
type StorageMoverSourceEndpoint struct {
	Name      string
	Args      StorageMoverSourceEndpointArgs
	state     *storageMoverSourceEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageMoverSourceEndpoint].
func (smse *StorageMoverSourceEndpoint) Type() string {
	return "azurerm_storage_mover_source_endpoint"
}

// LocalName returns the local name for [StorageMoverSourceEndpoint].
func (smse *StorageMoverSourceEndpoint) LocalName() string {
	return smse.Name
}

// Configuration returns the configuration (args) for [StorageMoverSourceEndpoint].
func (smse *StorageMoverSourceEndpoint) Configuration() interface{} {
	return smse.Args
}

// DependOn is used for other resources to depend on [StorageMoverSourceEndpoint].
func (smse *StorageMoverSourceEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(smse)
}

// Dependencies returns the list of resources [StorageMoverSourceEndpoint] depends_on.
func (smse *StorageMoverSourceEndpoint) Dependencies() terra.Dependencies {
	return smse.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageMoverSourceEndpoint].
func (smse *StorageMoverSourceEndpoint) LifecycleManagement() *terra.Lifecycle {
	return smse.Lifecycle
}

// Attributes returns the attributes for [StorageMoverSourceEndpoint].
func (smse *StorageMoverSourceEndpoint) Attributes() storageMoverSourceEndpointAttributes {
	return storageMoverSourceEndpointAttributes{ref: terra.ReferenceResource(smse)}
}

// ImportState imports the given attribute values into [StorageMoverSourceEndpoint]'s state.
func (smse *StorageMoverSourceEndpoint) ImportState(av io.Reader) error {
	smse.state = &storageMoverSourceEndpointState{}
	if err := json.NewDecoder(av).Decode(smse.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smse.Type(), smse.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageMoverSourceEndpoint] has state.
func (smse *StorageMoverSourceEndpoint) State() (*storageMoverSourceEndpointState, bool) {
	return smse.state, smse.state != nil
}

// StateMust returns the state for [StorageMoverSourceEndpoint]. Panics if the state is nil.
func (smse *StorageMoverSourceEndpoint) StateMust() *storageMoverSourceEndpointState {
	if smse.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smse.Type(), smse.LocalName()))
	}
	return smse.state
}

// StorageMoverSourceEndpointArgs contains the configurations for azurerm_storage_mover_source_endpoint.
type StorageMoverSourceEndpointArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Export: string, optional
	Export terra.StringValue `hcl:"export,attr"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NfsVersion: string, optional
	NfsVersion terra.StringValue `hcl:"nfs_version,attr"`
	// StorageMoverId: string, required
	StorageMoverId terra.StringValue `hcl:"storage_mover_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *storagemoversourceendpoint.Timeouts `hcl:"timeouts,block"`
}
type storageMoverSourceEndpointAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_storage_mover_source_endpoint.
func (smse storageMoverSourceEndpointAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(smse.ref.Append("description"))
}

// Export returns a reference to field export of azurerm_storage_mover_source_endpoint.
func (smse storageMoverSourceEndpointAttributes) Export() terra.StringValue {
	return terra.ReferenceAsString(smse.ref.Append("export"))
}

// Host returns a reference to field host of azurerm_storage_mover_source_endpoint.
func (smse storageMoverSourceEndpointAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(smse.ref.Append("host"))
}

// Id returns a reference to field id of azurerm_storage_mover_source_endpoint.
func (smse storageMoverSourceEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smse.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_storage_mover_source_endpoint.
func (smse storageMoverSourceEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smse.ref.Append("name"))
}

// NfsVersion returns a reference to field nfs_version of azurerm_storage_mover_source_endpoint.
func (smse storageMoverSourceEndpointAttributes) NfsVersion() terra.StringValue {
	return terra.ReferenceAsString(smse.ref.Append("nfs_version"))
}

// StorageMoverId returns a reference to field storage_mover_id of azurerm_storage_mover_source_endpoint.
func (smse storageMoverSourceEndpointAttributes) StorageMoverId() terra.StringValue {
	return terra.ReferenceAsString(smse.ref.Append("storage_mover_id"))
}

func (smse storageMoverSourceEndpointAttributes) Timeouts() storagemoversourceendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagemoversourceendpoint.TimeoutsAttributes](smse.ref.Append("timeouts"))
}

type storageMoverSourceEndpointState struct {
	Description    string                                    `json:"description"`
	Export         string                                    `json:"export"`
	Host           string                                    `json:"host"`
	Id             string                                    `json:"id"`
	Name           string                                    `json:"name"`
	NfsVersion     string                                    `json:"nfs_version"`
	StorageMoverId string                                    `json:"storage_mover_id"`
	Timeouts       *storagemoversourceendpoint.TimeoutsState `json:"timeouts"`
}
