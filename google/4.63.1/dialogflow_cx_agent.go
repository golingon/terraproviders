// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dialogflowcxagent "github.com/golingon/terraproviders/google/4.63.1/dialogflowcxagent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowCxAgent creates a new instance of [DialogflowCxAgent].
func NewDialogflowCxAgent(name string, args DialogflowCxAgentArgs) *DialogflowCxAgent {
	return &DialogflowCxAgent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowCxAgent)(nil)

// DialogflowCxAgent represents the Terraform resource google_dialogflow_cx_agent.
type DialogflowCxAgent struct {
	Name      string
	Args      DialogflowCxAgentArgs
	state     *dialogflowCxAgentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowCxAgent].
func (dca *DialogflowCxAgent) Type() string {
	return "google_dialogflow_cx_agent"
}

// LocalName returns the local name for [DialogflowCxAgent].
func (dca *DialogflowCxAgent) LocalName() string {
	return dca.Name
}

// Configuration returns the configuration (args) for [DialogflowCxAgent].
func (dca *DialogflowCxAgent) Configuration() interface{} {
	return dca.Args
}

// DependOn is used for other resources to depend on [DialogflowCxAgent].
func (dca *DialogflowCxAgent) DependOn() terra.Reference {
	return terra.ReferenceResource(dca)
}

// Dependencies returns the list of resources [DialogflowCxAgent] depends_on.
func (dca *DialogflowCxAgent) Dependencies() terra.Dependencies {
	return dca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowCxAgent].
func (dca *DialogflowCxAgent) LifecycleManagement() *terra.Lifecycle {
	return dca.Lifecycle
}

// Attributes returns the attributes for [DialogflowCxAgent].
func (dca *DialogflowCxAgent) Attributes() dialogflowCxAgentAttributes {
	return dialogflowCxAgentAttributes{ref: terra.ReferenceResource(dca)}
}

// ImportState imports the given attribute values into [DialogflowCxAgent]'s state.
func (dca *DialogflowCxAgent) ImportState(av io.Reader) error {
	dca.state = &dialogflowCxAgentState{}
	if err := json.NewDecoder(av).Decode(dca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dca.Type(), dca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowCxAgent] has state.
func (dca *DialogflowCxAgent) State() (*dialogflowCxAgentState, bool) {
	return dca.state, dca.state != nil
}

// StateMust returns the state for [DialogflowCxAgent]. Panics if the state is nil.
func (dca *DialogflowCxAgent) StateMust() *dialogflowCxAgentState {
	if dca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dca.Type(), dca.LocalName()))
	}
	return dca.state
}

// DialogflowCxAgentArgs contains the configurations for google_dialogflow_cx_agent.
type DialogflowCxAgentArgs struct {
	// AvatarUri: string, optional
	AvatarUri terra.StringValue `hcl:"avatar_uri,attr"`
	// DefaultLanguageCode: string, required
	DefaultLanguageCode terra.StringValue `hcl:"default_language_code,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// EnableSpellCorrection: bool, optional
	EnableSpellCorrection terra.BoolValue `hcl:"enable_spell_correction,attr"`
	// EnableStackdriverLogging: bool, optional
	EnableStackdriverLogging terra.BoolValue `hcl:"enable_stackdriver_logging,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SecuritySettings: string, optional
	SecuritySettings terra.StringValue `hcl:"security_settings,attr"`
	// SupportedLanguageCodes: list of string, optional
	SupportedLanguageCodes terra.ListValue[terra.StringValue] `hcl:"supported_language_codes,attr"`
	// TimeZone: string, required
	TimeZone terra.StringValue `hcl:"time_zone,attr" validate:"required"`
	// SpeechToTextSettings: optional
	SpeechToTextSettings *dialogflowcxagent.SpeechToTextSettings `hcl:"speech_to_text_settings,block"`
	// Timeouts: optional
	Timeouts *dialogflowcxagent.Timeouts `hcl:"timeouts,block"`
}
type dialogflowCxAgentAttributes struct {
	ref terra.Reference
}

// AvatarUri returns a reference to field avatar_uri of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) AvatarUri() terra.StringValue {
	return terra.ReferenceAsString(dca.ref.Append("avatar_uri"))
}

// DefaultLanguageCode returns a reference to field default_language_code of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) DefaultLanguageCode() terra.StringValue {
	return terra.ReferenceAsString(dca.ref.Append("default_language_code"))
}

// Description returns a reference to field description of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dca.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dca.ref.Append("display_name"))
}

// EnableSpellCorrection returns a reference to field enable_spell_correction of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) EnableSpellCorrection() terra.BoolValue {
	return terra.ReferenceAsBool(dca.ref.Append("enable_spell_correction"))
}

// EnableStackdriverLogging returns a reference to field enable_stackdriver_logging of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) EnableStackdriverLogging() terra.BoolValue {
	return terra.ReferenceAsBool(dca.ref.Append("enable_stackdriver_logging"))
}

// Id returns a reference to field id of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dca.ref.Append("id"))
}

// Location returns a reference to field location of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dca.ref.Append("location"))
}

// Name returns a reference to field name of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dca.ref.Append("name"))
}

// Project returns a reference to field project of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dca.ref.Append("project"))
}

// SecuritySettings returns a reference to field security_settings of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) SecuritySettings() terra.StringValue {
	return terra.ReferenceAsString(dca.ref.Append("security_settings"))
}

// StartFlow returns a reference to field start_flow of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) StartFlow() terra.StringValue {
	return terra.ReferenceAsString(dca.ref.Append("start_flow"))
}

// SupportedLanguageCodes returns a reference to field supported_language_codes of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) SupportedLanguageCodes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dca.ref.Append("supported_language_codes"))
}

// TimeZone returns a reference to field time_zone of google_dialogflow_cx_agent.
func (dca dialogflowCxAgentAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(dca.ref.Append("time_zone"))
}

func (dca dialogflowCxAgentAttributes) SpeechToTextSettings() terra.ListValue[dialogflowcxagent.SpeechToTextSettingsAttributes] {
	return terra.ReferenceAsList[dialogflowcxagent.SpeechToTextSettingsAttributes](dca.ref.Append("speech_to_text_settings"))
}

func (dca dialogflowCxAgentAttributes) Timeouts() dialogflowcxagent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowcxagent.TimeoutsAttributes](dca.ref.Append("timeouts"))
}

type dialogflowCxAgentState struct {
	AvatarUri                string                                        `json:"avatar_uri"`
	DefaultLanguageCode      string                                        `json:"default_language_code"`
	Description              string                                        `json:"description"`
	DisplayName              string                                        `json:"display_name"`
	EnableSpellCorrection    bool                                          `json:"enable_spell_correction"`
	EnableStackdriverLogging bool                                          `json:"enable_stackdriver_logging"`
	Id                       string                                        `json:"id"`
	Location                 string                                        `json:"location"`
	Name                     string                                        `json:"name"`
	Project                  string                                        `json:"project"`
	SecuritySettings         string                                        `json:"security_settings"`
	StartFlow                string                                        `json:"start_flow"`
	SupportedLanguageCodes   []string                                      `json:"supported_language_codes"`
	TimeZone                 string                                        `json:"time_zone"`
	SpeechToTextSettings     []dialogflowcxagent.SpeechToTextSettingsState `json:"speech_to_text_settings"`
	Timeouts                 *dialogflowcxagent.TimeoutsState              `json:"timeouts"`
}
