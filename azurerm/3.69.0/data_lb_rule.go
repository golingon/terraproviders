// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datalbrule "github.com/golingon/terraproviders/azurerm/3.69.0/datalbrule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLbRule creates a new instance of [DataLbRule].
func NewDataLbRule(name string, args DataLbRuleArgs) *DataLbRule {
	return &DataLbRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLbRule)(nil)

// DataLbRule represents the Terraform data resource azurerm_lb_rule.
type DataLbRule struct {
	Name string
	Args DataLbRuleArgs
}

// DataSource returns the Terraform object type for [DataLbRule].
func (lr *DataLbRule) DataSource() string {
	return "azurerm_lb_rule"
}

// LocalName returns the local name for [DataLbRule].
func (lr *DataLbRule) LocalName() string {
	return lr.Name
}

// Configuration returns the configuration (args) for [DataLbRule].
func (lr *DataLbRule) Configuration() interface{} {
	return lr.Args
}

// Attributes returns the attributes for [DataLbRule].
func (lr *DataLbRule) Attributes() dataLbRuleAttributes {
	return dataLbRuleAttributes{ref: terra.ReferenceDataResource(lr)}
}

// DataLbRuleArgs contains the configurations for azurerm_lb_rule.
type DataLbRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoadbalancerId: string, required
	LoadbalancerId terra.StringValue `hcl:"loadbalancer_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datalbrule.Timeouts `hcl:"timeouts,block"`
}
type dataLbRuleAttributes struct {
	ref terra.Reference
}

// BackendAddressPoolId returns a reference to field backend_address_pool_id of azurerm_lb_rule.
func (lr dataLbRuleAttributes) BackendAddressPoolId() terra.StringValue {
	return terra.ReferenceAsString(lr.ref.Append("backend_address_pool_id"))
}

// BackendPort returns a reference to field backend_port of azurerm_lb_rule.
func (lr dataLbRuleAttributes) BackendPort() terra.NumberValue {
	return terra.ReferenceAsNumber(lr.ref.Append("backend_port"))
}

// DisableOutboundSnat returns a reference to field disable_outbound_snat of azurerm_lb_rule.
func (lr dataLbRuleAttributes) DisableOutboundSnat() terra.BoolValue {
	return terra.ReferenceAsBool(lr.ref.Append("disable_outbound_snat"))
}

// EnableFloatingIp returns a reference to field enable_floating_ip of azurerm_lb_rule.
func (lr dataLbRuleAttributes) EnableFloatingIp() terra.BoolValue {
	return terra.ReferenceAsBool(lr.ref.Append("enable_floating_ip"))
}

// EnableTcpReset returns a reference to field enable_tcp_reset of azurerm_lb_rule.
func (lr dataLbRuleAttributes) EnableTcpReset() terra.BoolValue {
	return terra.ReferenceAsBool(lr.ref.Append("enable_tcp_reset"))
}

// FrontendIpConfigurationName returns a reference to field frontend_ip_configuration_name of azurerm_lb_rule.
func (lr dataLbRuleAttributes) FrontendIpConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(lr.ref.Append("frontend_ip_configuration_name"))
}

// FrontendPort returns a reference to field frontend_port of azurerm_lb_rule.
func (lr dataLbRuleAttributes) FrontendPort() terra.NumberValue {
	return terra.ReferenceAsNumber(lr.ref.Append("frontend_port"))
}

// Id returns a reference to field id of azurerm_lb_rule.
func (lr dataLbRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lr.ref.Append("id"))
}

// IdleTimeoutInMinutes returns a reference to field idle_timeout_in_minutes of azurerm_lb_rule.
func (lr dataLbRuleAttributes) IdleTimeoutInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(lr.ref.Append("idle_timeout_in_minutes"))
}

// LoadDistribution returns a reference to field load_distribution of azurerm_lb_rule.
func (lr dataLbRuleAttributes) LoadDistribution() terra.StringValue {
	return terra.ReferenceAsString(lr.ref.Append("load_distribution"))
}

// LoadbalancerId returns a reference to field loadbalancer_id of azurerm_lb_rule.
func (lr dataLbRuleAttributes) LoadbalancerId() terra.StringValue {
	return terra.ReferenceAsString(lr.ref.Append("loadbalancer_id"))
}

// Name returns a reference to field name of azurerm_lb_rule.
func (lr dataLbRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lr.ref.Append("name"))
}

// ProbeId returns a reference to field probe_id of azurerm_lb_rule.
func (lr dataLbRuleAttributes) ProbeId() terra.StringValue {
	return terra.ReferenceAsString(lr.ref.Append("probe_id"))
}

// Protocol returns a reference to field protocol of azurerm_lb_rule.
func (lr dataLbRuleAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(lr.ref.Append("protocol"))
}

func (lr dataLbRuleAttributes) Timeouts() datalbrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalbrule.TimeoutsAttributes](lr.ref.Append("timeouts"))
}
