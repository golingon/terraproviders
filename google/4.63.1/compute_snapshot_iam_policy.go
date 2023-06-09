// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeSnapshotIamPolicy creates a new instance of [ComputeSnapshotIamPolicy].
func NewComputeSnapshotIamPolicy(name string, args ComputeSnapshotIamPolicyArgs) *ComputeSnapshotIamPolicy {
	return &ComputeSnapshotIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSnapshotIamPolicy)(nil)

// ComputeSnapshotIamPolicy represents the Terraform resource google_compute_snapshot_iam_policy.
type ComputeSnapshotIamPolicy struct {
	Name      string
	Args      ComputeSnapshotIamPolicyArgs
	state     *computeSnapshotIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeSnapshotIamPolicy].
func (csip *ComputeSnapshotIamPolicy) Type() string {
	return "google_compute_snapshot_iam_policy"
}

// LocalName returns the local name for [ComputeSnapshotIamPolicy].
func (csip *ComputeSnapshotIamPolicy) LocalName() string {
	return csip.Name
}

// Configuration returns the configuration (args) for [ComputeSnapshotIamPolicy].
func (csip *ComputeSnapshotIamPolicy) Configuration() interface{} {
	return csip.Args
}

// DependOn is used for other resources to depend on [ComputeSnapshotIamPolicy].
func (csip *ComputeSnapshotIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(csip)
}

// Dependencies returns the list of resources [ComputeSnapshotIamPolicy] depends_on.
func (csip *ComputeSnapshotIamPolicy) Dependencies() terra.Dependencies {
	return csip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeSnapshotIamPolicy].
func (csip *ComputeSnapshotIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return csip.Lifecycle
}

// Attributes returns the attributes for [ComputeSnapshotIamPolicy].
func (csip *ComputeSnapshotIamPolicy) Attributes() computeSnapshotIamPolicyAttributes {
	return computeSnapshotIamPolicyAttributes{ref: terra.ReferenceResource(csip)}
}

// ImportState imports the given attribute values into [ComputeSnapshotIamPolicy]'s state.
func (csip *ComputeSnapshotIamPolicy) ImportState(av io.Reader) error {
	csip.state = &computeSnapshotIamPolicyState{}
	if err := json.NewDecoder(av).Decode(csip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csip.Type(), csip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeSnapshotIamPolicy] has state.
func (csip *ComputeSnapshotIamPolicy) State() (*computeSnapshotIamPolicyState, bool) {
	return csip.state, csip.state != nil
}

// StateMust returns the state for [ComputeSnapshotIamPolicy]. Panics if the state is nil.
func (csip *ComputeSnapshotIamPolicy) StateMust() *computeSnapshotIamPolicyState {
	if csip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csip.Type(), csip.LocalName()))
	}
	return csip.state
}

// ComputeSnapshotIamPolicyArgs contains the configurations for google_compute_snapshot_iam_policy.
type ComputeSnapshotIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type computeSnapshotIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_snapshot_iam_policy.
func (csip computeSnapshotIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(csip.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_snapshot_iam_policy.
func (csip computeSnapshotIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csip.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_snapshot_iam_policy.
func (csip computeSnapshotIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_compute_snapshot_iam_policy.
func (csip computeSnapshotIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(csip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_compute_snapshot_iam_policy.
func (csip computeSnapshotIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csip.ref.Append("project"))
}

type computeSnapshotIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
