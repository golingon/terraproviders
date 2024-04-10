// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package automationpowershell72module

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ModuleLink struct {
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
	// Hash: optional
	Hash *Hash `hcl:"hash,block"`
}

type Hash struct {
	// Algorithm: string, required
	Algorithm terra.StringValue `hcl:"algorithm,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
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

type ModuleLinkAttributes struct {
	ref terra.Reference
}

func (ml ModuleLinkAttributes) InternalRef() (terra.Reference, error) {
	return ml.ref, nil
}

func (ml ModuleLinkAttributes) InternalWithRef(ref terra.Reference) ModuleLinkAttributes {
	return ModuleLinkAttributes{ref: ref}
}

func (ml ModuleLinkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ml.ref.InternalTokens()
}

func (ml ModuleLinkAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(ml.ref.Append("uri"))
}

func (ml ModuleLinkAttributes) Hash() terra.ListValue[HashAttributes] {
	return terra.ReferenceAsList[HashAttributes](ml.ref.Append("hash"))
}

type HashAttributes struct {
	ref terra.Reference
}

func (h HashAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h HashAttributes) InternalWithRef(ref terra.Reference) HashAttributes {
	return HashAttributes{ref: ref}
}

func (h HashAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h HashAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("algorithm"))
}

func (h HashAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("value"))
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

type ModuleLinkState struct {
	Uri  string      `json:"uri"`
	Hash []HashState `json:"hash"`
}

type HashState struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
