// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	maintenanceassignmentvirtualmachinescaleset "github.com/golingon/terraproviders/azurerm/3.65.0/maintenanceassignmentvirtualmachinescaleset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMaintenanceAssignmentVirtualMachineScaleSet creates a new instance of [MaintenanceAssignmentVirtualMachineScaleSet].
func NewMaintenanceAssignmentVirtualMachineScaleSet(name string, args MaintenanceAssignmentVirtualMachineScaleSetArgs) *MaintenanceAssignmentVirtualMachineScaleSet {
	return &MaintenanceAssignmentVirtualMachineScaleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MaintenanceAssignmentVirtualMachineScaleSet)(nil)

// MaintenanceAssignmentVirtualMachineScaleSet represents the Terraform resource azurerm_maintenance_assignment_virtual_machine_scale_set.
type MaintenanceAssignmentVirtualMachineScaleSet struct {
	Name      string
	Args      MaintenanceAssignmentVirtualMachineScaleSetArgs
	state     *maintenanceAssignmentVirtualMachineScaleSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MaintenanceAssignmentVirtualMachineScaleSet].
func (mavmss *MaintenanceAssignmentVirtualMachineScaleSet) Type() string {
	return "azurerm_maintenance_assignment_virtual_machine_scale_set"
}

// LocalName returns the local name for [MaintenanceAssignmentVirtualMachineScaleSet].
func (mavmss *MaintenanceAssignmentVirtualMachineScaleSet) LocalName() string {
	return mavmss.Name
}

// Configuration returns the configuration (args) for [MaintenanceAssignmentVirtualMachineScaleSet].
func (mavmss *MaintenanceAssignmentVirtualMachineScaleSet) Configuration() interface{} {
	return mavmss.Args
}

// DependOn is used for other resources to depend on [MaintenanceAssignmentVirtualMachineScaleSet].
func (mavmss *MaintenanceAssignmentVirtualMachineScaleSet) DependOn() terra.Reference {
	return terra.ReferenceResource(mavmss)
}

// Dependencies returns the list of resources [MaintenanceAssignmentVirtualMachineScaleSet] depends_on.
func (mavmss *MaintenanceAssignmentVirtualMachineScaleSet) Dependencies() terra.Dependencies {
	return mavmss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MaintenanceAssignmentVirtualMachineScaleSet].
func (mavmss *MaintenanceAssignmentVirtualMachineScaleSet) LifecycleManagement() *terra.Lifecycle {
	return mavmss.Lifecycle
}

// Attributes returns the attributes for [MaintenanceAssignmentVirtualMachineScaleSet].
func (mavmss *MaintenanceAssignmentVirtualMachineScaleSet) Attributes() maintenanceAssignmentVirtualMachineScaleSetAttributes {
	return maintenanceAssignmentVirtualMachineScaleSetAttributes{ref: terra.ReferenceResource(mavmss)}
}

// ImportState imports the given attribute values into [MaintenanceAssignmentVirtualMachineScaleSet]'s state.
func (mavmss *MaintenanceAssignmentVirtualMachineScaleSet) ImportState(av io.Reader) error {
	mavmss.state = &maintenanceAssignmentVirtualMachineScaleSetState{}
	if err := json.NewDecoder(av).Decode(mavmss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mavmss.Type(), mavmss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MaintenanceAssignmentVirtualMachineScaleSet] has state.
func (mavmss *MaintenanceAssignmentVirtualMachineScaleSet) State() (*maintenanceAssignmentVirtualMachineScaleSetState, bool) {
	return mavmss.state, mavmss.state != nil
}

// StateMust returns the state for [MaintenanceAssignmentVirtualMachineScaleSet]. Panics if the state is nil.
func (mavmss *MaintenanceAssignmentVirtualMachineScaleSet) StateMust() *maintenanceAssignmentVirtualMachineScaleSetState {
	if mavmss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mavmss.Type(), mavmss.LocalName()))
	}
	return mavmss.state
}

// MaintenanceAssignmentVirtualMachineScaleSetArgs contains the configurations for azurerm_maintenance_assignment_virtual_machine_scale_set.
type MaintenanceAssignmentVirtualMachineScaleSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaintenanceConfigurationId: string, required
	MaintenanceConfigurationId terra.StringValue `hcl:"maintenance_configuration_id,attr" validate:"required"`
	// VirtualMachineScaleSetId: string, required
	VirtualMachineScaleSetId terra.StringValue `hcl:"virtual_machine_scale_set_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *maintenanceassignmentvirtualmachinescaleset.Timeouts `hcl:"timeouts,block"`
}
type maintenanceAssignmentVirtualMachineScaleSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_maintenance_assignment_virtual_machine_scale_set.
func (mavmss maintenanceAssignmentVirtualMachineScaleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mavmss.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_maintenance_assignment_virtual_machine_scale_set.
func (mavmss maintenanceAssignmentVirtualMachineScaleSetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mavmss.ref.Append("location"))
}

// MaintenanceConfigurationId returns a reference to field maintenance_configuration_id of azurerm_maintenance_assignment_virtual_machine_scale_set.
func (mavmss maintenanceAssignmentVirtualMachineScaleSetAttributes) MaintenanceConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(mavmss.ref.Append("maintenance_configuration_id"))
}

// VirtualMachineScaleSetId returns a reference to field virtual_machine_scale_set_id of azurerm_maintenance_assignment_virtual_machine_scale_set.
func (mavmss maintenanceAssignmentVirtualMachineScaleSetAttributes) VirtualMachineScaleSetId() terra.StringValue {
	return terra.ReferenceAsString(mavmss.ref.Append("virtual_machine_scale_set_id"))
}

func (mavmss maintenanceAssignmentVirtualMachineScaleSetAttributes) Timeouts() maintenanceassignmentvirtualmachinescaleset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[maintenanceassignmentvirtualmachinescaleset.TimeoutsAttributes](mavmss.ref.Append("timeouts"))
}

type maintenanceAssignmentVirtualMachineScaleSetState struct {
	Id                         string                                                     `json:"id"`
	Location                   string                                                     `json:"location"`
	MaintenanceConfigurationId string                                                     `json:"maintenance_configuration_id"`
	VirtualMachineScaleSetId   string                                                     `json:"virtual_machine_scale_set_id"`
	Timeouts                   *maintenanceassignmentvirtualmachinescaleset.TimeoutsState `json:"timeouts"`
}