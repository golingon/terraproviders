// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vpngatewayconnection

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Routing struct {
	// AssociatedRouteTable: string, required
	AssociatedRouteTable terra.StringValue `hcl:"associated_route_table,attr" validate:"required"`
	// InboundRouteMapId: string, optional
	InboundRouteMapId terra.StringValue `hcl:"inbound_route_map_id,attr"`
	// OutboundRouteMapId: string, optional
	OutboundRouteMapId terra.StringValue `hcl:"outbound_route_map_id,attr"`
	// PropagatedRouteTable: optional
	PropagatedRouteTable *PropagatedRouteTable `hcl:"propagated_route_table,block"`
}

type PropagatedRouteTable struct {
	// Labels: set of string, optional
	Labels terra.SetValue[terra.StringValue] `hcl:"labels,attr"`
	// RouteTableIds: list of string, required
	RouteTableIds terra.ListValue[terra.StringValue] `hcl:"route_table_ids,attr" validate:"required"`
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

type TrafficSelectorPolicy struct {
	// LocalAddressRanges: set of string, required
	LocalAddressRanges terra.SetValue[terra.StringValue] `hcl:"local_address_ranges,attr" validate:"required"`
	// RemoteAddressRanges: set of string, required
	RemoteAddressRanges terra.SetValue[terra.StringValue] `hcl:"remote_address_ranges,attr" validate:"required"`
}

type VpnLink struct {
	// BandwidthMbps: number, optional
	BandwidthMbps terra.NumberValue `hcl:"bandwidth_mbps,attr"`
	// BgpEnabled: bool, optional
	BgpEnabled terra.BoolValue `hcl:"bgp_enabled,attr"`
	// ConnectionMode: string, optional
	ConnectionMode terra.StringValue `hcl:"connection_mode,attr"`
	// EgressNatRuleIds: set of string, optional
	EgressNatRuleIds terra.SetValue[terra.StringValue] `hcl:"egress_nat_rule_ids,attr"`
	// IngressNatRuleIds: set of string, optional
	IngressNatRuleIds terra.SetValue[terra.StringValue] `hcl:"ingress_nat_rule_ids,attr"`
	// LocalAzureIpAddressEnabled: bool, optional
	LocalAzureIpAddressEnabled terra.BoolValue `hcl:"local_azure_ip_address_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyBasedTrafficSelectorEnabled: bool, optional
	PolicyBasedTrafficSelectorEnabled terra.BoolValue `hcl:"policy_based_traffic_selector_enabled,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// RatelimitEnabled: bool, optional
	RatelimitEnabled terra.BoolValue `hcl:"ratelimit_enabled,attr"`
	// RouteWeight: number, optional
	RouteWeight terra.NumberValue `hcl:"route_weight,attr"`
	// SharedKey: string, optional
	SharedKey terra.StringValue `hcl:"shared_key,attr"`
	// VpnSiteLinkId: string, required
	VpnSiteLinkId terra.StringValue `hcl:"vpn_site_link_id,attr" validate:"required"`
	// CustomBgpAddress: min=0
	CustomBgpAddress []CustomBgpAddress `hcl:"custom_bgp_address,block" validate:"min=0"`
	// IpsecPolicy: min=0
	IpsecPolicy []IpsecPolicy `hcl:"ipsec_policy,block" validate:"min=0"`
}

type CustomBgpAddress struct {
	// IpAddress: string, required
	IpAddress terra.StringValue `hcl:"ip_address,attr" validate:"required"`
	// IpConfigurationId: string, required
	IpConfigurationId terra.StringValue `hcl:"ip_configuration_id,attr" validate:"required"`
}

type IpsecPolicy struct {
	// DhGroup: string, required
	DhGroup terra.StringValue `hcl:"dh_group,attr" validate:"required"`
	// EncryptionAlgorithm: string, required
	EncryptionAlgorithm terra.StringValue `hcl:"encryption_algorithm,attr" validate:"required"`
	// IkeEncryptionAlgorithm: string, required
	IkeEncryptionAlgorithm terra.StringValue `hcl:"ike_encryption_algorithm,attr" validate:"required"`
	// IkeIntegrityAlgorithm: string, required
	IkeIntegrityAlgorithm terra.StringValue `hcl:"ike_integrity_algorithm,attr" validate:"required"`
	// IntegrityAlgorithm: string, required
	IntegrityAlgorithm terra.StringValue `hcl:"integrity_algorithm,attr" validate:"required"`
	// PfsGroup: string, required
	PfsGroup terra.StringValue `hcl:"pfs_group,attr" validate:"required"`
	// SaDataSizeKb: number, required
	SaDataSizeKb terra.NumberValue `hcl:"sa_data_size_kb,attr" validate:"required"`
	// SaLifetimeSec: number, required
	SaLifetimeSec terra.NumberValue `hcl:"sa_lifetime_sec,attr" validate:"required"`
}

type RoutingAttributes struct {
	ref terra.Reference
}

func (r RoutingAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RoutingAttributes) InternalWithRef(ref terra.Reference) RoutingAttributes {
	return RoutingAttributes{ref: ref}
}

func (r RoutingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RoutingAttributes) AssociatedRouteTable() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("associated_route_table"))
}

