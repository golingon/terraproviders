// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package elasticbeanstalkapplication

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AppversionLifecycle struct {
	// DeleteSourceFromS3: bool, optional
	DeleteSourceFromS3 terra.BoolValue `hcl:"delete_source_from_s3,attr"`
	// MaxAgeInDays: number, optional
	MaxAgeInDays terra.NumberValue `hcl:"max_age_in_days,attr"`
	// MaxCount: number, optional
	MaxCount terra.NumberValue `hcl:"max_count,attr"`
	// ServiceRole: string, required
	ServiceRole terra.StringValue `hcl:"service_role,attr" validate:"required"`
}

type AppversionLifecycleAttributes struct {
	ref terra.Reference
}

func (al AppversionLifecycleAttributes) InternalRef() terra.Reference {
	return al.ref
}

func (al AppversionLifecycleAttributes) InternalWithRef(ref terra.Reference) AppversionLifecycleAttributes {
	return AppversionLifecycleAttributes{ref: ref}
}

func (al AppversionLifecycleAttributes) InternalTokens() hclwrite.Tokens {
	return al.ref.InternalTokens()
}

func (al AppversionLifecycleAttributes) DeleteSourceFromS3() terra.BoolValue {
	return terra.ReferenceBool(al.ref.Append("delete_source_from_s3"))
}

func (al AppversionLifecycleAttributes) MaxAgeInDays() terra.NumberValue {
	return terra.ReferenceNumber(al.ref.Append("max_age_in_days"))
}

func (al AppversionLifecycleAttributes) MaxCount() terra.NumberValue {
	return terra.ReferenceNumber(al.ref.Append("max_count"))
}

func (al AppversionLifecycleAttributes) ServiceRole() terra.StringValue {
	return terra.ReferenceString(al.ref.Append("service_role"))
}

type AppversionLifecycleState struct {
	DeleteSourceFromS3 bool    `json:"delete_source_from_s3"`
	MaxAgeInDays       float64 `json:"max_age_in_days"`
	MaxCount           float64 `json:"max_count"`
	ServiceRole        string  `json:"service_role"`
}
