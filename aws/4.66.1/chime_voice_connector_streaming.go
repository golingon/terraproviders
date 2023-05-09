// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	chimevoiceconnectorstreaming "github.com/golingon/terraproviders/aws/4.66.1/chimevoiceconnectorstreaming"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewChimeVoiceConnectorStreaming creates a new instance of [ChimeVoiceConnectorStreaming].
func NewChimeVoiceConnectorStreaming(name string, args ChimeVoiceConnectorStreamingArgs) *ChimeVoiceConnectorStreaming {
	return &ChimeVoiceConnectorStreaming{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ChimeVoiceConnectorStreaming)(nil)

// ChimeVoiceConnectorStreaming represents the Terraform resource aws_chime_voice_connector_streaming.
type ChimeVoiceConnectorStreaming struct {
	Name      string
	Args      ChimeVoiceConnectorStreamingArgs
	state     *chimeVoiceConnectorStreamingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ChimeVoiceConnectorStreaming].
func (cvcs *ChimeVoiceConnectorStreaming) Type() string {
	return "aws_chime_voice_connector_streaming"
}

// LocalName returns the local name for [ChimeVoiceConnectorStreaming].
func (cvcs *ChimeVoiceConnectorStreaming) LocalName() string {
	return cvcs.Name
}

// Configuration returns the configuration (args) for [ChimeVoiceConnectorStreaming].
func (cvcs *ChimeVoiceConnectorStreaming) Configuration() interface{} {
	return cvcs.Args
}

// DependOn is used for other resources to depend on [ChimeVoiceConnectorStreaming].
func (cvcs *ChimeVoiceConnectorStreaming) DependOn() terra.Reference {
	return terra.ReferenceResource(cvcs)
}

// Dependencies returns the list of resources [ChimeVoiceConnectorStreaming] depends_on.
func (cvcs *ChimeVoiceConnectorStreaming) Dependencies() terra.Dependencies {
	return cvcs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ChimeVoiceConnectorStreaming].
func (cvcs *ChimeVoiceConnectorStreaming) LifecycleManagement() *terra.Lifecycle {
	return cvcs.Lifecycle
}

// Attributes returns the attributes for [ChimeVoiceConnectorStreaming].
func (cvcs *ChimeVoiceConnectorStreaming) Attributes() chimeVoiceConnectorStreamingAttributes {
	return chimeVoiceConnectorStreamingAttributes{ref: terra.ReferenceResource(cvcs)}
}

// ImportState imports the given attribute values into [ChimeVoiceConnectorStreaming]'s state.
func (cvcs *ChimeVoiceConnectorStreaming) ImportState(av io.Reader) error {
	cvcs.state = &chimeVoiceConnectorStreamingState{}
	if err := json.NewDecoder(av).Decode(cvcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cvcs.Type(), cvcs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ChimeVoiceConnectorStreaming] has state.
func (cvcs *ChimeVoiceConnectorStreaming) State() (*chimeVoiceConnectorStreamingState, bool) {
	return cvcs.state, cvcs.state != nil
}

// StateMust returns the state for [ChimeVoiceConnectorStreaming]. Panics if the state is nil.
func (cvcs *ChimeVoiceConnectorStreaming) StateMust() *chimeVoiceConnectorStreamingState {
	if cvcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cvcs.Type(), cvcs.LocalName()))
	}
	return cvcs.state
}

// ChimeVoiceConnectorStreamingArgs contains the configurations for aws_chime_voice_connector_streaming.
type ChimeVoiceConnectorStreamingArgs struct {
	// DataRetention: number, required
	DataRetention terra.NumberValue `hcl:"data_retention,attr" validate:"required"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// StreamingNotificationTargets: set of string, optional
	StreamingNotificationTargets terra.SetValue[terra.StringValue] `hcl:"streaming_notification_targets,attr"`
	// VoiceConnectorId: string, required
	VoiceConnectorId terra.StringValue `hcl:"voice_connector_id,attr" validate:"required"`
	// MediaInsightsConfiguration: optional
	MediaInsightsConfiguration *chimevoiceconnectorstreaming.MediaInsightsConfiguration `hcl:"media_insights_configuration,block"`
}
type chimeVoiceConnectorStreamingAttributes struct {
	ref terra.Reference
}

// DataRetention returns a reference to field data_retention of aws_chime_voice_connector_streaming.
func (cvcs chimeVoiceConnectorStreamingAttributes) DataRetention() terra.NumberValue {
	return terra.ReferenceAsNumber(cvcs.ref.Append("data_retention"))
}

// Disabled returns a reference to field disabled of aws_chime_voice_connector_streaming.
func (cvcs chimeVoiceConnectorStreamingAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(cvcs.ref.Append("disabled"))
}

// Id returns a reference to field id of aws_chime_voice_connector_streaming.
func (cvcs chimeVoiceConnectorStreamingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cvcs.ref.Append("id"))
}

// StreamingNotificationTargets returns a reference to field streaming_notification_targets of aws_chime_voice_connector_streaming.
func (cvcs chimeVoiceConnectorStreamingAttributes) StreamingNotificationTargets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cvcs.ref.Append("streaming_notification_targets"))
}

// VoiceConnectorId returns a reference to field voice_connector_id of aws_chime_voice_connector_streaming.
func (cvcs chimeVoiceConnectorStreamingAttributes) VoiceConnectorId() terra.StringValue {
	return terra.ReferenceAsString(cvcs.ref.Append("voice_connector_id"))
}

func (cvcs chimeVoiceConnectorStreamingAttributes) MediaInsightsConfiguration() terra.ListValue[chimevoiceconnectorstreaming.MediaInsightsConfigurationAttributes] {
	return terra.ReferenceAsList[chimevoiceconnectorstreaming.MediaInsightsConfigurationAttributes](cvcs.ref.Append("media_insights_configuration"))
}

type chimeVoiceConnectorStreamingState struct {
	DataRetention                float64                                                        `json:"data_retention"`
	Disabled                     bool                                                           `json:"disabled"`
	Id                           string                                                         `json:"id"`
	StreamingNotificationTargets []string                                                       `json:"streaming_notification_targets"`
	VoiceConnectorId             string                                                         `json:"voice_connector_id"`
	MediaInsightsConfiguration   []chimevoiceconnectorstreaming.MediaInsightsConfigurationState `json:"media_insights_configuration"`
}