func (r RoutingAttributes) InboundRouteMapId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("inbound_route_map_id"))
}

func (r RoutingAttributes) OutboundRouteMapId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("outbound_route_map_id"))
}

func (r RoutingAttributes) PropagatedRouteTable() terra.ListValue[PropagatedRouteTableAttributes] {
	return terra.ReferenceAsList[PropagatedRouteTableAttributes](r.ref.Append("propagated_route_table"))
}

type PropagatedRouteTableAttributes struct {
	ref terra.Reference
}

func (prt PropagatedRouteTableAttributes) InternalRef() (terra.Reference, error) {
	return prt.ref, nil
}

func (prt PropagatedRouteTableAttributes) InternalWithRef(ref terra.Reference) PropagatedRouteTableAttributes {
	return PropagatedRouteTableAttributes{ref: ref}
}

func (prt PropagatedRouteTableAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return prt.ref.InternalTokens()
}

func (prt PropagatedRouteTableAttributes) Labels() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](prt.ref.Append("labels"))
}

func (prt PropagatedRouteTableAttributes) RouteTableIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](prt.ref.Append("route_table_ids"))
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

type TrafficSelectorPolicyAttributes struct {
	ref terra.Reference
}

func (tsp TrafficSelectorPolicyAttributes) InternalRef() (terra.Reference, error) {
	return tsp.ref, nil
}

func (tsp TrafficSelectorPolicyAttributes) InternalWithRef(ref terra.Reference) TrafficSelectorPolicyAttributes {
	return TrafficSelectorPolicyAttributes{ref: ref}
}

func (tsp TrafficSelectorPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tsp.ref.InternalTokens()
}

func (tsp TrafficSelectorPolicyAttributes) LocalAddressRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tsp.ref.Append("local_address_ranges"))
}

func (tsp TrafficSelectorPolicyAttributes) RemoteAddressRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tsp.ref.Append("remote_address_ranges"))
}

type VpnLinkAttributes struct {
	ref terra.Reference
}

func (vl VpnLinkAttributes) InternalRef() (terra.Reference, error) {
	return vl.ref, nil
}

func (vl VpnLinkAttributes) InternalWithRef(ref terra.Reference) VpnLinkAttributes {
	return VpnLinkAttributes{ref: ref}
}

func (vl VpnLinkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vl.ref.InternalTokens()
}

func (vl VpnLinkAttributes) BandwidthMbps() terra.NumberValue {
	return terra.ReferenceAsNumber(vl.ref.Append("bandwidth_mbps"))
}

func (vl VpnLinkAttributes) BgpEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vl.ref.Append("bgp_enabled"))
}

func (vl VpnLinkAttributes) ConnectionMode() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("connection_mode"))
}

func (vl VpnLinkAttributes) EgressNatRuleIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vl.ref.Append("egress_nat_rule_ids"))
}

func (vl VpnLinkAttributes) IngressNatRuleIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vl.ref.Append("ingress_nat_rule_ids"))
}

func (vl VpnLinkAttributes) LocalAzureIpAddressEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vl.ref.Append("local_azure_ip_address_enabled"))
}

func (vl VpnLinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("name"))
}

func (vl VpnLinkAttributes) PolicyBasedTrafficSelectorEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vl.ref.Append("policy_based_traffic_selector_enabled"))
}

func (vl VpnLinkAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("protocol"))
}

func (vl VpnLinkAttributes) RatelimitEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vl.ref.Append("ratelimit_enabled"))
}

func (vl VpnLinkAttributes) RouteWeight() terra.NumberValue {
	return terra.ReferenceAsNumber(vl.ref.Append("route_weight"))
}

func (vl VpnLinkAttributes) SharedKey() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("shared_key"))
}

func (vl VpnLinkAttributes) VpnSiteLinkId() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("vpn_site_link_id"))
}

func (vl VpnLinkAttributes) CustomBgpAddress() terra.SetValue[CustomBgpAddressAttributes] {
	return terra.ReferenceAsSet[CustomBgpAddressAttributes](vl.ref.Append("custom_bgp_address"))
}

func (vl VpnLinkAttributes) IpsecPolicy() terra.ListValue[IpsecPolicyAttributes] {
	return terra.ReferenceAsList[IpsecPolicyAttributes](vl.ref.Append("ipsec_policy"))
}

