// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_app_service_environment_v3

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_app_service_environment_v3.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aasev *DataSource) DataSource() string {
	return "azurerm_app_service_environment_v3"
}

// LocalName returns the local name for [DataSource].
func (aasev *DataSource) LocalName() string {
	return aasev.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aasev *DataSource) Configuration() interface{} {
	return aasev.Args
}

// Attributes returns the attributes for [DataSource].
func (aasev *DataSource) Attributes() dataAzurermAppServiceEnvironmentV3Attributes {
	return dataAzurermAppServiceEnvironmentV3Attributes{ref: terra.ReferenceDataSource(aasev)}
}

// DataArgs contains the configurations for azurerm_app_service_environment_v3.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermAppServiceEnvironmentV3Attributes struct {
	ref terra.Reference
}

// AllowNewPrivateEndpointConnections returns a reference to field allow_new_private_endpoint_connections of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) AllowNewPrivateEndpointConnections() terra.BoolValue {
	return terra.ReferenceAsBool(aasev.ref.Append("allow_new_private_endpoint_connections"))
}

// DedicatedHostCount returns a reference to field dedicated_host_count of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) DedicatedHostCount() terra.NumberValue {
	return terra.ReferenceAsNumber(aasev.ref.Append("dedicated_host_count"))
}

// DnsSuffix returns a reference to field dns_suffix of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) DnsSuffix() terra.StringValue {
	return terra.ReferenceAsString(aasev.ref.Append("dns_suffix"))
}

// ExternalInboundIpAddresses returns a reference to field external_inbound_ip_addresses of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) ExternalInboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aasev.ref.Append("external_inbound_ip_addresses"))
}

// Id returns a reference to field id of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aasev.ref.Append("id"))
}

// InternalInboundIpAddresses returns a reference to field internal_inbound_ip_addresses of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) InternalInboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aasev.ref.Append("internal_inbound_ip_addresses"))
}

// InternalLoadBalancingMode returns a reference to field internal_load_balancing_mode of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) InternalLoadBalancingMode() terra.StringValue {
	return terra.ReferenceAsString(aasev.ref.Append("internal_load_balancing_mode"))
}

// IpSslAddressCount returns a reference to field ip_ssl_address_count of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) IpSslAddressCount() terra.NumberValue {
	return terra.ReferenceAsNumber(aasev.ref.Append("ip_ssl_address_count"))
}

// LinuxOutboundIpAddresses returns a reference to field linux_outbound_ip_addresses of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) LinuxOutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aasev.ref.Append("linux_outbound_ip_addresses"))
}

// Location returns a reference to field location of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aasev.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aasev.ref.Append("name"))
}

// PricingTier returns a reference to field pricing_tier of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) PricingTier() terra.StringValue {
	return terra.ReferenceAsString(aasev.ref.Append("pricing_tier"))
}

// RemoteDebuggingEnabled returns a reference to field remote_debugging_enabled of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) RemoteDebuggingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aasev.ref.Append("remote_debugging_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aasev.ref.Append("resource_group_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(aasev.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aasev.ref.Append("tags"))
}

// WindowsOutboundIpAddresses returns a reference to field windows_outbound_ip_addresses of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) WindowsOutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aasev.ref.Append("windows_outbound_ip_addresses"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_app_service_environment_v3.
func (aasev dataAzurermAppServiceEnvironmentV3Attributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(aasev.ref.Append("zone_redundant"))
}

func (aasev dataAzurermAppServiceEnvironmentV3Attributes) ClusterSetting() terra.ListValue[DataClusterSettingAttributes] {
	return terra.ReferenceAsList[DataClusterSettingAttributes](aasev.ref.Append("cluster_setting"))
}

func (aasev dataAzurermAppServiceEnvironmentV3Attributes) InboundNetworkDependencies() terra.ListValue[DataInboundNetworkDependenciesAttributes] {
	return terra.ReferenceAsList[DataInboundNetworkDependenciesAttributes](aasev.ref.Append("inbound_network_dependencies"))
}

func (aasev dataAzurermAppServiceEnvironmentV3Attributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](aasev.ref.Append("timeouts"))
}
