// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewComputeInstanceIamPolicy creates a new instance of [ComputeInstanceIamPolicy].
func NewComputeInstanceIamPolicy(name string, args ComputeInstanceIamPolicyArgs) *ComputeInstanceIamPolicy {
	return &ComputeInstanceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInstanceIamPolicy)(nil)

// ComputeInstanceIamPolicy represents the Terraform resource google_compute_instance_iam_policy.
type ComputeInstanceIamPolicy struct {
	Name      string
	Args      ComputeInstanceIamPolicyArgs
	state     *computeInstanceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeInstanceIamPolicy].
func (ciip *ComputeInstanceIamPolicy) Type() string {
	return "google_compute_instance_iam_policy"
}

// LocalName returns the local name for [ComputeInstanceIamPolicy].
func (ciip *ComputeInstanceIamPolicy) LocalName() string {
	return ciip.Name
}

// Configuration returns the configuration (args) for [ComputeInstanceIamPolicy].
func (ciip *ComputeInstanceIamPolicy) Configuration() interface{} {
	return ciip.Args
}

// DependOn is used for other resources to depend on [ComputeInstanceIamPolicy].
func (ciip *ComputeInstanceIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ciip)
}

// Dependencies returns the list of resources [ComputeInstanceIamPolicy] depends_on.
func (ciip *ComputeInstanceIamPolicy) Dependencies() terra.Dependencies {
	return ciip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeInstanceIamPolicy].
func (ciip *ComputeInstanceIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return ciip.Lifecycle
}

// Attributes returns the attributes for [ComputeInstanceIamPolicy].
func (ciip *ComputeInstanceIamPolicy) Attributes() computeInstanceIamPolicyAttributes {
	return computeInstanceIamPolicyAttributes{ref: terra.ReferenceResource(ciip)}
}

// ImportState imports the given attribute values into [ComputeInstanceIamPolicy]'s state.
func (ciip *ComputeInstanceIamPolicy) ImportState(av io.Reader) error {
	ciip.state = &computeInstanceIamPolicyState{}
	if err := json.NewDecoder(av).Decode(ciip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ciip.Type(), ciip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeInstanceIamPolicy] has state.
func (ciip *ComputeInstanceIamPolicy) State() (*computeInstanceIamPolicyState, bool) {
	return ciip.state, ciip.state != nil
}

// StateMust returns the state for [ComputeInstanceIamPolicy]. Panics if the state is nil.
func (ciip *ComputeInstanceIamPolicy) StateMust() *computeInstanceIamPolicyState {
	if ciip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ciip.Type(), ciip.LocalName()))
	}
	return ciip.state
}

// ComputeInstanceIamPolicyArgs contains the configurations for google_compute_instance_iam_policy.
type ComputeInstanceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
}
type computeInstanceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_instance_iam_policy.
func (ciip computeInstanceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_instance_iam_policy.
func (ciip computeInstanceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of google_compute_instance_iam_policy.
func (ciip computeInstanceIamPolicyAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("instance_name"))
}

// PolicyData returns a reference to field policy_data of google_compute_instance_iam_policy.
func (ciip computeInstanceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_compute_instance_iam_policy.
func (ciip computeInstanceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("project"))
}

// Zone returns a reference to field zone of google_compute_instance_iam_policy.
func (ciip computeInstanceIamPolicyAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("zone"))
}

type computeInstanceIamPolicyState struct {
	Etag         string `json:"etag"`
	Id           string `json:"id"`
	InstanceName string `json:"instance_name"`
	PolicyData   string `json:"policy_data"`
	Project      string `json:"project"`
	Zone         string `json:"zone"`
}
