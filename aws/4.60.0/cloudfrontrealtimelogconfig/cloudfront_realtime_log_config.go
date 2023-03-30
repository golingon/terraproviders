// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudfrontrealtimelogconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Endpoint struct {
	// StreamType: string, required
	StreamType terra.StringValue `hcl:"stream_type,attr" validate:"required"`
	// KinesisStreamConfig: required
	KinesisStreamConfig *KinesisStreamConfig `hcl:"kinesis_stream_config,block" validate:"required"`
}

type KinesisStreamConfig struct {
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// StreamArn: string, required
	StreamArn terra.StringValue `hcl:"stream_arn,attr" validate:"required"`
}

type EndpointAttributes struct {
	ref terra.Reference
}

func (e EndpointAttributes) InternalRef() terra.Reference {
	return e.ref
}

func (e EndpointAttributes) InternalWithRef(ref terra.Reference) EndpointAttributes {
	return EndpointAttributes{ref: ref}
}

func (e EndpointAttributes) InternalTokens() hclwrite.Tokens {
	return e.ref.InternalTokens()
}

func (e EndpointAttributes) StreamType() terra.StringValue {
	return terra.ReferenceString(e.ref.Append("stream_type"))
}

func (e EndpointAttributes) KinesisStreamConfig() terra.ListValue[KinesisStreamConfigAttributes] {
	return terra.ReferenceList[KinesisStreamConfigAttributes](e.ref.Append("kinesis_stream_config"))
}

type KinesisStreamConfigAttributes struct {
	ref terra.Reference
}

func (ksc KinesisStreamConfigAttributes) InternalRef() terra.Reference {
	return ksc.ref
}

func (ksc KinesisStreamConfigAttributes) InternalWithRef(ref terra.Reference) KinesisStreamConfigAttributes {
	return KinesisStreamConfigAttributes{ref: ref}
}

func (ksc KinesisStreamConfigAttributes) InternalTokens() hclwrite.Tokens {
	return ksc.ref.InternalTokens()
}

func (ksc KinesisStreamConfigAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceString(ksc.ref.Append("role_arn"))
}

func (ksc KinesisStreamConfigAttributes) StreamArn() terra.StringValue {
	return terra.ReferenceString(ksc.ref.Append("stream_arn"))
}

type EndpointState struct {
	StreamType          string                     `json:"stream_type"`
	KinesisStreamConfig []KinesisStreamConfigState `json:"kinesis_stream_config"`
}

type KinesisStreamConfigState struct {
	RoleArn   string `json:"role_arn"`
	StreamArn string `json:"stream_arn"`
}
