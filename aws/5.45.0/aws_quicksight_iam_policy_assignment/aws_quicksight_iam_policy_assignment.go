// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_quicksight_iam_policy_assignment

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

// Resource represents the Terraform resource aws_quicksight_iam_policy_assignment.
type Resource struct {
	Name      string
	Args      Args
	state     *awsQuicksightIamPolicyAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aqipa *Resource) Type() string {
	return "aws_quicksight_iam_policy_assignment"
}

// LocalName returns the local name for [Resource].
func (aqipa *Resource) LocalName() string {
	return aqipa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aqipa *Resource) Configuration() interface{} {
	return aqipa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aqipa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aqipa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aqipa *Resource) Dependencies() terra.Dependencies {
	return aqipa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aqipa *Resource) LifecycleManagement() *terra.Lifecycle {
	return aqipa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aqipa *Resource) Attributes() awsQuicksightIamPolicyAssignmentAttributes {
	return awsQuicksightIamPolicyAssignmentAttributes{ref: terra.ReferenceResource(aqipa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aqipa *Resource) ImportState(state io.Reader) error {
	aqipa.state = &awsQuicksightIamPolicyAssignmentState{}
	if err := json.NewDecoder(state).Decode(aqipa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aqipa.Type(), aqipa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aqipa *Resource) State() (*awsQuicksightIamPolicyAssignmentState, bool) {
	return aqipa.state, aqipa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aqipa *Resource) StateMust() *awsQuicksightIamPolicyAssignmentState {
	if aqipa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aqipa.Type(), aqipa.LocalName()))
	}
	return aqipa.state
}

// Args contains the configurations for aws_quicksight_iam_policy_assignment.
type Args struct {
	// AssignmentName: string, required
	AssignmentName terra.StringValue `hcl:"assignment_name,attr" validate:"required"`
	// AssignmentStatus: string, required
	AssignmentStatus terra.StringValue `hcl:"assignment_status,attr" validate:"required"`
	// AwsAccountId: string, optional
	AwsAccountId terra.StringValue `hcl:"aws_account_id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// PolicyArn: string, optional
	PolicyArn terra.StringValue `hcl:"policy_arn,attr"`
	// Identities: min=0
	Identities []Identities `hcl:"identities,block" validate:"min=0"`
}

type awsQuicksightIamPolicyAssignmentAttributes struct {
	ref terra.Reference
}

// AssignmentId returns a reference to field assignment_id of aws_quicksight_iam_policy_assignment.
func (aqipa awsQuicksightIamPolicyAssignmentAttributes) AssignmentId() terra.StringValue {
	return terra.ReferenceAsString(aqipa.ref.Append("assignment_id"))
}

// AssignmentName returns a reference to field assignment_name of aws_quicksight_iam_policy_assignment.
func (aqipa awsQuicksightIamPolicyAssignmentAttributes) AssignmentName() terra.StringValue {
	return terra.ReferenceAsString(aqipa.ref.Append("assignment_name"))
}

// AssignmentStatus returns a reference to field assignment_status of aws_quicksight_iam_policy_assignment.
func (aqipa awsQuicksightIamPolicyAssignmentAttributes) AssignmentStatus() terra.StringValue {
	return terra.ReferenceAsString(aqipa.ref.Append("assignment_status"))
}

// AwsAccountId returns a reference to field aws_account_id of aws_quicksight_iam_policy_assignment.
func (aqipa awsQuicksightIamPolicyAssignmentAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(aqipa.ref.Append("aws_account_id"))
}

// Id returns a reference to field id of aws_quicksight_iam_policy_assignment.
func (aqipa awsQuicksightIamPolicyAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aqipa.ref.Append("id"))
}

// Namespace returns a reference to field namespace of aws_quicksight_iam_policy_assignment.
func (aqipa awsQuicksightIamPolicyAssignmentAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(aqipa.ref.Append("namespace"))
}

// PolicyArn returns a reference to field policy_arn of aws_quicksight_iam_policy_assignment.
func (aqipa awsQuicksightIamPolicyAssignmentAttributes) PolicyArn() terra.StringValue {
	return terra.ReferenceAsString(aqipa.ref.Append("policy_arn"))
}

func (aqipa awsQuicksightIamPolicyAssignmentAttributes) Identities() terra.ListValue[IdentitiesAttributes] {
	return terra.ReferenceAsList[IdentitiesAttributes](aqipa.ref.Append("identities"))
}

type awsQuicksightIamPolicyAssignmentState struct {
	AssignmentId     string            `json:"assignment_id"`
	AssignmentName   string            `json:"assignment_name"`
	AssignmentStatus string            `json:"assignment_status"`
	AwsAccountId     string            `json:"aws_account_id"`
	Id               string            `json:"id"`
	Namespace        string            `json:"namespace"`
	PolicyArn        string            `json:"policy_arn"`
	Identities       []IdentitiesState `json:"identities"`
}
