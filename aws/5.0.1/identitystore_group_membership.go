// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentitystoreGroupMembership creates a new instance of [IdentitystoreGroupMembership].
func NewIdentitystoreGroupMembership(name string, args IdentitystoreGroupMembershipArgs) *IdentitystoreGroupMembership {
	return &IdentitystoreGroupMembership{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentitystoreGroupMembership)(nil)

// IdentitystoreGroupMembership represents the Terraform resource aws_identitystore_group_membership.
type IdentitystoreGroupMembership struct {
	Name      string
	Args      IdentitystoreGroupMembershipArgs
	state     *identitystoreGroupMembershipState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentitystoreGroupMembership].
func (igm *IdentitystoreGroupMembership) Type() string {
	return "aws_identitystore_group_membership"
}

// LocalName returns the local name for [IdentitystoreGroupMembership].
func (igm *IdentitystoreGroupMembership) LocalName() string {
	return igm.Name
}

// Configuration returns the configuration (args) for [IdentitystoreGroupMembership].
func (igm *IdentitystoreGroupMembership) Configuration() interface{} {
	return igm.Args
}

// DependOn is used for other resources to depend on [IdentitystoreGroupMembership].
func (igm *IdentitystoreGroupMembership) DependOn() terra.Reference {
	return terra.ReferenceResource(igm)
}

// Dependencies returns the list of resources [IdentitystoreGroupMembership] depends_on.
func (igm *IdentitystoreGroupMembership) Dependencies() terra.Dependencies {
	return igm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentitystoreGroupMembership].
func (igm *IdentitystoreGroupMembership) LifecycleManagement() *terra.Lifecycle {
	return igm.Lifecycle
}

// Attributes returns the attributes for [IdentitystoreGroupMembership].
func (igm *IdentitystoreGroupMembership) Attributes() identitystoreGroupMembershipAttributes {
	return identitystoreGroupMembershipAttributes{ref: terra.ReferenceResource(igm)}
}

// ImportState imports the given attribute values into [IdentitystoreGroupMembership]'s state.
func (igm *IdentitystoreGroupMembership) ImportState(av io.Reader) error {
	igm.state = &identitystoreGroupMembershipState{}
	if err := json.NewDecoder(av).Decode(igm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", igm.Type(), igm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentitystoreGroupMembership] has state.
func (igm *IdentitystoreGroupMembership) State() (*identitystoreGroupMembershipState, bool) {
	return igm.state, igm.state != nil
}

// StateMust returns the state for [IdentitystoreGroupMembership]. Panics if the state is nil.
func (igm *IdentitystoreGroupMembership) StateMust() *identitystoreGroupMembershipState {
	if igm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", igm.Type(), igm.LocalName()))
	}
	return igm.state
}

// IdentitystoreGroupMembershipArgs contains the configurations for aws_identitystore_group_membership.
type IdentitystoreGroupMembershipArgs struct {
	// GroupId: string, required
	GroupId terra.StringValue `hcl:"group_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityStoreId: string, required
	IdentityStoreId terra.StringValue `hcl:"identity_store_id,attr" validate:"required"`
	// MemberId: string, required
	MemberId terra.StringValue `hcl:"member_id,attr" validate:"required"`
}
type identitystoreGroupMembershipAttributes struct {
	ref terra.Reference
}

// GroupId returns a reference to field group_id of aws_identitystore_group_membership.
func (igm identitystoreGroupMembershipAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(igm.ref.Append("group_id"))
}

// Id returns a reference to field id of aws_identitystore_group_membership.
func (igm identitystoreGroupMembershipAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(igm.ref.Append("id"))
}

// IdentityStoreId returns a reference to field identity_store_id of aws_identitystore_group_membership.
func (igm identitystoreGroupMembershipAttributes) IdentityStoreId() terra.StringValue {
	return terra.ReferenceAsString(igm.ref.Append("identity_store_id"))
}

// MemberId returns a reference to field member_id of aws_identitystore_group_membership.
func (igm identitystoreGroupMembershipAttributes) MemberId() terra.StringValue {
	return terra.ReferenceAsString(igm.ref.Append("member_id"))
}

// MembershipId returns a reference to field membership_id of aws_identitystore_group_membership.
func (igm identitystoreGroupMembershipAttributes) MembershipId() terra.StringValue {
	return terra.ReferenceAsString(igm.ref.Append("membership_id"))
}

type identitystoreGroupMembershipState struct {
	GroupId         string `json:"group_id"`
	Id              string `json:"id"`
	IdentityStoreId string `json:"identity_store_id"`
	MemberId        string `json:"member_id"`
	MembershipId    string `json:"membership_id"`
}
