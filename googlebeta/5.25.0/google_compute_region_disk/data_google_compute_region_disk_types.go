// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_region_disk

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataAsyncPrimaryDiskAttributes struct {
	ref terra.Reference
}

func (apd DataAsyncPrimaryDiskAttributes) InternalRef() (terra.Reference, error) {
	return apd.ref, nil
}

func (apd DataAsyncPrimaryDiskAttributes) InternalWithRef(ref terra.Reference) DataAsyncPrimaryDiskAttributes {
	return DataAsyncPrimaryDiskAttributes{ref: ref}
}

func (apd DataAsyncPrimaryDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return apd.ref.InternalTokens()
}

func (apd DataAsyncPrimaryDiskAttributes) Disk() terra.StringValue {
	return terra.ReferenceAsString(apd.ref.Append("disk"))
}

type DataDiskEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (dek DataDiskEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return dek.ref, nil
}

func (dek DataDiskEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) DataDiskEncryptionKeyAttributes {
	return DataDiskEncryptionKeyAttributes{ref: ref}
}

func (dek DataDiskEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dek.ref.InternalTokens()
}

func (dek DataDiskEncryptionKeyAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("kms_key_name"))
}

func (dek DataDiskEncryptionKeyAttributes) RawKey() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("raw_key"))
}

func (dek DataDiskEncryptionKeyAttributes) Sha256() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("sha256"))
}

type DataGuestOsFeaturesAttributes struct {
	ref terra.Reference
}

func (gof DataGuestOsFeaturesAttributes) InternalRef() (terra.Reference, error) {
	return gof.ref, nil
}

func (gof DataGuestOsFeaturesAttributes) InternalWithRef(ref terra.Reference) DataGuestOsFeaturesAttributes {
	return DataGuestOsFeaturesAttributes{ref: ref}
}

func (gof DataGuestOsFeaturesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gof.ref.InternalTokens()
}

func (gof DataGuestOsFeaturesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(gof.ref.Append("type"))
}

type DataSourceSnapshotEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (ssek DataSourceSnapshotEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return ssek.ref, nil
}

func (ssek DataSourceSnapshotEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) DataSourceSnapshotEncryptionKeyAttributes {
	return DataSourceSnapshotEncryptionKeyAttributes{ref: ref}
}

func (ssek DataSourceSnapshotEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssek.ref.InternalTokens()
}

func (ssek DataSourceSnapshotEncryptionKeyAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(ssek.ref.Append("kms_key_name"))
}

func (ssek DataSourceSnapshotEncryptionKeyAttributes) RawKey() terra.StringValue {
	return terra.ReferenceAsString(ssek.ref.Append("raw_key"))
}

func (ssek DataSourceSnapshotEncryptionKeyAttributes) Sha256() terra.StringValue {
	return terra.ReferenceAsString(ssek.ref.Append("sha256"))
}

type DataAsyncPrimaryDiskState struct {
	Disk string `json:"disk"`
}

type DataDiskEncryptionKeyState struct {
	KmsKeyName string `json:"kms_key_name"`
	RawKey     string `json:"raw_key"`
	Sha256     string `json:"sha256"`
}

type DataGuestOsFeaturesState struct {
	Type string `json:"type"`
}

type DataSourceSnapshotEncryptionKeyState struct {
	KmsKeyName string `json:"kms_key_name"`
	RawKey     string `json:"raw_key"`
	Sha256     string `json:"sha256"`
}
