// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storagecontainer "github.com/golingon/terraproviders/azurerm/3.49.0/storagecontainer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageContainer creates a new instance of [StorageContainer].
func NewStorageContainer(name string, args StorageContainerArgs) *StorageContainer {
	return &StorageContainer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageContainer)(nil)

// StorageContainer represents the Terraform resource azurerm_storage_container.
type StorageContainer struct {
	Name      string
	Args      StorageContainerArgs
	state     *storageContainerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageContainer].
func (sc *StorageContainer) Type() string {
	return "azurerm_storage_container"
}

// LocalName returns the local name for [StorageContainer].
func (sc *StorageContainer) LocalName() string {
	return sc.Name
}

// Configuration returns the configuration (args) for [StorageContainer].
func (sc *StorageContainer) Configuration() interface{} {
	return sc.Args
}

// DependOn is used for other resources to depend on [StorageContainer].
func (sc *StorageContainer) DependOn() terra.Reference {
	return terra.ReferenceResource(sc)
}

// Dependencies returns the list of resources [StorageContainer] depends_on.
func (sc *StorageContainer) Dependencies() terra.Dependencies {
	return sc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageContainer].
func (sc *StorageContainer) LifecycleManagement() *terra.Lifecycle {
	return sc.Lifecycle
}

// Attributes returns the attributes for [StorageContainer].
func (sc *StorageContainer) Attributes() storageContainerAttributes {
	return storageContainerAttributes{ref: terra.ReferenceResource(sc)}
}

// ImportState imports the given attribute values into [StorageContainer]'s state.
func (sc *StorageContainer) ImportState(av io.Reader) error {
	sc.state = &storageContainerState{}
	if err := json.NewDecoder(av).Decode(sc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sc.Type(), sc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageContainer] has state.
func (sc *StorageContainer) State() (*storageContainerState, bool) {
	return sc.state, sc.state != nil
}

// StateMust returns the state for [StorageContainer]. Panics if the state is nil.
func (sc *StorageContainer) StateMust() *storageContainerState {
	if sc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sc.Type(), sc.LocalName()))
	}
	return sc.state
}

// StorageContainerArgs contains the configurations for azurerm_storage_container.
type StorageContainerArgs struct {
	// ContainerAccessType: string, optional
	ContainerAccessType terra.StringValue `hcl:"container_access_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *storagecontainer.Timeouts `hcl:"timeouts,block"`
}
type storageContainerAttributes struct {
	ref terra.Reference
}

// ContainerAccessType returns a reference to field container_access_type of azurerm_storage_container.
func (sc storageContainerAttributes) ContainerAccessType() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("container_access_type"))
}

// HasImmutabilityPolicy returns a reference to field has_immutability_policy of azurerm_storage_container.
func (sc storageContainerAttributes) HasImmutabilityPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("has_immutability_policy"))
}

// HasLegalHold returns a reference to field has_legal_hold of azurerm_storage_container.
func (sc storageContainerAttributes) HasLegalHold() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("has_legal_hold"))
}

// Id returns a reference to field id of azurerm_storage_container.
func (sc storageContainerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_storage_container.
func (sc storageContainerAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sc.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_storage_container.
func (sc storageContainerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("name"))
}

// ResourceManagerId returns a reference to field resource_manager_id of azurerm_storage_container.
func (sc storageContainerAttributes) ResourceManagerId() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("resource_manager_id"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_storage_container.
func (sc storageContainerAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("storage_account_name"))
}

func (sc storageContainerAttributes) Timeouts() storagecontainer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagecontainer.TimeoutsAttributes](sc.ref.Append("timeouts"))
}

type storageContainerState struct {
	ContainerAccessType   string                          `json:"container_access_type"`
	HasImmutabilityPolicy bool                            `json:"has_immutability_policy"`
	HasLegalHold          bool                            `json:"has_legal_hold"`
	Id                    string                          `json:"id"`
	Metadata              map[string]string               `json:"metadata"`
	Name                  string                          `json:"name"`
	ResourceManagerId     string                          `json:"resource_manager_id"`
	StorageAccountName    string                          `json:"storage_account_name"`
	Timeouts              *storagecontainer.TimeoutsState `json:"timeouts"`
}
