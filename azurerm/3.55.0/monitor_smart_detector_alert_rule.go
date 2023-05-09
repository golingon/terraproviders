// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitorsmartdetectoralertrule "github.com/golingon/terraproviders/azurerm/3.55.0/monitorsmartdetectoralertrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorSmartDetectorAlertRule creates a new instance of [MonitorSmartDetectorAlertRule].
func NewMonitorSmartDetectorAlertRule(name string, args MonitorSmartDetectorAlertRuleArgs) *MonitorSmartDetectorAlertRule {
	return &MonitorSmartDetectorAlertRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorSmartDetectorAlertRule)(nil)

// MonitorSmartDetectorAlertRule represents the Terraform resource azurerm_monitor_smart_detector_alert_rule.
type MonitorSmartDetectorAlertRule struct {
	Name      string
	Args      MonitorSmartDetectorAlertRuleArgs
	state     *monitorSmartDetectorAlertRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorSmartDetectorAlertRule].
func (msdar *MonitorSmartDetectorAlertRule) Type() string {
	return "azurerm_monitor_smart_detector_alert_rule"
}

// LocalName returns the local name for [MonitorSmartDetectorAlertRule].
func (msdar *MonitorSmartDetectorAlertRule) LocalName() string {
	return msdar.Name
}

// Configuration returns the configuration (args) for [MonitorSmartDetectorAlertRule].
func (msdar *MonitorSmartDetectorAlertRule) Configuration() interface{} {
	return msdar.Args
}

// DependOn is used for other resources to depend on [MonitorSmartDetectorAlertRule].
func (msdar *MonitorSmartDetectorAlertRule) DependOn() terra.Reference {
	return terra.ReferenceResource(msdar)
}

// Dependencies returns the list of resources [MonitorSmartDetectorAlertRule] depends_on.
func (msdar *MonitorSmartDetectorAlertRule) Dependencies() terra.Dependencies {
	return msdar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorSmartDetectorAlertRule].
func (msdar *MonitorSmartDetectorAlertRule) LifecycleManagement() *terra.Lifecycle {
	return msdar.Lifecycle
}

// Attributes returns the attributes for [MonitorSmartDetectorAlertRule].
func (msdar *MonitorSmartDetectorAlertRule) Attributes() monitorSmartDetectorAlertRuleAttributes {
	return monitorSmartDetectorAlertRuleAttributes{ref: terra.ReferenceResource(msdar)}
}

