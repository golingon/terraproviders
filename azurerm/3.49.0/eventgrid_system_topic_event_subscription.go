// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventgridsystemtopiceventsubscription "github.com/golingon/terraproviders/azurerm/3.49.0/eventgridsystemtopiceventsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEventgridSystemTopicEventSubscription(name string, args EventgridSystemTopicEventSubscriptionArgs) *EventgridSystemTopicEventSubscription {
	return &EventgridSystemTopicEventSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventgridSystemTopicEventSubscription)(nil)

type EventgridSystemTopicEventSubscription struct {
	Name  string
	Args  EventgridSystemTopicEventSubscriptionArgs
	state *eventgridSystemTopicEventSubscriptionState
}

func (estes *EventgridSystemTopicEventSubscription) Type() string {
	return "azurerm_eventgrid_system_topic_event_subscription"
}

func (estes *EventgridSystemTopicEventSubscription) LocalName() string {
	return estes.Name
}

func (estes *EventgridSystemTopicEventSubscription) Configuration() interface{} {
	return estes.Args
}

func (estes *EventgridSystemTopicEventSubscription) Attributes() eventgridSystemTopicEventSubscriptionAttributes {
	return eventgridSystemTopicEventSubscriptionAttributes{ref: terra.ReferenceResource(estes)}
}

func (estes *EventgridSystemTopicEventSubscription) ImportState(av io.Reader) error {
	estes.state = &eventgridSystemTopicEventSubscriptionState{}
	if err := json.NewDecoder(av).Decode(estes.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", estes.Type(), estes.LocalName(), err)
	}
	return nil
}

func (estes *EventgridSystemTopicEventSubscription) State() (*eventgridSystemTopicEventSubscriptionState, bool) {
	return estes.state, estes.state != nil
}

func (estes *EventgridSystemTopicEventSubscription) StateMust() *eventgridSystemTopicEventSubscriptionState {
	if estes.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", estes.Type(), estes.LocalName()))
	}
	return estes.state
}

func (estes *EventgridSystemTopicEventSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(estes)
}

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
	// DependsOn contains resources that EventgridSystemTopicEventSubscription depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type eventgridSystemTopicEventSubscriptionAttributes struct {
	ref terra.Reference
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) AdvancedFilteringOnArraysEnabled() terra.BoolValue {
	return terra.ReferenceBool(estes.ref.Append("advanced_filtering_on_arrays_enabled"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) EventDeliverySchema() terra.StringValue {
	return terra.ReferenceString(estes.ref.Append("event_delivery_schema"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) EventhubEndpointId() terra.StringValue {
	return terra.ReferenceString(estes.ref.Append("eventhub_endpoint_id"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) ExpirationTimeUtc() terra.StringValue {
	return terra.ReferenceString(estes.ref.Append("expiration_time_utc"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) HybridConnectionEndpointId() terra.StringValue {
	return terra.ReferenceString(estes.ref.Append("hybrid_connection_endpoint_id"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(estes.ref.Append("id"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) IncludedEventTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](estes.ref.Append("included_event_types"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) Labels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](estes.ref.Append("labels"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(estes.ref.Append("name"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(estes.ref.Append("resource_group_name"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) ServiceBusQueueEndpointId() terra.StringValue {
	return terra.ReferenceString(estes.ref.Append("service_bus_queue_endpoint_id"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) ServiceBusTopicEndpointId() terra.StringValue {
	return terra.ReferenceString(estes.ref.Append("service_bus_topic_endpoint_id"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) SystemTopic() terra.StringValue {
	return terra.ReferenceString(estes.ref.Append("system_topic"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) AdvancedFilter() terra.ListValue[eventgridsystemtopiceventsubscription.AdvancedFilterAttributes] {
	return terra.ReferenceList[eventgridsystemtopiceventsubscription.AdvancedFilterAttributes](estes.ref.Append("advanced_filter"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) AzureFunctionEndpoint() terra.ListValue[eventgridsystemtopiceventsubscription.AzureFunctionEndpointAttributes] {
	return terra.ReferenceList[eventgridsystemtopiceventsubscription.AzureFunctionEndpointAttributes](estes.ref.Append("azure_function_endpoint"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) DeadLetterIdentity() terra.ListValue[eventgridsystemtopiceventsubscription.DeadLetterIdentityAttributes] {
	return terra.ReferenceList[eventgridsystemtopiceventsubscription.DeadLetterIdentityAttributes](estes.ref.Append("dead_letter_identity"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) DeliveryIdentity() terra.ListValue[eventgridsystemtopiceventsubscription.DeliveryIdentityAttributes] {
	return terra.ReferenceList[eventgridsystemtopiceventsubscription.DeliveryIdentityAttributes](estes.ref.Append("delivery_identity"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) DeliveryProperty() terra.ListValue[eventgridsystemtopiceventsubscription.DeliveryPropertyAttributes] {
	return terra.ReferenceList[eventgridsystemtopiceventsubscription.DeliveryPropertyAttributes](estes.ref.Append("delivery_property"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) RetryPolicy() terra.ListValue[eventgridsystemtopiceventsubscription.RetryPolicyAttributes] {
	return terra.ReferenceList[eventgridsystemtopiceventsubscription.RetryPolicyAttributes](estes.ref.Append("retry_policy"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) StorageBlobDeadLetterDestination() terra.ListValue[eventgridsystemtopiceventsubscription.StorageBlobDeadLetterDestinationAttributes] {
	return terra.ReferenceList[eventgridsystemtopiceventsubscription.StorageBlobDeadLetterDestinationAttributes](estes.ref.Append("storage_blob_dead_letter_destination"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) StorageQueueEndpoint() terra.ListValue[eventgridsystemtopiceventsubscription.StorageQueueEndpointAttributes] {
	return terra.ReferenceList[eventgridsystemtopiceventsubscription.StorageQueueEndpointAttributes](estes.ref.Append("storage_queue_endpoint"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) SubjectFilter() terra.ListValue[eventgridsystemtopiceventsubscription.SubjectFilterAttributes] {
	return terra.ReferenceList[eventgridsystemtopiceventsubscription.SubjectFilterAttributes](estes.ref.Append("subject_filter"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) Timeouts() eventgridsystemtopiceventsubscription.TimeoutsAttributes {
	return terra.ReferenceSingle[eventgridsystemtopiceventsubscription.TimeoutsAttributes](estes.ref.Append("timeouts"))
}

func (estes eventgridSystemTopicEventSubscriptionAttributes) WebhookEndpoint() terra.ListValue[eventgridsystemtopiceventsubscription.WebhookEndpointAttributes] {
	return terra.ReferenceList[eventgridsystemtopiceventsubscription.WebhookEndpointAttributes](estes.ref.Append("webhook_endpoint"))
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
