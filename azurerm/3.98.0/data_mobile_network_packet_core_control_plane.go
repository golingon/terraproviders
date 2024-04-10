// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datamobilenetworkpacketcorecontrolplane "github.com/golingon/terraproviders/azurerm/3.98.0/datamobilenetworkpacketcorecontrolplane"
)

// NewDataMobileNetworkPacketCoreControlPlane creates a new instance of [DataMobileNetworkPacketCoreControlPlane].
func NewDataMobileNetworkPacketCoreControlPlane(name string, args DataMobileNetworkPacketCoreControlPlaneArgs) *DataMobileNetworkPacketCoreControlPlane {
	return &DataMobileNetworkPacketCoreControlPlane{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMobileNetworkPacketCoreControlPlane)(nil)

// DataMobileNetworkPacketCoreControlPlane represents the Terraform data resource azurerm_mobile_network_packet_core_control_plane.
type DataMobileNetworkPacketCoreControlPlane struct {
	Name string
	Args DataMobileNetworkPacketCoreControlPlaneArgs
}

// DataSource returns the Terraform object type for [DataMobileNetworkPacketCoreControlPlane].
func (mnpccp *DataMobileNetworkPacketCoreControlPlane) DataSource() string {
	return "azurerm_mobile_network_packet_core_control_plane"
}

// LocalName returns the local name for [DataMobileNetworkPacketCoreControlPlane].
func (mnpccp *DataMobileNetworkPacketCoreControlPlane) LocalName() string {
	return mnpccp.Name
}

// Configuration returns the configuration (args) for [DataMobileNetworkPacketCoreControlPlane].
func (mnpccp *DataMobileNetworkPacketCoreControlPlane) Configuration() interface{} {
	return mnpccp.Args
}

// Attributes returns the attributes for [DataMobileNetworkPacketCoreControlPlane].
func (mnpccp *DataMobileNetworkPacketCoreControlPlane) Attributes() dataMobileNetworkPacketCoreControlPlaneAttributes {
	return dataMobileNetworkPacketCoreControlPlaneAttributes{ref: terra.ReferenceDataResource(mnpccp)}
}

// DataMobileNetworkPacketCoreControlPlaneArgs contains the configurations for azurerm_mobile_network_packet_core_control_plane.
type DataMobileNetworkPacketCoreControlPlaneArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datamobilenetworkpacketcorecontrolplane.Identity `hcl:"identity,block" validate:"min=0"`
	// LocalDiagnosticsAccess: min=0
	LocalDiagnosticsAccess []datamobilenetworkpacketcorecontrolplane.LocalDiagnosticsAccess `hcl:"local_diagnostics_access,block" validate:"min=0"`
	// Platform: min=0
	Platform []datamobilenetworkpacketcorecontrolplane.Platform `hcl:"platform,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamobilenetworkpacketcorecontrolplane.Timeouts `hcl:"timeouts,block"`
}
type dataMobileNetworkPacketCoreControlPlaneAttributes struct {
	ref terra.Reference
}

// ControlPlaneAccessIpv4Address returns a reference to field control_plane_access_ipv4_address of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) ControlPlaneAccessIpv4Address() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("control_plane_access_ipv4_address"))
}

// ControlPlaneAccessIpv4Gateway returns a reference to field control_plane_access_ipv4_gateway of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) ControlPlaneAccessIpv4Gateway() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("control_plane_access_ipv4_gateway"))
}

// ControlPlaneAccessIpv4Subnet returns a reference to field control_plane_access_ipv4_subnet of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) ControlPlaneAccessIpv4Subnet() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("control_plane_access_ipv4_subnet"))
}

// ControlPlaneAccessName returns a reference to field control_plane_access_name of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) ControlPlaneAccessName() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("control_plane_access_name"))
}

// CoreNetworkTechnology returns a reference to field core_network_technology of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) CoreNetworkTechnology() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("core_network_technology"))
}

// Id returns a reference to field id of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("id"))
}

// InteroperabilitySettingsJson returns a reference to field interoperability_settings_json of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) InteroperabilitySettingsJson() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("interoperability_settings_json"))
}

// Location returns a reference to field location of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("resource_group_name"))
}

// SiteIds returns a reference to field site_ids of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) SiteIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mnpccp.ref.Append("site_ids"))
}

// Sku returns a reference to field sku of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("sku"))
}

// SoftwareVersion returns a reference to field software_version of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) SoftwareVersion() terra.StringValue {
	return terra.ReferenceAsString(mnpccp.ref.Append("software_version"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mnpccp.ref.Append("tags"))
}

// UserEquipmentMtuInBytes returns a reference to field user_equipment_mtu_in_bytes of azurerm_mobile_network_packet_core_control_plane.
func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) UserEquipmentMtuInBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(mnpccp.ref.Append("user_equipment_mtu_in_bytes"))
}

func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) Identity() terra.ListValue[datamobilenetworkpacketcorecontrolplane.IdentityAttributes] {
	return terra.ReferenceAsList[datamobilenetworkpacketcorecontrolplane.IdentityAttributes](mnpccp.ref.Append("identity"))
}

func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) LocalDiagnosticsAccess() terra.ListValue[datamobilenetworkpacketcorecontrolplane.LocalDiagnosticsAccessAttributes] {
	return terra.ReferenceAsList[datamobilenetworkpacketcorecontrolplane.LocalDiagnosticsAccessAttributes](mnpccp.ref.Append("local_diagnostics_access"))
}

func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) Platform() terra.ListValue[datamobilenetworkpacketcorecontrolplane.PlatformAttributes] {
	return terra.ReferenceAsList[datamobilenetworkpacketcorecontrolplane.PlatformAttributes](mnpccp.ref.Append("platform"))
}

func (mnpccp dataMobileNetworkPacketCoreControlPlaneAttributes) Timeouts() datamobilenetworkpacketcorecontrolplane.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamobilenetworkpacketcorecontrolplane.TimeoutsAttributes](mnpccp.ref.Append("timeouts"))
}
