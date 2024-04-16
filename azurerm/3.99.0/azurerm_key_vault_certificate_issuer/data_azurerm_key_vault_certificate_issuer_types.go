// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_key_vault_certificate_issuer

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataAdminAttributes struct {
	ref terra.Reference
}

func (a DataAdminAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a DataAdminAttributes) InternalWithRef(ref terra.Reference) DataAdminAttributes {
	return DataAdminAttributes{ref: ref}
}

func (a DataAdminAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a DataAdminAttributes) EmailAddress() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("email_address"))
}

func (a DataAdminAttributes) FirstName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("first_name"))
}

func (a DataAdminAttributes) LastName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("last_name"))
}

func (a DataAdminAttributes) Phone() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("phone"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataAdminState struct {
	EmailAddress string `json:"email_address"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Phone        string `json:"phone"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
