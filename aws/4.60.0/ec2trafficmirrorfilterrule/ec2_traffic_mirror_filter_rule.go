// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ec2trafficmirrorfilterrule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DestinationPortRange struct {
	// FromPort: number, optional
	FromPort terra.NumberValue `hcl:"from_port,attr"`
	// ToPort: number, optional
	ToPort terra.NumberValue `hcl:"to_port,attr"`
}

type SourcePortRange struct {
	// FromPort: number, optional
	FromPort terra.NumberValue `hcl:"from_port,attr"`
	// ToPort: number, optional
	ToPort terra.NumberValue `hcl:"to_port,attr"`
}

type DestinationPortRangeAttributes struct {
	ref terra.Reference
}

func (dpr DestinationPortRangeAttributes) InternalRef() terra.Reference {
	return dpr.ref
}

func (dpr DestinationPortRangeAttributes) InternalWithRef(ref terra.Reference) DestinationPortRangeAttributes {
	return DestinationPortRangeAttributes{ref: ref}
}

func (dpr DestinationPortRangeAttributes) InternalTokens() hclwrite.Tokens {
	return dpr.ref.InternalTokens()
}

func (dpr DestinationPortRangeAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(dpr.ref.Append("from_port"))
}

func (dpr DestinationPortRangeAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(dpr.ref.Append("to_port"))
}

type SourcePortRangeAttributes struct {
	ref terra.Reference
}

func (spr SourcePortRangeAttributes) InternalRef() terra.Reference {
	return spr.ref
}

func (spr SourcePortRangeAttributes) InternalWithRef(ref terra.Reference) SourcePortRangeAttributes {
	return SourcePortRangeAttributes{ref: ref}
}

func (spr SourcePortRangeAttributes) InternalTokens() hclwrite.Tokens {
	return spr.ref.InternalTokens()
}

func (spr SourcePortRangeAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(spr.ref.Append("from_port"))
}

func (spr SourcePortRangeAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(spr.ref.Append("to_port"))
}

type DestinationPortRangeState struct {
	FromPort float64 `json:"from_port"`
	ToPort   float64 `json:"to_port"`
}

type SourcePortRangeState struct {
	FromPort float64 `json:"from_port"`
	ToPort   float64 `json:"to_port"`
}
