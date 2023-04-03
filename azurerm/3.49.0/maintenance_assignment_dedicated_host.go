// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	maintenanceassignmentdedicatedhost "github.com/golingon/terraproviders/azurerm/3.49.0/maintenanceassignmentdedicatedhost"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMaintenanceAssignmentDedicatedHost creates a new instance of [MaintenanceAssignmentDedicatedHost].
func NewMaintenanceAssignmentDedicatedHost(name string, args MaintenanceAssignmentDedicatedHostArgs) *MaintenanceAssignmentDedicatedHost {
	return &MaintenanceAssignmentDedicatedHost{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MaintenanceAssignmentDedicatedHost)(nil)

// MaintenanceAssignmentDedicatedHost represents the Terraform resource azurerm_maintenance_assignment_dedicated_host.
type MaintenanceAssignmentDedicatedHost struct {
	Name      string
	Args      MaintenanceAssignmentDedicatedHostArgs
	state     *maintenanceAssignmentDedicatedHostState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MaintenanceAssignmentDedicatedHost].
func (madh *MaintenanceAssignmentDedicatedHost) Type() string {
	return "azurerm_maintenance_assignment_dedicated_host"
}

// LocalName returns the local name for [MaintenanceAssignmentDedicatedHost].
func (madh *MaintenanceAssignmentDedicatedHost) LocalName() string {
	return madh.Name
}

// Configuration returns the configuration (args) for [MaintenanceAssignmentDedicatedHost].
func (madh *MaintenanceAssignmentDedicatedHost) Configuration() interface{} {
	return madh.Args
}

// DependOn is used for other resources to depend on [MaintenanceAssignmentDedicatedHost].
func (madh *MaintenanceAssignmentDedicatedHost) DependOn() terra.Reference {
	return terra.ReferenceResource(madh)
}

// Dependencies returns the list of resources [MaintenanceAssignmentDedicatedHost] depends_on.
func (madh *MaintenanceAssignmentDedicatedHost) Dependencies() terra.Dependencies {
	return madh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MaintenanceAssignmentDedicatedHost].
func (madh *MaintenanceAssignmentDedicatedHost) LifecycleManagement() *terra.Lifecycle {
	return madh.Lifecycle
}

// Attributes returns the attributes for [MaintenanceAssignmentDedicatedHost].
func (madh *MaintenanceAssignmentDedicatedHost) Attributes() maintenanceAssignmentDedicatedHostAttributes {
	return maintenanceAssignmentDedicatedHostAttributes{ref: terra.ReferenceResource(madh)}
}

// ImportState imports the given attribute values into [MaintenanceAssignmentDedicatedHost]'s state.
func (madh *MaintenanceAssignmentDedicatedHost) ImportState(av io.Reader) error {
	madh.state = &maintenanceAssignmentDedicatedHostState{}
	if err := json.NewDecoder(av).Decode(madh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", madh.Type(), madh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MaintenanceAssignmentDedicatedHost] has state.
func (madh *MaintenanceAssignmentDedicatedHost) State() (*maintenanceAssignmentDedicatedHostState, bool) {
	return madh.state, madh.state != nil
}

// StateMust returns the state for [MaintenanceAssignmentDedicatedHost]. Panics if the state is nil.
func (madh *MaintenanceAssignmentDedicatedHost) StateMust() *maintenanceAssignmentDedicatedHostState {
	if madh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", madh.Type(), madh.LocalName()))
	}
	return madh.state
}

// MaintenanceAssignmentDedicatedHostArgs contains the configurations for azurerm_maintenance_assignment_dedicated_host.
type MaintenanceAssignmentDedicatedHostArgs struct {
	// DedicatedHostId: string, required
	DedicatedHostId terra.StringValue `hcl:"dedicated_host_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaintenanceConfigurationId: string, required
	MaintenanceConfigurationId terra.StringValue `hcl:"maintenance_configuration_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *maintenanceassignmentdedicatedhost.Timeouts `hcl:"timeouts,block"`
}
type maintenanceAssignmentDedicatedHostAttributes struct {
	ref terra.Reference
}

// DedicatedHostId returns a reference to field dedicated_host_id of azurerm_maintenance_assignment_dedicated_host.
func (madh maintenanceAssignmentDedicatedHostAttributes) DedicatedHostId() terra.StringValue {
	return terra.ReferenceAsString(madh.ref.Append("dedicated_host_id"))
}

// Id returns a reference to field id of azurerm_maintenance_assignment_dedicated_host.
func (madh maintenanceAssignmentDedicatedHostAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(madh.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_maintenance_assignment_dedicated_host.
func (madh maintenanceAssignmentDedicatedHostAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(madh.ref.Append("location"))
}

// MaintenanceConfigurationId returns a reference to field maintenance_configuration_id of azurerm_maintenance_assignment_dedicated_host.
func (madh maintenanceAssignmentDedicatedHostAttributes) MaintenanceConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(madh.ref.Append("maintenance_configuration_id"))
}

func (madh maintenanceAssignmentDedicatedHostAttributes) Timeouts() maintenanceassignmentdedicatedhost.TimeoutsAttributes {
	return terra.ReferenceAsSingle[maintenanceassignmentdedicatedhost.TimeoutsAttributes](madh.ref.Append("timeouts"))
}

type maintenanceAssignmentDedicatedHostState struct {
	DedicatedHostId            string                                            `json:"dedicated_host_id"`
	Id                         string                                            `json:"id"`
	Location                   string                                            `json:"location"`
	MaintenanceConfigurationId string                                            `json:"maintenance_configuration_id"`
	Timeouts                   *maintenanceassignmentdedicatedhost.TimeoutsState `json:"timeouts"`
}
