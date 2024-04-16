// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_logic_app_integration_account_agreement

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type GuestIdentity struct {
	// Qualifier: string, required
	Qualifier terra.StringValue `hcl:"qualifier,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type HostIdentity struct {
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

type GuestIdentityAttributes struct {
	ref terra.Reference
}

func (gi GuestIdentityAttributes) InternalRef() (terra.Reference, error) {
	return gi.ref, nil
}

func (gi GuestIdentityAttributes) InternalWithRef(ref terra.Reference) GuestIdentityAttributes {
	return GuestIdentityAttributes{ref: ref}
}

func (gi GuestIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gi.ref.InternalTokens()
}

func (gi GuestIdentityAttributes) Qualifier() terra.StringValue {
	return terra.ReferenceAsString(gi.ref.Append("qualifier"))
}

func (gi GuestIdentityAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(gi.ref.Append("value"))
}

type HostIdentityAttributes struct {
	ref terra.Reference
}

func (hi HostIdentityAttributes) InternalRef() (terra.Reference, error) {
	return hi.ref, nil
}

func (hi HostIdentityAttributes) InternalWithRef(ref terra.Reference) HostIdentityAttributes {
	return HostIdentityAttributes{ref: ref}
}

func (hi HostIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hi.ref.InternalTokens()
}

func (hi HostIdentityAttributes) Qualifier() terra.StringValue {
	return terra.ReferenceAsString(hi.ref.Append("qualifier"))
}

func (hi HostIdentityAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(hi.ref.Append("value"))
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

type GuestIdentityState struct {
	Qualifier string `json:"qualifier"`
	Value     string `json:"value"`
}

type HostIdentityState struct {
	Qualifier string `json:"qualifier"`
	Value     string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
