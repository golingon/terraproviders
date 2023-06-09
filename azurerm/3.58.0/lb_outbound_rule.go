// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	lboutboundrule "github.com/golingon/terraproviders/azurerm/3.58.0/lboutboundrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLbOutboundRule creates a new instance of [LbOutboundRule].
func NewLbOutboundRule(name string, args LbOutboundRuleArgs) *LbOutboundRule {
	return &LbOutboundRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LbOutboundRule)(nil)

// LbOutboundRule represents the Terraform resource azurerm_lb_outbound_rule.
type LbOutboundRule struct {
	Name      string
	Args      LbOutboundRuleArgs
	state     *lbOutboundRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LbOutboundRule].
func (lor *LbOutboundRule) Type() string {
	return "azurerm_lb_outbound_rule"
}

// LocalName returns the local name for [LbOutboundRule].
func (lor *LbOutboundRule) LocalName() string {
	return lor.Name
}

// Configuration returns the configuration (args) for [LbOutboundRule].
func (lor *LbOutboundRule) Configuration() interface{} {
	return lor.Args
}

// DependOn is used for other resources to depend on [LbOutboundRule].
func (lor *LbOutboundRule) DependOn() terra.Reference {
	return terra.ReferenceResource(lor)
}

// Dependencies returns the list of resources [LbOutboundRule] depends_on.
func (lor *LbOutboundRule) Dependencies() terra.Dependencies {
	return lor.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LbOutboundRule].
func (lor *LbOutboundRule) LifecycleManagement() *terra.Lifecycle {
	return lor.Lifecycle
}

// Attributes returns the attributes for [LbOutboundRule].
func (lor *LbOutboundRule) Attributes() lbOutboundRuleAttributes {
	return lbOutboundRuleAttributes{ref: terra.ReferenceResource(lor)}
}

// ImportState imports the given attribute values into [LbOutboundRule]'s state.
func (lor *LbOutboundRule) ImportState(av io.Reader) error {
	lor.state = &lbOutboundRuleState{}
	if err := json.NewDecoder(av).Decode(lor.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lor.Type(), lor.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LbOutboundRule] has state.
func (lor *LbOutboundRule) State() (*lbOutboundRuleState, bool) {
	return lor.state, lor.state != nil
}

// StateMust returns the state for [LbOutboundRule]. Panics if the state is nil.
func (lor *LbOutboundRule) StateMust() *lbOutboundRuleState {
	if lor.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lor.Type(), lor.LocalName()))
	}
	return lor.state
}

// LbOutboundRuleArgs contains the configurations for azurerm_lb_outbound_rule.
type LbOutboundRuleArgs struct {
	// AllocatedOutboundPorts: number, optional
	AllocatedOutboundPorts terra.NumberValue `hcl:"allocated_outbound_ports,attr"`
	// BackendAddressPoolId: string, required
	BackendAddressPoolId terra.StringValue `hcl:"backend_address_pool_id,attr" validate:"required"`
	// EnableTcpReset: bool, optional
	EnableTcpReset terra.BoolValue `hcl:"enable_tcp_reset,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdleTimeoutInMinutes: number, optional
	IdleTimeoutInMinutes terra.NumberValue `hcl:"idle_timeout_in_minutes,attr"`
	// LoadbalancerId: string, required
	LoadbalancerId terra.StringValue `hcl:"loadbalancer_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// FrontendIpConfiguration: min=0
	FrontendIpConfiguration []lboutboundrule.FrontendIpConfiguration `hcl:"frontend_ip_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *lboutboundrule.Timeouts `hcl:"timeouts,block"`
}
type lbOutboundRuleAttributes struct {
	ref terra.Reference
}

// AllocatedOutboundPorts returns a reference to field allocated_outbound_ports of azurerm_lb_outbound_rule.
func (lor lbOutboundRuleAttributes) AllocatedOutboundPorts() terra.NumberValue {
	return terra.ReferenceAsNumber(lor.ref.Append("allocated_outbound_ports"))
}

// BackendAddressPoolId returns a reference to field backend_address_pool_id of azurerm_lb_outbound_rule.
func (lor lbOutboundRuleAttributes) BackendAddressPoolId() terra.StringValue {
	return terra.ReferenceAsString(lor.ref.Append("backend_address_pool_id"))
}

// EnableTcpReset returns a reference to field enable_tcp_reset of azurerm_lb_outbound_rule.
func (lor lbOutboundRuleAttributes) EnableTcpReset() terra.BoolValue {
	return terra.ReferenceAsBool(lor.ref.Append("enable_tcp_reset"))
}

// Id returns a reference to field id of azurerm_lb_outbound_rule.
func (lor lbOutboundRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lor.ref.Append("id"))
}

// IdleTimeoutInMinutes returns a reference to field idle_timeout_in_minutes of azurerm_lb_outbound_rule.
func (lor lbOutboundRuleAttributes) IdleTimeoutInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(lor.ref.Append("idle_timeout_in_minutes"))
}

// LoadbalancerId returns a reference to field loadbalancer_id of azurerm_lb_outbound_rule.
func (lor lbOutboundRuleAttributes) LoadbalancerId() terra.StringValue {
	return terra.ReferenceAsString(lor.ref.Append("loadbalancer_id"))
}

// Name returns a reference to field name of azurerm_lb_outbound_rule.
func (lor lbOutboundRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lor.ref.Append("name"))
}

// Protocol returns a reference to field protocol of azurerm_lb_outbound_rule.
func (lor lbOutboundRuleAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(lor.ref.Append("protocol"))
}

func (lor lbOutboundRuleAttributes) FrontendIpConfiguration() terra.ListValue[lboutboundrule.FrontendIpConfigurationAttributes] {
	return terra.ReferenceAsList[lboutboundrule.FrontendIpConfigurationAttributes](lor.ref.Append("frontend_ip_configuration"))
}

func (lor lbOutboundRuleAttributes) Timeouts() lboutboundrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lboutboundrule.TimeoutsAttributes](lor.ref.Append("timeouts"))
}

type lbOutboundRuleState struct {
	AllocatedOutboundPorts  float64                                       `json:"allocated_outbound_ports"`
	BackendAddressPoolId    string                                        `json:"backend_address_pool_id"`
	EnableTcpReset          bool                                          `json:"enable_tcp_reset"`
	Id                      string                                        `json:"id"`
	IdleTimeoutInMinutes    float64                                       `json:"idle_timeout_in_minutes"`
	LoadbalancerId          string                                        `json:"loadbalancer_id"`
	Name                    string                                        `json:"name"`
	Protocol                string                                        `json:"protocol"`
	FrontendIpConfiguration []lboutboundrule.FrontendIpConfigurationState `json:"frontend_ip_configuration"`
	Timeouts                *lboutboundrule.TimeoutsState                 `json:"timeouts"`
}
