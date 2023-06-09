// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mediaservicesaccount

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Encryption struct {
	// KeyVaultKeyIdentifier: string, optional
	KeyVaultKeyIdentifier terra.StringValue `hcl:"key_vault_key_identifier,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// EncryptionManagedIdentity: optional
	ManagedIdentity *EncryptionManagedIdentity `hcl:"managed_identity,block"`
}

type EncryptionManagedIdentity struct {
	// UseSystemAssignedIdentity: bool, optional
	UseSystemAssignedIdentity terra.BoolValue `hcl:"use_system_assigned_identity,attr"`
	// UserAssignedIdentityId: string, optional
	UserAssignedIdentityId terra.StringValue `hcl:"user_assigned_identity_id,attr"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type KeyDeliveryAccessControl struct {
	// DefaultAction: string, optional
	DefaultAction terra.StringValue `hcl:"default_action,attr"`
	// IpAllowList: set of string, optional
	IpAllowList terra.SetValue[terra.StringValue] `hcl:"ip_allow_list,attr"`
}

type StorageAccount struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// IsPrimary: bool, optional
	IsPrimary terra.BoolValue `hcl:"is_primary,attr"`
	// StorageAccountManagedIdentity: optional
	ManagedIdentity *StorageAccountManagedIdentity `hcl:"managed_identity,block"`
}

type StorageAccountManagedIdentity struct {
	// UseSystemAssignedIdentity: bool, optional
	UseSystemAssignedIdentity terra.BoolValue `hcl:"use_system_assigned_identity,attr"`
	// UserAssignedIdentityId: string, optional
	UserAssignedIdentityId terra.StringValue `hcl:"user_assigned_identity_id,attr"`
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

func (e EncryptionAttributes) CurrentKeyIdentifier() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("current_key_identifier"))
}

func (e EncryptionAttributes) KeyVaultKeyIdentifier() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("key_vault_key_identifier"))
}

func (e EncryptionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("type"))
}

func (e EncryptionAttributes) ManagedIdentity() terra.ListValue[EncryptionManagedIdentityAttributes] {
	return terra.ReferenceAsList[EncryptionManagedIdentityAttributes](e.ref.Append("managed_identity"))
}

type EncryptionManagedIdentityAttributes struct {
	ref terra.Reference
}

func (mi EncryptionManagedIdentityAttributes) InternalRef() (terra.Reference, error) {
	return mi.ref, nil
}

func (mi EncryptionManagedIdentityAttributes) InternalWithRef(ref terra.Reference) EncryptionManagedIdentityAttributes {
	return EncryptionManagedIdentityAttributes{ref: ref}
}

func (mi EncryptionManagedIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mi.ref.InternalTokens()
}

func (mi EncryptionManagedIdentityAttributes) UseSystemAssignedIdentity() terra.BoolValue {
	return terra.ReferenceAsBool(mi.ref.Append("use_system_assigned_identity"))
}

func (mi EncryptionManagedIdentityAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("user_assigned_identity_id"))
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

type KeyDeliveryAccessControlAttributes struct {
	ref terra.Reference
}

func (kdac KeyDeliveryAccessControlAttributes) InternalRef() (terra.Reference, error) {
	return kdac.ref, nil
}

func (kdac KeyDeliveryAccessControlAttributes) InternalWithRef(ref terra.Reference) KeyDeliveryAccessControlAttributes {
	return KeyDeliveryAccessControlAttributes{ref: ref}
}

func (kdac KeyDeliveryAccessControlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kdac.ref.InternalTokens()
}

func (kdac KeyDeliveryAccessControlAttributes) DefaultAction() terra.StringValue {
	return terra.ReferenceAsString(kdac.ref.Append("default_action"))
}

func (kdac KeyDeliveryAccessControlAttributes) IpAllowList() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kdac.ref.Append("ip_allow_list"))
}

type StorageAccountAttributes struct {
	ref terra.Reference
}

func (sa StorageAccountAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa StorageAccountAttributes) InternalWithRef(ref terra.Reference) StorageAccountAttributes {
	return StorageAccountAttributes{ref: ref}
}

func (sa StorageAccountAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa StorageAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("id"))
}

func (sa StorageAccountAttributes) IsPrimary() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("is_primary"))
}

func (sa StorageAccountAttributes) ManagedIdentity() terra.ListValue[StorageAccountManagedIdentityAttributes] {
	return terra.ReferenceAsList[StorageAccountManagedIdentityAttributes](sa.ref.Append("managed_identity"))
}

type StorageAccountManagedIdentityAttributes struct {
	ref terra.Reference
}

func (mi StorageAccountManagedIdentityAttributes) InternalRef() (terra.Reference, error) {
	return mi.ref, nil
}

func (mi StorageAccountManagedIdentityAttributes) InternalWithRef(ref terra.Reference) StorageAccountManagedIdentityAttributes {
	return StorageAccountManagedIdentityAttributes{ref: ref}
}

func (mi StorageAccountManagedIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mi.ref.InternalTokens()
}

func (mi StorageAccountManagedIdentityAttributes) UseSystemAssignedIdentity() terra.BoolValue {
	return terra.ReferenceAsBool(mi.ref.Append("use_system_assigned_identity"))
}

func (mi StorageAccountManagedIdentityAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("user_assigned_identity_id"))
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
	CurrentKeyIdentifier  string                           `json:"current_key_identifier"`
	KeyVaultKeyIdentifier string                           `json:"key_vault_key_identifier"`
	Type                  string                           `json:"type"`
	ManagedIdentity       []EncryptionManagedIdentityState `json:"managed_identity"`
}

type EncryptionManagedIdentityState struct {
	UseSystemAssignedIdentity bool   `json:"use_system_assigned_identity"`
	UserAssignedIdentityId    string `json:"user_assigned_identity_id"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type KeyDeliveryAccessControlState struct {
	DefaultAction string   `json:"default_action"`
	IpAllowList   []string `json:"ip_allow_list"`
}

type StorageAccountState struct {
	Id              string                               `json:"id"`
	IsPrimary       bool                                 `json:"is_primary"`
	ManagedIdentity []StorageAccountManagedIdentityState `json:"managed_identity"`
}

type StorageAccountManagedIdentityState struct {
	UseSystemAssignedIdentity bool   `json:"use_system_assigned_identity"`
	UserAssignedIdentityId    string `json:"user_assigned_identity_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
