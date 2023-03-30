// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package gameliftalias

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RoutingStrategy struct {
	// FleetId: string, optional
	FleetId terra.StringValue `hcl:"fleet_id,attr"`
	// Message: string, optional
	Message terra.StringValue `hcl:"message,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type RoutingStrategyAttributes struct {
	ref terra.Reference
}

func (rs RoutingStrategyAttributes) InternalRef() terra.Reference {
	return rs.ref
}

func (rs RoutingStrategyAttributes) InternalWithRef(ref terra.Reference) RoutingStrategyAttributes {
	return RoutingStrategyAttributes{ref: ref}
}

func (rs RoutingStrategyAttributes) InternalTokens() hclwrite.Tokens {
	return rs.ref.InternalTokens()
}

func (rs RoutingStrategyAttributes) FleetId() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("fleet_id"))
}

func (rs RoutingStrategyAttributes) Message() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("message"))
}

func (rs RoutingStrategyAttributes) Type() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("type"))
}

type RoutingStrategyState struct {
	FleetId string `json:"fleet_id"`
	Message string `json:"message"`
	Type    string `json:"type"`
}
