// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ivschatloggingconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DestinationConfiguration struct {
	// CloudwatchLogs: optional
	CloudwatchLogs *CloudwatchLogs `hcl:"cloudwatch_logs,block"`
	// Firehose: optional
	Firehose *Firehose `hcl:"firehose,block"`
	// S3: optional
	S3 *S3 `hcl:"s3,block"`
}

type CloudwatchLogs struct {
	// LogGroupName: string, required
	LogGroupName terra.StringValue `hcl:"log_group_name,attr" validate:"required"`
}

type Firehose struct {
	// DeliveryStreamName: string, required
	DeliveryStreamName terra.StringValue `hcl:"delivery_stream_name,attr" validate:"required"`
}

type S3 struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DestinationConfigurationAttributes struct {
	ref terra.Reference
}

func (dc DestinationConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DestinationConfigurationAttributes) InternalWithRef(ref terra.Reference) DestinationConfigurationAttributes {
	return DestinationConfigurationAttributes{ref: ref}
}

func (dc DestinationConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DestinationConfigurationAttributes) CloudwatchLogs() terra.ListValue[CloudwatchLogsAttributes] {
	return terra.ReferenceAsList[CloudwatchLogsAttributes](dc.ref.Append("cloudwatch_logs"))
}

func (dc DestinationConfigurationAttributes) Firehose() terra.ListValue[FirehoseAttributes] {
	return terra.ReferenceAsList[FirehoseAttributes](dc.ref.Append("firehose"))
}

func (dc DestinationConfigurationAttributes) S3() terra.ListValue[S3Attributes] {
	return terra.ReferenceAsList[S3Attributes](dc.ref.Append("s3"))
}

type CloudwatchLogsAttributes struct {
	ref terra.Reference
}

func (cl CloudwatchLogsAttributes) InternalRef() (terra.Reference, error) {
	return cl.ref, nil
}

func (cl CloudwatchLogsAttributes) InternalWithRef(ref terra.Reference) CloudwatchLogsAttributes {
	return CloudwatchLogsAttributes{ref: ref}
}

func (cl CloudwatchLogsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cl.ref.InternalTokens()
}

func (cl CloudwatchLogsAttributes) LogGroupName() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("log_group_name"))
}

type FirehoseAttributes struct {
	ref terra.Reference
}

func (f FirehoseAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FirehoseAttributes) InternalWithRef(ref terra.Reference) FirehoseAttributes {
	return FirehoseAttributes{ref: ref}
}

func (f FirehoseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FirehoseAttributes) DeliveryStreamName() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("delivery_stream_name"))
}

type S3Attributes struct {
	ref terra.Reference
}

func (s S3Attributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s S3Attributes) InternalWithRef(ref terra.Reference) S3Attributes {
	return S3Attributes{ref: ref}
}

func (s S3Attributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s S3Attributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("bucket_name"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type DestinationConfigurationState struct {
	CloudwatchLogs []CloudwatchLogsState `json:"cloudwatch_logs"`
	Firehose       []FirehoseState       `json:"firehose"`
	S3             []S3State             `json:"s3"`
}

type CloudwatchLogsState struct {
	LogGroupName string `json:"log_group_name"`
}

type FirehoseState struct {
	DeliveryStreamName string `json:"delivery_stream_name"`
}

type S3State struct {
	BucketName string `json:"bucket_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
