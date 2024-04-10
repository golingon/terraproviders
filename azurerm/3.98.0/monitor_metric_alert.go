// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	monitormetricalert "github.com/golingon/terraproviders/azurerm/3.98.0/monitormetricalert"
	"io"
)

// NewMonitorMetricAlert creates a new instance of [MonitorMetricAlert].
func NewMonitorMetricAlert(name string, args MonitorMetricAlertArgs) *MonitorMetricAlert {
	return &MonitorMetricAlert{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorMetricAlert)(nil)

// MonitorMetricAlert represents the Terraform resource azurerm_monitor_metric_alert.
type MonitorMetricAlert struct {
	Name      string
	Args      MonitorMetricAlertArgs
	state     *monitorMetricAlertState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorMetricAlert].
func (mma *MonitorMetricAlert) Type() string {
	return "azurerm_monitor_metric_alert"
}

// LocalName returns the local name for [MonitorMetricAlert].
func (mma *MonitorMetricAlert) LocalName() string {
	return mma.Name
}

// Configuration returns the configuration (args) for [MonitorMetricAlert].
func (mma *MonitorMetricAlert) Configuration() interface{} {
	return mma.Args
}

// DependOn is used for other resources to depend on [MonitorMetricAlert].
func (mma *MonitorMetricAlert) DependOn() terra.Reference {
	return terra.ReferenceResource(mma)
}

// Dependencies returns the list of resources [MonitorMetricAlert] depends_on.
func (mma *MonitorMetricAlert) Dependencies() terra.Dependencies {
	return mma.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorMetricAlert].
func (mma *MonitorMetricAlert) LifecycleManagement() *terra.Lifecycle {
	return mma.Lifecycle
}

// Attributes returns the attributes for [MonitorMetricAlert].
func (mma *MonitorMetricAlert) Attributes() monitorMetricAlertAttributes {
	return monitorMetricAlertAttributes{ref: terra.ReferenceResource(mma)}
}

// ImportState imports the given attribute values into [MonitorMetricAlert]'s state.
func (mma *MonitorMetricAlert) ImportState(av io.Reader) error {
	mma.state = &monitorMetricAlertState{}
	if err := json.NewDecoder(av).Decode(mma.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mma.Type(), mma.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorMetricAlert] has state.
func (mma *MonitorMetricAlert) State() (*monitorMetricAlertState, bool) {
	return mma.state, mma.state != nil
}

// StateMust returns the state for [MonitorMetricAlert]. Panics if the state is nil.
func (mma *MonitorMetricAlert) StateMust() *monitorMetricAlertState {
	if mma.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mma.Type(), mma.LocalName()))
	}
	return mma.state
}

