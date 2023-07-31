// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mariadbvirtualnetworkrule "github.com/golingon/terraproviders/azurerm/3.67.0/mariadbvirtualnetworkrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMariadbVirtualNetworkRule creates a new instance of [MariadbVirtualNetworkRule].
func NewMariadbVirtualNetworkRule(name string, args MariadbVirtualNetworkRuleArgs) *MariadbVirtualNetworkRule {
	return &MariadbVirtualNetworkRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MariadbVirtualNetworkRule)(nil)

// MariadbVirtualNetworkRule represents the Terraform resource azurerm_mariadb_virtual_network_rule.
type MariadbVirtualNetworkRule struct {
	Name      string
	Args      MariadbVirtualNetworkRuleArgs
	state     *mariadbVirtualNetworkRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MariadbVirtualNetworkRule].
func (mvnr *MariadbVirtualNetworkRule) Type() string {
	return "azurerm_mariadb_virtual_network_rule"
}

// LocalName returns the local name for [MariadbVirtualNetworkRule].
func (mvnr *MariadbVirtualNetworkRule) LocalName() string {
	return mvnr.Name
}

// Configuration returns the configuration (args) for [MariadbVirtualNetworkRule].
func (mvnr *MariadbVirtualNetworkRule) Configuration() interface{} {
	return mvnr.Args
}

// DependOn is used for other resources to depend on [MariadbVirtualNetworkRule].
func (mvnr *MariadbVirtualNetworkRule) DependOn() terra.Reference {
	return terra.ReferenceResource(mvnr)
}

// Dependencies returns the list of resources [MariadbVirtualNetworkRule] depends_on.
func (mvnr *MariadbVirtualNetworkRule) Dependencies() terra.Dependencies {
	return mvnr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MariadbVirtualNetworkRule].
func (mvnr *MariadbVirtualNetworkRule) LifecycleManagement() *terra.Lifecycle {
	return mvnr.Lifecycle
}

// Attributes returns the attributes for [MariadbVirtualNetworkRule].
func (mvnr *MariadbVirtualNetworkRule) Attributes() mariadbVirtualNetworkRuleAttributes {
	return mariadbVirtualNetworkRuleAttributes{ref: terra.ReferenceResource(mvnr)}
}

// ImportState imports the given attribute values into [MariadbVirtualNetworkRule]'s state.
func (mvnr *MariadbVirtualNetworkRule) ImportState(av io.Reader) error {
	mvnr.state = &mariadbVirtualNetworkRuleState{}
	if err := json.NewDecoder(av).Decode(mvnr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mvnr.Type(), mvnr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MariadbVirtualNetworkRule] has state.
func (mvnr *MariadbVirtualNetworkRule) State() (*mariadbVirtualNetworkRuleState, bool) {
	return mvnr.state, mvnr.state != nil
}

// StateMust returns the state for [MariadbVirtualNetworkRule]. Panics if the state is nil.
func (mvnr *MariadbVirtualNetworkRule) StateMust() *mariadbVirtualNetworkRuleState {
	if mvnr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mvnr.Type(), mvnr.LocalName()))
	}
	return mvnr.state
}

// MariadbVirtualNetworkRuleArgs contains the configurations for azurerm_mariadb_virtual_network_rule.
type MariadbVirtualNetworkRuleArgs struct {
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
	Timeouts *mariadbvirtualnetworkrule.Timeouts `hcl:"timeouts,block"`
}
type mariadbVirtualNetworkRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mariadb_virtual_network_rule.
func (mvnr mariadbVirtualNetworkRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mvnr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mariadb_virtual_network_rule.
func (mvnr mariadbVirtualNetworkRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mvnr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mariadb_virtual_network_rule.
func (mvnr mariadbVirtualNetworkRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mvnr.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mariadb_virtual_network_rule.
func (mvnr mariadbVirtualNetworkRuleAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(mvnr.ref.Append("server_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_mariadb_virtual_network_rule.
func (mvnr mariadbVirtualNetworkRuleAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(mvnr.ref.Append("subnet_id"))
}

func (mvnr mariadbVirtualNetworkRuleAttributes) Timeouts() mariadbvirtualnetworkrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mariadbvirtualnetworkrule.TimeoutsAttributes](mvnr.ref.Append("timeouts"))
}

type mariadbVirtualNetworkRuleState struct {
	Id                string                                   `json:"id"`
	Name              string                                   `json:"name"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	ServerName        string                                   `json:"server_name"`
	SubnetId          string                                   `json:"subnet_id"`
	Timeouts          *mariadbvirtualnetworkrule.TimeoutsState `json:"timeouts"`
}
