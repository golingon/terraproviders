// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudidentitygroupmembership "github.com/golingon/terraproviders/googlebeta/4.62.0/cloudidentitygroupmembership"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudIdentityGroupMembership creates a new instance of [CloudIdentityGroupMembership].
func NewCloudIdentityGroupMembership(name string, args CloudIdentityGroupMembershipArgs) *CloudIdentityGroupMembership {
	return &CloudIdentityGroupMembership{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudIdentityGroupMembership)(nil)

// CloudIdentityGroupMembership represents the Terraform resource google_cloud_identity_group_membership.
type CloudIdentityGroupMembership struct {
	Name      string
	Args      CloudIdentityGroupMembershipArgs
	state     *cloudIdentityGroupMembershipState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudIdentityGroupMembership].
func (cigm *CloudIdentityGroupMembership) Type() string {
	return "google_cloud_identity_group_membership"
}

// LocalName returns the local name for [CloudIdentityGroupMembership].
func (cigm *CloudIdentityGroupMembership) LocalName() string {
	return cigm.Name
}

// Configuration returns the configuration (args) for [CloudIdentityGroupMembership].
func (cigm *CloudIdentityGroupMembership) Configuration() interface{} {
	return cigm.Args
}

// DependOn is used for other resources to depend on [CloudIdentityGroupMembership].
func (cigm *CloudIdentityGroupMembership) DependOn() terra.Reference {
	return terra.ReferenceResource(cigm)
}

// Dependencies returns the list of resources [CloudIdentityGroupMembership] depends_on.
func (cigm *CloudIdentityGroupMembership) Dependencies() terra.Dependencies {
	return cigm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudIdentityGroupMembership].
func (cigm *CloudIdentityGroupMembership) LifecycleManagement() *terra.Lifecycle {
	return cigm.Lifecycle
}

// Attributes returns the attributes for [CloudIdentityGroupMembership].
func (cigm *CloudIdentityGroupMembership) Attributes() cloudIdentityGroupMembershipAttributes {
	return cloudIdentityGroupMembershipAttributes{ref: terra.ReferenceResource(cigm)}
}

// ImportState imports the given attribute values into [CloudIdentityGroupMembership]'s state.
func (cigm *CloudIdentityGroupMembership) ImportState(av io.Reader) error {
	cigm.state = &cloudIdentityGroupMembershipState{}
	if err := json.NewDecoder(av).Decode(cigm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cigm.Type(), cigm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudIdentityGroupMembership] has state.
func (cigm *CloudIdentityGroupMembership) State() (*cloudIdentityGroupMembershipState, bool) {
	return cigm.state, cigm.state != nil
}

// StateMust returns the state for [CloudIdentityGroupMembership]. Panics if the state is nil.
func (cigm *CloudIdentityGroupMembership) StateMust() *cloudIdentityGroupMembershipState {
	if cigm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cigm.Type(), cigm.LocalName()))
	}
	return cigm.state
}

// CloudIdentityGroupMembershipArgs contains the configurations for google_cloud_identity_group_membership.
type CloudIdentityGroupMembershipArgs struct {
	// Group: string, required
	Group terra.StringValue `hcl:"group,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MemberKey: optional
	MemberKey *cloudidentitygroupmembership.MemberKey `hcl:"member_key,block"`
	// PreferredMemberKey: optional
	PreferredMemberKey *cloudidentitygroupmembership.PreferredMemberKey `hcl:"preferred_member_key,block"`
	// Roles: min=1
	Roles []cloudidentitygroupmembership.Roles `hcl:"roles,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *cloudidentitygroupmembership.Timeouts `hcl:"timeouts,block"`
}
type cloudIdentityGroupMembershipAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_cloud_identity_group_membership.
func (cigm cloudIdentityGroupMembershipAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("create_time"))
}

// Group returns a reference to field group of google_cloud_identity_group_membership.
func (cigm cloudIdentityGroupMembershipAttributes) Group() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("group"))
}

// Id returns a reference to field id of google_cloud_identity_group_membership.
func (cigm cloudIdentityGroupMembershipAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("id"))
}

// Name returns a reference to field name of google_cloud_identity_group_membership.
func (cigm cloudIdentityGroupMembershipAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("name"))
}

// Type returns a reference to field type of google_cloud_identity_group_membership.
func (cigm cloudIdentityGroupMembershipAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("type"))
}

// UpdateTime returns a reference to field update_time of google_cloud_identity_group_membership.
func (cigm cloudIdentityGroupMembershipAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("update_time"))
}

func (cigm cloudIdentityGroupMembershipAttributes) MemberKey() terra.ListValue[cloudidentitygroupmembership.MemberKeyAttributes] {
	return terra.ReferenceAsList[cloudidentitygroupmembership.MemberKeyAttributes](cigm.ref.Append("member_key"))
}

func (cigm cloudIdentityGroupMembershipAttributes) PreferredMemberKey() terra.ListValue[cloudidentitygroupmembership.PreferredMemberKeyAttributes] {
	return terra.ReferenceAsList[cloudidentitygroupmembership.PreferredMemberKeyAttributes](cigm.ref.Append("preferred_member_key"))
}

func (cigm cloudIdentityGroupMembershipAttributes) Roles() terra.SetValue[cloudidentitygroupmembership.RolesAttributes] {
	return terra.ReferenceAsSet[cloudidentitygroupmembership.RolesAttributes](cigm.ref.Append("roles"))
}

func (cigm cloudIdentityGroupMembershipAttributes) Timeouts() cloudidentitygroupmembership.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudidentitygroupmembership.TimeoutsAttributes](cigm.ref.Append("timeouts"))
}

type cloudIdentityGroupMembershipState struct {
	CreateTime         string                                                 `json:"create_time"`
	Group              string                                                 `json:"group"`
	Id                 string                                                 `json:"id"`
	Name               string                                                 `json:"name"`
	Type               string                                                 `json:"type"`
	UpdateTime         string                                                 `json:"update_time"`
	MemberKey          []cloudidentitygroupmembership.MemberKeyState          `json:"member_key"`
	PreferredMemberKey []cloudidentitygroupmembership.PreferredMemberKeyState `json:"preferred_member_key"`
	Roles              []cloudidentitygroupmembership.RolesState              `json:"roles"`
	Timeouts           *cloudidentitygroupmembership.TimeoutsState            `json:"timeouts"`
}
