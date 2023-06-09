// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package chimesdkvoicesiprule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type TargetApplications struct {
	// AwsRegion: string, required
	AwsRegion terra.StringValue `hcl:"aws_region,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// SipMediaApplicationId: string, required
	SipMediaApplicationId terra.StringValue `hcl:"sip_media_application_id,attr" validate:"required"`
}

type TargetApplicationsAttributes struct {
	ref terra.Reference
}

func (ta TargetApplicationsAttributes) InternalRef() (terra.Reference, error) {
	return ta.ref, nil
}

func (ta TargetApplicationsAttributes) InternalWithRef(ref terra.Reference) TargetApplicationsAttributes {
	return TargetApplicationsAttributes{ref: ref}
}

func (ta TargetApplicationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ta.ref.InternalTokens()
}

func (ta TargetApplicationsAttributes) AwsRegion() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("aws_region"))
}

func (ta TargetApplicationsAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(ta.ref.Append("priority"))
}

func (ta TargetApplicationsAttributes) SipMediaApplicationId() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("sip_media_application_id"))
}

type TargetApplicationsState struct {
	AwsRegion             string  `json:"aws_region"`
	Priority              float64 `json:"priority"`
	SipMediaApplicationId string  `json:"sip_media_application_id"`
}
