// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datasecretsmanagersecretrotation

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RotationRules struct{}

type RotationRulesAttributes struct {
	ref terra.Reference
}

func (rr RotationRulesAttributes) InternalRef() (terra.Reference, error) {
	return rr.ref, nil
}

func (rr RotationRulesAttributes) InternalWithRef(ref terra.Reference) RotationRulesAttributes {
	return RotationRulesAttributes{ref: ref}
}

func (rr RotationRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rr.ref.InternalTokens()
}

func (rr RotationRulesAttributes) AutomaticallyAfterDays() terra.NumberValue {
	return terra.ReferenceAsNumber(rr.ref.Append("automatically_after_days"))
}

type RotationRulesState struct {
	AutomaticallyAfterDays float64 `json:"automatically_after_days"`
}
