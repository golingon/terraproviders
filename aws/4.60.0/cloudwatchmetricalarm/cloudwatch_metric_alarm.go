// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudwatchmetricalarm

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MetricQuery struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// Expression: string, optional
	Expression terra.StringValue `hcl:"expression,attr"`
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// Period: number, optional
	Period terra.NumberValue `hcl:"period,attr"`
	// ReturnData: bool, optional
	ReturnData terra.BoolValue `hcl:"return_data,attr"`
	// Metric: optional
	Metric *Metric `hcl:"metric,block"`
}

type Metric struct {
	// Dimensions: map of string, optional
	Dimensions terra.MapValue[terra.StringValue] `hcl:"dimensions,attr"`
	// MetricName: string, required
	MetricName terra.StringValue `hcl:"metric_name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Period: number, required
	Period terra.NumberValue `hcl:"period,attr" validate:"required"`
	// Stat: string, required
	Stat terra.StringValue `hcl:"stat,attr" validate:"required"`
	// Unit: string, optional
	Unit terra.StringValue `hcl:"unit,attr"`
}

type MetricQueryAttributes struct {
	ref terra.Reference
}

func (mq MetricQueryAttributes) InternalRef() terra.Reference {
	return mq.ref
}

func (mq MetricQueryAttributes) InternalWithRef(ref terra.Reference) MetricQueryAttributes {
	return MetricQueryAttributes{ref: ref}
}

func (mq MetricQueryAttributes) InternalTokens() hclwrite.Tokens {
	return mq.ref.InternalTokens()
}

func (mq MetricQueryAttributes) AccountId() terra.StringValue {
	return terra.ReferenceString(mq.ref.Append("account_id"))
}

func (mq MetricQueryAttributes) Expression() terra.StringValue {
	return terra.ReferenceString(mq.ref.Append("expression"))
}

func (mq MetricQueryAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mq.ref.Append("id"))
}

func (mq MetricQueryAttributes) Label() terra.StringValue {
	return terra.ReferenceString(mq.ref.Append("label"))
}

func (mq MetricQueryAttributes) Period() terra.NumberValue {
	return terra.ReferenceNumber(mq.ref.Append("period"))
}

func (mq MetricQueryAttributes) ReturnData() terra.BoolValue {
	return terra.ReferenceBool(mq.ref.Append("return_data"))
}

func (mq MetricQueryAttributes) Metric() terra.ListValue[MetricAttributes] {
	return terra.ReferenceList[MetricAttributes](mq.ref.Append("metric"))
}

type MetricAttributes struct {
	ref terra.Reference
}

func (m MetricAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m MetricAttributes) InternalWithRef(ref terra.Reference) MetricAttributes {
	return MetricAttributes{ref: ref}
}

func (m MetricAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m MetricAttributes) Dimensions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](m.ref.Append("dimensions"))
}

func (m MetricAttributes) MetricName() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("metric_name"))
}

func (m MetricAttributes) Namespace() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("namespace"))
}

func (m MetricAttributes) Period() terra.NumberValue {
	return terra.ReferenceNumber(m.ref.Append("period"))
}

func (m MetricAttributes) Stat() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("stat"))
}

func (m MetricAttributes) Unit() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("unit"))
}

type MetricQueryState struct {
	AccountId  string        `json:"account_id"`
	Expression string        `json:"expression"`
	Id         string        `json:"id"`
	Label      string        `json:"label"`
	Period     float64       `json:"period"`
	ReturnData bool          `json:"return_data"`
	Metric     []MetricState `json:"metric"`
}

type MetricState struct {
	Dimensions map[string]string `json:"dimensions"`
	MetricName string            `json:"metric_name"`
	Namespace  string            `json:"namespace"`
	Period     float64           `json:"period"`
	Stat       string            `json:"stat"`
	Unit       string            `json:"unit"`
}
