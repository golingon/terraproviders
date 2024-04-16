// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread_directory_role_member

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

// Resource represents the Terraform resource azuread_directory_role_member.
type Resource struct {
	Name      string
	Args      Args
	state     *azureadDirectoryRoleMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adrm *Resource) Type() string {
	return "azuread_directory_role_member"
}

// LocalName returns the local name for [Resource].
func (adrm *Resource) LocalName() string {
	return adrm.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adrm *Resource) Configuration() interface{} {
	return adrm.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adrm *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adrm)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adrm *Resource) Dependencies() terra.Dependencies {
	return adrm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adrm *Resource) LifecycleManagement() *terra.Lifecycle {
	return adrm.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adrm *Resource) Attributes() azureadDirectoryRoleMemberAttributes {
	return azureadDirectoryRoleMemberAttributes{ref: terra.ReferenceResource(adrm)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adrm *Resource) ImportState(state io.Reader) error {
	adrm.state = &azureadDirectoryRoleMemberState{}
	if err := json.NewDecoder(state).Decode(adrm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adrm.Type(), adrm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adrm *Resource) State() (*azureadDirectoryRoleMemberState, bool) {
	return adrm.state, adrm.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adrm *Resource) StateMust() *azureadDirectoryRoleMemberState {
	if adrm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adrm.Type(), adrm.LocalName()))
	}
	return adrm.state
}

// Args contains the configurations for azuread_directory_role_member.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MemberObjectId: string, optional
	MemberObjectId terra.StringValue `hcl:"member_object_id,attr"`
	// RoleObjectId: string, optional
	RoleObjectId terra.StringValue `hcl:"role_object_id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azureadDirectoryRoleMemberAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azuread_directory_role_member.
func (adrm azureadDirectoryRoleMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adrm.ref.Append("id"))
}

// MemberObjectId returns a reference to field member_object_id of azuread_directory_role_member.
func (adrm azureadDirectoryRoleMemberAttributes) MemberObjectId() terra.StringValue {
	return terra.ReferenceAsString(adrm.ref.Append("member_object_id"))
}

// RoleObjectId returns a reference to field role_object_id of azuread_directory_role_member.
func (adrm azureadDirectoryRoleMemberAttributes) RoleObjectId() terra.StringValue {
	return terra.ReferenceAsString(adrm.ref.Append("role_object_id"))
}

func (adrm azureadDirectoryRoleMemberAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adrm.ref.Append("timeouts"))
}

type azureadDirectoryRoleMemberState struct {
	Id             string         `json:"id"`
	MemberObjectId string         `json:"member_object_id"`
	RoleObjectId   string         `json:"role_object_id"`
	Timeouts       *TimeoutsState `json:"timeouts"`
}
