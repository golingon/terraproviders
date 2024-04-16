// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_iam_role_policy_attachment

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

// Resource represents the Terraform resource aws_iam_role_policy_attachment.
type Resource struct {
	Name      string
	Args      Args
	state     *awsIamRolePolicyAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (airpa *Resource) Type() string {
	return "aws_iam_role_policy_attachment"
}

// LocalName returns the local name for [Resource].
func (airpa *Resource) LocalName() string {
	return airpa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (airpa *Resource) Configuration() interface{} {
	return airpa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (airpa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(airpa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (airpa *Resource) Dependencies() terra.Dependencies {
	return airpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (airpa *Resource) LifecycleManagement() *terra.Lifecycle {
	return airpa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (airpa *Resource) Attributes() awsIamRolePolicyAttachmentAttributes {
	return awsIamRolePolicyAttachmentAttributes{ref: terra.ReferenceResource(airpa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (airpa *Resource) ImportState(state io.Reader) error {
	airpa.state = &awsIamRolePolicyAttachmentState{}
	if err := json.NewDecoder(state).Decode(airpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", airpa.Type(), airpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (airpa *Resource) State() (*awsIamRolePolicyAttachmentState, bool) {
	return airpa.state, airpa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (airpa *Resource) StateMust() *awsIamRolePolicyAttachmentState {
	if airpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", airpa.Type(), airpa.LocalName()))
	}
	return airpa.state
}

// Args contains the configurations for aws_iam_role_policy_attachment.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyArn: string, required
	PolicyArn terra.StringValue `hcl:"policy_arn,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
}

type awsIamRolePolicyAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_iam_role_policy_attachment.
func (airpa awsIamRolePolicyAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(airpa.ref.Append("id"))
}

// PolicyArn returns a reference to field policy_arn of aws_iam_role_policy_attachment.
func (airpa awsIamRolePolicyAttachmentAttributes) PolicyArn() terra.StringValue {
	return terra.ReferenceAsString(airpa.ref.Append("policy_arn"))
}

// Role returns a reference to field role of aws_iam_role_policy_attachment.
func (airpa awsIamRolePolicyAttachmentAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(airpa.ref.Append("role"))
}

type awsIamRolePolicyAttachmentState struct {
	Id        string `json:"id"`
	PolicyArn string `json:"policy_arn"`
	Role      string `json:"role"`
}
