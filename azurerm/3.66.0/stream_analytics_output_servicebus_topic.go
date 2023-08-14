// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsoutputservicebustopic "github.com/golingon/terraproviders/azurerm/3.66.0/streamanalyticsoutputservicebustopic"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsOutputServicebusTopic creates a new instance of [StreamAnalyticsOutputServicebusTopic].
func NewStreamAnalyticsOutputServicebusTopic(name string, args StreamAnalyticsOutputServicebusTopicArgs) *StreamAnalyticsOutputServicebusTopic {
	return &StreamAnalyticsOutputServicebusTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsOutputServicebusTopic)(nil)

// StreamAnalyticsOutputServicebusTopic represents the Terraform resource azurerm_stream_analytics_output_servicebus_topic.
type StreamAnalyticsOutputServicebusTopic struct {
	Name      string
	Args      StreamAnalyticsOutputServicebusTopicArgs
	state     *streamAnalyticsOutputServicebusTopicState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsOutputServicebusTopic].
func (saost *StreamAnalyticsOutputServicebusTopic) Type() string {
	return "azurerm_stream_analytics_output_servicebus_topic"
}

// LocalName returns the local name for [StreamAnalyticsOutputServicebusTopic].
func (saost *StreamAnalyticsOutputServicebusTopic) LocalName() string {
	return saost.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsOutputServicebusTopic].
func (saost *StreamAnalyticsOutputServicebusTopic) Configuration() interface{} {
	return saost.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsOutputServicebusTopic].
func (saost *StreamAnalyticsOutputServicebusTopic) DependOn() terra.Reference {
	return terra.ReferenceResource(saost)
}

// Dependencies returns the list of resources [StreamAnalyticsOutputServicebusTopic] depends_on.
func (saost *StreamAnalyticsOutputServicebusTopic) Dependencies() terra.Dependencies {
	return saost.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsOutputServicebusTopic].
func (saost *StreamAnalyticsOutputServicebusTopic) LifecycleManagement() *terra.Lifecycle {
	return saost.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsOutputServicebusTopic].
func (saost *StreamAnalyticsOutputServicebusTopic) Attributes() streamAnalyticsOutputServicebusTopicAttributes {
	return streamAnalyticsOutputServicebusTopicAttributes{ref: terra.ReferenceResource(saost)}
}

// ImportState imports the given attribute values into [StreamAnalyticsOutputServicebusTopic]'s state.
func (saost *StreamAnalyticsOutputServicebusTopic) ImportState(av io.Reader) error {
	saost.state = &streamAnalyticsOutputServicebusTopicState{}
	if err := json.NewDecoder(av).Decode(saost.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saost.Type(), saost.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsOutputServicebusTopic] has state.
func (saost *StreamAnalyticsOutputServicebusTopic) State() (*streamAnalyticsOutputServicebusTopicState, bool) {
	return saost.state, saost.state != nil
}

// StateMust returns the state for [StreamAnalyticsOutputServicebusTopic]. Panics if the state is nil.
func (saost *StreamAnalyticsOutputServicebusTopic) StateMust() *streamAnalyticsOutputServicebusTopicState {
	if saost.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saost.Type(), saost.LocalName()))
	}
	return saost.state
}

