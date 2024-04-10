// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package batchschedulingpolicy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type FairSharePolicy struct {
	// ComputeReservation: number, optional
	ComputeReservation terra.NumberValue `hcl:"compute_reservation,attr"`
	// ShareDecaySeconds: number, optional
	ShareDecaySeconds terra.NumberValue `hcl:"share_decay_seconds,attr"`
	// ShareDistribution: min=0,max=500
	ShareDistribution []ShareDistribution `hcl:"share_distribution,block" validate:"min=0,max=500"`
}

type ShareDistribution struct {
	// ShareIdentifier: string, required
	ShareIdentifier terra.StringValue `hcl:"share_identifier,attr" validate:"required"`
	// WeightFactor: number, optional
	WeightFactor terra.NumberValue `hcl:"weight_factor,attr"`
}

type FairSharePolicyAttributes struct {
	ref terra.Reference
}

func (fsp FairSharePolicyAttributes) InternalRef() (terra.Reference, error) {
	return fsp.ref, nil
}

func (fsp FairSharePolicyAttributes) InternalWithRef(ref terra.Reference) FairSharePolicyAttributes {
	return FairSharePolicyAttributes{ref: ref}
}

func (fsp FairSharePolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fsp.ref.InternalTokens()
}

func (fsp FairSharePolicyAttributes) ComputeReservation() terra.NumberValue {
	return terra.ReferenceAsNumber(fsp.ref.Append("compute_reservation"))
}

func (fsp FairSharePolicyAttributes) ShareDecaySeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(fsp.ref.Append("share_decay_seconds"))
}

func (fsp FairSharePolicyAttributes) ShareDistribution() terra.SetValue[ShareDistributionAttributes] {
	return terra.ReferenceAsSet[ShareDistributionAttributes](fsp.ref.Append("share_distribution"))
}

type ShareDistributionAttributes struct {
	ref terra.Reference
}

func (sd ShareDistributionAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd ShareDistributionAttributes) InternalWithRef(ref terra.Reference) ShareDistributionAttributes {
	return ShareDistributionAttributes{ref: ref}
}

func (sd ShareDistributionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd ShareDistributionAttributes) ShareIdentifier() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("share_identifier"))
}

func (sd ShareDistributionAttributes) WeightFactor() terra.NumberValue {
	return terra.ReferenceAsNumber(sd.ref.Append("weight_factor"))
}

type FairSharePolicyState struct {
	ComputeReservation float64                  `json:"compute_reservation"`
	ShareDecaySeconds  float64                  `json:"share_decay_seconds"`
	ShareDistribution  []ShareDistributionState `json:"share_distribution"`
}

type ShareDistributionState struct {
	ShareIdentifier string  `json:"share_identifier"`
	WeightFactor    float64 `json:"weight_factor"`
}
