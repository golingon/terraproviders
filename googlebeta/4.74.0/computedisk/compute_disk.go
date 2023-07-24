// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package computedisk

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AsyncPrimaryDisk struct {
	// Disk: string, required
	Disk terra.StringValue `hcl:"disk,attr" validate:"required"`
}

type DiskEncryptionKey struct {
	// KmsKeySelfLink: string, optional
	KmsKeySelfLink terra.StringValue `hcl:"kms_key_self_link,attr"`
	// KmsKeyServiceAccount: string, optional
	KmsKeyServiceAccount terra.StringValue `hcl:"kms_key_service_account,attr"`
	// RawKey: string, optional
	RawKey terra.StringValue `hcl:"raw_key,attr"`
	// RsaEncryptedKey: string, optional
	RsaEncryptedKey terra.StringValue `hcl:"rsa_encrypted_key,attr"`
}

type GuestOsFeatures struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type SourceImageEncryptionKey struct {
	// KmsKeySelfLink: string, optional
	KmsKeySelfLink terra.StringValue `hcl:"kms_key_self_link,attr"`
	// KmsKeyServiceAccount: string, optional
	KmsKeyServiceAccount terra.StringValue `hcl:"kms_key_service_account,attr"`
	// RawKey: string, optional
	RawKey terra.StringValue `hcl:"raw_key,attr"`
}

type SourceSnapshotEncryptionKey struct {
	// KmsKeySelfLink: string, optional
	KmsKeySelfLink terra.StringValue `hcl:"kms_key_self_link,attr"`
	// KmsKeyServiceAccount: string, optional
	KmsKeyServiceAccount terra.StringValue `hcl:"kms_key_service_account,attr"`
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

type AsyncPrimaryDiskAttributes struct {
	ref terra.Reference
}

func (apd AsyncPrimaryDiskAttributes) InternalRef() (terra.Reference, error) {
	return apd.ref, nil
}

func (apd AsyncPrimaryDiskAttributes) InternalWithRef(ref terra.Reference) AsyncPrimaryDiskAttributes {
	return AsyncPrimaryDiskAttributes{ref: ref}
}

func (apd AsyncPrimaryDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return apd.ref.InternalTokens()
}

func (apd AsyncPrimaryDiskAttributes) Disk() terra.StringValue {
	return terra.ReferenceAsString(apd.ref.Append("disk"))
}

type DiskEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (dek DiskEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return dek.ref, nil
}

func (dek DiskEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) DiskEncryptionKeyAttributes {
	return DiskEncryptionKeyAttributes{ref: ref}
}

func (dek DiskEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dek.ref.InternalTokens()
}

func (dek DiskEncryptionKeyAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("kms_key_self_link"))
}

func (dek DiskEncryptionKeyAttributes) KmsKeyServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("kms_key_service_account"))
}

func (dek DiskEncryptionKeyAttributes) RawKey() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("raw_key"))
}

func (dek DiskEncryptionKeyAttributes) RsaEncryptedKey() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("rsa_encrypted_key"))
}

func (dek DiskEncryptionKeyAttributes) Sha256() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("sha256"))
}

type GuestOsFeaturesAttributes struct {
	ref terra.Reference
}

func (gof GuestOsFeaturesAttributes) InternalRef() (terra.Reference, error) {
	return gof.ref, nil
}

func (gof GuestOsFeaturesAttributes) InternalWithRef(ref terra.Reference) GuestOsFeaturesAttributes {
	return GuestOsFeaturesAttributes{ref: ref}
}

func (gof GuestOsFeaturesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gof.ref.InternalTokens()
}

func (gof GuestOsFeaturesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(gof.ref.Append("type"))
}

type SourceImageEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (siek SourceImageEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return siek.ref, nil
}

func (siek SourceImageEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) SourceImageEncryptionKeyAttributes {
	return SourceImageEncryptionKeyAttributes{ref: ref}
}

func (siek SourceImageEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return siek.ref.InternalTokens()
}

func (siek SourceImageEncryptionKeyAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(siek.ref.Append("kms_key_self_link"))
}

func (siek SourceImageEncryptionKeyAttributes) KmsKeyServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(siek.ref.Append("kms_key_service_account"))
}

func (siek SourceImageEncryptionKeyAttributes) RawKey() terra.StringValue {
	return terra.ReferenceAsString(siek.ref.Append("raw_key"))
}

func (siek SourceImageEncryptionKeyAttributes) Sha256() terra.StringValue {
	return terra.ReferenceAsString(siek.ref.Append("sha256"))
}

type SourceSnapshotEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (ssek SourceSnapshotEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return ssek.ref, nil
}

func (ssek SourceSnapshotEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) SourceSnapshotEncryptionKeyAttributes {
	return SourceSnapshotEncryptionKeyAttributes{ref: ref}
}

func (ssek SourceSnapshotEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssek.ref.InternalTokens()
}

func (ssek SourceSnapshotEncryptionKeyAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(ssek.ref.Append("kms_key_self_link"))
}

func (ssek SourceSnapshotEncryptionKeyAttributes) KmsKeyServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(ssek.ref.Append("kms_key_service_account"))
}

func (ssek SourceSnapshotEncryptionKeyAttributes) RawKey() terra.StringValue {
	return terra.ReferenceAsString(ssek.ref.Append("raw_key"))
}

func (ssek SourceSnapshotEncryptionKeyAttributes) Sha256() terra.StringValue {
	return terra.ReferenceAsString(ssek.ref.Append("sha256"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AsyncPrimaryDiskState struct {
	Disk string `json:"disk"`
}

type DiskEncryptionKeyState struct {
	KmsKeySelfLink       string `json:"kms_key_self_link"`
	KmsKeyServiceAccount string `json:"kms_key_service_account"`
	RawKey               string `json:"raw_key"`
	RsaEncryptedKey      string `json:"rsa_encrypted_key"`
	Sha256               string `json:"sha256"`
}

type GuestOsFeaturesState struct {
	Type string `json:"type"`
}

type SourceImageEncryptionKeyState struct {
	KmsKeySelfLink       string `json:"kms_key_self_link"`
	KmsKeyServiceAccount string `json:"kms_key_service_account"`
	RawKey               string `json:"raw_key"`
	Sha256               string `json:"sha256"`
}

type SourceSnapshotEncryptionKeyState struct {
	KmsKeySelfLink       string `json:"kms_key_self_link"`
	KmsKeyServiceAccount string `json:"kms_key_service_account"`
	RawKey               string `json:"raw_key"`
	Sha256               string `json:"sha256"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
