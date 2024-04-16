// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_region_backend_service_iam_binding

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

// Resource represents the Terraform resource google_compute_region_backend_service_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeRegionBackendServiceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcrbsib *Resource) Type() string {
	return "google_compute_region_backend_service_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gcrbsib *Resource) LocalName() string {
	return gcrbsib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcrbsib *Resource) Configuration() interface{} {
	return gcrbsib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcrbsib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrbsib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcrbsib *Resource) Dependencies() terra.Dependencies {
	return gcrbsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcrbsib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcrbsib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcrbsib *Resource) Attributes() googleComputeRegionBackendServiceIamBindingAttributes {
	return googleComputeRegionBackendServiceIamBindingAttributes{ref: terra.ReferenceResource(gcrbsib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcrbsib *Resource) ImportState(state io.Reader) error {
	gcrbsib.state = &googleComputeRegionBackendServiceIamBindingState{}
	if err := json.NewDecoder(state).Decode(gcrbsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrbsib.Type(), gcrbsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcrbsib *Resource) State() (*googleComputeRegionBackendServiceIamBindingState, bool) {
	return gcrbsib.state, gcrbsib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcrbsib *Resource) StateMust() *googleComputeRegionBackendServiceIamBindingState {
	if gcrbsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrbsib.Type(), gcrbsib.LocalName()))
	}
	return gcrbsib.state
}

// Args contains the configurations for google_compute_region_backend_service_iam_binding.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleComputeRegionBackendServiceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_region_backend_service_iam_binding.
func (gcrbsib googleComputeRegionBackendServiceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gcrbsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_region_backend_service_iam_binding.
func (gcrbsib googleComputeRegionBackendServiceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrbsib.ref.Append("id"))
}

// Members returns a reference to field members of google_compute_region_backend_service_iam_binding.
func (gcrbsib googleComputeRegionBackendServiceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gcrbsib.ref.Append("members"))
}

// Name returns a reference to field name of google_compute_region_backend_service_iam_binding.
func (gcrbsib googleComputeRegionBackendServiceIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcrbsib.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_backend_service_iam_binding.
func (gcrbsib googleComputeRegionBackendServiceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcrbsib.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_backend_service_iam_binding.
func (gcrbsib googleComputeRegionBackendServiceIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcrbsib.ref.Append("region"))
}

// Role returns a reference to field role of google_compute_region_backend_service_iam_binding.
func (gcrbsib googleComputeRegionBackendServiceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gcrbsib.ref.Append("role"))
}

func (gcrbsib googleComputeRegionBackendServiceIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gcrbsib.ref.Append("condition"))
}

type googleComputeRegionBackendServiceIamBindingState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Members   []string         `json:"members"`
	Name      string           `json:"name"`
	Project   string           `json:"project"`
	Region    string           `json:"region"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
