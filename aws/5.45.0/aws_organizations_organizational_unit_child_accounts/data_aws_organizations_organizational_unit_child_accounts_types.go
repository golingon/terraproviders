// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_organizations_organizational_unit_child_accounts

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataAccountsAttributes struct {
	ref terra.Reference
}

func (a DataAccountsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a DataAccountsAttributes) InternalWithRef(ref terra.Reference) DataAccountsAttributes {
	return DataAccountsAttributes{ref: ref}
}

func (a DataAccountsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a DataAccountsAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("arn"))
}

func (a DataAccountsAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("email"))
}

func (a DataAccountsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("id"))
}

func (a DataAccountsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

func (a DataAccountsAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("status"))
}

type DataAccountsState struct {
	Arn    string `json:"arn"`
	Email  string `json:"email"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
