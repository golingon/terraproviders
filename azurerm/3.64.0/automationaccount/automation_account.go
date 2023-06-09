// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package automationaccount

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PrivateEndpointConnection struct{}

type Encryption struct {
	// KeySource: string, optional
	KeySource terra.StringValue `hcl:"key_source,attr"`
	// KeyVaultKeyId: string, required
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr" validate:"required"`
	// UserAssignedIdentityId: string, optional
	UserAssignedIdentityId terra.StringValue `hcl:"user_assigned_identity_id,attr"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
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

type PrivateEndpointConnectionAttributes struct {
	ref terra.Reference
}

func (pec PrivateEndpointConnectionAttributes) InternalRef() (terra.Reference, error) {
	return pec.ref, nil
}

func (pec PrivateEndpointConnectionAttributes) InternalWithRef(ref terra.Reference) PrivateEndpointConnectionAttributes {
	return PrivateEndpointConnectionAttributes{ref: ref}
}

func (pec PrivateEndpointConnectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pec.ref.InternalTokens()
}

func (pec PrivateEndpointConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("id"))
}

func (pec PrivateEndpointConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("name"))
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

func (e EncryptionAttributes) KeySource() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("key_source"))
}

func (e EncryptionAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("key_vault_key_id"))
}

func (e EncryptionAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("user_assigned_identity_id"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
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

type PrivateEndpointConnectionState struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type EncryptionState struct {
	KeySource              string `json:"key_source"`
	KeyVaultKeyId          string `json:"key_vault_key_id"`
	UserAssignedIdentityId string `json:"user_assigned_identity_id"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