// MonitorMetricAlertArgs contains the configurations for azurerm_monitor_metric_alert.
type MonitorMetricAlertArgs struct {
	// AutoMitigate: bool, optional
	AutoMitigate terra.BoolValue `hcl:"auto_mitigate,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Frequency: string, optional
	Frequency terra.StringValue `hcl:"frequency,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Scopes: set of string, required
	Scopes terra.SetValue[terra.StringValue] `hcl:"scopes,attr" validate:"required"`
	// Severity: number, optional
	Severity terra.NumberValue `hcl:"severity,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TargetResourceLocation: string, optional
	TargetResourceLocation terra.StringValue `hcl:"target_resource_location,attr"`
	// TargetResourceType: string, optional
	TargetResourceType terra.StringValue `hcl:"target_resource_type,attr"`
	// WindowSize: string, optional
	WindowSize terra.StringValue `hcl:"window_size,attr"`
	// Action: min=0
	Action []monitormetricalert.Action `hcl:"action,block" validate:"min=0"`
	// ApplicationInsightsWebTestLocationAvailabilityCriteria: optional
	ApplicationInsightsWebTestLocationAvailabilityCriteria *monitormetricalert.ApplicationInsightsWebTestLocationAvailabilityCriteria `hcl:"application_insights_web_test_location_availability_criteria,block"`
	// Criteria: min=0
	Criteria []monitormetricalert.Criteria `hcl:"criteria,block" validate:"min=0"`
	// DynamicCriteria: optional
	DynamicCriteria *monitormetricalert.DynamicCriteria `hcl:"dynamic_criteria,block"`
	// Timeouts: optional
	Timeouts *monitormetricalert.Timeouts `hcl:"timeouts,block"`
}
type monitorMetricAlertAttributes struct {
	ref terra.Reference
}

// AutoMitigate returns a reference to field auto_mitigate of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) AutoMitigate() terra.BoolValue {
	return terra.ReferenceAsBool(mma.ref.Append("auto_mitigate"))
}

// Description returns a reference to field description of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mma.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(mma.ref.Append("enabled"))
}

// Frequency returns a reference to field frequency of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(mma.ref.Append("frequency"))
}

// Id returns a reference to field id of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mma.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mma.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mma.ref.Append("resource_group_name"))
}

// Scopes returns a reference to field scopes of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) Scopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mma.ref.Append("scopes"))
}

// Severity returns a reference to field severity of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) Severity() terra.NumberValue {
	return terra.ReferenceAsNumber(mma.ref.Append("severity"))
}

// Tags returns a reference to field tags of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mma.ref.Append("tags"))
}

// TargetResourceLocation returns a reference to field target_resource_location of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) TargetResourceLocation() terra.StringValue {
	return terra.ReferenceAsString(mma.ref.Append("target_resource_location"))
}

// TargetResourceType returns a reference to field target_resource_type of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) TargetResourceType() terra.StringValue {
	return terra.ReferenceAsString(mma.ref.Append("target_resource_type"))
}

// WindowSize returns a reference to field window_size of azurerm_monitor_metric_alert.
func (mma monitorMetricAlertAttributes) WindowSize() terra.StringValue {
	return terra.ReferenceAsString(mma.ref.Append("window_size"))
}

func (mma monitorMetricAlertAttributes) Action() terra.SetValue[monitormetricalert.ActionAttributes] {
	return terra.ReferenceAsSet[monitormetricalert.ActionAttributes](mma.ref.Append("action"))
}

func (mma monitorMetricAlertAttributes) ApplicationInsightsWebTestLocationAvailabilityCriteria() terra.ListValue[monitormetricalert.ApplicationInsightsWebTestLocationAvailabilityCriteriaAttributes] {
	return terra.ReferenceAsList[monitormetricalert.ApplicationInsightsWebTestLocationAvailabilityCriteriaAttributes](mma.ref.Append("application_insights_web_test_location_availability_criteria"))
}

func (mma monitorMetricAlertAttributes) Criteria() terra.ListValue[monitormetricalert.CriteriaAttributes] {
	return terra.ReferenceAsList[monitormetricalert.CriteriaAttributes](mma.ref.Append("criteria"))
}

func (mma monitorMetricAlertAttributes) DynamicCriteria() terra.ListValue[monitormetricalert.DynamicCriteriaAttributes] {
	return terra.ReferenceAsList[monitormetricalert.DynamicCriteriaAttributes](mma.ref.Append("dynamic_criteria"))
}

func (mma monitorMetricAlertAttributes) Timeouts() monitormetricalert.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitormetricalert.TimeoutsAttributes](mma.ref.Append("timeouts"))
}

type monitorMetricAlertState struct {
	AutoMitigate                                           bool                                                                             `json:"auto_mitigate"`
	Description                                            string                                                                           `json:"description"`
	Enabled                                                bool                                                                             `json:"enabled"`
	Frequency                                              string                                                                           `json:"frequency"`
	Id                                                     string                                                                           `json:"id"`
	Name                                                   string                                                                           `json:"name"`
	ResourceGroupName                                      string                                                                           `json:"resource_group_name"`
	Scopes                                                 []string                                                                         `json:"scopes"`
	Severity                                               float64                                                                          `json:"severity"`
	Tags                                                   map[string]string                                                                `json:"tags"`
	TargetResourceLocation                                 string                                                                           `json:"target_resource_location"`
	TargetResourceType                                     string                                                                           `json:"target_resource_type"`
	WindowSize                                             string                                                                           `json:"window_size"`
	Action                                                 []monitormetricalert.ActionState                                                 `json:"action"`
	ApplicationInsightsWebTestLocationAvailabilityCriteria []monitormetricalert.ApplicationInsightsWebTestLocationAvailabilityCriteriaState `json:"application_insights_web_test_location_availability_criteria"`
	Criteria                                               []monitormetricalert.CriteriaState                                               `json:"criteria"`
	DynamicCriteria                                        []monitormetricalert.DynamicCriteriaState                                        `json:"dynamic_criteria"`
	Timeouts                                               *monitormetricalert.TimeoutsState                                                `json:"timeouts"`
}
