// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datadogmonitortagrule "github.com/golingon/terraproviders/azurerm/3.58.0/datadogmonitortagrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatadogMonitorTagRule creates a new instance of [DatadogMonitorTagRule].
func NewDatadogMonitorTagRule(name string, args DatadogMonitorTagRuleArgs) *DatadogMonitorTagRule {
	return &DatadogMonitorTagRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatadogMonitorTagRule)(nil)

// DatadogMonitorTagRule represents the Terraform resource azurerm_datadog_monitor_tag_rule.
type DatadogMonitorTagRule struct {
	Name      string
	Args      DatadogMonitorTagRuleArgs
	state     *datadogMonitorTagRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatadogMonitorTagRule].
func (dmtr *DatadogMonitorTagRule) Type() string {
	return "azurerm_datadog_monitor_tag_rule"
}

// LocalName returns the local name for [DatadogMonitorTagRule].
func (dmtr *DatadogMonitorTagRule) LocalName() string {
	return dmtr.Name
}

// Configuration returns the configuration (args) for [DatadogMonitorTagRule].
func (dmtr *DatadogMonitorTagRule) Configuration() interface{} {
	return dmtr.Args
}

// DependOn is used for other resources to depend on [DatadogMonitorTagRule].
func (dmtr *DatadogMonitorTagRule) DependOn() terra.Reference {
	return terra.ReferenceResource(dmtr)
}

// Dependencies returns the list of resources [DatadogMonitorTagRule] depends_on.
func (dmtr *DatadogMonitorTagRule) Dependencies() terra.Dependencies {
	return dmtr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatadogMonitorTagRule].
func (dmtr *DatadogMonitorTagRule) LifecycleManagement() *terra.Lifecycle {
	return dmtr.Lifecycle
}

// Attributes returns the attributes for [DatadogMonitorTagRule].
func (dmtr *DatadogMonitorTagRule) Attributes() datadogMonitorTagRuleAttributes {
	return datadogMonitorTagRuleAttributes{ref: terra.ReferenceResource(dmtr)}
}

// ImportState imports the given attribute values into [DatadogMonitorTagRule]'s state.
func (dmtr *DatadogMonitorTagRule) ImportState(av io.Reader) error {
	dmtr.state = &datadogMonitorTagRuleState{}
	if err := json.NewDecoder(av).Decode(dmtr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmtr.Type(), dmtr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatadogMonitorTagRule] has state.
func (dmtr *DatadogMonitorTagRule) State() (*datadogMonitorTagRuleState, bool) {
	return dmtr.state, dmtr.state != nil
}

// StateMust returns the state for [DatadogMonitorTagRule]. Panics if the state is nil.
func (dmtr *DatadogMonitorTagRule) StateMust() *datadogMonitorTagRuleState {
	if dmtr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmtr.Type(), dmtr.LocalName()))
	}
	return dmtr.state
}

// DatadogMonitorTagRuleArgs contains the configurations for azurerm_datadog_monitor_tag_rule.
type DatadogMonitorTagRuleArgs struct {
	// DatadogMonitorId: string, required
	DatadogMonitorId terra.StringValue `hcl:"datadog_monitor_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Log: min=0
	Log []datadogmonitortagrule.Log `hcl:"log,block" validate:"min=0"`
	// Metric: min=0
	Metric []datadogmonitortagrule.Metric `hcl:"metric,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadogmonitortagrule.Timeouts `hcl:"timeouts,block"`
}
type datadogMonitorTagRuleAttributes struct {
	ref terra.Reference
}

// DatadogMonitorId returns a reference to field datadog_monitor_id of azurerm_datadog_monitor_tag_rule.
func (dmtr datadogMonitorTagRuleAttributes) DatadogMonitorId() terra.StringValue {
	return terra.ReferenceAsString(dmtr.ref.Append("datadog_monitor_id"))
}

// Id returns a reference to field id of azurerm_datadog_monitor_tag_rule.
func (dmtr datadogMonitorTagRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmtr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_datadog_monitor_tag_rule.
func (dmtr datadogMonitorTagRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dmtr.ref.Append("name"))
}

func (dmtr datadogMonitorTagRuleAttributes) Log() terra.ListValue[datadogmonitortagrule.LogAttributes] {
	return terra.ReferenceAsList[datadogmonitortagrule.LogAttributes](dmtr.ref.Append("log"))
}

func (dmtr datadogMonitorTagRuleAttributes) Metric() terra.ListValue[datadogmonitortagrule.MetricAttributes] {
	return terra.ReferenceAsList[datadogmonitortagrule.MetricAttributes](dmtr.ref.Append("metric"))
}

func (dmtr datadogMonitorTagRuleAttributes) Timeouts() datadogmonitortagrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadogmonitortagrule.TimeoutsAttributes](dmtr.ref.Append("timeouts"))
}

type datadogMonitorTagRuleState struct {
	DatadogMonitorId string                               `json:"datadog_monitor_id"`
	Id               string                               `json:"id"`
	Name             string                               `json:"name"`
	Log              []datadogmonitortagrule.LogState     `json:"log"`
	Metric           []datadogmonitortagrule.MetricState  `json:"metric"`
	Timeouts         *datadogmonitortagrule.TimeoutsState `json:"timeouts"`
}