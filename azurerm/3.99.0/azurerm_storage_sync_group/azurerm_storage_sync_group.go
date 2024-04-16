// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_storage_sync_group

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_storage_sync_group.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermStorageSyncGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (assg *Resource) Type() string {
	return "azurerm_storage_sync_group"
}

// LocalName returns the local name for [Resource].
func (assg *Resource) LocalName() string {
	return assg.Name
}

// Configuration returns the configuration (args) for [Resource].
func (assg *Resource) Configuration() interface{} {
	return assg.Args
}

// DependOn is used for other resources to depend on [Resource].
func (assg *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(assg)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (assg *Resource) Dependencies() terra.Dependencies {
	return assg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (assg *Resource) LifecycleManagement() *terra.Lifecycle {
	return assg.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (assg *Resource) Attributes() azurermStorageSyncGroupAttributes {
	return azurermStorageSyncGroupAttributes{ref: terra.ReferenceResource(assg)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (assg *Resource) ImportState(state io.Reader) error {
	assg.state = &azurermStorageSyncGroupState{}
	if err := json.NewDecoder(state).Decode(assg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", assg.Type(), assg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (assg *Resource) State() (*azurermStorageSyncGroupState, bool) {
	return assg.state, assg.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (assg *Resource) StateMust() *azurermStorageSyncGroupState {
	if assg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", assg.Type(), assg.LocalName()))
	}
	return assg.state
}

// Args contains the configurations for azurerm_storage_sync_group.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageSyncId: string, required
	StorageSyncId terra.StringValue `hcl:"storage_sync_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermStorageSyncGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_sync_group.
func (assg azurermStorageSyncGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(assg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_storage_sync_group.
func (assg azurermStorageSyncGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(assg.ref.Append("name"))
}

// StorageSyncId returns a reference to field storage_sync_id of azurerm_storage_sync_group.
func (assg azurermStorageSyncGroupAttributes) StorageSyncId() terra.StringValue {
	return terra.ReferenceAsString(assg.ref.Append("storage_sync_id"))
}

func (assg azurermStorageSyncGroupAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](assg.ref.Append("timeouts"))
}

type azurermStorageSyncGroupState struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	StorageSyncId string         `json:"storage_sync_id"`
	Timeouts      *TimeoutsState `json:"timeouts"`
}
