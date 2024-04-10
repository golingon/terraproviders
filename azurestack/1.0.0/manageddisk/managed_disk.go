// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package manageddisk

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Encryption struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// DiskEncryptionKey: optional
	DiskEncryptionKey *DiskEncryptionKey `hcl:"disk_encryption_key,block"`
	// KeyEncryptionKey: optional
	KeyEncryptionKey *KeyEncryptionKey `hcl:"key_encryption_key,block"`
}

type DiskEncryptionKey struct {
	// SecretUrl: string, required
	SecretUrl terra.StringValue `hcl:"secret_url,attr" validate:"required"`
	// SourceVaultId: string, required
	SourceVaultId terra.StringValue `hcl:"source_vault_id,attr" validate:"required"`
}

type KeyEncryptionKey struct {
	// KeyUrl: string, required
	KeyUrl terra.StringValue `hcl:"key_url,attr" validate:"required"`
	// SourceVaultId: string, required
	SourceVaultId terra.StringValue `hcl:"source_vault_id,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type EncryptionAttributes struct {
	ref terra.Reference
}

func (e EncryptionAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EncryptionAttributes) InternalWithRef(ref terra.Reference) EncryptionAttributes {
	return EncryptionAttributes{ref: ref}
}

func (e EncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EncryptionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("enabled"))
}

func (e EncryptionAttributes) DiskEncryptionKey() terra.ListValue[DiskEncryptionKeyAttributes] {
	return terra.ReferenceAsList[DiskEncryptionKeyAttributes](e.ref.Append("disk_encryption_key"))
}

func (e EncryptionAttributes) KeyEncryptionKey() terra.ListValue[KeyEncryptionKeyAttributes] {
	return terra.ReferenceAsList[KeyEncryptionKeyAttributes](e.ref.Append("key_encryption_key"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type EncryptionState struct {
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
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
