// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datavpcpeeringconnection

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CidrBlockSet struct{}

type PeerCidrBlockSet struct{}

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

type CidrBlockSetAttributes struct {
	ref terra.Reference
}

func (cbs CidrBlockSetAttributes) InternalRef() (terra.Reference, error) {
	return cbs.ref, nil
}

func (cbs CidrBlockSetAttributes) InternalWithRef(ref terra.Reference) CidrBlockSetAttributes {
	return CidrBlockSetAttributes{ref: ref}
}

func (cbs CidrBlockSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cbs.ref.InternalTokens()
}

func (cbs CidrBlockSetAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("cidr_block"))
}

type PeerCidrBlockSetAttributes struct {
	ref terra.Reference
}

func (pcbs PeerCidrBlockSetAttributes) InternalRef() (terra.Reference, error) {
	return pcbs.ref, nil
}

func (pcbs PeerCidrBlockSetAttributes) InternalWithRef(ref terra.Reference) PeerCidrBlockSetAttributes {
	return PeerCidrBlockSetAttributes{ref: ref}
}

func (pcbs PeerCidrBlockSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pcbs.ref.InternalTokens()
}

func (pcbs PeerCidrBlockSetAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(pcbs.ref.Append("cidr_block"))
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

type CidrBlockSetState struct {
	CidrBlock string `json:"cidr_block"`
}

type PeerCidrBlockSetState struct {
	CidrBlock string `json:"cidr_block"`
}

type FilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}