// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_snapshot

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataSnapshotEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (sek DataSnapshotEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return sek.ref, nil
}

func (sek DataSnapshotEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) DataSnapshotEncryptionKeyAttributes {
	return DataSnapshotEncryptionKeyAttributes{ref: ref}
}

func (sek DataSnapshotEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sek.ref.InternalTokens()
}

func (sek DataSnapshotEncryptionKeyAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(sek.ref.Append("kms_key_self_link"))
}

func (sek DataSnapshotEncryptionKeyAttributes) KmsKeyServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(sek.ref.Append("kms_key_service_account"))
}

func (sek DataSnapshotEncryptionKeyAttributes) RawKey() terra.StringValue {
	return terra.ReferenceAsString(sek.ref.Append("raw_key"))
}

func (sek DataSnapshotEncryptionKeyAttributes) Sha256() terra.StringValue {
	return terra.ReferenceAsString(sek.ref.Append("sha256"))
}

type DataSourceDiskEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (sdek DataSourceDiskEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return sdek.ref, nil
}

func (sdek DataSourceDiskEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) DataSourceDiskEncryptionKeyAttributes {
	return DataSourceDiskEncryptionKeyAttributes{ref: ref}
}

func (sdek DataSourceDiskEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sdek.ref.InternalTokens()
}

func (sdek DataSourceDiskEncryptionKeyAttributes) KmsKeyServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(sdek.ref.Append("kms_key_service_account"))
}

func (sdek DataSourceDiskEncryptionKeyAttributes) RawKey() terra.StringValue {
	return terra.ReferenceAsString(sdek.ref.Append("raw_key"))
}

type DataSnapshotEncryptionKeyState struct {
	KmsKeySelfLink       string `json:"kms_key_self_link"`
	KmsKeyServiceAccount string `json:"kms_key_service_account"`
	RawKey               string `json:"raw_key"`
	Sha256               string `json:"sha256"`
}

type DataSourceDiskEncryptionKeyState struct {
	KmsKeyServiceAccount string `json:"kms_key_service_account"`
	RawKey               string `json:"raw_key"`
}
