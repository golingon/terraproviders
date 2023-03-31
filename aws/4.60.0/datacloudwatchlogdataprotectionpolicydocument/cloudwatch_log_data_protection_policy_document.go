// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacloudwatchlogdataprotectionpolicydocument

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Statement struct {
	// DataIdentifiers: set of string, required
	DataIdentifiers terra.SetValue[terra.StringValue] `hcl:"data_identifiers,attr" validate:"required"`
	// Sid: string, optional
	Sid terra.StringValue `hcl:"sid,attr"`
	// Operation: required
	Operation *Operation `hcl:"operation,block" validate:"required"`
}

type Operation struct {
	// Audit: optional
	Audit *Audit `hcl:"audit,block"`
	// Deidentify: optional
	Deidentify *Deidentify `hcl:"deidentify,block"`
}

type Audit struct {
	// FindingsDestination: required
	FindingsDestination *FindingsDestination `hcl:"findings_destination,block" validate:"required"`
}

type FindingsDestination struct {
	// CloudwatchLogs: optional
	CloudwatchLogs *CloudwatchLogs `hcl:"cloudwatch_logs,block"`
	// Firehose: optional
	Firehose *Firehose `hcl:"firehose,block"`
	// S3: optional
	S3 *S3 `hcl:"s3,block"`
}

type CloudwatchLogs struct {
	// LogGroup: string, required
	LogGroup terra.StringValue `hcl:"log_group,attr" validate:"required"`
}

type Firehose struct {
	// DeliveryStream: string, required
	DeliveryStream terra.StringValue `hcl:"delivery_stream,attr" validate:"required"`
}

type S3 struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
}

type Deidentify struct {
	// MaskConfig: required
	MaskConfig *MaskConfig `hcl:"mask_config,block" validate:"required"`
}

type MaskConfig struct{}

type StatementAttributes struct {
	ref terra.Reference
}

func (s StatementAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s StatementAttributes) InternalWithRef(ref terra.Reference) StatementAttributes {
	return StatementAttributes{ref: ref}
}

func (s StatementAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s StatementAttributes) DataIdentifiers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("data_identifiers"))
}

func (s StatementAttributes) Sid() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("sid"))
}

func (s StatementAttributes) Operation() terra.ListValue[OperationAttributes] {
	return terra.ReferenceAsList[OperationAttributes](s.ref.Append("operation"))
}

type OperationAttributes struct {
	ref terra.Reference
}

func (o OperationAttributes) InternalRef() terra.Reference {
	return o.ref
}

func (o OperationAttributes) InternalWithRef(ref terra.Reference) OperationAttributes {
	return OperationAttributes{ref: ref}
}

func (o OperationAttributes) InternalTokens() hclwrite.Tokens {
	return o.ref.InternalTokens()
}

func (o OperationAttributes) Audit() terra.ListValue[AuditAttributes] {
	return terra.ReferenceAsList[AuditAttributes](o.ref.Append("audit"))
}

func (o OperationAttributes) Deidentify() terra.ListValue[DeidentifyAttributes] {
	return terra.ReferenceAsList[DeidentifyAttributes](o.ref.Append("deidentify"))
}

type AuditAttributes struct {
	ref terra.Reference
}

func (a AuditAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AuditAttributes) InternalWithRef(ref terra.Reference) AuditAttributes {
	return AuditAttributes{ref: ref}
}

func (a AuditAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AuditAttributes) FindingsDestination() terra.ListValue[FindingsDestinationAttributes] {
	return terra.ReferenceAsList[FindingsDestinationAttributes](a.ref.Append("findings_destination"))
}

type FindingsDestinationAttributes struct {
	ref terra.Reference
}

func (fd FindingsDestinationAttributes) InternalRef() terra.Reference {
	return fd.ref
}

func (fd FindingsDestinationAttributes) InternalWithRef(ref terra.Reference) FindingsDestinationAttributes {
	return FindingsDestinationAttributes{ref: ref}
}

func (fd FindingsDestinationAttributes) InternalTokens() hclwrite.Tokens {
	return fd.ref.InternalTokens()
}

func (fd FindingsDestinationAttributes) CloudwatchLogs() terra.ListValue[CloudwatchLogsAttributes] {
	return terra.ReferenceAsList[CloudwatchLogsAttributes](fd.ref.Append("cloudwatch_logs"))
}

func (fd FindingsDestinationAttributes) Firehose() terra.ListValue[FirehoseAttributes] {
	return terra.ReferenceAsList[FirehoseAttributes](fd.ref.Append("firehose"))
}

func (fd FindingsDestinationAttributes) S3() terra.ListValue[S3Attributes] {
	return terra.ReferenceAsList[S3Attributes](fd.ref.Append("s3"))
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

type FirehoseAttributes struct {
	ref terra.Reference
}

func (f FirehoseAttributes) InternalRef() terra.Reference {
	return f.ref
}

func (f FirehoseAttributes) InternalWithRef(ref terra.Reference) FirehoseAttributes {
	return FirehoseAttributes{ref: ref}
}

func (f FirehoseAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FirehoseAttributes) DeliveryStream() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("delivery_stream"))
}

type S3Attributes struct {
	ref terra.Reference
}

func (s S3Attributes) InternalRef() terra.Reference {
	return s.ref
}

func (s S3Attributes) InternalWithRef(ref terra.Reference) S3Attributes {
	return S3Attributes{ref: ref}
}

func (s S3Attributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s S3Attributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("bucket"))
}

type DeidentifyAttributes struct {
	ref terra.Reference
}

func (d DeidentifyAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d DeidentifyAttributes) InternalWithRef(ref terra.Reference) DeidentifyAttributes {
	return DeidentifyAttributes{ref: ref}
}

func (d DeidentifyAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d DeidentifyAttributes) MaskConfig() terra.ListValue[MaskConfigAttributes] {
	return terra.ReferenceAsList[MaskConfigAttributes](d.ref.Append("mask_config"))
}

type MaskConfigAttributes struct {
	ref terra.Reference
}

func (mc MaskConfigAttributes) InternalRef() terra.Reference {
	return mc.ref
}

func (mc MaskConfigAttributes) InternalWithRef(ref terra.Reference) MaskConfigAttributes {
	return MaskConfigAttributes{ref: ref}
}

func (mc MaskConfigAttributes) InternalTokens() hclwrite.Tokens {
	return mc.ref.InternalTokens()
}

type StatementState struct {
	DataIdentifiers []string         `json:"data_identifiers"`
	Sid             string           `json:"sid"`
	Operation       []OperationState `json:"operation"`
}

type OperationState struct {
	Audit      []AuditState      `json:"audit"`
	Deidentify []DeidentifyState `json:"deidentify"`
}

type AuditState struct {
	FindingsDestination []FindingsDestinationState `json:"findings_destination"`
}

type FindingsDestinationState struct {
	CloudwatchLogs []CloudwatchLogsState `json:"cloudwatch_logs"`
	Firehose       []FirehoseState       `json:"firehose"`
	S3             []S3State             `json:"s3"`
}

type CloudwatchLogsState struct {
	LogGroup string `json:"log_group"`
}

type FirehoseState struct {
	DeliveryStream string `json:"delivery_stream"`
}

type S3State struct {
	Bucket string `json:"bucket"`
}

type DeidentifyState struct {
	MaskConfig []MaskConfigState `json:"mask_config"`
}

type MaskConfigState struct{}
