// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitoractionruleactiongroup "github.com/golingon/terraproviders/azurerm/3.65.0/monitoractionruleactiongroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorActionRuleActionGroup creates a new instance of [MonitorActionRuleActionGroup].
func NewMonitorActionRuleActionGroup(name string, args MonitorActionRuleActionGroupArgs) *MonitorActionRuleActionGroup {
	return &MonitorActionRuleActionGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorActionRuleActionGroup)(nil)

// MonitorActionRuleActionGroup represents the Terraform resource azurerm_monitor_action_rule_action_group.
type MonitorActionRuleActionGroup struct {
	Name      string
	Args      MonitorActionRuleActionGroupArgs
	state     *monitorActionRuleActionGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorActionRuleActionGroup].
func (marag *MonitorActionRuleActionGroup) Type() string {
	return "azurerm_monitor_action_rule_action_group"
}

// LocalName returns the local name for [MonitorActionRuleActionGroup].
func (marag *MonitorActionRuleActionGroup) LocalName() string {
	return marag.Name
}

// Configuration returns the configuration (args) for [MonitorActionRuleActionGroup].
func (marag *MonitorActionRuleActionGroup) Configuration() interface{} {
	return marag.Args
}

// DependOn is used for other resources to depend on [MonitorActionRuleActionGroup].
func (marag *MonitorActionRuleActionGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(marag)
}

// Dependencies returns the list of resources [MonitorActionRuleActionGroup] depends_on.
func (marag *MonitorActionRuleActionGroup) Dependencies() terra.Dependencies {
	return marag.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorActionRuleActionGroup].
func (marag *MonitorActionRuleActionGroup) LifecycleManagement() *terra.Lifecycle {
	return marag.Lifecycle
}

// Attributes returns the attributes for [MonitorActionRuleActionGroup].
func (marag *MonitorActionRuleActionGroup) Attributes() monitorActionRuleActionGroupAttributes {
	return monitorActionRuleActionGroupAttributes{ref: terra.ReferenceResource(marag)}
}

// ImportState imports the given attribute values into [MonitorActionRuleActionGroup]'s state.
func (marag *MonitorActionRuleActionGroup) ImportState(av io.Reader) error {
	marag.state = &monitorActionRuleActionGroupState{}
	if err := json.NewDecoder(av).Decode(marag.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", marag.Type(), marag.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorActionRuleActionGroup] has state.
func (marag *MonitorActionRuleActionGroup) State() (*monitorActionRuleActionGroupState, bool) {
	return marag.state, marag.state != nil
}

// StateMust returns the state for [MonitorActionRuleActionGroup]. Panics if the state is nil.
func (marag *MonitorActionRuleActionGroup) StateMust() *monitorActionRuleActionGroupState {
	if marag.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", marag.Type(), marag.LocalName()))
	}
	return marag.state
}

// MonitorActionRuleActionGroupArgs contains the configurations for azurerm_monitor_action_rule_action_group.
type MonitorActionRuleActionGroupArgs struct {
	// ActionGroupId: string, required
	ActionGroupId terra.StringValue `hcl:"action_group_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Condition: optional
	Condition *monitoractionruleactiongroup.Condition `hcl:"condition,block"`
	// Scope: optional
	Scope *monitoractionruleactiongroup.Scope `hcl:"scope,block"`
	// Timeouts: optional
	Timeouts *monitoractionruleactiongroup.Timeouts `hcl:"timeouts,block"`
}
type monitorActionRuleActionGroupAttributes struct {
	ref terra.Reference
}

// ActionGroupId returns a reference to field action_group_id of azurerm_monitor_action_rule_action_group.
func (marag monitorActionRuleActionGroupAttributes) ActionGroupId() terra.StringValue {
	return terra.ReferenceAsString(marag.ref.Append("action_group_id"))
}

// Description returns a reference to field description of azurerm_monitor_action_rule_action_group.
func (marag monitorActionRuleActionGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(marag.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_monitor_action_rule_action_group.
func (marag monitorActionRuleActionGroupAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(marag.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_monitor_action_rule_action_group.
func (marag monitorActionRuleActionGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(marag.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_monitor_action_rule_action_group.
func (marag monitorActionRuleActionGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(marag.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_action_rule_action_group.
func (marag monitorActionRuleActionGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(marag.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_monitor_action_rule_action_group.
func (marag monitorActionRuleActionGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](marag.ref.Append("tags"))
}

func (marag monitorActionRuleActionGroupAttributes) Condition() terra.ListValue[monitoractionruleactiongroup.ConditionAttributes] {
	return terra.ReferenceAsList[monitoractionruleactiongroup.ConditionAttributes](marag.ref.Append("condition"))
}

func (marag monitorActionRuleActionGroupAttributes) Scope() terra.ListValue[monitoractionruleactiongroup.ScopeAttributes] {
	return terra.ReferenceAsList[monitoractionruleactiongroup.ScopeAttributes](marag.ref.Append("scope"))
}

func (marag monitorActionRuleActionGroupAttributes) Timeouts() monitoractionruleactiongroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoractionruleactiongroup.TimeoutsAttributes](marag.ref.Append("timeouts"))
}

type monitorActionRuleActionGroupState struct {
	ActionGroupId     string                                        `json:"action_group_id"`
	Description       string                                        `json:"description"`
	Enabled           bool                                          `json:"enabled"`
	Id                string                                        `json:"id"`
	Name              string                                        `json:"name"`
	ResourceGroupName string                                        `json:"resource_group_name"`
	Tags              map[string]string                             `json:"tags"`
	Condition         []monitoractionruleactiongroup.ConditionState `json:"condition"`
	Scope             []monitoractionruleactiongroup.ScopeState     `json:"scope"`
	Timeouts          *monitoractionruleactiongroup.TimeoutsState   `json:"timeouts"`
}
