// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appserviceenvironmentv3 "github.com/golingon/terraproviders/azurerm/3.52.0/appserviceenvironmentv3"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceEnvironmentV3 creates a new instance of [AppServiceEnvironmentV3].
func NewAppServiceEnvironmentV3(name string, args AppServiceEnvironmentV3Args) *AppServiceEnvironmentV3 {
	return &AppServiceEnvironmentV3{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceEnvironmentV3)(nil)

// AppServiceEnvironmentV3 represents the Terraform resource azurerm_app_service_environment_v3.
type AppServiceEnvironmentV3 struct {
	Name      string
	Args      AppServiceEnvironmentV3Args
	state     *appServiceEnvironmentV3State
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceEnvironmentV3].
func (asev *AppServiceEnvironmentV3) Type() string {
	return "azurerm_app_service_environment_v3"
}

// LocalName returns the local name for [AppServiceEnvironmentV3].
func (asev *AppServiceEnvironmentV3) LocalName() string {
	return asev.Name
}

// Configuration returns the configuration (args) for [AppServiceEnvironmentV3].
func (asev *AppServiceEnvironmentV3) Configuration() interface{} {
	return asev.Args
}

// DependOn is used for other resources to depend on [AppServiceEnvironmentV3].
func (asev *AppServiceEnvironmentV3) DependOn() terra.Reference {
	return terra.ReferenceResource(asev)
}

// Dependencies returns the list of resources [AppServiceEnvironmentV3] depends_on.
func (asev *AppServiceEnvironmentV3) Dependencies() terra.Dependencies {
	return asev.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceEnvironmentV3].
func (asev *AppServiceEnvironmentV3) LifecycleManagement() *terra.Lifecycle {
	return asev.Lifecycle
}

// Attributes returns the attributes for [AppServiceEnvironmentV3].
func (asev *AppServiceEnvironmentV3) Attributes() appServiceEnvironmentV3Attributes {
	return appServiceEnvironmentV3Attributes{ref: terra.ReferenceResource(asev)}
}

// ImportState imports the given attribute values into [AppServiceEnvironmentV3]'s state.
func (asev *AppServiceEnvironmentV3) ImportState(av io.Reader) error {
	asev.state = &appServiceEnvironmentV3State{}
	if err := json.NewDecoder(av).Decode(asev.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asev.Type(), asev.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceEnvironmentV3] has state.
func (asev *AppServiceEnvironmentV3) State() (*appServiceEnvironmentV3State, bool) {
	return asev.state, asev.state != nil
}

// StateMust returns the state for [AppServiceEnvironmentV3]. Panics if the state is nil.
func (asev *AppServiceEnvironmentV3) StateMust() *appServiceEnvironmentV3State {
	if asev.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asev.Type(), asev.LocalName()))
	}
	return asev.state
}

