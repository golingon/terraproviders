// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	monitoringalertpolicy "github.com/golingon/terraproviders/googlebeta/4.72.1/monitoringalertpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitoringAlertPolicy creates a new instance of [MonitoringAlertPolicy].
func NewMonitoringAlertPolicy(name string, args MonitoringAlertPolicyArgs) *MonitoringAlertPolicy {
	return &MonitoringAlertPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitoringAlertPolicy)(nil)

// MonitoringAlertPolicy represents the Terraform resource google_monitoring_alert_policy.
type MonitoringAlertPolicy struct {
	Name      string
	Args      MonitoringAlertPolicyArgs
	state     *monitoringAlertPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitoringAlertPolicy].
func (_map *MonitoringAlertPolicy) Type() string {
	return "google_monitoring_alert_policy"
}

// LocalName returns the local name for [MonitoringAlertPolicy].
func (_map *MonitoringAlertPolicy) LocalName() string {
	return _map.Name
}

// Configuration returns the configuration (args) for [MonitoringAlertPolicy].
func (_map *MonitoringAlertPolicy) Configuration() interface{} {
	return _map.Args
}

// DependOn is used for other resources to depend on [MonitoringAlertPolicy].
func (_map *MonitoringAlertPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(_map)
}

// Dependencies returns the list of resources [MonitoringAlertPolicy] depends_on.
func (_map *MonitoringAlertPolicy) Dependencies() terra.Dependencies {
	return _map.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitoringAlertPolicy].
func (_map *MonitoringAlertPolicy) LifecycleManagement() *terra.Lifecycle {
	return _map.Lifecycle
}

// Attributes returns the attributes for [MonitoringAlertPolicy].
func (_map *MonitoringAlertPolicy) Attributes() monitoringAlertPolicyAttributes {
	return monitoringAlertPolicyAttributes{ref: terra.ReferenceResource(_map)}
}

// ImportState imports the given attribute values into [MonitoringAlertPolicy]'s state.
func (_map *MonitoringAlertPolicy) ImportState(av io.Reader) error {
	_map.state = &monitoringAlertPolicyState{}
	if err := json.NewDecoder(av).Decode(_map.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", _map.Type(), _map.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitoringAlertPolicy] has state.
func (_map *MonitoringAlertPolicy) State() (*monitoringAlertPolicyState, bool) {
	return _map.state, _map.state != nil
}

// StateMust returns the state for [MonitoringAlertPolicy]. Panics if the state is nil.
func (_map *MonitoringAlertPolicy) StateMust() *monitoringAlertPolicyState {
	if _map.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", _map.Type(), _map.LocalName()))
	}
	return _map.state
}

// MonitoringAlertPolicyArgs contains the configurations for google_monitoring_alert_policy.
type MonitoringAlertPolicyArgs struct {
	// Combiner: string, required
	Combiner terra.StringValue `hcl:"combiner,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NotificationChannels: list of string, optional
	NotificationChannels terra.ListValue[terra.StringValue] `hcl:"notification_channels,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// UserLabels: map of string, optional
	UserLabels terra.MapValue[terra.StringValue] `hcl:"user_labels,attr"`
	// CreationRecord: min=0
	CreationRecord []monitoringalertpolicy.CreationRecord `hcl:"creation_record,block" validate:"min=0"`
	// AlertStrategy: optional
	AlertStrategy *monitoringalertpolicy.AlertStrategy `hcl:"alert_strategy,block"`
	// Conditions: min=1
	Conditions []monitoringalertpolicy.Conditions `hcl:"conditions,block" validate:"min=1"`
	// Documentation: optional
	Documentation *monitoringalertpolicy.Documentation `hcl:"documentation,block"`
	// Timeouts: optional
	Timeouts *monitoringalertpolicy.Timeouts `hcl:"timeouts,block"`
}
type monitoringAlertPolicyAttributes struct {
	ref terra.Reference
}

// Combiner returns a reference to field combiner of google_monitoring_alert_policy.
func (_map monitoringAlertPolicyAttributes) Combiner() terra.StringValue {
	return terra.ReferenceAsString(_map.ref.Append("combiner"))
}

// DisplayName returns a reference to field display_name of google_monitoring_alert_policy.
func (_map monitoringAlertPolicyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(_map.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of google_monitoring_alert_policy.
func (_map monitoringAlertPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(_map.ref.Append("enabled"))
}

// Id returns a reference to field id of google_monitoring_alert_policy.
func (_map monitoringAlertPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(_map.ref.Append("id"))
}

// Name returns a reference to field name of google_monitoring_alert_policy.
func (_map monitoringAlertPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(_map.ref.Append("name"))
}

// NotificationChannels returns a reference to field notification_channels of google_monitoring_alert_policy.
func (_map monitoringAlertPolicyAttributes) NotificationChannels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](_map.ref.Append("notification_channels"))
}

// Project returns a reference to field project of google_monitoring_alert_policy.
func (_map monitoringAlertPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(_map.ref.Append("project"))
}

// UserLabels returns a reference to field user_labels of google_monitoring_alert_policy.
func (_map monitoringAlertPolicyAttributes) UserLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](_map.ref.Append("user_labels"))
}

func (_map monitoringAlertPolicyAttributes) CreationRecord() terra.ListValue[monitoringalertpolicy.CreationRecordAttributes] {
	return terra.ReferenceAsList[monitoringalertpolicy.CreationRecordAttributes](_map.ref.Append("creation_record"))
}

func (_map monitoringAlertPolicyAttributes) AlertStrategy() terra.ListValue[monitoringalertpolicy.AlertStrategyAttributes] {
	return terra.ReferenceAsList[monitoringalertpolicy.AlertStrategyAttributes](_map.ref.Append("alert_strategy"))
}

func (_map monitoringAlertPolicyAttributes) Conditions() terra.ListValue[monitoringalertpolicy.ConditionsAttributes] {
	return terra.ReferenceAsList[monitoringalertpolicy.ConditionsAttributes](_map.ref.Append("conditions"))
}

func (_map monitoringAlertPolicyAttributes) Documentation() terra.ListValue[monitoringalertpolicy.DocumentationAttributes] {
	return terra.ReferenceAsList[monitoringalertpolicy.DocumentationAttributes](_map.ref.Append("documentation"))
}

func (_map monitoringAlertPolicyAttributes) Timeouts() monitoringalertpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoringalertpolicy.TimeoutsAttributes](_map.ref.Append("timeouts"))
}

type monitoringAlertPolicyState struct {
	Combiner             string                                      `json:"combiner"`
	DisplayName          string                                      `json:"display_name"`
	Enabled              bool                                        `json:"enabled"`
	Id                   string                                      `json:"id"`
	Name                 string                                      `json:"name"`
	NotificationChannels []string                                    `json:"notification_channels"`
	Project              string                                      `json:"project"`
	UserLabels           map[string]string                           `json:"user_labels"`
	CreationRecord       []monitoringalertpolicy.CreationRecordState `json:"creation_record"`
	AlertStrategy        []monitoringalertpolicy.AlertStrategyState  `json:"alert_strategy"`
	Conditions           []monitoringalertpolicy.ConditionsState     `json:"conditions"`
	Documentation        []monitoringalertpolicy.DocumentationState  `json:"documentation"`
	Timeouts             *monitoringalertpolicy.TimeoutsState        `json:"timeouts"`
}
