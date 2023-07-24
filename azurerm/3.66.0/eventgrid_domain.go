// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventgriddomain "github.com/golingon/terraproviders/azurerm/3.66.0/eventgriddomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventgridDomain creates a new instance of [EventgridDomain].
func NewEventgridDomain(name string, args EventgridDomainArgs) *EventgridDomain {
	return &EventgridDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventgridDomain)(nil)

// EventgridDomain represents the Terraform resource azurerm_eventgrid_domain.
type EventgridDomain struct {
	Name      string
	Args      EventgridDomainArgs
	state     *eventgridDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventgridDomain].
func (ed *EventgridDomain) Type() string {
	return "azurerm_eventgrid_domain"
}

// LocalName returns the local name for [EventgridDomain].
func (ed *EventgridDomain) LocalName() string {
	return ed.Name
}

// Configuration returns the configuration (args) for [EventgridDomain].
func (ed *EventgridDomain) Configuration() interface{} {
	return ed.Args
}

// DependOn is used for other resources to depend on [EventgridDomain].
func (ed *EventgridDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(ed)
}

// Dependencies returns the list of resources [EventgridDomain] depends_on.
func (ed *EventgridDomain) Dependencies() terra.Dependencies {
	return ed.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventgridDomain].
func (ed *EventgridDomain) LifecycleManagement() *terra.Lifecycle {
	return ed.Lifecycle
}

// Attributes returns the attributes for [EventgridDomain].
func (ed *EventgridDomain) Attributes() eventgridDomainAttributes {
	return eventgridDomainAttributes{ref: terra.ReferenceResource(ed)}
}

// ImportState imports the given attribute values into [EventgridDomain]'s state.
func (ed *EventgridDomain) ImportState(av io.Reader) error {
	ed.state = &eventgridDomainState{}
	if err := json.NewDecoder(av).Decode(ed.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ed.Type(), ed.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventgridDomain] has state.
func (ed *EventgridDomain) State() (*eventgridDomainState, bool) {
	return ed.state, ed.state != nil
}

// StateMust returns the state for [EventgridDomain]. Panics if the state is nil.
func (ed *EventgridDomain) StateMust() *eventgridDomainState {
	if ed.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ed.Type(), ed.LocalName()))
	}
	return ed.state
}

// EventgridDomainArgs contains the configurations for azurerm_eventgrid_domain.
type EventgridDomainArgs struct {
	// AutoCreateTopicWithFirstSubscription: bool, optional
	AutoCreateTopicWithFirstSubscription terra.BoolValue `hcl:"auto_create_topic_with_first_subscription,attr"`
	// AutoDeleteTopicWithLastSubscription: bool, optional
	AutoDeleteTopicWithLastSubscription terra.BoolValue `hcl:"auto_delete_topic_with_last_subscription,attr"`
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
	InboundIpRule []eventgriddomain.InboundIpRule `hcl:"inbound_ip_rule,block" validate:"min=0"`
	// Identity: optional
	Identity *eventgriddomain.Identity `hcl:"identity,block"`
	// InputMappingDefaultValues: optional
	InputMappingDefaultValues *eventgriddomain.InputMappingDefaultValues `hcl:"input_mapping_default_values,block"`
	// InputMappingFields: optional
	InputMappingFields *eventgriddomain.InputMappingFields `hcl:"input_mapping_fields,block"`
	// Timeouts: optional
	Timeouts *eventgriddomain.Timeouts `hcl:"timeouts,block"`
}
type eventgridDomainAttributes struct {
	ref terra.Reference
}

// AutoCreateTopicWithFirstSubscription returns a reference to field auto_create_topic_with_first_subscription of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) AutoCreateTopicWithFirstSubscription() terra.BoolValue {
	return terra.ReferenceAsBool(ed.ref.Append("auto_create_topic_with_first_subscription"))
}

// AutoDeleteTopicWithLastSubscription returns a reference to field auto_delete_topic_with_last_subscription of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) AutoDeleteTopicWithLastSubscription() terra.BoolValue {
	return terra.ReferenceAsBool(ed.ref.Append("auto_delete_topic_with_last_subscription"))
}

// Endpoint returns a reference to field endpoint of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("endpoint"))
}

// Id returns a reference to field id of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("id"))
}

// InputSchema returns a reference to field input_schema of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) InputSchema() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("input_schema"))
}

// LocalAuthEnabled returns a reference to field local_auth_enabled of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ed.ref.Append("local_auth_enabled"))
}

// Location returns a reference to field location of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("name"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("primary_access_key"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ed.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("secondary_access_key"))
}

// Tags returns a reference to field tags of azurerm_eventgrid_domain.
func (ed eventgridDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ed.ref.Append("tags"))
}

func (ed eventgridDomainAttributes) InboundIpRule() terra.ListValue[eventgriddomain.InboundIpRuleAttributes] {
	return terra.ReferenceAsList[eventgriddomain.InboundIpRuleAttributes](ed.ref.Append("inbound_ip_rule"))
}

func (ed eventgridDomainAttributes) Identity() terra.ListValue[eventgriddomain.IdentityAttributes] {
	return terra.ReferenceAsList[eventgriddomain.IdentityAttributes](ed.ref.Append("identity"))
}

func (ed eventgridDomainAttributes) InputMappingDefaultValues() terra.ListValue[eventgriddomain.InputMappingDefaultValuesAttributes] {
	return terra.ReferenceAsList[eventgriddomain.InputMappingDefaultValuesAttributes](ed.ref.Append("input_mapping_default_values"))
}

func (ed eventgridDomainAttributes) InputMappingFields() terra.ListValue[eventgriddomain.InputMappingFieldsAttributes] {
	return terra.ReferenceAsList[eventgriddomain.InputMappingFieldsAttributes](ed.ref.Append("input_mapping_fields"))
}

func (ed eventgridDomainAttributes) Timeouts() eventgriddomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventgriddomain.TimeoutsAttributes](ed.ref.Append("timeouts"))
}

type eventgridDomainState struct {
	AutoCreateTopicWithFirstSubscription bool                                             `json:"auto_create_topic_with_first_subscription"`
	AutoDeleteTopicWithLastSubscription  bool                                             `json:"auto_delete_topic_with_last_subscription"`
	Endpoint                             string                                           `json:"endpoint"`
	Id                                   string                                           `json:"id"`
	InputSchema                          string                                           `json:"input_schema"`
	LocalAuthEnabled                     bool                                             `json:"local_auth_enabled"`
	Location                             string                                           `json:"location"`
	Name                                 string                                           `json:"name"`
	PrimaryAccessKey                     string                                           `json:"primary_access_key"`
	PublicNetworkAccessEnabled           bool                                             `json:"public_network_access_enabled"`
	ResourceGroupName                    string                                           `json:"resource_group_name"`
	SecondaryAccessKey                   string                                           `json:"secondary_access_key"`
	Tags                                 map[string]string                                `json:"tags"`
	InboundIpRule                        []eventgriddomain.InboundIpRuleState             `json:"inbound_ip_rule"`
	Identity                             []eventgriddomain.IdentityState                  `json:"identity"`
	InputMappingDefaultValues            []eventgriddomain.InputMappingDefaultValuesState `json:"input_mapping_default_values"`
	InputMappingFields                   []eventgriddomain.InputMappingFieldsState        `json:"input_mapping_fields"`
	Timeouts                             *eventgriddomain.TimeoutsState                   `json:"timeouts"`
}
