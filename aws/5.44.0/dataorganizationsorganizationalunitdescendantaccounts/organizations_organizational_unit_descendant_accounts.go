// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dataorganizationsorganizationalunitdescendantaccounts

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Accounts struct{}

type AccountsAttributes struct {
	ref terra.Reference
}

func (a AccountsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AccountsAttributes) InternalWithRef(ref terra.Reference) AccountsAttributes {
	return AccountsAttributes{ref: ref}
}

func (a AccountsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AccountsAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("arn"))
}

func (a AccountsAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("email"))
}

func (a AccountsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("id"))
}

func (a AccountsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

func (a AccountsAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("status"))
}

type AccountsState struct {
	Arn    string `json:"arn"`
	Email  string `json:"email"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
