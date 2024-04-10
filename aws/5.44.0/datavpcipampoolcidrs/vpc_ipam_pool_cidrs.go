// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datavpcipampoolcidrs

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type IpamPoolCidrs struct{}

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

type IpamPoolCidrsAttributes struct {
	ref terra.Reference
}

func (ipc IpamPoolCidrsAttributes) InternalRef() (terra.Reference, error) {
	return ipc.ref, nil
}

func (ipc IpamPoolCidrsAttributes) InternalWithRef(ref terra.Reference) IpamPoolCidrsAttributes {
	return IpamPoolCidrsAttributes{ref: ref}
}

func (ipc IpamPoolCidrsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ipc.ref.InternalTokens()
}

func (ipc IpamPoolCidrsAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(ipc.ref.Append("cidr"))
}

func (ipc IpamPoolCidrsAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ipc.ref.Append("state"))
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

type IpamPoolCidrsState struct {
	Cidr  string `json:"cidr"`
	State string `json:"state"`
}

type FilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
