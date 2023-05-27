// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	administrativeunitrolemember "github.com/golingon/terraproviders/azuread/2.39.0/administrativeunitrolemember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAdministrativeUnitRoleMember creates a new instance of [AdministrativeUnitRoleMember].
func NewAdministrativeUnitRoleMember(name string, args AdministrativeUnitRoleMemberArgs) *AdministrativeUnitRoleMember {
	return &AdministrativeUnitRoleMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AdministrativeUnitRoleMember)(nil)

// AdministrativeUnitRoleMember represents the Terraform resource azuread_administrative_unit_role_member.
type AdministrativeUnitRoleMember struct {
	Name      string
	Args      AdministrativeUnitRoleMemberArgs
	state     *administrativeUnitRoleMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AdministrativeUnitRoleMember].
func (aurm *AdministrativeUnitRoleMember) Type() string {
	return "azuread_administrative_unit_role_member"
}

// LocalName returns the local name for [AdministrativeUnitRoleMember].
func (aurm *AdministrativeUnitRoleMember) LocalName() string {
	return aurm.Name
}

// Configuration returns the configuration (args) for [AdministrativeUnitRoleMember].
func (aurm *AdministrativeUnitRoleMember) Configuration() interface{} {
	return aurm.Args
}

// DependOn is used for other resources to depend on [AdministrativeUnitRoleMember].
func (aurm *AdministrativeUnitRoleMember) DependOn() terra.Reference {
	return terra.ReferenceResource(aurm)
}

// Dependencies returns the list of resources [AdministrativeUnitRoleMember] depends_on.
func (aurm *AdministrativeUnitRoleMember) Dependencies() terra.Dependencies {
	return aurm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AdministrativeUnitRoleMember].
func (aurm *AdministrativeUnitRoleMember) LifecycleManagement() *terra.Lifecycle {
	return aurm.Lifecycle
}

// Attributes returns the attributes for [AdministrativeUnitRoleMember].
func (aurm *AdministrativeUnitRoleMember) Attributes() administrativeUnitRoleMemberAttributes {
	return administrativeUnitRoleMemberAttributes{ref: terra.ReferenceResource(aurm)}
}

// ImportState imports the given attribute values into [AdministrativeUnitRoleMember]'s state.
func (aurm *AdministrativeUnitRoleMember) ImportState(av io.Reader) error {
	aurm.state = &administrativeUnitRoleMemberState{}
	if err := json.NewDecoder(av).Decode(aurm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aurm.Type(), aurm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AdministrativeUnitRoleMember] has state.
func (aurm *AdministrativeUnitRoleMember) State() (*administrativeUnitRoleMemberState, bool) {
	return aurm.state, aurm.state != nil
}

// StateMust returns the state for [AdministrativeUnitRoleMember]. Panics if the state is nil.
func (aurm *AdministrativeUnitRoleMember) StateMust() *administrativeUnitRoleMemberState {
	if aurm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aurm.Type(), aurm.LocalName()))
	}
	return aurm.state
}

// AdministrativeUnitRoleMemberArgs contains the configurations for azuread_administrative_unit_role_member.
type AdministrativeUnitRoleMemberArgs struct {
	// AdministrativeUnitObjectId: string, required
	AdministrativeUnitObjectId terra.StringValue `hcl:"administrative_unit_object_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MemberObjectId: string, required
	MemberObjectId terra.StringValue `hcl:"member_object_id,attr" validate:"required"`
	// RoleObjectId: string, required
	RoleObjectId terra.StringValue `hcl:"role_object_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *administrativeunitrolemember.Timeouts `hcl:"timeouts,block"`
}
type administrativeUnitRoleMemberAttributes struct {
	ref terra.Reference
}

// AdministrativeUnitObjectId returns a reference to field administrative_unit_object_id of azuread_administrative_unit_role_member.
func (aurm administrativeUnitRoleMemberAttributes) AdministrativeUnitObjectId() terra.StringValue {
	return terra.ReferenceAsString(aurm.ref.Append("administrative_unit_object_id"))
}

// Id returns a reference to field id of azuread_administrative_unit_role_member.
func (aurm administrativeUnitRoleMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aurm.ref.Append("id"))
}

// MemberObjectId returns a reference to field member_object_id of azuread_administrative_unit_role_member.
func (aurm administrativeUnitRoleMemberAttributes) MemberObjectId() terra.StringValue {
	return terra.ReferenceAsString(aurm.ref.Append("member_object_id"))
}

// RoleObjectId returns a reference to field role_object_id of azuread_administrative_unit_role_member.
func (aurm administrativeUnitRoleMemberAttributes) RoleObjectId() terra.StringValue {
	return terra.ReferenceAsString(aurm.ref.Append("role_object_id"))
}

func (aurm administrativeUnitRoleMemberAttributes) Timeouts() administrativeunitrolemember.TimeoutsAttributes {
	return terra.ReferenceAsSingle[administrativeunitrolemember.TimeoutsAttributes](aurm.ref.Append("timeouts"))
}

type administrativeUnitRoleMemberState struct {
	AdministrativeUnitObjectId string                                      `json:"administrative_unit_object_id"`
	Id                         string                                      `json:"id"`
	MemberObjectId             string                                      `json:"member_object_id"`
	RoleObjectId               string                                      `json:"role_object_id"`
	Timeouts                   *administrativeunitrolemember.TimeoutsState `json:"timeouts"`
}
