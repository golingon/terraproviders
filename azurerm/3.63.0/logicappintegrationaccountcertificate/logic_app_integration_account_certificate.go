// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package logicappintegrationaccountcertificate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type KeyVaultKey struct {
	// KeyName: string, required
	KeyName terra.StringValue `hcl:"key_name,attr" validate:"required"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// KeyVersion: string, optional
	KeyVersion terra.StringValue `hcl:"key_version,attr"`
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

type KeyVaultKeyAttributes struct {
	ref terra.Reference
}

func (kvk KeyVaultKeyAttributes) InternalRef() (terra.Reference, error) {
	return kvk.ref, nil
}

func (kvk KeyVaultKeyAttributes) InternalWithRef(ref terra.Reference) KeyVaultKeyAttributes {
	return KeyVaultKeyAttributes{ref: ref}
}

func (kvk KeyVaultKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kvk.ref.InternalTokens()
}

func (kvk KeyVaultKeyAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("key_name"))
}

func (kvk KeyVaultKeyAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("key_vault_id"))
}

func (kvk KeyVaultKeyAttributes) KeyVersion() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("key_version"))
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

type KeyVaultKeyState struct {
	KeyName    string `json:"key_name"`
	KeyVaultId string `json:"key_vault_id"`
	KeyVersion string `json:"key_version"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
