// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread_administrative_unit

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

// Resource represents the Terraform resource azuread_administrative_unit.
type Resource struct {
	Name      string
	Args      Args
	state     *azureadAdministrativeUnitState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aau *Resource) Type() string {
	return "azuread_administrative_unit"
}

// LocalName returns the local name for [Resource].
func (aau *Resource) LocalName() string {
	return aau.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aau *Resource) Configuration() interface{} {
	return aau.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aau *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aau)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aau *Resource) Dependencies() terra.Dependencies {
	return aau.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aau *Resource) LifecycleManagement() *terra.Lifecycle {
	return aau.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aau *Resource) Attributes() azureadAdministrativeUnitAttributes {
	return azureadAdministrativeUnitAttributes{ref: terra.ReferenceResource(aau)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aau *Resource) ImportState(state io.Reader) error {
	aau.state = &azureadAdministrativeUnitState{}
	if err := json.NewDecoder(state).Decode(aau.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aau.Type(), aau.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aau *Resource) State() (*azureadAdministrativeUnitState, bool) {
	return aau.state, aau.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aau *Resource) StateMust() *azureadAdministrativeUnitState {
	if aau.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aau.Type(), aau.LocalName()))
	}
	return aau.state
}

// Args contains the configurations for azuread_administrative_unit.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// HiddenMembershipEnabled: bool, optional
	HiddenMembershipEnabled terra.BoolValue `hcl:"hidden_membership_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, optional
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr"`
	// PreventDuplicateNames: bool, optional
	PreventDuplicateNames terra.BoolValue `hcl:"prevent_duplicate_names,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azureadAdministrativeUnitAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azuread_administrative_unit.
func (aau azureadAdministrativeUnitAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aau.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_administrative_unit.
func (aau azureadAdministrativeUnitAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aau.ref.Append("display_name"))
}

// HiddenMembershipEnabled returns a reference to field hidden_membership_enabled of azuread_administrative_unit.
func (aau azureadAdministrativeUnitAttributes) HiddenMembershipEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aau.ref.Append("hidden_membership_enabled"))
}

// Id returns a reference to field id of azuread_administrative_unit.
func (aau azureadAdministrativeUnitAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aau.ref.Append("id"))
}

// Members returns a reference to field members of azuread_administrative_unit.
func (aau azureadAdministrativeUnitAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aau.ref.Append("members"))
}

// ObjectId returns a reference to field object_id of azuread_administrative_unit.
func (aau azureadAdministrativeUnitAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(aau.ref.Append("object_id"))
}

// PreventDuplicateNames returns a reference to field prevent_duplicate_names of azuread_administrative_unit.
func (aau azureadAdministrativeUnitAttributes) PreventDuplicateNames() terra.BoolValue {
	return terra.ReferenceAsBool(aau.ref.Append("prevent_duplicate_names"))
}

func (aau azureadAdministrativeUnitAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aau.ref.Append("timeouts"))
}

type azureadAdministrativeUnitState struct {
	Description             string         `json:"description"`
	DisplayName             string         `json:"display_name"`
	HiddenMembershipEnabled bool           `json:"hidden_membership_enabled"`
	Id                      string         `json:"id"`
	Members                 []string       `json:"members"`
	ObjectId                string         `json:"object_id"`
	PreventDuplicateNames   bool           `json:"prevent_duplicate_names"`
	Timeouts                *TimeoutsState `json:"timeouts"`
}
