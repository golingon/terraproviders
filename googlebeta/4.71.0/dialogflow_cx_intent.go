// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dialogflowcxintent "github.com/golingon/terraproviders/googlebeta/4.71.0/dialogflowcxintent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowCxIntent creates a new instance of [DialogflowCxIntent].
func NewDialogflowCxIntent(name string, args DialogflowCxIntentArgs) *DialogflowCxIntent {
	return &DialogflowCxIntent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowCxIntent)(nil)

// DialogflowCxIntent represents the Terraform resource google_dialogflow_cx_intent.
type DialogflowCxIntent struct {
	Name      string
	Args      DialogflowCxIntentArgs
	state     *dialogflowCxIntentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowCxIntent].
func (dci *DialogflowCxIntent) Type() string {
	return "google_dialogflow_cx_intent"
}

// LocalName returns the local name for [DialogflowCxIntent].
func (dci *DialogflowCxIntent) LocalName() string {
	return dci.Name
}

// Configuration returns the configuration (args) for [DialogflowCxIntent].
func (dci *DialogflowCxIntent) Configuration() interface{} {
	return dci.Args
}

// DependOn is used for other resources to depend on [DialogflowCxIntent].
func (dci *DialogflowCxIntent) DependOn() terra.Reference {
	return terra.ReferenceResource(dci)
}

// Dependencies returns the list of resources [DialogflowCxIntent] depends_on.
func (dci *DialogflowCxIntent) Dependencies() terra.Dependencies {
	return dci.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowCxIntent].
func (dci *DialogflowCxIntent) LifecycleManagement() *terra.Lifecycle {
	return dci.Lifecycle
}

// Attributes returns the attributes for [DialogflowCxIntent].
func (dci *DialogflowCxIntent) Attributes() dialogflowCxIntentAttributes {
	return dialogflowCxIntentAttributes{ref: terra.ReferenceResource(dci)}
}

// ImportState imports the given attribute values into [DialogflowCxIntent]'s state.
func (dci *DialogflowCxIntent) ImportState(av io.Reader) error {
	dci.state = &dialogflowCxIntentState{}
	if err := json.NewDecoder(av).Decode(dci.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dci.Type(), dci.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowCxIntent] has state.
func (dci *DialogflowCxIntent) State() (*dialogflowCxIntentState, bool) {
	return dci.state, dci.state != nil
}

// StateMust returns the state for [DialogflowCxIntent]. Panics if the state is nil.
func (dci *DialogflowCxIntent) StateMust() *dialogflowCxIntentState {
	if dci.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dci.Type(), dci.LocalName()))
	}
	return dci.state
}

// DialogflowCxIntentArgs contains the configurations for google_dialogflow_cx_intent.
type DialogflowCxIntentArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsFallback: bool, optional
	IsFallback terra.BoolValue `hcl:"is_fallback,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// LanguageCode: string, optional
	LanguageCode terra.StringValue `hcl:"language_code,attr"`
	// Parent: string, optional
	Parent terra.StringValue `hcl:"parent,attr"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// Parameters: min=0
	Parameters []dialogflowcxintent.Parameters `hcl:"parameters,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dialogflowcxintent.Timeouts `hcl:"timeouts,block"`
	// TrainingPhrases: min=0
	TrainingPhrases []dialogflowcxintent.TrainingPhrases `hcl:"training_phrases,block" validate:"min=0"`
}
type dialogflowCxIntentAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_dialogflow_cx_intent.
func (dci dialogflowCxIntentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dialogflow_cx_intent.
func (dci dialogflowCxIntentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("display_name"))
}

// Id returns a reference to field id of google_dialogflow_cx_intent.
func (dci dialogflowCxIntentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("id"))
}

// IsFallback returns a reference to field is_fallback of google_dialogflow_cx_intent.
func (dci dialogflowCxIntentAttributes) IsFallback() terra.BoolValue {
	return terra.ReferenceAsBool(dci.ref.Append("is_fallback"))
}

// Labels returns a reference to field labels of google_dialogflow_cx_intent.
func (dci dialogflowCxIntentAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dci.ref.Append("labels"))
}

// LanguageCode returns a reference to field language_code of google_dialogflow_cx_intent.
func (dci dialogflowCxIntentAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("language_code"))
}

// Name returns a reference to field name of google_dialogflow_cx_intent.
func (dci dialogflowCxIntentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("name"))
}

// Parent returns a reference to field parent of google_dialogflow_cx_intent.
func (dci dialogflowCxIntentAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("parent"))
}

// Priority returns a reference to field priority of google_dialogflow_cx_intent.
func (dci dialogflowCxIntentAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(dci.ref.Append("priority"))
}

func (dci dialogflowCxIntentAttributes) Parameters() terra.ListValue[dialogflowcxintent.ParametersAttributes] {
	return terra.ReferenceAsList[dialogflowcxintent.ParametersAttributes](dci.ref.Append("parameters"))
}

func (dci dialogflowCxIntentAttributes) Timeouts() dialogflowcxintent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowcxintent.TimeoutsAttributes](dci.ref.Append("timeouts"))
}

func (dci dialogflowCxIntentAttributes) TrainingPhrases() terra.ListValue[dialogflowcxintent.TrainingPhrasesAttributes] {
	return terra.ReferenceAsList[dialogflowcxintent.TrainingPhrasesAttributes](dci.ref.Append("training_phrases"))
}

type dialogflowCxIntentState struct {
	Description     string                                    `json:"description"`
	DisplayName     string                                    `json:"display_name"`
	Id              string                                    `json:"id"`
	IsFallback      bool                                      `json:"is_fallback"`
	Labels          map[string]string                         `json:"labels"`
	LanguageCode    string                                    `json:"language_code"`
	Name            string                                    `json:"name"`
	Parent          string                                    `json:"parent"`
	Priority        float64                                   `json:"priority"`
	Parameters      []dialogflowcxintent.ParametersState      `json:"parameters"`
	Timeouts        *dialogflowcxintent.TimeoutsState         `json:"timeouts"`
	TrainingPhrases []dialogflowcxintent.TrainingPhrasesState `json:"training_phrases"`
}
