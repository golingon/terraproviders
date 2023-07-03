// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dialogflowcxflow "github.com/golingon/terraproviders/googlebeta/4.71.0/dialogflowcxflow"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowCxFlow creates a new instance of [DialogflowCxFlow].
func NewDialogflowCxFlow(name string, args DialogflowCxFlowArgs) *DialogflowCxFlow {
	return &DialogflowCxFlow{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowCxFlow)(nil)

// DialogflowCxFlow represents the Terraform resource google_dialogflow_cx_flow.
type DialogflowCxFlow struct {
	Name      string
	Args      DialogflowCxFlowArgs
	state     *dialogflowCxFlowState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowCxFlow].
func (dcf *DialogflowCxFlow) Type() string {
	return "google_dialogflow_cx_flow"
}

// LocalName returns the local name for [DialogflowCxFlow].
func (dcf *DialogflowCxFlow) LocalName() string {
	return dcf.Name
}

// Configuration returns the configuration (args) for [DialogflowCxFlow].
func (dcf *DialogflowCxFlow) Configuration() interface{} {
	return dcf.Args
}

// DependOn is used for other resources to depend on [DialogflowCxFlow].
func (dcf *DialogflowCxFlow) DependOn() terra.Reference {
	return terra.ReferenceResource(dcf)
}

// Dependencies returns the list of resources [DialogflowCxFlow] depends_on.
func (dcf *DialogflowCxFlow) Dependencies() terra.Dependencies {
	return dcf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowCxFlow].
func (dcf *DialogflowCxFlow) LifecycleManagement() *terra.Lifecycle {
	return dcf.Lifecycle
}

// Attributes returns the attributes for [DialogflowCxFlow].
func (dcf *DialogflowCxFlow) Attributes() dialogflowCxFlowAttributes {
	return dialogflowCxFlowAttributes{ref: terra.ReferenceResource(dcf)}
}

// ImportState imports the given attribute values into [DialogflowCxFlow]'s state.
func (dcf *DialogflowCxFlow) ImportState(av io.Reader) error {
	dcf.state = &dialogflowCxFlowState{}
	if err := json.NewDecoder(av).Decode(dcf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcf.Type(), dcf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowCxFlow] has state.
func (dcf *DialogflowCxFlow) State() (*dialogflowCxFlowState, bool) {
	return dcf.state, dcf.state != nil
}

// StateMust returns the state for [DialogflowCxFlow]. Panics if the state is nil.
func (dcf *DialogflowCxFlow) StateMust() *dialogflowCxFlowState {
	if dcf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcf.Type(), dcf.LocalName()))
	}
	return dcf.state
}

// DialogflowCxFlowArgs contains the configurations for google_dialogflow_cx_flow.
type DialogflowCxFlowArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LanguageCode: string, optional
	LanguageCode terra.StringValue `hcl:"language_code,attr"`
	// Parent: string, optional
	Parent terra.StringValue `hcl:"parent,attr"`
	// TransitionRouteGroups: list of string, optional
	TransitionRouteGroups terra.ListValue[terra.StringValue] `hcl:"transition_route_groups,attr"`
	// EventHandlers: min=0
	EventHandlers []dialogflowcxflow.EventHandlers `hcl:"event_handlers,block" validate:"min=0"`
	// NluSettings: optional
	NluSettings *dialogflowcxflow.NluSettings `hcl:"nlu_settings,block"`
	// Timeouts: optional
	Timeouts *dialogflowcxflow.Timeouts `hcl:"timeouts,block"`
	// TransitionRoutes: min=0
	TransitionRoutes []dialogflowcxflow.TransitionRoutes `hcl:"transition_routes,block" validate:"min=0"`
}
type dialogflowCxFlowAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_dialogflow_cx_flow.
func (dcf dialogflowCxFlowAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dcf.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dialogflow_cx_flow.
func (dcf dialogflowCxFlowAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dcf.ref.Append("display_name"))
}

// Id returns a reference to field id of google_dialogflow_cx_flow.
func (dcf dialogflowCxFlowAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcf.ref.Append("id"))
}

// LanguageCode returns a reference to field language_code of google_dialogflow_cx_flow.
func (dcf dialogflowCxFlowAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(dcf.ref.Append("language_code"))
}

// Name returns a reference to field name of google_dialogflow_cx_flow.
func (dcf dialogflowCxFlowAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcf.ref.Append("name"))
}

// Parent returns a reference to field parent of google_dialogflow_cx_flow.
func (dcf dialogflowCxFlowAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(dcf.ref.Append("parent"))
}

// TransitionRouteGroups returns a reference to field transition_route_groups of google_dialogflow_cx_flow.
func (dcf dialogflowCxFlowAttributes) TransitionRouteGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dcf.ref.Append("transition_route_groups"))
}

func (dcf dialogflowCxFlowAttributes) EventHandlers() terra.ListValue[dialogflowcxflow.EventHandlersAttributes] {
	return terra.ReferenceAsList[dialogflowcxflow.EventHandlersAttributes](dcf.ref.Append("event_handlers"))
}

func (dcf dialogflowCxFlowAttributes) NluSettings() terra.ListValue[dialogflowcxflow.NluSettingsAttributes] {
	return terra.ReferenceAsList[dialogflowcxflow.NluSettingsAttributes](dcf.ref.Append("nlu_settings"))
}

func (dcf dialogflowCxFlowAttributes) Timeouts() dialogflowcxflow.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowcxflow.TimeoutsAttributes](dcf.ref.Append("timeouts"))
}

func (dcf dialogflowCxFlowAttributes) TransitionRoutes() terra.ListValue[dialogflowcxflow.TransitionRoutesAttributes] {
	return terra.ReferenceAsList[dialogflowcxflow.TransitionRoutesAttributes](dcf.ref.Append("transition_routes"))
}

type dialogflowCxFlowState struct {
	Description           string                                   `json:"description"`
	DisplayName           string                                   `json:"display_name"`
	Id                    string                                   `json:"id"`
	LanguageCode          string                                   `json:"language_code"`
	Name                  string                                   `json:"name"`
	Parent                string                                   `json:"parent"`
	TransitionRouteGroups []string                                 `json:"transition_route_groups"`
	EventHandlers         []dialogflowcxflow.EventHandlersState    `json:"event_handlers"`
	NluSettings           []dialogflowcxflow.NluSettingsState      `json:"nlu_settings"`
	Timeouts              *dialogflowcxflow.TimeoutsState          `json:"timeouts"`
	TransitionRoutes      []dialogflowcxflow.TransitionRoutesState `json:"transition_routes"`
}
