// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package routetable

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Route struct {
	// AddressPrefix: string, optional
	AddressPrefix terra.StringValue `hcl:"address_prefix,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NextHopInIpAddress: string, optional
	NextHopInIpAddress terra.StringValue `hcl:"next_hop_in_ip_address,attr"`
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

func (r RouteAttributes) AddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("address_prefix"))
}

func (r RouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r RouteAttributes) NextHopInIpAddress() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("next_hop_in_ip_address"))
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
	AddressPrefix      string `json:"address_prefix"`
	Name               string `json:"name"`
	NextHopInIpAddress string `json:"next_hop_in_ip_address"`
	NextHopType        string `json:"next_hop_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
