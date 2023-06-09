// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsoutputeventhub "github.com/golingon/terraproviders/azurerm/3.55.0/streamanalyticsoutputeventhub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsOutputEventhub creates a new instance of [StreamAnalyticsOutputEventhub].
func NewStreamAnalyticsOutputEventhub(name string, args StreamAnalyticsOutputEventhubArgs) *StreamAnalyticsOutputEventhub {
	return &StreamAnalyticsOutputEventhub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsOutputEventhub)(nil)

// StreamAnalyticsOutputEventhub represents the Terraform resource azurerm_stream_analytics_output_eventhub.
type StreamAnalyticsOutputEventhub struct {
	Name      string
	Args      StreamAnalyticsOutputEventhubArgs
	state     *streamAnalyticsOutputEventhubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsOutputEventhub].
func (saoe *StreamAnalyticsOutputEventhub) Type() string {
	return "azurerm_stream_analytics_output_eventhub"
}

// LocalName returns the local name for [StreamAnalyticsOutputEventhub].
func (saoe *StreamAnalyticsOutputEventhub) LocalName() string {
	return saoe.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsOutputEventhub].
func (saoe *StreamAnalyticsOutputEventhub) Configuration() interface{} {
	return saoe.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsOutputEventhub].
func (saoe *StreamAnalyticsOutputEventhub) DependOn() terra.Reference {
	return terra.ReferenceResource(saoe)
}

// Dependencies returns the list of resources [StreamAnalyticsOutputEventhub] depends_on.
func (saoe *StreamAnalyticsOutputEventhub) Dependencies() terra.Dependencies {
	return saoe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsOutputEventhub].
func (saoe *StreamAnalyticsOutputEventhub) LifecycleManagement() *terra.Lifecycle {
	return saoe.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsOutputEventhub].
func (saoe *StreamAnalyticsOutputEventhub) Attributes() streamAnalyticsOutputEventhubAttributes {
	return streamAnalyticsOutputEventhubAttributes{ref: terra.ReferenceResource(saoe)}
}

// ImportState imports the given attribute values into [StreamAnalyticsOutputEventhub]'s state.
func (saoe *StreamAnalyticsOutputEventhub) ImportState(av io.Reader) error {
	saoe.state = &streamAnalyticsOutputEventhubState{}
	if err := json.NewDecoder(av).Decode(saoe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saoe.Type(), saoe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsOutputEventhub] has state.
func (saoe *StreamAnalyticsOutputEventhub) State() (*streamAnalyticsOutputEventhubState, bool) {
	return saoe.state, saoe.state != nil
}

// StateMust returns the state for [StreamAnalyticsOutputEventhub]. Panics if the state is nil.
func (saoe *StreamAnalyticsOutputEventhub) StateMust() *streamAnalyticsOutputEventhubState {
	if saoe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saoe.Type(), saoe.LocalName()))
	}
	return saoe.state
}

// StreamAnalyticsOutputEventhubArgs contains the configurations for azurerm_stream_analytics_output_eventhub.
type StreamAnalyticsOutputEventhubArgs struct {
	// AuthenticationMode: string, optional
	AuthenticationMode terra.StringValue `hcl:"authentication_mode,attr"`
	// EventhubName: string, required
	EventhubName terra.StringValue `hcl:"eventhub_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PartitionKey: string, optional
	PartitionKey terra.StringValue `hcl:"partition_key,attr"`
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
	// Serialization: required
	Serialization *streamanalyticsoutputeventhub.Serialization `hcl:"serialization,block" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsoutputeventhub.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsOutputEventhubAttributes struct {
	ref terra.Reference
}

// AuthenticationMode returns a reference to field authentication_mode of azurerm_stream_analytics_output_eventhub.
func (saoe streamAnalyticsOutputEventhubAttributes) AuthenticationMode() terra.StringValue {
	return terra.ReferenceAsString(saoe.ref.Append("authentication_mode"))
}

// EventhubName returns a reference to field eventhub_name of azurerm_stream_analytics_output_eventhub.
func (saoe streamAnalyticsOutputEventhubAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(saoe.ref.Append("eventhub_name"))
}

// Id returns a reference to field id of azurerm_stream_analytics_output_eventhub.
func (saoe streamAnalyticsOutputEventhubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saoe.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_output_eventhub.
func (saoe streamAnalyticsOutputEventhubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(saoe.ref.Append("name"))
}

// PartitionKey returns a reference to field partition_key of azurerm_stream_analytics_output_eventhub.
func (saoe streamAnalyticsOutputEventhubAttributes) PartitionKey() terra.StringValue {
	return terra.ReferenceAsString(saoe.ref.Append("partition_key"))
}

// PropertyColumns returns a reference to field property_columns of azurerm_stream_analytics_output_eventhub.
func (saoe streamAnalyticsOutputEventhubAttributes) PropertyColumns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](saoe.ref.Append("property_columns"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_output_eventhub.
func (saoe streamAnalyticsOutputEventhubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(saoe.ref.Append("resource_group_name"))
}

// ServicebusNamespace returns a reference to field servicebus_namespace of azurerm_stream_analytics_output_eventhub.
func (saoe streamAnalyticsOutputEventhubAttributes) ServicebusNamespace() terra.StringValue {
	return terra.ReferenceAsString(saoe.ref.Append("servicebus_namespace"))
}

// SharedAccessPolicyKey returns a reference to field shared_access_policy_key of azurerm_stream_analytics_output_eventhub.
func (saoe streamAnalyticsOutputEventhubAttributes) SharedAccessPolicyKey() terra.StringValue {
	return terra.ReferenceAsString(saoe.ref.Append("shared_access_policy_key"))
}

// SharedAccessPolicyName returns a reference to field shared_access_policy_name of azurerm_stream_analytics_output_eventhub.
func (saoe streamAnalyticsOutputEventhubAttributes) SharedAccessPolicyName() terra.StringValue {
	return terra.ReferenceAsString(saoe.ref.Append("shared_access_policy_name"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_output_eventhub.
func (saoe streamAnalyticsOutputEventhubAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(saoe.ref.Append("stream_analytics_job_name"))
}

func (saoe streamAnalyticsOutputEventhubAttributes) Serialization() terra.ListValue[streamanalyticsoutputeventhub.SerializationAttributes] {
	return terra.ReferenceAsList[streamanalyticsoutputeventhub.SerializationAttributes](saoe.ref.Append("serialization"))
}

func (saoe streamAnalyticsOutputEventhubAttributes) Timeouts() streamanalyticsoutputeventhub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsoutputeventhub.TimeoutsAttributes](saoe.ref.Append("timeouts"))
}

type streamAnalyticsOutputEventhubState struct {
	AuthenticationMode     string                                             `json:"authentication_mode"`
	EventhubName           string                                             `json:"eventhub_name"`
	Id                     string                                             `json:"id"`
	Name                   string                                             `json:"name"`
	PartitionKey           string                                             `json:"partition_key"`
	PropertyColumns        []string                                           `json:"property_columns"`
	ResourceGroupName      string                                             `json:"resource_group_name"`
	ServicebusNamespace    string                                             `json:"servicebus_namespace"`
	SharedAccessPolicyKey  string                                             `json:"shared_access_policy_key"`
	SharedAccessPolicyName string                                             `json:"shared_access_policy_name"`
	StreamAnalyticsJobName string                                             `json:"stream_analytics_job_name"`
	Serialization          []streamanalyticsoutputeventhub.SerializationState `json:"serialization"`
	Timeouts               *streamanalyticsoutputeventhub.TimeoutsState       `json:"timeouts"`
}
