// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datasnapshot

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EncryptionSettings struct {
	// DiskEncryptionKey: min=0
	DiskEncryptionKey []DiskEncryptionKey `hcl:"disk_encryption_key,block" validate:"min=0"`
	// KeyEncryptionKey: min=0
	KeyEncryptionKey []KeyEncryptionKey `hcl:"key_encryption_key,block" validate:"min=0"`
}

type DiskEncryptionKey struct{}

type KeyEncryptionKey struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type EncryptionSettingsAttributes struct {
	ref terra.Reference
}

func (es EncryptionSettingsAttributes) InternalRef() (terra.Reference, error) {
	return es.ref, nil
}

func (es EncryptionSettingsAttributes) InternalWithRef(ref terra.Reference) EncryptionSettingsAttributes {
	return EncryptionSettingsAttributes{ref: ref}
}

func (es EncryptionSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return es.ref.InternalTokens()
}

func (es EncryptionSettingsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(es.ref.Append("enabled"))
}

func (es EncryptionSettingsAttributes) DiskEncryptionKey() terra.ListValue[DiskEncryptionKeyAttributes] {
	return terra.ReferenceAsList[DiskEncryptionKeyAttributes](es.ref.Append("disk_encryption_key"))
}

func (es EncryptionSettingsAttributes) KeyEncryptionKey() terra.ListValue[KeyEncryptionKeyAttributes] {
	return terra.ReferenceAsList[KeyEncryptionKeyAttributes](es.ref.Append("key_encryption_key"))
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

func (dek DiskEncryptionKeyAttributes) SecretUrl() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("secret_url"))
}

func (dek DiskEncryptionKeyAttributes) SourceVaultId() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("source_vault_id"))
}

type KeyEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (kek KeyEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return kek.ref, nil
}

func (kek KeyEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) KeyEncryptionKeyAttributes {
	return KeyEncryptionKeyAttributes{ref: ref}
}

func (kek KeyEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kek.ref.InternalTokens()
}

func (kek KeyEncryptionKeyAttributes) KeyUrl() terra.StringValue {
	return terra.ReferenceAsString(kek.ref.Append("key_url"))
}

func (kek KeyEncryptionKeyAttributes) SourceVaultId() terra.StringValue {
	return terra.ReferenceAsString(kek.ref.Append("source_vault_id"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type EncryptionSettingsState struct {
	Enabled           bool                     `json:"enabled"`
	DiskEncryptionKey []DiskEncryptionKeyState `json:"disk_encryption_key"`
	KeyEncryptionKey  []KeyEncryptionKeyState  `json:"key_encryption_key"`
}

type DiskEncryptionKeyState struct {
	SecretUrl     string `json:"secret_url"`
	SourceVaultId string `json:"source_vault_id"`
}

type KeyEncryptionKeyState struct {
	KeyUrl        string `json:"key_url"`
	SourceVaultId string `json:"source_vault_id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
