// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	monitoringdashboard "github.com/golingon/terraproviders/google/5.24.0/monitoringdashboard"
	"io"
)

// NewMonitoringDashboard creates a new instance of [MonitoringDashboard].
func NewMonitoringDashboard(name string, args MonitoringDashboardArgs) *MonitoringDashboard {
	return &MonitoringDashboard{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitoringDashboard)(nil)

// MonitoringDashboard represents the Terraform resource google_monitoring_dashboard.
type MonitoringDashboard struct {
	Name      string
	Args      MonitoringDashboardArgs
	state     *monitoringDashboardState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitoringDashboard].
func (md *MonitoringDashboard) Type() string {
	return "google_monitoring_dashboard"
}

// LocalName returns the local name for [MonitoringDashboard].
func (md *MonitoringDashboard) LocalName() string {
	return md.Name
}

// Configuration returns the configuration (args) for [MonitoringDashboard].
func (md *MonitoringDashboard) Configuration() interface{} {
	return md.Args
}

// DependOn is used for other resources to depend on [MonitoringDashboard].
func (md *MonitoringDashboard) DependOn() terra.Reference {
	return terra.ReferenceResource(md)
}

// Dependencies returns the list of resources [MonitoringDashboard] depends_on.
func (md *MonitoringDashboard) Dependencies() terra.Dependencies {
	return md.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitoringDashboard].
func (md *MonitoringDashboard) LifecycleManagement() *terra.Lifecycle {
	return md.Lifecycle
}

// Attributes returns the attributes for [MonitoringDashboard].
func (md *MonitoringDashboard) Attributes() monitoringDashboardAttributes {
	return monitoringDashboardAttributes{ref: terra.ReferenceResource(md)}
}

// ImportState imports the given attribute values into [MonitoringDashboard]'s state.
func (md *MonitoringDashboard) ImportState(av io.Reader) error {
	md.state = &monitoringDashboardState{}
	if err := json.NewDecoder(av).Decode(md.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", md.Type(), md.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitoringDashboard] has state.
func (md *MonitoringDashboard) State() (*monitoringDashboardState, bool) {
	return md.state, md.state != nil
}

// StateMust returns the state for [MonitoringDashboard]. Panics if the state is nil.
func (md *MonitoringDashboard) StateMust() *monitoringDashboardState {
	if md.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", md.Type(), md.LocalName()))
	}
	return md.state
}

// MonitoringDashboardArgs contains the configurations for google_monitoring_dashboard.
type MonitoringDashboardArgs struct {
	// DashboardJson: string, required
	DashboardJson terra.StringValue `hcl:"dashboard_json,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *monitoringdashboard.Timeouts `hcl:"timeouts,block"`
}
type monitoringDashboardAttributes struct {
	ref terra.Reference
}

// DashboardJson returns a reference to field dashboard_json of google_monitoring_dashboard.
func (md monitoringDashboardAttributes) DashboardJson() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("dashboard_json"))
}

// Id returns a reference to field id of google_monitoring_dashboard.
func (md monitoringDashboardAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("id"))
}

// Project returns a reference to field project of google_monitoring_dashboard.
func (md monitoringDashboardAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("project"))
}

func (md monitoringDashboardAttributes) Timeouts() monitoringdashboard.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoringdashboard.TimeoutsAttributes](md.ref.Append("timeouts"))
}

type monitoringDashboardState struct {
	DashboardJson string                             `json:"dashboard_json"`
	Id            string                             `json:"id"`
	Project       string                             `json:"project"`
	Timeouts      *monitoringdashboard.TimeoutsState `json:"timeouts"`
}
