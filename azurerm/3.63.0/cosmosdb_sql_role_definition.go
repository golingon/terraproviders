// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbsqlroledefinition "github.com/golingon/terraproviders/azurerm/3.63.0/cosmosdbsqlroledefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbSqlRoleDefinition creates a new instance of [CosmosdbSqlRoleDefinition].
func NewCosmosdbSqlRoleDefinition(name string, args CosmosdbSqlRoleDefinitionArgs) *CosmosdbSqlRoleDefinition {
	return &CosmosdbSqlRoleDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbSqlRoleDefinition)(nil)

// CosmosdbSqlRoleDefinition represents the Terraform resource azurerm_cosmosdb_sql_role_definition.
type CosmosdbSqlRoleDefinition struct {
	Name      string
	Args      CosmosdbSqlRoleDefinitionArgs
	state     *cosmosdbSqlRoleDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbSqlRoleDefinition].
func (csrd *CosmosdbSqlRoleDefinition) Type() string {
	return "azurerm_cosmosdb_sql_role_definition"
}

// LocalName returns the local name for [CosmosdbSqlRoleDefinition].
func (csrd *CosmosdbSqlRoleDefinition) LocalName() string {
	return csrd.Name
}

// Configuration returns the configuration (args) for [CosmosdbSqlRoleDefinition].
func (csrd *CosmosdbSqlRoleDefinition) Configuration() interface{} {
	return csrd.Args
}

// DependOn is used for other resources to depend on [CosmosdbSqlRoleDefinition].
func (csrd *CosmosdbSqlRoleDefinition) DependOn() terra.Reference {
	return terra.ReferenceResource(csrd)
}

// Dependencies returns the list of resources [CosmosdbSqlRoleDefinition] depends_on.
func (csrd *CosmosdbSqlRoleDefinition) Dependencies() terra.Dependencies {
	return csrd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbSqlRoleDefinition].
func (csrd *CosmosdbSqlRoleDefinition) LifecycleManagement() *terra.Lifecycle {
	return csrd.Lifecycle
}

// Attributes returns the attributes for [CosmosdbSqlRoleDefinition].
func (csrd *CosmosdbSqlRoleDefinition) Attributes() cosmosdbSqlRoleDefinitionAttributes {
	return cosmosdbSqlRoleDefinitionAttributes{ref: terra.ReferenceResource(csrd)}
}

// ImportState imports the given attribute values into [CosmosdbSqlRoleDefinition]'s state.
func (csrd *CosmosdbSqlRoleDefinition) ImportState(av io.Reader) error {
	csrd.state = &cosmosdbSqlRoleDefinitionState{}
	if err := json.NewDecoder(av).Decode(csrd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csrd.Type(), csrd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbSqlRoleDefinition] has state.
func (csrd *CosmosdbSqlRoleDefinition) State() (*cosmosdbSqlRoleDefinitionState, bool) {
	return csrd.state, csrd.state != nil
}

// StateMust returns the state for [CosmosdbSqlRoleDefinition]. Panics if the state is nil.
func (csrd *CosmosdbSqlRoleDefinition) StateMust() *cosmosdbSqlRoleDefinitionState {
	if csrd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csrd.Type(), csrd.LocalName()))
	}
	return csrd.state
}

// CosmosdbSqlRoleDefinitionArgs contains the configurations for azurerm_cosmosdb_sql_role_definition.
type CosmosdbSqlRoleDefinitionArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// AssignableScopes: set of string, required
	AssignableScopes terra.SetValue[terra.StringValue] `hcl:"assignable_scopes,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RoleDefinitionId: string, optional
	RoleDefinitionId terra.StringValue `hcl:"role_definition_id,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Permissions: min=1
	Permissions []cosmosdbsqlroledefinition.Permissions `hcl:"permissions,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *cosmosdbsqlroledefinition.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbSqlRoleDefinitionAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_sql_role_definition.
func (csrd cosmosdbSqlRoleDefinitionAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("account_name"))
}

// AssignableScopes returns a reference to field assignable_scopes of azurerm_cosmosdb_sql_role_definition.
func (csrd cosmosdbSqlRoleDefinitionAttributes) AssignableScopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csrd.ref.Append("assignable_scopes"))
}

// Id returns a reference to field id of azurerm_cosmosdb_sql_role_definition.
func (csrd cosmosdbSqlRoleDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_sql_role_definition.
func (csrd cosmosdbSqlRoleDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_sql_role_definition.
func (csrd cosmosdbSqlRoleDefinitionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("resource_group_name"))
}

// RoleDefinitionId returns a reference to field role_definition_id of azurerm_cosmosdb_sql_role_definition.
func (csrd cosmosdbSqlRoleDefinitionAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("role_definition_id"))
}

// Type returns a reference to field type of azurerm_cosmosdb_sql_role_definition.
func (csrd cosmosdbSqlRoleDefinitionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("type"))
}

func (csrd cosmosdbSqlRoleDefinitionAttributes) Permissions() terra.SetValue[cosmosdbsqlroledefinition.PermissionsAttributes] {
	return terra.ReferenceAsSet[cosmosdbsqlroledefinition.PermissionsAttributes](csrd.ref.Append("permissions"))
}

func (csrd cosmosdbSqlRoleDefinitionAttributes) Timeouts() cosmosdbsqlroledefinition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbsqlroledefinition.TimeoutsAttributes](csrd.ref.Append("timeouts"))
}

type cosmosdbSqlRoleDefinitionState struct {
	AccountName       string                                       `json:"account_name"`
	AssignableScopes  []string                                     `json:"assignable_scopes"`
	Id                string                                       `json:"id"`
	Name              string                                       `json:"name"`
	ResourceGroupName string                                       `json:"resource_group_name"`
	RoleDefinitionId  string                                       `json:"role_definition_id"`
	Type              string                                       `json:"type"`
	Permissions       []cosmosdbsqlroledefinition.PermissionsState `json:"permissions"`
	Timeouts          *cosmosdbsqlroledefinition.TimeoutsState     `json:"timeouts"`
}
