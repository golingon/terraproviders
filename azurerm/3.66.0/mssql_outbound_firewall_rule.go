// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqloutboundfirewallrule "github.com/golingon/terraproviders/azurerm/3.66.0/mssqloutboundfirewallrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlOutboundFirewallRule creates a new instance of [MssqlOutboundFirewallRule].
func NewMssqlOutboundFirewallRule(name string, args MssqlOutboundFirewallRuleArgs) *MssqlOutboundFirewallRule {
	return &MssqlOutboundFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlOutboundFirewallRule)(nil)

// MssqlOutboundFirewallRule represents the Terraform resource azurerm_mssql_outbound_firewall_rule.
type MssqlOutboundFirewallRule struct {
	Name      string
	Args      MssqlOutboundFirewallRuleArgs
	state     *mssqlOutboundFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlOutboundFirewallRule].
func (mofr *MssqlOutboundFirewallRule) Type() string {
	return "azurerm_mssql_outbound_firewall_rule"
}

// LocalName returns the local name for [MssqlOutboundFirewallRule].
func (mofr *MssqlOutboundFirewallRule) LocalName() string {
	return mofr.Name
}

// Configuration returns the configuration (args) for [MssqlOutboundFirewallRule].
func (mofr *MssqlOutboundFirewallRule) Configuration() interface{} {
	return mofr.Args
}

// DependOn is used for other resources to depend on [MssqlOutboundFirewallRule].
func (mofr *MssqlOutboundFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(mofr)
}

// Dependencies returns the list of resources [MssqlOutboundFirewallRule] depends_on.
func (mofr *MssqlOutboundFirewallRule) Dependencies() terra.Dependencies {
	return mofr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlOutboundFirewallRule].
func (mofr *MssqlOutboundFirewallRule) LifecycleManagement() *terra.Lifecycle {
	return mofr.Lifecycle
}

// Attributes returns the attributes for [MssqlOutboundFirewallRule].
func (mofr *MssqlOutboundFirewallRule) Attributes() mssqlOutboundFirewallRuleAttributes {
	return mssqlOutboundFirewallRuleAttributes{ref: terra.ReferenceResource(mofr)}
}

// ImportState imports the given attribute values into [MssqlOutboundFirewallRule]'s state.
func (mofr *MssqlOutboundFirewallRule) ImportState(av io.Reader) error {
	mofr.state = &mssqlOutboundFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(mofr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mofr.Type(), mofr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlOutboundFirewallRule] has state.
func (mofr *MssqlOutboundFirewallRule) State() (*mssqlOutboundFirewallRuleState, bool) {
	return mofr.state, mofr.state != nil
}

// StateMust returns the state for [MssqlOutboundFirewallRule]. Panics if the state is nil.
func (mofr *MssqlOutboundFirewallRule) StateMust() *mssqlOutboundFirewallRuleState {
	if mofr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mofr.Type(), mofr.LocalName()))
	}
	return mofr.state
}

// MssqlOutboundFirewallRuleArgs contains the configurations for azurerm_mssql_outbound_firewall_rule.
type MssqlOutboundFirewallRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *mssqloutboundfirewallrule.Timeouts `hcl:"timeouts,block"`
}
type mssqlOutboundFirewallRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mssql_outbound_firewall_rule.
func (mofr mssqlOutboundFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mofr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mssql_outbound_firewall_rule.
func (mofr mssqlOutboundFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mofr.ref.Append("name"))
}

// ServerId returns a reference to field server_id of azurerm_mssql_outbound_firewall_rule.
func (mofr mssqlOutboundFirewallRuleAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(mofr.ref.Append("server_id"))
}

func (mofr mssqlOutboundFirewallRuleAttributes) Timeouts() mssqloutboundfirewallrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqloutboundfirewallrule.TimeoutsAttributes](mofr.ref.Append("timeouts"))
}

type mssqlOutboundFirewallRuleState struct {
	Id       string                                   `json:"id"`
	Name     string                                   `json:"name"`
	ServerId string                                   `json:"server_id"`
	Timeouts *mssqloutboundfirewallrule.TimeoutsState `json:"timeouts"`
}
