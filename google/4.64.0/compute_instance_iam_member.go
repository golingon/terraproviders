// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeinstanceiammember "github.com/golingon/terraproviders/google/4.64.0/computeinstanceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeInstanceIamMember creates a new instance of [ComputeInstanceIamMember].
func NewComputeInstanceIamMember(name string, args ComputeInstanceIamMemberArgs) *ComputeInstanceIamMember {
	return &ComputeInstanceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInstanceIamMember)(nil)

// ComputeInstanceIamMember represents the Terraform resource google_compute_instance_iam_member.
type ComputeInstanceIamMember struct {
	Name      string
	Args      ComputeInstanceIamMemberArgs
	state     *computeInstanceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeInstanceIamMember].
func (ciim *ComputeInstanceIamMember) Type() string {
	return "google_compute_instance_iam_member"
}

// LocalName returns the local name for [ComputeInstanceIamMember].
func (ciim *ComputeInstanceIamMember) LocalName() string {
	return ciim.Name
}

// Configuration returns the configuration (args) for [ComputeInstanceIamMember].
func (ciim *ComputeInstanceIamMember) Configuration() interface{} {
	return ciim.Args
}

// DependOn is used for other resources to depend on [ComputeInstanceIamMember].
func (ciim *ComputeInstanceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ciim)
}

// Dependencies returns the list of resources [ComputeInstanceIamMember] depends_on.
func (ciim *ComputeInstanceIamMember) Dependencies() terra.Dependencies {
	return ciim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeInstanceIamMember].
func (ciim *ComputeInstanceIamMember) LifecycleManagement() *terra.Lifecycle {
	return ciim.Lifecycle
}

// Attributes returns the attributes for [ComputeInstanceIamMember].
func (ciim *ComputeInstanceIamMember) Attributes() computeInstanceIamMemberAttributes {
	return computeInstanceIamMemberAttributes{ref: terra.ReferenceResource(ciim)}
}

// ImportState imports the given attribute values into [ComputeInstanceIamMember]'s state.
func (ciim *ComputeInstanceIamMember) ImportState(av io.Reader) error {
	ciim.state = &computeInstanceIamMemberState{}
	if err := json.NewDecoder(av).Decode(ciim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ciim.Type(), ciim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeInstanceIamMember] has state.
func (ciim *ComputeInstanceIamMember) State() (*computeInstanceIamMemberState, bool) {
	return ciim.state, ciim.state != nil
}

// StateMust returns the state for [ComputeInstanceIamMember]. Panics if the state is nil.
func (ciim *ComputeInstanceIamMember) StateMust() *computeInstanceIamMemberState {
	if ciim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ciim.Type(), ciim.LocalName()))
	}
	return ciim.state
}

// ComputeInstanceIamMemberArgs contains the configurations for google_compute_instance_iam_member.
type ComputeInstanceIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Condition: optional
	Condition *computeinstanceiammember.Condition `hcl:"condition,block"`
}
type computeInstanceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_instance_iam_member.
func (ciim computeInstanceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_instance_iam_member.
func (ciim computeInstanceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of google_compute_instance_iam_member.
func (ciim computeInstanceIamMemberAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("instance_name"))
}

// Member returns a reference to field member of google_compute_instance_iam_member.
func (ciim computeInstanceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("member"))
}

// Project returns a reference to field project of google_compute_instance_iam_member.
func (ciim computeInstanceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_instance_iam_member.
func (ciim computeInstanceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("role"))
}

// Zone returns a reference to field zone of google_compute_instance_iam_member.
func (ciim computeInstanceIamMemberAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(ciim.ref.Append("zone"))
}

func (ciim computeInstanceIamMemberAttributes) Condition() terra.ListValue[computeinstanceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[computeinstanceiammember.ConditionAttributes](ciim.ref.Append("condition"))
}

type computeInstanceIamMemberState struct {
	Etag         string                                    `json:"etag"`
	Id           string                                    `json:"id"`
	InstanceName string                                    `json:"instance_name"`
	Member       string                                    `json:"member"`
	Project      string                                    `json:"project"`
	Role         string                                    `json:"role"`
	Zone         string                                    `json:"zone"`
	Condition    []computeinstanceiammember.ConditionState `json:"condition"`
}
