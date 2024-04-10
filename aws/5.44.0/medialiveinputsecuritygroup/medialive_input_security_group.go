// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package medialiveinputsecuritygroup

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type WhitelistRules struct {
	// Cidr: string, required
	Cidr terra.StringValue `hcl:"cidr,attr" validate:"required"`
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

type WhitelistRulesAttributes struct {
	ref terra.Reference
}

func (wr WhitelistRulesAttributes) InternalRef() (terra.Reference, error) {
	return wr.ref, nil
}

func (wr WhitelistRulesAttributes) InternalWithRef(ref terra.Reference) WhitelistRulesAttributes {
	return WhitelistRulesAttributes{ref: ref}
}

func (wr WhitelistRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wr.ref.InternalTokens()
}

func (wr WhitelistRulesAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(wr.ref.Append("cidr"))
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type WhitelistRulesState struct {
	Cidr string `json:"cidr"`
}
