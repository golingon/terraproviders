// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_iam_group_policy_attachment

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

// Resource represents the Terraform resource aws_iam_group_policy_attachment.
type Resource struct {
	Name      string
	Args      Args
	state     *awsIamGroupPolicyAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aigpa *Resource) Type() string {
	return "aws_iam_group_policy_attachment"
}

// LocalName returns the local name for [Resource].
func (aigpa *Resource) LocalName() string {
	return aigpa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aigpa *Resource) Configuration() interface{} {
	return aigpa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aigpa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aigpa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aigpa *Resource) Dependencies() terra.Dependencies {
	return aigpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aigpa *Resource) LifecycleManagement() *terra.Lifecycle {
	return aigpa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aigpa *Resource) Attributes() awsIamGroupPolicyAttachmentAttributes {
	return awsIamGroupPolicyAttachmentAttributes{ref: terra.ReferenceResource(aigpa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aigpa *Resource) ImportState(state io.Reader) error {
	aigpa.state = &awsIamGroupPolicyAttachmentState{}
	if err := json.NewDecoder(state).Decode(aigpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aigpa.Type(), aigpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aigpa *Resource) State() (*awsIamGroupPolicyAttachmentState, bool) {
	return aigpa.state, aigpa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aigpa *Resource) StateMust() *awsIamGroupPolicyAttachmentState {
	if aigpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aigpa.Type(), aigpa.LocalName()))
	}
	return aigpa.state
}

// Args contains the configurations for aws_iam_group_policy_attachment.
type Args struct {
	// Group: string, required
	Group terra.StringValue `hcl:"group,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyArn: string, required
	PolicyArn terra.StringValue `hcl:"policy_arn,attr" validate:"required"`
}

type awsIamGroupPolicyAttachmentAttributes struct {
	ref terra.Reference
}

// Group returns a reference to field group of aws_iam_group_policy_attachment.
func (aigpa awsIamGroupPolicyAttachmentAttributes) Group() terra.StringValue {
	return terra.ReferenceAsString(aigpa.ref.Append("group"))
}

// Id returns a reference to field id of aws_iam_group_policy_attachment.
func (aigpa awsIamGroupPolicyAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aigpa.ref.Append("id"))
}

// PolicyArn returns a reference to field policy_arn of aws_iam_group_policy_attachment.
func (aigpa awsIamGroupPolicyAttachmentAttributes) PolicyArn() terra.StringValue {
	return terra.ReferenceAsString(aigpa.ref.Append("policy_arn"))
}

type awsIamGroupPolicyAttachmentState struct {
	Group     string `json:"group"`
	Id        string `json:"id"`
	PolicyArn string `json:"policy_arn"`
}
