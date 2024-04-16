// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_cloud_run_service_iam_binding

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

// Resource represents the Terraform resource google_cloud_run_service_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleCloudRunServiceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcrsib *Resource) Type() string {
	return "google_cloud_run_service_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gcrsib *Resource) LocalName() string {
	return gcrsib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcrsib *Resource) Configuration() interface{} {
	return gcrsib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcrsib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrsib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcrsib *Resource) Dependencies() terra.Dependencies {
	return gcrsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcrsib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcrsib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcrsib *Resource) Attributes() googleCloudRunServiceIamBindingAttributes {
	return googleCloudRunServiceIamBindingAttributes{ref: terra.ReferenceResource(gcrsib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcrsib *Resource) ImportState(state io.Reader) error {
	gcrsib.state = &googleCloudRunServiceIamBindingState{}
	if err := json.NewDecoder(state).Decode(gcrsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrsib.Type(), gcrsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcrsib *Resource) State() (*googleCloudRunServiceIamBindingState, bool) {
	return gcrsib.state, gcrsib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcrsib *Resource) StateMust() *googleCloudRunServiceIamBindingState {
	if gcrsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrsib.Type(), gcrsib.LocalName()))
	}
	return gcrsib.state
}

// Args contains the configurations for google_cloud_run_service_iam_binding.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleCloudRunServiceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_run_service_iam_binding.
func (gcrsib googleCloudRunServiceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gcrsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_run_service_iam_binding.
func (gcrsib googleCloudRunServiceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrsib.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_service_iam_binding.
func (gcrsib googleCloudRunServiceIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gcrsib.ref.Append("location"))
}

// Members returns a reference to field members of google_cloud_run_service_iam_binding.
func (gcrsib googleCloudRunServiceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gcrsib.ref.Append("members"))
}

// Project returns a reference to field project of google_cloud_run_service_iam_binding.
func (gcrsib googleCloudRunServiceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcrsib.ref.Append("project"))
}

// Role returns a reference to field role of google_cloud_run_service_iam_binding.
func (gcrsib googleCloudRunServiceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gcrsib.ref.Append("role"))
}

// Service returns a reference to field service of google_cloud_run_service_iam_binding.
func (gcrsib googleCloudRunServiceIamBindingAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(gcrsib.ref.Append("service"))
}

func (gcrsib googleCloudRunServiceIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gcrsib.ref.Append("condition"))
}

type googleCloudRunServiceIamBindingState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Location  string           `json:"location"`
	Members   []string         `json:"members"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	Service   string           `json:"service"`
	Condition []ConditionState `json:"condition"`
}
