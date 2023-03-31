// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpnconnection "github.com/golingon/terraproviders/aws/4.60.0/vpnconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpnConnection creates a new instance of [VpnConnection].
func NewVpnConnection(name string, args VpnConnectionArgs) *VpnConnection {
	return &VpnConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpnConnection)(nil)

// VpnConnection represents the Terraform resource aws_vpn_connection.
type VpnConnection struct {
	Name      string
	Args      VpnConnectionArgs
	state     *vpnConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpnConnection].
func (vc *VpnConnection) Type() string {
	return "aws_vpn_connection"
}

// LocalName returns the local name for [VpnConnection].
func (vc *VpnConnection) LocalName() string {
	return vc.Name
}

// Configuration returns the configuration (args) for [VpnConnection].
func (vc *VpnConnection) Configuration() interface{} {
	return vc.Args
}

// DependOn is used for other resources to depend on [VpnConnection].
func (vc *VpnConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(vc)
}

// Dependencies returns the list of resources [VpnConnection] depends_on.
func (vc *VpnConnection) Dependencies() terra.Dependencies {
	return vc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpnConnection].
func (vc *VpnConnection) LifecycleManagement() *terra.Lifecycle {
	return vc.Lifecycle
}

// Attributes returns the attributes for [VpnConnection].
func (vc *VpnConnection) Attributes() vpnConnectionAttributes {
	return vpnConnectionAttributes{ref: terra.ReferenceResource(vc)}
}

// ImportState imports the given attribute values into [VpnConnection]'s state.
func (vc *VpnConnection) ImportState(av io.Reader) error {
	vc.state = &vpnConnectionState{}
	if err := json.NewDecoder(av).Decode(vc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vc.Type(), vc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpnConnection] has state.
func (vc *VpnConnection) State() (*vpnConnectionState, bool) {
	return vc.state, vc.state != nil
}

// StateMust returns the state for [VpnConnection]. Panics if the state is nil.
func (vc *VpnConnection) StateMust() *vpnConnectionState {
	if vc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vc.Type(), vc.LocalName()))
	}
	return vc.state
}

