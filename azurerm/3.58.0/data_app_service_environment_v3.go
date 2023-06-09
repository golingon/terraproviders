// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataappserviceenvironmentv3 "github.com/golingon/terraproviders/azurerm/3.58.0/dataappserviceenvironmentv3"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppServiceEnvironmentV3 creates a new instance of [DataAppServiceEnvironmentV3].
func NewDataAppServiceEnvironmentV3(name string, args DataAppServiceEnvironmentV3Args) *DataAppServiceEnvironmentV3 {
	return &DataAppServiceEnvironmentV3{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppServiceEnvironmentV3)(nil)

// DataAppServiceEnvironmentV3 represents the Terraform data resource azurerm_app_service_environment_v3.
type DataAppServiceEnvironmentV3 struct {
	Name string
	Args DataAppServiceEnvironmentV3Args
}

// DataSource returns the Terraform object type for [DataAppServiceEnvironmentV3].
func (asev *DataAppServiceEnvironmentV3) DataSource() string {
	return "azurerm_app_service_environment_v3"
}

// LocalName returns the local name for [DataAppServiceEnvironmentV3].
func (asev *DataAppServiceEnvironmentV3) LocalName() string {
	return asev.Name
}

// Configuration returns the configuration (args) for [DataAppServiceEnvironmentV3].
func (asev *DataAppServiceEnvironmentV3) Configuration() interface{} {
	return asev.Args
}

// Attributes returns the attributes for [DataAppServiceEnvironmentV3].
func (asev *DataAppServiceEnvironmentV3) Attributes() dataAppServiceEnvironmentV3Attributes {
	return dataAppServiceEnvironmentV3Attributes{ref: terra.ReferenceDataResource(asev)}
}

// DataAppServiceEnvironmentV3Args contains the configurations for azurerm_app_service_environment_v3.
type DataAppServiceEnvironmentV3Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ClusterSetting: min=0
	ClusterSetting []dataappserviceenvironmentv3.ClusterSetting `hcl:"cluster_setting,block" validate:"min=0"`
	// InboundNetworkDependencies: min=0
	InboundNetworkDependencies []dataappserviceenvironmentv3.InboundNetworkDependencies `hcl:"inbound_network_dependencies,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataappserviceenvironmentv3.Timeouts `hcl:"timeouts,block"`
}
type dataAppServiceEnvironmentV3Attributes struct {
	ref terra.Reference
}

// AllowNewPrivateEndpointConnections returns a reference to field allow_new_private_endpoint_connections of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) AllowNewPrivateEndpointConnections() terra.BoolValue {
	return terra.ReferenceAsBool(asev.ref.Append("allow_new_private_endpoint_connections"))
}

// DedicatedHostCount returns a reference to field dedicated_host_count of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) DedicatedHostCount() terra.NumberValue {
	return terra.ReferenceAsNumber(asev.ref.Append("dedicated_host_count"))
}

// DnsSuffix returns a reference to field dns_suffix of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) DnsSuffix() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("dns_suffix"))
}

// ExternalInboundIpAddresses returns a reference to field external_inbound_ip_addresses of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) ExternalInboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asev.ref.Append("external_inbound_ip_addresses"))
}

// Id returns a reference to field id of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("id"))
}

// InternalInboundIpAddresses returns a reference to field internal_inbound_ip_addresses of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) InternalInboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asev.ref.Append("internal_inbound_ip_addresses"))
}

// InternalLoadBalancingMode returns a reference to field internal_load_balancing_mode of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) InternalLoadBalancingMode() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("internal_load_balancing_mode"))
}

// IpSslAddressCount returns a reference to field ip_ssl_address_count of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) IpSslAddressCount() terra.NumberValue {
	return terra.ReferenceAsNumber(asev.ref.Append("ip_ssl_address_count"))
}

// LinuxOutboundIpAddresses returns a reference to field linux_outbound_ip_addresses of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) LinuxOutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asev.ref.Append("linux_outbound_ip_addresses"))
}

// Location returns a reference to field location of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) Location() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("name"))
}

// PricingTier returns a reference to field pricing_tier of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) PricingTier() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("pricing_tier"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("resource_group_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asev.ref.Append("tags"))
}

// WindowsOutboundIpAddresses returns a reference to field windows_outbound_ip_addresses of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) WindowsOutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asev.ref.Append("windows_outbound_ip_addresses"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_app_service_environment_v3.
func (asev dataAppServiceEnvironmentV3Attributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(asev.ref.Append("zone_redundant"))
}

func (asev dataAppServiceEnvironmentV3Attributes) ClusterSetting() terra.ListValue[dataappserviceenvironmentv3.ClusterSettingAttributes] {
	return terra.ReferenceAsList[dataappserviceenvironmentv3.ClusterSettingAttributes](asev.ref.Append("cluster_setting"))
}

func (asev dataAppServiceEnvironmentV3Attributes) InboundNetworkDependencies() terra.ListValue[dataappserviceenvironmentv3.InboundNetworkDependenciesAttributes] {
	return terra.ReferenceAsList[dataappserviceenvironmentv3.InboundNetworkDependenciesAttributes](asev.ref.Append("inbound_network_dependencies"))
}

func (asev dataAppServiceEnvironmentV3Attributes) Timeouts() dataappserviceenvironmentv3.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataappserviceenvironmentv3.TimeoutsAttributes](asev.ref.Append("timeouts"))
}
