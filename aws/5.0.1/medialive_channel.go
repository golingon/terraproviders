// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	medialivechannel "github.com/golingon/terraproviders/aws/5.0.1/medialivechannel"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMedialiveChannel creates a new instance of [MedialiveChannel].
func NewMedialiveChannel(name string, args MedialiveChannelArgs) *MedialiveChannel {
	return &MedialiveChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MedialiveChannel)(nil)

// MedialiveChannel represents the Terraform resource aws_medialive_channel.
type MedialiveChannel struct {
	Name      string
	Args      MedialiveChannelArgs
	state     *medialiveChannelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MedialiveChannel].
func (mc *MedialiveChannel) Type() string {
	return "aws_medialive_channel"
}

// LocalName returns the local name for [MedialiveChannel].
func (mc *MedialiveChannel) LocalName() string {
	return mc.Name
}

// Configuration returns the configuration (args) for [MedialiveChannel].
func (mc *MedialiveChannel) Configuration() interface{} {
	return mc.Args
}

// DependOn is used for other resources to depend on [MedialiveChannel].
func (mc *MedialiveChannel) DependOn() terra.Reference {
	return terra.ReferenceResource(mc)
}

// Dependencies returns the list of resources [MedialiveChannel] depends_on.
func (mc *MedialiveChannel) Dependencies() terra.Dependencies {
	return mc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MedialiveChannel].
func (mc *MedialiveChannel) LifecycleManagement() *terra.Lifecycle {
	return mc.Lifecycle
}

// Attributes returns the attributes for [MedialiveChannel].
func (mc *MedialiveChannel) Attributes() medialiveChannelAttributes {
	return medialiveChannelAttributes{ref: terra.ReferenceResource(mc)}
}

// ImportState imports the given attribute values into [MedialiveChannel]'s state.
func (mc *MedialiveChannel) ImportState(av io.Reader) error {
	mc.state = &medialiveChannelState{}
	if err := json.NewDecoder(av).Decode(mc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mc.Type(), mc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MedialiveChannel] has state.
func (mc *MedialiveChannel) State() (*medialiveChannelState, bool) {
	return mc.state, mc.state != nil
}

// StateMust returns the state for [MedialiveChannel]. Panics if the state is nil.
func (mc *MedialiveChannel) StateMust() *medialiveChannelState {
	if mc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mc.Type(), mc.LocalName()))
	}
	return mc.state
}

