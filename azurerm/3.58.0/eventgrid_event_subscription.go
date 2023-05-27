// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventgrideventsubscription "github.com/golingon/terraproviders/azurerm/3.58.0/eventgrideventsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventgridEventSubscription creates a new instance of [EventgridEventSubscription].
func NewEventgridEventSubscription(name string, args EventgridEventSubscriptionArgs) *EventgridEventSubscription {
	return &EventgridEventSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventgridEventSubscription)(nil)

// EventgridEventSubscription represents the Terraform resource azurerm_eventgrid_event_subscription.
type EventgridEventSubscription struct {
	Name      string
	Args      EventgridEventSubscriptionArgs
	state     *eventgridEventSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventgridEventSubscription].
func (ees *EventgridEventSubscription) Type() string {
	return "azurerm_eventgrid_event_subscription"
}

// LocalName returns the local name for [EventgridEventSubscription].
func (ees *EventgridEventSubscription) LocalName() string {
	return ees.Name
}

// Configuration returns the configuration (args) for [EventgridEventSubscription].
func (ees *EventgridEventSubscription) Configuration() interface{} {
	return ees.Args
}

// DependOn is used for other resources to depend on [EventgridEventSubscription].
func (ees *EventgridEventSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(ees)
}

// Dependencies returns the list of resources [EventgridEventSubscription] depends_on.
func (ees *EventgridEventSubscription) Dependencies() terra.Dependencies {
	return ees.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventgridEventSubscription].
func (ees *EventgridEventSubscription) LifecycleManagement() *terra.Lifecycle {
	return ees.Lifecycle
}

// Attributes returns the attributes for [EventgridEventSubscription].
func (ees *EventgridEventSubscription) Attributes() eventgridEventSubscriptionAttributes {
	return eventgridEventSubscriptionAttributes{ref: terra.ReferenceResource(ees)}
}

// ImportState imports the given attribute values into [EventgridEventSubscription]'s state.
func (ees *EventgridEventSubscription) ImportState(av io.Reader) error {
	ees.state = &eventgridEventSubscriptionState{}
	if err := json.NewDecoder(av).Decode(ees.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ees.Type(), ees.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventgridEventSubscription] has state.
func (ees *EventgridEventSubscription) State() (*eventgridEventSubscriptionState, bool) {
	return ees.state, ees.state != nil
}

// StateMust returns the state for [EventgridEventSubscription]. Panics if the state is nil.
func (ees *EventgridEventSubscription) StateMust() *eventgridEventSubscriptionState {
	if ees.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ees.Type(), ees.LocalName()))
	}
	return ees.state
}

// EventgridEventSubscriptionArgs contains the configurations for azurerm_eventgrid_event_subscription.
type EventgridEventSubscriptionArgs struct {
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
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// ServiceBusQueueEndpointId: string, optional
	ServiceBusQueueEndpointId terra.StringValue `hcl:"service_bus_queue_endpoint_id,attr"`
	// ServiceBusTopicEndpointId: string, optional
	ServiceBusTopicEndpointId terra.StringValue `hcl:"service_bus_topic_endpoint_id,attr"`
	// AdvancedFilter: optional
	AdvancedFilter *eventgrideventsubscription.AdvancedFilter `hcl:"advanced_filter,block"`
	// AzureFunctionEndpoint: optional
	AzureFunctionEndpoint *eventgrideventsubscription.AzureFunctionEndpoint `hcl:"azure_function_endpoint,block"`
	// DeadLetterIdentity: optional
	DeadLetterIdentity *eventgrideventsubscription.DeadLetterIdentity `hcl:"dead_letter_identity,block"`
	// DeliveryIdentity: optional
	DeliveryIdentity *eventgrideventsubscription.DeliveryIdentity `hcl:"delivery_identity,block"`
	// DeliveryProperty: min=0
	DeliveryProperty []eventgrideventsubscription.DeliveryProperty `hcl:"delivery_property,block" validate:"min=0"`
	// RetryPolicy: optional
	RetryPolicy *eventgrideventsubscription.RetryPolicy `hcl:"retry_policy,block"`
	// StorageBlobDeadLetterDestination: optional
	StorageBlobDeadLetterDestination *eventgrideventsubscription.StorageBlobDeadLetterDestination `hcl:"storage_blob_dead_letter_destination,block"`
	// StorageQueueEndpoint: optional
	StorageQueueEndpoint *eventgrideventsubscription.StorageQueueEndpoint `hcl:"storage_queue_endpoint,block"`
	// SubjectFilter: optional
	SubjectFilter *eventgrideventsubscription.SubjectFilter `hcl:"subject_filter,block"`
	// Timeouts: optional
	Timeouts *eventgrideventsubscription.Timeouts `hcl:"timeouts,block"`
	// WebhookEndpoint: optional
	WebhookEndpoint *eventgrideventsubscription.WebhookEndpoint `hcl:"webhook_endpoint,block"`
}
type eventgridEventSubscriptionAttributes struct {
	ref terra.Reference
}

