// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataprocjob "github.com/golingon/terraproviders/google/4.59.0/dataprocjob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataprocJob(name string, args DataprocJobArgs) *DataprocJob {
	return &DataprocJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocJob)(nil)

type DataprocJob struct {
	Name  string
	Args  DataprocJobArgs
	state *dataprocJobState
}

func (dj *DataprocJob) Type() string {
	return "google_dataproc_job"
}

func (dj *DataprocJob) LocalName() string {
	return dj.Name
}

func (dj *DataprocJob) Configuration() interface{} {
	return dj.Args
}

func (dj *DataprocJob) Attributes() dataprocJobAttributes {
	return dataprocJobAttributes{ref: terra.ReferenceResource(dj)}
}

func (dj *DataprocJob) ImportState(av io.Reader) error {
	dj.state = &dataprocJobState{}
	if err := json.NewDecoder(av).Decode(dj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dj.Type(), dj.LocalName(), err)
	}
	return nil
}

func (dj *DataprocJob) State() (*dataprocJobState, bool) {
	return dj.state, dj.state != nil
}

func (dj *DataprocJob) StateMust() *dataprocJobState {
	if dj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dj.Type(), dj.LocalName()))
	}
	return dj.state
}

func (dj *DataprocJob) DependOn() terra.Reference {
	return terra.ReferenceResource(dj)
}

type DataprocJobArgs struct {
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
	// Status: min=0
	Status []dataprocjob.Status `hcl:"status,block" validate:"min=0"`
	// HadoopConfig: optional
	HadoopConfig *dataprocjob.HadoopConfig `hcl:"hadoop_config,block"`
	// HiveConfig: optional
	HiveConfig *dataprocjob.HiveConfig `hcl:"hive_config,block"`
	// PigConfig: optional
	PigConfig *dataprocjob.PigConfig `hcl:"pig_config,block"`
	// Placement: required
	Placement *dataprocjob.Placement `hcl:"placement,block" validate:"required"`
	// PrestoConfig: optional
	PrestoConfig *dataprocjob.PrestoConfig `hcl:"presto_config,block"`
	// PysparkConfig: optional
	PysparkConfig *dataprocjob.PysparkConfig `hcl:"pyspark_config,block"`
	// Reference: optional
	Reference *dataprocjob.Reference `hcl:"reference,block"`
	// Scheduling: optional
	Scheduling *dataprocjob.Scheduling `hcl:"scheduling,block"`
	// SparkConfig: optional
	SparkConfig *dataprocjob.SparkConfig `hcl:"spark_config,block"`
	// SparksqlConfig: optional
	SparksqlConfig *dataprocjob.SparksqlConfig `hcl:"sparksql_config,block"`
	// Timeouts: optional
	Timeouts *dataprocjob.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that DataprocJob depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataprocJobAttributes struct {
	ref terra.Reference
}

func (dj dataprocJobAttributes) DriverControlsFilesUri() terra.StringValue {
	return terra.ReferenceString(dj.ref.Append("driver_controls_files_uri"))
}

func (dj dataprocJobAttributes) DriverOutputResourceUri() terra.StringValue {
	return terra.ReferenceString(dj.ref.Append("driver_output_resource_uri"))
}

func (dj dataprocJobAttributes) ForceDelete() terra.BoolValue {
	return terra.ReferenceBool(dj.ref.Append("force_delete"))
}

func (dj dataprocJobAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dj.ref.Append("id"))
}

func (dj dataprocJobAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dj.ref.Append("labels"))
}

func (dj dataprocJobAttributes) Project() terra.StringValue {
	return terra.ReferenceString(dj.ref.Append("project"))
}

func (dj dataprocJobAttributes) Region() terra.StringValue {
	return terra.ReferenceString(dj.ref.Append("region"))
}

func (dj dataprocJobAttributes) Status() terra.ListValue[dataprocjob.StatusAttributes] {
	return terra.ReferenceList[dataprocjob.StatusAttributes](dj.ref.Append("status"))
}

