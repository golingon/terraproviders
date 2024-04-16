// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_kinesis_firehose_delivery_stream

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_kinesis_firehose_delivery_stream.
type Resource struct {
	Name      string
	Args      Args
	state     *awsKinesisFirehoseDeliveryStreamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (akfds *Resource) Type() string {
	return "aws_kinesis_firehose_delivery_stream"
}

// LocalName returns the local name for [Resource].
func (akfds *Resource) LocalName() string {
	return akfds.Name
}

// Configuration returns the configuration (args) for [Resource].
func (akfds *Resource) Configuration() interface{} {
	return akfds.Args
}

// DependOn is used for other resources to depend on [Resource].
func (akfds *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(akfds)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (akfds *Resource) Dependencies() terra.Dependencies {
	return akfds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (akfds *Resource) LifecycleManagement() *terra.Lifecycle {
	return akfds.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (akfds *Resource) Attributes() awsKinesisFirehoseDeliveryStreamAttributes {
	return awsKinesisFirehoseDeliveryStreamAttributes{ref: terra.ReferenceResource(akfds)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (akfds *Resource) ImportState(state io.Reader) error {
	akfds.state = &awsKinesisFirehoseDeliveryStreamState{}
	if err := json.NewDecoder(state).Decode(akfds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akfds.Type(), akfds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (akfds *Resource) State() (*awsKinesisFirehoseDeliveryStreamState, bool) {
	return akfds.state, akfds.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (akfds *Resource) StateMust() *awsKinesisFirehoseDeliveryStreamState {
	if akfds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akfds.Type(), akfds.LocalName()))
	}
	return akfds.state
}

// Args contains the configurations for aws_kinesis_firehose_delivery_stream.
type Args struct {
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
	ElasticsearchConfiguration *ElasticsearchConfiguration `hcl:"elasticsearch_configuration,block"`
	// ExtendedS3Configuration: optional
	ExtendedS3Configuration *ExtendedS3Configuration `hcl:"extended_s3_configuration,block"`
	// HttpEndpointConfiguration: optional
	HttpEndpointConfiguration *HttpEndpointConfiguration `hcl:"http_endpoint_configuration,block"`
	// KinesisSourceConfiguration: optional
	KinesisSourceConfiguration *KinesisSourceConfiguration `hcl:"kinesis_source_configuration,block"`
	// MskSourceConfiguration: optional
	MskSourceConfiguration *MskSourceConfiguration `hcl:"msk_source_configuration,block"`
	// OpensearchConfiguration: optional
	OpensearchConfiguration *OpensearchConfiguration `hcl:"opensearch_configuration,block"`
	// OpensearchserverlessConfiguration: optional
	OpensearchserverlessConfiguration *OpensearchserverlessConfiguration `hcl:"opensearchserverless_configuration,block"`
	// RedshiftConfiguration: optional
	RedshiftConfiguration *RedshiftConfiguration `hcl:"redshift_configuration,block"`
	// ServerSideEncryption: optional
	ServerSideEncryption *ServerSideEncryption `hcl:"server_side_encryption,block"`
	// SplunkConfiguration: optional
	SplunkConfiguration *SplunkConfiguration `hcl:"splunk_configuration,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsKinesisFirehoseDeliveryStreamAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kinesis_firehose_delivery_stream.
func (akfds awsKinesisFirehoseDeliveryStreamAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(akfds.ref.Append("arn"))
}

// Destination returns a reference to field destination of aws_kinesis_firehose_delivery_stream.
func (akfds awsKinesisFirehoseDeliveryStreamAttributes) Destination() terra.StringValue {
	return terra.ReferenceAsString(akfds.ref.Append("destination"))
}

// DestinationId returns a reference to field destination_id of aws_kinesis_firehose_delivery_stream.
func (akfds awsKinesisFirehoseDeliveryStreamAttributes) DestinationId() terra.StringValue {
	return terra.ReferenceAsString(akfds.ref.Append("destination_id"))
}

// Id returns a reference to field id of aws_kinesis_firehose_delivery_stream.
func (akfds awsKinesisFirehoseDeliveryStreamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akfds.ref.Append("id"))
}

// Name returns a reference to field name of aws_kinesis_firehose_delivery_stream.
func (akfds awsKinesisFirehoseDeliveryStreamAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akfds.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_kinesis_firehose_delivery_stream.
func (akfds awsKinesisFirehoseDeliveryStreamAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akfds.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kinesis_firehose_delivery_stream.
func (akfds awsKinesisFirehoseDeliveryStreamAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akfds.ref.Append("tags_all"))
}

// VersionId returns a reference to field version_id of aws_kinesis_firehose_delivery_stream.
func (akfds awsKinesisFirehoseDeliveryStreamAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(akfds.ref.Append("version_id"))
}

func (akfds awsKinesisFirehoseDeliveryStreamAttributes) ElasticsearchConfiguration() terra.ListValue[ElasticsearchConfigurationAttributes] {
	return terra.ReferenceAsList[ElasticsearchConfigurationAttributes](akfds.ref.Append("elasticsearch_configuration"))
}

func (akfds awsKinesisFirehoseDeliveryStreamAttributes) ExtendedS3Configuration() terra.ListValue[ExtendedS3ConfigurationAttributes] {
	return terra.ReferenceAsList[ExtendedS3ConfigurationAttributes](akfds.ref.Append("extended_s3_configuration"))
}

func (akfds awsKinesisFirehoseDeliveryStreamAttributes) HttpEndpointConfiguration() terra.ListValue[HttpEndpointConfigurationAttributes] {
	return terra.ReferenceAsList[HttpEndpointConfigurationAttributes](akfds.ref.Append("http_endpoint_configuration"))
}

func (akfds awsKinesisFirehoseDeliveryStreamAttributes) KinesisSourceConfiguration() terra.ListValue[KinesisSourceConfigurationAttributes] {
	return terra.ReferenceAsList[KinesisSourceConfigurationAttributes](akfds.ref.Append("kinesis_source_configuration"))
}

func (akfds awsKinesisFirehoseDeliveryStreamAttributes) MskSourceConfiguration() terra.ListValue[MskSourceConfigurationAttributes] {
	return terra.ReferenceAsList[MskSourceConfigurationAttributes](akfds.ref.Append("msk_source_configuration"))
}

func (akfds awsKinesisFirehoseDeliveryStreamAttributes) OpensearchConfiguration() terra.ListValue[OpensearchConfigurationAttributes] {
	return terra.ReferenceAsList[OpensearchConfigurationAttributes](akfds.ref.Append("opensearch_configuration"))
}

func (akfds awsKinesisFirehoseDeliveryStreamAttributes) OpensearchserverlessConfiguration() terra.ListValue[OpensearchserverlessConfigurationAttributes] {
	return terra.ReferenceAsList[OpensearchserverlessConfigurationAttributes](akfds.ref.Append("opensearchserverless_configuration"))
}

func (akfds awsKinesisFirehoseDeliveryStreamAttributes) RedshiftConfiguration() terra.ListValue[RedshiftConfigurationAttributes] {
	return terra.ReferenceAsList[RedshiftConfigurationAttributes](akfds.ref.Append("redshift_configuration"))
}

func (akfds awsKinesisFirehoseDeliveryStreamAttributes) ServerSideEncryption() terra.ListValue[ServerSideEncryptionAttributes] {
	return terra.ReferenceAsList[ServerSideEncryptionAttributes](akfds.ref.Append("server_side_encryption"))
}

func (akfds awsKinesisFirehoseDeliveryStreamAttributes) SplunkConfiguration() terra.ListValue[SplunkConfigurationAttributes] {
	return terra.ReferenceAsList[SplunkConfigurationAttributes](akfds.ref.Append("splunk_configuration"))
}

func (akfds awsKinesisFirehoseDeliveryStreamAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](akfds.ref.Append("timeouts"))
}

type awsKinesisFirehoseDeliveryStreamState struct {
	Arn                               string                                   `json:"arn"`
	Destination                       string                                   `json:"destination"`
	DestinationId                     string                                   `json:"destination_id"`
	Id                                string                                   `json:"id"`
	Name                              string                                   `json:"name"`
	Tags                              map[string]string                        `json:"tags"`
	TagsAll                           map[string]string                        `json:"tags_all"`
	VersionId                         string                                   `json:"version_id"`
	ElasticsearchConfiguration        []ElasticsearchConfigurationState        `json:"elasticsearch_configuration"`
	ExtendedS3Configuration           []ExtendedS3ConfigurationState           `json:"extended_s3_configuration"`
	HttpEndpointConfiguration         []HttpEndpointConfigurationState         `json:"http_endpoint_configuration"`
	KinesisSourceConfiguration        []KinesisSourceConfigurationState        `json:"kinesis_source_configuration"`
	MskSourceConfiguration            []MskSourceConfigurationState            `json:"msk_source_configuration"`
	OpensearchConfiguration           []OpensearchConfigurationState           `json:"opensearch_configuration"`
	OpensearchserverlessConfiguration []OpensearchserverlessConfigurationState `json:"opensearchserverless_configuration"`
	RedshiftConfiguration             []RedshiftConfigurationState             `json:"redshift_configuration"`
	ServerSideEncryption              []ServerSideEncryptionState              `json:"server_side_encryption"`
	SplunkConfiguration               []SplunkConfigurationState               `json:"splunk_configuration"`
	Timeouts                          *TimeoutsState                           `json:"timeouts"`
}
