// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storagemovertargetendpoint "github.com/golingon/terraproviders/azurerm/3.65.0/storagemovertargetendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageMoverTargetEndpoint creates a new instance of [StorageMoverTargetEndpoint].
func NewStorageMoverTargetEndpoint(name string, args StorageMoverTargetEndpointArgs) *StorageMoverTargetEndpoint {
	return &StorageMoverTargetEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageMoverTargetEndpoint)(nil)

// StorageMoverTargetEndpoint represents the Terraform resource azurerm_storage_mover_target_endpoint.
type StorageMoverTargetEndpoint struct {
	Name      string
	Args      StorageMoverTargetEndpointArgs
	state     *storageMoverTargetEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageMoverTargetEndpoint].
func (smte *StorageMoverTargetEndpoint) Type() string {
	return "azurerm_storage_mover_target_endpoint"
}

// LocalName returns the local name for [StorageMoverTargetEndpoint].
func (smte *StorageMoverTargetEndpoint) LocalName() string {
	return smte.Name
}

// Configuration returns the configuration (args) for [StorageMoverTargetEndpoint].
func (smte *StorageMoverTargetEndpoint) Configuration() interface{} {
	return smte.Args
}

// DependOn is used for other resources to depend on [StorageMoverTargetEndpoint].
func (smte *StorageMoverTargetEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(smte)
}

// Dependencies returns the list of resources [StorageMoverTargetEndpoint] depends_on.
func (smte *StorageMoverTargetEndpoint) Dependencies() terra.Dependencies {
	return smte.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageMoverTargetEndpoint].
func (smte *StorageMoverTargetEndpoint) LifecycleManagement() *terra.Lifecycle {
	return smte.Lifecycle
}

// Attributes returns the attributes for [StorageMoverTargetEndpoint].
func (smte *StorageMoverTargetEndpoint) Attributes() storageMoverTargetEndpointAttributes {
	return storageMoverTargetEndpointAttributes{ref: terra.ReferenceResource(smte)}
}

// ImportState imports the given attribute values into [StorageMoverTargetEndpoint]'s state.
func (smte *StorageMoverTargetEndpoint) ImportState(av io.Reader) error {
	smte.state = &storageMoverTargetEndpointState{}
	if err := json.NewDecoder(av).Decode(smte.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smte.Type(), smte.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageMoverTargetEndpoint] has state.
func (smte *StorageMoverTargetEndpoint) State() (*storageMoverTargetEndpointState, bool) {
	return smte.state, smte.state != nil
}

// StateMust returns the state for [StorageMoverTargetEndpoint]. Panics if the state is nil.
func (smte *StorageMoverTargetEndpoint) StateMust() *storageMoverTargetEndpointState {
	if smte.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smte.Type(), smte.LocalName()))
	}
	return smte.state
}

// StorageMoverTargetEndpointArgs contains the configurations for azurerm_storage_mover_target_endpoint.
type StorageMoverTargetEndpointArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// StorageContainerName: string, required
	StorageContainerName terra.StringValue `hcl:"storage_container_name,attr" validate:"required"`
	// StorageMoverId: string, required
	StorageMoverId terra.StringValue `hcl:"storage_mover_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *storagemovertargetendpoint.Timeouts `hcl:"timeouts,block"`
}
type storageMoverTargetEndpointAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_storage_mover_target_endpoint.
func (smte storageMoverTargetEndpointAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(smte.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_storage_mover_target_endpoint.
func (smte storageMoverTargetEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smte.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_storage_mover_target_endpoint.
func (smte storageMoverTargetEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smte.ref.Append("name"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_storage_mover_target_endpoint.
func (smte storageMoverTargetEndpointAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(smte.ref.Append("storage_account_id"))
}

// StorageContainerName returns a reference to field storage_container_name of azurerm_storage_mover_target_endpoint.
func (smte storageMoverTargetEndpointAttributes) StorageContainerName() terra.StringValue {
	return terra.ReferenceAsString(smte.ref.Append("storage_container_name"))
}

// StorageMoverId returns a reference to field storage_mover_id of azurerm_storage_mover_target_endpoint.
func (smte storageMoverTargetEndpointAttributes) StorageMoverId() terra.StringValue {
	return terra.ReferenceAsString(smte.ref.Append("storage_mover_id"))
}

func (smte storageMoverTargetEndpointAttributes) Timeouts() storagemovertargetendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagemovertargetendpoint.TimeoutsAttributes](smte.ref.Append("timeouts"))
}

type storageMoverTargetEndpointState struct {
	Description          string                                    `json:"description"`
	Id                   string                                    `json:"id"`
	Name                 string                                    `json:"name"`
	StorageAccountId     string                                    `json:"storage_account_id"`
	StorageContainerName string                                    `json:"storage_container_name"`
	StorageMoverId       string                                    `json:"storage_mover_id"`
	Timeouts             *storagemovertargetendpoint.TimeoutsState `json:"timeouts"`
}
