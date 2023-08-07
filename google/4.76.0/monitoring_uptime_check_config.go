// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	monitoringuptimecheckconfig "github.com/golingon/terraproviders/google/4.76.0/monitoringuptimecheckconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitoringUptimeCheckConfig creates a new instance of [MonitoringUptimeCheckConfig].
func NewMonitoringUptimeCheckConfig(name string, args MonitoringUptimeCheckConfigArgs) *MonitoringUptimeCheckConfig {
	return &MonitoringUptimeCheckConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitoringUptimeCheckConfig)(nil)

// MonitoringUptimeCheckConfig represents the Terraform resource google_monitoring_uptime_check_config.
type MonitoringUptimeCheckConfig struct {
	Name      string
	Args      MonitoringUptimeCheckConfigArgs
	state     *monitoringUptimeCheckConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitoringUptimeCheckConfig].
func (mucc *MonitoringUptimeCheckConfig) Type() string {
	return "google_monitoring_uptime_check_config"
}

// LocalName returns the local name for [MonitoringUptimeCheckConfig].
func (mucc *MonitoringUptimeCheckConfig) LocalName() string {
	return mucc.Name
}

// Configuration returns the configuration (args) for [MonitoringUptimeCheckConfig].
func (mucc *MonitoringUptimeCheckConfig) Configuration() interface{} {
	return mucc.Args
}

// DependOn is used for other resources to depend on [MonitoringUptimeCheckConfig].
func (mucc *MonitoringUptimeCheckConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(mucc)
}

// Dependencies returns the list of resources [MonitoringUptimeCheckConfig] depends_on.
func (mucc *MonitoringUptimeCheckConfig) Dependencies() terra.Dependencies {
	return mucc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitoringUptimeCheckConfig].
func (mucc *MonitoringUptimeCheckConfig) LifecycleManagement() *terra.Lifecycle {
	return mucc.Lifecycle
}

// Attributes returns the attributes for [MonitoringUptimeCheckConfig].
func (mucc *MonitoringUptimeCheckConfig) Attributes() monitoringUptimeCheckConfigAttributes {
	return monitoringUptimeCheckConfigAttributes{ref: terra.ReferenceResource(mucc)}
}

// ImportState imports the given attribute values into [MonitoringUptimeCheckConfig]'s state.
func (mucc *MonitoringUptimeCheckConfig) ImportState(av io.Reader) error {
	mucc.state = &monitoringUptimeCheckConfigState{}
	if err := json.NewDecoder(av).Decode(mucc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mucc.Type(), mucc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitoringUptimeCheckConfig] has state.
func (mucc *MonitoringUptimeCheckConfig) State() (*monitoringUptimeCheckConfigState, bool) {
	return mucc.state, mucc.state != nil
}

// StateMust returns the state for [MonitoringUptimeCheckConfig]. Panics if the state is nil.
func (mucc *MonitoringUptimeCheckConfig) StateMust() *monitoringUptimeCheckConfigState {
	if mucc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mucc.Type(), mucc.LocalName()))
	}
	return mucc.state
}

// MonitoringUptimeCheckConfigArgs contains the configurations for google_monitoring_uptime_check_config.
type MonitoringUptimeCheckConfigArgs struct {
	// CheckerType: string, optional
	CheckerType terra.StringValue `hcl:"checker_type,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Period: string, optional
	Period terra.StringValue `hcl:"period,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SelectedRegions: list of string, optional
	SelectedRegions terra.ListValue[terra.StringValue] `hcl:"selected_regions,attr"`
	// Timeout: string, required
	Timeout terra.StringValue `hcl:"timeout,attr" validate:"required"`
	// ContentMatchers: min=0
	ContentMatchers []monitoringuptimecheckconfig.ContentMatchers `hcl:"content_matchers,block" validate:"min=0"`
	// HttpCheck: optional
	HttpCheck *monitoringuptimecheckconfig.HttpCheck `hcl:"http_check,block"`
	// MonitoredResource: optional
	MonitoredResource *monitoringuptimecheckconfig.MonitoredResource `hcl:"monitored_resource,block"`
	// ResourceGroup: optional
	ResourceGroup *monitoringuptimecheckconfig.ResourceGroup `hcl:"resource_group,block"`
	// TcpCheck: optional
	TcpCheck *monitoringuptimecheckconfig.TcpCheck `hcl:"tcp_check,block"`
	// Timeouts: optional
	Timeouts *monitoringuptimecheckconfig.Timeouts `hcl:"timeouts,block"`
}
type monitoringUptimeCheckConfigAttributes struct {
	ref terra.Reference
}

// CheckerType returns a reference to field checker_type of google_monitoring_uptime_check_config.
func (mucc monitoringUptimeCheckConfigAttributes) CheckerType() terra.StringValue {
	return terra.ReferenceAsString(mucc.ref.Append("checker_type"))
}

