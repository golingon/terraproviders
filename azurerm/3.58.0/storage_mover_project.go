// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storagemoverproject "github.com/golingon/terraproviders/azurerm/3.58.0/storagemoverproject"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageMoverProject creates a new instance of [StorageMoverProject].
func NewStorageMoverProject(name string, args StorageMoverProjectArgs) *StorageMoverProject {
	return &StorageMoverProject{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageMoverProject)(nil)

// StorageMoverProject represents the Terraform resource azurerm_storage_mover_project.
type StorageMoverProject struct {
	Name      string
	Args      StorageMoverProjectArgs
	state     *storageMoverProjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageMoverProject].
func (smp *StorageMoverProject) Type() string {
	return "azurerm_storage_mover_project"
}

// LocalName returns the local name for [StorageMoverProject].
func (smp *StorageMoverProject) LocalName() string {
	return smp.Name
}

// Configuration returns the configuration (args) for [StorageMoverProject].
func (smp *StorageMoverProject) Configuration() interface{} {
	return smp.Args
}

// DependOn is used for other resources to depend on [StorageMoverProject].
func (smp *StorageMoverProject) DependOn() terra.Reference {
	return terra.ReferenceResource(smp)
}

// Dependencies returns the list of resources [StorageMoverProject] depends_on.
func (smp *StorageMoverProject) Dependencies() terra.Dependencies {
	return smp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageMoverProject].
func (smp *StorageMoverProject) LifecycleManagement() *terra.Lifecycle {
	return smp.Lifecycle
}

// Attributes returns the attributes for [StorageMoverProject].
func (smp *StorageMoverProject) Attributes() storageMoverProjectAttributes {
	return storageMoverProjectAttributes{ref: terra.ReferenceResource(smp)}
}

// ImportState imports the given attribute values into [StorageMoverProject]'s state.
func (smp *StorageMoverProject) ImportState(av io.Reader) error {
	smp.state = &storageMoverProjectState{}
	if err := json.NewDecoder(av).Decode(smp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smp.Type(), smp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageMoverProject] has state.
func (smp *StorageMoverProject) State() (*storageMoverProjectState, bool) {
	return smp.state, smp.state != nil
}

// StateMust returns the state for [StorageMoverProject]. Panics if the state is nil.
func (smp *StorageMoverProject) StateMust() *storageMoverProjectState {
	if smp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smp.Type(), smp.LocalName()))
	}
	return smp.state
}

// StorageMoverProjectArgs contains the configurations for azurerm_storage_mover_project.
type StorageMoverProjectArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageMoverId: string, required
	StorageMoverId terra.StringValue `hcl:"storage_mover_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *storagemoverproject.Timeouts `hcl:"timeouts,block"`
}
type storageMoverProjectAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_storage_mover_project.
func (smp storageMoverProjectAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(smp.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_storage_mover_project.
func (smp storageMoverProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_storage_mover_project.
func (smp storageMoverProjectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smp.ref.Append("name"))
}

// StorageMoverId returns a reference to field storage_mover_id of azurerm_storage_mover_project.
func (smp storageMoverProjectAttributes) StorageMoverId() terra.StringValue {
	return terra.ReferenceAsString(smp.ref.Append("storage_mover_id"))
}

func (smp storageMoverProjectAttributes) Timeouts() storagemoverproject.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagemoverproject.TimeoutsAttributes](smp.ref.Append("timeouts"))
}

type storageMoverProjectState struct {
	Description    string                             `json:"description"`
	Id             string                             `json:"id"`
	Name           string                             `json:"name"`
	StorageMoverId string                             `json:"storage_mover_id"`
	Timeouts       *storagemoverproject.TimeoutsState `json:"timeouts"`
}
