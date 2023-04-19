// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregiondiskiambinding "github.com/golingon/terraproviders/google/4.62.0/computeregiondiskiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionDiskIamBinding creates a new instance of [ComputeRegionDiskIamBinding].
func NewComputeRegionDiskIamBinding(name string, args ComputeRegionDiskIamBindingArgs) *ComputeRegionDiskIamBinding {
	return &ComputeRegionDiskIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionDiskIamBinding)(nil)

// ComputeRegionDiskIamBinding represents the Terraform resource google_compute_region_disk_iam_binding.
type ComputeRegionDiskIamBinding struct {
	Name      string
	Args      ComputeRegionDiskIamBindingArgs
	state     *computeRegionDiskIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionDiskIamBinding].
func (crdib *ComputeRegionDiskIamBinding) Type() string {
	return "google_compute_region_disk_iam_binding"
}

// LocalName returns the local name for [ComputeRegionDiskIamBinding].
func (crdib *ComputeRegionDiskIamBinding) LocalName() string {
	return crdib.Name
}

// Configuration returns the configuration (args) for [ComputeRegionDiskIamBinding].
func (crdib *ComputeRegionDiskIamBinding) Configuration() interface{} {
	return crdib.Args
}

// DependOn is used for other resources to depend on [ComputeRegionDiskIamBinding].
func (crdib *ComputeRegionDiskIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(crdib)
}

// Dependencies returns the list of resources [ComputeRegionDiskIamBinding] depends_on.
func (crdib *ComputeRegionDiskIamBinding) Dependencies() terra.Dependencies {
	return crdib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionDiskIamBinding].
func (crdib *ComputeRegionDiskIamBinding) LifecycleManagement() *terra.Lifecycle {
	return crdib.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionDiskIamBinding].
func (crdib *ComputeRegionDiskIamBinding) Attributes() computeRegionDiskIamBindingAttributes {
	return computeRegionDiskIamBindingAttributes{ref: terra.ReferenceResource(crdib)}
}

// ImportState imports the given attribute values into [ComputeRegionDiskIamBinding]'s state.
func (crdib *ComputeRegionDiskIamBinding) ImportState(av io.Reader) error {
	crdib.state = &computeRegionDiskIamBindingState{}
	if err := json.NewDecoder(av).Decode(crdib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crdib.Type(), crdib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionDiskIamBinding] has state.
func (crdib *ComputeRegionDiskIamBinding) State() (*computeRegionDiskIamBindingState, bool) {
	return crdib.state, crdib.state != nil
}

// StateMust returns the state for [ComputeRegionDiskIamBinding]. Panics if the state is nil.
func (crdib *ComputeRegionDiskIamBinding) StateMust() *computeRegionDiskIamBindingState {
	if crdib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crdib.Type(), crdib.LocalName()))
	}
	return crdib.state
}

// ComputeRegionDiskIamBindingArgs contains the configurations for google_compute_region_disk_iam_binding.
type ComputeRegionDiskIamBindingArgs struct {
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
	Condition *computeregiondiskiambinding.Condition `hcl:"condition,block"`
}
type computeRegionDiskIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_region_disk_iam_binding.
func (crdib computeRegionDiskIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crdib.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_region_disk_iam_binding.
func (crdib computeRegionDiskIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crdib.ref.Append("id"))
}

// Members returns a reference to field members of google_compute_region_disk_iam_binding.
func (crdib computeRegionDiskIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crdib.ref.Append("members"))
}

// Name returns a reference to field name of google_compute_region_disk_iam_binding.
func (crdib computeRegionDiskIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crdib.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_disk_iam_binding.
func (crdib computeRegionDiskIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crdib.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_disk_iam_binding.
func (crdib computeRegionDiskIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crdib.ref.Append("region"))
}

// Role returns a reference to field role of google_compute_region_disk_iam_binding.
func (crdib computeRegionDiskIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(crdib.ref.Append("role"))
}

func (crdib computeRegionDiskIamBindingAttributes) Condition() terra.ListValue[computeregiondiskiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[computeregiondiskiambinding.ConditionAttributes](crdib.ref.Append("condition"))
}

type computeRegionDiskIamBindingState struct {
	Etag      string                                       `json:"etag"`
	Id        string                                       `json:"id"`
	Members   []string                                     `json:"members"`
	Name      string                                       `json:"name"`
	Project   string                                       `json:"project"`
	Region    string                                       `json:"region"`
	Role      string                                       `json:"role"`
	Condition []computeregiondiskiambinding.ConditionState `json:"condition"`
}
