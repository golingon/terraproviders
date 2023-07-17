// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computesubnetworkiambinding "github.com/golingon/terraproviders/google/4.73.1/computesubnetworkiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeSubnetworkIamBinding creates a new instance of [ComputeSubnetworkIamBinding].
func NewComputeSubnetworkIamBinding(name string, args ComputeSubnetworkIamBindingArgs) *ComputeSubnetworkIamBinding {
	return &ComputeSubnetworkIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSubnetworkIamBinding)(nil)

// ComputeSubnetworkIamBinding represents the Terraform resource google_compute_subnetwork_iam_binding.
type ComputeSubnetworkIamBinding struct {
	Name      string
	Args      ComputeSubnetworkIamBindingArgs
	state     *computeSubnetworkIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeSubnetworkIamBinding].
func (csib *ComputeSubnetworkIamBinding) Type() string {
	return "google_compute_subnetwork_iam_binding"
}

// LocalName returns the local name for [ComputeSubnetworkIamBinding].
func (csib *ComputeSubnetworkIamBinding) LocalName() string {
	return csib.Name
}

// Configuration returns the configuration (args) for [ComputeSubnetworkIamBinding].
func (csib *ComputeSubnetworkIamBinding) Configuration() interface{} {
	return csib.Args
}

// DependOn is used for other resources to depend on [ComputeSubnetworkIamBinding].
func (csib *ComputeSubnetworkIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(csib)
}

// Dependencies returns the list of resources [ComputeSubnetworkIamBinding] depends_on.
func (csib *ComputeSubnetworkIamBinding) Dependencies() terra.Dependencies {
	return csib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeSubnetworkIamBinding].
func (csib *ComputeSubnetworkIamBinding) LifecycleManagement() *terra.Lifecycle {
	return csib.Lifecycle
}

// Attributes returns the attributes for [ComputeSubnetworkIamBinding].
func (csib *ComputeSubnetworkIamBinding) Attributes() computeSubnetworkIamBindingAttributes {
	return computeSubnetworkIamBindingAttributes{ref: terra.ReferenceResource(csib)}
}

// ImportState imports the given attribute values into [ComputeSubnetworkIamBinding]'s state.
func (csib *ComputeSubnetworkIamBinding) ImportState(av io.Reader) error {
	csib.state = &computeSubnetworkIamBindingState{}
	if err := json.NewDecoder(av).Decode(csib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csib.Type(), csib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeSubnetworkIamBinding] has state.
func (csib *ComputeSubnetworkIamBinding) State() (*computeSubnetworkIamBindingState, bool) {
	return csib.state, csib.state != nil
}

// StateMust returns the state for [ComputeSubnetworkIamBinding]. Panics if the state is nil.
func (csib *ComputeSubnetworkIamBinding) StateMust() *computeSubnetworkIamBindingState {
	if csib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csib.Type(), csib.LocalName()))
	}
	return csib.state
}

// ComputeSubnetworkIamBindingArgs contains the configurations for google_compute_subnetwork_iam_binding.
type ComputeSubnetworkIamBindingArgs struct {
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
	// Subnetwork: string, required
	Subnetwork terra.StringValue `hcl:"subnetwork,attr" validate:"required"`
	// Condition: optional
	Condition *computesubnetworkiambinding.Condition `hcl:"condition,block"`
}
type computeSubnetworkIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_subnetwork_iam_binding.
func (csib computeSubnetworkIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(csib.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_subnetwork_iam_binding.
func (csib computeSubnetworkIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csib.ref.Append("id"))
}

// Members returns a reference to field members of google_compute_subnetwork_iam_binding.
func (csib computeSubnetworkIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csib.ref.Append("members"))
}

// Project returns a reference to field project of google_compute_subnetwork_iam_binding.
func (csib computeSubnetworkIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csib.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_subnetwork_iam_binding.
func (csib computeSubnetworkIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(csib.ref.Append("region"))
}

// Role returns a reference to field role of google_compute_subnetwork_iam_binding.
func (csib computeSubnetworkIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(csib.ref.Append("role"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_subnetwork_iam_binding.
func (csib computeSubnetworkIamBindingAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(csib.ref.Append("subnetwork"))
}

func (csib computeSubnetworkIamBindingAttributes) Condition() terra.ListValue[computesubnetworkiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[computesubnetworkiambinding.ConditionAttributes](csib.ref.Append("condition"))
}

type computeSubnetworkIamBindingState struct {
	Etag       string                                       `json:"etag"`
	Id         string                                       `json:"id"`
	Members    []string                                     `json:"members"`
	Project    string                                       `json:"project"`
	Region     string                                       `json:"region"`
	Role       string                                       `json:"role"`
	Subnetwork string                                       `json:"subnetwork"`
	Condition  []computesubnetworkiambinding.ConditionState `json:"condition"`
}
