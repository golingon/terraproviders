// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_data_pipeline_pipeline

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

// Resource represents the Terraform resource google_data_pipeline_pipeline.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataPipelinePipelineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdpp *Resource) Type() string {
	return "google_data_pipeline_pipeline"
}

// LocalName returns the local name for [Resource].
func (gdpp *Resource) LocalName() string {
	return gdpp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdpp *Resource) Configuration() interface{} {
	return gdpp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdpp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdpp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdpp *Resource) Dependencies() terra.Dependencies {
	return gdpp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdpp *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdpp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdpp *Resource) Attributes() googleDataPipelinePipelineAttributes {
	return googleDataPipelinePipelineAttributes{ref: terra.ReferenceResource(gdpp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdpp *Resource) ImportState(state io.Reader) error {
	gdpp.state = &googleDataPipelinePipelineState{}
	if err := json.NewDecoder(state).Decode(gdpp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdpp.Type(), gdpp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdpp *Resource) State() (*googleDataPipelinePipelineState, bool) {
	return gdpp.state, gdpp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdpp *Resource) StateMust() *googleDataPipelinePipelineState {
	if gdpp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdpp.Type(), gdpp.LocalName()))
	}
	return gdpp.state
}

// Args contains the configurations for google_data_pipeline_pipeline.
type Args struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PipelineSources: map of string, optional
	PipelineSources terra.MapValue[terra.StringValue] `hcl:"pipeline_sources,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// SchedulerServiceAccountEmail: string, optional
	SchedulerServiceAccountEmail terra.StringValue `hcl:"scheduler_service_account_email,attr"`
	// State: string, required
	State terra.StringValue `hcl:"state,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// ScheduleInfo: optional
	ScheduleInfo *ScheduleInfo `hcl:"schedule_info,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// Workload: optional
	Workload *Workload `hcl:"workload,block"`
}

type googleDataPipelinePipelineAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gdpp.ref.Append("create_time"))
}

// DisplayName returns a reference to field display_name of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gdpp.ref.Append("display_name"))
}

// Id returns a reference to field id of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdpp.ref.Append("id"))
}

// JobCount returns a reference to field job_count of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) JobCount() terra.NumberValue {
	return terra.ReferenceAsNumber(gdpp.ref.Append("job_count"))
}

// LastUpdateTime returns a reference to field last_update_time of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) LastUpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gdpp.ref.Append("last_update_time"))
}

// Name returns a reference to field name of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gdpp.ref.Append("name"))
}

// PipelineSources returns a reference to field pipeline_sources of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) PipelineSources() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gdpp.ref.Append("pipeline_sources"))
}

// Project returns a reference to field project of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdpp.ref.Append("project"))
}

// Region returns a reference to field region of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gdpp.ref.Append("region"))
}

// SchedulerServiceAccountEmail returns a reference to field scheduler_service_account_email of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) SchedulerServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(gdpp.ref.Append("scheduler_service_account_email"))
}

// State returns a reference to field state of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gdpp.ref.Append("state"))
}

// Type returns a reference to field type of google_data_pipeline_pipeline.
func (gdpp googleDataPipelinePipelineAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(gdpp.ref.Append("type"))
}

func (gdpp googleDataPipelinePipelineAttributes) ScheduleInfo() terra.ListValue[ScheduleInfoAttributes] {
	return terra.ReferenceAsList[ScheduleInfoAttributes](gdpp.ref.Append("schedule_info"))
}

func (gdpp googleDataPipelinePipelineAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gdpp.ref.Append("timeouts"))
}

func (gdpp googleDataPipelinePipelineAttributes) Workload() terra.ListValue[WorkloadAttributes] {
	return terra.ReferenceAsList[WorkloadAttributes](gdpp.ref.Append("workload"))
}

type googleDataPipelinePipelineState struct {
	CreateTime                   string              `json:"create_time"`
	DisplayName                  string              `json:"display_name"`
	Id                           string              `json:"id"`
	JobCount                     float64             `json:"job_count"`
	LastUpdateTime               string              `json:"last_update_time"`
	Name                         string              `json:"name"`
	PipelineSources              map[string]string   `json:"pipeline_sources"`
	Project                      string              `json:"project"`
	Region                       string              `json:"region"`
	SchedulerServiceAccountEmail string              `json:"scheduler_service_account_email"`
	State                        string              `json:"state"`
	Type                         string              `json:"type"`
	ScheduleInfo                 []ScheduleInfoState `json:"schedule_info"`
	Timeouts                     *TimeoutsState      `json:"timeouts"`
	Workload                     []WorkloadState     `json:"workload"`
}
