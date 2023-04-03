// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsstreaminputeventhub "github.com/golingon/terraproviders/azurerm/3.49.0/streamanalyticsstreaminputeventhub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsStreamInputEventhub creates a new instance of [StreamAnalyticsStreamInputEventhub].
func NewStreamAnalyticsStreamInputEventhub(name string, args StreamAnalyticsStreamInputEventhubArgs) *StreamAnalyticsStreamInputEventhub {
	return &StreamAnalyticsStreamInputEventhub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsStreamInputEventhub)(nil)

// StreamAnalyticsStreamInputEventhub represents the Terraform resource azurerm_stream_analytics_stream_input_eventhub.
type StreamAnalyticsStreamInputEventhub struct {
	Name      string
	Args      StreamAnalyticsStreamInputEventhubArgs
	state     *streamAnalyticsStreamInputEventhubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsStreamInputEventhub].
func (sasie *StreamAnalyticsStreamInputEventhub) Type() string {
	return "azurerm_stream_analytics_stream_input_eventhub"
}

// LocalName returns the local name for [StreamAnalyticsStreamInputEventhub].
func (sasie *StreamAnalyticsStreamInputEventhub) LocalName() string {
	return sasie.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsStreamInputEventhub].
func (sasie *StreamAnalyticsStreamInputEventhub) Configuration() interface{} {
	return sasie.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsStreamInputEventhub].
func (sasie *StreamAnalyticsStreamInputEventhub) DependOn() terra.Reference {
	return terra.ReferenceResource(sasie)
}

// Dependencies returns the list of resources [StreamAnalyticsStreamInputEventhub] depends_on.
func (sasie *StreamAnalyticsStreamInputEventhub) Dependencies() terra.Dependencies {
	return sasie.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsStreamInputEventhub].
func (sasie *StreamAnalyticsStreamInputEventhub) LifecycleManagement() *terra.Lifecycle {
	return sasie.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsStreamInputEventhub].
func (sasie *StreamAnalyticsStreamInputEventhub) Attributes() streamAnalyticsStreamInputEventhubAttributes {
	return streamAnalyticsStreamInputEventhubAttributes{ref: terra.ReferenceResource(sasie)}
}

// ImportState imports the given attribute values into [StreamAnalyticsStreamInputEventhub]'s state.
func (sasie *StreamAnalyticsStreamInputEventhub) ImportState(av io.Reader) error {
	sasie.state = &streamAnalyticsStreamInputEventhubState{}
	if err := json.NewDecoder(av).Decode(sasie.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sasie.Type(), sasie.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsStreamInputEventhub] has state.
func (sasie *StreamAnalyticsStreamInputEventhub) State() (*streamAnalyticsStreamInputEventhubState, bool) {
	return sasie.state, sasie.state != nil
}

// StateMust returns the state for [StreamAnalyticsStreamInputEventhub]. Panics if the state is nil.
func (sasie *StreamAnalyticsStreamInputEventhub) StateMust() *streamAnalyticsStreamInputEventhubState {
	if sasie.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sasie.Type(), sasie.LocalName()))
	}
	return sasie.state
}

