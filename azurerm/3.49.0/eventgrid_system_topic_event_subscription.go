// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventgridsystemtopiceventsubscription "github.com/golingon/terraproviders/azurerm/3.49.0/eventgridsystemtopiceventsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventgridSystemTopicEventSubscription creates a new instance of [EventgridSystemTopicEventSubscription].
func NewEventgridSystemTopicEventSubscription(name string, args EventgridSystemTopicEventSubscriptionArgs) *EventgridSystemTopicEventSubscription {
	return &EventgridSystemTopicEventSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventgridSystemTopicEventSubscription)(nil)

// EventgridSystemTopicEventSubscription represents the Terraform resource azurerm_eventgrid_system_topic_event_subscription.
type EventgridSystemTopicEventSubscription struct {
	Name      string
	Args      EventgridSystemTopicEventSubscriptionArgs
	state     *eventgridSystemTopicEventSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventgridSystemTopicEventSubscription].
func (estes *EventgridSystemTopicEventSubscription) Type() string {
	return "azurerm_eventgrid_system_topic_event_subscription"
}

// LocalName returns the local name for [EventgridSystemTopicEventSubscription].
func (estes *EventgridSystemTopicEventSubscription) LocalName() string {
	return estes.Name
}

// Configuration returns the configuration (args) for [EventgridSystemTopicEventSubscription].
func (estes *EventgridSystemTopicEventSubscription) Configuration() interface{} {
	return estes.Args
}

// DependOn is used for other resources to depend on [EventgridSystemTopicEventSubscription].
func (estes *EventgridSystemTopicEventSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(estes)
}

// Dependencies returns the list of resources [EventgridSystemTopicEventSubscription] depends_on.
func (estes *EventgridSystemTopicEventSubscription) Dependencies() terra.Dependencies {
	return estes.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventgridSystemTopicEventSubscription].
func (estes *EventgridSystemTopicEventSubscription) LifecycleManagement() *terra.Lifecycle {
	return estes.Lifecycle
}

// Attributes returns the attributes for [EventgridSystemTopicEventSubscription].
func (estes *EventgridSystemTopicEventSubscription) Attributes() eventgridSystemTopicEventSubscriptionAttributes {
	return eventgridSystemTopicEventSubscriptionAttributes{ref: terra.ReferenceResource(estes)}
}

