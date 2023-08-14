// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ivschannel "github.com/golingon/terraproviders/aws/5.9.0/ivschannel"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIvsChannel creates a new instance of [IvsChannel].
func NewIvsChannel(name string, args IvsChannelArgs) *IvsChannel {
	return &IvsChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IvsChannel)(nil)

// IvsChannel represents the Terraform resource aws_ivs_channel.
type IvsChannel struct {
	Name      string
	Args      IvsChannelArgs
	state     *ivsChannelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IvsChannel].
func (ic *IvsChannel) Type() string {
	return "aws_ivs_channel"
}

// LocalName returns the local name for [IvsChannel].
func (ic *IvsChannel) LocalName() string {
	return ic.Name
}

// Configuration returns the configuration (args) for [IvsChannel].
func (ic *IvsChannel) Configuration() interface{} {
	return ic.Args
}

// DependOn is used for other resources to depend on [IvsChannel].
func (ic *IvsChannel) DependOn() terra.Reference {
	return terra.ReferenceResource(ic)
}

// Dependencies returns the list of resources [IvsChannel] depends_on.
func (ic *IvsChannel) Dependencies() terra.Dependencies {
	return ic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IvsChannel].
func (ic *IvsChannel) LifecycleManagement() *terra.Lifecycle {
	return ic.Lifecycle
}

// Attributes returns the attributes for [IvsChannel].
func (ic *IvsChannel) Attributes() ivsChannelAttributes {
	return ivsChannelAttributes{ref: terra.ReferenceResource(ic)}
}

// ImportState imports the given attribute values into [IvsChannel]'s state.
func (ic *IvsChannel) ImportState(av io.Reader) error {
	ic.state = &ivsChannelState{}
	if err := json.NewDecoder(av).Decode(ic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ic.Type(), ic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IvsChannel] has state.
func (ic *IvsChannel) State() (*ivsChannelState, bool) {
	return ic.state, ic.state != nil
}

// StateMust returns the state for [IvsChannel]. Panics if the state is nil.
func (ic *IvsChannel) StateMust() *ivsChannelState {
	if ic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ic.Type(), ic.LocalName()))
	}
	return ic.state
}

// IvsChannelArgs contains the configurations for aws_ivs_channel.
type IvsChannelArgs struct {
	// Authorized: bool, optional
	Authorized terra.BoolValue `hcl:"authorized,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LatencyMode: string, optional
	LatencyMode terra.StringValue `hcl:"latency_mode,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// RecordingConfigurationArn: string, optional
	RecordingConfigurationArn terra.StringValue `hcl:"recording_configuration_arn,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Timeouts: optional
	Timeouts *ivschannel.Timeouts `hcl:"timeouts,block"`
}
type ivsChannelAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ivs_channel.
func (ic ivsChannelAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("arn"))
}

// Authorized returns a reference to field authorized of aws_ivs_channel.
func (ic ivsChannelAttributes) Authorized() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("authorized"))
}

// Id returns a reference to field id of aws_ivs_channel.
func (ic ivsChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("id"))
}

// IngestEndpoint returns a reference to field ingest_endpoint of aws_ivs_channel.
func (ic ivsChannelAttributes) IngestEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("ingest_endpoint"))
}

// LatencyMode returns a reference to field latency_mode of aws_ivs_channel.
func (ic ivsChannelAttributes) LatencyMode() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("latency_mode"))
}

// Name returns a reference to field name of aws_ivs_channel.
func (ic ivsChannelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("name"))
}

// PlaybackUrl returns a reference to field playback_url of aws_ivs_channel.
func (ic ivsChannelAttributes) PlaybackUrl() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("playback_url"))
}

// RecordingConfigurationArn returns a reference to field recording_configuration_arn of aws_ivs_channel.
func (ic ivsChannelAttributes) RecordingConfigurationArn() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("recording_configuration_arn"))
}

// Tags returns a reference to field tags of aws_ivs_channel.
func (ic ivsChannelAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ic.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ivs_channel.
func (ic ivsChannelAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ic.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_ivs_channel.
func (ic ivsChannelAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("type"))
}

func (ic ivsChannelAttributes) Timeouts() ivschannel.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ivschannel.TimeoutsAttributes](ic.ref.Append("timeouts"))
}

type ivsChannelState struct {
	Arn                       string                    `json:"arn"`
	Authorized                bool                      `json:"authorized"`
	Id                        string                    `json:"id"`
	IngestEndpoint            string                    `json:"ingest_endpoint"`
	LatencyMode               string                    `json:"latency_mode"`
	Name                      string                    `json:"name"`
	PlaybackUrl               string                    `json:"playback_url"`
	RecordingConfigurationArn string                    `json:"recording_configuration_arn"`
	Tags                      map[string]string         `json:"tags"`
	TagsAll                   map[string]string         `json:"tags_all"`
	Type                      string                    `json:"type"`
	Timeouts                  *ivschannel.TimeoutsState `json:"timeouts"`
}