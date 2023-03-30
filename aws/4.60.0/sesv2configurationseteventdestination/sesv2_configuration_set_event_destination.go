// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sesv2configurationseteventdestination

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EventDestination struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// MatchingEventTypes: list of string, required
	MatchingEventTypes terra.ListValue[terra.StringValue] `hcl:"matching_event_types,attr" validate:"required"`
	// CloudWatchDestination: optional
	CloudWatchDestination *CloudWatchDestination `hcl:"cloud_watch_destination,block"`
	// KinesisFirehoseDestination: optional
	KinesisFirehoseDestination *KinesisFirehoseDestination `hcl:"kinesis_firehose_destination,block"`
	// PinpointDestination: optional
	PinpointDestination *PinpointDestination `hcl:"pinpoint_destination,block"`
	// SnsDestination: optional
	SnsDestination *SnsDestination `hcl:"sns_destination,block"`
}

type CloudWatchDestination struct {
	// DimensionConfiguration: min=1
	DimensionConfiguration []DimensionConfiguration `hcl:"dimension_configuration,block" validate:"min=1"`
}

type DimensionConfiguration struct {
	// DefaultDimensionValue: string, required
	DefaultDimensionValue terra.StringValue `hcl:"default_dimension_value,attr" validate:"required"`
	// DimensionName: string, required
	DimensionName terra.StringValue `hcl:"dimension_name,attr" validate:"required"`
	// DimensionValueSource: string, required
	DimensionValueSource terra.StringValue `hcl:"dimension_value_source,attr" validate:"required"`
}

type KinesisFirehoseDestination struct {
	// DeliveryStreamArn: string, required
	DeliveryStreamArn terra.StringValue `hcl:"delivery_stream_arn,attr" validate:"required"`
	// IamRoleArn: string, required
	IamRoleArn terra.StringValue `hcl:"iam_role_arn,attr" validate:"required"`
}

type PinpointDestination struct {
	// ApplicationArn: string, required
	ApplicationArn terra.StringValue `hcl:"application_arn,attr" validate:"required"`
}

type SnsDestination struct {
	// TopicArn: string, required
	TopicArn terra.StringValue `hcl:"topic_arn,attr" validate:"required"`
}

type EventDestinationAttributes struct {
	ref terra.Reference
}

func (ed EventDestinationAttributes) InternalRef() terra.Reference {
	return ed.ref
}

func (ed EventDestinationAttributes) InternalWithRef(ref terra.Reference) EventDestinationAttributes {
	return EventDestinationAttributes{ref: ref}
}

func (ed EventDestinationAttributes) InternalTokens() hclwrite.Tokens {
	return ed.ref.InternalTokens()
}

func (ed EventDestinationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(ed.ref.Append("enabled"))
}

func (ed EventDestinationAttributes) MatchingEventTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ed.ref.Append("matching_event_types"))
}

func (ed EventDestinationAttributes) CloudWatchDestination() terra.ListValue[CloudWatchDestinationAttributes] {
	return terra.ReferenceList[CloudWatchDestinationAttributes](ed.ref.Append("cloud_watch_destination"))
}

func (ed EventDestinationAttributes) KinesisFirehoseDestination() terra.ListValue[KinesisFirehoseDestinationAttributes] {
	return terra.ReferenceList[KinesisFirehoseDestinationAttributes](ed.ref.Append("kinesis_firehose_destination"))
}

func (ed EventDestinationAttributes) PinpointDestination() terra.ListValue[PinpointDestinationAttributes] {
	return terra.ReferenceList[PinpointDestinationAttributes](ed.ref.Append("pinpoint_destination"))
}

func (ed EventDestinationAttributes) SnsDestination() terra.ListValue[SnsDestinationAttributes] {
	return terra.ReferenceList[SnsDestinationAttributes](ed.ref.Append("sns_destination"))
}

type CloudWatchDestinationAttributes struct {
	ref terra.Reference
}

func (cwd CloudWatchDestinationAttributes) InternalRef() terra.Reference {
	return cwd.ref
}

func (cwd CloudWatchDestinationAttributes) InternalWithRef(ref terra.Reference) CloudWatchDestinationAttributes {
	return CloudWatchDestinationAttributes{ref: ref}
}

func (cwd CloudWatchDestinationAttributes) InternalTokens() hclwrite.Tokens {
	return cwd.ref.InternalTokens()
}

