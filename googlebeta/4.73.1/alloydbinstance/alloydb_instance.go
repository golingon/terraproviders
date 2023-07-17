// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package alloydbinstance

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MachineConfig struct {
	// CpuCount: number, optional
	CpuCount terra.NumberValue `hcl:"cpu_count,attr"`
}

type ReadPoolConfig struct {
	// NodeCount: number, optional
	NodeCount terra.NumberValue `hcl:"node_count,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type MachineConfigAttributes struct {
	ref terra.Reference
}

func (mc MachineConfigAttributes) InternalRef() (terra.Reference, error) {
	return mc.ref, nil
}

func (mc MachineConfigAttributes) InternalWithRef(ref terra.Reference) MachineConfigAttributes {
	return MachineConfigAttributes{ref: ref}
}

func (mc MachineConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mc.ref.InternalTokens()
}

func (mc MachineConfigAttributes) CpuCount() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("cpu_count"))
}

type ReadPoolConfigAttributes struct {
	ref terra.Reference
}

func (rpc ReadPoolConfigAttributes) InternalRef() (terra.Reference, error) {
	return rpc.ref, nil
}

func (rpc ReadPoolConfigAttributes) InternalWithRef(ref terra.Reference) ReadPoolConfigAttributes {
	return ReadPoolConfigAttributes{ref: ref}
}

func (rpc ReadPoolConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rpc.ref.InternalTokens()
}

func (rpc ReadPoolConfigAttributes) NodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(rpc.ref.Append("node_count"))
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

type MachineConfigState struct {
	CpuCount float64 `json:"cpu_count"`
}

type ReadPoolConfigState struct {
	NodeCount float64 `json:"node_count"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}