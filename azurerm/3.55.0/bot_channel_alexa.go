// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botchannelalexa "github.com/golingon/terraproviders/azurerm/3.55.0/botchannelalexa"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBotChannelAlexa creates a new instance of [BotChannelAlexa].
func NewBotChannelAlexa(name string, args BotChannelAlexaArgs) *BotChannelAlexa {
	return &BotChannelAlexa{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotChannelAlexa)(nil)

// BotChannelAlexa represents the Terraform resource azurerm_bot_channel_alexa.
type BotChannelAlexa struct {
	Name      string
	Args      BotChannelAlexaArgs
	state     *botChannelAlexaState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotChannelAlexa].
func (bca *BotChannelAlexa) Type() string {
	return "azurerm_bot_channel_alexa"
}

// LocalName returns the local name for [BotChannelAlexa].
func (bca *BotChannelAlexa) LocalName() string {
	return bca.Name
}

// Configuration returns the configuration (args) for [BotChannelAlexa].
func (bca *BotChannelAlexa) Configuration() interface{} {
	return bca.Args
}

// DependOn is used for other resources to depend on [BotChannelAlexa].
func (bca *BotChannelAlexa) DependOn() terra.Reference {
	return terra.ReferenceResource(bca)
}

// Dependencies returns the list of resources [BotChannelAlexa] depends_on.
func (bca *BotChannelAlexa) Dependencies() terra.Dependencies {
	return bca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotChannelAlexa].
func (bca *BotChannelAlexa) LifecycleManagement() *terra.Lifecycle {
	return bca.Lifecycle
}

// Attributes returns the attributes for [BotChannelAlexa].
func (bca *BotChannelAlexa) Attributes() botChannelAlexaAttributes {
	return botChannelAlexaAttributes{ref: terra.ReferenceResource(bca)}
}

// ImportState imports the given attribute values into [BotChannelAlexa]'s state.
func (bca *BotChannelAlexa) ImportState(av io.Reader) error {
	bca.state = &botChannelAlexaState{}
	if err := json.NewDecoder(av).Decode(bca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bca.Type(), bca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotChannelAlexa] has state.
func (bca *BotChannelAlexa) State() (*botChannelAlexaState, bool) {
	return bca.state, bca.state != nil
}

// StateMust returns the state for [BotChannelAlexa]. Panics if the state is nil.
func (bca *BotChannelAlexa) StateMust() *botChannelAlexaState {
	if bca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bca.Type(), bca.LocalName()))
	}
	return bca.state
}

// BotChannelAlexaArgs contains the configurations for azurerm_bot_channel_alexa.
type BotChannelAlexaArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkillId: string, required
	SkillId terra.StringValue `hcl:"skill_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *botchannelalexa.Timeouts `hcl:"timeouts,block"`
}
type botChannelAlexaAttributes struct {
	ref terra.Reference
}

// BotName returns a reference to field bot_name of azurerm_bot_channel_alexa.
func (bca botChannelAlexaAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(bca.ref.Append("bot_name"))
}

// Id returns a reference to field id of azurerm_bot_channel_alexa.
func (bca botChannelAlexaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bca.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_bot_channel_alexa.
func (bca botChannelAlexaAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bca.ref.Append("location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_channel_alexa.
func (bca botChannelAlexaAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bca.ref.Append("resource_group_name"))
}

// SkillId returns a reference to field skill_id of azurerm_bot_channel_alexa.
func (bca botChannelAlexaAttributes) SkillId() terra.StringValue {
	return terra.ReferenceAsString(bca.ref.Append("skill_id"))
}

func (bca botChannelAlexaAttributes) Timeouts() botchannelalexa.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botchannelalexa.TimeoutsAttributes](bca.ref.Append("timeouts"))
}

type botChannelAlexaState struct {
	BotName           string                         `json:"bot_name"`
	Id                string                         `json:"id"`
	Location          string                         `json:"location"`
	ResourceGroupName string                         `json:"resource_group_name"`
	SkillId           string                         `json:"skill_id"`
	Timeouts          *botchannelalexa.TimeoutsState `json:"timeouts"`
}
