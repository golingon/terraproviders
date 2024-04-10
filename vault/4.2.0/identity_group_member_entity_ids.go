// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewIdentityGroupMemberEntityIds creates a new instance of [IdentityGroupMemberEntityIds].
func NewIdentityGroupMemberEntityIds(name string, args IdentityGroupMemberEntityIdsArgs) *IdentityGroupMemberEntityIds {
	return &IdentityGroupMemberEntityIds{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityGroupMemberEntityIds)(nil)

// IdentityGroupMemberEntityIds represents the Terraform resource vault_identity_group_member_entity_ids.
type IdentityGroupMemberEntityIds struct {
	Name      string
	Args      IdentityGroupMemberEntityIdsArgs
	state     *identityGroupMemberEntityIdsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityGroupMemberEntityIds].
func (igmei *IdentityGroupMemberEntityIds) Type() string {
	return "vault_identity_group_member_entity_ids"
}

// LocalName returns the local name for [IdentityGroupMemberEntityIds].
func (igmei *IdentityGroupMemberEntityIds) LocalName() string {
	return igmei.Name
}

// Configuration returns the configuration (args) for [IdentityGroupMemberEntityIds].
func (igmei *IdentityGroupMemberEntityIds) Configuration() interface{} {
	return igmei.Args
}

// DependOn is used for other resources to depend on [IdentityGroupMemberEntityIds].
func (igmei *IdentityGroupMemberEntityIds) DependOn() terra.Reference {
	return terra.ReferenceResource(igmei)
}

// Dependencies returns the list of resources [IdentityGroupMemberEntityIds] depends_on.
func (igmei *IdentityGroupMemberEntityIds) Dependencies() terra.Dependencies {
	return igmei.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityGroupMemberEntityIds].
func (igmei *IdentityGroupMemberEntityIds) LifecycleManagement() *terra.Lifecycle {
	return igmei.Lifecycle
}

// Attributes returns the attributes for [IdentityGroupMemberEntityIds].
func (igmei *IdentityGroupMemberEntityIds) Attributes() identityGroupMemberEntityIdsAttributes {
	return identityGroupMemberEntityIdsAttributes{ref: terra.ReferenceResource(igmei)}
}

// ImportState imports the given attribute values into [IdentityGroupMemberEntityIds]'s state.
func (igmei *IdentityGroupMemberEntityIds) ImportState(av io.Reader) error {
	igmei.state = &identityGroupMemberEntityIdsState{}
	if err := json.NewDecoder(av).Decode(igmei.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", igmei.Type(), igmei.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityGroupMemberEntityIds] has state.
func (igmei *IdentityGroupMemberEntityIds) State() (*identityGroupMemberEntityIdsState, bool) {
	return igmei.state, igmei.state != nil
}

// StateMust returns the state for [IdentityGroupMemberEntityIds]. Panics if the state is nil.
func (igmei *IdentityGroupMemberEntityIds) StateMust() *identityGroupMemberEntityIdsState {
	if igmei.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", igmei.Type(), igmei.LocalName()))
	}
	return igmei.state
}

// IdentityGroupMemberEntityIdsArgs contains the configurations for vault_identity_group_member_entity_ids.
type IdentityGroupMemberEntityIdsArgs struct {
	// Exclusive: bool, optional
	Exclusive terra.BoolValue `hcl:"exclusive,attr"`
	// GroupId: string, required
	GroupId terra.StringValue `hcl:"group_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MemberEntityIds: set of string, optional
	MemberEntityIds terra.SetValue[terra.StringValue] `hcl:"member_entity_ids,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}
type identityGroupMemberEntityIdsAttributes struct {
	ref terra.Reference
}

// Exclusive returns a reference to field exclusive of vault_identity_group_member_entity_ids.
func (igmei identityGroupMemberEntityIdsAttributes) Exclusive() terra.BoolValue {
	return terra.ReferenceAsBool(igmei.ref.Append("exclusive"))
}

// GroupId returns a reference to field group_id of vault_identity_group_member_entity_ids.
func (igmei identityGroupMemberEntityIdsAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(igmei.ref.Append("group_id"))
}

// Id returns a reference to field id of vault_identity_group_member_entity_ids.
func (igmei identityGroupMemberEntityIdsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(igmei.ref.Append("id"))
}

// MemberEntityIds returns a reference to field member_entity_ids of vault_identity_group_member_entity_ids.
func (igmei identityGroupMemberEntityIdsAttributes) MemberEntityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](igmei.ref.Append("member_entity_ids"))
}

// Namespace returns a reference to field namespace of vault_identity_group_member_entity_ids.
func (igmei identityGroupMemberEntityIdsAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(igmei.ref.Append("namespace"))
}

type identityGroupMemberEntityIdsState struct {
	Exclusive       bool     `json:"exclusive"`
	GroupId         string   `json:"group_id"`
	Id              string   `json:"id"`
	MemberEntityIds []string `json:"member_entity_ids"`
	Namespace       string   `json:"namespace"`
}
