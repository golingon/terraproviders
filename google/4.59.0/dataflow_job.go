// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataflowjob "github.com/golingon/terraproviders/google/4.59.0/dataflowjob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataflowJob creates a new instance of [DataflowJob].
func NewDataflowJob(name string, args DataflowJobArgs) *DataflowJob {
	return &DataflowJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataflowJob)(nil)

// DataflowJob represents the Terraform resource google_dataflow_job.
type DataflowJob struct {
	Name      string
	Args      DataflowJobArgs
	state     *dataflowJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataflowJob].
func (dj *DataflowJob) Type() string {
	return "google_dataflow_job"
}

// LocalName returns the local name for [DataflowJob].
func (dj *DataflowJob) LocalName() string {
	return dj.Name
}

// Configuration returns the configuration (args) for [DataflowJob].
func (dj *DataflowJob) Configuration() interface{} {
	return dj.Args
}

// DependOn is used for other resources to depend on [DataflowJob].
func (dj *DataflowJob) DependOn() terra.Reference {
	return terra.ReferenceResource(dj)
}

// Dependencies returns the list of resources [DataflowJob] depends_on.
func (dj *DataflowJob) Dependencies() terra.Dependencies {
	return dj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataflowJob].
func (dj *DataflowJob) LifecycleManagement() *terra.Lifecycle {
	return dj.Lifecycle
}

// Attributes returns the attributes for [DataflowJob].
func (dj *DataflowJob) Attributes() dataflowJobAttributes {
	return dataflowJobAttributes{ref: terra.ReferenceResource(dj)}
}

// ImportState imports the given attribute values into [DataflowJob]'s state.
func (dj *DataflowJob) ImportState(av io.Reader) error {
	dj.state = &dataflowJobState{}
	if err := json.NewDecoder(av).Decode(dj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dj.Type(), dj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataflowJob] has state.
func (dj *DataflowJob) State() (*dataflowJobState, bool) {
	return dj.state, dj.state != nil
}

// StateMust returns the state for [DataflowJob]. Panics if the state is nil.
func (dj *DataflowJob) StateMust() *dataflowJobState {
	if dj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dj.Type(), dj.LocalName()))
	}
	return dj.state
}

// DataflowJobArgs contains the configurations for google_dataflow_job.
type DataflowJobArgs struct {
	// AdditionalExperiments: set of string, optional
	AdditionalExperiments terra.SetValue[terra.StringValue] `hcl:"additional_experiments,attr"`
	// EnableStreamingEngine: bool, optional
	EnableStreamingEngine terra.BoolValue `hcl:"enable_streaming_engine,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpConfiguration: string, optional
	IpConfiguration terra.StringValue `hcl:"ip_configuration,attr"`
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// MachineType: string, optional
	MachineType terra.StringValue `hcl:"machine_type,attr"`
	// MaxWorkers: number, optional
	MaxWorkers terra.NumberValue `hcl:"max_workers,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// OnDelete: string, optional
	OnDelete terra.StringValue `hcl:"on_delete,attr"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ServiceAccountEmail: string, optional
	ServiceAccountEmail terra.StringValue `hcl:"service_account_email,attr"`
	// SkipWaitOnJobTermination: bool, optional
	SkipWaitOnJobTermination terra.BoolValue `hcl:"skip_wait_on_job_termination,attr"`
	// Subnetwork: string, optional
	Subnetwork terra.StringValue `hcl:"subnetwork,attr"`
	// TempGcsLocation: string, required
	TempGcsLocation terra.StringValue `hcl:"temp_gcs_location,attr" validate:"required"`
	// TemplateGcsPath: string, required
	TemplateGcsPath terra.StringValue `hcl:"template_gcs_path,attr" validate:"required"`
	// TransformNameMapping: map of string, optional
	TransformNameMapping terra.MapValue[terra.StringValue] `hcl:"transform_name_mapping,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Timeouts: optional
	Timeouts *dataflowjob.Timeouts `hcl:"timeouts,block"`
}
type dataflowJobAttributes struct {
	ref terra.Reference
}

// AdditionalExperiments returns a reference to field additional_experiments of google_dataflow_job.
func (dj dataflowJobAttributes) AdditionalExperiments() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dj.ref.Append("additional_experiments"))
}

// EnableStreamingEngine returns a reference to field enable_streaming_engine of google_dataflow_job.
func (dj dataflowJobAttributes) EnableStreamingEngine() terra.BoolValue {
	return terra.ReferenceAsBool(dj.ref.Append("enable_streaming_engine"))
}

// Id returns a reference to field id of google_dataflow_job.
func (dj dataflowJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("id"))
}

