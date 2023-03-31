// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package evidentlyproject

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DataDelivery struct {
	// CloudwatchLogs: optional
	CloudwatchLogs *CloudwatchLogs `hcl:"cloudwatch_logs,block"`
	// S3Destination: optional
	S3Destination *S3Destination `hcl:"s3_destination,block"`
}

type CloudwatchLogs struct {
	// LogGroup: string, optional
	LogGroup terra.StringValue `hcl:"log_group,attr"`
}

type S3Destination struct {
	// Bucket: string, optional
	Bucket terra.StringValue `hcl:"bucket,attr"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DataDeliveryAttributes struct {
	ref terra.Reference
}

func (dd DataDeliveryAttributes) InternalRef() terra.Reference {
	return dd.ref
}

func (dd DataDeliveryAttributes) InternalWithRef(ref terra.Reference) DataDeliveryAttributes {
	return DataDeliveryAttributes{ref: ref}
}

func (dd DataDeliveryAttributes) InternalTokens() hclwrite.Tokens {
	return dd.ref.InternalTokens()
}

func (dd DataDeliveryAttributes) CloudwatchLogs() terra.ListValue[CloudwatchLogsAttributes] {
	return terra.ReferenceAsList[CloudwatchLogsAttributes](dd.ref.Append("cloudwatch_logs"))
}

func (dd DataDeliveryAttributes) S3Destination() terra.ListValue[S3DestinationAttributes] {
	return terra.ReferenceAsList[S3DestinationAttributes](dd.ref.Append("s3_destination"))
}

type CloudwatchLogsAttributes struct {
	ref terra.Reference
}

func (cl CloudwatchLogsAttributes) InternalRef() terra.Reference {
	return cl.ref
}

func (cl CloudwatchLogsAttributes) InternalWithRef(ref terra.Reference) CloudwatchLogsAttributes {
	return CloudwatchLogsAttributes{ref: ref}
}

func (cl CloudwatchLogsAttributes) InternalTokens() hclwrite.Tokens {
	return cl.ref.InternalTokens()
}

func (cl CloudwatchLogsAttributes) LogGroup() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("log_group"))
}

type S3DestinationAttributes struct {
	ref terra.Reference
}

func (sd S3DestinationAttributes) InternalRef() terra.Reference {
	return sd.ref
}

func (sd S3DestinationAttributes) InternalWithRef(ref terra.Reference) S3DestinationAttributes {
	return S3DestinationAttributes{ref: ref}
}

func (sd S3DestinationAttributes) InternalTokens() hclwrite.Tokens {
	return sd.ref.InternalTokens()
}

func (sd S3DestinationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("bucket"))
}

func (sd S3DestinationAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("prefix"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
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

type DataDeliveryState struct {
	CloudwatchLogs []CloudwatchLogsState `json:"cloudwatch_logs"`
	S3Destination  []S3DestinationState  `json:"s3_destination"`
}

type CloudwatchLogsState struct {
	LogGroup string `json:"log_group"`
}

type S3DestinationState struct {
	Bucket string `json:"bucket"`
	Prefix string `json:"prefix"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
