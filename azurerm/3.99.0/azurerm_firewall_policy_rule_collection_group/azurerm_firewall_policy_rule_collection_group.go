// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_firewall_policy_rule_collection_group

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_firewall_policy_rule_collection_group.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermFirewallPolicyRuleCollectionGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (afprcg *Resource) Type() string {
	return "azurerm_firewall_policy_rule_collection_group"
}

// LocalName returns the local name for [Resource].
func (afprcg *Resource) LocalName() string {
	return afprcg.Name
}

// Configuration returns the configuration (args) for [Resource].
func (afprcg *Resource) Configuration() interface{} {
	return afprcg.Args
}

// DependOn is used for other resources to depend on [Resource].
func (afprcg *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(afprcg)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (afprcg *Resource) Dependencies() terra.Dependencies {
	return afprcg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (afprcg *Resource) LifecycleManagement() *terra.Lifecycle {
	return afprcg.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (afprcg *Resource) Attributes() azurermFirewallPolicyRuleCollectionGroupAttributes {
	return azurermFirewallPolicyRuleCollectionGroupAttributes{ref: terra.ReferenceResource(afprcg)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (afprcg *Resource) ImportState(state io.Reader) error {
	afprcg.state = &azurermFirewallPolicyRuleCollectionGroupState{}
	if err := json.NewDecoder(state).Decode(afprcg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", afprcg.Type(), afprcg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (afprcg *Resource) State() (*azurermFirewallPolicyRuleCollectionGroupState, bool) {
	return afprcg.state, afprcg.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (afprcg *Resource) StateMust() *azurermFirewallPolicyRuleCollectionGroupState {
	if afprcg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", afprcg.Type(), afprcg.LocalName()))
	}
	return afprcg.state
}

// Args contains the configurations for azurerm_firewall_policy_rule_collection_group.
type Args struct {
	// FirewallPolicyId: string, required
	FirewallPolicyId terra.StringValue `hcl:"firewall_policy_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// ApplicationRuleCollection: min=0
	ApplicationRuleCollection []ApplicationRuleCollection `hcl:"application_rule_collection,block" validate:"min=0"`
	// NatRuleCollection: min=0
	NatRuleCollection []NatRuleCollection `hcl:"nat_rule_collection,block" validate:"min=0"`
	// NetworkRuleCollection: min=0
	NetworkRuleCollection []NetworkRuleCollection `hcl:"network_rule_collection,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermFirewallPolicyRuleCollectionGroupAttributes struct {
	ref terra.Reference
}

// FirewallPolicyId returns a reference to field firewall_policy_id of azurerm_firewall_policy_rule_collection_group.
func (afprcg azurermFirewallPolicyRuleCollectionGroupAttributes) FirewallPolicyId() terra.StringValue {
	return terra.ReferenceAsString(afprcg.ref.Append("firewall_policy_id"))
}

// Id returns a reference to field id of azurerm_firewall_policy_rule_collection_group.
func (afprcg azurermFirewallPolicyRuleCollectionGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(afprcg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_firewall_policy_rule_collection_group.
func (afprcg azurermFirewallPolicyRuleCollectionGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(afprcg.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_firewall_policy_rule_collection_group.
func (afprcg azurermFirewallPolicyRuleCollectionGroupAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(afprcg.ref.Append("priority"))
}

func (afprcg azurermFirewallPolicyRuleCollectionGroupAttributes) ApplicationRuleCollection() terra.ListValue[ApplicationRuleCollectionAttributes] {
	return terra.ReferenceAsList[ApplicationRuleCollectionAttributes](afprcg.ref.Append("application_rule_collection"))
}

func (afprcg azurermFirewallPolicyRuleCollectionGroupAttributes) NatRuleCollection() terra.ListValue[NatRuleCollectionAttributes] {
	return terra.ReferenceAsList[NatRuleCollectionAttributes](afprcg.ref.Append("nat_rule_collection"))
}

func (afprcg azurermFirewallPolicyRuleCollectionGroupAttributes) NetworkRuleCollection() terra.ListValue[NetworkRuleCollectionAttributes] {
	return terra.ReferenceAsList[NetworkRuleCollectionAttributes](afprcg.ref.Append("network_rule_collection"))
}

func (afprcg azurermFirewallPolicyRuleCollectionGroupAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](afprcg.ref.Append("timeouts"))
}

type azurermFirewallPolicyRuleCollectionGroupState struct {
	FirewallPolicyId          string                           `json:"firewall_policy_id"`
	Id                        string                           `json:"id"`
	Name                      string                           `json:"name"`
	Priority                  float64                          `json:"priority"`
	ApplicationRuleCollection []ApplicationRuleCollectionState `json:"application_rule_collection"`
	NatRuleCollection         []NatRuleCollectionState         `json:"nat_rule_collection"`
	NetworkRuleCollection     []NetworkRuleCollectionState     `json:"network_rule_collection"`
	Timeouts                  *TimeoutsState                   `json:"timeouts"`
}
