// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewIdentityGroup creates a new instance of [IdentityGroup].
func NewIdentityGroup(name string, args IdentityGroupArgs) *IdentityGroup {
	return &IdentityGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityGroup)(nil)

// IdentityGroup represents the Terraform resource vault_identity_group.
type IdentityGroup struct {
	Name      string
	Args      IdentityGroupArgs
	state     *identityGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityGroup].
func (ig *IdentityGroup) Type() string {
	return "vault_identity_group"
}

// LocalName returns the local name for [IdentityGroup].
func (ig *IdentityGroup) LocalName() string {
	return ig.Name
}

// Configuration returns the configuration (args) for [IdentityGroup].
func (ig *IdentityGroup) Configuration() interface{} {
	return ig.Args
}

// DependOn is used for other resources to depend on [IdentityGroup].
func (ig *IdentityGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ig)
}

// Dependencies returns the list of resources [IdentityGroup] depends_on.
func (ig *IdentityGroup) Dependencies() terra.Dependencies {
	return ig.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityGroup].
func (ig *IdentityGroup) LifecycleManagement() *terra.Lifecycle {
	return ig.Lifecycle
}

// Attributes returns the attributes for [IdentityGroup].
func (ig *IdentityGroup) Attributes() identityGroupAttributes {
	return identityGroupAttributes{ref: terra.ReferenceResource(ig)}
}

// ImportState imports the given attribute values into [IdentityGroup]'s state.
func (ig *IdentityGroup) ImportState(av io.Reader) error {
	ig.state = &identityGroupState{}
	if err := json.NewDecoder(av).Decode(ig.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ig.Type(), ig.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityGroup] has state.
func (ig *IdentityGroup) State() (*identityGroupState, bool) {
	return ig.state, ig.state != nil
}

// StateMust returns the state for [IdentityGroup]. Panics if the state is nil.
func (ig *IdentityGroup) StateMust() *identityGroupState {
	if ig.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ig.Type(), ig.LocalName()))
	}
	return ig.state
}

// IdentityGroupArgs contains the configurations for vault_identity_group.
type IdentityGroupArgs struct {
	// ExternalMemberEntityIds: bool, optional
	ExternalMemberEntityIds terra.BoolValue `hcl:"external_member_entity_ids,attr"`
	// ExternalMemberGroupIds: bool, optional
	ExternalMemberGroupIds terra.BoolValue `hcl:"external_member_group_ids,attr"`
	// ExternalPolicies: bool, optional
	ExternalPolicies terra.BoolValue `hcl:"external_policies,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MemberEntityIds: set of string, optional
	MemberEntityIds terra.SetValue[terra.StringValue] `hcl:"member_entity_ids,attr"`
	// MemberGroupIds: set of string, optional
	MemberGroupIds terra.SetValue[terra.StringValue] `hcl:"member_group_ids,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Policies: set of string, optional
	Policies terra.SetValue[terra.StringValue] `hcl:"policies,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}
type identityGroupAttributes struct {
	ref terra.Reference
}

// ExternalMemberEntityIds returns a reference to field external_member_entity_ids of vault_identity_group.
func (ig identityGroupAttributes) ExternalMemberEntityIds() terra.BoolValue {
	return terra.ReferenceAsBool(ig.ref.Append("external_member_entity_ids"))
}

// ExternalMemberGroupIds returns a reference to field external_member_group_ids of vault_identity_group.
func (ig identityGroupAttributes) ExternalMemberGroupIds() terra.BoolValue {
	return terra.ReferenceAsBool(ig.ref.Append("external_member_group_ids"))
}

// ExternalPolicies returns a reference to field external_policies of vault_identity_group.
func (ig identityGroupAttributes) ExternalPolicies() terra.BoolValue {
	return terra.ReferenceAsBool(ig.ref.Append("external_policies"))
}

// Id returns a reference to field id of vault_identity_group.
func (ig identityGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("id"))
}

// MemberEntityIds returns a reference to field member_entity_ids of vault_identity_group.
func (ig identityGroupAttributes) MemberEntityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ig.ref.Append("member_entity_ids"))
}

// MemberGroupIds returns a reference to field member_group_ids of vault_identity_group.
func (ig identityGroupAttributes) MemberGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ig.ref.Append("member_group_ids"))
}

// Metadata returns a reference to field metadata of vault_identity_group.
func (ig identityGroupAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ig.ref.Append("metadata"))
}

// Name returns a reference to field name of vault_identity_group.
func (ig identityGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_identity_group.
func (ig identityGroupAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("namespace"))
}

// Policies returns a reference to field policies of vault_identity_group.
func (ig identityGroupAttributes) Policies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ig.ref.Append("policies"))
}

// Type returns a reference to field type of vault_identity_group.
func (ig identityGroupAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("type"))
}

type identityGroupState struct {
	ExternalMemberEntityIds bool              `json:"external_member_entity_ids"`
	ExternalMemberGroupIds  bool              `json:"external_member_group_ids"`
	ExternalPolicies        bool              `json:"external_policies"`
	Id                      string            `json:"id"`
	MemberEntityIds         []string          `json:"member_entity_ids"`
	MemberGroupIds          []string          `json:"member_group_ids"`
	Metadata                map[string]string `json:"metadata"`
	Name                    string            `json:"name"`
	Namespace               string            `json:"namespace"`
	Policies                []string          `json:"policies"`
	Type                    string            `json:"type"`
}
