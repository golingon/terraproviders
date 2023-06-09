// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package efsaccesspoint

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
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
	// CreationInfo: optional
	CreationInfo *CreationInfo `hcl:"creation_info,block"`
}

type CreationInfo struct {
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

func (rd RootDirectoryAttributes) CreationInfo() terra.ListValue[CreationInfoAttributes] {
	return terra.ReferenceAsList[CreationInfoAttributes](rd.ref.Append("creation_info"))
}

type CreationInfoAttributes struct {
	ref terra.Reference
}

func (ci CreationInfoAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci CreationInfoAttributes) InternalWithRef(ref terra.Reference) CreationInfoAttributes {
	return CreationInfoAttributes{ref: ref}
}

func (ci CreationInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ci.ref.InternalTokens()
}

func (ci CreationInfoAttributes) OwnerGid() terra.NumberValue {
	return terra.ReferenceAsNumber(ci.ref.Append("owner_gid"))
}

func (ci CreationInfoAttributes) OwnerUid() terra.NumberValue {
	return terra.ReferenceAsNumber(ci.ref.Append("owner_uid"))
}

func (ci CreationInfoAttributes) Permissions() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("permissions"))
}

type PosixUserState struct {
	Gid           float64   `json:"gid"`
	SecondaryGids []float64 `json:"secondary_gids"`
	Uid           float64   `json:"uid"`
}

type RootDirectoryState struct {
	Path         string              `json:"path"`
	CreationInfo []CreationInfoState `json:"creation_info"`
}

type CreationInfoState struct {
	OwnerGid    float64 `json:"owner_gid"`
	OwnerUid    float64 `json:"owner_uid"`
	Permissions string  `json:"permissions"`
}
