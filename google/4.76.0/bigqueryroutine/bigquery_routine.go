// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package bigqueryroutine

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Arguments struct {
	// ArgumentKind: string, optional
	ArgumentKind terra.StringValue `hcl:"argument_kind,attr"`
	// DataType: string, optional
	DataType terra.StringValue `hcl:"data_type,attr"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
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

type ArgumentsAttributes struct {
	ref terra.Reference
}

func (a ArgumentsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ArgumentsAttributes) InternalWithRef(ref terra.Reference) ArgumentsAttributes {
	return ArgumentsAttributes{ref: ref}
}

func (a ArgumentsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ArgumentsAttributes) ArgumentKind() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("argument_kind"))
}

func (a ArgumentsAttributes) DataType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("data_type"))
}

func (a ArgumentsAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("mode"))
}

func (a ArgumentsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
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

type ArgumentsState struct {
	ArgumentKind string `json:"argument_kind"`
	DataType     string `json:"data_type"`
	Mode         string `json:"mode"`
	Name         string `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}