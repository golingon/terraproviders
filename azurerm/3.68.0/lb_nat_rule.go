// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	lbnatrule "github.com/golingon/terraproviders/azurerm/3.68.0/lbnatrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLbNatRule creates a new instance of [LbNatRule].
func NewLbNatRule(name string, args LbNatRuleArgs) *LbNatRule {
	return &LbNatRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LbNatRule)(nil)

// LbNatRule represents the Terraform resource azurerm_lb_nat_rule.
type LbNatRule struct {
	Name      string
	Args      LbNatRuleArgs
	state     *lbNatRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LbNatRule].
func (lnr *LbNatRule) Type() string {
	return "azurerm_lb_nat_rule"
}

// LocalName returns the local name for [LbNatRule].
func (lnr *LbNatRule) LocalName() string {
	return lnr.Name
}

// Configuration returns the configuration (args) for [LbNatRule].
func (lnr *LbNatRule) Configuration() interface{} {
	return lnr.Args
}

// DependOn is used for other resources to depend on [LbNatRule].
func (lnr *LbNatRule) DependOn() terra.Reference {
	return terra.ReferenceResource(lnr)
}

// Dependencies returns the list of resources [LbNatRule] depends_on.
func (lnr *LbNatRule) Dependencies() terra.Dependencies {
	return lnr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LbNatRule].
func (lnr *LbNatRule) LifecycleManagement() *terra.Lifecycle {
	return lnr.Lifecycle
}

// Attributes returns the attributes for [LbNatRule].
func (lnr *LbNatRule) Attributes() lbNatRuleAttributes {
	return lbNatRuleAttributes{ref: terra.ReferenceResource(lnr)}
}

// ImportState imports the given attribute values into [LbNatRule]'s state.
func (lnr *LbNatRule) ImportState(av io.Reader) error {
	lnr.state = &lbNatRuleState{}
	if err := json.NewDecoder(av).Decode(lnr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lnr.Type(), lnr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LbNatRule] has state.
func (lnr *LbNatRule) State() (*lbNatRuleState, bool) {
	return lnr.state, lnr.state != nil
}

// StateMust returns the state for [LbNatRule]. Panics if the state is nil.
func (lnr *LbNatRule) StateMust() *lbNatRuleState {
	if lnr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lnr.Type(), lnr.LocalName()))
	}
	return lnr.state
}

// LbNatRuleArgs contains the configurations for azurerm_lb_nat_rule.
type LbNatRuleArgs struct {
	// BackendAddressPoolId: string, optional
	BackendAddressPoolId terra.StringValue `hcl:"backend_address_pool_id,attr"`
	// BackendPort: number, required
	BackendPort terra.NumberValue `hcl:"backend_port,attr" validate:"required"`
	// EnableFloatingIp: bool, optional
	EnableFloatingIp terra.BoolValue `hcl:"enable_floating_ip,attr"`
	// EnableTcpReset: bool, optional
	EnableTcpReset terra.BoolValue `hcl:"enable_tcp_reset,attr"`
	// FrontendIpConfigurationName: string, required
	FrontendIpConfigurationName terra.StringValue `hcl:"frontend_ip_configuration_name,attr" validate:"required"`
	// FrontendPort: number, optional
	FrontendPort terra.NumberValue `hcl:"frontend_port,attr"`
	// FrontendPortEnd: number, optional
	FrontendPortEnd terra.NumberValue `hcl:"frontend_port_end,attr"`
	// FrontendPortStart: number, optional
	FrontendPortStart terra.NumberValue `hcl:"frontend_port_start,attr"`
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
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *lbnatrule.Timeouts `hcl:"timeouts,block"`
}
type lbNatRuleAttributes struct {
	ref terra.Reference
}

// BackendAddressPoolId returns a reference to field backend_address_pool_id of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) BackendAddressPoolId() terra.StringValue {
	return terra.ReferenceAsString(lnr.ref.Append("backend_address_pool_id"))
}

