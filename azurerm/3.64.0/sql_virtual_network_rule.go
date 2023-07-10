// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sqlvirtualnetworkrule "github.com/golingon/terraproviders/azurerm/3.64.0/sqlvirtualnetworkrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlVirtualNetworkRule creates a new instance of [SqlVirtualNetworkRule].
func NewSqlVirtualNetworkRule(name string, args SqlVirtualNetworkRuleArgs) *SqlVirtualNetworkRule {
	return &SqlVirtualNetworkRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlVirtualNetworkRule)(nil)

// SqlVirtualNetworkRule represents the Terraform resource azurerm_sql_virtual_network_rule.
type SqlVirtualNetworkRule struct {
	Name      string
	Args      SqlVirtualNetworkRuleArgs
	state     *sqlVirtualNetworkRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlVirtualNetworkRule].
func (svnr *SqlVirtualNetworkRule) Type() string {
	return "azurerm_sql_virtual_network_rule"
}

// LocalName returns the local name for [SqlVirtualNetworkRule].
func (svnr *SqlVirtualNetworkRule) LocalName() string {
	return svnr.Name
}

// Configuration returns the configuration (args) for [SqlVirtualNetworkRule].
func (svnr *SqlVirtualNetworkRule) Configuration() interface{} {
	return svnr.Args
}

// DependOn is used for other resources to depend on [SqlVirtualNetworkRule].
func (svnr *SqlVirtualNetworkRule) DependOn() terra.Reference {
	return terra.ReferenceResource(svnr)
}

// Dependencies returns the list of resources [SqlVirtualNetworkRule] depends_on.
func (svnr *SqlVirtualNetworkRule) Dependencies() terra.Dependencies {
	return svnr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlVirtualNetworkRule].
func (svnr *SqlVirtualNetworkRule) LifecycleManagement() *terra.Lifecycle {
	return svnr.Lifecycle
}

// Attributes returns the attributes for [SqlVirtualNetworkRule].
func (svnr *SqlVirtualNetworkRule) Attributes() sqlVirtualNetworkRuleAttributes {
	return sqlVirtualNetworkRuleAttributes{ref: terra.ReferenceResource(svnr)}
}

// ImportState imports the given attribute values into [SqlVirtualNetworkRule]'s state.
func (svnr *SqlVirtualNetworkRule) ImportState(av io.Reader) error {
	svnr.state = &sqlVirtualNetworkRuleState{}
	if err := json.NewDecoder(av).Decode(svnr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", svnr.Type(), svnr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlVirtualNetworkRule] has state.
func (svnr *SqlVirtualNetworkRule) State() (*sqlVirtualNetworkRuleState, bool) {
	return svnr.state, svnr.state != nil
}

// StateMust returns the state for [SqlVirtualNetworkRule]. Panics if the state is nil.
func (svnr *SqlVirtualNetworkRule) StateMust() *sqlVirtualNetworkRuleState {
	if svnr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", svnr.Type(), svnr.LocalName()))
	}
	return svnr.state
}

// SqlVirtualNetworkRuleArgs contains the configurations for azurerm_sql_virtual_network_rule.
type SqlVirtualNetworkRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreMissingVnetServiceEndpoint: bool, optional
	IgnoreMissingVnetServiceEndpoint terra.BoolValue `hcl:"ignore_missing_vnet_service_endpoint,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *sqlvirtualnetworkrule.Timeouts `hcl:"timeouts,block"`
}
type sqlVirtualNetworkRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sql_virtual_network_rule.
func (svnr sqlVirtualNetworkRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(svnr.ref.Append("id"))
}

// IgnoreMissingVnetServiceEndpoint returns a reference to field ignore_missing_vnet_service_endpoint of azurerm_sql_virtual_network_rule.
func (svnr sqlVirtualNetworkRuleAttributes) IgnoreMissingVnetServiceEndpoint() terra.BoolValue {
	return terra.ReferenceAsBool(svnr.ref.Append("ignore_missing_vnet_service_endpoint"))
}

// Name returns a reference to field name of azurerm_sql_virtual_network_rule.
func (svnr sqlVirtualNetworkRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(svnr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_virtual_network_rule.
func (svnr sqlVirtualNetworkRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(svnr.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_sql_virtual_network_rule.
func (svnr sqlVirtualNetworkRuleAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(svnr.ref.Append("server_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_sql_virtual_network_rule.
func (svnr sqlVirtualNetworkRuleAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(svnr.ref.Append("subnet_id"))
}

func (svnr sqlVirtualNetworkRuleAttributes) Timeouts() sqlvirtualnetworkrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlvirtualnetworkrule.TimeoutsAttributes](svnr.ref.Append("timeouts"))
}

type sqlVirtualNetworkRuleState struct {
	Id                               string                               `json:"id"`
	IgnoreMissingVnetServiceEndpoint bool                                 `json:"ignore_missing_vnet_service_endpoint"`
	Name                             string                               `json:"name"`
	ResourceGroupName                string                               `json:"resource_group_name"`
	ServerName                       string                               `json:"server_name"`
	SubnetId                         string                               `json:"subnet_id"`
	Timeouts                         *sqlvirtualnetworkrule.TimeoutsState `json:"timeouts"`
}
