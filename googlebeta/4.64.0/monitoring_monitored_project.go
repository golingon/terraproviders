// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	monitoringmonitoredproject "github.com/golingon/terraproviders/googlebeta/4.64.0/monitoringmonitoredproject"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitoringMonitoredProject creates a new instance of [MonitoringMonitoredProject].
func NewMonitoringMonitoredProject(name string, args MonitoringMonitoredProjectArgs) *MonitoringMonitoredProject {
	return &MonitoringMonitoredProject{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitoringMonitoredProject)(nil)

// MonitoringMonitoredProject represents the Terraform resource google_monitoring_monitored_project.
type MonitoringMonitoredProject struct {
	Name      string
	Args      MonitoringMonitoredProjectArgs
	state     *monitoringMonitoredProjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitoringMonitoredProject].
func (mmp *MonitoringMonitoredProject) Type() string {
	return "google_monitoring_monitored_project"
}

// LocalName returns the local name for [MonitoringMonitoredProject].
func (mmp *MonitoringMonitoredProject) LocalName() string {
	return mmp.Name
}

// Configuration returns the configuration (args) for [MonitoringMonitoredProject].
func (mmp *MonitoringMonitoredProject) Configuration() interface{} {
	return mmp.Args
}

// DependOn is used for other resources to depend on [MonitoringMonitoredProject].
func (mmp *MonitoringMonitoredProject) DependOn() terra.Reference {
	return terra.ReferenceResource(mmp)
}

// Dependencies returns the list of resources [MonitoringMonitoredProject] depends_on.
func (mmp *MonitoringMonitoredProject) Dependencies() terra.Dependencies {
	return mmp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitoringMonitoredProject].
func (mmp *MonitoringMonitoredProject) LifecycleManagement() *terra.Lifecycle {
	return mmp.Lifecycle
}

// Attributes returns the attributes for [MonitoringMonitoredProject].
func (mmp *MonitoringMonitoredProject) Attributes() monitoringMonitoredProjectAttributes {
	return monitoringMonitoredProjectAttributes{ref: terra.ReferenceResource(mmp)}
}

// ImportState imports the given attribute values into [MonitoringMonitoredProject]'s state.
func (mmp *MonitoringMonitoredProject) ImportState(av io.Reader) error {
	mmp.state = &monitoringMonitoredProjectState{}
	if err := json.NewDecoder(av).Decode(mmp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mmp.Type(), mmp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitoringMonitoredProject] has state.
func (mmp *MonitoringMonitoredProject) State() (*monitoringMonitoredProjectState, bool) {
	return mmp.state, mmp.state != nil
}

// StateMust returns the state for [MonitoringMonitoredProject]. Panics if the state is nil.
func (mmp *MonitoringMonitoredProject) StateMust() *monitoringMonitoredProjectState {
	if mmp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mmp.Type(), mmp.LocalName()))
	}
	return mmp.state
}

// MonitoringMonitoredProjectArgs contains the configurations for google_monitoring_monitored_project.
type MonitoringMonitoredProjectArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MetricsScope: string, required
	MetricsScope terra.StringValue `hcl:"metrics_scope,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *monitoringmonitoredproject.Timeouts `hcl:"timeouts,block"`
}
type monitoringMonitoredProjectAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_monitoring_monitored_project.
func (mmp monitoringMonitoredProjectAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(mmp.ref.Append("create_time"))
}

// Id returns a reference to field id of google_monitoring_monitored_project.
func (mmp monitoringMonitoredProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mmp.ref.Append("id"))
}

// MetricsScope returns a reference to field metrics_scope of google_monitoring_monitored_project.
func (mmp monitoringMonitoredProjectAttributes) MetricsScope() terra.StringValue {
	return terra.ReferenceAsString(mmp.ref.Append("metrics_scope"))
}

// Name returns a reference to field name of google_monitoring_monitored_project.
func (mmp monitoringMonitoredProjectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mmp.ref.Append("name"))
}

func (mmp monitoringMonitoredProjectAttributes) Timeouts() monitoringmonitoredproject.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoringmonitoredproject.TimeoutsAttributes](mmp.ref.Append("timeouts"))
}

type monitoringMonitoredProjectState struct {
	CreateTime   string                                    `json:"create_time"`
	Id           string                                    `json:"id"`
	MetricsScope string                                    `json:"metrics_scope"`
	Name         string                                    `json:"name"`
	Timeouts     *monitoringmonitoredproject.TimeoutsState `json:"timeouts"`
}
