// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googleworkspace

import (
	"encoding/json"
	"fmt"
	groupmembers "github.com/golingon/terraproviders/googleworkspace/0.7.0/groupmembers"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGroupMembers creates a new instance of [GroupMembers].
func NewGroupMembers(name string, args GroupMembersArgs) *GroupMembers {
	return &GroupMembers{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GroupMembers)(nil)

// GroupMembers represents the Terraform resource googleworkspace_group_members.
type GroupMembers struct {
	Name      string
	Args      GroupMembersArgs
	state     *groupMembersState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GroupMembers].
func (gm *GroupMembers) Type() string {
	return "googleworkspace_group_members"
}

// LocalName returns the local name for [GroupMembers].
func (gm *GroupMembers) LocalName() string {
	return gm.Name
}

// Configuration returns the configuration (args) for [GroupMembers].
func (gm *GroupMembers) Configuration() interface{} {
	return gm.Args
}

// DependOn is used for other resources to depend on [GroupMembers].
func (gm *GroupMembers) DependOn() terra.Reference {
	return terra.ReferenceResource(gm)
}

// Dependencies returns the list of resources [GroupMembers] depends_on.
func (gm *GroupMembers) Dependencies() terra.Dependencies {
	return gm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GroupMembers].
func (gm *GroupMembers) LifecycleManagement() *terra.Lifecycle {
	return gm.Lifecycle
}

// Attributes returns the attributes for [GroupMembers].
func (gm *GroupMembers) Attributes() groupMembersAttributes {
	return groupMembersAttributes{ref: terra.ReferenceResource(gm)}
}

// ImportState imports the given attribute values into [GroupMembers]'s state.
func (gm *GroupMembers) ImportState(av io.Reader) error {
	gm.state = &groupMembersState{}
	if err := json.NewDecoder(av).Decode(gm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gm.Type(), gm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GroupMembers] has state.
func (gm *GroupMembers) State() (*groupMembersState, bool) {
	return gm.state, gm.state != nil
}

// StateMust returns the state for [GroupMembers]. Panics if the state is nil.
func (gm *GroupMembers) StateMust() *groupMembersState {
	if gm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gm.Type(), gm.LocalName()))
	}
	return gm.state
}

// GroupMembersArgs contains the configurations for googleworkspace_group_members.
type GroupMembersArgs struct {
	// GroupId: string, required
	GroupId terra.StringValue `hcl:"group_id,attr" validate:"required"`
	// Members: min=0
	Members []groupmembers.Members `hcl:"members,block" validate:"min=0"`
}
type groupMembersAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of googleworkspace_group_members.
func (gm groupMembersAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("etag"))
}

// GroupId returns a reference to field group_id of googleworkspace_group_members.
func (gm groupMembersAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("group_id"))
}

// Id returns a reference to field id of googleworkspace_group_members.
func (gm groupMembersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("id"))
}

func (gm groupMembersAttributes) Members() terra.SetValue[groupmembers.MembersAttributes] {
	return terra.ReferenceAsSet[groupmembers.MembersAttributes](gm.ref.Append("members"))
}

type groupMembersState struct {
	Etag    string                      `json:"etag"`
	GroupId string                      `json:"group_id"`
	Id      string                      `json:"id"`
	Members []groupmembers.MembersState `json:"members"`
}