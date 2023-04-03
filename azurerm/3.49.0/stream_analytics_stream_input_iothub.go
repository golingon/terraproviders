// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsstreaminputiothub "github.com/golingon/terraproviders/azurerm/3.49.0/streamanalyticsstreaminputiothub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsStreamInputIothub creates a new instance of [StreamAnalyticsStreamInputIothub].
func NewStreamAnalyticsStreamInputIothub(name string, args StreamAnalyticsStreamInputIothubArgs) *StreamAnalyticsStreamInputIothub {
	return &StreamAnalyticsStreamInputIothub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsStreamInputIothub)(nil)

// StreamAnalyticsStreamInputIothub represents the Terraform resource azurerm_stream_analytics_stream_input_iothub.
type StreamAnalyticsStreamInputIothub struct {
	Name      string
	Args      StreamAnalyticsStreamInputIothubArgs
	state     *streamAnalyticsStreamInputIothubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsStreamInputIothub].
func (sasii *StreamAnalyticsStreamInputIothub) Type() string {
	return "azurerm_stream_analytics_stream_input_iothub"
}

// LocalName returns the local name for [StreamAnalyticsStreamInputIothub].
func (sasii *StreamAnalyticsStreamInputIothub) LocalName() string {
	return sasii.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsStreamInputIothub].
func (sasii *StreamAnalyticsStreamInputIothub) Configuration() interface{} {
	return sasii.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsStreamInputIothub].
func (sasii *StreamAnalyticsStreamInputIothub) DependOn() terra.Reference {
	return terra.ReferenceResource(sasii)
}

// Dependencies returns the list of resources [StreamAnalyticsStreamInputIothub] depends_on.
func (sasii *StreamAnalyticsStreamInputIothub) Dependencies() terra.Dependencies {
	return sasii.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsStreamInputIothub].
func (sasii *StreamAnalyticsStreamInputIothub) LifecycleManagement() *terra.Lifecycle {
	return sasii.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsStreamInputIothub].
func (sasii *StreamAnalyticsStreamInputIothub) Attributes() streamAnalyticsStreamInputIothubAttributes {
	return streamAnalyticsStreamInputIothubAttributes{ref: terra.ReferenceResource(sasii)}
}

// ImportState imports the given attribute values into [StreamAnalyticsStreamInputIothub]'s state.
func (sasii *StreamAnalyticsStreamInputIothub) ImportState(av io.Reader) error {
	sasii.state = &streamAnalyticsStreamInputIothubState{}
	if err := json.NewDecoder(av).Decode(sasii.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sasii.Type(), sasii.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsStreamInputIothub] has state.
func (sasii *StreamAnalyticsStreamInputIothub) State() (*streamAnalyticsStreamInputIothubState, bool) {
	return sasii.state, sasii.state != nil
}

// StateMust returns the state for [StreamAnalyticsStreamInputIothub]. Panics if the state is nil.
func (sasii *StreamAnalyticsStreamInputIothub) StateMust() *streamAnalyticsStreamInputIothubState {
	if sasii.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sasii.Type(), sasii.LocalName()))
	}
	return sasii.state
}

