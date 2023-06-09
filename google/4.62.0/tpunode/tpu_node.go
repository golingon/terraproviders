// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package tpunode

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type NetworkEndpoints struct{}

type SchedulingConfig struct {
	// Preemptible: bool, required
	Preemptible terra.BoolValue `hcl:"preemptible,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type NetworkEndpointsAttributes struct {
	ref terra.Reference
}

func (ne NetworkEndpointsAttributes) InternalRef() (terra.Reference, error) {
	return ne.ref, nil
}

func (ne NetworkEndpointsAttributes) InternalWithRef(ref terra.Reference) NetworkEndpointsAttributes {
	return NetworkEndpointsAttributes{ref: ref}
}

func (ne NetworkEndpointsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ne.ref.InternalTokens()
}

func (ne NetworkEndpointsAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ne.ref.Append("ip_address"))
}

func (ne NetworkEndpointsAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ne.ref.Append("port"))
}

type SchedulingConfigAttributes struct {
	ref terra.Reference
}

func (sc SchedulingConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SchedulingConfigAttributes) InternalWithRef(ref terra.Reference) SchedulingConfigAttributes {
	return SchedulingConfigAttributes{ref: ref}
}

func (sc SchedulingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SchedulingConfigAttributes) Preemptible() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("preemptible"))
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

type NetworkEndpointsState struct {
	IpAddress string  `json:"ip_address"`
	Port      float64 `json:"port"`
}

type SchedulingConfigState struct {
	Preemptible bool `json:"preemptible"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
