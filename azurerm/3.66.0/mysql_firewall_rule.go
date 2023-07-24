// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mysqlfirewallrule "github.com/golingon/terraproviders/azurerm/3.66.0/mysqlfirewallrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMysqlFirewallRule creates a new instance of [MysqlFirewallRule].
func NewMysqlFirewallRule(name string, args MysqlFirewallRuleArgs) *MysqlFirewallRule {
	return &MysqlFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlFirewallRule)(nil)

// MysqlFirewallRule represents the Terraform resource azurerm_mysql_firewall_rule.
type MysqlFirewallRule struct {
	Name      string
	Args      MysqlFirewallRuleArgs
	state     *mysqlFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MysqlFirewallRule].
func (mfr *MysqlFirewallRule) Type() string {
	return "azurerm_mysql_firewall_rule"
}

// LocalName returns the local name for [MysqlFirewallRule].
func (mfr *MysqlFirewallRule) LocalName() string {
	return mfr.Name
}

// Configuration returns the configuration (args) for [MysqlFirewallRule].
func (mfr *MysqlFirewallRule) Configuration() interface{} {
	return mfr.Args
}

// DependOn is used for other resources to depend on [MysqlFirewallRule].
func (mfr *MysqlFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(mfr)
}

// Dependencies returns the list of resources [MysqlFirewallRule] depends_on.
func (mfr *MysqlFirewallRule) Dependencies() terra.Dependencies {
	return mfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MysqlFirewallRule].
func (mfr *MysqlFirewallRule) LifecycleManagement() *terra.Lifecycle {
	return mfr.Lifecycle
}

// Attributes returns the attributes for [MysqlFirewallRule].
func (mfr *MysqlFirewallRule) Attributes() mysqlFirewallRuleAttributes {
	return mysqlFirewallRuleAttributes{ref: terra.ReferenceResource(mfr)}
}

// ImportState imports the given attribute values into [MysqlFirewallRule]'s state.
func (mfr *MysqlFirewallRule) ImportState(av io.Reader) error {
	mfr.state = &mysqlFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(mfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mfr.Type(), mfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MysqlFirewallRule] has state.
func (mfr *MysqlFirewallRule) State() (*mysqlFirewallRuleState, bool) {
	return mfr.state, mfr.state != nil
}

// StateMust returns the state for [MysqlFirewallRule]. Panics if the state is nil.
func (mfr *MysqlFirewallRule) StateMust() *mysqlFirewallRuleState {
	if mfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mfr.Type(), mfr.LocalName()))
	}
	return mfr.state
}

// MysqlFirewallRuleArgs contains the configurations for azurerm_mysql_firewall_rule.
type MysqlFirewallRuleArgs struct {
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
	Timeouts *mysqlfirewallrule.Timeouts `hcl:"timeouts,block"`
}
type mysqlFirewallRuleAttributes struct {
	ref terra.Reference
}

// EndIpAddress returns a reference to field end_ip_address of azurerm_mysql_firewall_rule.
func (mfr mysqlFirewallRuleAttributes) EndIpAddress() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("end_ip_address"))
}

// Id returns a reference to field id of azurerm_mysql_firewall_rule.
func (mfr mysqlFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mysql_firewall_rule.
func (mfr mysqlFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_firewall_rule.
func (mfr mysqlFirewallRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mysql_firewall_rule.
func (mfr mysqlFirewallRuleAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("server_name"))
}

// StartIpAddress returns a reference to field start_ip_address of azurerm_mysql_firewall_rule.
func (mfr mysqlFirewallRuleAttributes) StartIpAddress() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("start_ip_address"))
}

func (mfr mysqlFirewallRuleAttributes) Timeouts() mysqlfirewallrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mysqlfirewallrule.TimeoutsAttributes](mfr.ref.Append("timeouts"))
}

type mysqlFirewallRuleState struct {
	EndIpAddress      string                           `json:"end_ip_address"`
	Id                string                           `json:"id"`
	Name              string                           `json:"name"`
	ResourceGroupName string                           `json:"resource_group_name"`
	ServerName        string                           `json:"server_name"`
	StartIpAddress    string                           `json:"start_ip_address"`
	Timeouts          *mysqlfirewallrule.TimeoutsState `json:"timeouts"`
}
