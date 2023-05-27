// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsstreaminputeventhubv2 "github.com/golingon/terraproviders/azurerm/3.58.0/streamanalyticsstreaminputeventhubv2"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsStreamInputEventhubV2 creates a new instance of [StreamAnalyticsStreamInputEventhubV2].
func NewStreamAnalyticsStreamInputEventhubV2(name string, args StreamAnalyticsStreamInputEventhubV2Args) *StreamAnalyticsStreamInputEventhubV2 {
	return &StreamAnalyticsStreamInputEventhubV2{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsStreamInputEventhubV2)(nil)

// StreamAnalyticsStreamInputEventhubV2 represents the Terraform resource azurerm_stream_analytics_stream_input_eventhub_v2.
type StreamAnalyticsStreamInputEventhubV2 struct {
	Name      string
	Args      StreamAnalyticsStreamInputEventhubV2Args
	state     *streamAnalyticsStreamInputEventhubV2State
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsStreamInputEventhubV2].
func (sasiev *StreamAnalyticsStreamInputEventhubV2) Type() string {
	return "azurerm_stream_analytics_stream_input_eventhub_v2"
}

// LocalName returns the local name for [StreamAnalyticsStreamInputEventhubV2].
func (sasiev *StreamAnalyticsStreamInputEventhubV2) LocalName() string {
	return sasiev.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsStreamInputEventhubV2].
func (sasiev *StreamAnalyticsStreamInputEventhubV2) Configuration() interface{} {
	return sasiev.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsStreamInputEventhubV2].
func (sasiev *StreamAnalyticsStreamInputEventhubV2) DependOn() terra.Reference {
	return terra.ReferenceResource(sasiev)
}

// Dependencies returns the list of resources [StreamAnalyticsStreamInputEventhubV2] depends_on.
func (sasiev *StreamAnalyticsStreamInputEventhubV2) Dependencies() terra.Dependencies {
	return sasiev.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsStreamInputEventhubV2].
func (sasiev *StreamAnalyticsStreamInputEventhubV2) LifecycleManagement() *terra.Lifecycle {
	return sasiev.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsStreamInputEventhubV2].
func (sasiev *StreamAnalyticsStreamInputEventhubV2) Attributes() streamAnalyticsStreamInputEventhubV2Attributes {
	return streamAnalyticsStreamInputEventhubV2Attributes{ref: terra.ReferenceResource(sasiev)}
}

// ImportState imports the given attribute values into [StreamAnalyticsStreamInputEventhubV2]'s state.
func (sasiev *StreamAnalyticsStreamInputEventhubV2) ImportState(av io.Reader) error {
	sasiev.state = &streamAnalyticsStreamInputEventhubV2State{}
	if err := json.NewDecoder(av).Decode(sasiev.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sasiev.Type(), sasiev.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsStreamInputEventhubV2] has state.
func (sasiev *StreamAnalyticsStreamInputEventhubV2) State() (*streamAnalyticsStreamInputEventhubV2State, bool) {
	return sasiev.state, sasiev.state != nil
}

// StateMust returns the state for [StreamAnalyticsStreamInputEventhubV2]. Panics if the state is nil.
func (sasiev *StreamAnalyticsStreamInputEventhubV2) StateMust() *streamAnalyticsStreamInputEventhubV2State {
	if sasiev.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sasiev.Type(), sasiev.LocalName()))
	}
	return sasiev.state
}

// StreamAnalyticsStreamInputEventhubV2Args contains the configurations for azurerm_stream_analytics_stream_input_eventhub_v2.
type StreamAnalyticsStreamInputEventhubV2Args struct {
	// AuthenticationMode: string, optional
	AuthenticationMode terra.StringValue `hcl:"authentication_mode,attr"`
	// EventhubConsumerGroupName: string, optional
	EventhubConsumerGroupName terra.StringValue `hcl:"eventhub_consumer_group_name,attr"`
	// EventhubName: string, required
	EventhubName terra.StringValue `hcl:"eventhub_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PartitionKey: string, optional
	PartitionKey terra.StringValue `hcl:"partition_key,attr"`
	// ServicebusNamespace: string, required
	ServicebusNamespace terra.StringValue `hcl:"servicebus_namespace,attr" validate:"required"`
	// SharedAccessPolicyKey: string, optional
	SharedAccessPolicyKey terra.StringValue `hcl:"shared_access_policy_key,attr"`
	// SharedAccessPolicyName: string, optional
	SharedAccessPolicyName terra.StringValue `hcl:"shared_access_policy_name,attr"`
	// StreamAnalyticsJobId: string, required
	StreamAnalyticsJobId terra.StringValue `hcl:"stream_analytics_job_id,attr" validate:"required"`
	// Serialization: required
	Serialization *streamanalyticsstreaminputeventhubv2.Serialization `hcl:"serialization,block" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsstreaminputeventhubv2.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsStreamInputEventhubV2Attributes struct {
	ref terra.Reference
}

// AuthenticationMode returns a reference to field authentication_mode of azurerm_stream_analytics_stream_input_eventhub_v2.
func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) AuthenticationMode() terra.StringValue {
	return terra.ReferenceAsString(sasiev.ref.Append("authentication_mode"))
}