// AdvancedFilteringOnArraysEnabled returns a reference to field advanced_filtering_on_arrays_enabled of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) AdvancedFilteringOnArraysEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ees.ref.Append("advanced_filtering_on_arrays_enabled"))
}

// EventDeliverySchema returns a reference to field event_delivery_schema of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) EventDeliverySchema() terra.StringValue {
	return terra.ReferenceAsString(ees.ref.Append("event_delivery_schema"))
}

// EventhubEndpointId returns a reference to field eventhub_endpoint_id of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) EventhubEndpointId() terra.StringValue {
	return terra.ReferenceAsString(ees.ref.Append("eventhub_endpoint_id"))
}

// ExpirationTimeUtc returns a reference to field expiration_time_utc of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) ExpirationTimeUtc() terra.StringValue {
	return terra.ReferenceAsString(ees.ref.Append("expiration_time_utc"))
}

// HybridConnectionEndpointId returns a reference to field hybrid_connection_endpoint_id of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) HybridConnectionEndpointId() terra.StringValue {
	return terra.ReferenceAsString(ees.ref.Append("hybrid_connection_endpoint_id"))
}

// Id returns a reference to field id of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ees.ref.Append("id"))
}

// IncludedEventTypes returns a reference to field included_event_types of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) IncludedEventTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ees.ref.Append("included_event_types"))
}

// Labels returns a reference to field labels of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) Labels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ees.ref.Append("labels"))
}

// Name returns a reference to field name of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ees.ref.Append("name"))
}

// Scope returns a reference to field scope of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(ees.ref.Append("scope"))
}

// ServiceBusQueueEndpointId returns a reference to field service_bus_queue_endpoint_id of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) ServiceBusQueueEndpointId() terra.StringValue {
	return terra.ReferenceAsString(ees.ref.Append("service_bus_queue_endpoint_id"))
}

// ServiceBusTopicEndpointId returns a reference to field service_bus_topic_endpoint_id of azurerm_eventgrid_event_subscription.
func (ees eventgridEventSubscriptionAttributes) ServiceBusTopicEndpointId() terra.StringValue {
	return terra.ReferenceAsString(ees.ref.Append("service_bus_topic_endpoint_id"))
}

func (ees eventgridEventSubscriptionAttributes) AdvancedFilter() terra.ListValue[eventgrideventsubscription.AdvancedFilterAttributes] {
	return terra.ReferenceAsList[eventgrideventsubscription.AdvancedFilterAttributes](ees.ref.Append("advanced_filter"))
}

func (ees eventgridEventSubscriptionAttributes) AzureFunctionEndpoint() terra.ListValue[eventgrideventsubscription.AzureFunctionEndpointAttributes] {
	return terra.ReferenceAsList[eventgrideventsubscription.AzureFunctionEndpointAttributes](ees.ref.Append("azure_function_endpoint"))
}

func (ees eventgridEventSubscriptionAttributes) DeadLetterIdentity() terra.ListValue[eventgrideventsubscription.DeadLetterIdentityAttributes] {
	return terra.ReferenceAsList[eventgrideventsubscription.DeadLetterIdentityAttributes](ees.ref.Append("dead_letter_identity"))
}

func (ees eventgridEventSubscriptionAttributes) DeliveryIdentity() terra.ListValue[eventgrideventsubscription.DeliveryIdentityAttributes] {
	return terra.ReferenceAsList[eventgrideventsubscription.DeliveryIdentityAttributes](ees.ref.Append("delivery_identity"))
}

func (ees eventgridEventSubscriptionAttributes) DeliveryProperty() terra.ListValue[eventgrideventsubscription.DeliveryPropertyAttributes] {
	return terra.ReferenceAsList[eventgrideventsubscription.DeliveryPropertyAttributes](ees.ref.Append("delivery_property"))
}

