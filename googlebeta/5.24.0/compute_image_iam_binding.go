// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computeimageiambinding "github.com/golingon/terraproviders/googlebeta/5.24.0/computeimageiambinding"
	"io"
)

// NewComputeImageIamBinding creates a new instance of [ComputeImageIamBinding].
func NewComputeImageIamBinding(name string, args ComputeImageIamBindingArgs) *ComputeImageIamBinding {
	return &ComputeImageIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeImageIamBinding)(nil)

// ComputeImageIamBinding represents the Terraform resource google_compute_image_iam_binding.
type ComputeImageIamBinding struct {
	Name      string
	Args      ComputeImageIamBindingArgs
	state     *computeImageIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeImageIamBinding].
func (ciib *ComputeImageIamBinding) Type() string {
	return "google_compute_image_iam_binding"
}

// LocalName returns the local name for [ComputeImageIamBinding].
func (ciib *ComputeImageIamBinding) LocalName() string {
	return ciib.Name
}

// Configuration returns the configuration (args) for [ComputeImageIamBinding].
func (ciib *ComputeImageIamBinding) Configuration() interface{} {
	return ciib.Args
}

// DependOn is used for other resources to depend on [ComputeImageIamBinding].
func (ciib *ComputeImageIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ciib)
}

// Dependencies returns the list of resources [ComputeImageIamBinding] depends_on.
func (ciib *ComputeImageIamBinding) Dependencies() terra.Dependencies {
	return ciib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeImageIamBinding].
func (ciib *ComputeImageIamBinding) LifecycleManagement() *terra.Lifecycle {
	return ciib.Lifecycle
}

// Attributes returns the attributes for [ComputeImageIamBinding].
func (ciib *ComputeImageIamBinding) Attributes() computeImageIamBindingAttributes {
	return computeImageIamBindingAttributes{ref: terra.ReferenceResource(ciib)}
}

// ImportState imports the given attribute values into [ComputeImageIamBinding]'s state.
func (ciib *ComputeImageIamBinding) ImportState(av io.Reader) error {
	ciib.state = &computeImageIamBindingState{}
	if err := json.NewDecoder(av).Decode(ciib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ciib.Type(), ciib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeImageIamBinding] has state.
func (ciib *ComputeImageIamBinding) State() (*computeImageIamBindingState, bool) {
	return ciib.state, ciib.state != nil
}

// StateMust returns the state for [ComputeImageIamBinding]. Panics if the state is nil.
func (ciib *ComputeImageIamBinding) StateMust() *computeImageIamBindingState {
	if ciib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ciib.Type(), ciib.LocalName()))
	}
	return ciib.state
}

// ComputeImageIamBindingArgs contains the configurations for google_compute_image_iam_binding.
type ComputeImageIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Image: string, required
	Image terra.StringValue `hcl:"image,attr" validate:"required"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *computeimageiambinding.Condition `hcl:"condition,block"`
}
type computeImageIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_image_iam_binding.
func (ciib computeImageIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ciib.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_image_iam_binding.
func (ciib computeImageIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ciib.ref.Append("id"))
}

// Image returns a reference to field image of google_compute_image_iam_binding.
func (ciib computeImageIamBindingAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(ciib.ref.Append("image"))
}

// Members returns a reference to field members of google_compute_image_iam_binding.
func (ciib computeImageIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ciib.ref.Append("members"))
}

// Project returns a reference to field project of google_compute_image_iam_binding.
func (ciib computeImageIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ciib.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_image_iam_binding.
func (ciib computeImageIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ciib.ref.Append("role"))
}

func (ciib computeImageIamBindingAttributes) Condition() terra.ListValue[computeimageiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[computeimageiambinding.ConditionAttributes](ciib.ref.Append("condition"))
}

type computeImageIamBindingState struct {
	Etag      string                                  `json:"etag"`
	Id        string                                  `json:"id"`
	Image     string                                  `json:"image"`
	Members   []string                                `json:"members"`
	Project   string                                  `json:"project"`
	Role      string                                  `json:"role"`
	Condition []computeimageiambinding.ConditionState `json:"condition"`
}
