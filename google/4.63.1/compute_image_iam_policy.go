// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeImageIamPolicy creates a new instance of [ComputeImageIamPolicy].
func NewComputeImageIamPolicy(name string, args ComputeImageIamPolicyArgs) *ComputeImageIamPolicy {
	return &ComputeImageIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeImageIamPolicy)(nil)

// ComputeImageIamPolicy represents the Terraform resource google_compute_image_iam_policy.
type ComputeImageIamPolicy struct {
	Name      string
	Args      ComputeImageIamPolicyArgs
	state     *computeImageIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeImageIamPolicy].
func (ciip *ComputeImageIamPolicy) Type() string {
	return "google_compute_image_iam_policy"
}

// LocalName returns the local name for [ComputeImageIamPolicy].
func (ciip *ComputeImageIamPolicy) LocalName() string {
	return ciip.Name
}

// Configuration returns the configuration (args) for [ComputeImageIamPolicy].
func (ciip *ComputeImageIamPolicy) Configuration() interface{} {
	return ciip.Args
}

// DependOn is used for other resources to depend on [ComputeImageIamPolicy].
func (ciip *ComputeImageIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ciip)
}

// Dependencies returns the list of resources [ComputeImageIamPolicy] depends_on.
func (ciip *ComputeImageIamPolicy) Dependencies() terra.Dependencies {
	return ciip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeImageIamPolicy].
func (ciip *ComputeImageIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return ciip.Lifecycle
}

// Attributes returns the attributes for [ComputeImageIamPolicy].
func (ciip *ComputeImageIamPolicy) Attributes() computeImageIamPolicyAttributes {
	return computeImageIamPolicyAttributes{ref: terra.ReferenceResource(ciip)}
}

// ImportState imports the given attribute values into [ComputeImageIamPolicy]'s state.
func (ciip *ComputeImageIamPolicy) ImportState(av io.Reader) error {
	ciip.state = &computeImageIamPolicyState{}
	if err := json.NewDecoder(av).Decode(ciip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ciip.Type(), ciip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeImageIamPolicy] has state.
func (ciip *ComputeImageIamPolicy) State() (*computeImageIamPolicyState, bool) {
	return ciip.state, ciip.state != nil
}

// StateMust returns the state for [ComputeImageIamPolicy]. Panics if the state is nil.
func (ciip *ComputeImageIamPolicy) StateMust() *computeImageIamPolicyState {
	if ciip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ciip.Type(), ciip.LocalName()))
	}
	return ciip.state
}

// ComputeImageIamPolicyArgs contains the configurations for google_compute_image_iam_policy.
type ComputeImageIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Image: string, required
	Image terra.StringValue `hcl:"image,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type computeImageIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_image_iam_policy.
func (ciip computeImageIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_image_iam_policy.
func (ciip computeImageIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("id"))
}

// Image returns a reference to field image of google_compute_image_iam_policy.
func (ciip computeImageIamPolicyAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("image"))
}

// PolicyData returns a reference to field policy_data of google_compute_image_iam_policy.
func (ciip computeImageIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_compute_image_iam_policy.
func (ciip computeImageIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("project"))
}

type computeImageIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Image      string `json:"image"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}