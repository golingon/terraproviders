// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	lbbackendaddresspool "github.com/golingon/terraproviders/azurerm/3.49.0/lbbackendaddresspool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewLbBackendAddressPool(name string, args LbBackendAddressPoolArgs) *LbBackendAddressPool {
	return &LbBackendAddressPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LbBackendAddressPool)(nil)

type LbBackendAddressPool struct {
	Name  string
	Args  LbBackendAddressPoolArgs
	state *lbBackendAddressPoolState
}

func (lbap *LbBackendAddressPool) Type() string {
	return "azurerm_lb_backend_address_pool"
}

func (lbap *LbBackendAddressPool) LocalName() string {
	return lbap.Name
}

func (lbap *LbBackendAddressPool) Configuration() interface{} {
	return lbap.Args
}

func (lbap *LbBackendAddressPool) Attributes() lbBackendAddressPoolAttributes {
	return lbBackendAddressPoolAttributes{ref: terra.ReferenceResource(lbap)}
}

func (lbap *LbBackendAddressPool) ImportState(av io.Reader) error {
	lbap.state = &lbBackendAddressPoolState{}
	if err := json.NewDecoder(av).Decode(lbap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lbap.Type(), lbap.LocalName(), err)
	}
	return nil
}

func (lbap *LbBackendAddressPool) State() (*lbBackendAddressPoolState, bool) {
	return lbap.state, lbap.state != nil
}

func (lbap *LbBackendAddressPool) StateMust() *lbBackendAddressPoolState {
	if lbap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lbap.Type(), lbap.LocalName()))
	}
	return lbap.state
}

func (lbap *LbBackendAddressPool) DependOn() terra.Reference {
	return terra.ReferenceResource(lbap)
}

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
	// DependsOn contains resources that LbBackendAddressPool depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type lbBackendAddressPoolAttributes struct {
	ref terra.Reference
}

func (lbap lbBackendAddressPoolAttributes) BackendIpConfigurations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](lbap.ref.Append("backend_ip_configurations"))
}

func (lbap lbBackendAddressPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceString(lbap.ref.Append("id"))
}

func (lbap lbBackendAddressPoolAttributes) InboundNatRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](lbap.ref.Append("inbound_nat_rules"))
}

func (lbap lbBackendAddressPoolAttributes) LoadBalancingRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](lbap.ref.Append("load_balancing_rules"))
}

func (lbap lbBackendAddressPoolAttributes) LoadbalancerId() terra.StringValue {
	return terra.ReferenceString(lbap.ref.Append("loadbalancer_id"))
}

func (lbap lbBackendAddressPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceString(lbap.ref.Append("name"))
}

func (lbap lbBackendAddressPoolAttributes) OutboundRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](lbap.ref.Append("outbound_rules"))
}

func (lbap lbBackendAddressPoolAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceString(lbap.ref.Append("virtual_network_id"))
}

func (lbap lbBackendAddressPoolAttributes) Timeouts() lbbackendaddresspool.TimeoutsAttributes {
	return terra.ReferenceSingle[lbbackendaddresspool.TimeoutsAttributes](lbap.ref.Append("timeouts"))
}

func (lbap lbBackendAddressPoolAttributes) TunnelInterface() terra.ListValue[lbbackendaddresspool.TunnelInterfaceAttributes] {
	return terra.ReferenceList[lbbackendaddresspool.TunnelInterfaceAttributes](lbap.ref.Append("tunnel_interface"))
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
