// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualnetwork "github.com/golingon/terraproviders/azurerm/3.49.0/virtualnetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualNetwork creates a new instance of [VirtualNetwork].
func NewVirtualNetwork(name string, args VirtualNetworkArgs) *VirtualNetwork {
	return &VirtualNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualNetwork)(nil)

// VirtualNetwork represents the Terraform resource azurerm_virtual_network.
type VirtualNetwork struct {
	Name      string
	Args      VirtualNetworkArgs
	state     *virtualNetworkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualNetwork].
func (vn *VirtualNetwork) Type() string {
	return "azurerm_virtual_network"
}

// LocalName returns the local name for [VirtualNetwork].
func (vn *VirtualNetwork) LocalName() string {
	return vn.Name
}

// Configuration returns the configuration (args) for [VirtualNetwork].
func (vn *VirtualNetwork) Configuration() interface{} {
	return vn.Args
}

// DependOn is used for other resources to depend on [VirtualNetwork].
func (vn *VirtualNetwork) DependOn() terra.Reference {
	return terra.ReferenceResource(vn)
}

// Dependencies returns the list of resources [VirtualNetwork] depends_on.
func (vn *VirtualNetwork) Dependencies() terra.Dependencies {
	return vn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualNetwork].
func (vn *VirtualNetwork) LifecycleManagement() *terra.Lifecycle {
	return vn.Lifecycle
}

// Attributes returns the attributes for [VirtualNetwork].
func (vn *VirtualNetwork) Attributes() virtualNetworkAttributes {
	return virtualNetworkAttributes{ref: terra.ReferenceResource(vn)}
}

// ImportState imports the given attribute values into [VirtualNetwork]'s state.
func (vn *VirtualNetwork) ImportState(av io.Reader) error {
	vn.state = &virtualNetworkState{}
	if err := json.NewDecoder(av).Decode(vn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vn.Type(), vn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualNetwork] has state.
func (vn *VirtualNetwork) State() (*virtualNetworkState, bool) {
	return vn.state, vn.state != nil
}

// StateMust returns the state for [VirtualNetwork]. Panics if the state is nil.
func (vn *VirtualNetwork) StateMust() *virtualNetworkState {
	if vn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vn.Type(), vn.LocalName()))
	}
	return vn.state
}

// VirtualNetworkArgs contains the configurations for azurerm_virtual_network.
type VirtualNetworkArgs struct {
	// AddressSpace: list of string, required
	AddressSpace terra.ListValue[terra.StringValue] `hcl:"address_space,attr" validate:"required"`
	// BgpCommunity: string, optional
	BgpCommunity terra.StringValue `hcl:"bgp_community,attr"`
	// DnsServers: list of string, optional
	DnsServers terra.ListValue[terra.StringValue] `hcl:"dns_servers,attr"`
	// EdgeZone: string, optional
	EdgeZone terra.StringValue `hcl:"edge_zone,attr"`
	// FlowTimeoutInMinutes: number, optional
	FlowTimeoutInMinutes terra.NumberValue `hcl:"flow_timeout_in_minutes,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Subnet: min=0
	Subnet []virtualnetwork.Subnet `hcl:"subnet,block" validate:"min=0"`
	// DdosProtectionPlan: optional
	DdosProtectionPlan *virtualnetwork.DdosProtectionPlan `hcl:"ddos_protection_plan,block"`
	// Timeouts: optional
	Timeouts *virtualnetwork.Timeouts `hcl:"timeouts,block"`
}
type virtualNetworkAttributes struct {
	ref terra.Reference
}

// AddressSpace returns a reference to field address_space of azurerm_virtual_network.
func (vn virtualNetworkAttributes) AddressSpace() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vn.ref.Append("address_space"))
}

// BgpCommunity returns a reference to field bgp_community of azurerm_virtual_network.
func (vn virtualNetworkAttributes) BgpCommunity() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("bgp_community"))
}

// DnsServers returns a reference to field dns_servers of azurerm_virtual_network.
func (vn virtualNetworkAttributes) DnsServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vn.ref.Append("dns_servers"))
}

// EdgeZone returns a reference to field edge_zone of azurerm_virtual_network.
func (vn virtualNetworkAttributes) EdgeZone() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("edge_zone"))
}

// FlowTimeoutInMinutes returns a reference to field flow_timeout_in_minutes of azurerm_virtual_network.
func (vn virtualNetworkAttributes) FlowTimeoutInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(vn.ref.Append("flow_timeout_in_minutes"))
}

// Guid returns a reference to field guid of azurerm_virtual_network.
func (vn virtualNetworkAttributes) Guid() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("guid"))
}

// Id returns a reference to field id of azurerm_virtual_network.
func (vn virtualNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_virtual_network.
func (vn virtualNetworkAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_network.
func (vn virtualNetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_network.
func (vn virtualNetworkAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_virtual_network.
func (vn virtualNetworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vn.ref.Append("tags"))
}

func (vn virtualNetworkAttributes) Subnet() terra.SetValue[virtualnetwork.SubnetAttributes] {
	return terra.ReferenceAsSet[virtualnetwork.SubnetAttributes](vn.ref.Append("subnet"))
}

func (vn virtualNetworkAttributes) DdosProtectionPlan() terra.ListValue[virtualnetwork.DdosProtectionPlanAttributes] {
	return terra.ReferenceAsList[virtualnetwork.DdosProtectionPlanAttributes](vn.ref.Append("ddos_protection_plan"))
}

func (vn virtualNetworkAttributes) Timeouts() virtualnetwork.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualnetwork.TimeoutsAttributes](vn.ref.Append("timeouts"))
}

type virtualNetworkState struct {
	AddressSpace         []string                                 `json:"address_space"`
	BgpCommunity         string                                   `json:"bgp_community"`
	DnsServers           []string                                 `json:"dns_servers"`
	EdgeZone             string                                   `json:"edge_zone"`
	FlowTimeoutInMinutes float64                                  `json:"flow_timeout_in_minutes"`
	Guid                 string                                   `json:"guid"`
	Id                   string                                   `json:"id"`
	Location             string                                   `json:"location"`
	Name                 string                                   `json:"name"`
	ResourceGroupName    string                                   `json:"resource_group_name"`
	Tags                 map[string]string                        `json:"tags"`
	Subnet               []virtualnetwork.SubnetState             `json:"subnet"`
	DdosProtectionPlan   []virtualnetwork.DdosProtectionPlanState `json:"ddos_protection_plan"`
	Timeouts             *virtualnetwork.TimeoutsState            `json:"timeouts"`
}
