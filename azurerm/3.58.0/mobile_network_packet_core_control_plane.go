// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mobilenetworkpacketcorecontrolplane "github.com/golingon/terraproviders/azurerm/3.58.0/mobilenetworkpacketcorecontrolplane"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMobileNetworkPacketCoreControlPlane creates a new instance of [MobileNetworkPacketCoreControlPlane].
func NewMobileNetworkPacketCoreControlPlane(name string, args MobileNetworkPacketCoreControlPlaneArgs) *MobileNetworkPacketCoreControlPlane {
	return &MobileNetworkPacketCoreControlPlane{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MobileNetworkPacketCoreControlPlane)(nil)

// MobileNetworkPacketCoreControlPlane represents the Terraform resource azurerm_mobile_network_packet_core_control_plane.
type MobileNetworkPacketCoreControlPlane struct {
	Name      string
	Args      MobileNetworkPacketCoreControlPlaneArgs
	state     *mobileNetworkPacketCoreControlPlaneState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MobileNetworkPacketCoreControlPlane].
func (mnpccp *MobileNetworkPacketCoreControlPlane) Type() string {
	return "azurerm_mobile_network_packet_core_control_plane"
}

// LocalName returns the local name for [MobileNetworkPacketCoreControlPlane].
func (mnpccp *MobileNetworkPacketCoreControlPlane) LocalName() string {
	return mnpccp.Name
}

// Configuration returns the configuration (args) for [MobileNetworkPacketCoreControlPlane].
func (mnpccp *MobileNetworkPacketCoreControlPlane) Configuration() interface{} {
	return mnpccp.Args
}

// DependOn is used for other resources to depend on [MobileNetworkPacketCoreControlPlane].
func (mnpccp *MobileNetworkPacketCoreControlPlane) DependOn() terra.Reference {
	return terra.ReferenceResource(mnpccp)
}

// Dependencies returns the list of resources [MobileNetworkPacketCoreControlPlane] depends_on.
func (mnpccp *MobileNetworkPacketCoreControlPlane) Dependencies() terra.Dependencies {
	return mnpccp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MobileNetworkPacketCoreControlPlane].
func (mnpccp *MobileNetworkPacketCoreControlPlane) LifecycleManagement() *terra.Lifecycle {
	return mnpccp.Lifecycle
}

// Attributes returns the attributes for [MobileNetworkPacketCoreControlPlane].
func (mnpccp *MobileNetworkPacketCoreControlPlane) Attributes() mobileNetworkPacketCoreControlPlaneAttributes {
	return mobileNetworkPacketCoreControlPlaneAttributes{ref: terra.ReferenceResource(mnpccp)}
}

// ImportState imports the given attribute values into [MobileNetworkPacketCoreControlPlane]'s state.
func (mnpccp *MobileNetworkPacketCoreControlPlane) ImportState(av io.Reader) error {
	mnpccp.state = &mobileNetworkPacketCoreControlPlaneState{}
	if err := json.NewDecoder(av).Decode(mnpccp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mnpccp.Type(), mnpccp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MobileNetworkPacketCoreControlPlane] has state.
func (mnpccp *MobileNetworkPacketCoreControlPlane) State() (*mobileNetworkPacketCoreControlPlaneState, bool) {
	return mnpccp.state, mnpccp.state != nil
}

// StateMust returns the state for [MobileNetworkPacketCoreControlPlane]. Panics if the state is nil.
func (mnpccp *MobileNetworkPacketCoreControlPlane) StateMust() *mobileNetworkPacketCoreControlPlaneState {
	if mnpccp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mnpccp.Type(), mnpccp.LocalName()))
	}
	return mnpccp.state
}

// MobileNetworkPacketCoreControlPlaneArgs contains the configurations for azurerm_mobile_network_packet_core_control_plane.
type MobileNetworkPacketCoreControlPlaneArgs struct {
	// ControlPlaneAccessIpv4Address: string, optional
	ControlPlaneAccessIpv4Address terra.StringValue `hcl:"control_plane_access_ipv4_address,attr"`
	// ControlPlaneAccessIpv4Gateway: string, optional
	ControlPlaneAccessIpv4Gateway terra.StringValue `hcl:"control_plane_access_ipv4_gateway,attr"`
	// ControlPlaneAccessIpv4Subnet: string, optional
	ControlPlaneAccessIpv4Subnet terra.StringValue `hcl:"control_plane_access_ipv4_subnet,attr"`
	// ControlPlaneAccessName: string, optional
	ControlPlaneAccessName terra.StringValue `hcl:"control_plane_access_name,attr"`
	// CoreNetworkTechnology: string, optional
	CoreNetworkTechnology terra.StringValue `hcl:"core_network_technology,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InteroperabilitySettingsJson: string, optional
	InteroperabilitySettingsJson terra.StringValue `hcl:"interoperability_settings_json,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SiteIds: list of string, required
	SiteIds terra.ListValue[terra.StringValue] `hcl:"site_ids,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// SoftwareVersion: string, optional
	SoftwareVersion terra.StringValue `hcl:"software_version,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// UserEquipmentMtuInBytes: number, optional
	UserEquipmentMtuInBytes terra.NumberValue `hcl:"user_equipment_mtu_in_bytes,attr"`
	// Identity: optional
	Identity *mobilenetworkpacketcorecontrolplane.Identity `hcl:"identity,block"`
	// LocalDiagnosticsAccess: required
	LocalDiagnosticsAccess *mobilenetworkpacketcorecontrolplane.LocalDiagnosticsAccess `hcl:"local_diagnostics_access,block" validate:"required"`
	// Platform: optional
	Platform *mobilenetworkpacketcorecontrolplane.Platform `hcl:"platform,block"`
	// Timeouts: optional
	Timeouts *mobilenetworkpacketcorecontrolplane.Timeouts `hcl:"timeouts,block"`
}
type mobileNetworkPacketCoreControlPlaneAttributes struct {
	ref terra.Reference
}

