// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	eventarctrigger "github.com/golingon/terraproviders/google/4.73.1/eventarctrigger"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventarcTrigger creates a new instance of [EventarcTrigger].
func NewEventarcTrigger(name string, args EventarcTriggerArgs) *EventarcTrigger {
	return &EventarcTrigger{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventarcTrigger)(nil)

// EventarcTrigger represents the Terraform resource google_eventarc_trigger.
type EventarcTrigger struct {
	Name      string
	Args      EventarcTriggerArgs
	state     *eventarcTriggerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventarcTrigger].
func (et *EventarcTrigger) Type() string {
	return "google_eventarc_trigger"
}

// LocalName returns the local name for [EventarcTrigger].
func (et *EventarcTrigger) LocalName() string {
	return et.Name
}

// Configuration returns the configuration (args) for [EventarcTrigger].
func (et *EventarcTrigger) Configuration() interface{} {
	return et.Args
}

// DependOn is used for other resources to depend on [EventarcTrigger].
func (et *EventarcTrigger) DependOn() terra.Reference {
	return terra.ReferenceResource(et)
}

// Dependencies returns the list of resources [EventarcTrigger] depends_on.
func (et *EventarcTrigger) Dependencies() terra.Dependencies {
	return et.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventarcTrigger].
func (et *EventarcTrigger) LifecycleManagement() *terra.Lifecycle {
	return et.Lifecycle
}

// Attributes returns the attributes for [EventarcTrigger].
func (et *EventarcTrigger) Attributes() eventarcTriggerAttributes {
	return eventarcTriggerAttributes{ref: terra.ReferenceResource(et)}
}

// ImportState imports the given attribute values into [EventarcTrigger]'s state.
func (et *EventarcTrigger) ImportState(av io.Reader) error {
	et.state = &eventarcTriggerState{}
	if err := json.NewDecoder(av).Decode(et.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", et.Type(), et.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventarcTrigger] has state.
func (et *EventarcTrigger) State() (*eventarcTriggerState, bool) {
	return et.state, et.state != nil
}

// StateMust returns the state for [EventarcTrigger]. Panics if the state is nil.
func (et *EventarcTrigger) StateMust() *eventarcTriggerState {
	if et.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", et.Type(), et.LocalName()))
	}
	return et.state
}

// EventarcTriggerArgs contains the configurations for google_eventarc_trigger.
type EventarcTriggerArgs struct {
	// Channel: string, optional
	Channel terra.StringValue `hcl:"channel,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServiceAccount: string, optional
	ServiceAccount terra.StringValue `hcl:"service_account,attr"`
	// Destination: required
	Destination *eventarctrigger.Destination `hcl:"destination,block" validate:"required"`
	// MatchingCriteria: min=1
	MatchingCriteria []eventarctrigger.MatchingCriteria `hcl:"matching_criteria,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *eventarctrigger.Timeouts `hcl:"timeouts,block"`
	// Transport: optional
	Transport *eventarctrigger.Transport `hcl:"transport,block"`
}
type eventarcTriggerAttributes struct {
	ref terra.Reference
}

// Channel returns a reference to field channel of google_eventarc_trigger.
func (et eventarcTriggerAttributes) Channel() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("channel"))
}

// Conditions returns a reference to field conditions of google_eventarc_trigger.
func (et eventarcTriggerAttributes) Conditions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](et.ref.Append("conditions"))
}

// CreateTime returns a reference to field create_time of google_eventarc_trigger.
func (et eventarcTriggerAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("create_time"))
}

// Etag returns a reference to field etag of google_eventarc_trigger.
func (et eventarcTriggerAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("etag"))
}

// Id returns a reference to field id of google_eventarc_trigger.
func (et eventarcTriggerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("id"))
}

// Labels returns a reference to field labels of google_eventarc_trigger.
func (et eventarcTriggerAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](et.ref.Append("labels"))
}

// Location returns a reference to field location of google_eventarc_trigger.
func (et eventarcTriggerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("location"))
}

// Name returns a reference to field name of google_eventarc_trigger.
func (et eventarcTriggerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("name"))
}

// Project returns a reference to field project of google_eventarc_trigger.
func (et eventarcTriggerAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("project"))
}

// ServiceAccount returns a reference to field service_account of google_eventarc_trigger.
func (et eventarcTriggerAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("service_account"))
}

// Uid returns a reference to field uid of google_eventarc_trigger.
func (et eventarcTriggerAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_eventarc_trigger.
func (et eventarcTriggerAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("update_time"))
}

func (et eventarcTriggerAttributes) Destination() terra.ListValue[eventarctrigger.DestinationAttributes] {
	return terra.ReferenceAsList[eventarctrigger.DestinationAttributes](et.ref.Append("destination"))
}

func (et eventarcTriggerAttributes) MatchingCriteria() terra.SetValue[eventarctrigger.MatchingCriteriaAttributes] {
	return terra.ReferenceAsSet[eventarctrigger.MatchingCriteriaAttributes](et.ref.Append("matching_criteria"))
}

func (et eventarcTriggerAttributes) Timeouts() eventarctrigger.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventarctrigger.TimeoutsAttributes](et.ref.Append("timeouts"))
}

func (et eventarcTriggerAttributes) Transport() terra.ListValue[eventarctrigger.TransportAttributes] {
	return terra.ReferenceAsList[eventarctrigger.TransportAttributes](et.ref.Append("transport"))
}

type eventarcTriggerState struct {
	Channel          string                                  `json:"channel"`
	Conditions       map[string]string                       `json:"conditions"`
	CreateTime       string                                  `json:"create_time"`
	Etag             string                                  `json:"etag"`
	Id               string                                  `json:"id"`
	Labels           map[string]string                       `json:"labels"`
	Location         string                                  `json:"location"`
	Name             string                                  `json:"name"`
	Project          string                                  `json:"project"`
	ServiceAccount   string                                  `json:"service_account"`
	Uid              string                                  `json:"uid"`
	UpdateTime       string                                  `json:"update_time"`
	Destination      []eventarctrigger.DestinationState      `json:"destination"`
	MatchingCriteria []eventarctrigger.MatchingCriteriaState `json:"matching_criteria"`
	Timeouts         *eventarctrigger.TimeoutsState          `json:"timeouts"`
	Transport        []eventarctrigger.TransportState        `json:"transport"`
}
