// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package s3bucketnotification

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type LambdaFunction struct {
	// Events: set of string, required
	Events terra.SetValue[terra.StringValue] `hcl:"events,attr" validate:"required"`
	// FilterPrefix: string, optional
	FilterPrefix terra.StringValue `hcl:"filter_prefix,attr"`
	// FilterSuffix: string, optional
	FilterSuffix terra.StringValue `hcl:"filter_suffix,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LambdaFunctionArn: string, optional
	LambdaFunctionArn terra.StringValue `hcl:"lambda_function_arn,attr"`
}

type Queue struct {
	// Events: set of string, required
	Events terra.SetValue[terra.StringValue] `hcl:"events,attr" validate:"required"`
	// FilterPrefix: string, optional
	FilterPrefix terra.StringValue `hcl:"filter_prefix,attr"`
	// FilterSuffix: string, optional
	FilterSuffix terra.StringValue `hcl:"filter_suffix,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// QueueArn: string, required
	QueueArn terra.StringValue `hcl:"queue_arn,attr" validate:"required"`
}

type Topic struct {
	// Events: set of string, required
	Events terra.SetValue[terra.StringValue] `hcl:"events,attr" validate:"required"`
	// FilterPrefix: string, optional
	FilterPrefix terra.StringValue `hcl:"filter_prefix,attr"`
	// FilterSuffix: string, optional
	FilterSuffix terra.StringValue `hcl:"filter_suffix,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TopicArn: string, required
	TopicArn terra.StringValue `hcl:"topic_arn,attr" validate:"required"`
}

type LambdaFunctionAttributes struct {
	ref terra.Reference
}

func (lf LambdaFunctionAttributes) InternalRef() (terra.Reference, error) {
	return lf.ref, nil
}

func (lf LambdaFunctionAttributes) InternalWithRef(ref terra.Reference) LambdaFunctionAttributes {
	return LambdaFunctionAttributes{ref: ref}
}

func (lf LambdaFunctionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lf.ref.InternalTokens()
}

func (lf LambdaFunctionAttributes) Events() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lf.ref.Append("events"))
}

func (lf LambdaFunctionAttributes) FilterPrefix() terra.StringValue {
	return terra.ReferenceAsString(lf.ref.Append("filter_prefix"))
}

func (lf LambdaFunctionAttributes) FilterSuffix() terra.StringValue {
	return terra.ReferenceAsString(lf.ref.Append("filter_suffix"))
}

func (lf LambdaFunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lf.ref.Append("id"))
}

func (lf LambdaFunctionAttributes) LambdaFunctionArn() terra.StringValue {
	return terra.ReferenceAsString(lf.ref.Append("lambda_function_arn"))
}

type QueueAttributes struct {
	ref terra.Reference
}

func (q QueueAttributes) InternalRef() (terra.Reference, error) {
	return q.ref, nil
}

func (q QueueAttributes) InternalWithRef(ref terra.Reference) QueueAttributes {
	return QueueAttributes{ref: ref}
}

func (q QueueAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return q.ref.InternalTokens()
}

func (q QueueAttributes) Events() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](q.ref.Append("events"))
}

func (q QueueAttributes) FilterPrefix() terra.StringValue {
	return terra.ReferenceAsString(q.ref.Append("filter_prefix"))
}

func (q QueueAttributes) FilterSuffix() terra.StringValue {
	return terra.ReferenceAsString(q.ref.Append("filter_suffix"))
}

func (q QueueAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(q.ref.Append("id"))
}

func (q QueueAttributes) QueueArn() terra.StringValue {
	return terra.ReferenceAsString(q.ref.Append("queue_arn"))
}

type TopicAttributes struct {
	ref terra.Reference
}

func (t TopicAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TopicAttributes) InternalWithRef(ref terra.Reference) TopicAttributes {
	return TopicAttributes{ref: ref}
}

func (t TopicAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TopicAttributes) Events() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](t.ref.Append("events"))
}

func (t TopicAttributes) FilterPrefix() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("filter_prefix"))
}

func (t TopicAttributes) FilterSuffix() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("filter_suffix"))
}

func (t TopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("id"))
}

func (t TopicAttributes) TopicArn() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("topic_arn"))
}

type LambdaFunctionState struct {
	Events            []string `json:"events"`
	FilterPrefix      string   `json:"filter_prefix"`
	FilterSuffix      string   `json:"filter_suffix"`
	Id                string   `json:"id"`
	LambdaFunctionArn string   `json:"lambda_function_arn"`
}

type QueueState struct {
	Events       []string `json:"events"`
	FilterPrefix string   `json:"filter_prefix"`
	FilterSuffix string   `json:"filter_suffix"`
	Id           string   `json:"id"`
	QueueArn     string   `json:"queue_arn"`
}

type TopicState struct {
	Events       []string `json:"events"`
	FilterPrefix string   `json:"filter_prefix"`
	FilterSuffix string   `json:"filter_suffix"`
	Id           string   `json:"id"`
	TopicArn     string   `json:"topic_arn"`
}
