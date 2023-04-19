// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	administrativeunitmember "github.com/golingon/terraproviders/azuread/2.37.1/administrativeunitmember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAdministrativeUnitMember creates a new instance of [AdministrativeUnitMember].
func NewAdministrativeUnitMember(name string, args AdministrativeUnitMemberArgs) *AdministrativeUnitMember {
	return &AdministrativeUnitMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AdministrativeUnitMember)(nil)

// AdministrativeUnitMember represents the Terraform resource azuread_administrative_unit_member.
type AdministrativeUnitMember struct {
	Name      string
	Args      AdministrativeUnitMemberArgs
	state     *administrativeUnitMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AdministrativeUnitMember].
func (aum *AdministrativeUnitMember) Type() string {
	return "azuread_administrative_unit_member"
}

// LocalName returns the local name for [AdministrativeUnitMember].
func (aum *AdministrativeUnitMember) LocalName() string {
	return aum.Name
}

// Configuration returns the configuration (args) for [AdministrativeUnitMember].
func (aum *AdministrativeUnitMember) Configuration() interface{} {
	return aum.Args
}

// DependOn is used for other resources to depend on [AdministrativeUnitMember].
func (aum *AdministrativeUnitMember) DependOn() terra.Reference {
	return terra.ReferenceResource(aum)
}

// Dependencies returns the list of resources [AdministrativeUnitMember] depends_on.
func (aum *AdministrativeUnitMember) Dependencies() terra.Dependencies {
	return aum.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AdministrativeUnitMember].
func (aum *AdministrativeUnitMember) LifecycleManagement() *terra.Lifecycle {
	return aum.Lifecycle
}

// Attributes returns the attributes for [AdministrativeUnitMember].
func (aum *AdministrativeUnitMember) Attributes() administrativeUnitMemberAttributes {
	return administrativeUnitMemberAttributes{ref: terra.ReferenceResource(aum)}
}

// ImportState imports the given attribute values into [AdministrativeUnitMember]'s state.
func (aum *AdministrativeUnitMember) ImportState(av io.Reader) error {
	aum.state = &administrativeUnitMemberState{}
	if err := json.NewDecoder(av).Decode(aum.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aum.Type(), aum.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AdministrativeUnitMember] has state.
func (aum *AdministrativeUnitMember) State() (*administrativeUnitMemberState, bool) {
	return aum.state, aum.state != nil
}

// StateMust returns the state for [AdministrativeUnitMember]. Panics if the state is nil.
func (aum *AdministrativeUnitMember) StateMust() *administrativeUnitMemberState {
	if aum.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aum.Type(), aum.LocalName()))
	}
	return aum.state
}

// AdministrativeUnitMemberArgs contains the configurations for azuread_administrative_unit_member.
type AdministrativeUnitMemberArgs struct {
	// AdministrativeUnitObjectId: string, optional
	AdministrativeUnitObjectId terra.StringValue `hcl:"administrative_unit_object_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MemberObjectId: string, optional
	MemberObjectId terra.StringValue `hcl:"member_object_id,attr"`
	// Timeouts: optional
	Timeouts *administrativeunitmember.Timeouts `hcl:"timeouts,block"`
}
type administrativeUnitMemberAttributes struct {
	ref terra.Reference
}

// AdministrativeUnitObjectId returns a reference to field administrative_unit_object_id of azuread_administrative_unit_member.
func (aum administrativeUnitMemberAttributes) AdministrativeUnitObjectId() terra.StringValue {
	return terra.ReferenceAsString(aum.ref.Append("administrative_unit_object_id"))
}

// Id returns a reference to field id of azuread_administrative_unit_member.
func (aum administrativeUnitMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aum.ref.Append("id"))
}

// MemberObjectId returns a reference to field member_object_id of azuread_administrative_unit_member.
func (aum administrativeUnitMemberAttributes) MemberObjectId() terra.StringValue {
	return terra.ReferenceAsString(aum.ref.Append("member_object_id"))
}

func (aum administrativeUnitMemberAttributes) Timeouts() administrativeunitmember.TimeoutsAttributes {
	return terra.ReferenceAsSingle[administrativeunitmember.TimeoutsAttributes](aum.ref.Append("timeouts"))
}

type administrativeUnitMemberState struct {
	AdministrativeUnitObjectId string                                  `json:"administrative_unit_object_id"`
	Id                         string                                  `json:"id"`
	MemberObjectId             string                                  `json:"member_object_id"`
	Timeouts                   *administrativeunitmember.TimeoutsState `json:"timeouts"`
}
