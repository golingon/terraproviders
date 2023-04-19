// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package globalacceleratorlistener

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PortRange struct {
	// FromPort: number, optional
	FromPort terra.NumberValue `hcl:"from_port,attr"`
	// ToPort: number, optional
	ToPort terra.NumberValue `hcl:"to_port,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type PortRangeAttributes struct {
	ref terra.Reference
}

func (pr PortRangeAttributes) InternalRef() (terra.Reference, error) {
	return pr.ref, nil
}

func (pr PortRangeAttributes) InternalWithRef(ref terra.Reference) PortRangeAttributes {
	return PortRangeAttributes{ref: ref}
}

func (pr PortRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pr.ref.InternalTokens()
}

func (pr PortRangeAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(pr.ref.Append("from_port"))
}

func (pr PortRangeAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(pr.ref.Append("to_port"))
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

type PortRangeState struct {
	FromPort float64 `json:"from_port"`
	ToPort   float64 `json:"to_port"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}