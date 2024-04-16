// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_servicebus_namespace

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CustomerManagedKey struct {
	// IdentityId: string, required
	IdentityId terra.StringValue `hcl:"identity_id,attr" validate:"required"`
	// InfrastructureEncryptionEnabled: bool, optional
	InfrastructureEncryptionEnabled terra.BoolValue `hcl:"infrastructure_encryption_enabled,attr"`
	// KeyVaultKeyId: string, required
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr" validate:"required"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type NetworkRuleSet struct {
	// DefaultAction: string, optional
	DefaultAction terra.StringValue `hcl:"default_action,attr"`
	// IpRules: set of string, optional
	IpRules terra.SetValue[terra.StringValue] `hcl:"ip_rules,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// TrustedServicesAllowed: bool, optional
	TrustedServicesAllowed terra.BoolValue `hcl:"trusted_services_allowed,attr"`
	// NetworkRuleSetNetworkRules: min=0
	NetworkRules []NetworkRuleSetNetworkRules `hcl:"network_rules,block" validate:"min=0"`
}

type NetworkRuleSetNetworkRules struct {
	// IgnoreMissingVnetServiceEndpoint: bool, optional
	IgnoreMissingVnetServiceEndpoint terra.BoolValue `hcl:"ignore_missing_vnet_service_endpoint,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
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

func (cmk CustomerManagedKeyAttributes) IdentityId() terra.StringValue {
	return terra.ReferenceAsString(cmk.ref.Append("identity_id"))
}

func (cmk CustomerManagedKeyAttributes) InfrastructureEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cmk.ref.Append("infrastructure_encryption_enabled"))
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

type NetworkRuleSetAttributes struct {
	ref terra.Reference
}

func (nrs NetworkRuleSetAttributes) InternalRef() (terra.Reference, error) {
	return nrs.ref, nil
}

func (nrs NetworkRuleSetAttributes) InternalWithRef(ref terra.Reference) NetworkRuleSetAttributes {
	return NetworkRuleSetAttributes{ref: ref}
}

func (nrs NetworkRuleSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nrs.ref.InternalTokens()
}

func (nrs NetworkRuleSetAttributes) DefaultAction() terra.StringValue {
	return terra.ReferenceAsString(nrs.ref.Append("default_action"))
}

func (nrs NetworkRuleSetAttributes) IpRules() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nrs.ref.Append("ip_rules"))
}

func (nrs NetworkRuleSetAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(nrs.ref.Append("public_network_access_enabled"))
}

func (nrs NetworkRuleSetAttributes) TrustedServicesAllowed() terra.BoolValue {
	return terra.ReferenceAsBool(nrs.ref.Append("trusted_services_allowed"))
}

func (nrs NetworkRuleSetAttributes) NetworkRules() terra.SetValue[NetworkRuleSetNetworkRulesAttributes] {
	return terra.ReferenceAsSet[NetworkRuleSetNetworkRulesAttributes](nrs.ref.Append("network_rules"))
}

type NetworkRuleSetNetworkRulesAttributes struct {
	ref terra.Reference
}

func (nr NetworkRuleSetNetworkRulesAttributes) InternalRef() (terra.Reference, error) {
	return nr.ref, nil
}

func (nr NetworkRuleSetNetworkRulesAttributes) InternalWithRef(ref terra.Reference) NetworkRuleSetNetworkRulesAttributes {
	return NetworkRuleSetNetworkRulesAttributes{ref: ref}
}

func (nr NetworkRuleSetNetworkRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nr.ref.InternalTokens()
}

func (nr NetworkRuleSetNetworkRulesAttributes) IgnoreMissingVnetServiceEndpoint() terra.BoolValue {
	return terra.ReferenceAsBool(nr.ref.Append("ignore_missing_vnet_service_endpoint"))
}

func (nr NetworkRuleSetNetworkRulesAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(nr.ref.Append("subnet_id"))
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
	IdentityId                      string `json:"identity_id"`
	InfrastructureEncryptionEnabled bool   `json:"infrastructure_encryption_enabled"`
	KeyVaultKeyId                   string `json:"key_vault_key_id"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type NetworkRuleSetState struct {
	DefaultAction              string                            `json:"default_action"`
	IpRules                    []string                          `json:"ip_rules"`
	PublicNetworkAccessEnabled bool                              `json:"public_network_access_enabled"`
	TrustedServicesAllowed     bool                              `json:"trusted_services_allowed"`
	NetworkRules               []NetworkRuleSetNetworkRulesState `json:"network_rules"`
}

type NetworkRuleSetNetworkRulesState struct {
	IgnoreMissingVnetServiceEndpoint bool   `json:"ignore_missing_vnet_service_endpoint"`
	SubnetId                         string `json:"subnet_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
