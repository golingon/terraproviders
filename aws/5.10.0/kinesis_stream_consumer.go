// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKinesisStreamConsumer creates a new instance of [KinesisStreamConsumer].
func NewKinesisStreamConsumer(name string, args KinesisStreamConsumerArgs) *KinesisStreamConsumer {
	return &KinesisStreamConsumer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KinesisStreamConsumer)(nil)

// KinesisStreamConsumer represents the Terraform resource aws_kinesis_stream_consumer.
type KinesisStreamConsumer struct {
	Name      string
	Args      KinesisStreamConsumerArgs
	state     *kinesisStreamConsumerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KinesisStreamConsumer].
func (ksc *KinesisStreamConsumer) Type() string {
	return "aws_kinesis_stream_consumer"
}

// LocalName returns the local name for [KinesisStreamConsumer].
func (ksc *KinesisStreamConsumer) LocalName() string {
	return ksc.Name
}

// Configuration returns the configuration (args) for [KinesisStreamConsumer].
func (ksc *KinesisStreamConsumer) Configuration() interface{} {
	return ksc.Args
}

// DependOn is used for other resources to depend on [KinesisStreamConsumer].
func (ksc *KinesisStreamConsumer) DependOn() terra.Reference {
	return terra.ReferenceResource(ksc)
}

// Dependencies returns the list of resources [KinesisStreamConsumer] depends_on.
func (ksc *KinesisStreamConsumer) Dependencies() terra.Dependencies {
	return ksc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KinesisStreamConsumer].
func (ksc *KinesisStreamConsumer) LifecycleManagement() *terra.Lifecycle {
	return ksc.Lifecycle
}

// Attributes returns the attributes for [KinesisStreamConsumer].
func (ksc *KinesisStreamConsumer) Attributes() kinesisStreamConsumerAttributes {
	return kinesisStreamConsumerAttributes{ref: terra.ReferenceResource(ksc)}
}

// ImportState imports the given attribute values into [KinesisStreamConsumer]'s state.
func (ksc *KinesisStreamConsumer) ImportState(av io.Reader) error {
	ksc.state = &kinesisStreamConsumerState{}
	if err := json.NewDecoder(av).Decode(ksc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ksc.Type(), ksc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KinesisStreamConsumer] has state.
func (ksc *KinesisStreamConsumer) State() (*kinesisStreamConsumerState, bool) {
	return ksc.state, ksc.state != nil
}

// StateMust returns the state for [KinesisStreamConsumer]. Panics if the state is nil.
func (ksc *KinesisStreamConsumer) StateMust() *kinesisStreamConsumerState {
	if ksc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ksc.Type(), ksc.LocalName()))
	}
	return ksc.state
}

// KinesisStreamConsumerArgs contains the configurations for aws_kinesis_stream_consumer.
type KinesisStreamConsumerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StreamArn: string, required
	StreamArn terra.StringValue `hcl:"stream_arn,attr" validate:"required"`
}
type kinesisStreamConsumerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kinesis_stream_consumer.
func (ksc kinesisStreamConsumerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("arn"))
}

// CreationTimestamp returns a reference to field creation_timestamp of aws_kinesis_stream_consumer.
func (ksc kinesisStreamConsumerAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("creation_timestamp"))
}

// Id returns a reference to field id of aws_kinesis_stream_consumer.
func (ksc kinesisStreamConsumerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("id"))
}

// Name returns a reference to field name of aws_kinesis_stream_consumer.
func (ksc kinesisStreamConsumerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("name"))
}

// StreamArn returns a reference to field stream_arn of aws_kinesis_stream_consumer.
func (ksc kinesisStreamConsumerAttributes) StreamArn() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("stream_arn"))
}

type kinesisStreamConsumerState struct {
	Arn               string `json:"arn"`
	CreationTimestamp string `json:"creation_timestamp"`
	Id                string `json:"id"`
	Name              string `json:"name"`
	StreamArn         string `json:"stream_arn"`
}
