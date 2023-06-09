// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeimageiammember "github.com/golingon/terraproviders/google/4.64.0/computeimageiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeImageIamMember creates a new instance of [ComputeImageIamMember].
func NewComputeImageIamMember(name string, args ComputeImageIamMemberArgs) *ComputeImageIamMember {
	return &ComputeImageIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeImageIamMember)(nil)

// ComputeImageIamMember represents the Terraform resource google_compute_image_iam_member.
type ComputeImageIamMember struct {
	Name      string
	Args      ComputeImageIamMemberArgs
	state     *computeImageIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeImageIamMember].
func (ciim *ComputeImageIamMember) Type() string {
	return "google_compute_image_iam_member"
}

// LocalName returns the local name for [ComputeImageIamMember].
func (ciim *ComputeImageIamMember) LocalName() string {
	return ciim.Name
}

// Configuration returns the configuration (args) for [ComputeImageIamMember].
func (ciim *ComputeImageIamMember) Configuration() interface{} {
	return ciim.Args
}

// DependOn is used for other resources to depend on [ComputeImageIamMember].
func (ciim *ComputeImageIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ciim)
}

// Dependencies returns the list of resources [ComputeImageIamMember] depends_on.
func (ciim *ComputeImageIamMember) Dependencies() terra.Dependencies {
	return ciim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeImageIamMember].
func (ciim *ComputeImageIamMember) LifecycleManagement() *terra.Lifecycle {
	return ciim.Lifecycle
}

// Attributes returns the attributes for [ComputeImageIamMember].
func (ciim *ComputeImageIamMember) Attributes() computeImageIamMemberAttributes {
	return computeImageIamMemberAttributes{ref: terra.ReferenceResource(ciim)}
}

// ImportState imports the given attribute values into [ComputeImageIamMember]'s state.
func (ciim *ComputeImageIamMember) ImportState(av io.Reader) error {
	ciim.state = &computeImageIamMemberState{}
	if err := json.NewDecoder(av).Decode(ciim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ciim.Type(), ciim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeImageIamMember] has state.
func (ciim *ComputeImageIamMember) State() (*computeImageIamMemberState, bool) {
	return ciim.state, ciim.state != nil
}

// StateMust returns the state for [ComputeImageIamMember]. Panics if the state is nil.
func (ciim *ComputeImageIamMember) StateMust() *computeImageIamMemberState {
	if ciim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ciim.Type(), ciim.LocalName()))
	}
	return ciim.state
}

// ComputeImageIamMemberArgs contains the configurations for google_compute_image_iam_member.
type ComputeImageIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Image: string, required
	Image terra.StringValue `hcl:"image,attr" validate:"required"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *computeimageiammember.Condition `hcl:"condition,block"`
}
type computeImageIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_image_iam_member.
func (ciim computeImageIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_image_iam_member.
func (ciim computeImageIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("id"))
}

// Image returns a reference to field image of google_compute_image_iam_member.
func (ciim computeImageIamMemberAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("image"))
}

// Member returns a reference to field member of google_compute_image_iam_member.
func (ciim computeImageIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("member"))
}

// Project returns a reference to field project of google_compute_image_iam_member.
func (ciim computeImageIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_image_iam_member.
func (ciim computeImageIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("role"))
}

func (ciim computeImageIamMemberAttributes) Condition() terra.ListValue[computeimageiammember.ConditionAttributes] {
	return terra.ReferenceAsList[computeimageiammember.ConditionAttributes](ciim.ref.Append("condition"))
}

type computeImageIamMemberState struct {
	Etag      string                                 `json:"etag"`
	Id        string                                 `json:"id"`
	Image     string                                 `json:"image"`
	Member    string                                 `json:"member"`
	Project   string                                 `json:"project"`
	Role      string                                 `json:"role"`
	Condition []computeimageiammember.ConditionState `json:"condition"`
}
