// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIamGroupPolicyAttachment creates a new instance of [IamGroupPolicyAttachment].
func NewIamGroupPolicyAttachment(name string, args IamGroupPolicyAttachmentArgs) *IamGroupPolicyAttachment {
	return &IamGroupPolicyAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamGroupPolicyAttachment)(nil)

// IamGroupPolicyAttachment represents the Terraform resource aws_iam_group_policy_attachment.
type IamGroupPolicyAttachment struct {
	Name      string
	Args      IamGroupPolicyAttachmentArgs
	state     *iamGroupPolicyAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamGroupPolicyAttachment].
func (igpa *IamGroupPolicyAttachment) Type() string {
	return "aws_iam_group_policy_attachment"
}

// LocalName returns the local name for [IamGroupPolicyAttachment].
func (igpa *IamGroupPolicyAttachment) LocalName() string {
	return igpa.Name
}

// Configuration returns the configuration (args) for [IamGroupPolicyAttachment].
func (igpa *IamGroupPolicyAttachment) Configuration() interface{} {
	return igpa.Args
}

// DependOn is used for other resources to depend on [IamGroupPolicyAttachment].
func (igpa *IamGroupPolicyAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(igpa)
}

// Dependencies returns the list of resources [IamGroupPolicyAttachment] depends_on.
func (igpa *IamGroupPolicyAttachment) Dependencies() terra.Dependencies {
	return igpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamGroupPolicyAttachment].
func (igpa *IamGroupPolicyAttachment) LifecycleManagement() *terra.Lifecycle {
	return igpa.Lifecycle
}

// Attributes returns the attributes for [IamGroupPolicyAttachment].
func (igpa *IamGroupPolicyAttachment) Attributes() iamGroupPolicyAttachmentAttributes {
	return iamGroupPolicyAttachmentAttributes{ref: terra.ReferenceResource(igpa)}
}

// ImportState imports the given attribute values into [IamGroupPolicyAttachment]'s state.
func (igpa *IamGroupPolicyAttachment) ImportState(av io.Reader) error {
	igpa.state = &iamGroupPolicyAttachmentState{}
	if err := json.NewDecoder(av).Decode(igpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", igpa.Type(), igpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamGroupPolicyAttachment] has state.
func (igpa *IamGroupPolicyAttachment) State() (*iamGroupPolicyAttachmentState, bool) {
	return igpa.state, igpa.state != nil
}

// StateMust returns the state for [IamGroupPolicyAttachment]. Panics if the state is nil.
func (igpa *IamGroupPolicyAttachment) StateMust() *iamGroupPolicyAttachmentState {
	if igpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", igpa.Type(), igpa.LocalName()))
	}
	return igpa.state
}

// IamGroupPolicyAttachmentArgs contains the configurations for aws_iam_group_policy_attachment.
type IamGroupPolicyAttachmentArgs struct {
	// Group: string, required
	Group terra.StringValue `hcl:"group,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyArn: string, required
	PolicyArn terra.StringValue `hcl:"policy_arn,attr" validate:"required"`
}
type iamGroupPolicyAttachmentAttributes struct {
	ref terra.Reference
}

// Group returns a reference to field group of aws_iam_group_policy_attachment.
func (igpa iamGroupPolicyAttachmentAttributes) Group() terra.StringValue {
	return terra.ReferenceAsString(igpa.ref.Append("group"))
}

// Id returns a reference to field id of aws_iam_group_policy_attachment.
func (igpa iamGroupPolicyAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(igpa.ref.Append("id"))
}

// PolicyArn returns a reference to field policy_arn of aws_iam_group_policy_attachment.
func (igpa iamGroupPolicyAttachmentAttributes) PolicyArn() terra.StringValue {
	return terra.ReferenceAsString(igpa.ref.Append("policy_arn"))
}

type iamGroupPolicyAttachmentState struct {
	Group     string `json:"group"`
	Id        string `json:"id"`
	PolicyArn string `json:"policy_arn"`
}
