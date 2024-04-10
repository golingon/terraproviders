// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vmwareenginenetwork

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type VpcNetworks struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type VpcNetworksAttributes struct {
	ref terra.Reference
}

func (vn VpcNetworksAttributes) InternalRef() (terra.Reference, error) {
	return vn.ref, nil
}

func (vn VpcNetworksAttributes) InternalWithRef(ref terra.Reference) VpcNetworksAttributes {
	return VpcNetworksAttributes{ref: ref}
}

func (vn VpcNetworksAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vn.ref.InternalTokens()
}

func (vn VpcNetworksAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("network"))
}

func (vn VpcNetworksAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("type"))
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

type VpcNetworksState struct {
	Network string `json:"network"`
	Type    string `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
