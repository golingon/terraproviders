// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cognitiveaccount

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CustomerManagedKey struct {
	// IdentityClientId: string, optional
	IdentityClientId terra.StringValue `hcl:"identity_client_id,attr"`
	// KeyVaultKeyId: string, required
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr" validate:"required"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type NetworkAcls struct {
	// DefaultAction: string, required
	DefaultAction terra.StringValue `hcl:"default_action,attr" validate:"required"`
	// IpRules: set of string, optional
	IpRules terra.SetValue[terra.StringValue] `hcl:"ip_rules,attr"`
	// VirtualNetworkRules: min=0
	VirtualNetworkRules []VirtualNetworkRules `hcl:"virtual_network_rules,block" validate:"min=0"`
}

type VirtualNetworkRules struct {
	// IgnoreMissingVnetServiceEndpoint: bool, optional
	IgnoreMissingVnetServiceEndpoint terra.BoolValue `hcl:"ignore_missing_vnet_service_endpoint,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
}

type Storage struct {
	// IdentityClientId: string, optional
	IdentityClientId terra.StringValue `hcl:"identity_client_id,attr"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
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

type CustomerManagedKeyAttributes struct {
	ref terra.Reference
}

func (cmk CustomerManagedKeyAttributes) InternalRef() (terra.Reference, error) {
	return cmk.ref, nil
}

func (cmk CustomerManagedKeyAttributes) InternalWithRef(ref terra.Reference) CustomerManagedKeyAttributes {
	return CustomerManagedKeyAttributes{ref: ref}
}

func (cmk CustomerManagedKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cmk.ref.InternalTokens()
}

func (cmk CustomerManagedKeyAttributes) IdentityClientId() terra.StringValue {
	return terra.ReferenceAsString(cmk.ref.Append("identity_client_id"))
}

func (cmk CustomerManagedKeyAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(cmk.ref.Append("key_vault_key_id"))
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

type NetworkAclsAttributes struct {
	ref terra.Reference
}

func (na NetworkAclsAttributes) InternalRef() (terra.Reference, error) {
	return na.ref, nil
}

func (na NetworkAclsAttributes) InternalWithRef(ref terra.Reference) NetworkAclsAttributes {
	return NetworkAclsAttributes{ref: ref}
}

func (na NetworkAclsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return na.ref.InternalTokens()
}

func (na NetworkAclsAttributes) DefaultAction() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("default_action"))
}

func (na NetworkAclsAttributes) IpRules() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](na.ref.Append("ip_rules"))
}

func (na NetworkAclsAttributes) VirtualNetworkRules() terra.SetValue[VirtualNetworkRulesAttributes] {
	return terra.ReferenceAsSet[VirtualNetworkRulesAttributes](na.ref.Append("virtual_network_rules"))
}

type VirtualNetworkRulesAttributes struct {
	ref terra.Reference
}

func (vnr VirtualNetworkRulesAttributes) InternalRef() (terra.Reference, error) {
	return vnr.ref, nil
}

func (vnr VirtualNetworkRulesAttributes) InternalWithRef(ref terra.Reference) VirtualNetworkRulesAttributes {
	return VirtualNetworkRulesAttributes{ref: ref}
}

func (vnr VirtualNetworkRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vnr.ref.InternalTokens()
}

func (vnr VirtualNetworkRulesAttributes) IgnoreMissingVnetServiceEndpoint() terra.BoolValue {
	return terra.ReferenceAsBool(vnr.ref.Append("ignore_missing_vnet_service_endpoint"))
}

func (vnr VirtualNetworkRulesAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(vnr.ref.Append("subnet_id"))
}

type StorageAttributes struct {
	ref terra.Reference
}

func (s StorageAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StorageAttributes) InternalWithRef(ref terra.Reference) StorageAttributes {
	return StorageAttributes{ref: ref}
}

func (s StorageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StorageAttributes) IdentityClientId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("identity_client_id"))
}

func (s StorageAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("storage_account_id"))
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

type CustomerManagedKeyState struct {
	IdentityClientId string `json:"identity_client_id"`
	KeyVaultKeyId    string `json:"key_vault_key_id"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type NetworkAclsState struct {
	DefaultAction       string                     `json:"default_action"`
	IpRules             []string                   `json:"ip_rules"`
	VirtualNetworkRules []VirtualNetworkRulesState `json:"virtual_network_rules"`
}

type VirtualNetworkRulesState struct {
	IgnoreMissingVnetServiceEndpoint bool   `json:"ignore_missing_vnet_service_endpoint"`
	SubnetId                         string `json:"subnet_id"`
}

type StorageState struct {
	IdentityClientId string `json:"identity_client_id"`
	StorageAccountId string `json:"storage_account_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}