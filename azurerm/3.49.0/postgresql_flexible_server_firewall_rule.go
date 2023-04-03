// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	postgresqlflexibleserverfirewallrule "github.com/golingon/terraproviders/azurerm/3.49.0/postgresqlflexibleserverfirewallrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPostgresqlFlexibleServerFirewallRule creates a new instance of [PostgresqlFlexibleServerFirewallRule].
func NewPostgresqlFlexibleServerFirewallRule(name string, args PostgresqlFlexibleServerFirewallRuleArgs) *PostgresqlFlexibleServerFirewallRule {
	return &PostgresqlFlexibleServerFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PostgresqlFlexibleServerFirewallRule)(nil)

// PostgresqlFlexibleServerFirewallRule represents the Terraform resource azurerm_postgresql_flexible_server_firewall_rule.
type PostgresqlFlexibleServerFirewallRule struct {
	Name      string
	Args      PostgresqlFlexibleServerFirewallRuleArgs
	state     *postgresqlFlexibleServerFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PostgresqlFlexibleServerFirewallRule].
func (pfsfr *PostgresqlFlexibleServerFirewallRule) Type() string {
	return "azurerm_postgresql_flexible_server_firewall_rule"
}

// LocalName returns the local name for [PostgresqlFlexibleServerFirewallRule].
func (pfsfr *PostgresqlFlexibleServerFirewallRule) LocalName() string {
	return pfsfr.Name
}

// Configuration returns the configuration (args) for [PostgresqlFlexibleServerFirewallRule].
func (pfsfr *PostgresqlFlexibleServerFirewallRule) Configuration() interface{} {
	return pfsfr.Args
}

// DependOn is used for other resources to depend on [PostgresqlFlexibleServerFirewallRule].
func (pfsfr *PostgresqlFlexibleServerFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(pfsfr)
}

// Dependencies returns the list of resources [PostgresqlFlexibleServerFirewallRule] depends_on.
func (pfsfr *PostgresqlFlexibleServerFirewallRule) Dependencies() terra.Dependencies {
	return pfsfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PostgresqlFlexibleServerFirewallRule].
func (pfsfr *PostgresqlFlexibleServerFirewallRule) LifecycleManagement() *terra.Lifecycle {
	return pfsfr.Lifecycle
}

// Attributes returns the attributes for [PostgresqlFlexibleServerFirewallRule].
func (pfsfr *PostgresqlFlexibleServerFirewallRule) Attributes() postgresqlFlexibleServerFirewallRuleAttributes {
	return postgresqlFlexibleServerFirewallRuleAttributes{ref: terra.ReferenceResource(pfsfr)}
}

// ImportState imports the given attribute values into [PostgresqlFlexibleServerFirewallRule]'s state.
func (pfsfr *PostgresqlFlexibleServerFirewallRule) ImportState(av io.Reader) error {
	pfsfr.state = &postgresqlFlexibleServerFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(pfsfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pfsfr.Type(), pfsfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PostgresqlFlexibleServerFirewallRule] has state.
func (pfsfr *PostgresqlFlexibleServerFirewallRule) State() (*postgresqlFlexibleServerFirewallRuleState, bool) {
	return pfsfr.state, pfsfr.state != nil
}

// StateMust returns the state for [PostgresqlFlexibleServerFirewallRule]. Panics if the state is nil.
func (pfsfr *PostgresqlFlexibleServerFirewallRule) StateMust() *postgresqlFlexibleServerFirewallRuleState {
	if pfsfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pfsfr.Type(), pfsfr.LocalName()))
	}
	return pfsfr.state
}

// PostgresqlFlexibleServerFirewallRuleArgs contains the configurations for azurerm_postgresql_flexible_server_firewall_rule.
type PostgresqlFlexibleServerFirewallRuleArgs struct {
	// EndIpAddress: string, required
	EndIpAddress terra.StringValue `hcl:"end_ip_address,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// StartIpAddress: string, required
	StartIpAddress terra.StringValue `hcl:"start_ip_address,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *postgresqlflexibleserverfirewallrule.Timeouts `hcl:"timeouts,block"`
}
type postgresqlFlexibleServerFirewallRuleAttributes struct {
	ref terra.Reference
}

// EndIpAddress returns a reference to field end_ip_address of azurerm_postgresql_flexible_server_firewall_rule.
func (pfsfr postgresqlFlexibleServerFirewallRuleAttributes) EndIpAddress() terra.StringValue {
	return terra.ReferenceAsString(pfsfr.ref.Append("end_ip_address"))
}

// Id returns a reference to field id of azurerm_postgresql_flexible_server_firewall_rule.
func (pfsfr postgresqlFlexibleServerFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pfsfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_postgresql_flexible_server_firewall_rule.
func (pfsfr postgresqlFlexibleServerFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pfsfr.ref.Append("name"))
}

// ServerId returns a reference to field server_id of azurerm_postgresql_flexible_server_firewall_rule.
func (pfsfr postgresqlFlexibleServerFirewallRuleAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(pfsfr.ref.Append("server_id"))
}

// StartIpAddress returns a reference to field start_ip_address of azurerm_postgresql_flexible_server_firewall_rule.
func (pfsfr postgresqlFlexibleServerFirewallRuleAttributes) StartIpAddress() terra.StringValue {
	return terra.ReferenceAsString(pfsfr.ref.Append("start_ip_address"))
}

func (pfsfr postgresqlFlexibleServerFirewallRuleAttributes) Timeouts() postgresqlflexibleserverfirewallrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[postgresqlflexibleserverfirewallrule.TimeoutsAttributes](pfsfr.ref.Append("timeouts"))
}

type postgresqlFlexibleServerFirewallRuleState struct {
	EndIpAddress   string                                              `json:"end_ip_address"`
	Id             string                                              `json:"id"`
	Name           string                                              `json:"name"`
	ServerId       string                                              `json:"server_id"`
	StartIpAddress string                                              `json:"start_ip_address"`
	Timeouts       *postgresqlflexibleserverfirewallrule.TimeoutsState `json:"timeouts"`
}
