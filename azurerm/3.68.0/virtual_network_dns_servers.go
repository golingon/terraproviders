// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualnetworkdnsservers "github.com/golingon/terraproviders/azurerm/3.68.0/virtualnetworkdnsservers"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualNetworkDnsServers creates a new instance of [VirtualNetworkDnsServers].
func NewVirtualNetworkDnsServers(name string, args VirtualNetworkDnsServersArgs) *VirtualNetworkDnsServers {
	return &VirtualNetworkDnsServers{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualNetworkDnsServers)(nil)

// VirtualNetworkDnsServers represents the Terraform resource azurerm_virtual_network_dns_servers.
type VirtualNetworkDnsServers struct {
	Name      string
	Args      VirtualNetworkDnsServersArgs
	state     *virtualNetworkDnsServersState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualNetworkDnsServers].
func (vnds *VirtualNetworkDnsServers) Type() string {
	return "azurerm_virtual_network_dns_servers"
}

// LocalName returns the local name for [VirtualNetworkDnsServers].
func (vnds *VirtualNetworkDnsServers) LocalName() string {
	return vnds.Name
}

// Configuration returns the configuration (args) for [VirtualNetworkDnsServers].
func (vnds *VirtualNetworkDnsServers) Configuration() interface{} {
	return vnds.Args
}

// DependOn is used for other resources to depend on [VirtualNetworkDnsServers].
func (vnds *VirtualNetworkDnsServers) DependOn() terra.Reference {
	return terra.ReferenceResource(vnds)
}

// Dependencies returns the list of resources [VirtualNetworkDnsServers] depends_on.
func (vnds *VirtualNetworkDnsServers) Dependencies() terra.Dependencies {
	return vnds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualNetworkDnsServers].
func (vnds *VirtualNetworkDnsServers) LifecycleManagement() *terra.Lifecycle {
	return vnds.Lifecycle
}

// Attributes returns the attributes for [VirtualNetworkDnsServers].
func (vnds *VirtualNetworkDnsServers) Attributes() virtualNetworkDnsServersAttributes {
	return virtualNetworkDnsServersAttributes{ref: terra.ReferenceResource(vnds)}
}

// ImportState imports the given attribute values into [VirtualNetworkDnsServers]'s state.
func (vnds *VirtualNetworkDnsServers) ImportState(av io.Reader) error {
	vnds.state = &virtualNetworkDnsServersState{}
	if err := json.NewDecoder(av).Decode(vnds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vnds.Type(), vnds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualNetworkDnsServers] has state.
func (vnds *VirtualNetworkDnsServers) State() (*virtualNetworkDnsServersState, bool) {
	return vnds.state, vnds.state != nil
}

// StateMust returns the state for [VirtualNetworkDnsServers]. Panics if the state is nil.
func (vnds *VirtualNetworkDnsServers) StateMust() *virtualNetworkDnsServersState {
	if vnds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vnds.Type(), vnds.LocalName()))
	}
	return vnds.state
}

// VirtualNetworkDnsServersArgs contains the configurations for azurerm_virtual_network_dns_servers.
type VirtualNetworkDnsServersArgs struct {
	// DnsServers: list of string, optional
	DnsServers terra.ListValue[terra.StringValue] `hcl:"dns_servers,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VirtualNetworkId: string, required
	VirtualNetworkId terra.StringValue `hcl:"virtual_network_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *virtualnetworkdnsservers.Timeouts `hcl:"timeouts,block"`
}
type virtualNetworkDnsServersAttributes struct {
	ref terra.Reference
}

// DnsServers returns a reference to field dns_servers of azurerm_virtual_network_dns_servers.
func (vnds virtualNetworkDnsServersAttributes) DnsServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vnds.ref.Append("dns_servers"))
}

// Id returns a reference to field id of azurerm_virtual_network_dns_servers.
func (vnds virtualNetworkDnsServersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vnds.ref.Append("id"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_virtual_network_dns_servers.
func (vnds virtualNetworkDnsServersAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(vnds.ref.Append("virtual_network_id"))
}

func (vnds virtualNetworkDnsServersAttributes) Timeouts() virtualnetworkdnsservers.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualnetworkdnsservers.TimeoutsAttributes](vnds.ref.Append("timeouts"))
}

type virtualNetworkDnsServersState struct {
	DnsServers       []string                                `json:"dns_servers"`
	Id               string                                  `json:"id"`
	VirtualNetworkId string                                  `json:"virtual_network_id"`
	Timeouts         *virtualnetworkdnsservers.TimeoutsState `json:"timeouts"`
}