// AppServiceEnvironmentV3Args contains the configurations for azurerm_app_service_environment_v3.
type AppServiceEnvironmentV3Args struct {
	// AllowNewPrivateEndpointConnections: bool, optional
	AllowNewPrivateEndpointConnections terra.BoolValue `hcl:"allow_new_private_endpoint_connections,attr"`
	// DedicatedHostCount: number, optional
	DedicatedHostCount terra.NumberValue `hcl:"dedicated_host_count,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InternalLoadBalancingMode: string, optional
	InternalLoadBalancingMode terra.StringValue `hcl:"internal_load_balancing_mode,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ZoneRedundant: bool, optional
	ZoneRedundant terra.BoolValue `hcl:"zone_redundant,attr"`
	// InboundNetworkDependencies: min=0
	InboundNetworkDependencies []appserviceenvironmentv3.InboundNetworkDependencies `hcl:"inbound_network_dependencies,block" validate:"min=0"`
	// ClusterSetting: min=0
	ClusterSetting []appserviceenvironmentv3.ClusterSetting `hcl:"cluster_setting,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *appserviceenvironmentv3.Timeouts `hcl:"timeouts,block"`
}
type appServiceEnvironmentV3Attributes struct {
	ref terra.Reference
}

// AllowNewPrivateEndpointConnections returns a reference to field allow_new_private_endpoint_connections of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) AllowNewPrivateEndpointConnections() terra.BoolValue {
	return terra.ReferenceAsBool(asev.ref.Append("allow_new_private_endpoint_connections"))
}

// DedicatedHostCount returns a reference to field dedicated_host_count of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) DedicatedHostCount() terra.NumberValue {
	return terra.ReferenceAsNumber(asev.ref.Append("dedicated_host_count"))
}

// DnsSuffix returns a reference to field dns_suffix of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) DnsSuffix() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("dns_suffix"))
}

// ExternalInboundIpAddresses returns a reference to field external_inbound_ip_addresses of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) ExternalInboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asev.ref.Append("external_inbound_ip_addresses"))
}

// Id returns a reference to field id of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("id"))
}

// InternalInboundIpAddresses returns a reference to field internal_inbound_ip_addresses of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) InternalInboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asev.ref.Append("internal_inbound_ip_addresses"))
}

// InternalLoadBalancingMode returns a reference to field internal_load_balancing_mode of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) InternalLoadBalancingMode() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("internal_load_balancing_mode"))
}

// IpSslAddressCount returns a reference to field ip_ssl_address_count of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) IpSslAddressCount() terra.NumberValue {
	return terra.ReferenceAsNumber(asev.ref.Append("ip_ssl_address_count"))
}

// LinuxOutboundIpAddresses returns a reference to field linux_outbound_ip_addresses of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) LinuxOutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asev.ref.Append("linux_outbound_ip_addresses"))
}

// Location returns a reference to field location of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) Location() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("name"))
}

// PricingTier returns a reference to field pricing_tier of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) PricingTier() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("pricing_tier"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("resource_group_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(asev.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asev.ref.Append("tags"))
}

// WindowsOutboundIpAddresses returns a reference to field windows_outbound_ip_addresses of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) WindowsOutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asev.ref.Append("windows_outbound_ip_addresses"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_app_service_environment_v3.
func (asev appServiceEnvironmentV3Attributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(asev.ref.Append("zone_redundant"))
}

func (asev appServiceEnvironmentV3Attributes) InboundNetworkDependencies() terra.ListValue[appserviceenvironmentv3.InboundNetworkDependenciesAttributes] {
	return terra.ReferenceAsList[appserviceenvironmentv3.InboundNetworkDependenciesAttributes](asev.ref.Append("inbound_network_dependencies"))
}

func (asev appServiceEnvironmentV3Attributes) ClusterSetting() terra.ListValue[appserviceenvironmentv3.ClusterSettingAttributes] {
	return terra.ReferenceAsList[appserviceenvironmentv3.ClusterSettingAttributes](asev.ref.Append("cluster_setting"))
}

func (asev appServiceEnvironmentV3Attributes) Timeouts() appserviceenvironmentv3.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appserviceenvironmentv3.TimeoutsAttributes](asev.ref.Append("timeouts"))
}

type appServiceEnvironmentV3State struct {
	AllowNewPrivateEndpointConnections bool                                                      `json:"allow_new_private_endpoint_connections"`
	DedicatedHostCount                 float64                                                   `json:"dedicated_host_count"`
	DnsSuffix                          string                                                    `json:"dns_suffix"`
	ExternalInboundIpAddresses         []string                                                  `json:"external_inbound_ip_addresses"`
	Id                                 string                                                    `json:"id"`
	InternalInboundIpAddresses         []string                                                  `json:"internal_inbound_ip_addresses"`
	InternalLoadBalancingMode          string                                                    `json:"internal_load_balancing_mode"`
	IpSslAddressCount                  float64                                                   `json:"ip_ssl_address_count"`
	LinuxOutboundIpAddresses           []string                                                  `json:"linux_outbound_ip_addresses"`
	Location                           string                                                    `json:"location"`
	Name                               string                                                    `json:"name"`
	PricingTier                        string                                                    `json:"pricing_tier"`
	ResourceGroupName                  string                                                    `json:"resource_group_name"`
	SubnetId                           string                                                    `json:"subnet_id"`
	Tags                               map[string]string                                         `json:"tags"`
	WindowsOutboundIpAddresses         []string                                                  `json:"windows_outbound_ip_addresses"`
	ZoneRedundant                      bool                                                      `json:"zone_redundant"`
	InboundNetworkDependencies         []appserviceenvironmentv3.InboundNetworkDependenciesState `json:"inbound_network_dependencies"`
	ClusterSetting                     []appserviceenvironmentv3.ClusterSettingState             `json:"cluster_setting"`
	Timeouts                           *appserviceenvironmentv3.TimeoutsState                    `json:"timeouts"`
}
