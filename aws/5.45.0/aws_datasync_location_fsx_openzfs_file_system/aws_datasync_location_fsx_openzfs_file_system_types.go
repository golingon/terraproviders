// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_datasync_location_fsx_openzfs_file_system

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Protocol struct {
	// ProtocolNfs: required
	Nfs *ProtocolNfs `hcl:"nfs,block" validate:"required"`
}

type ProtocolNfs struct {
	// ProtocolNfsMountOptions: required
	MountOptions *ProtocolNfsMountOptions `hcl:"mount_options,block" validate:"required"`
}

type ProtocolNfsMountOptions struct {
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}

type ProtocolAttributes struct {
	ref terra.Reference
}

func (p ProtocolAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ProtocolAttributes) InternalWithRef(ref terra.Reference) ProtocolAttributes {
	return ProtocolAttributes{ref: ref}
}

func (p ProtocolAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ProtocolAttributes) Nfs() terra.ListValue[ProtocolNfsAttributes] {
	return terra.ReferenceAsList[ProtocolNfsAttributes](p.ref.Append("nfs"))
}

type ProtocolNfsAttributes struct {
	ref terra.Reference
}

func (n ProtocolNfsAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n ProtocolNfsAttributes) InternalWithRef(ref terra.Reference) ProtocolNfsAttributes {
	return ProtocolNfsAttributes{ref: ref}
}

func (n ProtocolNfsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n ProtocolNfsAttributes) MountOptions() terra.ListValue[ProtocolNfsMountOptionsAttributes] {
	return terra.ReferenceAsList[ProtocolNfsMountOptionsAttributes](n.ref.Append("mount_options"))
}

type ProtocolNfsMountOptionsAttributes struct {
	ref terra.Reference
}

func (mo ProtocolNfsMountOptionsAttributes) InternalRef() (terra.Reference, error) {
	return mo.ref, nil
}

func (mo ProtocolNfsMountOptionsAttributes) InternalWithRef(ref terra.Reference) ProtocolNfsMountOptionsAttributes {
	return ProtocolNfsMountOptionsAttributes{ref: ref}
}

func (mo ProtocolNfsMountOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mo.ref.InternalTokens()
}

func (mo ProtocolNfsMountOptionsAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("version"))
}

type ProtocolState struct {
	Nfs []ProtocolNfsState `json:"nfs"`
}

type ProtocolNfsState struct {
	MountOptions []ProtocolNfsMountOptionsState `json:"mount_options"`
}

type ProtocolNfsMountOptionsState struct {
	Version string `json:"version"`
}