// StreamAnalyticsStreamInputEventhubArgs contains the configurations for azurerm_stream_analytics_stream_input_eventhub.
type StreamAnalyticsStreamInputEventhubArgs struct {
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
	// Serialization: required
	Serialization *streamanalyticsstreaminputeventhub.Serialization `hcl:"serialization,block" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsstreaminputeventhub.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsStreamInputEventhubAttributes struct {
	ref terra.Reference
}

// AuthenticationMode returns a reference to field authentication_mode of azurerm_stream_analytics_stream_input_eventhub.
func (sasie streamAnalyticsStreamInputEventhubAttributes) AuthenticationMode() terra.StringValue {
	return terra.ReferenceAsString(sasie.ref.Append("authentication_mode"))
}

// EventhubConsumerGroupName returns a reference to field eventhub_consumer_group_name of azurerm_stream_analytics_stream_input_eventhub.
func (sasie streamAnalyticsStreamInputEventhubAttributes) EventhubConsumerGroupName() terra.StringValue {
	return terra.ReferenceAsString(sasie.ref.Append("eventhub_consumer_group_name"))
}

// EventhubName returns a reference to field eventhub_name of azurerm_stream_analytics_stream_input_eventhub.
func (sasie streamAnalyticsStreamInputEventhubAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(sasie.ref.Append("eventhub_name"))
}

// Id returns a reference to field id of azurerm_stream_analytics_stream_input_eventhub.
func (sasie streamAnalyticsStreamInputEventhubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sasie.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_stream_input_eventhub.
func (sasie streamAnalyticsStreamInputEventhubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sasie.ref.Append("name"))
}

// PartitionKey returns a reference to field partition_key of azurerm_stream_analytics_stream_input_eventhub.
func (sasie streamAnalyticsStreamInputEventhubAttributes) PartitionKey() terra.StringValue {
	return terra.ReferenceAsString(sasie.ref.Append("partition_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_stream_input_eventhub.
func (sasie streamAnalyticsStreamInputEventhubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sasie.ref.Append("resource_group_name"))
}

// ServicebusNamespace returns a reference to field servicebus_namespace of azurerm_stream_analytics_stream_input_eventhub.
func (sasie streamAnalyticsStreamInputEventhubAttributes) ServicebusNamespace() terra.StringValue {
	return terra.ReferenceAsString(sasie.ref.Append("servicebus_namespace"))
}

// SharedAccessPolicyKey returns a reference to field shared_access_policy_key of azurerm_stream_analytics_stream_input_eventhub.
func (sasie streamAnalyticsStreamInputEventhubAttributes) SharedAccessPolicyKey() terra.StringValue {
	return terra.ReferenceAsString(sasie.ref.Append("shared_access_policy_key"))
}

// SharedAccessPolicyName returns a reference to field shared_access_policy_name of azurerm_stream_analytics_stream_input_eventhub.
func (sasie streamAnalyticsStreamInputEventhubAttributes) SharedAccessPolicyName() terra.StringValue {
	return terra.ReferenceAsString(sasie.ref.Append("shared_access_policy_name"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_stream_input_eventhub.
func (sasie streamAnalyticsStreamInputEventhubAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(sasie.ref.Append("stream_analytics_job_name"))
}

func (sasie streamAnalyticsStreamInputEventhubAttributes) Serialization() terra.ListValue[streamanalyticsstreaminputeventhub.SerializationAttributes] {
	return terra.ReferenceAsList[streamanalyticsstreaminputeventhub.SerializationAttributes](sasie.ref.Append("serialization"))
}

func (sasie streamAnalyticsStreamInputEventhubAttributes) Timeouts() streamanalyticsstreaminputeventhub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsstreaminputeventhub.TimeoutsAttributes](sasie.ref.Append("timeouts"))
}

type streamAnalyticsStreamInputEventhubState struct {
	AuthenticationMode        string                                                  `json:"authentication_mode"`
	EventhubConsumerGroupName string                                                  `json:"eventhub_consumer_group_name"`
	EventhubName              string                                                  `json:"eventhub_name"`
	Id                        string                                                  `json:"id"`
	Name                      string                                                  `json:"name"`
	PartitionKey              string                                                  `json:"partition_key"`
	ResourceGroupName         string                                                  `json:"resource_group_name"`
	ServicebusNamespace       string                                                  `json:"servicebus_namespace"`
	SharedAccessPolicyKey     string                                                  `json:"shared_access_policy_key"`
	SharedAccessPolicyName    string                                                  `json:"shared_access_policy_name"`
	StreamAnalyticsJobName    string                                                  `json:"stream_analytics_job_name"`
	Serialization             []streamanalyticsstreaminputeventhub.SerializationState `json:"serialization"`
	Timeouts                  *streamanalyticsstreaminputeventhub.TimeoutsState       `json:"timeouts"`
}