func (cwd CloudWatchDestinationAttributes) DimensionConfiguration() terra.ListValue[DimensionConfigurationAttributes] {
	return terra.ReferenceList[DimensionConfigurationAttributes](cwd.ref.Append("dimension_configuration"))
}

type DimensionConfigurationAttributes struct {
	ref terra.Reference
}

func (dc DimensionConfigurationAttributes) InternalRef() terra.Reference {
	return dc.ref
}

func (dc DimensionConfigurationAttributes) InternalWithRef(ref terra.Reference) DimensionConfigurationAttributes {
	return DimensionConfigurationAttributes{ref: ref}
}

func (dc DimensionConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return dc.ref.InternalTokens()
}

func (dc DimensionConfigurationAttributes) DefaultDimensionValue() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("default_dimension_value"))
}

func (dc DimensionConfigurationAttributes) DimensionName() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("dimension_name"))
}

func (dc DimensionConfigurationAttributes) DimensionValueSource() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("dimension_value_source"))
}

type KinesisFirehoseDestinationAttributes struct {
	ref terra.Reference
}

func (kfd KinesisFirehoseDestinationAttributes) InternalRef() terra.Reference {
	return kfd.ref
}

func (kfd KinesisFirehoseDestinationAttributes) InternalWithRef(ref terra.Reference) KinesisFirehoseDestinationAttributes {
	return KinesisFirehoseDestinationAttributes{ref: ref}
}

func (kfd KinesisFirehoseDestinationAttributes) InternalTokens() hclwrite.Tokens {
	return kfd.ref.InternalTokens()
}

func (kfd KinesisFirehoseDestinationAttributes) DeliveryStreamArn() terra.StringValue {
	return terra.ReferenceString(kfd.ref.Append("delivery_stream_arn"))
}

func (kfd KinesisFirehoseDestinationAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceString(kfd.ref.Append("iam_role_arn"))
}

type PinpointDestinationAttributes struct {
	ref terra.Reference
}

func (pd PinpointDestinationAttributes) InternalRef() terra.Reference {
	return pd.ref
}

func (pd PinpointDestinationAttributes) InternalWithRef(ref terra.Reference) PinpointDestinationAttributes {
	return PinpointDestinationAttributes{ref: ref}
}

func (pd PinpointDestinationAttributes) InternalTokens() hclwrite.Tokens {
	return pd.ref.InternalTokens()
}

func (pd PinpointDestinationAttributes) ApplicationArn() terra.StringValue {
	return terra.ReferenceString(pd.ref.Append("application_arn"))
}

type SnsDestinationAttributes struct {
	ref terra.Reference
}

func (sd SnsDestinationAttributes) InternalRef() terra.Reference {
	return sd.ref
}

func (sd SnsDestinationAttributes) InternalWithRef(ref terra.Reference) SnsDestinationAttributes {
	return SnsDestinationAttributes{ref: ref}
}

func (sd SnsDestinationAttributes) InternalTokens() hclwrite.Tokens {
	return sd.ref.InternalTokens()
}

func (sd SnsDestinationAttributes) TopicArn() terra.StringValue {
	return terra.ReferenceString(sd.ref.Append("topic_arn"))
}

type EventDestinationState struct {
	Enabled                    bool                              `json:"enabled"`
	MatchingEventTypes         []string                          `json:"matching_event_types"`
	CloudWatchDestination      []CloudWatchDestinationState      `json:"cloud_watch_destination"`
	KinesisFirehoseDestination []KinesisFirehoseDestinationState `json:"kinesis_firehose_destination"`
	PinpointDestination        []PinpointDestinationState        `json:"pinpoint_destination"`
	SnsDestination             []SnsDestinationState             `json:"sns_destination"`
}

type CloudWatchDestinationState struct {
	DimensionConfiguration []DimensionConfigurationState `json:"dimension_configuration"`
}

type DimensionConfigurationState struct {
	DefaultDimensionValue string `json:"default_dimension_value"`
	DimensionName         string `json:"dimension_name"`
	DimensionValueSource  string `json:"dimension_value_source"`
}

type KinesisFirehoseDestinationState struct {
	DeliveryStreamArn string `json:"delivery_stream_arn"`
	IamRoleArn        string `json:"iam_role_arn"`
}

type PinpointDestinationState struct {
	ApplicationArn string `json:"application_arn"`
}

type SnsDestinationState struct {
	TopicArn string `json:"topic_arn"`
}
