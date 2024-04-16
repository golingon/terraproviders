// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudfront_realtime_log_config

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Endpoint struct {
	// StreamType: string, required
	StreamType terra.StringValue `hcl:"stream_type,attr" validate:"required"`
	// EndpointKinesisStreamConfig: required
	KinesisStreamConfig *EndpointKinesisStreamConfig `hcl:"kinesis_stream_config,block" validate:"required"`
}

type EndpointKinesisStreamConfig struct {
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// StreamArn: string, required
	StreamArn terra.StringValue `hcl:"stream_arn,attr" validate:"required"`
}

type EndpointAttributes struct {
	ref terra.Reference
}

func (e EndpointAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EndpointAttributes) InternalWithRef(ref terra.Reference) EndpointAttributes {
	return EndpointAttributes{ref: ref}
}

func (e EndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EndpointAttributes) StreamType() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("stream_type"))
}

func (e EndpointAttributes) KinesisStreamConfig() terra.ListValue[EndpointKinesisStreamConfigAttributes] {
	return terra.ReferenceAsList[EndpointKinesisStreamConfigAttributes](e.ref.Append("kinesis_stream_config"))
}

type EndpointKinesisStreamConfigAttributes struct {
	ref terra.Reference
}

func (ksc EndpointKinesisStreamConfigAttributes) InternalRef() (terra.Reference, error) {
	return ksc.ref, nil
}

func (ksc EndpointKinesisStreamConfigAttributes) InternalWithRef(ref terra.Reference) EndpointKinesisStreamConfigAttributes {
	return EndpointKinesisStreamConfigAttributes{ref: ref}
}

func (ksc EndpointKinesisStreamConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ksc.ref.InternalTokens()
}

func (ksc EndpointKinesisStreamConfigAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("role_arn"))
}

func (ksc EndpointKinesisStreamConfigAttributes) StreamArn() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("stream_arn"))
}

type EndpointState struct {
	StreamType          string                             `json:"stream_type"`
	KinesisStreamConfig []EndpointKinesisStreamConfigState `json:"kinesis_stream_config"`
}

type EndpointKinesisStreamConfigState struct {
	RoleArn   string `json:"role_arn"`
	StreamArn string `json:"stream_arn"`
}
