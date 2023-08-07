// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package connectuser

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type IdentityInfo struct {
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// FirstName: string, optional
	FirstName terra.StringValue `hcl:"first_name,attr"`
	// LastName: string, optional
	LastName terra.StringValue `hcl:"last_name,attr"`
}

type PhoneConfig struct {
	// AfterContactWorkTimeLimit: number, optional
	AfterContactWorkTimeLimit terra.NumberValue `hcl:"after_contact_work_time_limit,attr"`
	// AutoAccept: bool, optional
	AutoAccept terra.BoolValue `hcl:"auto_accept,attr"`
	// DeskPhoneNumber: string, optional
	DeskPhoneNumber terra.StringValue `hcl:"desk_phone_number,attr"`
	// PhoneType: string, required
	PhoneType terra.StringValue `hcl:"phone_type,attr" validate:"required"`
}

type IdentityInfoAttributes struct {
	ref terra.Reference
}

func (ii IdentityInfoAttributes) InternalRef() (terra.Reference, error) {
	return ii.ref, nil
}

func (ii IdentityInfoAttributes) InternalWithRef(ref terra.Reference) IdentityInfoAttributes {
	return IdentityInfoAttributes{ref: ref}
}

func (ii IdentityInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ii.ref.InternalTokens()
}

func (ii IdentityInfoAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(ii.ref.Append("email"))
}

func (ii IdentityInfoAttributes) FirstName() terra.StringValue {
	return terra.ReferenceAsString(ii.ref.Append("first_name"))
}

func (ii IdentityInfoAttributes) LastName() terra.StringValue {
	return terra.ReferenceAsString(ii.ref.Append("last_name"))
}

type PhoneConfigAttributes struct {
	ref terra.Reference
}

func (pc PhoneConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PhoneConfigAttributes) InternalWithRef(ref terra.Reference) PhoneConfigAttributes {
	return PhoneConfigAttributes{ref: ref}
}

func (pc PhoneConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PhoneConfigAttributes) AfterContactWorkTimeLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(pc.ref.Append("after_contact_work_time_limit"))
}

func (pc PhoneConfigAttributes) AutoAccept() terra.BoolValue {
	return terra.ReferenceAsBool(pc.ref.Append("auto_accept"))
}

func (pc PhoneConfigAttributes) DeskPhoneNumber() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("desk_phone_number"))
}

func (pc PhoneConfigAttributes) PhoneType() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("phone_type"))
}

type IdentityInfoState struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type PhoneConfigState struct {
	AfterContactWorkTimeLimit float64 `json:"after_contact_work_time_limit"`
	AutoAccept                bool    `json:"auto_accept"`
	DeskPhoneNumber           string  `json:"desk_phone_number"`
	PhoneType                 string  `json:"phone_type"`
}
