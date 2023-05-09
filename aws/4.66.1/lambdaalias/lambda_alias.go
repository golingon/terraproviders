// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lambdaalias

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RoutingConfig struct {
	// AdditionalVersionWeights: map of number, optional
	AdditionalVersionWeights terra.MapValue[terra.NumberValue] `hcl:"additional_version_weights,attr"`
}

type RoutingConfigAttributes struct {
	ref terra.Reference
}

func (rc RoutingConfigAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RoutingConfigAttributes) InternalWithRef(ref terra.Reference) RoutingConfigAttributes {
	return RoutingConfigAttributes{ref: ref}
}

func (rc RoutingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RoutingConfigAttributes) AdditionalVersionWeights() terra.MapValue[terra.NumberValue] {
	return terra.ReferenceAsMap[terra.NumberValue](rc.ref.Append("additional_version_weights"))
}

type RoutingConfigState struct {
	AdditionalVersionWeights map[string]float64 `json:"additional_version_weights"`
}