// DisplayName returns a reference to field display_name of google_monitoring_uptime_check_config.
func (mucc monitoringUptimeCheckConfigAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mucc.ref.Append("display_name"))
}

// Id returns a reference to field id of google_monitoring_uptime_check_config.
func (mucc monitoringUptimeCheckConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mucc.ref.Append("id"))
}

// Name returns a reference to field name of google_monitoring_uptime_check_config.
func (mucc monitoringUptimeCheckConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mucc.ref.Append("name"))
}

// Period returns a reference to field period of google_monitoring_uptime_check_config.
func (mucc monitoringUptimeCheckConfigAttributes) Period() terra.StringValue {
	return terra.ReferenceAsString(mucc.ref.Append("period"))
}

// Project returns a reference to field project of google_monitoring_uptime_check_config.
func (mucc monitoringUptimeCheckConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(mucc.ref.Append("project"))
}

// SelectedRegions returns a reference to field selected_regions of google_monitoring_uptime_check_config.
func (mucc monitoringUptimeCheckConfigAttributes) SelectedRegions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mucc.ref.Append("selected_regions"))
}

// Timeout returns a reference to field timeout of google_monitoring_uptime_check_config.
func (mucc monitoringUptimeCheckConfigAttributes) Timeout() terra.StringValue {
	return terra.ReferenceAsString(mucc.ref.Append("timeout"))
}

// UptimeCheckId returns a reference to field uptime_check_id of google_monitoring_uptime_check_config.
func (mucc monitoringUptimeCheckConfigAttributes) UptimeCheckId() terra.StringValue {
	return terra.ReferenceAsString(mucc.ref.Append("uptime_check_id"))
}

func (mucc monitoringUptimeCheckConfigAttributes) ContentMatchers() terra.ListValue[monitoringuptimecheckconfig.ContentMatchersAttributes] {
	return terra.ReferenceAsList[monitoringuptimecheckconfig.ContentMatchersAttributes](mucc.ref.Append("content_matchers"))
}

func (mucc monitoringUptimeCheckConfigAttributes) HttpCheck() terra.ListValue[monitoringuptimecheckconfig.HttpCheckAttributes] {
	return terra.ReferenceAsList[monitoringuptimecheckconfig.HttpCheckAttributes](mucc.ref.Append("http_check"))
}

func (mucc monitoringUptimeCheckConfigAttributes) MonitoredResource() terra.ListValue[monitoringuptimecheckconfig.MonitoredResourceAttributes] {
	return terra.ReferenceAsList[monitoringuptimecheckconfig.MonitoredResourceAttributes](mucc.ref.Append("monitored_resource"))
}

func (mucc monitoringUptimeCheckConfigAttributes) ResourceGroup() terra.ListValue[monitoringuptimecheckconfig.ResourceGroupAttributes] {
	return terra.ReferenceAsList[monitoringuptimecheckconfig.ResourceGroupAttributes](mucc.ref.Append("resource_group"))
}

func (mucc monitoringUptimeCheckConfigAttributes) TcpCheck() terra.ListValue[monitoringuptimecheckconfig.TcpCheckAttributes] {
	return terra.ReferenceAsList[monitoringuptimecheckconfig.TcpCheckAttributes](mucc.ref.Append("tcp_check"))
}

func (mucc monitoringUptimeCheckConfigAttributes) Timeouts() monitoringuptimecheckconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoringuptimecheckconfig.TimeoutsAttributes](mucc.ref.Append("timeouts"))
}

type monitoringUptimeCheckConfigState struct {
	CheckerType       string                                               `json:"checker_type"`
	DisplayName       string                                               `json:"display_name"`
	Id                string                                               `json:"id"`
	Name              string                                               `json:"name"`
	Period            string                                               `json:"period"`
	Project           string                                               `json:"project"`
	SelectedRegions   []string                                             `json:"selected_regions"`
	Timeout           string                                               `json:"timeout"`
	UptimeCheckId     string                                               `json:"uptime_check_id"`
	ContentMatchers   []monitoringuptimecheckconfig.ContentMatchersState   `json:"content_matchers"`
	HttpCheck         []monitoringuptimecheckconfig.HttpCheckState         `json:"http_check"`
	MonitoredResource []monitoringuptimecheckconfig.MonitoredResourceState `json:"monitored_resource"`
	ResourceGroup     []monitoringuptimecheckconfig.ResourceGroupState     `json:"resource_group"`
	TcpCheck          []monitoringuptimecheckconfig.TcpCheckState          `json:"tcp_check"`
	Timeouts          *monitoringuptimecheckconfig.TimeoutsState           `json:"timeouts"`
}
