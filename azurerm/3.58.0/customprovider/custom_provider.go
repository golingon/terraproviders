// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package customprovider

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Action struct {
	// Endpoint: string, required
	Endpoint terra.StringValue `hcl:"endpoint,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type ResourceType struct {
	// Endpoint: string, required
	Endpoint terra.StringValue `hcl:"endpoint,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoutingType: string, optional
	RoutingType terra.StringValue `hcl:"routing_type,attr"`
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

type Validation struct {
	// Specification: string, required
	Specification terra.StringValue `hcl:"specification,attr" validate:"required"`
}

type ActionAttributes struct {
	ref terra.Reference
}

func (a ActionAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ActionAttributes) InternalWithRef(ref terra.Reference) ActionAttributes {
	return ActionAttributes{ref: ref}
}

func (a ActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ActionAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("endpoint"))
}

func (a ActionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

type ResourceTypeAttributes struct {
	ref terra.Reference
}

func (rt ResourceTypeAttributes) InternalRef() (terra.Reference, error) {
	return rt.ref, nil
}

func (rt ResourceTypeAttributes) InternalWithRef(ref terra.Reference) ResourceTypeAttributes {
	return ResourceTypeAttributes{ref: ref}
}

func (rt ResourceTypeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rt.ref.InternalTokens()
}

func (rt ResourceTypeAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("endpoint"))
}

func (rt ResourceTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("name"))
}

func (rt ResourceTypeAttributes) RoutingType() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("routing_type"))
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

type ValidationAttributes struct {
	ref terra.Reference
}

func (v ValidationAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v ValidationAttributes) InternalWithRef(ref terra.Reference) ValidationAttributes {
	return ValidationAttributes{ref: ref}
}

func (v ValidationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v ValidationAttributes) Specification() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("specification"))
}

type ActionState struct {
	Endpoint string `json:"endpoint"`
	Name     string `json:"name"`
}

type ResourceTypeState struct {
	Endpoint    string `json:"endpoint"`
	Name        string `json:"name"`
	RoutingType string `json:"routing_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type ValidationState struct {
	Specification string `json:"specification"`
}
