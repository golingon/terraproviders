// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dataworkspacesbundle

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ComputeType struct{}

type RootStorage struct{}

type UserStorage struct{}

type ComputeTypeAttributes struct {
	ref terra.Reference
}

func (ct ComputeTypeAttributes) InternalRef() (terra.Reference, error) {
	return ct.ref, nil
}

func (ct ComputeTypeAttributes) InternalWithRef(ref terra.Reference) ComputeTypeAttributes {
	return ComputeTypeAttributes{ref: ref}
}

func (ct ComputeTypeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ct.ref.InternalTokens()
}

func (ct ComputeTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("name"))
}

type RootStorageAttributes struct {
	ref terra.Reference
}

func (rs RootStorageAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs RootStorageAttributes) InternalWithRef(ref terra.Reference) RootStorageAttributes {
	return RootStorageAttributes{ref: ref}
}

func (rs RootStorageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs RootStorageAttributes) Capacity() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("capacity"))
}

type UserStorageAttributes struct {
	ref terra.Reference
}

func (us UserStorageAttributes) InternalRef() (terra.Reference, error) {
	return us.ref, nil
}

func (us UserStorageAttributes) InternalWithRef(ref terra.Reference) UserStorageAttributes {
	return UserStorageAttributes{ref: ref}
}

func (us UserStorageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return us.ref.InternalTokens()
}

func (us UserStorageAttributes) Capacity() terra.StringValue {
	return terra.ReferenceAsString(us.ref.Append("capacity"))
}

type ComputeTypeState struct {
	Name string `json:"name"`
}

type RootStorageState struct {
	Capacity string `json:"capacity"`
}

type UserStorageState struct {
	Capacity string `json:"capacity"`
}
