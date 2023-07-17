// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mysqlvirtualnetworkrule "github.com/golingon/terraproviders/azurerm/3.65.0/mysqlvirtualnetworkrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMysqlVirtualNetworkRule creates a new instance of [MysqlVirtualNetworkRule].
func NewMysqlVirtualNetworkRule(name string, args MysqlVirtualNetworkRuleArgs) *MysqlVirtualNetworkRule {
	return &MysqlVirtualNetworkRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlVirtualNetworkRule)(nil)

// MysqlVirtualNetworkRule represents the Terraform resource azurerm_mysql_virtual_network_rule.
type MysqlVirtualNetworkRule struct {
	Name      string
	Args      MysqlVirtualNetworkRuleArgs
	state     *mysqlVirtualNetworkRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MysqlVirtualNetworkRule].
func (mvnr *MysqlVirtualNetworkRule) Type() string {
	return "azurerm_mysql_virtual_network_rule"
}

// LocalName returns the local name for [MysqlVirtualNetworkRule].
func (mvnr *MysqlVirtualNetworkRule) LocalName() string {
	return mvnr.Name
}

// Configuration returns the configuration (args) for [MysqlVirtualNetworkRule].
func (mvnr *MysqlVirtualNetworkRule) Configuration() interface{} {
	return mvnr.Args
}

// DependOn is used for other resources to depend on [MysqlVirtualNetworkRule].
func (mvnr *MysqlVirtualNetworkRule) DependOn() terra.Reference {
	return terra.ReferenceResource(mvnr)
}

// Dependencies returns the list of resources [MysqlVirtualNetworkRule] depends_on.
func (mvnr *MysqlVirtualNetworkRule) Dependencies() terra.Dependencies {
	return mvnr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MysqlVirtualNetworkRule].
func (mvnr *MysqlVirtualNetworkRule) LifecycleManagement() *terra.Lifecycle {
	return mvnr.Lifecycle
}

// Attributes returns the attributes for [MysqlVirtualNetworkRule].
func (mvnr *MysqlVirtualNetworkRule) Attributes() mysqlVirtualNetworkRuleAttributes {
	return mysqlVirtualNetworkRuleAttributes{ref: terra.ReferenceResource(mvnr)}
}

// ImportState imports the given attribute values into [MysqlVirtualNetworkRule]'s state.
func (mvnr *MysqlVirtualNetworkRule) ImportState(av io.Reader) error {
	mvnr.state = &mysqlVirtualNetworkRuleState{}
	if err := json.NewDecoder(av).Decode(mvnr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mvnr.Type(), mvnr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MysqlVirtualNetworkRule] has state.
func (mvnr *MysqlVirtualNetworkRule) State() (*mysqlVirtualNetworkRuleState, bool) {
	return mvnr.state, mvnr.state != nil
}

// StateMust returns the state for [MysqlVirtualNetworkRule]. Panics if the state is nil.
func (mvnr *MysqlVirtualNetworkRule) StateMust() *mysqlVirtualNetworkRuleState {
	if mvnr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mvnr.Type(), mvnr.LocalName()))
	}
	return mvnr.state
}

// MysqlVirtualNetworkRuleArgs contains the configurations for azurerm_mysql_virtual_network_rule.
type MysqlVirtualNetworkRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *mysqlvirtualnetworkrule.Timeouts `hcl:"timeouts,block"`
}
type mysqlVirtualNetworkRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mysql_virtual_network_rule.
func (mvnr mysqlVirtualNetworkRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mvnr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mysql_virtual_network_rule.
func (mvnr mysqlVirtualNetworkRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mvnr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_virtual_network_rule.
func (mvnr mysqlVirtualNetworkRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mvnr.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mysql_virtual_network_rule.
func (mvnr mysqlVirtualNetworkRuleAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(mvnr.ref.Append("server_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_mysql_virtual_network_rule.
func (mvnr mysqlVirtualNetworkRuleAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(mvnr.ref.Append("subnet_id"))
}

func (mvnr mysqlVirtualNetworkRuleAttributes) Timeouts() mysqlvirtualnetworkrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mysqlvirtualnetworkrule.TimeoutsAttributes](mvnr.ref.Append("timeouts"))
}

type mysqlVirtualNetworkRuleState struct {
	Id                string                                 `json:"id"`
	Name              string                                 `json:"name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	ServerName        string                                 `json:"server_name"`
	SubnetId          string                                 `json:"subnet_id"`
	Timeouts          *mysqlvirtualnetworkrule.TimeoutsState `json:"timeouts"`
}
