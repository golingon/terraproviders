// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_sesv2_configuration_set_event_destination

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type EventDestination struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// MatchingEventTypes: list of string, required
	MatchingEventTypes terra.ListValue[terra.StringValue] `hcl:"matching_event_types,attr" validate:"required"`
	// EventDestinationCloudWatchDestination: optional
	CloudWatchDestination *EventDestinationCloudWatchDestination `hcl:"cloud_watch_destination,block"`
	// EventDestinationKinesisFirehoseDestination: optional
	KinesisFirehoseDestination *EventDestinationKinesisFirehoseDestination `hcl:"kinesis_firehose_destination,block"`
	// EventDestinationPinpointDestination: optional
	PinpointDestination *EventDestinationPinpointDestination `hcl:"pinpoint_destination,block"`
	// EventDestinationSnsDestination: optional
	SnsDestination *EventDestinationSnsDestination `hcl:"sns_destination,block"`
}

type EventDestinationCloudWatchDestination struct {
	// EventDestinationCloudWatchDestinationDimensionConfiguration: min=1
	DimensionConfiguration []EventDestinationCloudWatchDestinationDimensionConfiguration `hcl:"dimension_configuration,block" validate:"min=1"`
}

type EventDestinationCloudWatchDestinationDimensionConfiguration struct {
	// DefaultDimensionValue: string, required
	DefaultDimensionValue terra.StringValue `hcl:"default_dimension_value,attr" validate:"required"`
	// DimensionName: string, required
	DimensionName terra.StringValue `hcl:"dimension_name,attr" validate:"required"`
	// DimensionValueSource: string, required
	DimensionValueSource terra.StringValue `hcl:"dimension_value_source,attr" validate:"required"`
}

type EventDestinationKinesisFirehoseDestination struct {
	// DeliveryStreamArn: string, required
	DeliveryStreamArn terra.StringValue `hcl:"delivery_stream_arn,attr" validate:"required"`
	// IamRoleArn: string, required
	IamRoleArn terra.StringValue `hcl:"iam_role_arn,attr" validate:"required"`
}

type EventDestinationPinpointDestination struct {
	// ApplicationArn: string, required
	ApplicationArn terra.StringValue `hcl:"application_arn,attr" validate:"required"`
}

type EventDestinationSnsDestination struct {
	// TopicArn: string, required
	TopicArn terra.StringValue `hcl:"topic_arn,attr" validate:"required"`
}

type EventDestinationAttributes struct {
	ref terra.Reference
}

func (ed EventDestinationAttributes) InternalRef() (terra.Reference, error) {
	return ed.ref, nil
}

func (ed EventDestinationAttributes) InternalWithRef(ref terra.Reference) EventDestinationAttributes {
	return EventDestinationAttributes{ref: ref}
}

func (ed EventDestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ed.ref.InternalTokens()
}

func (ed EventDestinationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ed.ref.Append("enabled"))
}

func (ed EventDestinationAttributes) MatchingEventTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ed.ref.Append("matching_event_types"))
}

func (ed EventDestinationAttributes) CloudWatchDestination() terra.ListValue[EventDestinationCloudWatchDestinationAttributes] {
	return terra.ReferenceAsList[EventDestinationCloudWatchDestinationAttributes](ed.ref.Append("cloud_watch_destination"))
}

func (ed EventDestinationAttributes) KinesisFirehoseDestination() terra.ListValue[EventDestinationKinesisFirehoseDestinationAttributes] {
	return terra.ReferenceAsList[EventDestinationKinesisFirehoseDestinationAttributes](ed.ref.Append("kinesis_firehose_destination"))
}

func (ed EventDestinationAttributes) PinpointDestination() terra.ListValue[EventDestinationPinpointDestinationAttributes] {
	return terra.ReferenceAsList[EventDestinationPinpointDestinationAttributes](ed.ref.Append("pinpoint_destination"))
}

func (ed EventDestinationAttributes) SnsDestination() terra.ListValue[EventDestinationSnsDestinationAttributes] {
	return terra.ReferenceAsList[EventDestinationSnsDestinationAttributes](ed.ref.Append("sns_destination"))
}

type EventDestinationCloudWatchDestinationAttributes struct {
	ref terra.Reference
}

func (cwd EventDestinationCloudWatchDestinationAttributes) InternalRef() (terra.Reference, error) {
	return cwd.ref, nil
}

func (cwd EventDestinationCloudWatchDestinationAttributes) InternalWithRef(ref terra.Reference) EventDestinationCloudWatchDestinationAttributes {
	return EventDestinationCloudWatchDestinationAttributes{ref: ref}
}

func (cwd EventDestinationCloudWatchDestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cwd.ref.InternalTokens()
}