// ControlPlaneAccessIpv4Address returns a reference to field control_plane_access_ipv4_address of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) ControlPlaneAccessIpv4Address() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("control_plane_access_ipv4_address"))
}

// ControlPlaneAccessIpv4Gateway returns a reference to field control_plane_access_ipv4_gateway of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) ControlPlaneAccessIpv4Gateway() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("control_plane_access_ipv4_gateway"))
}

// ControlPlaneAccessIpv4Subnet returns a reference to field control_plane_access_ipv4_subnet of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) ControlPlaneAccessIpv4Subnet() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("control_plane_access_ipv4_subnet"))
}

// ControlPlaneAccessName returns a reference to field control_plane_access_name of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) ControlPlaneAccessName() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("control_plane_access_name"))
}

// CoreNetworkTechnology returns a reference to field core_network_technology of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) CoreNetworkTechnology() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("core_network_technology"))
}

// Id returns a reference to field id of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("id"))
}

// InteroperabilitySettingsJson returns a reference to field interoperability_settings_json of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) InteroperabilitySettingsJson() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("interoperability_settings_json"))
}

// Location returns a reference to field location of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("resource_group_name"))
}

// SiteIds returns a reference to field site_ids of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) SiteIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mnpccp.ref.Append("site_ids"))
}

// Sku returns a reference to field sku of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("sku"))
}

// SoftwareVersion returns a reference to field software_version of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) SoftwareVersion() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("software_version"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mnpccp.ref.Append("tags"))
}

// UserEquipmentMtuInBytes returns a reference to field user_equipment_mtu_in_bytes of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) UserEquipmentMtuInBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(mnpccp.ref.Append("user_equipment_mtu_in_bytes"))
}

func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) Identity() terra.ListValue[mobilenetworkpacketcorecontrolplane.IdentityAttributes] {
	return terra.ReferenceAsList[mobilenetworkpacketcorecontrolplane.IdentityAttributes](mnpccp.ref.Append("identity"))
}

func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) LocalDiagnosticsAccess() terra.ListValue[mobilenetworkpacketcorecontrolplane.LocalDiagnosticsAccessAttributes] {
	return terra.ReferenceAsList[mobilenetworkpacketcorecontrolplane.LocalDiagnosticsAccessAttributes](mnpccp.ref.Append("local_diagnostics_access"))
}

func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) Platform() terra.ListValue[mobilenetworkpacketcorecontrolplane.PlatformAttributes] {
	return terra.ReferenceAsList[mobilenetworkpacketcorecontrolplane.PlatformAttributes](mnpccp.ref.Append("platform"))
}

func (mnpccp mobileNetworkPacketCoreControlPlaneAttributes) Timeouts() mobilenetworkpacketcorecontrolplane.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mobilenetworkpacketcorecontrolplane.TimeoutsAttributes](mnpccp.ref.Append("timeouts"))
}

type mobileNetworkPacketCoreControlPlaneState struct {
	ControlPlaneAccessIpv4Address string                                                            `json:"control_plane_access_ipv4_address"`
	ControlPlaneAccessIpv4Gateway string                                                            `json:"control_plane_access_ipv4_gateway"`
	ControlPlaneAccessIpv4Subnet  string                                                            `json:"control_plane_access_ipv4_subnet"`
	ControlPlaneAccessName        string                                                            `json:"control_plane_access_name"`
	CoreNetworkTechnology         string                                                            `json:"core_network_technology"`
	Id                            string                                                            `json:"id"`
	InteroperabilitySettingsJson  string                                                            `json:"interoperability_settings_json"`
	Location                      string                                                            `json:"location"`
	Name                          string                                                            `json:"name"`
	ResourceGroupName             string                                                            `json:"resource_group_name"`
	SiteIds                       []string                                                          `json:"site_ids"`
	Sku                           string                                                            `json:"sku"`
	SoftwareVersion               string                                                            `json:"software_version"`
	Tags                          map[string]string                                                 `json:"tags"`
	UserEquipmentMtuInBytes       float64                                                           `json:"user_equipment_mtu_in_bytes"`
	Identity                      []mobilenetworkpacketcorecontrolplane.IdentityState               `json:"identity"`
	LocalDiagnosticsAccess        []mobilenetworkpacketcorecontrolplane.LocalDiagnosticsAccessState `json:"local_diagnostics_access"`
	Platform                      []mobilenetworkpacketcorecontrolplane.PlatformState               `json:"platform"`
	Timeouts                      *mobilenetworkpacketcorecontrolplane.TimeoutsState                `json:"timeouts"`
}