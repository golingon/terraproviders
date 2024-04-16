// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread_directory_role_eligibility_schedule_request

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

// Resource represents the Terraform resource azuread_directory_role_eligibility_schedule_request.
type Resource struct {
	Name      string
	Args      Args
	state     *azureadDirectoryRoleEligibilityScheduleRequestState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adresr *Resource) Type() string {
	return "azuread_directory_role_eligibility_schedule_request"
}

// LocalName returns the local name for [Resource].
func (adresr *Resource) LocalName() string {
	return adresr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adresr *Resource) Configuration() interface{} {
	return adresr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adresr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adresr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adresr *Resource) Dependencies() terra.Dependencies {
	return adresr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adresr *Resource) LifecycleManagement() *terra.Lifecycle {
	return adresr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adresr *Resource) Attributes() azureadDirectoryRoleEligibilityScheduleRequestAttributes {
	return azureadDirectoryRoleEligibilityScheduleRequestAttributes{ref: terra.ReferenceResource(adresr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adresr *Resource) ImportState(state io.Reader) error {
	adresr.state = &azureadDirectoryRoleEligibilityScheduleRequestState{}
	if err := json.NewDecoder(state).Decode(adresr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adresr.Type(), adresr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adresr *Resource) State() (*azureadDirectoryRoleEligibilityScheduleRequestState, bool) {
	return adresr.state, adresr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adresr *Resource) StateMust() *azureadDirectoryRoleEligibilityScheduleRequestState {
	if adresr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adresr.Type(), adresr.LocalName()))
	}
	return adresr.state
}

// Args contains the configurations for azuread_directory_role_eligibility_schedule_request.
type Args struct {
	// DirectoryScopeId: string, required
	DirectoryScopeId terra.StringValue `hcl:"directory_scope_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Justification: string, required
	Justification terra.StringValue `hcl:"justification,attr" validate:"required"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
	// RoleDefinitionId: string, required
	RoleDefinitionId terra.StringValue `hcl:"role_definition_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azureadDirectoryRoleEligibilityScheduleRequestAttributes struct {
	ref terra.Reference
}

// DirectoryScopeId returns a reference to field directory_scope_id of azuread_directory_role_eligibility_schedule_request.
func (adresr azureadDirectoryRoleEligibilityScheduleRequestAttributes) DirectoryScopeId() terra.StringValue {
	return terra.ReferenceAsString(adresr.ref.Append("directory_scope_id"))
}

// Id returns a reference to field id of azuread_directory_role_eligibility_schedule_request.
func (adresr azureadDirectoryRoleEligibilityScheduleRequestAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adresr.ref.Append("id"))
}

// Justification returns a reference to field justification of azuread_directory_role_eligibility_schedule_request.
func (adresr azureadDirectoryRoleEligibilityScheduleRequestAttributes) Justification() terra.StringValue {
	return terra.ReferenceAsString(adresr.ref.Append("justification"))
}

// PrincipalId returns a reference to field principal_id of azuread_directory_role_eligibility_schedule_request.
func (adresr azureadDirectoryRoleEligibilityScheduleRequestAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(adresr.ref.Append("principal_id"))
}

// RoleDefinitionId returns a reference to field role_definition_id of azuread_directory_role_eligibility_schedule_request.
func (adresr azureadDirectoryRoleEligibilityScheduleRequestAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(adresr.ref.Append("role_definition_id"))
}

func (adresr azureadDirectoryRoleEligibilityScheduleRequestAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adresr.ref.Append("timeouts"))
}

type azureadDirectoryRoleEligibilityScheduleRequestState struct {
	DirectoryScopeId string         `json:"directory_scope_id"`
	Id               string         `json:"id"`
	Justification    string         `json:"justification"`
	PrincipalId      string         `json:"principal_id"`
	RoleDefinitionId string         `json:"role_definition_id"`
	Timeouts         *TimeoutsState `json:"timeouts"`
}
