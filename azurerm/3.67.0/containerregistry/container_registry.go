// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package containerregistry

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Encryption struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// IdentityClientId: string, optional
	IdentityClientId terra.StringValue `hcl:"identity_client_id,attr"`
	// KeyVaultKeyId: string, optional
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr"`
}

type NetworkRuleSet struct {
	// DefaultAction: string, optional
	DefaultAction terra.StringValue `hcl:"default_action,attr"`
	// IpRule: min=0
	IpRule []IpRule `hcl:"ip_rule,block" validate:"min=0"`
	// VirtualNetwork: min=0
	VirtualNetwork []VirtualNetwork `hcl:"virtual_network,block" validate:"min=0"`
}

type IpRule struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// IpRange: string, optional
	IpRange terra.StringValue `hcl:"ip_range,attr"`
}

type VirtualNetwork struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
}

type RetentionPolicy struct {
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type TrustPolicy struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type Georeplications struct {
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// RegionalEndpointEnabled: bool, optional
	RegionalEndpointEnabled terra.BoolValue `hcl:"regional_endpoint_enabled,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ZoneRedundancyEnabled: bool, optional
	ZoneRedundancyEnabled terra.BoolValue `hcl:"zone_redundancy_enabled,attr"`
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

func (e EncryptionAttributes) IdentityClientId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("identity_client_id"))
}

func (e EncryptionAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("key_vault_key_id"))
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

func (nrs NetworkRuleSetAttributes) IpRule() terra.SetValue[IpRuleAttributes] {
	return terra.ReferenceAsSet[IpRuleAttributes](nrs.ref.Append("ip_rule"))
}

func (nrs NetworkRuleSetAttributes) VirtualNetwork() terra.SetValue[VirtualNetworkAttributes] {
	return terra.ReferenceAsSet[VirtualNetworkAttributes](nrs.ref.Append("virtual_network"))
}

type IpRuleAttributes struct {
	ref terra.Reference
}

func (ir IpRuleAttributes) InternalRef() (terra.Reference, error) {
	return ir.ref, nil
}

func (ir IpRuleAttributes) InternalWithRef(ref terra.Reference) IpRuleAttributes {
	return IpRuleAttributes{ref: ref}
}

func (ir IpRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ir.ref.InternalTokens()
}

func (ir IpRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("action"))
}

func (ir IpRuleAttributes) IpRange() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("ip_range"))
}

type VirtualNetworkAttributes struct {
	ref terra.Reference
}

func (vn VirtualNetworkAttributes) InternalRef() (terra.Reference, error) {
	return vn.ref, nil
}

func (vn VirtualNetworkAttributes) InternalWithRef(ref terra.Reference) VirtualNetworkAttributes {
	return VirtualNetworkAttributes{ref: ref}
}

func (vn VirtualNetworkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vn.ref.InternalTokens()
}

func (vn VirtualNetworkAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("action"))
}

func (vn VirtualNetworkAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("subnet_id"))
}

type RetentionPolicyAttributes struct {
	ref terra.Reference
}

func (rp RetentionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp RetentionPolicyAttributes) InternalWithRef(ref terra.Reference) RetentionPolicyAttributes {
	return RetentionPolicyAttributes{ref: ref}
}

func (rp RetentionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp RetentionPolicyAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("days"))
}

func (rp RetentionPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("enabled"))
}

type TrustPolicyAttributes struct {
	ref terra.Reference
}

func (tp TrustPolicyAttributes) InternalRef() (terra.Reference, error) {
	return tp.ref, nil
}

func (tp TrustPolicyAttributes) InternalWithRef(ref terra.Reference) TrustPolicyAttributes {
	return TrustPolicyAttributes{ref: ref}
}

func (tp TrustPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tp.ref.InternalTokens()
}

func (tp TrustPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(tp.ref.Append("enabled"))
}

type GeoreplicationsAttributes struct {
	ref terra.Reference
}

func (g GeoreplicationsAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GeoreplicationsAttributes) InternalWithRef(ref terra.Reference) GeoreplicationsAttributes {
	return GeoreplicationsAttributes{ref: ref}
}

func (g GeoreplicationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GeoreplicationsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("location"))
}

func (g GeoreplicationsAttributes) RegionalEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("regional_endpoint_enabled"))
}

func (g GeoreplicationsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](g.ref.Append("tags"))
}

func (g GeoreplicationsAttributes) ZoneRedundancyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("zone_redundancy_enabled"))
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

type EncryptionState struct {
	Enabled          bool   `json:"enabled"`
	IdentityClientId string `json:"identity_client_id"`
	KeyVaultKeyId    string `json:"key_vault_key_id"`
}

type NetworkRuleSetState struct {
	DefaultAction  string                `json:"default_action"`
	IpRule         []IpRuleState         `json:"ip_rule"`
	VirtualNetwork []VirtualNetworkState `json:"virtual_network"`
}

type IpRuleState struct {
	Action  string `json:"action"`
	IpRange string `json:"ip_range"`
}

type VirtualNetworkState struct {
	Action   string `json:"action"`
	SubnetId string `json:"subnet_id"`
}

type RetentionPolicyState struct {
	Days    float64 `json:"days"`
	Enabled bool    `json:"enabled"`
}

type TrustPolicyState struct {
	Enabled bool `json:"enabled"`
}

type GeoreplicationsState struct {
	Location                string            `json:"location"`
	RegionalEndpointEnabled bool              `json:"regional_endpoint_enabled"`
	Tags                    map[string]string `json:"tags"`
	ZoneRedundancyEnabled   bool              `json:"zone_redundancy_enabled"`
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
