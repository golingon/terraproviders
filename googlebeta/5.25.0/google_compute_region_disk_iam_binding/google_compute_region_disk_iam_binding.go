// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_region_disk_iam_binding

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

// Resource represents the Terraform resource google_compute_region_disk_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeRegionDiskIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcrdib *Resource) Type() string {
	return "google_compute_region_disk_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gcrdib *Resource) LocalName() string {
	return gcrdib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcrdib *Resource) Configuration() interface{} {
	return gcrdib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcrdib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrdib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcrdib *Resource) Dependencies() terra.Dependencies {
	return gcrdib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcrdib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcrdib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcrdib *Resource) Attributes() googleComputeRegionDiskIamBindingAttributes {
	return googleComputeRegionDiskIamBindingAttributes{ref: terra.ReferenceResource(gcrdib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcrdib *Resource) ImportState(state io.Reader) error {
	gcrdib.state = &googleComputeRegionDiskIamBindingState{}
	if err := json.NewDecoder(state).Decode(gcrdib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrdib.Type(), gcrdib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcrdib *Resource) State() (*googleComputeRegionDiskIamBindingState, bool) {
	return gcrdib.state, gcrdib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcrdib *Resource) StateMust() *googleComputeRegionDiskIamBindingState {
	if gcrdib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrdib.Type(), gcrdib.LocalName()))
	}
	return gcrdib.state
}

// Args contains the configurations for google_compute_region_disk_iam_binding.
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

type googleComputeRegionDiskIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_region_disk_iam_binding.
func (gcrdib googleComputeRegionDiskIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gcrdib.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_region_disk_iam_binding.
func (gcrdib googleComputeRegionDiskIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrdib.ref.Append("id"))
}

// Members returns a reference to field members of google_compute_region_disk_iam_binding.
func (gcrdib googleComputeRegionDiskIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gcrdib.ref.Append("members"))
}

// Name returns a reference to field name of google_compute_region_disk_iam_binding.
func (gcrdib googleComputeRegionDiskIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcrdib.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_disk_iam_binding.
func (gcrdib googleComputeRegionDiskIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcrdib.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_disk_iam_binding.
func (gcrdib googleComputeRegionDiskIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcrdib.ref.Append("region"))
}

// Role returns a reference to field role of google_compute_region_disk_iam_binding.
func (gcrdib googleComputeRegionDiskIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gcrdib.ref.Append("role"))
}

func (gcrdib googleComputeRegionDiskIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gcrdib.ref.Append("condition"))
}

type googleComputeRegionDiskIamBindingState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Members   []string         `json:"members"`
	Name      string           `json:"name"`
	Project   string           `json:"project"`
	Region    string           `json:"region"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
