// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package logicappintegrationaccountpartner

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BusinessIdentity struct {
	// Qualifier: string, required
	Qualifier terra.StringValue `hcl:"qualifier,attr" validate:"required"`
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

type BusinessIdentityAttributes struct {
	ref terra.Reference
}

func (bi BusinessIdentityAttributes) InternalRef() (terra.Reference, error) {
	return bi.ref, nil
}

func (bi BusinessIdentityAttributes) InternalWithRef(ref terra.Reference) BusinessIdentityAttributes {
	return BusinessIdentityAttributes{ref: ref}
}

func (bi BusinessIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bi.ref.InternalTokens()
}

func (bi BusinessIdentityAttributes) Qualifier() terra.StringValue {
	return terra.ReferenceAsString(bi.ref.Append("qualifier"))
}

func (bi BusinessIdentityAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(bi.ref.Append("value"))
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

type BusinessIdentityState struct {
	Qualifier string `json:"qualifier"`
	Value     string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
