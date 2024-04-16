// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_snapshot_iam_member

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

// Resource represents the Terraform resource google_compute_snapshot_iam_member.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeSnapshotIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcsim *Resource) Type() string {
	return "google_compute_snapshot_iam_member"
}

// LocalName returns the local name for [Resource].
func (gcsim *Resource) LocalName() string {
	return gcsim.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcsim *Resource) Configuration() interface{} {
	return gcsim.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcsim *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcsim)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcsim *Resource) Dependencies() terra.Dependencies {
	return gcsim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcsim *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcsim.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcsim *Resource) Attributes() googleComputeSnapshotIamMemberAttributes {
	return googleComputeSnapshotIamMemberAttributes{ref: terra.ReferenceResource(gcsim)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcsim *Resource) ImportState(state io.Reader) error {
	gcsim.state = &googleComputeSnapshotIamMemberState{}
	if err := json.NewDecoder(state).Decode(gcsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcsim.Type(), gcsim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcsim *Resource) State() (*googleComputeSnapshotIamMemberState, bool) {
	return gcsim.state, gcsim.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcsim *Resource) StateMust() *googleComputeSnapshotIamMemberState {
	if gcsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcsim.Type(), gcsim.LocalName()))
	}
	return gcsim.state
}

// Args contains the configurations for google_compute_snapshot_iam_member.
type Args struct {
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
	Condition *Condition `hcl:"condition,block"`
}

type googleComputeSnapshotIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_snapshot_iam_member.
func (gcsim googleComputeSnapshotIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gcsim.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_snapshot_iam_member.
func (gcsim googleComputeSnapshotIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcsim.ref.Append("id"))
}

// Member returns a reference to field member of google_compute_snapshot_iam_member.
func (gcsim googleComputeSnapshotIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(gcsim.ref.Append("member"))
}

// Name returns a reference to field name of google_compute_snapshot_iam_member.
func (gcsim googleComputeSnapshotIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcsim.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_snapshot_iam_member.
func (gcsim googleComputeSnapshotIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcsim.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_snapshot_iam_member.
func (gcsim googleComputeSnapshotIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gcsim.ref.Append("role"))
}

func (gcsim googleComputeSnapshotIamMemberAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gcsim.ref.Append("condition"))
}

type googleComputeSnapshotIamMemberState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Member    string           `json:"member"`
	Name      string           `json:"name"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
