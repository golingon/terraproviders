// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataproc_job

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

// Resource represents the Terraform resource google_dataproc_job.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataprocJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdj *Resource) Type() string {
	return "google_dataproc_job"
}

// LocalName returns the local name for [Resource].
func (gdj *Resource) LocalName() string {
	return gdj.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdj *Resource) Configuration() interface{} {
	return gdj.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdj *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdj)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdj *Resource) Dependencies() terra.Dependencies {
	return gdj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdj *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdj.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdj *Resource) Attributes() googleDataprocJobAttributes {
	return googleDataprocJobAttributes{ref: terra.ReferenceResource(gdj)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdj *Resource) ImportState(state io.Reader) error {
	gdj.state = &googleDataprocJobState{}
	if err := json.NewDecoder(state).Decode(gdj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdj.Type(), gdj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdj *Resource) State() (*googleDataprocJobState, bool) {
	return gdj.state, gdj.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdj *Resource) StateMust() *googleDataprocJobState {
	if gdj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdj.Type(), gdj.LocalName()))
	}
	return gdj.state
}

// Args contains the configurations for google_dataproc_job.
type Args struct {
	// ForceDelete: bool, optional
	ForceDelete terra.BoolValue `hcl:"force_delete,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// HadoopConfig: optional
	HadoopConfig *HadoopConfig `hcl:"hadoop_config,block"`
	// HiveConfig: optional
	HiveConfig *HiveConfig `hcl:"hive_config,block"`
	// PigConfig: optional
	PigConfig *PigConfig `hcl:"pig_config,block"`
	// Placement: required
	Placement *Placement `hcl:"placement,block" validate:"required"`
	// PrestoConfig: optional
	PrestoConfig *PrestoConfig `hcl:"presto_config,block"`
	// PysparkConfig: optional
	PysparkConfig *PysparkConfig `hcl:"pyspark_config,block"`
	// Reference: optional
	Reference *Reference `hcl:"reference,block"`
	// Scheduling: optional
	Scheduling *Scheduling `hcl:"scheduling,block"`
	// SparkConfig: optional
	SparkConfig *SparkConfig `hcl:"spark_config,block"`
	// SparksqlConfig: optional
	SparksqlConfig *SparksqlConfig `hcl:"sparksql_config,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleDataprocJobAttributes struct {
	ref terra.Reference
}

// DriverControlsFilesUri returns a reference to field driver_controls_files_uri of google_dataproc_job.
func (gdj googleDataprocJobAttributes) DriverControlsFilesUri() terra.StringValue {
	return terra.ReferenceAsString(gdj.ref.Append("driver_controls_files_uri"))
}

// DriverOutputResourceUri returns a reference to field driver_output_resource_uri of google_dataproc_job.
func (gdj googleDataprocJobAttributes) DriverOutputResourceUri() terra.StringValue {
	return terra.ReferenceAsString(gdj.ref.Append("driver_output_resource_uri"))
}

// EffectiveLabels returns a reference to field effective_labels of google_dataproc_job.
func (gdj googleDataprocJobAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gdj.ref.Append("effective_labels"))
}

// ForceDelete returns a reference to field force_delete of google_dataproc_job.
func (gdj googleDataprocJobAttributes) ForceDelete() terra.BoolValue {
	return terra.ReferenceAsBool(gdj.ref.Append("force_delete"))
}

// Id returns a reference to field id of google_dataproc_job.
func (gdj googleDataprocJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdj.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dataproc_job.
func (gdj googleDataprocJobAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gdj.ref.Append("labels"))
}

// Project returns a reference to field project of google_dataproc_job.
func (gdj googleDataprocJobAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdj.ref.Append("project"))
}

// Region returns a reference to field region of google_dataproc_job.
func (gdj googleDataprocJobAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gdj.ref.Append("region"))
}

// TerraformLabels returns a reference to field terraform_labels of google_dataproc_job.
func (gdj googleDataprocJobAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gdj.ref.Append("terraform_labels"))
}

