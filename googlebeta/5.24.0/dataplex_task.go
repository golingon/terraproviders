// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	dataplextask "github.com/golingon/terraproviders/googlebeta/5.24.0/dataplextask"
	"io"
)

// NewDataplexTask creates a new instance of [DataplexTask].
func NewDataplexTask(name string, args DataplexTaskArgs) *DataplexTask {
	return &DataplexTask{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexTask)(nil)

// DataplexTask represents the Terraform resource google_dataplex_task.
type DataplexTask struct {
	Name      string
	Args      DataplexTaskArgs
	state     *dataplexTaskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexTask].
func (dt *DataplexTask) Type() string {
	return "google_dataplex_task"
}

// LocalName returns the local name for [DataplexTask].
func (dt *DataplexTask) LocalName() string {
	return dt.Name
}

// Configuration returns the configuration (args) for [DataplexTask].
func (dt *DataplexTask) Configuration() interface{} {
	return dt.Args
}

// DependOn is used for other resources to depend on [DataplexTask].
func (dt *DataplexTask) DependOn() terra.Reference {
	return terra.ReferenceResource(dt)
}

// Dependencies returns the list of resources [DataplexTask] depends_on.
func (dt *DataplexTask) Dependencies() terra.Dependencies {
	return dt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexTask].
func (dt *DataplexTask) LifecycleManagement() *terra.Lifecycle {
	return dt.Lifecycle
}

// Attributes returns the attributes for [DataplexTask].
func (dt *DataplexTask) Attributes() dataplexTaskAttributes {
	return dataplexTaskAttributes{ref: terra.ReferenceResource(dt)}
}

// ImportState imports the given attribute values into [DataplexTask]'s state.
func (dt *DataplexTask) ImportState(av io.Reader) error {
	dt.state = &dataplexTaskState{}
	if err := json.NewDecoder(av).Decode(dt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dt.Type(), dt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexTask] has state.
func (dt *DataplexTask) State() (*dataplexTaskState, bool) {
	return dt.state, dt.state != nil
}

// StateMust returns the state for [DataplexTask]. Panics if the state is nil.
func (dt *DataplexTask) StateMust() *dataplexTaskState {
	if dt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dt.Type(), dt.LocalName()))
	}
	return dt.state
}

// DataplexTaskArgs contains the configurations for google_dataplex_task.
type DataplexTaskArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Lake: string, optional
	Lake terra.StringValue `hcl:"lake,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TaskId: string, optional
	TaskId terra.StringValue `hcl:"task_id,attr"`
	// ExecutionStatus: min=0
	ExecutionStatus []dataplextask.ExecutionStatus `hcl:"execution_status,block" validate:"min=0"`
	// ExecutionSpec: required
	ExecutionSpec *dataplextask.ExecutionSpec `hcl:"execution_spec,block" validate:"required"`
	// Notebook: optional
	Notebook *dataplextask.Notebook `hcl:"notebook,block"`
	// Spark: optional
	Spark *dataplextask.Spark `hcl:"spark,block"`
	// Timeouts: optional
	Timeouts *dataplextask.Timeouts `hcl:"timeouts,block"`
	// TriggerSpec: required
	TriggerSpec *dataplextask.TriggerSpec `hcl:"trigger_spec,block" validate:"required"`
}
type dataplexTaskAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_dataplex_task.
func (dt dataplexTaskAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("create_time"))
}

// Description returns a reference to field description of google_dataplex_task.
func (dt dataplexTaskAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dataplex_task.
func (dt dataplexTaskAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("display_name"))
}

// EffectiveLabels returns a reference to field effective_labels of google_dataplex_task.
func (dt dataplexTaskAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dt.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_dataplex_task.
func (dt dataplexTaskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dataplex_task.
func (dt dataplexTaskAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dt.ref.Append("labels"))
}

// Lake returns a reference to field lake of google_dataplex_task.
func (dt dataplexTaskAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_task.
func (dt dataplexTaskAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("location"))
}

// Name returns a reference to field name of google_dataplex_task.
func (dt dataplexTaskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("name"))
}

// Project returns a reference to field project of google_dataplex_task.
func (dt dataplexTaskAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("project"))
}

// State returns a reference to field state of google_dataplex_task.
func (dt dataplexTaskAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("state"))
}

// TaskId returns a reference to field task_id of google_dataplex_task.
func (dt dataplexTaskAttributes) TaskId() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("task_id"))
}

// TerraformLabels returns a reference to field terraform_labels of google_dataplex_task.
func (dt dataplexTaskAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dt.ref.Append("terraform_labels"))
}

// Uid returns a reference to field uid of google_dataplex_task.
func (dt dataplexTaskAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_dataplex_task.
func (dt dataplexTaskAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("update_time"))
}

func (dt dataplexTaskAttributes) ExecutionStatus() terra.ListValue[dataplextask.ExecutionStatusAttributes] {
	return terra.ReferenceAsList[dataplextask.ExecutionStatusAttributes](dt.ref.Append("execution_status"))
}

func (dt dataplexTaskAttributes) ExecutionSpec() terra.ListValue[dataplextask.ExecutionSpecAttributes] {
	return terra.ReferenceAsList[dataplextask.ExecutionSpecAttributes](dt.ref.Append("execution_spec"))
}

func (dt dataplexTaskAttributes) Notebook() terra.ListValue[dataplextask.NotebookAttributes] {
	return terra.ReferenceAsList[dataplextask.NotebookAttributes](dt.ref.Append("notebook"))
}

func (dt dataplexTaskAttributes) Spark() terra.ListValue[dataplextask.SparkAttributes] {
	return terra.ReferenceAsList[dataplextask.SparkAttributes](dt.ref.Append("spark"))
}

func (dt dataplexTaskAttributes) Timeouts() dataplextask.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataplextask.TimeoutsAttributes](dt.ref.Append("timeouts"))
}

func (dt dataplexTaskAttributes) TriggerSpec() terra.ListValue[dataplextask.TriggerSpecAttributes] {
	return terra.ReferenceAsList[dataplextask.TriggerSpecAttributes](dt.ref.Append("trigger_spec"))
}

type dataplexTaskState struct {
	CreateTime      string                              `json:"create_time"`
	Description     string                              `json:"description"`
	DisplayName     string                              `json:"display_name"`
	EffectiveLabels map[string]string                   `json:"effective_labels"`
	Id              string                              `json:"id"`
	Labels          map[string]string                   `json:"labels"`
	Lake            string                              `json:"lake"`
	Location        string                              `json:"location"`
	Name            string                              `json:"name"`
	Project         string                              `json:"project"`
	State           string                              `json:"state"`
	TaskId          string                              `json:"task_id"`
	TerraformLabels map[string]string                   `json:"terraform_labels"`
	Uid             string                              `json:"uid"`
	UpdateTime      string                              `json:"update_time"`
	ExecutionStatus []dataplextask.ExecutionStatusState `json:"execution_status"`
	ExecutionSpec   []dataplextask.ExecutionSpecState   `json:"execution_spec"`
	Notebook        []dataplextask.NotebookState        `json:"notebook"`
	Spark           []dataplextask.SparkState           `json:"spark"`
	Timeouts        *dataplextask.TimeoutsState         `json:"timeouts"`
	TriggerSpec     []dataplextask.TriggerSpecState     `json:"trigger_spec"`
}
