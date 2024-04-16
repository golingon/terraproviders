// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_traffic_manager_external_endpoint

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CustomHeader struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type Subnet struct {
	// First: string, required
	First terra.StringValue `hcl:"first,attr" validate:"required"`
	// Last: string, optional
	Last terra.StringValue `hcl:"last,attr"`
	// Scope: number, optional
	Scope terra.NumberValue `hcl:"scope,attr"`
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

type CustomHeaderAttributes struct {
	ref terra.Reference
}

func (ch CustomHeaderAttributes) InternalRef() (terra.Reference, error) {
	return ch.ref, nil
}

func (ch CustomHeaderAttributes) InternalWithRef(ref terra.Reference) CustomHeaderAttributes {
	return CustomHeaderAttributes{ref: ref}
}

func (ch CustomHeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ch.ref.InternalTokens()
}

func (ch CustomHeaderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("name"))
}

func (ch CustomHeaderAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("value"))
}

type SubnetAttributes struct {
	ref terra.Reference
}

func (s SubnetAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SubnetAttributes) InternalWithRef(ref terra.Reference) SubnetAttributes {
	return SubnetAttributes{ref: ref}
}

func (s SubnetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SubnetAttributes) First() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("first"))
}

func (s SubnetAttributes) Last() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("last"))
}

func (s SubnetAttributes) Scope() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("scope"))
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

type CustomHeaderState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SubnetState struct {
	First string  `json:"first"`
	Last  string  `json:"last"`
	Scope float64 `json:"scope"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
