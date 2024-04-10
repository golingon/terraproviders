// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	databricksvirtualnetworkpeering "github.com/golingon/terraproviders/azurerm/3.98.0/databricksvirtualnetworkpeering"
	"io"
)

// NewDatabricksVirtualNetworkPeering creates a new instance of [DatabricksVirtualNetworkPeering].
func NewDatabricksVirtualNetworkPeering(name string, args DatabricksVirtualNetworkPeeringArgs) *DatabricksVirtualNetworkPeering {
	return &DatabricksVirtualNetworkPeering{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatabricksVirtualNetworkPeering)(nil)

// DatabricksVirtualNetworkPeering represents the Terraform resource azurerm_databricks_virtual_network_peering.
type DatabricksVirtualNetworkPeering struct {
	Name      string
	Args      DatabricksVirtualNetworkPeeringArgs
	state     *databricksVirtualNetworkPeeringState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatabricksVirtualNetworkPeering].
func (dvnp *DatabricksVirtualNetworkPeering) Type() string {
	return "azurerm_databricks_virtual_network_peering"
}

// LocalName returns the local name for [DatabricksVirtualNetworkPeering].
func (dvnp *DatabricksVirtualNetworkPeering) LocalName() string {
	return dvnp.Name
}

// Configuration returns the configuration (args) for [DatabricksVirtualNetworkPeering].
func (dvnp *DatabricksVirtualNetworkPeering) Configuration() interface{} {
	return dvnp.Args
}

// DependOn is used for other resources to depend on [DatabricksVirtualNetworkPeering].
func (dvnp *DatabricksVirtualNetworkPeering) DependOn() terra.Reference {
	return terra.ReferenceResource(dvnp)
}

// Dependencies returns the list of resources [DatabricksVirtualNetworkPeering] depends_on.
func (dvnp *DatabricksVirtualNetworkPeering) Dependencies() terra.Dependencies {
	return dvnp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatabricksVirtualNetworkPeering].
func (dvnp *DatabricksVirtualNetworkPeering) LifecycleManagement() *terra.Lifecycle {
	return dvnp.Lifecycle
}

// Attributes returns the attributes for [DatabricksVirtualNetworkPeering].
func (dvnp *DatabricksVirtualNetworkPeering) Attributes() databricksVirtualNetworkPeeringAttributes {
	return databricksVirtualNetworkPeeringAttributes{ref: terra.ReferenceResource(dvnp)}
}

// ImportState imports the given attribute values into [DatabricksVirtualNetworkPeering]'s state.
func (dvnp *DatabricksVirtualNetworkPeering) ImportState(av io.Reader) error {
	dvnp.state = &databricksVirtualNetworkPeeringState{}
	if err := json.NewDecoder(av).Decode(dvnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dvnp.Type(), dvnp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatabricksVirtualNetworkPeering] has state.
func (dvnp *DatabricksVirtualNetworkPeering) State() (*databricksVirtualNetworkPeeringState, bool) {
	return dvnp.state, dvnp.state != nil
}

// StateMust returns the state for [DatabricksVirtualNetworkPeering]. Panics if the state is nil.
func (dvnp *DatabricksVirtualNetworkPeering) StateMust() *databricksVirtualNetworkPeeringState {
	if dvnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dvnp.Type(), dvnp.LocalName()))
	}
	return dvnp.state
}

