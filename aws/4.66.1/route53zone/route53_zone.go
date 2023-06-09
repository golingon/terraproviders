// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package route53zone

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Vpc struct {
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// VpcRegion: string, optional
	VpcRegion terra.StringValue `hcl:"vpc_region,attr"`
}

type VpcAttributes struct {
	ref terra.Reference
}

func (v VpcAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v VpcAttributes) InternalWithRef(ref terra.Reference) VpcAttributes {
	return VpcAttributes{ref: ref}
}

func (v VpcAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v VpcAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("vpc_id"))
}

func (v VpcAttributes) VpcRegion() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("vpc_region"))
}

type VpcState struct {
	VpcId     string `json:"vpc_id"`
	VpcRegion string `json:"vpc_region"`
}
