// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botchanneldirectline "github.com/golingon/terraproviders/azurerm/3.66.0/botchanneldirectline"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBotChannelDirectline creates a new instance of [BotChannelDirectline].
func NewBotChannelDirectline(name string, args BotChannelDirectlineArgs) *BotChannelDirectline {
	return &BotChannelDirectline{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotChannelDirectline)(nil)

// BotChannelDirectline represents the Terraform resource azurerm_bot_channel_directline.
type BotChannelDirectline struct {
	Name      string
	Args      BotChannelDirectlineArgs
	state     *botChannelDirectlineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotChannelDirectline].
func (bcd *BotChannelDirectline) Type() string {
	return "azurerm_bot_channel_directline"
}

// LocalName returns the local name for [BotChannelDirectline].
func (bcd *BotChannelDirectline) LocalName() string {
	return bcd.Name
}

// Configuration returns the configuration (args) for [BotChannelDirectline].
func (bcd *BotChannelDirectline) Configuration() interface{} {
	return bcd.Args
}

// DependOn is used for other resources to depend on [BotChannelDirectline].
func (bcd *BotChannelDirectline) DependOn() terra.Reference {
	return terra.ReferenceResource(bcd)
}

// Dependencies returns the list of resources [BotChannelDirectline] depends_on.
func (bcd *BotChannelDirectline) Dependencies() terra.Dependencies {
	return bcd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotChannelDirectline].
func (bcd *BotChannelDirectline) LifecycleManagement() *terra.Lifecycle {
	return bcd.Lifecycle
}

// Attributes returns the attributes for [BotChannelDirectline].
func (bcd *BotChannelDirectline) Attributes() botChannelDirectlineAttributes {
	return botChannelDirectlineAttributes{ref: terra.ReferenceResource(bcd)}
}

// ImportState imports the given attribute values into [BotChannelDirectline]'s state.
func (bcd *BotChannelDirectline) ImportState(av io.Reader) error {
	bcd.state = &botChannelDirectlineState{}
	if err := json.NewDecoder(av).Decode(bcd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcd.Type(), bcd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotChannelDirectline] has state.
func (bcd *BotChannelDirectline) State() (*botChannelDirectlineState, bool) {
	return bcd.state, bcd.state != nil
}

// StateMust returns the state for [BotChannelDirectline]. Panics if the state is nil.
func (bcd *BotChannelDirectline) StateMust() *botChannelDirectlineState {
	if bcd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcd.Type(), bcd.LocalName()))
	}
	return bcd.state
}

// BotChannelDirectlineArgs contains the configurations for azurerm_bot_channel_directline.
type BotChannelDirectlineArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Site: min=1
	Site []botchanneldirectline.Site `hcl:"site,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *botchanneldirectline.Timeouts `hcl:"timeouts,block"`
}
type botChannelDirectlineAttributes struct {
	ref terra.Reference
}

// BotName returns a reference to field bot_name of azurerm_bot_channel_directline.
func (bcd botChannelDirectlineAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(bcd.ref.Append("bot_name"))
}

// Id returns a reference to field id of azurerm_bot_channel_directline.
func (bcd botChannelDirectlineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_bot_channel_directline.
func (bcd botChannelDirectlineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcd.ref.Append("location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_channel_directline.
func (bcd botChannelDirectlineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bcd.ref.Append("resource_group_name"))
}

func (bcd botChannelDirectlineAttributes) Site() terra.SetValue[botchanneldirectline.SiteAttributes] {
	return terra.ReferenceAsSet[botchanneldirectline.SiteAttributes](bcd.ref.Append("site"))
}

func (bcd botChannelDirectlineAttributes) Timeouts() botchanneldirectline.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botchanneldirectline.TimeoutsAttributes](bcd.ref.Append("timeouts"))
}

type botChannelDirectlineState struct {
	BotName           string                              `json:"bot_name"`
	Id                string                              `json:"id"`
	Location          string                              `json:"location"`
	ResourceGroupName string                              `json:"resource_group_name"`
	Site              []botchanneldirectline.SiteState    `json:"site"`
	Timeouts          *botchanneldirectline.TimeoutsState `json:"timeouts"`
}
