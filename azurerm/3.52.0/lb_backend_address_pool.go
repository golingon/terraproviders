// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	lbbackendaddresspool "github.com/golingon/terraproviders/azurerm/3.52.0/lbbackendaddresspool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLbBackendAddressPool creates a new instance of [LbBackendAddressPool].
func NewLbBackendAddressPool(name string, args LbBackendAddressPoolArgs) *LbBackendAddressPool {
	return &LbBackendAddressPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LbBackendAddressPool)(nil)

// LbBackendAddressPool represents the Terraform resource azurerm_lb_backend_address_pool.
type LbBackendAddressPool struct {
	Name      string
	Args      LbBackendAddressPoolArgs
	state     *lbBackendAddressPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LbBackendAddressPool].
func (lbap *LbBackendAddressPool) Type() string {
	return "azurerm_lb_backend_address_pool"
}

// LocalName returns the local name for [LbBackendAddressPool].
func (lbap *LbBackendAddressPool) LocalName() string {
	return lbap.Name
}

// Configuration returns the configuration (args) for [LbBackendAddressPool].
func (lbap *LbBackendAddressPool) Configuration() interface{} {
	return lbap.Args
}

// DependOn is used for other resources to depend on [LbBackendAddressPool].
func (lbap *LbBackendAddressPool) DependOn() terra.Reference {
	return terra.ReferenceResource(lbap)
}

// Dependencies returns the list of resources [LbBackendAddressPool] depends_on.
func (lbap *LbBackendAddressPool) Dependencies() terra.Dependencies {
	return lbap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LbBackendAddressPool].
func (lbap *LbBackendAddressPool) LifecycleManagement() *terra.Lifecycle {
	return lbap.Lifecycle
}

// Attributes returns the attributes for [LbBackendAddressPool].
func (lbap *LbBackendAddressPool) Attributes() lbBackendAddressPoolAttributes {
	return lbBackendAddressPoolAttributes{ref: terra.ReferenceResource(lbap)}
}

// ImportState imports the given attribute values into [LbBackendAddressPool]'s state.
func (lbap *LbBackendAddressPool) ImportState(av io.Reader) error {
	lbap.state = &lbBackendAddressPoolState{}
	if err := json.NewDecoder(av).Decode(lbap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lbap.Type(), lbap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LbBackendAddressPool] has state.
func (lbap *LbBackendAddressPool) State() (*lbBackendAddressPoolState, bool) {
	return lbap.state, lbap.state != nil
}

// StateMust returns the state for [LbBackendAddressPool]. Panics if the state is nil.
func (lbap *LbBackendAddressPool) StateMust() *lbBackendAddressPoolState {
	if lbap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lbap.Type(), lbap.LocalName()))
	}
	return lbap.state
}

// LbBackendAddressPoolArgs contains the configurations for azurerm_lb_backend_address_pool.
type LbBackendAddressPoolArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoadbalancerId: string, required
	LoadbalancerId terra.StringValue `hcl:"loadbalancer_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// VirtualNetworkId: string, optional
	VirtualNetworkId terra.StringValue `hcl:"virtual_network_id,attr"`
	// Timeouts: optional
	Timeouts *lbbackendaddresspool.Timeouts `hcl:"timeouts,block"`
	// TunnelInterface: min=0
	TunnelInterface []lbbackendaddresspool.TunnelInterface `hcl:"tunnel_interface,block" validate:"min=0"`
}
type lbBackendAddressPoolAttributes struct {
	ref terra.Reference
}

// BackendIpConfigurations returns a reference to field backend_ip_configurations of azurerm_lb_backend_address_pool.
func (lbap lbBackendAddressPoolAttributes) BackendIpConfigurations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lbap.ref.Append("backend_ip_configurations"))
}

// Id returns a reference to field id of azurerm_lb_backend_address_pool.
func (lbap lbBackendAddressPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lbap.ref.Append("id"))
}

// InboundNatRules returns a reference to field inbound_nat_rules of azurerm_lb_backend_address_pool.
func (lbap lbBackendAddressPoolAttributes) InboundNatRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lbap.ref.Append("inbound_nat_rules"))
}

// LoadBalancingRules returns a reference to field load_balancing_rules of azurerm_lb_backend_address_pool.
func (lbap lbBackendAddressPoolAttributes) LoadBalancingRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lbap.ref.Append("load_balancing_rules"))
}

// LoadbalancerId returns a reference to field loadbalancer_id of azurerm_lb_backend_address_pool.
func (lbap lbBackendAddressPoolAttributes) LoadbalancerId() terra.StringValue {
	return terra.ReferenceAsString(lbap.ref.Append("loadbalancer_id"))
}

// Name returns a reference to field name of azurerm_lb_backend_address_pool.
func (lbap lbBackendAddressPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lbap.ref.Append("name"))
}

// OutboundRules returns a reference to field outbound_rules of azurerm_lb_backend_address_pool.
func (lbap lbBackendAddressPoolAttributes) OutboundRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lbap.ref.Append("outbound_rules"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_lb_backend_address_pool.
func (lbap lbBackendAddressPoolAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(lbap.ref.Append("virtual_network_id"))
}

func (lbap lbBackendAddressPoolAttributes) Timeouts() lbbackendaddresspool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lbbackendaddresspool.TimeoutsAttributes](lbap.ref.Append("timeouts"))
}

func (lbap lbBackendAddressPoolAttributes) TunnelInterface() terra.ListValue[lbbackendaddresspool.TunnelInterfaceAttributes] {
	return terra.ReferenceAsList[lbbackendaddresspool.TunnelInterfaceAttributes](lbap.ref.Append("tunnel_interface"))
}

type lbBackendAddressPoolState struct {
	BackendIpConfigurations []string                                    `json:"backend_ip_configurations"`
	Id                      string                                      `json:"id"`
	InboundNatRules         []string                                    `json:"inbound_nat_rules"`
	LoadBalancingRules      []string                                    `json:"load_balancing_rules"`
	LoadbalancerId          string                                      `json:"loadbalancer_id"`
	Name                    string                                      `json:"name"`
	OutboundRules           []string                                    `json:"outbound_rules"`
	VirtualNetworkId        string                                      `json:"virtual_network_id"`
	Timeouts                *lbbackendaddresspool.TimeoutsState         `json:"timeouts"`
	TunnelInterface         []lbbackendaddresspool.TunnelInterfaceState `json:"tunnel_interface"`
}
