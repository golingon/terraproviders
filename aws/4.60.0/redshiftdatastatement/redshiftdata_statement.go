// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package redshiftdatastatement

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Parameters struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
}

type ParametersAttributes struct {
	ref terra.Reference
}

func (p ParametersAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p ParametersAttributes) InternalWithRef(ref terra.Reference) ParametersAttributes {
	return ParametersAttributes{ref: ref}
}

func (p ParametersAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p ParametersAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p ParametersAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("value"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

type ParametersState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
}
