// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package neptunecluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ServerlessV2ScalingConfiguration struct {
	// MaxCapacity: number, optional
	MaxCapacity terra.NumberValue `hcl:"max_capacity,attr"`
	// MinCapacity: number, optional
	MinCapacity terra.NumberValue `hcl:"min_capacity,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ServerlessV2ScalingConfigurationAttributes struct {
	ref terra.Reference
}

func (svsc ServerlessV2ScalingConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return svsc.ref, nil
}

func (svsc ServerlessV2ScalingConfigurationAttributes) InternalWithRef(ref terra.Reference) ServerlessV2ScalingConfigurationAttributes {
	return ServerlessV2ScalingConfigurationAttributes{ref: ref}
}

func (svsc ServerlessV2ScalingConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return svsc.ref.InternalTokens()
}

func (svsc ServerlessV2ScalingConfigurationAttributes) MaxCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(svsc.ref.Append("max_capacity"))
}

func (svsc ServerlessV2ScalingConfigurationAttributes) MinCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(svsc.ref.Append("min_capacity"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ServerlessV2ScalingConfigurationState struct {
	MaxCapacity float64 `json:"max_capacity"`
	MinCapacity float64 `json:"min_capacity"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
