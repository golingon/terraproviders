// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataproc_cluster_iam_binding

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

// Resource represents the Terraform resource google_dataproc_cluster_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataprocClusterIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdcib *Resource) Type() string {
	return "google_dataproc_cluster_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gdcib *Resource) LocalName() string {
	return gdcib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdcib *Resource) Configuration() interface{} {
	return gdcib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdcib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdcib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdcib *Resource) Dependencies() terra.Dependencies {
	return gdcib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdcib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdcib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdcib *Resource) Attributes() googleDataprocClusterIamBindingAttributes {
	return googleDataprocClusterIamBindingAttributes{ref: terra.ReferenceResource(gdcib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdcib *Resource) ImportState(state io.Reader) error {
	gdcib.state = &googleDataprocClusterIamBindingState{}
	if err := json.NewDecoder(state).Decode(gdcib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdcib.Type(), gdcib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdcib *Resource) State() (*googleDataprocClusterIamBindingState, bool) {
	return gdcib.state, gdcib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdcib *Resource) StateMust() *googleDataprocClusterIamBindingState {
	if gdcib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdcib.Type(), gdcib.LocalName()))
	}
	return gdcib.state
}

// Args contains the configurations for google_dataproc_cluster_iam_binding.
type Args struct {
	// Cluster: string, required
	Cluster terra.StringValue `hcl:"cluster,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
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

type googleDataprocClusterIamBindingAttributes struct {
	ref terra.Reference
}

// Cluster returns a reference to field cluster of google_dataproc_cluster_iam_binding.
func (gdcib googleDataprocClusterIamBindingAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(gdcib.ref.Append("cluster"))
}

// Etag returns a reference to field etag of google_dataproc_cluster_iam_binding.
func (gdcib googleDataprocClusterIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gdcib.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_cluster_iam_binding.
func (gdcib googleDataprocClusterIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdcib.ref.Append("id"))
}

// Members returns a reference to field members of google_dataproc_cluster_iam_binding.
func (gdcib googleDataprocClusterIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gdcib.ref.Append("members"))
}

// Project returns a reference to field project of google_dataproc_cluster_iam_binding.
func (gdcib googleDataprocClusterIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdcib.ref.Append("project"))
}

// Region returns a reference to field region of google_dataproc_cluster_iam_binding.
func (gdcib googleDataprocClusterIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gdcib.ref.Append("region"))
}

// Role returns a reference to field role of google_dataproc_cluster_iam_binding.
func (gdcib googleDataprocClusterIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gdcib.ref.Append("role"))
}

func (gdcib googleDataprocClusterIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gdcib.ref.Append("condition"))
}

type googleDataprocClusterIamBindingState struct {
	Cluster   string           `json:"cluster"`
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Members   []string         `json:"members"`
	Project   string           `json:"project"`
	Region    string           `json:"region"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
