// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	kinesisfirehosedeliverystream "github.com/golingon/terraproviders/aws/5.8.0/kinesisfirehosedeliverystream"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKinesisFirehoseDeliveryStream creates a new instance of [KinesisFirehoseDeliveryStream].
func NewKinesisFirehoseDeliveryStream(name string, args KinesisFirehoseDeliveryStreamArgs) *KinesisFirehoseDeliveryStream {
	return &KinesisFirehoseDeliveryStream{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KinesisFirehoseDeliveryStream)(nil)

// KinesisFirehoseDeliveryStream represents the Terraform resource aws_kinesis_firehose_delivery_stream.
type KinesisFirehoseDeliveryStream struct {
	Name      string
	Args      KinesisFirehoseDeliveryStreamArgs
	state     *kinesisFirehoseDeliveryStreamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KinesisFirehoseDeliveryStream].
func (kfds *KinesisFirehoseDeliveryStream) Type() string {
	return "aws_kinesis_firehose_delivery_stream"
}

// LocalName returns the local name for [KinesisFirehoseDeliveryStream].
func (kfds *KinesisFirehoseDeliveryStream) LocalName() string {
	return kfds.Name
}

// Configuration returns the configuration (args) for [KinesisFirehoseDeliveryStream].
func (kfds *KinesisFirehoseDeliveryStream) Configuration() interface{} {
	return kfds.Args
}

// DependOn is used for other resources to depend on [KinesisFirehoseDeliveryStream].
func (kfds *KinesisFirehoseDeliveryStream) DependOn() terra.Reference {
	return terra.ReferenceResource(kfds)
}

// Dependencies returns the list of resources [KinesisFirehoseDeliveryStream] depends_on.
func (kfds *KinesisFirehoseDeliveryStream) Dependencies() terra.Dependencies {
	return kfds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KinesisFirehoseDeliveryStream].
func (kfds *KinesisFirehoseDeliveryStream) LifecycleManagement() *terra.Lifecycle {
	return kfds.Lifecycle
}

// Attributes returns the attributes for [KinesisFirehoseDeliveryStream].
func (kfds *KinesisFirehoseDeliveryStream) Attributes() kinesisFirehoseDeliveryStreamAttributes {
	return kinesisFirehoseDeliveryStreamAttributes{ref: terra.ReferenceResource(kfds)}
}

// ImportState imports the given attribute values into [KinesisFirehoseDeliveryStream]'s state.
func (kfds *KinesisFirehoseDeliveryStream) ImportState(av io.Reader) error {
	kfds.state = &kinesisFirehoseDeliveryStreamState{}
	if err := json.NewDecoder(av).Decode(kfds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kfds.Type(), kfds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KinesisFirehoseDeliveryStream] has state.
func (kfds *KinesisFirehoseDeliveryStream) State() (*kinesisFirehoseDeliveryStreamState, bool) {
	return kfds.state, kfds.state != nil
}

// StateMust returns the state for [KinesisFirehoseDeliveryStream]. Panics if the state is nil.
func (kfds *KinesisFirehoseDeliveryStream) StateMust() *kinesisFirehoseDeliveryStreamState {
	if kfds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kfds.Type(), kfds.LocalName()))
	}
	return kfds.state
}

