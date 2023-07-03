// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventgridsystemtopic "github.com/golingon/terraproviders/azurerm/3.63.0/eventgridsystemtopic"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventgridSystemTopic creates a new instance of [EventgridSystemTopic].
func NewEventgridSystemTopic(name string, args EventgridSystemTopicArgs) *EventgridSystemTopic {
	return &EventgridSystemTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventgridSystemTopic)(nil)

// EventgridSystemTopic represents the Terraform resource azurerm_eventgrid_system_topic.
type EventgridSystemTopic struct {
	Name      string
	Args      EventgridSystemTopicArgs
	state     *eventgridSystemTopicState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventgridSystemTopic].
func (est *EventgridSystemTopic) Type() string {
	return "azurerm_eventgrid_system_topic"
}

// LocalName returns the local name for [EventgridSystemTopic].
func (est *EventgridSystemTopic) LocalName() string {
	return est.Name
}

// Configuration returns the configuration (args) for [EventgridSystemTopic].
func (est *EventgridSystemTopic) Configuration() interface{} {
	return est.Args
}

// DependOn is used for other resources to depend on [EventgridSystemTopic].
func (est *EventgridSystemTopic) DependOn() terra.Reference {
	return terra.ReferenceResource(est)
}

// Dependencies returns the list of resources [EventgridSystemTopic] depends_on.
func (est *EventgridSystemTopic) Dependencies() terra.Dependencies {
	return est.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventgridSystemTopic].
func (est *EventgridSystemTopic) LifecycleManagement() *terra.Lifecycle {
	return est.Lifecycle
}

// Attributes returns the attributes for [EventgridSystemTopic].
func (est *EventgridSystemTopic) Attributes() eventgridSystemTopicAttributes {
	return eventgridSystemTopicAttributes{ref: terra.ReferenceResource(est)}
}

// ImportState imports the given attribute values into [EventgridSystemTopic]'s state.
func (est *EventgridSystemTopic) ImportState(av io.Reader) error {
	est.state = &eventgridSystemTopicState{}
	if err := json.NewDecoder(av).Decode(est.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", est.Type(), est.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventgridSystemTopic] has state.
func (est *EventgridSystemTopic) State() (*eventgridSystemTopicState, bool) {
	return est.state, est.state != nil
}

// StateMust returns the state for [EventgridSystemTopic]. Panics if the state is nil.
func (est *EventgridSystemTopic) StateMust() *eventgridSystemTopicState {
	if est.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", est.Type(), est.LocalName()))
	}
	return est.state
}

// EventgridSystemTopicArgs contains the configurations for azurerm_eventgrid_system_topic.
type EventgridSystemTopicArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SourceArmResourceId: string, required
	SourceArmResourceId terra.StringValue `hcl:"source_arm_resource_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TopicType: string, required
	TopicType terra.StringValue `hcl:"topic_type,attr" validate:"required"`
	// Identity: optional
	Identity *eventgridsystemtopic.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *eventgridsystemtopic.Timeouts `hcl:"timeouts,block"`
}
type eventgridSystemTopicAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_eventgrid_system_topic.
func (est eventgridSystemTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_eventgrid_system_topic.
func (est eventgridSystemTopicAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("location"))
}

// MetricArmResourceId returns a reference to field metric_arm_resource_id of azurerm_eventgrid_system_topic.
func (est eventgridSystemTopicAttributes) MetricArmResourceId() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("metric_arm_resource_id"))
}

// Name returns a reference to field name of azurerm_eventgrid_system_topic.
func (est eventgridSystemTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventgrid_system_topic.
func (est eventgridSystemTopicAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("resource_group_name"))
}

// SourceArmResourceId returns a reference to field source_arm_resource_id of azurerm_eventgrid_system_topic.
func (est eventgridSystemTopicAttributes) SourceArmResourceId() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("source_arm_resource_id"))
}

// Tags returns a reference to field tags of azurerm_eventgrid_system_topic.
func (est eventgridSystemTopicAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](est.ref.Append("tags"))
}

// TopicType returns a reference to field topic_type of azurerm_eventgrid_system_topic.
func (est eventgridSystemTopicAttributes) TopicType() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("topic_type"))
}

func (est eventgridSystemTopicAttributes) Identity() terra.ListValue[eventgridsystemtopic.IdentityAttributes] {
	return terra.ReferenceAsList[eventgridsystemtopic.IdentityAttributes](est.ref.Append("identity"))
}

func (est eventgridSystemTopicAttributes) Timeouts() eventgridsystemtopic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventgridsystemtopic.TimeoutsAttributes](est.ref.Append("timeouts"))
}

type eventgridSystemTopicState struct {
	Id                  string                               `json:"id"`
	Location            string                               `json:"location"`
	MetricArmResourceId string                               `json:"metric_arm_resource_id"`
	Name                string                               `json:"name"`
	ResourceGroupName   string                               `json:"resource_group_name"`
	SourceArmResourceId string                               `json:"source_arm_resource_id"`
	Tags                map[string]string                    `json:"tags"`
	TopicType           string                               `json:"topic_type"`
	Identity            []eventgridsystemtopic.IdentityState `json:"identity"`
	Timeouts            *eventgridsystemtopic.TimeoutsState  `json:"timeouts"`
}
