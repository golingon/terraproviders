// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitoralertprocessingrulesuppression "github.com/golingon/terraproviders/azurerm/3.49.0/monitoralertprocessingrulesuppression"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewMonitorAlertProcessingRuleSuppression(name string, args MonitorAlertProcessingRuleSuppressionArgs) *MonitorAlertProcessingRuleSuppression {
	return &MonitorAlertProcessingRuleSuppression{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorAlertProcessingRuleSuppression)(nil)

type MonitorAlertProcessingRuleSuppression struct {
	Name  string
	Args  MonitorAlertProcessingRuleSuppressionArgs
	state *monitorAlertProcessingRuleSuppressionState
}

func (maprs *MonitorAlertProcessingRuleSuppression) Type() string {
	return "azurerm_monitor_alert_processing_rule_suppression"
}

func (maprs *MonitorAlertProcessingRuleSuppression) LocalName() string {
	return maprs.Name
}

func (maprs *MonitorAlertProcessingRuleSuppression) Configuration() interface{} {
	return maprs.Args
}

func (maprs *MonitorAlertProcessingRuleSuppression) Attributes() monitorAlertProcessingRuleSuppressionAttributes {
	return monitorAlertProcessingRuleSuppressionAttributes{ref: terra.ReferenceResource(maprs)}
}

func (maprs *MonitorAlertProcessingRuleSuppression) ImportState(av io.Reader) error {
	maprs.state = &monitorAlertProcessingRuleSuppressionState{}
	if err := json.NewDecoder(av).Decode(maprs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", maprs.Type(), maprs.LocalName(), err)
	}
	return nil
}

func (maprs *MonitorAlertProcessingRuleSuppression) State() (*monitorAlertProcessingRuleSuppressionState, bool) {
	return maprs.state, maprs.state != nil
}

func (maprs *MonitorAlertProcessingRuleSuppression) StateMust() *monitorAlertProcessingRuleSuppressionState {
	if maprs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", maprs.Type(), maprs.LocalName()))
	}
	return maprs.state
}

func (maprs *MonitorAlertProcessingRuleSuppression) DependOn() terra.Reference {
	return terra.ReferenceResource(maprs)
}

type MonitorAlertProcessingRuleSuppressionArgs struct {
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
	Condition *monitoralertprocessingrulesuppression.Condition `hcl:"condition,block"`
	// Schedule: optional
	Schedule *monitoralertprocessingrulesuppression.Schedule `hcl:"schedule,block"`
	// Timeouts: optional
	Timeouts *monitoralertprocessingrulesuppression.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that MonitorAlertProcessingRuleSuppression depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type monitorAlertProcessingRuleSuppressionAttributes struct {
	ref terra.Reference
}

func (maprs monitorAlertProcessingRuleSuppressionAttributes) Description() terra.StringValue {
	return terra.ReferenceString(maprs.ref.Append("description"))
}

func (maprs monitorAlertProcessingRuleSuppressionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(maprs.ref.Append("enabled"))
}

func (maprs monitorAlertProcessingRuleSuppressionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(maprs.ref.Append("id"))
}

func (maprs monitorAlertProcessingRuleSuppressionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(maprs.ref.Append("name"))
}

func (maprs monitorAlertProcessingRuleSuppressionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(maprs.ref.Append("resource_group_name"))
}

func (maprs monitorAlertProcessingRuleSuppressionAttributes) Scopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](maprs.ref.Append("scopes"))
}

func (maprs monitorAlertProcessingRuleSuppressionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](maprs.ref.Append("tags"))
}

func (maprs monitorAlertProcessingRuleSuppressionAttributes) Condition() terra.ListValue[monitoralertprocessingrulesuppression.ConditionAttributes] {
	return terra.ReferenceList[monitoralertprocessingrulesuppression.ConditionAttributes](maprs.ref.Append("condition"))
}

func (maprs monitorAlertProcessingRuleSuppressionAttributes) Schedule() terra.ListValue[monitoralertprocessingrulesuppression.ScheduleAttributes] {
	return terra.ReferenceList[monitoralertprocessingrulesuppression.ScheduleAttributes](maprs.ref.Append("schedule"))
}

func (maprs monitorAlertProcessingRuleSuppressionAttributes) Timeouts() monitoralertprocessingrulesuppression.TimeoutsAttributes {
	return terra.ReferenceSingle[monitoralertprocessingrulesuppression.TimeoutsAttributes](maprs.ref.Append("timeouts"))
}

type monitorAlertProcessingRuleSuppressionState struct {
	Description       string                                                 `json:"description"`
	Enabled           bool                                                   `json:"enabled"`
	Id                string                                                 `json:"id"`
	Name              string                                                 `json:"name"`
	ResourceGroupName string                                                 `json:"resource_group_name"`
	Scopes            []string                                               `json:"scopes"`
	Tags              map[string]string                                      `json:"tags"`
	Condition         []monitoralertprocessingrulesuppression.ConditionState `json:"condition"`
	Schedule          []monitoralertprocessingrulesuppression.ScheduleState  `json:"schedule"`
	Timeouts          *monitoralertprocessingrulesuppression.TimeoutsState   `json:"timeouts"`
}
