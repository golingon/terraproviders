// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitorautoscalesetting "github.com/golingon/terraproviders/azurerm/3.49.0/monitorautoscalesetting"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewMonitorAutoscaleSetting(name string, args MonitorAutoscaleSettingArgs) *MonitorAutoscaleSetting {
	return &MonitorAutoscaleSetting{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorAutoscaleSetting)(nil)

type MonitorAutoscaleSetting struct {
	Name  string
	Args  MonitorAutoscaleSettingArgs
	state *monitorAutoscaleSettingState
}

func (mas *MonitorAutoscaleSetting) Type() string {
	return "azurerm_monitor_autoscale_setting"
}

func (mas *MonitorAutoscaleSetting) LocalName() string {
	return mas.Name
}

func (mas *MonitorAutoscaleSetting) Configuration() interface{} {
	return mas.Args
}

func (mas *MonitorAutoscaleSetting) Attributes() monitorAutoscaleSettingAttributes {
	return monitorAutoscaleSettingAttributes{ref: terra.ReferenceResource(mas)}
}

func (mas *MonitorAutoscaleSetting) ImportState(av io.Reader) error {
	mas.state = &monitorAutoscaleSettingState{}
	if err := json.NewDecoder(av).Decode(mas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mas.Type(), mas.LocalName(), err)
	}
	return nil
}

func (mas *MonitorAutoscaleSetting) State() (*monitorAutoscaleSettingState, bool) {
	return mas.state, mas.state != nil
}

func (mas *MonitorAutoscaleSetting) StateMust() *monitorAutoscaleSettingState {
	if mas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mas.Type(), mas.LocalName()))
	}
	return mas.state
}

func (mas *MonitorAutoscaleSetting) DependOn() terra.Reference {
	return terra.ReferenceResource(mas)
}

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
	// Profile: min=1,max=20
	Profile []monitorautoscalesetting.Profile `hcl:"profile,block" validate:"min=1,max=20"`
	// Timeouts: optional
	Timeouts *monitorautoscalesetting.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that MonitorAutoscaleSetting depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type monitorAutoscaleSettingAttributes struct {
	ref terra.Reference
}

func (mas monitorAutoscaleSettingAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(mas.ref.Append("enabled"))
}

func (mas monitorAutoscaleSettingAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mas.ref.Append("id"))
}

func (mas monitorAutoscaleSettingAttributes) Location() terra.StringValue {
	return terra.ReferenceString(mas.ref.Append("location"))
}

func (mas monitorAutoscaleSettingAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mas.ref.Append("name"))
}

func (mas monitorAutoscaleSettingAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(mas.ref.Append("resource_group_name"))
}

func (mas monitorAutoscaleSettingAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](mas.ref.Append("tags"))
}

func (mas monitorAutoscaleSettingAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceString(mas.ref.Append("target_resource_id"))
}

func (mas monitorAutoscaleSettingAttributes) Notification() terra.ListValue[monitorautoscalesetting.NotificationAttributes] {
	return terra.ReferenceList[monitorautoscalesetting.NotificationAttributes](mas.ref.Append("notification"))
}

func (mas monitorAutoscaleSettingAttributes) Profile() terra.ListValue[monitorautoscalesetting.ProfileAttributes] {
	return terra.ReferenceList[monitorautoscalesetting.ProfileAttributes](mas.ref.Append("profile"))
}

func (mas monitorAutoscaleSettingAttributes) Timeouts() monitorautoscalesetting.TimeoutsAttributes {
	return terra.ReferenceSingle[monitorautoscalesetting.TimeoutsAttributes](mas.ref.Append("timeouts"))
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
	Profile           []monitorautoscalesetting.ProfileState      `json:"profile"`
	Timeouts          *monitorautoscalesetting.TimeoutsState      `json:"timeouts"`
}
