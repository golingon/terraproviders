// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	pimeligibleroleassignment "github.com/golingon/terraproviders/azurerm/3.69.0/pimeligibleroleassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPimEligibleRoleAssignment creates a new instance of [PimEligibleRoleAssignment].
func NewPimEligibleRoleAssignment(name string, args PimEligibleRoleAssignmentArgs) *PimEligibleRoleAssignment {
	return &PimEligibleRoleAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PimEligibleRoleAssignment)(nil)

// PimEligibleRoleAssignment represents the Terraform resource azurerm_pim_eligible_role_assignment.
type PimEligibleRoleAssignment struct {
	Name      string
	Args      PimEligibleRoleAssignmentArgs
	state     *pimEligibleRoleAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PimEligibleRoleAssignment].
func (pera *PimEligibleRoleAssignment) Type() string {
	return "azurerm_pim_eligible_role_assignment"
}

// LocalName returns the local name for [PimEligibleRoleAssignment].
func (pera *PimEligibleRoleAssignment) LocalName() string {
	return pera.Name
}

// Configuration returns the configuration (args) for [PimEligibleRoleAssignment].
func (pera *PimEligibleRoleAssignment) Configuration() interface{} {
	return pera.Args
}

// DependOn is used for other resources to depend on [PimEligibleRoleAssignment].
func (pera *PimEligibleRoleAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(pera)
}

// Dependencies returns the list of resources [PimEligibleRoleAssignment] depends_on.
func (pera *PimEligibleRoleAssignment) Dependencies() terra.Dependencies {
	return pera.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PimEligibleRoleAssignment].
func (pera *PimEligibleRoleAssignment) LifecycleManagement() *terra.Lifecycle {
	return pera.Lifecycle
}

// Attributes returns the attributes for [PimEligibleRoleAssignment].
func (pera *PimEligibleRoleAssignment) Attributes() pimEligibleRoleAssignmentAttributes {
	return pimEligibleRoleAssignmentAttributes{ref: terra.ReferenceResource(pera)}
}

// ImportState imports the given attribute values into [PimEligibleRoleAssignment]'s state.
func (pera *PimEligibleRoleAssignment) ImportState(av io.Reader) error {
	pera.state = &pimEligibleRoleAssignmentState{}
	if err := json.NewDecoder(av).Decode(pera.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pera.Type(), pera.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PimEligibleRoleAssignment] has state.
func (pera *PimEligibleRoleAssignment) State() (*pimEligibleRoleAssignmentState, bool) {
	return pera.state, pera.state != nil
}

// StateMust returns the state for [PimEligibleRoleAssignment]. Panics if the state is nil.
func (pera *PimEligibleRoleAssignment) StateMust() *pimEligibleRoleAssignmentState {
	if pera.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pera.Type(), pera.LocalName()))
	}
	return pera.state
}

// PimEligibleRoleAssignmentArgs contains the configurations for azurerm_pim_eligible_role_assignment.
type PimEligibleRoleAssignmentArgs struct {
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
	Schedule *pimeligibleroleassignment.Schedule `hcl:"schedule,block"`
	// Ticket: optional
	Ticket *pimeligibleroleassignment.Ticket `hcl:"ticket,block"`
	// Timeouts: optional
	Timeouts *pimeligibleroleassignment.Timeouts `hcl:"timeouts,block"`
}
type pimEligibleRoleAssignmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_pim_eligible_role_assignment.
func (pera pimEligibleRoleAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pera.ref.Append("id"))
}

// Justification returns a reference to field justification of azurerm_pim_eligible_role_assignment.
func (pera pimEligibleRoleAssignmentAttributes) Justification() terra.StringValue {
	return terra.ReferenceAsString(pera.ref.Append("justification"))
}

// PrincipalId returns a reference to field principal_id of azurerm_pim_eligible_role_assignment.
func (pera pimEligibleRoleAssignmentAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(pera.ref.Append("principal_id"))
}

// PrincipalType returns a reference to field principal_type of azurerm_pim_eligible_role_assignment.
func (pera pimEligibleRoleAssignmentAttributes) PrincipalType() terra.StringValue {
	return terra.ReferenceAsString(pera.ref.Append("principal_type"))
}

// RoleDefinitionId returns a reference to field role_definition_id of azurerm_pim_eligible_role_assignment.
func (pera pimEligibleRoleAssignmentAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(pera.ref.Append("role_definition_id"))
}

// Scope returns a reference to field scope of azurerm_pim_eligible_role_assignment.
func (pera pimEligibleRoleAssignmentAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(pera.ref.Append("scope"))
}

func (pera pimEligibleRoleAssignmentAttributes) Schedule() terra.ListValue[pimeligibleroleassignment.ScheduleAttributes] {
	return terra.ReferenceAsList[pimeligibleroleassignment.ScheduleAttributes](pera.ref.Append("schedule"))
}

func (pera pimEligibleRoleAssignmentAttributes) Ticket() terra.ListValue[pimeligibleroleassignment.TicketAttributes] {
	return terra.ReferenceAsList[pimeligibleroleassignment.TicketAttributes](pera.ref.Append("ticket"))
}

func (pera pimEligibleRoleAssignmentAttributes) Timeouts() pimeligibleroleassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[pimeligibleroleassignment.TimeoutsAttributes](pera.ref.Append("timeouts"))
}

type pimEligibleRoleAssignmentState struct {
	Id               string                                    `json:"id"`
	Justification    string                                    `json:"justification"`
	PrincipalId      string                                    `json:"principal_id"`
	PrincipalType    string                                    `json:"principal_type"`
	RoleDefinitionId string                                    `json:"role_definition_id"`
	Scope            string                                    `json:"scope"`
	Schedule         []pimeligibleroleassignment.ScheduleState `json:"schedule"`
	Ticket           []pimeligibleroleassignment.TicketState   `json:"ticket"`
	Timeouts         *pimeligibleroleassignment.TimeoutsState  `json:"timeouts"`
}