// MedialiveChannelArgs contains the configurations for aws_medialive_channel.
type MedialiveChannelArgs struct {
	// ChannelClass: string, required
	ChannelClass terra.StringValue `hcl:"channel_class,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogLevel: string, optional
	LogLevel terra.StringValue `hcl:"log_level,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// StartChannel: bool, optional
	StartChannel terra.BoolValue `hcl:"start_channel,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// CdiInputSpecification: optional
	CdiInputSpecification *medialivechannel.CdiInputSpecification `hcl:"cdi_input_specification,block"`
	// Destinations: min=1
	Destinations []medialivechannel.Destinations `hcl:"destinations,block" validate:"min=1"`
	// EncoderSettings: required
	EncoderSettings *medialivechannel.EncoderSettings `hcl:"encoder_settings,block" validate:"required"`
	// InputAttachments: min=1
	InputAttachments []medialivechannel.InputAttachments `hcl:"input_attachments,block" validate:"min=1"`
	// InputSpecification: required
	InputSpecification *medialivechannel.InputSpecification `hcl:"input_specification,block" validate:"required"`
	// Maintenance: optional
	Maintenance *medialivechannel.Maintenance `hcl:"maintenance,block"`
	// Timeouts: optional
	Timeouts *medialivechannel.Timeouts `hcl:"timeouts,block"`
	// Vpc: optional
	Vpc *medialivechannel.Vpc `hcl:"vpc,block"`
}
type medialiveChannelAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_medialive_channel.
func (mc medialiveChannelAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("arn"))
}

// ChannelClass returns a reference to field channel_class of aws_medialive_channel.
func (mc medialiveChannelAttributes) ChannelClass() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("channel_class"))
}

// ChannelId returns a reference to field channel_id of aws_medialive_channel.
func (mc medialiveChannelAttributes) ChannelId() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("channel_id"))
}

// Id returns a reference to field id of aws_medialive_channel.
func (mc medialiveChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("id"))
}

// LogLevel returns a reference to field log_level of aws_medialive_channel.
func (mc medialiveChannelAttributes) LogLevel() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("log_level"))
}

// Name returns a reference to field name of aws_medialive_channel.
func (mc medialiveChannelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_medialive_channel.
func (mc medialiveChannelAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("role_arn"))
}

// StartChannel returns a reference to field start_channel of aws_medialive_channel.
func (mc medialiveChannelAttributes) StartChannel() terra.BoolValue {
	return terra.ReferenceAsBool(mc.ref.Append("start_channel"))
}

// Tags returns a reference to field tags of aws_medialive_channel.
func (mc medialiveChannelAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_medialive_channel.
func (mc medialiveChannelAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("tags_all"))
}

func (mc medialiveChannelAttributes) CdiInputSpecification() terra.ListValue[medialivechannel.CdiInputSpecificationAttributes] {
	return terra.ReferenceAsList[medialivechannel.CdiInputSpecificationAttributes](mc.ref.Append("cdi_input_specification"))
}

func (mc medialiveChannelAttributes) Destinations() terra.SetValue[medialivechannel.DestinationsAttributes] {
	return terra.ReferenceAsSet[medialivechannel.DestinationsAttributes](mc.ref.Append("destinations"))
}

func (mc medialiveChannelAttributes) EncoderSettings() terra.ListValue[medialivechannel.EncoderSettingsAttributes] {
	return terra.ReferenceAsList[medialivechannel.EncoderSettingsAttributes](mc.ref.Append("encoder_settings"))
}

func (mc medialiveChannelAttributes) InputAttachments() terra.SetValue[medialivechannel.InputAttachmentsAttributes] {
	return terra.ReferenceAsSet[medialivechannel.InputAttachmentsAttributes](mc.ref.Append("input_attachments"))
}

func (mc medialiveChannelAttributes) InputSpecification() terra.ListValue[medialivechannel.InputSpecificationAttributes] {
	return terra.ReferenceAsList[medialivechannel.InputSpecificationAttributes](mc.ref.Append("input_specification"))
}

func (mc medialiveChannelAttributes) Maintenance() terra.ListValue[medialivechannel.MaintenanceAttributes] {
	return terra.ReferenceAsList[medialivechannel.MaintenanceAttributes](mc.ref.Append("maintenance"))
}

func (mc medialiveChannelAttributes) Timeouts() medialivechannel.TimeoutsAttributes {
	return terra.ReferenceAsSingle[medialivechannel.TimeoutsAttributes](mc.ref.Append("timeouts"))
}

func (mc medialiveChannelAttributes) Vpc() terra.ListValue[medialivechannel.VpcAttributes] {
	return terra.ReferenceAsList[medialivechannel.VpcAttributes](mc.ref.Append("vpc"))
}

type medialiveChannelState struct {
	Arn                   string                                        `json:"arn"`
	ChannelClass          string                                        `json:"channel_class"`
	ChannelId             string                                        `json:"channel_id"`
	Id                    string                                        `json:"id"`
	LogLevel              string                                        `json:"log_level"`
	Name                  string                                        `json:"name"`
	RoleArn               string                                        `json:"role_arn"`
	StartChannel          bool                                          `json:"start_channel"`
	Tags                  map[string]string                             `json:"tags"`
	TagsAll               map[string]string                             `json:"tags_all"`
	CdiInputSpecification []medialivechannel.CdiInputSpecificationState `json:"cdi_input_specification"`
	Destinations          []medialivechannel.DestinationsState          `json:"destinations"`
	EncoderSettings       []medialivechannel.EncoderSettingsState       `json:"encoder_settings"`
	InputAttachments      []medialivechannel.InputAttachmentsState      `json:"input_attachments"`
	InputSpecification    []medialivechannel.InputSpecificationState    `json:"input_specification"`
	Maintenance           []medialivechannel.MaintenanceState           `json:"maintenance"`
	Timeouts              *medialivechannel.TimeoutsState               `json:"timeouts"`
	Vpc                   []medialivechannel.VpcState                   `json:"vpc"`
}