type CustomBgpAddressAttributes struct {
	ref terra.Reference
}

func (cba CustomBgpAddressAttributes) InternalRef() (terra.Reference, error) {
	return cba.ref, nil
}

func (cba CustomBgpAddressAttributes) InternalWithRef(ref terra.Reference) CustomBgpAddressAttributes {
	return CustomBgpAddressAttributes{ref: ref}
}

func (cba CustomBgpAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cba.ref.InternalTokens()
}

func (cba CustomBgpAddressAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(cba.ref.Append("ip_address"))
}

func (cba CustomBgpAddressAttributes) IpConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(cba.ref.Append("ip_configuration_id"))
}

type IpsecPolicyAttributes struct {
	ref terra.Reference
}

func (ip IpsecPolicyAttributes) InternalRef() (terra.Reference, error) {
	return ip.ref, nil
}

func (ip IpsecPolicyAttributes) InternalWithRef(ref terra.Reference) IpsecPolicyAttributes {
	return IpsecPolicyAttributes{ref: ref}
}

func (ip IpsecPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ip.ref.InternalTokens()
}

func (ip IpsecPolicyAttributes) DhGroup() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("dh_group"))
}

func (ip IpsecPolicyAttributes) EncryptionAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("encryption_algorithm"))
}

func (ip IpsecPolicyAttributes) IkeEncryptionAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("ike_encryption_algorithm"))
}

func (ip IpsecPolicyAttributes) IkeIntegrityAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("ike_integrity_algorithm"))
}

func (ip IpsecPolicyAttributes) IntegrityAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("integrity_algorithm"))
}

func (ip IpsecPolicyAttributes) PfsGroup() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("pfs_group"))
}

func (ip IpsecPolicyAttributes) SaDataSizeKb() terra.NumberValue {
	return terra.ReferenceAsNumber(ip.ref.Append("sa_data_size_kb"))
}

func (ip IpsecPolicyAttributes) SaLifetimeSec() terra.NumberValue {
	return terra.ReferenceAsNumber(ip.ref.Append("sa_lifetime_sec"))
}

type RoutingState struct {
	AssociatedRouteTable string                      `json:"associated_route_table"`
	InboundRouteMapId    string                      `json:"inbound_route_map_id"`
	OutboundRouteMapId   string                      `json:"outbound_route_map_id"`
	PropagatedRouteTable []PropagatedRouteTableState `json:"propagated_route_table"`
}

type PropagatedRouteTableState struct {
	Labels        []string `json:"labels"`
	RouteTableIds []string `json:"route_table_ids"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type TrafficSelectorPolicyState struct {
	LocalAddressRanges  []string `json:"local_address_ranges"`
	RemoteAddressRanges []string `json:"remote_address_ranges"`
}

type VpnLinkState struct {
	BandwidthMbps                     float64                 `json:"bandwidth_mbps"`
	BgpEnabled                        bool                    `json:"bgp_enabled"`
	ConnectionMode                    string                  `json:"connection_mode"`
	EgressNatRuleIds                  []string                `json:"egress_nat_rule_ids"`
	IngressNatRuleIds                 []string                `json:"ingress_nat_rule_ids"`
	LocalAzureIpAddressEnabled        bool                    `json:"local_azure_ip_address_enabled"`
	Name                              string                  `json:"name"`
	PolicyBasedTrafficSelectorEnabled bool                    `json:"policy_based_traffic_selector_enabled"`
	Protocol                          string                  `json:"protocol"`
	RatelimitEnabled                  bool                    `json:"ratelimit_enabled"`
	RouteWeight                       float64                 `json:"route_weight"`
	SharedKey                         string                  `json:"shared_key"`
	VpnSiteLinkId                     string                  `json:"vpn_site_link_id"`
	CustomBgpAddress                  []CustomBgpAddressState `json:"custom_bgp_address"`
	IpsecPolicy                       []IpsecPolicyState      `json:"ipsec_policy"`
}

type CustomBgpAddressState struct {
	IpAddress         string `json:"ip_address"`
	IpConfigurationId string `json:"ip_configuration_id"`
}

type IpsecPolicyState struct {
	DhGroup                string  `json:"dh_group"`
	EncryptionAlgorithm    string  `json:"encryption_algorithm"`
	IkeEncryptionAlgorithm string  `json:"ike_encryption_algorithm"`
	IkeIntegrityAlgorithm  string  `json:"ike_integrity_algorithm"`
	IntegrityAlgorithm     string  `json:"integrity_algorithm"`
	PfsGroup               string  `json:"pfs_group"`
	SaDataSizeKb           float64 `json:"sa_data_size_kb"`
	SaLifetimeSec          float64 `json:"sa_lifetime_sec"`
}