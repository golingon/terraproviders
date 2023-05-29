// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package gameliftgamesessionqueue

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PlayerLatencyPolicy struct {
	// MaximumIndividualPlayerLatencyMilliseconds: number, required
	MaximumIndividualPlayerLatencyMilliseconds terra.NumberValue `hcl:"maximum_individual_player_latency_milliseconds,attr" validate:"required"`
	// PolicyDurationSeconds: number, optional
	PolicyDurationSeconds terra.NumberValue `hcl:"policy_duration_seconds,attr"`
}

type PlayerLatencyPolicyAttributes struct {
	ref terra.Reference
}

func (plp PlayerLatencyPolicyAttributes) InternalRef() (terra.Reference, error) {
	return plp.ref, nil
}

func (plp PlayerLatencyPolicyAttributes) InternalWithRef(ref terra.Reference) PlayerLatencyPolicyAttributes {
	return PlayerLatencyPolicyAttributes{ref: ref}
}

func (plp PlayerLatencyPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return plp.ref.InternalTokens()
}

func (plp PlayerLatencyPolicyAttributes) MaximumIndividualPlayerLatencyMilliseconds() terra.NumberValue {
	return terra.ReferenceAsNumber(plp.ref.Append("maximum_individual_player_latency_milliseconds"))
}

func (plp PlayerLatencyPolicyAttributes) PolicyDurationSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(plp.ref.Append("policy_duration_seconds"))
}

type PlayerLatencyPolicyState struct {
	MaximumIndividualPlayerLatencyMilliseconds float64 `json:"maximum_individual_player_latency_milliseconds"`
	PolicyDurationSeconds                      float64 `json:"policy_duration_seconds"`
}