// ImportState imports the given attribute values into [EventgridSystemTopicEventSubscription]'s state.
func (estes *EventgridSystemTopicEventSubscription) ImportState(av io.Reader) error {
	estes.state = &eventgridSystemTopicEventSubscriptionState{}
	if err := json.NewDecoder(av).Decode(estes.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", estes.Type(), estes.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventgridSystemTopicEventSubscription] has state.
func (estes *EventgridSystemTopicEventSubscription) State() (*eventgridSystemTopicEventSubscriptionState, bool) {
	return estes.state, estes.state != nil
}

// StateMust returns the state for [EventgridSystemTopicEventSubscription]. Panics if the state is nil.
func (estes *EventgridSystemTopicEventSubscription) StateMust() *eventgridSystemTopicEventSubscriptionState {
	if estes.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", estes.Type(), estes.LocalName()))
	}
	return estes.state
}

// EventgridSystemTopicEventSubscriptionArgs contains the configurations for azurerm_eventgrid_system_topic_event_subscription.
type EventgridSystemTopicEventSubscriptionArgs struct {
	// AdvancedFilteringOnArraysEnabled: bool, optional
	AdvancedFilteringOnArraysEnabled terra.BoolValue `hcl:"advanced_filtering_on_arrays_enabled,attr"`
	// EventDeliverySchema: string, optional
	EventDeliverySchema terra.StringValue `hcl:"event_delivery_schema,attr"`
	// EventhubEndpointId: string, optional
	EventhubEndpointId terra.StringValue `hcl:"eventhub_endpoint_id,attr"`
	// ExpirationTimeUtc: string, optional
	ExpirationTimeUtc terra.StringValue `hcl:"expiration_time_utc,attr"`
	// HybridConnectionEndpointId: string, optional
	HybridConnectionEndpointId terra.StringValue `hcl:"hybrid_connection_endpoint_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludedEventTypes: list of string, optional
	IncludedEventTypes terra.ListValue[terra.StringValue] `hcl:"included_event_types,attr"`
	// Labels: list of string, optional
	Labels terra.ListValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceBusQueueEndpointId: string, optional
	ServiceBusQueueEndpointId terra.StringValue `hcl:"service_bus_queue_endpoint_id,attr"`
	// ServiceBusTopicEndpointId: string, optional
	ServiceBusTopicEndpointId terra.StringValue `hcl:"service_bus_topic_endpoint_id,attr"`
	// SystemTopic: string, required
	SystemTopic terra.StringValue `hcl:"system_topic,attr" validate:"required"`
	// AdvancedFilter: optional
	AdvancedFilter *eventgridsystemtopiceventsubscription.AdvancedFilter `hcl:"advanced_filter,block"`
	// AzureFunctionEndpoint: optional
	AzureFunctionEndpoint *eventgridsystemtopiceventsubscription.AzureFunctionEndpoint `hcl:"azure_function_endpoint,block"`
	// DeadLetterIdentity: optional
	DeadLetterIdentity *eventgridsystemtopiceventsubscription.DeadLetterIdentity `hcl:"dead_letter_identity,block"`
	// DeliveryIdentity: optional
	DeliveryIdentity *eventgridsystemtopiceventsubscription.DeliveryIdentity `hcl:"delivery_identity,block"`
	// DeliveryProperty: min=0
	DeliveryProperty []eventgridsystemtopiceventsubscription.DeliveryProperty `hcl:"delivery_property,block" validate:"min=0"`
	// RetryPolicy: optional
	RetryPolicy *eventgridsystemtopiceventsubscription.RetryPolicy `hcl:"retry_policy,block"`
	// StorageBlobDeadLetterDestination: optional
	StorageBlobDeadLetterDestination *eventgridsystemtopiceventsubscription.StorageBlobDeadLetterDestination `hcl:"storage_blob_dead_letter_destination,block"`
	// StorageQueueEndpoint: optional
	StorageQueueEndpoint *eventgridsystemtopiceventsubscription.StorageQueueEndpoint `hcl:"storage_queue_endpoint,block"`
	// SubjectFilter: optional
	SubjectFilter *eventgridsystemtopiceventsubscription.SubjectFilter `hcl:"subject_filter,block"`
	// Timeouts: optional
	Timeouts *eventgridsystemtopiceventsubscription.Timeouts `hcl:"timeouts,block"`
	// WebhookEndpoint: optional
	WebhookEndpoint *eventgridsystemtopiceventsubscription.WebhookEndpoint `hcl:"webhook_endpoint,block"`
}
type eventgridSystemTopicEventSubscriptionAttributes struct {
	ref terra.Reference
}

// AdvancedFilteringOnArraysEnabled returns a reference to field advanced_filtering_on_arrays_enabled of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) AdvancedFilteringOnArraysEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(estes.ref.Append("advanced_filtering_on_arrays_enabled"))
}

// EventDeliverySchema returns a reference to field event_delivery_schema of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) EventDeliverySchema() terra.StringValue {
	return terra.ReferenceAsString(estes.ref.Append("event_delivery_schema"))
}

// EventhubEndpointId returns a reference to field eventhub_endpoint_id of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) EventhubEndpointId() terra.StringValue {
	return terra.ReferenceAsString(estes.ref.Append("eventhub_endpoint_id"))
}

// ExpirationTimeUtc returns a reference to field expiration_time_utc of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) ExpirationTimeUtc() terra.StringValue {
	return terra.ReferenceAsString(estes.ref.Append("expiration_time_utc"))
}

// HybridConnectionEndpointId returns a reference to field hybrid_connection_endpoint_id of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) HybridConnectionEndpointId() terra.StringValue {
	return terra.ReferenceAsString(estes.ref.Append("hybrid_connection_endpoint_id"))
}

// Id returns a reference to field id of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(estes.ref.Append("id"))
}

// IncludedEventTypes returns a reference to field included_event_types of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) IncludedEventTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](estes.ref.Append("included_event_types"))
}

// Labels returns a reference to field labels of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) Labels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](estes.ref.Append("labels"))
}

// Name returns a reference to field name of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(estes.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(estes.ref.Append("resource_group_name"))
}

// ServiceBusQueueEndpointId returns a reference to field service_bus_queue_endpoint_id of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) ServiceBusQueueEndpointId() terra.StringValue {
	return terra.ReferenceAsString(estes.ref.Append("service_bus_queue_endpoint_id"))
}

// ServiceBusTopicEndpointId returns a reference to field service_bus_topic_endpoint_id of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) ServiceBusTopicEndpointId() terra.StringValue {
	return terra.ReferenceAsString(estes.ref.Append("service_bus_topic_endpoint_id"))
}

// SystemTopic returns a reference to field system_topic of azurerm_eventgrid_system_topic_event_subscription.
func (estes eventgridSystemTopicEventSubscriptionAttributes) SystemTopic() terra.StringValue {
	return terra.ReferenceAsString(estes.ref.Append("system_topic"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) AdvancedFilter() terra.ListValue[eventgridsystemtopiceventsubscription.AdvancedFilterAttributes] {
	return terra.ReferenceAsList[eventgridsystemtopiceventsubscription.AdvancedFilterAttributes](estes.ref.Append("advanced_filter"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) AzureFunctionEndpoint() terra.ListValue[eventgridsystemtopiceventsubscription.AzureFunctionEndpointAttributes] {
	return terra.ReferenceAsList[eventgridsystemtopiceventsubscription.AzureFunctionEndpointAttributes](estes.ref.Append("azure_function_endpoint"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) DeadLetterIdentity() terra.ListValue[eventgridsystemtopiceventsubscription.DeadLetterIdentityAttributes] {
	return terra.ReferenceAsList[eventgridsystemtopiceventsubscription.DeadLetterIdentityAttributes](estes.ref.Append("dead_letter_identity"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) DeliveryIdentity() terra.ListValue[eventgridsystemtopiceventsubscription.DeliveryIdentityAttributes] {
	return terra.ReferenceAsList[eventgridsystemtopiceventsubscription.DeliveryIdentityAttributes](estes.ref.Append("delivery_identity"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) DeliveryProperty() terra.ListValue[eventgridsystemtopiceventsubscription.DeliveryPropertyAttributes] {
	return terra.ReferenceAsList[eventgridsystemtopiceventsubscription.DeliveryPropertyAttributes](estes.ref.Append("delivery_property"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) RetryPolicy() terra.ListValue[eventgridsystemtopiceventsubscription.RetryPolicyAttributes] {
	return terra.ReferenceAsList[eventgridsystemtopiceventsubscription.RetryPolicyAttributes](estes.ref.Append("retry_policy"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) StorageBlobDeadLetterDestination() terra.ListValue[eventgridsystemtopiceventsubscription.StorageBlobDeadLetterDestinationAttributes] {
	return terra.ReferenceAsList[eventgridsystemtopiceventsubscription.StorageBlobDeadLetterDestinationAttributes](estes.ref.Append("storage_blob_dead_letter_destination"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) StorageQueueEndpoint() terra.ListValue[eventgridsystemtopiceventsubscription.StorageQueueEndpointAttributes] {
	return terra.ReferenceAsList[eventgridsystemtopiceventsubscription.StorageQueueEndpointAttributes](estes.ref.Append("storage_queue_endpoint"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) SubjectFilter() terra.ListValue[eventgridsystemtopiceventsubscription.SubjectFilterAttributes] {
	return terra.ReferenceAsList[eventgridsystemtopiceventsubscription.SubjectFilterAttributes](estes.ref.Append("subject_filter"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) Timeouts() eventgridsystemtopiceventsubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventgridsystemtopiceventsubscription.TimeoutsAttributes](estes.ref.Append("timeouts"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) WebhookEndpoint() terra.ListValue[eventgridsystemtopiceventsubscription.WebhookEndpointAttributes] {
	return terra.ReferenceAsList[eventgridsystemtopiceventsubscription.WebhookEndpointAttributes](estes.ref.Append("webhook_endpoint"))
}

type eventgridSystemTopicEventSubscriptionState struct {
	AdvancedFilteringOnArraysEnabled bool                                                                          `json:"advanced_filtering_on_arrays_enabled"`
	EventDeliverySchema              string                                                                        `json:"event_delivery_schema"`
	EventhubEndpointId               string                                                                        `json:"eventhub_endpoint_id"`
	ExpirationTimeUtc                string                                                                        `json:"expiration_time_utc"`
	HybridConnectionEndpointId       string                                                                        `json:"hybrid_connection_endpoint_id"`
	Id                               string                                                                        `json:"id"`
	IncludedEventTypes               []string                                                                      `json:"included_event_types"`
	Labels                           []string                                                                      `json:"labels"`
	Name                             string                                                                        `json:"name"`
	ResourceGroupName                string                                                                        `json:"resource_group_name"`
	ServiceBusQueueEndpointId        string                                                                        `json:"service_bus_queue_endpoint_id"`
	ServiceBusTopicEndpointId        string                                                                        `json:"service_bus_topic_endpoint_id"`
	SystemTopic                      string                                                                        `json:"system_topic"`
	AdvancedFilter                   []eventgridsystemtopiceventsubscription.AdvancedFilterState                   `json:"advanced_filter"`
	AzureFunctionEndpoint            []eventgridsystemtopiceventsubscription.AzureFunctionEndpointState            `json:"azure_function_endpoint"`
	DeadLetterIdentity               []eventgridsystemtopiceventsubscription.DeadLetterIdentityState               `json:"dead_letter_identity"`
	DeliveryIdentity                 []eventgridsystemtopiceventsubscription.DeliveryIdentityState                 `json:"delivery_identity"`
	DeliveryProperty                 []eventgridsystemtopiceventsubscription.DeliveryPropertyState                 `json:"delivery_property"`
	RetryPolicy                      []eventgridsystemtopiceventsubscription.RetryPolicyState                      `json:"retry_policy"`
	StorageBlobDeadLetterDestination []eventgridsystemtopiceventsubscription.StorageBlobDeadLetterDestinationState `json:"storage_blob_dead_letter_destination"`
	StorageQueueEndpoint             []eventgridsystemtopiceventsubscription.StorageQueueEndpointState             `json:"storage_queue_endpoint"`
	SubjectFilter                    []eventgridsystemtopiceventsubscription.SubjectFilterState                    `json:"subject_filter"`
	Timeouts                         *eventgridsystemtopiceventsubscription.TimeoutsState                          `json:"timeouts"`
	WebhookEndpoint                  []eventgridsystemtopiceventsubscription.WebhookEndpointState                  `json:"webhook_endpoint"`
}
