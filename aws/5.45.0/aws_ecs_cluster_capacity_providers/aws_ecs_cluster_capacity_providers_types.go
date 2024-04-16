// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ecs_cluster_capacity_providers

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DefaultCapacityProviderStrategy struct {
	// Base: number, optional
	Base terra.NumberValue `hcl:"base,attr"`
	// CapacityProvider: string, required
	CapacityProvider terra.StringValue `hcl:"capacity_provider,attr" validate:"required"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
}

type DefaultCapacityProviderStrategyAttributes struct {
	ref terra.Reference
}

func (dcps DefaultCapacityProviderStrategyAttributes) InternalRef() (terra.Reference, error) {
	return dcps.ref, nil
}

func (dcps DefaultCapacityProviderStrategyAttributes) InternalWithRef(ref terra.Reference) DefaultCapacityProviderStrategyAttributes {
	return DefaultCapacityProviderStrategyAttributes{ref: ref}
}

func (dcps DefaultCapacityProviderStrategyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dcps.ref.InternalTokens()
}

func (dcps DefaultCapacityProviderStrategyAttributes) Base() terra.NumberValue {
	return terra.ReferenceAsNumber(dcps.ref.Append("base"))
}

func (dcps DefaultCapacityProviderStrategyAttributes) CapacityProvider() terra.StringValue {
	return terra.ReferenceAsString(dcps.ref.Append("capacity_provider"))
}

func (dcps DefaultCapacityProviderStrategyAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(dcps.ref.Append("weight"))
}

type DefaultCapacityProviderStrategyState struct {
	Base             float64 `json:"base"`
	CapacityProvider string  `json:"capacity_provider"`
	Weight           float64 `json:"weight"`
}