// VpnConnectionArgs contains the configurations for aws_vpn_connection.
type VpnConnectionArgs struct {
	// CustomerGatewayId: string, required
	CustomerGatewayId terra.StringValue `hcl:"customer_gateway_id,attr" validate:"required"`
	// EnableAcceleration: bool, optional
	EnableAcceleration terra.BoolValue `hcl:"enable_acceleration,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalIpv4NetworkCidr: string, optional
	LocalIpv4NetworkCidr terra.StringValue `hcl:"local_ipv4_network_cidr,attr"`
	// LocalIpv6NetworkCidr: string, optional
	LocalIpv6NetworkCidr terra.StringValue `hcl:"local_ipv6_network_cidr,attr"`
	// OutsideIpAddressType: string, optional
	OutsideIpAddressType terra.StringValue `hcl:"outside_ip_address_type,attr"`
	// RemoteIpv4NetworkCidr: string, optional
	RemoteIpv4NetworkCidr terra.StringValue `hcl:"remote_ipv4_network_cidr,attr"`
	// RemoteIpv6NetworkCidr: string, optional
	RemoteIpv6NetworkCidr terra.StringValue `hcl:"remote_ipv6_network_cidr,attr"`
	// StaticRoutesOnly: bool, optional
	StaticRoutesOnly terra.BoolValue `hcl:"static_routes_only,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitGatewayId: string, optional
	TransitGatewayId terra.StringValue `hcl:"transit_gateway_id,attr"`
	// TransportTransitGatewayAttachmentId: string, optional
	TransportTransitGatewayAttachmentId terra.StringValue `hcl:"transport_transit_gateway_attachment_id,attr"`
	// Tunnel1DpdTimeoutAction: string, optional
	Tunnel1DpdTimeoutAction terra.StringValue `hcl:"tunnel1_dpd_timeout_action,attr"`
	// Tunnel1DpdTimeoutSeconds: number, optional
	Tunnel1DpdTimeoutSeconds terra.NumberValue `hcl:"tunnel1_dpd_timeout_seconds,attr"`
	// Tunnel1IkeVersions: set of string, optional
	Tunnel1IkeVersions terra.SetValue[terra.StringValue] `hcl:"tunnel1_ike_versions,attr"`
	// Tunnel1InsideCidr: string, optional
	Tunnel1InsideCidr terra.StringValue `hcl:"tunnel1_inside_cidr,attr"`
	// Tunnel1InsideIpv6Cidr: string, optional
	Tunnel1InsideIpv6Cidr terra.StringValue `hcl:"tunnel1_inside_ipv6_cidr,attr"`
	// Tunnel1Phase1DhGroupNumbers: set of number, optional
	Tunnel1Phase1DhGroupNumbers terra.SetValue[terra.NumberValue] `hcl:"tunnel1_phase1_dh_group_numbers,attr"`
	// Tunnel1Phase1EncryptionAlgorithms: set of string, optional
	Tunnel1Phase1EncryptionAlgorithms terra.SetValue[terra.StringValue] `hcl:"tunnel1_phase1_encryption_algorithms,attr"`
	// Tunnel1Phase1IntegrityAlgorithms: set of string, optional
	Tunnel1Phase1IntegrityAlgorithms terra.SetValue[terra.StringValue] `hcl:"tunnel1_phase1_integrity_algorithms,attr"`
	// Tunnel1Phase1LifetimeSeconds: number, optional
	Tunnel1Phase1LifetimeSeconds terra.NumberValue `hcl:"tunnel1_phase1_lifetime_seconds,attr"`
	// Tunnel1Phase2DhGroupNumbers: set of number, optional
	Tunnel1Phase2DhGroupNumbers terra.SetValue[terra.NumberValue] `hcl:"tunnel1_phase2_dh_group_numbers,attr"`
	// Tunnel1Phase2EncryptionAlgorithms: set of string, optional
	Tunnel1Phase2EncryptionAlgorithms terra.SetValue[terra.StringValue] `hcl:"tunnel1_phase2_encryption_algorithms,attr"`
	// Tunnel1Phase2IntegrityAlgorithms: set of string, optional
	Tunnel1Phase2IntegrityAlgorithms terra.SetValue[terra.StringValue] `hcl:"tunnel1_phase2_integrity_algorithms,attr"`
	// Tunnel1Phase2LifetimeSeconds: number, optional
	Tunnel1Phase2LifetimeSeconds terra.NumberValue `hcl:"tunnel1_phase2_lifetime_seconds,attr"`
	// Tunnel1PresharedKey: string, optional
	Tunnel1PresharedKey terra.StringValue `hcl:"tunnel1_preshared_key,attr"`
	// Tunnel1RekeyFuzzPercentage: number, optional
	Tunnel1RekeyFuzzPercentage terra.NumberValue `hcl:"tunnel1_rekey_fuzz_percentage,attr"`
	// Tunnel1RekeyMarginTimeSeconds: number, optional
	Tunnel1RekeyMarginTimeSeconds terra.NumberValue `hcl:"tunnel1_rekey_margin_time_seconds,attr"`
	// Tunnel1ReplayWindowSize: number, optional
	Tunnel1ReplayWindowSize terra.NumberValue `hcl:"tunnel1_replay_window_size,attr"`
	// Tunnel1StartupAction: string, optional
	Tunnel1StartupAction terra.StringValue `hcl:"tunnel1_startup_action,attr"`
	// Tunnel2DpdTimeoutAction: string, optional
	Tunnel2DpdTimeoutAction terra.StringValue `hcl:"tunnel2_dpd_timeout_action,attr"`
	// Tunnel2DpdTimeoutSeconds: number, optional
	Tunnel2DpdTimeoutSeconds terra.NumberValue `hcl:"tunnel2_dpd_timeout_seconds,attr"`
	// Tunnel2IkeVersions: set of string, optional
	Tunnel2IkeVersions terra.SetValue[terra.StringValue] `hcl:"tunnel2_ike_versions,attr"`
	// Tunnel2InsideCidr: string, optional
	Tunnel2InsideCidr terra.StringValue `hcl:"tunnel2_inside_cidr,attr"`
	// Tunnel2InsideIpv6Cidr: string, optional
	Tunnel2InsideIpv6Cidr terra.StringValue `hcl:"tunnel2_inside_ipv6_cidr,attr"`
	// Tunnel2Phase1DhGroupNumbers: set of number, optional
	Tunnel2Phase1DhGroupNumbers terra.SetValue[terra.NumberValue] `hcl:"tunnel2_phase1_dh_group_numbers,attr"`
	// Tunnel2Phase1EncryptionAlgorithms: set of string, optional
	Tunnel2Phase1EncryptionAlgorithms terra.SetValue[terra.StringValue] `hcl:"tunnel2_phase1_encryption_algorithms,attr"`
	// Tunnel2Phase1IntegrityAlgorithms: set of string, optional
	Tunnel2Phase1IntegrityAlgorithms terra.SetValue[terra.StringValue] `hcl:"tunnel2_phase1_integrity_algorithms,attr"`
	// Tunnel2Phase1LifetimeSeconds: number, optional
	Tunnel2Phase1LifetimeSeconds terra.NumberValue `hcl:"tunnel2_phase1_lifetime_seconds,attr"`
	// Tunnel2Phase2DhGroupNumbers: set of number, optional
	Tunnel2Phase2DhGroupNumbers terra.SetValue[terra.NumberValue] `hcl:"tunnel2_phase2_dh_group_numbers,attr"`
	// Tunnel2Phase2EncryptionAlgorithms: set of string, optional
	Tunnel2Phase2EncryptionAlgorithms terra.SetValue[terra.StringValue] `hcl:"tunnel2_phase2_encryption_algorithms,attr"`
	// Tunnel2Phase2IntegrityAlgorithms: set of string, optional
	Tunnel2Phase2IntegrityAlgorithms terra.SetValue[terra.StringValue] `hcl:"tunnel2_phase2_integrity_algorithms,attr"`
	// Tunnel2Phase2LifetimeSeconds: number, optional
	Tunnel2Phase2LifetimeSeconds terra.NumberValue `hcl:"tunnel2_phase2_lifetime_seconds,attr"`
	// Tunnel2PresharedKey: string, optional
	Tunnel2PresharedKey terra.StringValue `hcl:"tunnel2_preshared_key,attr"`
	// Tunnel2RekeyFuzzPercentage: number, optional
	Tunnel2RekeyFuzzPercentage terra.NumberValue `hcl:"tunnel2_rekey_fuzz_percentage,attr"`
	// Tunnel2RekeyMarginTimeSeconds: number, optional
	Tunnel2RekeyMarginTimeSeconds terra.NumberValue `hcl:"tunnel2_rekey_margin_time_seconds,attr"`
	// Tunnel2ReplayWindowSize: number, optional
	Tunnel2ReplayWindowSize terra.NumberValue `hcl:"tunnel2_replay_window_size,attr"`
	// Tunnel2StartupAction: string, optional
	Tunnel2StartupAction terra.StringValue `hcl:"tunnel2_startup_action,attr"`
	// TunnelInsideIpVersion: string, optional
	TunnelInsideIpVersion terra.StringValue `hcl:"tunnel_inside_ip_version,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// VpnGatewayId: string, optional
	VpnGatewayId terra.StringValue `hcl:"vpn_gateway_id,attr"`
	// Routes: min=0
	Routes []vpnconnection.Routes `hcl:"routes,block" validate:"min=0"`
	// VgwTelemetry: min=0
	VgwTelemetry []vpnconnection.VgwTelemetry `hcl:"vgw_telemetry,block" validate:"min=0"`
	// Tunnel1LogOptions: optional
	Tunnel1LogOptions *vpnconnection.Tunnel1LogOptions `hcl:"tunnel1_log_options,block"`
	// Tunnel2LogOptions: optional
	Tunnel2LogOptions *vpnconnection.Tunnel2LogOptions `hcl:"tunnel2_log_options,block"`
}
type vpnConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpn_connection.
func (vc vpnConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("arn"))
}

// CoreNetworkArn returns a reference to field core_network_arn of aws_vpn_connection.
func (vc vpnConnectionAttributes) CoreNetworkArn() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("core_network_arn"))
}

// CoreNetworkAttachmentArn returns a reference to field core_network_attachment_arn of aws_vpn_connection.
func (vc vpnConnectionAttributes) CoreNetworkAttachmentArn() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("core_network_attachment_arn"))
}

// CustomerGatewayConfiguration returns a reference to field customer_gateway_configuration of aws_vpn_connection.
func (vc vpnConnectionAttributes) CustomerGatewayConfiguration() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("customer_gateway_configuration"))
}

// CustomerGatewayId returns a reference to field customer_gateway_id of aws_vpn_connection.
func (vc vpnConnectionAttributes) CustomerGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("customer_gateway_id"))
}

// EnableAcceleration returns a reference to field enable_acceleration of aws_vpn_connection.
func (vc vpnConnectionAttributes) EnableAcceleration() terra.BoolValue {
	return terra.ReferenceAsBool(vc.ref.Append("enable_acceleration"))
}

// Id returns a reference to field id of aws_vpn_connection.
func (vc vpnConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("id"))
}

// LocalIpv4NetworkCidr returns a reference to field local_ipv4_network_cidr of aws_vpn_connection.
func (vc vpnConnectionAttributes) LocalIpv4NetworkCidr() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("local_ipv4_network_cidr"))
}

// LocalIpv6NetworkCidr returns a reference to field local_ipv6_network_cidr of aws_vpn_connection.
func (vc vpnConnectionAttributes) LocalIpv6NetworkCidr() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("local_ipv6_network_cidr"))
}

// OutsideIpAddressType returns a reference to field outside_ip_address_type of aws_vpn_connection.
func (vc vpnConnectionAttributes) OutsideIpAddressType() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("outside_ip_address_type"))
}

// RemoteIpv4NetworkCidr returns a reference to field remote_ipv4_network_cidr of aws_vpn_connection.
func (vc vpnConnectionAttributes) RemoteIpv4NetworkCidr() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("remote_ipv4_network_cidr"))
}

// RemoteIpv6NetworkCidr returns a reference to field remote_ipv6_network_cidr of aws_vpn_connection.
func (vc vpnConnectionAttributes) RemoteIpv6NetworkCidr() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("remote_ipv6_network_cidr"))
}

// StaticRoutesOnly returns a reference to field static_routes_only of aws_vpn_connection.
func (vc vpnConnectionAttributes) StaticRoutesOnly() terra.BoolValue {
	return terra.ReferenceAsBool(vc.ref.Append("static_routes_only"))
}

// Tags returns a reference to field tags of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpn_connection.
func (vc vpnConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vc.ref.Append("tags_all"))
}

// TransitGatewayAttachmentId returns a reference to field transit_gateway_attachment_id of aws_vpn_connection.
func (vc vpnConnectionAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("transit_gateway_attachment_id"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_vpn_connection.
func (vc vpnConnectionAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("transit_gateway_id"))
}

// TransportTransitGatewayAttachmentId returns a reference to field transport_transit_gateway_attachment_id of aws_vpn_connection.
func (vc vpnConnectionAttributes) TransportTransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("transport_transit_gateway_attachment_id"))
}

// Tunnel1Address returns a reference to field tunnel1_address of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1Address() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel1_address"))
}

// Tunnel1BgpAsn returns a reference to field tunnel1_bgp_asn of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1BgpAsn() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel1_bgp_asn"))
}

// Tunnel1BgpHoldtime returns a reference to field tunnel1_bgp_holdtime of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1BgpHoldtime() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel1_bgp_holdtime"))
}

// Tunnel1CgwInsideAddress returns a reference to field tunnel1_cgw_inside_address of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1CgwInsideAddress() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel1_cgw_inside_address"))
}

// Tunnel1DpdTimeoutAction returns a reference to field tunnel1_dpd_timeout_action of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1DpdTimeoutAction() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel1_dpd_timeout_action"))
}

// Tunnel1DpdTimeoutSeconds returns a reference to field tunnel1_dpd_timeout_seconds of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1DpdTimeoutSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel1_dpd_timeout_seconds"))
}

// Tunnel1IkeVersions returns a reference to field tunnel1_ike_versions of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1IkeVersions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("tunnel1_ike_versions"))
}

// Tunnel1InsideCidr returns a reference to field tunnel1_inside_cidr of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1InsideCidr() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel1_inside_cidr"))
}

// Tunnel1InsideIpv6Cidr returns a reference to field tunnel1_inside_ipv6_cidr of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1InsideIpv6Cidr() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel1_inside_ipv6_cidr"))
}

// Tunnel1Phase1DhGroupNumbers returns a reference to field tunnel1_phase1_dh_group_numbers of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1Phase1DhGroupNumbers() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](vc.ref.Append("tunnel1_phase1_dh_group_numbers"))
}

// Tunnel1Phase1EncryptionAlgorithms returns a reference to field tunnel1_phase1_encryption_algorithms of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1Phase1EncryptionAlgorithms() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("tunnel1_phase1_encryption_algorithms"))
}

// Tunnel1Phase1IntegrityAlgorithms returns a reference to field tunnel1_phase1_integrity_algorithms of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1Phase1IntegrityAlgorithms() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("tunnel1_phase1_integrity_algorithms"))
}

// Tunnel1Phase1LifetimeSeconds returns a reference to field tunnel1_phase1_lifetime_seconds of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1Phase1LifetimeSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel1_phase1_lifetime_seconds"))
}

// Tunnel1Phase2DhGroupNumbers returns a reference to field tunnel1_phase2_dh_group_numbers of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1Phase2DhGroupNumbers() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](vc.ref.Append("tunnel1_phase2_dh_group_numbers"))
}

// Tunnel1Phase2EncryptionAlgorithms returns a reference to field tunnel1_phase2_encryption_algorithms of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1Phase2EncryptionAlgorithms() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("tunnel1_phase2_encryption_algorithms"))
}

// Tunnel1Phase2IntegrityAlgorithms returns a reference to field tunnel1_phase2_integrity_algorithms of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1Phase2IntegrityAlgorithms() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("tunnel1_phase2_integrity_algorithms"))
}

// Tunnel1Phase2LifetimeSeconds returns a reference to field tunnel1_phase2_lifetime_seconds of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1Phase2LifetimeSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel1_phase2_lifetime_seconds"))
}

// Tunnel1PresharedKey returns a reference to field tunnel1_preshared_key of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1PresharedKey() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel1_preshared_key"))
}

// Tunnel1RekeyFuzzPercentage returns a reference to field tunnel1_rekey_fuzz_percentage of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1RekeyFuzzPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel1_rekey_fuzz_percentage"))
}

// Tunnel1RekeyMarginTimeSeconds returns a reference to field tunnel1_rekey_margin_time_seconds of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1RekeyMarginTimeSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel1_rekey_margin_time_seconds"))
}

// Tunnel1ReplayWindowSize returns a reference to field tunnel1_replay_window_size of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1ReplayWindowSize() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel1_replay_window_size"))
}

// Tunnel1StartupAction returns a reference to field tunnel1_startup_action of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1StartupAction() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel1_startup_action"))
}

// Tunnel1VgwInsideAddress returns a reference to field tunnel1_vgw_inside_address of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel1VgwInsideAddress() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel1_vgw_inside_address"))
}

// Tunnel2Address returns a reference to field tunnel2_address of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2Address() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel2_address"))
}

// Tunnel2BgpAsn returns a reference to field tunnel2_bgp_asn of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2BgpAsn() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel2_bgp_asn"))
}

// Tunnel2BgpHoldtime returns a reference to field tunnel2_bgp_holdtime of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2BgpHoldtime() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel2_bgp_holdtime"))
}

// Tunnel2CgwInsideAddress returns a reference to field tunnel2_cgw_inside_address of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2CgwInsideAddress() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel2_cgw_inside_address"))
}

// Tunnel2DpdTimeoutAction returns a reference to field tunnel2_dpd_timeout_action of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2DpdTimeoutAction() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel2_dpd_timeout_action"))
}

// Tunnel2DpdTimeoutSeconds returns a reference to field tunnel2_dpd_timeout_seconds of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2DpdTimeoutSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel2_dpd_timeout_seconds"))
}

// Tunnel2IkeVersions returns a reference to field tunnel2_ike_versions of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2IkeVersions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("tunnel2_ike_versions"))
}

// Tunnel2InsideCidr returns a reference to field tunnel2_inside_cidr of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2InsideCidr() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel2_inside_cidr"))
}

// Tunnel2InsideIpv6Cidr returns a reference to field tunnel2_inside_ipv6_cidr of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2InsideIpv6Cidr() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel2_inside_ipv6_cidr"))
}

// Tunnel2Phase1DhGroupNumbers returns a reference to field tunnel2_phase1_dh_group_numbers of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2Phase1DhGroupNumbers() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](vc.ref.Append("tunnel2_phase1_dh_group_numbers"))
}

// Tunnel2Phase1EncryptionAlgorithms returns a reference to field tunnel2_phase1_encryption_algorithms of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2Phase1EncryptionAlgorithms() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("tunnel2_phase1_encryption_algorithms"))
}

// Tunnel2Phase1IntegrityAlgorithms returns a reference to field tunnel2_phase1_integrity_algorithms of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2Phase1IntegrityAlgorithms() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("tunnel2_phase1_integrity_algorithms"))
}

// Tunnel2Phase1LifetimeSeconds returns a reference to field tunnel2_phase1_lifetime_seconds of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2Phase1LifetimeSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel2_phase1_lifetime_seconds"))
}

// Tunnel2Phase2DhGroupNumbers returns a reference to field tunnel2_phase2_dh_group_numbers of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2Phase2DhGroupNumbers() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](vc.ref.Append("tunnel2_phase2_dh_group_numbers"))
}

// Tunnel2Phase2EncryptionAlgorithms returns a reference to field tunnel2_phase2_encryption_algorithms of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2Phase2EncryptionAlgorithms() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("tunnel2_phase2_encryption_algorithms"))
}

// Tunnel2Phase2IntegrityAlgorithms returns a reference to field tunnel2_phase2_integrity_algorithms of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2Phase2IntegrityAlgorithms() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("tunnel2_phase2_integrity_algorithms"))
}

// Tunnel2Phase2LifetimeSeconds returns a reference to field tunnel2_phase2_lifetime_seconds of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2Phase2LifetimeSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel2_phase2_lifetime_seconds"))
}

// Tunnel2PresharedKey returns a reference to field tunnel2_preshared_key of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2PresharedKey() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel2_preshared_key"))
}

// Tunnel2RekeyFuzzPercentage returns a reference to field tunnel2_rekey_fuzz_percentage of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2RekeyFuzzPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel2_rekey_fuzz_percentage"))
}

// Tunnel2RekeyMarginTimeSeconds returns a reference to field tunnel2_rekey_margin_time_seconds of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2RekeyMarginTimeSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel2_rekey_margin_time_seconds"))
}

// Tunnel2ReplayWindowSize returns a reference to field tunnel2_replay_window_size of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2ReplayWindowSize() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("tunnel2_replay_window_size"))
}

// Tunnel2StartupAction returns a reference to field tunnel2_startup_action of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2StartupAction() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel2_startup_action"))
}

// Tunnel2VgwInsideAddress returns a reference to field tunnel2_vgw_inside_address of aws_vpn_connection.
func (vc vpnConnectionAttributes) Tunnel2VgwInsideAddress() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel2_vgw_inside_address"))
}

// TunnelInsideIpVersion returns a reference to field tunnel_inside_ip_version of aws_vpn_connection.
func (vc vpnConnectionAttributes) TunnelInsideIpVersion() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("tunnel_inside_ip_version"))
}

// Type returns a reference to field type of aws_vpn_connection.
func (vc vpnConnectionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("type"))
}

// VpnGatewayId returns a reference to field vpn_gateway_id of aws_vpn_connection.
func (vc vpnConnectionAttributes) VpnGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("vpn_gateway_id"))
}

func (vc vpnConnectionAttributes) Routes() terra.SetValue[vpnconnection.RoutesAttributes] {
	return terra.ReferenceAsSet[vpnconnection.RoutesAttributes](vc.ref.Append("routes"))
}

func (vc vpnConnectionAttributes) VgwTelemetry() terra.SetValue[vpnconnection.VgwTelemetryAttributes] {
	return terra.ReferenceAsSet[vpnconnection.VgwTelemetryAttributes](vc.ref.Append("vgw_telemetry"))
}

func (vc vpnConnectionAttributes) Tunnel1LogOptions() terra.ListValue[vpnconnection.Tunnel1LogOptionsAttributes] {
	return terra.ReferenceAsList[vpnconnection.Tunnel1LogOptionsAttributes](vc.ref.Append("tunnel1_log_options"))
}

func (vc vpnConnectionAttributes) Tunnel2LogOptions() terra.ListValue[vpnconnection.Tunnel2LogOptionsAttributes] {
	return terra.ReferenceAsList[vpnconnection.Tunnel2LogOptionsAttributes](vc.ref.Append("tunnel2_log_options"))
}

type vpnConnectionState struct {
	Arn                                 string                                 `json:"arn"`
	CoreNetworkArn                      string                                 `json:"core_network_arn"`
	CoreNetworkAttachmentArn            string                                 `json:"core_network_attachment_arn"`
	CustomerGatewayConfiguration        string                                 `json:"customer_gateway_configuration"`
	CustomerGatewayId                   string                                 `json:"customer_gateway_id"`
	EnableAcceleration                  bool                                   `json:"enable_acceleration"`
	Id                                  string                                 `json:"id"`
	LocalIpv4NetworkCidr                string                                 `json:"local_ipv4_network_cidr"`
	LocalIpv6NetworkCidr                string                                 `json:"local_ipv6_network_cidr"`
	OutsideIpAddressType                string                                 `json:"outside_ip_address_type"`
	RemoteIpv4NetworkCidr               string                                 `json:"remote_ipv4_network_cidr"`
	RemoteIpv6NetworkCidr               string                                 `json:"remote_ipv6_network_cidr"`
	StaticRoutesOnly                    bool                                   `json:"static_routes_only"`
	Tags                                map[string]string                      `json:"tags"`
	TagsAll                             map[string]string                      `json:"tags_all"`
	TransitGatewayAttachmentId          string                                 `json:"transit_gateway_attachment_id"`
	TransitGatewayId                    string                                 `json:"transit_gateway_id"`
	TransportTransitGatewayAttachmentId string                                 `json:"transport_transit_gateway_attachment_id"`
	Tunnel1Address                      string                                 `json:"tunnel1_address"`
	Tunnel1BgpAsn                       string                                 `json:"tunnel1_bgp_asn"`
	Tunnel1BgpHoldtime                  float64                                `json:"tunnel1_bgp_holdtime"`
	Tunnel1CgwInsideAddress             string                                 `json:"tunnel1_cgw_inside_address"`
	Tunnel1DpdTimeoutAction             string                                 `json:"tunnel1_dpd_timeout_action"`
	Tunnel1DpdTimeoutSeconds            float64                                `json:"tunnel1_dpd_timeout_seconds"`
	Tunnel1IkeVersions                  []string                               `json:"tunnel1_ike_versions"`
	Tunnel1InsideCidr                   string                                 `json:"tunnel1_inside_cidr"`
	Tunnel1InsideIpv6Cidr               string                                 `json:"tunnel1_inside_ipv6_cidr"`
	Tunnel1Phase1DhGroupNumbers         []float64                              `json:"tunnel1_phase1_dh_group_numbers"`
	Tunnel1Phase1EncryptionAlgorithms   []string                               `json:"tunnel1_phase1_encryption_algorithms"`
	Tunnel1Phase1IntegrityAlgorithms    []string                               `json:"tunnel1_phase1_integrity_algorithms"`
	Tunnel1Phase1LifetimeSeconds        float64                                `json:"tunnel1_phase1_lifetime_seconds"`
	Tunnel1Phase2DhGroupNumbers         []float64                              `json:"tunnel1_phase2_dh_group_numbers"`
	Tunnel1Phase2EncryptionAlgorithms   []string                               `json:"tunnel1_phase2_encryption_algorithms"`
	Tunnel1Phase2IntegrityAlgorithms    []string                               `json:"tunnel1_phase2_integrity_algorithms"`
	Tunnel1Phase2LifetimeSeconds        float64                                `json:"tunnel1_phase2_lifetime_seconds"`
	Tunnel1PresharedKey                 string                                 `json:"tunnel1_preshared_key"`
	Tunnel1RekeyFuzzPercentage          float64                                `json:"tunnel1_rekey_fuzz_percentage"`
	Tunnel1RekeyMarginTimeSeconds       float64                                `json:"tunnel1_rekey_margin_time_seconds"`
	Tunnel1ReplayWindowSize             float64                                `json:"tunnel1_replay_window_size"`
	Tunnel1StartupAction                string                                 `json:"tunnel1_startup_action"`
	Tunnel1VgwInsideAddress             string                                 `json:"tunnel1_vgw_inside_address"`
	Tunnel2Address                      string                                 `json:"tunnel2_address"`
	Tunnel2BgpAsn                       string                                 `json:"tunnel2_bgp_asn"`
	Tunnel2BgpHoldtime                  float64                                `json:"tunnel2_bgp_holdtime"`
	Tunnel2CgwInsideAddress             string                                 `json:"tunnel2_cgw_inside_address"`
	Tunnel2DpdTimeoutAction             string                                 `json:"tunnel2_dpd_timeout_action"`
	Tunnel2DpdTimeoutSeconds            float64                                `json:"tunnel2_dpd_timeout_seconds"`
	Tunnel2IkeVersions                  []string                               `json:"tunnel2_ike_versions"`
	Tunnel2InsideCidr                   string                                 `json:"tunnel2_inside_cidr"`
	Tunnel2InsideIpv6Cidr               string                                 `json:"tunnel2_inside_ipv6_cidr"`
	Tunnel2Phase1DhGroupNumbers         []float64                              `json:"tunnel2_phase1_dh_group_numbers"`
	Tunnel2Phase1EncryptionAlgorithms   []string                               `json:"tunnel2_phase1_encryption_algorithms"`
	Tunnel2Phase1IntegrityAlgorithms    []string                               `json:"tunnel2_phase1_integrity_algorithms"`
	Tunnel2Phase1LifetimeSeconds        float64                                `json:"tunnel2_phase1_lifetime_seconds"`
	Tunnel2Phase2DhGroupNumbers         []float64                              `json:"tunnel2_phase2_dh_group_numbers"`
	Tunnel2Phase2EncryptionAlgorithms   []string                               `json:"tunnel2_phase2_encryption_algorithms"`
	Tunnel2Phase2IntegrityAlgorithms    []string                               `json:"tunnel2_phase2_integrity_algorithms"`
	Tunnel2Phase2LifetimeSeconds        float64                                `json:"tunnel2_phase2_lifetime_seconds"`
	Tunnel2PresharedKey                 string                                 `json:"tunnel2_preshared_key"`
	Tunnel2RekeyFuzzPercentage          float64                                `json:"tunnel2_rekey_fuzz_percentage"`
	Tunnel2RekeyMarginTimeSeconds       float64                                `json:"tunnel2_rekey_margin_time_seconds"`
	Tunnel2ReplayWindowSize             float64                                `json:"tunnel2_replay_window_size"`
	Tunnel2StartupAction                string                                 `json:"tunnel2_startup_action"`
	Tunnel2VgwInsideAddress             string                                 `json:"tunnel2_vgw_inside_address"`
	TunnelInsideIpVersion               string                                 `json:"tunnel_inside_ip_version"`
	Type                                string                                 `json:"type"`
	VpnGatewayId                        string                                 `json:"vpn_gateway_id"`
	Routes                              []vpnconnection.RoutesState            `json:"routes"`
	VgwTelemetry                        []vpnconnection.VgwTelemetryState      `json:"vgw_telemetry"`
	Tunnel1LogOptions                   []vpnconnection.Tunnel1LogOptionsState `json:"tunnel1_log_options"`
	Tunnel2LogOptions                   []vpnconnection.Tunnel2LogOptionsState `json:"tunnel2_log_options"`
}
