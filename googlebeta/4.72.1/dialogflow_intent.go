// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dialogflowintent "github.com/golingon/terraproviders/googlebeta/4.72.1/dialogflowintent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowIntent creates a new instance of [DialogflowIntent].
func NewDialogflowIntent(name string, args DialogflowIntentArgs) *DialogflowIntent {
	return &DialogflowIntent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowIntent)(nil)

// DialogflowIntent represents the Terraform resource google_dialogflow_intent.
type DialogflowIntent struct {
	Name      string
	Args      DialogflowIntentArgs
	state     *dialogflowIntentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowIntent].
func (di *DialogflowIntent) Type() string {
	return "google_dialogflow_intent"
}

// LocalName returns the local name for [DialogflowIntent].
func (di *DialogflowIntent) LocalName() string {
	return di.Name
}

// Configuration returns the configuration (args) for [DialogflowIntent].
func (di *DialogflowIntent) Configuration() interface{} {
	return di.Args
}

// DependOn is used for other resources to depend on [DialogflowIntent].
func (di *DialogflowIntent) DependOn() terra.Reference {
	return terra.ReferenceResource(di)
}

// Dependencies returns the list of resources [DialogflowIntent] depends_on.
func (di *DialogflowIntent) Dependencies() terra.Dependencies {
	return di.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowIntent].
func (di *DialogflowIntent) LifecycleManagement() *terra.Lifecycle {
	return di.Lifecycle
}

// Attributes returns the attributes for [DialogflowIntent].
func (di *DialogflowIntent) Attributes() dialogflowIntentAttributes {
	return dialogflowIntentAttributes{ref: terra.ReferenceResource(di)}
}

// ImportState imports the given attribute values into [DialogflowIntent]'s state.
func (di *DialogflowIntent) ImportState(av io.Reader) error {
	di.state = &dialogflowIntentState{}
	if err := json.NewDecoder(av).Decode(di.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", di.Type(), di.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowIntent] has state.
func (di *DialogflowIntent) State() (*dialogflowIntentState, bool) {
	return di.state, di.state != nil
}

// StateMust returns the state for [DialogflowIntent]. Panics if the state is nil.
func (di *DialogflowIntent) StateMust() *dialogflowIntentState {
	if di.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", di.Type(), di.LocalName()))
	}
	return di.state
}

// DialogflowIntentArgs contains the configurations for google_dialogflow_intent.
type DialogflowIntentArgs struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// DefaultResponsePlatforms: list of string, optional
	DefaultResponsePlatforms terra.ListValue[terra.StringValue] `hcl:"default_response_platforms,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Events: list of string, optional
	Events terra.ListValue[terra.StringValue] `hcl:"events,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InputContextNames: list of string, optional
	InputContextNames terra.ListValue[terra.StringValue] `hcl:"input_context_names,attr"`
	// IsFallback: bool, optional
	IsFallback terra.BoolValue `hcl:"is_fallback,attr"`
	// MlDisabled: bool, optional
	MlDisabled terra.BoolValue `hcl:"ml_disabled,attr"`
	// ParentFollowupIntentName: string, optional
	ParentFollowupIntentName terra.StringValue `hcl:"parent_followup_intent_name,attr"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ResetContexts: bool, optional
	ResetContexts terra.BoolValue `hcl:"reset_contexts,attr"`
	// WebhookState: string, optional
	WebhookState terra.StringValue `hcl:"webhook_state,attr"`
	// FollowupIntentInfo: min=0
	FollowupIntentInfo []dialogflowintent.FollowupIntentInfo `hcl:"followup_intent_info,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dialogflowintent.Timeouts `hcl:"timeouts,block"`
}
type dialogflowIntentAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of google_dialogflow_intent.
func (di dialogflowIntentAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("action"))
}

