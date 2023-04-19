// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botchannelslack "github.com/golingon/terraproviders/azurerm/3.52.0/botchannelslack"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBotChannelSlack creates a new instance of [BotChannelSlack].
func NewBotChannelSlack(name string, args BotChannelSlackArgs) *BotChannelSlack {
	return &BotChannelSlack{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotChannelSlack)(nil)

// BotChannelSlack represents the Terraform resource azurerm_bot_channel_slack.
type BotChannelSlack struct {
	Name      string
	Args      BotChannelSlackArgs
	state     *botChannelSlackState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotChannelSlack].
func (bcs *BotChannelSlack) Type() string {
	return "azurerm_bot_channel_slack"
}

// LocalName returns the local name for [BotChannelSlack].
func (bcs *BotChannelSlack) LocalName() string {
	return bcs.Name
}

// Configuration returns the configuration (args) for [BotChannelSlack].
func (bcs *BotChannelSlack) Configuration() interface{} {
	return bcs.Args
}

// DependOn is used for other resources to depend on [BotChannelSlack].
func (bcs *BotChannelSlack) DependOn() terra.Reference {
	return terra.ReferenceResource(bcs)
}

// Dependencies returns the list of resources [BotChannelSlack] depends_on.
func (bcs *BotChannelSlack) Dependencies() terra.Dependencies {
	return bcs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotChannelSlack].
func (bcs *BotChannelSlack) LifecycleManagement() *terra.Lifecycle {
	return bcs.Lifecycle
}

// Attributes returns the attributes for [BotChannelSlack].
func (bcs *BotChannelSlack) Attributes() botChannelSlackAttributes {
	return botChannelSlackAttributes{ref: terra.ReferenceResource(bcs)}
}

// ImportState imports the given attribute values into [BotChannelSlack]'s state.
func (bcs *BotChannelSlack) ImportState(av io.Reader) error {
	bcs.state = &botChannelSlackState{}
	if err := json.NewDecoder(av).Decode(bcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcs.Type(), bcs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotChannelSlack] has state.
func (bcs *BotChannelSlack) State() (*botChannelSlackState, bool) {
	return bcs.state, bcs.state != nil
}

// StateMust returns the state for [BotChannelSlack]. Panics if the state is nil.
func (bcs *BotChannelSlack) StateMust() *botChannelSlackState {
	if bcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcs.Type(), bcs.LocalName()))
	}
	return bcs.state
}

// BotChannelSlackArgs contains the configurations for azurerm_bot_channel_slack.
type BotChannelSlackArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LandingPageUrl: string, optional
	LandingPageUrl terra.StringValue `hcl:"landing_page_url,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SigningSecret: string, optional
	SigningSecret terra.StringValue `hcl:"signing_secret,attr"`
	// VerificationToken: string, required
	VerificationToken terra.StringValue `hcl:"verification_token,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *botchannelslack.Timeouts `hcl:"timeouts,block"`
}
type botChannelSlackAttributes struct {
	ref terra.Reference
}

// BotName returns a reference to field bot_name of azurerm_bot_channel_slack.
func (bcs botChannelSlackAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(bcs.ref.Append("bot_name"))
}

// ClientId returns a reference to field client_id of azurerm_bot_channel_slack.
func (bcs botChannelSlackAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(bcs.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of azurerm_bot_channel_slack.
func (bcs botChannelSlackAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(bcs.ref.Append("client_secret"))
}

// Id returns a reference to field id of azurerm_bot_channel_slack.
func (bcs botChannelSlackAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcs.ref.Append("id"))
}

// LandingPageUrl returns a reference to field landing_page_url of azurerm_bot_channel_slack.
func (bcs botChannelSlackAttributes) LandingPageUrl() terra.StringValue {
	return terra.ReferenceAsString(bcs.ref.Append("landing_page_url"))
}

// Location returns a reference to field location of azurerm_bot_channel_slack.
func (bcs botChannelSlackAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcs.ref.Append("location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_channel_slack.
func (bcs botChannelSlackAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bcs.ref.Append("resource_group_name"))
}

// SigningSecret returns a reference to field signing_secret of azurerm_bot_channel_slack.
func (bcs botChannelSlackAttributes) SigningSecret() terra.StringValue {
	return terra.ReferenceAsString(bcs.ref.Append("signing_secret"))
}

// VerificationToken returns a reference to field verification_token of azurerm_bot_channel_slack.
func (bcs botChannelSlackAttributes) VerificationToken() terra.StringValue {
	return terra.ReferenceAsString(bcs.ref.Append("verification_token"))
}

func (bcs botChannelSlackAttributes) Timeouts() botchannelslack.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botchannelslack.TimeoutsAttributes](bcs.ref.Append("timeouts"))
}

type botChannelSlackState struct {
	BotName           string                         `json:"bot_name"`
	ClientId          string                         `json:"client_id"`
	ClientSecret      string                         `json:"client_secret"`
	Id                string                         `json:"id"`
	LandingPageUrl    string                         `json:"landing_page_url"`
	Location          string                         `json:"location"`
	ResourceGroupName string                         `json:"resource_group_name"`
	SigningSecret     string                         `json:"signing_secret"`
	VerificationToken string                         `json:"verification_token"`
	Timeouts          *botchannelslack.TimeoutsState `json:"timeouts"`
}
