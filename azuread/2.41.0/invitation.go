// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	invitation "github.com/golingon/terraproviders/azuread/2.41.0/invitation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewInvitation creates a new instance of [Invitation].
func NewInvitation(name string, args InvitationArgs) *Invitation {
	return &Invitation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Invitation)(nil)

// Invitation represents the Terraform resource azuread_invitation.
type Invitation struct {
	Name      string
	Args      InvitationArgs
	state     *invitationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Invitation].
func (i *Invitation) Type() string {
	return "azuread_invitation"
}

// LocalName returns the local name for [Invitation].
func (i *Invitation) LocalName() string {
	return i.Name
}

// Configuration returns the configuration (args) for [Invitation].
func (i *Invitation) Configuration() interface{} {
	return i.Args
}

// DependOn is used for other resources to depend on [Invitation].
func (i *Invitation) DependOn() terra.Reference {
	return terra.ReferenceResource(i)
}

// Dependencies returns the list of resources [Invitation] depends_on.
func (i *Invitation) Dependencies() terra.Dependencies {
	return i.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Invitation].
func (i *Invitation) LifecycleManagement() *terra.Lifecycle {
	return i.Lifecycle
}

// Attributes returns the attributes for [Invitation].
func (i *Invitation) Attributes() invitationAttributes {
	return invitationAttributes{ref: terra.ReferenceResource(i)}
}

// ImportState imports the given attribute values into [Invitation]'s state.
func (i *Invitation) ImportState(av io.Reader) error {
	i.state = &invitationState{}
	if err := json.NewDecoder(av).Decode(i.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", i.Type(), i.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Invitation] has state.
func (i *Invitation) State() (*invitationState, bool) {
	return i.state, i.state != nil
}

// StateMust returns the state for [Invitation]. Panics if the state is nil.
func (i *Invitation) StateMust() *invitationState {
	if i.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", i.Type(), i.LocalName()))
	}
	return i.state
}

// InvitationArgs contains the configurations for azuread_invitation.
type InvitationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RedirectUrl: string, required
	RedirectUrl terra.StringValue `hcl:"redirect_url,attr" validate:"required"`
	// UserDisplayName: string, optional
	UserDisplayName terra.StringValue `hcl:"user_display_name,attr"`
	// UserEmailAddress: string, required
	UserEmailAddress terra.StringValue `hcl:"user_email_address,attr" validate:"required"`
	// UserType: string, optional
	UserType terra.StringValue `hcl:"user_type,attr"`
	// Message: optional
	Message *invitation.Message `hcl:"message,block"`
	// Timeouts: optional
	Timeouts *invitation.Timeouts `hcl:"timeouts,block"`
}
type invitationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azuread_invitation.
func (i invitationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("id"))
}

// RedeemUrl returns a reference to field redeem_url of azuread_invitation.
func (i invitationAttributes) RedeemUrl() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("redeem_url"))
}

// RedirectUrl returns a reference to field redirect_url of azuread_invitation.
func (i invitationAttributes) RedirectUrl() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("redirect_url"))
}

// UserDisplayName returns a reference to field user_display_name of azuread_invitation.
func (i invitationAttributes) UserDisplayName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("user_display_name"))
}

// UserEmailAddress returns a reference to field user_email_address of azuread_invitation.
func (i invitationAttributes) UserEmailAddress() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("user_email_address"))
}

// UserId returns a reference to field user_id of azuread_invitation.
func (i invitationAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("user_id"))
}

// UserType returns a reference to field user_type of azuread_invitation.
func (i invitationAttributes) UserType() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("user_type"))
}

func (i invitationAttributes) Message() terra.ListValue[invitation.MessageAttributes] {
	return terra.ReferenceAsList[invitation.MessageAttributes](i.ref.Append("message"))
}

func (i invitationAttributes) Timeouts() invitation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[invitation.TimeoutsAttributes](i.ref.Append("timeouts"))
}

type invitationState struct {
	Id               string                    `json:"id"`
	RedeemUrl        string                    `json:"redeem_url"`
	RedirectUrl      string                    `json:"redirect_url"`
	UserDisplayName  string                    `json:"user_display_name"`
	UserEmailAddress string                    `json:"user_email_address"`
	UserId           string                    `json:"user_id"`
	UserType         string                    `json:"user_type"`
	Message          []invitation.MessageState `json:"message"`
	Timeouts         *invitation.TimeoutsState `json:"timeouts"`
}
