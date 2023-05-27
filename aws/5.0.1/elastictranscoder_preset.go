// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elastictranscoderpreset "github.com/golingon/terraproviders/aws/5.0.1/elastictranscoderpreset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElastictranscoderPreset creates a new instance of [ElastictranscoderPreset].
func NewElastictranscoderPreset(name string, args ElastictranscoderPresetArgs) *ElastictranscoderPreset {
	return &ElastictranscoderPreset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElastictranscoderPreset)(nil)

// ElastictranscoderPreset represents the Terraform resource aws_elastictranscoder_preset.
type ElastictranscoderPreset struct {
	Name      string
	Args      ElastictranscoderPresetArgs
	state     *elastictranscoderPresetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElastictranscoderPreset].
func (ep *ElastictranscoderPreset) Type() string {
	return "aws_elastictranscoder_preset"
}

// LocalName returns the local name for [ElastictranscoderPreset].
func (ep *ElastictranscoderPreset) LocalName() string {
	return ep.Name
}

// Configuration returns the configuration (args) for [ElastictranscoderPreset].
func (ep *ElastictranscoderPreset) Configuration() interface{} {
	return ep.Args
}

// DependOn is used for other resources to depend on [ElastictranscoderPreset].
func (ep *ElastictranscoderPreset) DependOn() terra.Reference {
	return terra.ReferenceResource(ep)
}

// Dependencies returns the list of resources [ElastictranscoderPreset] depends_on.
func (ep *ElastictranscoderPreset) Dependencies() terra.Dependencies {
	return ep.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElastictranscoderPreset].
func (ep *ElastictranscoderPreset) LifecycleManagement() *terra.Lifecycle {
	return ep.Lifecycle
}

// Attributes returns the attributes for [ElastictranscoderPreset].
func (ep *ElastictranscoderPreset) Attributes() elastictranscoderPresetAttributes {
	return elastictranscoderPresetAttributes{ref: terra.ReferenceResource(ep)}
}

// ImportState imports the given attribute values into [ElastictranscoderPreset]'s state.
func (ep *ElastictranscoderPreset) ImportState(av io.Reader) error {
	ep.state = &elastictranscoderPresetState{}
	if err := json.NewDecoder(av).Decode(ep.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ep.Type(), ep.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElastictranscoderPreset] has state.
func (ep *ElastictranscoderPreset) State() (*elastictranscoderPresetState, bool) {
	return ep.state, ep.state != nil
}

// StateMust returns the state for [ElastictranscoderPreset]. Panics if the state is nil.
func (ep *ElastictranscoderPreset) StateMust() *elastictranscoderPresetState {
	if ep.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ep.Type(), ep.LocalName()))
	}
	return ep.state
}

// ElastictranscoderPresetArgs contains the configurations for aws_elastictranscoder_preset.
type ElastictranscoderPresetArgs struct {
	// Container: string, required
	Container terra.StringValue `hcl:"container,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// VideoCodecOptions: map of string, optional
	VideoCodecOptions terra.MapValue[terra.StringValue] `hcl:"video_codec_options,attr"`
	// Audio: optional
	Audio *elastictranscoderpreset.Audio `hcl:"audio,block"`
	// AudioCodecOptions: optional
	AudioCodecOptions *elastictranscoderpreset.AudioCodecOptions `hcl:"audio_codec_options,block"`
	// Thumbnails: optional
	Thumbnails *elastictranscoderpreset.Thumbnails `hcl:"thumbnails,block"`
	// Video: optional
	Video *elastictranscoderpreset.Video `hcl:"video,block"`
	// VideoWatermarks: min=0
	VideoWatermarks []elastictranscoderpreset.VideoWatermarks `hcl:"video_watermarks,block" validate:"min=0"`
}
type elastictranscoderPresetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_elastictranscoder_preset.
func (ep elastictranscoderPresetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("arn"))
}

// Container returns a reference to field container of aws_elastictranscoder_preset.
func (ep elastictranscoderPresetAttributes) Container() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("container"))
}

// Description returns a reference to field description of aws_elastictranscoder_preset.
func (ep elastictranscoderPresetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("description"))
}

// Id returns a reference to field id of aws_elastictranscoder_preset.
func (ep elastictranscoderPresetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("id"))
}

// Name returns a reference to field name of aws_elastictranscoder_preset.
func (ep elastictranscoderPresetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("name"))
}

// Type returns a reference to field type of aws_elastictranscoder_preset.
func (ep elastictranscoderPresetAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("type"))
}

// VideoCodecOptions returns a reference to field video_codec_options of aws_elastictranscoder_preset.
func (ep elastictranscoderPresetAttributes) VideoCodecOptions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ep.ref.Append("video_codec_options"))
}

func (ep elastictranscoderPresetAttributes) Audio() terra.ListValue[elastictranscoderpreset.AudioAttributes] {
	return terra.ReferenceAsList[elastictranscoderpreset.AudioAttributes](ep.ref.Append("audio"))
}

func (ep elastictranscoderPresetAttributes) AudioCodecOptions() terra.ListValue[elastictranscoderpreset.AudioCodecOptionsAttributes] {
	return terra.ReferenceAsList[elastictranscoderpreset.AudioCodecOptionsAttributes](ep.ref.Append("audio_codec_options"))
}

func (ep elastictranscoderPresetAttributes) Thumbnails() terra.ListValue[elastictranscoderpreset.ThumbnailsAttributes] {
	return terra.ReferenceAsList[elastictranscoderpreset.ThumbnailsAttributes](ep.ref.Append("thumbnails"))
}

func (ep elastictranscoderPresetAttributes) Video() terra.ListValue[elastictranscoderpreset.VideoAttributes] {
	return terra.ReferenceAsList[elastictranscoderpreset.VideoAttributes](ep.ref.Append("video"))
}

func (ep elastictranscoderPresetAttributes) VideoWatermarks() terra.SetValue[elastictranscoderpreset.VideoWatermarksAttributes] {
	return terra.ReferenceAsSet[elastictranscoderpreset.VideoWatermarksAttributes](ep.ref.Append("video_watermarks"))
}

type elastictranscoderPresetState struct {
	Arn               string                                           `json:"arn"`
	Container         string                                           `json:"container"`
	Description       string                                           `json:"description"`
	Id                string                                           `json:"id"`
	Name              string                                           `json:"name"`
	Type              string                                           `json:"type"`
	VideoCodecOptions map[string]string                                `json:"video_codec_options"`
	Audio             []elastictranscoderpreset.AudioState             `json:"audio"`
	AudioCodecOptions []elastictranscoderpreset.AudioCodecOptionsState `json:"audio_codec_options"`
	Thumbnails        []elastictranscoderpreset.ThumbnailsState        `json:"thumbnails"`
	Video             []elastictranscoderpreset.VideoState             `json:"video"`
	VideoWatermarks   []elastictranscoderpreset.VideoWatermarksState   `json:"video_watermarks"`
}
