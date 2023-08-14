// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventgridtopic "github.com/golingon/terraproviders/azurerm/3.69.0/eventgridtopic"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventgridTopic creates a new instance of [EventgridTopic].
func NewEventgridTopic(name string, args EventgridTopicArgs) *EventgridTopic {
	return &EventgridTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventgridTopic)(nil)

// EventgridTopic represents the Terraform resource azurerm_eventgrid_topic.
type EventgridTopic struct {
	Name      string
	Args      EventgridTopicArgs
	state     *eventgridTopicState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventgridTopic].
func (et *EventgridTopic) Type() string {
	return "azurerm_eventgrid_topic"
}

// LocalName returns the local name for [EventgridTopic].
func (et *EventgridTopic) LocalName() string {
	return et.Name
}

// Configuration returns the configuration (args) for [EventgridTopic].
func (et *EventgridTopic) Configuration() interface{} {
	return et.Args
}

// DependOn is used for other resources to depend on [EventgridTopic].
func (et *EventgridTopic) DependOn() terra.Reference {
	return terra.ReferenceResource(et)
}

// Dependencies returns the list of resources [EventgridTopic] depends_on.
func (et *EventgridTopic) Dependencies() terra.Dependencies {
	return et.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventgridTopic].
func (et *EventgridTopic) LifecycleManagement() *terra.Lifecycle {
	return et.Lifecycle
}

// Attributes returns the attributes for [EventgridTopic].
func (et *EventgridTopic) Attributes() eventgridTopicAttributes {
	return eventgridTopicAttributes{ref: terra.ReferenceResource(et)}
}

// ImportState imports the given attribute values into [EventgridTopic]'s state.
func (et *EventgridTopic) ImportState(av io.Reader) error {
	et.state = &eventgridTopicState{}
	if err := json.NewDecoder(av).Decode(et.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", et.Type(), et.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventgridTopic] has state.
func (et *EventgridTopic) State() (*eventgridTopicState, bool) {
	return et.state, et.state != nil
}

// StateMust returns the state for [EventgridTopic]. Panics if the state is nil.
func (et *EventgridTopic) StateMust() *eventgridTopicState {
	if et.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", et.Type(), et.LocalName()))
	}
	return et.state
}

// EventgridTopicArgs contains the configurations for azurerm_eventgrid_topic.
type EventgridTopicArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InputSchema: string, optional
	InputSchema terra.StringValue `hcl:"input_schema,attr"`
	// LocalAuthEnabled: bool, optional
	LocalAuthEnabled terra.BoolValue `hcl:"local_auth_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// InboundIpRule: min=0
	InboundIpRule []eventgridtopic.InboundIpRule `hcl:"inbound_ip_rule,block" validate:"min=0"`
	// Identity: optional
	Identity *eventgridtopic.Identity `hcl:"identity,block"`
	// InputMappingDefaultValues: optional
	InputMappingDefaultValues *eventgridtopic.InputMappingDefaultValues `hcl:"input_mapping_default_values,block"`
	// InputMappingFields: optional
	InputMappingFields *eventgridtopic.InputMappingFields `hcl:"input_mapping_fields,block"`
	// Timeouts: optional
	Timeouts *eventgridtopic.Timeouts `hcl:"timeouts,block"`
}
type eventgridTopicAttributes struct {
	ref terra.Reference
}

// Endpoint returns a reference to field endpoint of azurerm_eventgrid_topic.
func (et eventgridTopicAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("endpoint"))
}

// Id returns a reference to field id of azurerm_eventgrid_topic.
func (et eventgridTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("id"))
}

// InputSchema returns a reference to field input_schema of azurerm_eventgrid_topic.
func (et eventgridTopicAttributes) InputSchema() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("input_schema"))
}

// LocalAuthEnabled returns a reference to field local_auth_enabled of azurerm_eventgrid_topic.
func (et eventgridTopicAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(et.ref.Append("local_auth_enabled"))
}

// Location returns a reference to field location of azurerm_eventgrid_topic.
func (et eventgridTopicAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_eventgrid_topic.
func (et eventgridTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("name"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_eventgrid_topic.
func (et eventgridTopicAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("primary_access_key"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_eventgrid_topic.
func (et eventgridTopicAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(et.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventgrid_topic.
func (et eventgridTopicAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_eventgrid_topic.
func (et eventgridTopicAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("secondary_access_key"))
}

// Tags returns a reference to field tags of azurerm_eventgrid_topic.
func (et eventgridTopicAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](et.ref.Append("tags"))
}

func (et eventgridTopicAttributes) InboundIpRule() terra.ListValue[eventgridtopic.InboundIpRuleAttributes] {
	return terra.ReferenceAsList[eventgridtopic.InboundIpRuleAttributes](et.ref.Append("inbound_ip_rule"))
}

func (et eventgridTopicAttributes) Identity() terra.ListValue[eventgridtopic.IdentityAttributes] {
	return terra.ReferenceAsList[eventgridtopic.IdentityAttributes](et.ref.Append("identity"))
}

func (et eventgridTopicAttributes) InputMappingDefaultValues() terra.ListValue[eventgridtopic.InputMappingDefaultValuesAttributes] {
	return terra.ReferenceAsList[eventgridtopic.InputMappingDefaultValuesAttributes](et.ref.Append("input_mapping_default_values"))
}

func (et eventgridTopicAttributes) InputMappingFields() terra.ListValue[eventgridtopic.InputMappingFieldsAttributes] {
	return terra.ReferenceAsList[eventgridtopic.InputMappingFieldsAttributes](et.ref.Append("input_mapping_fields"))
}

func (et eventgridTopicAttributes) Timeouts() eventgridtopic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventgridtopic.TimeoutsAttributes](et.ref.Append("timeouts"))
}

type eventgridTopicState struct {
	Endpoint                   string                                          `json:"endpoint"`
	Id                         string                                          `json:"id"`
	InputSchema                string                                          `json:"input_schema"`
	LocalAuthEnabled           bool                                            `json:"local_auth_enabled"`
	Location                   string                                          `json:"location"`
	Name                       string                                          `json:"name"`
	PrimaryAccessKey           string                                          `json:"primary_access_key"`
	PublicNetworkAccessEnabled bool                                            `json:"public_network_access_enabled"`
	ResourceGroupName          string                                          `json:"resource_group_name"`
	SecondaryAccessKey         string                                          `json:"secondary_access_key"`
	Tags                       map[string]string                               `json:"tags"`
	InboundIpRule              []eventgridtopic.InboundIpRuleState             `json:"inbound_ip_rule"`
	Identity                   []eventgridtopic.IdentityState                  `json:"identity"`
	InputMappingDefaultValues  []eventgridtopic.InputMappingDefaultValuesState `json:"input_mapping_default_values"`
	InputMappingFields         []eventgridtopic.InputMappingFieldsState        `json:"input_mapping_fields"`
	Timeouts                   *eventgridtopic.TimeoutsState                   `json:"timeouts"`
}
