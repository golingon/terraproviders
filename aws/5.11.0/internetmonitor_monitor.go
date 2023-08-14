// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	internetmonitormonitor "github.com/golingon/terraproviders/aws/5.11.0/internetmonitormonitor"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewInternetmonitorMonitor creates a new instance of [InternetmonitorMonitor].
func NewInternetmonitorMonitor(name string, args InternetmonitorMonitorArgs) *InternetmonitorMonitor {
	return &InternetmonitorMonitor{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*InternetmonitorMonitor)(nil)

// InternetmonitorMonitor represents the Terraform resource aws_internetmonitor_monitor.
type InternetmonitorMonitor struct {
	Name      string
	Args      InternetmonitorMonitorArgs
	state     *internetmonitorMonitorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [InternetmonitorMonitor].
func (im *InternetmonitorMonitor) Type() string {
	return "aws_internetmonitor_monitor"
}

// LocalName returns the local name for [InternetmonitorMonitor].
func (im *InternetmonitorMonitor) LocalName() string {
	return im.Name
}

// Configuration returns the configuration (args) for [InternetmonitorMonitor].
func (im *InternetmonitorMonitor) Configuration() interface{} {
	return im.Args
}

// DependOn is used for other resources to depend on [InternetmonitorMonitor].
func (im *InternetmonitorMonitor) DependOn() terra.Reference {
	return terra.ReferenceResource(im)
}

// Dependencies returns the list of resources [InternetmonitorMonitor] depends_on.
func (im *InternetmonitorMonitor) Dependencies() terra.Dependencies {
	return im.DependsOn
}

// LifecycleManagement returns the lifecycle block for [InternetmonitorMonitor].
func (im *InternetmonitorMonitor) LifecycleManagement() *terra.Lifecycle {
	return im.Lifecycle
}

// Attributes returns the attributes for [InternetmonitorMonitor].
func (im *InternetmonitorMonitor) Attributes() internetmonitorMonitorAttributes {
	return internetmonitorMonitorAttributes{ref: terra.ReferenceResource(im)}
}

// ImportState imports the given attribute values into [InternetmonitorMonitor]'s state.
func (im *InternetmonitorMonitor) ImportState(av io.Reader) error {
	im.state = &internetmonitorMonitorState{}
	if err := json.NewDecoder(av).Decode(im.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", im.Type(), im.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [InternetmonitorMonitor] has state.
func (im *InternetmonitorMonitor) State() (*internetmonitorMonitorState, bool) {
	return im.state, im.state != nil
}

// StateMust returns the state for [InternetmonitorMonitor]. Panics if the state is nil.
func (im *InternetmonitorMonitor) StateMust() *internetmonitorMonitorState {
	if im.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", im.Type(), im.LocalName()))
	}
	return im.state
}

// InternetmonitorMonitorArgs contains the configurations for aws_internetmonitor_monitor.
type InternetmonitorMonitorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxCityNetworksToMonitor: number, optional
	MaxCityNetworksToMonitor terra.NumberValue `hcl:"max_city_networks_to_monitor,attr"`
	// MonitorName: string, required
	MonitorName terra.StringValue `hcl:"monitor_name,attr" validate:"required"`
	// Resources: set of string, optional
	Resources terra.SetValue[terra.StringValue] `hcl:"resources,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TrafficPercentageToMonitor: number, optional
	TrafficPercentageToMonitor terra.NumberValue `hcl:"traffic_percentage_to_monitor,attr"`
	// HealthEventsConfig: optional
	HealthEventsConfig *internetmonitormonitor.HealthEventsConfig `hcl:"health_events_config,block"`
	// InternetMeasurementsLogDelivery: optional
	InternetMeasurementsLogDelivery *internetmonitormonitor.InternetMeasurementsLogDelivery `hcl:"internet_measurements_log_delivery,block"`
}
type internetmonitorMonitorAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_internetmonitor_monitor.
func (im internetmonitorMonitorAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(im.ref.Append("arn"))
}

// Id returns a reference to field id of aws_internetmonitor_monitor.
func (im internetmonitorMonitorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(im.ref.Append("id"))
}

// MaxCityNetworksToMonitor returns a reference to field max_city_networks_to_monitor of aws_internetmonitor_monitor.
func (im internetmonitorMonitorAttributes) MaxCityNetworksToMonitor() terra.NumberValue {
	return terra.ReferenceAsNumber(im.ref.Append("max_city_networks_to_monitor"))
}

// MonitorName returns a reference to field monitor_name of aws_internetmonitor_monitor.
func (im internetmonitorMonitorAttributes) MonitorName() terra.StringValue {
	return terra.ReferenceAsString(im.ref.Append("monitor_name"))
}

// Resources returns a reference to field resources of aws_internetmonitor_monitor.
func (im internetmonitorMonitorAttributes) Resources() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](im.ref.Append("resources"))
}

// Status returns a reference to field status of aws_internetmonitor_monitor.
func (im internetmonitorMonitorAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(im.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_internetmonitor_monitor.
func (im internetmonitorMonitorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](im.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_internetmonitor_monitor.
func (im internetmonitorMonitorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](im.ref.Append("tags_all"))
}

// TrafficPercentageToMonitor returns a reference to field traffic_percentage_to_monitor of aws_internetmonitor_monitor.
func (im internetmonitorMonitorAttributes) TrafficPercentageToMonitor() terra.NumberValue {
	return terra.ReferenceAsNumber(im.ref.Append("traffic_percentage_to_monitor"))
}

func (im internetmonitorMonitorAttributes) HealthEventsConfig() terra.ListValue[internetmonitormonitor.HealthEventsConfigAttributes] {
	return terra.ReferenceAsList[internetmonitormonitor.HealthEventsConfigAttributes](im.ref.Append("health_events_config"))
}

func (im internetmonitorMonitorAttributes) InternetMeasurementsLogDelivery() terra.ListValue[internetmonitormonitor.InternetMeasurementsLogDeliveryAttributes] {
	return terra.ReferenceAsList[internetmonitormonitor.InternetMeasurementsLogDeliveryAttributes](im.ref.Append("internet_measurements_log_delivery"))
}

type internetmonitorMonitorState struct {
	Arn                             string                                                        `json:"arn"`
	Id                              string                                                        `json:"id"`
	MaxCityNetworksToMonitor        float64                                                       `json:"max_city_networks_to_monitor"`
	MonitorName                     string                                                        `json:"monitor_name"`
	Resources                       []string                                                      `json:"resources"`
	Status                          string                                                        `json:"status"`
	Tags                            map[string]string                                             `json:"tags"`
	TagsAll                         map[string]string                                             `json:"tags_all"`
	TrafficPercentageToMonitor      float64                                                       `json:"traffic_percentage_to_monitor"`
	HealthEventsConfig              []internetmonitormonitor.HealthEventsConfigState              `json:"health_events_config"`
	InternetMeasurementsLogDelivery []internetmonitormonitor.InternetMeasurementsLogDeliveryState `json:"internet_measurements_log_delivery"`
}