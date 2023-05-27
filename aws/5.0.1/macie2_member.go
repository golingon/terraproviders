// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	macie2member "github.com/golingon/terraproviders/aws/5.0.1/macie2member"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMacie2Member creates a new instance of [Macie2Member].
func NewMacie2Member(name string, args Macie2MemberArgs) *Macie2Member {
	return &Macie2Member{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Macie2Member)(nil)

// Macie2Member represents the Terraform resource aws_macie2_member.
type Macie2Member struct {
	Name      string
	Args      Macie2MemberArgs
	state     *macie2MemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Macie2Member].
func (mm *Macie2Member) Type() string {
	return "aws_macie2_member"
}

// LocalName returns the local name for [Macie2Member].
func (mm *Macie2Member) LocalName() string {
	return mm.Name
}

// Configuration returns the configuration (args) for [Macie2Member].
func (mm *Macie2Member) Configuration() interface{} {
	return mm.Args
}

// DependOn is used for other resources to depend on [Macie2Member].
func (mm *Macie2Member) DependOn() terra.Reference {
	return terra.ReferenceResource(mm)
}

// Dependencies returns the list of resources [Macie2Member] depends_on.
func (mm *Macie2Member) Dependencies() terra.Dependencies {
	return mm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Macie2Member].
func (mm *Macie2Member) LifecycleManagement() *terra.Lifecycle {
	return mm.Lifecycle
}

// Attributes returns the attributes for [Macie2Member].
func (mm *Macie2Member) Attributes() macie2MemberAttributes {
	return macie2MemberAttributes{ref: terra.ReferenceResource(mm)}
}

// ImportState imports the given attribute values into [Macie2Member]'s state.
func (mm *Macie2Member) ImportState(av io.Reader) error {
	mm.state = &macie2MemberState{}
	if err := json.NewDecoder(av).Decode(mm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mm.Type(), mm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Macie2Member] has state.
func (mm *Macie2Member) State() (*macie2MemberState, bool) {
	return mm.state, mm.state != nil
}

// StateMust returns the state for [Macie2Member]. Panics if the state is nil.
func (mm *Macie2Member) StateMust() *macie2MemberState {
	if mm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mm.Type(), mm.LocalName()))
	}
	return mm.state
}

// Macie2MemberArgs contains the configurations for aws_macie2_member.
type Macie2MemberArgs struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InvitationDisableEmailNotification: bool, optional
	InvitationDisableEmailNotification terra.BoolValue `hcl:"invitation_disable_email_notification,attr"`
	// InvitationMessage: string, optional
	InvitationMessage terra.StringValue `hcl:"invitation_message,attr"`
	// Invite: bool, optional
	Invite terra.BoolValue `hcl:"invite,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *macie2member.Timeouts `hcl:"timeouts,block"`
}
type macie2MemberAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_macie2_member.
func (mm macie2MemberAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("account_id"))
}

// AdministratorAccountId returns a reference to field administrator_account_id of aws_macie2_member.
func (mm macie2MemberAttributes) AdministratorAccountId() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("administrator_account_id"))
}

// Arn returns a reference to field arn of aws_macie2_member.
func (mm macie2MemberAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("arn"))
}

// Email returns a reference to field email of aws_macie2_member.
func (mm macie2MemberAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("email"))
}

// Id returns a reference to field id of aws_macie2_member.
func (mm macie2MemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("id"))
}

// InvitationDisableEmailNotification returns a reference to field invitation_disable_email_notification of aws_macie2_member.
func (mm macie2MemberAttributes) InvitationDisableEmailNotification() terra.BoolValue {
	return terra.ReferenceAsBool(mm.ref.Append("invitation_disable_email_notification"))
}

// InvitationMessage returns a reference to field invitation_message of aws_macie2_member.
func (mm macie2MemberAttributes) InvitationMessage() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("invitation_message"))
}

// Invite returns a reference to field invite of aws_macie2_member.
func (mm macie2MemberAttributes) Invite() terra.BoolValue {
	return terra.ReferenceAsBool(mm.ref.Append("invite"))
}

// InvitedAt returns a reference to field invited_at of aws_macie2_member.
func (mm macie2MemberAttributes) InvitedAt() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("invited_at"))
}

// MasterAccountId returns a reference to field master_account_id of aws_macie2_member.
func (mm macie2MemberAttributes) MasterAccountId() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("master_account_id"))
}

// RelationshipStatus returns a reference to field relationship_status of aws_macie2_member.
func (mm macie2MemberAttributes) RelationshipStatus() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("relationship_status"))
}

// Status returns a reference to field status of aws_macie2_member.
func (mm macie2MemberAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_macie2_member.
func (mm macie2MemberAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mm.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_macie2_member.
func (mm macie2MemberAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mm.ref.Append("tags_all"))
}

// UpdatedAt returns a reference to field updated_at of aws_macie2_member.
func (mm macie2MemberAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("updated_at"))
}

func (mm macie2MemberAttributes) Timeouts() macie2member.TimeoutsAttributes {
	return terra.ReferenceAsSingle[macie2member.TimeoutsAttributes](mm.ref.Append("timeouts"))
}

type macie2MemberState struct {
	AccountId                          string                      `json:"account_id"`
	AdministratorAccountId             string                      `json:"administrator_account_id"`
	Arn                                string                      `json:"arn"`
	Email                              string                      `json:"email"`
	Id                                 string                      `json:"id"`
	InvitationDisableEmailNotification bool                        `json:"invitation_disable_email_notification"`
	InvitationMessage                  string                      `json:"invitation_message"`
	Invite                             bool                        `json:"invite"`
	InvitedAt                          string                      `json:"invited_at"`
	MasterAccountId                    string                      `json:"master_account_id"`
	RelationshipStatus                 string                      `json:"relationship_status"`
	Status                             string                      `json:"status"`
	Tags                               map[string]string           `json:"tags"`
	TagsAll                            map[string]string           `json:"tags_all"`
	UpdatedAt                          string                      `json:"updated_at"`
	Timeouts                           *macie2member.TimeoutsState `json:"timeouts"`
}
