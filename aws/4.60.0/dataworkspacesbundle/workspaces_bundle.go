// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataworkspacesbundle

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ComputeType struct{}

type RootStorage struct{}

type UserStorage struct{}

type ComputeTypeAttributes struct {
	ref terra.Reference
}

func (ct ComputeTypeAttributes) InternalRef() terra.Reference {
	return ct.ref
}

func (ct ComputeTypeAttributes) InternalWithRef(ref terra.Reference) ComputeTypeAttributes {
	return ComputeTypeAttributes{ref: ref}
}

func (ct ComputeTypeAttributes) InternalTokens() hclwrite.Tokens {
	return ct.ref.InternalTokens()
}

func (ct ComputeTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ct.ref.Append("name"))
}

type RootStorageAttributes struct {
	ref terra.Reference
}

func (rs RootStorageAttributes) InternalRef() terra.Reference {
	return rs.ref
}

func (rs RootStorageAttributes) InternalWithRef(ref terra.Reference) RootStorageAttributes {
	return RootStorageAttributes{ref: ref}
}

func (rs RootStorageAttributes) InternalTokens() hclwrite.Tokens {
	return rs.ref.InternalTokens()
}

func (rs RootStorageAttributes) Capacity() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("capacity"))
}

type UserStorageAttributes struct {
	ref terra.Reference
}

func (us UserStorageAttributes) InternalRef() terra.Reference {
	return us.ref
}

func (us UserStorageAttributes) InternalWithRef(ref terra.Reference) UserStorageAttributes {
	return UserStorageAttributes{ref: ref}
}

func (us UserStorageAttributes) InternalTokens() hclwrite.Tokens {
	return us.ref.InternalTokens()
}

func (us UserStorageAttributes) Capacity() terra.StringValue {
	return terra.ReferenceString(us.ref.Append("capacity"))
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
