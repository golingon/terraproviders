// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computemachineimageiammember "github.com/golingon/terraproviders/googlebeta/5.24.0/computemachineimageiammember"
	"io"
)

// NewComputeMachineImageIamMember creates a new instance of [ComputeMachineImageIamMember].
func NewComputeMachineImageIamMember(name string, args ComputeMachineImageIamMemberArgs) *ComputeMachineImageIamMember {
	return &ComputeMachineImageIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeMachineImageIamMember)(nil)

// ComputeMachineImageIamMember represents the Terraform resource google_compute_machine_image_iam_member.
type ComputeMachineImageIamMember struct {
	Name      string
	Args      ComputeMachineImageIamMemberArgs
	state     *computeMachineImageIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeMachineImageIamMember].
func (cmiim *ComputeMachineImageIamMember) Type() string {
	return "google_compute_machine_image_iam_member"
}

// LocalName returns the local name for [ComputeMachineImageIamMember].
func (cmiim *ComputeMachineImageIamMember) LocalName() string {
	return cmiim.Name
}

// Configuration returns the configuration (args) for [ComputeMachineImageIamMember].
func (cmiim *ComputeMachineImageIamMember) Configuration() interface{} {
	return cmiim.Args
}

// DependOn is used for other resources to depend on [ComputeMachineImageIamMember].
func (cmiim *ComputeMachineImageIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(cmiim)
}

// Dependencies returns the list of resources [ComputeMachineImageIamMember] depends_on.
func (cmiim *ComputeMachineImageIamMember) Dependencies() terra.Dependencies {
	return cmiim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeMachineImageIamMember].
func (cmiim *ComputeMachineImageIamMember) LifecycleManagement() *terra.Lifecycle {
	return cmiim.Lifecycle
}

// Attributes returns the attributes for [ComputeMachineImageIamMember].
func (cmiim *ComputeMachineImageIamMember) Attributes() computeMachineImageIamMemberAttributes {
	return computeMachineImageIamMemberAttributes{ref: terra.ReferenceResource(cmiim)}
}

// ImportState imports the given attribute values into [ComputeMachineImageIamMember]'s state.
func (cmiim *ComputeMachineImageIamMember) ImportState(av io.Reader) error {
	cmiim.state = &computeMachineImageIamMemberState{}
	if err := json.NewDecoder(av).Decode(cmiim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmiim.Type(), cmiim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeMachineImageIamMember] has state.
func (cmiim *ComputeMachineImageIamMember) State() (*computeMachineImageIamMemberState, bool) {
	return cmiim.state, cmiim.state != nil
}

// StateMust returns the state for [ComputeMachineImageIamMember]. Panics if the state is nil.
func (cmiim *ComputeMachineImageIamMember) StateMust() *computeMachineImageIamMemberState {
	if cmiim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmiim.Type(), cmiim.LocalName()))
	}
	return cmiim.state
}

// ComputeMachineImageIamMemberArgs contains the configurations for google_compute_machine_image_iam_member.
type ComputeMachineImageIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MachineImage: string, required
	MachineImage terra.StringValue `hcl:"machine_image,attr" validate:"required"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *computemachineimageiammember.Condition `hcl:"condition,block"`
}
type computeMachineImageIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_machine_image_iam_member.
func (cmiim computeMachineImageIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cmiim.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_machine_image_iam_member.
func (cmiim computeMachineImageIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmiim.ref.Append("id"))
}

// MachineImage returns a reference to field machine_image of google_compute_machine_image_iam_member.
func (cmiim computeMachineImageIamMemberAttributes) MachineImage() terra.StringValue {
	return terra.ReferenceAsString(cmiim.ref.Append("machine_image"))
}

// Member returns a reference to field member of google_compute_machine_image_iam_member.
func (cmiim computeMachineImageIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(cmiim.ref.Append("member"))
}

// Project returns a reference to field project of google_compute_machine_image_iam_member.
func (cmiim computeMachineImageIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cmiim.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_machine_image_iam_member.
func (cmiim computeMachineImageIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cmiim.ref.Append("role"))
}

func (cmiim computeMachineImageIamMemberAttributes) Condition() terra.ListValue[computemachineimageiammember.ConditionAttributes] {
	return terra.ReferenceAsList[computemachineimageiammember.ConditionAttributes](cmiim.ref.Append("condition"))
}

type computeMachineImageIamMemberState struct {
	Etag         string                                        `json:"etag"`
	Id           string                                        `json:"id"`
	MachineImage string                                        `json:"machine_image"`
	Member       string                                        `json:"member"`
	Project      string                                        `json:"project"`
	Role         string                                        `json:"role"`
	Condition    []computemachineimageiammember.ConditionState `json:"condition"`
}
