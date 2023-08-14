// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityhubMember creates a new instance of [SecurityhubMember].
func NewSecurityhubMember(name string, args SecurityhubMemberArgs) *SecurityhubMember {
	return &SecurityhubMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityhubMember)(nil)

// SecurityhubMember represents the Terraform resource aws_securityhub_member.
type SecurityhubMember struct {
	Name      string
	Args      SecurityhubMemberArgs
	state     *securityhubMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityhubMember].
func (sm *SecurityhubMember) Type() string {
	return "aws_securityhub_member"
}

// LocalName returns the local name for [SecurityhubMember].
func (sm *SecurityhubMember) LocalName() string {
	return sm.Name
}

// Configuration returns the configuration (args) for [SecurityhubMember].
func (sm *SecurityhubMember) Configuration() interface{} {
	return sm.Args
}

// DependOn is used for other resources to depend on [SecurityhubMember].
func (sm *SecurityhubMember) DependOn() terra.Reference {
	return terra.ReferenceResource(sm)
}

// Dependencies returns the list of resources [SecurityhubMember] depends_on.
func (sm *SecurityhubMember) Dependencies() terra.Dependencies {
	return sm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityhubMember].
func (sm *SecurityhubMember) LifecycleManagement() *terra.Lifecycle {
	return sm.Lifecycle
}

// Attributes returns the attributes for [SecurityhubMember].
func (sm *SecurityhubMember) Attributes() securityhubMemberAttributes {
	return securityhubMemberAttributes{ref: terra.ReferenceResource(sm)}
}

// ImportState imports the given attribute values into [SecurityhubMember]'s state.
func (sm *SecurityhubMember) ImportState(av io.Reader) error {
	sm.state = &securityhubMemberState{}
	if err := json.NewDecoder(av).Decode(sm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sm.Type(), sm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityhubMember] has state.
func (sm *SecurityhubMember) State() (*securityhubMemberState, bool) {
	return sm.state, sm.state != nil
}

// StateMust returns the state for [SecurityhubMember]. Panics if the state is nil.
func (sm *SecurityhubMember) StateMust() *securityhubMemberState {
	if sm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sm.Type(), sm.LocalName()))
	}
	return sm.state
}

// SecurityhubMemberArgs contains the configurations for aws_securityhub_member.
type SecurityhubMemberArgs struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Invite: bool, optional
	Invite terra.BoolValue `hcl:"invite,attr"`
}
type securityhubMemberAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_securityhub_member.
func (sm securityhubMemberAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("account_id"))
}

// Email returns a reference to field email of aws_securityhub_member.
func (sm securityhubMemberAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("email"))
}

// Id returns a reference to field id of aws_securityhub_member.
func (sm securityhubMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("id"))
}

// Invite returns a reference to field invite of aws_securityhub_member.
func (sm securityhubMemberAttributes) Invite() terra.BoolValue {
	return terra.ReferenceAsBool(sm.ref.Append("invite"))
}

// MasterId returns a reference to field master_id of aws_securityhub_member.
func (sm securityhubMemberAttributes) MasterId() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("master_id"))
}

// MemberStatus returns a reference to field member_status of aws_securityhub_member.
func (sm securityhubMemberAttributes) MemberStatus() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("member_status"))
}

type securityhubMemberState struct {
	AccountId    string `json:"account_id"`
	Email        string `json:"email"`
	Id           string `json:"id"`
	Invite       bool   `json:"invite"`
	MasterId     string `json:"master_id"`
	MemberStatus string `json:"member_status"`
}