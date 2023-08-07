// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botchannelline "github.com/golingon/terraproviders/azurerm/3.68.0/botchannelline"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBotChannelLine creates a new instance of [BotChannelLine].
func NewBotChannelLine(name string, args BotChannelLineArgs) *BotChannelLine {
	return &BotChannelLine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotChannelLine)(nil)

// BotChannelLine represents the Terraform resource azurerm_bot_channel_line.
type BotChannelLine struct {
	Name      string
	Args      BotChannelLineArgs
	state     *botChannelLineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotChannelLine].
func (bcl *BotChannelLine) Type() string {
	return "azurerm_bot_channel_line"
}

// LocalName returns the local name for [BotChannelLine].
func (bcl *BotChannelLine) LocalName() string {
	return bcl.Name
}

// Configuration returns the configuration (args) for [BotChannelLine].
func (bcl *BotChannelLine) Configuration() interface{} {
	return bcl.Args
}

// DependOn is used for other resources to depend on [BotChannelLine].
func (bcl *BotChannelLine) DependOn() terra.Reference {
	return terra.ReferenceResource(bcl)
}

// Dependencies returns the list of resources [BotChannelLine] depends_on.
func (bcl *BotChannelLine) Dependencies() terra.Dependencies {
	return bcl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotChannelLine].
func (bcl *BotChannelLine) LifecycleManagement() *terra.Lifecycle {
	return bcl.Lifecycle
}

// Attributes returns the attributes for [BotChannelLine].
func (bcl *BotChannelLine) Attributes() botChannelLineAttributes {
	return botChannelLineAttributes{ref: terra.ReferenceResource(bcl)}
}

// ImportState imports the given attribute values into [BotChannelLine]'s state.
func (bcl *BotChannelLine) ImportState(av io.Reader) error {
	bcl.state = &botChannelLineState{}
	if err := json.NewDecoder(av).Decode(bcl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcl.Type(), bcl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotChannelLine] has state.
func (bcl *BotChannelLine) State() (*botChannelLineState, bool) {
	return bcl.state, bcl.state != nil
}

// StateMust returns the state for [BotChannelLine]. Panics if the state is nil.
func (bcl *BotChannelLine) StateMust() *botChannelLineState {
	if bcl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcl.Type(), bcl.LocalName()))
	}
	return bcl.state
}

// BotChannelLineArgs contains the configurations for azurerm_bot_channel_line.
type BotChannelLineArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// LineChannel: min=1
	LineChannel []botchannelline.LineChannel `hcl:"line_channel,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *botchannelline.Timeouts `hcl:"timeouts,block"`
}
type botChannelLineAttributes struct {
	ref terra.Reference
}

// BotName returns a reference to field bot_name of azurerm_bot_channel_line.
func (bcl botChannelLineAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(bcl.ref.Append("bot_name"))
}

// Id returns a reference to field id of azurerm_bot_channel_line.
func (bcl botChannelLineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcl.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_bot_channel_line.
func (bcl botChannelLineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcl.ref.Append("location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_channel_line.
func (bcl botChannelLineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bcl.ref.Append("resource_group_name"))
}

func (bcl botChannelLineAttributes) LineChannel() terra.SetValue[botchannelline.LineChannelAttributes] {
	return terra.ReferenceAsSet[botchannelline.LineChannelAttributes](bcl.ref.Append("line_channel"))
}

func (bcl botChannelLineAttributes) Timeouts() botchannelline.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botchannelline.TimeoutsAttributes](bcl.ref.Append("timeouts"))
}

type botChannelLineState struct {
	BotName           string                            `json:"bot_name"`
	Id                string                            `json:"id"`
	Location          string                            `json:"location"`
	ResourceGroupName string                            `json:"resource_group_name"`
	LineChannel       []botchannelline.LineChannelState `json:"line_channel"`
	Timeouts          *botchannelline.TimeoutsState     `json:"timeouts"`
}
