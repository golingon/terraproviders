// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataform_repository_workflow_config

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

// Resource represents the Terraform resource google_dataform_repository_workflow_config.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataformRepositoryWorkflowConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdrwc *Resource) Type() string {
	return "google_dataform_repository_workflow_config"
}

// LocalName returns the local name for [Resource].
func (gdrwc *Resource) LocalName() string {
	return gdrwc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdrwc *Resource) Configuration() interface{} {
	return gdrwc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdrwc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdrwc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdrwc *Resource) Dependencies() terra.Dependencies {
	return gdrwc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdrwc *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdrwc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdrwc *Resource) Attributes() googleDataformRepositoryWorkflowConfigAttributes {
	return googleDataformRepositoryWorkflowConfigAttributes{ref: terra.ReferenceResource(gdrwc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdrwc *Resource) ImportState(state io.Reader) error {
	gdrwc.state = &googleDataformRepositoryWorkflowConfigState{}
	if err := json.NewDecoder(state).Decode(gdrwc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdrwc.Type(), gdrwc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdrwc *Resource) State() (*googleDataformRepositoryWorkflowConfigState, bool) {
	return gdrwc.state, gdrwc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdrwc *Resource) StateMust() *googleDataformRepositoryWorkflowConfigState {
	if gdrwc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdrwc.Type(), gdrwc.LocalName()))
	}
	return gdrwc.state
}

// Args contains the configurations for google_dataform_repository_workflow_config.
type Args struct {
	// CronSchedule: string, optional
	CronSchedule terra.StringValue `hcl:"cron_schedule,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ReleaseConfig: string, required
	ReleaseConfig terra.StringValue `hcl:"release_config,attr" validate:"required"`
	// Repository: string, optional
	Repository terra.StringValue `hcl:"repository,attr"`
	// TimeZone: string, optional
	TimeZone terra.StringValue `hcl:"time_zone,attr"`
	// InvocationConfig: optional
	InvocationConfig *InvocationConfig `hcl:"invocation_config,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleDataformRepositoryWorkflowConfigAttributes struct {
	ref terra.Reference
}

// CronSchedule returns a reference to field cron_schedule of google_dataform_repository_workflow_config.
func (gdrwc googleDataformRepositoryWorkflowConfigAttributes) CronSchedule() terra.StringValue {
	return terra.ReferenceAsString(gdrwc.ref.Append("cron_schedule"))
}

// Id returns a reference to field id of google_dataform_repository_workflow_config.
func (gdrwc googleDataformRepositoryWorkflowConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdrwc.ref.Append("id"))
}

// Name returns a reference to field name of google_dataform_repository_workflow_config.
func (gdrwc googleDataformRepositoryWorkflowConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gdrwc.ref.Append("name"))
}

// Project returns a reference to field project of google_dataform_repository_workflow_config.
func (gdrwc googleDataformRepositoryWorkflowConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdrwc.ref.Append("project"))
}

// Region returns a reference to field region of google_dataform_repository_workflow_config.
func (gdrwc googleDataformRepositoryWorkflowConfigAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gdrwc.ref.Append("region"))
}

// ReleaseConfig returns a reference to field release_config of google_dataform_repository_workflow_config.
func (gdrwc googleDataformRepositoryWorkflowConfigAttributes) ReleaseConfig() terra.StringValue {
	return terra.ReferenceAsString(gdrwc.ref.Append("release_config"))
}

// Repository returns a reference to field repository of google_dataform_repository_workflow_config.
func (gdrwc googleDataformRepositoryWorkflowConfigAttributes) Repository() terra.StringValue {
	return terra.ReferenceAsString(gdrwc.ref.Append("repository"))
}

// TimeZone returns a reference to field time_zone of google_dataform_repository_workflow_config.
func (gdrwc googleDataformRepositoryWorkflowConfigAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(gdrwc.ref.Append("time_zone"))
}

func (gdrwc googleDataformRepositoryWorkflowConfigAttributes) RecentScheduledExecutionRecords() terra.ListValue[RecentScheduledExecutionRecordsAttributes] {
	return terra.ReferenceAsList[RecentScheduledExecutionRecordsAttributes](gdrwc.ref.Append("recent_scheduled_execution_records"))
}

func (gdrwc googleDataformRepositoryWorkflowConfigAttributes) InvocationConfig() terra.ListValue[InvocationConfigAttributes] {
	return terra.ReferenceAsList[InvocationConfigAttributes](gdrwc.ref.Append("invocation_config"))
}

func (gdrwc googleDataformRepositoryWorkflowConfigAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gdrwc.ref.Append("timeouts"))
}

type googleDataformRepositoryWorkflowConfigState struct {
	CronSchedule                    string                                 `json:"cron_schedule"`
	Id                              string                                 `json:"id"`
	Name                            string                                 `json:"name"`
	Project                         string                                 `json:"project"`
	Region                          string                                 `json:"region"`
	ReleaseConfig                   string                                 `json:"release_config"`
	Repository                      string                                 `json:"repository"`
	TimeZone                        string                                 `json:"time_zone"`
	RecentScheduledExecutionRecords []RecentScheduledExecutionRecordsState `json:"recent_scheduled_execution_records"`
	InvocationConfig                []InvocationConfigState                `json:"invocation_config"`
	Timeouts                        *TimeoutsState                         `json:"timeouts"`
}
