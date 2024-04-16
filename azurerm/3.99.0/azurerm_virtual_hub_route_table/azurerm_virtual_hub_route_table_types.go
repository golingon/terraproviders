// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_virtual_hub_route_table

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Route struct {
	// Destinations: set of string, required
	Destinations terra.SetValue[terra.StringValue] `hcl:"destinations,attr" validate:"required"`
	// DestinationsType: string, required
	DestinationsType terra.StringValue `hcl:"destinations_type,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NextHop: string, required
	NextHop terra.StringValue `hcl:"next_hop,attr" validate:"required"`
	// NextHopType: string, optional
	NextHopType terra.StringValue `hcl:"next_hop_type,attr"`
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

type RouteAttributes struct {
	ref terra.Reference
}

func (r RouteAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RouteAttributes) InternalWithRef(ref terra.Reference) RouteAttributes {
	return RouteAttributes{ref: ref}
}

func (r RouteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RouteAttributes) Destinations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](r.ref.Append("destinations"))
}

func (r RouteAttributes) DestinationsType() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("destinations_type"))
}

func (r RouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r RouteAttributes) NextHop() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("next_hop"))
}

func (r RouteAttributes) NextHopType() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("next_hop_type"))
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

type RouteState struct {
	Destinations     []string `json:"destinations"`
	DestinationsType string   `json:"destinations_type"`
	Name             string   `json:"name"`
	NextHop          string   `json:"next_hop"`
	NextHopType      string   `json:"next_hop_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
