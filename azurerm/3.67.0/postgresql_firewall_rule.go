// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	postgresqlfirewallrule "github.com/golingon/terraproviders/azurerm/3.67.0/postgresqlfirewallrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPostgresqlFirewallRule creates a new instance of [PostgresqlFirewallRule].
func NewPostgresqlFirewallRule(name string, args PostgresqlFirewallRuleArgs) *PostgresqlFirewallRule {
	return &PostgresqlFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PostgresqlFirewallRule)(nil)

// PostgresqlFirewallRule represents the Terraform resource azurerm_postgresql_firewall_rule.
type PostgresqlFirewallRule struct {
	Name      string
	Args      PostgresqlFirewallRuleArgs
	state     *postgresqlFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PostgresqlFirewallRule].
func (pfr *PostgresqlFirewallRule) Type() string {
	return "azurerm_postgresql_firewall_rule"
}

// LocalName returns the local name for [PostgresqlFirewallRule].
func (pfr *PostgresqlFirewallRule) LocalName() string {
	return pfr.Name
}

// Configuration returns the configuration (args) for [PostgresqlFirewallRule].
func (pfr *PostgresqlFirewallRule) Configuration() interface{} {
	return pfr.Args
}

// DependOn is used for other resources to depend on [PostgresqlFirewallRule].
func (pfr *PostgresqlFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(pfr)
}

// Dependencies returns the list of resources [PostgresqlFirewallRule] depends_on.
func (pfr *PostgresqlFirewallRule) Dependencies() terra.Dependencies {
	return pfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PostgresqlFirewallRule].
func (pfr *PostgresqlFirewallRule) LifecycleManagement() *terra.Lifecycle {
	return pfr.Lifecycle
}

// Attributes returns the attributes for [PostgresqlFirewallRule].
func (pfr *PostgresqlFirewallRule) Attributes() postgresqlFirewallRuleAttributes {
	return postgresqlFirewallRuleAttributes{ref: terra.ReferenceResource(pfr)}
}

// ImportState imports the given attribute values into [PostgresqlFirewallRule]'s state.
func (pfr *PostgresqlFirewallRule) ImportState(av io.Reader) error {
	pfr.state = &postgresqlFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(pfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pfr.Type(), pfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PostgresqlFirewallRule] has state.
func (pfr *PostgresqlFirewallRule) State() (*postgresqlFirewallRuleState, bool) {
	return pfr.state, pfr.state != nil
}

// StateMust returns the state for [PostgresqlFirewallRule]. Panics if the state is nil.
func (pfr *PostgresqlFirewallRule) StateMust() *postgresqlFirewallRuleState {
	if pfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pfr.Type(), pfr.LocalName()))
	}
	return pfr.state
}

// PostgresqlFirewallRuleArgs contains the configurations for azurerm_postgresql_firewall_rule.
type PostgresqlFirewallRuleArgs struct {
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
	Timeouts *postgresqlfirewallrule.Timeouts `hcl:"timeouts,block"`
}
type postgresqlFirewallRuleAttributes struct {
	ref terra.Reference
}

// EndIpAddress returns a reference to field end_ip_address of azurerm_postgresql_firewall_rule.
func (pfr postgresqlFirewallRuleAttributes) EndIpAddress() terra.StringValue {
	return terra.ReferenceAsString(pfr.ref.Append("end_ip_address"))
}

// Id returns a reference to field id of azurerm_postgresql_firewall_rule.
func (pfr postgresqlFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_postgresql_firewall_rule.
func (pfr postgresqlFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pfr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_postgresql_firewall_rule.
func (pfr postgresqlFirewallRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pfr.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_postgresql_firewall_rule.
func (pfr postgresqlFirewallRuleAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(pfr.ref.Append("server_name"))
}

// StartIpAddress returns a reference to field start_ip_address of azurerm_postgresql_firewall_rule.
func (pfr postgresqlFirewallRuleAttributes) StartIpAddress() terra.StringValue {
	return terra.ReferenceAsString(pfr.ref.Append("start_ip_address"))
}

func (pfr postgresqlFirewallRuleAttributes) Timeouts() postgresqlfirewallrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[postgresqlfirewallrule.TimeoutsAttributes](pfr.ref.Append("timeouts"))
}

type postgresqlFirewallRuleState struct {
	EndIpAddress      string                                `json:"end_ip_address"`
	Id                string                                `json:"id"`
	Name              string                                `json:"name"`
	ResourceGroupName string                                `json:"resource_group_name"`
	ServerName        string                                `json:"server_name"`
	StartIpAddress    string                                `json:"start_ip_address"`
	Timeouts          *postgresqlfirewallrule.TimeoutsState `json:"timeouts"`
}
