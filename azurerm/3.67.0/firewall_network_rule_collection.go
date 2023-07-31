// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	firewallnetworkrulecollection "github.com/golingon/terraproviders/azurerm/3.67.0/firewallnetworkrulecollection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirewallNetworkRuleCollection creates a new instance of [FirewallNetworkRuleCollection].
func NewFirewallNetworkRuleCollection(name string, args FirewallNetworkRuleCollectionArgs) *FirewallNetworkRuleCollection {
	return &FirewallNetworkRuleCollection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirewallNetworkRuleCollection)(nil)

// FirewallNetworkRuleCollection represents the Terraform resource azurerm_firewall_network_rule_collection.
type FirewallNetworkRuleCollection struct {
	Name      string
	Args      FirewallNetworkRuleCollectionArgs
	state     *firewallNetworkRuleCollectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirewallNetworkRuleCollection].
func (fnrc *FirewallNetworkRuleCollection) Type() string {
	return "azurerm_firewall_network_rule_collection"
}

// LocalName returns the local name for [FirewallNetworkRuleCollection].
func (fnrc *FirewallNetworkRuleCollection) LocalName() string {
	return fnrc.Name
}

// Configuration returns the configuration (args) for [FirewallNetworkRuleCollection].
func (fnrc *FirewallNetworkRuleCollection) Configuration() interface{} {
	return fnrc.Args
}

// DependOn is used for other resources to depend on [FirewallNetworkRuleCollection].
func (fnrc *FirewallNetworkRuleCollection) DependOn() terra.Reference {
	return terra.ReferenceResource(fnrc)
}

// Dependencies returns the list of resources [FirewallNetworkRuleCollection] depends_on.
func (fnrc *FirewallNetworkRuleCollection) Dependencies() terra.Dependencies {
	return fnrc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirewallNetworkRuleCollection].
func (fnrc *FirewallNetworkRuleCollection) LifecycleManagement() *terra.Lifecycle {
	return fnrc.Lifecycle
}

// Attributes returns the attributes for [FirewallNetworkRuleCollection].
func (fnrc *FirewallNetworkRuleCollection) Attributes() firewallNetworkRuleCollectionAttributes {
	return firewallNetworkRuleCollectionAttributes{ref: terra.ReferenceResource(fnrc)}
}

// ImportState imports the given attribute values into [FirewallNetworkRuleCollection]'s state.
func (fnrc *FirewallNetworkRuleCollection) ImportState(av io.Reader) error {
	fnrc.state = &firewallNetworkRuleCollectionState{}
	if err := json.NewDecoder(av).Decode(fnrc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fnrc.Type(), fnrc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirewallNetworkRuleCollection] has state.
func (fnrc *FirewallNetworkRuleCollection) State() (*firewallNetworkRuleCollectionState, bool) {
	return fnrc.state, fnrc.state != nil
}

// StateMust returns the state for [FirewallNetworkRuleCollection]. Panics if the state is nil.
func (fnrc *FirewallNetworkRuleCollection) StateMust() *firewallNetworkRuleCollectionState {
	if fnrc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fnrc.Type(), fnrc.LocalName()))
	}
	return fnrc.state
}

// FirewallNetworkRuleCollectionArgs contains the configurations for azurerm_firewall_network_rule_collection.
type FirewallNetworkRuleCollectionArgs struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// AzureFirewallName: string, required
	AzureFirewallName terra.StringValue `hcl:"azure_firewall_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Rule: min=1
	Rule []firewallnetworkrulecollection.Rule `hcl:"rule,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *firewallnetworkrulecollection.Timeouts `hcl:"timeouts,block"`
}
type firewallNetworkRuleCollectionAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of azurerm_firewall_network_rule_collection.
func (fnrc firewallNetworkRuleCollectionAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(fnrc.ref.Append("action"))
}

// AzureFirewallName returns a reference to field azure_firewall_name of azurerm_firewall_network_rule_collection.
func (fnrc firewallNetworkRuleCollectionAttributes) AzureFirewallName() terra.StringValue {
	return terra.ReferenceAsString(fnrc.ref.Append("azure_firewall_name"))
}

// Id returns a reference to field id of azurerm_firewall_network_rule_collection.
func (fnrc firewallNetworkRuleCollectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fnrc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_firewall_network_rule_collection.
func (fnrc firewallNetworkRuleCollectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fnrc.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_firewall_network_rule_collection.
func (fnrc firewallNetworkRuleCollectionAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(fnrc.ref.Append("priority"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_firewall_network_rule_collection.
func (fnrc firewallNetworkRuleCollectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(fnrc.ref.Append("resource_group_name"))
}

func (fnrc firewallNetworkRuleCollectionAttributes) Rule() terra.ListValue[firewallnetworkrulecollection.RuleAttributes] {
	return terra.ReferenceAsList[firewallnetworkrulecollection.RuleAttributes](fnrc.ref.Append("rule"))
}

func (fnrc firewallNetworkRuleCollectionAttributes) Timeouts() firewallnetworkrulecollection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firewallnetworkrulecollection.TimeoutsAttributes](fnrc.ref.Append("timeouts"))
}

type firewallNetworkRuleCollectionState struct {
	Action            string                                       `json:"action"`
	AzureFirewallName string                                       `json:"azure_firewall_name"`
	Id                string                                       `json:"id"`
	Name              string                                       `json:"name"`
	Priority          float64                                      `json:"priority"`
	ResourceGroupName string                                       `json:"resource_group_name"`
	Rule              []firewallnetworkrulecollection.RuleState    `json:"rule"`
	Timeouts          *firewallnetworkrulecollection.TimeoutsState `json:"timeouts"`
}