func (ees eventgridEventSubscriptionAttributes) RetryPolicy() terra.ListValue[eventgrideventsubscription.RetryPolicyAttributes] {
	return terra.ReferenceAsList[eventgrideventsubscription.RetryPolicyAttributes](ees.ref.Append("retry_policy"))
}

func (ees eventgridEventSubscriptionAttributes) StorageBlobDeadLetterDestination() terra.ListValue[eventgrideventsubscription.StorageBlobDeadLetterDestinationAttributes] {
	return terra.ReferenceAsList[eventgrideventsubscription.StorageBlobDeadLetterDestinationAttributes](ees.ref.Append("storage_blob_dead_letter_destination"))
}

func (ees eventgridEventSubscriptionAttributes) StorageQueueEndpoint() terra.ListValue[eventgrideventsubscription.StorageQueueEndpointAttributes] {
	return terra.ReferenceAsList[eventgrideventsubscription.StorageQueueEndpointAttributes](ees.ref.Append("storage_queue_endpoint"))
}

func (ees eventgridEventSubscriptionAttributes) SubjectFilter() terra.ListValue[eventgrideventsubscription.SubjectFilterAttributes] {
	return terra.ReferenceAsList[eventgrideventsubscription.SubjectFilterAttributes](ees.ref.Append("subject_filter"))
}

func (ees eventgridEventSubscriptionAttributes) Timeouts() eventgrideventsubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventgrideventsubscription.TimeoutsAttributes](ees.ref.Append("timeouts"))
}

func (ees eventgridEventSubscriptionAttributes) WebhookEndpoint() terra.ListValue[eventgrideventsubscription.WebhookEndpointAttributes] {
	return terra.ReferenceAsList[eventgrideventsubscription.WebhookEndpointAttributes](ees.ref.Append("webhook_endpoint"))
}

type eventgridEventSubscriptionState struct {
	AdvancedFilteringOnArraysEnabled bool                                                               `json:"advanced_filtering_on_arrays_enabled"`
	EventDeliverySchema              string                                                             `json:"event_delivery_schema"`
	EventhubEndpointId               string                                                             `json:"eventhub_endpoint_id"`
	ExpirationTimeUtc                string                                                             `json:"expiration_time_utc"`
	HybridConnectionEndpointId       string                                                             `json:"hybrid_connection_endpoint_id"`
	Id                               string                                                             `json:"id"`
	IncludedEventTypes               []string                                                           `json:"included_event_types"`
	Labels                           []string                                                           `json:"labels"`
	Name                             string                                                             `json:"name"`
	Scope                            string                                                             `json:"scope"`
	ServiceBusQueueEndpointId        string                                                             `json:"service_bus_queue_endpoint_id"`
	ServiceBusTopicEndpointId        string                                                             `json:"service_bus_topic_endpoint_id"`
	AdvancedFilter                   []eventgrideventsubscription.AdvancedFilterState                   `json:"advanced_filter"`
	AzureFunctionEndpoint            []eventgrideventsubscription.AzureFunctionEndpointState            `json:"azure_function_endpoint"`
	DeadLetterIdentity               []eventgrideventsubscription.DeadLetterIdentityState               `json:"dead_letter_identity"`
	DeliveryIdentity                 []eventgrideventsubscription.DeliveryIdentityState                 `json:"delivery_identity"`
	DeliveryProperty                 []eventgrideventsubscription.DeliveryPropertyState                 `json:"delivery_property"`
	RetryPolicy                      []eventgrideventsubscription.RetryPolicyState                      `json:"retry_policy"`
	StorageBlobDeadLetterDestination []eventgrideventsubscription.StorageBlobDeadLetterDestinationState `json:"storage_blob_dead_letter_destination"`
	StorageQueueEndpoint             []eventgrideventsubscription.StorageQueueEndpointState             `json:"storage_queue_endpoint"`
	SubjectFilter                    []eventgrideventsubscription.SubjectFilterState                    `json:"subject_filter"`
	Timeouts                         *eventgrideventsubscription.TimeoutsState                          `json:"timeouts"`
	WebhookEndpoint                  []eventgrideventsubscription.WebhookEndpointState                  `json:"webhook_endpoint"`
}
