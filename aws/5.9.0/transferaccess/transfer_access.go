// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package transferaccess

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type HomeDirectoryMappings struct {
	// Entry: string, required
	Entry terra.StringValue `hcl:"entry,attr" validate:"required"`
	// Target: string, required
	Target terra.StringValue `hcl:"target,attr" validate:"required"`
}

type PosixProfile struct {
	// Gid: number, required
	Gid terra.NumberValue `hcl:"gid,attr" validate:"required"`
	// SecondaryGids: set of number, optional
	SecondaryGids terra.SetValue[terra.NumberValue] `hcl:"secondary_gids,attr"`
	// Uid: number, required
	Uid terra.NumberValue `hcl:"uid,attr" validate:"required"`
}

type HomeDirectoryMappingsAttributes struct {
	ref terra.Reference
}

func (hdm HomeDirectoryMappingsAttributes) InternalRef() (terra.Reference, error) {
	return hdm.ref, nil
}

func (hdm HomeDirectoryMappingsAttributes) InternalWithRef(ref terra.Reference) HomeDirectoryMappingsAttributes {
	return HomeDirectoryMappingsAttributes{ref: ref}
}

func (hdm HomeDirectoryMappingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hdm.ref.InternalTokens()
}

func (hdm HomeDirectoryMappingsAttributes) Entry() terra.StringValue {
	return terra.ReferenceAsString(hdm.ref.Append("entry"))
}

func (hdm HomeDirectoryMappingsAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(hdm.ref.Append("target"))
}

type PosixProfileAttributes struct {
	ref terra.Reference
}

func (pp PosixProfileAttributes) InternalRef() (terra.Reference, error) {
	return pp.ref, nil
}

func (pp PosixProfileAttributes) InternalWithRef(ref terra.Reference) PosixProfileAttributes {
	return PosixProfileAttributes{ref: ref}
}

func (pp PosixProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pp.ref.InternalTokens()
}

func (pp PosixProfileAttributes) Gid() terra.NumberValue {
	return terra.ReferenceAsNumber(pp.ref.Append("gid"))
}

func (pp PosixProfileAttributes) SecondaryGids() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](pp.ref.Append("secondary_gids"))
}

func (pp PosixProfileAttributes) Uid() terra.NumberValue {
	return terra.ReferenceAsNumber(pp.ref.Append("uid"))
}

type HomeDirectoryMappingsState struct {
	Entry  string `json:"entry"`
	Target string `json:"target"`
}

type PosixProfileState struct {
	Gid           float64   `json:"gid"`
	SecondaryGids []float64 `json:"secondary_gids"`
	Uid           float64   `json:"uid"`
}
