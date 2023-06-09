// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package keyvaultcertificateissuer

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Admin struct {
	// EmailAddress: string, required
	EmailAddress terra.StringValue `hcl:"email_address,attr" validate:"required"`
	// FirstName: string, optional
	FirstName terra.StringValue `hcl:"first_name,attr"`
	// LastName: string, optional
	LastName terra.StringValue `hcl:"last_name,attr"`
	// Phone: string, optional
	Phone terra.StringValue `hcl:"phone,attr"`
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

type AdminAttributes struct {
	ref terra.Reference
}

func (a AdminAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AdminAttributes) InternalWithRef(ref terra.Reference) AdminAttributes {
	return AdminAttributes{ref: ref}
}

func (a AdminAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AdminAttributes) EmailAddress() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("email_address"))
}

func (a AdminAttributes) FirstName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("first_name"))
}

func (a AdminAttributes) LastName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("last_name"))
}

func (a AdminAttributes) Phone() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("phone"))
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

type AdminState struct {
	EmailAddress string `json:"email_address"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Phone        string `json:"phone"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
