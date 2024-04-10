// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	rumappmonitor "github.com/golingon/terraproviders/aws/5.44.0/rumappmonitor"
	"io"
)

// NewRumAppMonitor creates a new instance of [RumAppMonitor].
func NewRumAppMonitor(name string, args RumAppMonitorArgs) *RumAppMonitor {
	return &RumAppMonitor{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RumAppMonitor)(nil)

// RumAppMonitor represents the Terraform resource aws_rum_app_monitor.
type RumAppMonitor struct {
	Name      string
	Args      RumAppMonitorArgs
	state     *rumAppMonitorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RumAppMonitor].
func (ram *RumAppMonitor) Type() string {
	return "aws_rum_app_monitor"
}

// LocalName returns the local name for [RumAppMonitor].
func (ram *RumAppMonitor) LocalName() string {
	return ram.Name
}

// Configuration returns the configuration (args) for [RumAppMonitor].
func (ram *RumAppMonitor) Configuration() interface{} {
	return ram.Args
}

// DependOn is used for other resources to depend on [RumAppMonitor].
func (ram *RumAppMonitor) DependOn() terra.Reference {
	return terra.ReferenceResource(ram)
}

// Dependencies returns the list of resources [RumAppMonitor] depends_on.
func (ram *RumAppMonitor) Dependencies() terra.Dependencies {
	return ram.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RumAppMonitor].
func (ram *RumAppMonitor) LifecycleManagement() *terra.Lifecycle {
	return ram.Lifecycle
}

// Attributes returns the attributes for [RumAppMonitor].
func (ram *RumAppMonitor) Attributes() rumAppMonitorAttributes {
	return rumAppMonitorAttributes{ref: terra.ReferenceResource(ram)}
}

// ImportState imports the given attribute values into [RumAppMonitor]'s state.
func (ram *RumAppMonitor) ImportState(av io.Reader) error {
	ram.state = &rumAppMonitorState{}
	if err := json.NewDecoder(av).Decode(ram.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ram.Type(), ram.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RumAppMonitor] has state.
func (ram *RumAppMonitor) State() (*rumAppMonitorState, bool) {
	return ram.state, ram.state != nil
}

// StateMust returns the state for [RumAppMonitor]. Panics if the state is nil.
func (ram *RumAppMonitor) StateMust() *rumAppMonitorState {
	if ram.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ram.Type(), ram.LocalName()))
	}
	return ram.state
}

// RumAppMonitorArgs contains the configurations for aws_rum_app_monitor.
type RumAppMonitorArgs struct {
	// CwLogEnabled: bool, optional
	CwLogEnabled terra.BoolValue `hcl:"cw_log_enabled,attr"`
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AppMonitorConfiguration: optional
	AppMonitorConfiguration *rumappmonitor.AppMonitorConfiguration `hcl:"app_monitor_configuration,block"`
	// CustomEvents: optional
	CustomEvents *rumappmonitor.CustomEvents `hcl:"custom_events,block"`
}
type rumAppMonitorAttributes struct {
	ref terra.Reference
}

// AppMonitorId returns a reference to field app_monitor_id of aws_rum_app_monitor.
func (ram rumAppMonitorAttributes) AppMonitorId() terra.StringValue {
	return terra.ReferenceAsString(ram.ref.Append("app_monitor_id"))
}

// Arn returns a reference to field arn of aws_rum_app_monitor.
func (ram rumAppMonitorAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ram.ref.Append("arn"))
}

// CwLogEnabled returns a reference to field cw_log_enabled of aws_rum_app_monitor.
func (ram rumAppMonitorAttributes) CwLogEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ram.ref.Append("cw_log_enabled"))
}

// CwLogGroup returns a reference to field cw_log_group of aws_rum_app_monitor.
func (ram rumAppMonitorAttributes) CwLogGroup() terra.StringValue {
	return terra.ReferenceAsString(ram.ref.Append("cw_log_group"))
}

// Domain returns a reference to field domain of aws_rum_app_monitor.
func (ram rumAppMonitorAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(ram.ref.Append("domain"))
}

// Id returns a reference to field id of aws_rum_app_monitor.
func (ram rumAppMonitorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ram.ref.Append("id"))
}

// Name returns a reference to field name of aws_rum_app_monitor.
func (ram rumAppMonitorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ram.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_rum_app_monitor.
func (ram rumAppMonitorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ram.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_rum_app_monitor.
func (ram rumAppMonitorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ram.ref.Append("tags_all"))
}

func (ram rumAppMonitorAttributes) AppMonitorConfiguration() terra.ListValue[rumappmonitor.AppMonitorConfigurationAttributes] {
	return terra.ReferenceAsList[rumappmonitor.AppMonitorConfigurationAttributes](ram.ref.Append("app_monitor_configuration"))
}

func (ram rumAppMonitorAttributes) CustomEvents() terra.ListValue[rumappmonitor.CustomEventsAttributes] {
	return terra.ReferenceAsList[rumappmonitor.CustomEventsAttributes](ram.ref.Append("custom_events"))
}

type rumAppMonitorState struct {
	AppMonitorId            string                                       `json:"app_monitor_id"`
	Arn                     string                                       `json:"arn"`
	CwLogEnabled            bool                                         `json:"cw_log_enabled"`
	CwLogGroup              string                                       `json:"cw_log_group"`
	Domain                  string                                       `json:"domain"`
	Id                      string                                       `json:"id"`
	Name                    string                                       `json:"name"`
	Tags                    map[string]string                            `json:"tags"`
	TagsAll                 map[string]string                            `json:"tags_all"`
	AppMonitorConfiguration []rumappmonitor.AppMonitorConfigurationState `json:"app_monitor_configuration"`
	CustomEvents            []rumappmonitor.CustomEventsState            `json:"custom_events"`
}
