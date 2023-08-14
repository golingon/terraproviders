// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitoractionrulesuppression "github.com/golingon/terraproviders/azurerm/3.69.0/monitoractionrulesuppression"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorActionRuleSuppression creates a new instance of [MonitorActionRuleSuppression].
func NewMonitorActionRuleSuppression(name string, args MonitorActionRuleSuppressionArgs) *MonitorActionRuleSuppression {
	return &MonitorActionRuleSuppression{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorActionRuleSuppression)(nil)

// MonitorActionRuleSuppression represents the Terraform resource azurerm_monitor_action_rule_suppression.
type MonitorActionRuleSuppression struct {
	Name      string
	Args      MonitorActionRuleSuppressionArgs
	state     *monitorActionRuleSuppressionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorActionRuleSuppression].
func (mars *MonitorActionRuleSuppression) Type() string {
	return "azurerm_monitor_action_rule_suppression"
}

// LocalName returns the local name for [MonitorActionRuleSuppression].
func (mars *MonitorActionRuleSuppression) LocalName() string {
	return mars.Name
}

// Configuration returns the configuration (args) for [MonitorActionRuleSuppression].
func (mars *MonitorActionRuleSuppression) Configuration() interface{} {
	return mars.Args
}

// DependOn is used for other resources to depend on [MonitorActionRuleSuppression].
func (mars *MonitorActionRuleSuppression) DependOn() terra.Reference {
	return terra.ReferenceResource(mars)
}

// Dependencies returns the list of resources [MonitorActionRuleSuppression] depends_on.
func (mars *MonitorActionRuleSuppression) Dependencies() terra.Dependencies {
	return mars.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorActionRuleSuppression].
func (mars *MonitorActionRuleSuppression) LifecycleManagement() *terra.Lifecycle {
	return mars.Lifecycle
}

// Attributes returns the attributes for [MonitorActionRuleSuppression].
func (mars *MonitorActionRuleSuppression) Attributes() monitorActionRuleSuppressionAttributes {
	return monitorActionRuleSuppressionAttributes{ref: terra.ReferenceResource(mars)}
}

// ImportState imports the given attribute values into [MonitorActionRuleSuppression]'s state.
func (mars *MonitorActionRuleSuppression) ImportState(av io.Reader) error {
	mars.state = &monitorActionRuleSuppressionState{}
	if err := json.NewDecoder(av).Decode(mars.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mars.Type(), mars.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorActionRuleSuppression] has state.
func (mars *MonitorActionRuleSuppression) State() (*monitorActionRuleSuppressionState, bool) {
	return mars.state, mars.state != nil
}

// StateMust returns the state for [MonitorActionRuleSuppression]. Panics if the state is nil.
func (mars *MonitorActionRuleSuppression) StateMust() *monitorActionRuleSuppressionState {
	if mars.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mars.Type(), mars.LocalName()))
	}
	return mars.state
}

// MonitorActionRuleSuppressionArgs contains the configurations for azurerm_monitor_action_rule_suppression.
type MonitorActionRuleSuppressionArgs struct {
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
	Condition *monitoractionrulesuppression.Condition `hcl:"condition,block"`
	// Scope: optional
	Scope *monitoractionrulesuppression.Scope `hcl:"scope,block"`
	// Suppression: required
	Suppression *monitoractionrulesuppression.Suppression `hcl:"suppression,block" validate:"required"`
	// Timeouts: optional
	Timeouts *monitoractionrulesuppression.Timeouts `hcl:"timeouts,block"`
}
type monitorActionRuleSuppressionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_monitor_action_rule_suppression.
func (mars monitorActionRuleSuppressionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mars.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_monitor_action_rule_suppression.
func (mars monitorActionRuleSuppressionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(mars.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_monitor_action_rule_suppression.
func (mars monitorActionRuleSuppressionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mars.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_monitor_action_rule_suppression.
func (mars monitorActionRuleSuppressionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mars.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_action_rule_suppression.
func (mars monitorActionRuleSuppressionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mars.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_monitor_action_rule_suppression.
func (mars monitorActionRuleSuppressionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mars.ref.Append("tags"))
}

func (mars monitorActionRuleSuppressionAttributes) Condition() terra.ListValue[monitoractionrulesuppression.ConditionAttributes] {
	return terra.ReferenceAsList[monitoractionrulesuppression.ConditionAttributes](mars.ref.Append("condition"))
}

func (mars monitorActionRuleSuppressionAttributes) Scope() terra.ListValue[monitoractionrulesuppression.ScopeAttributes] {
	return terra.ReferenceAsList[monitoractionrulesuppression.ScopeAttributes](mars.ref.Append("scope"))
}

func (mars monitorActionRuleSuppressionAttributes) Suppression() terra.ListValue[monitoractionrulesuppression.SuppressionAttributes] {
	return terra.ReferenceAsList[monitoractionrulesuppression.SuppressionAttributes](mars.ref.Append("suppression"))
}

func (mars monitorActionRuleSuppressionAttributes) Timeouts() monitoractionrulesuppression.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoractionrulesuppression.TimeoutsAttributes](mars.ref.Append("timeouts"))
}

type monitorActionRuleSuppressionState struct {
	Description       string                                          `json:"description"`
	Enabled           bool                                            `json:"enabled"`
	Id                string                                          `json:"id"`
	Name              string                                          `json:"name"`
	ResourceGroupName string                                          `json:"resource_group_name"`
	Tags              map[string]string                               `json:"tags"`
	Condition         []monitoractionrulesuppression.ConditionState   `json:"condition"`
	Scope             []monitoractionrulesuppression.ScopeState       `json:"scope"`
	Suppression       []monitoractionrulesuppression.SuppressionState `json:"suppression"`
	Timeouts          *monitoractionrulesuppression.TimeoutsState     `json:"timeouts"`
}
