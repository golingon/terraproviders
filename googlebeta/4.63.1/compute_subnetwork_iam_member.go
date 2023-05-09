// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computesubnetworkiammember "github.com/golingon/terraproviders/googlebeta/4.63.1/computesubnetworkiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeSubnetworkIamMember creates a new instance of [ComputeSubnetworkIamMember].
func NewComputeSubnetworkIamMember(name string, args ComputeSubnetworkIamMemberArgs) *ComputeSubnetworkIamMember {
	return &ComputeSubnetworkIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSubnetworkIamMember)(nil)

// ComputeSubnetworkIamMember represents the Terraform resource google_compute_subnetwork_iam_member.
type ComputeSubnetworkIamMember struct {
	Name      string
	Args      ComputeSubnetworkIamMemberArgs
	state     *computeSubnetworkIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeSubnetworkIamMember].
func (csim *ComputeSubnetworkIamMember) Type() string {
	return "google_compute_subnetwork_iam_member"
}

// LocalName returns the local name for [ComputeSubnetworkIamMember].
func (csim *ComputeSubnetworkIamMember) LocalName() string {
	return csim.Name
}

// Configuration returns the configuration (args) for [ComputeSubnetworkIamMember].
func (csim *ComputeSubnetworkIamMember) Configuration() interface{} {
	return csim.Args
}

// DependOn is used for other resources to depend on [ComputeSubnetworkIamMember].
func (csim *ComputeSubnetworkIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(csim)
}

// Dependencies returns the list of resources [ComputeSubnetworkIamMember] depends_on.
func (csim *ComputeSubnetworkIamMember) Dependencies() terra.Dependencies {
	return csim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeSubnetworkIamMember].
func (csim *ComputeSubnetworkIamMember) LifecycleManagement() *terra.Lifecycle {
	return csim.Lifecycle
}

// Attributes returns the attributes for [ComputeSubnetworkIamMember].
func (csim *ComputeSubnetworkIamMember) Attributes() computeSubnetworkIamMemberAttributes {
	return computeSubnetworkIamMemberAttributes{ref: terra.ReferenceResource(csim)}
}

// ImportState imports the given attribute values into [ComputeSubnetworkIamMember]'s state.
func (csim *ComputeSubnetworkIamMember) ImportState(av io.Reader) error {
	csim.state = &computeSubnetworkIamMemberState{}
	if err := json.NewDecoder(av).Decode(csim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csim.Type(), csim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeSubnetworkIamMember] has state.
func (csim *ComputeSubnetworkIamMember) State() (*computeSubnetworkIamMemberState, bool) {
	return csim.state, csim.state != nil
}

// StateMust returns the state for [ComputeSubnetworkIamMember]. Panics if the state is nil.
func (csim *ComputeSubnetworkIamMember) StateMust() *computeSubnetworkIamMemberState {
	if csim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csim.Type(), csim.LocalName()))
	}
	return csim.state
}

// ComputeSubnetworkIamMemberArgs contains the configurations for google_compute_subnetwork_iam_member.
type ComputeSubnetworkIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Subnetwork: string, required
	Subnetwork terra.StringValue `hcl:"subnetwork,attr" validate:"required"`
	// Condition: optional
	Condition *computesubnetworkiammember.Condition `hcl:"condition,block"`
}
type computeSubnetworkIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_subnetwork_iam_member.
func (csim computeSubnetworkIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_subnetwork_iam_member.
func (csim computeSubnetworkIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("id"))
}

// Member returns a reference to field member of google_compute_subnetwork_iam_member.
func (csim computeSubnetworkIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("member"))
}

// Project returns a reference to field project of google_compute_subnetwork_iam_member.
func (csim computeSubnetworkIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_subnetwork_iam_member.
func (csim computeSubnetworkIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("region"))
}

// Role returns a reference to field role of google_compute_subnetwork_iam_member.
func (csim computeSubnetworkIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("role"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_subnetwork_iam_member.
func (csim computeSubnetworkIamMemberAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(csim.ref.Append("subnetwork"))
}

func (csim computeSubnetworkIamMemberAttributes) Condition() terra.ListValue[computesubnetworkiammember.ConditionAttributes] {
	return terra.ReferenceAsList[computesubnetworkiammember.ConditionAttributes](csim.ref.Append("condition"))
}

type computeSubnetworkIamMemberState struct {
	Etag       string                                      `json:"etag"`
	Id         string                                      `json:"id"`
	Member     string                                      `json:"member"`
	Project    string                                      `json:"project"`
	Region     string                                      `json:"region"`
	Role       string                                      `json:"role"`
	Subnetwork string                                      `json:"subnetwork"`
	Condition  []computesubnetworkiammember.ConditionState `json:"condition"`
}
