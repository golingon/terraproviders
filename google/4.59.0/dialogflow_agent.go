// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dialogflowagent "github.com/golingon/terraproviders/google/4.59.0/dialogflowagent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDialogflowAgent(name string, args DialogflowAgentArgs) *DialogflowAgent {
	return &DialogflowAgent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowAgent)(nil)

type DialogflowAgent struct {
	Name  string
	Args  DialogflowAgentArgs
	state *dialogflowAgentState
}

func (da *DialogflowAgent) Type() string {
	return "google_dialogflow_agent"
}

func (da *DialogflowAgent) LocalName() string {
	return da.Name
}

func (da *DialogflowAgent) Configuration() interface{} {
	return da.Args
}

func (da *DialogflowAgent) Attributes() dialogflowAgentAttributes {
	return dialogflowAgentAttributes{ref: terra.ReferenceResource(da)}
}

func (da *DialogflowAgent) ImportState(av io.Reader) error {
	da.state = &dialogflowAgentState{}
	if err := json.NewDecoder(av).Decode(da.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", da.Type(), da.LocalName(), err)
	}
	return nil
}

func (da *DialogflowAgent) State() (*dialogflowAgentState, bool) {
	return da.state, da.state != nil
}

func (da *DialogflowAgent) StateMust() *dialogflowAgentState {
	if da.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", da.Type(), da.LocalName()))
	}
	return da.state
}

func (da *DialogflowAgent) DependOn() terra.Reference {
	return terra.ReferenceResource(da)
}

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
	// DependsOn contains resources that DialogflowAgent depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dialogflowAgentAttributes struct {
	ref terra.Reference
}

func (da dialogflowAgentAttributes) ApiVersion() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("api_version"))
}

func (da dialogflowAgentAttributes) AvatarUri() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("avatar_uri"))
}

func (da dialogflowAgentAttributes) AvatarUriBackend() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("avatar_uri_backend"))
}

func (da dialogflowAgentAttributes) ClassificationThreshold() terra.NumberValue {
	return terra.ReferenceNumber(da.ref.Append("classification_threshold"))
}

func (da dialogflowAgentAttributes) DefaultLanguageCode() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("default_language_code"))
}

func (da dialogflowAgentAttributes) Description() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("description"))
}

func (da dialogflowAgentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("display_name"))
}

func (da dialogflowAgentAttributes) EnableLogging() terra.BoolValue {
	return terra.ReferenceBool(da.ref.Append("enable_logging"))
}

func (da dialogflowAgentAttributes) Id() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("id"))
}

func (da dialogflowAgentAttributes) MatchMode() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("match_mode"))
}

func (da dialogflowAgentAttributes) Project() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("project"))
}

func (da dialogflowAgentAttributes) SupportedLanguageCodes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](da.ref.Append("supported_language_codes"))
}

func (da dialogflowAgentAttributes) Tier() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("tier"))
}

func (da dialogflowAgentAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("time_zone"))
}

func (da dialogflowAgentAttributes) Timeouts() dialogflowagent.TimeoutsAttributes {
	return terra.ReferenceSingle[dialogflowagent.TimeoutsAttributes](da.ref.Append("timeouts"))
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
