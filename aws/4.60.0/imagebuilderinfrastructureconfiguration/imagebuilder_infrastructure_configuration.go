// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package imagebuilderinfrastructureconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type InstanceMetadataOptions struct {
	// HttpPutResponseHopLimit: number, optional
	HttpPutResponseHopLimit terra.NumberValue `hcl:"http_put_response_hop_limit,attr"`
	// HttpTokens: string, optional
	HttpTokens terra.StringValue `hcl:"http_tokens,attr"`
}

type Logging struct {
	// S3Logs: required
	S3Logs *S3Logs `hcl:"s3_logs,block" validate:"required"`
}

type S3Logs struct {
	// S3BucketName: string, required
	S3BucketName terra.StringValue `hcl:"s3_bucket_name,attr" validate:"required"`
	// S3KeyPrefix: string, optional
	S3KeyPrefix terra.StringValue `hcl:"s3_key_prefix,attr"`
}

type InstanceMetadataOptionsAttributes struct {
	ref terra.Reference
}

func (imo InstanceMetadataOptionsAttributes) InternalRef() (terra.Reference, error) {
	return imo.ref, nil
}

func (imo InstanceMetadataOptionsAttributes) InternalWithRef(ref terra.Reference) InstanceMetadataOptionsAttributes {
	return InstanceMetadataOptionsAttributes{ref: ref}
}

func (imo InstanceMetadataOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return imo.ref.InternalTokens()
}

func (imo InstanceMetadataOptionsAttributes) HttpPutResponseHopLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(imo.ref.Append("http_put_response_hop_limit"))
}

func (imo InstanceMetadataOptionsAttributes) HttpTokens() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("http_tokens"))
}

type LoggingAttributes struct {
	ref terra.Reference
}

func (l LoggingAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LoggingAttributes) InternalWithRef(ref terra.Reference) LoggingAttributes {
	return LoggingAttributes{ref: ref}
}

func (l LoggingAttributes) InternalTokens() hclwrite.Tokens {
	return l.ref.InternalTokens()
}

func (l LoggingAttributes) S3Logs() terra.ListValue[S3LogsAttributes] {
	return terra.ReferenceAsList[S3LogsAttributes](l.ref.Append("s3_logs"))
}

type S3LogsAttributes struct {
	ref terra.Reference
}

func (sl S3LogsAttributes) InternalRef() (terra.Reference, error) {
	return sl.ref, nil
}

func (sl S3LogsAttributes) InternalWithRef(ref terra.Reference) S3LogsAttributes {
	return S3LogsAttributes{ref: ref}
}

func (sl S3LogsAttributes) InternalTokens() hclwrite.Tokens {
	return sl.ref.InternalTokens()
}

func (sl S3LogsAttributes) S3BucketName() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("s3_bucket_name"))
}

func (sl S3LogsAttributes) S3KeyPrefix() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("s3_key_prefix"))
}

type InstanceMetadataOptionsState struct {
	HttpPutResponseHopLimit float64 `json:"http_put_response_hop_limit"`
	HttpTokens              string  `json:"http_tokens"`
}

type LoggingState struct {
	S3Logs []S3LogsState `json:"s3_logs"`
}

type S3LogsState struct {
	S3BucketName string `json:"s3_bucket_name"`
	S3KeyPrefix  string `json:"s3_key_prefix"`
}
