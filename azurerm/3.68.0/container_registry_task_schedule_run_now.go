// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	containerregistrytaskschedulerunnow "github.com/golingon/terraproviders/azurerm/3.68.0/containerregistrytaskschedulerunnow"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerRegistryTaskScheduleRunNow creates a new instance of [ContainerRegistryTaskScheduleRunNow].
func NewContainerRegistryTaskScheduleRunNow(name string, args ContainerRegistryTaskScheduleRunNowArgs) *ContainerRegistryTaskScheduleRunNow {
	return &ContainerRegistryTaskScheduleRunNow{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerRegistryTaskScheduleRunNow)(nil)

// ContainerRegistryTaskScheduleRunNow represents the Terraform resource azurerm_container_registry_task_schedule_run_now.
type ContainerRegistryTaskScheduleRunNow struct {
	Name      string
	Args      ContainerRegistryTaskScheduleRunNowArgs
	state     *containerRegistryTaskScheduleRunNowState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerRegistryTaskScheduleRunNow].
func (crtsrn *ContainerRegistryTaskScheduleRunNow) Type() string {
	return "azurerm_container_registry_task_schedule_run_now"
}

// LocalName returns the local name for [ContainerRegistryTaskScheduleRunNow].
func (crtsrn *ContainerRegistryTaskScheduleRunNow) LocalName() string {
	return crtsrn.Name
}

// Configuration returns the configuration (args) for [ContainerRegistryTaskScheduleRunNow].
func (crtsrn *ContainerRegistryTaskScheduleRunNow) Configuration() interface{} {
	return crtsrn.Args
}

// DependOn is used for other resources to depend on [ContainerRegistryTaskScheduleRunNow].
func (crtsrn *ContainerRegistryTaskScheduleRunNow) DependOn() terra.Reference {
	return terra.ReferenceResource(crtsrn)
}

// Dependencies returns the list of resources [ContainerRegistryTaskScheduleRunNow] depends_on.
func (crtsrn *ContainerRegistryTaskScheduleRunNow) Dependencies() terra.Dependencies {
	return crtsrn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerRegistryTaskScheduleRunNow].
func (crtsrn *ContainerRegistryTaskScheduleRunNow) LifecycleManagement() *terra.Lifecycle {
	return crtsrn.Lifecycle
}

// Attributes returns the attributes for [ContainerRegistryTaskScheduleRunNow].
func (crtsrn *ContainerRegistryTaskScheduleRunNow) Attributes() containerRegistryTaskScheduleRunNowAttributes {
	return containerRegistryTaskScheduleRunNowAttributes{ref: terra.ReferenceResource(crtsrn)}
}

// ImportState imports the given attribute values into [ContainerRegistryTaskScheduleRunNow]'s state.
func (crtsrn *ContainerRegistryTaskScheduleRunNow) ImportState(av io.Reader) error {
	crtsrn.state = &containerRegistryTaskScheduleRunNowState{}
	if err := json.NewDecoder(av).Decode(crtsrn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crtsrn.Type(), crtsrn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerRegistryTaskScheduleRunNow] has state.
func (crtsrn *ContainerRegistryTaskScheduleRunNow) State() (*containerRegistryTaskScheduleRunNowState, bool) {
	return crtsrn.state, crtsrn.state != nil
}

// StateMust returns the state for [ContainerRegistryTaskScheduleRunNow]. Panics if the state is nil.
func (crtsrn *ContainerRegistryTaskScheduleRunNow) StateMust() *containerRegistryTaskScheduleRunNowState {
	if crtsrn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crtsrn.Type(), crtsrn.LocalName()))
	}
	return crtsrn.state
}

// ContainerRegistryTaskScheduleRunNowArgs contains the configurations for azurerm_container_registry_task_schedule_run_now.
type ContainerRegistryTaskScheduleRunNowArgs struct {
	// ContainerRegistryTaskId: string, required
	ContainerRegistryTaskId terra.StringValue `hcl:"container_registry_task_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *containerregistrytaskschedulerunnow.Timeouts `hcl:"timeouts,block"`
}
type containerRegistryTaskScheduleRunNowAttributes struct {
	ref terra.Reference
}

// ContainerRegistryTaskId returns a reference to field container_registry_task_id of azurerm_container_registry_task_schedule_run_now.
func (crtsrn containerRegistryTaskScheduleRunNowAttributes) ContainerRegistryTaskId() terra.StringValue {
	return terra.ReferenceAsString(crtsrn.ref.Append("container_registry_task_id"))
}

// Id returns a reference to field id of azurerm_container_registry_task_schedule_run_now.
func (crtsrn containerRegistryTaskScheduleRunNowAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crtsrn.ref.Append("id"))
}

func (crtsrn containerRegistryTaskScheduleRunNowAttributes) Timeouts() containerregistrytaskschedulerunnow.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerregistrytaskschedulerunnow.TimeoutsAttributes](crtsrn.ref.Append("timeouts"))
}

type containerRegistryTaskScheduleRunNowState struct {
	ContainerRegistryTaskId string                                             `json:"container_registry_task_id"`
	Id                      string                                             `json:"id"`
	Timeouts                *containerregistrytaskschedulerunnow.TimeoutsState `json:"timeouts"`
}