func (dj dataprocJobAttributes) HadoopConfig() terra.ListValue[dataprocjob.HadoopConfigAttributes] {
	return terra.ReferenceList[dataprocjob.HadoopConfigAttributes](dj.ref.Append("hadoop_config"))
}

func (dj dataprocJobAttributes) HiveConfig() terra.ListValue[dataprocjob.HiveConfigAttributes] {
	return terra.ReferenceList[dataprocjob.HiveConfigAttributes](dj.ref.Append("hive_config"))
}

func (dj dataprocJobAttributes) PigConfig() terra.ListValue[dataprocjob.PigConfigAttributes] {
	return terra.ReferenceList[dataprocjob.PigConfigAttributes](dj.ref.Append("pig_config"))
}

func (dj dataprocJobAttributes) Placement() terra.ListValue[dataprocjob.PlacementAttributes] {
	return terra.ReferenceList[dataprocjob.PlacementAttributes](dj.ref.Append("placement"))
}

func (dj dataprocJobAttributes) PrestoConfig() terra.ListValue[dataprocjob.PrestoConfigAttributes] {
	return terra.ReferenceList[dataprocjob.PrestoConfigAttributes](dj.ref.Append("presto_config"))
}

func (dj dataprocJobAttributes) PysparkConfig() terra.ListValue[dataprocjob.PysparkConfigAttributes] {
	return terra.ReferenceList[dataprocjob.PysparkConfigAttributes](dj.ref.Append("pyspark_config"))
}

func (dj dataprocJobAttributes) Reference() terra.ListValue[dataprocjob.ReferenceAttributes] {
	return terra.ReferenceList[dataprocjob.ReferenceAttributes](dj.ref.Append("reference"))
}

func (dj dataprocJobAttributes) Scheduling() terra.ListValue[dataprocjob.SchedulingAttributes] {
	return terra.ReferenceList[dataprocjob.SchedulingAttributes](dj.ref.Append("scheduling"))
}

func (dj dataprocJobAttributes) SparkConfig() terra.ListValue[dataprocjob.SparkConfigAttributes] {
	return terra.ReferenceList[dataprocjob.SparkConfigAttributes](dj.ref.Append("spark_config"))
}

func (dj dataprocJobAttributes) SparksqlConfig() terra.ListValue[dataprocjob.SparksqlConfigAttributes] {
	return terra.ReferenceList[dataprocjob.SparksqlConfigAttributes](dj.ref.Append("sparksql_config"))
}

func (dj dataprocJobAttributes) Timeouts() dataprocjob.TimeoutsAttributes {
	return terra.ReferenceSingle[dataprocjob.TimeoutsAttributes](dj.ref.Append("timeouts"))
}

type dataprocJobState struct {
	DriverControlsFilesUri  string                            `json:"driver_controls_files_uri"`
	DriverOutputResourceUri string                            `json:"driver_output_resource_uri"`
	ForceDelete             bool                              `json:"force_delete"`
	Id                      string                            `json:"id"`
	Labels                  map[string]string                 `json:"labels"`
	Project                 string                            `json:"project"`
	Region                  string                            `json:"region"`
	Status                  []dataprocjob.StatusState         `json:"status"`
	HadoopConfig            []dataprocjob.HadoopConfigState   `json:"hadoop_config"`
	HiveConfig              []dataprocjob.HiveConfigState     `json:"hive_config"`
	PigConfig               []dataprocjob.PigConfigState      `json:"pig_config"`
	Placement               []dataprocjob.PlacementState      `json:"placement"`
	PrestoConfig            []dataprocjob.PrestoConfigState   `json:"presto_config"`
	PysparkConfig           []dataprocjob.PysparkConfigState  `json:"pyspark_config"`
	Reference               []dataprocjob.ReferenceState      `json:"reference"`
	Scheduling              []dataprocjob.SchedulingState     `json:"scheduling"`
	SparkConfig             []dataprocjob.SparkConfigState    `json:"spark_config"`
	SparksqlConfig          []dataprocjob.SparksqlConfigState `json:"sparksql_config"`
	Timeouts                *dataprocjob.TimeoutsState        `json:"timeouts"`
}
