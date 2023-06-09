// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	roledefinition "github.com/golingon/terraproviders/azurerm/3.55.0/roledefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoleDefinition creates a new instance of [RoleDefinition].
func NewRoleDefinition(name string, args RoleDefinitionArgs) *RoleDefinition {
	return &RoleDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RoleDefinition)(nil)

// RoleDefinition represents the Terraform resource azurerm_role_definition.
type RoleDefinition struct {
	Name      string
	Args      RoleDefinitionArgs
	state     *roleDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RoleDefinition].
func (rd *RoleDefinition) Type() string {
	return "azurerm_role_definition"
}

// LocalName returns the local name for [RoleDefinition].
func (rd *RoleDefinition) LocalName() string {
	return rd.Name
}

// Configuration returns the configuration (args) for [RoleDefinition].
func (rd *RoleDefinition) Configuration() interface{} {
	return rd.Args
}

// DependOn is used for other resources to depend on [RoleDefinition].
func (rd *RoleDefinition) DependOn() terra.Reference {
	return terra.ReferenceResource(rd)
}

// Dependencies returns the list of resources [RoleDefinition] depends_on.
func (rd *RoleDefinition) Dependencies() terra.Dependencies {
	return rd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RoleDefinition].
func (rd *RoleDefinition) LifecycleManagement() *terra.Lifecycle {
	return rd.Lifecycle
}

// Attributes returns the attributes for [RoleDefinition].
func (rd *RoleDefinition) Attributes() roleDefinitionAttributes {
	return roleDefinitionAttributes{ref: terra.ReferenceResource(rd)}
}

// ImportState imports the given attribute values into [RoleDefinition]'s state.
func (rd *RoleDefinition) ImportState(av io.Reader) error {
	rd.state = &roleDefinitionState{}
	if err := json.NewDecoder(av).Decode(rd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rd.Type(), rd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RoleDefinition] has state.
func (rd *RoleDefinition) State() (*roleDefinitionState, bool) {
	return rd.state, rd.state != nil
}

// StateMust returns the state for [RoleDefinition]. Panics if the state is nil.
func (rd *RoleDefinition) StateMust() *roleDefinitionState {
	if rd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rd.Type(), rd.LocalName()))
	}
	return rd.state
}

// RoleDefinitionArgs contains the configurations for azurerm_role_definition.
type RoleDefinitionArgs struct {
	// AssignableScopes: list of string, optional
	AssignableScopes terra.ListValue[terra.StringValue] `hcl:"assignable_scopes,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleDefinitionId: string, optional
	RoleDefinitionId terra.StringValue `hcl:"role_definition_id,attr"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// Permissions: min=0
	Permissions []roledefinition.Permissions `hcl:"permissions,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *roledefinition.Timeouts `hcl:"timeouts,block"`
}
type roleDefinitionAttributes struct {
	ref terra.Reference
}

// AssignableScopes returns a reference to field assignable_scopes of azurerm_role_definition.
func (rd roleDefinitionAttributes) AssignableScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rd.ref.Append("assignable_scopes"))
}

// Description returns a reference to field description of azurerm_role_definition.
func (rd roleDefinitionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_role_definition.
func (rd roleDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_role_definition.
func (rd roleDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("name"))
}

// RoleDefinitionId returns a reference to field role_definition_id of azurerm_role_definition.
func (rd roleDefinitionAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("role_definition_id"))
}

// RoleDefinitionResourceId returns a reference to field role_definition_resource_id of azurerm_role_definition.
func (rd roleDefinitionAttributes) RoleDefinitionResourceId() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("role_definition_resource_id"))
}

// Scope returns a reference to field scope of azurerm_role_definition.
func (rd roleDefinitionAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("scope"))
}

func (rd roleDefinitionAttributes) Permissions() terra.ListValue[roledefinition.PermissionsAttributes] {
	return terra.ReferenceAsList[roledefinition.PermissionsAttributes](rd.ref.Append("permissions"))
}

func (rd roleDefinitionAttributes) Timeouts() roledefinition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[roledefinition.TimeoutsAttributes](rd.ref.Append("timeouts"))
}

type roleDefinitionState struct {
	AssignableScopes         []string                          `json:"assignable_scopes"`
	Description              string                            `json:"description"`
	Id                       string                            `json:"id"`
	Name                     string                            `json:"name"`
	RoleDefinitionId         string                            `json:"role_definition_id"`
	RoleDefinitionResourceId string                            `json:"role_definition_resource_id"`
	Scope                    string                            `json:"scope"`
	Permissions              []roledefinition.PermissionsState `json:"permissions"`
	Timeouts                 *roledefinition.TimeoutsState     `json:"timeouts"`
}