func (cwd EventDestinationCloudWatchDestinationAttributes) DimensionConfiguration() terra.ListValue[EventDestinationCloudWatchDestinationDimensionConfigurationAttributes] {
	return terra.ReferenceAsList[EventDestinationCloudWatchDestinationDimensionConfigurationAttributes](cwd.ref.Append("dimension_configuration"))
}

type EventDestinationCloudWatchDestinationDimensionConfigurationAttributes struct {
	ref terra.Reference
}

func (dc EventDestinationCloudWatchDestinationDimensionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc EventDestinationCloudWatchDestinationDimensionConfigurationAttributes) InternalWithRef(ref terra.Reference) EventDestinationCloudWatchDestinationDimensionConfigurationAttributes {
	return EventDestinationCloudWatchDestinationDimensionConfigurationAttributes{ref: ref}
}

func (dc EventDestinationCloudWatchDestinationDimensionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc EventDestinationCloudWatchDestinationDimensionConfigurationAttributes) DefaultDimensionValue() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("default_dimension_value"))
}

func (dc EventDestinationCloudWatchDestinationDimensionConfigurationAttributes) DimensionName() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("dimension_name"))
}

func (dc EventDestinationCloudWatchDestinationDimensionConfigurationAttributes) DimensionValueSource() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("dimension_value_source"))
}

type EventDestinationKinesisFirehoseDestinationAttributes struct {
	ref terra.Reference
}

func (kfd EventDestinationKinesisFirehoseDestinationAttributes) InternalRef() (terra.Reference, error) {
	return kfd.ref, nil
}

func (kfd EventDestinationKinesisFirehoseDestinationAttributes) InternalWithRef(ref terra.Reference) EventDestinationKinesisFirehoseDestinationAttributes {
	return EventDestinationKinesisFirehoseDestinationAttributes{ref: ref}
}

func (kfd EventDestinationKinesisFirehoseDestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kfd.ref.InternalTokens()
}

func (kfd EventDestinationKinesisFirehoseDestinationAttributes) DeliveryStreamArn() terra.StringValue {
	return terra.ReferenceAsString(kfd.ref.Append("delivery_stream_arn"))
}

func (kfd EventDestinationKinesisFirehoseDestinationAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(kfd.ref.Append("iam_role_arn"))
}

type EventDestinationPinpointDestinationAttributes struct {
	ref terra.Reference
}

func (pd EventDestinationPinpointDestinationAttributes) InternalRef() (terra.Reference, error) {
	return pd.ref, nil
}

func (pd EventDestinationPinpointDestinationAttributes) InternalWithRef(ref terra.Reference) EventDestinationPinpointDestinationAttributes {
	return EventDestinationPinpointDestinationAttributes{ref: ref}
}

func (pd EventDestinationPinpointDestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pd.ref.InternalTokens()
}

func (pd EventDestinationPinpointDestinationAttributes) ApplicationArn() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("application_arn"))
}

type EventDestinationSnsDestinationAttributes struct {
	ref terra.Reference
}

func (sd EventDestinationSnsDestinationAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd EventDestinationSnsDestinationAttributes) InternalWithRef(ref terra.Reference) EventDestinationSnsDestinationAttributes {
	return EventDestinationSnsDestinationAttributes{ref: ref}
}

func (sd EventDestinationSnsDestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd EventDestinationSnsDestinationAttributes) TopicArn() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("topic_arn"))
}

type EventDestinationState struct {
	Enabled                    bool                                              `json:"enabled"`
	MatchingEventTypes         []string                                          `json:"matching_event_types"`
	CloudWatchDestination      []EventDestinationCloudWatchDestinationState      `json:"cloud_watch_destination"`
	KinesisFirehoseDestination []EventDestinationKinesisFirehoseDestinationState `json:"kinesis_firehose_destination"`
	PinpointDestination        []EventDestinationPinpointDestinationState        `json:"pinpoint_destination"`
	SnsDestination             []EventDestinationSnsDestinationState             `json:"sns_destination"`
}

type EventDestinationCloudWatchDestinationState struct {
	DimensionConfiguration []EventDestinationCloudWatchDestinationDimensionConfigurationState `json:"dimension_configuration"`
}

type EventDestinationCloudWatchDestinationDimensionConfigurationState struct {
	DefaultDimensionValue string `json:"default_dimension_value"`
	DimensionName         string `json:"dimension_name"`
	DimensionValueSource  string `json:"dimension_value_source"`
}

type EventDestinationKinesisFirehoseDestinationState struct {
	DeliveryStreamArn string `json:"delivery_stream_arn"`
	IamRoleArn        string `json:"iam_role_arn"`
}

type EventDestinationPinpointDestinationState struct {
	ApplicationArn string `json:"application_arn"`
}

type EventDestinationSnsDestinationState struct {
	TopicArn string `json:"topic_arn"`
}
