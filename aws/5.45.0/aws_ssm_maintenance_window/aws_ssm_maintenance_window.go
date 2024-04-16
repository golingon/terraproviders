// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ssm_maintenance_window

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_ssm_maintenance_window.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSsmMaintenanceWindowState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asmw *Resource) Type() string {
	return "aws_ssm_maintenance_window"
}

// LocalName returns the local name for [Resource].
func (asmw *Resource) LocalName() string {
	return asmw.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asmw *Resource) Configuration() interface{} {
	return asmw.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asmw *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asmw)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asmw *Resource) Dependencies() terra.Dependencies {
	return asmw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asmw *Resource) LifecycleManagement() *terra.Lifecycle {
	return asmw.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asmw *Resource) Attributes() awsSsmMaintenanceWindowAttributes {
	return awsSsmMaintenanceWindowAttributes{ref: terra.ReferenceResource(asmw)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asmw *Resource) ImportState(state io.Reader) error {
	asmw.state = &awsSsmMaintenanceWindowState{}
	if err := json.NewDecoder(state).Decode(asmw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asmw.Type(), asmw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asmw *Resource) State() (*awsSsmMaintenanceWindowState, bool) {
	return asmw.state, asmw.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asmw *Resource) StateMust() *awsSsmMaintenanceWindowState {
	if asmw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asmw.Type(), asmw.LocalName()))
	}
	return asmw.state
}

// Args contains the configurations for aws_ssm_maintenance_window.
type Args struct {
	// AllowUnassociatedTargets: bool, optional
	AllowUnassociatedTargets terra.BoolValue `hcl:"allow_unassociated_targets,attr"`
	// Cutoff: number, required
	Cutoff terra.NumberValue `hcl:"cutoff,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Duration: number, required
	Duration terra.NumberValue `hcl:"duration,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EndDate: string, optional
	EndDate terra.StringValue `hcl:"end_date,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Schedule: string, required
	Schedule terra.StringValue `hcl:"schedule,attr" validate:"required"`
	// ScheduleOffset: number, optional
	ScheduleOffset terra.NumberValue `hcl:"schedule_offset,attr"`
	// ScheduleTimezone: string, optional
	ScheduleTimezone terra.StringValue `hcl:"schedule_timezone,attr"`
	// StartDate: string, optional
	StartDate terra.StringValue `hcl:"start_date,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsSsmMaintenanceWindowAttributes struct {
	ref terra.Reference
}

// AllowUnassociatedTargets returns a reference to field allow_unassociated_targets of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) AllowUnassociatedTargets() terra.BoolValue {
	return terra.ReferenceAsBool(asmw.ref.Append("allow_unassociated_targets"))
}

// Cutoff returns a reference to field cutoff of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) Cutoff() terra.NumberValue {
	return terra.ReferenceAsNumber(asmw.ref.Append("cutoff"))
}

// Description returns a reference to field description of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(asmw.ref.Append("description"))
}

// Duration returns a reference to field duration of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) Duration() terra.NumberValue {
	return terra.ReferenceAsNumber(asmw.ref.Append("duration"))
}

// Enabled returns a reference to field enabled of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(asmw.ref.Append("enabled"))
}

// EndDate returns a reference to field end_date of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) EndDate() terra.StringValue {
	return terra.ReferenceAsString(asmw.ref.Append("end_date"))
}

// Id returns a reference to field id of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asmw.ref.Append("id"))
}

// Name returns a reference to field name of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asmw.ref.Append("name"))
}

// Schedule returns a reference to field schedule of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) Schedule() terra.StringValue {
	return terra.ReferenceAsString(asmw.ref.Append("schedule"))
}

// ScheduleOffset returns a reference to field schedule_offset of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) ScheduleOffset() terra.NumberValue {
	return terra.ReferenceAsNumber(asmw.ref.Append("schedule_offset"))
}

// ScheduleTimezone returns a reference to field schedule_timezone of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) ScheduleTimezone() terra.StringValue {
	return terra.ReferenceAsString(asmw.ref.Append("schedule_timezone"))
}

// StartDate returns a reference to field start_date of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) StartDate() terra.StringValue {
	return terra.ReferenceAsString(asmw.ref.Append("start_date"))
}

// Tags returns a reference to field tags of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asmw.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ssm_maintenance_window.
func (asmw awsSsmMaintenanceWindowAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asmw.ref.Append("tags_all"))
}

type awsSsmMaintenanceWindowState struct {
	AllowUnassociatedTargets bool              `json:"allow_unassociated_targets"`
	Cutoff                   float64           `json:"cutoff"`
	Description              string            `json:"description"`
	Duration                 float64           `json:"duration"`
	Enabled                  bool              `json:"enabled"`
	EndDate                  string            `json:"end_date"`
	Id                       string            `json:"id"`
	Name                     string            `json:"name"`
	Schedule                 string            `json:"schedule"`
	ScheduleOffset           float64           `json:"schedule_offset"`
	ScheduleTimezone         string            `json:"schedule_timezone"`
	StartDate                string            `json:"start_date"`
	Tags                     map[string]string `json:"tags"`
	TagsAll                  map[string]string `json:"tags_all"`
}
