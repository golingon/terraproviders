// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storagemover "github.com/golingon/terraproviders/azurerm/3.66.0/storagemover"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageMover creates a new instance of [StorageMover].
func NewStorageMover(name string, args StorageMoverArgs) *StorageMover {
	return &StorageMover{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageMover)(nil)

// StorageMover represents the Terraform resource azurerm_storage_mover.
type StorageMover struct {
	Name      string
	Args      StorageMoverArgs
	state     *storageMoverState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageMover].
func (sm *StorageMover) Type() string {
	return "azurerm_storage_mover"
}

// LocalName returns the local name for [StorageMover].
func (sm *StorageMover) LocalName() string {
	return sm.Name
}

// Configuration returns the configuration (args) for [StorageMover].
func (sm *StorageMover) Configuration() interface{} {
	return sm.Args
}

// DependOn is used for other resources to depend on [StorageMover].
func (sm *StorageMover) DependOn() terra.Reference {
	return terra.ReferenceResource(sm)
}

// Dependencies returns the list of resources [StorageMover] depends_on.
func (sm *StorageMover) Dependencies() terra.Dependencies {
	return sm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageMover].
func (sm *StorageMover) LifecycleManagement() *terra.Lifecycle {
	return sm.Lifecycle
}

// Attributes returns the attributes for [StorageMover].
func (sm *StorageMover) Attributes() storageMoverAttributes {
	return storageMoverAttributes{ref: terra.ReferenceResource(sm)}
}

// ImportState imports the given attribute values into [StorageMover]'s state.
func (sm *StorageMover) ImportState(av io.Reader) error {
	sm.state = &storageMoverState{}
	if err := json.NewDecoder(av).Decode(sm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sm.Type(), sm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageMover] has state.
func (sm *StorageMover) State() (*storageMoverState, bool) {
	return sm.state, sm.state != nil
}

// StateMust returns the state for [StorageMover]. Panics if the state is nil.
func (sm *StorageMover) StateMust() *storageMoverState {
	if sm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sm.Type(), sm.LocalName()))
	}
	return sm.state
}

// StorageMoverArgs contains the configurations for azurerm_storage_mover.
type StorageMoverArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *storagemover.Timeouts `hcl:"timeouts,block"`
}
type storageMoverAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_storage_mover.
func (sm storageMoverAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_storage_mover.
func (sm storageMoverAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_storage_mover.
func (sm storageMoverAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_storage_mover.
func (sm storageMoverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_storage_mover.
func (sm storageMoverAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_storage_mover.
func (sm storageMoverAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sm.ref.Append("tags"))
}

func (sm storageMoverAttributes) Timeouts() storagemover.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagemover.TimeoutsAttributes](sm.ref.Append("timeouts"))
}

type storageMoverState struct {
	Description       string                      `json:"description"`
	Id                string                      `json:"id"`
	Location          string                      `json:"location"`
	Name              string                      `json:"name"`
	ResourceGroupName string                      `json:"resource_group_name"`
	Tags              map[string]string           `json:"tags"`
	Timeouts          *storagemover.TimeoutsState `json:"timeouts"`
}