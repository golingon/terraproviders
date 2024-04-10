// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataIvsStreamKey creates a new instance of [DataIvsStreamKey].
func NewDataIvsStreamKey(name string, args DataIvsStreamKeyArgs) *DataIvsStreamKey {
	return &DataIvsStreamKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIvsStreamKey)(nil)

// DataIvsStreamKey represents the Terraform data resource aws_ivs_stream_key.
type DataIvsStreamKey struct {
	Name string
	Args DataIvsStreamKeyArgs
}

// DataSource returns the Terraform object type for [DataIvsStreamKey].
func (isk *DataIvsStreamKey) DataSource() string {
	return "aws_ivs_stream_key"
}

// LocalName returns the local name for [DataIvsStreamKey].
func (isk *DataIvsStreamKey) LocalName() string {
	return isk.Name
}

// Configuration returns the configuration (args) for [DataIvsStreamKey].
func (isk *DataIvsStreamKey) Configuration() interface{} {
	return isk.Args
}

// Attributes returns the attributes for [DataIvsStreamKey].
func (isk *DataIvsStreamKey) Attributes() dataIvsStreamKeyAttributes {
	return dataIvsStreamKeyAttributes{ref: terra.ReferenceDataResource(isk)}
}

// DataIvsStreamKeyArgs contains the configurations for aws_ivs_stream_key.
type DataIvsStreamKeyArgs struct {
	// ChannelArn: string, required
	ChannelArn terra.StringValue `hcl:"channel_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataIvsStreamKeyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ivs_stream_key.
func (isk dataIvsStreamKeyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(isk.ref.Append("arn"))
}

// ChannelArn returns a reference to field channel_arn of aws_ivs_stream_key.
func (isk dataIvsStreamKeyAttributes) ChannelArn() terra.StringValue {
	return terra.ReferenceAsString(isk.ref.Append("channel_arn"))
}

// Id returns a reference to field id of aws_ivs_stream_key.
func (isk dataIvsStreamKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(isk.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_ivs_stream_key.
func (isk dataIvsStreamKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](isk.ref.Append("tags"))
}

// Value returns a reference to field value of aws_ivs_stream_key.
func (isk dataIvsStreamKeyAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(isk.ref.Append("value"))
}
