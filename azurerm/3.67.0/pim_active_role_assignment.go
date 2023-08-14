// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	pimactiveroleassignment "github.com/golingon/terraproviders/azurerm/3.67.0/pimactiveroleassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPimActiveRoleAssignment creates a new instance of [PimActiveRoleAssignment].
func NewPimActiveRoleAssignment(name string, args PimActiveRoleAssignmentArgs) *PimActiveRoleAssignment {
	return &PimActiveRoleAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PimActiveRoleAssignment)(nil)

// PimActiveRoleAssignment represents the Terraform resource azurerm_pim_active_role_assignment.
type PimActiveRoleAssignment struct {
	Name      string
	Args      PimActiveRoleAssignmentArgs
	state     *pimActiveRoleAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PimActiveRoleAssignment].
func (para *PimActiveRoleAssignment) Type() string {
	return "azurerm_pim_active_role_assignment"
}

// LocalName returns the local name for [PimActiveRoleAssignment].
func (para *PimActiveRoleAssignment) LocalName() string {
	return para.Name
}

// Configuration returns the configuration (args) for [PimActiveRoleAssignment].
func (para *PimActiveRoleAssignment) Configuration() interface{} {
	return para.Args
}

// DependOn is used for other resources to depend on [PimActiveRoleAssignment].
func (para *PimActiveRoleAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(para)
}

// Dependencies returns the list of resources [PimActiveRoleAssignment] depends_on.
func (para *PimActiveRoleAssignment) Dependencies() terra.Dependencies {
	return para.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PimActiveRoleAssignment].
func (para *PimActiveRoleAssignment) LifecycleManagement() *terra.Lifecycle {
	return para.Lifecycle
}

// Attributes returns the attributes for [PimActiveRoleAssignment].
func (para *PimActiveRoleAssignment) Attributes() pimActiveRoleAssignmentAttributes {
	return pimActiveRoleAssignmentAttributes{ref: terra.ReferenceResource(para)}
}

// ImportState imports the given attribute values into [PimActiveRoleAssignment]'s state.
func (para *PimActiveRoleAssignment) ImportState(av io.Reader) error {
	para.state = &pimActiveRoleAssignmentState{}
	if err := json.NewDecoder(av).Decode(para.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", para.Type(), para.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PimActiveRoleAssignment] has state.
func (para *PimActiveRoleAssignment) State() (*pimActiveRoleAssignmentState, bool) {
	return para.state, para.state != nil
}

// StateMust returns the state for [PimActiveRoleAssignment]. Panics if the state is nil.
func (para *PimActiveRoleAssignment) StateMust() *pimActiveRoleAssignmentState {
	if para.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", para.Type(), para.LocalName()))
	}
	return para.state
}

// PimActiveRoleAssignmentArgs contains the configurations for azurerm_pim_active_role_assignment.
type PimActiveRoleAssignmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Justification: string, optional
	Justification terra.StringValue `hcl:"justification,attr"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
	// RoleDefinitionId: string, required
	RoleDefinitionId terra.StringValue `hcl:"role_definition_id,attr" validate:"required"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// Schedule: optional
	Schedule *pimactiveroleassignment.Schedule `hcl:"schedule,block"`
	// Ticket: optional
	Ticket *pimactiveroleassignment.Ticket `hcl:"ticket,block"`
	// Timeouts: optional
	Timeouts *pimactiveroleassignment.Timeouts `hcl:"timeouts,block"`
}
type pimActiveRoleAssignmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_pim_active_role_assignment.
func (para pimActiveRoleAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(para.ref.Append("id"))
}

// Justification returns a reference to field justification of azurerm_pim_active_role_assignment.
func (para pimActiveRoleAssignmentAttributes) Justification() terra.StringValue {
	return terra.ReferenceAsString(para.ref.Append("justification"))
}

// PrincipalId returns a reference to field principal_id of azurerm_pim_active_role_assignment.
func (para pimActiveRoleAssignmentAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(para.ref.Append("principal_id"))
}

// PrincipalType returns a reference to field principal_type of azurerm_pim_active_role_assignment.
func (para pimActiveRoleAssignmentAttributes) PrincipalType() terra.StringValue {
	return terra.ReferenceAsString(para.ref.Append("principal_type"))
}

// RoleDefinitionId returns a reference to field role_definition_id of azurerm_pim_active_role_assignment.
func (para pimActiveRoleAssignmentAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(para.ref.Append("role_definition_id"))
}

// Scope returns a reference to field scope of azurerm_pim_active_role_assignment.
func (para pimActiveRoleAssignmentAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(para.ref.Append("scope"))
}

func (para pimActiveRoleAssignmentAttributes) Schedule() terra.ListValue[pimactiveroleassignment.ScheduleAttributes] {
	return terra.ReferenceAsList[pimactiveroleassignment.ScheduleAttributes](para.ref.Append("schedule"))
}

func (para pimActiveRoleAssignmentAttributes) Ticket() terra.ListValue[pimactiveroleassignment.TicketAttributes] {
	return terra.ReferenceAsList[pimactiveroleassignment.TicketAttributes](para.ref.Append("ticket"))
}

func (para pimActiveRoleAssignmentAttributes) Timeouts() pimactiveroleassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[pimactiveroleassignment.TimeoutsAttributes](para.ref.Append("timeouts"))
}

type pimActiveRoleAssignmentState struct {
	Id               string                                  `json:"id"`
	Justification    string                                  `json:"justification"`
	PrincipalId      string                                  `json:"principal_id"`
	PrincipalType    string                                  `json:"principal_type"`
	RoleDefinitionId string                                  `json:"role_definition_id"`
	Scope            string                                  `json:"scope"`
	Schedule         []pimactiveroleassignment.ScheduleState `json:"schedule"`
	Ticket           []pimactiveroleassignment.TicketState   `json:"ticket"`
	Timeouts         *pimactiveroleassignment.TimeoutsState  `json:"timeouts"`
}