// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataproc_job_iam_binding

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

// Resource represents the Terraform resource google_dataproc_job_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataprocJobIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdjib *Resource) Type() string {
	return "google_dataproc_job_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gdjib *Resource) LocalName() string {
	return gdjib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdjib *Resource) Configuration() interface{} {
	return gdjib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdjib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdjib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdjib *Resource) Dependencies() terra.Dependencies {
	return gdjib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdjib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdjib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdjib *Resource) Attributes() googleDataprocJobIamBindingAttributes {
	return googleDataprocJobIamBindingAttributes{ref: terra.ReferenceResource(gdjib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdjib *Resource) ImportState(state io.Reader) error {
	gdjib.state = &googleDataprocJobIamBindingState{}
	if err := json.NewDecoder(state).Decode(gdjib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdjib.Type(), gdjib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdjib *Resource) State() (*googleDataprocJobIamBindingState, bool) {
	return gdjib.state, gdjib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdjib *Resource) StateMust() *googleDataprocJobIamBindingState {
	if gdjib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdjib.Type(), gdjib.LocalName()))
	}
	return gdjib.state
}

// Args contains the configurations for google_dataproc_job_iam_binding.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// JobId: string, required
	JobId terra.StringValue `hcl:"job_id,attr" validate:"required"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleDataprocJobIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_job_iam_binding.
func (gdjib googleDataprocJobIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gdjib.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_job_iam_binding.
func (gdjib googleDataprocJobIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdjib.ref.Append("id"))
}

// JobId returns a reference to field job_id of google_dataproc_job_iam_binding.
func (gdjib googleDataprocJobIamBindingAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(gdjib.ref.Append("job_id"))
}

// Members returns a reference to field members of google_dataproc_job_iam_binding.
func (gdjib googleDataprocJobIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gdjib.ref.Append("members"))
}

// Project returns a reference to field project of google_dataproc_job_iam_binding.
func (gdjib googleDataprocJobIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdjib.ref.Append("project"))
}

// Region returns a reference to field region of google_dataproc_job_iam_binding.
func (gdjib googleDataprocJobIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gdjib.ref.Append("region"))
}

// Role returns a reference to field role of google_dataproc_job_iam_binding.
func (gdjib googleDataprocJobIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gdjib.ref.Append("role"))
}

func (gdjib googleDataprocJobIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gdjib.ref.Append("condition"))
}

type googleDataprocJobIamBindingState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	JobId     string           `json:"job_id"`
	Members   []string         `json:"members"`
	Project   string           `json:"project"`
	Region    string           `json:"region"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
