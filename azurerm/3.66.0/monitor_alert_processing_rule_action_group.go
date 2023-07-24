// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitoralertprocessingruleactiongroup "github.com/golingon/terraproviders/azurerm/3.66.0/monitoralertprocessingruleactiongroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorAlertProcessingRuleActionGroup creates a new instance of [MonitorAlertProcessingRuleActionGroup].
func NewMonitorAlertProcessingRuleActionGroup(name string, args MonitorAlertProcessingRuleActionGroupArgs) *MonitorAlertProcessingRuleActionGroup {
	return &MonitorAlertProcessingRuleActionGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorAlertProcessingRuleActionGroup)(nil)

// MonitorAlertProcessingRuleActionGroup represents the Terraform resource azurerm_monitor_alert_processing_rule_action_group.
type MonitorAlertProcessingRuleActionGroup struct {
	Name      string
	Args      MonitorAlertProcessingRuleActionGroupArgs
	state     *monitorAlertProcessingRuleActionGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorAlertProcessingRuleActionGroup].
func (maprag *MonitorAlertProcessingRuleActionGroup) Type() string {
	return "azurerm_monitor_alert_processing_rule_action_group"
}

// LocalName returns the local name for [MonitorAlertProcessingRuleActionGroup].
func (maprag *MonitorAlertProcessingRuleActionGroup) LocalName() string {
	return maprag.Name
}

// Configuration returns the configuration (args) for [MonitorAlertProcessingRuleActionGroup].
func (maprag *MonitorAlertProcessingRuleActionGroup) Configuration() interface{} {
	return maprag.Args
}

// DependOn is used for other resources to depend on [MonitorAlertProcessingRuleActionGroup].
func (maprag *MonitorAlertProcessingRuleActionGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(maprag)
}

// Dependencies returns the list of resources [MonitorAlertProcessingRuleActionGroup] depends_on.
func (maprag *MonitorAlertProcessingRuleActionGroup) Dependencies() terra.Dependencies {
	return maprag.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorAlertProcessingRuleActionGroup].
func (maprag *MonitorAlertProcessingRuleActionGroup) LifecycleManagement() *terra.Lifecycle {
	return maprag.Lifecycle
}

// Attributes returns the attributes for [MonitorAlertProcessingRuleActionGroup].
func (maprag *MonitorAlertProcessingRuleActionGroup) Attributes() monitorAlertProcessingRuleActionGroupAttributes {
	return monitorAlertProcessingRuleActionGroupAttributes{ref: terra.ReferenceResource(maprag)}
}

// ImportState imports the given attribute values into [MonitorAlertProcessingRuleActionGroup]'s state.
func (maprag *MonitorAlertProcessingRuleActionGroup) ImportState(av io.Reader) error {
	maprag.state = &monitorAlertProcessingRuleActionGroupState{}
	if err := json.NewDecoder(av).Decode(maprag.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", maprag.Type(), maprag.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorAlertProcessingRuleActionGroup] has state.
func (maprag *MonitorAlertProcessingRuleActionGroup) State() (*monitorAlertProcessingRuleActionGroupState, bool) {
	return maprag.state, maprag.state != nil
}

// StateMust returns the state for [MonitorAlertProcessingRuleActionGroup]. Panics if the state is nil.
func (maprag *MonitorAlertProcessingRuleActionGroup) StateMust() *monitorAlertProcessingRuleActionGroupState {
	if maprag.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", maprag.Type(), maprag.LocalName()))
	}
	return maprag.state
}

// MonitorAlertProcessingRuleActionGroupArgs contains the configurations for azurerm_monitor_alert_processing_rule_action_group.
type MonitorAlertProcessingRuleActionGroupArgs struct {
	// AddActionGroupIds: list of string, required
	AddActionGroupIds terra.ListValue[terra.StringValue] `hcl:"add_action_group_ids,attr" validate:"required"`
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
	// Scopes: list of string, required
	Scopes terra.ListValue[terra.StringValue] `hcl:"scopes,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Condition: optional
	Condition *monitoralertprocessingruleactiongroup.Condition `hcl:"condition,block"`
	// Schedule: optional
	Schedule *monitoralertprocessingruleactiongroup.Schedule `hcl:"schedule,block"`
	// Timeouts: optional
	Timeouts *monitoralertprocessingruleactiongroup.Timeouts `hcl:"timeouts,block"`
}
type monitorAlertProcessingRuleActionGroupAttributes struct {
	ref terra.Reference
}

// AddActionGroupIds returns a reference to field add_action_group_ids of azurerm_monitor_alert_processing_rule_action_group.
func (maprag monitorAlertProcessingRuleActionGroupAttributes) AddActionGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](maprag.ref.Append("add_action_group_ids"))
}

// Description returns a reference to field description of azurerm_monitor_alert_processing_rule_action_group.
func (maprag monitorAlertProcessingRuleActionGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(maprag.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_monitor_alert_processing_rule_action_group.
func (maprag monitorAlertProcessingRuleActionGroupAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(maprag.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_monitor_alert_processing_rule_action_group.
func (maprag monitorAlertProcessingRuleActionGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(maprag.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_monitor_alert_processing_rule_action_group.
func (maprag monitorAlertProcessingRuleActionGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(maprag.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_alert_processing_rule_action_group.
func (maprag monitorAlertProcessingRuleActionGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(maprag.ref.Append("resource_group_name"))
}

// Scopes returns a reference to field scopes of azurerm_monitor_alert_processing_rule_action_group.
func (maprag monitorAlertProcessingRuleActionGroupAttributes) Scopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](maprag.ref.Append("scopes"))
}

// Tags returns a reference to field tags of azurerm_monitor_alert_processing_rule_action_group.
func (maprag monitorAlertProcessingRuleActionGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](maprag.ref.Append("tags"))
}

func (maprag monitorAlertProcessingRuleActionGroupAttributes) Condition() terra.ListValue[monitoralertprocessingruleactiongroup.ConditionAttributes] {
	return terra.ReferenceAsList[monitoralertprocessingruleactiongroup.ConditionAttributes](maprag.ref.Append("condition"))
}

func (maprag monitorAlertProcessingRuleActionGroupAttributes) Schedule() terra.ListValue[monitoralertprocessingruleactiongroup.ScheduleAttributes] {
	return terra.ReferenceAsList[monitoralertprocessingruleactiongroup.ScheduleAttributes](maprag.ref.Append("schedule"))
}

func (maprag monitorAlertProcessingRuleActionGroupAttributes) Timeouts() monitoralertprocessingruleactiongroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoralertprocessingruleactiongroup.TimeoutsAttributes](maprag.ref.Append("timeouts"))
}

type monitorAlertProcessingRuleActionGroupState struct {
	AddActionGroupIds []string                                               `json:"add_action_group_ids"`
	Description       string                                                 `json:"description"`
	Enabled           bool                                                   `json:"enabled"`
	Id                string                                                 `json:"id"`
	Name              string                                                 `json:"name"`
	ResourceGroupName string                                                 `json:"resource_group_name"`
	Scopes            []string                                               `json:"scopes"`
	Tags              map[string]string                                      `json:"tags"`
	Condition         []monitoralertprocessingruleactiongroup.ConditionState `json:"condition"`
	Schedule          []monitoralertprocessingruleactiongroup.ScheduleState  `json:"schedule"`
	Timeouts          *monitoralertprocessingruleactiongroup.TimeoutsState   `json:"timeouts"`
}
