// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sqlfirewallrule "github.com/golingon/terraproviders/azurerm/3.52.0/sqlfirewallrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlFirewallRule creates a new instance of [SqlFirewallRule].
func NewSqlFirewallRule(name string, args SqlFirewallRuleArgs) *SqlFirewallRule {
	return &SqlFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlFirewallRule)(nil)

// SqlFirewallRule represents the Terraform resource azurerm_sql_firewall_rule.
type SqlFirewallRule struct {
	Name      string
	Args      SqlFirewallRuleArgs
	state     *sqlFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlFirewallRule].
func (sfr *SqlFirewallRule) Type() string {
	return "azurerm_sql_firewall_rule"
}

// LocalName returns the local name for [SqlFirewallRule].
func (sfr *SqlFirewallRule) LocalName() string {
	return sfr.Name
}

// Configuration returns the configuration (args) for [SqlFirewallRule].
func (sfr *SqlFirewallRule) Configuration() interface{} {
	return sfr.Args
}

// DependOn is used for other resources to depend on [SqlFirewallRule].
func (sfr *SqlFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(sfr)
}

// Dependencies returns the list of resources [SqlFirewallRule] depends_on.
func (sfr *SqlFirewallRule) Dependencies() terra.Dependencies {
	return sfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlFirewallRule].
func (sfr *SqlFirewallRule) LifecycleManagement() *terra.Lifecycle {
	return sfr.Lifecycle
}

// Attributes returns the attributes for [SqlFirewallRule].
func (sfr *SqlFirewallRule) Attributes() sqlFirewallRuleAttributes {
	return sqlFirewallRuleAttributes{ref: terra.ReferenceResource(sfr)}
}

// ImportState imports the given attribute values into [SqlFirewallRule]'s state.
func (sfr *SqlFirewallRule) ImportState(av io.Reader) error {
	sfr.state = &sqlFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(sfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sfr.Type(), sfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlFirewallRule] has state.
func (sfr *SqlFirewallRule) State() (*sqlFirewallRuleState, bool) {
	return sfr.state, sfr.state != nil
}

// StateMust returns the state for [SqlFirewallRule]. Panics if the state is nil.
func (sfr *SqlFirewallRule) StateMust() *sqlFirewallRuleState {
	if sfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sfr.Type(), sfr.LocalName()))
	}
	return sfr.state
}

// SqlFirewallRuleArgs contains the configurations for azurerm_sql_firewall_rule.
type SqlFirewallRuleArgs struct {
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
	Timeouts *sqlfirewallrule.Timeouts `hcl:"timeouts,block"`
}
type sqlFirewallRuleAttributes struct {
	ref terra.Reference
}

// EndIpAddress returns a reference to field end_ip_address of azurerm_sql_firewall_rule.
func (sfr sqlFirewallRuleAttributes) EndIpAddress() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("end_ip_address"))
}

// Id returns a reference to field id of azurerm_sql_firewall_rule.
func (sfr sqlFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_sql_firewall_rule.
func (sfr sqlFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_firewall_rule.
func (sfr sqlFirewallRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_sql_firewall_rule.
func (sfr sqlFirewallRuleAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("server_name"))
}

// StartIpAddress returns a reference to field start_ip_address of azurerm_sql_firewall_rule.
func (sfr sqlFirewallRuleAttributes) StartIpAddress() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("start_ip_address"))
}

func (sfr sqlFirewallRuleAttributes) Timeouts() sqlfirewallrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlfirewallrule.TimeoutsAttributes](sfr.ref.Append("timeouts"))
}

type sqlFirewallRuleState struct {
	EndIpAddress      string                         `json:"end_ip_address"`
	Id                string                         `json:"id"`
	Name              string                         `json:"name"`
	ResourceGroupName string                         `json:"resource_group_name"`
	ServerName        string                         `json:"server_name"`
	StartIpAddress    string                         `json:"start_ip_address"`
	Timeouts          *sqlfirewallrule.TimeoutsState `json:"timeouts"`
}
