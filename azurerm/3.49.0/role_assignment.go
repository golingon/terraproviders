// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	roleassignment "github.com/golingon/terraproviders/azurerm/3.49.0/roleassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoleAssignment creates a new instance of [RoleAssignment].
func NewRoleAssignment(name string, args RoleAssignmentArgs) *RoleAssignment {
	return &RoleAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RoleAssignment)(nil)

// RoleAssignment represents the Terraform resource azurerm_role_assignment.
type RoleAssignment struct {
	Name      string
	Args      RoleAssignmentArgs
	state     *roleAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RoleAssignment].
func (ra *RoleAssignment) Type() string {
	return "azurerm_role_assignment"
}

// LocalName returns the local name for [RoleAssignment].
func (ra *RoleAssignment) LocalName() string {
	return ra.Name
}

// Configuration returns the configuration (args) for [RoleAssignment].
func (ra *RoleAssignment) Configuration() interface{} {
	return ra.Args
}

// DependOn is used for other resources to depend on [RoleAssignment].
func (ra *RoleAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(ra)
}

// Dependencies returns the list of resources [RoleAssignment] depends_on.
func (ra *RoleAssignment) Dependencies() terra.Dependencies {
	return ra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RoleAssignment].
func (ra *RoleAssignment) LifecycleManagement() *terra.Lifecycle {
	return ra.Lifecycle
}

// Attributes returns the attributes for [RoleAssignment].
func (ra *RoleAssignment) Attributes() roleAssignmentAttributes {
	return roleAssignmentAttributes{ref: terra.ReferenceResource(ra)}
}

// ImportState imports the given attribute values into [RoleAssignment]'s state.
func (ra *RoleAssignment) ImportState(av io.Reader) error {
	ra.state = &roleAssignmentState{}
	if err := json.NewDecoder(av).Decode(ra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ra.Type(), ra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RoleAssignment] has state.
func (ra *RoleAssignment) State() (*roleAssignmentState, bool) {
	return ra.state, ra.state != nil
}

// StateMust returns the state for [RoleAssignment]. Panics if the state is nil.
func (ra *RoleAssignment) StateMust() *roleAssignmentState {
	if ra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ra.Type(), ra.LocalName()))
	}
	return ra.state
}

// RoleAssignmentArgs contains the configurations for azurerm_role_assignment.
type RoleAssignmentArgs struct {
	// Condition: string, optional
	Condition terra.StringValue `hcl:"condition,attr"`
	// ConditionVersion: string, optional
	ConditionVersion terra.StringValue `hcl:"condition_version,attr"`
	// DelegatedManagedIdentityResourceId: string, optional
	DelegatedManagedIdentityResourceId terra.StringValue `hcl:"delegated_managed_identity_resource_id,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
	// RoleDefinitionId: string, optional
	RoleDefinitionId terra.StringValue `hcl:"role_definition_id,attr"`
	// RoleDefinitionName: string, optional
	RoleDefinitionName terra.StringValue `hcl:"role_definition_name,attr"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// SkipServicePrincipalAadCheck: bool, optional
	SkipServicePrincipalAadCheck terra.BoolValue `hcl:"skip_service_principal_aad_check,attr"`
	// Timeouts: optional
	Timeouts *roleassignment.Timeouts `hcl:"timeouts,block"`
}
type roleAssignmentAttributes struct {
	ref terra.Reference
}

// Condition returns a reference to field condition of azurerm_role_assignment.
func (ra roleAssignmentAttributes) Condition() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("condition"))
}

// ConditionVersion returns a reference to field condition_version of azurerm_role_assignment.
func (ra roleAssignmentAttributes) ConditionVersion() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("condition_version"))
}

// DelegatedManagedIdentityResourceId returns a reference to field delegated_managed_identity_resource_id of azurerm_role_assignment.
func (ra roleAssignmentAttributes) DelegatedManagedIdentityResourceId() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("delegated_managed_identity_resource_id"))
}

// Description returns a reference to field description of azurerm_role_assignment.
func (ra roleAssignmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_role_assignment.
func (ra roleAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_role_assignment.
func (ra roleAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("name"))
}

// PrincipalId returns a reference to field principal_id of azurerm_role_assignment.
func (ra roleAssignmentAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("principal_id"))
}

// PrincipalType returns a reference to field principal_type of azurerm_role_assignment.
func (ra roleAssignmentAttributes) PrincipalType() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("principal_type"))
}

// RoleDefinitionId returns a reference to field role_definition_id of azurerm_role_assignment.
func (ra roleAssignmentAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("role_definition_id"))
}

// RoleDefinitionName returns a reference to field role_definition_name of azurerm_role_assignment.
func (ra roleAssignmentAttributes) RoleDefinitionName() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("role_definition_name"))
}

// Scope returns a reference to field scope of azurerm_role_assignment.
func (ra roleAssignmentAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("scope"))
}

// SkipServicePrincipalAadCheck returns a reference to field skip_service_principal_aad_check of azurerm_role_assignment.
func (ra roleAssignmentAttributes) SkipServicePrincipalAadCheck() terra.BoolValue {
	return terra.ReferenceAsBool(ra.ref.Append("skip_service_principal_aad_check"))
}

func (ra roleAssignmentAttributes) Timeouts() roleassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[roleassignment.TimeoutsAttributes](ra.ref.Append("timeouts"))
}

type roleAssignmentState struct {
	Condition                          string                        `json:"condition"`
	ConditionVersion                   string                        `json:"condition_version"`
	DelegatedManagedIdentityResourceId string                        `json:"delegated_managed_identity_resource_id"`
	Description                        string                        `json:"description"`
	Id                                 string                        `json:"id"`
	Name                               string                        `json:"name"`
	PrincipalId                        string                        `json:"principal_id"`
	PrincipalType                      string                        `json:"principal_type"`
	RoleDefinitionId                   string                        `json:"role_definition_id"`
	RoleDefinitionName                 string                        `json:"role_definition_name"`
	Scope                              string                        `json:"scope"`
	SkipServicePrincipalAadCheck       bool                          `json:"skip_service_principal_aad_check"`
	Timeouts                           *roleassignment.TimeoutsState `json:"timeouts"`
}
