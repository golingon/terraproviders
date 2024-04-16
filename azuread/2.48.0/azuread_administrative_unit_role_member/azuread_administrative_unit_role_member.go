// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread_administrative_unit_role_member

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

// Resource represents the Terraform resource azuread_administrative_unit_role_member.
type Resource struct {
	Name      string
	Args      Args
	state     *azureadAdministrativeUnitRoleMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aaurm *Resource) Type() string {
	return "azuread_administrative_unit_role_member"
}

// LocalName returns the local name for [Resource].
func (aaurm *Resource) LocalName() string {
	return aaurm.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aaurm *Resource) Configuration() interface{} {
	return aaurm.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aaurm *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aaurm)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aaurm *Resource) Dependencies() terra.Dependencies {
	return aaurm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aaurm *Resource) LifecycleManagement() *terra.Lifecycle {
	return aaurm.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aaurm *Resource) Attributes() azureadAdministrativeUnitRoleMemberAttributes {
	return azureadAdministrativeUnitRoleMemberAttributes{ref: terra.ReferenceResource(aaurm)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aaurm *Resource) ImportState(state io.Reader) error {
	aaurm.state = &azureadAdministrativeUnitRoleMemberState{}
	if err := json.NewDecoder(state).Decode(aaurm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aaurm.Type(), aaurm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aaurm *Resource) State() (*azureadAdministrativeUnitRoleMemberState, bool) {
	return aaurm.state, aaurm.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aaurm *Resource) StateMust() *azureadAdministrativeUnitRoleMemberState {
	if aaurm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aaurm.Type(), aaurm.LocalName()))
	}
	return aaurm.state
}

// Args contains the configurations for azuread_administrative_unit_role_member.
type Args struct {
	// AdministrativeUnitObjectId: string, required
	AdministrativeUnitObjectId terra.StringValue `hcl:"administrative_unit_object_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MemberObjectId: string, required
	MemberObjectId terra.StringValue `hcl:"member_object_id,attr" validate:"required"`
	// RoleObjectId: string, required
	RoleObjectId terra.StringValue `hcl:"role_object_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azureadAdministrativeUnitRoleMemberAttributes struct {
	ref terra.Reference
}

// AdministrativeUnitObjectId returns a reference to field administrative_unit_object_id of azuread_administrative_unit_role_member.
func (aaurm azureadAdministrativeUnitRoleMemberAttributes) AdministrativeUnitObjectId() terra.StringValue {
	return terra.ReferenceAsString(aaurm.ref.Append("administrative_unit_object_id"))
}

// Id returns a reference to field id of azuread_administrative_unit_role_member.
func (aaurm azureadAdministrativeUnitRoleMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aaurm.ref.Append("id"))
}

// MemberObjectId returns a reference to field member_object_id of azuread_administrative_unit_role_member.
func (aaurm azureadAdministrativeUnitRoleMemberAttributes) MemberObjectId() terra.StringValue {
	return terra.ReferenceAsString(aaurm.ref.Append("member_object_id"))
}

// RoleObjectId returns a reference to field role_object_id of azuread_administrative_unit_role_member.
func (aaurm azureadAdministrativeUnitRoleMemberAttributes) RoleObjectId() terra.StringValue {
	return terra.ReferenceAsString(aaurm.ref.Append("role_object_id"))
}

func (aaurm azureadAdministrativeUnitRoleMemberAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aaurm.ref.Append("timeouts"))
}

type azureadAdministrativeUnitRoleMemberState struct {
	AdministrativeUnitObjectId string         `json:"administrative_unit_object_id"`
	Id                         string         `json:"id"`
	MemberObjectId             string         `json:"member_object_id"`
	RoleObjectId               string         `json:"role_object_id"`
	Timeouts                   *TimeoutsState `json:"timeouts"`
}
