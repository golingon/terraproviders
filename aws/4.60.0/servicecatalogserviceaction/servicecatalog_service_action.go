// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package servicecatalogserviceaction

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Definition struct {
	// AssumeRole: string, optional
	AssumeRole terra.StringValue `hcl:"assume_role,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: string, optional
	Parameters terra.StringValue `hcl:"parameters,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
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

type DefinitionAttributes struct {
	ref terra.Reference
}

func (d DefinitionAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d DefinitionAttributes) InternalWithRef(ref terra.Reference) DefinitionAttributes {
	return DefinitionAttributes{ref: ref}
}

func (d DefinitionAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d DefinitionAttributes) AssumeRole() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("assume_role"))
}

func (d DefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("name"))
}

func (d DefinitionAttributes) Parameters() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("parameters"))
}

func (d DefinitionAttributes) Type() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("type"))
}

func (d DefinitionAttributes) Version() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("version"))
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
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type DefinitionState struct {
	AssumeRole string `json:"assume_role"`
	Name       string `json:"name"`
	Parameters string `json:"parameters"`
	Type       string `json:"type"`
	Version    string `json:"version"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
