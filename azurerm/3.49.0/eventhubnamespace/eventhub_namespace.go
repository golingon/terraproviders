// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package eventhubnamespace

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type NetworkRulesets struct {
	// DefaultAction: string, optional
	DefaultAction terra.StringValue `hcl:"default_action,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// TrustedServiceAccessEnabled: bool, optional
	TrustedServiceAccessEnabled terra.BoolValue `hcl:"trusted_service_access_enabled,attr"`
	// IpRule: min=0
	IpRule []IpRule `hcl:"ip_rule,block" validate:"min=0"`
	// VirtualNetworkRule: min=0
	VirtualNetworkRule []VirtualNetworkRule `hcl:"virtual_network_rule,block" validate:"min=0"`
}

type IpRule struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// IpMask: string, optional
	IpMask terra.StringValue `hcl:"ip_mask,attr"`
}

type VirtualNetworkRule struct {
	// IgnoreMissingVirtualNetworkServiceEndpoint: bool, optional
	IgnoreMissingVirtualNetworkServiceEndpoint terra.BoolValue `hcl:"ignore_missing_virtual_network_service_endpoint,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
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

type NetworkRulesetsAttributes struct {
	ref terra.Reference
}

func (nr NetworkRulesetsAttributes) InternalRef() terra.Reference {
	return nr.ref
}

func (nr NetworkRulesetsAttributes) InternalWithRef(ref terra.Reference) NetworkRulesetsAttributes {
	return NetworkRulesetsAttributes{ref: ref}
}

func (nr NetworkRulesetsAttributes) InternalTokens() hclwrite.Tokens {
	return nr.ref.InternalTokens()
}

func (nr NetworkRulesetsAttributes) DefaultAction() terra.StringValue {
	return terra.ReferenceString(nr.ref.Append("default_action"))
}

func (nr NetworkRulesetsAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceBool(nr.ref.Append("public_network_access_enabled"))
}

func (nr NetworkRulesetsAttributes) TrustedServiceAccessEnabled() terra.BoolValue {
	return terra.ReferenceBool(nr.ref.Append("trusted_service_access_enabled"))
}

func (nr NetworkRulesetsAttributes) IpRule() terra.ListValue[IpRuleAttributes] {
	return terra.ReferenceList[IpRuleAttributes](nr.ref.Append("ip_rule"))
}

func (nr NetworkRulesetsAttributes) VirtualNetworkRule() terra.SetValue[VirtualNetworkRuleAttributes] {
	return terra.ReferenceSet[VirtualNetworkRuleAttributes](nr.ref.Append("virtual_network_rule"))
}

type IpRuleAttributes struct {
	ref terra.Reference
}

func (ir IpRuleAttributes) InternalRef() terra.Reference {
	return ir.ref
}

func (ir IpRuleAttributes) InternalWithRef(ref terra.Reference) IpRuleAttributes {
	return IpRuleAttributes{ref: ref}
}

func (ir IpRuleAttributes) InternalTokens() hclwrite.Tokens {
	return ir.ref.InternalTokens()
}

func (ir IpRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceString(ir.ref.Append("action"))
}

func (ir IpRuleAttributes) IpMask() terra.StringValue {
	return terra.ReferenceString(ir.ref.Append("ip_mask"))
}

type VirtualNetworkRuleAttributes struct {
	ref terra.Reference
}

func (vnr VirtualNetworkRuleAttributes) InternalRef() terra.Reference {
	return vnr.ref
}

func (vnr VirtualNetworkRuleAttributes) InternalWithRef(ref terra.Reference) VirtualNetworkRuleAttributes {
	return VirtualNetworkRuleAttributes{ref: ref}
}

func (vnr VirtualNetworkRuleAttributes) InternalTokens() hclwrite.Tokens {
	return vnr.ref.InternalTokens()
}

func (vnr VirtualNetworkRuleAttributes) IgnoreMissingVirtualNetworkServiceEndpoint() terra.BoolValue {
	return terra.ReferenceBool(vnr.ref.Append("ignore_missing_virtual_network_service_endpoint"))
}

func (vnr VirtualNetworkRuleAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(vnr.ref.Append("subnet_id"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() terra.Reference {
	return i.ref
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() hclwrite.Tokens {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("type"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type NetworkRulesetsState struct {
	DefaultAction               string                    `json:"default_action"`
	PublicNetworkAccessEnabled  bool                      `json:"public_network_access_enabled"`
	TrustedServiceAccessEnabled bool                      `json:"trusted_service_access_enabled"`
	IpRule                      []IpRuleState             `json:"ip_rule"`
	VirtualNetworkRule          []VirtualNetworkRuleState `json:"virtual_network_rule"`
}

type IpRuleState struct {
	Action string `json:"action"`
	IpMask string `json:"ip_mask"`
}

type VirtualNetworkRuleState struct {
	IgnoreMissingVirtualNetworkServiceEndpoint bool   `json:"ignore_missing_virtual_network_service_endpoint"`
	SubnetId                                   string `json:"subnet_id"`
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
