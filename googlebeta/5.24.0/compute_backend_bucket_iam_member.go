// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computebackendbucketiammember "github.com/golingon/terraproviders/googlebeta/5.24.0/computebackendbucketiammember"
	"io"
)

// NewComputeBackendBucketIamMember creates a new instance of [ComputeBackendBucketIamMember].
func NewComputeBackendBucketIamMember(name string, args ComputeBackendBucketIamMemberArgs) *ComputeBackendBucketIamMember {
	return &ComputeBackendBucketIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeBackendBucketIamMember)(nil)

// ComputeBackendBucketIamMember represents the Terraform resource google_compute_backend_bucket_iam_member.
type ComputeBackendBucketIamMember struct {
	Name      string
	Args      ComputeBackendBucketIamMemberArgs
	state     *computeBackendBucketIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeBackendBucketIamMember].
func (cbbim *ComputeBackendBucketIamMember) Type() string {
	return "google_compute_backend_bucket_iam_member"
}

// LocalName returns the local name for [ComputeBackendBucketIamMember].
func (cbbim *ComputeBackendBucketIamMember) LocalName() string {
	return cbbim.Name
}

// Configuration returns the configuration (args) for [ComputeBackendBucketIamMember].
func (cbbim *ComputeBackendBucketIamMember) Configuration() interface{} {
	return cbbim.Args
}

// DependOn is used for other resources to depend on [ComputeBackendBucketIamMember].
func (cbbim *ComputeBackendBucketIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(cbbim)
}

// Dependencies returns the list of resources [ComputeBackendBucketIamMember] depends_on.
func (cbbim *ComputeBackendBucketIamMember) Dependencies() terra.Dependencies {
	return cbbim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeBackendBucketIamMember].
func (cbbim *ComputeBackendBucketIamMember) LifecycleManagement() *terra.Lifecycle {
	return cbbim.Lifecycle
}

// Attributes returns the attributes for [ComputeBackendBucketIamMember].
func (cbbim *ComputeBackendBucketIamMember) Attributes() computeBackendBucketIamMemberAttributes {
	return computeBackendBucketIamMemberAttributes{ref: terra.ReferenceResource(cbbim)}
}

// ImportState imports the given attribute values into [ComputeBackendBucketIamMember]'s state.
func (cbbim *ComputeBackendBucketIamMember) ImportState(av io.Reader) error {
	cbbim.state = &computeBackendBucketIamMemberState{}
	if err := json.NewDecoder(av).Decode(cbbim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbbim.Type(), cbbim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeBackendBucketIamMember] has state.
func (cbbim *ComputeBackendBucketIamMember) State() (*computeBackendBucketIamMemberState, bool) {
	return cbbim.state, cbbim.state != nil
}

// StateMust returns the state for [ComputeBackendBucketIamMember]. Panics if the state is nil.
func (cbbim *ComputeBackendBucketIamMember) StateMust() *computeBackendBucketIamMemberState {
	if cbbim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbbim.Type(), cbbim.LocalName()))
	}
	return cbbim.state
}

// ComputeBackendBucketIamMemberArgs contains the configurations for google_compute_backend_bucket_iam_member.
type ComputeBackendBucketIamMemberArgs struct {
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
	Condition *computebackendbucketiammember.Condition `hcl:"condition,block"`
}
type computeBackendBucketIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_backend_bucket_iam_member.
func (cbbim computeBackendBucketIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cbbim.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_backend_bucket_iam_member.
func (cbbim computeBackendBucketIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbbim.ref.Append("id"))
}

// Member returns a reference to field member of google_compute_backend_bucket_iam_member.
func (cbbim computeBackendBucketIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(cbbim.ref.Append("member"))
}

// Name returns a reference to field name of google_compute_backend_bucket_iam_member.
func (cbbim computeBackendBucketIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbbim.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_backend_bucket_iam_member.
func (cbbim computeBackendBucketIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbbim.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_backend_bucket_iam_member.
func (cbbim computeBackendBucketIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cbbim.ref.Append("role"))
}

func (cbbim computeBackendBucketIamMemberAttributes) Condition() terra.ListValue[computebackendbucketiammember.ConditionAttributes] {
	return terra.ReferenceAsList[computebackendbucketiammember.ConditionAttributes](cbbim.ref.Append("condition"))
}

type computeBackendBucketIamMemberState struct {
	Etag      string                                         `json:"etag"`
	Id        string                                         `json:"id"`
	Member    string                                         `json:"member"`
	Name      string                                         `json:"name"`
	Project   string                                         `json:"project"`
	Role      string                                         `json:"role"`
	Condition []computebackendbucketiammember.ConditionState `json:"condition"`
}
