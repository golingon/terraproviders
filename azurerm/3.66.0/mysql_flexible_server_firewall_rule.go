// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mysqlflexibleserverfirewallrule "github.com/golingon/terraproviders/azurerm/3.66.0/mysqlflexibleserverfirewallrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMysqlFlexibleServerFirewallRule creates a new instance of [MysqlFlexibleServerFirewallRule].
func NewMysqlFlexibleServerFirewallRule(name string, args MysqlFlexibleServerFirewallRuleArgs) *MysqlFlexibleServerFirewallRule {
	return &MysqlFlexibleServerFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlFlexibleServerFirewallRule)(nil)

// MysqlFlexibleServerFirewallRule represents the Terraform resource azurerm_mysql_flexible_server_firewall_rule.
type MysqlFlexibleServerFirewallRule struct {
	Name      string
	Args      MysqlFlexibleServerFirewallRuleArgs
	state     *mysqlFlexibleServerFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MysqlFlexibleServerFirewallRule].
func (mfsfr *MysqlFlexibleServerFirewallRule) Type() string {
	return "azurerm_mysql_flexible_server_firewall_rule"
}

// LocalName returns the local name for [MysqlFlexibleServerFirewallRule].
func (mfsfr *MysqlFlexibleServerFirewallRule) LocalName() string {
	return mfsfr.Name
}

// Configuration returns the configuration (args) for [MysqlFlexibleServerFirewallRule].
func (mfsfr *MysqlFlexibleServerFirewallRule) Configuration() interface{} {
	return mfsfr.Args
}

// DependOn is used for other resources to depend on [MysqlFlexibleServerFirewallRule].
func (mfsfr *MysqlFlexibleServerFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(mfsfr)
}

// Dependencies returns the list of resources [MysqlFlexibleServerFirewallRule] depends_on.
func (mfsfr *MysqlFlexibleServerFirewallRule) Dependencies() terra.Dependencies {
	return mfsfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MysqlFlexibleServerFirewallRule].
func (mfsfr *MysqlFlexibleServerFirewallRule) LifecycleManagement() *terra.Lifecycle {
	return mfsfr.Lifecycle
}

// Attributes returns the attributes for [MysqlFlexibleServerFirewallRule].
func (mfsfr *MysqlFlexibleServerFirewallRule) Attributes() mysqlFlexibleServerFirewallRuleAttributes {
	return mysqlFlexibleServerFirewallRuleAttributes{ref: terra.ReferenceResource(mfsfr)}
}

// ImportState imports the given attribute values into [MysqlFlexibleServerFirewallRule]'s state.
func (mfsfr *MysqlFlexibleServerFirewallRule) ImportState(av io.Reader) error {
	mfsfr.state = &mysqlFlexibleServerFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(mfsfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mfsfr.Type(), mfsfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MysqlFlexibleServerFirewallRule] has state.
func (mfsfr *MysqlFlexibleServerFirewallRule) State() (*mysqlFlexibleServerFirewallRuleState, bool) {
	return mfsfr.state, mfsfr.state != nil
}

// StateMust returns the state for [MysqlFlexibleServerFirewallRule]. Panics if the state is nil.
func (mfsfr *MysqlFlexibleServerFirewallRule) StateMust() *mysqlFlexibleServerFirewallRuleState {
	if mfsfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mfsfr.Type(), mfsfr.LocalName()))
	}
	return mfsfr.state
}

// MysqlFlexibleServerFirewallRuleArgs contains the configurations for azurerm_mysql_flexible_server_firewall_rule.
type MysqlFlexibleServerFirewallRuleArgs struct {
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
	Timeouts *mysqlflexibleserverfirewallrule.Timeouts `hcl:"timeouts,block"`
}
type mysqlFlexibleServerFirewallRuleAttributes struct {
	ref terra.Reference
}

// EndIpAddress returns a reference to field end_ip_address of azurerm_mysql_flexible_server_firewall_rule.
func (mfsfr mysqlFlexibleServerFirewallRuleAttributes) EndIpAddress() terra.StringValue {
	return terra.ReferenceAsString(mfsfr.ref.Append("end_ip_address"))
}

// Id returns a reference to field id of azurerm_mysql_flexible_server_firewall_rule.
func (mfsfr mysqlFlexibleServerFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mfsfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mysql_flexible_server_firewall_rule.
func (mfsfr mysqlFlexibleServerFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mfsfr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_flexible_server_firewall_rule.
func (mfsfr mysqlFlexibleServerFirewallRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mfsfr.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mysql_flexible_server_firewall_rule.
func (mfsfr mysqlFlexibleServerFirewallRuleAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(mfsfr.ref.Append("server_name"))
}

// StartIpAddress returns a reference to field start_ip_address of azurerm_mysql_flexible_server_firewall_rule.
func (mfsfr mysqlFlexibleServerFirewallRuleAttributes) StartIpAddress() terra.StringValue {
	return terra.ReferenceAsString(mfsfr.ref.Append("start_ip_address"))
}

func (mfsfr mysqlFlexibleServerFirewallRuleAttributes) Timeouts() mysqlflexibleserverfirewallrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mysqlflexibleserverfirewallrule.TimeoutsAttributes](mfsfr.ref.Append("timeouts"))
}

type mysqlFlexibleServerFirewallRuleState struct {
	EndIpAddress      string                                         `json:"end_ip_address"`
	Id                string                                         `json:"id"`
	Name              string                                         `json:"name"`
	ResourceGroupName string                                         `json:"resource_group_name"`
	ServerName        string                                         `json:"server_name"`
	StartIpAddress    string                                         `json:"start_ip_address"`
	Timeouts          *mysqlflexibleserverfirewallrule.TimeoutsState `json:"timeouts"`
}
