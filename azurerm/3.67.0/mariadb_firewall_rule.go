// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mariadbfirewallrule "github.com/golingon/terraproviders/azurerm/3.67.0/mariadbfirewallrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMariadbFirewallRule creates a new instance of [MariadbFirewallRule].
func NewMariadbFirewallRule(name string, args MariadbFirewallRuleArgs) *MariadbFirewallRule {
	return &MariadbFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MariadbFirewallRule)(nil)

// MariadbFirewallRule represents the Terraform resource azurerm_mariadb_firewall_rule.
type MariadbFirewallRule struct {
	Name      string
	Args      MariadbFirewallRuleArgs
	state     *mariadbFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MariadbFirewallRule].
func (mfr *MariadbFirewallRule) Type() string {
	return "azurerm_mariadb_firewall_rule"
}

// LocalName returns the local name for [MariadbFirewallRule].
func (mfr *MariadbFirewallRule) LocalName() string {
	return mfr.Name
}

// Configuration returns the configuration (args) for [MariadbFirewallRule].
func (mfr *MariadbFirewallRule) Configuration() interface{} {
	return mfr.Args
}

// DependOn is used for other resources to depend on [MariadbFirewallRule].
func (mfr *MariadbFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(mfr)
}

// Dependencies returns the list of resources [MariadbFirewallRule] depends_on.
func (mfr *MariadbFirewallRule) Dependencies() terra.Dependencies {
	return mfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MariadbFirewallRule].
func (mfr *MariadbFirewallRule) LifecycleManagement() *terra.Lifecycle {
	return mfr.Lifecycle
}

// Attributes returns the attributes for [MariadbFirewallRule].
func (mfr *MariadbFirewallRule) Attributes() mariadbFirewallRuleAttributes {
	return mariadbFirewallRuleAttributes{ref: terra.ReferenceResource(mfr)}
}

// ImportState imports the given attribute values into [MariadbFirewallRule]'s state.
func (mfr *MariadbFirewallRule) ImportState(av io.Reader) error {
	mfr.state = &mariadbFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(mfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mfr.Type(), mfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MariadbFirewallRule] has state.
func (mfr *MariadbFirewallRule) State() (*mariadbFirewallRuleState, bool) {
	return mfr.state, mfr.state != nil
}

// StateMust returns the state for [MariadbFirewallRule]. Panics if the state is nil.
func (mfr *MariadbFirewallRule) StateMust() *mariadbFirewallRuleState {
	if mfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mfr.Type(), mfr.LocalName()))
	}
	return mfr.state
}

// MariadbFirewallRuleArgs contains the configurations for azurerm_mariadb_firewall_rule.
type MariadbFirewallRuleArgs struct {
	// EndIpAddress: string, required
	EndIpAddress terra.StringValue `hcl:"end_ip_address,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// StartIpAddress: string, required
	StartIpAddress terra.StringValue `hcl:"start_ip_address,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *mariadbfirewallrule.Timeouts `hcl:"timeouts,block"`
}
type mariadbFirewallRuleAttributes struct {
	ref terra.Reference
}

// EndIpAddress returns a reference to field end_ip_address of azurerm_mariadb_firewall_rule.
func (mfr mariadbFirewallRuleAttributes) EndIpAddress() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("end_ip_address"))
}

// Id returns a reference to field id of azurerm_mariadb_firewall_rule.
func (mfr mariadbFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mariadb_firewall_rule.
func (mfr mariadbFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mariadb_firewall_rule.
func (mfr mariadbFirewallRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mariadb_firewall_rule.
func (mfr mariadbFirewallRuleAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("server_name"))
}

// StartIpAddress returns a reference to field start_ip_address of azurerm_mariadb_firewall_rule.
func (mfr mariadbFirewallRuleAttributes) StartIpAddress() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("start_ip_address"))
}

func (mfr mariadbFirewallRuleAttributes) Timeouts() mariadbfirewallrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mariadbfirewallrule.TimeoutsAttributes](mfr.ref.Append("timeouts"))
}

type mariadbFirewallRuleState struct {
	EndIpAddress      string                             `json:"end_ip_address"`
	Id                string                             `json:"id"`
	Name              string                             `json:"name"`
	ResourceGroupName string                             `json:"resource_group_name"`
	ServerName        string                             `json:"server_name"`
	StartIpAddress    string                             `json:"start_ip_address"`
	Timeouts          *mariadbfirewallrule.TimeoutsState `json:"timeouts"`
}
