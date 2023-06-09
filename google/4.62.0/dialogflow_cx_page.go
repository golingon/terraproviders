// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dialogflowcxpage "github.com/golingon/terraproviders/google/4.62.0/dialogflowcxpage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowCxPage creates a new instance of [DialogflowCxPage].
func NewDialogflowCxPage(name string, args DialogflowCxPageArgs) *DialogflowCxPage {
	return &DialogflowCxPage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowCxPage)(nil)

// DialogflowCxPage represents the Terraform resource google_dialogflow_cx_page.
type DialogflowCxPage struct {
	Name      string
	Args      DialogflowCxPageArgs
	state     *dialogflowCxPageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowCxPage].
func (dcp *DialogflowCxPage) Type() string {
	return "google_dialogflow_cx_page"
}

// LocalName returns the local name for [DialogflowCxPage].
func (dcp *DialogflowCxPage) LocalName() string {
	return dcp.Name
}

// Configuration returns the configuration (args) for [DialogflowCxPage].
func (dcp *DialogflowCxPage) Configuration() interface{} {
	return dcp.Args
}

// DependOn is used for other resources to depend on [DialogflowCxPage].
func (dcp *DialogflowCxPage) DependOn() terra.Reference {
	return terra.ReferenceResource(dcp)
}

// Dependencies returns the list of resources [DialogflowCxPage] depends_on.
func (dcp *DialogflowCxPage) Dependencies() terra.Dependencies {
	return dcp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowCxPage].
func (dcp *DialogflowCxPage) LifecycleManagement() *terra.Lifecycle {
	return dcp.Lifecycle
}

// Attributes returns the attributes for [DialogflowCxPage].
func (dcp *DialogflowCxPage) Attributes() dialogflowCxPageAttributes {
	return dialogflowCxPageAttributes{ref: terra.ReferenceResource(dcp)}
}

// ImportState imports the given attribute values into [DialogflowCxPage]'s state.
func (dcp *DialogflowCxPage) ImportState(av io.Reader) error {
	dcp.state = &dialogflowCxPageState{}
	if err := json.NewDecoder(av).Decode(dcp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcp.Type(), dcp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowCxPage] has state.
func (dcp *DialogflowCxPage) State() (*dialogflowCxPageState, bool) {
	return dcp.state, dcp.state != nil
}

// StateMust returns the state for [DialogflowCxPage]. Panics if the state is nil.
func (dcp *DialogflowCxPage) StateMust() *dialogflowCxPageState {
	if dcp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcp.Type(), dcp.LocalName()))
	}
	return dcp.state
}

// DialogflowCxPageArgs contains the configurations for google_dialogflow_cx_page.
type DialogflowCxPageArgs struct {
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
	// EntryFulfillment: optional
	EntryFulfillment *dialogflowcxpage.EntryFulfillment `hcl:"entry_fulfillment,block"`
	// EventHandlers: min=0
	EventHandlers []dialogflowcxpage.EventHandlers `hcl:"event_handlers,block" validate:"min=0"`
	// Form: optional
	Form *dialogflowcxpage.Form `hcl:"form,block"`
	// Timeouts: optional
	Timeouts *dialogflowcxpage.Timeouts `hcl:"timeouts,block"`
	// TransitionRoutes: min=0
	TransitionRoutes []dialogflowcxpage.TransitionRoutes `hcl:"transition_routes,block" validate:"min=0"`
}
type dialogflowCxPageAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_dialogflow_cx_page.
func (dcp dialogflowCxPageAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dcp.ref.Append("display_name"))
}

// Id returns a reference to field id of google_dialogflow_cx_page.
func (dcp dialogflowCxPageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcp.ref.Append("id"))
}

// LanguageCode returns a reference to field language_code of google_dialogflow_cx_page.
func (dcp dialogflowCxPageAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(dcp.ref.Append("language_code"))
}

// Name returns a reference to field name of google_dialogflow_cx_page.
func (dcp dialogflowCxPageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcp.ref.Append("name"))
}

// Parent returns a reference to field parent of google_dialogflow_cx_page.
func (dcp dialogflowCxPageAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(dcp.ref.Append("parent"))
}

// TransitionRouteGroups returns a reference to field transition_route_groups of google_dialogflow_cx_page.
func (dcp dialogflowCxPageAttributes) TransitionRouteGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dcp.ref.Append("transition_route_groups"))
}

func (dcp dialogflowCxPageAttributes) EntryFulfillment() terra.ListValue[dialogflowcxpage.EntryFulfillmentAttributes] {
	return terra.ReferenceAsList[dialogflowcxpage.EntryFulfillmentAttributes](dcp.ref.Append("entry_fulfillment"))
}

func (dcp dialogflowCxPageAttributes) EventHandlers() terra.ListValue[dialogflowcxpage.EventHandlersAttributes] {
	return terra.ReferenceAsList[dialogflowcxpage.EventHandlersAttributes](dcp.ref.Append("event_handlers"))
}

func (dcp dialogflowCxPageAttributes) Form() terra.ListValue[dialogflowcxpage.FormAttributes] {
	return terra.ReferenceAsList[dialogflowcxpage.FormAttributes](dcp.ref.Append("form"))
}

func (dcp dialogflowCxPageAttributes) Timeouts() dialogflowcxpage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowcxpage.TimeoutsAttributes](dcp.ref.Append("timeouts"))
}

func (dcp dialogflowCxPageAttributes) TransitionRoutes() terra.ListValue[dialogflowcxpage.TransitionRoutesAttributes] {
	return terra.ReferenceAsList[dialogflowcxpage.TransitionRoutesAttributes](dcp.ref.Append("transition_routes"))
}

type dialogflowCxPageState struct {
	DisplayName           string                                   `json:"display_name"`
	Id                    string                                   `json:"id"`
	LanguageCode          string                                   `json:"language_code"`
	Name                  string                                   `json:"name"`
	Parent                string                                   `json:"parent"`
	TransitionRouteGroups []string                                 `json:"transition_route_groups"`
	EntryFulfillment      []dialogflowcxpage.EntryFulfillmentState `json:"entry_fulfillment"`
	EventHandlers         []dialogflowcxpage.EventHandlersState    `json:"event_handlers"`
	Form                  []dialogflowcxpage.FormState             `json:"form"`
	Timeouts              *dialogflowcxpage.TimeoutsState          `json:"timeouts"`
	TransitionRoutes      []dialogflowcxpage.TransitionRoutesState `json:"transition_routes"`
}
