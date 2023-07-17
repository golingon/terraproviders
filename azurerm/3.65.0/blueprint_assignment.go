// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	blueprintassignment "github.com/golingon/terraproviders/azurerm/3.65.0/blueprintassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBlueprintAssignment creates a new instance of [BlueprintAssignment].
func NewBlueprintAssignment(name string, args BlueprintAssignmentArgs) *BlueprintAssignment {
	return &BlueprintAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BlueprintAssignment)(nil)

// BlueprintAssignment represents the Terraform resource azurerm_blueprint_assignment.
type BlueprintAssignment struct {
	Name      string
	Args      BlueprintAssignmentArgs
	state     *blueprintAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BlueprintAssignment].
func (ba *BlueprintAssignment) Type() string {
	return "azurerm_blueprint_assignment"
}

// LocalName returns the local name for [BlueprintAssignment].
func (ba *BlueprintAssignment) LocalName() string {
	return ba.Name
}

// Configuration returns the configuration (args) for [BlueprintAssignment].
func (ba *BlueprintAssignment) Configuration() interface{} {
	return ba.Args
}

// DependOn is used for other resources to depend on [BlueprintAssignment].
func (ba *BlueprintAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(ba)
}

// Dependencies returns the list of resources [BlueprintAssignment] depends_on.
func (ba *BlueprintAssignment) Dependencies() terra.Dependencies {
	return ba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BlueprintAssignment].
func (ba *BlueprintAssignment) LifecycleManagement() *terra.Lifecycle {
	return ba.Lifecycle
}

// Attributes returns the attributes for [BlueprintAssignment].
func (ba *BlueprintAssignment) Attributes() blueprintAssignmentAttributes {
	return blueprintAssignmentAttributes{ref: terra.ReferenceResource(ba)}
}

// ImportState imports the given attribute values into [BlueprintAssignment]'s state.
func (ba *BlueprintAssignment) ImportState(av io.Reader) error {
	ba.state = &blueprintAssignmentState{}
	if err := json.NewDecoder(av).Decode(ba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ba.Type(), ba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BlueprintAssignment] has state.
func (ba *BlueprintAssignment) State() (*blueprintAssignmentState, bool) {
	return ba.state, ba.state != nil
}

// StateMust returns the state for [BlueprintAssignment]. Panics if the state is nil.
func (ba *BlueprintAssignment) StateMust() *blueprintAssignmentState {
	if ba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ba.Type(), ba.LocalName()))
	}
	return ba.state
}

// BlueprintAssignmentArgs contains the configurations for azurerm_blueprint_assignment.
type BlueprintAssignmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// LockExcludeActions: list of string, optional
	LockExcludeActions terra.ListValue[terra.StringValue] `hcl:"lock_exclude_actions,attr"`
	// LockExcludePrincipals: list of string, optional
	LockExcludePrincipals terra.ListValue[terra.StringValue] `hcl:"lock_exclude_principals,attr"`
	// LockMode: string, optional
	LockMode terra.StringValue `hcl:"lock_mode,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParameterValues: string, optional
	ParameterValues terra.StringValue `hcl:"parameter_values,attr"`
	// ResourceGroups: string, optional
	ResourceGroups terra.StringValue `hcl:"resource_groups,attr"`
	// TargetSubscriptionId: string, required
	TargetSubscriptionId terra.StringValue `hcl:"target_subscription_id,attr" validate:"required"`
	// VersionId: string, required
	VersionId terra.StringValue `hcl:"version_id,attr" validate:"required"`
	// Identity: required
	Identity *blueprintassignment.Identity `hcl:"identity,block" validate:"required"`
	// Timeouts: optional
	Timeouts *blueprintassignment.Timeouts `hcl:"timeouts,block"`
}
type blueprintAssignmentAttributes struct {
	ref terra.Reference
}

// BlueprintName returns a reference to field blueprint_name of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) BlueprintName() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("blueprint_name"))
}

// Description returns a reference to field description of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("location"))
}

// LockExcludeActions returns a reference to field lock_exclude_actions of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) LockExcludeActions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ba.ref.Append("lock_exclude_actions"))
}

// LockExcludePrincipals returns a reference to field lock_exclude_principals of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) LockExcludePrincipals() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ba.ref.Append("lock_exclude_principals"))
}

// LockMode returns a reference to field lock_mode of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) LockMode() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("lock_mode"))
}

// Name returns a reference to field name of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("name"))
}

// ParameterValues returns a reference to field parameter_values of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) ParameterValues() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("parameter_values"))
}

// ResourceGroups returns a reference to field resource_groups of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) ResourceGroups() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("resource_groups"))
}

// TargetSubscriptionId returns a reference to field target_subscription_id of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) TargetSubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("target_subscription_id"))
}

// Type returns a reference to field type of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("type"))
}

// VersionId returns a reference to field version_id of azurerm_blueprint_assignment.
func (ba blueprintAssignmentAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("version_id"))
}

func (ba blueprintAssignmentAttributes) Identity() terra.ListValue[blueprintassignment.IdentityAttributes] {
	return terra.ReferenceAsList[blueprintassignment.IdentityAttributes](ba.ref.Append("identity"))
}

func (ba blueprintAssignmentAttributes) Timeouts() blueprintassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[blueprintassignment.TimeoutsAttributes](ba.ref.Append("timeouts"))
}

type blueprintAssignmentState struct {
	BlueprintName         string                              `json:"blueprint_name"`
	Description           string                              `json:"description"`
	DisplayName           string                              `json:"display_name"`
	Id                    string                              `json:"id"`
	Location              string                              `json:"location"`
	LockExcludeActions    []string                            `json:"lock_exclude_actions"`
	LockExcludePrincipals []string                            `json:"lock_exclude_principals"`
	LockMode              string                              `json:"lock_mode"`
	Name                  string                              `json:"name"`
	ParameterValues       string                              `json:"parameter_values"`
	ResourceGroups        string                              `json:"resource_groups"`
	TargetSubscriptionId  string                              `json:"target_subscription_id"`
	Type                  string                              `json:"type"`
	VersionId             string                              `json:"version_id"`
	Identity              []blueprintassignment.IdentityState `json:"identity"`
	Timeouts              *blueprintassignment.TimeoutsState  `json:"timeouts"`
}