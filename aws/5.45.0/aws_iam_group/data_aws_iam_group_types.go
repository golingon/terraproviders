// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_iam_group

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataUsersAttributes struct {
	ref terra.Reference
}

func (u DataUsersAttributes) InternalRef() (terra.Reference, error) {
	return u.ref, nil
}

func (u DataUsersAttributes) InternalWithRef(ref terra.Reference) DataUsersAttributes {
	return DataUsersAttributes{ref: ref}
}

func (u DataUsersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return u.ref.InternalTokens()
}

func (u DataUsersAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("arn"))
}

func (u DataUsersAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("path"))
}

func (u DataUsersAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("user_id"))
}

func (u DataUsersAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("user_name"))
}

type DataUsersState struct {
	Arn      string `json:"arn"`
	Path     string `json:"path"`
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}
