// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_devopsguru_notification_channel

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataFilters struct{}

type DataSns struct{}

type DataFiltersAttributes struct {
	ref terra.Reference
}

func (f DataFiltersAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f DataFiltersAttributes) InternalWithRef(ref terra.Reference) DataFiltersAttributes {
	return DataFiltersAttributes{ref: ref}
}

func (f DataFiltersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f DataFiltersAttributes) MessageTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](f.ref.Append("message_types"))
}

func (f DataFiltersAttributes) Severities() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](f.ref.Append("severities"))
}

type DataSnsAttributes struct {
	ref terra.Reference
}

func (s DataSnsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s DataSnsAttributes) InternalWithRef(ref terra.Reference) DataSnsAttributes {
	return DataSnsAttributes{ref: ref}
}

func (s DataSnsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s DataSnsAttributes) TopicArn() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("topic_arn"))
}

type DataFiltersState struct {
	MessageTypes []string `json:"message_types"`
	Severities   []string `json:"severities"`
}

type DataSnsState struct {
	TopicArn string `json:"topic_arn"`
}
