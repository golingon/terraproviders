// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	groupmember "github.com/golingon/terraproviders/azuread/2.39.0/groupmember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGroupMember creates a new instance of [GroupMember].
func NewGroupMember(name string, args GroupMemberArgs) *GroupMember {
	return &GroupMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GroupMember)(nil)

// GroupMember represents the Terraform resource azuread_group_member.
type GroupMember struct {
	Name      string
	Args      GroupMemberArgs
	state     *groupMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GroupMember].
func (gm *GroupMember) Type() string {
	return "azuread_group_member"
}

// LocalName returns the local name for [GroupMember].
func (gm *GroupMember) LocalName() string {
	return gm.Name
}

// Configuration returns the configuration (args) for [GroupMember].
func (gm *GroupMember) Configuration() interface{} {
	return gm.Args
}

// DependOn is used for other resources to depend on [GroupMember].
func (gm *GroupMember) DependOn() terra.Reference {
	return terra.ReferenceResource(gm)
}

// Dependencies returns the list of resources [GroupMember] depends_on.
func (gm *GroupMember) Dependencies() terra.Dependencies {
	return gm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GroupMember].
func (gm *GroupMember) LifecycleManagement() *terra.Lifecycle {
	return gm.Lifecycle
}

// Attributes returns the attributes for [GroupMember].
func (gm *GroupMember) Attributes() groupMemberAttributes {
	return groupMemberAttributes{ref: terra.ReferenceResource(gm)}
}

// ImportState imports the given attribute values into [GroupMember]'s state.
func (gm *GroupMember) ImportState(av io.Reader) error {
	gm.state = &groupMemberState{}
	if err := json.NewDecoder(av).Decode(gm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gm.Type(), gm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GroupMember] has state.
func (gm *GroupMember) State() (*groupMemberState, bool) {
	return gm.state, gm.state != nil
}

// StateMust returns the state for [GroupMember]. Panics if the state is nil.
func (gm *GroupMember) StateMust() *groupMemberState {
	if gm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gm.Type(), gm.LocalName()))
	}
	return gm.state
}

// GroupMemberArgs contains the configurations for azuread_group_member.
type GroupMemberArgs struct {
	// GroupObjectId: string, required
	GroupObjectId terra.StringValue `hcl:"group_object_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MemberObjectId: string, required
	MemberObjectId terra.StringValue `hcl:"member_object_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *groupmember.Timeouts `hcl:"timeouts,block"`
}
type groupMemberAttributes struct {
	ref terra.Reference
}

// GroupObjectId returns a reference to field group_object_id of azuread_group_member.
func (gm groupMemberAttributes) GroupObjectId() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("group_object_id"))
}

// Id returns a reference to field id of azuread_group_member.
func (gm groupMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("id"))
}

// MemberObjectId returns a reference to field member_object_id of azuread_group_member.
func (gm groupMemberAttributes) MemberObjectId() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("member_object_id"))
}

func (gm groupMemberAttributes) Timeouts() groupmember.TimeoutsAttributes {
	return terra.ReferenceAsSingle[groupmember.TimeoutsAttributes](gm.ref.Append("timeouts"))
}

type groupMemberState struct {
	GroupObjectId  string                     `json:"group_object_id"`
	Id             string                     `json:"id"`
	MemberObjectId string                     `json:"member_object_id"`
	Timeouts       *groupmember.TimeoutsState `json:"timeouts"`
}
