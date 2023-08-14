// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	kinesisvideostream "github.com/golingon/terraproviders/aws/5.12.0/kinesisvideostream"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKinesisVideoStream creates a new instance of [KinesisVideoStream].
func NewKinesisVideoStream(name string, args KinesisVideoStreamArgs) *KinesisVideoStream {
	return &KinesisVideoStream{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KinesisVideoStream)(nil)

// KinesisVideoStream represents the Terraform resource aws_kinesis_video_stream.
type KinesisVideoStream struct {
	Name      string
	Args      KinesisVideoStreamArgs
	state     *kinesisVideoStreamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KinesisVideoStream].
func (kvs *KinesisVideoStream) Type() string {
	return "aws_kinesis_video_stream"
}

// LocalName returns the local name for [KinesisVideoStream].
func (kvs *KinesisVideoStream) LocalName() string {
	return kvs.Name
}

// Configuration returns the configuration (args) for [KinesisVideoStream].
func (kvs *KinesisVideoStream) Configuration() interface{} {
	return kvs.Args
}

// DependOn is used for other resources to depend on [KinesisVideoStream].
func (kvs *KinesisVideoStream) DependOn() terra.Reference {
	return terra.ReferenceResource(kvs)
}

// Dependencies returns the list of resources [KinesisVideoStream] depends_on.
func (kvs *KinesisVideoStream) Dependencies() terra.Dependencies {
	return kvs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KinesisVideoStream].
func (kvs *KinesisVideoStream) LifecycleManagement() *terra.Lifecycle {
	return kvs.Lifecycle
}

// Attributes returns the attributes for [KinesisVideoStream].
func (kvs *KinesisVideoStream) Attributes() kinesisVideoStreamAttributes {
	return kinesisVideoStreamAttributes{ref: terra.ReferenceResource(kvs)}
}

// ImportState imports the given attribute values into [KinesisVideoStream]'s state.
func (kvs *KinesisVideoStream) ImportState(av io.Reader) error {
	kvs.state = &kinesisVideoStreamState{}
	if err := json.NewDecoder(av).Decode(kvs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kvs.Type(), kvs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KinesisVideoStream] has state.
func (kvs *KinesisVideoStream) State() (*kinesisVideoStreamState, bool) {
	return kvs.state, kvs.state != nil
}

// StateMust returns the state for [KinesisVideoStream]. Panics if the state is nil.
func (kvs *KinesisVideoStream) StateMust() *kinesisVideoStreamState {
	if kvs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kvs.Type(), kvs.LocalName()))
	}
	return kvs.state
}

// KinesisVideoStreamArgs contains the configurations for aws_kinesis_video_stream.
type KinesisVideoStreamArgs struct {
	// DataRetentionInHours: number, optional
	DataRetentionInHours terra.NumberValue `hcl:"data_retention_in_hours,attr"`
	// DeviceName: string, optional
	DeviceName terra.StringValue `hcl:"device_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// MediaType: string, optional
	MediaType terra.StringValue `hcl:"media_type,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *kinesisvideostream.Timeouts `hcl:"timeouts,block"`
}
type kinesisVideoStreamAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kinesis_video_stream.
func (kvs kinesisVideoStreamAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("arn"))
}

// CreationTime returns a reference to field creation_time of aws_kinesis_video_stream.
func (kvs kinesisVideoStreamAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("creation_time"))
}

// DataRetentionInHours returns a reference to field data_retention_in_hours of aws_kinesis_video_stream.
func (kvs kinesisVideoStreamAttributes) DataRetentionInHours() terra.NumberValue {
	return terra.ReferenceAsNumber(kvs.ref.Append("data_retention_in_hours"))
}

// DeviceName returns a reference to field device_name of aws_kinesis_video_stream.
func (kvs kinesisVideoStreamAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("device_name"))
}

// Id returns a reference to field id of aws_kinesis_video_stream.
func (kvs kinesisVideoStreamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_kinesis_video_stream.
func (kvs kinesisVideoStreamAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("kms_key_id"))
}

// MediaType returns a reference to field media_type of aws_kinesis_video_stream.
func (kvs kinesisVideoStreamAttributes) MediaType() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("media_type"))
}

// Name returns a reference to field name of aws_kinesis_video_stream.
func (kvs kinesisVideoStreamAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_kinesis_video_stream.
func (kvs kinesisVideoStreamAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kvs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kinesis_video_stream.
func (kvs kinesisVideoStreamAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kvs.ref.Append("tags_all"))
}

// Version returns a reference to field version of aws_kinesis_video_stream.
func (kvs kinesisVideoStreamAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("version"))
}

func (kvs kinesisVideoStreamAttributes) Timeouts() kinesisvideostream.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kinesisvideostream.TimeoutsAttributes](kvs.ref.Append("timeouts"))
}

type kinesisVideoStreamState struct {
	Arn                  string                            `json:"arn"`
	CreationTime         string                            `json:"creation_time"`
	DataRetentionInHours float64                           `json:"data_retention_in_hours"`
	DeviceName           string                            `json:"device_name"`
	Id                   string                            `json:"id"`
	KmsKeyId             string                            `json:"kms_key_id"`
	MediaType            string                            `json:"media_type"`
	Name                 string                            `json:"name"`
	Tags                 map[string]string                 `json:"tags"`
	TagsAll              map[string]string                 `json:"tags_all"`
	Version              string                            `json:"version"`
	Timeouts             *kinesisvideostream.TimeoutsState `json:"timeouts"`
}