// IpConfiguration returns a reference to field ip_configuration of google_dataflow_job.
func (dj dataflowJobAttributes) IpConfiguration() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("ip_configuration"))
}

// JobId returns a reference to field job_id of google_dataflow_job.
func (dj dataflowJobAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("job_id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_dataflow_job.
func (dj dataflowJobAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("kms_key_name"))
}

// Labels returns a reference to field labels of google_dataflow_job.
func (dj dataflowJobAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dj.ref.Append("labels"))
}

// MachineType returns a reference to field machine_type of google_dataflow_job.
func (dj dataflowJobAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("machine_type"))
}

// MaxWorkers returns a reference to field max_workers of google_dataflow_job.
func (dj dataflowJobAttributes) MaxWorkers() terra.NumberValue {
	return terra.ReferenceAsNumber(dj.ref.Append("max_workers"))
}

// Name returns a reference to field name of google_dataflow_job.
func (dj dataflowJobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("name"))
}

// Network returns a reference to field network of google_dataflow_job.
func (dj dataflowJobAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("network"))
}

// OnDelete returns a reference to field on_delete of google_dataflow_job.
func (dj dataflowJobAttributes) OnDelete() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("on_delete"))
}

// Parameters returns a reference to field parameters of google_dataflow_job.
func (dj dataflowJobAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dj.ref.Append("parameters"))
}

// Project returns a reference to field project of google_dataflow_job.
func (dj dataflowJobAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("project"))
}

// Region returns a reference to field region of google_dataflow_job.
func (dj dataflowJobAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("region"))
}

// ServiceAccountEmail returns a reference to field service_account_email of google_dataflow_job.
func (dj dataflowJobAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("service_account_email"))
}

// SkipWaitOnJobTermination returns a reference to field skip_wait_on_job_termination of google_dataflow_job.
func (dj dataflowJobAttributes) SkipWaitOnJobTermination() terra.BoolValue {
	return terra.ReferenceAsBool(dj.ref.Append("skip_wait_on_job_termination"))
}

// State returns a reference to field state of google_dataflow_job.
func (dj dataflowJobAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("state"))
}

// Subnetwork returns a reference to field subnetwork of google_dataflow_job.
func (dj dataflowJobAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("subnetwork"))
}

// TempGcsLocation returns a reference to field temp_gcs_location of google_dataflow_job.
func (dj dataflowJobAttributes) TempGcsLocation() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("temp_gcs_location"))
}

// TemplateGcsPath returns a reference to field template_gcs_path of google_dataflow_job.
func (dj dataflowJobAttributes) TemplateGcsPath() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("template_gcs_path"))
}

// TransformNameMapping returns a reference to field transform_name_mapping of google_dataflow_job.
func (dj dataflowJobAttributes) TransformNameMapping() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dj.ref.Append("transform_name_mapping"))
}

// Type returns a reference to field type of google_dataflow_job.
func (dj dataflowJobAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("type"))
}

// Zone returns a reference to field zone of google_dataflow_job.
func (dj dataflowJobAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(dj.ref.Append("zone"))
}

func (dj dataflowJobAttributes) Timeouts() dataflowjob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataflowjob.TimeoutsAttributes](dj.ref.Append("timeouts"))
}

type dataflowJobState struct {
	AdditionalExperiments    []string                   `json:"additional_experiments"`
	EnableStreamingEngine    bool                       `json:"enable_streaming_engine"`
	Id                       string                     `json:"id"`
	IpConfiguration          string                     `json:"ip_configuration"`
	JobId                    string                     `json:"job_id"`
	KmsKeyName               string                     `json:"kms_key_name"`
	Labels                   map[string]string          `json:"labels"`
	MachineType              string                     `json:"machine_type"`
	MaxWorkers               float64                    `json:"max_workers"`
	Name                     string                     `json:"name"`
	Network                  string                     `json:"network"`
	OnDelete                 string                     `json:"on_delete"`
	Parameters               map[string]string          `json:"parameters"`
	Project                  string                     `json:"project"`
	Region                   string                     `json:"region"`
	ServiceAccountEmail      string                     `json:"service_account_email"`
	SkipWaitOnJobTermination bool                       `json:"skip_wait_on_job_termination"`
	State                    string                     `json:"state"`
	Subnetwork               string                     `json:"subnetwork"`
	TempGcsLocation          string                     `json:"temp_gcs_location"`
	TemplateGcsPath          string                     `json:"template_gcs_path"`
	TransformNameMapping     map[string]string          `json:"transform_name_mapping"`
	Type                     string                     `json:"type"`
	Zone                     string                     `json:"zone"`
	Timeouts                 *dataflowjob.TimeoutsState `json:"timeouts"`
}
