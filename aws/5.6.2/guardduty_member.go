// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	guarddutymember "github.com/golingon/terraproviders/aws/5.6.2/guarddutymember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGuarddutyMember creates a new instance of [GuarddutyMember].
func NewGuarddutyMember(name string, args GuarddutyMemberArgs) *GuarddutyMember {
	return &GuarddutyMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GuarddutyMember)(nil)

// GuarddutyMember represents the Terraform resource aws_guardduty_member.
type GuarddutyMember struct {
	Name      string
	Args      GuarddutyMemberArgs
	state     *guarddutyMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GuarddutyMember].
func (gm *GuarddutyMember) Type() string {
	return "aws_guardduty_member"
}

// LocalName returns the local name for [GuarddutyMember].
func (gm *GuarddutyMember) LocalName() string {
	return gm.Name
}

// Configuration returns the configuration (args) for [GuarddutyMember].
func (gm *GuarddutyMember) Configuration() interface{} {
	return gm.Args
}

// DependOn is used for other resources to depend on [GuarddutyMember].
func (gm *GuarddutyMember) DependOn() terra.Reference {
	return terra.ReferenceResource(gm)
}

// Dependencies returns the list of resources [GuarddutyMember] depends_on.
func (gm *GuarddutyMember) Dependencies() terra.Dependencies {
	return gm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GuarddutyMember].
func (gm *GuarddutyMember) LifecycleManagement() *terra.Lifecycle {
	return gm.Lifecycle
}

// Attributes returns the attributes for [GuarddutyMember].
func (gm *GuarddutyMember) Attributes() guarddutyMemberAttributes {
	return guarddutyMemberAttributes{ref: terra.ReferenceResource(gm)}
}

// ImportState imports the given attribute values into [GuarddutyMember]'s state.
func (gm *GuarddutyMember) ImportState(av io.Reader) error {
	gm.state = &guarddutyMemberState{}
	if err := json.NewDecoder(av).Decode(gm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gm.Type(), gm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GuarddutyMember] has state.
func (gm *GuarddutyMember) State() (*guarddutyMemberState, bool) {
	return gm.state, gm.state != nil
}

// StateMust returns the state for [GuarddutyMember]. Panics if the state is nil.
func (gm *GuarddutyMember) StateMust() *guarddutyMemberState {
	if gm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gm.Type(), gm.LocalName()))
	}
	return gm.state
}

// GuarddutyMemberArgs contains the configurations for aws_guardduty_member.
type GuarddutyMemberArgs struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
	// DetectorId: string, required
	DetectorId terra.StringValue `hcl:"detector_id,attr" validate:"required"`
	// DisableEmailNotification: bool, optional
	DisableEmailNotification terra.BoolValue `hcl:"disable_email_notification,attr"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InvitationMessage: string, optional
	InvitationMessage terra.StringValue `hcl:"invitation_message,attr"`
	// Invite: bool, optional
	Invite terra.BoolValue `hcl:"invite,attr"`
	// Timeouts: optional
	Timeouts *guarddutymember.Timeouts `hcl:"timeouts,block"`
}
type guarddutyMemberAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_guardduty_member.
func (gm guarddutyMemberAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("account_id"))
}

// DetectorId returns a reference to field detector_id of aws_guardduty_member.
func (gm guarddutyMemberAttributes) DetectorId() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("detector_id"))
}

// DisableEmailNotification returns a reference to field disable_email_notification of aws_guardduty_member.
func (gm guarddutyMemberAttributes) DisableEmailNotification() terra.BoolValue {
	return terra.ReferenceAsBool(gm.ref.Append("disable_email_notification"))
}

// Email returns a reference to field email of aws_guardduty_member.
func (gm guarddutyMemberAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("email"))
}

// Id returns a reference to field id of aws_guardduty_member.
func (gm guarddutyMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("id"))
}

// InvitationMessage returns a reference to field invitation_message of aws_guardduty_member.
func (gm guarddutyMemberAttributes) InvitationMessage() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("invitation_message"))
}

// Invite returns a reference to field invite of aws_guardduty_member.
func (gm guarddutyMemberAttributes) Invite() terra.BoolValue {
	return terra.ReferenceAsBool(gm.ref.Append("invite"))
}

// RelationshipStatus returns a reference to field relationship_status of aws_guardduty_member.
func (gm guarddutyMemberAttributes) RelationshipStatus() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("relationship_status"))
}

func (gm guarddutyMemberAttributes) Timeouts() guarddutymember.TimeoutsAttributes {
	return terra.ReferenceAsSingle[guarddutymember.TimeoutsAttributes](gm.ref.Append("timeouts"))
}

type guarddutyMemberState struct {
	AccountId                string                         `json:"account_id"`
	DetectorId               string                         `json:"detector_id"`
	DisableEmailNotification bool                           `json:"disable_email_notification"`
	Email                    string                         `json:"email"`
	Id                       string                         `json:"id"`
	InvitationMessage        string                         `json:"invitation_message"`
	Invite                   bool                           `json:"invite"`
	RelationshipStatus       string                         `json:"relationship_status"`
	Timeouts                 *guarddutymember.TimeoutsState `json:"timeouts"`
}
