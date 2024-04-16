// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dialogflow_cx_page

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

// Resource represents the Terraform resource google_dialogflow_cx_page.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDialogflowCxPageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdcp *Resource) Type() string {
	return "google_dialogflow_cx_page"
}

// LocalName returns the local name for [Resource].
func (gdcp *Resource) LocalName() string {
	return gdcp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdcp *Resource) Configuration() interface{} {
	return gdcp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdcp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdcp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdcp *Resource) Dependencies() terra.Dependencies {
	return gdcp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdcp *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdcp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdcp *Resource) Attributes() googleDialogflowCxPageAttributes {
	return googleDialogflowCxPageAttributes{ref: terra.ReferenceResource(gdcp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdcp *Resource) ImportState(state io.Reader) error {
	gdcp.state = &googleDialogflowCxPageState{}
	if err := json.NewDecoder(state).Decode(gdcp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdcp.Type(), gdcp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdcp *Resource) State() (*googleDialogflowCxPageState, bool) {
	return gdcp.state, gdcp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdcp *Resource) StateMust() *googleDialogflowCxPageState {
	if gdcp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdcp.Type(), gdcp.LocalName()))
	}
	return gdcp.state
}

// Args contains the configurations for google_dialogflow_cx_page.
type Args struct {
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
	// AdvancedSettings: optional
	AdvancedSettings *AdvancedSettings `hcl:"advanced_settings,block"`
	// EntryFulfillment: optional
	EntryFulfillment *EntryFulfillment `hcl:"entry_fulfillment,block"`
	// EventHandlers: min=0
	EventHandlers []EventHandlers `hcl:"event_handlers,block" validate:"min=0"`
	// Form: optional
	Form *Form `hcl:"form,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// TransitionRoutes: min=0
	TransitionRoutes []TransitionRoutes `hcl:"transition_routes,block" validate:"min=0"`
}

type googleDialogflowCxPageAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_dialogflow_cx_page.
func (gdcp googleDialogflowCxPageAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gdcp.ref.Append("display_name"))
}

// Id returns a reference to field id of google_dialogflow_cx_page.
func (gdcp googleDialogflowCxPageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdcp.ref.Append("id"))
}

// LanguageCode returns a reference to field language_code of google_dialogflow_cx_page.
func (gdcp googleDialogflowCxPageAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(gdcp.ref.Append("language_code"))
}

// Name returns a reference to field name of google_dialogflow_cx_page.
func (gdcp googleDialogflowCxPageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gdcp.ref.Append("name"))
}

// Parent returns a reference to field parent of google_dialogflow_cx_page.
func (gdcp googleDialogflowCxPageAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(gdcp.ref.Append("parent"))
}

// TransitionRouteGroups returns a reference to field transition_route_groups of google_dialogflow_cx_page.
func (gdcp googleDialogflowCxPageAttributes) TransitionRouteGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gdcp.ref.Append("transition_route_groups"))
}

func (gdcp googleDialogflowCxPageAttributes) AdvancedSettings() terra.ListValue[AdvancedSettingsAttributes] {
	return terra.ReferenceAsList[AdvancedSettingsAttributes](gdcp.ref.Append("advanced_settings"))
}

func (gdcp googleDialogflowCxPageAttributes) EntryFulfillment() terra.ListValue[EntryFulfillmentAttributes] {
	return terra.ReferenceAsList[EntryFulfillmentAttributes](gdcp.ref.Append("entry_fulfillment"))
}

func (gdcp googleDialogflowCxPageAttributes) EventHandlers() terra.ListValue[EventHandlersAttributes] {
	return terra.ReferenceAsList[EventHandlersAttributes](gdcp.ref.Append("event_handlers"))
}

func (gdcp googleDialogflowCxPageAttributes) Form() terra.ListValue[FormAttributes] {
	return terra.ReferenceAsList[FormAttributes](gdcp.ref.Append("form"))
}

func (gdcp googleDialogflowCxPageAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gdcp.ref.Append("timeouts"))
}

func (gdcp googleDialogflowCxPageAttributes) TransitionRoutes() terra.ListValue[TransitionRoutesAttributes] {
	return terra.ReferenceAsList[TransitionRoutesAttributes](gdcp.ref.Append("transition_routes"))
}

type googleDialogflowCxPageState struct {
	DisplayName           string                  `json:"display_name"`
	Id                    string                  `json:"id"`
	LanguageCode          string                  `json:"language_code"`
	Name                  string                  `json:"name"`
	Parent                string                  `json:"parent"`
	TransitionRouteGroups []string                `json:"transition_route_groups"`
	AdvancedSettings      []AdvancedSettingsState `json:"advanced_settings"`
	EntryFulfillment      []EntryFulfillmentState `json:"entry_fulfillment"`
	EventHandlers         []EventHandlersState    `json:"event_handlers"`
	Form                  []FormState             `json:"form"`
	Timeouts              *TimeoutsState          `json:"timeouts"`
	TransitionRoutes      []TransitionRoutesState `json:"transition_routes"`
}
