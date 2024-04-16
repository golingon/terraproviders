// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_lexv2models_bot_locale

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

// Resource represents the Terraform resource aws_lexv2models_bot_locale.
type Resource struct {
	Name      string
	Args      Args
	state     *awsLexv2ModelsBotLocaleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (albl *Resource) Type() string {
	return "aws_lexv2models_bot_locale"
}

// LocalName returns the local name for [Resource].
func (albl *Resource) LocalName() string {
	return albl.Name
}

// Configuration returns the configuration (args) for [Resource].
func (albl *Resource) Configuration() interface{} {
	return albl.Args
}

// DependOn is used for other resources to depend on [Resource].
func (albl *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(albl)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (albl *Resource) Dependencies() terra.Dependencies {
	return albl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (albl *Resource) LifecycleManagement() *terra.Lifecycle {
	return albl.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (albl *Resource) Attributes() awsLexv2ModelsBotLocaleAttributes {
	return awsLexv2ModelsBotLocaleAttributes{ref: terra.ReferenceResource(albl)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (albl *Resource) ImportState(state io.Reader) error {
	albl.state = &awsLexv2ModelsBotLocaleState{}
	if err := json.NewDecoder(state).Decode(albl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", albl.Type(), albl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (albl *Resource) State() (*awsLexv2ModelsBotLocaleState, bool) {
	return albl.state, albl.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (albl *Resource) StateMust() *awsLexv2ModelsBotLocaleState {
	if albl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", albl.Type(), albl.LocalName()))
	}
	return albl.state
}

// Args contains the configurations for aws_lexv2models_bot_locale.
type Args struct {
	// BotId: string, required
	BotId terra.StringValue `hcl:"bot_id,attr" validate:"required"`
	// BotVersion: string, required
	BotVersion terra.StringValue `hcl:"bot_version,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// LocaleId: string, required
	LocaleId terra.StringValue `hcl:"locale_id,attr" validate:"required"`
	// NLuIntentConfidenceThreshold: number, required
	NLuIntentConfidenceThreshold terra.NumberValue `hcl:"n_lu_intent_confidence_threshold,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// VoiceSettings: min=0
	VoiceSettings []VoiceSettings `hcl:"voice_settings,block" validate:"min=0"`
}

type awsLexv2ModelsBotLocaleAttributes struct {
	ref terra.Reference
}

// BotId returns a reference to field bot_id of aws_lexv2models_bot_locale.
func (albl awsLexv2ModelsBotLocaleAttributes) BotId() terra.StringValue {
	return terra.ReferenceAsString(albl.ref.Append("bot_id"))
}

// BotVersion returns a reference to field bot_version of aws_lexv2models_bot_locale.
func (albl awsLexv2ModelsBotLocaleAttributes) BotVersion() terra.StringValue {
	return terra.ReferenceAsString(albl.ref.Append("bot_version"))
}

// Description returns a reference to field description of aws_lexv2models_bot_locale.
func (albl awsLexv2ModelsBotLocaleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(albl.ref.Append("description"))
}

// Id returns a reference to field id of aws_lexv2models_bot_locale.
func (albl awsLexv2ModelsBotLocaleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(albl.ref.Append("id"))
}

// LocaleId returns a reference to field locale_id of aws_lexv2models_bot_locale.
func (albl awsLexv2ModelsBotLocaleAttributes) LocaleId() terra.StringValue {
	return terra.ReferenceAsString(albl.ref.Append("locale_id"))
}

// NLuIntentConfidenceThreshold returns a reference to field n_lu_intent_confidence_threshold of aws_lexv2models_bot_locale.
func (albl awsLexv2ModelsBotLocaleAttributes) NLuIntentConfidenceThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(albl.ref.Append("n_lu_intent_confidence_threshold"))
}

// Name returns a reference to field name of aws_lexv2models_bot_locale.
func (albl awsLexv2ModelsBotLocaleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(albl.ref.Append("name"))
}

func (albl awsLexv2ModelsBotLocaleAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](albl.ref.Append("timeouts"))
}

func (albl awsLexv2ModelsBotLocaleAttributes) VoiceSettings() terra.ListValue[VoiceSettingsAttributes] {
	return terra.ReferenceAsList[VoiceSettingsAttributes](albl.ref.Append("voice_settings"))
}

type awsLexv2ModelsBotLocaleState struct {
	BotId                        string               `json:"bot_id"`
	BotVersion                   string               `json:"bot_version"`
	Description                  string               `json:"description"`
	Id                           string               `json:"id"`
	LocaleId                     string               `json:"locale_id"`
	NLuIntentConfidenceThreshold float64              `json:"n_lu_intent_confidence_threshold"`
	Name                         string               `json:"name"`
	Timeouts                     *TimeoutsState       `json:"timeouts"`
	VoiceSettings                []VoiceSettingsState `json:"voice_settings"`
}