// ImportState imports the given attribute values into [MonitorSmartDetectorAlertRule]'s state.
func (msdar *MonitorSmartDetectorAlertRule) ImportState(av io.Reader) error {
	msdar.state = &monitorSmartDetectorAlertRuleState{}
	if err := json.NewDecoder(av).Decode(msdar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", msdar.Type(), msdar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorSmartDetectorAlertRule] has state.
func (msdar *MonitorSmartDetectorAlertRule) State() (*monitorSmartDetectorAlertRuleState, bool) {
	return msdar.state, msdar.state != nil
}

// StateMust returns the state for [MonitorSmartDetectorAlertRule]. Panics if the state is nil.
func (msdar *MonitorSmartDetectorAlertRule) StateMust() *monitorSmartDetectorAlertRuleState {
	if msdar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", msdar.Type(), msdar.LocalName()))
	}
	return msdar.state
}

// MonitorSmartDetectorAlertRuleArgs contains the configurations for azurerm_monitor_smart_detector_alert_rule.
type MonitorSmartDetectorAlertRuleArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DetectorType: string, required
	DetectorType terra.StringValue `hcl:"detector_type,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Frequency: string, required
	Frequency terra.StringValue `hcl:"frequency,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ScopeResourceIds: set of string, required
	ScopeResourceIds terra.SetValue[terra.StringValue] `hcl:"scope_resource_ids,attr" validate:"required"`
	// Severity: string, required
	Severity terra.StringValue `hcl:"severity,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ThrottlingDuration: string, optional
	ThrottlingDuration terra.StringValue `hcl:"throttling_duration,attr"`
	// ActionGroup: required
	ActionGroup *monitorsmartdetectoralertrule.ActionGroup `hcl:"action_group,block" validate:"required"`
	// Timeouts: optional
	Timeouts *monitorsmartdetectoralertrule.Timeouts `hcl:"timeouts,block"`
}
type monitorSmartDetectorAlertRuleAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_monitor_smart_detector_alert_rule.
func (msdar monitorSmartDetectorAlertRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(msdar.ref.Append("description"))
}

// DetectorType returns a reference to field detector_type of azurerm_monitor_smart_detector_alert_rule.
func (msdar monitorSmartDetectorAlertRuleAttributes) DetectorType() terra.StringValue {
	return terra.ReferenceAsString(msdar.ref.Append("detector_type"))
}

// Enabled returns a reference to field enabled of azurerm_monitor_smart_detector_alert_rule.
func (msdar monitorSmartDetectorAlertRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(msdar.ref.Append("enabled"))
}

// Frequency returns a reference to field frequency of azurerm_monitor_smart_detector_alert_rule.
func (msdar monitorSmartDetectorAlertRuleAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(msdar.ref.Append("frequency"))
}

// Id returns a reference to field id of azurerm_monitor_smart_detector_alert_rule.
func (msdar monitorSmartDetectorAlertRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msdar.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_monitor_smart_detector_alert_rule.
func (msdar monitorSmartDetectorAlertRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(msdar.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_smart_detector_alert_rule.
func (msdar monitorSmartDetectorAlertRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(msdar.ref.Append("resource_group_name"))
}

// ScopeResourceIds returns a reference to field scope_resource_ids of azurerm_monitor_smart_detector_alert_rule.
func (msdar monitorSmartDetectorAlertRuleAttributes) ScopeResourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](msdar.ref.Append("scope_resource_ids"))
}

// Severity returns a reference to field severity of azurerm_monitor_smart_detector_alert_rule.
func (msdar monitorSmartDetectorAlertRuleAttributes) Severity() terra.StringValue {
	return terra.ReferenceAsString(msdar.ref.Append("severity"))
}

// Tags returns a reference to field tags of azurerm_monitor_smart_detector_alert_rule.
func (msdar monitorSmartDetectorAlertRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](msdar.ref.Append("tags"))
}

// ThrottlingDuration returns a reference to field throttling_duration of azurerm_monitor_smart_detector_alert_rule.
func (msdar monitorSmartDetectorAlertRuleAttributes) ThrottlingDuration() terra.StringValue {
	return terra.ReferenceAsString(msdar.ref.Append("throttling_duration"))
}

func (msdar monitorSmartDetectorAlertRuleAttributes) ActionGroup() terra.ListValue[monitorsmartdetectoralertrule.ActionGroupAttributes] {
	return terra.ReferenceAsList[monitorsmartdetectoralertrule.ActionGroupAttributes](msdar.ref.Append("action_group"))
}

func (msdar monitorSmartDetectorAlertRuleAttributes) Timeouts() monitorsmartdetectoralertrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitorsmartdetectoralertrule.TimeoutsAttributes](msdar.ref.Append("timeouts"))
}

type monitorSmartDetectorAlertRuleState struct {
	Description        string                                           `json:"description"`
	DetectorType       string                                           `json:"detector_type"`
	Enabled            bool                                             `json:"enabled"`
	Frequency          string                                           `json:"frequency"`
	Id                 string                                           `json:"id"`
	Name               string                                           `json:"name"`
	ResourceGroupName  string                                           `json:"resource_group_name"`
	ScopeResourceIds   []string                                         `json:"scope_resource_ids"`
	Severity           string                                           `json:"severity"`
	Tags               map[string]string                                `json:"tags"`
	ThrottlingDuration string                                           `json:"throttling_duration"`
	ActionGroup        []monitorsmartdetectoralertrule.ActionGroupState `json:"action_group"`
	Timeouts           *monitorsmartdetectoralertrule.TimeoutsState     `json:"timeouts"`
}
