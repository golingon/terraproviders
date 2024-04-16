// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_efs_access_point

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type PosixUser struct {
	// Gid: number, required
	Gid terra.NumberValue `hcl:"gid,attr" validate:"required"`
	// SecondaryGids: set of number, optional
	SecondaryGids terra.SetValue[terra.NumberValue] `hcl:"secondary_gids,attr"`
	// Uid: number, required
	Uid terra.NumberValue `hcl:"uid,attr" validate:"required"`
}

type RootDirectory struct {
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// RootDirectoryCreationInfo: optional
	CreationInfo *RootDirectoryCreationInfo `hcl:"creation_info,block"`
}

type RootDirectoryCreationInfo struct {
	// OwnerGid: number, required
	OwnerGid terra.NumberValue `hcl:"owner_gid,attr" validate:"required"`
	// OwnerUid: number, required
	OwnerUid terra.NumberValue `hcl:"owner_uid,attr" validate:"required"`
	// Permissions: string, required
	Permissions terra.StringValue `hcl:"permissions,attr" validate:"required"`
}

type PosixUserAttributes struct {
	ref terra.Reference
}

func (pu PosixUserAttributes) InternalRef() (terra.Reference, error) {
	return pu.ref, nil
}

func (pu PosixUserAttributes) InternalWithRef(ref terra.Reference) PosixUserAttributes {
	return PosixUserAttributes{ref: ref}
}

func (pu PosixUserAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pu.ref.InternalTokens()
}

func (pu PosixUserAttributes) Gid() terra.NumberValue {
	return terra.ReferenceAsNumber(pu.ref.Append("gid"))
}

func (pu PosixUserAttributes) SecondaryGids() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](pu.ref.Append("secondary_gids"))
}

func (pu PosixUserAttributes) Uid() terra.NumberValue {
	return terra.ReferenceAsNumber(pu.ref.Append("uid"))
}

type RootDirectoryAttributes struct {
	ref terra.Reference
}

func (rd RootDirectoryAttributes) InternalRef() (terra.Reference, error) {
	return rd.ref, nil
}

func (rd RootDirectoryAttributes) InternalWithRef(ref terra.Reference) RootDirectoryAttributes {
	return RootDirectoryAttributes{ref: ref}
}

func (rd RootDirectoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rd.ref.InternalTokens()
}

func (rd RootDirectoryAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("path"))
}

func (rd RootDirectoryAttributes) CreationInfo() terra.ListValue[RootDirectoryCreationInfoAttributes] {
	return terra.ReferenceAsList[RootDirectoryCreationInfoAttributes](rd.ref.Append("creation_info"))
}

type RootDirectoryCreationInfoAttributes struct {
	ref terra.Reference
}

func (ci RootDirectoryCreationInfoAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci RootDirectoryCreationInfoAttributes) InternalWithRef(ref terra.Reference) RootDirectoryCreationInfoAttributes {
	return RootDirectoryCreationInfoAttributes{ref: ref}
}

func (ci RootDirectoryCreationInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ci.ref.InternalTokens()
}

func (ci RootDirectoryCreationInfoAttributes) OwnerGid() terra.NumberValue {
	return terra.ReferenceAsNumber(ci.ref.Append("owner_gid"))
}

func (ci RootDirectoryCreationInfoAttributes) OwnerUid() terra.NumberValue {
	return terra.ReferenceAsNumber(ci.ref.Append("owner_uid"))
}

func (ci RootDirectoryCreationInfoAttributes) Permissions() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("permissions"))
}

type PosixUserState struct {
	Gid           float64   `json:"gid"`
	SecondaryGids []float64 `json:"secondary_gids"`
	Uid           float64   `json:"uid"`
}

type RootDirectoryState struct {
	Path         string                           `json:"path"`
	CreationInfo []RootDirectoryCreationInfoState `json:"creation_info"`
}

type RootDirectoryCreationInfoState struct {
	OwnerGid    float64 `json:"owner_gid"`
	OwnerUid    float64 `json:"owner_uid"`
	Permissions string  `json:"permissions"`
}