// StreamAnalyticsOutputServicebusTopicArgs contains the configurations for azurerm_stream_analytics_output_servicebus_topic.
type StreamAnalyticsOutputServicebusTopicArgs struct {
	// AuthenticationMode: string, optional
	AuthenticationMode terra.StringValue `hcl:"authentication_mode,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PropertyColumns: list of string, optional
	PropertyColumns terra.ListValue[terra.StringValue] `hcl:"property_columns,attr"`
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
	// TopicName: string, required
	TopicName terra.StringValue `hcl:"topic_name,attr" validate:"required"`
	// Serialization: required
	Serialization *streamanalyticsoutputservicebustopic.Serialization `hcl:"serialization,block" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsoutputservicebustopic.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsOutputServicebusTopicAttributes struct {
	ref terra.Reference
}

// AuthenticationMode returns a reference to field authentication_mode of azurerm_stream_analytics_output_servicebus_topic.
func (saost streamAnalyticsOutputServicebusTopicAttributes) AuthenticationMode() terra.StringValue {
	return terra.ReferenceAsString(saost.ref.Append("authentication_mode"))
}

// Id returns a reference to field id of azurerm_stream_analytics_output_servicebus_topic.
func (saost streamAnalyticsOutputServicebusTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saost.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_output_servicebus_topic.
func (saost streamAnalyticsOutputServicebusTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(saost.ref.Append("name"))
}

// PropertyColumns returns a reference to field property_columns of azurerm_stream_analytics_output_servicebus_topic.
func (saost streamAnalyticsOutputServicebusTopicAttributes) PropertyColumns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](saost.ref.Append("property_columns"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_output_servicebus_topic.
func (saost streamAnalyticsOutputServicebusTopicAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(saost.ref.Append("resource_group_name"))
}

// ServicebusNamespace returns a reference to field servicebus_namespace of azurerm_stream_analytics_output_servicebus_topic.
func (saost streamAnalyticsOutputServicebusTopicAttributes) ServicebusNamespace() terra.StringValue {
	return terra.ReferenceAsString(saost.ref.Append("servicebus_namespace"))
}

// SharedAccessPolicyKey returns a reference to field shared_access_policy_key of azurerm_stream_analytics_output_servicebus_topic.
func (saost streamAnalyticsOutputServicebusTopicAttributes) SharedAccessPolicyKey() terra.StringValue {
	return terra.ReferenceAsString(saost.ref.Append("shared_access_policy_key"))
}

// SharedAccessPolicyName returns a reference to field shared_access_policy_name of azurerm_stream_analytics_output_servicebus_topic.
func (saost streamAnalyticsOutputServicebusTopicAttributes) SharedAccessPolicyName() terra.StringValue {
	return terra.ReferenceAsString(saost.ref.Append("shared_access_policy_name"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_output_servicebus_topic.
func (saost streamAnalyticsOutputServicebusTopicAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(saost.ref.Append("stream_analytics_job_name"))
}

// SystemPropertyColumns returns a reference to field system_property_columns of azurerm_stream_analytics_output_servicebus_topic.
func (saost streamAnalyticsOutputServicebusTopicAttributes) SystemPropertyColumns() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](saost.ref.Append("system_property_columns"))
}

// TopicName returns a reference to field topic_name of azurerm_stream_analytics_output_servicebus_topic.
func (saost streamAnalyticsOutputServicebusTopicAttributes) TopicName() terra.StringValue {
	return terra.ReferenceAsString(saost.ref.Append("topic_name"))
}

func (saost streamAnalyticsOutputServicebusTopicAttributes) Serialization() terra.ListValue[streamanalyticsoutputservicebustopic.SerializationAttributes] {
	return terra.ReferenceAsList[streamanalyticsoutputservicebustopic.SerializationAttributes](saost.ref.Append("serialization"))
}

func (saost streamAnalyticsOutputServicebusTopicAttributes) Timeouts() streamanalyticsoutputservicebustopic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsoutputservicebustopic.TimeoutsAttributes](saost.ref.Append("timeouts"))
}

type streamAnalyticsOutputServicebusTopicState struct {
	AuthenticationMode     string                                                    `json:"authentication_mode"`
	Id                     string                                                    `json:"id"`
	Name                   string                                                    `json:"name"`
	PropertyColumns        []string                                                  `json:"property_columns"`
	ResourceGroupName      string                                                    `json:"resource_group_name"`
	ServicebusNamespace    string                                                    `json:"servicebus_namespace"`
	SharedAccessPolicyKey  string                                                    `json:"shared_access_policy_key"`
	SharedAccessPolicyName string                                                    `json:"shared_access_policy_name"`
	StreamAnalyticsJobName string                                                    `json:"stream_analytics_job_name"`
	SystemPropertyColumns  map[string]string                                         `json:"system_property_columns"`
	TopicName              string                                                    `json:"topic_name"`
	Serialization          []streamanalyticsoutputservicebustopic.SerializationState `json:"serialization"`
	Timeouts               *streamanalyticsoutputservicebustopic.TimeoutsState       `json:"timeouts"`
}