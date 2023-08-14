// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botchannelmsteams "github.com/golingon/terraproviders/azurerm/3.69.0/botchannelmsteams"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBotChannelMsTeams creates a new instance of [BotChannelMsTeams].
func NewBotChannelMsTeams(name string, args BotChannelMsTeamsArgs) *BotChannelMsTeams {
	return &BotChannelMsTeams{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotChannelMsTeams)(nil)

// BotChannelMsTeams represents the Terraform resource azurerm_bot_channel_ms_teams.
type BotChannelMsTeams struct {
	Name      string
	Args      BotChannelMsTeamsArgs
	state     *botChannelMsTeamsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotChannelMsTeams].
func (bcmt *BotChannelMsTeams) Type() string {
	return "azurerm_bot_channel_ms_teams"
}

// LocalName returns the local name for [BotChannelMsTeams].
func (bcmt *BotChannelMsTeams) LocalName() string {
	return bcmt.Name
}

// Configuration returns the configuration (args) for [BotChannelMsTeams].
func (bcmt *BotChannelMsTeams) Configuration() interface{} {
	return bcmt.Args
}

// DependOn is used for other resources to depend on [BotChannelMsTeams].
func (bcmt *BotChannelMsTeams) DependOn() terra.Reference {
	return terra.ReferenceResource(bcmt)
}

// Dependencies returns the list of resources [BotChannelMsTeams] depends_on.
func (bcmt *BotChannelMsTeams) Dependencies() terra.Dependencies {
	return bcmt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotChannelMsTeams].
func (bcmt *BotChannelMsTeams) LifecycleManagement() *terra.Lifecycle {
	return bcmt.Lifecycle
}

// Attributes returns the attributes for [BotChannelMsTeams].
func (bcmt *BotChannelMsTeams) Attributes() botChannelMsTeamsAttributes {
	return botChannelMsTeamsAttributes{ref: terra.ReferenceResource(bcmt)}
}

// ImportState imports the given attribute values into [BotChannelMsTeams]'s state.
func (bcmt *BotChannelMsTeams) ImportState(av io.Reader) error {
	bcmt.state = &botChannelMsTeamsState{}
	if err := json.NewDecoder(av).Decode(bcmt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcmt.Type(), bcmt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotChannelMsTeams] has state.
func (bcmt *BotChannelMsTeams) State() (*botChannelMsTeamsState, bool) {
	return bcmt.state, bcmt.state != nil
}

// StateMust returns the state for [BotChannelMsTeams]. Panics if the state is nil.
func (bcmt *BotChannelMsTeams) StateMust() *botChannelMsTeamsState {
	if bcmt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcmt.Type(), bcmt.LocalName()))
	}
	return bcmt.state
}

// BotChannelMsTeamsArgs contains the configurations for azurerm_bot_channel_ms_teams.
type BotChannelMsTeamsArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// CallingWebHook: string, optional
	CallingWebHook terra.StringValue `hcl:"calling_web_hook,attr"`
	// EnableCalling: bool, optional
	EnableCalling terra.BoolValue `hcl:"enable_calling,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *botchannelmsteams.Timeouts `hcl:"timeouts,block"`
}
type botChannelMsTeamsAttributes struct {
	ref terra.Reference
}

// BotName returns a reference to field bot_name of azurerm_bot_channel_ms_teams.
func (bcmt botChannelMsTeamsAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(bcmt.ref.Append("bot_name"))
}

// CallingWebHook returns a reference to field calling_web_hook of azurerm_bot_channel_ms_teams.
func (bcmt botChannelMsTeamsAttributes) CallingWebHook() terra.StringValue {
	return terra.ReferenceAsString(bcmt.ref.Append("calling_web_hook"))
}

// EnableCalling returns a reference to field enable_calling of azurerm_bot_channel_ms_teams.
func (bcmt botChannelMsTeamsAttributes) EnableCalling() terra.BoolValue {
	return terra.ReferenceAsBool(bcmt.ref.Append("enable_calling"))
}

// Id returns a reference to field id of azurerm_bot_channel_ms_teams.
func (bcmt botChannelMsTeamsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcmt.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_bot_channel_ms_teams.
func (bcmt botChannelMsTeamsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcmt.ref.Append("location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_channel_ms_teams.
func (bcmt botChannelMsTeamsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bcmt.ref.Append("resource_group_name"))
}

func (bcmt botChannelMsTeamsAttributes) Timeouts() botchannelmsteams.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botchannelmsteams.TimeoutsAttributes](bcmt.ref.Append("timeouts"))
}

type botChannelMsTeamsState struct {
	BotName           string                           `json:"bot_name"`
	CallingWebHook    string                           `json:"calling_web_hook"`
	EnableCalling     bool                             `json:"enable_calling"`
	Id                string                           `json:"id"`
	Location          string                           `json:"location"`
	ResourceGroupName string                           `json:"resource_group_name"`
	Timeouts          *botchannelmsteams.TimeoutsState `json:"timeouts"`
}