// DefaultResponsePlatforms returns a reference to field default_response_platforms of google_dialogflow_intent.
func (di dialogflowIntentAttributes) DefaultResponsePlatforms() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](di.ref.Append("default_response_platforms"))
}

// DisplayName returns a reference to field display_name of google_dialogflow_intent.
func (di dialogflowIntentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("display_name"))
}

// Events returns a reference to field events of google_dialogflow_intent.
func (di dialogflowIntentAttributes) Events() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](di.ref.Append("events"))
}

// Id returns a reference to field id of google_dialogflow_intent.
func (di dialogflowIntentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("id"))
}

// InputContextNames returns a reference to field input_context_names of google_dialogflow_intent.
func (di dialogflowIntentAttributes) InputContextNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](di.ref.Append("input_context_names"))
}

// IsFallback returns a reference to field is_fallback of google_dialogflow_intent.
func (di dialogflowIntentAttributes) IsFallback() terra.BoolValue {
	return terra.ReferenceAsBool(di.ref.Append("is_fallback"))
}

// MlDisabled returns a reference to field ml_disabled of google_dialogflow_intent.
func (di dialogflowIntentAttributes) MlDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(di.ref.Append("ml_disabled"))
}

// Name returns a reference to field name of google_dialogflow_intent.
func (di dialogflowIntentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("name"))
}

// ParentFollowupIntentName returns a reference to field parent_followup_intent_name of google_dialogflow_intent.
func (di dialogflowIntentAttributes) ParentFollowupIntentName() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("parent_followup_intent_name"))
}

// Priority returns a reference to field priority of google_dialogflow_intent.
func (di dialogflowIntentAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(di.ref.Append("priority"))
}

// Project returns a reference to field project of google_dialogflow_intent.
func (di dialogflowIntentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("project"))
}

// ResetContexts returns a reference to field reset_contexts of google_dialogflow_intent.
func (di dialogflowIntentAttributes) ResetContexts() terra.BoolValue {
	return terra.ReferenceAsBool(di.ref.Append("reset_contexts"))
}

// RootFollowupIntentName returns a reference to field root_followup_intent_name of google_dialogflow_intent.
func (di dialogflowIntentAttributes) RootFollowupIntentName() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("root_followup_intent_name"))
}

// WebhookState returns a reference to field webhook_state of google_dialogflow_intent.
func (di dialogflowIntentAttributes) WebhookState() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("webhook_state"))
}

func (di dialogflowIntentAttributes) FollowupIntentInfo() terra.ListValue[dialogflowintent.FollowupIntentInfoAttributes] {
	return terra.ReferenceAsList[dialogflowintent.FollowupIntentInfoAttributes](di.ref.Append("followup_intent_info"))
}

func (di dialogflowIntentAttributes) Timeouts() dialogflowintent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowintent.TimeoutsAttributes](di.ref.Append("timeouts"))
}

type dialogflowIntentState struct {
	Action                   string                                     `json:"action"`
	DefaultResponsePlatforms []string                                   `json:"default_response_platforms"`
	DisplayName              string                                     `json:"display_name"`
	Events                   []string                                   `json:"events"`
	Id                       string                                     `json:"id"`
	InputContextNames        []string                                   `json:"input_context_names"`
	IsFallback               bool                                       `json:"is_fallback"`
	MlDisabled               bool                                       `json:"ml_disabled"`
	Name                     string                                     `json:"name"`
	ParentFollowupIntentName string                                     `json:"parent_followup_intent_name"`
	Priority                 float64                                    `json:"priority"`
	Project                  string                                     `json:"project"`
	ResetContexts            bool                                       `json:"reset_contexts"`
	RootFollowupIntentName   string                                     `json:"root_followup_intent_name"`
	WebhookState             string                                     `json:"webhook_state"`
	FollowupIntentInfo       []dialogflowintent.FollowupIntentInfoState `json:"followup_intent_info"`
	Timeouts                 *dialogflowintent.TimeoutsState            `json:"timeouts"`
}
