// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	pointtositevpngateway "github.com/golingon/terraproviders/azurerm/3.69.0/pointtositevpngateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPointToSiteVpnGateway creates a new instance of [PointToSiteVpnGateway].
func NewPointToSiteVpnGateway(name string, args PointToSiteVpnGatewayArgs) *PointToSiteVpnGateway {
	return &PointToSiteVpnGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PointToSiteVpnGateway)(nil)

// PointToSiteVpnGateway represents the Terraform resource azurerm_point_to_site_vpn_gateway.
type PointToSiteVpnGateway struct {
	Name      string
	Args      PointToSiteVpnGatewayArgs
	state     *pointToSiteVpnGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PointToSiteVpnGateway].
func (ptsvg *PointToSiteVpnGateway) Type() string {
	return "azurerm_point_to_site_vpn_gateway"
}

// LocalName returns the local name for [PointToSiteVpnGateway].
func (ptsvg *PointToSiteVpnGateway) LocalName() string {
	return ptsvg.Name
}

// Configuration returns the configuration (args) for [PointToSiteVpnGateway].
func (ptsvg *PointToSiteVpnGateway) Configuration() interface{} {
	return ptsvg.Args
}

// DependOn is used for other resources to depend on [PointToSiteVpnGateway].
func (ptsvg *PointToSiteVpnGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(ptsvg)
}

// Dependencies returns the list of resources [PointToSiteVpnGateway] depends_on.
func (ptsvg *PointToSiteVpnGateway) Dependencies() terra.Dependencies {
	return ptsvg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PointToSiteVpnGateway].
func (ptsvg *PointToSiteVpnGateway) LifecycleManagement() *terra.Lifecycle {
	return ptsvg.Lifecycle
}

// Attributes returns the attributes for [PointToSiteVpnGateway].
func (ptsvg *PointToSiteVpnGateway) Attributes() pointToSiteVpnGatewayAttributes {
	return pointToSiteVpnGatewayAttributes{ref: terra.ReferenceResource(ptsvg)}
}

// ImportState imports the given attribute values into [PointToSiteVpnGateway]'s state.
func (ptsvg *PointToSiteVpnGateway) ImportState(av io.Reader) error {
	ptsvg.state = &pointToSiteVpnGatewayState{}
	if err := json.NewDecoder(av).Decode(ptsvg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ptsvg.Type(), ptsvg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PointToSiteVpnGateway] has state.
func (ptsvg *PointToSiteVpnGateway) State() (*pointToSiteVpnGatewayState, bool) {
	return ptsvg.state, ptsvg.state != nil
}

// StateMust returns the state for [PointToSiteVpnGateway]. Panics if the state is nil.
func (ptsvg *PointToSiteVpnGateway) StateMust() *pointToSiteVpnGatewayState {
	if ptsvg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ptsvg.Type(), ptsvg.LocalName()))
	}
	return ptsvg.state
}

// PointToSiteVpnGatewayArgs contains the configurations for azurerm_point_to_site_vpn_gateway.
type PointToSiteVpnGatewayArgs struct {
	// DnsServers: list of string, optional
	DnsServers terra.ListValue[terra.StringValue] `hcl:"dns_servers,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RoutingPreferenceInternetEnabled: bool, optional
	RoutingPreferenceInternetEnabled terra.BoolValue `hcl:"routing_preference_internet_enabled,attr"`
	// ScaleUnit: number, required
	ScaleUnit terra.NumberValue `hcl:"scale_unit,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualHubId: string, required
	VirtualHubId terra.StringValue `hcl:"virtual_hub_id,attr" validate:"required"`
	// VpnServerConfigurationId: string, required
	VpnServerConfigurationId terra.StringValue `hcl:"vpn_server_configuration_id,attr" validate:"required"`
	// ConnectionConfiguration: required
	ConnectionConfiguration *pointtositevpngateway.ConnectionConfiguration `hcl:"connection_configuration,block" validate:"required"`
	// Timeouts: optional
	Timeouts *pointtositevpngateway.Timeouts `hcl:"timeouts,block"`
}
type pointToSiteVpnGatewayAttributes struct {
	ref terra.Reference
}

// DnsServers returns a reference to field dns_servers of azurerm_point_to_site_vpn_gateway.
func (ptsvg pointToSiteVpnGatewayAttributes) DnsServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ptsvg.ref.Append("dns_servers"))
}

// Id returns a reference to field id of azurerm_point_to_site_vpn_gateway.
func (ptsvg pointToSiteVpnGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ptsvg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_point_to_site_vpn_gateway.
func (ptsvg pointToSiteVpnGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ptsvg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_point_to_site_vpn_gateway.
func (ptsvg pointToSiteVpnGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ptsvg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_point_to_site_vpn_gateway.
func (ptsvg pointToSiteVpnGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ptsvg.ref.Append("resource_group_name"))
}

// RoutingPreferenceInternetEnabled returns a reference to field routing_preference_internet_enabled of azurerm_point_to_site_vpn_gateway.
func (ptsvg pointToSiteVpnGatewayAttributes) RoutingPreferenceInternetEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ptsvg.ref.Append("routing_preference_internet_enabled"))
}

// ScaleUnit returns a reference to field scale_unit of azurerm_point_to_site_vpn_gateway.
func (ptsvg pointToSiteVpnGatewayAttributes) ScaleUnit() terra.NumberValue {
	return terra.ReferenceAsNumber(ptsvg.ref.Append("scale_unit"))
}

// Tags returns a reference to field tags of azurerm_point_to_site_vpn_gateway.
func (ptsvg pointToSiteVpnGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ptsvg.ref.Append("tags"))
}

// VirtualHubId returns a reference to field virtual_hub_id of azurerm_point_to_site_vpn_gateway.
func (ptsvg pointToSiteVpnGatewayAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceAsString(ptsvg.ref.Append("virtual_hub_id"))
}

// VpnServerConfigurationId returns a reference to field vpn_server_configuration_id of azurerm_point_to_site_vpn_gateway.
func (ptsvg pointToSiteVpnGatewayAttributes) VpnServerConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(ptsvg.ref.Append("vpn_server_configuration_id"))
}

func (ptsvg pointToSiteVpnGatewayAttributes) ConnectionConfiguration() terra.ListValue[pointtositevpngateway.ConnectionConfigurationAttributes] {
	return terra.ReferenceAsList[pointtositevpngateway.ConnectionConfigurationAttributes](ptsvg.ref.Append("connection_configuration"))
}

func (ptsvg pointToSiteVpnGatewayAttributes) Timeouts() pointtositevpngateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[pointtositevpngateway.TimeoutsAttributes](ptsvg.ref.Append("timeouts"))
}

type pointToSiteVpnGatewayState struct {
	DnsServers                       []string                                             `json:"dns_servers"`
	Id                               string                                               `json:"id"`
	Location                         string                                               `json:"location"`
	Name                             string                                               `json:"name"`
	ResourceGroupName                string                                               `json:"resource_group_name"`
	RoutingPreferenceInternetEnabled bool                                                 `json:"routing_preference_internet_enabled"`
	ScaleUnit                        float64                                              `json:"scale_unit"`
	Tags                             map[string]string                                    `json:"tags"`
	VirtualHubId                     string                                               `json:"virtual_hub_id"`
	VpnServerConfigurationId         string                                               `json:"vpn_server_configuration_id"`
	ConnectionConfiguration          []pointtositevpngateway.ConnectionConfigurationState `json:"connection_configuration"`
	Timeouts                         *pointtositevpngateway.TimeoutsState                 `json:"timeouts"`
}
