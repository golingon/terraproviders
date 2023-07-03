// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computesnapshotiammember "github.com/golingon/terraproviders/google/4.71.0/computesnapshotiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeSnapshotIamMember creates a new instance of [ComputeSnapshotIamMember].
func NewComputeSnapshotIamMember(name string, args ComputeSnapshotIamMemberArgs) *ComputeSnapshotIamMember {
	return &ComputeSnapshotIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSnapshotIamMember)(nil)

// ComputeSnapshotIamMember represents the Terraform resource google_compute_snapshot_iam_member.
type ComputeSnapshotIamMember struct {
	Name      string
	Args      ComputeSnapshotIamMemberArgs
	state     *computeSnapshotIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeSnapshotIamMember].
func (csim *ComputeSnapshotIamMember) Type() string {
	return "google_compute_snapshot_iam_member"
}

// LocalName returns the local name for [ComputeSnapshotIamMember].
func (csim *ComputeSnapshotIamMember) LocalName() string {
	return csim.Name
}

// Configuration returns the configuration (args) for [ComputeSnapshotIamMember].
func (csim *ComputeSnapshotIamMember) Configuration() interface{} {
	return csim.Args
}

// DependOn is used for other resources to depend on [ComputeSnapshotIamMember].
func (csim *ComputeSnapshotIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(csim)
}

// Dependencies returns the list of resources [ComputeSnapshotIamMember] depends_on.
func (csim *ComputeSnapshotIamMember) Dependencies() terra.Dependencies {
	return csim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeSnapshotIamMember].
func (csim *ComputeSnapshotIamMember) LifecycleManagement() *terra.Lifecycle {
	return csim.Lifecycle
}

// Attributes returns the attributes for [ComputeSnapshotIamMember].
func (csim *ComputeSnapshotIamMember) Attributes() computeSnapshotIamMemberAttributes {
	return computeSnapshotIamMemberAttributes{ref: terra.ReferenceResource(csim)}
}

// ImportState imports the given attribute values into [ComputeSnapshotIamMember]'s state.
func (csim *ComputeSnapshotIamMember) ImportState(av io.Reader) error {
	csim.state = &computeSnapshotIamMemberState{}
	if err := json.NewDecoder(av).Decode(csim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csim.Type(), csim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeSnapshotIamMember] has state.
func (csim *ComputeSnapshotIamMember) State() (*computeSnapshotIamMemberState, bool) {
	return csim.state, csim.state != nil
}

// StateMust returns the state for [ComputeSnapshotIamMember]. Panics if the state is nil.
func (csim *ComputeSnapshotIamMember) StateMust() *computeSnapshotIamMemberState {
	if csim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csim.Type(), csim.LocalName()))
	}
	return csim.state
}

// ComputeSnapshotIamMemberArgs contains the configurations for google_compute_snapshot_iam_member.
type ComputeSnapshotIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *computesnapshotiammember.Condition `hcl:"condition,block"`
}
type computeSnapshotIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_snapshot_iam_member.
func (csim computeSnapshotIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_snapshot_iam_member.
func (csim computeSnapshotIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("id"))
}

// Member returns a reference to field member of google_compute_snapshot_iam_member.
func (csim computeSnapshotIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("member"))
}

// Name returns a reference to field name of google_compute_snapshot_iam_member.
func (csim computeSnapshotIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_snapshot_iam_member.
func (csim computeSnapshotIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_snapshot_iam_member.
func (csim computeSnapshotIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("role"))
}

func (csim computeSnapshotIamMemberAttributes) Condition() terra.ListValue[computesnapshotiammember.ConditionAttributes] {
	return terra.ReferenceAsList[computesnapshotiammember.ConditionAttributes](csim.ref.Append("condition"))
}

type computeSnapshotIamMemberState struct {
	Etag      string                                    `json:"etag"`
	Id        string                                    `json:"id"`
	Member    string                                    `json:"member"`
	Name      string                                    `json:"name"`
	Project   string                                    `json:"project"`
	Role      string                                    `json:"role"`
	Condition []computesnapshotiammember.ConditionState `json:"condition"`
}
