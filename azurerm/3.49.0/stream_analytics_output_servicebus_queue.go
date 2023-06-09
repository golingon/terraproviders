// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsoutputservicebusqueue "github.com/golingon/terraproviders/azurerm/3.49.0/streamanalyticsoutputservicebusqueue"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsOutputServicebusQueue creates a new instance of [StreamAnalyticsOutputServicebusQueue].
func NewStreamAnalyticsOutputServicebusQueue(name string, args StreamAnalyticsOutputServicebusQueueArgs) *StreamAnalyticsOutputServicebusQueue {
	return &StreamAnalyticsOutputServicebusQueue{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsOutputServicebusQueue)(nil)

// StreamAnalyticsOutputServicebusQueue represents the Terraform resource azurerm_stream_analytics_output_servicebus_queue.
type StreamAnalyticsOutputServicebusQueue struct {
	Name      string
	Args      StreamAnalyticsOutputServicebusQueueArgs
	state     *streamAnalyticsOutputServicebusQueueState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsOutputServicebusQueue].
func (saosq *StreamAnalyticsOutputServicebusQueue) Type() string {
	return "azurerm_stream_analytics_output_servicebus_queue"
}

// LocalName returns the local name for [StreamAnalyticsOutputServicebusQueue].
func (saosq *StreamAnalyticsOutputServicebusQueue) LocalName() string {
	return saosq.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsOutputServicebusQueue].
func (saosq *StreamAnalyticsOutputServicebusQueue) Configuration() interface{} {
	return saosq.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsOutputServicebusQueue].
func (saosq *StreamAnalyticsOutputServicebusQueue) DependOn() terra.Reference {
	return terra.ReferenceResource(saosq)
}

// Dependencies returns the list of resources [StreamAnalyticsOutputServicebusQueue] depends_on.
func (saosq *StreamAnalyticsOutputServicebusQueue) Dependencies() terra.Dependencies {
	return saosq.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsOutputServicebusQueue].
func (saosq *StreamAnalyticsOutputServicebusQueue) LifecycleManagement() *terra.Lifecycle {
	return saosq.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsOutputServicebusQueue].
func (saosq *StreamAnalyticsOutputServicebusQueue) Attributes() streamAnalyticsOutputServicebusQueueAttributes {
	return streamAnalyticsOutputServicebusQueueAttributes{ref: terra.ReferenceResource(saosq)}
}

// ImportState imports the given attribute values into [StreamAnalyticsOutputServicebusQueue]'s state.
func (saosq *StreamAnalyticsOutputServicebusQueue) ImportState(av io.Reader) error {
	saosq.state = &streamAnalyticsOutputServicebusQueueState{}
	if err := json.NewDecoder(av).Decode(saosq.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saosq.Type(), saosq.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsOutputServicebusQueue] has state.
func (saosq *StreamAnalyticsOutputServicebusQueue) State() (*streamAnalyticsOutputServicebusQueueState, bool) {
	return saosq.state, saosq.state != nil
}

// StateMust returns the state for [StreamAnalyticsOutputServicebusQueue]. Panics if the state is nil.
func (saosq *StreamAnalyticsOutputServicebusQueue) StateMust() *streamAnalyticsOutputServicebusQueueState {
	if saosq.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saosq.Type(), saosq.LocalName()))
	}
	return saosq.state
}

