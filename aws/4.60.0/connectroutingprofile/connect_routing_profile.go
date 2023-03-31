// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package connectroutingprofile

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type QueueConfigsAssociated struct{}

type MediaConcurrencies struct {
	// Channel: string, required
	Channel terra.StringValue `hcl:"channel,attr" validate:"required"`
	// Concurrency: number, required
	Concurrency terra.NumberValue `hcl:"concurrency,attr" validate:"required"`
}

type QueueConfigs struct {
	// Channel: string, required
	Channel terra.StringValue `hcl:"channel,attr" validate:"required"`
	// Delay: number, required
	Delay terra.NumberValue `hcl:"delay,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// QueueId: string, required
	QueueId terra.StringValue `hcl:"queue_id,attr" validate:"required"`
}

type QueueConfigsAssociatedAttributes struct {
	ref terra.Reference
}

func (qca QueueConfigsAssociatedAttributes) InternalRef() terra.Reference {
	return qca.ref
}

func (qca QueueConfigsAssociatedAttributes) InternalWithRef(ref terra.Reference) QueueConfigsAssociatedAttributes {
	return QueueConfigsAssociatedAttributes{ref: ref}
}

func (qca QueueConfigsAssociatedAttributes) InternalTokens() hclwrite.Tokens {
	return qca.ref.InternalTokens()
}

func (qca QueueConfigsAssociatedAttributes) Channel() terra.StringValue {
	return terra.ReferenceAsString(qca.ref.Append("channel"))
}

func (qca QueueConfigsAssociatedAttributes) Delay() terra.NumberValue {
	return terra.ReferenceAsNumber(qca.ref.Append("delay"))
}

func (qca QueueConfigsAssociatedAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(qca.ref.Append("priority"))
}

func (qca QueueConfigsAssociatedAttributes) QueueArn() terra.StringValue {
	return terra.ReferenceAsString(qca.ref.Append("queue_arn"))
}

func (qca QueueConfigsAssociatedAttributes) QueueId() terra.StringValue {
	return terra.ReferenceAsString(qca.ref.Append("queue_id"))
}

func (qca QueueConfigsAssociatedAttributes) QueueName() terra.StringValue {
	return terra.ReferenceAsString(qca.ref.Append("queue_name"))
}

type MediaConcurrenciesAttributes struct {
	ref terra.Reference
}

func (mc MediaConcurrenciesAttributes) InternalRef() terra.Reference {
	return mc.ref
}

func (mc MediaConcurrenciesAttributes) InternalWithRef(ref terra.Reference) MediaConcurrenciesAttributes {
	return MediaConcurrenciesAttributes{ref: ref}
}

func (mc MediaConcurrenciesAttributes) InternalTokens() hclwrite.Tokens {
	return mc.ref.InternalTokens()
}

func (mc MediaConcurrenciesAttributes) Channel() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("channel"))
}

func (mc MediaConcurrenciesAttributes) Concurrency() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("concurrency"))
}

type QueueConfigsAttributes struct {
	ref terra.Reference
}

func (qc QueueConfigsAttributes) InternalRef() terra.Reference {
	return qc.ref
}

func (qc QueueConfigsAttributes) InternalWithRef(ref terra.Reference) QueueConfigsAttributes {
	return QueueConfigsAttributes{ref: ref}
}

func (qc QueueConfigsAttributes) InternalTokens() hclwrite.Tokens {
	return qc.ref.InternalTokens()
}

func (qc QueueConfigsAttributes) Channel() terra.StringValue {
	return terra.ReferenceAsString(qc.ref.Append("channel"))
}

func (qc QueueConfigsAttributes) Delay() terra.NumberValue {
	return terra.ReferenceAsNumber(qc.ref.Append("delay"))
}

func (qc QueueConfigsAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(qc.ref.Append("priority"))
}

func (qc QueueConfigsAttributes) QueueArn() terra.StringValue {
	return terra.ReferenceAsString(qc.ref.Append("queue_arn"))
}

func (qc QueueConfigsAttributes) QueueId() terra.StringValue {
	return terra.ReferenceAsString(qc.ref.Append("queue_id"))
}

func (qc QueueConfigsAttributes) QueueName() terra.StringValue {
	return terra.ReferenceAsString(qc.ref.Append("queue_name"))
}

type QueueConfigsAssociatedState struct {
	Channel   string  `json:"channel"`
	Delay     float64 `json:"delay"`
	Priority  float64 `json:"priority"`
	QueueArn  string  `json:"queue_arn"`
	QueueId   string  `json:"queue_id"`
	QueueName string  `json:"queue_name"`
}

type MediaConcurrenciesState struct {
	Channel     string  `json:"channel"`
	Concurrency float64 `json:"concurrency"`
}

type QueueConfigsState struct {
	Channel   string  `json:"channel"`
	Delay     float64 `json:"delay"`
	Priority  float64 `json:"priority"`
	QueueArn  string  `json:"queue_arn"`
	QueueId   string  `json:"queue_id"`
	QueueName string  `json:"queue_name"`
}
