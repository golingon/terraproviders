// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dialogflowagent "github.com/golingon/terraproviders/googlebeta/4.62.0/dialogflowagent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowAgent creates a new instance of [DialogflowAgent].
func NewDialogflowAgent(name string, args DialogflowAgentArgs) *DialogflowAgent {
	return &DialogflowAgent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowAgent)(nil)

// DialogflowAgent represents the Terraform resource google_dialogflow_agent.
type DialogflowAgent struct {
	Name      string
	Args      DialogflowAgentArgs
	state     *dialogflowAgentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowAgent].
func (da *DialogflowAgent) Type() string {
	return "google_dialogflow_agent"
}

// LocalName returns the local name for [DialogflowAgent].
func (da *DialogflowAgent) LocalName() string {
	return da.Name
}

// Configuration returns the configuration (args) for [DialogflowAgent].
func (da *DialogflowAgent) Configuration() interface{} {
	return da.Args
}

// DependOn is used for other resources to depend on [DialogflowAgent].
func (da *DialogflowAgent) DependOn() terra.Reference {
	return terra.ReferenceResource(da)
}

// Dependencies returns the list of resources [DialogflowAgent] depends_on.
func (da *DialogflowAgent) Dependencies() terra.Dependencies {
	return da.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowAgent].
func (da *DialogflowAgent) LifecycleManagement() *terra.Lifecycle {
	return da.Lifecycle
}

// Attributes returns the attributes for [DialogflowAgent].
func (da *DialogflowAgent) Attributes() dialogflowAgentAttributes {
	return dialogflowAgentAttributes{ref: terra.ReferenceResource(da)}
}

// ImportState imports the given attribute values into [DialogflowAgent]'s state.
func (da *DialogflowAgent) ImportState(av io.Reader) error {
	da.state = &dialogflowAgentState{}
	if err := json.NewDecoder(av).Decode(da.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", da.Type(), da.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowAgent] has state.
func (da *DialogflowAgent) State() (*dialogflowAgentState, bool) {
	return da.state, da.state != nil
}

// StateMust returns the state for [DialogflowAgent]. Panics if the state is nil.
func (da *DialogflowAgent) StateMust() *dialogflowAgentState {
	if da.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", da.Type(), da.LocalName()))
	}
	return da.state
}

// DialogflowAgentArgs contains the configurations for google_dialogflow_agent.
type DialogflowAgentArgs struct {
	// ApiVersion: string, optional
	ApiVersion terra.StringValue `hcl:"api_version,attr"`
	// AvatarUri: string, optional
	AvatarUri terra.StringValue `hcl:"avatar_uri,attr"`
	// ClassificationThreshold: number, optional
	ClassificationThreshold terra.NumberValue `hcl:"classification_threshold,attr"`
	// DefaultLanguageCode: string, required
	DefaultLanguageCode terra.StringValue `hcl:"default_language_code,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// EnableLogging: bool, optional
	EnableLogging terra.BoolValue `hcl:"enable_logging,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MatchMode: string, optional
	MatchMode terra.StringValue `hcl:"match_mode,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SupportedLanguageCodes: list of string, optional
	SupportedLanguageCodes terra.ListValue[terra.StringValue] `hcl:"supported_language_codes,attr"`
	// Tier: string, optional
	Tier terra.StringValue `hcl:"tier,attr"`
	// TimeZone: string, required
	TimeZone terra.StringValue `hcl:"time_zone,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dialogflowagent.Timeouts `hcl:"timeouts,block"`
}
type dialogflowAgentAttributes struct {
	ref terra.Reference
}

// ApiVersion returns a reference to field api_version of google_dialogflow_agent.
func (da dialogflowAgentAttributes) ApiVersion() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("api_version"))
}

// AvatarUri returns a reference to field avatar_uri of google_dialogflow_agent.
func (da dialogflowAgentAttributes) AvatarUri() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("avatar_uri"))
}

// AvatarUriBackend returns a reference to field avatar_uri_backend of google_dialogflow_agent.
func (da dialogflowAgentAttributes) AvatarUriBackend() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("avatar_uri_backend"))
}

// ClassificationThreshold returns a reference to field classification_threshold of google_dialogflow_agent.
func (da dialogflowAgentAttributes) ClassificationThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(da.ref.Append("classification_threshold"))
}

// DefaultLanguageCode returns a reference to field default_language_code of google_dialogflow_agent.
func (da dialogflowAgentAttributes) DefaultLanguageCode() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("default_language_code"))
}

// Description returns a reference to field description of google_dialogflow_agent.
func (da dialogflowAgentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dialogflow_agent.
func (da dialogflowAgentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("display_name"))
}

// EnableLogging returns a reference to field enable_logging of google_dialogflow_agent.
func (da dialogflowAgentAttributes) EnableLogging() terra.BoolValue {
	return terra.ReferenceAsBool(da.ref.Append("enable_logging"))
}

// Id returns a reference to field id of google_dialogflow_agent.
func (da dialogflowAgentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("id"))
}

// MatchMode returns a reference to field match_mode of google_dialogflow_agent.
func (da dialogflowAgentAttributes) MatchMode() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("match_mode"))
}

// Project returns a reference to field project of google_dialogflow_agent.
func (da dialogflowAgentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("project"))
}

// SupportedLanguageCodes returns a reference to field supported_language_codes of google_dialogflow_agent.
func (da dialogflowAgentAttributes) SupportedLanguageCodes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](da.ref.Append("supported_language_codes"))
}

// Tier returns a reference to field tier of google_dialogflow_agent.
func (da dialogflowAgentAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("tier"))
}

// TimeZone returns a reference to field time_zone of google_dialogflow_agent.
func (da dialogflowAgentAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("time_zone"))
}

func (da dialogflowAgentAttributes) Timeouts() dialogflowagent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowagent.TimeoutsAttributes](da.ref.Append("timeouts"))
}

type dialogflowAgentState struct {
	ApiVersion              string                         `json:"api_version"`
	AvatarUri               string                         `json:"avatar_uri"`
	AvatarUriBackend        string                         `json:"avatar_uri_backend"`
	ClassificationThreshold float64                        `json:"classification_threshold"`
	DefaultLanguageCode     string                         `json:"default_language_code"`
	Description             string                         `json:"description"`
	DisplayName             string                         `json:"display_name"`
	EnableLogging           bool                           `json:"enable_logging"`
	Id                      string                         `json:"id"`
	MatchMode               string                         `json:"match_mode"`
	Project                 string                         `json:"project"`
	SupportedLanguageCodes  []string                       `json:"supported_language_codes"`
	Tier                    string                         `json:"tier"`
	TimeZone                string                         `json:"time_zone"`
	Timeouts                *dialogflowagent.TimeoutsState `json:"timeouts"`
}