// KinesisFirehoseDeliveryStreamArgs contains the configurations for aws_kinesis_firehose_delivery_stream.
type KinesisFirehoseDeliveryStreamArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Destination: string, required
	Destination terra.StringValue `hcl:"destination,attr" validate:"required"`
	// DestinationId: string, optional
	DestinationId terra.StringValue `hcl:"destination_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VersionId: string, optional
	VersionId terra.StringValue `hcl:"version_id,attr"`
	// ElasticsearchConfiguration: optional
	ElasticsearchConfiguration *kinesisfirehosedeliverystream.ElasticsearchConfiguration `hcl:"elasticsearch_configuration,block"`
	// ExtendedS3Configuration: optional
	ExtendedS3Configuration *kinesisfirehosedeliverystream.ExtendedS3Configuration `hcl:"extended_s3_configuration,block"`
	// HttpEndpointConfiguration: optional
	HttpEndpointConfiguration *kinesisfirehosedeliverystream.HttpEndpointConfiguration `hcl:"http_endpoint_configuration,block"`
	// KinesisSourceConfiguration: optional
	KinesisSourceConfiguration *kinesisfirehosedeliverystream.KinesisSourceConfiguration `hcl:"kinesis_source_configuration,block"`
	// OpensearchConfiguration: optional
	OpensearchConfiguration *kinesisfirehosedeliverystream.OpensearchConfiguration `hcl:"opensearch_configuration,block"`
	// RedshiftConfiguration: optional
	RedshiftConfiguration *kinesisfirehosedeliverystream.RedshiftConfiguration `hcl:"redshift_configuration,block"`
	// ServerSideEncryption: optional
	ServerSideEncryption *kinesisfirehosedeliverystream.ServerSideEncryption `hcl:"server_side_encryption,block"`
	// SplunkConfiguration: optional
	SplunkConfiguration *kinesisfirehosedeliverystream.SplunkConfiguration `hcl:"splunk_configuration,block"`
	// Timeouts: optional
	Timeouts *kinesisfirehosedeliverystream.Timeouts `hcl:"timeouts,block"`
}
type kinesisFirehoseDeliveryStreamAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kinesis_firehose_delivery_stream.
func (kfds kinesisFirehoseDeliveryStreamAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(kfds.ref.Append("arn"))
}

// Destination returns a reference to field destination of aws_kinesis_firehose_delivery_stream.
func (kfds kinesisFirehoseDeliveryStreamAttributes) Destination() terra.StringValue {
	return terra.ReferenceAsString(kfds.ref.Append("destination"))
}

// DestinationId returns a reference to field destination_id of aws_kinesis_firehose_delivery_stream.
func (kfds kinesisFirehoseDeliveryStreamAttributes) DestinationId() terra.StringValue {
	return terra.ReferenceAsString(kfds.ref.Append("destination_id"))
}

// Id returns a reference to field id of aws_kinesis_firehose_delivery_stream.
func (kfds kinesisFirehoseDeliveryStreamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kfds.ref.Append("id"))
}

// Name returns a reference to field name of aws_kinesis_firehose_delivery_stream.
func (kfds kinesisFirehoseDeliveryStreamAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kfds.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_kinesis_firehose_delivery_stream.
func (kfds kinesisFirehoseDeliveryStreamAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kfds.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kinesis_firehose_delivery_stream.
func (kfds kinesisFirehoseDeliveryStreamAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kfds.ref.Append("tags_all"))
}

// VersionId returns a reference to field version_id of aws_kinesis_firehose_delivery_stream.
func (kfds kinesisFirehoseDeliveryStreamAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(kfds.ref.Append("version_id"))
}

func (kfds kinesisFirehoseDeliveryStreamAttributes) ElasticsearchConfiguration() terra.ListValue[kinesisfirehosedeliverystream.ElasticsearchConfigurationAttributes] {
	return terra.ReferenceAsList[kinesisfirehosedeliverystream.ElasticsearchConfigurationAttributes](kfds.ref.Append("elasticsearch_configuration"))
}

func (kfds kinesisFirehoseDeliveryStreamAttributes) ExtendedS3Configuration() terra.ListValue[kinesisfirehosedeliverystream.ExtendedS3ConfigurationAttributes] {
	return terra.ReferenceAsList[kinesisfirehosedeliverystream.ExtendedS3ConfigurationAttributes](kfds.ref.Append("extended_s3_configuration"))
}

func (kfds kinesisFirehoseDeliveryStreamAttributes) HttpEndpointConfiguration() terra.ListValue[kinesisfirehosedeliverystream.HttpEndpointConfigurationAttributes] {
	return terra.ReferenceAsList[kinesisfirehosedeliverystream.HttpEndpointConfigurationAttributes](kfds.ref.Append("http_endpoint_configuration"))
}

func (kfds kinesisFirehoseDeliveryStreamAttributes) KinesisSourceConfiguration() terra.ListValue[kinesisfirehosedeliverystream.KinesisSourceConfigurationAttributes] {
	return terra.ReferenceAsList[kinesisfirehosedeliverystream.KinesisSourceConfigurationAttributes](kfds.ref.Append("kinesis_source_configuration"))
}

func (kfds kinesisFirehoseDeliveryStreamAttributes) OpensearchConfiguration() terra.ListValue[kinesisfirehosedeliverystream.OpensearchConfigurationAttributes] {
	return terra.ReferenceAsList[kinesisfirehosedeliverystream.OpensearchConfigurationAttributes](kfds.ref.Append("opensearch_configuration"))
}

func (kfds kinesisFirehoseDeliveryStreamAttributes) RedshiftConfiguration() terra.ListValue[kinesisfirehosedeliverystream.RedshiftConfigurationAttributes] {
	return terra.ReferenceAsList[kinesisfirehosedeliverystream.RedshiftConfigurationAttributes](kfds.ref.Append("redshift_configuration"))
}

func (kfds kinesisFirehoseDeliveryStreamAttributes) ServerSideEncryption() terra.ListValue[kinesisfirehosedeliverystream.ServerSideEncryptionAttributes] {
	return terra.ReferenceAsList[kinesisfirehosedeliverystream.ServerSideEncryptionAttributes](kfds.ref.Append("server_side_encryption"))
}

func (kfds kinesisFirehoseDeliveryStreamAttributes) SplunkConfiguration() terra.ListValue[kinesisfirehosedeliverystream.SplunkConfigurationAttributes] {
	return terra.ReferenceAsList[kinesisfirehosedeliverystream.SplunkConfigurationAttributes](kfds.ref.Append("splunk_configuration"))
}

func (kfds kinesisFirehoseDeliveryStreamAttributes) Timeouts() kinesisfirehosedeliverystream.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kinesisfirehosedeliverystream.TimeoutsAttributes](kfds.ref.Append("timeouts"))
}

type kinesisFirehoseDeliveryStreamState struct {
	Arn                        string                                                          `json:"arn"`
	Destination                string                                                          `json:"destination"`
	DestinationId              string                                                          `json:"destination_id"`
	Id                         string                                                          `json:"id"`
	Name                       string                                                          `json:"name"`
	Tags                       map[string]string                                               `json:"tags"`
	TagsAll                    map[string]string                                               `json:"tags_all"`
	VersionId                  string                                                          `json:"version_id"`
	ElasticsearchConfiguration []kinesisfirehosedeliverystream.ElasticsearchConfigurationState `json:"elasticsearch_configuration"`
	ExtendedS3Configuration    []kinesisfirehosedeliverystream.ExtendedS3ConfigurationState    `json:"extended_s3_configuration"`
	HttpEndpointConfiguration  []kinesisfirehosedeliverystream.HttpEndpointConfigurationState  `json:"http_endpoint_configuration"`
	KinesisSourceConfiguration []kinesisfirehosedeliverystream.KinesisSourceConfigurationState `json:"kinesis_source_configuration"`
	OpensearchConfiguration    []kinesisfirehosedeliverystream.OpensearchConfigurationState    `json:"opensearch_configuration"`
	RedshiftConfiguration      []kinesisfirehosedeliverystream.RedshiftConfigurationState      `json:"redshift_configuration"`
	ServerSideEncryption       []kinesisfirehosedeliverystream.ServerSideEncryptionState       `json:"server_side_encryption"`
	SplunkConfiguration        []kinesisfirehosedeliverystream.SplunkConfigurationState        `json:"splunk_configuration"`
	Timeouts                   *kinesisfirehosedeliverystream.TimeoutsState                    `json:"timeouts"`
}
