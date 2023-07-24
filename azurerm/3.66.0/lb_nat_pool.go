// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	lbnatpool "github.com/golingon/terraproviders/azurerm/3.66.0/lbnatpool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLbNatPool creates a new instance of [LbNatPool].
func NewLbNatPool(name string, args LbNatPoolArgs) *LbNatPool {
	return &LbNatPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LbNatPool)(nil)

// LbNatPool represents the Terraform resource azurerm_lb_nat_pool.
type LbNatPool struct {
	Name      string
	Args      LbNatPoolArgs
	state     *lbNatPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LbNatPool].
func (lnp *LbNatPool) Type() string {
	return "azurerm_lb_nat_pool"
}

// LocalName returns the local name for [LbNatPool].
func (lnp *LbNatPool) LocalName() string {
	return lnp.Name
}

// Configuration returns the configuration (args) for [LbNatPool].
func (lnp *LbNatPool) Configuration() interface{} {
	return lnp.Args
}

// DependOn is used for other resources to depend on [LbNatPool].
func (lnp *LbNatPool) DependOn() terra.Reference {
	return terra.ReferenceResource(lnp)
}

// Dependencies returns the list of resources [LbNatPool] depends_on.
func (lnp *LbNatPool) Dependencies() terra.Dependencies {
	return lnp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LbNatPool].
func (lnp *LbNatPool) LifecycleManagement() *terra.Lifecycle {
	return lnp.Lifecycle
}

// Attributes returns the attributes for [LbNatPool].
func (lnp *LbNatPool) Attributes() lbNatPoolAttributes {
	return lbNatPoolAttributes{ref: terra.ReferenceResource(lnp)}
}

// ImportState imports the given attribute values into [LbNatPool]'s state.
func (lnp *LbNatPool) ImportState(av io.Reader) error {
	lnp.state = &lbNatPoolState{}
	if err := json.NewDecoder(av).Decode(lnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lnp.Type(), lnp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LbNatPool] has state.
func (lnp *LbNatPool) State() (*lbNatPoolState, bool) {
	return lnp.state, lnp.state != nil
}

// StateMust returns the state for [LbNatPool]. Panics if the state is nil.
func (lnp *LbNatPool) StateMust() *lbNatPoolState {
	if lnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lnp.Type(), lnp.LocalName()))
	}
	return lnp.state
}

// LbNatPoolArgs contains the configurations for azurerm_lb_nat_pool.
type LbNatPoolArgs struct {
	// BackendPort: number, required
	BackendPort terra.NumberValue `hcl:"backend_port,attr" validate:"required"`
	// FloatingIpEnabled: bool, optional
	FloatingIpEnabled terra.BoolValue `hcl:"floating_ip_enabled,attr"`
	// FrontendIpConfigurationName: string, required
	FrontendIpConfigurationName terra.StringValue `hcl:"frontend_ip_configuration_name,attr" validate:"required"`
	// FrontendPortEnd: number, required
	FrontendPortEnd terra.NumberValue `hcl:"frontend_port_end,attr" validate:"required"`
	// FrontendPortStart: number, required
	FrontendPortStart terra.NumberValue `hcl:"frontend_port_start,attr" validate:"required"`
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
	// TcpResetEnabled: bool, optional
	TcpResetEnabled terra.BoolValue `hcl:"tcp_reset_enabled,attr"`
	// Timeouts: optional
	Timeouts *lbnatpool.Timeouts `hcl:"timeouts,block"`
}
type lbNatPoolAttributes struct {
	ref terra.Reference
}

// BackendPort returns a reference to field backend_port of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) BackendPort() terra.NumberValue {
	return terra.ReferenceAsNumber(lnp.ref.Append("backend_port"))
}

// FloatingIpEnabled returns a reference to field floating_ip_enabled of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) FloatingIpEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lnp.ref.Append("floating_ip_enabled"))
}

// FrontendIpConfigurationId returns a reference to field frontend_ip_configuration_id of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) FrontendIpConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(lnp.ref.Append("frontend_ip_configuration_id"))
}

// FrontendIpConfigurationName returns a reference to field frontend_ip_configuration_name of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) FrontendIpConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(lnp.ref.Append("frontend_ip_configuration_name"))
}

// FrontendPortEnd returns a reference to field frontend_port_end of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) FrontendPortEnd() terra.NumberValue {
	return terra.ReferenceAsNumber(lnp.ref.Append("frontend_port_end"))
}

// FrontendPortStart returns a reference to field frontend_port_start of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) FrontendPortStart() terra.NumberValue {
	return terra.ReferenceAsNumber(lnp.ref.Append("frontend_port_start"))
}

// Id returns a reference to field id of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lnp.ref.Append("id"))
}

// IdleTimeoutInMinutes returns a reference to field idle_timeout_in_minutes of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) IdleTimeoutInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(lnp.ref.Append("idle_timeout_in_minutes"))
}

// LoadbalancerId returns a reference to field loadbalancer_id of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) LoadbalancerId() terra.StringValue {
	return terra.ReferenceAsString(lnp.ref.Append("loadbalancer_id"))
}

// Name returns a reference to field name of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lnp.ref.Append("name"))
}

// Protocol returns a reference to field protocol of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(lnp.ref.Append("protocol"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lnp.ref.Append("resource_group_name"))
}

// TcpResetEnabled returns a reference to field tcp_reset_enabled of azurerm_lb_nat_pool.
func (lnp lbNatPoolAttributes) TcpResetEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lnp.ref.Append("tcp_reset_enabled"))
}

func (lnp lbNatPoolAttributes) Timeouts() lbnatpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lbnatpool.TimeoutsAttributes](lnp.ref.Append("timeouts"))
}

type lbNatPoolState struct {
	BackendPort                 float64                  `json:"backend_port"`
	FloatingIpEnabled           bool                     `json:"floating_ip_enabled"`
	FrontendIpConfigurationId   string                   `json:"frontend_ip_configuration_id"`
	FrontendIpConfigurationName string                   `json:"frontend_ip_configuration_name"`
	FrontendPortEnd             float64                  `json:"frontend_port_end"`
	FrontendPortStart           float64                  `json:"frontend_port_start"`
	Id                          string                   `json:"id"`
	IdleTimeoutInMinutes        float64                  `json:"idle_timeout_in_minutes"`
	LoadbalancerId              string                   `json:"loadbalancer_id"`
	Name                        string                   `json:"name"`
	Protocol                    string                   `json:"protocol"`
	ResourceGroupName           string                   `json:"resource_group_name"`
	TcpResetEnabled             bool                     `json:"tcp_reset_enabled"`
	Timeouts                    *lbnatpool.TimeoutsState `json:"timeouts"`
}
