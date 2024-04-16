// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_route_table

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataRouteAttributes struct {
	ref terra.Reference
}

func (r DataRouteAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r DataRouteAttributes) InternalWithRef(ref terra.Reference) DataRouteAttributes {
	return DataRouteAttributes{ref: ref}
}

func (r DataRouteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r DataRouteAttributes) AddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("address_prefix"))
}

func (r DataRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r DataRouteAttributes) NextHopInIpAddress() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("next_hop_in_ip_address"))
}

func (r DataRouteAttributes) NextHopType() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("next_hop_type"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataRouteState struct {
	AddressPrefix      string `json:"address_prefix"`
	Name               string `json:"name"`
	NextHopInIpAddress string `json:"next_hop_in_ip_address"`
	NextHopType        string `json:"next_hop_type"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
