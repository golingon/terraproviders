// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package shieldproactiveengagement

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type EmergencyContact struct {
	// ContactNotes: string, optional
	ContactNotes terra.StringValue `hcl:"contact_notes,attr"`
	// EmailAddress: string, required
	EmailAddress terra.StringValue `hcl:"email_address,attr" validate:"required"`
	// PhoneNumber: string, optional
	PhoneNumber terra.StringValue `hcl:"phone_number,attr"`
}

type EmergencyContactAttributes struct {
	ref terra.Reference
}

func (ec EmergencyContactAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EmergencyContactAttributes) InternalWithRef(ref terra.Reference) EmergencyContactAttributes {
	return EmergencyContactAttributes{ref: ref}
}

func (ec EmergencyContactAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EmergencyContactAttributes) ContactNotes() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("contact_notes"))
}

func (ec EmergencyContactAttributes) EmailAddress() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("email_address"))
}

func (ec EmergencyContactAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("phone_number"))
}

type EmergencyContactState struct {
	ContactNotes string `json:"contact_notes"`
	EmailAddress string `json:"email_address"`
	PhoneNumber  string `json:"phone_number"`
}
