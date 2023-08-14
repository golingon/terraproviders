// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	workflowsworkflow "github.com/golingon/terraproviders/google/4.77.0/workflowsworkflow"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWorkflowsWorkflow creates a new instance of [WorkflowsWorkflow].
func NewWorkflowsWorkflow(name string, args WorkflowsWorkflowArgs) *WorkflowsWorkflow {
	return &WorkflowsWorkflow{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkflowsWorkflow)(nil)

// WorkflowsWorkflow represents the Terraform resource google_workflows_workflow.
type WorkflowsWorkflow struct {
	Name      string
	Args      WorkflowsWorkflowArgs
	state     *workflowsWorkflowState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkflowsWorkflow].
func (ww *WorkflowsWorkflow) Type() string {
	return "google_workflows_workflow"
}

// LocalName returns the local name for [WorkflowsWorkflow].
func (ww *WorkflowsWorkflow) LocalName() string {
	return ww.Name
}

// Configuration returns the configuration (args) for [WorkflowsWorkflow].
func (ww *WorkflowsWorkflow) Configuration() interface{} {
	return ww.Args
}

// DependOn is used for other resources to depend on [WorkflowsWorkflow].
func (ww *WorkflowsWorkflow) DependOn() terra.Reference {
	return terra.ReferenceResource(ww)
}

// Dependencies returns the list of resources [WorkflowsWorkflow] depends_on.
func (ww *WorkflowsWorkflow) Dependencies() terra.Dependencies {
	return ww.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkflowsWorkflow].
func (ww *WorkflowsWorkflow) LifecycleManagement() *terra.Lifecycle {
	return ww.Lifecycle
}

// Attributes returns the attributes for [WorkflowsWorkflow].
func (ww *WorkflowsWorkflow) Attributes() workflowsWorkflowAttributes {
	return workflowsWorkflowAttributes{ref: terra.ReferenceResource(ww)}
}

// ImportState imports the given attribute values into [WorkflowsWorkflow]'s state.
func (ww *WorkflowsWorkflow) ImportState(av io.Reader) error {
	ww.state = &workflowsWorkflowState{}
	if err := json.NewDecoder(av).Decode(ww.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ww.Type(), ww.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkflowsWorkflow] has state.
func (ww *WorkflowsWorkflow) State() (*workflowsWorkflowState, bool) {
	return ww.state, ww.state != nil
}

// StateMust returns the state for [WorkflowsWorkflow]. Panics if the state is nil.
func (ww *WorkflowsWorkflow) StateMust() *workflowsWorkflowState {
	if ww.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ww.Type(), ww.LocalName()))
	}
	return ww.state
}

// WorkflowsWorkflowArgs contains the configurations for google_workflows_workflow.
type WorkflowsWorkflowArgs struct {
	// CryptoKeyName: string, optional
	CryptoKeyName terra.StringValue `hcl:"crypto_key_name,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ServiceAccount: string, optional
	ServiceAccount terra.StringValue `hcl:"service_account,attr"`
	// SourceContents: string, optional
	SourceContents terra.StringValue `hcl:"source_contents,attr"`
	// Timeouts: optional
	Timeouts *workflowsworkflow.Timeouts `hcl:"timeouts,block"`
}
type workflowsWorkflowAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("create_time"))
}

// CryptoKeyName returns a reference to field crypto_key_name of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) CryptoKeyName() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("crypto_key_name"))
}

// Description returns a reference to field description of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("description"))
}

// Id returns a reference to field id of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("id"))
}

// Labels returns a reference to field labels of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ww.ref.Append("labels"))
}

// Name returns a reference to field name of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("name_prefix"))
}

// Project returns a reference to field project of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("project"))
}

// Region returns a reference to field region of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("region"))
}

// RevisionId returns a reference to field revision_id of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) RevisionId() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("revision_id"))
}

// ServiceAccount returns a reference to field service_account of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("service_account"))
}

// SourceContents returns a reference to field source_contents of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) SourceContents() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("source_contents"))
}

// State returns a reference to field state of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("state"))
}

// UpdateTime returns a reference to field update_time of google_workflows_workflow.
func (ww workflowsWorkflowAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("update_time"))
}

func (ww workflowsWorkflowAttributes) Timeouts() workflowsworkflow.TimeoutsAttributes {
	return terra.ReferenceAsSingle[workflowsworkflow.TimeoutsAttributes](ww.ref.Append("timeouts"))
}

type workflowsWorkflowState struct {
	CreateTime     string                           `json:"create_time"`
	CryptoKeyName  string                           `json:"crypto_key_name"`
	Description    string                           `json:"description"`
	Id             string                           `json:"id"`
	Labels         map[string]string                `json:"labels"`
	Name           string                           `json:"name"`
	NamePrefix     string                           `json:"name_prefix"`
	Project        string                           `json:"project"`
	Region         string                           `json:"region"`
	RevisionId     string                           `json:"revision_id"`
	ServiceAccount string                           `json:"service_account"`
	SourceContents string                           `json:"source_contents"`
	State          string                           `json:"state"`
	UpdateTime     string                           `json:"update_time"`
	Timeouts       *workflowsworkflow.TimeoutsState `json:"timeouts"`
}
