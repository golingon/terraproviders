// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package synapselinkedservice

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type IntegrationRuntime struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type IntegrationRuntimeAttributes struct {
	ref terra.Reference
}

func (ir IntegrationRuntimeAttributes) InternalRef() (terra.Reference, error) {
	return ir.ref, nil
}

func (ir IntegrationRuntimeAttributes) InternalWithRef(ref terra.Reference) IntegrationRuntimeAttributes {
	return IntegrationRuntimeAttributes{ref: ref}
}

func (ir IntegrationRuntimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ir.ref.InternalTokens()
}

func (ir IntegrationRuntimeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("name"))
}

func (ir IntegrationRuntimeAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ir.ref.Append("parameters"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type IntegrationRuntimeState struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