// StreamAnalyticsOutputServicebusQueueArgs contains the configurations for azurerm_stream_analytics_output_servicebus_queue.
type StreamAnalyticsOutputServicebusQueueArgs struct {
	// AuthenticationMode: string, optional
	AuthenticationMode terra.StringValue `hcl:"authentication_mode,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PropertyColumns: list of string, optional
	PropertyColumns terra.ListValue[terra.StringValue] `hcl:"property_columns,attr"`
	// QueueName: string, required
	QueueName terra.StringValue `hcl:"queue_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServicebusNamespace: string, required
	ServicebusNamespace terra.StringValue `hcl:"servicebus_namespace,attr" validate:"required"`
	// SharedAccessPolicyKey: string, optional
	SharedAccessPolicyKey terra.StringValue `hcl:"shared_access_policy_key,attr"`
	// SharedAccessPolicyName: string, optional
	SharedAccessPolicyName terra.StringValue `hcl:"shared_access_policy_name,attr"`
	// StreamAnalyticsJobName: string, required
	StreamAnalyticsJobName terra.StringValue `hcl:"stream_analytics_job_name,attr" validate:"required"`
	// SystemPropertyColumns: map of string, optional
	SystemPropertyColumns terra.MapValue[terra.StringValue] `hcl:"system_property_columns,attr"`
	// Serialization: required
	Serialization *streamanalyticsoutputservicebusqueue.Serialization `hcl:"serialization,block" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsoutputservicebusqueue.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsOutputServicebusQueueAttributes struct {
	ref terra.Reference
}

// AuthenticationMode returns a reference to field authentication_mode of azurerm_stream_analytics_output_servicebus_queue.
func (saosq streamAnalyticsOutputServicebusQueueAttributes) AuthenticationMode() terra.StringValue {
	return terra.ReferenceAsString(saosq.ref.Append("authentication_mode"))
}

// Id returns a reference to field id of azurerm_stream_analytics_output_servicebus_queue.
func (saosq streamAnalyticsOutputServicebusQueueAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saosq.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_output_servicebus_queue.
func (saosq streamAnalyticsOutputServicebusQueueAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(saosq.ref.Append("name"))
}

// PropertyColumns returns a reference to field property_columns of azurerm_stream_analytics_output_servicebus_queue.
func (saosq streamAnalyticsOutputServicebusQueueAttributes) PropertyColumns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](saosq.ref.Append("property_columns"))
}

// QueueName returns a reference to field queue_name of azurerm_stream_analytics_output_servicebus_queue.
func (saosq streamAnalyticsOutputServicebusQueueAttributes) QueueName() terra.StringValue {
	return terra.ReferenceAsString(saosq.ref.Append("queue_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_output_servicebus_queue.
func (saosq streamAnalyticsOutputServicebusQueueAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(saosq.ref.Append("resource_group_name"))
}

// ServicebusNamespace returns a reference to field servicebus_namespace of azurerm_stream_analytics_output_servicebus_queue.
func (saosq streamAnalyticsOutputServicebusQueueAttributes) ServicebusNamespace() terra.StringValue {
	return terra.ReferenceAsString(saosq.ref.Append("servicebus_namespace"))
}

// SharedAccessPolicyKey returns a reference to field shared_access_policy_key of azurerm_stream_analytics_output_servicebus_queue.
func (saosq streamAnalyticsOutputServicebusQueueAttributes) SharedAccessPolicyKey() terra.StringValue {
	return terra.ReferenceAsString(saosq.ref.Append("shared_access_policy_key"))
}

// SharedAccessPolicyName returns a reference to field shared_access_policy_name of azurerm_stream_analytics_output_servicebus_queue.
func (saosq streamAnalyticsOutputServicebusQueueAttributes) SharedAccessPolicyName() terra.StringValue {
	return terra.ReferenceAsString(saosq.ref.Append("shared_access_policy_name"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_output_servicebus_queue.
func (saosq streamAnalyticsOutputServicebusQueueAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(saosq.ref.Append("stream_analytics_job_name"))
}

// SystemPropertyColumns returns a reference to field system_property_columns of azurerm_stream_analytics_output_servicebus_queue.
func (saosq streamAnalyticsOutputServicebusQueueAttributes) SystemPropertyColumns() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](saosq.ref.Append("system_property_columns"))
}

func (saosq streamAnalyticsOutputServicebusQueueAttributes) Serialization() terra.ListValue[streamanalyticsoutputservicebusqueue.SerializationAttributes] {
	return terra.ReferenceAsList[streamanalyticsoutputservicebusqueue.SerializationAttributes](saosq.ref.Append("serialization"))
}

func (saosq streamAnalyticsOutputServicebusQueueAttributes) Timeouts() streamanalyticsoutputservicebusqueue.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsoutputservicebusqueue.TimeoutsAttributes](saosq.ref.Append("timeouts"))
}

type streamAnalyticsOutputServicebusQueueState struct {
	AuthenticationMode     string                                                    `json:"authentication_mode"`
	Id                     string                                                    `json:"id"`
	Name                   string                                                    `json:"name"`
	PropertyColumns        []string                                                  `json:"property_columns"`
	QueueName              string                                                    `json:"queue_name"`
	ResourceGroupName      string                                                    `json:"resource_group_name"`
	ServicebusNamespace    string                                                    `json:"servicebus_namespace"`
	SharedAccessPolicyKey  string                                                    `json:"shared_access_policy_key"`
	SharedAccessPolicyName string                                                    `json:"shared_access_policy_name"`
	StreamAnalyticsJobName string                                                    `json:"stream_analytics_job_name"`
	SystemPropertyColumns  map[string]string                                         `json:"system_property_columns"`
	Serialization          []streamanalyticsoutputservicebusqueue.SerializationState `json:"serialization"`
	Timeouts               *streamanalyticsoutputservicebusqueue.TimeoutsState       `json:"timeouts"`
}