// EventhubConsumerGroupName returns a reference to field eventhub_consumer_group_name of azurerm_stream_analytics_stream_input_eventhub_v2.
func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) EventhubConsumerGroupName() terra.StringValue {
	return terra.ReferenceAsString(sasiev.ref.Append("eventhub_consumer_group_name"))
}

// EventhubName returns a reference to field eventhub_name of azurerm_stream_analytics_stream_input_eventhub_v2.
func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(sasiev.ref.Append("eventhub_name"))
}

// Id returns a reference to field id of azurerm_stream_analytics_stream_input_eventhub_v2.
func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sasiev.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_stream_input_eventhub_v2.
func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sasiev.ref.Append("name"))
}

// PartitionKey returns a reference to field partition_key of azurerm_stream_analytics_stream_input_eventhub_v2.
func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) PartitionKey() terra.StringValue {
	return terra.ReferenceAsString(sasiev.ref.Append("partition_key"))
}

// ServicebusNamespace returns a reference to field servicebus_namespace of azurerm_stream_analytics_stream_input_eventhub_v2.
func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) ServicebusNamespace() terra.StringValue {
	return terra.ReferenceAsString(sasiev.ref.Append("servicebus_namespace"))
}

// SharedAccessPolicyKey returns a reference to field shared_access_policy_key of azurerm_stream_analytics_stream_input_eventhub_v2.
func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) SharedAccessPolicyKey() terra.StringValue {
	return terra.ReferenceAsString(sasiev.ref.Append("shared_access_policy_key"))
}

// SharedAccessPolicyName returns a reference to field shared_access_policy_name of azurerm_stream_analytics_stream_input_eventhub_v2.
func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) SharedAccessPolicyName() terra.StringValue {
	return terra.ReferenceAsString(sasiev.ref.Append("shared_access_policy_name"))
}

// StreamAnalyticsJobId returns a reference to field stream_analytics_job_id of azurerm_stream_analytics_stream_input_eventhub_v2.
func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) StreamAnalyticsJobId() terra.StringValue {
	return terra.ReferenceAsString(sasiev.ref.Append("stream_analytics_job_id"))
}

func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) Serialization() terra.ListValue[streamanalyticsstreaminputeventhubv2.SerializationAttributes] {
	return terra.ReferenceAsList[streamanalyticsstreaminputeventhubv2.SerializationAttributes](sasiev.ref.Append("serialization"))
}

func (sasiev streamAnalyticsStreamInputEventhubV2Attributes) Timeouts() streamanalyticsstreaminputeventhubv2.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsstreaminputeventhubv2.TimeoutsAttributes](sasiev.ref.Append("timeouts"))
}

type streamAnalyticsStreamInputEventhubV2State struct {
	AuthenticationMode        string                                                    `json:"authentication_mode"`
	EventhubConsumerGroupName string                                                    `json:"eventhub_consumer_group_name"`
	EventhubName              string                                                    `json:"eventhub_name"`
	Id                        string                                                    `json:"id"`
	Name                      string                                                    `json:"name"`
	PartitionKey              string                                                    `json:"partition_key"`
	ServicebusNamespace       string                                                    `json:"servicebus_namespace"`
	SharedAccessPolicyKey     string                                                    `json:"shared_access_policy_key"`
	SharedAccessPolicyName    string                                                    `json:"shared_access_policy_name"`
	StreamAnalyticsJobId      string                                                    `json:"stream_analytics_job_id"`
	Serialization             []streamanalyticsstreaminputeventhubv2.SerializationState `json:"serialization"`
	Timeouts                  *streamanalyticsstreaminputeventhubv2.TimeoutsState       `json:"timeouts"`
}