func (gdj googleDataprocJobAttributes) Status() terra.ListValue[StatusAttributes] {
	return terra.ReferenceAsList[StatusAttributes](gdj.ref.Append("status"))
}

func (gdj googleDataprocJobAttributes) HadoopConfig() terra.ListValue[HadoopConfigAttributes] {
	return terra.ReferenceAsList[HadoopConfigAttributes](gdj.ref.Append("hadoop_config"))
}

func (gdj googleDataprocJobAttributes) HiveConfig() terra.ListValue[HiveConfigAttributes] {
	return terra.ReferenceAsList[HiveConfigAttributes](gdj.ref.Append("hive_config"))
}

func (gdj googleDataprocJobAttributes) PigConfig() terra.ListValue[PigConfigAttributes] {
	return terra.ReferenceAsList[PigConfigAttributes](gdj.ref.Append("pig_config"))
}

func (gdj googleDataprocJobAttributes) Placement() terra.ListValue[PlacementAttributes] {
	return terra.ReferenceAsList[PlacementAttributes](gdj.ref.Append("placement"))
}

func (gdj googleDataprocJobAttributes) PrestoConfig() terra.ListValue[PrestoConfigAttributes] {
	return terra.ReferenceAsList[PrestoConfigAttributes](gdj.ref.Append("presto_config"))
}

func (gdj googleDataprocJobAttributes) PysparkConfig() terra.ListValue[PysparkConfigAttributes] {
	return terra.ReferenceAsList[PysparkConfigAttributes](gdj.ref.Append("pyspark_config"))
}

func (gdj googleDataprocJobAttributes) Reference() terra.ListValue[ReferenceAttributes] {
	return terra.ReferenceAsList[ReferenceAttributes](gdj.ref.Append("reference"))
}

func (gdj googleDataprocJobAttributes) Scheduling() terra.ListValue[SchedulingAttributes] {
	return terra.ReferenceAsList[SchedulingAttributes](gdj.ref.Append("scheduling"))
}

func (gdj googleDataprocJobAttributes) SparkConfig() terra.ListValue[SparkConfigAttributes] {
	return terra.ReferenceAsList[SparkConfigAttributes](gdj.ref.Append("spark_config"))
}

func (gdj googleDataprocJobAttributes) SparksqlConfig() terra.ListValue[SparksqlConfigAttributes] {
	return terra.ReferenceAsList[SparksqlConfigAttributes](gdj.ref.Append("sparksql_config"))
}

func (gdj googleDataprocJobAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gdj.ref.Append("timeouts"))
}

type googleDataprocJobState struct {
	DriverControlsFilesUri  string                `json:"driver_controls_files_uri"`
	DriverOutputResourceUri string                `json:"driver_output_resource_uri"`
	EffectiveLabels         map[string]string     `json:"effective_labels"`
	ForceDelete             bool                  `json:"force_delete"`
	Id                      string                `json:"id"`
	Labels                  map[string]string     `json:"labels"`
	Project                 string                `json:"project"`
	Region                  string                `json:"region"`
	TerraformLabels         map[string]string     `json:"terraform_labels"`
	Status                  []StatusState         `json:"status"`
	HadoopConfig            []HadoopConfigState   `json:"hadoop_config"`
	HiveConfig              []HiveConfigState     `json:"hive_config"`
	PigConfig               []PigConfigState      `json:"pig_config"`
	Placement               []PlacementState      `json:"placement"`
	PrestoConfig            []PrestoConfigState   `json:"presto_config"`
	PysparkConfig           []PysparkConfigState  `json:"pyspark_config"`
	Reference               []ReferenceState      `json:"reference"`
	Scheduling              []SchedulingState     `json:"scheduling"`
	SparkConfig             []SparkConfigState    `json:"spark_config"`
	SparksqlConfig          []SparksqlConfigState `json:"sparksql_config"`
	Timeouts                *TimeoutsState        `json:"timeouts"`
}