// DatabricksVirtualNetworkPeeringArgs contains the configurations for azurerm_databricks_virtual_network_peering.
type DatabricksVirtualNetworkPeeringArgs struct {
	// AllowForwardedTraffic: bool, optional
	AllowForwardedTraffic terra.BoolValue `hcl:"allow_forwarded_traffic,attr"`
	// AllowGatewayTransit: bool, optional
	AllowGatewayTransit terra.BoolValue `hcl:"allow_gateway_transit,attr"`
	// AllowVirtualNetworkAccess: bool, optional
	AllowVirtualNetworkAccess terra.BoolValue `hcl:"allow_virtual_network_access,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RemoteAddressSpacePrefixes: list of string, required
	RemoteAddressSpacePrefixes terra.ListValue[terra.StringValue] `hcl:"remote_address_space_prefixes,attr" validate:"required"`
	// RemoteVirtualNetworkId: string, required
	RemoteVirtualNetworkId terra.StringValue `hcl:"remote_virtual_network_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// UseRemoteGateways: bool, optional
	UseRemoteGateways terra.BoolValue `hcl:"use_remote_gateways,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *databricksvirtualnetworkpeering.Timeouts `hcl:"timeouts,block"`
}
type databricksVirtualNetworkPeeringAttributes struct {
	ref terra.Reference
}

// AddressSpacePrefixes returns a reference to field address_space_prefixes of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) AddressSpacePrefixes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dvnp.ref.Append("address_space_prefixes"))
}

// AllowForwardedTraffic returns a reference to field allow_forwarded_traffic of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) AllowForwardedTraffic() terra.BoolValue {
	return terra.ReferenceAsBool(dvnp.ref.Append("allow_forwarded_traffic"))
}

// AllowGatewayTransit returns a reference to field allow_gateway_transit of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) AllowGatewayTransit() terra.BoolValue {
	return terra.ReferenceAsBool(dvnp.ref.Append("allow_gateway_transit"))
}

// AllowVirtualNetworkAccess returns a reference to field allow_virtual_network_access of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) AllowVirtualNetworkAccess() terra.BoolValue {
	return terra.ReferenceAsBool(dvnp.ref.Append("allow_virtual_network_access"))
}

// Id returns a reference to field id of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dvnp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dvnp.ref.Append("name"))
}

// RemoteAddressSpacePrefixes returns a reference to field remote_address_space_prefixes of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) RemoteAddressSpacePrefixes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dvnp.ref.Append("remote_address_space_prefixes"))
}

// RemoteVirtualNetworkId returns a reference to field remote_virtual_network_id of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) RemoteVirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(dvnp.ref.Append("remote_virtual_network_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dvnp.ref.Append("resource_group_name"))
}

// UseRemoteGateways returns a reference to field use_remote_gateways of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) UseRemoteGateways() terra.BoolValue {
	return terra.ReferenceAsBool(dvnp.ref.Append("use_remote_gateways"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(dvnp.ref.Append("virtual_network_id"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_databricks_virtual_network_peering.
func (dvnp databricksVirtualNetworkPeeringAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(dvnp.ref.Append("workspace_id"))
}

func (dvnp databricksVirtualNetworkPeeringAttributes) Timeouts() databricksvirtualnetworkpeering.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databricksvirtualnetworkpeering.TimeoutsAttributes](dvnp.ref.Append("timeouts"))
}

type databricksVirtualNetworkPeeringState struct {
	AddressSpacePrefixes       []string                                       `json:"address_space_prefixes"`
	AllowForwardedTraffic      bool                                           `json:"allow_forwarded_traffic"`
	AllowGatewayTransit        bool                                           `json:"allow_gateway_transit"`
	AllowVirtualNetworkAccess  bool                                           `json:"allow_virtual_network_access"`
	Id                         string                                         `json:"id"`
	Name                       string                                         `json:"name"`
	RemoteAddressSpacePrefixes []string                                       `json:"remote_address_space_prefixes"`
	RemoteVirtualNetworkId     string                                         `json:"remote_virtual_network_id"`
	ResourceGroupName          string                                         `json:"resource_group_name"`
	UseRemoteGateways          bool                                           `json:"use_remote_gateways"`
	VirtualNetworkId           string                                         `json:"virtual_network_id"`
	WorkspaceId                string                                         `json:"workspace_id"`
	Timeouts                   *databricksvirtualnetworkpeering.TimeoutsState `json:"timeouts"`
}
