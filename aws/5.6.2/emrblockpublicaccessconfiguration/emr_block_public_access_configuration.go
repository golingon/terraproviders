// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package emrblockpublicaccessconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PermittedPublicSecurityGroupRuleRange struct {
	// MaxRange: number, required
	MaxRange terra.NumberValue `hcl:"max_range,attr" validate:"required"`
	// MinRange: number, required
	MinRange terra.NumberValue `hcl:"min_range,attr" validate:"required"`
}

type PermittedPublicSecurityGroupRuleRangeAttributes struct {
	ref terra.Reference
}

func (ppsgrr PermittedPublicSecurityGroupRuleRangeAttributes) InternalRef() (terra.Reference, error) {
	return ppsgrr.ref, nil
}

func (ppsgrr PermittedPublicSecurityGroupRuleRangeAttributes) InternalWithRef(ref terra.Reference) PermittedPublicSecurityGroupRuleRangeAttributes {
	return PermittedPublicSecurityGroupRuleRangeAttributes{ref: ref}
}

func (ppsgrr PermittedPublicSecurityGroupRuleRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ppsgrr.ref.InternalTokens()
}

func (ppsgrr PermittedPublicSecurityGroupRuleRangeAttributes) MaxRange() terra.NumberValue {
	return terra.ReferenceAsNumber(ppsgrr.ref.Append("max_range"))
}

func (ppsgrr PermittedPublicSecurityGroupRuleRangeAttributes) MinRange() terra.NumberValue {
	return terra.ReferenceAsNumber(ppsgrr.ref.Append("min_range"))
}

type PermittedPublicSecurityGroupRuleRangeState struct {
	MaxRange float64 `json:"max_range"`
	MinRange float64 `json:"min_range"`
}
