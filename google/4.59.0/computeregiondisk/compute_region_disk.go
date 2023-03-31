// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package computeregiondisk

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DiskEncryptionKey struct {
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
	// RawKey: string, optional
	RawKey terra.StringValue `hcl:"raw_key,attr"`
}

type SourceSnapshotEncryptionKey struct {
	// RawKey: string, optional
	RawKey terra.StringValue `hcl:"raw_key,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DiskEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (dek DiskEncryptionKeyAttributes) InternalRef() terra.Reference {
	return dek.ref
}

func (dek DiskEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) DiskEncryptionKeyAttributes {
	return DiskEncryptionKeyAttributes{ref: ref}
}

func (dek DiskEncryptionKeyAttributes) InternalTokens() hclwrite.Tokens {
	return dek.ref.InternalTokens()
}

func (dek DiskEncryptionKeyAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceString(dek.ref.Append("kms_key_name"))
}

func (dek DiskEncryptionKeyAttributes) RawKey() terra.StringValue {
	return terra.ReferenceString(dek.ref.Append("raw_key"))
}

func (dek DiskEncryptionKeyAttributes) Sha256() terra.StringValue {
	return terra.ReferenceString(dek.ref.Append("sha256"))
}

type SourceSnapshotEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (ssek SourceSnapshotEncryptionKeyAttributes) InternalRef() terra.Reference {
	return ssek.ref
}

func (ssek SourceSnapshotEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) SourceSnapshotEncryptionKeyAttributes {
	return SourceSnapshotEncryptionKeyAttributes{ref: ref}
}

func (ssek SourceSnapshotEncryptionKeyAttributes) InternalTokens() hclwrite.Tokens {
	return ssek.ref.InternalTokens()
}

func (ssek SourceSnapshotEncryptionKeyAttributes) RawKey() terra.StringValue {
	return terra.ReferenceString(ssek.ref.Append("raw_key"))
}

func (ssek SourceSnapshotEncryptionKeyAttributes) Sha256() terra.StringValue {
	return terra.ReferenceString(ssek.ref.Append("sha256"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type DiskEncryptionKeyState struct {
	KmsKeyName string `json:"kms_key_name"`
	RawKey     string `json:"raw_key"`
	Sha256     string `json:"sha256"`
}

type SourceSnapshotEncryptionKeyState struct {
	RawKey string `json:"raw_key"`
	Sha256 string `json:"sha256"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