// BackendIpConfigurationId returns a reference to field backend_ip_configuration_id of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) BackendIpConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(lnr.ref.Append("backend_ip_configuration_id"))
}

// BackendPort returns a reference to field backend_port of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) BackendPort() terra.NumberValue {
	return terra.ReferenceAsNumber(lnr.ref.Append("backend_port"))
}

// EnableFloatingIp returns a reference to field enable_floating_ip of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) EnableFloatingIp() terra.BoolValue {
	return terra.ReferenceAsBool(lnr.ref.Append("enable_floating_ip"))
}

// EnableTcpReset returns a reference to field enable_tcp_reset of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) EnableTcpReset() terra.BoolValue {
	return terra.ReferenceAsBool(lnr.ref.Append("enable_tcp_reset"))
}

// FrontendIpConfigurationId returns a reference to field frontend_ip_configuration_id of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) FrontendIpConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(lnr.ref.Append("frontend_ip_configuration_id"))
}

// FrontendIpConfigurationName returns a reference to field frontend_ip_configuration_name of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) FrontendIpConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(lnr.ref.Append("frontend_ip_configuration_name"))
}

// FrontendPort returns a reference to field frontend_port of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) FrontendPort() terra.NumberValue {
	return terra.ReferenceAsNumber(lnr.ref.Append("frontend_port"))
}

// FrontendPortEnd returns a reference to field frontend_port_end of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) FrontendPortEnd() terra.NumberValue {
	return terra.ReferenceAsNumber(lnr.ref.Append("frontend_port_end"))
}

// FrontendPortStart returns a reference to field frontend_port_start of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) FrontendPortStart() terra.NumberValue {
	return terra.ReferenceAsNumber(lnr.ref.Append("frontend_port_start"))
}

// Id returns a reference to field id of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lnr.ref.Append("id"))
}

// IdleTimeoutInMinutes returns a reference to field idle_timeout_in_minutes of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) IdleTimeoutInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(lnr.ref.Append("idle_timeout_in_minutes"))
}

// LoadbalancerId returns a reference to field loadbalancer_id of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) LoadbalancerId() terra.StringValue {
	return terra.ReferenceAsString(lnr.ref.Append("loadbalancer_id"))
}

// Name returns a reference to field name of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lnr.ref.Append("name"))
}

// Protocol returns a reference to field protocol of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(lnr.ref.Append("protocol"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_lb_nat_rule.
func (lnr lbNatRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lnr.ref.Append("resource_group_name"))
}

func (lnr lbNatRuleAttributes) Timeouts() lbnatrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lbnatrule.TimeoutsAttributes](lnr.ref.Append("timeouts"))
}

type lbNatRuleState struct {
	BackendAddressPoolId        string                   `json:"backend_address_pool_id"`
	BackendIpConfigurationId    string                   `json:"backend_ip_configuration_id"`
	BackendPort                 float64                  `json:"backend_port"`
	EnableFloatingIp            bool                     `json:"enable_floating_ip"`
	EnableTcpReset              bool                     `json:"enable_tcp_reset"`
	FrontendIpConfigurationId   string                   `json:"frontend_ip_configuration_id"`
	FrontendIpConfigurationName string                   `json:"frontend_ip_configuration_name"`
	FrontendPort                float64                  `json:"frontend_port"`
	FrontendPortEnd             float64                  `json:"frontend_port_end"`
	FrontendPortStart           float64                  `json:"frontend_port_start"`
	Id                          string                   `json:"id"`
	IdleTimeoutInMinutes        float64                  `json:"idle_timeout_in_minutes"`
	LoadbalancerId              string                   `json:"loadbalancer_id"`
	Name                        string                   `json:"name"`
	Protocol                    string                   `json:"protocol"`
	ResourceGroupName           string                   `json:"resource_group_name"`
	Timeouts                    *lbnatrule.TimeoutsState `json:"timeouts"`
}
