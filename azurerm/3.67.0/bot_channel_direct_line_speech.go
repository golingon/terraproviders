// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botchanneldirectlinespeech "github.com/golingon/terraproviders/azurerm/3.67.0/botchanneldirectlinespeech"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBotChannelDirectLineSpeech creates a new instance of [BotChannelDirectLineSpeech].
func NewBotChannelDirectLineSpeech(name string, args BotChannelDirectLineSpeechArgs) *BotChannelDirectLineSpeech {
	return &BotChannelDirectLineSpeech{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotChannelDirectLineSpeech)(nil)

// BotChannelDirectLineSpeech represents the Terraform resource azurerm_bot_channel_direct_line_speech.
type BotChannelDirectLineSpeech struct {
	Name      string
	Args      BotChannelDirectLineSpeechArgs
	state     *botChannelDirectLineSpeechState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotChannelDirectLineSpeech].
func (bcdls *BotChannelDirectLineSpeech) Type() string {
	return "azurerm_bot_channel_direct_line_speech"
}

// LocalName returns the local name for [BotChannelDirectLineSpeech].
func (bcdls *BotChannelDirectLineSpeech) LocalName() string {
	return bcdls.Name
}

// Configuration returns the configuration (args) for [BotChannelDirectLineSpeech].
func (bcdls *BotChannelDirectLineSpeech) Configuration() interface{} {
	return bcdls.Args
}

// DependOn is used for other resources to depend on [BotChannelDirectLineSpeech].
func (bcdls *BotChannelDirectLineSpeech) DependOn() terra.Reference {
	return terra.ReferenceResource(bcdls)
}

// Dependencies returns the list of resources [BotChannelDirectLineSpeech] depends_on.
func (bcdls *BotChannelDirectLineSpeech) Dependencies() terra.Dependencies {
	return bcdls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotChannelDirectLineSpeech].
func (bcdls *BotChannelDirectLineSpeech) LifecycleManagement() *terra.Lifecycle {
	return bcdls.Lifecycle
}

// Attributes returns the attributes for [BotChannelDirectLineSpeech].
func (bcdls *BotChannelDirectLineSpeech) Attributes() botChannelDirectLineSpeechAttributes {
	return botChannelDirectLineSpeechAttributes{ref: terra.ReferenceResource(bcdls)}
}

// ImportState imports the given attribute values into [BotChannelDirectLineSpeech]'s state.
func (bcdls *BotChannelDirectLineSpeech) ImportState(av io.Reader) error {
	bcdls.state = &botChannelDirectLineSpeechState{}
	if err := json.NewDecoder(av).Decode(bcdls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcdls.Type(), bcdls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotChannelDirectLineSpeech] has state.
func (bcdls *BotChannelDirectLineSpeech) State() (*botChannelDirectLineSpeechState, bool) {
	return bcdls.state, bcdls.state != nil
}

// StateMust returns the state for [BotChannelDirectLineSpeech]. Panics if the state is nil.
func (bcdls *BotChannelDirectLineSpeech) StateMust() *botChannelDirectLineSpeechState {
	if bcdls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcdls.Type(), bcdls.LocalName()))
	}
	return bcdls.state
}

// BotChannelDirectLineSpeechArgs contains the configurations for azurerm_bot_channel_direct_line_speech.
type BotChannelDirectLineSpeechArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// CognitiveServiceAccessKey: string, required
	CognitiveServiceAccessKey terra.StringValue `hcl:"cognitive_service_access_key,attr" validate:"required"`
	// CognitiveServiceLocation: string, required
	CognitiveServiceLocation terra.StringValue `hcl:"cognitive_service_location,attr" validate:"required"`
	// CustomSpeechModelId: string, optional
	CustomSpeechModelId terra.StringValue `hcl:"custom_speech_model_id,attr"`
	// CustomVoiceDeploymentId: string, optional
	CustomVoiceDeploymentId terra.StringValue `hcl:"custom_voice_deployment_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *botchanneldirectlinespeech.Timeouts `hcl:"timeouts,block"`
}
type botChannelDirectLineSpeechAttributes struct {
	ref terra.Reference
}

// BotName returns a reference to field bot_name of azurerm_bot_channel_direct_line_speech.
func (bcdls botChannelDirectLineSpeechAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(bcdls.ref.Append("bot_name"))
}

// CognitiveServiceAccessKey returns a reference to field cognitive_service_access_key of azurerm_bot_channel_direct_line_speech.
func (bcdls botChannelDirectLineSpeechAttributes) CognitiveServiceAccessKey() terra.StringValue {
	return terra.ReferenceAsString(bcdls.ref.Append("cognitive_service_access_key"))
}

// CognitiveServiceLocation returns a reference to field cognitive_service_location of azurerm_bot_channel_direct_line_speech.
func (bcdls botChannelDirectLineSpeechAttributes) CognitiveServiceLocation() terra.StringValue {
	return terra.ReferenceAsString(bcdls.ref.Append("cognitive_service_location"))
}

// CustomSpeechModelId returns a reference to field custom_speech_model_id of azurerm_bot_channel_direct_line_speech.
func (bcdls botChannelDirectLineSpeechAttributes) CustomSpeechModelId() terra.StringValue {
	return terra.ReferenceAsString(bcdls.ref.Append("custom_speech_model_id"))
}

// CustomVoiceDeploymentId returns a reference to field custom_voice_deployment_id of azurerm_bot_channel_direct_line_speech.
func (bcdls botChannelDirectLineSpeechAttributes) CustomVoiceDeploymentId() terra.StringValue {
	return terra.ReferenceAsString(bcdls.ref.Append("custom_voice_deployment_id"))
}

// Id returns a reference to field id of azurerm_bot_channel_direct_line_speech.
func (bcdls botChannelDirectLineSpeechAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcdls.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_bot_channel_direct_line_speech.
func (bcdls botChannelDirectLineSpeechAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcdls.ref.Append("location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_channel_direct_line_speech.
func (bcdls botChannelDirectLineSpeechAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bcdls.ref.Append("resource_group_name"))
}

func (bcdls botChannelDirectLineSpeechAttributes) Timeouts() botchanneldirectlinespeech.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botchanneldirectlinespeech.TimeoutsAttributes](bcdls.ref.Append("timeouts"))
}

type botChannelDirectLineSpeechState struct {
	BotName                   string                                    `json:"bot_name"`
	CognitiveServiceAccessKey string                                    `json:"cognitive_service_access_key"`
	CognitiveServiceLocation  string                                    `json:"cognitive_service_location"`
	CustomSpeechModelId       string                                    `json:"custom_speech_model_id"`
	CustomVoiceDeploymentId   string                                    `json:"custom_voice_deployment_id"`
	Id                        string                                    `json:"id"`
	Location                  string                                    `json:"location"`
	ResourceGroupName         string                                    `json:"resource_group_name"`
	Timeouts                  *botchanneldirectlinespeech.TimeoutsState `json:"timeouts"`
}
