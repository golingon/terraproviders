// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_bot_channel_ms_teams

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_bot_channel_ms_teams.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermBotChannelMsTeamsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (abcmt *Resource) Type() string {
	return "azurerm_bot_channel_ms_teams"
}

// LocalName returns the local name for [Resource].
func (abcmt *Resource) LocalName() string {
	return abcmt.Name
}

// Configuration returns the configuration (args) for [Resource].
func (abcmt *Resource) Configuration() interface{} {
	return abcmt.Args
}

// DependOn is used for other resources to depend on [Resource].
func (abcmt *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(abcmt)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (abcmt *Resource) Dependencies() terra.Dependencies {
	return abcmt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (abcmt *Resource) LifecycleManagement() *terra.Lifecycle {
	return abcmt.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (abcmt *Resource) Attributes() azurermBotChannelMsTeamsAttributes {
	return azurermBotChannelMsTeamsAttributes{ref: terra.ReferenceResource(abcmt)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (abcmt *Resource) ImportState(state io.Reader) error {
	abcmt.state = &azurermBotChannelMsTeamsState{}
	if err := json.NewDecoder(state).Decode(abcmt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", abcmt.Type(), abcmt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (abcmt *Resource) State() (*azurermBotChannelMsTeamsState, bool) {
	return abcmt.state, abcmt.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (abcmt *Resource) StateMust() *azurermBotChannelMsTeamsState {
	if abcmt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", abcmt.Type(), abcmt.LocalName()))
	}
	return abcmt.state
}

// Args contains the configurations for azurerm_bot_channel_ms_teams.
type Args struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// CallingWebHook: string, optional
	CallingWebHook terra.StringValue `hcl:"calling_web_hook,attr"`
	// DeploymentEnvironment: string, optional
	DeploymentEnvironment terra.StringValue `hcl:"deployment_environment,attr"`
	// EnableCalling: bool, optional
	EnableCalling terra.BoolValue `hcl:"enable_calling,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermBotChannelMsTeamsAttributes struct {
	ref terra.Reference
}

// BotName returns a reference to field bot_name of azurerm_bot_channel_ms_teams.
func (abcmt azurermBotChannelMsTeamsAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(abcmt.ref.Append("bot_name"))
}

// CallingWebHook returns a reference to field calling_web_hook of azurerm_bot_channel_ms_teams.
func (abcmt azurermBotChannelMsTeamsAttributes) CallingWebHook() terra.StringValue {
	return terra.ReferenceAsString(abcmt.ref.Append("calling_web_hook"))
}

// DeploymentEnvironment returns a reference to field deployment_environment of azurerm_bot_channel_ms_teams.
func (abcmt azurermBotChannelMsTeamsAttributes) DeploymentEnvironment() terra.StringValue {
	return terra.ReferenceAsString(abcmt.ref.Append("deployment_environment"))
}

// EnableCalling returns a reference to field enable_calling of azurerm_bot_channel_ms_teams.
func (abcmt azurermBotChannelMsTeamsAttributes) EnableCalling() terra.BoolValue {
	return terra.ReferenceAsBool(abcmt.ref.Append("enable_calling"))
}

// Id returns a reference to field id of azurerm_bot_channel_ms_teams.
func (abcmt azurermBotChannelMsTeamsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(abcmt.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_bot_channel_ms_teams.
func (abcmt azurermBotChannelMsTeamsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(abcmt.ref.Append("location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_channel_ms_teams.
func (abcmt azurermBotChannelMsTeamsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(abcmt.ref.Append("resource_group_name"))
}

func (abcmt azurermBotChannelMsTeamsAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](abcmt.ref.Append("timeouts"))
}

type azurermBotChannelMsTeamsState struct {
	BotName               string         `json:"bot_name"`
	CallingWebHook        string         `json:"calling_web_hook"`
	DeploymentEnvironment string         `json:"deployment_environment"`
	EnableCalling         bool           `json:"enable_calling"`
	Id                    string         `json:"id"`
	Location              string         `json:"location"`
	ResourceGroupName     string         `json:"resource_group_name"`
	Timeouts              *TimeoutsState `json:"timeouts"`
}
