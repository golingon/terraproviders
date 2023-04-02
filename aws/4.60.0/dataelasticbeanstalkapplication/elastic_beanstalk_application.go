// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataelasticbeanstalkapplication

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AppversionLifecycle struct{}

type AppversionLifecycleAttributes struct {
	ref terra.Reference
}

func (al AppversionLifecycleAttributes) InternalRef() (terra.Reference, error) {
	return al.ref, nil
}

func (al AppversionLifecycleAttributes) InternalWithRef(ref terra.Reference) AppversionLifecycleAttributes {
	return AppversionLifecycleAttributes{ref: ref}
}

func (al AppversionLifecycleAttributes) InternalTokens() hclwrite.Tokens {
	return al.ref.InternalTokens()
}

func (al AppversionLifecycleAttributes) DeleteSourceFromS3() terra.BoolValue {
	return terra.ReferenceAsBool(al.ref.Append("delete_source_from_s3"))
}

func (al AppversionLifecycleAttributes) MaxAgeInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(al.ref.Append("max_age_in_days"))
}

func (al AppversionLifecycleAttributes) MaxCount() terra.NumberValue {
	return terra.ReferenceAsNumber(al.ref.Append("max_count"))
}

func (al AppversionLifecycleAttributes) ServiceRole() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("service_role"))
}

type AppversionLifecycleState struct {
	DeleteSourceFromS3 bool    `json:"delete_source_from_s3"`
	MaxAgeInDays       float64 `json:"max_age_in_days"`
	MaxCount           float64 `json:"max_count"`
	ServiceRole        string  `json:"service_role"`
}
