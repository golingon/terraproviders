// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataprocworkflowtemplate "github.com/golingon/terraproviders/googlebeta/4.73.1/dataprocworkflowtemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocWorkflowTemplate creates a new instance of [DataprocWorkflowTemplate].
func NewDataprocWorkflowTemplate(name string, args DataprocWorkflowTemplateArgs) *DataprocWorkflowTemplate {
	return &DataprocWorkflowTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocWorkflowTemplate)(nil)

// DataprocWorkflowTemplate represents the Terraform resource google_dataproc_workflow_template.
type DataprocWorkflowTemplate struct {
	Name      string
	Args      DataprocWorkflowTemplateArgs
	state     *dataprocWorkflowTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocWorkflowTemplate].
func (dwt *DataprocWorkflowTemplate) Type() string {
	return "google_dataproc_workflow_template"
}

// LocalName returns the local name for [DataprocWorkflowTemplate].
func (dwt *DataprocWorkflowTemplate) LocalName() string {
	return dwt.Name
}

// Configuration returns the configuration (args) for [DataprocWorkflowTemplate].
func (dwt *DataprocWorkflowTemplate) Configuration() interface{} {
	return dwt.Args
}

// DependOn is used for other resources to depend on [DataprocWorkflowTemplate].
func (dwt *DataprocWorkflowTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(dwt)
}

// Dependencies returns the list of resources [DataprocWorkflowTemplate] depends_on.
func (dwt *DataprocWorkflowTemplate) Dependencies() terra.Dependencies {
	return dwt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocWorkflowTemplate].
func (dwt *DataprocWorkflowTemplate) LifecycleManagement() *terra.Lifecycle {
	return dwt.Lifecycle
}

// Attributes returns the attributes for [DataprocWorkflowTemplate].
func (dwt *DataprocWorkflowTemplate) Attributes() dataprocWorkflowTemplateAttributes {
	return dataprocWorkflowTemplateAttributes{ref: terra.ReferenceResource(dwt)}
}

// ImportState imports the given attribute values into [DataprocWorkflowTemplate]'s state.
func (dwt *DataprocWorkflowTemplate) ImportState(av io.Reader) error {
	dwt.state = &dataprocWorkflowTemplateState{}
	if err := json.NewDecoder(av).Decode(dwt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dwt.Type(), dwt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocWorkflowTemplate] has state.
func (dwt *DataprocWorkflowTemplate) State() (*dataprocWorkflowTemplateState, bool) {
	return dwt.state, dwt.state != nil
}

// StateMust returns the state for [DataprocWorkflowTemplate]. Panics if the state is nil.
func (dwt *DataprocWorkflowTemplate) StateMust() *dataprocWorkflowTemplateState {
	if dwt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dwt.Type(), dwt.LocalName()))
	}
	return dwt.state
}

// DataprocWorkflowTemplateArgs contains the configurations for google_dataproc_workflow_template.
type DataprocWorkflowTemplateArgs struct {
	// DagTimeout: string, optional
	DagTimeout terra.StringValue `hcl:"dag_timeout,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Version: number, optional
	Version terra.NumberValue `hcl:"version,attr"`
	// Jobs: min=1
	Jobs []dataprocworkflowtemplate.Jobs `hcl:"jobs,block" validate:"min=1"`
	// Parameters: min=0
	Parameters []dataprocworkflowtemplate.Parameters `hcl:"parameters,block" validate:"min=0"`
	// Placement: required
	Placement *dataprocworkflowtemplate.Placement `hcl:"placement,block" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprocworkflowtemplate.Timeouts `hcl:"timeouts,block"`
}
type dataprocWorkflowTemplateAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_dataproc_workflow_template.
func (dwt dataprocWorkflowTemplateAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(dwt.ref.Append("create_time"))
}

// DagTimeout returns a reference to field dag_timeout of google_dataproc_workflow_template.
func (dwt dataprocWorkflowTemplateAttributes) DagTimeout() terra.StringValue {
	return terra.ReferenceAsString(dwt.ref.Append("dag_timeout"))
}

// Id returns a reference to field id of google_dataproc_workflow_template.
func (dwt dataprocWorkflowTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dwt.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dataproc_workflow_template.
func (dwt dataprocWorkflowTemplateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dwt.ref.Append("labels"))
}

// Location returns a reference to field location of google_dataproc_workflow_template.
func (dwt dataprocWorkflowTemplateAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dwt.ref.Append("location"))
}

// Name returns a reference to field name of google_dataproc_workflow_template.
func (dwt dataprocWorkflowTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dwt.ref.Append("name"))
}

// Project returns a reference to field project of google_dataproc_workflow_template.
func (dwt dataprocWorkflowTemplateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dwt.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_dataproc_workflow_template.
func (dwt dataprocWorkflowTemplateAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(dwt.ref.Append("update_time"))
}

// Version returns a reference to field version of google_dataproc_workflow_template.
func (dwt dataprocWorkflowTemplateAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(dwt.ref.Append("version"))
}

func (dwt dataprocWorkflowTemplateAttributes) Jobs() terra.ListValue[dataprocworkflowtemplate.JobsAttributes] {
	return terra.ReferenceAsList[dataprocworkflowtemplate.JobsAttributes](dwt.ref.Append("jobs"))
}

func (dwt dataprocWorkflowTemplateAttributes) Parameters() terra.ListValue[dataprocworkflowtemplate.ParametersAttributes] {
	return terra.ReferenceAsList[dataprocworkflowtemplate.ParametersAttributes](dwt.ref.Append("parameters"))
}

func (dwt dataprocWorkflowTemplateAttributes) Placement() terra.ListValue[dataprocworkflowtemplate.PlacementAttributes] {
	return terra.ReferenceAsList[dataprocworkflowtemplate.PlacementAttributes](dwt.ref.Append("placement"))
}

func (dwt dataprocWorkflowTemplateAttributes) Timeouts() dataprocworkflowtemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprocworkflowtemplate.TimeoutsAttributes](dwt.ref.Append("timeouts"))
}

type dataprocWorkflowTemplateState struct {
	CreateTime string                                     `json:"create_time"`
	DagTimeout string                                     `json:"dag_timeout"`
	Id         string                                     `json:"id"`
	Labels     map[string]string                          `json:"labels"`
	Location   string                                     `json:"location"`
	Name       string                                     `json:"name"`
	Project    string                                     `json:"project"`
	UpdateTime string                                     `json:"update_time"`
	Version    float64                                    `json:"version"`
	Jobs       []dataprocworkflowtemplate.JobsState       `json:"jobs"`
	Parameters []dataprocworkflowtemplate.ParametersState `json:"parameters"`
	Placement  []dataprocworkflowtemplate.PlacementState  `json:"placement"`
	Timeouts   *dataprocworkflowtemplate.TimeoutsState    `json:"timeouts"`
}
