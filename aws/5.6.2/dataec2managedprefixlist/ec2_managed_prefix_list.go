// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataec2managedprefixlist

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Entries struct{}

type Filter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type EntriesAttributes struct {
	ref terra.Reference
}

func (e EntriesAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EntriesAttributes) InternalWithRef(ref terra.Reference) EntriesAttributes {
	return EntriesAttributes{ref: ref}
}

func (e EntriesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EntriesAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("cidr"))
}

func (e EntriesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("description"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f FilterAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](f.ref.Append("values"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type EntriesState struct {
	Cidr        string `json:"cidr"`
	Description string `json:"description"`
}

type FilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
