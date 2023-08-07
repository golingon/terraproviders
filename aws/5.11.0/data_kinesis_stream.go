// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datakinesisstream "github.com/golingon/terraproviders/aws/5.11.0/datakinesisstream"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKinesisStream creates a new instance of [DataKinesisStream].
func NewDataKinesisStream(name string, args DataKinesisStreamArgs) *DataKinesisStream {
	return &DataKinesisStream{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKinesisStream)(nil)

// DataKinesisStream represents the Terraform data resource aws_kinesis_stream.
type DataKinesisStream struct {
	Name string
	Args DataKinesisStreamArgs
}

// DataSource returns the Terraform object type for [DataKinesisStream].
func (ks *DataKinesisStream) DataSource() string {
	return "aws_kinesis_stream"
}

// LocalName returns the local name for [DataKinesisStream].
func (ks *DataKinesisStream) LocalName() string {
	return ks.Name
}

// Configuration returns the configuration (args) for [DataKinesisStream].
func (ks *DataKinesisStream) Configuration() interface{} {
	return ks.Args
}

// Attributes returns the attributes for [DataKinesisStream].
func (ks *DataKinesisStream) Attributes() dataKinesisStreamAttributes {
	return dataKinesisStreamAttributes{ref: terra.ReferenceDataResource(ks)}
}

// DataKinesisStreamArgs contains the configurations for aws_kinesis_stream.
type DataKinesisStreamArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// StreamModeDetails: min=0
	StreamModeDetails []datakinesisstream.StreamModeDetails `hcl:"stream_mode_details,block" validate:"min=0"`
}
type dataKinesisStreamAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kinesis_stream.
func (ks dataKinesisStreamAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("arn"))
}

// ClosedShards returns a reference to field closed_shards of aws_kinesis_stream.
func (ks dataKinesisStreamAttributes) ClosedShards() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ks.ref.Append("closed_shards"))
}

// CreationTimestamp returns a reference to field creation_timestamp of aws_kinesis_stream.
func (ks dataKinesisStreamAttributes) CreationTimestamp() terra.NumberValue {
	return terra.ReferenceAsNumber(ks.ref.Append("creation_timestamp"))
}

// Id returns a reference to field id of aws_kinesis_stream.
func (ks dataKinesisStreamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("id"))
}

// Name returns a reference to field name of aws_kinesis_stream.
func (ks dataKinesisStreamAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("name"))
}

// OpenShards returns a reference to field open_shards of aws_kinesis_stream.
func (ks dataKinesisStreamAttributes) OpenShards() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ks.ref.Append("open_shards"))
}

// RetentionPeriod returns a reference to field retention_period of aws_kinesis_stream.
func (ks dataKinesisStreamAttributes) RetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(ks.ref.Append("retention_period"))
}

// ShardLevelMetrics returns a reference to field shard_level_metrics of aws_kinesis_stream.
func (ks dataKinesisStreamAttributes) ShardLevelMetrics() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ks.ref.Append("shard_level_metrics"))
}

// Status returns a reference to field status of aws_kinesis_stream.
func (ks dataKinesisStreamAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_kinesis_stream.
func (ks dataKinesisStreamAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ks.ref.Append("tags"))
}

func (ks dataKinesisStreamAttributes) StreamModeDetails() terra.ListValue[datakinesisstream.StreamModeDetailsAttributes] {
	return terra.ReferenceAsList[datakinesisstream.StreamModeDetailsAttributes](ks.ref.Append("stream_mode_details"))
}
