// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	kinesisstream "github.com/golingon/terraproviders/aws/4.63.0/kinesisstream"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKinesisStream creates a new instance of [KinesisStream].
func NewKinesisStream(name string, args KinesisStreamArgs) *KinesisStream {
	return &KinesisStream{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KinesisStream)(nil)

// KinesisStream represents the Terraform resource aws_kinesis_stream.
type KinesisStream struct {
	Name      string
	Args      KinesisStreamArgs
	state     *kinesisStreamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KinesisStream].
func (ks *KinesisStream) Type() string {
	return "aws_kinesis_stream"
}

// LocalName returns the local name for [KinesisStream].
func (ks *KinesisStream) LocalName() string {
	return ks.Name
}

// Configuration returns the configuration (args) for [KinesisStream].
func (ks *KinesisStream) Configuration() interface{} {
	return ks.Args
}

// DependOn is used for other resources to depend on [KinesisStream].
func (ks *KinesisStream) DependOn() terra.Reference {
	return terra.ReferenceResource(ks)
}

// Dependencies returns the list of resources [KinesisStream] depends_on.
func (ks *KinesisStream) Dependencies() terra.Dependencies {
	return ks.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KinesisStream].
func (ks *KinesisStream) LifecycleManagement() *terra.Lifecycle {
	return ks.Lifecycle
}

// Attributes returns the attributes for [KinesisStream].
func (ks *KinesisStream) Attributes() kinesisStreamAttributes {
	return kinesisStreamAttributes{ref: terra.ReferenceResource(ks)}
}

// ImportState imports the given attribute values into [KinesisStream]'s state.
func (ks *KinesisStream) ImportState(av io.Reader) error {
	ks.state = &kinesisStreamState{}
	if err := json.NewDecoder(av).Decode(ks.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ks.Type(), ks.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KinesisStream] has state.
func (ks *KinesisStream) State() (*kinesisStreamState, bool) {
	return ks.state, ks.state != nil
}

// StateMust returns the state for [KinesisStream]. Panics if the state is nil.
func (ks *KinesisStream) StateMust() *kinesisStreamState {
	if ks.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ks.Type(), ks.LocalName()))
	}
	return ks.state
}

// KinesisStreamArgs contains the configurations for aws_kinesis_stream.
type KinesisStreamArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// EncryptionType: string, optional
	EncryptionType terra.StringValue `hcl:"encryption_type,attr"`
	// EnforceConsumerDeletion: bool, optional
	EnforceConsumerDeletion terra.BoolValue `hcl:"enforce_consumer_deletion,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RetentionPeriod: number, optional
	RetentionPeriod terra.NumberValue `hcl:"retention_period,attr"`
	// ShardCount: number, optional
	ShardCount terra.NumberValue `hcl:"shard_count,attr"`
	// ShardLevelMetrics: set of string, optional
	ShardLevelMetrics terra.SetValue[terra.StringValue] `hcl:"shard_level_metrics,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// StreamModeDetails: optional
	StreamModeDetails *kinesisstream.StreamModeDetails `hcl:"stream_mode_details,block"`
	// Timeouts: optional
	Timeouts *kinesisstream.Timeouts `hcl:"timeouts,block"`
}
type kinesisStreamAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kinesis_stream.
func (ks kinesisStreamAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("arn"))
}

// EncryptionType returns a reference to field encryption_type of aws_kinesis_stream.
func (ks kinesisStreamAttributes) EncryptionType() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("encryption_type"))
}

// EnforceConsumerDeletion returns a reference to field enforce_consumer_deletion of aws_kinesis_stream.
func (ks kinesisStreamAttributes) EnforceConsumerDeletion() terra.BoolValue {
	return terra.ReferenceAsBool(ks.ref.Append("enforce_consumer_deletion"))
}

// Id returns a reference to field id of aws_kinesis_stream.
func (ks kinesisStreamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_kinesis_stream.
func (ks kinesisStreamAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("kms_key_id"))
}

// Name returns a reference to field name of aws_kinesis_stream.
func (ks kinesisStreamAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("name"))
}

// RetentionPeriod returns a reference to field retention_period of aws_kinesis_stream.
func (ks kinesisStreamAttributes) RetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(ks.ref.Append("retention_period"))
}

// ShardCount returns a reference to field shard_count of aws_kinesis_stream.
func (ks kinesisStreamAttributes) ShardCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ks.ref.Append("shard_count"))
}

// ShardLevelMetrics returns a reference to field shard_level_metrics of aws_kinesis_stream.
func (ks kinesisStreamAttributes) ShardLevelMetrics() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ks.ref.Append("shard_level_metrics"))
}

// Tags returns a reference to field tags of aws_kinesis_stream.
func (ks kinesisStreamAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ks.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kinesis_stream.
func (ks kinesisStreamAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ks.ref.Append("tags_all"))
}

func (ks kinesisStreamAttributes) StreamModeDetails() terra.ListValue[kinesisstream.StreamModeDetailsAttributes] {
	return terra.ReferenceAsList[kinesisstream.StreamModeDetailsAttributes](ks.ref.Append("stream_mode_details"))
}

func (ks kinesisStreamAttributes) Timeouts() kinesisstream.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kinesisstream.TimeoutsAttributes](ks.ref.Append("timeouts"))
}

type kinesisStreamState struct {
	Arn                     string                                 `json:"arn"`
	EncryptionType          string                                 `json:"encryption_type"`
	EnforceConsumerDeletion bool                                   `json:"enforce_consumer_deletion"`
	Id                      string                                 `json:"id"`
	KmsKeyId                string                                 `json:"kms_key_id"`
	Name                    string                                 `json:"name"`
	RetentionPeriod         float64                                `json:"retention_period"`
	ShardCount              float64                                `json:"shard_count"`
	ShardLevelMetrics       []string                               `json:"shard_level_metrics"`
	Tags                    map[string]string                      `json:"tags"`
	TagsAll                 map[string]string                      `json:"tags_all"`
	StreamModeDetails       []kinesisstream.StreamModeDetailsState `json:"stream_mode_details"`
	Timeouts                *kinesisstream.TimeoutsState           `json:"timeouts"`
}