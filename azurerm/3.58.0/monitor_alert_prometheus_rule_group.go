// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitoralertprometheusrulegroup "github.com/golingon/terraproviders/azurerm/3.58.0/monitoralertprometheusrulegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorAlertPrometheusRuleGroup creates a new instance of [MonitorAlertPrometheusRuleGroup].
func NewMonitorAlertPrometheusRuleGroup(name string, args MonitorAlertPrometheusRuleGroupArgs) *MonitorAlertPrometheusRuleGroup {
	return &MonitorAlertPrometheusRuleGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorAlertPrometheusRuleGroup)(nil)

// MonitorAlertPrometheusRuleGroup represents the Terraform resource azurerm_monitor_alert_prometheus_rule_group.
type MonitorAlertPrometheusRuleGroup struct {
	Name      string
	Args      MonitorAlertPrometheusRuleGroupArgs
	state     *monitorAlertPrometheusRuleGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorAlertPrometheusRuleGroup].
func (maprg *MonitorAlertPrometheusRuleGroup) Type() string {
	return "azurerm_monitor_alert_prometheus_rule_group"
}

// LocalName returns the local name for [MonitorAlertPrometheusRuleGroup].
func (maprg *MonitorAlertPrometheusRuleGroup) LocalName() string {
	return maprg.Name
}

// Configuration returns the configuration (args) for [MonitorAlertPrometheusRuleGroup].
func (maprg *MonitorAlertPrometheusRuleGroup) Configuration() interface{} {
	return maprg.Args
}

// DependOn is used for other resources to depend on [MonitorAlertPrometheusRuleGroup].
func (maprg *MonitorAlertPrometheusRuleGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(maprg)
}

// Dependencies returns the list of resources [MonitorAlertPrometheusRuleGroup] depends_on.
func (maprg *MonitorAlertPrometheusRuleGroup) Dependencies() terra.Dependencies {
	return maprg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorAlertPrometheusRuleGroup].
func (maprg *MonitorAlertPrometheusRuleGroup) LifecycleManagement() *terra.Lifecycle {
	return maprg.Lifecycle
}

// Attributes returns the attributes for [MonitorAlertPrometheusRuleGroup].
func (maprg *MonitorAlertPrometheusRuleGroup) Attributes() monitorAlertPrometheusRuleGroupAttributes {
	return monitorAlertPrometheusRuleGroupAttributes{ref: terra.ReferenceResource(maprg)}
}

// ImportState imports the given attribute values into [MonitorAlertPrometheusRuleGroup]'s state.
func (maprg *MonitorAlertPrometheusRuleGroup) ImportState(av io.Reader) error {
	maprg.state = &monitorAlertPrometheusRuleGroupState{}
	if err := json.NewDecoder(av).Decode(maprg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", maprg.Type(), maprg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorAlertPrometheusRuleGroup] has state.
func (maprg *MonitorAlertPrometheusRuleGroup) State() (*monitorAlertPrometheusRuleGroupState, bool) {
	return maprg.state, maprg.state != nil
}

// StateMust returns the state for [MonitorAlertPrometheusRuleGroup]. Panics if the state is nil.
func (maprg *MonitorAlertPrometheusRuleGroup) StateMust() *monitorAlertPrometheusRuleGroupState {
	if maprg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", maprg.Type(), maprg.LocalName()))
	}
	return maprg.state
}

// MonitorAlertPrometheusRuleGroupArgs contains the configurations for azurerm_monitor_alert_prometheus_rule_group.
type MonitorAlertPrometheusRuleGroupArgs struct {
	// ClusterName: string, optional
	ClusterName terra.StringValue `hcl:"cluster_name,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Interval: string, optional
	Interval terra.StringValue `hcl:"interval,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RuleGroupEnabled: bool, optional
	RuleGroupEnabled terra.BoolValue `hcl:"rule_group_enabled,attr"`
	// Scopes: list of string, required
	Scopes terra.ListValue[terra.StringValue] `hcl:"scopes,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Rule: min=1
	Rule []monitoralertprometheusrulegroup.Rule `hcl:"rule,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *monitoralertprometheusrulegroup.Timeouts `hcl:"timeouts,block"`
}
type monitorAlertPrometheusRuleGroupAttributes struct {
	ref terra.Reference
}

// ClusterName returns a reference to field cluster_name of azurerm_monitor_alert_prometheus_rule_group.
func (maprg monitorAlertPrometheusRuleGroupAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(maprg.ref.Append("cluster_name"))
}

// Description returns a reference to field description of azurerm_monitor_alert_prometheus_rule_group.
func (maprg monitorAlertPrometheusRuleGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(maprg.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_monitor_alert_prometheus_rule_group.
func (maprg monitorAlertPrometheusRuleGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(maprg.ref.Append("id"))
}

// Interval returns a reference to field interval of azurerm_monitor_alert_prometheus_rule_group.
func (maprg monitorAlertPrometheusRuleGroupAttributes) Interval() terra.StringValue {
	return terra.ReferenceAsString(maprg.ref.Append("interval"))
}

// Location returns a reference to field location of azurerm_monitor_alert_prometheus_rule_group.
func (maprg monitorAlertPrometheusRuleGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(maprg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_monitor_alert_prometheus_rule_group.
func (maprg monitorAlertPrometheusRuleGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(maprg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_alert_prometheus_rule_group.
func (maprg monitorAlertPrometheusRuleGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(maprg.ref.Append("resource_group_name"))
}

// RuleGroupEnabled returns a reference to field rule_group_enabled of azurerm_monitor_alert_prometheus_rule_group.
func (maprg monitorAlertPrometheusRuleGroupAttributes) RuleGroupEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(maprg.ref.Append("rule_group_enabled"))
}

// Scopes returns a reference to field scopes of azurerm_monitor_alert_prometheus_rule_group.
func (maprg monitorAlertPrometheusRuleGroupAttributes) Scopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](maprg.ref.Append("scopes"))
}

// Tags returns a reference to field tags of azurerm_monitor_alert_prometheus_rule_group.
func (maprg monitorAlertPrometheusRuleGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](maprg.ref.Append("tags"))
}

func (maprg monitorAlertPrometheusRuleGroupAttributes) Rule() terra.ListValue[monitoralertprometheusrulegroup.RuleAttributes] {
	return terra.ReferenceAsList[monitoralertprometheusrulegroup.RuleAttributes](maprg.ref.Append("rule"))
}

func (maprg monitorAlertPrometheusRuleGroupAttributes) Timeouts() monitoralertprometheusrulegroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoralertprometheusrulegroup.TimeoutsAttributes](maprg.ref.Append("timeouts"))
}

type monitorAlertPrometheusRuleGroupState struct {
	ClusterName       string                                         `json:"cluster_name"`
	Description       string                                         `json:"description"`
	Id                string                                         `json:"id"`
	Interval          string                                         `json:"interval"`
	Location          string                                         `json:"location"`
	Name              string                                         `json:"name"`
	ResourceGroupName string                                         `json:"resource_group_name"`
	RuleGroupEnabled  bool                                           `json:"rule_group_enabled"`
	Scopes            []string                                       `json:"scopes"`
	Tags              map[string]string                              `json:"tags"`
	Rule              []monitoralertprometheusrulegroup.RuleState    `json:"rule"`
	Timeouts          *monitoralertprometheusrulegroup.TimeoutsState `json:"timeouts"`
}
