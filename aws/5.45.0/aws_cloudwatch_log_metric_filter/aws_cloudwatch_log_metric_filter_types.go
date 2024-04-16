// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudwatch_log_metric_filter

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type MetricTransformation struct {
	// DefaultValue: string, optional
	DefaultValue terra.StringValue `hcl:"default_value,attr"`
	// Dimensions: map of string, optional
	Dimensions terra.MapValue[terra.StringValue] `hcl:"dimensions,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
	// Unit: string, optional
	Unit terra.StringValue `hcl:"unit,attr"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type MetricTransformationAttributes struct {
	ref terra.Reference
}

func (mt MetricTransformationAttributes) InternalRef() (terra.Reference, error) {
	return mt.ref, nil
}

func (mt MetricTransformationAttributes) InternalWithRef(ref terra.Reference) MetricTransformationAttributes {
	return MetricTransformationAttributes{ref: ref}
}

func (mt MetricTransformationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mt.ref.InternalTokens()
}

func (mt MetricTransformationAttributes) DefaultValue() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("default_value"))
}

func (mt MetricTransformationAttributes) Dimensions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mt.ref.Append("dimensions"))
}

func (mt MetricTransformationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("name"))
}

func (mt MetricTransformationAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("namespace"))
}

func (mt MetricTransformationAttributes) Unit() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("unit"))
}

func (mt MetricTransformationAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("value"))
}

type MetricTransformationState struct {
	DefaultValue string            `json:"default_value"`
	Dimensions   map[string]string `json:"dimensions"`
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Unit         string            `json:"unit"`
	Value        string            `json:"value"`
}
