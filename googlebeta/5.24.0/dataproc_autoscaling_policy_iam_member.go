// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	dataprocautoscalingpolicyiammember "github.com/golingon/terraproviders/googlebeta/5.24.0/dataprocautoscalingpolicyiammember"
	"io"
)

// NewDataprocAutoscalingPolicyIamMember creates a new instance of [DataprocAutoscalingPolicyIamMember].
func NewDataprocAutoscalingPolicyIamMember(name string, args DataprocAutoscalingPolicyIamMemberArgs) *DataprocAutoscalingPolicyIamMember {
	return &DataprocAutoscalingPolicyIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocAutoscalingPolicyIamMember)(nil)

// DataprocAutoscalingPolicyIamMember represents the Terraform resource google_dataproc_autoscaling_policy_iam_member.
type DataprocAutoscalingPolicyIamMember struct {
	Name      string
	Args      DataprocAutoscalingPolicyIamMemberArgs
	state     *dataprocAutoscalingPolicyIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocAutoscalingPolicyIamMember].
func (dapim *DataprocAutoscalingPolicyIamMember) Type() string {
	return "google_dataproc_autoscaling_policy_iam_member"
}

// LocalName returns the local name for [DataprocAutoscalingPolicyIamMember].
func (dapim *DataprocAutoscalingPolicyIamMember) LocalName() string {
	return dapim.Name
}

// Configuration returns the configuration (args) for [DataprocAutoscalingPolicyIamMember].
func (dapim *DataprocAutoscalingPolicyIamMember) Configuration() interface{} {
	return dapim.Args
}

// DependOn is used for other resources to depend on [DataprocAutoscalingPolicyIamMember].
func (dapim *DataprocAutoscalingPolicyIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dapim)
}

// Dependencies returns the list of resources [DataprocAutoscalingPolicyIamMember] depends_on.
func (dapim *DataprocAutoscalingPolicyIamMember) Dependencies() terra.Dependencies {
	return dapim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocAutoscalingPolicyIamMember].
func (dapim *DataprocAutoscalingPolicyIamMember) LifecycleManagement() *terra.Lifecycle {
	return dapim.Lifecycle
}

// Attributes returns the attributes for [DataprocAutoscalingPolicyIamMember].
func (dapim *DataprocAutoscalingPolicyIamMember) Attributes() dataprocAutoscalingPolicyIamMemberAttributes {
	return dataprocAutoscalingPolicyIamMemberAttributes{ref: terra.ReferenceResource(dapim)}
}

// ImportState imports the given attribute values into [DataprocAutoscalingPolicyIamMember]'s state.
func (dapim *DataprocAutoscalingPolicyIamMember) ImportState(av io.Reader) error {
	dapim.state = &dataprocAutoscalingPolicyIamMemberState{}
	if err := json.NewDecoder(av).Decode(dapim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dapim.Type(), dapim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocAutoscalingPolicyIamMember] has state.
func (dapim *DataprocAutoscalingPolicyIamMember) State() (*dataprocAutoscalingPolicyIamMemberState, bool) {
	return dapim.state, dapim.state != nil
}

// StateMust returns the state for [DataprocAutoscalingPolicyIamMember]. Panics if the state is nil.
func (dapim *DataprocAutoscalingPolicyIamMember) StateMust() *dataprocAutoscalingPolicyIamMemberState {
	if dapim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dapim.Type(), dapim.LocalName()))
	}
	return dapim.state
}

// DataprocAutoscalingPolicyIamMemberArgs contains the configurations for google_dataproc_autoscaling_policy_iam_member.
type DataprocAutoscalingPolicyIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// PolicyId: string, required
	PolicyId terra.StringValue `hcl:"policy_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *dataprocautoscalingpolicyiammember.Condition `hcl:"condition,block"`
}
type dataprocAutoscalingPolicyIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_autoscaling_policy_iam_member.
func (dapim dataprocAutoscalingPolicyIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dapim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_autoscaling_policy_iam_member.
func (dapim dataprocAutoscalingPolicyIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dapim.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_autoscaling_policy_iam_member.
func (dapim dataprocAutoscalingPolicyIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dapim.ref.Append("location"))
}

// Member returns a reference to field member of google_dataproc_autoscaling_policy_iam_member.
func (dapim dataprocAutoscalingPolicyIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dapim.ref.Append("member"))
}

// PolicyId returns a reference to field policy_id of google_dataproc_autoscaling_policy_iam_member.
func (dapim dataprocAutoscalingPolicyIamMemberAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(dapim.ref.Append("policy_id"))
}

// Project returns a reference to field project of google_dataproc_autoscaling_policy_iam_member.
func (dapim dataprocAutoscalingPolicyIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dapim.ref.Append("project"))
}

// Role returns a reference to field role of google_dataproc_autoscaling_policy_iam_member.
func (dapim dataprocAutoscalingPolicyIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dapim.ref.Append("role"))
}

func (dapim dataprocAutoscalingPolicyIamMemberAttributes) Condition() terra.ListValue[dataprocautoscalingpolicyiammember.ConditionAttributes] {
	return terra.ReferenceAsList[dataprocautoscalingpolicyiammember.ConditionAttributes](dapim.ref.Append("condition"))
}

type dataprocAutoscalingPolicyIamMemberState struct {
	Etag      string                                              `json:"etag"`
	Id        string                                              `json:"id"`
	Location  string                                              `json:"location"`
	Member    string                                              `json:"member"`
	PolicyId  string                                              `json:"policy_id"`
	Project   string                                              `json:"project"`
	Role      string                                              `json:"role"`
	Condition []dataprocautoscalingpolicyiammember.ConditionState `json:"condition"`
}
