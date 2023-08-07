// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botchannelfacebook "github.com/golingon/terraproviders/azurerm/3.68.0/botchannelfacebook"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBotChannelFacebook creates a new instance of [BotChannelFacebook].
func NewBotChannelFacebook(name string, args BotChannelFacebookArgs) *BotChannelFacebook {
	return &BotChannelFacebook{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotChannelFacebook)(nil)

// BotChannelFacebook represents the Terraform resource azurerm_bot_channel_facebook.
type BotChannelFacebook struct {
	Name      string
	Args      BotChannelFacebookArgs
	state     *botChannelFacebookState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotChannelFacebook].
func (bcf *BotChannelFacebook) Type() string {
	return "azurerm_bot_channel_facebook"
}

// LocalName returns the local name for [BotChannelFacebook].
func (bcf *BotChannelFacebook) LocalName() string {
	return bcf.Name
}

// Configuration returns the configuration (args) for [BotChannelFacebook].
func (bcf *BotChannelFacebook) Configuration() interface{} {
	return bcf.Args
}

// DependOn is used for other resources to depend on [BotChannelFacebook].
func (bcf *BotChannelFacebook) DependOn() terra.Reference {
	return terra.ReferenceResource(bcf)
}

// Dependencies returns the list of resources [BotChannelFacebook] depends_on.
func (bcf *BotChannelFacebook) Dependencies() terra.Dependencies {
	return bcf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotChannelFacebook].
func (bcf *BotChannelFacebook) LifecycleManagement() *terra.Lifecycle {
	return bcf.Lifecycle
}

// Attributes returns the attributes for [BotChannelFacebook].
func (bcf *BotChannelFacebook) Attributes() botChannelFacebookAttributes {
	return botChannelFacebookAttributes{ref: terra.ReferenceResource(bcf)}
}

// ImportState imports the given attribute values into [BotChannelFacebook]'s state.
func (bcf *BotChannelFacebook) ImportState(av io.Reader) error {
	bcf.state = &botChannelFacebookState{}
	if err := json.NewDecoder(av).Decode(bcf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcf.Type(), bcf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotChannelFacebook] has state.
func (bcf *BotChannelFacebook) State() (*botChannelFacebookState, bool) {
	return bcf.state, bcf.state != nil
}

// StateMust returns the state for [BotChannelFacebook]. Panics if the state is nil.
func (bcf *BotChannelFacebook) StateMust() *botChannelFacebookState {
	if bcf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcf.Type(), bcf.LocalName()))
	}
	return bcf.state
}

// BotChannelFacebookArgs contains the configurations for azurerm_bot_channel_facebook.
type BotChannelFacebookArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// FacebookApplicationId: string, required
	FacebookApplicationId terra.StringValue `hcl:"facebook_application_id,attr" validate:"required"`
	// FacebookApplicationSecret: string, required
	FacebookApplicationSecret terra.StringValue `hcl:"facebook_application_secret,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Page: min=1
	Page []botchannelfacebook.Page `hcl:"page,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *botchannelfacebook.Timeouts `hcl:"timeouts,block"`
}
type botChannelFacebookAttributes struct {
	ref terra.Reference
}

// BotName returns a reference to field bot_name of azurerm_bot_channel_facebook.
func (bcf botChannelFacebookAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(bcf.ref.Append("bot_name"))
}

// FacebookApplicationId returns a reference to field facebook_application_id of azurerm_bot_channel_facebook.
func (bcf botChannelFacebookAttributes) FacebookApplicationId() terra.StringValue {
	return terra.ReferenceAsString(bcf.ref.Append("facebook_application_id"))
}

// FacebookApplicationSecret returns a reference to field facebook_application_secret of azurerm_bot_channel_facebook.
func (bcf botChannelFacebookAttributes) FacebookApplicationSecret() terra.StringValue {
	return terra.ReferenceAsString(bcf.ref.Append("facebook_application_secret"))
}

// Id returns a reference to field id of azurerm_bot_channel_facebook.
func (bcf botChannelFacebookAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcf.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_bot_channel_facebook.
func (bcf botChannelFacebookAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcf.ref.Append("location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_channel_facebook.
func (bcf botChannelFacebookAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bcf.ref.Append("resource_group_name"))
}

func (bcf botChannelFacebookAttributes) Page() terra.SetValue[botchannelfacebook.PageAttributes] {
	return terra.ReferenceAsSet[botchannelfacebook.PageAttributes](bcf.ref.Append("page"))
}

func (bcf botChannelFacebookAttributes) Timeouts() botchannelfacebook.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botchannelfacebook.TimeoutsAttributes](bcf.ref.Append("timeouts"))
}

type botChannelFacebookState struct {
	BotName                   string                            `json:"bot_name"`
	FacebookApplicationId     string                            `json:"facebook_application_id"`
	FacebookApplicationSecret string                            `json:"facebook_application_secret"`
	Id                        string                            `json:"id"`
	Location                  string                            `json:"location"`
	ResourceGroupName         string                            `json:"resource_group_name"`
	Page                      []botchannelfacebook.PageState    `json:"page"`
	Timeouts                  *botchannelfacebook.TimeoutsState `json:"timeouts"`
}