// StreamAnalyticsStreamInputIothubArgs contains the configurations for azurerm_stream_analytics_stream_input_iothub.
type StreamAnalyticsStreamInputIothubArgs struct {
	// Endpoint: string, required
	Endpoint terra.StringValue `hcl:"endpoint,attr" validate:"required"`
	// EventhubConsumerGroupName: string, required
	EventhubConsumerGroupName terra.StringValue `hcl:"eventhub_consumer_group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubNamespace: string, required
	IothubNamespace terra.StringValue `hcl:"iothub_namespace,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SharedAccessPolicyKey: string, required
	SharedAccessPolicyKey terra.StringValue `hcl:"shared_access_policy_key,attr" validate:"required"`
	// SharedAccessPolicyName: string, required
	SharedAccessPolicyName terra.StringValue `hcl:"shared_access_policy_name,attr" validate:"required"`
	// StreamAnalyticsJobName: string, required
	StreamAnalyticsJobName terra.StringValue `hcl:"stream_analytics_job_name,attr" validate:"required"`
	// Serialization: required
	Serialization *streamanalyticsstreaminputiothub.Serialization `hcl:"serialization,block" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsstreaminputiothub.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsStreamInputIothubAttributes struct {
	ref terra.Reference
}

// Endpoint returns a reference to field endpoint of azurerm_stream_analytics_stream_input_iothub.
func (sasii streamAnalyticsStreamInputIothubAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(sasii.ref.Append("endpoint"))
}

// EventhubConsumerGroupName returns a reference to field eventhub_consumer_group_name of azurerm_stream_analytics_stream_input_iothub.
func (sasii streamAnalyticsStreamInputIothubAttributes) EventhubConsumerGroupName() terra.StringValue {
	return terra.ReferenceAsString(sasii.ref.Append("eventhub_consumer_group_name"))
}

// Id returns a reference to field id of azurerm_stream_analytics_stream_input_iothub.
func (sasii streamAnalyticsStreamInputIothubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sasii.ref.Append("id"))
}

// IothubNamespace returns a reference to field iothub_namespace of azurerm_stream_analytics_stream_input_iothub.
func (sasii streamAnalyticsStreamInputIothubAttributes) IothubNamespace() terra.StringValue {
	return terra.ReferenceAsString(sasii.ref.Append("iothub_namespace"))
}

// Name returns a reference to field name of azurerm_stream_analytics_stream_input_iothub.
func (sasii streamAnalyticsStreamInputIothubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sasii.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_stream_input_iothub.
func (sasii streamAnalyticsStreamInputIothubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sasii.ref.Append("resource_group_name"))
}

// SharedAccessPolicyKey returns a reference to field shared_access_policy_key of azurerm_stream_analytics_stream_input_iothub.
func (sasii streamAnalyticsStreamInputIothubAttributes) SharedAccessPolicyKey() terra.StringValue {
	return terra.ReferenceAsString(sasii.ref.Append("shared_access_policy_key"))
}

// SharedAccessPolicyName returns a reference to field shared_access_policy_name of azurerm_stream_analytics_stream_input_iothub.
func (sasii streamAnalyticsStreamInputIothubAttributes) SharedAccessPolicyName() terra.StringValue {
	return terra.ReferenceAsString(sasii.ref.Append("shared_access_policy_name"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_stream_input_iothub.
func (sasii streamAnalyticsStreamInputIothubAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(sasii.ref.Append("stream_analytics_job_name"))
}

func (sasii streamAnalyticsStreamInputIothubAttributes) Serialization() terra.ListValue[streamanalyticsstreaminputiothub.SerializationAttributes] {
	return terra.ReferenceAsList[streamanalyticsstreaminputiothub.SerializationAttributes](sasii.ref.Append("serialization"))
}

func (sasii streamAnalyticsStreamInputIothubAttributes) Timeouts() streamanalyticsstreaminputiothub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsstreaminputiothub.TimeoutsAttributes](sasii.ref.Append("timeouts"))
}

type streamAnalyticsStreamInputIothubState struct {
	Endpoint                  string                                                `json:"endpoint"`
	EventhubConsumerGroupName string                                                `json:"eventhub_consumer_group_name"`
	Id                        string                                                `json:"id"`
	IothubNamespace           string                                                `json:"iothub_namespace"`
	Name                      string                                                `json:"name"`
	ResourceGroupName         string                                                `json:"resource_group_name"`
	SharedAccessPolicyKey     string                                                `json:"shared_access_policy_key"`
	SharedAccessPolicyName    string                                                `json:"shared_access_policy_name"`
	StreamAnalyticsJobName    string                                                `json:"stream_analytics_job_name"`
	Serialization             []streamanalyticsstreaminputiothub.SerializationState `json:"serialization"`
	Timeouts                  *streamanalyticsstreaminputiothub.TimeoutsState       `json:"timeouts"`
}
