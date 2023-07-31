// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitorautoscalesetting "github.com/golingon/terraproviders/azurerm/3.67.0/monitorautoscalesetting"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorAutoscaleSetting creates a new instance of [MonitorAutoscaleSetting].
func NewMonitorAutoscaleSetting(name string, args MonitorAutoscaleSettingArgs) *MonitorAutoscaleSetting {
	return &MonitorAutoscaleSetting{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorAutoscaleSetting)(nil)

// MonitorAutoscaleSetting represents the Terraform resource azurerm_monitor_autoscale_setting.
type MonitorAutoscaleSetting struct {
	Name      string
	Args      MonitorAutoscaleSettingArgs
	state     *monitorAutoscaleSettingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorAutoscaleSetting].
func (mas *MonitorAutoscaleSetting) Type() string {
	return "azurerm_monitor_autoscale_setting"
}

// LocalName returns the local name for [MonitorAutoscaleSetting].
func (mas *MonitorAutoscaleSetting) LocalName() string {
	return mas.Name
}

// Configuration returns the configuration (args) for [MonitorAutoscaleSetting].
func (mas *MonitorAutoscaleSetting) Configuration() interface{} {
	return mas.Args
}

// DependOn is used for other resources to depend on [MonitorAutoscaleSetting].
func (mas *MonitorAutoscaleSetting) DependOn() terra.Reference {
	return terra.ReferenceResource(mas)
}

// Dependencies returns the list of resources [MonitorAutoscaleSetting] depends_on.
func (mas *MonitorAutoscaleSetting) Dependencies() terra.Dependencies {
	return mas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorAutoscaleSetting].
func (mas *MonitorAutoscaleSetting) LifecycleManagement() *terra.Lifecycle {
	return mas.Lifecycle
}

// Attributes returns the attributes for [MonitorAutoscaleSetting].
func (mas *MonitorAutoscaleSetting) Attributes() monitorAutoscaleSettingAttributes {
	return monitorAutoscaleSettingAttributes{ref: terra.ReferenceResource(mas)}
}

// ImportState imports the given attribute values into [MonitorAutoscaleSetting]'s state.
func (mas *MonitorAutoscaleSetting) ImportState(av io.Reader) error {
	mas.state = &monitorAutoscaleSettingState{}
	if err := json.NewDecoder(av).Decode(mas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mas.Type(), mas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorAutoscaleSetting] has state.
func (mas *MonitorAutoscaleSetting) State() (*monitorAutoscaleSettingState, bool) {
	return mas.state, mas.state != nil
}

// StateMust returns the state for [MonitorAutoscaleSetting]. Panics if the state is nil.
func (mas *MonitorAutoscaleSetting) StateMust() *monitorAutoscaleSettingState {
	if mas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mas.Type(), mas.LocalName()))
	}
	return mas.state
}

// MonitorAutoscaleSettingArgs contains the configurations for azurerm_monitor_autoscale_setting.
type MonitorAutoscaleSettingArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// Notification: optional
	Notification *monitorautoscalesetting.Notification `hcl:"notification,block"`
	// Predictive: optional
	Predictive *monitorautoscalesetting.Predictive `hcl:"predictive,block"`
	// Profile: min=1,max=20
	Profile []monitorautoscalesetting.Profile `hcl:"profile,block" validate:"min=1,max=20"`
	// Timeouts: optional
	Timeouts *monitorautoscalesetting.Timeouts `hcl:"timeouts,block"`
}
type monitorAutoscaleSettingAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_monitor_autoscale_setting.
func (mas monitorAutoscaleSettingAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(mas.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_monitor_autoscale_setting.
func (mas monitorAutoscaleSettingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mas.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_monitor_autoscale_setting.
func (mas monitorAutoscaleSettingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mas.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_monitor_autoscale_setting.
func (mas monitorAutoscaleSettingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mas.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_autoscale_setting.
func (mas monitorAutoscaleSettingAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mas.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_monitor_autoscale_setting.
func (mas monitorAutoscaleSettingAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mas.ref.Append("tags"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_monitor_autoscale_setting.
func (mas monitorAutoscaleSettingAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(mas.ref.Append("target_resource_id"))
}

func (mas monitorAutoscaleSettingAttributes) Notification() terra.ListValue[monitorautoscalesetting.NotificationAttributes] {
	return terra.ReferenceAsList[monitorautoscalesetting.NotificationAttributes](mas.ref.Append("notification"))
}

func (mas monitorAutoscaleSettingAttributes) Predictive() terra.ListValue[monitorautoscalesetting.PredictiveAttributes] {
	return terra.ReferenceAsList[monitorautoscalesetting.PredictiveAttributes](mas.ref.Append("predictive"))
}

func (mas monitorAutoscaleSettingAttributes) Profile() terra.ListValue[monitorautoscalesetting.ProfileAttributes] {
	return terra.ReferenceAsList[monitorautoscalesetting.ProfileAttributes](mas.ref.Append("profile"))
}

func (mas monitorAutoscaleSettingAttributes) Timeouts() monitorautoscalesetting.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitorautoscalesetting.TimeoutsAttributes](mas.ref.Append("timeouts"))
}

type monitorAutoscaleSettingState struct {
	Enabled           bool                                        `json:"enabled"`
	Id                string                                      `json:"id"`
	Location          string                                      `json:"location"`
	Name              string                                      `json:"name"`
	ResourceGroupName string                                      `json:"resource_group_name"`
	Tags              map[string]string                           `json:"tags"`
	TargetResourceId  string                                      `json:"target_resource_id"`
	Notification      []monitorautoscalesetting.NotificationState `json:"notification"`
	Predictive        []monitorautoscalesetting.PredictiveState   `json:"predictive"`
	Profile           []monitorautoscalesetting.ProfileState      `json:"profile"`
	Timeouts          *monitorautoscalesetting.TimeoutsState      `json:"timeouts"`
}
