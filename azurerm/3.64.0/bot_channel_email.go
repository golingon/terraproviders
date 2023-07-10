// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botchannelemail "github.com/golingon/terraproviders/azurerm/3.64.0/botchannelemail"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBotChannelEmail creates a new instance of [BotChannelEmail].
func NewBotChannelEmail(name string, args BotChannelEmailArgs) *BotChannelEmail {
	return &BotChannelEmail{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotChannelEmail)(nil)

// BotChannelEmail represents the Terraform resource azurerm_bot_channel_email.
type BotChannelEmail struct {
	Name      string
	Args      BotChannelEmailArgs
	state     *botChannelEmailState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotChannelEmail].
func (bce *BotChannelEmail) Type() string {
	return "azurerm_bot_channel_email"
}

// LocalName returns the local name for [BotChannelEmail].
func (bce *BotChannelEmail) LocalName() string {
	return bce.Name
}

// Configuration returns the configuration (args) for [BotChannelEmail].
func (bce *BotChannelEmail) Configuration() interface{} {
	return bce.Args
}

// DependOn is used for other resources to depend on [BotChannelEmail].
func (bce *BotChannelEmail) DependOn() terra.Reference {
	return terra.ReferenceResource(bce)
}

// Dependencies returns the list of resources [BotChannelEmail] depends_on.
func (bce *BotChannelEmail) Dependencies() terra.Dependencies {
	return bce.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotChannelEmail].
func (bce *BotChannelEmail) LifecycleManagement() *terra.Lifecycle {
	return bce.Lifecycle
}

// Attributes returns the attributes for [BotChannelEmail].
func (bce *BotChannelEmail) Attributes() botChannelEmailAttributes {
	return botChannelEmailAttributes{ref: terra.ReferenceResource(bce)}
}

// ImportState imports the given attribute values into [BotChannelEmail]'s state.
func (bce *BotChannelEmail) ImportState(av io.Reader) error {
	bce.state = &botChannelEmailState{}
	if err := json.NewDecoder(av).Decode(bce.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bce.Type(), bce.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotChannelEmail] has state.
func (bce *BotChannelEmail) State() (*botChannelEmailState, bool) {
	return bce.state, bce.state != nil
}

// StateMust returns the state for [BotChannelEmail]. Panics if the state is nil.
func (bce *BotChannelEmail) StateMust() *botChannelEmailState {
	if bce.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bce.Type(), bce.LocalName()))
	}
	return bce.state
}

// BotChannelEmailArgs contains the configurations for azurerm_bot_channel_email.
type BotChannelEmailArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// EmailAddress: string, required
	EmailAddress terra.StringValue `hcl:"email_address,attr" validate:"required"`
	// EmailPassword: string, required
	EmailPassword terra.StringValue `hcl:"email_password,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *botchannelemail.Timeouts `hcl:"timeouts,block"`
}
type botChannelEmailAttributes struct {
	ref terra.Reference
}

// BotName returns a reference to field bot_name of azurerm_bot_channel_email.
func (bce botChannelEmailAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("bot_name"))
}

// EmailAddress returns a reference to field email_address of azurerm_bot_channel_email.
func (bce botChannelEmailAttributes) EmailAddress() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("email_address"))
}

// EmailPassword returns a reference to field email_password of azurerm_bot_channel_email.
func (bce botChannelEmailAttributes) EmailPassword() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("email_password"))
}

// Id returns a reference to field id of azurerm_bot_channel_email.
func (bce botChannelEmailAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_bot_channel_email.
func (bce botChannelEmailAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_channel_email.
func (bce botChannelEmailAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("resource_group_name"))
}

func (bce botChannelEmailAttributes) Timeouts() botchannelemail.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botchannelemail.TimeoutsAttributes](bce.ref.Append("timeouts"))
}

type botChannelEmailState struct {
	BotName           string                         `json:"bot_name"`
	EmailAddress      string                         `json:"email_address"`
	EmailPassword     string                         `json:"email_password"`
	Id                string                         `json:"id"`
	Location          string                         `json:"location"`
	ResourceGroupName string                         `json:"resource_group_name"`
	Timeouts          *botchannelemail.TimeoutsState `json:"timeouts"`
}
