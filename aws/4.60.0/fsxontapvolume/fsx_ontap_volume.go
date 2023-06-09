// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package fsxontapvolume

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type TieringPolicy struct {
	// CoolingPeriod: number, optional
	CoolingPeriod terra.NumberValue `hcl:"cooling_period,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type TieringPolicyAttributes struct {
	ref terra.Reference
}

func (tp TieringPolicyAttributes) InternalRef() (terra.Reference, error) {
	return tp.ref, nil
}

func (tp TieringPolicyAttributes) InternalWithRef(ref terra.Reference) TieringPolicyAttributes {
	return TieringPolicyAttributes{ref: ref}
}

func (tp TieringPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tp.ref.InternalTokens()
}

func (tp TieringPolicyAttributes) CoolingPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(tp.ref.Append("cooling_period"))
}

func (tp TieringPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tp.ref.Append("name"))
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

type TieringPolicyState struct {
	CoolingPeriod float64 `json:"cooling_period"`
	Name          string  `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
