// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataiamgroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Users struct{}

type UsersAttributes struct {
	ref terra.Reference
}

func (u UsersAttributes) InternalRef() (terra.Reference, error) {
	return u.ref, nil
}

func (u UsersAttributes) InternalWithRef(ref terra.Reference) UsersAttributes {
	return UsersAttributes{ref: ref}
}

func (u UsersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return u.ref.InternalTokens()
}

func (u UsersAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("arn"))
}

func (u UsersAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("path"))
}

func (u UsersAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("user_id"))
}

func (u UsersAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("user_name"))
}

type UsersState struct {
	Arn      string `json:"arn"`
	Path     string `json:"path"`
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}
