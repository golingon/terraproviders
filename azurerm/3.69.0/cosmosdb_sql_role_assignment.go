// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbsqlroleassignment "github.com/golingon/terraproviders/azurerm/3.69.0/cosmosdbsqlroleassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbSqlRoleAssignment creates a new instance of [CosmosdbSqlRoleAssignment].
func NewCosmosdbSqlRoleAssignment(name string, args CosmosdbSqlRoleAssignmentArgs) *CosmosdbSqlRoleAssignment {
	return &CosmosdbSqlRoleAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbSqlRoleAssignment)(nil)

// CosmosdbSqlRoleAssignment represents the Terraform resource azurerm_cosmosdb_sql_role_assignment.
type CosmosdbSqlRoleAssignment struct {
	Name      string
	Args      CosmosdbSqlRoleAssignmentArgs
	state     *cosmosdbSqlRoleAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbSqlRoleAssignment].
func (csra *CosmosdbSqlRoleAssignment) Type() string {
	return "azurerm_cosmosdb_sql_role_assignment"
}

// LocalName returns the local name for [CosmosdbSqlRoleAssignment].
func (csra *CosmosdbSqlRoleAssignment) LocalName() string {
	return csra.Name
}

// Configuration returns the configuration (args) for [CosmosdbSqlRoleAssignment].
func (csra *CosmosdbSqlRoleAssignment) Configuration() interface{} {
	return csra.Args
}

// DependOn is used for other resources to depend on [CosmosdbSqlRoleAssignment].
func (csra *CosmosdbSqlRoleAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(csra)
}

// Dependencies returns the list of resources [CosmosdbSqlRoleAssignment] depends_on.
func (csra *CosmosdbSqlRoleAssignment) Dependencies() terra.Dependencies {
	return csra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbSqlRoleAssignment].
func (csra *CosmosdbSqlRoleAssignment) LifecycleManagement() *terra.Lifecycle {
	return csra.Lifecycle
}

// Attributes returns the attributes for [CosmosdbSqlRoleAssignment].
func (csra *CosmosdbSqlRoleAssignment) Attributes() cosmosdbSqlRoleAssignmentAttributes {
	return cosmosdbSqlRoleAssignmentAttributes{ref: terra.ReferenceResource(csra)}
}

// ImportState imports the given attribute values into [CosmosdbSqlRoleAssignment]'s state.
func (csra *CosmosdbSqlRoleAssignment) ImportState(av io.Reader) error {
	csra.state = &cosmosdbSqlRoleAssignmentState{}
	if err := json.NewDecoder(av).Decode(csra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csra.Type(), csra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbSqlRoleAssignment] has state.
func (csra *CosmosdbSqlRoleAssignment) State() (*cosmosdbSqlRoleAssignmentState, bool) {
	return csra.state, csra.state != nil
}

// StateMust returns the state for [CosmosdbSqlRoleAssignment]. Panics if the state is nil.
func (csra *CosmosdbSqlRoleAssignment) StateMust() *cosmosdbSqlRoleAssignmentState {
	if csra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csra.Type(), csra.LocalName()))
	}
	return csra.state
}

// CosmosdbSqlRoleAssignmentArgs contains the configurations for azurerm_cosmosdb_sql_role_assignment.
type CosmosdbSqlRoleAssignmentArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RoleDefinitionId: string, required
	RoleDefinitionId terra.StringValue `hcl:"role_definition_id,attr" validate:"required"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cosmosdbsqlroleassignment.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbSqlRoleAssignmentAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_sql_role_assignment.
func (csra cosmosdbSqlRoleAssignmentAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(csra.ref.Append("account_name"))
}

// Id returns a reference to field id of azurerm_cosmosdb_sql_role_assignment.
func (csra cosmosdbSqlRoleAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csra.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_sql_role_assignment.
func (csra cosmosdbSqlRoleAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csra.ref.Append("name"))
}

// PrincipalId returns a reference to field principal_id of azurerm_cosmosdb_sql_role_assignment.
func (csra cosmosdbSqlRoleAssignmentAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(csra.ref.Append("principal_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_sql_role_assignment.
func (csra cosmosdbSqlRoleAssignmentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(csra.ref.Append("resource_group_name"))
}

// RoleDefinitionId returns a reference to field role_definition_id of azurerm_cosmosdb_sql_role_assignment.
func (csra cosmosdbSqlRoleAssignmentAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(csra.ref.Append("role_definition_id"))
}

// Scope returns a reference to field scope of azurerm_cosmosdb_sql_role_assignment.
func (csra cosmosdbSqlRoleAssignmentAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(csra.ref.Append("scope"))
}

func (csra cosmosdbSqlRoleAssignmentAttributes) Timeouts() cosmosdbsqlroleassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbsqlroleassignment.TimeoutsAttributes](csra.ref.Append("timeouts"))
}

type cosmosdbSqlRoleAssignmentState struct {
	AccountName       string                                   `json:"account_name"`
	Id                string                                   `json:"id"`
	Name              string                                   `json:"name"`
	PrincipalId       string                                   `json:"principal_id"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	RoleDefinitionId  string                                   `json:"role_definition_id"`
	Scope             string                                   `json:"scope"`
	Timeouts          *cosmosdbsqlroleassignment.TimeoutsState `json:"timeouts"`
}
