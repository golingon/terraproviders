// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datafactorydatasetpostgresql

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SchemaColumn struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
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

type SchemaColumnAttributes struct {
	ref terra.Reference
}

func (sc SchemaColumnAttributes) InternalRef() terra.Reference {
	return sc.ref
}

func (sc SchemaColumnAttributes) InternalWithRef(ref terra.Reference) SchemaColumnAttributes {
	return SchemaColumnAttributes{ref: ref}
}

func (sc SchemaColumnAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc SchemaColumnAttributes) Description() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("description"))
}

func (sc SchemaColumnAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("name"))
}

func (sc SchemaColumnAttributes) Type() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("type"))
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

type SchemaColumnState struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
